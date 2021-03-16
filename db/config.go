/*
creatTime: 2020/11/7 23:22
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package db

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func ReadDbConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	fileInfo, _ := os.Stat(path)
	b := make([]byte, fileInfo.Size())
	_, _ = file.Read(b)
	c := Config{}
	_ = json.Unmarshal(b, &c)
	return c, nil
}

func GetLocalIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Print(fmt.Errorf("System initialization failed: " + err.Error()))
		time.Sleep(60 * time.Second)
		return ""
	}
	defer conn.Close()
	local := conn.LocalAddr().(*net.UDPAddr)
	return local.IP.String()
}
