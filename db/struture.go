/*
creatTime: 2020/12/17
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

import "github.com/JustKeepSilence/gdb/cmap"

/*
The data structure and errors returned by GDB
*/

const (
	timeFormatString = "2006-01-02 15:04:05"
)

// structure

type Config struct {
	GdbConfigs `json:"gdbConfigs"`
	LogConfigs `json:"logConfigs"`
}

type GdbConfigs struct {
	Port            int64  `json:"port"`
	DbPath          string `json:"dbPath"`
	ItemDbPath      string `json:"itemDbPath"`
	IP              string `json:"ip"`
	ApplicationName string `json:"applicationName"`
	Authorization   bool   `json:"authorization"`
	Mode            string `json:"mode"`
	HttpsConfigs    `json:"httpsConfigs"`
}

type LogConfigs struct {
	LogWriting  bool   `json:"logWriting"`
	Level       string `json:"level"`
	ExpiredTime int    `json:"expiredTime"`
}

type HttpsConfigs struct {
	Ca                    bool   `json:"ca"`
	SelfSignedCa          bool   `json:"selfSignedCa"`
	CaCertificateName     string `json:"caCertificateName"`
	ServerCertificateName string `json:"serverCertificateName"`
	ServerKeyName         string `json:"serverKeyName"`
}

// ResponseData common
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Rows struct {
	EffectedRows int `json:"effectedRows"`
}

type Cols struct {
	EffectedCols int `json:"effectedCols"`
}

// groups

type AddedGroupInfo struct {
	GroupName   string   `json:"groupName" binding:"required"`
	ColumnNames []string `json:"columnNames" binding:"required"`
}

type AddedGroupInfos struct {
	GroupInfos []AddedGroupInfo `json:"groupInfos" binding:"required"`
}

type GroupNamesInfo struct {
	GroupNames []string `json:"groupNames"`
}

type GroupPropertyInfo struct {
	ItemCount       string   `json:"itemCount"`
	ItemColumnNames []string `json:"itemColumnNames"`
}

type QueryGroupPropertyInfo struct {
	GroupName string `json:"groupName" binding:"required"`
	Condition string `json:"condition" binding:"required"`
}

type UpdatedGroupNameInfo struct {
	OldGroupName string `json:"oldGroupName"`
	NewGroupName string `json:"newGroupName"`
}

type UpdatedGroupNamesInfo struct {
	Infos []UpdatedGroupNameInfo `json:"infos"`
}

type UpdatedGroupColumnNamesInfo struct {
	GroupName      string   `json:"groupName"`
	OldColumnNames []string `json:"oldColumnNames"`
	NewColumnNames []string `json:"newColumnNames"`
}

type DeletedGroupColumnNamesInfo struct {
	GroupName   string   `json:"groupName" binding:"required"`
	ColumnNames []string `json:"columnNames" binding:"required"`
}

type AddedGroupColumnsInfo struct {
	GroupName     string   `json:"groupName" binding:"required"`
	ColumnNames   []string `json:"columnNames" binding:"required"`
	DefaultValues []string `json:"defaultValues" binding:"required"`
}

// items

type AddedItemsInfo struct {
	GroupName  string              `json:"groupName" binding:"required"`
	ItemValues []map[string]string `json:"itemValues" binding:"required"`
}

type DeletedItemsInfo struct {
	GroupName string `json:"groupName"`
	Condition string `json:"condition"`
}

type ItemsInfo struct {
	GroupName   string `json:"groupName" binding:"required"`
	Condition   string `json:"condition" binding:"required"`
	Clause      string `json:"clause"`
	ColumnNames string `json:"columnNames" binding:"required"`
	StartRow    int    `json:"startRow"`
	RowCount    int    `json:"rowCount"`
}

type GdbItems struct {
	ItemValues []map[string]string `json:"itemValues"`
}

type GdbItemsWithCount struct {
	ItemCount int64 `json:"itemCount"`
	GdbItems
}

type CheckItemsInfo struct {
	GroupName string   `json:"groupName"`
	ItemNames []string `json:"itemNames"`
}

// data

type ItemValue struct {
	ItemName string `json:"itemName" binding:"required"`
	Value    string `json:"value" binding:"required"`
}

type HistoricalItemValue struct {
	ItemName   string   `json:"itemName" binding:"required"`
	Values     []string `json:"values" binding:"required"`
	TimeStamps []string `json:"timeStamps" binding:"required"`
}

type BatchWriteString struct {
	ItemValues []ItemValue `json:"itemValues" binding:"required"`
}

type BatchWriteHistoricalString struct {
	HistoricalItemValues []HistoricalItemValue `json:"historicalItemValues" binding:"required"`
}

type QueryRealTimeDataString struct {
	ItemNames []string `json:"itemNames"` // ItemNames
}

type GdbRealTimeData struct {
	RealTimeData cmap.ConcurrentMap `json:"realTimeData"`
}

type QueryHistoricalDataString struct {
	ItemNames  []string `json:"itemNames" binding:"required"`  // ItemNames
	StartTimes []int    `json:"startTimes" binding:"required"` // startTime Unix TimeStamp
	EndTimes   []int    `json:"endTimes" binding:"required"`   // endTime Unix TimeStamp
	Intervals  []int    `json:"intervals" binding:"required"`  // interval
}

