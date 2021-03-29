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
	if err := db.InitialDbServer(); err != nil {
		log.Print(fmt.Errorf("System initialization failed: " + err.Error()))
		time.Sleep(60 * time.Second)
		return
	}
}
