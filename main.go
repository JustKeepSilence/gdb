package main

import (
	"fmt"
	"github.com/JustKeepSilence/gdb/config"
	"github.com/JustKeepSilence/gdb/sqlite"
	"github.com/JustKeepSilence/gdb/utils"
	"github.com/JustKeepSilence/gdb/web"
	"time"
)

func main() {
	startReadConfigTime := time.Now()
	dbConfigs, err := config.ReadDbConfig("./configs/config.json")
	if err != nil {
		utils.WriteError("", "", "", err.Error())
		return
	}
	dbPath := dbConfigs.Path
	var dbIp string
	dbIp = dbConfigs.IP
	dbPort := dbConfigs.Port
	if len(dbIp) == 0 {
		// not config
		dbIp = utils.GetLocalIp()
	}
	if err := sqlite.InitialSQLite(); err != nil {
		utils.WriteError("", "", "", err.Error())
		return
	}
	endReadConfigTimer := time.Now()
	fmt.Printf("%s: Read the configuration successfully, time consuming :%dms\n", time.Now().Format(utils.TimeFormatString),
		endReadConfigTimer.Sub(startReadConfigTime).Milliseconds())
	if err := web.InitialDbServer(dbIp, dbPort, dbPath, startReadConfigTime); err != nil {
		utils.WriteError("", "", "", err.Error())
		return
	}
}
