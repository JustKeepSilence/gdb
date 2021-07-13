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

func ExampleNewGdb() {
	//initial db with sqlite3
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gdb.hisTimeDuration)
	}
	// initial db with mysql
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, &Options{
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
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, &Options{
		DriverName:   "mysql",
		Dsn:          "root:admin@123@tcp(192.168.0.166:3306)/itemDb",
		UseInnerStop: true,
		RtGdb: &RedisRt{
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
}

func ExampleGdb_AddGroups() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.AddGroups(AddedGroupInfo{
			GroupName:   "1DCS",
			ColumnNames: []string{"unit", "descriptions"},
		}, AddedGroupInfo{
			GroupName:   "2DCS",
			ColumnNames: nil,
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.EffectedRows, r.Times)
		}
	}
}

func ExampleGdb_DeleteGroups() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.DeleteGroups(GroupNamesInfo{GroupNames: []string{"2DCS"}}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.EffectedRows, r.Times)
		}
	}
}

func ExampleGdb_GetGroups() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.GetGroups(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.GroupNames)
		}
	}
}

func ExampleGdb_GetGroupProperty() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.GetGroupProperty("1DCS", "1=1"); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.ItemCount, r.ItemColumnNames)
		}
	}
}

func ExampleGdb_UpdateGroupNames() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.UpdateGroupNames(UpdatedGroupNameInfo{
			OldGroupName: "1DCS",
			NewGroupName: "3DCS",
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_UpdateGroupColumnNames() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.UpdateGroupColumnNames(UpdatedGroupColumnNamesInfo{
			GroupName:      "3DCS",
			OldColumnNames: []string{"unit", "descriptions"},
			NewColumnNames: []string{"units", "description"},
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedCols)
		}
	}
}

func ExampleGdb_DeleteGroupColumns() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.DeleteGroupColumns(DeletedGroupColumnNamesInfo{
			GroupName:   "3DCS",
			ColumnNames: []string{"units"},
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedCols)
		}
	}
}

func ExampleGdb_AddGroupColumns() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.AddGroupColumns(AddedGroupColumnsInfo{
			GroupName:     "3DCS",
			ColumnNames:   []string{"unit"},
			DefaultValues: []string{"m/s"},
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedCols)
		}
	}
}

