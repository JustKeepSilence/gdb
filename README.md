# GDB
GDB is a real-time database encapsulated based on [goleveldb](https://pkg.go.dev/github.com/syndtr/goleveldb/leveldb)
it can be used to obtain and store large amount of historical data,it provides rest interface and web client , and it 
allows you to generate your own data based on existing data by coding js on web client.If you need deal with big data,
you will love GDB.

## Installation
Gdb is a cgo project, to run or build it, you need gcc ,and install [GO](https://golang.org/) (**version 1.16+ is required**), then set GO111MODULE=ON
```sh
go get github.com/JustKeepSilence/gdb
```

Then import gdb in your own code
```go
import "github.com/JustKeepSilence/gdb/db"
```

## Quick Start
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

## Build GDB
If you want to use gdb with other language or web application, you need build GDB.
First you need change directory to gdb\main, and build with the following command.
```sh
go build -o ../db.exe main.go
```
you can custom your own configs in config.json.
```json
{
  "ip": "", 
  "port": 9000,
  "dbPath": "./leveldb",
  "itemDbPath": "./itemDb",
  "applicationName": "db.exe"
}
```
```gotemplate
ip: ip of gdb web application and restful interface,if empty string, gdb will get ip of local machine
port: port of gdb web application and restful interface, default 9000
dbPath: path of leveldb to store data
itemDbPath: path of SQLite to store items
applicationName: name of gdb application, we need this to trace the running info of porgram
```
Notes: you need set db.exe,config.json, and dist folder in the same path to sure gdb work normally.

## Restful 
If you use other language to access gdb, you can use resutful interface, before use
you need [build gdb](https://github.com/JustKeepSilence/gdb#build-gdb), and after running the application,you can interact with 
gdb by the restful interface easily.Here is the examples of python.For more details see
[document](https://justkeepsilence.gitbook.io/gdb/)
```python
import json
import time
from itertools import repeat

import requests


if __name__ == "__main__":
    ip = "http://192.168.1.2:9000"
    # add groups
    group_infos = [{"groupName": "1DCS", "columnNames": "groupName,type,description,unit,source".split(",")}]
    requests.post(url=f"{ip}/group/addGroups", data=json.dumps(group_infos, ensure_ascii=False))
    # add items
    items = [{"itemName": "testItem1", "groupName": "1DCS", "type": "","description": "testItem1", "unit": "", "source": ""},
             {"itemName": "testItem2", "type": "","groupName": "1DCS", "description": "testItem2", "unit": "", "source": ""},
             {"itemName": "testItem3", "type": "","groupName": "1DCS", "description": "testItem3", "unit": "", "source": ""},
             {"itemName": "testItem4", "type": "","groupName": "1DCS", "description": "testItem4", "unit": "", "source": ""},
             {"itemName": "testItem5", "type": "","groupName": "1DCS", "description": "testItem5", "unit": "", "source": ""},
             {"itemName": "testItem6", "type": "","groupName": "1DCS", "description": "testItem6", "unit": "", "source": ""},
             {"itemName": "testItem7", "type": "","groupName": "1DCS", "description": "testItem7", "unit": "", "source": ""},
             {"itemName": "testItem8", "type": "","groupName": "1DCS", "description": "testItem8", "unit": "", "source": ""}]
    requests.post(f"{ip}/item/addItems", data=json.dumps({"groupName": "1DCS", "values": items}))
    # write realTime data without timeStamps
    item_values = [{"itemName": f"testItem{i}", "value": str(i * 100)} for i in range(1, 6)]
    requests.post(url=f"{ip}/data/batchWrite", data=json.dumps({"groupName": "1DCS", "itemValues": item_values, "withTimeStamp": False}))
    # write data with timestamp
    t = int(time.time())
    item_values = [{"itemName": f"testItem{i}", "value": str(i * 100), "timeStamp": str(t)} for i in range(6, 9)]
    requests.post(url=f"{ip}/data/batchWrite", data=json.dumps({"groupName": "1DCS", "itemValues": item_values, "withTimeStamp": True}))
    # get realTime data, return the latest updated data
    requests.post(url=f"{ip}/data/getRealTimeData", data=json.dumps({"itemNames": [f"testItem{i}" for i in range(1, 9)]}))
    # get historical data with timestamp
    requests.post(url=f"{ip}/data/getHistoricalDataWithStamp",
                  data=json.dumps({"itemNames": [f"testItem{i}" for i in range(6, 9)],
                                   "timeStamps": list(repeat([t], 3))}))

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