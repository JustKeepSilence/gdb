/*
creatTime: 2020/11/9 20:53
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package db

import (
	"gdb/cmap"
	"gdb/compare"
	"gdb/sqlite"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Gdb struct {
	RtDb       *leveldb.DB        // the pointer of realTime database
	HisDb      *leveldb.DB        // the pointer of history database
	InfoDb     *leveldb.DB        // the pointer of gdb info database
	RtDbFilter cmap.ConcurrentMap // key filter
	DbPath     string
}

const (
	TimeKey         = "currentTimeStamp"
	Ram             = "ram"
	WrittenItems    = "writtenItems"
	Speed           = "speed"
	initialUserName = "admin"
	initialUserInfo = `{"passWord": "admin@123", "roles": ["super_user"]}`
)

func (gdb *Gdb) InitialDb(flag int) error {
	switch flag {
	case 0:
		var err error
		gdb.RtDb, err = leveldb.OpenFile(gdb.DbPath+"\\realTimeData", &opt.Options{
			Filter:   filter.NewBloomFilter(10),
			Comparer: compare.GdbComparer,
		})
		if err != nil {
			return connectionDBError{"fail in opening  " + gdb.DbPath + "\\realTimeData: " + err.Error()}
		}
		if gdb.RtDb == nil {
			return connectionDBError{"fail in opening " + gdb.DbPath + "\\realTimeData: null db pointer"}
		}
		return nil
	case 1:
		var err error
		gdb.HisDb, err = leveldb.OpenFile(gdb.DbPath+"\\historicalData", &opt.Options{
			Filter:   filter.NewBloomFilter(10),
			Comparer: compare.GdbComparer,
		})
		if err != nil {
			return connectionDBError{"fail in opening " + gdb.DbPath + "\\historicalData: " + err.Error()}
		}
		if gdb.HisDb == nil {
			return connectionDBError{"fail in opening " + gdb.DbPath + "\\historicalData : null db pointer"}
		}
		return nil
	default:
		var err error
		gdb.RtDb, err = leveldb.OpenFile(gdb.DbPath+"\\realTimeData", &opt.Options{
			Filter:   filter.NewBloomFilter(10),
			Comparer: compare.GdbComparer,
		})
		if err != nil {
			return connectionDBError{"fail in opening " + gdb.DbPath + "\\realTimeData: " + err.Error()}
		}
		if gdb.RtDb == nil {
			return connectionDBError{"fail in opening " + gdb.DbPath + "\\realTimeData: null db pointer"}
		}
		gdb.HisDb, err = leveldb.OpenFile(gdb.DbPath+"\\historicalData", &opt.Options{
			Filter:   filter.NewBloomFilter(10),
			Comparer: compare.GdbComparer,
		})
		if err != nil {
			return connectionDBError{"fail in opening " + gdb.DbPath + "\\historicalData: " + err.Error()}
		}
		if gdb.HisDb == nil {
			return connectionDBError{"fail in opening " + gdb.DbPath + "\\historicalData: null db pointer"}
		}
		gdb.InfoDb, err = leveldb.OpenFile(gdb.DbPath+"\\InfoData", &opt.Options{
			Filter:   filter.NewBloomFilter(10),
			Comparer: compare.GdbComparer,
		})
		if err != nil {
			return connectionDBError{"fail in opening " + gdb.DbPath + "\\InfoData: " + err.Error()}
		}
		if gdb.InfoDb == nil {
			return connectionDBError{"fail in opening " + gdb.DbPath + "\\InfoData: null db pointer"}
		}
		_ = gdb.InfoDb.Put([]byte(initialUserName), []byte(initialUserInfo), nil) // initial user info
		gdb.RtDbFilter = cmap.New()
		// add all items in SQLite to bloom filter
		groups, _ := sqlite.Query("select groupName from group_cfg")
		for _, group := range groups {
			groupName := group["groupName"]
			items, _ := sqlite.Query("select itemName from '" + groupName + "'")
			for _, item := range items {
				itemName := item["itemName"]
				gdb.RtDbFilter.Set(itemName, struct{}{}) // add key to filter, don't lock
			}
		}
		return nil
	}
}

func NewGdb(dbPath string) *Gdb {
	return &Gdb{
		DbPath: dbPath,
	}
}
