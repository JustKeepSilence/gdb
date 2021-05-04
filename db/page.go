/*
creatTime: 2020/12/10
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io/ioutil"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// user login
func (gdb *Gdb) userLogin(info authInfo) (userToken, error) {
	userName := info.UserName
	if r, err := query(gdb.ItemDbPath, "select passWord, isLogin from user_cfg where userName='"+userName+"'"); err != nil || len(r) == 0 {
		return userToken{}, userNameError{"userNameError: " + userName}
	} else {
		if r[0]["passWord"] != info.PassWord {
			return userToken{}, userNameError{"userNameError: " + userName}
		} else {
			b := []byte(userName + "@seu" + time.Now().Format(timeFormatString) + "JustKeepSilence")
			token := fmt.Sprintf("%x", md5.Sum(b)) // result is 32-bit lowercase
			if _, err := updateItem(gdb.ItemDbPath, "update user_cfg set token='"+token+"', isLogin='true'"+" where userName='"+userName+"'"); err != nil {
				return userToken{}, err
			} else {
				return userToken{token}, nil
			}
		}
	}
}

func (gdb *Gdb) userLogout(userName string) (Rows, error) {
	if _, err := updateItem(gdb.ItemDbPath, "update user_cfg set token='' where userName='"+userName+"'"); err != nil {
		return Rows{}, err
	} else {
		return Rows{1}, nil
	}
}

func (gdb *Gdb) getUserInfo(userName string) (UserInfo, error) {
	if r, err := query(gdb.ItemDbPath, "select role from user_cfg where userName='"+userName+"'"); err != nil || len(r) == 0 {
		return UserInfo{}, err
	} else {
		return UserInfo{
			UserName: UserName{userName},
			Role:     []string{r[0]["role"]},
		}, nil
	}
}

func (gdb *Gdb) addUsers(info addedUserInfo) (Rows, error) {
	sqlTemplate := template.Must(template.New("addUserTemplate").Parse(`insert into user_cfg (userName, passWord, role) 
								values ('{{.Name}}', '{{.PassWord}}', '{{.Role}}')`))
	var b bytes.Buffer
	if err := sqlTemplate.Execute(&b, info); err != nil {
		return Rows{}, err
	} else {
		sqlString := b.String()
		if _, err := updateItem(gdb.ItemDbPath, sqlString); err != nil {
			return Rows{}, err
		} else {
			return Rows{1}, nil
		}
	}
}

func (gdb *Gdb) deleteUsers(name UserName) (Rows, error) {
	if _, err := updateItem(gdb.ItemDbPath, "delete from user_cfg where userName='"+name.Name+"'"); err != nil {
		return Rows{}, err
	} else {
		return Rows{1}, nil
	}
}

func (gdb *Gdb) updateUsers(info updatedUserInfo) (Rows, error) {
	id, oldUserName, newUserName := info.Id, info.OldUserName, info.NewUserName
	if oldUserName == newUserName {
		if _, err := updateItem(gdb.ItemDbPath, "update user_cfg set role='"+info.Role+"' where id="+strconv.Itoa(id)); err != nil {
			return Rows{}, err
		} else {
			return Rows{1}, nil
		}
	} else {
		if _, err := updateItem(gdb.ItemDbPath, "update user_cfg set role='"+info.Role+"', userName='"+newUserName+"' where id="+strconv.Itoa(id)); err != nil {
			return Rows{}, err
		} else {
			return Rows{1}, nil
		}
	}
}

// AddItemsByExcel add items by excel
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
		items.ItemValues = values
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
	if b, err := ioutil.ReadFile("./uploadFiles/" + fileName); err != nil {
		return "", err
	} else {
		c := strings.Replace(fmt.Sprintf("%s", b), "\r\n", "<br />", -1)
		c1 := strings.Replace(c, "\n", "<br />", -1)
		return c1, nil
	}
}

func (gdb *Gdb) getLogs(info queryLogsInfo) (LogsInfo, error) {
	var queryStringTemplate, queryCountStringTemplate string
	if info.Level == "all" {
		queryStringTemplate = `select * from log_cfg where (insertTime > '{{.StartTime}}' and insertTime < '{{.EndTime}}') and (requestUser = '{{.Name}}' or requestUser = '') Limit {{.RowCount}} offset {{.StartRow}}`
		queryCountStringTemplate = `select count(*) as count from log_cfg where (insertTime > '{{.StartTime}}' and insertTime < '{{.EndTime}}') and (requestUser = '{{.Name}}' or requestUser = '')`
	} else {
		queryStringTemplate = `select * from log_cfg where (insertTime > '{{.StartTime}}' and insertTime < '{{.EndTime}}') and level='{{.Level}}' and (requestUser = '{{.Name}}' or requestUser = '') Limit {{.RowCount}} offset {{.StartRow}}`
		queryCountStringTemplate = `select count(*) as count from log_cfg where (insertTime > '{{.StartTime}}' and insertTime < '{{.EndTime}}') and level='{{.Level}}' and (requestUser = '{{.Name}}' or requestUser = '')`
	}
	var b, qb bytes.Buffer
	sqlTemplate := template.Must(template.New("sqlTemplate").Parse(queryStringTemplate))
	sqlQueryTemplate := template.Must(template.New("sqlQueryTemplate").Parse(queryCountStringTemplate))
	if err := sqlTemplate.Execute(&b, info); err != nil {
		return LogsInfo{}, err
	} else {
		if err := sqlQueryTemplate.Execute(&qb, info); err != nil {
			return LogsInfo{}, err
		} else {
			if result, err := query(gdb.ItemDbPath, b.String()); err != nil {
				return LogsInfo{}, err
			} else {
				if c, err := query(gdb.ItemDbPath, qb.String()); err != nil {
					return LogsInfo{}, err
				} else {
					if count, err := strconv.Atoi(c[0]["count"]); err != nil {
						return LogsInfo{}, err
					} else {
						return LogsInfo{result, count}, nil
					}
				}
			}
		}
	}
}

func (gdb *Gdb) deleteLogs(info deletedLogInfo) (Rows, error) {
	id, startTime, endTime, condition := info.Id, info.StartTime, info.EndTime, info.UserNameCondition
	var sqlString string
	if len(strings.Trim(id, " ")) != 0 {
		sqlString = "delete from log_cfg where id = '" + id + "'"
	} else {
		sqlString = "delete from log_cfg where (insertTime > '" + startTime + "' and insertTime <'" + endTime + "') and ( " + condition + ")"
	}
	if row, err := updateItem(gdb.ItemDbPath, sqlString); err != nil {
		return Rows{}, err
	} else {
		return Rows{EffectedRows: int(row)}, nil
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
