/*
createTime: 2021/5/5
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"time"
)

// examples

func ExampleGdb_AddGroups() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	// you can not add itemName column
	if _, err := gdb.AddGroups(AddedGroupInfo{
		GroupName:   "1DCS",
		ColumnNames: []string{"description"},
	}); err != nil {
		log.Fatal(err)
	}
}

func ExampleGdb_DeleteGroups() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if _, err := gdb.DeleteGroups(GroupNamesInfo{GroupNames: []string{"1DCS"}}); err != nil {
		log.Fatal(err)
	}
}

func ExampleGdb_GetGroups() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if r, err := gdb.GetGroups(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("update group name: ", r.GroupNames)
	}
}

func ExampleGdb_GetGroupProperty() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if r, err := gdb.GetGroupProperty("1DCS", "1=1"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("get group property: ", r.ItemColumnNames, r.ItemCount)
	}
}

func ExampleGdb_UpdateGroupNames() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if _, err := gdb.UpdateGroupNames(UpdatedGroupNameInfo{
		OldGroupName: "1DCS",
		NewGroupName: "2DCS",
	}); err != nil {
		log.Fatal(err)
	} else {
		if r, err := gdb.GetGroups(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("update group name: ", r.GroupNames)
		}
	}
}

func ExampleGdb_UpdateGroupColumnNames() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if _, err := gdb.UpdateGroupColumnNames(UpdatedGroupColumnNamesInfo{
		GroupName:      "2DCS",
		OldColumnNames: []string{"Column1", "Column2"},
		NewColumnNames: []string{"Column11", "Column22"},
	}); err != nil {
		log.Fatal(err)
	} else {
		if r, err := gdb.GetGroupProperty("2DCS", "1=1"); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("update group column name:", r.ItemColumnNames)
		}
	}
}

func ExampleGdb_DeleteGroupColumns() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if _, err := gdb.DeleteGroupColumns(DeletedGroupColumnNamesInfo{
		GroupName:   "2DCS",
		ColumnNames: []string{"Column3"},
	}); err != nil {
		log.Fatal(err)
	} else {
		if r, err := gdb.GetGroupProperty("2DCS", "1=1"); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("delete group columns:", r.ItemColumnNames)
		}
	}
}

func ExampleGdb_AddGroupColumns() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if _, err := gdb.AddGroupColumns(AddedGroupColumnsInfo{
		GroupName:     "2DCS",
		ColumnNames:   []string{"Column3", "Column4"},
		DefaultValues: []string{"", ""},
	}); err != nil {
		log.Fatal(err)
	} else {
		if r, err := gdb.GetGroupProperty("2DCS", "1=1"); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("add group columns:", r.ItemColumnNames)
		}
	}
}

func ExampleGdb_AddItems() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	// By default group  has columns itemName , description and dataType
	if _, err := gdb.AddItems(AddedItemsInfo{
		GroupName: "1DCS",
		ItemValues: []map[string]string{{"itemName": "x", "description": "x", "dataType": "float64"},
			{"itemName": "y", "description": "y", "dataType": "float64"}, {"itemName": "z", "description": "z", "dataType": "float64"},
			{"itemName": "item1", "description": "item1", "dataType": "float64"}, {"itemName": "item2", "description": "item2", "dataType": "float64"}},
	}); err != nil {
		log.Fatal(err)
	}
}

func ExampleGdb_GetItems() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if r, err := gdb.GetItems(ItemsInfo{
		GroupName:   "1DCS",
		Condition:   "1=1",
		ColumnNames: "*",
		StartRow:    -1,
		RowCount:    0,
	}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(r.ItemValues)
	}
}

func ExampleGdb_DeleteItems() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if _, err := gdb.DeleteItems(DeletedItemsInfo{
		GroupName: "1DCS",
		Condition: "itemName like '%item%'",
	}); err != nil {
		log.Fatal(err)
	}
}

func ExampleGdb_UpdateItems() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if _, err := gdb.UpdateItems(UpdatedItemsInfo{
		GroupName: "1DCS",
		Clause:    "description='x1'",
		Condition: "itemName='x'",
	}); err != nil {
		log.Fatal(err)
	}
}

func ExampleGdb_CheckItems() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if err := gdb.CheckItems("1DCS", "x", "x1"); err != nil {
		fmt.Println(err)
	}
}

func ExampleGdb_CleanGroupItems() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if _, err := gdb.CleanGroupItems("1DCS"); err != nil {
		log.Fatal(err)
	}
}

func ExampleGdb_BatchWrite() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if _, err := gdb.BatchWrite([]ItemValue{{ItemName: "x", Value: 1.0, GroupName: "1DCS"}, {ItemName: "y", Value: 2.0, GroupName: "1DCS"}}...); err != nil {
		log.Fatal(err)
	}
}

func ExampleGdb_BatchWriteHistoricalData() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	var xData, yData []interface{}
	var ts []int
	now := time.Now()
	fmt.Println("now: ", now.Format("2006-01-02 15:04:05"))
	r := rand.New(rand.NewSource(99))
	for i := 0; i < 3600; i++ {
		x := float64(r.Intn(3600)) * math.Pi
		y := 2 * x
		t := now.Add(time.Second*time.Duration(i)).Unix() + 8*3600
		xData = append(xData, x)
		yData = append(yData, y)
		ts = append(ts, int(t))
	}
	if err := gdb.BatchWriteHistoricalData([]HistoricalItemValue{{ItemName: "x", Values: xData, TimeStamps: ts, GroupName: "1DCS"}, {ItemName: "y", Values: yData, TimeStamps: ts, GroupName: "1DCS"}}...); err != nil {
		log.Fatal(err)
	}
}

func ExampleGdb_GetRealTimeData() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if r, err := gdb.GetRealTimeData([]string{"1DCS", "1DCS"}, "x", "y"); err != nil {
		log.Fatal(err)
	} else {
		d, _ := json.Marshal(r)
		fmt.Println(string(d))
	}
}

func ExampleGdb_GetHistoricalData() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	now := time.Now()
	stX := int(now.Add(time.Minute*5).Unix() + 8*3600)
	etX := int(now.Add(time.Minute*25).Unix() + 8*3600)
	stY := int(now.Add(time.Minute*35).Unix() + 8*3600)
	etY := int(now.Add(time.Minute*55).Unix() + 8*3600)
	// Obtain the historical data of x and y in startTime stX, endTime etX, with an interval of 2s,
	// and startTime as stY, endTime as etY, and an interval of 10s.
	if r, err := gdb.GetHistoricalData([]string{"1DCS", "1DCS"}, []string{"x", "y"}, []int{stX, stY}, []int{etX, etY}, []int{2, 10}); err != nil {
		log.Fatal(err)
	} else {
		d, _ := json.Marshal(r)
		_ = ioutil.WriteFile("./hX.txt", d, 0644)
	}
}

func ExampleGdb_GetRawHistoricalData() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	if r, err := gdb.GetRawHistoricalData([]string{"1DCS"}, "x"); err != nil {
		log.Fatal(err)
	} else {
		d, _ := json.Marshal(r)
		_ = ioutil.WriteFile("./rawX.txt", d, 0644)
	}
}

func ExampleGdb_GetHistoricalDataWithStamp() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	now := time.Now()
	stX := int(now.Add(time.Minute*5).Unix() + 8*3600)
	etX := int(now.Add(time.Minute*25).Unix() + 8*3600)
	stY := int(now.Add(time.Minute*35).Unix() + 8*3600)
	etY := int(now.Add(time.Minute*55).Unix() + 8*3600)
	// Obtain the historical data of x with timeStamp stX, etX, y with timeStamp stY, etY
	if r, err := gdb.GetHistoricalDataWithStamp([]string{"1DCS", "1DCS"}, []string{"x", "y"}, [][]int{{stX, etX}, {stY, etY}}...); err != nil {
		log.Fatal(err)
	} else {
		d, _ := json.Marshal(r)
		fmt.Println(string(d))
	}
}

func ExampleGdb_GetHistoricalDataWithCondition() {
	gdb, _ := NewGdb("./leveldb", "./itemDb")
	now := time.Now()
	stX := int(now.Add(time.Minute*5).Unix() + 8*3600)
	etX := int(now.Add(time.Minute*25).Unix() + 8*3600)
	stY := int(now.Add(time.Minute*35).Unix() + 8*3600)
	etY := int(now.Add(time.Minute*55).Unix() + 8*3600)
	// Obtain the historical data of x and y in startTime stX, endTime etX, with an interval of 2s,
	// and startTime as stY, endTime as etY, and an interval of 10s under the given condition and deadZone
	if r, err := gdb.GetHistoricalDataWithCondition([]string{"1DCS", "1DCS"}, []string{"x", "y"}, []int{stX, stY}, []int{etX, etY}, []int{2, 10}, `item["x"] > 0 && item["y"] > 1000`, []DeadZone{}...); err != nil {
		log.Fatal(err)
	} else {
		d, _ := json.Marshal(r)
		_ = ioutil.WriteFile("./f.txt", d, 0644)
	}
}
