/*
createTime: 2021/6/25
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

// group model in gin

type AddedGroupInfo struct {
	GroupName   string   `json:"groupName" binding:"required"`
	ColumnNames []string `json:"columnNames" binding:"required"`
}

type addedGroupInfos struct {
	GroupInfos []AddedGroupInfo `json:"groupInfos" binding:"required"`
}

type GroupNamesInfo struct {
	GroupNames []string `json:"groupNames" binding:"required"`
}

type queryGroupPropertyInfo struct {
	GroupName string `json:"groupName" binding:"required"`
	Condition string `json:"condition" binding:"required"`
}

type UpdatedGroupNameInfo struct {
	OldGroupName string `json:"oldGroupName" binding:"required"`
	NewGroupName string `json:"newGroupName" binding:"required"`
}

type updatedGroupNamesInfo struct {
	Infos []UpdatedGroupNameInfo `json:"infos" binding:"required"`
}

type AddedGroupColumnsInfo struct {
	GroupName     string   `json:"groupName" binding:"required"`
	ColumnNames   []string `json:"columnNames" binding:"required"`
	DefaultValues []string `json:"defaultValues" binding:"required"`
}

type DeletedGroupColumnNamesInfo struct {
	GroupName   string   `json:"groupName" binding:"required"`
	ColumnNames []string `json:"columnNames" binding:"required"`
}

type GroupPropertyInfo struct {
	ItemCount       string   `json:"itemCount"`
	ItemColumnNames []string `json:"itemColumnNames"`
}

type UpdatedGroupColumnNamesInfo struct {
	GroupName      string   `json:"groupName" binding:"required"`
	OldColumnNames []string `json:"oldColumnNames" binding:"required"`
	NewColumnNames []string `json:"newColumnNames" binding:"required"`
}
