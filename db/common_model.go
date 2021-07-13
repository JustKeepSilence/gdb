/*
createTime: 2021/6/28
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

// common model in gin

import (
	"github.com/JustKeepSilence/gdb/cmap"
)

/*
The data structure and errors returned by GDB
*/

// ResponseData common
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TimeCols struct {
	EffectedCols int   `json:"effectedCols"`
	Times        int64 `json:"times"`
}

type TimeMap struct {
	Result cmap.ConcurrentMap `json:"result"`
	Times  int64              `json:"times"`
}

type TimeRows struct {
	EffectedRows int   `json:"effectedRows"`
	Times        int64 `json:"times"`
}
