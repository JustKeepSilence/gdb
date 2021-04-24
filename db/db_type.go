/*
creatTime: 2020/11/9 20:53
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package db

import (
	"github.com/JustKeepSilence/gdb/cmap"
	"github.com/JustKeepSilence/gdb/compare"
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
	initialUserInfo = `{"passWord": "685a6b21dc732a9702a96e6731811ec9", "roles": ["super_user"]}` // md5 of admin@123@seu
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
	if r, err := gdb.infoDb.Get([]byte(initialUserName), nil); err != nil {
		return err
	} else {
		if r == nil {
			_ = gdb.infoDb.Put([]byte(initialUserName), []byte(initialUserInfo), nil) // initial user info
		}
	}
	gdb.rtDbFilter = cmap.New()
	// add all items in SQLite to bloom filter
	groups, _ := query(gdb.ItemDbPath, "select groupName from group_cfg")
	for _, group := range groups {
		groupName := group["groupName"]
		items, _ := query(gdb.ItemDbPath, "select itemName from '"+groupName+"'")
		for _, item := range items {
			itemName := item["itemName"]
			gdb.rtDbFilter.Set(itemName, struct{}{}) // add key to filter, don't lock
		}
	}
	return nil

}

func (gdb *Gdb) initialSQLite() error {
	sqlCreateGroupCfgTable := `create table if not exists group_cfg (id integer not null primary key, groupName text UNIQUE)`                                                                                                                                                                           // add group_cfg table
	sqlAddCalc := `insert into group_cfg (groupName) values ('calc')`                                                                                                                                                                                                                                   // add calc group
	sqlAddCalcTable := `create table if not exists calc (id integer not null primary key, itemName text UNIQUE, description text)`                                                                                                                                                                      // add calc group
	sqlAddCalcCfgTable := `create table if not exists calc_cfg (id integer not null primary key, description text, expression text, status text default 'false', duration text default 10, errorMessage text default '', createTime text, updatedTime text default '')`                                 //  add calc cfg table
	sqlAddLogCfgTable := `create table if not exists log_cfg (id integer not null primary key, logType text default 'error', requestString text default '', requestMethod text default 'post', requestUrl text default '', logMessage text, insertTime  NUMERIC DEFAULT (datetime('now','localtime')))` // create log table
	// columns are id, logType,  requestString, requestMethod, requestUrl, logMessage, insertTime
	sqlAddUserCfgTable := `create table if not exists user_cfg (id integer not null primary key, userName text UNIQUE, role text)`
	_, _ = updateItem(gdb.ItemDbPath, sqlAddCalc)
	if err := updateItems(gdb.ItemDbPath, []string{sqlCreateGroupCfgTable, sqlAddCalcTable, sqlAddCalcCfgTable, sqlAddLogCfgTable, sqlAddUserCfgTable}...); err != nil {
		return err
	}
	if r, err := query(gdb.ItemDbPath, "select 1 from user_cfg where userName='admin' limit 1"); err != nil {
		return err
	} else {
		if len(r) == 0 {
			if _, err := updateItem(gdb.ItemDbPath, `insert into user_cfg (userName, role) values ('admin', 'super_user')`); err != nil {
				return err
			} else {
				return nil
			}
		} else {
			return nil
		}
	}
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
