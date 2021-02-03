package main

import (
	"fmt"
	"github.com/JustKeepSilence/gdb/config"
	"github.com/JustKeepSilence/gdb/utils"
	"github.com/JustKeepSilence/gdb/web"
	"os"
	"time"
)

func main() {
	startReadConfigTime := time.Now()
	dbConfigs, err := config.ReadDbConfig("./config.json")
	if err != nil {
		utils.WriteError("", "", "", err.Error())
		return
	}
	dbPath, itemDbPath := dbConfigs.DbPath, dbConfigs.ItemDbPath
	var dbIp string
	dbIp = dbConfigs.IP
	dbPort := dbConfigs.Port
	if len(dbIp) == 0 {
		// not config
		dbIp = utils.GetLocalIp()
	}
	// generate log file
	_, err = os.Lstat("./logs/gdb_log.log")
	if !os.IsExist(err) {
		_ = os.MkdirAll("./logs", 0766)
		_, _ = os.Create("./logs/gdb_log.log")
	}
	// generate uploadFiles folder
	_, err = os.Lstat("./uploadFiles")
	if !os.IsExist(err) {
		_ = os.MkdirAll("./uploadFiles", 0766)
	}
	endReadConfigTimer := time.Now()
	fmt.Printf("%s: Read the configuration successfully, time consuming :%dms\n", time.Now().Format(utils.TimeFormatString),
		endReadConfigTimer.Sub(startReadConfigTime).Milliseconds())
	if err := web.InitialDbServer(dbIp, dbPort, dbPath, itemDbPath, startReadConfigTime); err != nil {
		utils.WriteError("", "", "", err.Error())
		return
	}
}
