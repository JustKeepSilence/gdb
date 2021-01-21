/*
creatTime: 2020/11/9 20:53
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package db

import (
	"gdb/cmap"
	"gdb/sqlite"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type LevelDb struct {
	RtDb       *leveldb.DB        // the pointer of realTime database
	HisDb      *leveldb.DB        // the pointer of history database
	InfoDb     *leveldb.DB        // the pointer of gdb info database
	RtDbFilter cmap.ConcurrentMap // key filter
}

const (
	TimeKey         = "currentTimeStamp"
	Ram             = "ram"
	WrittenItems    = "writtenItems"
	Speed           = "speed"
	initialUserName = "seu"
	initialUserInfo = `{"passWord": "seu@123", "roles": ["super_user"]}`
)

func (ldb *LevelDb) InitialDb(dbPath string, flag int) error {
	switch flag {
	case 0:
		var err error
		ldb.RtDb, err = leveldb.OpenFile(dbPath+"\\realTimeData", &opt.Options{
			Filter: filter.NewBloomFilter(10),
		})
		if err != nil {
			return connectionDBError{"fail in opening  " + dbPath + "\\realTimeData: " + err.Error()}
		}
		if ldb.RtDb == nil {
			return connectionDBError{"fail in opening " + dbPath + "\\realTimeData: null db pointer"}
		}
		return nil
	case 1:
		var err error
		ldb.HisDb, err = leveldb.OpenFile(dbPath+"\\historicalData", &opt.Options{
			Filter: filter.NewBloomFilter(10),
		})
		if err != nil {
			return connectionDBError{"fail in opening " + dbPath + "\\historicalData: " + err.Error()}
		}
		if ldb.HisDb == nil {
			return connectionDBError{"fail in opening " + dbPath + "\\historicalData : null db pointer"}
		}
		return nil
	default:
		var err error
		ldb.RtDb, err = leveldb.OpenFile(dbPath+"\\realTimeData", &opt.Options{
			Filter: filter.NewBloomFilter(10),
		})
		if err != nil {
			return connectionDBError{"fail in opening " + dbPath + "\\realTimeData: " + err.Error()}
		}
		if ldb.RtDb == nil {
			return connectionDBError{"fail in opening " + dbPath + "\\realTimeData: null db pointer"}
		}
		ldb.HisDb, err = leveldb.OpenFile(dbPath+"\\historicalData", &opt.Options{
			Filter: filter.NewBloomFilter(10),
		})
		if err != nil {
			return connectionDBError{"fail in opening " + dbPath + "\\historicalData: " + err.Error()}
		}
		if ldb.HisDb == nil {
			return connectionDBError{"fail in opening " + dbPath + "\\historicalData: null db pointer"}
		}
		ldb.InfoDb, err = leveldb.OpenFile(dbPath+"\\InfoData", &opt.Options{
			Filter: filter.NewBloomFilter(10),
		})
		if err != nil {
			return connectionDBError{"fail in opening " + dbPath + "\\InfoData: " + err.Error()}
		}
		if ldb.InfoDb == nil {
			return connectionDBError{"fail in opening " + dbPath + "\\InfoData: null db pointer"}
		}
		_ = ldb.InfoDb.Put([]byte(initialUserName), []byte(initialUserInfo), nil) // initial user info
		ldb.RtDbFilter = cmap.New()
		// add all items in SQLite to bloom filter
		groups, _ := sqlite.Query("select groupName from group_cfg")
		for _, group := range groups {
			groupName := group["groupName"]
			items, _ := sqlite.Query("select itemName from '" + groupName + "'")
			for _, item := range items {
				itemName := item["itemName"]
				ldb.RtDbFilter.Set(itemName, struct{}{}) // add key to filter, don't lock
			}
		}
		return nil
	}
}
