// +build gdbClient

/*
creatTime: 2020/12/10
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
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
	if r, err := query(gdb.ItemDbPath, "select passWord from user_cfg where userName='"+userName+"'"); err != nil || len(r) == 0 {
		return userToken{}, userNameError{"userNameError: " + userName}
	} else {
		if r[0]["passWord"] != info.PassWord {
			return userToken{}, userNameError{"userNameError: " + userName}
		} else {
			b := []byte(userName + "@seu" + time.Now().Format(timeFormatString) + "JustKeepSilence")
			token := fmt.Sprintf("%x", md5.Sum(b)) // result is 32-bit lowercase
			if _, err := updateItem(gdb.ItemDbPath, "update user_cfg set token='"+token+"'"+" where userName='"+userName+"'"); err != nil {
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

func (gdb *Gdb) getUserInfo(n string) (userInfo, error) {
	if r, err := query(gdb.ItemDbPath, "select role from user_cfg where userName='"+n+"'"); err != nil || len(r) == 0 {
		return userInfo{}, err
	} else {
		return userInfo{
			UserName: n,
			Role:     []string{r[0]["role"]},
		}, nil
	}
}

func (gdb *Gdb) addUsers(info addedUserInfo) (Rows, error) {
	if !strings.Contains(roles, info.Role) {
		return Rows{}, fmt.Errorf("role must be visitor or common_user or super_user")
	}
	sqlTemplate := template.Must(template.New("addUserTemplate").Parse(`insert into user_cfg (userName, passWord, role) 
								values ('{{.Name}}', '{{.PassWord}}', '{{.Role}}')`))
	var b bytes.Buffer
	if err := sqlTemplate.Execute(&b, info); err != nil {
		return Rows{}, err
	} else {
		sqlString := b.String() // addUsers
		var routeSqlString string
		var routeRoles []string
		switch info.Role {
		case "super_user":
			routeRole := strings.Replace(superUserRoutes, "userName", info.Name, -1)
			routeSqlString = "insert into route_cfg (userName, routeRoles) values ('" + info.Name + "', '" + routeRole + "')"
			routeRoles = append(routeRoles, routeRole)
			break
		case "common_user":
			for _, route := range commonUserRoutes {
				routeRole := "p," + info.Name + "," + toTitle(route) + "," + "POST"
				routeRoles = append(routeRoles, routeRole)
			}
			break
		default:
			// visitor
			for _, route := range visitorUserRoutes {
				routeRole := "p," + info.Name + "," + toTitle(route) + "," + "POST"
				routeRoles = append(routeRoles, routeRole)
			}
			break
		}
		r, _ := json.Marshal(routeRoles)
		routeSqlString = "insert into route_cfg (userName, routeRoles) values ('" + info.Name + "', '" + string(r) + "')"
		if err := updateItems(gdb.ItemDbPath, sqlString, routeSqlString); err != nil {
			return Rows{}, err
		} else {
			// add policy to model
			m := gdb.gdbAdapter.e.GetModel()
			for _, ast := range m["p"] {
				for _, role := range routeRoles {
					ast.Policy = append(ast.Policy, strings.Split(role, ",")[1:])
				}
			}
			return Rows{1}, nil
		}
	}
}

func (gdb *Gdb) deleteUsers(name userName) (Rows, error) {
	if err := updateItems(gdb.ItemDbPath, "delete from user_cfg where userName='"+name.Name+"'", "delete from route_cfg where userName='"+name.Name+"'"); err != nil {
		return Rows{}, err
	} else {
		// remove policy from model
		if err := gdb.gdbAdapter.e.LoadPolicy(); err != nil {
			return Rows{}, err
		} else {
			return Rows{1}, nil
		}
	}
}

func (gdb *Gdb) updateUsers(info updatedUserInfo) (Rows, error) {
	if info.NewPassWord == "" {
		r, _ := query(gdb.ItemDbPath, "select passWord from user_cfg where userName='"+info.UserName+"'")
		info.NewPassWord = r[0]["passWord"]
	}
	if info.NewUserName == "" {
		info.NewUserName = info.UserName
	}
	if info.NewRole == "" {
		r, _ := query(gdb.ItemDbPath, "select role from user_cfg where userName='"+info.UserName+"'")
		info.NewRole = r[0]["role"]
	}
	sqlTemplate := template.Must(template.New("updateUserTemplate").Parse(`update user_cfg set role='{{.NewRole}}', userName='{{.NewUserName}}',
								 passWord='{{.NewPassWord}}' where userName='{{.UserName}}'`))
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

// addItemsByExcel add items by excel
func (gdb *Gdb) addItemsByExcel(groupName, filePath string) (Rows, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return Rows{-1}, excelError{"excelError: " + err.Error()}
	} else {
		// open excel successfully
		sheetName := f.GetSheetList()[0] // use first worksheet
		rows, err := f.Rows(sheetName)   // get all rows
		var headers []string             // headers
		var items AddedItemsInfo
		var values []map[string]string
		if err != nil {
			return Rows{-1}, excelError{"excelError: " + err.Error()}
		} else {
			// get rows successfully
			count := 0
			for rows.Next() {
				if count == 0 {
					// check headers
					h, err := rows.Columns() // columns of excel
					if err != nil {
						return Rows{-1}, excelError{"excelError: " + err.Error()}
					} else {
						// get headers successfully
						cols, err := gdb.GetGroupProperty(groupName, "1=1")
						if err != nil {
							return Rows{-1}, err
						}
						headers = cols.ItemColumnNames // columns of database
						if !equal(h, headers) {
							return Rows{-1}, excelError{"excelError: Inconsistent header"}
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

func (gdb *Gdb) importHistoryByExcel(fileName, groupName string, itemNames []string, sheetNames ...string) error {
	if f, err := excelize.OpenFile(fileName); err != nil {
		return excelError{"excelError: " + err.Error()}
	} else {
		dataTypes := []string{}
		for _, itemName := range itemNames {
			if t, ok := gdb.rtDbFilter.Get(itemName + joiner + groupName); !ok {
				return fmt.Errorf("item " + itemName + " not existed")
			} else {
				dataTypes = append(dataTypes, t.(string))
			}
		}
		infos := []HistoricalItemValue{}
		for index := 0; index < len(itemNames); index++ {
			sheetName, itemName := sheetNames[index], itemNames[index]
			if rows, err := f.Rows(sheetName); err != nil {
				return err
			} else {
				info := HistoricalItemValue{ItemName: itemName}
				var values []string
				var timeStamps []int
				for rows.Next() {
					// first row is timeStamp, second is value
					if c, err := rows.Columns(); err != nil {
						return err
					} else {
						values = append(values, c[1])
						if t, err := time.Parse(timeFormatString, c[0]); err != nil {
							return err
						} else {
							timeStamps = append(timeStamps, int(t.Unix()))
						}
					}
				}
				if v, err := convertValues(dataTypes[index], values...); err != nil {
					return err
				} else {
					info.Values = v
					info.TimeStamps = timeStamps
					info.GroupName = groupName
					infos = append(infos, info)
				}
			}
		}
		if err := gdb.BatchWriteHistoricalData(infos...); err != nil {
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
		return string(b), nil
	}
}

func (gdb *Gdb) getLogs(info queryLogsInfo) (logsInfo, error) {
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
		return logsInfo{}, err
	} else {
		if err := sqlQueryTemplate.Execute(&qb, info); err != nil {
			return logsInfo{}, err
		} else {
			if result, err := query(gdb.ItemDbPath, b.String()); err != nil {
				return logsInfo{}, err
			} else {
				if c, err := query(gdb.ItemDbPath, qb.String()); err != nil {
					return logsInfo{}, err
				} else {
					if count, err := strconv.Atoi(c[0]["count"]); err != nil {
						return logsInfo{}, err
					} else {
						return logsInfo{result, count}, nil
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

func (gdb *Gdb) getRoutes() ([]map[string]string, error) {
	if rows, err := query(gdb.ItemDbPath, "select userName, routeRoles from route_cfg where 1=1"); err != nil {
		return nil, err
	} else {
		for _, row := range rows {
			name := row["userName"]
			if r, err := query(gdb.ItemDbPath, "select role from user_cfg where userName='"+name+"'"); err != nil {
				return nil, err
			} else {
				if len(r) == 0 {
					row["role"] = ""
				} else {
					row["role"] = r[0]["role"]
				}
			}

		}
		return rows, nil
	}
}

func (gdb *Gdb) deleteRoutes(name string, routes ...string) error {
	for _, route := range routes {
		if _, err := gdb.e.RemovePolicy(name, route, "POST"); err != nil {
			return err
		}
	}
	return nil
}

// add routes to existed user
func (gdb *Gdb) addRoutes(name string, routes ...string) error {
	for _, route := range routes {
		if _, err := gdb.e.AddPolicy(name, route, "POST"); err != nil {
			return err
		}
	}
	return nil
}

// add routes to new user
func (gdb *Gdb) addUserRoutes(name string, routes ...string) error {
	routeRoles := []string{}
	for _, route := range routes {
		routeRole := "p," + name + "," + toTitle(route) + "," + "POST"
		routeRoles = append(routeRoles, routeRole)
	}
	r, _ := json.Marshal(routeRoles)
	routeSqlString := "insert into route_cfg (userName, routeRoles) values ('" + name + "', '" + string(r) + "')"
	if _, err := updateItem(gdb.ItemDbPath, routeSqlString); err != nil {
		return err
	} else {
		// add policy to model
		m := gdb.gdbAdapter.e.GetModel()
		for _, ast := range m["p"] {
			for _, role := range routeRoles {
				ast.Policy = append(ast.Policy, strings.Split(role, ",")[1:])
			}
		}
		return nil
	}
}

// check whether routes exist, if all routes exist, return [], true
func (gdb *Gdb) checkRoutes(name string, routes ...string) ([]int32, bool) {
	index := []int32{}
	for i, route := range routes {
		if !gdb.e.HasPolicy(name, toTitle(route), "POST") {
			index = append(index, int32(i))
		}
	}
	if len(index) == 0 {
		return index, false
	} else {
		return index, true
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
