package main

import (
	"fmt"
	"github.com/JustKeepSilence/gdb/db"
	"log"
	"os"
	"time"
)

func main() {
	startReadConfigTime := time.Now()
	dbConfigs, err := db.ReadDbConfig("./config.json")
	if err != nil {
		log.Print(fmt.Errorf("System initialization failed: " + err.Error()))
		time.Sleep(60 * time.Second)
		return
	}
	dbPath, itemDbPath, authorization := dbConfigs.DbPath, dbConfigs.ItemDbPath, dbConfigs.Authorization
	var dbIp string
	dbIp = dbConfigs.IP
	dbPort := dbConfigs.Port
	if len(dbIp) == 0 {
		// not config
		dbIp = db.GetLocalIp()
	}
	// generate uploadFiles folder
	_, err = os.Lstat("./uploadFiles")
	if !os.IsExist(err) {
		_ = os.MkdirAll("./uploadFiles", 0766)
	}
	endReadConfigTimer := time.Now()
	fmt.Printf("%s: Read the configuration successfully, time consuming :%dms\n", time.Now().Format("2006-01-02 15:04:05"),
		endReadConfigTimer.Sub(startReadConfigTime).Milliseconds())
	if err := db.InitialDbServer(dbIp, dbPort, dbPath, itemDbPath, startReadConfigTime, authorization); err != nil {
		log.Print(fmt.Errorf("System initialization failed: " + err.Error()))
		time.Sleep(60 * time.Second)
		return
	}
}