type QuerySpeedHistoryDataString struct {
	ItemName   string `json:"itemName" binding:"required"`
	StartTimes []int  `json:"startTimes" binding:"required"` // startTime Unix TimeStamp
	EndTimes   []int  `json:"endTimes" binding:"required"`   // endTime Unix TimeStamp
	Interval   int    `json:"interval" binding:"required"`   // interval
}

type QueryHistoricalDataWithTimeStampString struct {
	ItemNames  []string `json:"itemNames"`  // ItemNames
	TimeStamps [][]int  `json:"timeStamps"` // time stamp
}

type DeadZone struct {
	ItemName      string `json:"itemName"`
	DeadZoneCount int    `json:"deadZoneCount"`
}

type QueryHistoricalDataWithConditionString struct {
	ItemNames       []string   `json:"itemNames"`       // ItemNames
	TimeStamps      [][]int    `json:"timeStamps"`      // time stamp
	StartTimes      []int      `json:"startTimes"`      // startTime Unix TimeStamp
	EndTimes        []int      `json:"endTimes"`        // endTime Unix TimeStamp
	Intervals       []int      `json:"intervals"`       // interval
	FilterCondition string     `json:"filterCondition"` // filter condition: item["itemNames1"] > 100
	DeadZones       []DeadZone `json:"deadZones"`       // deadZone filter condition
}

type GdbHistoricalData struct {
	HistoricalData cmap.ConcurrentMap `json:"historicalData"`
}

type GdbInfoData struct {
	Info cmap.ConcurrentMap `json:"info"`
}

// page

type httpsFile struct {
	File     []int  `json:"file" binding:"required"`
	FileName string `json:"fileName" binding:"required"`
}

type authInfo struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

type userToken struct {
	Token string `json:"token"`
}

type UserName struct {
	Name string `json:"name" binding:"required"`
}

type UserInfo struct {
	UserName `json:"userName" binding:"required"`
	Role     []string `json:"role" binding:"required"`
}

type updatedUserInfo struct {
	Id          int    `json:"id" binding:"required"`
	OldUserName string `json:"OldUserName" binding:"required"`
	NewUserName string `json:"NewUserName" binding:"required"`
	Role        string `json:"role" binding:"required"`
}

type addedUserInfo struct {
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required"`
	PassWord string `json:"passWord" binding:"required"`
}

type queryLogsInfo struct {
	Level     string `json:"level" binding:"required"`
	StartTime string `json:"startTime" binding:"required"`
	EndTime   string `json:"endTime" binding:"required"`
	StartRow  int    `json:"startRow"`
	RowCount  int    `json:"rowCount" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

type LogsInfo struct {
	Infos []map[string]string `json:"infos"`
	Count int                 `json:"count"`
}

type LogMessage struct {
	RequestUrl    string `json:"requestUrl"`
	RequestMethod string `json:"requestMethod"`
	UserAgent     string `json:"userAgent"`
	RequestBody   string `json:"requestBody"`
	RemoteAddress string `json:"remoteAddress"`
	Message       string `json:"message"`
}

// calc

type addedCalcItemInfo struct {
	Expression  string `json:"expression"`
	Flag        string `json:"flag"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
}

type CalculationResult struct {
	Result interface{} `json:"result"`
}

type queryCalcItemsInfo struct {
	Condition string `json:"condition"`
}

type CalcItemsInfo struct {
	Infos []map[string]string `json:"infos"`
}

type updatedCalcInfo struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Expression  string `json:"expression"`
	Duration    string `json:"duration"`
}

type calcId struct {
	Id []string `json:"id"`
}

type fileInfo struct {
	FileName  string `json:"fileName" binding:"required"`
	GroupName string `json:"groupName"`
}

type historyFileInfo struct {
	FileName   string   `json:"fileName"`
	ItemNames  []string `json:"itemNames"`
	SheetNames []string `json:"sheetNames"`
}

type deletedLogInfo struct {
	Id                string `json:"id"`
	StartTime         string `json:"startTime"`
	EndTime           string `json:"endTime"`
	UserNameCondition string `json:"userNameCondition"`
}

// errors, some errors imported , some not

type conditionError struct {
	ErrorInfo string
}

func (ce conditionError) Error() string {
	return ce.ErrorInfo
}

type groupNameError struct {
	ErrorInfo string
}

func (ge groupNameError) Error() string {
	return ge.ErrorInfo
}

type ColumnNameError struct {
	ErrorInfo string
}

func (cn ColumnNameError) Error() string {
	return cn.ErrorInfo
}

type snError struct {
	ErrorInfo string
}

func (se snError) Error() string {
	return se.ErrorInfo
}

type itemError struct {
	ErrorInfo string
}

func (ie itemError) Error() string {
	return ie.ErrorInfo
}

type runTimeError struct {
	ErrorInfo string
}

func (rt runTimeError) Error() string {
	return rt.ErrorInfo
}

type ExcelError struct {
	ErrorInfo string
}

func (oe ExcelError) Error() string {
	return oe.ErrorInfo
}

type userNameError struct {
	ErrorInfo string
}

func (un userNameError) Error() string {
	return un.ErrorInfo
}

type connectionDBError struct {
	ErrorInfo string
}

func (cd connectionDBError) Error() string {
	return cd.ErrorInfo
}

type updateItemError struct {
	ErrorInfo string
}

func (ui updateItemError) Error() string {
	return ui.ErrorInfo
}
