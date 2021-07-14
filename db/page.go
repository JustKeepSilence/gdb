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
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// user login
func (gdb *Gdb) userLogin(info authInfo) (userToken, error) {
	userName := info.UserName
	if r, err := gdb.query("select passWord from user_cfg where userName='" + userName + "'"); err != nil || len(r) == 0 {
		return userToken{}, fmt.Errorf("userName error:" + userName)
	} else {
		if r[0]["passWord"] != info.PassWord {
			return userToken{}, fmt.Errorf("passWord error")
		} else {
			b := convertStringToByte(userName + "@seu" + time.Now().Format(timeFormatString) + "JustKeepSilence")
			token := fmt.Sprintf("%x", md5.Sum(b)) // result is 32-bit lowercase
			if _, err := gdb.updateItem("update user_cfg set token='" + token + "'" + " where userName='" + userName + "'"); err != nil {
				return userToken{}, err
			} else {
				return userToken{token}, nil
			}
		}
	}
}

func (gdb *Gdb) userLogout(userName string) (TimeRows, error) {
	st := time.Now()
	if _, err := gdb.updateItem("update user_cfg set token='' where userName='" + userName + "'"); err != nil {
		return TimeRows{}, err
	} else {
		return TimeRows{EffectedRows: 1, Times: time.Since(st).Milliseconds()}, nil
	}
}

func (gdb *Gdb) getUserInfo(n string) (userInfo, error) {
	if r, err := gdb.query("select role from user_cfg where userName='" + n + "'"); err != nil || len(r) == 0 {
		return userInfo{}, err
	} else {
		return userInfo{
			UserName: n,
			Role:     []string{r[0]["role"]},
		}, nil
	}
}

func (gdb *Gdb) addUsers(info addedUserInfo) (TimeRows, error) {
	st := time.Now()
	if !strings.Contains(roles, info.Role) {
		return TimeRows{}, fmt.Errorf("role must be visitor or common_user or super_user")
	}
	sqlTemplate := template.Must(template.New("addUserTemplate").Parse(`insert into user_cfg (userName, passWord, role) 
								values ('{{.Name}}', '{{.PassWord}}', '{{.Role}}')`))
	var b bytes.Buffer
	if err := sqlTemplate.Execute(&b, info); err != nil {
		return TimeRows{}, err
	} else {
		sqlString := b.String() // addUsers
		var routeSqlString string
		var routeRoles []string
		switch info.Role {
		case "super_user":
			routeRoles = append(routeRoles, "p,"+info.Name+",all, POST")
			break
		case "common_user":
			for _, route := range commonUserRoutes {
				routeRole := "p," + info.Name + "," + strings.Title(route) + "," + "POST"
				routeRoles = append(routeRoles, routeRole)
			}
			break
		default:
			// visitor
			for _, route := range visitorUserRoutes {
				routeRole := "p," + info.Name + "," + strings.Title(route) + "," + "POST"
				routeRoles = append(routeRoles, routeRole)
			}
			break
		}
		r, _ := json.Marshal(routeRoles)
		routeSqlString = "insert into route_cfg (userName, routeRoles) values ('" + info.Name + "', '" + string(r) + "')"
		if err := gdb.updateItems(sqlString, routeSqlString); err != nil {
			return TimeRows{}, err
		} else {
			// add policy to model
			m := gdb.e.GetModel()
			for _, ast := range m["p"] {
				for _, role := range routeRoles {
					ast.Policy = append(ast.Policy, strings.Split(role, ",")[1:])
				}
			}
			return TimeRows{EffectedRows: 1, Times: time.Since(st).Milliseconds()}, nil
		}
	}
}

func (gdb *Gdb) deleteUsers(name userName) (TimeRows, error) {
	st := time.Now()
	if name.Name == "admin" {
		return TimeRows{}, fmt.Errorf("you can't delete admin user")
	}
	if err := gdb.updateItems("delete from user_cfg where userName='"+name.Name+"'", "delete from route_cfg where userName='"+name.Name+"'"); err != nil {
		return TimeRows{}, err
	} else {
		// remove policy from model
		if err := gdb.e.LoadPolicy(); err != nil {
			return TimeRows{}, err
		} else {
			return TimeRows{EffectedRows: 1, Times: time.Since(st).Milliseconds()}, nil
		}
	}
}

