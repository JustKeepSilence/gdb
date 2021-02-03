/*
creatTime: 2020/11/9 20:53
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package db

import (
	"github.com/JustKeepSilence/gdb/cmap"
	"github.com/JustKeepSilence/gdb/compare"
	"github.com/JustKeepSilence/gdb/sqlite"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"os"
)

type Gdb struct {
	rtDb       *leveldb.DB        // the pointer of realTime database
	hisDb      *leveldb.DB        // the pointer of history database
	infoDb     *leveldb.DB        // the pointer of gdb info database
	rtDbFilter cmap.ConcurrentMap // key filter
	DbPath     string             // path of leveldb
	ItemDbPath string             // path of item of leveldb
}

const (
	TimeKey         = "currentTimeStamp"
	Ram             = "ram"
	WrittenItems    = "writtenItems"
	Speed           = "speed"
	initialUserName = "admin"
	initialUserInfo = `{"passWord": "admin@123", "roles": ["super_user"]}`
)

func (gdb *Gdb) initialDb() error {
	var err error
	gdb.rtDb, err = leveldb.OpenFile(gdb.DbPath+"\\realTimeData", &opt.Options{
		Filter:   filter.NewBloomFilter(10),
		Comparer: compare.GdbComparer,
	})
	if err != nil {
		return connectionDBError{"fail in opening " + gdb.DbPath + "\\realTimeData: " + err.Error()}
	}
	if gdb.rtDb == nil {
		return connectionDBError{"fail in opening " + gdb.DbPath + "\\realTimeData: null db pointer"}
	}
	gdb.hisDb, err = leveldb.OpenFile(gdb.DbPath+"\\historicalData", &opt.Options{
		Filter:   filter.NewBloomFilter(10),
		Comparer: compare.GdbComparer,
	})
	if err != nil {
		return connectionDBError{"fail in opening " + gdb.DbPath + "\\historicalData: " + err.Error()}
	}
	if gdb.hisDb == nil {
		return connectionDBError{"fail in opening " + gdb.DbPath + "\\historicalData: null db pointer"}
	}
	gdb.infoDb, err = leveldb.OpenFile(gdb.DbPath+"\\InfoData", &opt.Options{
		Filter:   filter.NewBloomFilter(10),
		Comparer: compare.GdbComparer,
	})
	if err != nil {
		return connectionDBError{"fail in opening " + gdb.DbPath + "\\InfoData: " + err.Error()}
	}
	if gdb.infoDb == nil {
		return connectionDBError{"fail in opening " + gdb.DbPath + "\\InfoData: null db pointer"}
	}
	_ = gdb.infoDb.Put([]byte(initialUserName), []byte(initialUserInfo), nil) // initial user info
	gdb.rtDbFilter = cmap.New()
	// add all items in SQLite to bloom filter
	groups, _ := sqlite.Query(gdb.ItemDbPath, "select groupName from group_cfg")
	for _, group := range groups {
		groupName := group["groupName"]
		items, _ := sqlite.Query(gdb.ItemDbPath, "select itemName from '"+groupName+"'")
		for _, item := range items {
			itemName := item["itemName"]
			gdb.rtDbFilter.Set(itemName, struct{}{}) // add key to filter, don't lock
		}
	}
	return nil

}

func (gdb *Gdb) initialSQLite() error {
	sqlCreateGroupCfg := `create table if not exists group_cfg (id integer not null primary key, groupName text UNIQUE)`
	if err := sqlite.UpdateItems(gdb.ItemDbPath, []string{sqlCreateGroupCfg}...); err != nil {
		return err
	}
	sqlAddCalc := `insert into group_cfg (groupName) values ('calc')` // add calc group
	_ = sqlite.UpdateItems(gdb.ItemDbPath, []string{sqlAddCalc}...)
	sqlAddCalcTable := `create table if not exists calc (id integer not null primary key, itemName text UNIQUE, description text)`                                                                                                                                      // add calc group
	sqlAddCalcCfgTable := `create table if not exists calc_cfg (id integer not null primary key, description text, expression text, status text default 'false', duration text default 10, errorMessage text default '', createTime text, updatedTime text default '')` //  add calc cfg table
	if err := sqlite.UpdateItems(gdb.ItemDbPath, []string{sqlAddCalcTable, sqlAddCalcCfgTable}...); err != nil {
		return err
	}
	return nil
}

func NewGdb(dbPath, itemDbPath string) (*Gdb, error) {
	// check whether the given path exist
	if !fileExist(dbPath) {
		if err := os.MkdirAll(dbPath, 0766); err != nil {
			return nil, err
		}
	}
	if !fileExist(itemDbPath) {
		if err := os.MkdirAll(itemDbPath, 0766); err != nil {
			return nil, err
		}
	}
	g := &Gdb{
		DbPath:     dbPath,
		ItemDbPath: itemDbPath + "/ldb.db",
	}
	if err := g.initialSQLite(); err != nil {
		return nil, err
	}
	if err := g.initialDb(); err != nil {
		return nil, err
	}
	return g, nil
}

func fileExist(path string) bool {
	_, err := os.Lstat(path)
	b := os.IsExist(err)
	return b
}
