/*
creatTime: 2020/12/17
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

/*
The data structure and errors returned by GDB
*/

const (
	timeFormatString = "2006-01-02 15:04:05"
)

// structure

type Config struct {
	Port            int64  `json:"port"`
	DbPath          string `json:"dbPath"`
	ItemDbPath      string `json:"itemDbPath"`
	IP              string `json:"ip"`
	ApplicationName string `json:"applicationName"`
	Authorization   bool   `json:"authorization"`
}

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// function return effectedRow
// {"effectedRows": 1}
type Rows struct {
	EffectedRows int `json:"effectedRows"`
}

// function return effectedCol
// {"effectedCols": 1}
type Cols struct {
	EffectedCols int `json:"effectedCols"`
}

type Items struct {
	ItemCount  int64               `json:"itemCount"`
	ItemValues []map[string]string `json:"itemValues"`
}

// function return GroupName, every item of slice is the name of group
// {"groupNames": ["1DCS", "2DCS"]}
type GroupNameInfos struct {
	GroupNames []string `json:"groupNames"`
}

// function return GroupProperty {"itemCount": 10, "itemColumnNames": ["units", "type"]}
type GroupPropertyInfo struct {
	ItemCount       string   `json:"itemCount"`
	ItemColumnNames []string `json:"itemColumnNames"`
}

type GetGroupPropertyInfo struct {
	GroupName string `json:"groupName" binding:"required"`
	Condition string `json:"condition" binding:"required"`
}

type AddGroupInfo struct {
	GroupName   string   `json:"groupName" binding:"required"`
	ColumnNames []string `json:"columnNames" binding:"required"`
}

type AddGroupInfos struct {
	GroupInfos []AddGroupInfo `json:"groupInfos" binding:"required"`
}

type DeletedGroupInfo struct {
	GroupNames  []string `json:"groupNames" binding:"required"`
	ColumnNames string   `json:"columnNames" binding:"required"`
}

type UpdatedGroupInfo struct {
	OldGroupName string `json:"oldGroupName"`
	NewGroupName string `json:"newGroupName"`
}

type UpdatedGroupColumnInfo struct {
	GroupName      string   `json:"groupName"`
	OldColumnNames []string `json:"oldColumnNames"`
	NewColumnNames []string `json:"newColumnNames"`
}

type DeleteGroupColumnInfo struct {
	GroupName   string   `json:"groupName"`
	ColumnNames []string `json:"columnNames"`
}

type AddGroupColumnInfo struct {
	GroupName     string   `json:"groupName"`
	ColumnNames   []string `json:"columnNames"`
	DefaultValues []string `json:"defaultValues"`
}

type DeletedHistoricalDataInfo struct {
	ItemNames  []string    `json:"itemNames"`
	TimeStamps []TimeStamp `json:"timeStamps"`
}

type TimeStamp struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type HistoricalDataInfo struct {
	ItemNames  []string `json:"itemNames"`  // ItemNames
	StartTimes []int    `json:"startTimes"` // startTime Unix TimeStamp
	EndTimes   []int    `json:"endTimes"`   // endTime Unix TimeStamp
	Intervals  []int    `json:"intervals"`  // interval
}

type HistoricalDataInfoWithCondition struct {
	ItemNames       []string   `json:"itemNames"`       // ItemNames
	TimeStamps      [][]int    `json:"timeStamps"`      // time stamp
	StartTimes      []int      `json:"startTimes"`      // startTime Unix TimeStamp
	EndTimes        []int      `json:"endTimes"`        // endTime Unix TimeStamp
	Intervals       []int      `json:"intervals"`       // interval
	FilterCondition string     `json:"filterCondition"` // filter condition: item["itemNames1"] > 100
	DeadZones       []DeadZone `json:"deadZones"`       // deadZone filter condition
}

type HistoricalDataInfoWithTimeStamp struct {
	ItemNames  []string `json:"itemNames"`  // ItemNames
	TimeStamps [][]int  `json:"timeStamps"` // time stamp
}

type DeadZone struct {
	ItemName      string `json:"itemName"`
	DeadZoneCount int    `json:"deadZoneCount"`
}

type calcInfo struct {
	Expression  string `json:"expression"`
	Flag        string `json:"flag"`
	Duration    string `json:"duration"`
	CreateTime  string `json:"createTime"`
	UpdatedTime string `json:"updatedTime"`
	Description string `json:"description"`
}

type AddItemInfo struct {
	GroupName string              `json:"groupName"`
	Values    []map[string]string `json:"values"`
}

type ItemInfo struct {
	GroupName string `json:"groupName"`
	Column    string `json:"column"`
	StartRow  int    `json:"startRow"`
	RowCount  int    `json:"rowCount"`
	Condition string `json:"condition"`
	Clause    string `json:"clause"`
}

type authInfo struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

type BatchWriteString struct {
	GroupName     string      `json:"groupName"`
	ItemValues    []ItemValue `json:"itemValues"`
	WithTimeStamp bool        `json:"withTimeStamp"`
}

type ItemValue struct {
	ItemName  string `json:"itemName"`
	Value     string `json:"value"`
	TimeStamp string `json:"timeStamp"`
}

type RealTimeDataString struct {
	ItemNames []string `json:"itemNames"` // ItemNames
}

type fileInfo struct {
	FileName  string `json:"fileName"`
	GroupName string `json:"groupName"`
}

type userInfo struct {
	PassWord string   `json:"passWord"`
	Roles    []string `json:"roles"`
}

type updatedCalculationInfo struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Expression  string `json:"expression"`
	Duration    string `json:"duration"`
}

type getLogsInfo struct {
	LogType   string `json:"logType"`
	Condition string `json:"condition"` // used to search according to message
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type userToken struct {
	Token string `json:"token"`
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