func (gdb *Gdb) updateUsers(info updatedUserInfo) (TimeRows, error) {
	st := time.Now()
	if info.NewPassWord == "" {
		r, _ := gdb.query("select passWord from user_cfg where userName='" + info.UserName + "'")
		info.NewPassWord = r[0]["passWord"]
	}
	if info.NewUserName == "" {
		info.NewUserName = info.UserName
	}
	if info.NewRole == "" {
		r, _ := gdb.query("select role from user_cfg where userName='" + info.UserName + "'")
		info.NewRole = r[0]["role"]
	}
	sqlTemplate := template.Must(template.New("updateUserTemplate").Parse(`update user_cfg set role='{{.NewRole}}', userName='{{.NewUserName}}',
								 passWord='{{.NewPassWord}}' where userName='{{.UserName}}'`))
	var routeSqlString string
	var routeRoles []string
	switch info.NewRole {
	case "super_user":
		routeRoles = append(routeRoles, "p,"+info.NewUserName+",all,POST")
		break
	case "common_user":
		for _, route := range commonUserRoutes {
			routeRole := "p," + info.NewUserName + "," + strings.Title(route) + "," + "POST"
			routeRoles = append(routeRoles, routeRole)
		}
		break
	case "visitor":
		// visitor
		for _, route := range visitorUserRoutes {
			routeRole := "p," + info.NewUserName + "," + strings.Title(route) + "," + "POST"
			routeRoles = append(routeRoles, routeRole)
		}
		break
	default:
		return TimeRows{}, fmt.Errorf("userRole can only be " + roles)
	}
	r, _ := json.Marshal(routeRoles)
	routeSqlString = "update route_cfg set userName='" + info.NewUserName + "', routeRoles='" + string(r) + "' where userName='" + info.UserName + "'"
	var b bytes.Buffer
	if err := sqlTemplate.Execute(&b, info); err != nil {
		return TimeRows{}, err
	} else {
		sqlString := b.String()
		if err := gdb.updateItems(sqlString, routeSqlString); err != nil {
			return TimeRows{}, err
		} else {
			// update policy
			if err := gdb.e.LoadPolicy(); err != nil {
				return TimeRows{}, err
			}
			return TimeRows{EffectedRows: 1, Times: time.Since(st).Milliseconds()}, nil
		}
	}
}

