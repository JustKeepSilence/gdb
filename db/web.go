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
	"github.com/deckarep/golang-set"
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

var json = jsonIter.ConfigCompatibleWithStandardLibrary // see: https://github.com/json-iterator/go

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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getGroupsHandler(c *gin.Context) {
	responseData, err := gdb.GetGroups()
	if err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()), gin.H{})
	} else {
		r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", ""})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", ""})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getRealTimeDataHandler(c *gin.Context) {
	request := c.Request
	var g queryRealTimeDataString
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		itemNames := g.ItemNames
		responseData, err := gdb.GetRealTimeData(g.GroupNames, itemNames...)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", gdbRealTimeData{responseData}}); err != nil {
				gdb.string(c, 500, "%s", []byte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getHistoricalDataHandler(c *gin.Context) {
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
		responseData, err := gdb.GetHistoricalData(g.GroupNames, itemNames, startTimes, endTimes, intervals)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", gdbHistoricalData{HistoricalData: responseData}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getHistoricalDataWithConditionHandler(c *gin.Context) {
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
		responseData, err := gdb.GetHistoricalDataWithCondition(g.GroupNames, itemNames, startTimes, endTimes, intervals, filterCondition, deadZones...)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", gdbHistoricalData{responseData}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getHistoricalDataWithStampHandler(c *gin.Context) {
	request := c.Request
	g := queryHistoricalDataWithTimeStampString{}
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		itemNames := g.ItemNames
		timeStamps := g.TimeStamps
		responseData, err := gdb.GetHistoricalDataWithStamp(g.GroupNames, itemNames, timeStamps...)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", gdbHistoricalData{responseData}})
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
		r, _ := json.Marshal(ResponseData{200, "", gdbInfoData{responseData}})
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
		responseData, err := gdb.getDbInfoHistory(g.ItemName, g.StartTimes, g.EndTimes, g.Intervals)
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", gin.H{"historicalData": responseData}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	if responseData, err := gdb.getRoutes(); err != nil {
		gdb.string(c, 500, "%s", []byte(err.Error()), nil)
	} else {
		r, _ := json.Marshal(ResponseData{200, "", map[string][]map[string]string{"routes": responseData}})
		gdb.string(c, 200, "%s", r, nil)
	}
}

func (gdb *Gdb) deleteRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := routesInfo{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if err := gdb.deleteRoutes(g.Name, g.Routes...); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), err)
		} else {
			responseData := Rows{EffectedRows: len(g.Routes)}
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) addRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := routesInfo{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if err := gdb.addRoutes(g.Name, g.Routes...); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), err)
		} else {
			responseData := Rows{EffectedRows: len(g.Routes)}
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) addUserRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := routesInfo{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if err := gdb.addUserRoutes(g.Name, g.Routes...); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), err)
		} else {
			responseData := Rows{EffectedRows: 1}
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteUserRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := userName{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if _, err := updateItem(gdb.ItemDbPath, "delete from route_cfg where userName='"+g.Name+"'"); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), err)
		} else {
			_ = gdb.e.LoadPolicy()
			responseData := Rows{EffectedRows: 1}
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getAllRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	r, _ := json.Marshal(ResponseData{200, "", map[string][][]string{"routes": {allRoutes, commonUserRoutes, visitorUserRoutes}}})
	gdb.string(c, 200, "%s", r, nil)
}

func (gdb *Gdb) checkRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	g := routesInfo{}
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, _ := gdb.checkRoutes(g.Name, g.Routes...)
		r, _ := json.Marshal(ResponseData{200, "", gin.H{"result": responseData}})
		gdb.string(c, 200, "%s", r, g)
	}
}

