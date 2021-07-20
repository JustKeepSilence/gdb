// +build gdbServer

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
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/syndtr/goleveldb/leveldb"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

//go:embed templateFiles
var dFiles embed.FS

// group handler, add json binding to all request data, for details see:
// https://github.com/gin-gonic/gin#model-binding-and-validation
func (gdb *Gdb) addGroupsHandler(c *gin.Context) {
	var g addedGroupInfos
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.AddGroups(g.GroupInfos...) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteGroupsHandler(c *gin.Context) {
	var g GroupNamesInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.DeleteGroups(g)
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getGroupsHandler(c *gin.Context) {
	responseData, err := gdb.GetGroups()
	if err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(err.Error()), gin.H{})
	} else {
		r, _ := json.Marshal(ResponseData{200, "", responseData})
		gdb.string(c, 200, "%s", r, gin.H{})
	}
}

func (gdb *Gdb) getGroupPropertyHandler(c *gin.Context) {
	var g queryGroupPropertyInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		groupName, condition := g.GroupName, g.Condition
		responseData, err := gdb.GetGroupProperty(groupName, condition)
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) updateGroupNamesHandler(c *gin.Context) {
	var g updatedGroupNamesInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.UpdateGroupNames(g.Infos...)
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) updateGroupColumnNamesHandler(c *gin.Context) {
	var g UpdatedGroupColumnNamesInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.UpdateGroupColumnNames(g)
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteGroupColumnsHandler(c *gin.Context) {
	var g DeletedGroupColumnNamesInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.DeleteGroupColumns(g)
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) addGroupColumnsHandler(c *gin.Context) {
	var g AddedGroupColumnsInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.AddGroupColumns(g)
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}

}

func (gdb *Gdb) cleanGroupItemsHandler(c *gin.Context) {
	var g GroupNamesInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.CleanGroupItems(g.GroupNames...); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

// item handler
func (gdb *Gdb) addItemsHandler(c *gin.Context) {
	var g AddedItemsInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.AddItems(g) // add items
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}

}

func (gdb *Gdb) deleteItemsHandler(c *gin.Context) {
	var g DeletedItemsInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.DeleteItems(g) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}

}

