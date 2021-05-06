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
	"strings"
)

/*
Operating the group in leveldb is essentially operating the group_cfg table
and the corresponding group table in the SQLite database
			group_cfg
		group1   group2   group3.....
     c11,c12..  c21,c2.. c31,c32...
every group has two default columns:id and itemName, so when adding groups, needn't to
specify the information of theses to columns
*/

const (
	/* Predefined keywords, the name of the group can not be repeated in the first place
	   Among them, the first 4 lines are the keywords of go, gdb is the name of the database
	   and the following are all keywords of SQL. At the same time, it should be noted that
	   the name of the table cannot be an empty string.
	*/
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

// AddGroups add group to gdb
func (gdb *Gdb) AddGroups(groupInfos ...AddedGroupInfo) (Rows, error) {
	groupNames := [][]string{}
	columnNames := [][]string{}
	for _, groupInfo := range groupInfos {
		groupName := strings.Trim(groupInfo.GroupName, " ")
		if strings.Contains(gdbKeyWords, strings.ToLower(groupName)) {
			return Rows{}, groupNameError{"groupNameError:" + groupName} // illegal groupName
		}
		index, r := checkColumnNames(groupInfo.ColumnNames...) // check columnName
		if index != -1 {
			return Rows{}, columnNameError{"columnNameError:" + groupInfo.ColumnNames[index]}
		}
		columnNames = append(columnNames, r)
		groupNames = append(groupNames, []string{groupName})
	}
	err1 := insertItems(gdb.ItemDbPath, "insert into group_cfg (groupName) values(?)", groupNames...) // add group to group_cfg
	if err1 != nil {
		// fail adding
		return Rows{}, err1
	}
	for i := 0; i < len(groupNames); i++ {
		sb := strings.Builder{}
		sb.Write([]byte("create table if not exists '"))
		sb.Write([]byte(groupNames[i][0]))
		// have custom columns
		if item := columnNames[i]; len(item) != 0 {
			sb.Write([]byte("' (id integer not null primary key, itemName text UNIQUE,")) // system three columns :id,pointName,groupName
			sb.Write([]byte(strings.Join(columnNames[i], " text, ")))
			sb.Write([]byte(" text)"))
		} else {
			// no custom columns
			sb.Write([]byte("' (id integer not null primary key, itemName text UNIQUE )"))
		}
		createTableSqlString := sb.String()
		if err := updateItems(gdb.ItemDbPath, []string{createTableSqlString}...); err != nil {
			// fail in creating table
			gdb.rollBack(groupNames...)
			return Rows{}, err
		}
	}
	// add group successfully
	return Rows{len(groupNames)}, nil
}

// DeleteGroups delete group from gdb
func (gdb *Gdb) DeleteGroups(groupInfos GroupNamesInfo) (Rows, error) {
	c := 0
	groupNames := groupInfos.GroupNames
	// you can't delete calc group
	if From(groupNames).Contains("calc") {
		return Rows{}, fmt.Errorf("groupError: you can't delete calc group")
	}
	for _, groupName := range groupNames {
		// delete columns from group_cfg
		err := updateItems(gdb.ItemDbPath, []string{"delete from group_cfg where groupName='" + groupName + "'"}...)
		if err != nil {
			return Rows{}, err
		}
		// drop table
		if err := updateItems(gdb.ItemDbPath, []string{"drop table '" + groupName + "'"}...); err != nil {
			return Rows{}, err
		}
		c++
	}
	return Rows{c}, nil
}

// GetGroups get group name
func (gdb *Gdb) GetGroups() (GroupNamesInfo, error) {
	r, err := query(gdb.ItemDbPath, "select groupName from group_cfg")
	if err != nil {
		return GroupNamesInfo{}, err
	}
	groupNames := []string{}
	for _, item := range r {
		groupNames = append(groupNames, item["groupName"])
	}
	return GroupNamesInfo{groupNames}, nil
}

// GetGroupProperty get the column name and item count of the given groupName and condition
// condition is the where clause of sqlite
func (gdb *Gdb) GetGroupProperty(groupName, condition string) (GroupPropertyInfo, error) {
	c, err := query(gdb.ItemDbPath, "PRAGMA table_info(["+groupName+"])") // get column names of given table
	if err != nil {
		return GroupPropertyInfo{}, err
	}
	itemCount, err := query(gdb.ItemDbPath, "select count(*) as count from '"+groupName+"' where "+condition)
	if err != nil {
		return GroupPropertyInfo{}, err
	}
	columnNames := []string{}
	for i := 0; i < len(c); i++ {
		columnNames = append(columnNames, c[i]["name"])
	}
	return GroupPropertyInfo{ItemCount: itemCount[0]["count"], ItemColumnNames: columnNames[1:]}, nil
}

// UpdateGroupNames update groupNames
func (gdb *Gdb) UpdateGroupNames(groupInfos ...UpdatedGroupNameInfo) (Rows, error) {
	c := 0
	sqlStrings := []string{}
	for _, groupInfo := range groupInfos {
		oldGroupName := groupInfo.OldGroupName
		newGroupName := groupInfo.NewGroupName
		if strings.Contains(gdbKeyWords, strings.ToLower(newGroupName)) {
			return Rows{}, groupNameError{"groupNameError:" + newGroupName} // illegal groupName
		}
		sqlStrings = append(sqlStrings, "update group_cfg set groupName='"+newGroupName+"' where groupName='"+oldGroupName+"'")
		c++ // update successfully
	}
	if err := updateItems(gdb.ItemDbPath, sqlStrings...); err != nil {
		return Rows{}, err
	}
	for _, groupInfo := range groupInfos {
		oldGroupName := groupInfo.OldGroupName
		newGroupName := groupInfo.NewGroupName
		//alter table name
		if err := updateItems(gdb.ItemDbPath, []string{"alter table '" + oldGroupName + "' rename to '" + newGroupName + "'"}...); err != nil {
			// rollback
			_ = updateItems(gdb.ItemDbPath, []string{"update group_cfg set groupName='" + oldGroupName + "' where groupName='" + newGroupName + "'"}...)
			return Rows{}, err
		}
	}
	return Rows{c}, nil
}

// UpdateGroupColumnNames update column names of group
func (gdb *Gdb) UpdateGroupColumnNames(info UpdatedGroupColumnNamesInfo) (Cols, error) {
	oldColumnNames, newColumnNames, groupName := info.OldColumnNames, info.NewColumnNames, info.GroupName
	if len(oldColumnNames) != len(newColumnNames) {
		return Cols{}, columnNameError{"columnNameError: inconsistent columnNames"}
	}
	index, addedColumnNames := checkColumnNames(newColumnNames...) // check whether new columnName is valid
	if index != -1 {
		return Cols{}, columnNameError{"columnNameError: invalid columnName " + addedColumnNames[index]}
	}
	index, checkedOldColumnNames := checkColumnNames(oldColumnNames...) // check whether to modify id, itemName
	if index != -1 {
		return Cols{}, columnNameError{"columnError: can't modify column " + oldColumnNames[index]}
	}
	sqlStrings := []string{}
	for i := 0; i < len(checkedOldColumnNames); i++ {
		sqlStrings = append(sqlStrings, "alter table '"+groupName+"' rename column '"+checkedOldColumnNames[i]+"' to '"+addedColumnNames[i]+"'")
	}
	err := updateItems(gdb.ItemDbPath, sqlStrings...)
	if err != nil {
		return Cols{}, err
	}
	return Cols{len(newColumnNames)}, nil
}

// DeleteGroupColumns delete columns from group
func (gdb *Gdb) DeleteGroupColumns(info DeletedGroupColumnNamesInfo) (Cols, error) {
	groupName, deletedColumnNames := info.GroupName, info.ColumnNames
	if contains(deletedColumnNames...) {
		return Cols{}, columnNameError{"columnNameError"}
	}
	r, err := gdb.GetGroupProperty(groupName, "1=1") // get existed columns of given group
	if err != nil {
		return Cols{}, err
	}
	// try drop t1_backup
	_ = updateItems(gdb.ItemDbPath, []string{"drop table 't1_backup'"}...)
	cs := mapset.NewSet(convertStringToInterface(r.ItemColumnNames...)...)  // existed columns, include itemName, not include id
	ds := mapset.NewSet(convertStringToInterface(deletedColumnNames...)...) // deleted groups
	// check whether the column to be deleted exist
	if !ds.IsSubset(cs) {
		return Cols{}, columnNameError{"columnNameError: some columns don't exist"}
	}
	rs := cs.Difference(ds)
	rs.Remove("itemName")
	remainedColumnNames := rs.ToSlice()
	newColumnNames := []string{} // remained columns in new table
	for _, name := range remainedColumnNames {
		newColumnNames = append(newColumnNames, name.(string))
	}
	sb := strings.Builder{}
	sb.Write([]byte("CREATE TEMPORARY TABLE t1_backup("))
	if len(newColumnNames) == 0 {
		sb.Write([]byte("id text, itemName text "))
	} else {
		sb.Write([]byte("id text, itemName text, "))
	}
	sb.Write([]byte(strings.Join(newColumnNames, " text, ")))
	if len(newColumnNames) == 0 {
		sb.Write([]byte("); insert into t1_backup select id, itemName "))
	} else {
		sb.Write([]byte("); insert into t1_backup select id, itemName, "))
		sb.Write([]byte(strings.Join(newColumnNames, ",")))
	}
	sb.Write([]byte(" from '"))
	sb.Write([]byte(groupName))
	sb.Write([]byte("'; drop table '"))
	sb.Write([]byte(groupName))
	sb.Write([]byte("'; create table if not exists '"))
	sb.Write([]byte(groupName))
	if item := newColumnNames; len(item) != 0 {
		sb.Write([]byte("' (id integer not null primary key, itemName text UNIQUE,"))
		sb.Write([]byte(strings.Join(newColumnNames, " text, ")))
		sb.Write([]byte(" text)"))
	} else {
		sb.Write([]byte("' (id integer not null primary key, itemName text UNIQUE )"))
	}
	if len(newColumnNames) == 0 {
		sb.Write([]byte("; insert into '" + groupName + "' select id, itemName "))
	} else {
		sb.Write([]byte("; insert into '" + groupName + "' select id, itemName, " + strings.Join(newColumnNames, ",")))
	}
	sb.Write([]byte(" from 't1_backup' "))
	sb.Write([]byte("; drop table t1_backup"))
	sqlString := sb.String()
	if err := updateItems(gdb.ItemDbPath, []string{sqlString}...); err != nil {
		return Cols{}, err
	}
	return Cols{len(deletedColumnNames)}, nil
}

// AddGroupColumns add columns to group, all columns type are text
func (gdb *Gdb) AddGroupColumns(info AddedGroupColumnsInfo) (Cols, error) {
	groupName, addedColumnNames, defaultValues := info.GroupName, info.ColumnNames, info.DefaultValues
	if len(addedColumnNames) != len(defaultValues) {
		return Cols{}, fmt.Errorf("inconsist of columnNames and defaultValues")
	}
	sqlStrings := []string{}
	for index, name := range addedColumnNames {
		sqlStrings = append(sqlStrings, "alter table '"+groupName+"' add column '"+name+"' text default '"+defaultValues[index]+"'")
	}
	if err := updateItems(gdb.ItemDbPath, sqlStrings...); err != nil {
		return Cols{}, err
	}
	return Cols{len(addedColumnNames)}, nil
}

func (gdb *Gdb) CleanGroupItems(groupNames ...string) (Rows, error) {
	sqliteString := []string{}
	for _, groupName := range groupNames {
		sqliteString = append(sqliteString, "delete from '"+groupName+"'")
	}
	if err := updateItems(gdb.ItemDbPath, sqliteString...); err != nil {
		return Rows{}, err
	} else {
		return Rows{EffectedRows: len(groupNames)}, nil
	}
}

// rollback when failing creating table: firstly delete column in group_cfg and then drop table
func (gdb *Gdb) rollBack(groupNames ...[]string) {
	var deletedGroupNames []string
	for j := 0; j < len(groupNames); j++ {
		deletedGroupNames = append(deletedGroupNames, "'"+groupNames[j][0]+"'")
		_ = updateItems(gdb.ItemDbPath, []string{"drop table '" + groupNames[j][0] + "'"}...) // delete added table
		// An error indicates that the table does not exist
	}
	// Delete columns that have been added to group_cfg
	deleteGroupCfg := "delete from group_cfg where groupName=" + strings.Join(deletedGroupNames, " or groupName=")
	_ = updateItems(gdb.ItemDbPath, []string{deleteGroupCfg}...)
}

// check whether column name is valid, trim ‘ and empty string between the column name
// if all column names are valid the the index is -1, column names can't be one of id, itemName, groupName and empty string
func checkColumnNames(columnNames ...string) (int, []string) {
	r := []string{}
	for index, columnName := range columnNames {
		c := strings.Trim(strings.Replace(columnName, "'", "", -1), " ") // 去除两端空格和'
		if strings.ToLower(c) == "id" || strings.ToLower(c) == "itemname" || len(c) == 0 {
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
		if t == "itemName" || t == "id" || len(t) == 0 {
			return true
		}
	}
	return false
}

// convert string slice to interface slice
func convertStringToInterface(s ...string) []interface{} {
	v := make([]interface{}, len(s))
	for _, s2 := range s {
		v = append(v, s2)
	}
	return v
}
