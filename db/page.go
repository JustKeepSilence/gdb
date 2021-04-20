/*
creatTime: 2020/12/10
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

import (
	"crypto/md5"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"os"
	"strings"
	"time"
)

// user login
func (gdb *Gdb) userLogin(info authInfo, userAgent string) (userToken, error) {
	userName := info.UserName
	v, err := gdb.infoDb.Get([]byte(userName), nil)
	if err != nil || v == nil {
		return userToken{}, userNameError{"userNameError: " + userName}
	} else {
		ui := gdbUserInfo{}
		err := Json.Unmarshal(v, &ui)
		if err != nil {
			return userToken{}, fmt.Errorf("fail parsing userInfo: " + err.Error())
		} else {
			if fmt.Sprintf("%s", ui.PassWord) != info.PassWord {
				return userToken{}, userNameError{"userNameError: " + userName}
			} else {
				// correct userInfo, generate token
				b := []byte("seu" + time.Now().Format(timeFormatString) + "JustKeepSilence")
				token := fmt.Sprintf("%x", md5.Sum(b))                                                    // result is 32-bit lowercase
				_ = gdb.infoDb.Put([]byte(userName+"_token"+"_"+token+"_"+userAgent), []byte(token), nil) // write token to gdb
				return userToken{token}, nil
			}
		}
	}
}

func (gdb *Gdb) userLogout(userName, userAgent, token string) (Rows, error) {
	if err := gdb.infoDb.Delete([]byte(userName+"_token"+"_"+token+"_"+userAgent), nil); err != nil {
		return Rows{}, err
	} else {
		return Rows{1}, nil
	}
}

func (gdb *Gdb) getUserInfo(userName string) (UserInfo, error) {
	v, err := gdb.infoDb.Get([]byte(userName), nil)
	if err != nil {
		return UserInfo{}, userNameError{"userNameError: " + userName}
	} else {
		ui := gdbUserInfo{}
		err := Json.Unmarshal(v, &ui)
		if err != nil {
			return UserInfo{}, fmt.Errorf("fail parsing userInfo: " + err.Error())
		} else {
			return UserInfo{
				UserName: UserName{userName},
				Role:     ui.Roles,
			}, nil
		}
	}
}

// add items by excel
func (gdb *Gdb) AddItemsByExcel(groupName, filePath string) (Rows, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return Rows{-1}, ExcelError{"ExcelError: " + err.Error()}
	} else {
		// open excel successfully
		sheetName := f.GetSheetList()[0] // use first worksheet
		rows, err := f.Rows(sheetName)   // get all rows
		var headers []string             // headers
		var items AddedItemsInfo
		var values []map[string]string
		if err != nil {
			return Rows{-1}, ExcelError{"ExcelError: " + err.Error()}
		} else {
			// get rows successfully
			count := 0
			for rows.Next() {
				if count == 0 {
					// check headers
					h, err := rows.Columns() // columns of excel
					if err != nil {
						return Rows{-1}, ExcelError{"ExcelError: " + err.Error()}
					} else {
						// get headers successfully
						cols, err := gdb.GetGroupProperty(groupName, "1=1")
						if err != nil {
							return Rows{-1}, err
						}
						headers = cols.ItemColumnNames // columns of database
						if !equal(h, headers) {
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
		items.GdbItems.ItemValues = values
		if r, err := gdb.AddItems(items); err != nil {
			return Rows{-1}, err
		} else {
			return r, nil
		}
	}
}

func (gdb *Gdb) ImportHistoryByExcel(fileName string, itemNames []string, sheetNames ...string) error {
	if f, err := excelize.OpenFile(fileName); err != nil {
		return ExcelError{"ExcelError: " + err.Error()}
	} else {
		infos := []HistoricalItemValue{}
		for index := 0; index < len(itemNames); index++ {
			sheetName, itemName := sheetNames[index], itemNames[index]
			if rows, err := f.Rows(sheetName); err != nil {
				return err
			} else {
				info := HistoricalItemValue{ItemName: itemName}
				var values, timeStamps []string
				for rows.Next() {
					// first row is timeStamp, second is time
					if c, err := rows.Columns(); err != nil {
						return err
					} else {
						values = append(values, c[1])
						if t, err := time.Parse(timeFormatString, c[0]); err != nil {
							return err
						} else {
							timeStamps = append(timeStamps, fmt.Sprintf("%d", t.Unix()))
						}
					}
				}
				info.Values = values
				info.TimeStamps = timeStamps
				infos = append(infos, info)
			}
		}
		if err := gdb.BatchWriteHistoricalData(BatchWriteHistoricalString{HistoricalItemValues: infos}); err != nil {
			return err
		} else {
			return nil
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

func (gdb *Gdb) getLogs(logType, condition, startTime, endTime string) (LogsInfo, error) {
	var queryString string
	st, et := strings.Trim(startTime, " "), strings.Trim(endTime, " ")
	if logType == "all" {
		if len(st) == 0 && len(et) == 0 {
			queryString = "select * from log_cfg where logMessage like '%" + condition + "%' order by insertTime desc"
		} else if len(st) == 0 && len(et) != 0 {
			queryString = "select * from log_cfg where insertTime <= '" + et + "' and logMessage like '%" + condition + "%' order by insertTime desc"
		} else {
			queryString = "select * from log_cfg where insertTime >= '" + st + "' and logMessage like '%" + condition + "%' order by insertTime desc"
		}
	} else {
		if len(st) == 0 && len(et) == 0 {
			queryString = "select * from log_cfg where logType='" + logType + "' " + " and logMessage like '%" + condition + "%' order by insertTime desc"
		} else if len(st) == 0 && len(et) != 0 {
			queryString = "select * from log_cfg where logType='" + logType + "' " + " and insertTime <= '" + et + "' and logMessage like '%" + condition + "%' order by insertTime desc"
		} else {
			queryString = "select * from log_cfg where logType='" + logType + "' " + " and insertTime >= '" + st + "' and logMessage like '%" + condition + "%' order by insertTime desc"
		}
	}
	if result, err := query(gdb.ItemDbPath, queryString); err != nil {
		return LogsInfo{}, err
	} else {
		return LogsInfo{result}, nil
	}
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
