/*
creatTime: 2020/12/8
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

import (
	"gdb/sqlite"
	"github.com/syndtr/goleveldb/leveldb"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

/*
ItemHandler
*/

func (ldb *LevelDb) AddItems(itemInfo AddItemInfo) (Rows, error) {
	groupName := itemInfo.GroupName
	itemValues := itemInfo.Values
	c, err := sqlite.Query("PRAGMA table_info([" + groupName + "])") // get column names
	if err != nil {
		return Rows{-1}, err
	}
	c = c[1:]                      // abort first item id
	columnNames := []string{}      // without '
	addedColumnNames := []string{} // with '
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
	itemNames := []string{}
	for _, itemValue := range itemValues {
		t := []string{}
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
				return Rows{-1}, ColumnNameError{"ColumnNameError: " + columnNames[i]}
			}

		}
		addedItemValues = append(addedItemValues, t)
	}
	if err := sqlite.InsertItems(insertSqlString, addedItemValues...); err != nil {
		return Rows{-1}, err
	}
	for _, itemName := range itemNames {
		ldb.RtDbFilter.Set(itemName, struct{}{})
	}
	// initial write realTime data, all key write ''
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		batch := leveldb.Batch{}
		for i := 0; i < len(itemNames); i++ {
			batch.Put([]byte(itemNames[i]), []byte("0"))
		}
		_ = ldb.RtDb.Write(&batch, nil)
		return
	}()
	wg.Wait()
	return Rows{len(itemValues)}, nil
}

func (ldb *LevelDb) DeleteItems(itemInfo ItemInfo) (Rows, error) {
	groupName := itemInfo.GroupName
	condition := itemInfo.Condition
	item, err := sqlite.Query("select itemName from '" + groupName + "' where " + condition)
	if len(item) == 0 {
		return Rows{}, conditionError{"conditionError: " + condition}
	}
	if err != nil {
		return Rows{}, err
	}
	rows, err := sqlite.UpdateItem("delete from '" + groupName + "' where " + condition)
	if err != nil {
		return Rows{}, err
	}
	ldb.RtDbFilter.Remove(item[0]["itemName"]) // remove key from bloom filter
	return Rows{int(rows)}, nil
}

func GetItems(itemInfo ItemInfo) ([]map[string]string, error) {
	groupName := itemInfo.GroupName // groupName
	columns := itemInfo.Column      // column
	startRow := itemInfo.StartRow   // startRow
	condition := itemInfo.Condition // condition
	var rows []map[string]string
	var err error
	if startRow == -1 {
		// all rows
		rows, err = sqlite.Query("select " + columns + " from '" + groupName + "' where " + condition)
		if err != nil {
			return nil, err
		}
	} else {
		// Limit query
		rowCount := itemInfo.RowCount
		rows, err = sqlite.Query("select " + columns + " from '" + groupName + "' where " + condition + " Limit " + strconv.Itoa(rowCount) + " offset " + strconv.Itoa(startRow))
		if err != nil {
			return nil, err
		}
	}
	return rows, nil
}

func GetItemsWithCount(itemInfo ItemInfo) (Items, error) {
	condition := itemInfo.Condition
	groupName := itemInfo.GroupName
	c, err := sqlite.Query("select count(*) as count from '" + groupName + "' where " + condition)
	if err != nil {
		return Items{}, err
	}
	itemValues, err := GetItems(itemInfo)
	if err != nil {
		return Items{}, nil
	}
	count, err := strconv.ParseInt(c[0]["count"], 10, 64)
	if err != nil {
		return Items{}, nil
	}
	return Items{count, itemValues}, nil
}

func (ldb *LevelDb) UpdateItems(itemInfo ItemInfo) (Rows, error) {
	groupName := itemInfo.GroupName
	condition := itemInfo.Condition
	clause := itemInfo.Clause
	// Firstly determine whether to update itemName
	regPoint := regexp.MustCompile(`(?i)itemName='(.*?)'`) // Match the content after itemName
	if !regPoint.Match([]byte(clause)) {
		// no itemName
		// update SQLite
		rows, err := sqlite.UpdateItem("update '" + groupName + "' set " + clause + " where " + condition)
		if err != nil {
			return Rows{}, err
		}
		return Rows{int(rows)}, nil
	} else {
		return Rows{}, updateItemError{"updateItemError: can't update itemName!"}
	}
}
