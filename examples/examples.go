package examples

// examples about how to use gdb in go, for more details you
// can see document:https://pkg.go.dev/github.com/JustKeepSilence/gdb

import (
	"encoding/json"
	"fmt"
	"github.com/JustKeepSilence/gdb/db"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"time"
)

func initialDb() (*db.Gdb, error) {
	if gdb, err := db.NewGdb("./leveldb", "./itemDb"); err != nil {
		return nil, err
	} else {
		return gdb, nil
	}
}

func group() {
	if gdb, err := initialDb(); err != nil {
		log.Fatal(err)
	} else {
		// add groups
		if _, err := gdb.AddGroups(db.AddedGroupInfo{
			GroupName:   "1DCS",
			ColumnNames: []string{"Column1", "Column2", "Column3"}, // column name can't be itemName
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("add group 1DCS successfully")
		}
		//delete groups
		if _, err := gdb.DeleteGroups(db.GroupNamesInfo{GroupNames: []string{"1DCS"}}); err != nil {
			log.Fatal(err)
		}
		// get groups
		if r, err := gdb.GetGroups(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("get groups:", r.GroupNames)
		}
		// get group property
		if r, err := gdb.GetGroupProperty("1DCS", "1=1"); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("get group property: ", r.ItemColumnNames, r.ItemCount)
		}
		// update group name
		if _, err := gdb.UpdateGroupNames(db.UpdatedGroupNameInfo{
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
		// update group column name
		if _, err := gdb.UpdateGroupColumnNames(db.UpdatedGroupColumnNamesInfo{
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
		// delete group columns
		if _, err := gdb.DeleteGroupColumns(db.DeletedGroupColumnNamesInfo{
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
		// add group columns
		if _, err := gdb.AddGroupColumns(db.AddedGroupColumnsInfo{
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
}

func item() {
	if gdb, err := initialDb(); err != nil {
		log.Fatal(err)
	} else {
		if _, err := gdb.AddGroups(db.AddedGroupInfo{
			GroupName:   "1DCS",
			ColumnNames: []string{"description"},
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("add 1DCS successfully!")
		}
		// add items
		if _, err := gdb.AddItems(db.AddedItemsInfo{
			GroupName: "1DCS",
			ItemValues: []map[string]string{{"itemName": "x", "description": "x", "dataType": "float64"},
				{"itemName": "y", "description": "y", "dataType": "float64"}, {"itemName": "z", "description": "z", "dataType": "float64"},
				{"itemName": "item1", "description": "item1", "dataType": "float64"}, {"itemName": "item2", "description": "item2", "dataType": "float64"}},
		}); err != nil {
			log.Fatal(err)
		} else {
			// startRow = -1 represent get all items with the given condition
			if r, err := gdb.GetItems(db.ItemsInfo{
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
			if r, err := gdb.GetItems(db.ItemsInfo{
				GroupName:   "1DCS",
				Condition:   "itemName like '%item%'",
				ColumnNames: "itemName",
				StartRow:    -1,
				RowCount:    0,
			}); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(r.ItemValues)
			}
		}
		// delete items
		if _, err := gdb.DeleteItems(db.DeletedItemsInfo{
			GroupName: "1DCS",
			Condition: "itemName like '%item%'",
		}); err != nil {
			log.Fatal(err)
		} else {
			if r, err := gdb.GetItems(db.ItemsInfo{
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
		// update items
		if _, err := gdb.UpdateItems(db.UpdatedItemsInfo{
			GroupName: "1DCS",
			Clause:    "description='x1'",
			Condition: "itemName='x'",
		}); err != nil {
			log.Fatal(err)
		} else {
			if r, err := gdb.GetItems(db.ItemsInfo{
				GroupName:   "1DCS",
				Condition:   "1=1",
				ColumnNames: "itemName, description",
				StartRow:    -1,
				RowCount:    0,
			}); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(r.ItemValues)
			}
		}
		// check items
		if err := gdb.CheckItems("1DCS", "x", "x1"); err != nil {
			fmt.Println(err)
		}
		// clean group items
		if _, err := gdb.CleanGroupItems("1DCS"); err != nil {
			log.Fatal(err)
		} else {
			if r, err := gdb.GetItems(db.ItemsInfo{
				GroupName:   "1DCS",
				Condition:   "1=1",
				ColumnNames: "*",
				StartRow:    -1,
				RowCount:    0,
			}); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(r)
			}
		}
	}
}

func data() {
	if gdb, err := initialDb(); err != nil {
		log.Fatal(err)
	} else {
		// batch write, y = 2 * x
		if _, err := gdb.BatchWrite([]db.ItemValue{{ItemName: "x", Value: 1.0, GroupName: "1DCS"}, {ItemName: "y", Value: 2.0, GroupName: "1DCS"}}...); err != nil {
			log.Fatal(err)
		} else {
			// get latest updated value of given items
			if r, err := gdb.GetRealTimeData([]string{"1DCS", "1DCS"}, "x", "y"); err != nil {
				log.Fatal(err)
			} else {
				d, _ := json.Marshal(r)
				fmt.Println(string(d))
			}
		}
		// write historical data
		// mock one hour historical data
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
		if err := gdb.BatchWriteHistoricalData([]db.HistoricalItemValue{{ItemName: "x", Values: xData, TimeStamps: ts, GroupName: "1CS"}, {ItemName: "y", Values: yData, TimeStamps: ts, GroupName: "1DCS"}}...); err != nil {
			log.Fatal(err)
		} else {
			// get raw historical data for debugging
			if r, err := gdb.GetRawHistoricalData([]string{"1DCS"}, "x"); err != nil {
				log.Fatal(err)
			} else {
				d, _ := json.Marshal(r)
				_ = ioutil.WriteFile("./rawX.txt", d, 0644)
			}
			// get historical data with given itemName, startTime, endTime and intervals
			stX := int(now.Add(time.Minute*5).Unix() + 8*3600)
			etX := int(now.Add(time.Minute*25).Unix() + 8*3600)
			stY := int(now.Add(time.Minute*35).Unix() + 8*3600)
			etY := int(now.Add(time.Minute*55).Unix() + 8*3600)
			if r, err := gdb.GetHistoricalData([]string{"1DCS", "1DCS"}, []string{"x", "y"}, []int{stX, stY}, []int{etX, etY}, []int{2, 10}); err != nil {
				log.Fatal(err)
			} else {
				d, _ := json.Marshal(r)
				_ = ioutil.WriteFile("./hX.txt", d, 0644)
			}
			// get historical data with given itemName
			if r, err := gdb.GetHistoricalDataWithStamp([]string{"1DCS", "1DCS"}, []string{"x", "y"}, [][]int{{stX, etX}, {stY, etY}}...); err != nil {
				log.Fatal(err)
			} else {
				d, _ := json.Marshal(r)
				fmt.Println(string(d))
			}
			// get historical data with condition
			if r, err := gdb.GetHistoricalDataWithCondition([]string{"1DCS", "1DCS"}, []string{"x", "y"}, []int{stX, stY}, []int{etX, etY}, []int{2, 10}, `item["x"] > 0 && item["y"] > 1000`, []db.DeadZone{}...); err != nil {
				log.Fatal(err)
			} else {
				d, _ := json.Marshal(r)
				_ = ioutil.WriteFile("./f.txt", d, 0644)
			}
		}
	}
}
