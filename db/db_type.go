/*
creatTime: 2020/11/9
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/JustKeepSilence/gdb/cmap"
	"github.com/JustKeepSilence/gdb/compare"
	"github.com/JustKeepSilence/gdb/memap"
	"github.com/VictoriaMetrics/fastcache"
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	jsonIter "github.com/json-iterator/go"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Gdb define the struct of gdb, you should use NewGdb to get new instance of gdb
type Gdb struct {
	rtDb RtGdb // realTime database to store realTime data of gdb,you can implement RtGdb interface to specify your own realTime database to store realTime data, built-in is fast cache or redis
	// history database field
	hisDb map[string]map[string]*leveldb.DB
	// memory database field, used to cache historical data for a period of time
	intRmHisDb    memap.IntMemConcurrentMap //  memory history database to store temporary historical data of int
	floatRmHisDb  memap.FloatMemConcurrentMap
	stringRmHisDb memap.StringMemConcurrentMap
	boolRmHisDb   memap.BoolMemConcurrentMap
	// info database to record running information of gdb
	floatInfoDb  *leveldb.DB
	intInfoDb    *leveldb.DB
	stringInfoDb *leveldb.DB
	boolInfoDb   *leveldb.DB
	systemInfoDb *leveldb.DB        // store memory , cpu usage
	itemDb       *sql.DB            // item database to store item of gdb, can be sqlite or mysql
	rtDbFilter   cmap.ConcurrentMap // key filter of gdb
	delimiter    string
	dbPath       string           // path of leveldb
	dsn          string           // dsn to connect item database
	e            *casbin.Enforcer // route permission
	itemDbName   string
	driverName   string
	groupNames   []string
	sqlField
	rtTimeDuration  time.Duration // time durations to sync RealTime data in memory to other database or file system
	hisTimeDuration time.Duration // time durations to sync history data in memory to leveldb
	mu              sync.RWMutex
	syncStatus      bool // record whether is syncing history data, true ==> is syncing, false ==> no
	isReloading     bool // record whether db is reload, if is reloading, all write operation will fail
}

type Options struct {
	DriverName   string // driverName of itemDb, default is sqlite3
	Dsn          string // dsn to connect itemDb, default is "file:itemDb.db?_auth_user=seu&_auth_pass=admin@123&_auth_crypt=SHA1@_vacuum=1"
	UseInnerStop bool   // whether user inner stopService function, if true, we will use StopService to stopService, you can also define your own method by implement StopGdb interface,in your own method, you should call gdb.CloseGdb to ensure gdbService close successfully
	RtGdb               // realTime interface, default is FastCacheRt
}

type sqlField struct {
	key      string // primary key clause
	unique   string // default key clause
	dateTime string // dateTime clause
}

const (
	timeKey          = "currentTimeStamp"
	joiner           = "__"
	timeFormatString = "2006-01-02 15:04:05"
	ram              = "ram"
	cpu              = "cpu"
	speed            = "speed"
	roles            = "visitor, super_user, common_user"
	fileSize         = "fileSize"
	syncTime         = "syncTime"
	syncConsumeTime  = "syncConsumeTime"
)

var dataTypes = []string{"float32", "int32", "string", "bool"}

var json = jsonIter.ConfigCompatibleWithStandardLibrary // see: https://github.com/json-iterator/go

// initial history database of gdb
func (gdb *Gdb) initialDb(rtGdb RtGdb) error {
	var err error
	var delimiter string
	if runtime.GOOS == "windows" {
		delimiter = "\\"
	} else {
		delimiter = "/"
	}
	gdb.delimiter = delimiter
	// initial realTime database
	gdb.rtDb = rtGdb
	err = gdb.rtDb.Load() // load memory realTime database
	if err != nil {
		return err
	}
	// initial memory history database
	gdb.intRmHisDb = memap.NewInt()
	gdb.floatRmHisDb = memap.NewFloat()
	gdb.stringRmHisDb = memap.NewString()
	gdb.boolRmHisDb = memap.NewBool()
	// initial info database
	gdb.systemInfoDb, err = leveldb.OpenFile(gdb.dbPath+delimiter+"InfoData"+delimiter+"systemInfoData", &opt.Options{
		Filter:   filter.NewBloomFilter(10),
		Comparer: compare.GdbComparer,
	})
	if err != nil {
		return fmt.Errorf("fail in opening infoData database:" + err.Error())
	}
	gdb.floatInfoDb, err = leveldb.OpenFile(gdb.dbPath+delimiter+"InfoData"+delimiter+"floatInfoData", &opt.Options{
		Filter:   filter.NewBloomFilter(10),
		Comparer: compare.GdbComparer,
	})
	if err != nil {
		return fmt.Errorf("fail in opening infoData database:" + err.Error())
	}
	gdb.intInfoDb, err = leveldb.OpenFile(gdb.dbPath+delimiter+"InfoData"+delimiter+"intInfoData", &opt.Options{
		Filter:   filter.NewBloomFilter(10),
		Comparer: compare.GdbComparer,
	})
	if err != nil {
		return fmt.Errorf("fail in opening infoData database:" + err.Error())
	}
	gdb.stringInfoDb, err = leveldb.OpenFile(gdb.dbPath+delimiter+"InfoData"+delimiter+"stringInfoData", &opt.Options{
		Filter:   filter.NewBloomFilter(10),
		Comparer: compare.GdbComparer,
	})
	if err != nil {
		return fmt.Errorf("fail in opening infoData database:" + err.Error())
	}
	gdb.boolInfoDb, err = leveldb.OpenFile(gdb.dbPath+delimiter+"InfoData"+delimiter+"boolInfoData", &opt.Options{
		Filter:   filter.NewBloomFilter(10),
		Comparer: compare.GdbComparer,
	})
	if err != nil {
		return fmt.Errorf("fail in opening infoData database:" + err.Error())
	}
	gdb.rtDbFilter = cmap.New()
	// add all items in SQLite to bloom filter
	groups, _ := gdb.query("select groupName from group_cfg")
	groupNames := []string{}
	for _, group := range groups {
		groupName := group["groupName"]
		groupNames = append(groupNames, groupName)
		items, _ := gdb.query("select itemName, dataType from `" + groupName + "`")
		for _, item := range items {
			itemName := item["itemName"] + joiner + groupName // itemName in filter = itemName + "__" + groupName
			dataType := item["dataType"]
			gdb.rtDbFilter.Set(itemName, dataType) // add key to filter, don't lock
		}
	}
	gdb.groupNames = groupNames
	if err := gdb.generateDBPointer(groupNames); err != nil {
		return err
	}
	return nil
}

func (gdb *Gdb) generateDBPointer(groupNames []string) error {
	path, delimiter := gdb.dbPath, gdb.delimiter
	if gdb.hisDb == nil {
		h := map[string]map[string]*leveldb.DB{}
		for _, dataType := range dataTypes {
			t := map[string]*leveldb.DB{}
			for _, groupName := range groupNames {
				if db, err := leveldb.OpenFile(path+delimiter+"historicalData"+delimiter+dataType+delimiter+groupName, &opt.Options{Filter: filter.NewBloomFilter(10),
					Comparer: compare.GdbComparer}); err != nil {
					return err
				} else {
					t[groupName] = db
				}
			}
			h[dataType] = t
		}
		gdb.hisDb = h
	} else {
		for _, dataType := range dataTypes {
			for _, groupName := range groupNames {
				if db, err := leveldb.OpenFile(path+delimiter+"historicalData"+delimiter+dataType+delimiter+groupName, &opt.Options{Filter: filter.NewBloomFilter(10),
					Comparer: compare.GdbComparer}); err != nil {
					return err
				} else {
					gdb.hisDb[dataType][groupName] = db // add pointer
				}
			}
		}
	}
	return nil
}

// initial itemDb of gdb
func (gdb *Gdb) initialItemDb(clientMode bool, driverName, dsn string) error {
	if err := gdb.connectDataBase(driverName, dsn); err != nil {
		return err
	}
	// generate sql according to driverName
	switch driverName {
	case "sqlite3":
		gdb.key = " id integer not null primary key " // auto increment
		gdb.unique = " text "                         // unique type
		gdb.dateTime = " datetime default (datetime('now','localtime')) "
		break
	case "mysql":
		gdb.key = " id int primary key auto_increment"
		gdb.unique = " varchar(250) "
		gdb.dateTime = " TIMESTAMP DEFAULT CURRENT_TIMESTAMP "
		break
	default:
		// sql server
	}
	gdb.dsn = dsn
	sqlCreateGroupCfgTable := `create table if not exists group_cfg (` + gdb.key + `, groupName ` + gdb.unique + ` UNIQUE)` // add group_cfg table
	if clientMode {
		sqlAddCalc := `insert into group_cfg (groupName) values ('calc')`                                                                                                                                      // add calc group
		sqlAddCalcTable := `create table if not exists calc ( ` + gdb.key + `, itemName ` + gdb.unique + ` UNIQUE, dataType text ,description text)`                                                           // add calc group
		sqlAddCalcCfgTable := `create table if not exists calc_cfg ( ` + gdb.key + `, description text, expression text, status text , duration text , errorMessage text , createTime text, updatedTime text)` //  add calc cfg table
		sqlAddLogCfgTable := `create table if not exists log_cfg ( ` + gdb.key + `, logMessage text, requestUser text , level text, insertTime` + gdb.dateTime + `)`                                           // create log table
		// columns are id, logType,  requestString, requestMethod, requestUrl, logMessage, insertTime
		sqlAddUserCfgTable := `create table if not exists user_cfg ( ` + gdb.key + `, userName ` + gdb.unique + ` UNIQUE, passWord text, role text, token text)`
		sqlAddRouteTable := `create table if not exists route_cfg ( ` + gdb.key + `, userName ` + gdb.unique + ` UNIQUE, routeRoles text)` // routeRoles is ["p,userName,url,method"]
		var infoString string
		if driverName == "sqlite3" {
			infoString = "select name from sqlite_master where name='calc'"
		} else {
			infoString = "select table_name from information_schema.tables where table_schema='" + gdb.itemDbName + "' and TABLE_NAME='calc'"
		}
		if r, err := gdb.query(infoString); err != nil {
			return err
		} else {
			if len(r) == 0 {
				// no calc table
				if err := gdb.updateItems([]string{sqlCreateGroupCfgTable, sqlAddCalcTable, sqlAddCalcCfgTable, sqlAddLogCfgTable, sqlAddUserCfgTable, sqlAddCalc, sqlAddRouteTable}...); err != nil {
					return err
				}
			} else {
				// exist calc table
				if err := gdb.updateItems([]string{sqlCreateGroupCfgTable, sqlAddCalcTable, sqlAddCalcCfgTable, sqlAddLogCfgTable, sqlAddUserCfgTable, sqlAddRouteTable}...); err != nil {
					return err
				}
			}
		}
		if r, err := gdb.query("select 1 from user_cfg where userName='admin' limit 1"); err != nil {
			return err
		} else {
			_, _ = gdb.updateItem(`insert into route_cfg (userName, routeRoles) values ('admin', '["p,admin,all,POST"]')`)
			if len(r) == 0 {
				// passWord = md5(passWord + "@seu")
				if _, err := gdb.updateItem(`insert into user_cfg (userName, passWord ,role) values ('admin', '685a6b21dc732a9702a96e6731811ec9', 'super_user')`); err != nil {
					return err
				} else {
					return nil
				}
			} else {
				return nil
			}
		}
	} else {
		if _, err := gdb.updateItem(sqlCreateGroupCfgTable); err != nil {
			return err
		} else {
			return nil
		}
	}
}

// NewGdb used to create new instance of gdb
//
// dbPath: path of history data, if folder of path not exist, we will create this path
//
// dsn: dsn string to connect itemDb of gdb, at present we support sqlite3 and mysql.
// for sqlite3, dsn should like "./itemDb/item.db", for mysql, dsn should like "root:admin@123@tcp(192.168.0.166:3306)/itemDb"
//
// Note: for sqlite3, if itemDb given not exist, we will create it, but for mysql, you should create itemDb database in advance.
func NewGdb(dbPath string, rtTimeDuration, hisTimeDuration time.Duration, opt *Options) (*Gdb, error) {
	// check whether the given path exist
	if !fileExist(dbPath) {
		if err := os.MkdirAll(dbPath, 0766); err != nil {
			return nil, err
		}
	}
	var g *Gdb
	path, _ := filepath.Abs(dbPath) // get abs path of given path
	if runtime.GOOS == "windows" {
		// for windows, we need to use absolute path
		g = &Gdb{
			dbPath: path,
		}
	} else {
		g = &Gdb{
			dbPath: path,
		}
	}
	if err := g.initialItemDb(false, opt.DriverName, opt.Dsn); err != nil {
		return nil, err
	}
	if err := g.initialDb(opt.RtGdb); err != nil {
		return nil, err
	}
	g.rtTimeDuration = rtTimeDuration
	g.hisTimeDuration = hisTimeDuration
	sw := sync.WaitGroup{}
	if opt.UseInnerStop {
		sw.Add(3)
		go func() {
			defer sw.Done()
			_ = g.StopService()
		}()
	} else {
		sw.Add(2)
	}
	go func() {
		defer sw.Done()
		_ = g.syncRtData()
	}()
	go func() {
		defer sw.Done()
		_ = g.syncHisData()
	}()
	return g, nil
}

func fileExist(path string) bool {
	_, err := os.Lstat(path)
	b := os.IsExist(err)
	return b
}

// RtGdb realTime database interface, you can implement this interface to define your own way to store realTimeData
//
// we provide to ways to store realTimeData, that is fastCache and redis
type RtGdb interface {
	BatchWrite(keys, values [][]byte) error     // batch write key-values pairs to realTime database
	BatchFetch(keys [][]byte) ([][]byte, error) // batch get data according to the given keys from realTime database, if given key not exist you should set nil
	Sync() error                                // save realTime data to other db or file system
	Load(params ...interface{}) error           //  initial database and load realTime data to database from db or other file system
}

// FastCacheRt gdb uses fast cache internally to implement realTimeDatabase interface
type FastCacheRt struct {
	RealTimePath string
	cache        *fastcache.Cache
}

func (r *FastCacheRt) BatchWrite(keys, values [][]byte) error {
	g := errgroup.Group{}
	for i := 0; i < len(keys); i++ {
		index := i
		g.Go(func() error {
			r.cache.SetBig(keys[index], values[index])
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (r *FastCacheRt) BatchFetch(keys [][]byte) ([][]byte, error) {
	b := make([][]byte, len(keys))
	g := errgroup.Group{}
	for i := 0; i < len(keys); i++ {
		index := i
		g.Go(func() error {
			if ok := r.cache.Has(keys[index]); ok {
				b[index] = r.cache.GetBig(nil, keys[index])
			} else {
				b[index] = nil
			}
			return nil
		})
	}
	_ = g.Wait()
	return b, nil
}

func (r *FastCacheRt) Sync() error {
	if err := r.cache.SaveToFileConcurrent(r.RealTimePath, 4); err != nil {
		return err
	}
	return nil
}

func (r *FastCacheRt) Load(_ ...interface{}) error {
	r.cache = fastcache.LoadFromFileOrNew(r.RealTimePath, 100*1024*1024)
	return nil
}

// RedisRt implementation of using redis as a realTimeDatabase of gdb
type RedisRt struct {
	Ip       string
	Port     int
	PassWord string
	DbNum    int
	KeyName  string
	client   *redis.Client
}

func (r *RedisRt) BatchWrite(keys, values [][]byte) error {
	v := []string{}
	for i := 0; i < len(keys); i++ {
		v = append(v, string(keys[i]))
		v = append(v, string(values[i]))
	}
	result := r.client.HSet(context.Background(), r.KeyName, v)
	if _, err := result.Result(); err != nil {
		return err
	}
	return nil
}

func (r *RedisRt) BatchFetch(keys [][]byte) ([][]byte, error) {
	b := [][]byte{}
	for _, key := range keys {
		row := r.client.HGet(context.Background(), r.KeyName, string(key))
		if result, err := row.Result(); err != nil {
			if err == redis.Nil {
				b = append(b, nil)
			} else {
				return nil, err
			}

		} else {
			b = append(b, convertStringToByte(result))
		}
	}
	return b, nil
}

func (r *RedisRt) Sync() error {
	result := r.client.BgSave(context.Background())
	if _, err := result.Result(); err != nil {
		return err
	}
	return nil
}

func (r *RedisRt) Load(_ ...interface{}) error {
	r.client = redis.NewClient(&redis.Options{
		Addr:     r.Ip + ":" + strconv.Itoa(r.Port),
		Password: r.PassWord,
		DB:       r.DbNum,
	})
	return nil
}

// StopGdb defines the way to stop gdb
type StopGdb interface {
	StopService(params ...interface{}) error
}

func (gdb *Gdb) StopService(_ ...interface{}) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	<-ctx.Done()
	stop()
	return gdb.CloseGdb()
}

// DefaultOptions returns default options of gdb
//
// DriverName: sqlite3
//
// Dsn: ./itemDb.db
//
// UseInnerStop: true
//
//RtGdb: &FastCacheRt{RealTimePath: "./realTimeDb"}
func DefaultOptions() *Options {
	return &Options{
		DriverName:   "sqlite3",
		Dsn:          "./itemDb.db",
		UseInnerStop: true,
		RtGdb:        &FastCacheRt{RealTimePath: "./realTimeDb"},
	}
}
