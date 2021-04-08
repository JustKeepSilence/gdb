/*
creatTime: 2020/11/9 21:12
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package db

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/gin-gonic/gin"
	jsonIter "github.com/json-iterator/go"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"strconv"
	"strings"
	"time"
)

var Json = jsonIter.ConfigCompatibleWithStandardLibrary // see: https://github.com/json-iterator/go

// group handler, add json binding to all request data, for details see:
// https://github.com/gin-gonic/gin#model-binding-and-validation
func (gdb *Gdb) addGroupsHandler(c *gin.Context) {
	g := AddedGroupInfos{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.AddGroups(g.GroupInfos...) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) deleteGroupsHandler(c *gin.Context) {
	g := GroupNamesInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.DeleteGroups(g)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}

	}
}

func (gdb *Gdb) getGroupsHandler(c *gin.Context) {
	responseData, err := gdb.GetGroups()
	if err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()))
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		gdb.string(c, 200, "%s", r)
	}
}

func (gdb *Gdb) getGroupPropertyHandler(c *gin.Context) {
	g := QueryGroupPropertyInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		groupName, condition := g.GroupName, g.Condition
		responseData, err := gdb.GetGroupProperty(groupName, condition)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}

}

func (gdb *Gdb) updateGroupNamesHandler(c *gin.Context) {
	g := UpdatedGroupNamesInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.UpdateGroupNames(g.Infos...)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) updateGroupColumnNamesHandler(c *gin.Context) {
	g := UpdatedGroupColumnNamesInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.UpdateGroupColumnNames(g)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) deleteGroupColumnsHandler(c *gin.Context) {
	g := DeletedGroupColumnNamesInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.DeleteGroupColumns(g)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) addGroupColumnsHandler(c *gin.Context) {
	g := AddedGroupColumnsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.AddGroupColumns(g)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}

}

// item handler
func (gdb *Gdb) addItemsHandler(c *gin.Context) {
	g := AddedItemsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.AddItems(g) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}

}

func (gdb *Gdb) deleteItemsHandler(c *gin.Context) {
	g := DeletedItemsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.DeleteItems(g) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}

}

func (gdb *Gdb) getItemsHandler(c *gin.Context) {
	g := ItemsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.GetItems(g) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) handleGetItemsWithCount(c *gin.Context) {
	g := ItemsInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.GetItemsWithCount(g) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) updateItemsHandler(c *gin.Context) {
	g := ItemsInfoWithoutRow{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.UpdateItems(g) //
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

// data handler
func (gdb *Gdb) batchWriteHandler(c *gin.Context) {
	//startTime := time.Now()
	var endTime, endTime1 time.Time
	request := c.Request
	var batchWriteString BatchWriteString
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&batchWriteString); err != nil {
		gdb.string(c, 500, "%s", []byte(fmt.Sprintf("fail parsing string: %s", err)))
	} else {
		//kv := batchWriteString.ItemValues
		var responseData Rows
		//var err error
		endTime = time.Now()
		responseData, err := gdb.BatchWrite(batchWriteString)
		endTime1 = time.Now()
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			_ = gdb.infoDb.Put([]byte(WrittenItems), []byte(fmt.Sprintf("%d", len(batchWriteString.ItemValues))), nil)
			_ = gdb.infoDb.Put([]byte(Speed), []byte(fmt.Sprintf("%dms/%d", endTime1.Sub(endTime).Milliseconds(), len(batchWriteString.ItemValues))), nil)
			ts, _ := gdb.infoDb.Get([]byte(TimeKey), nil)
			_ = gdb.infoDb.Put([]byte(Speed+fmt.Sprintf("%s", ts)), []byte(fmt.Sprintf("%s", endTime1.Sub(endTime))), nil) // write history
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) getRealTimeDataHandler(c *gin.Context) {
	startTime := time.Now()
	request := c.Request
	var realTimeDataString QueryRealTimeDataString
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&realTimeDataString); err != nil {
		gdb.string(c, 500, "%s", []byte(fmt.Sprintf("fail parsing string: %s", err)))
	} else {
		itemNames := realTimeDataString.ItemNames
		endTime := time.Now()
		responseData, err := gdb.GetRealTimeData(itemNames...)
		endTime1 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,getting: %d ms\n", time.Now().Format(timeFormatString), endTime.Sub(startTime).Milliseconds(), endTime1.Sub(endTime).Milliseconds())
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) getHistoricalDataHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	historicalDataString := QueryHistoricalDataString{}
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&historicalDataString); err != nil {
		gdb.string(c, 500, "%s", []byte(fmt.Sprintf("fail parsing string: %s", err)))
	} else {
		itemNames := historicalDataString.ItemNames
		startTimes := historicalDataString.StartTimes
		endTimes := historicalDataString.EndTimes
		intervals := historicalDataString.Intervals
		endTime1 := time.Now()
		responseData, err := gdb.GetHistoricalData(itemNames, startTimes, endTimes, intervals)
		endTime2 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,getting: %d ms\n", time.Now().Format(timeFormatString), endTime1.Sub(startTime1).Milliseconds(), endTime2.Sub(endTime1).Milliseconds())
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}

}

func (gdb *Gdb) getHistoricalDataWithConditionHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	g := QueryHistoricalDataWithConditionString{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
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
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) getHistoricalDataWithStampHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	g := QueryHistoricalDataWithTimeStampString{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		itemNames := g.ItemNames
		timeStamps := g.TimeStamps
		endTime1 := time.Now()
		responseData, err := gdb.GetHistoricalDataWithStamp(itemNames, timeStamps...)
		endTime2 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,writing: %d ms\n", time.Now().Format(timeFormatString), endTime1.Sub(startTime1).Milliseconds(), endTime2.Sub(endTime1).Milliseconds())
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

// get db info : ram, writtenItems, timestamp, speed
func (gdb *Gdb) getDbInfoHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	responseData, err := gdb.getDbInfo()
	if err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()))
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		gdb.string(c, 200, "%s", r)
	}
}

func (gdb *Gdb) getDbSpeedHistoryHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := QuerySpeedHistoryDataString{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.getDbSpeedHistory(Speed, g.StartTimes, g.EndTimes, g.Interval)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) getRawDataHandler(c *gin.Context) {
	g := map[string][]string{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.GetRawHistoricalData(g["itemNames"]...) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

// page handler
func (gdb *Gdb) handleUserLogin(c *gin.Context) {
	g := authInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		token, err := gdb.userLogin(g, c.Request.Header.Get("User-Agent")) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", token})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) handleUserLogout(c *gin.Context) {
	userName := c.Param("userName")
	_, token, _ := c.Request.BasicAuth()
	if responseData, err := gdb.userLogout(userName, token, c.Request.Header.Get("User-Agent")); err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()))
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		gdb.string(c, 200, "%s", r)
	}
}

func (gdb *Gdb) handleGetUerInfo(c *gin.Context) {
	g := UserName{} // userName
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.getUserInfo(g.Name) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) handleUploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		gdb.string(c, 500, "%s", []byte(fmt.Sprintf("fail parsing string: %s", err)))
	} else {
		if err := c.SaveUploadedFile(file, "./uploadFiles/"+file.Filename); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", nil})
			c.String(200, fmt.Sprintf("%s", r))
		}
	}
}

func (gdb *Gdb) handleAddItemsByExcel(c *gin.Context) {
	g := fileInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		fileName, groupName := g.FileName, g.GroupName
		responseData, err := gdb.AddItemsByExcel(groupName, "./uploadFiles/"+fileName) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) getJsCodeHandler(c *gin.Context) {
	fileName := c.Param("fileName")
	request := c.Request
	defer request.Body.Close()
	responseData, err := getJsCode(fileName) // add groups
	if err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()))
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		gdb.string(c, 200, "%s", r)
	}
}

func (gdb *Gdb) getLogsHandler(c *gin.Context) {
	request := c.Request
	g := queryLogsInfo{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()))
	} else {
		logType, condition, startTime, endTime := g.LogType, g.Condition, g.StartTime, g.EndTime
		responseData, err := gdb.getLogs(logType, condition, startTime, endTime)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) addCalcItemHandler(c *gin.Context) {
	g := addedCalcItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.testCalculation(g.Expression) // test expression
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			// add item to calc_cfg
			createTime := time.Now().Format(timeFormatString)
			if _, err := updateItem(gdb.ItemDbPath, "insert into calc_cfg (description, expression, createTime, updatedTime, duration, status) values ('"+g.Description+"', '"+g.Expression+"' , '"+createTime+"', '"+createTime+"', '"+g.Duration+"', '"+g.Flag+"')"); err != nil {
				gdb.string(c, 500, "%s", []byte(err.Error()))
			} else {
				r, _ := Json.Marshal(ResponseData{200, "", responseData})
				gdb.string(c, 200, "%s", r)
			}
		}
	}
}

func (gdb *Gdb) getCalcItemsHandler(c *gin.Context) {
	g := queryCalcItemsInfo{} // condition
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.getCalculationItem(g.Condition) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) updateCalcItemHandler(c *gin.Context) {
	g := updatedCalcInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		responseData, err := gdb.testCalculation(g.Expression) // test expression
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			if _, err := gdb.updateCalculationItem(g); err != nil {
				gdb.string(c, 500, "%s", []byte(err.Error()))
			} else {
				r, _ := Json.Marshal(ResponseData{200, "", responseData})
				gdb.string(c, 200, "%s", r)
			}

		}
	}
}

func (gdb *Gdb) startCalculationItemHandler(c *gin.Context) {
	g := calcId{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		id := []string{}
		for _, item := range g.Id {
			id = append(id, "id = '"+item+"'")
		}
		_, err := updateItem(gdb.ItemDbPath, "update calc_cfg set status='true' where "+strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", Rows{len(g.Id)}})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) stopCalculationItemHandler(c *gin.Context) {
	g := calcId{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		id := []string{}
		for _, item := range g.Id {
			id = append(id, "id = '"+item+"'")
		}
		_, err := updateItem(gdb.ItemDbPath, "update calc_cfg set status='false' where "+strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", Rows{len(g.Id)}})
			gdb.string(c, 200, "%s", r)
		}
	}
}

func (gdb *Gdb) deleteCalculationItemHandler(c *gin.Context) {
	g := calcId{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form"))
	} else {
		id := []string{}
		for _, item := range g.Id {
			id = append(id, "id = '"+item+"'")
		}
		_, err := updateItem(gdb.ItemDbPath, "delete from where "+strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()))
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", Rows{len(g.Id)}})
			gdb.string(c, 200, "%s", r)
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
					_ = gdb.infoDb.Put([]byte(Ram), []byte(fmt.Sprintf("%.2f", float64(m)*float64(tm.Total)*10e-9)), nil) // write mem usage
				}
			}
			ts, _ := gdb.infoDb.Get([]byte(TimeKey), nil)
			cs, _ := strconv.ParseInt(fmt.Sprintf("%s", ts), 10, 64)
			// un fresh
			if time.Now().Unix()-cs > 5 {
				_ = gdb.infoDb.Put([]byte(WrittenItems), []byte(fmt.Sprintf("%d", 0)), nil)
				_ = gdb.infoDb.Put([]byte(Speed), []byte(fmt.Sprintf("0ms/0")), nil)
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

func (gdb *Gdb) writeLog(level logLevel, requestUrl, requestString, requestMethod, message, remoteAddress string) error {
	sqlStringBuilder := strings.Builder{}
	if level == 0 {
		// info
		sqlStringBuilder.Write([]byte("insert into log_cfg (logType, requestString, requestMethod, requestUrl, logMessage, remoteAddress) values ('"))
		sqlStringBuilder.Write([]byte("info', '"))
		sqlStringBuilder.Write([]byte(requestString + "', '"))
		sqlStringBuilder.Write([]byte(requestMethod + "', '"))
		sqlStringBuilder.Write([]byte(requestUrl + "', '"))
		sqlStringBuilder.Write([]byte(message + "', '"))
		sqlStringBuilder.Write([]byte(remoteAddress + "')"))
	} else {
		// default loType is error
		sqlStringBuilder.Write([]byte("insert into log_cfg (requestString, requestMethod, requestUrl, logMessage) values ('"))
		sqlStringBuilder.Write([]byte(requestString + "', '"))
		sqlStringBuilder.Write([]byte(requestMethod + "', '"))
		sqlStringBuilder.Write([]byte(requestUrl + "', '"))
		sqlStringBuilder.Write([]byte(message + "', '"))
		sqlStringBuilder.Write([]byte(remoteAddress + "')"))
	}
	sqlString := sqlStringBuilder.String() // for debugging
	_, err := updateItem(gdb.ItemDbPath, sqlString)
	if err != nil {
		return err
	}
	return nil
}

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
