/*
creatTime: 2020/11/7 23:22
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"runtime"
	"strings"
	"time"
)

func ReadDbConfig(path string) (Config, error) {
	if b, err := ioutil.ReadFile(path); err != nil {
		return Config{}, err
	} else {
		c := Config{}
		cf := handleJson(fmt.Sprintf("%s", b))
		if err := json.Unmarshal(cf, &c); err != nil {
			return Config{}, err
		} else {
			return c, nil
		}
	}
}

func getLocalIp() string {
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

// handle json to allow  // single line comments in json file
func handleJson(js string) []byte {
	var lines, configs []string
	var newLine string
	switch runtime.GOOS {
	case "windows":
		newLine = "\r\n"
	case "linux":
		newLine = "\n"
	default:
		newLine = "r" // mac
	}
	lines = strings.Split(js, newLine)
code:
	for _, line := range lines {
		if strings.HasPrefix(strings.Trim(line, " "), "//") {
			continue code
		} else {
			configs = append(configs, line)
		}

	}
	return []byte(strings.Join(configs, newLine))
}