func (gdb *Gdb) handleGetItemsWithCount(c *gin.Context) {
	var g ItemsInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.getItemsWithCount(g) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) updateItemsHandler(c *gin.Context) {
	var g UpdatedItemsInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.UpdateItems(g) //
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) checkItemsHandler(c *gin.Context) {
	var g checkItemsInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if err := gdb.CheckItems(g.GroupName, g.ItemNames...); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", ""})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) batchWriteFloatDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g floatItemValues
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.BatchWriteFloatData(g.GroupNames, g.ItemNames, g.ItemValues); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) batchWriteIntDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g intItemValues
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.BatchWriteIntData(g.GroupNames, g.ItemNames, g.ItemValues); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) batchWriteBoolDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g boolItemValues
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.BatchWriteBoolData(g.GroupNames, g.ItemNames, g.ItemValues); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) batchWriteStringDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g stringItemValues
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.BatchWriteStringData(g.GroupNames, g.ItemNames, g.ItemValues); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) batchWriteFloatHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g floatHItemValues
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.BatchWriteFloatHistoricalData(g.GroupNames, g.ItemNames, g.TimeStamps, g.ItemValues); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) batchWriteIntHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g intHItemValues
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.BatchWriteIntHistoricalData(g.GroupNames, g.ItemNames, g.TimeStamps, g.ItemValues); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) batchWriteStringHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g stringHItemValues
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.BatchWriteStringHistoricalData(g.GroupNames, g.ItemNames, g.TimeStamps, g.ItemValues); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) batchWriteBoolHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g boolHItemValues
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.BatchWriteBoolHistoricalData(g.GroupNames, g.ItemNames, g.TimeStamps, g.ItemValues); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getRealTimeDataHandler(c *gin.Context) {
	request := c.Request
	var g queryRealTimeDataString
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		responseData, err := gdb.GetRealTimeData(g.GroupNames, g.ItemNames)
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getFloatHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetFloatHistoricalData(g.GroupNames, g.ItemNames, g.StartTimes, g.EndTimes, g.Intervals); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getIntHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetIntHistoricalData(g.GroupNames, g.ItemNames, g.StartTimes, g.EndTimes, g.Intervals); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getStringHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetStringHistoricalData(g.GroupNames, g.ItemNames, g.StartTimes, g.EndTimes, g.Intervals); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getBoolHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetBoolHistoricalData(g.GroupNames, g.ItemNames, g.StartTimes, g.EndTimes, g.Intervals); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getFloatRawHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryRawHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetFloatRawHistoricalData(g.GroupNames, g.ItemNames); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getIntRawHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryRawHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetIntRawHistoricalData(g.GroupNames, g.ItemNames); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getStringRawHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryRawHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetStringRawHistoricalData(g.GroupNames, g.ItemNames); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getBoolRawHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryRawHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetBoolRawHistoricalData(g.GroupNames, g.ItemNames); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getFloatHistoricalDataWithStampHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataWithStampString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetFloatHistoricalDataWithStamp(g.GroupNames, g.ItemNames, g.TimeStamps); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getIntHistoricalDataWithStampHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataWithStampString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetIntHistoricalDataWithStamp(g.GroupNames, g.ItemNames, g.TimeStamps); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getStringHistoricalDataWithStampHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataWithStampString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetStringHistoricalDataWithStamp(g.GroupNames, g.ItemNames, g.TimeStamps); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getBoolHistoricalDataWithStampHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataWithStampString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetBoolHistoricalDataWithStamp(g.GroupNames, g.ItemNames, g.TimeStamps); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getFloatHistoricalDataWithConditionHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataWithConditionString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetFloatHistoricalDataWithCondition(g.GroupName, g.ItemNames, g.StartTime, g.EndTime, g.Interval, g.FilterCondition, g.DeadZones); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getIntHistoricalDataWithConditionHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataWithConditionString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetIntHistoricalDataWithCondition(g.GroupName, g.ItemNames, g.StartTime, g.EndTime, g.Interval, g.FilterCondition, g.DeadZones); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getStringHistoricalDataWithConditionHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataWithConditionString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetStringHistoricalDataWithCondition(g.GroupName, g.ItemNames, g.StartTime, g.EndTime, g.Interval, g.FilterCondition, g.DeadZones); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getBoolHistoricalDataWithConditionHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g queryHistoricalDataWithConditionString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.GetBoolHistoricalDataWithCondition(g.GroupName, g.ItemNames, g.StartTime, g.EndTime, g.Interval, g.FilterCondition, g.DeadZones); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) deleteFloatHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g deleteHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.DeleteFloatHistoricalData(g.GroupNames, g.ItemNames, g.StartTimes, g.EndTimes); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) deleteIntHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g deleteHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.DeleteIntHistoricalData(g.GroupNames, g.ItemNames, g.StartTimes, g.EndTimes); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) deleteStringHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g deleteHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.DeleteStringHistoricalData(g.GroupNames, g.ItemNames, g.StartTimes, g.EndTimes); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) deleteBoolHistoricalDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g deleteHistoricalDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.DeleteBoolHistoricalData(g.GroupNames, g.ItemNames, g.StartTimes, g.EndTimes); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) cleanItemDataHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g DeletedItemsInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), g)
	} else {
		if responseData, err := gdb.CleanItemData(g); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) reLoadDbHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	if responseData, err := gdb.ReLoadDb(); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(err.Error()), nil)
	} else {
		if r, err := json.Marshal(ResponseData{200, "", responseData}); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), nil)
		} else {
			gdb.string(c, 200, "%s", r, nil)
		}
	}
}

// get db info : ram, writtenItems, timestamp, speed
func (gdb *Gdb) getDbInfoHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	responseData, err := gdb.getDbInfo()
	if err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(err.Error()), gin.H{})
	} else {
		r, _ := json.Marshal(ResponseData{200, "", gdbInfoData{responseData}})
		gdb.string(c, 200, "%s", r, gin.H{})
	}
}