func ExampleGdb_CleanGroupItems() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.CleanGroupItems("3DCS"); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_AddItems() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.AddItems(AddedItemsInfo{
			GroupName: "3DCS",
			ItemValues: []map[string]string{{"itemName": "X", "dataType": "float32", "description": "", "unit": ""},
				{"itemName": "Y", "dataType": "float32", "description": "", "unit": ""}},
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_DeleteItems() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.DeleteItems(DeletedItemsInfo{
			GroupName: "3DCS",
			Condition: "1=1",
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_GetItems() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.GetItems(ItemsInfo{
			GroupName:   "3DCS",
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
}

func ExampleGdb_UpdateItems() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.UpdateItems(UpdatedItemsInfo{
			GroupName: "3DCS",
			Condition: "itemName='X'",
			Clause:    "description='X'",
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_CheckItems() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if err := gdb.CheckItems("3DCS", "X", "Z"); err != nil {
			log.Fatal(err)
		}
	}
}

func ExampleGdb_BatchWriteFloatData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.BatchWriteFloatData([]string{"3DCS", "4DCS"}, [][]string{{"X", "Y", "Z"}, {"X1", "Y1", "Z1"}}, [][]float32{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_BatchWriteIntData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.BatchWriteIntData([]string{"3DCS", "4DCS"}, [][]string{{"X", "Y", "Z"}, {"X1", "Y1", "Z1"}}, [][]int32{{1, 2, 3}, {4, 5, 6}}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_BatchWriteStringData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.BatchWriteStringData([]string{"3DCS", "4DCS"}, [][]string{{"X", "Y", "Z"}, {"X1", "Y1", "Z1"}}, [][]string{{"1.0", "2.0", "3.0"}, {"4.0", "5.0", "6.0"}}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_BatchWriteBoolData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.BatchWriteBoolData([]string{"3DCS", "4DCS"}, [][]string{{"X", "Y", "Z"}, {"X1", "Y1", "Z1"}}, [][]bool{{true, false, true}, {false, true, false}}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_BatchWriteFloatHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		//generate history data of month
		seconds := 24 * 3600 * 30
		now := time.Now()
		groupNames := []string{"3DCS", "4DCS"}
		itemNames := []string{"X", "X1"}
		var timeStamp []int32
		var xItemValue []float32
		var x1ItemValue []float32
		for i := 0; i < seconds; i++ {
			timeStamp = append(timeStamp, int32(now.Add(time.Duration(i)*time.Second).Unix()+8*3600))
			xItemValue = append(xItemValue, rand.Float32()*math.Pi)
			x1ItemValue = append(x1ItemValue, rand.Float32()*math.E)
		}
		if r, err := gdb.BatchWriteFloatHistoricalData(groupNames, itemNames, [][]int32{timeStamp, timeStamp}, [][]float32{xItemValue, x1ItemValue}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_BatchWriteIntHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		//generate history data of month
		seconds := 24 * 3600 * 30
		now := time.Now()
		groupNames := []string{"3DCS", "4DCS"}
		itemNames := []string{"X", "X1"}
		var timeStamp []int32
		var xItemValue []int32
		var x1ItemValue []int32
		for i := 0; i < seconds; i++ {
			timeStamp = append(timeStamp, int32(now.Add(time.Duration(i)*time.Second).Unix()+8*3600))
			xItemValue = append(xItemValue, rand.Int31())
			x1ItemValue = append(x1ItemValue, rand.Int31()*int32(i))
		}
		if r, err := gdb.BatchWriteIntHistoricalData(groupNames, itemNames, [][]int32{timeStamp, timeStamp}, [][]int32{xItemValue, x1ItemValue}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_BatchWriteStringHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		//generate history data of month
		seconds := 24 * 3600 * 30
		now := time.Now()
		groupNames := []string{"3DCS", "4DCS"}
		itemNames := []string{"X", "X1"}
		var timeStamp []int32
		var xItemValue []string
		var x1ItemValue []string
		for i := 0; i < seconds; i++ {
			timeStamp = append(timeStamp, int32(now.Add(time.Duration(i)*time.Second).Unix()+8*3600))
			xItemValue = append(xItemValue, now.Add(time.Duration(-i)*time.Second).Format("2006-01-02 15:04:05"))
			x1ItemValue = append(x1ItemValue, now.Add(time.Duration(i)*time.Second).Format("2006-01-02 15:04:05"))
		}
		if r, err := gdb.BatchWriteStringHistoricalData(groupNames, itemNames, [][]int32{timeStamp, timeStamp}, [][]string{xItemValue, x1ItemValue}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_BatchWriteBoolHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		//generate history data of month
		seconds := 24 * 3600 * 30
		now := time.Now()
		groupNames := []string{"3DCS", "4DCS"}
		itemNames := []string{"X", "X1"}
		var timeStamp []int32
		var xItemValue []bool
		var x1ItemValue []bool
		for i := 0; i < seconds; i++ {
			timeStamp = append(timeStamp, int32(now.Add(time.Duration(i)*time.Second).Unix()+8*3600))
			xItemValue = append(xItemValue, i%2 == 0)
			x1ItemValue = append(x1ItemValue, i%2 != 0)
		}
		if r, err := gdb.BatchWriteBoolHistoricalData(groupNames, itemNames, [][]int32{timeStamp, timeStamp}, [][]bool{xItemValue, x1ItemValue}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_GetRealTimeData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.GetRealTimeData([]string{"3DCS", "4DCS", "3DCS"}, []string{"X", "X1", "Y"}); err != nil {
			log.Fatal(err)
		} else {
			r1, _ := json.Marshal(r.RealTimeData)
			fmt.Println(string(r1))
		}
	}
	// Output:
	// {"X":1,"X1":4,"Y":null}
}

func ExampleGdb_GetFloatHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		st := int32(time.Now().Add(time.Hour*9*24).Unix() + 8*3600)
		et := int32(time.Now().Add(time.Hour*10*24).Unix() + 8*3600)
		if r, err := gdb.GetFloatHistoricalData([]string{"3DCS", "4DCS"}, []string{"Y", "X1"}, []int32{st, st}, []int32{et, et}, []int32{6 * 3600, 6 * 3600}); err != nil {
			log.Fatal(err)
		} else {
			r1, _ := json.Marshal(r.HistoricalData)
			fmt.Println(string(r1))
		}
	}
	// Output:
	// {"X1":[[1626838972,1626860572,1626882172,1626903772,1626838972,1626860572,1626882172,1626903772],[2.4913614,2.4554338,3.019772,2.8327777,2.4913614,2.4554338,3.019772,2.8327777]],"Y":[null,null]}
}

func ExampleGdb_GetIntHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		st := int32(time.Now().Add(time.Hour*9*24).Unix() + 8*3600)
		et := int32(time.Now().Add(time.Hour*10*24).Unix() + 8*3600)
		if r, err := gdb.GetIntHistoricalData([]string{"4DCS"}, []string{"xInt"}, []int32{st}, []int32{et}, []int32{6 * 3600}); err != nil {
			log.Fatal(err)
		} else {
			r1, _ := json.Marshal(r.HistoricalData)
			fmt.Println(string(r1))
		}
	}
	// Output:
	// {"xInt":[[1626840105,1626861705,1626883305,1626904905],[1932697017,132599088,1758527870,1151289996]]}
}

func ExampleGdb_GetStringHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		st := int32(time.Now().Add(time.Hour*9*24).Unix() + 8*3600)
		et := int32(time.Now().Add(time.Hour*10*24).Unix() + 8*3600)
		if r, err := gdb.GetStringHistoricalData([]string{"4DCS"}, []string{"xString"}, []int32{st}, []int32{et}, []int32{6 * 3600}); err != nil {
			log.Fatal(err)
		} else {
			r1, _ := json.Marshal(r.HistoricalData)
			fmt.Println(string(r1))
		}
	}
	// Output:
	// {"xString":[[1626839992,1626861592,1626883192,1626904792],["2021-07-02 15:59:52","2021-07-02 09:59:52","2021-07-02 03:59:52","2021-07-01 21:59:52"]]}
}

func ExampleGdb_GetBoolHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		st := int32(time.Now().Add(time.Hour*9*24).Unix() + 8*3600)
		et := int32(time.Now().Add(time.Hour*10*24).Unix() + 8*3600)
		if r, err := gdb.GetBoolHistoricalData([]string{"4DCS"}, []string{"xBool"}, []int32{st}, []int32{et}, []int32{6 * 3600}); err != nil {
			log.Fatal(err)
		} else {
			r1, _ := json.Marshal(r.HistoricalData)
			fmt.Println(string(r1))
		}
	}
	// Output:
	// {"xBool":[[1626840057,1626861657,1626883257,1626904857],[true,true,true,true]]}
}

func ExampleGdb_GetFloatRawHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.GetFloatRawHistoricalData([]string{"4DCS"}, []string{"X1"}); err != nil {
			log.Fatal(err)
		} else {
			v, _ := r.HistoricalData.Get("X1")
			r1, _ := json.Marshal(r)
			_ = ioutil.WriteFile("./h.txt", r1, 0766)
			fmt.Println(len(v.([]interface{})[0].([]int32)))
		}
	}
	// Output:
	// 2592000
}

func ExampleGdb_GetIntRawHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.GetIntRawHistoricalData([]string{"4DCS"}, []string{"xInt"}); err != nil {
			log.Fatal(err)
		} else {
			v, _ := r.HistoricalData.Get("xInt")
			r1, _ := json.Marshal(r)
			_ = ioutil.WriteFile("./h.txt", r1, 0766)
			fmt.Println(len(v.([]interface{})[0].([]int32)))
		}
	}
	// Output:
	// 2592000
}

