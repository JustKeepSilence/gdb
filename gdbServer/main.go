// +build gdbServer

package main

import (
	"fmt"
	"github.com/JustKeepSilence/gdb/db"
	"log"
	"os"
	"time"
)

func main() {
	// generate uploadFiles folder
	_, err := os.Lstat("./uploadFiles")
	if !os.IsExist(err) {
		_ = os.MkdirAll("./uploadFiles", 0766)
	}
	if configs, err := db.ReadDbConfig("./config.json"); err != nil {
		log.Println(err)
		time.Sleep(60 * time.Second)
		return
	} else {
		if err := db.StartDbServer(configs); err != nil {
			log.Println(fmt.Errorf("System initialization failed: " + err.Error()))
			time.Sleep(60 * time.Second)
			return
		}
	}
}
