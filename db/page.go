/*
creatTime: 2020/12/10
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	. "github.com/ahmetb/go-linq/v3"
	"os"
	"regexp"
	"strings"
)

// user login
func (gdb *Gdb) userLogin(info authInfo) error {
	userName := info.UserName
	v, err := gdb.infoDb.Get([]byte(userName), nil)
	if err != nil {
		return userNameError{"userNameError: " + userName}
	} else {
		ui := userInfo{}
		err := Json.Unmarshal(v, &ui)
		if err != nil {
			return fmt.Errorf("fail parsing userInfo: " + err.Error())
		} else {
			if fmt.Sprintf("%s", ui.PassWord) != info.PassWord {
				return userNameError{"userNameError: " + userName}
			}
		}
	}
	return nil
}

func (gdb *Gdb) getUserInfo(userName string) (map[string]interface{}, error) {
	v, err := gdb.infoDb.Get([]byte(userName), nil)
	if err != nil {
		return nil, userNameError{"userNameError: " + userName}
	} else {
		ui := userInfo{}
		err := Json.Unmarshal(v, &ui)
		if err != nil {
			return nil, fmt.Errorf("fail parsing userInfo: " + err.Error())
		} else {
			return map[string]interface{}{"userName": userName, "role": ui.Roles}, nil
		}
	}
}

// add items by excel ,web app
func (gdb *Gdb) addItemsByExcel(info fileInfo) (Rows, error) {
	fileName, groupName := info.FileName, info.GroupName
	f, err := excelize.OpenFile("./uploadFiles/" + fileName) // open excel
	if err != nil {
		return Rows{-1}, ExcelError{"ExcelError: " + err.Error()}
	} else {
		// open excel successfully
		sheetName := f.GetSheetList()[0] // use first worksheet
		rows, err := f.Rows(sheetName)   // get all rows
		var headers []string             // headers
		var items AddItemInfo
		var values []map[string]string
		if err != nil {
			return Rows{-1}, ExcelError{"ExcelError: " + err.Error()}
		} else {
			// get rows successfully
			count := 0
			for rows.Next() {
				if count == 0 {
					// check headers
					headers, err = rows.Columns()
					if err != nil {
						return Rows{-1}, ExcelError{"ExcelError: " + err.Error()}
					} else {
						// get headers successfully
						cols, err := gdb.GetGroupProperty([]string{groupName}...)
						if err != nil {
							return Rows{-1}, err
						}
						tableHeaders := cols[groupName].ItemColumnNames
						var h []string
						From(headers).Where(func(item interface{}) bool {
							return len(item.(string)) != 0
						}).ToSlice(&h)
						if !equal(h, tableHeaders) {
							return Rows{-1}, ExcelError{"ExcelError: Inconsistent header"}
						}
					}
				} else {
					c, _ := rows.Columns()
					value := map[string]string{}
					if len(c) < len(headers) {
						// see: https://github.com/360EntSecGroup-Skylar/excelize/issues/721
						e := len(headers) - len(c)
						for i := 0; i < e; i++ {
							c = append(c, " ")
						}
					}
					for i := 0; i < len(headers); i++ {
						value[headers[i]] = c[i]
					}
					values = append(values, value)
				}
				count++
			}
		}
		items.GroupName = groupName
		items.Values = values
		if r, err := gdb.AddItems(items); err != nil {
			return Rows{-1}, err
		} else {
			return r, nil
		}
	}
}

func getJsCode(fileName string) (string, error) {
	filePath := "./uploadFiles/" + fileName
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		return "", err
	}
	fileInfo, _ := os.Stat(filePath)
	b := make([]byte, fileInfo.Size())
	_, _ = file.Read(b)
	c := strings.Replace(fmt.Sprintf("%s", b), "\r\n", "<br />", -1)
	c1 := strings.Replace(c, "\n", "<br />", -1)
	return c1, nil
}

func getLogs() ([]logsInfo, error) {
	const logPath = "./logs/gdb_log.log"
	f, err := os.OpenFile(logPath, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	info, _ := f.Stat()
	content := make([]byte, info.Size())
	n, err := f.Read(content)
	if err != nil {
		return nil, err
	}
	if n != int(info.Size()) {
		return nil, fmt.Errorf("read Error")
	}
	reg := regexp.MustCompile("(?is){(.*?)}")
	if !reg.Match(content) {
		return nil, nil
	}
	matchedResults := reg.FindAllString(fmt.Sprintf("%s", content), -1)
	var result []logsInfo
	for _, matchedResult := range matchedResults {
		r := logsInfo{}
		_ = Json.Unmarshal([]byte(matchedResult), &r)
		result = append(result, r)
	}
	return result, nil
}

func equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if strings.Trim(a[i], " ") != b[i] {
			return false
		}
	}
	return true
}