func ExampleGdb_GetBoolRawHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.GetBoolRawHistoricalData([]string{"4DCS"}, []string{"xBool"}); err != nil {
			log.Fatal(err)
		} else {
			v, _ := r.HistoricalData.Get("xBool")
			r1, _ := json.Marshal(r)
			_ = ioutil.WriteFile("./h.txt", r1, 0766)
			fmt.Println(len(v.([]interface{})[0].([]int32)))
		}
	}
	// Output:
	// 2592000
}

func ExampleGdb_GetFloatHistoricalDataWithStamp() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		now := time.Now()
		ts := []int32{int32(now.Add(time.Hour*24*-30).Unix() + 8*3600)}
		for i := 0; i < 5; i++ {
			ts = append(ts, int32(now.Add(time.Hour*24*time.Duration(i)).Unix()+8*3600))
		}
		ts = append(ts, int32(now.Add(time.Hour*24*60).Unix()+8*3600)) // history of ts[0] and ts[len(ts) - 1] not exist, so we will not
		// return value of this two timeStamp
		if r, err := gdb.GetFloatHistoricalDataWithStamp([]string{"4DCS", "3DCS"}, []string{"X1", "Y"}, [][]int32{ts, ts}); err != nil {
			log.Fatal(err)
		} else {
			r1, _ := json.Marshal(r.HistoricalData)
			fmt.Println(string(r1))
		}
	}
	// Output:
	// {"X1":[[1626045377,1626131777,1626218177,1626304577,1626390977],[0.26911953,1.5107508,1.7978884,0.84435517,0.05435284]],"Y":[null,null]}
}

