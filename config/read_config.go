/*
creatTime: 2020/11/7 23:22
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port            int64  `json:"port"`
	DbPath          string `json:"dbPath"`
	ItemDbPath      string `json:"itemDbPath"`
	IP              string `json:"ip"`
	ApplicationName string `json:"applicationName"`
	Authorization   bool   `json:"authorization"`
}

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