// addItemsByExcel add items by excel
func (gdb *Gdb) addItemsByExcel(groupName, filePath string) (TimeRows, error) {
	st := time.Now()
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return TimeRows{}, err
	} else {
		// open excel successfully
		sheetName := f.GetSheetList()[0] // use first worksheet
		rows, err := f.Rows(sheetName)   // get all rows
		var headers []string             // headers
		var items AddedItemsInfo
		var values []map[string]string
		if err != nil {
			return TimeRows{}, err
		} else {
			// get rows successfully
			count := 0
			for rows.Next() {
				if count == 0 {
					// check headers
					h, err := rows.Columns() // columns of excel
					if err != nil {
						return TimeRows{}, err
					} else {
						// get headers successfully
						cols, err := gdb.GetGroupProperty(groupName, "1=1")
						if err != nil {
							return TimeRows{}, err
						}
						headers = cols.ItemColumnNames // columns of database
						if !equal(h, headers) {
							return TimeRows{}, fmt.Errorf("inconsistent header")
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
			return TimeRows{}, err
		} else {
			return TimeRows{EffectedRows: r.EffectedRows, Times: time.Since(st).Milliseconds()}, nil
		}
	}
}

func (gdb *Gdb) importHistoryByExcel(fileName, groupName string, itemNames []string, sheetNames ...string) (TimeRows, error) {
	st := time.Now()
	if len(itemNames) != len(sheetNames) {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
	if f, err := excelize.OpenFile(fileName); err != nil {
		return TimeRows{}, err
	} else {
		dataTypes := map[string]string{}
		for _, itemName := range itemNames {
			if t, ok := gdb.rtDbFilter.Get(itemName + joiner + groupName); !ok {
				return TimeRows{}, fmt.Errorf("item " + itemName + " not existed")
			} else {
				dataTypes[itemName] = t.(string)
			}
		}
		floatItems := floatHItemValues{}
		intItems := intHItemValues{}
		stringItems := stringHItemValues{}
		boolItems := boolHItemValues{}
		for index := 0; index < len(itemNames); index++ {
			sheetName, itemName := sheetNames[index], itemNames[index]
			if rows, err := f.GetRows(sheetName); err != nil {
				return TimeRows{}, err
			} else {
				timeStamps := make([]int32, len(rows))
				switch dataType := dataTypes[itemName]; dataType {
				case "float64":
					itemValues := make([]float32, len(rows))
					for i := 0; i < len(rows); i++ {
						row := rows[i]
						r, err := strconv.ParseFloat(row[1], 32)
						{
							if err != nil {
								return TimeRows{}, err
							}
							itemValues[i] = float32(r)
						}
						if t, err := time.Parse(timeFormatString, row[0]); err != nil {
							return TimeRows{}, err
						} else {
							timeStamps[i] = int32(t.Unix())
						}
					}
					floatItems.ItemNames = append(floatItems.ItemNames, itemName)
					floatItems.GroupNames = append(floatItems.GroupNames, groupName)
					floatItems.ItemValues = append(floatItems.ItemValues, itemValues)
					floatItems.TimeStamps = append(floatItems.TimeStamps, timeStamps)
				case "int64":
					itemValues := make([]int32, len(rows))
					for i := 0; i < len(rows); i++ {
						row := rows[i]
						r, err := strconv.ParseInt(row[1], 10, 32)
						{
							if err != nil {
								return TimeRows{}, err
							}
							itemValues[i] = int32(r)
						}
						if t, err := time.Parse(timeFormatString, row[0]); err != nil {
							return TimeRows{}, err
						} else {
							timeStamps[i] = int32(t.Unix())
						}
					}
					intItems.ItemNames = append(intItems.ItemNames, itemName)
					intItems.GroupNames = append(intItems.GroupNames, groupName)
					intItems.ItemValues = append(intItems.ItemValues, itemValues)
					intItems.TimeStamps = append(intItems.TimeStamps, timeStamps)
				case "string":
					itemValues := make([]string, len(rows))
					for i := 0; i < len(rows); i++ {
						row := rows[i]
						itemValues[i] = row[1]
						if t, err := time.Parse(timeFormatString, row[0]); err != nil {
							return TimeRows{}, err
						} else {
							timeStamps[i] = int32(t.Unix())
						}
					}
					stringItems.ItemNames = append(stringItems.ItemNames, itemName)
					stringItems.GroupNames = append(stringItems.GroupNames, groupName)
					stringItems.ItemValues = append(stringItems.ItemValues, itemValues)
					stringItems.TimeStamps = append(stringItems.TimeStamps, timeStamps)
				case "bool":
					itemValues := make([]bool, len(rows))
					for i := 0; i < len(rows); i++ {
						row := rows[i]
						r, err := strconv.ParseBool(row[1])
						{
							if err != nil {
								return TimeRows{}, err
							}
							itemValues[i] = r
						}
						if t, err := time.Parse(timeFormatString, row[0]); err != nil {
							return TimeRows{}, err
						} else {
							timeStamps[i] = int32(t.Unix())
						}
					}
					boolItems.ItemNames = append(boolItems.ItemNames, itemName)
					boolItems.GroupNames = append(boolItems.GroupNames, groupName)
					boolItems.ItemValues = append(boolItems.ItemValues, itemValues)
					boolItems.TimeStamps = append(boolItems.TimeStamps, timeStamps)
				default:
					return TimeRows{}, fmt.Errorf("unknown dataType " + dataType)
				}
			}
		}
		g := errgroup.Group{}
		counts := make([]int, 4)
		g.Go(func() error {
			if len(floatItems.ItemNames) != 0 {
				r, err := gdb.BatchWriteFloatHistoricalData(floatItems.GroupNames, floatItems.ItemNames, floatItems.TimeStamps, floatItems.ItemValues)
				{
					if err != nil {
						return err
					}
					counts[0] = r.EffectedRows
				}
			}
			return nil
		})
		g.Go(func() error {
			if len(intItems.ItemNames) != 0 {
				r, err := gdb.BatchWriteIntHistoricalData(intItems.GroupNames, intItems.ItemNames, intItems.TimeStamps, intItems.ItemValues)
				{
					if err != nil {
						return err
					}
					counts[1] = r.EffectedRows
				}
			}
			return nil
		})
		g.Go(func() error {
			if len(stringItems.ItemNames) != 0 {
				r, err := gdb.BatchWriteStringHistoricalData(stringItems.GroupNames, stringItems.ItemNames, stringItems.TimeStamps, stringItems.ItemValues)
				{
					if err != nil {
						return err
					}
					counts[2] = r.EffectedRows
				}
			}
			return nil
		})
		g.Go(func() error {
			if len(boolItems.ItemNames) != 0 {
				r, err := gdb.BatchWriteBoolHistoricalData(boolItems.GroupNames, boolItems.ItemNames, boolItems.TimeStamps, boolItems.ItemValues)
				{
					if err != nil {
						return err
					}
					counts[3] = r.EffectedRows
				}
			}
			return nil
		})
		if err := g.Wait(); err != nil {
			return TimeRows{}, err
		}
		return TimeRows{EffectedRows: counts[0] + counts[1] + counts[2] + counts[3], Times: time.Since(st).Milliseconds()}, nil
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
	if !strings.Contains("all, Info, Error", info.Level) {
		return logsInfo{}, fmt.Errorf("query log level can only be all or Info or Error")
	}
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
			if result, err := gdb.query(b.String()); err != nil {
				return logsInfo{}, err
			} else {
				if c, err := gdb.query(qb.String()); err != nil {
					return logsInfo{}, err
				} else {
					if count, err := strconv.Atoi(c[0]["count"]); err != nil {
						return logsInfo{}, err
					} else {
						return logsInfo{result, int64(count)}, nil
					}
				}
			}
		}
	}
}

func (gdb *Gdb) deleteLogs(info deletedLogInfo) (TimeRows, error) {
	id, startTime, endTime, condition := info.Id, info.StartTime, info.EndTime, info.UserNameCondition
	var sqlString string
	if len(strings.Trim(id, " ")) != 0 {
		sqlString = "delete from log_cfg where id = '" + id + "'"
	} else {
		sqlString = "delete from log_cfg where (insertTime > '" + startTime + "' and insertTime <'" + endTime + "') and ( " + condition + ")"
	}
	if row, err := gdb.updateItem(sqlString); err != nil {
		return TimeRows{}, err
	} else {
		return TimeRows{EffectedRows: int(row)}, nil
	}
}

func (gdb *Gdb) getRoutes() ([]map[string]string, error) {
	if rows, err := gdb.query("select userName, routeRoles from route_cfg where 1=1"); err != nil {
		return nil, err
	} else {
		for _, row := range rows {
			name := row["userName"]
			if r, err := gdb.query("select role from user_cfg where userName='" + name + "'"); err != nil {
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
		routeRole := "p," + name + "," + strings.Title(route) + "," + "POST"
		routeRoles = append(routeRoles, routeRole)
	}
	r, _ := json.Marshal(routeRoles)
	routeSqlString := "insert into route_cfg (userName, routeRoles) values ('" + name + "', '" + string(r) + "')"
	if _, err := gdb.updateItem(routeSqlString); err != nil {
		return err
	} else {
		// add policy to model
		m := gdb.e.GetModel()
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
		if !gdb.e.HasPolicy(name, strings.Title(route), "POST") {
			index = append(index, int32(i))
		}
	}
	if len(index) == 0 {
		return index, false
	} else {
		return index, true
	}
}

func (gdb *Gdb) getCmdInfo(p string) (string, error) {
	//return gdb.hisDb["float"].GetProperty("leveldb." + p)
	return "", nil
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
