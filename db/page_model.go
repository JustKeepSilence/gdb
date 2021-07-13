// +build gdbClient

/*
createTime: 2021/6/27
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import "github.com/JustKeepSilence/gdb/cmap"

// item model in gin

type authInfo struct {
	UserName string `json:"userName" binding:"required"`
	PassWord string `json:"passWord" binding:"required"`
}

type userName struct {
	Name string `json:"name" binding:"required"`
}

type addedUserInfo struct {
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required"`
	PassWord string `json:"passWord" binding:"required"`
}

type updatedUserInfo struct {
	UserName    string `json:"userName" binding:"required"`
	NewUserName string `json:"newUserName"`
	NewPassWord string `json:"newPassWord"`
	NewRole     string `json:"newRole"`
}

type httpsFile struct {
	File     []int64 `json:"file" binding:"required"`
	FileName string  `json:"fileName" binding:"required"`
}

type fileInfo struct {
	FileName  string `json:"fileName" binding:"required"`
	GroupName string `json:"groupName"`
}

type historyFileInfo struct {
	GroupName  string   `json:"groupName" binding:"required"`
	FileName   string   `json:"fileName" binding:"required"`
	ItemNames  []string `json:"itemNames" binding:"required"`
	SheetNames []string `json:"sheetNames" binding:"required"`
}

type queryLogsInfo struct {
	Level     string `json:"level" binding:"required"`
	StartTime string `json:"startTime" binding:"required"`
	EndTime   string `json:"endTime" binding:"required"`
	StartRow  int32  `json:"startRow"`
	RowCount  int32  `json:"rowCount" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

type deletedLogInfo struct {
	Id                string `json:"id"`
	StartTime         string `json:"startTime"`
	EndTime           string `json:"endTime"`
	UserNameCondition string `json:"userNameCondition"`
}

type querySpeedHistoryDataString struct {
	InfoType   string  `json:"infoType" binding:"required"`
	ItemName   string  `json:"itemName" binding:"required"`
	StartTimes []int32 `json:"startTimes" binding:"required"` // startTime Unix TimeStamp
	EndTimes   []int32 `json:"endTimes" binding:"required"`   // endTime Unix TimeStamp
	Intervals  []int32 `json:"intervals" binding:"required"`  // interval
}

type routesInfo struct {
	Name   string   `json:"name" binding:"required"`
	Routes []string `json:"routes" binding:"required"`
}

type gdbInfoData struct {
	Info cmap.ConcurrentMap `json:"info"`
}

type userToken struct {
	Token string `json:"token"`
}

type userInfo struct {
	UserName string   `json:"userName" binding:"required"`
	Role     []string `json:"role" binding:"required"`
}

type logsInfo struct {
	Infos []map[string]string `json:"infos"`
	Count int64               `json:"count"`
}

type logMessage struct {
	RequestUrl    string `json:"requestUrl"`
	RequestMethod string `json:"requestMethod"`
	UserAgent     string `json:"userAgent"`
	RequestBody   string `json:"requestBody"`
	RemoteAddress string `json:"remoteAddress"`
	Message       string `json:"message"`
}
