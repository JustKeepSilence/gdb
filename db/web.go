// +build gdbClient

/*
creatTime: 2020/11/9
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"embed"
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/gin-gonic/gin"
	jsonIter "github.com/json-iterator/go"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var Json = jsonIter.ConfigCompatibleWithStandardLibrary // see: https://github.com/json-iterator/go

//go:embed templateFiles
var dFiles embed.FS

// group handler, add json binding to all request data, for details see:
// https://github.com/gin-gonic/gin#model-binding-and-validation
func (gdb *Gdb) addGroupsHandler(c *gin.Context) {
	g := AddedGroupInfos{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.AddGroups(g.GroupInfos...) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteGroupsHandler(c *gin.Context) {
	g := GroupNamesInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.DeleteGroups(g)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}

	}
}

func (gdb *Gdb) getGroupsHandler(c *gin.Context) {
	responseData, err := gdb.GetGroups()
	if err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()), gin.H{})
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		gdb.string(c, 200, "%s", r, gin.H{})
	}
}

func (gdb *Gdb) getGroupPropertyHandler(c *gin.Context) {
	g := queryGroupPropertyInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		groupName, condition := g.GroupName, g.Condition
		responseData, err := gdb.GetGroupProperty(groupName, condition)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}

}

func (gdb *Gdb) updateGroupNamesHandler(c *gin.Context) {
	g := UpdatedGroupNamesInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.UpdateGroupNames(g.Infos...)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) updateGroupColumnNamesHandler(c *gin.Context) {
	g := UpdatedGroupColumnNamesInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.UpdateGroupColumnNames(g)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteGroupColumnsHandler(c *gin.Context) {
	g := DeletedGroupColumnNamesInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.DeleteGroupColumns(g)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) addGroupColumnsHandler(c *gin.Context) {
	g := AddedGroupColumnsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.AddGroupColumns(g)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}

}

func (gdb *Gdb) cleanGroupItemsHandler(c *gin.Context) {
	g := GroupNamesInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.CleanGroupItems(g.GroupNames...); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

// item handler
func (gdb *Gdb) addItemsHandler(c *gin.Context) {
	g := AddedItemsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.AddItems(g) // add items
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}

}

func (gdb *Gdb) deleteItemsHandler(c *gin.Context) {
	g := DeletedItemsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.DeleteItems(g) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}

}

func (gdb *Gdb) getItemsHandler(c *gin.Context) {
	g := ItemsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.GetItems(g) // get items
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) handleGetItemsWithCount(c *gin.Context) {
	g := ItemsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.getItemsWithCount(g) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) updateItemsHandler(c *gin.Context) {
	g := UpdatedItemsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.UpdateItems(g) //
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) checkItemsHandler(c *gin.Context) {
	g := CheckItemsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if err := gdb.CheckItems(g.GroupName, g.ItemNames...); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", ""})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

// data handler
func (gdb *Gdb) batchWriteHandler(c *gin.Context) {
	//startTime := time.Now()
	var endTime, endTime1 time.Time
	request := c.Request
	var g batchWriteString
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		//kv := batchWriteString.ItemValues
		var responseData Rows
		//var err error
		endTime = time.Now()
		responseData, err := gdb.BatchWrite(g.ItemValues...)
		endTime1 = time.Now()
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			_ = gdb.infoDb.Put([]byte(writtenItems), []byte(fmt.Sprintf("%d", len(g.ItemValues))), nil)
			_ = gdb.infoDb.Put([]byte(speed), []byte(fmt.Sprintf("%dms/%d", endTime1.Sub(endTime).Milliseconds(), len(g.ItemValues))), nil)
			ts, _ := gdb.infoDb.Get([]byte(timeKey), nil)
			_ = gdb.infoDb.Put([]byte(speed+fmt.Sprintf("%s", ts)), []byte(fmt.Sprintf("%s", endTime1.Sub(endTime))), nil) // write history
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) batchWriteHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := batchWriteHistoricalString{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if err := gdb.BatchWriteHistoricalData(g.HistoricalItemValues...); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", ""})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getRealTimeDataHandler(c *gin.Context) {
	startTime := time.Now()
	request := c.Request
	var g queryRealTimeDataString
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		itemNames := g.ItemNames
		endTime := time.Now()
		responseData, err := gdb.GetRealTimeData(itemNames...)
		endTime1 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,getting: %d ms\n", time.Now().Format(timeFormatString), endTime.Sub(startTime).Milliseconds(), endTime1.Sub(endTime).Milliseconds())
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", gdbRealTimeData{responseData}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getHistoricalDataHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	g := queryHistoricalDataString{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		itemNames := g.ItemNames
		startTimes := g.StartTimes
		endTimes := g.EndTimes
		intervals := g.Intervals
		endTime1 := time.Now()
		responseData, err := gdb.GetHistoricalData(itemNames, startTimes, endTimes, intervals)
		endTime2 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,getting: %d ms\n", time.Now().Format(timeFormatString), endTime1.Sub(startTime1).Milliseconds(), endTime2.Sub(endTime1).Milliseconds())
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", gdbHistoricalData{HistoricalData: responseData}})
			gdb.string(c, 200, "%s", r, g)
		}
	}

}

func (gdb *Gdb) getHistoricalDataWithConditionHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	g := queryHistoricalDataWithConditionString{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		itemNames := g.ItemNames
		startTimes := g.StartTimes
		endTimes := g.EndTimes
		intervals := g.Intervals
		filterCondition := g.FilterCondition
		deadZones := g.DeadZones
		endTime1 := time.Now()
		responseData, err := gdb.GetHistoricalDataWithCondition(itemNames, startTimes, endTimes, intervals, filterCondition, deadZones...)
		endTime2 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,getting: %d ms\n", time.Now().Format(timeFormatString), endTime1.Sub(startTime1).Milliseconds(), endTime2.Sub(endTime1).Milliseconds())
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", gdbHistoricalData{responseData}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getHistoricalDataWithStampHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	g := queryHistoricalDataWithTimeStampString{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		itemNames := g.ItemNames
		timeStamps := g.TimeStamps
		endTime1 := time.Now()
		responseData, err := gdb.GetHistoricalDataWithStamp(itemNames, timeStamps...)
		endTime2 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,writing: %d ms\n", time.Now().Format(timeFormatString), endTime1.Sub(startTime1).Milliseconds(), endTime2.Sub(endTime1).Milliseconds())
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", gdbHistoricalData{responseData}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

// get db info : ram, writtenItems, timestamp, speed
func (gdb *Gdb) getDbInfoHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	responseData, err := gdb.getDbInfo()
	if err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()), gin.H{})
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", gdbInfoData{responseData}})
		gdb.string(c, 200, "%s", r, gin.H{})
	}
}

func (gdb *Gdb) getDbInfoHistoryHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := querySpeedHistoryDataString{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.getDbInfoHistory(g.ItemName, g.StartTimes, g.EndTimes, g.Interval)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", gin.H{"historicalData": responseData}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getRawDataHandler(c *gin.Context) {
	g := queryRealTimeDataString{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.GetRawHistoricalData(g.ItemNames...) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

// page handler
func (gdb *Gdb) handleUserLogin(c *gin.Context) {
	g := authInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		token, err := gdb.userLogin(g) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", token})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) handleUserLogout(c *gin.Context) {
	userName, _, _ := c.Request.BasicAuth()
	if responseData, err := gdb.userLogout(userName); err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()), gin.H{})
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		gdb.string(c, 200, "%s", r, gin.H{})
	}
}

func (gdb *Gdb) getUerInfoHandler(c *gin.Context) {
	g := userName{} // userName
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.getUserInfo(g.Name) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getUsersHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	if responseData, err := query(gdb.ItemDbPath, "select id, userName, role from user_cfg"); err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()), gin.H{})
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", gin.H{"userInfos": responseData}})
		gdb.string(c, 200, "%s", r, gin.H{})
	}
}

func (gdb *Gdb) addUsersHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := addedUserInfo{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.addUsers(g); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteUsersHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := userName{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.deleteUsers(g); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) updateUsersHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := updatedUserInfo{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.updateUsers(g); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) handleUploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		gdb.string(c, 500, "%s", []byte(fmt.Sprintf("fail parsing string: %s", err)), gin.H{})
	} else {
		if err := c.SaveUploadedFile(file, "./uploadFiles/"+file.Filename); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), gin.H{})
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", nil})
			c.String(200, fmt.Sprintf("%s", r))
		}
	}
}

func (gdb *Gdb) handleHttpsUploadFile(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := httpsFile{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		files, b := g.File, []uint8{}
		for _, file := range files {
			b = append(b, uint8(file))
		}
		if err := ioutil.WriteFile("./uploadFiles/"+g.FileName, b, 0644); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", ""})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) handleAddItemsByExcelHandler(c *gin.Context) {
	g := fileInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		fileName, groupName := g.FileName, g.GroupName
		responseData, err := gdb.addItemsByExcel(groupName, "./uploadFiles/"+fileName) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) importHistoryByExcelHandler(c *gin.Context) {
	g := historyFileInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		fileName, itemNames, sheetNames := g.FileName, g.ItemNames, g.SheetNames
		if err := gdb.importHistoryByExcel("./uploadFiles/"+fileName, itemNames, sheetNames...); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", ""})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getJsCodeHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := fileInfo{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		fileName := g.FileName
		responseData, err := getJsCode(fileName) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getLogsHandler(c *gin.Context) {
	request := c.Request
	g := queryLogsInfo{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()), g)
	} else {
		responseData, err := gdb.getLogs(g)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteLogsHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := deletedLogInfo{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.deleteLogs(g); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) downloadFileHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := fileInfo{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		fileName := g.FileName
		contents := []int32{}
		if fileContent, err := dFiles.ReadFile("templateFiles/" + fileName); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			for _, c := range fileContent {
				contents = append(contents, int32(c))
			}
			r, _ := Json.Marshal(ResponseData{200, "", gin.H{"contents": contents}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) addCalcItemHandler(c *gin.Context) {
	g := addedCalcItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.testCalculation(g.Expression) // test expression
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			// add item to calc_cfg
			createTime := time.Now().Format(timeFormatString)
			if _, err := updateItem(gdb.ItemDbPath, "insert into calc_cfg (description, expression, createTime, updatedTime, duration, status) values ('"+g.Description+"', '"+g.Expression+"' , '"+createTime+"', '"+createTime+"', '"+g.Duration+"', '"+g.Flag+"')"); err != nil {
				gdb.string(c, 500, "%s", []byte(err.Error()), g)
			} else {
				r, _ := Json.Marshal(ResponseData{200, "", responseData})
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getCalcItemsHandler(c *gin.Context) {
	g := queryCalcItemsInfo{} // condition
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.getCalculationItem(g.Condition) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) updateCalcItemHandler(c *gin.Context) {
	g := updatedCalcInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.testCalculation(g.Expression) // test expression
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			if _, err := gdb.updateCalculationItem(g); err != nil {
				gdb.string(c, 500, "%s", []byte(err.Error()), g)
			} else {
				r, _ := Json.Marshal(ResponseData{200, "", responseData})
				gdb.string(c, 200, "%s", r, g)
			}

		}
	}
}

func (gdb *Gdb) startCalculationItemHandler(c *gin.Context) {
	g := calcId{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		id := []string{}
		for _, item := range g.Id {
			id = append(id, "id = '"+item+"'")
		}
		_, err := updateItem(gdb.ItemDbPath, "update calc_cfg set status='true' where "+strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", Rows{len(g.Id)}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) stopCalculationItemHandler(c *gin.Context) {
	g := calcId{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		id := []string{}
		for _, item := range g.Id {
			id = append(id, "id = '"+item+"'")
		}
		_, err := updateItem(gdb.ItemDbPath, "update calc_cfg set status='false' where "+strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", Rows{len(g.Id)}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteCalculationItemHandler(c *gin.Context) {
	g := calcId{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		id := []string{}
		for _, item := range g.Id {
			id = append(id, "id = '"+item+"'")
		}
		_, err := updateItem(gdb.ItemDbPath, "delete from where "+strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", Rows{len(g.Id)}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

// monitor
func (gdb *Gdb) getProcessInfo() error {
	tm, _ := mem.VirtualMemory() // total RAM of machine
	ps, _ := process.Processes()
	dbConfigs, err := ReadDbConfig("./config.json")
	appName := dbConfigs.ApplicationName
	if err != nil {
		return err
	}
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			for _, p := range ps {
				name, _ := p.Name()
				if name == appName {
					m, _ := p.MemoryPercent()
					v := fmt.Sprintf("%.2f", float64(m)*float64(tm.Total)*10e-9)
					_ = gdb.infoDb.Put([]byte(ram), []byte(v), nil) // write mem usage
					currentTimeStamp := int(time.Now().Unix()) + 8*3600
					ts := strconv.Itoa(currentTimeStamp)
					_ = gdb.infoDb.Put([]byte(ram+fmt.Sprintf("%s", ts)), []byte(v), nil) // write history
				}
			}
			ts, _ := gdb.infoDb.Get([]byte(timeKey), nil)
			cs, _ := strconv.ParseInt(fmt.Sprintf("%s", ts), 10, 64)
			n := time.Now()
			currentTimeStamp := time.Date(n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Second(), 0, time.UTC).Unix() // unix timestamp
			// un fresh
			if currentTimeStamp-cs > 60 {
				_ = gdb.infoDb.Put([]byte(writtenItems), []byte(fmt.Sprintf("%d", 0)), nil)
				_ = gdb.infoDb.Put([]byte(speed), []byte(fmt.Sprintf("0ms/0")), nil)
			}
		}
	}
}

// calc
func (gdb *Gdb) calc() error {
	startTicker := time.NewTicker(60 * time.Second) // every 60s update configs
	for {
		select {
		case startTime := <-startTicker.C:
			rows, _ := query(gdb.ItemDbPath, "select id, expression, status, duration from calc_cfg where 1=1")
			for _, row := range rows {
				go func(r map[string]string) {
					status, _ := strconv.ParseBool(r["status"]) // if calc
					if status {
						// calc
						d, _ := strconv.ParseInt(r["duration"], 10, 64)
						expression := r["expression"]
						loop := eventloop.NewEventLoop()
						t := time.NewTicker(time.Duration(d) * time.Second)
						for {
							select {
							case endTime := <-t.C:
								if int(endTime.Sub(startTime).Seconds()) == 60 {
									return
								}
								// run script
								loop.Run(func(vm *goja.Runtime) {
									vm.Set("getRtData", gdb.getRtData)           // get realTime data
									vm.Set("getHData", gdb.getHData)             // get history
									vm.Set("writeRtData", gdb.BatchWrite)        // write data
									vm.Set("getTimeStamp", gdb.getUnixTimeStamp) // get timeStamp of given time string
									vm.Set("getNowTime", gdb.getNowTime)         // get current Time
									vm.Set("getTime", gdb.getTime)               // get time
									program, _ := goja.Compile("main.js", expression, false)
									_, err := vm.RunProgram(program)
									if err != nil {
										_, _ = updateItem(gdb.ItemDbPath, "update calc_cfg set errorMessage='"+err.Error()+"' where id="+r["id"])
									}
								})
							}
						}
					}
				}(row)
			}
		}
	}
}

// log

func (gdb *Gdb) cleanLogs() error {
	if dbConfigs, err := ReadDbConfig("./config.json"); err != nil {
		return err
	} else {
		expiredTime := dbConfigs.ExpiredTime
		if expiredTime == 0 {
			expiredTime = 3600 // default 1 hours, 3600s
		}
		if d, err := time.ParseDuration(strconv.Itoa(expiredTime) + "s"); err != nil {
			return err
		} else {
			t := time.NewTicker(d)
			for {
				select {
				case <-t.C:
					// clean logs
					expiredDate := time.Now().Add(d * -1).Format(timeFormatString)
					sqlString := "delete from log_cfg where insertTime < '" + expiredDate + "'"
					if _, err := updateItem(gdb.ItemDbPath, sqlString); err != nil {
						return err
					} else {
						return nil
					}
				}
			}
		}
	}
}
