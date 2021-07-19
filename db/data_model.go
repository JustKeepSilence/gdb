/*
createTime: 2021/6/25
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import "github.com/JustKeepSilence/gdb/cmap"

// data model in gin

type floatItemValues struct {
	GroupNames []string    `json:"groupNames" binding:"required"`
	ItemNames  [][]string  `json:"itemNames" binding:"required"`
	ItemValues [][]float32 `json:"itemValues" binding:"required"`
}

type intItemValues struct {
	GroupNames []string   `json:"groupNames" binding:"required"`
	ItemNames  [][]string `json:"itemNames" binding:"required"`
	ItemValues [][]int32  `json:"itemValues" binding:"required"`
}

type boolItemValues struct {
	GroupNames []string   `json:"groupNames" binding:"required"`
	ItemNames  [][]string `json:"itemNames" binding:"required"`
	ItemValues [][]bool   `json:"itemValues" binding:"required"`
}

type stringItemValues struct {
	GroupNames []string   `json:"groupNames" binding:"required"`
	ItemNames  [][]string `json:"itemNames" binding:"required"`
	ItemValues [][]string `json:"itemValues" binding:"required"`
}

type floatHItemValues struct {
	GroupNames []string    `json:"groupNames" binding:"required"`
	ItemNames  []string    `json:"itemNames" binding:"required"`
	ItemValues [][]float32 `json:"itemValues" binding:"required"`
	TimeStamps [][]int32   `json:"timeStamps" binding:"required"`
}

type intHItemValues struct {
	GroupNames []string  `json:"groupNames" binding:"required"`
	ItemNames  []string  `json:"itemNames" binding:"required"`
	ItemValues [][]int32 `json:"itemValues" binding:"required"`
	TimeStamps [][]int32 `json:"timeStamps" binding:"required"`
}

type stringHItemValues struct {
	GroupNames []string   `json:"groupNames" binding:"required"`
	ItemNames  []string   `json:"itemNames" binding:"required"`
	ItemValues [][]string `json:"itemValues" binding:"required"`
	TimeStamps [][]int32  `json:"timeStamps" binding:"required"`
}

type boolHItemValues struct {
	GroupNames []string  `json:"groupNames" binding:"required"`
	ItemNames  []string  `json:"itemNames" binding:"required"`
	ItemValues [][]bool  `json:"itemValues" binding:"required"`
	TimeStamps [][]int32 `json:"timeStamps" binding:"required"`
}

type queryRealTimeDataString struct {
	GroupNames []string `json:"groupNames" binding:"required"`
	ItemNames  []string `json:"itemNames" binding:"required"` // ItemNames
}

type queryHistoricalDataString struct {
	GroupNames []string `json:"groupNames" binding:"required"`
	ItemNames  []string `json:"itemNames" binding:"required"`  // ItemNames
	StartTimes []int32  `json:"startTimes" binding:"required"` // startTime Unix TimeStamp
	EndTimes   []int32  `json:"endTimes" binding:"required"`   // endTime Unix TimeStamp
	Intervals  []int32  `json:"intervals" binding:"required"`  // interval
}

type queryRawHistoricalDataString struct {
	GroupNames []string `json:"groupNames" binding:"required"`
	ItemNames  []string `json:"itemNames" binding:"required"` // ItemNames
}

type queryHistoricalDataWithStampString struct {
	GroupNames []string  `json:"groupNames" binding:"required"`
	ItemNames  []string  `json:"itemNames" binding:"required"` // ItemNames
	TimeStamps [][]int32 `json:"timeStamps" binding:"required"`
}

type DeadZone struct {
	ItemName      string `json:"itemName"`
	DeadZoneCount int32  `json:"deadZoneCount"`
}

type queryHistoricalDataWithConditionString struct {
	GroupName       string     `json:"groupName" binding:"required"`
	ItemNames       []string   `json:"itemNames" binding:"required"`
	StartTime       int32      `json:"startTime" binding:"required"`
	EndTime         int32      `json:"endTime" binding:"required"`
	Interval        int32      `json:"interval" binding:"required"`
	FilterCondition string     `json:"filterCondition" binding:"required"`
	DeadZones       []DeadZone `json:"deadZones" binding:"required"`
}

type deleteHistoricalDataString struct {
	GroupNames []string `json:"groupNames" binding:"required"`
	ItemNames  []string `json:"itemNames" binding:"required"`
	StartTimes []int32  `json:"startTimes" binding:"required"`
	EndTimes   []int32  `json:"endTimes" binding:"required"`
}

type GdbRealTimeData struct {
	RealTimeData cmap.ConcurrentMap `json:"realTimeData"`
	Times        int64              `json:"times"`
}

type GdbHistoricalData struct {
	HistoricalData cmap.ConcurrentMap `json:"historicalData"`
	Times          int64              `json:"times"`
}
