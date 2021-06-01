/*
creatTime: 2020/12/17
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
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

type queryGroupPropertyInfo struct {
	GroupName string `json:"groupName" binding:"required"`
	Condition string `json:"condition" binding:"required"`
}

type UpdatedGroupNameInfo struct {
	OldGroupName string `json:"oldGroupName" binding:"required"`
	NewGroupName string `json:"newGroupName" binding:"required"`
}

type UpdatedGroupNamesInfo struct {
	Infos []UpdatedGroupNameInfo `json:"infos" binding:"required"`
}

type UpdatedGroupColumnNamesInfo struct {
	GroupName      string   `json:"groupName" binding:"required"`
	OldColumnNames []string `json:"oldColumnNames" binding:"required"`
	NewColumnNames []string `json:"newColumnNames" binding:"required"`
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
	ColumnNames string `json:"columnNames" binding:"required"`
	StartRow    int    `json:"startRow"`
	RowCount    int    `json:"rowCount"`
}

type UpdatedItemsInfo struct {
	GroupName string `json:"groupName" binding:"required"`
	Condition string `json:"condition" binding:"required"`
	Clause    string `json:"clause" binding:"required"`
}

type GdbItems struct {
	ItemValues []map[string]string `json:"itemValues"`
}

type gdbItemsWithCount struct {
	ItemCount int64 `json:"itemCount"`
	GdbItems
}

type CheckItemsInfo struct {
	GroupName string   `json:"groupName"`
	ItemNames []string `json:"itemNames"`
}

// data

type ItemValue struct {
	GroupName string      `json:"GroupName" binding:"required"`
	ItemName  string      `json:"itemName" binding:"required"`
	Value     interface{} `json:"value" binding:"required"`
}

type HistoricalItemValue struct {
	GroupName  string        `json:"groupName" binding:"required"`
	ItemName   string        `json:"itemName" binding:"required"`
	Values     []interface{} `json:"values" binding:"required"`
	TimeStamps []int         `json:"timeStamps" binding:"required"`
}

type batchWriteString struct {
	ItemValues []ItemValue `json:"itemValues" binding:"required"`
}

type batchWriteHistoricalString struct {
	HistoricalItemValues []HistoricalItemValue `json:"historicalItemValues" binding:"required"`
}

type queryRealTimeDataString struct {
	GroupNames []string `json:"groupNames"`
	ItemNames  []string `json:"itemNames"` // ItemNames
}

type gdbRealTimeData struct {
	RealTimeData cmap.ConcurrentMap `json:"realTimeData"`
}

type queryHistoricalDataString struct {
	GroupNames []string `json:"groupNames" binding:"required"`
	ItemNames  []string `json:"itemNames" binding:"required"`  // ItemNames
	StartTimes []int    `json:"startTimes" binding:"required"` // startTime Unix TimeStamp
	EndTimes   []int    `json:"endTimes" binding:"required"`   // endTime Unix TimeStamp
	Intervals  []int    `json:"intervals" binding:"required"`  // interval
}

type querySpeedHistoryDataString struct {
	ItemName   string `json:"itemName" binding:"required"`
	StartTimes []int  `json:"startTimes" binding:"required"` // startTime Unix TimeStamp
	EndTimes   []int  `json:"endTimes" binding:"required"`   // endTime Unix TimeStamp
	Intervals  []int  `json:"intervals" binding:"required"`  // interval
}

type queryHistoricalDataWithTimeStampString struct {
	GroupNames []string `json:"groupNames" binding:"required"`
	ItemNames  []string `json:"itemNames" binding:"required"`  // ItemNames
	TimeStamps [][]int  `json:"timeStamps" binding:"required"` // time stamp
}

type DeadZone struct {
	ItemName      string `json:"itemName" binding:"required"`
	DeadZoneCount int    `json:"deadZoneCount" binding:"required"`
}

type queryHistoricalDataWithConditionString struct {
	GroupNames      []string   `json:"groupNames" binding:"required"`
	ItemNames       []string   `json:"itemNames" binding:"required"`       // ItemNames
	StartTimes      []int      `json:"startTimes" binding:"required"`      // startTime Unix TimeStamp
	EndTimes        []int      `json:"endTimes" binding:"required"`        // endTime Unix TimeStamp
	Intervals       []int      `json:"intervals" binding:"required"`       // interval
	FilterCondition string     `json:"filterCondition" binding:"required"` // filter condition: item["itemNames1"] > 100
	DeadZones       []DeadZone `json:"deadZones"`                          // deadZone filter condition
}

type gdbHistoricalData struct {
	HistoricalData cmap.ConcurrentMap `json:"historicalData"`
}

type gdbInfoData struct {
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

type userName struct {
	Name string `json:"name" binding:"required"`
}

type userInfo struct {
	UserName string   `json:"userName" binding:"required"`
	Role     []string `json:"role" binding:"required"`
}

type updatedUserInfo struct {
	UserName    string `json:"userName" binding:"required"`
	NewUserName string `json:"newUserName"`
	NewPassWord string `json:"newPassWord"`
	NewRole     string `json:"newRole"`
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

type logsInfo struct {
	Infos []map[string]string `json:"infos"`
	Count int                 `json:"count"`
}

type logMessage struct {
	RequestUrl    string `json:"requestUrl"`
	RequestMethod string `json:"requestMethod"`
	UserAgent     string `json:"userAgent"`
	RequestBody   string `json:"requestBody"`
	RemoteAddress string `json:"remoteAddress"`
	Message       string `json:"message"`
}

// calc

type addedCalcItemInfo struct {
	Expression  string `json:"expression" binding:"required"`
	Flag        string `json:"flag"`
	Duration    string `json:"duration" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type calculationResult struct {
	Result interface{} `json:"result"`
}

type queryCalcItemsInfo struct {
	Condition string `json:"condition"`
}

type calcItemsInfo struct {
	Infos []map[string]string `json:"infos"`
}

type updatedCalcInfo struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Expression  string `json:"expression"`
	Duration    string `json:"duration"`
	UpdatedTime string
}

type calcId struct {
	Id []string `json:"id"`
}

type calcConfig struct {
	id         int64        // calc item id
	f          func() error // function to invoke js code
	expression string
	status     bool // whether to calc
	duration   int64
}

type messageCalcConfig struct {
	updatedInfos []updatedInfo
	addedInfos   []calcConfig
}

type updatedInfo struct {
	id           string
	newStatus    bool
	newDuration  int64
	updatedFiled string
	f            func() error
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

type deletedLogInfo struct {
	Id                string `json:"id"`
	StartTime         string `json:"startTime"`
	EndTime           string `json:"endTime"`
	UserNameCondition string `json:"userNameCondition"`
}

type routesInfo struct {
	Name   string   `json:"name" binding:"required"`
	Routes []string `json:"routes" binding:"required"`
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

type columnNameError struct {
	ErrorInfo string
}

func (cn columnNameError) Error() string {
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

type excelError struct {
	ErrorInfo string
}

func (oe excelError) Error() string {
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
