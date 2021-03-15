/*
creatTime: 2020/11/9 21:12
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package db

import (
	"fmt"
	"github.com/JustKeepSilence/gdb/config"
	"github.com/JustKeepSilence/gdb/sqlite"
	"github.com/JustKeepSilence/gdb/utils"
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

func (gdb *Gdb) handleError(c *gin.Context, message string) {
	serverError := ResponseData{500, message, nil}
	responseData, _ := Json.Marshal(serverError)
	b, _ := ioutil.ReadAll(c.Request.Body)
	_ = gdb.writeLog(Error, "post", message, c.Request.URL.String(), fmt.Sprintf("%s", b))
	c.String(500, fmt.Sprintf("%s", responseData))
}

// group handler, add json binding to all request data, for details see:
// https://github.com/gin-gonic/gin#model-binding-and-validation
func (gdb *Gdb) addGroupsHandler(c *gin.Context) {
	g := AddGroupInfos{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.AddGroups(g.GroupInfos...) // add groups
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) deleteGroupsHandler(c *gin.Context) {
	g := GroupNameInfos{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.DeleteGroups(g)
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}

	}
}

func (gdb *Gdb) getGroupsHandler(c *gin.Context) {
	responseData, err := gdb.GetGroups()
	if err != nil {
		gdb.handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		c.String(200, "%s", r)
	}
}

func (gdb *Gdb) getGroupPropertyHandler(c *gin.Context) {
	g := GetGroupPropertyInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		groupName, condition := g.GroupName, g.Condition
		responseData, err := gdb.GetGroupProperty(groupName, condition)
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

func (gdb *Gdb) updateGroupNamesHandler(c *gin.Context) {
	var g []UpdatedGroupInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.UpdateGroupNames(g...)
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) updateGroupColumnNamesHandler(c *gin.Context) {
	g := UpdatedGroupColumnInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.UpdateGroupColumnNames(g)
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) deleteGroupColumnsHandler(c *gin.Context) {
	g := DeleteGroupColumnInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.DeleteGroupColumns(g)
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) addGroupColumnsHandler(c *gin.Context) {
	g := AddGroupColumnInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.AddGroupColumns(g)
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

// item handler
func (gdb *Gdb) addItemsHandler(c *gin.Context) {
	g := AddItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.AddItems(g) // add groups
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

func (gdb *Gdb) deleteItemsHandler(c *gin.Context) {
	g := ItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.DeleteItems(g) // add groups
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

func (gdb *Gdb) getItemsHandler(c *gin.Context) {
	g := ItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.GetItems(g) // add groups
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) updateItemsHandler(c *gin.Context) {
	g := ItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.UpdateItems(g) //
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
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
		gdb.handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		//kv := batchWriteString.ItemValues
		var responseData Rows
		//var err error
		endTime = time.Now()
		responseData, err := gdb.BatchWrite(batchWriteString)
		endTime1 = time.Now()
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			_ = gdb.infoDb.Put([]byte(WrittenItems), []byte(fmt.Sprintf("%d", len(batchWriteString.ItemValues))), nil)
			_ = gdb.infoDb.Put([]byte(Speed), []byte(fmt.Sprintf("%dms/%d", endTime1.Sub(endTime).Milliseconds(), len(batchWriteString.ItemValues))), nil)
			ts, _ := gdb.infoDb.Get([]byte(TimeKey), nil)
			_ = gdb.infoDb.Put([]byte(Speed+fmt.Sprintf("%s", ts)), []byte(fmt.Sprintf("%s", endTime1.Sub(endTime))), nil) // write history
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) getRealTimeDataHandler(c *gin.Context) {
	startTime := time.Now()
	request := c.Request
	var realTimeDataString RealTimeDataString
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&realTimeDataString); err != nil {
		gdb.handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		itemNames := realTimeDataString.ItemNames
		endTime := time.Now()
		responseData, err := gdb.GetRealTimeData(itemNames...)
		endTime1 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,getting: %d ms\n", time.Now().Format(utils.TimeFormatString), endTime.Sub(startTime).Milliseconds(), endTime1.Sub(endTime).Milliseconds())
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) getHistoricalDataHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	historicalDataString := HistoricalDataInfo{}
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&historicalDataString); err != nil {
		gdb.handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		itemNames := historicalDataString.ItemNames
		startTimes := historicalDataString.StartTimes
		endTimes := historicalDataString.EndTimes
		intervals := historicalDataString.Intervals
		endTime1 := time.Now()
		responseData, err := gdb.GetHistoricalData(itemNames, startTimes, endTimes, intervals)
		endTime2 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,getting: %d ms\n", time.Now().Format(utils.TimeFormatString), endTime1.Sub(startTime1).Milliseconds(), endTime2.Sub(endTime1).Milliseconds())
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

func (gdb *Gdb) getHistoricalDataWithConditionHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	g := HistoricalDataInfoWithCondition{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
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
		fmt.Printf("[%s]: reading configs: %d ms,getting: %d ms\n", time.Now().Format(utils.TimeFormatString), endTime1.Sub(startTime1).Milliseconds(), endTime2.Sub(endTime1).Milliseconds())
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

// delete historical data, you should stop other operation when deleting historical data
func (gdb *Gdb) deleteHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	g := DeletedHistoricalDataInfo{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		itemNames := g.ItemNames
		timeStamps := g.TimeStamps
		if len(itemNames) != len(timeStamps) {
			gdb.handleError(c, "inconsistent length of itemNames and timeStamps")
		} else {
			if responseData, err := gdb.DeleteHistoricalData(itemNames, timeStamps); err != nil {
				gdb.handleError(c, err.Error())
			} else {
				r, _ := Json.Marshal(ResponseData{200, "", responseData})
				c.String(200, "%s", r)
			}
		}
	}
}

func (gdb *Gdb) getHistoricalDataWithStampHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	g := HistoricalDataInfoWithTimeStamp{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		itemNames := g.ItemNames
		timeStamps := g.TimeStamps
		endTime1 := time.Now()
		responseData, err := gdb.GetHistoricalDataWithStamp(itemNames, timeStamps...)
		endTime2 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,writing: %d ms\n", time.Now().Format(utils.TimeFormatString), endTime1.Sub(startTime1).Milliseconds(), endTime2.Sub(endTime1).Milliseconds())
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

// get db info : ram, writtenItems, timestamp, speed
func (gdb *Gdb) getDbInfoHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	responseData, err := gdb.getDbInfo()
	if err != nil {
		gdb.handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		c.String(200, "%s", r)
	}
}

func (gdb *Gdb) getDbSpeedHistoryHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := HistoricalDataInfo{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.getDbSpeedHistory(Speed, g.StartTimes, g.EndTimes, g.Intervals[0])
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) getRawDataHandler(c *gin.Context) {
	g := map[string][]string{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.GetRawHistoricalData(g["itemNames"]...) // add groups
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

// page handler
func (gdb *Gdb) handleUserLogin(c *gin.Context) {
	g := authInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		err := gdb.userLogin(g) // add groups
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", nil})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) handleGetUerInfo(c *gin.Context) {
	g := map[string]string{} // userName
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.getUserInfo(g["userName"]) // add groups
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) handleUploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		gdb.handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		if err := c.SaveUploadedFile(file, "./uploadFiles/"+file.Filename); err != nil {
			gdb.handleError(c, err.Error())
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
		gdb.handleError(c, "incorrect json form")
	} else {
		fileName, groupName := g.FileName, g.GroupName
		responseData, err := gdb.AddItemsByExcel(groupName, "./uploadFiles/"+fileName) // add groups
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) handleGetItemsWithCount(c *gin.Context) {
	g := ItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.GetItemsWithCount(g) // add groups
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) getJsCodeHandler(c *gin.Context) {
	fileName := c.Param("fileName")
	request := c.Request
	defer request.Body.Close()
	responseData, err := getJsCode(fileName) // add groups
	if err != nil {
		gdb.handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		c.String(200, "%s", r)
	}
}

func (gdb *Gdb) getLogsHandler(c *gin.Context) {
	request := c.Request
	g := getLogsInfo{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, err.Error())
	} else {
		logType, condition, startTime, endTime := g.LogType, g.Condition, g.StartTime, g.EndTime
		responseData, err := gdb.getLogs(logType, condition, startTime, endTime)
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) addCalcItemHandler(c *gin.Context) {
	g := calcInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.testCalculation(g.Expression) // test expression
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			// add item to calc_cfg
			createTime := time.Now().Format(utils.TimeFormatString)
			_, _ = sqlite.UpdateItem(gdb.ItemDbPath, "insert into calc_cfg (description, expression, createTime) values ('"+g.Description+"', '"+g.Expression+"' , '"+createTime+"')")
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) getCalcItemHandler(c *gin.Context) {
	g := map[string]string{} // condition
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.getCalculationItem(g["condition"]) // add groups
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) updateCalcItemHandler(c *gin.Context) {
	g := updatedCalculationInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.handleError(c, "incorrect json form")
	} else {
		responseData, err := gdb.testCalculation(g.Expression) // test expression
		if err != nil {
			gdb.handleError(c, err.Error())
		} else {
			_, _ = gdb.updateCalculationItem(g)
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) startCalculationItemHandler(c *gin.Context) {
	id := c.Param("id") // get id
	_, err := sqlite.UpdateItem(gdb.ItemDbPath, "update calc_cfg set status='true' where id="+id)
	if err != nil {
		gdb.handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", Rows{1}})
		c.String(200, "%s", r)
	}
}

func (gdb *Gdb) stopCalculationItemHandler(c *gin.Context) {
	id := c.Param("id") // get id
	_, err := sqlite.UpdateItem(gdb.ItemDbPath, "update calc_cfg set status='false' where id="+id)
	if err != nil {
		gdb.handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", Rows{1}})
		c.String(200, "%s", r)
	}
}

func (gdb *Gdb) deleteCalculationItemHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := sqlite.UpdateItem(gdb.ItemDbPath, "delete from calc_cfg where id="+id)
	if err != nil {
		gdb.handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", Rows{1}})
		c.String(200, "%s", r)
	}
}

// monitor
func (gdb *Gdb) getProcessInfo() error {
	tm, _ := mem.VirtualMemory() // total RAM of machine
	ps, _ := process.Processes()
	dbConfigs, err := config.ReadDbConfig("./config.json")
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
			rows, _ := sqlite.Query(gdb.ItemDbPath, "select id, expression, status, duration from calc_cfg where 1=1")
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
										_, _ = sqlite.UpdateItem(gdb.ItemDbPath, "update calc_cfg set errorMessage='"+err.Error()+"' where id="+r["id"])
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

func (gdb *Gdb) writeLog(level logLevel, requestMethod, message, requestString, requestUrl string) error {
	sqlStringBuilder := strings.Builder{}
	if level == 0 {
		// info
		sqlStringBuilder.Write([]byte("insert into log_cfg (logType, requestString, requestMethod, requestUrl, logMessage) values ('"))
		sqlStringBuilder.Write([]byte("info', '"))
		sqlStringBuilder.Write([]byte(requestString + "', '"))
		sqlStringBuilder.Write([]byte(requestMethod + "', '"))
		sqlStringBuilder.Write([]byte(requestUrl + "', '"))
		sqlStringBuilder.Write([]byte(message + "')"))
	} else {
		// default loType is error
		sqlStringBuilder.Write([]byte("insert into log_cfg (requestString, requestMethod, requestUrl, logMessage) values ('"))
		sqlStringBuilder.Write([]byte(requestString + "', '"))
		sqlStringBuilder.Write([]byte(requestMethod + "', '"))
		sqlStringBuilder.Write([]byte(requestUrl + "', '"))
		sqlStringBuilder.Write([]byte(message + "')"))
	}
	sqlString := sqlStringBuilder.String() // for debugging
	_, err := sqlite.UpdateItem(gdb.ItemDbPath, sqlString)
	if err != nil {
		return err
	}
	return nil
}