func ExampleGdb_GetIntHistoricalDataWithStamp() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		now := time.Now()
		ts := []int32{int32(now.Add(time.Hour*24*-30).Unix() + 8*3600)}
		for i := 0; i < 5; i++ {
			ts = append(ts, int32(now.Add(time.Hour*24*time.Duration(i)).Unix()+8*3600))
		}
		ts = append(ts, int32(now.Add(time.Hour*24*60).Unix()+8*3600)) // history of ts[0] and ts[len(ts) - 1] not exist, so we will not
		// return value of this two timeStamp
		if r, err := gdb.GetIntHistoricalDataWithStamp([]string{"4DCS"}, []string{"xInt"}, [][]int32{ts}); err != nil {
			log.Fatal(err)
		} else {
			r1, _ := json.Marshal(r.HistoricalData)
			fmt.Println(string(r1))
		}
	}
	// Output:
	// {"xInt":[[1626085279,1626171679,1626258079,1626344479,1626430879],[1208980727,2117437444,1398993897,781075096,2106205256]]}
}

func ExampleGdb_GetStringHistoricalDataWithStamp() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		now := time.Now()
		ts := []int32{int32(now.Add(time.Hour*24*-30).Unix() + 8*3600)}
		for i := 0; i < 5; i++ {
			ts = append(ts, int32(now.Add(time.Hour*24*time.Duration(i)).Unix()+8*3600))
		}
		ts = append(ts, int32(now.Add(time.Hour*24*60).Unix()+8*3600)) // history of ts[0] and ts[len(ts) - 1] not exist, so we will not
		// return value of this two timeStamp
		if r, err := gdb.GetStringHistoricalDataWithStamp([]string{"4DCS"}, []string{"xString"}, [][]int32{ts}); err != nil {
			log.Fatal(err)
		} else {
			r1, _ := json.Marshal(r.HistoricalData)
			fmt.Println(string(r1))
		}
	}
	// Output:
	// {"xString":[[1626085482,1626171882,1626258282,1626344682,1626431082],["2021-07-11 09:35:02","2021-07-10 09:35:02","2021-07-09 09:35:02","2021-07-08 09:35:02","2021-07-07 09:35:02"]]}
}