func (gdb *Gdb) getDbInfoHistoryHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g querySpeedHistoryDataString
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.getDbInfoHistory(g.InfoType, g.ItemName, g.StartTimes, g.EndTimes, g.Intervals)
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
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
		gdb.string(c, 500, "%s", convertStringToByte(err.Error()), nil)
	} else {
		r, _ := json.Marshal(ResponseData{200, "", map[string][]map[string]string{"routes": responseData}})
		gdb.string(c, 200, "%s", r, nil)
	}
}

func (gdb *Gdb) deleteRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g routesInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if err := gdb.deleteRoutes(g.Name, g.Routes...); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), err)
		} else {
			responseData := TimeRows{EffectedRows: len(g.Routes)}
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) addRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g routesInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if err := gdb.addRoutes(g.Name, g.Routes...); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), err)
		} else {
			responseData := TimeRows{EffectedRows: len(g.Routes)}
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) addUserRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g routesInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if err := gdb.addUserRoutes(g.Name, g.Routes...); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), err)
		} else {
			responseData := TimeRows{EffectedRows: 1}
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteUserRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g userName
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if _, err := gdb.updateItem("delete from route_cfg where userName='" + g.Name + "'"); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), err)
		} else {
			_ = gdb.e.LoadPolicy()
			responseData := TimeRows{EffectedRows: 1}
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getAllRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	r, _ := json.Marshal(ResponseData{200, "", map[string][][]string{"routes": {superUserRoutes, commonUserRoutes, visitorUserRoutes}}})
	gdb.string(c, 200, "%s", r, nil)
}

func (gdb *Gdb) checkRoutesHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g routesInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, _ := gdb.checkRoutes(g.Name, g.Routes...)
		r, _ := json.Marshal(ResponseData{200, "", gin.H{"result": responseData}})
		gdb.string(c, 200, "%s", r, g)
	}
}

func (gdb *Gdb) handleUserLogin(c *gin.Context) {
	var g authInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		token, err := gdb.userLogin(g) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
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
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		userName := g["name"].(string)
		if responseData, err := gdb.userLogout(userName); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), gin.H{})
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, gin.H{})
		}
	}
}

func (gdb *Gdb) getUerInfoHandler(c *gin.Context) {
	var g userName // userName
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.getUserInfo(g.Name) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getUsersHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	if responseData, err := gdb.query("select id, userName, role from user_cfg"); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(err.Error()), gin.H{})
	} else {
		r, _ := json.Marshal(ResponseData{200, "", gin.H{"userInfos": responseData}})
		gdb.string(c, 200, "%s", r, gin.H{})
	}
}

func (gdb *Gdb) addUsersHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g addedUserInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.addUsers(g); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteUsersHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g userName
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.deleteUsers(g); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) updateUsersHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g updatedUserInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.updateUsers(g); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) handleUploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(fmt.Sprintf("fail parsing string: %s", err)), gin.H{})
	} else {
		if err := c.SaveUploadedFile(file, "./uploadFiles/"+file.Filename); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), gin.H{})
		} else {
			r, _ := json.Marshal(ResponseData{200, "", nil})
			c.String(200, fmt.Sprintf("%s", r))
		}
	}
}

func (gdb *Gdb) handleHttpsUploadFile(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g httpsFile
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		files, b := g.File, []uint8{}
		for _, file := range files {
			b = append(b, uint8(file))
		}
		if err := ioutil.WriteFile("./uploadFiles/"+g.FileName, b, 0644); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", ""})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) handleAddItemsByExcelHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g fileInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		fileName, groupName := g.FileName, g.GroupName
		responseData, err := gdb.addItemsByExcel(groupName, "./uploadFiles/"+fileName) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) importHistoryByExcelHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g historyFileInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		fileName, itemNames, sheetNames := g.FileName, g.ItemNames, g.SheetNames
		if responseData, err := gdb.importHistoryByExcel("./uploadFiles/"+fileName, g.GroupName, itemNames, sheetNames...); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getJsCodeHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g fileInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		fileName := g.FileName
		responseData, err := getJsCode(fileName) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", map[string]string{"code": responseData}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getLogsHandler(c *gin.Context) {
	request := c.Request
	var g queryLogsInfo
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
	} else {
		responseData, err := gdb.getLogs(g)
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteLogsHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g deletedLogInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if responseData, err := gdb.deleteLogs(g); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) downloadFileHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	var g fileInfo
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		fileName := g.FileName
		contents := []int32{}
		if fileContent, err := dFiles.ReadFile("templateFiles/" + fileName); err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			for _, c := range fileContent {
				contents = append(contents, int32(c))
			}
			r, _ := json.Marshal(ResponseData{200, "", gin.H{"contents": contents}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getDbSizeHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	if responseData, err := gdb.getDbSize(); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), nil)
	} else {
		r, _ := json.Marshal(ResponseData{200, "", gin.H{"fileSize": responseData}})
		gdb.string(c, 200, "%s", r, nil)
	}
}

