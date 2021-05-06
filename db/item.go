/*
creatTime: 2020/12/8
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"regexp"
	"strconv"
	"strings"
)

// AddItems to gdb,you can not add existed items
func (gdb *Gdb) AddItems(itemInfo AddedItemsInfo) (Rows, error) {
	groupName := itemInfo.GroupName
	itemValues := itemInfo.ItemValues
	c, err := query(gdb.ItemDbPath, "PRAGMA table_info(["+groupName+"])") // get column names
	if err != nil {
		return Rows{-1}, err
	}
	c = c[1:]                     // abort first item id
	var columnNames []string      // without '
	var addedColumnNames []string // with '
	for i := 0; i < len(c); i++ {
		columnNames = append(columnNames, c[i]["name"])
		addedColumnNames = append(addedColumnNames, "'"+c[i]["name"]+"'")
	}
	sb := strings.Builder{}
	sb.Write([]byte("insert into '"))
	sb.Write([]byte(groupName))
	sb.Write([]byte("' ("))
	sb.Write([]byte(strings.Join(addedColumnNames, ",")))
	sb.Write([]byte(") "))
	sb.Write([]byte("values("))
	sb.Write([]byte(strings.TrimRight(strings.Repeat("?,", len(columnNames)), ","))) // groupName column
	sb.Write([]byte(")"))
	insertSqlString := sb.String()
	var addedItemValues [][]string
	var itemNames []string
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
			} else {
				// key not exist
				return Rows{-1}, columnNameError{"columnNameError: " + columnNames[i]}
			}

		}
		addedItemValues = append(addedItemValues, t)
	}
	if err := insertItems(gdb.ItemDbPath, insertSqlString, addedItemValues...); err != nil {
		return Rows{-1}, err
	}
	for _, itemName := range itemNames {
		gdb.rtDbFilter.Set(itemName, struct{}{})
	}
	// initial write realTime data, all key write ""
	batch := leveldb.Batch{}
	for i := 0; i < len(itemNames); i++ {
		batch.Put([]byte(itemNames[i]), []byte(""))
	}
	_ = gdb.rtDb.Write(&batch, nil)
	return Rows{len(itemValues)}, nil
}

// DeleteItems delete items from given group according to condition, condition is where clause
// in sqlite
func (gdb *Gdb) DeleteItems(itemInfo DeletedItemsInfo) (Rows, error) {
	groupName := itemInfo.GroupName
	condition := itemInfo.Condition
	item, err := query(gdb.ItemDbPath, "select itemName from '"+groupName+"' where "+condition)
	if len(item) == 0 {
		return Rows{}, conditionError{"conditionError: " + condition}
	}
	if err != nil {
		return Rows{}, err
	}
	rows, err := updateItem(gdb.ItemDbPath, "delete from '"+groupName+"' where "+condition)
	if err != nil {
		return Rows{}, err
	}
	gdb.rtDbFilter.Remove(item[0]["itemName"]) // remove key from bloom filter
	return Rows{int(rows)}, nil
}

// GetItems get items from gdb according to the given columnName, condition,startRow and rowCount.
// columnName is the columnName you want to get, and condition is where clause, startRow and rowCount
// is used to page query, if startRow is -1, that is get all items without pagination
func (gdb *Gdb) GetItems(itemInfo ItemsInfo) (GdbItems, error) {
	groupName := itemInfo.GroupName // groupName
	columns := itemInfo.ColumnNames // column
	startRow := itemInfo.StartRow   // startRow
	condition := itemInfo.Condition // condition
	var rows []map[string]string
	var err error
	if startRow == -1 {
		// all rows
		rows, err = query(gdb.ItemDbPath, "select "+columns+" from '"+groupName+"' where "+condition)
		if err != nil {
			return GdbItems{}, err
		}
	} else {
		// Limit query
		rowCount := itemInfo.RowCount
		rows, err = query(gdb.ItemDbPath, "select "+columns+" from '"+groupName+"' where "+condition+" Limit "+strconv.Itoa(rowCount)+" offset "+strconv.Itoa(startRow))
		if err != nil {
			return GdbItems{}, err
		}
	}
	return GdbItems{ItemValues: rows}, nil
}

func (gdb *Gdb) getItemsWithCount(itemInfo ItemsInfo) (gdbItemsWithCount, error) {
	condition := itemInfo.Condition
	groupName := itemInfo.GroupName
	c, err := query(gdb.ItemDbPath, "select count(*) as count from '"+groupName+"' where "+condition)
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
	return gdbItemsWithCount{count, itemValues}, nil
}

// UpdateItems update items in gdb according to given condition and clause.condition is where
// clause in sqlite and clause is set clause in sqlite
func (gdb *Gdb) UpdateItems(itemInfo UpdatedItemsInfo) (Rows, error) {
	groupName := itemInfo.GroupName
	condition := itemInfo.Condition
	clause := itemInfo.Clause
	// Firstly determine whether to update itemName
	regPoint := regexp.MustCompile(`(?i)itemName='(.*?)'`) // Match the content after itemName
	if !regPoint.Match([]byte(clause)) {
		// no itemName
		// update SQLite
		rows, err := updateItem(gdb.ItemDbPath, "update '"+groupName+"' set "+clause+" where "+condition)
		if err != nil {
			return Rows{}, err
		}
		return Rows{int(rows)}, nil
	} else {
		return Rows{}, updateItemError{"updateItemError: can't update itemName!"}
	}
}

// CheckItems check whether the given items existed in the given group
func (gdb *Gdb) CheckItems(groupName string, itemNames ...string) error {
	for _, itemName := range itemNames {
		sqlString := "select 1 from '" + groupName + "' where itemName='" + itemName + "' limit 1"
		if r, err := query(gdb.ItemDbPath, sqlString); err != nil {
			return err
		} else {
			if len(r) == 0 {
				return fmt.Errorf("itemName: " + itemName + " not existed")
			}
		}
	}
	return nil
}
