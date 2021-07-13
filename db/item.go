/*
creatTime: 2020/12/8
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	types = "int32, float32, bool, string"
)

// AddItems to gdb,you can not add duplicate items.you need not add value of id column, and for
// value of dataType, it can only be int32 or float32 or bool or string.If column has not
// default value, you MUST specify value
//
// The operation is atomic
func (gdb *Gdb) AddItems(itemInfo AddedItemsInfo) (TimeRows, error) {
	st := time.Now()
	groupName := itemInfo.GroupName
	itemValues := itemInfo.ItemValues
	var c []map[string]string
	if gdb.driverName == "sqlite3" {
		var err error
		c, err = gdb.query("PRAGMA table_info([" + groupName + "])") // get column names
		if err != nil {
			return TimeRows{EffectedRows: -1, Times: time.Since(st).Milliseconds()}, err
		}
	} else {
		var err error
		c, err = gdb.query("select column_name as name from information_schema.columns where table_schema='" + gdb.itemDbName + "' and table_name='" + groupName + "'")
		if err != nil {
			return TimeRows{EffectedRows: -1, Times: time.Since(st).Milliseconds()}, err
		}
	}
	c = c[1:]                     // abort first item id
	var columnNames []string      // without '
	var addedColumnNames []string // with '
	for i := 0; i < len(c); i++ {
		columnNames = append(columnNames, c[i]["name"])
		addedColumnNames = append(addedColumnNames, "`"+c[i]["name"]+"`")
	}
	sb := strings.Builder{}
	sb.Write(convertStringToByte("insert into `"))
	sb.Write(convertStringToByte(groupName))
	sb.Write(convertStringToByte("` ("))
	sb.Write(convertStringToByte(strings.Join(addedColumnNames, ",")))
	sb.Write(convertStringToByte(") "))
	sb.Write(convertStringToByte("values("))
	sb.Write(convertStringToByte(strings.TrimRight(strings.Repeat("?,", len(columnNames)), ","))) // groupName column
	sb.Write(convertStringToByte(")"))
	insertSqlString := sb.String()
	var addedItemValues [][]string
	var itemNames []string
	dataTypes := []string{}
	for _, itemValue := range itemValues {
		var t []string
		for i := 0; i < len(columnNames); i++ {
			cv, ok := itemValue[columnNames[i]]
			if ok {
				// key exist
				t = append(t, cv)
				if columnNames[i] == "itemName" {
					itemNames = append(itemNames, cv)
				}
				if columnNames[i] == "dataType" {
					// check dataTypes of item
					if strings.Trim(cv, " ") == "" || !strings.Contains(types, cv) {
						return TimeRows{EffectedRows: -1, Times: time.Since(st).Milliseconds()}, fmt.Errorf("dataType can't be empty, it must be " + types)
					} else {
						dataTypes = append(dataTypes, cv)
					}
				}
			} else {
				// key not exist
				return TimeRows{EffectedRows: -1, Times: time.Since(st).Milliseconds()}, fmt.Errorf("columnNameError: " + columnNames[i])
			}

		}
		addedItemValues = append(addedItemValues, t)
	}
	if err := gdb.insertItems(insertSqlString, addedItemValues...); err != nil {
		return TimeRows{EffectedRows: -1, Times: time.Since(st).Milliseconds()}, err
	}
	for index, itemName := range itemNames {
		gdb.rtDbFilter.Set(itemName+joiner+groupName, dataTypes[index])
	}
	return TimeRows{EffectedRows: len(itemValues), Times: time.Since(st).Milliseconds()}, nil
}

// DeleteItems delete items from given group according to condition, condition is where clause
// in SQL.
//
//NOTE:this operation will not delete history data of item, if you want to delete history data of
// item, you should use CleanItemData method
func (gdb *Gdb) DeleteItems(itemInfo DeletedItemsInfo) (TimeRows, error) {
	st := time.Now()
	groupName := itemInfo.GroupName
	condition := itemInfo.Condition
	item, err := gdb.query("select itemName from `" + groupName + "` where " + condition)
	if len(item) == 0 {
		return TimeRows{}, fmt.Errorf("conditionError: " + condition)
	}
	if err != nil {
		return TimeRows{}, err
	}
	rows, err := gdb.updateItem("delete from `" + groupName + "` where " + condition)
	if err != nil {
		return TimeRows{}, err
	}
	gdb.rtDbFilter.Remove(item[0]["itemName"] + joiner + groupName) // remove key from bloom filter
	return TimeRows{EffectedRows: int(rows), Times: time.Since(st).Milliseconds()}, nil
}

// GetItems get items from gdb according to the given columnName, condition,startRow and rowCount.
// columnName is the columnName you want to get, if you want to get all columns info, you should use *
// and condition is where clause, startRow and rowCount is used to page query, if startRow is -1,
// that is get all items without pagination
func (gdb *Gdb) GetItems(itemInfo ItemsInfo) (GdbItems, error) {
	groupName := itemInfo.GroupName // groupName
	columns := itemInfo.ColumnNames // column
	startRow := itemInfo.StartRow   // startRow
	condition := itemInfo.Condition // condition
	var rows []map[string]string
	var err error
	if startRow == -1 {
		// all rows
		rows, err = gdb.query("select " + columns + " from `" + groupName + "` where " + condition)
		if err != nil {
			return GdbItems{}, err
		}
	} else {
		// Limit gdb.query
		rowCount := itemInfo.RowCount
		rows, err = gdb.query("select " + columns + " from `" + groupName + "` where " + condition + " Limit " + strconv.Itoa(int(rowCount)) + " offset " + strconv.Itoa(int(startRow)))
		if err != nil {
			return GdbItems{}, err
		}
	}
	return GdbItems{ItemValues: rows}, nil
}

func (gdb *Gdb) getItemsWithCount(itemInfo ItemsInfo) (gdbItemsWithCount, error) {
	condition := itemInfo.Condition
	groupName := itemInfo.GroupName
	c, err := gdb.query("select count(*) as count from `" + groupName + "` where " + condition)
	if err != nil {
		return gdbItemsWithCount{}, err
	}
	itemValues, err := gdb.GetItems(itemInfo)
	if err != nil {
		return gdbItemsWithCount{}, nil
	}
	count, err := strconv.ParseInt(c[0]["count"], 10, 64)
	if err != nil {
		return gdbItemsWithCount{}, nil
	}
	return gdbItemsWithCount{int32(count), itemValues}, nil
}

// UpdateItems update items in gdb according to given condition and clause.condition is where
// clause in SQL and clause is set clause in SQL.You can't update id or itemName or dataType column
func (gdb *Gdb) UpdateItems(itemInfo UpdatedItemsInfo) (TimeRows, error) {
	st := time.Now()
	groupName := itemInfo.GroupName
	condition := itemInfo.Condition
	clause := itemInfo.Clause
	// Firstly determine whether to update itemName or dataType or id
	regPoint := regexp.MustCompile(`(?i)itemName='(.*?)'`) // Match the content after itemName
	regItemName := regexp.MustCompile(`(?i)dataType='(.*?)'`)
	regId := regexp.MustCompile(`(?i)id='(.*?)'`)
	if !regPoint.Match(convertStringToByte(clause)) && !regItemName.Match(convertStringToByte(clause)) && !regId.Match(convertStringToByte(clause)) {
		// no itemName
		// update SQLite
		rows, err := gdb.updateItem("update `" + groupName + "` set " + clause + " where " + condition)
		if err != nil {
			return TimeRows{}, err
		}
		return TimeRows{EffectedRows: int(rows), Times: time.Since(st).Milliseconds()}, nil
	} else {
		return TimeRows{}, fmt.Errorf("can't update itemName, dataType or id")
	}
}

// CheckItems check whether the given items existed in the given group
func (gdb *Gdb) CheckItems(groupName string, itemNames ...string) error {
	for _, itemName := range itemNames {
		if _, ok := gdb.rtDbFilter.Get(itemName + joiner + groupName); !ok {
			return fmt.Errorf("itemName: " + itemName + " not existed")
		}
	}
	return nil
}

//shrinkItemDb will shrink itemDb when using sqlite to store items if fileSize >= 0.5G
func (gdb *Gdb) shrinkItemDb() error {
	if gdb.driverName == "sqlite3" {
		t := time.NewTicker(time.Minute)
		for {
			select {
			case <-t.C:
				if size, err := dirSize(gdb.dsn); err != nil {
					fmt.Println("failing to get size of database:" + err.Error())
				} else {
					if size > 0.5 {
						// more than 0.5G
						if _, err := gdb.updateItem("vacuum ;"); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
