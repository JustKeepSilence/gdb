## GDB
GDB is a real-time database encapsulated based on [goleveldb](https://pkg.go.dev/github.com/syndtr/goleveldb/leveldb)
it can be used to obtain and store large amount of historical data in various ways(including gettting raw data, filtered data
with given condition, etc...),it provides rest, gRPC interface and web client , and it 
allows you to generate your own data based on existing data by coding js on web client.If you need deal with big data,
you will love GDB.

## Features
- High writing performance
- Fine control of historical data
- Simulate data based on existing data with js
- Token-based permission control
- restful and gRPC api
- support https api
- web application, desktop client

## Contents
- [Quick Start](#quick-start)
- [Installation](#installation)
- [Build GDB](#build-gdb)
- [Run With HTTPS](#run-with-https)
- [Restful API Examples](#restful-api-examples)
    - [Page](#page)
    - [Group](#group)
    - [Item](#item)
    - [Data](#data)
- [gRPC API Examples](#grpc-api-examples)
    - [Page](#page)
    - [Group](#group)
    - [Item](#item)
    - [Data](#data)
- [WebApplication](#web-application)
- [DesktopApplication](#desktop-application)
- [FAQ](#faq)

## Quick Start
### Integration with go language
If you are familiar with go language, you call [install](#installation) gdb and then use it in your
project to customize your own behavior,For more details,you can see [document](https://pkg.go.dev/github.com/JustKeepSilence/gdb),
```go
import (
	"encoding/json"
	"fmt"
	"github.com/JustKeepSilence/gdb/db"
	"log"
	"time"
)

func main() {
	dbPath := "./db"  // path of data
	itemDbPath := "./itemDb"  // path of itemDb
	g, err := db.NewGdb(dbPath, itemDbPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Open db successfully")
	//add groups:
	groupInfos := []db.AddGroupInfo{{
		GroupName:   "1DCS",
		ColumnNames: []string{"groupName", "type", "description", "unit", "source"},  // every group has two cols: id and itemName
	}}
	if _, err := g.AddGroups(groupInfos...);err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("add group successfully")
	}
	//add items
	if _, err  := g.AddItems(db.AddItemInfo{
		GroupName: "1DCS",
		Values: []map[string]string{{"itemName": "testItem1", "groupName": "1DCS", "type": "","description": "testItem1", "unit": "", "source": ""},
			{"itemName": "testItem2", "type": "","groupName": "1DCS", "description": "testItem2", "unit": "", "source": ""},
			{"itemName": "testItem3", "type": "","groupName": "1DCS", "description": "testItem3", "unit": "", "source": ""},
			{"itemName": "testItem4", "type": "","groupName": "1DCS", "description": "testItem4", "unit": "", "source": ""},
			{"itemName": "testItem5", "type": "","groupName": "1DCS", "description": "testItem5", "unit": "", "source": ""},
			{"itemName": "testItem6", "type": "","groupName": "1DCS", "description": "testItem6", "unit": "", "source": ""},
			{"itemName": "testItem7", "type": "","groupName": "1DCS", "description": "testItem7", "unit": "", "source": ""},
			{"itemName": "testItem8", "type": "","groupName": "1DCS", "description": "testItem8", "unit": "", "source": ""}},
	});err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("add items successfully!")
	}
	//add items by excel
	if _, err := g.AddItemsByExcel("1DCS", "./test.xlsx");err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("add items by excel successfully!")
	}
	//write realTime Data without timeStamp
	if _, err := g.BatchWrite(db.BatchWriteString{
		GroupName:     "1DCS",
		ItemValues:    []db.ItemValue{{
			ItemName: "testItem1",
			Value:    "-100",
		},{
			ItemName: "testItem2",
			Value: "0",
		},{
			ItemName: "testItem3",
			Value: "100",
		},{
			ItemName: "testItem4",
			Value: "200",
		},{
			ItemName: "testItem5",
			Value: "300",
		}},
		WithTimeStamp: false,
	});err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("Write successfully")
	}
	// write realTime data without timeStamp
	t := fmt.Sprintf("%d", time.Now().Add(-1 * time.Hour).Unix())  // unix timeStamp
	if _, err := g.BatchWrite(db.BatchWriteString{
		GroupName:     "1DCS",
		ItemValues:    []db.ItemValue{{
			ItemName: "testItem6",
			Value: "400",
			TimeStamp: t,
		},{
			ItemName: "testItem7",
			Value: "500",
			TimeStamp: t,
		},{
			ItemName: "testItem8",
			Value: "600",
			TimeStamp: t,
		}},
		WithTimeStamp: true,
	});err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("Write with timeStamp successfully")
	}
	// get realTime data, return the latest updated data
	itemNames := []string{"testItem1", "testItem2", "testItem3", "testItem4", "testItem5", "testItem6", "testItem7", "testItem8"}
	if c, err := g.GetRealTimeData(itemNames...);err!=nil{
		log.Fatal(err)
	}else{
		r, _ := json.Marshal(c)
		fmt.Println(fmt.Sprintf("%s", r))
	}
	if c, err := g.GetRawHistoricalData(itemNames...);err!=nil{
		log.Fatal(err)
	}else{
		r, _ := json.Marshal(c)
		fmt.Println(fmt.Sprintf("%s", r))
	}
	// get historical data with timeStamp
	timeStamps := [][]int{{1612413561}, {1612413561}, {1612413561}, {1612413561}, {1612413561}}
	if c, err := g.GetHistoricalDataWithStamp([]string{"testItem1", "testItem2", "testItem3", "testItem4", "testItem5"}, timeStamps...);err!=nil{
		log.Fatal(err)
	}else{
		r, _ := json.Marshal(c)
		fmt.Println(fmt.Sprintf("%s", r))
	}
}
```
### Integration with other language
If you are not familiar with go, and want to use gdb as back-end database only, you can [build-gdb](#build-gdb), then run 
gdb in your server.Also, you can [download](https://wws.lanzous.com/iHt95nkegha) the compiled binary file and run it directly.In this way, you can't customize 
your own behavior, but you can use restful or grpc api provided by gdb, as well as token-control for every api we 
provided.For more details you can see [restful-examples](#restful-api-examples)  or [grpc-examples](#grpc-api-examples) 
or [documents](https://app.gitbook.com/@justkeepsilence/s/gdb/~/settings/share)

## Installation
Gdb is a cgo project, to run or build it, you need [gcc](https://gcc.gnu.org/) ,and install [Go](https://golang.org/) (**version 1.15+ is required**), then set GO111MODULE=ON
```sh
go get github.com/JustKeepSilence/gdb
```

Notes: Since gdb uses [gin](https://github.com/gin-gonic/gin#grouping-routes) v1.6.3 internally, you should manually add the gin.CustomRecovery function
for details, see: https://github.com/gin-gonic/gin/issues/2615


Then import gdb in your own code
```go
import "github.com/JustKeepSilence/gdb/db"
```

## Build GDB
If you want to use gdb with other language or web application, you need build GDB.
First you need install gcc and Go, then change directory to gdb\main, and build with
the following command(on windows)
```sh
go build -tags=jsoniter -o ../db.exe main.go
```
you can custom your own configs in config.json.
```json
// Notes: you can use // single line comments in json file
{
  "gdbConfigs": {
    "ip": "",
    "port": 8082,
    "dbPath": "./leveldb",
    "itemDbPath": "./itemDb",
    "applicationName": "db.exe",
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
    "Level" : 1,
    "expiredTime": 86400
  }
}

```

Notes: you need set db.exe,config.json, and dist folder in the same path to sure gdb work normally.

## Run With HTTPS

gdb support https mode for restful nad gRPC,if you want to run with https mode, you need to set 
mode filed in configs.json to https, and custom your own https configs.Then put your own certificate
to ssl folder and put ssl put the same path as gdb executable program.

Notes: selfSignedCa is not allowed on windows at moment.

## Run With Authorization

gdb support token authorization mode, if you want to run with authorization mode,
you need to set authorization field to true.Then you need to add Authorization field to
header if you use restful, or to context if you use gRPC.For more details, you can see
[document](https://blog.csdn.net/qq_39778055/article/details/114844756)


## Restful API Examples
If you use other language to access gdb, you can use resutful interface, before use
you need [build gdb](https://github.com/JustKeepSilence/gdb#build-gdb) or [download](), and after running the application,you can interact with 
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
axios.post("/item/addItems", JSON.stringify({"groupName": "1DCS", "gdbItems": {"itemValues": [{"itemName": "YFJDY", "description": "", "unit": ""}]}))
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
axios.post("/data/getHistoricalData", JSON.stringify({"groupName":"1DCS","itemNames":["NMJL.UNIT2.20ACS:MAG50AN001SV_MA"],"startTimes":[1618843574],"endTimes":[1618929974],"intervals":[60]}))
{"code":200,"message":"","data":{"historicalData":{"NMJL.UNIT2.20ACS:MAG50AN001SV_MA":[[1618921777,1618922077,1618922257,1618922437,1618923817,1618924657,1618924837,1618925137,1618925317,1618925497,1618925677,1618926577,1618926937,1618927177,1618927357,1618927537,1618927897,1618928197,1618928377,1618928557,1618928737,1618928917,1618929397,1618929577,1618929757,1618929937],["-218","-250","-72","387","-319","-438","-156","139","-81","-124","-251","218","224","-215","-94","-148","440","38","-78","-418","-59","-275","-279","83","-96","478"]]}}}
// getDbInfo
axios.post("/data/getDbInfo")
{"code":200,"message":"","data":{"info":{"currentTimeStamp":null,"ram":"30.44","speed":"0ms/0","writtenItems":"0"}}}
// getDbSpeedHistory
axios.post("/data/getDbSpeedHistory", JSON.stringify({"starTimes": [1617861409547], "endTimes": [1617862009547], "intervale": 5}))
{"code":200,"message":"","data":{"speed":[null,null]}}

```

## Web Application
see [web application](https://github.com/JustKeepSilence/gdb-web-app) for more details

## Windows Desktop Application
see [windows ui](https://github.com/JustKeepSilence/gdb-windows-ui) for more details 

## FAQ
1. How to obtain the timeStamp consistent with gdb
```go
# The timestamp in gdb uses the unix timestamp that comes with go,timeZone is UTC
import (
 "time"
)
n := time.Now
timeStamp := time.Date(n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Minute(), 0, time.UTC)
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