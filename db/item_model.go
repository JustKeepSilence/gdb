/*
createTime: 2021/6/25
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

// item model in gin

type AddedItemsInfo struct {
	GroupName  string              `json:"groupName" binding:"required"`
	ItemValues []map[string]string `json:"itemValues" binding:"required"`
}

type DeletedItemsInfo struct {
	GroupName string `json:"groupName" binding:"required"`
	Condition string `json:"condition" binding:"required"`
}

type ItemsInfo struct {
	GroupName   string `json:"groupName" binding:"required"`
	Condition   string `json:"condition" binding:"required"`
	ColumnNames string `json:"columnNames" binding:"required"`
	StartRow    int32  `json:"startRow"`
	RowCount    int32  `json:"rowCount"`
}

type gdbItemsWithCount struct {
	ItemCount int32 `json:"itemCount"`
	GdbItems
}

type GdbItems struct {
	ItemValues []map[string]string `json:"itemValues"`
}

type checkItemsInfo struct {
	GroupName string   `json:"groupName"`
	ItemNames []string `json:"itemNames"`
}

type UpdatedItemsInfo struct {
	GroupName string `json:"groupName" binding:"required"`
	Condition string `json:"condition" binding:"required"`
	Clause    string `json:"clause" binding:"required"`
}
