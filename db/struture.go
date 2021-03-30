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
	LogWriting  bool     `json:"logWriting"`
	Level       logLevel `json:"level"`
	ExpiredTime int      `json:"expiredTime"`
}

type HttpsConfigs struct {
	Ca                    bool   `json:"ca"`
	SelfSignedCa          bool   `json:"selfSignedCa"`
	CaCertificateName     string `json:"caCertificateName"`
	ServerCertificateName string `json:"serverCertificateName"`
}

// common
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
	GroupName   string   `json:"groupNames" binding:"required"`
	ColumnNames []string `json:"columnNames" binding:"required"`
}

type AddedGroupColumnsInfo struct {
	GroupName     string   `json:"groupName"`
	ColumnNames   []string `json:"columnNames"`
	DefaultValues []string `json:"defaultValues"`
}

// items

type AddedItemsInfo struct {
	GroupName string `json:"groupName"`
	GdbItems
}

type DeletedItemsInfo struct {
	GroupName string `json:"groupName"`
	Condition string `json:"condition"`
}

type ItemsInfo struct {
	ItemsInfoWithoutRow
	ColumnNames string `json:"columnNames"`
	StartRow    int    `json:"startRow"`
	RowCount    int    `json:"rowCount"`
}

type ItemsInfoWithoutRow struct {
	GroupName string `json:"groupName"`
	Condition string `json:"condition"`
	Clause    string `json:"clause"`
}

type GdbItems struct {
	ItemValues []map[string]string `json:"itemValues"`
}

type GdbItemsWithCount struct {
	ItemCount int64 `json:"itemCount"`
	GdbItems
}

// data

type ItemValue struct {
	ItemName  string `json:"itemName"`
	Value     string `json:"value"`
	TimeStamp string `json:"timeStamp"`
}

type BatchWriteString struct {
	GroupName     string      `json:"groupName"`
	ItemValues    []ItemValue `json:"itemValues"`
	WithTimeStamp bool        `json:"withTimeStamp"`
}

type QueryRealTimeDataString struct {
	ItemNames []string `json:"itemNames"` // ItemNames
}

type GdbRealTimeData struct {
	RealTimeData cmap.ConcurrentMap `json:"realTimeData"`
}

type QueryHistoricalDataString struct {
	ItemNames  []string `json:"itemNames"`  // ItemNames
	StartTimes []int32  `json:"startTimes"` // startTime Unix TimeStamp
	EndTimes   []int32  `json:"endTimes"`   // endTime Unix TimeStamp
	Intervals  []int32  `json:"intervals"`  // interval
}

type QueryHistoricalDataWithTimeStampString struct {
	ItemNames  []string  `json:"itemNames"`  // ItemNames
	TimeStamps [][]int32 `json:"timeStamps"` // time stamp
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

type authInfo struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

type userToken struct {
	Token string `json:"token"`
}

type UserName struct {
	Name string `json:"name"`
}

type UserInfo struct {
	UserName
	Role []string `json:"role"`
}

type gdbUserInfo struct {
	PassWord string   `json:"passWord"`
	Roles    []string `json:"roles"`
}

type queryLogsInfo struct {
	LogType   string `json:"logType"`
	Condition string `json:"condition"` // used to search according to message
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type LogsInfo struct {
	Infos []map[string]string `json:"infos"`
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
	FileName  string `json:"fileName"`
	GroupName string `json:"groupName"`
}

type logLevel int

const (
	Info logLevel = iota
	Error
)

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