func (gdb *Gdb) getRawDataHandler(c *gin.Context) {
	g := queryRealTimeDataString{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.GetRawHistoricalData(g.GroupNames, g.ItemNames...) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", token})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) handleUserLogout(c *gin.Context) {
	g := gin.H{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		userName := g["name"].(string)
		if responseData, err := gdb.userLogout(userName); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), gin.H{})
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, gin.H{})
		}
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
		r, _ := json.Marshal(ResponseData{200, "", gin.H{"userInfos": responseData}})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", nil})
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
			r, _ := json.Marshal(ResponseData{200, "", ""})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
		if err := gdb.importHistoryByExcel("./uploadFiles/"+fileName, g.GroupName, itemNames, sheetNames...); err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", ""})
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
			r, _ := json.Marshal(ResponseData{200, "", map[string]string{"code": responseData}})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", gin.H{"contents": contents}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) testCalcItemHandler(c *gin.Context) {
	g := gin.H{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBind(&g); err != nil {
		gdb.string(c, 500, "%s", []byte("incorrect json form :"+err.Error()), g)
	} else {
		if expression, ok := g["expression"]; !ok {
			gdb.string(c, 500, "%s", []byte("json must contain expression field"), g)
		} else {
			if responseData, err := gdb.testCalculation(expression.(string)); err != nil {
				gdb.string(c, 500, "%s", []byte(err.Error()), g)
			} else {
				r, _ := json.Marshal(ResponseData{200, "", responseData})
				gdb.string(c, 200, "%s", r, g)
			}
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
				r, _ := json.Marshal(ResponseData{200, "", responseData})
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
			r, _ := json.Marshal(ResponseData{200, "", responseData})
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
				r, _ := json.Marshal(ResponseData{200, "", responseData})
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
		_, err := updateItem(gdb.ItemDbPath, "update calc_cfg set status='true', updatedTime = '"+time.Now().Format(timeFormatString)+"' where "+strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", Rows{len(g.Id)}})
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
		_, err := updateItem(gdb.ItemDbPath, "update calc_cfg set status='false', updatedTime='"+time.Now().Format(timeFormatString)+"' where "+strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", Rows{len(g.Id)}})
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
		_, err := updateItem(gdb.ItemDbPath, "delete from calc_cfg where "+strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", []byte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", Rows{len(g.Id)}})
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
	ch := make(chan map[string]calcConfig, 2)
	messageCh := make(chan messageCalcConfig, 2)
	// updated goroutine
	go func() {
		c := map[string]calcConfig{}
		//c := map[string]calcConfig{}
		flag := true // record whether it is the first time to run
		startTime := time.NewTicker(5 * time.Second)
		expressions := map[string]string{} // record expressions
		for {
			select {
			case <-startTime.C:
				if flag {
					rows, _ := query(gdb.ItemDbPath, "select id, expression, status, duration from calc_cfg where 1=1")
					for _, row := range rows {
						id, _ := strconv.ParseInt(row["id"], 10, 64)
						duration, _ := strconv.ParseInt(row["duration"], 10, 64)
						status, _ := strconv.ParseBool(row["status"]) // if calc
						f := gdb.getJsFunction(row["expression"], row["id"])
						c[row["id"]] = calcConfig{
							id:         id,
							f:          f,
							status:     status,
							expression: row["expression"],
							duration:   duration,
						}
					}
					ch <- c
					flag = false
				} else {
					for k, e := range expressions {
						c[k] = calcConfig{
							id:         c[k].id,
							f:          c[k].f,
							expression: e,
							status:     c[k].status,
							duration:   c[k].duration,
						}
					}
					rows, _ := query(gdb.ItemDbPath, "select id, expression, status, duration from calc_cfg where 1=1")
					m := messageCalcConfig{}
					index := []interface{}{}  // id in calc_cfg
					indexC := []interface{}{} // id in c
					for k := range c {
						indexC = append(indexC, k)
					}
					for _, row := range rows {
						id, _ := strconv.ParseInt(row["id"], 10, 64)
						duration, _ := strconv.ParseInt(row["duration"], 10, 64)
						status, _ := strconv.ParseBool(row["status"])
						expression := row["expression"]
						if r, ok := c[row["id"]]; ok {
							// calculate which parameters have been updated
							// dse : duration, status, expression
							// ds : duration, status
							// de : duration, expression
							// se : status, expression
							// d : duration
							// s: status
							// e: expression
							// del : delete item
							index = append(index, row["id"])
							temp := updatedInfo{}
							temp.id = row["id"]
							temp.newStatus = status
							temp.newDuration = duration
							f := gdb.getJsFunction(expression, row["id"])
							temp.f = f
							if r.status != status && r.duration != duration && r.expression != expression {
								// dse
								temp.updatedFiled = "dse"
							} else if r.status != status && r.duration != duration && r.expression == expression {
								// ds
								temp.updatedFiled = "ds"
							} else if r.status != status && r.duration == duration && r.expression != expression {
								// s
								temp.updatedFiled = "se"
							} else if r.status == status && r.duration != duration && r.expression != expression {
								// de
								temp.updatedFiled = "de"
							} else if r.status != status && r.duration == duration && r.expression == expression {
								// s
								temp.updatedFiled = "s"
							} else if r.status == status && r.duration != duration && r.expression == expression {
								// d
								temp.updatedFiled = "d"
							} else if r.status == status && r.duration == r.duration && r.expression != expression {
								// e
								temp.updatedFiled = "e"
							} else {
								// not update
								continue
							}
							c[row["id"]] = calcConfig{
								id:       c[row["id"]].id,
								f:        f,
								status:   status,
								duration: duration,
							} // update c
							expressions[row["id"]] = expression
							m.updatedInfos = append(m.updatedInfos, temp) // update m
						} else {
							// added infos
							info := calcConfig{
								id:       id,
								f:        gdb.getJsFunction(expression, row["id"]),
								status:   status,
								duration: duration,
							}
							m.addedInfos = append(m.addedInfos, info)
							c[row["id"]] = info
						}
					}
					// whether delete calc item
					dc := mapset.NewSet(indexC...).Difference(mapset.NewSet(index...)).ToSlice()
					for _, d := range dc {
						m.updatedInfos = append(m.updatedInfos, updatedInfo{
							id:           d.(string),
							updatedFiled: "del",
						})
					}
					messageCh <- m
				}
			}
		}
	}()
	// calc goroutine
	//rows := map[string]calcConfig{}
	rows := map[string]calcConfig{}
	configs := map[string]*time.Ticker{}
	for {
		select {
		case rows = <-ch:
			for _, row := range rows {
				r := row
				t := time.NewTicker(time.Duration(r.duration) * time.Second)
				configs[strconv.Itoa(int(r.id))] = t
				if r.status {
					go func(ts *time.Ticker) {
						for {
							select {
							case <-ts.C:
								if err := r.f(); err != nil {
									_, _ = updateItem(gdb.ItemDbPath, "update calc_cfg set errorMessage='"+strings.Replace(err.Error(), "'", `"`, -1)+"', status='false' where id="+strconv.Itoa(int(r.id)))
								}
							}
						}
					}(t)
				}
			}
		case ms := <-messageCh:
			for _, info := range ms.updatedInfos {
				in := info
				configs[in.id].Stop() // stop ticker
				u := in.updatedFiled
				f := in.f
				id := in.id
				rows[in.id] = calcConfig{
					id:       rows[in.id].id,
					f:        in.f,
					status:   in.newStatus,
					duration: in.newDuration,
				}
				if u != "del" {
					// not del
					if strings.Contains(u, "s") {
						// update status
						if in.newStatus {
							t := time.NewTicker(time.Duration(in.newDuration) * time.Second)
							configs[in.id] = t
							go func(ts *time.Ticker) {
								for {
									select {
									case <-ts.C:
										if err := f(); err != nil {
											_, _ = updateItem(gdb.ItemDbPath, "update calc_cfg set errorMessage='"+strings.Replace(err.Error(), "'", `"`, -1)+"', status='false' where id="+id)
										}
									}
								}
							}(t)
						}
					} else {
						// not update status
						t := time.NewTicker(time.Duration(in.newDuration) * time.Second)
						configs[in.id] = t
						go func(ts *time.Ticker) {
							for {
								select {
								case <-ts.C:
									if err := f(); err != nil {
										_, _ = updateItem(gdb.ItemDbPath, "update calc_cfg set errorMessage='"+strings.Replace(err.Error(), "'", `"`, -1)+"', status='false' where id="+id)
									}
								}
							}
						}(t)
					}
				}
			}
			for _, info := range ms.addedInfos {
				// added
				in := info
				id := in.id
				t := time.NewTicker(time.Duration(in.duration) * time.Second)
				configs[strconv.Itoa(int(in.id))] = t
				rows[strconv.Itoa(int(in.id))] = calcConfig{
					id:       in.id,
					f:        in.f,
					status:   in.status,
					duration: in.duration,
				}
				go func(ts *time.Ticker) {
					for {
						select {
						case <-ts.C:
							if err := in.f(); err != nil {
								_, _ = updateItem(gdb.ItemDbPath, "update calc_cfg set errorMessage='"+strings.Replace(err.Error(), "'", `"`, -1)+"', status='false' where id="+strconv.Itoa(int(id)))
							}
						}
					}
				}(t)
			}
		}
	}
}

func (gdb *Gdb) getJsFunction(expression, id string) func() error {
	return func() error {
		loop := eventloop.NewEventLoop()
		var err error
		loop.Run(func(vm *goja.Runtime) {
			vm.Set("getRtData", gdb.getRtData)
			vm.Set("getHData", gdb.getHData)
			vm.Set("getHDataWithTs", gdb.getHDataWithTs)
			vm.Set("writeRtData", gdb.writeRtData)
			vm.Set("getTimeStamp", gdb.getUnixTimeStamp)
			vm.Set("getNowTime", gdb.getNowTime)
			vm.Set("testItemValue", gdb.testItemValue)
			program, _ := goja.Compile(id+".js", expression, false)
			_, err = vm.RunProgram(program)
		})
		if err != nil {
			return err
			//_, _ = updateItem(gdb.ItemDbPath, "update calc_cfg set errorMessage='"+err.Error()+"' where id="+r["id"])
		} else {
			return nil
		}
	}
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
