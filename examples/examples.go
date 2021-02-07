package main

import (
	"encoding/json"
	"fmt"
	"github.com/JustKeepSilence/gdb/db"
	"log"
	"time"
)

func main() {
	dbPath := "./db"         // path of data
	itemDbPath := "./itemDb" // path of itemDb
	g, err := db.NewGdb(dbPath, itemDbPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Open db successfully")
	//add groups:
	groupInfos := []db.AddGroupInfo{{
		GroupName:   "1DCS",
		ColumnNames: []string{"groupName", "type", "description", "unit", "source"}, // every group has two cols: id and itemName
	}}
	if _, err := g.AddGroups(groupInfos...); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("add group successfully")
	}
	//add items
	if _, err := g.AddItems(db.AddItemInfo{
		GroupName: "1DCS",
		Values: []map[string]string{{"itemName": "testItem1", "groupName": "1DCS", "type": "", "description": "testItem1", "unit": "", "source": ""},
			{"itemName": "testItem2", "type": "", "groupName": "1DCS", "description": "testItem2", "unit": "", "source": ""},
			{"itemName": "testItem3", "type": "", "groupName": "1DCS", "description": "testItem3", "unit": "", "source": ""},
			{"itemName": "testItem4", "type": "", "groupName": "1DCS", "description": "testItem4", "unit": "", "source": ""},
			{"itemName": "testItem5", "type": "", "groupName": "1DCS", "description": "testItem5", "unit": "", "source": ""},
			{"itemName": "testItem6", "type": "", "groupName": "1DCS", "description": "testItem6", "unit": "", "source": ""},
			{"itemName": "testItem7", "type": "", "groupName": "1DCS", "description": "testItem7", "unit": "", "source": ""},
			{"itemName": "testItem8", "type": "", "groupName": "1DCS", "description": "testItem8", "unit": "", "source": ""}},
	}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("add items successfully!")
	}
	//add items by excel
	if _, err := g.AddItemsByExcel("1DCS", "./test.xlsx"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("add items by excel successfully!")
	}
	//write realTime Data without timeStamp
	if _, err := g.BatchWrite(db.BatchWriteString{
		GroupName: "1DCS",
		ItemValues: []db.ItemValue{{
			ItemName: "testItem1",
			Value:    "-100",
		}, {
			ItemName: "testItem2",
			Value:    "0",
		}, {
			ItemName: "testItem3",
			Value:    "100",
		}, {
			ItemName: "testItem4",
			Value:    "200",
		}, {
			ItemName: "testItem5",
			Value:    "300",
		}},
		WithTimeStamp: false,
	}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Write successfully")
	}
	// write realTime data without timeStamp
	t := fmt.Sprintf("%d", time.Now().Add(-1*time.Hour).Unix()) // unix timeStamp
	if _, err := g.BatchWrite(db.BatchWriteString{
		GroupName: "1DCS",
		ItemValues: []db.ItemValue{{
			ItemName:  "testItem6",
			Value:     "400",
			TimeStamp: t,
		}, {
			ItemName:  "testItem7",
			Value:     "500",
			TimeStamp: t,
		}, {
			ItemName:  "testItem8",
			Value:     "600",
			TimeStamp: t,
		}},
		WithTimeStamp: true,
	}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Write with timeStamp successfully")
	}
	// get realTime data, return the latest updated data
	itemNames := []string{"testItem1", "testItem2", "testItem3", "testItem4", "testItem5", "testItem6", "testItem7", "testItem8"}
	if c, err := g.GetRealTimeData(itemNames...); err != nil {
		log.Fatal(err)
	} else {
		r, _ := json.Marshal(c)
		fmt.Println(fmt.Sprintf("%s", r))
	}
	if c, err := g.GetRawHistoricalData(itemNames...); err != nil {
		log.Fatal(err)
	} else {
		r, _ := json.Marshal(c)
		fmt.Println(fmt.Sprintf("%s", r))
	}
	// get historical data with timeStamp
	timeStamps := [][]int{{1612413561}, {1612413561}, {1612413561}, {1612413561}, {1612413561}}
	if c, err := g.GetHistoricalDataWithStamp([]string{"testItem1", "testItem2", "testItem3", "testItem4", "testItem5"}, timeStamps...); err != nil {
		log.Fatal(err)
	} else {
		r, _ := json.Marshal(c)
		fmt.Println(fmt.Sprintf("%s", r))
	}
}
