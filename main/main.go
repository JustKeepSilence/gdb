package main

import (
	"encoding/json"
	"fmt"
	"github.com/JustKeepSilence/gdb/db"
	"github.com/JustKeepSilence/gdb/utils"
	"log"
	"os"
	"time"
)

func main() {
	startReadConfigTime := time.Now()
	dbConfigs, err := readDbConfig("./config.json")
	if err != nil {
		utils.WriteError("", "", "", err.Error())
		return
	}
	dbPath, itemDbPath, authorization := dbConfigs.DbPath, dbConfigs.ItemDbPath, dbConfigs.Authorization
	var dbIp string
	dbIp = dbConfigs.IP
	dbPort := dbConfigs.Port
	if len(dbIp) == 0 {
		// not config
		dbIp = utils.GetLocalIp()
	}
	// generate uploadFiles folder
	_, err = os.Lstat("./uploadFiles")
	if !os.IsExist(err) {
		_ = os.MkdirAll("./uploadFiles", 0766)
	}
	endReadConfigTimer := time.Now()
	fmt.Printf("%s: Read the configuration successfully, time consuming :%dms\n", time.Now().Format(utils.TimeFormatString),
		endReadConfigTimer.Sub(startReadConfigTime).Milliseconds())
	if err := db.InitialDbServer(dbIp, dbPort, dbPath, itemDbPath, startReadConfigTime, authorization); err != nil {
		//utils.WriteError("", "", "", err.Error())
		log.Print(fmt.Errorf("System initialization failed: " + err.Error()))
		time.Sleep(60 * time.Second)
	}
}

type config struct {
	Port            int64  `json:"port"`
	DbPath          string `json:"dbPath"`
	ItemDbPath      string `json:"itemDbPath"`
	IP              string `json:"ip"`
	ApplicationName string `json:"applicationName"`
	Authorization   bool   `json:"authorization"`
}

// get configs of db
func readDbConfig(path string) (config, error) {
	file, err := os.Open(path)
	if err != nil {
		return config{}, err
	}
	defer file.Close()
	fileInfo, _ := os.Stat(path)
	b := make([]byte, fileInfo.Size())
	_, _ = file.Read(b)
	c := config{}
	_ = json.Unmarshal(b, &c)
	return c, nil
}
