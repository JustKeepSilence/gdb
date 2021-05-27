## GDB
GDB is a real-time database encapsulated based on [goleveldb](https://pkg.go.dev/github.com/syndtr/goleveldb/leveldb)
it can be used to obtain and store large amount of historical data in various ways(including gettting raw data, filtered data
with given condition, etc...),it provides rest, gRPC interface and desktop client , and it 
allows you to generate your own data based on existing data by coding js on desktop client.If you need deal with big data,
you will love GDB.

[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/JustKeepSilence/gdb)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-gonic/gin)](https://goreportcard.com/report/github.com/JustKeepSilence/gdb)
[![Release](https://img.shields.io/github/v/release/JustKeepSilence/gdb)](https://github.com/JustKeepSilence/gdb/releases)
[![go-version](https://img.shields.io/github/go-mod/go-version/JustKeepSilence/gdb)]()
[![count](https://img.shields.io/github/languages/count/JustKeepSilence/gdb)]()


## Features
- High writing performance
- Fine control of historical data
- Simulate data based on existing data with js
- Token-based permission control
- restful and gRPC api
- support https api
- desktop client based on Electron
- fine system documentation and interface documentation

## Contents
- [Quick Start](#quick-start)
    - [Installation](#installation)
- [GdbServer](#GdbServer)
    - [Build GDB](#build-gdb)
    - [Download GDB](#download-gdb)
    - [Run With HTTPS Mode](#run-with-https-mode)
    - [Run With Authorization Mode](#run-with-authorization-mode)
- [Restful API Examples](#restful-api-examples)
    - [Page](#page)
    - [Group](#group)
    - [Item](#item)
    - [Data](#data)
- [gRPC API Examples](#grpc-api-examples)

- [GdbUI](#desktop-application)
- [FAQ](#faq)

## Quick Start
If you are familiar with go language, you call [install](#installation) gdb and then use it in your
project to customize your own behavior,For more details,you can see [document](https://pkg.go.dev/github.com/JustKeepSilence/gdb) or 
[examples](https://github.com/JustKeepSilence/gdb/tree/master/examples)
The Base of gdb is group and item, item is the subset of group, you need to add group to gdb, then add item to
group, after that, you can write realTime Data to item and get historical Data.

### Installation
Gdb is a cgo project, to run or build it, you need [gcc](https://gcc.gnu.org/) ,and install [Go](https://golang.org/) (**version 1.16+ is required**), then set GO111MODULE=ON
```sh
go get github.com/JustKeepSilence/gdb
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
    // initial gdb
	if gdb, err := db.NewGdb("./leveldb", "./itemDb");err!=nil{
	    log.Fatal(err)	
    }else{
    	// add groups
        if _, err := gdb.AddGroups(db.AddedGroupInfo{
            GroupName:   "1DCS",
            ColumnNames: []string{"Column1", "Column2", "Column3"},  // column name can't be itemName
            });err!=nil{
            log.Fatal(err)
        }
        // add items
        if _, err := gdb.AddItems(AddedItemsInfo{
            GroupName: "1DCS",
            ItemValues: []map[string]string{{"itemName": "x", "description": "x", "dataType": "float64"},
                {"itemName": "y", "description": "y", "dataType": "float64"}, {"itemName": "z", "description": "z", "dataType": "float64"},
                {"itemName": "item1", "description": "item1", "dataType": "float64"}, {"itemName": "item2", "description": "item2", "dataType": "float64"}},
        }); err != nil {
            log.Fatal(err)
        }
        // batch write, y = 2 * x
        if _, err := gdb.BatchWrite([]db.ItemValue{{ItemName: "x", Value: 1.0, GroupName: "1DCS"}, {ItemName: "y", Value: 2.0, GroupName: "1DCS"}}...);err!=nil{
            log.Fatal(err)
        }else{
        // get latest updated value of given items
            if r, err := gdb.GetRealTimeData([]string{"1DCS", "1DCS"}, "x", "y");err!=nil{
            log.Fatal(err)
        }else{
            d, _ := json.Marshal(r)
            fmt.Println(string(d))
        }
        }
        // write historical data
        // mock one hour historical data
        var xData, yData, ts []interface{}
        now := time.Now()
        fmt.Println("now: ", now.Format("2006-01-02 15:04:05"))
        r := rand.New(rand.NewSource(99))
        for i := 0; i < 3600; i++ {
            x := float64(r.Intn(3600)) * math.Pi
            y := 2 * x
            t := now.Add(time.Second * time.Duration(i)).Unix() + 8 * 3600
            xData = append(xData,x)
            yData = append(yData, y)
            ts = append(ts, int(t))
        }
        if err := gdb.BatchWriteHistoricalData([]db.HistoricalItemValue{{ItemName: "x", Values: xData, TimeStamps: ts, GroupName: "1DCS"}, {ItemName: "y", Values: yData, TimeStamps: ts, GroupName: "1DCS"}}...);err!=nil{
            log.Fatal(err)
        }else{
            // get raw historical data for debugging
            if r, err := gdb.GetRawHistoricalData([]string{"1DCS"}, "x");err!=nil{
            log.Fatal(err)
        }else{
            d, _ := json.Marshal(r)
            _ = ioutil.WriteFile("./rawX.txt", d, 0644)
        }
        // get historical data with given itemName, startTime, endTime and intervals
        stX := int(now.Add(time.Minute * 5).Unix() + 8 * 3600)
        etX := int(now.Add(time.Minute * 25).Unix() + 8 * 3600)
        stY := int(now.Add(time.Minute * 35).Unix() + 8 * 3600)
        etY := int(now.Add(time.Minute * 55).Unix() + 8 * 3600)
        if r, err := gdb.GetHistoricalData([]string{"1DCS", "1DCS"}, []string{"x", "y"}, []int{stX, stY}, []int{etX, etY}, []int{2, 10});err!=nil{
            log.Fatal(err)
        }else{
            d, _ := json.Marshal(r)
            _ = ioutil.WriteFile("./hX.txt", d, 0644)
        }
        // get historical data with given itemName
        if r, err := gdb.GetHistoricalDataWithStamp([]string{"1DCS", "1DCS"},[]string{"x", "y"}, [][]int{{stX, etX}, {stY, etY}}...);err!=nil{
            log.Fatal(err)
        }else{
            d, _ := json.Marshal(r)
            fmt.Println(string(d))
        }
        // get historical data with condition
        if r, err := gdb.GetHistoricalDataWithCondition([]string{"x", "y"}, []int{stX, stY}, []int{etX, etY}, []int{2, 10}, `item["x"] > 0 && item["y"] > 1000`, []db.DeadZone{}...);err!=nil{
            log.Fatal(err)
        }else{
            d, _ := json.Marshal(r)
            _ = ioutil.WriteFile("./f.txt", d, 0644)
            }
        }
    }
}
```
Notes: In order to reduce the size of the entire project, in this case only the core functions of gdb are included, 
unless you use gdbClient tags when compiling

## GdbServer
If you are not familiar with go, and want to use gdb as back-end database only, you can [build-gdb](#build-gdb), then run 
gdb in your server.Also, you can [download](#download-gdb) the compiled installer and run it directly.
In this way, you can't customize your own behavior, but you can use restful or grpc api provided by gdb, as well as 
token-control for every api we provided.For more details you can see [restful-examples](#restful-api-examples) or [grpc-examples](#grpc-api-examples) or [documents](https://app.gitbook.com/@justkeepsilence/s/gdb/~/settings/share)


### Build GDB
you need to [install](#installation)Gdb firstly, then change to gdb/main directory and run the following command
```sh
go build -tags=jsoniter -tags=gdbClient -o ../gdb
```
Notes: you must add gdbClient tags when building gdb Client, otherWise only core function without client will be compiled.
After that, you can customize your own config in config.json.For more details about config you can see https://github.com/JustKeepSilence/gdb/blob/master/config.json
```json
// Notes: you can use // single line comments in json file
{
  "gdbConfigs": {
    "ip": "",
    "port": 8082,
    "dbPath": "./leveldb",
    "itemDbPath": "./itemDb",
    "applicationName": "gdb",
    "authorization": true,
    "mode": "http",
    "httpsConfigs": {
      "ca": false,
      "selfSignedCa": false,
      "caCertificateName": "",
      "serverCertificateName": "gdbServer"
    }
  },
  "logConfigs": {
    "logWriting": true,
    "Level" : "Error",
    "expiredTime": 86400
  }
}

```

Notes: you need set gdb,config.json, and ssl folder in the same path to sure gdb work normally.

### Download Gdb
if you are not familiar with go at all, you can also directly download the compiled installer we provided,
the download url is: https://wws.lanzous.com/iYPy0ovihdi, download passWord is g65q

### Run With HTTPS Mode

gdb support https mode for restful nad gRPC,if you want to run with https mode, you need to set 
mode filed in configs.json to https, and customize your own https configs.Then put your own certificate
to ssl folder and put ssl folder the same path as gdb executable program. Or you can use default certificate 
we provided without ca root.

Notes: selfSignedCa is not allowed on windows at moment.And if you want to use ca root, you need 
to set ca field to true and set the caCertificateName field in config.json

### Run With Authorization Mode

gdb support token authorization mode, if you want to run with authorization mode,
you need to set authorization field to true.Then you need to add Authorization field to
header if you use restful, or to context if you use gRPC.For more details, you can see
[document](https://blog.csdn.net/qq_39778055/article/details/114844756)


## Restful API Examples
If you use other language to access gdb, you can use resutful interface, before use
you need [build gdb](#build-gdb) or [download](#download-gdb), and after running the application,you can interact with 
gdb by the restful interface easily.Here is the examples of JS(ES6).For more details see
[document](https://justkeepsilence.gitbook.io/gdb/)

### Page
```jsx
// userLogin, passWord is md5 of `${passWord}@seu`
axios.post("/page/userLogin", {userName: "admin", passWord: "685a6b21dc732a9702a96e6731811ec9"})
{"code":200,"message":"","data":{"token":"bc947ca95872df7993fb277072eaa12d"}}
// userLogOut
axios.post("/page/userLogOut", {userName: "admin"})
{"code":200,"message":"","data":{"effectedRows":1}}
// getUserInfo
axios.post("/page/getUserInfo", {"name": "admin"})
{"code":200,"message":"","data":{"userName":"admin","role":["super_user"]}}
// getUsers
axios.post("/page/getUsers")
{"code":200,"message":"","data":{"userInfos":[{"id":"1","role":"super_user","userName":"admin"}]}}
// addUsers
axios.post("/page/addUsers", {"name":"seu","role":"common_user","passWord":"685a6b21dc732a9702a96e6731811ec9"})
{"code":200,"message":"","data":{"effectedRows":1}}
// updateUsers
axios.post("/page/updateUsers", {"userName":"seu1","newUserName":"seu1","newPassWord":"685a6b21dc732a9702a96e6731811ec9","newRole":"common_user"})
{"code":200,"message":"","data":{"effectedRows":1}}
// deleteUsers
axios.post("/page/deleteUsers", {"name":"seu"})
{"code":200,"message":"","data":{"effectedRows":1}}
// getLogs
axios.post("/page/getLogs", {"level":"all","startTime":"2021-05-23 14:42:29","endTime":"2021-05-24 14:42:29","startRow":0,"rowCount":10,"name":"admin"})
{"code":200,"message":"","data":{"infos":[{"id":"2","insertTime":"2021-05-24 10:48:09","level":"Error","logMessage":"{\"requestUrl\":\"/page/userLogin\",\"requestMethod\":\"HTTP/1.1\",\"userAgent\":\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.165 Electron/8.3.0 Safari/537.36\",\"requestBody\":\"{\\\"userName\\\":\\\"seu\\\",\\\"passWord\\\":\\\"c3ea4503d60a70c5cf33720b8cf98716\\\"}\",\"remoteAddress\":\"192.168.0.102:61725\",\"message\":\"userNameError: seu\"}","requestUser":""},{"id":"3","insertTime":"2021-05-24 10:48:16","level":"Error","logMessage":"{\"requestUrl\":\"/page/userLogin\",\"requestMethod\":\"HTTP/1.1\",\"userAgent\":\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.165 Electron/8.3.0 Safari/537.36\",\"requestBody\":\"{\\\"userName\\\":\\\"seu\\\",\\\"passWord\\\":\\\"c3ea4503d60a70c5cf33720b8cf98716\\\"}\",\"remoteAddress\":\"192.168.0.102:61725\",\"message\":\"userNameError: seu\"}","requestUser":""},{"id":"4","insertTime":"2021-05-24 10:48:19","level":"Error","logMessage":"{\"requestUrl\":\"/page/userLogin\",\"requestMethod\":\"HTTP/1.1\",\"userAgent\":\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.165 Electron/8.3.0 Safari/537.36\",\"requestBody\":\"{\\\"userName\\\":\\\"seu\\\",\\\"passWord\\\":\\\"9c614c1c9b72324327cd78ed88f956a3\\\"}\",\"remoteAddress\":\"192.168.0.102:61725\",\"message\":\"userNameError: seu\"}","requestUser":""},{"id":"5","insertTime":"2021-05-24 10:48:23","level":"Error","logMessage":"{\"requestUrl\":\"/page/userLogin\",\"requestMethod\":\"HTTP/1.1\",\"userAgent\":\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.165 Electron/8.3.0 Safari/537.36\",\"requestBody\":\"{\\\"userName\\\":\\\"seu1\\\",\\\"passWord\\\":\\\"9c614c1c9b72324327cd78ed88f956a3\\\"}\",\"remoteAddress\":\"192.168.0.102:61725\",\"message\":\"userNameError: seu1\"}","requestUser":""},{"id":"6","insertTime":"2021-05-24 10:48:25","level":"Error","logMessage":"{\"requestUrl\":\"/page/userLogin\",\"requestMethod\":\"HTTP/1.1\",\"userAgent\":\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.165 Electron/8.3.0 Safari/537.36\",\"requestBody\":\"{\\\"userName\\\":\\\"seu1\\\",\\\"passWord\\\":\\\"c3ea4503d60a70c5cf33720b8cf98716\\\"}\",\"remoteAddress\":\"192.168.0.102:61725\",\"message\":\"userNameError: seu1\"}","requestUser":""},{"id":"7","insertTime":"2021-05-24 10:57:45","level":"Error","logMessage":"{\"requestUrl\":\"/page/userLogin\",\"requestMethod\":\"HTTP/1.1\",\"userAgent\":\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.165 Electron/8.3.0 Safari/537.36\",\"requestBody\":\"{\\\"userName\\\":\\\"seu\\\",\\\"passWord\\\":\\\"c3ea4503d60a70c5cf33720b8cf98716\\\"}\",\"remoteAddress\":\"192.168.0.102:61800\",\"message\":\"userNameError: seu\"}","requestUser":""},{"id":"8","insertTime":"2021-05-24 10:57:47","level":"Error","logMessage":"{\"requestUrl\":\"/page/userLogin\",\"requestMethod\":\"HTTP/1.1\",\"userAgent\":\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.165 Electron/8.3.0 Safari/537.36\",\"requestBody\":\"{\\\"userName\\\":\\\"seu\\\",\\\"passWord\\\":\\\"9c614c1c9b72324327cd78ed88f956a3\\\"}\",\"remoteAddress\":\"192.168.0.102:61800\",\"message\":\"userNameError: seu\"}","requestUser":""}],"count":7}}
// deleteLogs
axios.post("/page/deleteLogs", {"id":"8"})
{"code":200,"message":"","data":{"effectedRows":1}}
axios.post("/page/deleteLogs", {"startTime":"2021-05-23 15:12:06","endTime":"2021-05-24 15:12:06","userNameCondition":"requestUser='admin'"})
{"code":200,"message":"","data":{"effectedRows":10}}
// getDbInfo
axios.post("/page/getDbInfo")
{"code":200,"message":"","data":{"info":{"currentTimeStamp":"1621869963","ram":"169.14","speed":"51ms/6137","writtenItems":"6137"}}}
// getDbSpeedHistory
axios.post("/page/getDbInfoHistory", {"itemName":"speed","startTimes":[1621209000],"endTimes":[1621382400],"intervals":[3600]})
{"code":200,"message":"","data":{"historicalData":{"speed":[["1621209000","1621213881","1621218786","1621223692","1621228582","1621233489","1621238397","1621243288","1621248199","1621253118","1621258008","1621262917","1621267818","1621272727","1621277619","1621282529","1621287427","1621292329","1621297233","1621302117","1621307022","1621311938","1621316850","1621321740","1621326637","1621331533","1621336424","1621341332","1621346240","1621351149","1621356048","1621360950","1621365846","1621370744","1621375634","1621380559"],["81.0773ms","107.1021ms","71.0679ms","167.1601ms","56.0528ms","109.1051ms","72.067ms","35.033ms","60.0555ms","49.047ms","114.1066ms","136.1308ms","98.0942ms","159.1516ms","85.082ms","98.0938ms","56.0524ms","98.0942ms","79.0724ms","116.1117ms","181.1737ms","82.08ms","140.1347ms","91.0851ms","144.1385ms","84.0801ms","146.1406ms","120.113ms","52.0503ms","83.0774ms","155.1494ms","103.097ms","159.1507ms","66.0635ms","101.0969ms","49.0471ms"]]}}}
```

### Group
```jsx
// addGroups
axios.post("/group/addGroups", {"groupInfos": [{"groupName": "1DCS", "columnNames": ["description", "unit"]}]})
{"code":200,"message":"","data":{"effectedRows":1}}
// deleteGroups
axios.post("/group/deleteGroups", {"groupNames": ["1DCS"]})
{"code":200,"message":"","data":{"effectedRows":1}}
//getGroups
axios.post("/group/getGroups")
{"code":200,"message":"","data":{"groupNames":["calc","1DCS"]}}
// getGroupProperty
axios.post("/group/getGroupProperty", {"groupName": "1DCS", "condition": "1=1"})
{"code":200,"message":"","data":{"itemCount":"6387","itemColumnNames":["itemName","groupName","dataType","description","unit","source"]}}
// updateGroupNames
axios.post("/group/updateGroupNames", {"infos": [{"oldGroupName": "4DCS", "newGroupName": "5DCS"}]})
{"code":200,"message":"","data":{"effectedRows":1}}
// updateGroupColumnNames
axios.post("/group/updateGroupColumnNames", {"groupName": "1DCS", "newColumnNames": ["unit1"], "oldColumnNames": ["unit"]})
{"code":200,"message":"","data":{"effectedCols":1}}
// deleteGroupColumns
axios.post("/group/deleteGroupColumns", {"groupName": "1DCS", "columnNames": ["unit"]})
{"code":200,"message":"","data":{"effectedCols":1}}
// addGroupColumns
axios.post("/group/addGroupColumns", {"groupName":"5DCS","columnNames":["unit","description"],"defaultValues":["",""]})
{"code":200,"message":"","data":{"effectedCols":2}}
```

### Item
```jsx
// addItems
axios.post("/item/addItems", {"groupName":"5DCS","itemValues":[{"itemName":"item1","dataType":"float64","unit":"","description":""}]})
{"code":200,"message":"","data":{"effectedRows":1}}
// deleteItems
axios.post("/item/deleteItems", {"groupName":"5DCS","condition":"itemName='item1'"})
{"code":200,"message":"","data":{"effectedRows":1}}
// getItemsWithCount
axois.post("/item/getItemsWithCount", {"groupName":"calc","columnNames":"*","condition":"itemName like '%%'","startRow":0,"rowCount":10})
{"code":200,"message":"","data":{"itemCount":1,"itemValues":[{"dataType":"float64","description":"","id":"1","itemName":"item1"}]}}
// updateItems
axios.post("/item/updateItems", {"groupName": "1DCS", "condition": "id=1", "clause": "description=' ',unit='℃1'"})
{"code":200,"message":"","data":{"effectedRows":1}}
// checkItems
axios.post("/item/checkItems", {"groupName": "1DCS", "itemNames": ["NMJL.UNIT2.20ACS:MAG50AN001SV_MA"]})
{"code":500,"message":"itemName: NMJL.UNIT2.20ACS:MAG50AN001SV_MA not existed","data":""}
// cleanGroupItems
axios.post("/item/cleanGroupItems", {"groupNames": ["1DCS"]})
{"code":200,"message":"","data":{"effectedRows":1}}
```

### Data
```jsx
// batchWrite
axios.post("/data/batchWrite", {"itemValues": [{"itemName": "x", "value": 1.0, "groupName": "5DCS"}, {"itemName": "y", "value": 2.0, "groupName": "5DCS"}]})
{"code":200,"message":"","data":{"effectedRows":2}}
// batchWriteHistoricalData
'use strict';
const axios = require('axios')
const ip = "192.168.0.199:8082"
const count = 3600 // one hour
const now = new Date(2021,4,24,19,44,0)
const st = now.getTime() / 1000 + 8 * 3600
let xData = []
let yData = []
let ts = []
for (var i = 0; i < count; i++) {
  const x = Math.floor(Math.random() * count)
  const y = 2 * x
  xData.push(x)
  yData.push(y)
  ts.push(st + i)
}
axios.post(`http://${ip}/data/batchWriteHistoricalData`, { "historicalItemValues": [{ "groupName": "5DCS", "itemName": "x", "values": xData, "timeStamps": ts }, { "groupName": "5DCS", "itemName": "y", "values": yData, "timeStamps": ts }] }).then((data) => {
  console.log(data)
}).catch((err) => {
  console.log(err)
})
// getRealTimeData
axios.post(`http://${ip}/data/getRealTimeData`, { "groupNames": ["5DCS", "5DCS"], "itemNames": ["x", "y"] })
{ code: 200, message: '', data: { realTimeData: { x: 1, y: 2 } } }
axios.post(`http://${ip}/data/getRealTimeData`, { "groupNames": ["5DCS", "5DCS", "calc"], "itemNames": ["x", "y", "z"] })
{"code":200,"message":"","data":{"realTimeData":{"x":1,"y":2,"z":null}}}  // z is in calc group, but does not have realTimeData, so the result is null
// getHistoricalData
// normal 
'use strict';
const axios = require('axios')
const ip = "192.168.0.199:8082"
const now = new Date(2021, 4, 24, 19, 44, 0)
const st = now.getTime() / 1000 + 8 * 3600
axios.post(`http://${ip}/data/getHistoricalData`, { "groupNames": ["5DCS", "5DCS"], "itemNames": ["x", "y"], "startTimes": [st], "endTimes": [st + 3600], "intervals": [60] }).then(({ data }) => {
  console.log(JSON.stringify(data))
}).catch((err) => {
  console.log(err)
})
// some items does not have history in the given st and et
axios.post(`http://${ip}/data/getHistoricalData`, { "groupNames": ["5DCS", "5DCS", "calc"], "itemNames": ["x", "y", "z"], "startTimes": [st], "endTimes": [st + 3600], "intervals": [60] }).then(({ data }) => {
  console.log(JSON.stringify(data))
}).catch((err) => {
  console.log(err)
})  // item z does not havb e history, so the result of z is {"z": [null, null]}
// getHistoricalDataWithStamp
// normal 
'use strict';
const axios = require('axios')
const ip = "192.168.0.199:8082"
const now = new Date(2021, 4, 24, 19, 44, 0)
const st = now.getTime() / 1000 + 8 * 3600
let ts = []
for (var i = 0; i < 60; i++) {
  ts.push(st + i)
}
axios.post(`http://${ip}/data/getHistoricalDataWithStamp`, { "groupNames": ["5DCS", "5DCS"], "itemNames": ["x", "y"], "timeStamps": [ts,ts]}).then(({ data }) => {
  console.log(JSON.stringify(data))
}).catch((err) => {
  console.log(err)
})
// item does not have history in the given timeStamps
'use strict';
const axios = require('axios')
const ip = "192.168.0.199:8082"
const now = new Date(2021, 4, 24, 19, 44, 0)
const st = now.getTime() / 1000 + 8 * 3600
let ts = []
for (var i = 0; i < 60; i++) {
  ts.push(st + i)
}
ts.push(st + 360000)  // not have history
axios.post(`http://${ip}/data/getHistoricalDataWithStamp`, { "groupNames": ["5DCS", "5DCS", "calc"], "itemNames": ["x", "y", "z"], "timeStamps": [ts, ts, ts] }).then(({ data }) => {
  console.log(JSON.stringify(data))
}).catch((err) => {
  console.log(err)
})
{"code":200,"message":"","data":{"historicalData":{"x":[[1621885440,1621885441,1621885442,1621885443,1621885444,1621885445,1621885446,1621885447,1621885448,1621885449,1621885450,1621885451,1621885452,1621885453,1621885454,1621885455,1621885456,1621885457,1621885458,1621885459,1621885460,1621885461,1621885462,1621885463,1621885464,1621885465,1621885466,1621885467,1621885468,1621885469,1621885470,1621885471,1621885472,1621885473,1621885474,1621885475,1621885476,1621885477,1621885478,1621885479,1621885480,1621885481,1621885482,1621885483,1621885484,1621885485,1621885486,1621885487,1621885488,1621885489,1621885490,1621885491,1621885492,1621885493,1621885494,1621885495,1621885496,1621885497,1621885498,1621885499,1622245440],[2389,2864,1212,1854,1877,1819,951,2781,1815,2058,2533,1508,2544,15,499,377,2120,471,577,2327,1965,217,1681,2465,2025,1401,3524,1520,3521,28,3071,112,2575,3024,1162,3032,627,659,2691,1585,3113,229,2169,396,1888,2989,467,1870,3015,2635,619,3285,510,2089,3159,733,429,1102,2077,986,null]],"y":[[1621885440,1621885441,1621885442,1621885443,1621885444,1621885445,1621885446,1621885447,1621885448,1621885449,1621885450,1621885451,1621885452,1621885453,1621885454,1621885455,1621885456,1621885457,1621885458,1621885459,1621885460,1621885461,1621885462,1621885463,1621885464,1621885465,1621885466,1621885467,1621885468,1621885469,1621885470,1621885471,1621885472,1621885473,1621885474,1621885475,1621885476,1621885477,1621885478,1621885479,1621885480,1621885481,1621885482,1621885483,1621885484,1621885485,1621885486,1621885487,1621885488,1621885489,1621885490,1621885491,1621885492,1621885493,1621885494,1621885495,1621885496,1621885497,1621885498,1621885499,1622245440],[4778,5728,2424,3708,3754,3638,1902,5562,3630,4116,5066,3016,5088,30,998,754,4240,942,1154,4654,3930,434,3362,4930,4050,2802,7048,3040,7042,56,6142,224,5150,6048,2324,6064,1254,1318,5382,3170,6226,458,4338,792,3776,5978,934,3740,6030,5270,1238,6570,1020,4178,6318,1466,858,2204,4154,1972,null]],"z":[[1621885440,1621885441,1621885442,1621885443,1621885444,1621885445,1621885446,1621885447,1621885448,1621885449,1621885450,1621885451,1621885452,1621885453,1621885454,1621885455,1621885456,1621885457,1621885458,1621885459,1621885460,1621885461,1621885462,1621885463,1621885464,1621885465,1621885466,1621885467,1621885468,1621885469,1621885470,1621885471,1621885472,1621885473,1621885474,1621885475,1621885476,1621885477,1621885478,1621885479,1621885480,1621885481,1621885482,1621885483,1621885484,1621885485,1621885486,1621885487,1621885488,1621885489,1621885490,1621885491,1621885492,1621885493,1621885494,1621885495,1621885496,1621885497,1621885498,1621885499,1622245440],[null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null]]}}}
// getHistoricalDataWithCondition
/// normal 
'use strict';
const axios = require('axios')
const fs = require('fs')
const path = require('path')
const ip = "192.168.0.199:8082"
const now = new Date(2021, 4, 24, 19, 44, 0)
const st = now.getTime() / 1000 + 8 * 3600
axios.post(`http://${ip}/data/getHistoricalDataWithCondition`, { "groupNames": ["5DCS", "5DCS"], "itemNames": ["x", "y"], "startTimes": [st], "endTimes": [st + 3600], intervals: [10], "filterCondition": `item["x"] > 2000 && item["y"] > 1000` }).then(({ data }) => {
  fs.writeFile(path.resolve(__dirname, './fX.txt'), JSON.stringify(data), err => { })
  console.log(JSON.stringify(data))
}).catch((err) => {
  console.log(err)
})
// item does not have history
'use strict';
const axios = require('axios')
const fs = require('fs')
const path = require('path')
const ip = "192.168.0.199:8082"
const now = new Date(2021, 6, 24, 19, 44, 0)
const st = now.getTime() / 1000 + 8 * 3600
axios.post(`http://${ip}/data/getHistoricalDataWithCondition`, { "groupNames": ["5DCS", "5DCS", "calc"], "itemNames": ["x", "y", "z"], "startTimes": [st], "endTimes": [st + 3600], intervals: [10], "filterCondition": `item["x"] > 2000 && item["y"] > 1000` }).then(({ data }) => {
  fs.writeFile(path.resolve(__dirname, './fX.txt'), JSON.stringify(data), err => { })
  console.log(JSON.stringify(data))
}).catch((err) => {
  console.log(err)
})
{"code":200,"message":"","data":{"historicalData":{"x":[null,null],"y":[null,null],"z":[null,null]}}}
```

## gRPC API Examples
gdb support gRPC both in http and https mode.In https mode, you need to provide a certificate corresponding to the configuration file(config.json).In authorization
mode you also need to set authorization field in gRPC metaData.Here are examples of how to connect gdb client in go with gRPC.For more details you can see [serviceExamples](https://github.com/JustKeepSilence/gdb/blob/master/examples/service_examples.go)
or about how to connect gdb in gRPC mode in NodeJs, you can see [gdbUI](https://github.com/JustKeepSilence/gdbUI/blob/master/src/renderer/api/index.js)

```go
  if cred, err := credentials.NewClientTLSFromFile("./ssl/gdbServer.crt", "gdb.com");err!=nil{
    log.Fatal(err)
  }else{
    if conn, err := grpc.Dial(ip, grpc.WithTransportCredentials(cred));err!=nil{
        log.Fatal(err)
    }else{
        client := model.NewPageClient(conn)
        if r, err := client.UserLogin(context.Background(), &model.AuthInfo{
            UserName: "admin",
            PassWord: "685a6b21dc732a9702a96e6731811ec9",
        });err!=nil{
            log.Fatal(err)
        }else{
            fmt.Println(r.GetToken())
        }
    }
}
```

## Desktop Application
see [gdbUI](https://github.com/JustKeepSilence/gdbUI) for more details 

## FAQ
1. How to obtain the timeStamp consistent with gdb
```go
# The timestamp in gdb uses the unix timestamp that comes with go,timeZone is UTC
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
```js
new Date(2021, 1, 11 , 14, 26, 26).getTime()/1000 + 8 * 3600
```
### Python(In China)
```python
from datetime import datetime
int(datetime(2021, 2, 11, 14, 26, 26).timestamp()) + 8 * 3600
```