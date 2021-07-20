## GDB

GDB是基于[goleveldb](https://pkg.go.dev/github.com/syndtr/goleveldb/leveldb)
用Go语言编写的高效的实时-历史数据库.支持restful和gRPC协议，支持基于js的二次自定义模拟数据.
如果你需要存储大量的过程历史数据，它会是一个很适合的选择！

[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/JustKeepSilence/gdb)
[![Go Report Card](https://goreportcard.com/badge/github.com/JustKeepSilence/gdb)](https://goreportcard.com/report/github.com/JustKeepSilence/gdb)
[![Release](https://img.shields.io/github/v/release/JustKeepSilence/gdb)](https://github.com/JustKeepSilence/gdb/releases)
[![go-version](https://img.shields.io/github/go-mod/go-version/JustKeepSilence/gdb)]()
[![count](https://img.shields.io/github/languages/count/JustKeepSilence/gdb)]()

## GDB的特点
- 很高的读写性能
- 提供了多种获取历史数据的方法(批量获取，条件筛选)
- 支持基于JS的二次模拟数据的生成
- 内置支持https,restful,gRPC协议
- 完善的权限控制(token访问控制和基于casbin的路由权限控制)
- 基于Electron的客户端，方便使用
- 完善的使用文档

## 内容
- [设计理念](#设计理念)
    - [ItemDb](#itemDb)
    - [DataDb](#dataDb)
- [快速开始](#快速开始)
- [Gdb服务](#Gdb服务)
    - [编译GDB服务](#编译GDB)
    - [下载GDB服务](#下载GDB服务)
    - [HTTPS模式](#HTTPS模式)
    - [授权模式](#授权模式)
    - [路由控制](#路由控制)
- [Restful示例](#Restful示例)
- [gRPC示例](#gRPC示例)
- [GDB客户端](#GDB客户端)
- [FAQ](#faq)

## 设计理念
<img src="https://github.com/JustKeepSilence/gdb/blob/master/db/templateFiles/desginIdeas.png">
如图所示,整个gdb数据库可以分为两部分,即itemDb和dataDb.

### ItemDb
ItemDb用来存储数据库中的item.gdb数据库中的基本存储单元是group,每个group中可以有很多item,每个item都有
自己对应的实时数据和历史数据,每个组中的item都是彼此独立的,但是组名是唯一的(唯一标识).所以使用gdb的第一步
就是先向数据库中添加group,再向组中添加item,有了item之后,你就可以向数据库写入实时数据,gdb会自动帮你将数
据存储下来,并形成历史数据。

### DataDb
DataDb用来存储item的数据,包括历史数据和实时数据以及整个数据库运行过程中的运行信息数据.对于实时数据库部分
gdb内置了两种存储方式,分别是[fastCache](https://pkg.go.dev/github.com/VictoriaMetrics/fastcache?GOOS=windows#Cache.GetBig)
和redis,默认的方式为fastCache.如果在你的项目中,你想自定义实时数 的存储方式,你只需要实现[RtGdb](https://github.com/JustKeepSilence/gdb/blob/9047c12a6d592f60d51c7d7ab3292048340360aa/db/db_type.go#L347)
接口即可.对于历史数据,最近一段时间的历史数据会存储在内存中,并且会定期同步到disk中.

## 快速开始
在你自己的Go项目中使用gdb数据库,你需要先安装[Go](https://golang.org/)
环境(需要1.16以上)并开启GoMODULE模式.对于这种使用方式,我们只提供了gdb数据库的最核心功能,即读写实时以及历史数据的功能.但是这种使用方式
你可以对数据库进行任意的扩展,并且如果想实现权限或者其他功能,你可以参考我们的源代码.
```sh
go get github.com/JustKeepSilence/gdb@latest
```

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

## Gdb服务
如果你对Go语言不熟悉或者想使用gdb作为后端的数据库,那么你可以使用GDB服务,这种使用方式,你可以使用GDB的全部功能,但是同时意味着你不能对GDB进行随意的
扩展,但是我们相信您需要的绝大部分功能我们应该都已经涵盖了,如果有需求,可以联系我们. 为了使用GDB服务,你需要[编译](#编译GDB服务)或者[下载](#下载GDB服务)GDB服务,之后运行GDB服务即可.

### 编译GDB服务
首先,克隆我们的项目到本地
```shell
git clone https://github.com/JustKeepSilence/gdb.git
```
然后切换到gdb/gdbServer文件夹,运行下面的命令(需要安装Go环境)
```shell
go build -tags=jsoniter -tags=gdbServer -o ../gdb
```
注意:jsoniter条件编译可以省略,但是gdbServer条件编译不能省略,否则gdb服务会编译失败.编译成功之后可以自定义配置文件.关于配置文件详情可以查看[配置文件](#https://github.com/JustKeepSilence/gdb/blob/master/config.json)
```json
// Notes: 在配置文件中你可以使用//开始的单行注释
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
你可以选择我们默认的配置进行gdb服务的启动.但是你必须要将配置文件,gdb服务程序以及ssl文件夹放在同一个路径下才可以保gdb服务的成功运行.

### 下载GDB服务
如果你对go语言不熟悉或者没有安装go环境,你可以直接下载我们编译好的二进制文件,下载地址为:https://wws.lanzoui.com/icUGzpojb5e,下载密码是bwst

### HTTPS模式
GDB服务支持以https模式运行,如何你自己有证书文件,你只需要将证书文件放置到ssl文件夹中,并在配置文件中指定证书的名称即可.gdb默认的https模式使用的是
无签证的证书,在https-gRPC连接时的域名为gdb.com.

### 授权模式
开启GDB服务的授权模式,你只需要将配置文件中的authorization字段设置为true即可.这种模式下,所有的请求都需要在请求头部加上token信息.只有在授权模
式下,路由权限的管理才会起作用.

### 路由控制
当GDB服务以授权模式运行的时候,路由权限控制就会起作用.GDB服务的用户分为三类,分别是超级用户,普通用户和游客.用户的管理都可以通过[客户端](#GDB客户端)去完成.
其中对于group接口和item接口的路由权限,每类用户是定死的,你无法进行修改.但是对于data接口,即和读取数据相关的路由接口权限,可以在客户端进行修改.
对于每类用户的路由权限,你可以在这里查看[详情](#https://github.com/JustKeepSilence/gdb/blob/master/examples/restfulExamples/groupExamples.js)

### Restful示例
GDB服务支持标准的restful接口,返回的数据格式为

```json
{
  "code": 200,
  "message": "",
  "data": {}
}
```
其中code为请求的状态码,200为成功,500为失败,403为未授权.

message为请求过程中的信息,code为500则为请求失败的信息,200则为空字符串.

data为响应的数据,类型为object

关于restful的使用案例可以查看[RestFulExamples](https://github.com/JustKeepSilence/gdb/tree/master/examples/restfulExamples)

### gRPC示例
我们建议使用gRPC的方式和GDB服务进行通信,具体的示例可以查看[gRPCExamples](https://github.com/JustKeepSilence/gdb/tree/master/examples/gRPCExamples)

## GDB客户端
通过GDB客户端可以很方便的操作GDB服务,具体下载及使用可以查看[gdbUI](https://github.com/JustKeepSilence/gdbUI)

## FAQ
1. 如何与GDB中的时间戳保持同步
GDB使用的是标准的UNIX时间戳
```go
// The timestamp in gdb uses the unix timestamp that comes with go,timeZone is UTC
import (
 "time"
)
n := time.Now
timeStamp := time.Date(n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Minute(), 0, time.UTC)
timeStamp := n.Unix() + 8 * 3600
```
### C#
```C#
var t1 = new DateTime(2021, 2, 11, 14, 26, 26);
long timeStamp1 = (long)(t1 - new DateTime(1970, 1, 1, 0, 0, 0)).TotalSeconds;
```
### Js(In China)
NodeJS中你只需要使用moment.js就可以轻易的获取时间戳
```js
moment().unix() + 8 * 3600
```
### Python(In China)
```python
from datetime import datetime
int(datetime(2021, 2, 11, 14, 26, 26).timestamp()) + 8 * 3600
```