/*
creatTime: 2020/11/9 21:12
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package db

import (
	"fmt"
	"gdb/sqlite"
	"gdb/utils"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/gin-gonic/gin"
	jsonIter "github.com/json-iterator/go"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"io/ioutil"
	"strconv"
	"time"
)

var Json = jsonIter.ConfigCompatibleWithStandardLibrary // see: https://github.com/json-iterator/go

func handleError(c *gin.Context, message string) {
	serverError := ResponseData{500, message, nil}
	responseData, _ := Json.Marshal(serverError)
	b, _ := ioutil.ReadAll(c.Request.Body)
	utils.WriteError(c.Request.URL.String(), "POST", fmt.Sprintf("%s", b), message)
	c.String(500, fmt.Sprintf("%s", responseData))
}

// group handler
func (gdb *Gdb) AddGroupsHandler(c *gin.Context) {
	g := []AddGroupInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := AddGroups(g...) // add groups
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

func (gdb *Gdb) DeleteGroupsHandler(c *gin.Context) {
	g := DeletedGroupInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := DeleteGroups(g)
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

func (gdb *Gdb) GetGroupsHandler(c *gin.Context) {
	responseData, err := GetGroups()
	if err != nil {
		handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		c.String(200, "%s", r)
	}
}

func (gdb *Gdb) GetGroupPropertyHandler(c *gin.Context) {
	g := map[string][]string{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		groupNames := g["groupNames"]
		responseData, err := GetGroupProperty(groupNames...)
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

func (gdb *Gdb) UpdateGroupNamesHandler(c *gin.Context) {
	g := []UpdatedGroupInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := UpdateGroupNames(g...)
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) UpdateGroupColumnNamesHandler(c *gin.Context) {
	g := UpdatedGroupColumnInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := UpdateGroupColumnNames(g)
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) DeleteGroupColumnsHandler(c *gin.Context) {
	g := DeletedGroupColumnInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := DeleteGroupColumns(g)
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) AddGroupColumnsHandler(c *gin.Context) {
	g := AddGroupColumnInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := AddGroupColumns(g)
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

// item handler
func (gdb *Gdb) AddItemsHandler(c *gin.Context) {
	g := AddItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := gdb.AddItems(g) // add groups
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

func (gdb *Gdb) DeleteItemsHandler(c *gin.Context) {
	g := ItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := gdb.DeleteItems(g) // add groups
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

func (gdb *Gdb) GetItemsHandler(c *gin.Context) {
	g := ItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := GetItems(g) // add groups
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) UpdateItemsHandler(c *gin.Context) {
	g := ItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := gdb.UpdateItems(g) //
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

// data handler
func (gdb *Gdb) BatchWriteHandler(c *gin.Context) {
	//startTime := time.Now()
	var endTime, endTime1 time.Time
	request := c.Request
	var batchWriteString BatchWriteString
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&batchWriteString); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		kv := batchWriteString.ItemValues
		var responseData Rows
		var err error
		if len(kv) == 3 {
			// with time stamps
			if len(kv[0]) == len(kv[1]) && len(kv[1]) == len(kv[2]) {
				endTime = time.Now()
				responseData, err = gdb.BatchWrite(kv, true)
				endTime1 = time.Now()
			} else {
				err = fmt.Errorf("%s: The number of ItemNames, Values, and TimeStamps of the written data must be consistent", time.Now().Format(utils.TimeFormatString))
			}
		} else {
			// without time stamp
			if len(kv[0]) == len(kv[1]) {
				endTime = time.Now()
				responseData, err = gdb.BatchWrite(kv, false)
				endTime1 = time.Now()
				//fmt.Printf("[%s]: reading configs: %d ms,writing: %d ms\n", time.Now().Format(utils.TimeFormatString), endTime.Sub(startTime).Milliseconds(), endTime1.Sub(endTime).Milliseconds())
			} else {
				err = fmt.Errorf("%s: The number of ItemNames and Values must be consistent", time.Now().Format(utils.TimeFormatString))
			}
		}
		if err != nil {
			handleError(c, err.Error())
		} else {
			_ = gdb.InfoDb.Put([]byte(WrittenItems), []byte(fmt.Sprintf("%d", len(kv[0]))), nil)
			_ = gdb.InfoDb.Put([]byte(Speed), []byte(fmt.Sprintf("%dms/%d", endTime1.Sub(endTime).Milliseconds(), len(kv[0]))), nil)
			ts, _ := gdb.InfoDb.Get([]byte(TimeKey), nil)
			_ = gdb.InfoDb.Put([]byte(Speed+fmt.Sprintf("%s", ts)), []byte(fmt.Sprintf("%s", endTime1.Sub(endTime))), nil) // write history
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) GetRealTimeDataHandler(c *gin.Context) {
	startTime := time.Now()
	request := c.Request
	var realTimeDataString RealTimeDataString
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&realTimeDataString); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		itemNames := realTimeDataString.ItemNames
		endTime := time.Now()
		responseData, err := gdb.GetRealTimeData(itemNames...)
		endTime1 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,getting: %d ms\n", time.Now().Format(utils.TimeFormatString), endTime.Sub(startTime).Milliseconds(), endTime1.Sub(endTime).Milliseconds())
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) GetHistoricalDataHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	historicalDataString := HistoricalDataInfo{}
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&historicalDataString); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
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
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

func (gdb *Gdb) GetHistoricalDataWithConditionHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	g := HistoricalDataInfo{}
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
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
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

// delete historical data, you should stop other operation when deleting historical data
func (gdb *Gdb) DeleteHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	g := DeletedHistoricalDataInfo{}
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		itemNames := g.ItemNames
		timeStamps := g.TimeStamps
		if len(itemNames) != len(timeStamps) {
			handleError(c, "inconsistent length of itemNames and timeStamps")
		} else {
			if responseData, err := gdb.DeleteHistoricalData(itemNames, timeStamps); err != nil {
				handleError(c, err.Error())
			} else {
				r, _ := Json.Marshal(ResponseData{200, "", responseData})
				c.String(200, "%s", r)
			}
		}
	}
}

func (gdb *Gdb) GetHistoricalDataWithStampHandler(c *gin.Context) {
	startTime1 := time.Now()
	request := c.Request
	g := HistoricalDataInfo{}
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		itemNames := g.ItemNames
		timeStamps := g.TimeStamps
		endTime1 := time.Now()
		responseData, err := gdb.GetHistoricalDataWithStamp(itemNames, timeStamps...)
		endTime2 := time.Now()
		fmt.Printf("[%s]: reading configs: %d ms,writing: %d ms\n", time.Now().Format(utils.TimeFormatString), endTime1.Sub(startTime1).Milliseconds(), endTime2.Sub(endTime1).Milliseconds())
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

// get db info : ram, writtenItems, timestamp, speed
func (gdb *Gdb) GetDbInfoHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	responseData, err := gdb.getDbInfo()
	if err != nil {
		handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		c.String(200, "%s", r)
	}
}

func (gdb *Gdb) GetDbSpeedHistoryHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := HistoricalDataInfo{}
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := gdb.getDbSpeedHistory(Speed, g.StartTimes, g.EndTimes, g.Intervals[0])
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

// page handler
func (gdb *Gdb) HandleUserLogin(c *gin.Context) {
	g := authInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		err := gdb.userLogin(g) // add groups
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", nil})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) HandleGetUerInfo(c *gin.Context) {
	g := map[string]string{} // userName
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := gdb.getUserInfo(g["userName"]) // add groups
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}

}

func (gdb *Gdb) HandleUploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		if err := c.SaveUploadedFile(file, "./uploadFiles/"+file.Filename); err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", nil})
			c.String(200, fmt.Sprintf("%s", r))
		}
	}
}

func (gdb *Gdb) HandleAddItemsByExcel(c *gin.Context) {
	g := fileInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := gdb.addItemsByExcel(g) // add groups
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) HandleGetItemsWithCount(c *gin.Context) {
	g := ItemInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := GetItemsWithCount(g) // add groups
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) GetJsCodeHandler(c *gin.Context) {
	fileName := c.Param("fileName")
	request := c.Request
	defer request.Body.Close()
	responseData, err := getJsCode(fileName) // add groups
	if err != nil {
		handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		c.String(200, "%s", r)
	}
}

func (gdb *Gdb) GetLogsHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	responseData, err := getLogs() // add groups
	if err != nil {
		handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", responseData})
		c.String(200, "%s", r)
	}
}

func (gdb *Gdb) AddCalcItemHandler(c *gin.Context) {
	g := calcInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := gdb.testCalculation(g.Expression) // test expression
		if err != nil {
			handleError(c, err.Error())
		} else {
			// add item to calc_cfg
			createTime := time.Now().Format(utils.TimeFormatString)
			_, _ = sqlite.UpdateItem("insert into calc_cfg (description, expression, createTime) values ('" + g.Description + "', '" + g.Expression + "' , '" + createTime + "')")
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) GetCalcItemHandler(c *gin.Context) {
	g := map[string]string{} // condition
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := gdb.getCalculationItem(g["condition"]) // add groups
		if err != nil {
			handleError(c, err.Error())
		} else {
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) UpdateCalcItemHandler(c *gin.Context) {
	g := updatedCalculationInfo{}
	request := c.Request
	defer request.Body.Close()
	if err := Json.NewDecoder(request.Body).Decode(&g); err != nil {
		handleError(c, fmt.Sprintf("fail parsing string: %s", err))
	} else {
		responseData, err := gdb.testCalculation(g.Expression) // test expression
		if err != nil {
			handleError(c, err.Error())
		} else {
			_, _ = gdb.updateCalculationItem(g)
			r, _ := Json.Marshal(ResponseData{200, "", responseData})
			c.String(200, "%s", r)
		}
	}
}

func (gdb *Gdb) StartCalculationItemHandler(c *gin.Context) {
	id := c.Param("id") // get id
	_, err := sqlite.UpdateItem("update calc_cfg set status='true' where id=" + id)
	if err != nil {
		handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", Rows{1}})
		c.String(200, "%s", r)
	}
}

func (gdb *Gdb) StopCalculationItemHandler(c *gin.Context) {
	id := c.Param("id") // get id
	_, err := sqlite.UpdateItem("update calc_cfg set status='false' where id=" + id)
	if err != nil {
		handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", Rows{1}})
		c.String(200, "%s", r)
	}
}

func (gdb *Gdb) DeleteCalculationItemHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := sqlite.UpdateItem("delete from calc_cfg where id=" + id)
	if err != nil {
		handleError(c, err.Error())
	} else {
		r, _ := Json.Marshal(ResponseData{200, "", Rows{1}})
		c.String(200, "%s", r)
	}
}

// monitor
func (gdb *Gdb) GetProcessInfo() error {
	tm, _ := mem.VirtualMemory() // total RAM of machine
	ps, _ := process.Processes()
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			for _, p := range ps {
				name, _ := p.Name()
				if name == "db.exe" {
					m, _ := p.MemoryPercent()
					_ = gdb.InfoDb.Put([]byte(Ram), []byte(fmt.Sprintf("%.2f", float64(m)*float64(tm.Total)*10e-9)), nil) // write mem usage
				}
			}
			ts, _ := gdb.InfoDb.Get([]byte(TimeKey), nil)
			cs, _ := strconv.ParseInt(fmt.Sprintf("%s", ts), 10, 64)
			// un fresh
			if time.Now().Unix()-cs > 5 {
				_ = gdb.InfoDb.Put([]byte(WrittenItems), []byte(fmt.Sprintf("%d", 0)), nil)
				_ = gdb.InfoDb.Put([]byte(Speed), []byte(fmt.Sprintf("0ms/0")), nil)
			}
		}
	}
}

// calc
func (gdb *Gdb) Calc() error {
	startTicker := time.NewTicker(60 * time.Second) // every 60s update configs
	for {
		select {
		case startTime := <-startTicker.C:
			rows, _ := sqlite.Query("select id, expression, status, duration from calc_cfg where 1=1")
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
										_, _ = sqlite.UpdateItem("update calc_cfg set errorMessage='" + err.Error() + "' where id=" + r["id"])
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