func (gdb *Gdb) testCalcItemHandler(c *gin.Context) {
	g := gin.H{}
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		if expression, ok := g["expression"]; !ok {
			gdb.string(c, 500, "%s", convertStringToByte("json must contain expression field"), g)
		} else {
			if responseData, err := gdb.testCalculation(expression.(string)); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				r, _ := json.Marshal(ResponseData{200, "", responseData})
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) addCalcItemHandler(c *gin.Context) {
	var g addedCalcItemInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.testCalculation(g.Expression) // test expression
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			// add item to calc_cfg
			createTime := time.Now().Format(timeFormatString)
			if _, err := gdb.updateItem("insert into calc_cfg (description, expression, createTime, updatedTime, duration, status) values ('" + g.Description + "', '" + g.Expression + "' , '" + createTime + "', '" + createTime + "', '" + g.Duration + "', '" + g.Flag + "')"); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				r, _ := json.Marshal(ResponseData{200, "", responseData})
				gdb.string(c, 200, "%s", r, g)
			}
		}
	}
}

func (gdb *Gdb) getCalcItemsHandler(c *gin.Context) {
	var g queryCalcItemsInfo // condition
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.getCalculationItem(g.Condition) // add groups
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", responseData})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) updateCalcItemHandler(c *gin.Context) {
	var g updatedCalcInfo
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		responseData, err := gdb.testCalculation(g.Expression) // test expression
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			if _, err := gdb.updateCalculationItem(g); err != nil {
				gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
			} else {
				r, _ := json.Marshal(ResponseData{200, "", responseData})
				gdb.string(c, 200, "%s", r, g)
			}

		}
	}
}

func (gdb *Gdb) startCalculationItemHandler(c *gin.Context) {
	var g calcId
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		id := []string{}
		for _, item := range g.Id {
			id = append(id, "id = '"+item+"'")
		}
		st := time.Now()
		_, err := gdb.updateItem("update calc_cfg set status='true', updatedTime = '" + time.Now().Format(timeFormatString) + "' where " + strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", TimeRows{EffectedRows: len(g.Id), Times: time.Since(st).Milliseconds()}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) stopCalculationItemHandler(c *gin.Context) {
	var g calcId
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		id := []string{}
		for _, item := range g.Id {
			id = append(id, "id = '"+item+"'")
		}
		st := time.Now()
		_, err := gdb.updateItem("update calc_cfg set status='false', updatedTime='" + time.Now().Format(timeFormatString) + "' where " + strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", TimeRows{EffectedRows: len(g.Id), Times: time.Since(st).Milliseconds()}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) deleteCalculationItemHandler(c *gin.Context) {
	var g calcId
	request := c.Request
	defer request.Body.Close()
	if err := c.ShouldBindJSON(&g); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte("incorrect json form :"+err.Error()), g)
	} else {
		id := []string{}
		for _, item := range g.Id {
			id = append(id, "id = '"+item+"'")
		}
		st := time.Now()
		_, err := gdb.updateItem("delete from calc_cfg where " + strings.Join(id, " or "))
		if err != nil {
			gdb.string(c, 500, "%s", convertStringToByte(err.Error()), g)
		} else {
			r, _ := json.Marshal(ResponseData{200, "", TimeRows{EffectedRows: len(g.Id), Times: time.Since(st).Milliseconds()}})
			gdb.string(c, 200, "%s", r, g)
		}
	}
}