func ExampleGdb_GetBoolHistoricalDataWithStamp() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		now := time.Now()
		ts := []int32{int32(now.Add(time.Hour*24*-30).Unix() + 8*3600)}
		for i := 0; i < 5; i++ {
			ts = append(ts, int32(now.Add(time.Hour*24*time.Duration(i)).Unix()+8*3600))
		}
		ts = append(ts, int32(now.Add(time.Hour*24*60).Unix()+8*3600)) // history of ts[0] and ts[len(ts) - 1] not exist, so we will not
		// return value of this two timeStamp
		if r, err := gdb.GetBoolHistoricalDataWithStamp([]string{"4DCS"}, []string{"xBool"}, [][]int32{ts}); err != nil {
			log.Fatal(err)
		} else {
			r1, _ := json.Marshal(r.HistoricalData)
			fmt.Println(string(r1))
		}
	}
	// Output:
	// {"xBool":[[1626085566,1626171966,1626258366,1626344766,1626431166],[false,false,false,false,false]]}
}

func ExampleGdb_GetFloatHistoricalDataWithCondition() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		st, et := int32(1626201902), int32(1626288302)
		// without deadZone condition
		if r, err := gdb.GetFloatHistoricalDataWithCondition("4DCS", []string{"xFloat", "yFloat"}, []int32{st, st}, []int32{et, et}, []int32{10, 10}, `item["xFloat"]>= 1 && item["yFloat"]<= 4`, nil); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times)
			r1, _ := json.Marshal(r.HistoricalData)
			_ = ioutil.WriteFile("./hf2.json", r1, 0766)
		}
		// with deadZone condition
		if r, err := gdb.GetFloatHistoricalDataWithCondition("4DCS", []string{"xFloat", "yFloat"}, []int32{st, st}, []int32{et, et}, []int32{10, 10}, `item["xFloat"]>= 1 && item["yFloat"]<= 4`, []DeadZone{{ItemName: "xFloat", DeadZoneCount: 3}}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times)
			r1, _ := json.Marshal(r.HistoricalData)
			_ = ioutil.WriteFile("./hf2.json", r1, 0766)
		}
		// withOut filterCondition
		if r, err := gdb.GetFloatHistoricalDataWithCondition("4DCS", []string{"xFloat", "yFloat"}, []int32{st, st}, []int32{et, et}, []int32{10, 10}, `true`, []DeadZone{{ItemName: "xFloat", DeadZoneCount: 3}}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times)
			r1, _ := json.Marshal(r.HistoricalData)
			_ = ioutil.WriteFile("./hf2.json", r1, 0766)
		}
		// withOut filterCondition and deadZone condition == GetFloatHistoricalData
		if r, err := gdb.GetFloatHistoricalDataWithCondition("4DCS", []string{"xFloat", "yFloat"}, []int32{st, st}, []int32{et, et}, []int32{10, 10}, `true`, nil); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times)
			r1, _ := json.Marshal(r.HistoricalData)
			_ = ioutil.WriteFile("./hf2.json", r1, 0766)
		}
	}
}

func ExampleGdb_DeleteFloatHistoricalData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		st, et := int32(1626201902), int32(1626288302)
		if r, err := gdb.DeleteFloatHistoricalData([]string{"4DCS", "4DCS"}, []string{"xFloat", "yFloat"}, []int32{st, st}, []int32{et, et}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_ReLoadDb() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.ReLoadDb(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}

func ExampleGdb_CleanItemData() {
	if gdb, err := NewGdb("./historyDb", time.Hour, time.Minute*5, DefaultOptions()); err != nil {
		fmt.Println(err)
	} else {
		if r, err := gdb.CleanItemData(DeletedItemsInfo{
			GroupName: "4DCS",
			Condition: "1=1",
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.Times, r.EffectedRows)
		}
	}
}
