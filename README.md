## GDB
English | [Chinese](README-zh.md)

GDB is a historical database  based on [goleveldb](https://pkg.go.dev/github.com/syndtr/goleveldb/leveldb) with high performance.
It uses goLevelDb to store historical Data, use memory or Redis to store realTime Data and part history data.
It supports restful and gRPC protocol.If you want to need to store large amount of history data, you will love it. 

[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/JustKeepSilence/gdb)
[![Go Report Card](https://goreportcard.com/badge/github.com/JustKeepSilence/gdb)](https://goreportcard.com/report/github.com/JustKeepSilence/gdb)
[![Release](https://img.shields.io/github/v/release/JustKeepSilence/gdb)](https://github.com/JustKeepSilence/gdb/releases)
[![go-version](https://img.shields.io/github/go-mod/go-version/JustKeepSilence/gdb)]()
[![count](https://img.shields.io/github/languages/count/JustKeepSilence/gdb)]()


## Features
- High writing and reading performance
- Multiple ways to access history data(batchGet, conditionFilter...)
- Simulate data based on existing data with js code.
- Support restful ,gRPC, https protocol.
- Fine permission control(tokenBased permission control + casbinBased route permission control)
- Fine desktop client ensure you can use it even you don't have any programming knowledge
- Fine document

## Contents
- [Design Ideas](#Design-ideas)
    - [ItemDb](#itemDb)
    - [DataDb](#dataDb)
- [Quick Start](#quick-start)
- [GdbServer](#GdbServer)
    - [Build GDB](#build-gdb)
    - [Download GDB](#download-gdb)
    - [HTTPS Mode](#https-mode)
    - [Authorization Mode](#authorization-mode)
    - [Route Permission](#route-permission)
- [Restful API Examples](#restful-api-examples)
- [gRPC API Examples](#grpc-api-examples)
- [GdbUI](#gdbUI)
- [FAQ](#faq)

## Design Ideas
<img src="https://github.com/JustKeepSilence/gdb/blob/master/db/templateFiles/desginIdeas.png">
As shown about, whole gdb database can be divided in two parts, that is itemDb and dataDb

### ItemDb
ItemDb is used to store items in gdb.The root store unit in gdb is group, in one group, you can 
add many items, every item has its own realTime and history data in DataDb.Items in different group
is isolated.In group, we use itemName as the unique identifier.So the first step to use gdb is to 
add your own groups and items!

### DataDb
DataDb is used to store data of items, data in gdb can be divided into realTimeData and history Data.
when you write data to database, on the one hand, we will write data to realTime
database, for realTimeDataBase, we provide memory dataBase and redis to store realTimeData, 
if you use gdb in your own Go project, you can also implement [RtGdb](https://github.com/JustKeepSilence/gdb/blob/9047c12a6d592f60d51c7d7ab3292048340360aa/db/db_type.go#L347) interface to custom your own 
way to store realTimeData.On the other hand, we will write data to history dataBase in memory, 
and we will batchSync these data to disk period.


## Quick Start
To use gdb in your own Go project, you need install [Go](https://golang.org/) (**version 1.16+ is required**), then set GO111MODULE=ON
for these users, we only provide core functions of gdb,that is read and write data.But you can customer any functions you want by this way.
```sh
go get github.com/JustKeepSilence/gdb@latest
```
Then import gdb in your own code
```go
import "github.com/JustKeepSilence/gdb/db"
```

```go
import (
	"encoding/json"
	"fmt"
	"github.com/JustKeepSilence/gdb/db"
	"log"
	"math"
	"math/rand"
	"io/ioutil"
	"time"
)

func main()  {
    //initial db with sqlite3
	if gdb, err := db.NewGdb("./historyDb", time.Hour, time.Minute*5, db.DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gdb.hisTimeDuration)
	}
	
	// initial db with mysql
	if gdb, err := db.NewGdb("./historyDb", time.Hour, time.Minute*5, &db.Options{
		DriverName:   "mysql",
		Dsn:          "root:admin@123@tcp(192.168.0.166:3306)/itemDb",
		UseInnerStop: true,
		RtGdb:        &FastCacheRt{RealTimePath: "./realTimeDb"},
	}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gdb.hisTimeDuration)
	}
	
	// initial db with mysql + redis
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, &db.Options{
		DriverName:   "mysql",
		Dsn:          "root:admin@123@tcp(192.168.0.166:3306)/itemDb",
		UseInnerStop: true,
		RtGdb:        &RedisRt{
			Ip:       "192.168.0.199",
			Port:     6379,
			PassWord: "",
			DbNum:    0,
			KeyName:  "gdbRealTime",
		},
	}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gdb.hisTimeDuration)
	}
	
	if gdb, err := db.NewGdb("./historyDb", time.Hour, time.Minute*5, db.DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		// add groups
		if r, err := gdb.AddGroups(db.AddedGroupInfo{
			GroupName:   "3DCS",
			ColumnNames: []string{"unit", "description"},
		}, db.AddedGroupInfo{
			GroupName:   "4DCS",
			ColumnNames: nil,
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.EffectedRows, r.Times)
		}

		// add items to group
		if r, err := gdb.AddItems(db.AddedItemsInfo{
			GroupName: "3DCS",
			ItemValues: []map[string]string{{"itemName": "X", "dataType": "float32", "description": "", "unit": ""},
				{"itemName": "Y", "dataType": "float32", "description": "", "unit": ""},{"itemName": "Z", "dataType": "float32", "description": "", "unit": ""}},
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}

		if r, err := gdb.AddItems(db.AddedItemsInfo{
			GroupName: "4DCS",
			ItemValues: []map[string]string{{"itemName": "X1", "dataType": "float32"},
				{"itemName": "Y1", "dataType": "float32"},{"itemName": "Z1", "dataType": "float32"}},
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}

		// batchWriteFloatRealTimeData
		if r, err := gdb.BatchWriteFloatData([]string{"3DCS", "4DCS"}, [][]string{{"X", "Y", "Z"}, {"X1", "Y1", "Z1"}}, [][]float32{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}

		// getRealTimeData
		if r, err := gdb.GetRealTimeData([]string{"3DCS", "4DCS", "3DCS"}, []string{"X", "X1", "Y"});err!=nil{
			log.Fatal(err)
		}else{
			r1,_ := json.Marshal(r.RealTimeData)
			fmt.Println(string(r1))
		}

		// writeFloatHistoryData
		//generate history data of month
		seconds := 24 * 3600 * 30
		now := time.Now()
		groupNames := []string{"4DCS", "4DCS"}
		itemNames := []string{"X1", "Y1"}
		var timeStamp []int32
		var xItemValue []float32
		var yItemValue []float32
		for i := 0; i < seconds; i++ {
			timeStamp = append(timeStamp, int32(now.Add(time.Duration(i) * time.Second).Unix() + 8 * 3600))
			xItemValue = append(xItemValue, rand.Float32()*math.Pi)
			yItemValue = append(yItemValue, rand.Float32()*math.E)
		}
		if r, err := gdb.BatchWriteFloatHistoricalData(groupNames, itemNames, [][]int32{timeStamp, timeStamp}, [][]float32{xItemValue, yItemValue});err!=nil{
			log.Fatal(err)
		}else{
			fmt.Println(r.Times, r.EffectedRows)
		}

		// getFloatHistoryData
		st := int32(time.Now().Add( time.Hour * 9 * 24).Unix() + 8 * 3600)
		et := int32(time.Now().Add(time.Hour * 10 * 24).Unix() + 8 * 3600)
		if r, err := gdb.GetFloatHistoricalData([]string{"3DCS", "4DCS"}, []string{"Y", "X1"}, []int32{st, st}, []int32{et, et}, []int32{6 * 3600, 6 * 3600});err!=nil{
			log.Fatal(err)
		}else{
			r1, _ := json.Marshal(r.HistoricalData)
			fmt.Println(string(r1))
		}

		// getFloatRawHistoryData==>all history data
		if r, err := gdb.GetFloatRawHistoricalData([]string{"4DCS"}, []string{"X1"});err!=nil{
			log.Fatal(err)
		}else{
			v, _ := r.HistoricalData.Get("X1")
			r1, _ := json.Marshal(r)
			_ = ioutil.WriteFile("./h.txt", r1, 0766)
			fmt.Println(len(v.([]interface{})[0].([]int32)))
		}

		// getFloatHistory data with given timeStamps
		ts := []int32{int32(now.Add(time.Hour * 24 * -30).Unix() + 8 * 3600)}
		for i := 0; i < 5; i++ {
			ts = append(ts, int32(now.Add(time.Hour * 24 * time.Duration(i)).Unix() + 8 * 3600))
		}
		ts = append(ts, int32(now.Add(time.Hour * 24 * 60).Unix() + 8 * 3600))  // history of ts[0] and ts[len(ts) - 1] not exist, so we will not
		// return value of this two timeStamp
		if r, err := gdb.GetFloatHistoricalDataWithStamp([]string{"4DCS", "3DCS"}, []string{"X1", "Y"}, [][]int32{ts, ts});err!=nil{
			log.Fatal(err)
		}else{
			r1, _ := json.Marshal(r.HistoricalData)
			fmt.Println(string(r1))
		}

		// getFloatHistoryData with given condition
		st, et = int32(now.Add(time.Hour * 2).Unix() + 8 * 3600), int32(now.Add(time.Hour * 6).Unix() + 8 * 3600)
		// without deadZone condition
		if r, err := gdb.GetFloatHistoricalDataWithCondition("4DCS", []string{"X1", "Y1"}, st, et, 10, `item["X1"]>= 1 && item["Y1"]<= 4` ,nil);err!=nil{
			log.Fatal(err)
		}else{
			fmt.Println(r.Times)
			r1, _ := json.Marshal(r.HistoricalData)
			_ = ioutil.WriteFile("./hf2.json", r1, 0766)
		}
		// with deadZone condition
		if r, err := gdb.GetFloatHistoricalDataWithCondition("4DCS", []string{"X1", "Y1"}, st, et, 10, `item["X1"]>= 1 && item["Y1"]<= 4` ,[]db.DeadZone{{ItemName: "X1", DeadZoneCount: 3}});err!=nil{
			log.Fatal(err)
		}else{
			fmt.Println(r.Times)
			r1, _ := json.Marshal(r.HistoricalData)
			_ = ioutil.WriteFile("./hf2.json", r1, 0766)
		}
		// withOut filterCondition
		if r, err := gdb.GetFloatHistoricalDataWithCondition("4DCS", []string{"X1", "Y1"}, st, et, 10, `true` ,[]db.DeadZone{{ItemName: "X1", DeadZoneCount: 3}});err!=nil{
			log.Fatal(err)
		}else{
			fmt.Println(r.Times)
			r1, _ := json.Marshal(r.HistoricalData)
			_ = ioutil.WriteFile("./hf2.json", r1, 0766)
		}
		// withOut filterCondition and deadZone condition == GetFloatHistoricalData
		if r, err := gdb.GetFloatHistoricalDataWithCondition("4DCS", []string{"X1", "Y1"}, st, et, 10, `true` ,nil);err!=nil{
			log.Fatal(err)
		}else{
			fmt.Println(r.Times)
			r1, _ := json.Marshal(r.HistoricalData)
			_ = ioutil.WriteFile("./hf2.json", r1, 0766)
		}

		// deleteFloatHistoryData
		if r, err := gdb.DeleteFloatHistoricalData([]string{"4DCS", "4DCS"}, []string{"X1", "Y1"}, []int32{st, st}, []int32{et, et});err!=nil{
			log.Fatal(err)
		}else{
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}
```
Notes: In order to reduce the size of the entire project, in this case only the core functions of gdb are included, 
unless you use gdbServer tags when compiling

## GdbServer
If you are not familiar with go, and want to use gdb as back-end database only, you can [build-gdb](#build-gdb), then run 
gdb in your server.Also, you can [download](#download-gdb) the compiled installer and run it directly.
In this way, you can't customize your own behavior, but you can use restful or grpc api provided by gdb, as well as 
token-control for every api we provided.For more details you can see [restful-examples](#restful-api-examples) or [grpc-examples](#grpc-api-examples) 


### Build GDB
First, you need to clone gdb using the following command:
```shell
git clone https://github.com/JustKeepSilence/gdb.git
```
Then change to gdb/gdbServer directory, run the following command(need Go evn):
```shell
go build -tags=jsoniter -tags=gdbServer -o ../gdb
```
Notes: you must add gdbServer tags when building gdb Client, otherWise only core function without client will be compiled.
After that, you can customize your own config in config.json.For more details about config you can see [configFiles](#https://github.com/JustKeepSilence/gdb/blob/master/config.json)
```json
// Notes: you can use // single line comments in json file
{
  "baseConfigs": {
    "ip": "",
    "port": 8082,
    "dbPath": "./historyData",
    "applicationName": "gdb",
    "authorization": true,
    "mode": "http",
    "useRedis": false,
    "rtTimeDuration": 3600,
    "hisTimeDuration": 300
  },
  "itemDbConfigs": {
    "driverName": "sqlite3",
    "dsn": "file:itemDb.db?_auth_user=seu&_auth_pass=admin@123&_auth_crypt=SHA1@_vacuum=1"
  },
  "httpsConfigs": {
    "ca": false,
    "selfSignedCa": false,
    "caCertificateName": "",
    "serverCertificateName": "gdbServer.crt",
    "serverKeyName": "gdbServer.key"
  },
  "logConfigs": {
    "logWriting": true,
    "Level": "Error",
    "expiredTime": 3600
  },
  "redisConfigs": {
    "redisIp": "192.168.0.199",
    "redisPort": 6379,
    "redisPassWord": "",
    "redisDb": 0,
    "keyName": "gdbRealTime"
  }
}
```
Notes: you need set gdb,config.json, and ssl folder in the same path to sure gdb work normally.

### Download Gdb

if you are not familiar with go at all, you can also directly download the compiled installer we provided,
the download url is: https://wws.lanzoui.com/irUJ0rowr8d, download passWord is 7659

### HTTPS Mode

gdb support https mode for restful nad gRPC,if you want to run with https mode, you need to set 
mode filed in configs.json to https, and customize your own https configs.Then put your own certificate
to ssl folder and put ssl folder the same path as gdb executable program. Or you can use default certificate 
we provided without ca root.

Notes: selfSignedCa is not allowed on windows at moment.And if you want to use ca root, you need 
to set ca field to true and set the caCertificateName field in config.json

### Authorization Mode

gdb support token authorization mode, if you want to run with authorization mode, you need to set authorization field to
true.Then you need to add Authorization field to header for all requests, only under authorization mode, routes permission
will work.

### Route Permission
When the GDB service is running in the authorization mode, the routing permission control will work. The users of the 
GDB service are divided into three categories: super users, ordinary users and tourists. User management can be done 
through [Client](#gdbUI) to finish. For the routing permissions of the group interface and item interface, 
each type of user is fixed, and you cannot modify it. But for the data interface, that is, the routing interface 
permissions related to reading data, you can modify it on the client.For the routing permissions of each type of user, 
you can check here [details](#https://github.com/JustKeepSilence/gdb/blob/master/examples/restfulExamples/groupExamples.js)

## Restful API Examples
The GDB service supports a standard restful interface, and the returned data format is
```json
{
  "code": 200,
  "message": "",
  "data": {}
}
```
code is the status code of the request, 200 is success, 500 is failure, and 401 is unauthorized.

The message is the information in the request process, the code is 500 is the request failure information, and the 200 is the empty string.

data is the response data, the type is object

For restful use cases, please check [RestFulExamples](https://github.com/JustKeepSilence/gdb/tree/master/examples/restfulExamples)

## gRPC API Examples
We recommend you to interact with gdb by gRPC.For more details, you can see [gRPCExamples](https://github.com/JustKeepSilence/gdb/tree/master/examples/gRPCExamples)

## GdbUI
The GDB service can be easily operated through the GDB client. For specific download and use
please see [gdbUI](https://github.com/JustKeepSilence/gdbUI)

## FAQ
1. How to obtain the timeStamp consistent with gdb
```go
// The timestamp in gdb uses the unix timestamp that comes with go,timeZone is UTC
import (
 "time"
)
n := time.Now
timeStamp := time.Date(n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Minute(), 0, time.UTC)
timeStamp := n.Unix() + 8 * 3600
```
So, above this, here are some examples to show how to get the timeStamp consistent with gdb
### C#
```C#
var t1 = new DateTime(2021, 2, 11, 14, 26, 26);
long timeStamp1 = (long)(t1 - new DateTime(1970, 1, 1, 0, 0, 0)).TotalSeconds;
```
### Js(In China)
In node Js, you can use moment.js to handle time easily, to getTimeStamp of now 
```js
moment().unix() + 8 * 3600
```
### Python(In China)
```python
from datetime import datetime
int(datetime(2021, 2, 11, 14, 26, 26).timestamp()) + 8 * 3600
```