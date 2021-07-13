// +build gdbClient

/*
createTime: 2021/6/9
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"embed"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

//go:embed route.json
var routeFile embed.FS

// inner type of gdb, only for gdb service

var (
	superUserRoutes, commonUserRoutes, visitorUserRoutes []string
)

type Config struct {
	BaseConfigs   `json:"baseConfigs"`
	ItemDbConfigs `json:"itemDbConfigs"`
	HttpsConfigs  `json:"httpsConfigs"`
	LogConfigs    `json:"logConfigs"`
	RedisConfigs  `json:"redisConfigs"`
}

type BaseConfigs struct {
	IP              string `json:"ip"`
	Port            int64  `json:"port"`
	DbPath          string `json:"dbPath"`
	ApplicationName string `json:"applicationName"`
	Authorization   bool   `json:"authorization"`
	Mode            string `json:"mode"`
	UseRedis        bool   `json:"useRedis"`
	RtTimeDuration  int    `json:"rtTimeDuration"`
	HisTimeDuration int    `json:"hisTimeDuration"`
}

type ItemDbConfigs struct {
	DriverName string `json:"driverName"`
	Dsn        string `json:"dsn"`
}

type LogConfigs struct {
	LogWriting  bool   `json:"logWriting"`
	Level       string `json:"level"`
	ExpiredTime int    `json:"expiredTime"`
}

type HttpsConfigs struct {
	Ca                    bool   `json:"ca"`
	SelfSignedCa          bool   `json:"selfSignedCa"`
	CaCertificateName     string `json:"caCertificateName"`
	ServerCertificateName string `json:"serverCertificateName"`
	ServerKeyName         string `json:"serverKeyName"`
}

type RedisConfigs struct {
	RedisIp       string `json:"redisIp"`
	RedisPort     int    `json:"redisPort"`
	RedisPassWord string `json:"redisPassWord"`
	RedisDb       int    `json:"redisDb"`
	KeyName       string `json:"keyName"`
}

type Route struct {
	AllRoutes         []string `json:"allRoutes"`
	CommonUserRoutes  []string `json:"commonUserRoutes"`
	VisitorUserRoutes []string `json:"visitorUserRoutes"`
}

func innerNewGdb(dbPath string, useRedis bool, redisIP string, redisPort int, redisPassWord string,
	redisDb int, keyName string, driverName, dsn string, rtTimeDuration, hisTimeDuration int) (*Gdb, error) {
	// check whether the given path exist
	if !fileExist(dbPath) {
		if err := os.MkdirAll(dbPath, 0766); err != nil {
			return nil, err
		}
	}
	var g *Gdb
	var delimiter string
	path, _ := filepath.Abs(dbPath) // get abs path of given path
	if runtime.GOOS == "windows" {
		g = &Gdb{
			dbPath: path,
		}
		delimiter = "\\"
	} else {
		g = &Gdb{
			dbPath: path,
		}
		delimiter = "/"
	}
	// get routes
	if contents, err := routeFile.ReadFile("route.json"); err != nil {
		return nil, err
	} else {
		route := Route{}
		if err := json.Unmarshal(contents, &route); err != nil {
			return nil, err
		}
		superUserRoutes, commonUserRoutes, visitorUserRoutes = route.AllRoutes, route.CommonUserRoutes, route.VisitorUserRoutes
	}
	if err := g.initialItemDb(true, driverName, dsn); err != nil {
		return nil, err
	}
	rtPath := g.dbPath + delimiter + "realTimeData"
	if useRedis {
		// use redis
		rt := &RedisRt{
			Ip:       redisIP,
			Port:     redisPort,
			PassWord: redisPassWord,
			DbNum:    redisDb,
			KeyName:  keyName,
		}
		if err := g.initialDb(rt); err != nil {
			return nil, err
		}
	} else {
		rt := &FastCacheRt{
			RealTimePath: rtPath,
		}
		if err := g.initialDb(rt); err != nil {
			return nil, err
		}
	}
	g.rtTimeDuration = time.Duration(rtTimeDuration) * time.Second
	g.hisTimeDuration = time.Duration(hisTimeDuration) * time.Second
	if m, err := model.NewModelFromString(`
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = r.sub == p.sub && r.obj == p.obj || r.sub == "admin" || r.obj == "all"
	`); err != nil {
		return nil, err
	} else {
		if e, err := casbin.NewEnforcer(m, g); err != nil {
			return nil, err
		} else {
			g.e = e
			return g, nil
		}
	}
}
