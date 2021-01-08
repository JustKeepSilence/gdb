/*
creatTime: 2020/11/14 21:28
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	logPath = "./logs/gdb_log.log"
)

func WriteInfo(info string)  {
	f, _ := os.OpenFile(logPath, os.O_APPEND, 0666)
	log.SetOutput(f)
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
	}).Info(info)
}

func WriteError(requestUrl string, requestMethod string, requestString string, content string)  {
	f, _ := os.OpenFile(logPath, os.O_APPEND, 0666)
	log.SetOutput(f)
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"url": requestUrl,
		"method": requestMethod,
		"requestString": requestString,
	}).Error(content)
	fmt.Println(content)
}
