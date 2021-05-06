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
        if _, err := gdb.AddItems(db.AddedItemsInfo{
            GroupName:  "1DCS",
            ItemValues: []map[string]string{{"itemName": "x", "description": "x"}, {"itemName": "y", "description": "y"}, {"itemName": "z", "description": "z"},
            {"itemName": "item1", "description": "item1"}, {"itemName": "item2", "description": "item2"}},
            });err!=nil{
            log.Fatal(err)
        }
        // batch write, y = 2 * x
        if _, err := gdb.BatchWrite([]db.ItemValue{{ItemName: "x", Value: "1"}, {ItemName: "y", Value: "2"}}...);err!=nil{
        log.Fatal(err)
        }else{
        // get latest updated value of given items
            if r, err := gdb.GetRealTimeData("x", "y");err!=nil{
            log.Fatal(err)
        }else{
            d, _ := json.Marshal(r)
            fmt.Println(string(d))
        }
        }
        // write historical data
        // mock one hour historical data
        var xData, yData, ts []string
        now := time.Now()
        fmt.Println("now: ", now.Format("2006-01-02 15:04:05"))
        r := rand.New(rand.NewSource(99))
        for i := 0; i < 3600; i++ {
            x := r.Intn(3600)
            y := 2 * x
            t := now.Add(time.Second * time.Duration(i)).Unix() + 8 * 3600
            xData = append(xData, fmt.Sprintf("%d", x))
            yData = append(yData, fmt.Sprintf("%d", y))
            ts = append(ts, fmt.Sprintf("%d", t))
        }
        if err := gdb.BatchWriteHistoricalData([]db.HistoricalItemValue{{ItemName: "x", Values: xData, TimeStamps: ts}, {ItemName: "y", Values: yData, TimeStamps: ts}}...);err!=nil{
        log.Fatal(err)
        }else{
            // get raw historical data for debugging
            if r, err := gdb.GetRawHistoricalData("x");err!=nil{
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
        if r, err := gdb.GetHistoricalData([]string{"x", "y"}, []int{stX, stY}, []int{etX, etY}, []int{2, 10});err!=nil{
        log.Fatal(err)
        }else{
        d, _ := json.Marshal(r)
        _ = ioutil.WriteFile("./hX.txt", d, 0644)
        }
        // get historical data with given itemName
        if r, err := gdb.GetHistoricalDataWithStamp([]string{"x", "y"}, [][]int{{stX, etX}, {stY, etY}}...);err!=nil{
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
the download url is: https://wws.lanzous.com/iHt95nkegha, download passWord is 7bv4

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
axios.post("/page/userLogin", JSON.stringify({userName: "admin", passWord: "685a6b21dc732a9702a96e6731811ec9"}))
{"code":200,"message":"","data":{"token":"bc947ca95872df7993fb277072eaa12d"}}
// getUserInfo
axios.post("/page/getUserInfo", JSON.stringify({"name": "admin"}))
{"code":200,"message":"","data":{"name":"admin","role":["super_user"]}}
// userLogOut
axios.get("/page/userLogOut/admin")
{"code":200,"message":"","data":{"effectedRows":1}}
// upload excel file, you need to set Content-Type to multipart/form-data
const data = new FormData()
data.append('file', fileContent)    // field must be 'file'
axios({url: "/page/uploadFile", headers:{"Content-Type": "multipart/form-data"}, data, method: "post"})
// addItemsByExcel
axios.post("/page/addItemsByExcel", JSON.stringify({"fileName": "item.xlsx", "groupName": "1DCS"}))
{"code":200,"message":"","data":{"effectedRows":18112}}
// importHistoryByExcel
axios.post("/page/importHistoryByExcel", JSON.stringify({"fileName": "data.xlsx", "itemNames": ["YFLDY"], "sheetNames": ["Sheet1"]}))
{"code":200,"message":"","data":""}
// getDbInfo
axios.post("/page/getDbInfo")
{"code":200,"message":"","data":{"info":{"currentTimeStamp":null,"ram":"30.44","speed":"0ms/0","writtenItems":"0"}}}
// getDbSpeedHistory
axios.post("/page/getDbInfoHistory", JSON.stringify({"starTimes": [1617861409547], "endTimes": [1617862009547], "intervale": 5}))
{"code":200,"message":"","data":{"historicalData":{"speed":[null,null]}}}
```

### Group
```jsx
// addGroups
axios.post("/group/addGroups", JSON.stringify({"groupInfos": [{"groupName": "1DCS", "columnNames": ["description", "unit"]}]}))
// deleteGroups
axios.post("/group/deleteGroups", JSON.stringify({"groupNames": ["1DCS"]}))
//getGroups
axios.post("/group/getGroups")
{"code":200,"message":"","data":{"groupNames":["calc","1DCS"]}}
// getGroupProperty
axios.post("/group/getGroupProperty", JSON.stringify({"groupName": "1DCS", "condition": "1=1"}))
{"code":200,"message":"","data":{"itemCount":"6387","itemColumnNames":["itemName","groupName","dataType","description","unit","source"]}}
// updateGroupColumnNames
axios.post("/group/updateGroupColumnNames", JSON.stringify({"groupName": "1DCS", "newColumnNames": ["unit1"], "oldColumnNames": ["unit"]}))
{"code":200,"message":"","data":{"effectedCols":1}}
// deleteGroupColumns
axios.post("/group/deleteGroupColumns", JSON.stringify({"groupName": "1DCS", "columnNames": ["unit"]}))
{"code":200,"message":"","data":{"effectedCols":1}}
// addGroupColumns
axios.post("/group/addGroupColumns", JSON.stringify({"groupName": "1DCS", "columnNames": ["unit"], "defaultValues": [""]}))
{"code":200,"message":"","data":{"effectedCols":1}}
// cleanGroupItems
axios.post("/group/cleanGroupItems", JSON.stringify({"groupNames": ["1DCS"]}))
```

### Item
```jsx
// addItems
axios.post("/item/addItems", JSON.stringify({"groupName": "1DCS", "itemValues": [{"itemName": "YFJDY", "description": "", "unit": ""}]}))
{"code":200,"message":"","data":{"effectedRows":1}}
// deleteItems
axios.post("/item/deleteItems", JSON.stringify({"groupName": "1DCS", "condition": "itemName='YFJDY'"}))
{"code":200,"message":"","data":{"effectedRows":1}}
// getItems
axois.post("/item/getItems", JSON.stringify({"columnNames": "*", "condition": "itemName like '%%'", "groupName": "1DCS", "rowCount": 10, startRow: 0}))
{"code":200,"message":"","data":{"itemValues":[{"dataType":"5","description":"","groupName":"1DCS","id":"1","itemName":"JL1_TURBTEST1:PTJ.OUT","source":"WP1007/TURBTEST1:PTJ.OUT","unit":""},{"dataType":"5","description":"","groupName":"1DCS","id":"2","itemName":"JL1_1ETS:L009_8.BO01","source":"WP1007/1ETS:L009_8.BO01","unit":""},{"dataType":"5","description":"","groupName":"1DCS","id":"3","itemName":"JL1_1ETS:L001_8.BO01","source":"WP1007/1ETS:L001_8.BO01","unit":""},{"dataType":"5","description":"","groupName":"1DCS","id":"4","itemName":"JL1_1OA2:L128_13.BO01","source":"WP1007/1OA2:L128_13.BO01","unit":""},{"dataType":"5","description":"","groupName":"1DCS","id":"5","itemName":"JL1_1OA2:L128_2.BO01","source":"WP1007/1OA2:L128_2.BO01","unit":""},{"dataType":"5","description":"","groupName":"1DCS","id":"6","itemName":"JL1_1OA2:L128_16.BO01","source":"WP1007/1OA2:L128_16.BO01","unit":""},{"dataType":"5","description":"","groupName":"1DCS","id":"7","itemName":"JL1_1ETS:L001_4.BO01","source":"WP1007/1ETS:L001_4.BO01","unit":""},{"dataType":"5","description":"","groupName":"1DCS","id":"8","itemName":"JL1_1ETS:L001_1.BO01","source":"WP1007/1ETS:L001_1.BO01","unit":""},{"dataType":"5","description":"","groupName":"1DCS","id":"9","itemName":"JL1_1ETS:GEN_TRIP.COUT","source":"WP1007/1ETS:GEN_TRIP.COUT","unit":""},{"dataType":"5","description":"","groupName":"1DCS","id":"10","itemName":"JL1_1OA2:L128_6.BO01","source":"WP1007/1OA2:L128_6.BO01","unit":""}]}}
// getItemsWithCount
axois.post("/item/getItemsWithCount")
{"code":200,"message":"","data":{"itemCount":6387,"itemValues":[{"dataType":"5","description":"汽机第一级压力","groupName":"1DCS","id":"1","itemName":"JL1_TURBTEST1:PTJ.OUT","source":"WP1007/TURBTEST1:PTJ.OUT","unit":""},{"dataType":"5","description":"汽机已复置","groupName":"1DCS","id":"2","itemName":"JL1_1ETS:L009_8.BO01","source":"WP1007/1ETS:L009_8.BO01","unit":""},{"dataType":"5","description":"ETS在线试验#3通道","groupName":"1DCS","id":"3","itemName":"JL1_1ETS:L001_8.BO01","source":"WP1007/1ETS:L001_8.BO01","unit":""},{"dataType":"5","description":"阀切换过程中","groupName":"1DCS","id":"4","itemName":"JL1_1OA2:L128_13.BO01","source":"WP1007/1OA2:L128_13.BO01","unit":""},{"dataType":"5","description":"顺序阀","groupName":"1DCS","id":"5","itemName":"JL1_1OA2:L128_2.BO01","source":"WP1007/1OA2:L128_2.BO01","unit":""},{"dataType":"5","description":"顺序阀运行","groupName":"1DCS","id":"6","itemName":"JL1_1OA2:L128_16.BO01","source":"WP1007/1OA2:L128_16.BO01","unit":""},{"dataType":"5","description":"ETS在线试验#1通道","groupName":"1DCS","id":"7","itemName":"JL1_1ETS:L001_4.BO01","source":"WP1007/1ETS:L001_4.BO01","unit":""},{"dataType":"5","description":"ETS在线试验进入试验","groupName":"1DCS","id":"8","itemName":"JL1_1ETS:L001_1.BO01","source":"WP1007/1ETS:L001_1.BO01","unit":""},{"dataType":"5","description":"发电机遮断","groupName":"1DCS","id":"9","itemName":"JL1_1ETS:GEN_TRIP.COUT","source":"WP1007/1ETS:GEN_TRIP.COUT","unit":""},{"dataType":"5","description":"单阀","groupName":"1DCS","id":"10","itemName":"JL1_1OA2:L128_6.BO01","source":"WP1007/1OA2:L128_6.BO01","unit":""}]}}
// updateItems
axios.post("/item/updateItems", JSON.stringify({"groupName": "1DCS", "condition": "id=1", "clause": "description=' ',unit='℃1'"}))
{"code":200,"message":"","data":{"effectedRows":1}}
// checkItems
axios.post("/item/checkItems", JSON.stringify({"groupName": "1DCS", "itemNames": ["NMJL.UNIT2.20ACS:MAG50AN001SV_MA"]}))
{"code":500,"message":"itemName: NMJL.UNIT2.20ACS:MAG50AN001SV_MAnot existed","data":""}
```

### Data
```jsx
// batchWrite
axios.post("/data/batchWrite", JSON.stringify({"itemValues": [{"itemName":"NMJL.UNIT2.20FSSS21B:HFY24AP002ZD","value":"382"},{"itemName":"NMJL.UNIT2.20SCS10B:MAL10AA566ZC","value":"97"},{"itemName":"NMJL.UNIT2.20ACS:MAG50AN001SV_MA","value":"314"},{"itemName":"NMJL.UNIT2.20SCS08B:LAV20AP001NA","value":"-326"},{"itemName":"NMJL.UNIT2.20ECSCOMM_3A:20BTM01YC28","value":"224"},{"itemName":"FZSIS.2DCSAI.20DAS05A:LAB10DP101.PNT","value":"-266"},{"itemName":"FZSIS.2DCSAI.20DAS05A:LAB10CP101.PNT","value":"253"},{"itemName":"FZSIS.2DCSAI.20DAS05A:LAB11DP101.PNT","value":"80"},{"itemName":"FZSIS.2DCSAI.20DAS05A:LAV10CE101.PNT","value":"-70"},{"itemName":"FZSIS.2DCSAI.20DAS05A:LAF11CP101.PNT","value":"21"},{"itemName":"FZSIS.2DCSAI.20MCS07B:LAH10AA101GT.PNT","value":"-348"},{"itemName":"FZSIS.2DCSAI.20MCS07A:LAB12CP101.PNT","value":"-328"},{"itemName":"FZSIS.2DCSAI.20DAS05A:LAC10CE101.PNT","value":"143"},{"itemName":"FZSIS.2DCSAI.20MCS07A:LAB11CF101.PNT","value":"143"},{"itemName":"FZSIS.2DCSAI.20DAS05A:LAK11CS101.PNT","value":"-489"},{"itemName":"FZSIS.2DCSAI.20DAS05B:MKF10CE001.PNT","value":"86"}]}))
{"code":200,"message":"","data":{"effectedRows":1}}
// batchWriteHistoricalData
axios.post("/data/batchWriteHistoricalData", JSON.stringify({"historicalItemValues":[{"itemName":"YFJDY","values":["57.760","57.811","57.801","57.781","57.801","57.807","57.773","57.523","57.508"],"timeStamps":["1602347580","1602347581","1602347582","1602347583","1602347584","1602347585","1602347586","1602347587","1602347588"]}]}))
{"code":200,"message":"","data":""}
// getRealTimeData
axios.post("/data/getRealTimeData", JSON.stringify({"groupName": "1DCS", "itemNames": ["testItem1", "testItem2"]))
{"code":200,"message":"","data":{"realTimeData":{"testItem1": "10", "testItem2": "20"}}}
// getHistoricalData
axios.post("/data/getHistoricalData", JSON.stringify({"itemNames":["NMJL.UNIT2.20ACS:MAG50AN001SV_MA"],"startTimes":[1618843574],"endTimes":[1618929974],"intervals":[60]}))
{"code":200,"message":"","data":{"historicalData":{"NMJL.UNIT2.20ACS:MAG50AN001SV_MA":[[1618921777,1618922077,1618922257,1618922437,1618923817,1618924657,1618924837,1618925137,1618925317,1618925497,1618925677,1618926577,1618926937,1618927177,1618927357,1618927537,1618927897,1618928197,1618928377,1618928557,1618928737,1618928917,1618929397,1618929577,1618929757,1618929937],["-218","-250","-72","387","-319","-438","-156","139","-81","-124","-251","218","224","-215","-94","-148","440","38","-78","-418","-59","-275","-279","83","-96","478"]]}}}
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