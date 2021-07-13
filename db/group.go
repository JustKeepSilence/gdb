/*
creatTime: 2020/11/11
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"fmt"
	. "github.com/ahmetb/go-linq/v3"
	"github.com/deckarep/golang-set"
	"os"
	"strings"
	"time"
)

const (
	gdbKeyWords = `
				  break      default       func     interface   	select
				  case       defer         go       map         	struct
				  chan       else          goto     package     	switch
			      const      fallthrough   if       range       	type
				  continue   for           import   return      	var
                  gdb		 from          where    char        	varchar
                  int        smallint      numeric  real        	double 
				  precision  float         primary  key         	foreign 
				  references not           null     create      	table
				  insert     into          values   delete	    	update
			      set        where   	   drop 	alter 			add
 				  truncate   distinct      all      and 			or
			 	  join		 as 		   * 		order 			by
				  desc		 asc		   between  union 			except
				  is		 avg		   min		max				sum
				  count		 group		   having   realTimeData	historicalData 
    			  `
)

// AddGroups add group to gdb, you can't add duplicate group to gdb and the groupName can't be gdbKeyWords
//
//for columns you MUST NOT add column of id, itemName and dataType because these three columns will be automatically added.
// The operation is atomic
func (gdb *Gdb) AddGroups(groupInfos ...AddedGroupInfo) (TimeRows, error) {
	st := time.Now()
	groupNames := [][]string{}
	columnNames := [][]string{}
	gs := []string{}
	for _, groupInfo := range groupInfos {
		groupName := strings.Trim(groupInfo.GroupName, " ")
		gs = append(gs, groupName)
		if strings.Contains(gdbKeyWords, strings.ToLower(groupName)) {
			return TimeRows{}, fmt.Errorf("groupNameError:" + groupName) // illegal groupName
		}
		index, r := checkColumnNames(groupInfo.ColumnNames...)
		{
			if index != -1 {
				return TimeRows{}, fmt.Errorf("columnNameError:" + groupInfo.ColumnNames[index])
			}
			columnNames = append(columnNames, r)
			groupNames = append(groupNames, []string{groupName})
		} // check columnName
	}
	err1 := gdb.insertItems("insert into group_cfg (groupName) values(?)", groupNames...) // add group to group_cfg
	if err1 != nil {
		// fail adding
		return TimeRows{}, err1
	}
	for i := 0; i < len(groupNames); i++ {
		sb := strings.Builder{}
		sb.Write(convertStringToByte("create table if not exists `"))
		sb.Write(convertStringToByte(groupNames[i][0]))
		// have custom columns
		if item := columnNames[i]; len(item) != 0 {
			sb.Write(convertStringToByte("` ( " + gdb.key + ", itemName " + gdb.unique + " UNIQUE, " + "dataType " + gdb.unique + ", ")) // system three columns :id,pointName,groupName
			sb.Write(convertStringToByte(strings.Join(columnNames[i], " text, ")))
			sb.Write(convertStringToByte(" text)"))
		} else {
			// no custom columns
			sb.Write(convertStringToByte("` ( " + gdb.key + ", itemName " + gdb.unique + " UNIQUE, dataType text )"))
		}
		createTableSqlString := sb.String()
		if err := gdb.updateItems([]string{createTableSqlString}...); err != nil {
			// fail in creating table
			gdb.rollBack(groupNames...)
			return TimeRows{}, err
		}
	}
	// generate pointer
	if err := gdb.generateDBPointer(gs); err != nil {
		return TimeRows{}, err
	}
	// add groups to gdb
	gdb.groupNames = append(gdb.groupNames, gs...)
	// add group successfully
	return TimeRows{EffectedRows: len(groupNames), Times: time.Since(st).Milliseconds()}, nil
}

// DeleteGroups delete group from gdb,this operation will  delete items in itemDb as well as the historyData of whole group.
func (gdb *Gdb) DeleteGroups(groupInfos GroupNamesInfo) (TimeRows, error) {
	st := time.Now()
	c := 0
	groupNames := groupInfos.GroupNames
	// you can't delete calc group
	if From(groupNames).Contains("calc") {
		return TimeRows{}, fmt.Errorf("groupError: you can't delete calc group")
	}
	for _, groupName := range groupNames {
		// delete columns from group_cfg
		err := gdb.updateItems([]string{"delete from group_cfg where groupName='" + groupName + "'"}...)
		if err != nil {
			return TimeRows{}, err
		}
		// delete keys in filter==>you can't write data to gdb anymore
		items, _ := gdb.query("select itemName from `" + groupName + "`")
		for _, item := range items {
			itemName := item["itemName"]
			gdb.rtDbFilter.Remove(itemName + joiner + groupName)
		}
		// drop table
		if err := gdb.updateItems([]string{"drop table `" + groupName + "`"}...); err != nil {
			return TimeRows{}, err
		}
		// release db pointer
		for _, dataType := range dataTypes {
			// close db
			if err := gdb.hisDb[dataType][groupName].Close(); err != nil {
				return TimeRows{}, err
			}
			delete(gdb.hisDb[dataType], groupName) // release pointer
			// delete files
			if err := os.RemoveAll(gdb.dbPath + gdb.delimiter + "historicalData" + gdb.delimiter + dataType + gdb.delimiter + groupName); err != nil {
				return TimeRows{}, err
			}
		}
		c++
	}
	// delete groupNames in gdb
	gs := []string{}
	for _, groupName := range groupNames {
		if !From(gdb.groupNames).Contains(groupName) {
			gs = append(gs, groupName)
		}
	}
	gdb.groupNames = gs
	return TimeRows{EffectedRows: c, Times: time.Since(st).Milliseconds()}, nil
}

// GetGroups get group name
func (gdb *Gdb) GetGroups() (GroupNamesInfo, error) {
	return GroupNamesInfo{gdb.groupNames}, nil
}

// GetGroupProperty get the column name and item count of the given groupName and condition
// condition is the where clause of sqlite
func (gdb *Gdb) GetGroupProperty(groupName, condition string) (GroupPropertyInfo, error) {
	var c []map[string]string
	if gdb.driverName == "sqlite3" {
		var err error
		c, err = gdb.query("PRAGMA table_info([" + groupName + "])") // get column names
		if err != nil {
			return GroupPropertyInfo{}, err
		}
	} else {
		var err error
		c, err = gdb.query("select column_name as name from information_schema.columns where table_schema='" + gdb.itemDbName + "' and table_name='" + groupName + "'")
		if err != nil {
			return GroupPropertyInfo{}, err
		}
	}
	itemCount, err := gdb.query("select count(*) as count from `" + groupName + "` where " + condition)
	if err != nil {
		return GroupPropertyInfo{}, err
	}
	columnNames := []string{}
	for i := 0; i < len(c); i++ {
		columnNames = append(columnNames, c[i]["name"])
	}
	return GroupPropertyInfo{ItemCount: itemCount[0]["count"], ItemColumnNames: columnNames[1:]}, nil
}

// UpdateGroupNames update groupNames,the history data of oldGroup will migrate to new newGroup as well
func (gdb *Gdb) UpdateGroupNames(groupInfos ...UpdatedGroupNameInfo) (TimeRows, error) {
	st := time.Now()
	c := 0
	sqlStrings := []string{}
	for _, groupInfo := range groupInfos {
		oldGroupName := groupInfo.OldGroupName
		newGroupName := groupInfo.NewGroupName
		if strings.Contains(gdbKeyWords, strings.ToLower(newGroupName)) {
			return TimeRows{}, fmt.Errorf("groupNameError:" + newGroupName) // illegal groupName
		}
		if oldGroupName == "calc" {
			return TimeRows{}, fmt.Errorf("groupError: you can't update calc groupName")
		}
		sqlStrings = append(sqlStrings, "update group_cfg set groupName='"+newGroupName+"' where groupName='"+oldGroupName+"'")
		c++ // update successfully
	}
	if err := gdb.updateItems(sqlStrings...); err != nil {
		return TimeRows{}, err
	}
	for _, groupInfo := range groupInfos {
		oldGroupName := groupInfo.OldGroupName
		newGroupName := groupInfo.NewGroupName
		//alter table name
		if err := gdb.updateItems([]string{"alter table `" + oldGroupName + "` rename to `" + newGroupName + "`"}...); err != nil {
			// rollback
			_ = gdb.updateItems([]string{"update group_cfg set groupName='" + oldGroupName + "' where groupName='" + newGroupName + "'"}...)
			return TimeRows{}, err
		}
		// update itemNames in gdb.Filter and db pointer
		items, _ := gdb.query("select itemName, dataType from `" + newGroupName + "`")
		for _, item := range items {
			itemName := item["itemName"] + joiner + newGroupName // itemName in filter = itemName + "__" + groupName
			oldItemName := item["itemName"] + joiner + oldGroupName
			dataType := item["dataType"]
			gdb.rtDbFilter.Remove(oldItemName)     // remove old itemNames in filter
			gdb.rtDbFilter.Set(itemName, dataType) // add key to filter, don't lock
		}
		for _, dataType := range dataTypes {
			// close db
			if err := gdb.hisDb[dataType][oldGroupName].Close(); err != nil {
				return TimeRows{}, err
			}
			delete(gdb.hisDb[dataType], oldGroupName) // release db pointer
			// rename folder
			if err := os.Rename(gdb.dbPath+gdb.delimiter+"historicalData"+gdb.delimiter+dataType+gdb.delimiter+oldGroupName, gdb.dbPath+gdb.delimiter+"historicalData"+gdb.delimiter+dataType+gdb.delimiter+newGroupName); err != nil {
				return TimeRows{}, err
			}
			// update groupNames in gdb
			for index, name := range gdb.groupNames {
				if name == oldGroupName {
					gdb.groupNames[index] = newGroupName
				}
			}
		}
		// generate db pointer
		if err := gdb.generateDBPointer([]string{newGroupName}); err != nil {
			return TimeRows{}, err
		}
	}
	return TimeRows{EffectedRows: c, Times: time.Since(st).Milliseconds()}, nil
}

// UpdateGroupColumnNames update column names of group,you can't update columns of id, itemName, dataType's column name.
// the updatedColumnName MUST be existed
//
// The operation is atomic
func (gdb *Gdb) UpdateGroupColumnNames(info UpdatedGroupColumnNamesInfo) (TimeCols, error) {
	st := time.Now()
	oldColumnNames, newColumnNames, groupName := info.OldColumnNames, info.NewColumnNames, info.GroupName
	if len(oldColumnNames) != len(newColumnNames) {
		return TimeCols{}, fmt.Errorf("columnNameError: inconsistent columnNames")
	}
	index, addedColumnNames := checkColumnNames(newColumnNames...) // check whether new columnName is valid
	if index != -1 {
		return TimeCols{}, fmt.Errorf("columnNameError: invalid columnName " + addedColumnNames[index])
	}
	index, checkedOldColumnNames := checkColumnNames(oldColumnNames...) // check whether to modify id, itemName
	if index != -1 {
		return TimeCols{}, fmt.Errorf("columnError: can't modify column " + oldColumnNames[index])
	}
	sqlStrings := []string{}
	for i := 0; i < len(checkedOldColumnNames); i++ {
		sqlStrings = append(sqlStrings, "alter table `"+groupName+"` rename column `"+checkedOldColumnNames[i]+"` to `"+addedColumnNames[i]+"`")
	}
	err := gdb.updateItems(sqlStrings...)
	if err != nil {
		return TimeCols{}, err
	}
	return TimeCols{EffectedCols: len(newColumnNames), Times: time.Since(st).Milliseconds()}, nil
}

// DeleteGroupColumns delete columns from group, you can't delete id, itemName, dataType column
//
// The operation is atomic
func (gdb *Gdb) DeleteGroupColumns(info DeletedGroupColumnNamesInfo) (TimeCols, error) {
	st := time.Now()
	groupName, deletedColumnNames := info.GroupName, info.ColumnNames
	if contains(deletedColumnNames...) {
		return TimeCols{}, fmt.Errorf("columnNameError")
	}
	if gdb.driverName == "mysql" {
		sqlString := []string{}
		for _, columnName := range deletedColumnNames {
			sqlString = append(sqlString, "alter table `"+groupName+"` drop column `"+columnName+"`")
		}
		if err := gdb.updateItems(sqlString...); err != nil {
			return TimeCols{}, err
		}
		return TimeCols{EffectedCols: len(deletedColumnNames), Times: time.Since(st).Milliseconds()}, nil
	}
	r, err := gdb.GetGroupProperty(groupName, "1=1") // get existed columns of given group
	if err != nil {
		return TimeCols{}, err
	}
	// try drop t1_backup
	_ = gdb.updateItems([]string{"drop table t1_backup"}...)
	cs := mapset.NewSet(convertStringToInterface(r.ItemColumnNames...)...)  // existed columns, include itemName, not include id
	ds := mapset.NewSet(convertStringToInterface(deletedColumnNames...)...) // deleted groups
	// check whether the column to be deleted exist
	if !ds.IsSubset(cs) {
		return TimeCols{}, fmt.Errorf("columnNameError: some columns don't exist")
	}
	rs := cs.Difference(ds)
	rs.Remove("itemName")
	rs.Remove("dataType")
	remainedColumnNames := rs.ToSlice()
	newColumnNames := []string{} // remained columns in new table
	for _, name := range remainedColumnNames {
		newColumnNames = append(newColumnNames, name.(string))
	}
	sb := strings.Builder{}
	sb.Write(convertStringToByte("CREATE TEMPORARY TABLE t1_backup("))
	if len(newColumnNames) == 0 {
		sb.Write(convertStringToByte("id integer, itemName text, dataType text"))
	} else {
		sb.Write(convertStringToByte("id integer, itemName text, dataType text, "))
		sb.Write(convertStringToByte(strings.Join(newColumnNames, " text, ")))
		sb.Write(convertStringToByte(" text "))
	}
	if len(newColumnNames) == 0 {
		sb.Write(convertStringToByte("); insert into `t1_backup` select id, itemName, dataType "))
	} else {
		sb.Write(convertStringToByte("); insert into `t1_backup` select id, itemName, dataType, "))
		sb.Write(convertStringToByte(strings.Join(newColumnNames, ",")))
	}
	sb.Write(convertStringToByte(" from `"))
	sb.Write(convertStringToByte(groupName))
	sb.Write(convertStringToByte("`; drop table `"))
	sb.Write(convertStringToByte(groupName))
	sb.Write(convertStringToByte("`; create table if not exists `"))
	sb.Write(convertStringToByte(groupName))
	if item := newColumnNames; len(item) != 0 {
		sb.Write(convertStringToByte("` (" + gdb.key + ", itemName " + gdb.unique + " UNIQUE,dataType text, "))
		sb.Write(convertStringToByte(strings.Join(newColumnNames, " text, ")))
		sb.Write(convertStringToByte(" text)"))
	} else {
		sb.Write(convertStringToByte("` ( " + gdb.key + ", itemName " + gdb.unique + " UNIQUE, dataType text )"))
	}
	if len(newColumnNames) == 0 {
		sb.Write(convertStringToByte("; insert into `" + groupName + "` select id, itemName, dataType "))
	} else {
		sb.Write(convertStringToByte("; insert into `" + groupName + "` select id, itemName, dataType, " + strings.Join(newColumnNames, ",")))
	}
	sb.Write(convertStringToByte(" from `t1_backup` "))
	sb.Write(convertStringToByte("; drop table `t1_backup`"))
	sqlString := sb.String()
	if err := gdb.updateItems([]string{sqlString}...); err != nil {
		return TimeCols{}, err
	}
	return TimeCols{EffectedCols: len(deletedColumnNames), Times: time.Since(st).Milliseconds()}, nil
}

// AddGroupColumns add columns to group, all columns type are text default.And you can't add duplicate columns
func (gdb *Gdb) AddGroupColumns(info AddedGroupColumnsInfo) (TimeCols, error) {
	st := time.Now()
	groupName, addedColumnNames, defaultValues := info.GroupName, info.ColumnNames, info.DefaultValues
	if len(addedColumnNames) != len(defaultValues) {
		return TimeCols{}, fmt.Errorf("inconsist of columnNames and defaultValues")
	}
	sqlStrings := []string{}
	for index, name := range addedColumnNames {
		if len(strings.Trim(defaultValues[index], " ")) == 0 {
			sqlStrings = append(sqlStrings, "alter table `"+groupName+"` add column `"+name+"` "+gdb.unique+" default '' ")
		} else {
			sqlStrings = append(sqlStrings, "alter table `"+groupName+"` add column `"+name+"` "+gdb.unique+" default '"+defaultValues[index]+"'")
		}
	}
	if err := gdb.updateItems(sqlStrings...); err != nil {
		return TimeCols{}, err
	}
	return TimeCols{EffectedCols: len(addedColumnNames), Times: time.Since(st).Milliseconds()}, nil
}

// CleanGroupItems will delete all items and corresponding history data in the given groups,so after this operation
// you can't write or get data from this group any more,you MUST use this method carefully
func (gdb *Gdb) CleanGroupItems(groupNames ...string) (TimeRows, error) {
	sqliteString := []string{}
	for _, groupName := range groupNames {
		items, _ := gdb.query("select itemName from `" + groupName + "`")
		for _, item := range items {
			// delete keys in filter==>you can't write data to gdb
			gdb.rtDbFilter.Remove(item["itemName"] + joiner + groupName)
		}
		if gdb.driverName == "sqlite3" {
			sqliteString = append(sqliteString, "delete from `"+groupName+"`")
		} else {
			sqliteString = append(sqliteString, "truncate table `"+groupName+"`")
		}
	}
	if err := gdb.updateItems(sqliteString...); err != nil {
		return TimeRows{}, err
	} else {
		for _, groupName := range groupNames {
			for _, dataType := range dataTypes {
				// close db
				if err := gdb.hisDb[dataType][groupName].Close(); err != nil {
					return TimeRows{}, err
				}
				delete(gdb.hisDb[dataType], groupName) // release pointer
				// delete files
				if err := os.RemoveAll(gdb.dbPath + gdb.delimiter + "historicalData" + gdb.delimiter + dataType + gdb.delimiter + groupName); err != nil {
					return TimeRows{}, err
				}
			}
		}
		// generate pointer again
		if err := gdb.generateDBPointer(groupNames); err != nil {
			return TimeRows{}, err
		}
		return TimeRows{EffectedRows: len(groupNames)}, nil
	}
}

// rollback when failing creating table: firstly delete column in group_cfg and then drop table
func (gdb *Gdb) rollBack(groupNames ...[]string) {
	var deletedGroupNames []string
	for j := 0; j < len(groupNames); j++ {
		deletedGroupNames = append(deletedGroupNames, "`"+groupNames[j][0]+"`")
		_ = gdb.updateItems([]string{"drop table `" + groupNames[j][0] + "`"}...) // delete added table
		// An error indicates that the table does not exist
	}
	// Delete columns that have been added to group_cfg
	deleteGroupCfg := "delete from group_cfg where groupName=" + strings.Join(deletedGroupNames, " or groupName=")
	_ = gdb.updateItems([]string{deleteGroupCfg}...)
}

// check whether column name is valid, trim ‘ and empty string between the column name
// if all column names are valid the the index is -1, column names can't be one of id, itemName, groupName and empty string
func checkColumnNames(columnNames ...string) (int, []string) {
	r := []string{}
	for index, columnName := range columnNames {
		c := strings.Trim(strings.Replace(columnName, "'", "", -1), " ")
		if c == "id" || c == "itemName" || c == "dataType" || len(c) == 0 {
			return index, nil
		}
		r = append(r, c)
	}
	return -1, r
}

// check whether deleted column contains itemName, id or empty string
func contains(s ...string) bool {
	for _, s2 := range s {
		t := strings.Trim(s2, " ")
		if t == "itemName" || t == "id" || t == "dataType" || len(t) == 0 {
			return true
		}
	}
	return false
}