func (gdb *Gdb) getCmdInfoHandler(c *gin.Context) {
	request := c.Request
	defer request.Body.Close()
	param := c.Param("name")
	if responseData, err := gdb.getCmdInfo(param); err != nil {
		gdb.string(c, 500, "%s", convertStringToByte(err.Error()), nil)
	} else {
		r, _ := json.Marshal(ResponseData{200, "", responseData})
		gdb.string(c, 200, "%s", r, nil)
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
					currentTimeStamp := int(time.Now().Unix()) + 8*3600 // current ts
					ts := strconv.Itoa(currentTimeStamp)
					batch := &leveldb.Batch{}
					m, _ := p.MemoryPercent()
					v := fmt.Sprintf("%.2f", float64(m)*float64(tm.Total)*10e-9)
					batch.Put(convertStringToByte(ram), convertStringToByte(v))                       // ram
					batch.Put(convertStringToByte(ram+fmt.Sprintf("%s", ts)), convertStringToByte(v)) // write history data of ram
					c, _ := p.CPUPercent()
					{
						batch.Put(convertStringToByte(cpu), convertStringToByte(fmt.Sprintf("%.2f", c)))                       // cpu
						batch.Put(convertStringToByte(cpu+fmt.Sprintf("%s", ts)), convertStringToByte(fmt.Sprintf("%.2f", c))) // write history data of cpu
					}
					size, err := dirSize(gdb.dbPath)
					{
						if err != nil {
							return err
						}
						batch.Put(convertStringToByte(fileSize), convertStringToByte(strconv.FormatFloat(size, 'f', 2, 64))) // fileSize
					}
					batch.Put(convertStringToByte(timeKey), convertStringToByte(ts))
					if err := gdb.systemInfoDb.Write(batch, nil); err != nil {
						return err
					}
				}
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
		startTime := time.NewTicker(60 * time.Second)
		expressions := map[string]string{} // record expressions
		for {
			select {
			case <-startTime.C:
				if flag {
					rows, _ := gdb.query("select id, expression, status, duration from calc_cfg where 1=1")
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
					rows, _ := gdb.query("select id, expression, status, duration from calc_cfg where 1=1")
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
									_, _ = gdb.updateItem("update calc_cfg set errorMessage='" + strings.Replace(err.Error(), "'", `"`, -1) + "', status='false' where id=" + strconv.Itoa(int(r.id)))
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
											_, _ = gdb.updateItem("update calc_cfg set errorMessage='" + strings.Replace(err.Error(), "'", `"`, -1) + "', status='false' where id=" + id)
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
										_, _ = gdb.updateItem("update calc_cfg set errorMessage='" + strings.Replace(err.Error(), "'", `"`, -1) + "', status='false' where id=" + id)
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
								_, _ = gdb.updateItem("update calc_cfg set errorMessage='" + strings.Replace(err.Error(), "'", `"`, -1) + "', status='false' where id=" + strconv.Itoa(int(id)))
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
			vm.Set("getFloatHData", gdb.getFloatHData)
			vm.Set("getIntHData", gdb.getIntHData)
			vm.Set("getStringHData", gdb.getStringHData)
			vm.Set("getBoolHData", gdb.getBoolHData)
			vm.Set("getFloatHDataWithTs", gdb.getFloatHDataWithTs)
			vm.Set("getIntHDataWithTs", gdb.getIntHDataWithTs)
			vm.Set("getStringHDataWithTs", gdb.getStringHDataWithTs)
			vm.Set("getBoolHDataWithTs", gdb.getBoolHDataWithTs)
			vm.Set("writeFloatRtData", gdb.writeFloatRtData)
			vm.Set("writeIntRtData", gdb.writeIntRtData)
			vm.Set("writeStringRtData", gdb.writeStringRtData)
			vm.Set("writeBoolRtData", gdb.writeBoolRtData)
			vm.Set("getTimeStamp", gdb.getUnixTimeStamp)
			vm.Set("getNowTime", gdb.getNowTime)
			vm.Set("console", gdb.console)
			program, _ := goja.Compile(id+".js", expression, false)
			_, err = vm.RunProgram(program)
		})
		if err != nil {
			return err
			//_, _ = updateItem("update calc_cfg set errorMessage='"+err.Error()+"' where id="+r["id"])
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
					if _, err := gdb.updateItem(sqlString); err != nil {
						return err
					} else {
						return nil
					}
				}
			}
		}
	}
}
