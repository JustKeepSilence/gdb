// +build gdbServer

/*
createTime: 2021/6/27
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

// calc model in gin

type addedCalcItemInfo struct {
	Expression  string `json:"expression" binding:"required"`
	Flag        string `json:"flag"`
	Duration    string `json:"duration" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type queryCalcItemsInfo struct {
	Condition string `json:"condition"`
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

type calculationResult struct {
	Result interface{} `json:"result"`
}

type calcItemsInfo struct {
	Infos []map[string]string `json:"infos"`
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
