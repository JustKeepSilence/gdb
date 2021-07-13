/*
creatTime: 2020/11/10
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/sync/errgroup"
	"strings"
)

// connect SQLite, if not exist then create
func (gdb *Gdb) connectDataBase(driverName, dsn string) error {
	if db, err := sql.Open(driverName, dsn); err != nil {
		return fmt.Errorf("failded to connecting to database:" + err.Error())
	} else {
		if err := db.Ping(); err != nil {
			return err
		}
		gdb.itemDb = db
		gdb.driverName = driverName
		if driverName == "mysql" {
			gdb.itemDbName = strings.Split(strings.Split(dsn, "/")[1], "?")[0]
		}
		return nil
	}
}

// query data from SQLite, the format of return value is: [{columnName1: value1, columnName2: value2}]
func (gdb *Gdb) query(queryString string) ([]map[string]string, error) {
	rows, err := gdb.itemDb.Query(queryString)
	if err != nil {
		return nil, fmt.Errorf("sqlExecutionError:" + err.Error())
	}
	columnNames, _ := rows.Columns() // get column names
	var scanColumns []interface{}    // Store the address of the variable corresponding to each column to store the result of the Scan scan
	var rowsData []map[string]string // store the data of rows
	for i := 0; i < len(columnNames); i++ {
		var columnName sql.NullString                  // see:https://golang.org/pkg/database/sql/#NullString
		scanColumns = append(scanColumns, &columnName) //Create a new variable named columnName every time,
		// and put its address in, so that every time Scan is called, the number in scanColumns will be updated
		//correspondingly, scanColumns[index] in scanColumns is each index- in the corresponding database 1 column of data
	}
	for rows.Next() {
		err := rows.Scan(scanColumns...)
		if err != nil {
			// Write each row of data to scanColumns
			return nil, fmt.Errorf("sqlExecutionError:" + err.Error())
		}
		temp := make(map[string]string)
		for index, scanColumnValue := range scanColumns {
			ns := *(scanColumnValue.(*sql.NullString))
			if ns.Valid {
				// valid string
				temp[columnNames[index]] = ns.String
			} else {
				// invalid string
				temp[columnNames[index]] = "null"
			}
		}
		rowsData = append(rowsData, temp)
	}
	return rowsData, nil
}

// insert items into group
func (gdb *Gdb) insertItems(insertString string, rowValues ...[]string) error {
	tx, err := gdb.itemDb.Begin()
	if err != nil {
		return fmt.Errorf("sqlExecutionError:" + err.Error())
	}
	stmt, err := tx.Prepare(insertString)
	if err != nil {
		return fmt.Errorf("sqlExecutionError:" + err.Error())
	}
	eg := errgroup.Group{}
	for _, rowValue := range rowValues {
		insertedRows := rowValue
		eg.Go(func() error {
			var data []interface{} // data to be inserted
			for i := 0; i < len(insertedRows); i++ {
				data = append(data, insertedRows[i]) // can't use []string as []interface directly
			}
			_, err := stmt.Exec(data...)
			if err != nil {
				/* fail in inserting data
				tx.Rollback() is the key ,see https://github.com/mattn/go-sqlite3/issues/184
				*/
				//flagChan <- true
				return fmt.Errorf("sqlExecutionError:" + err.Error() + "\n")
			} else {
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		defer func(tx *sql.Tx) {
			_ = tx.Rollback()
		}(tx)
		defer func(stmt *sql.Stmt) {
			_ = stmt.Close()
		}(stmt)
		return err
	} else {
		if err := tx.Commit(); err != nil {
			defer func(tx *sql.Tx) {
				_ = tx.Rollback()
			}(tx)
			defer func(stmt *sql.Stmt) {
				_ = stmt.Close()
			}(stmt)
			return fmt.Errorf("sqlExecutionError:" + err.Error())
		} else {
			defer func(tx *sql.Tx) {
				_ = tx.Rollback()
			}(tx)
			defer func(stmt *sql.Stmt) {
				_ = stmt.Close()
			}(stmt)
			return nil
		}
	}
}

// update items with transaction
func (gdb *Gdb) updateItems(sqlStrings ...string) error {
	tx, err := gdb.itemDb.Begin()
	if err != nil {
		return fmt.Errorf("sqlExecutionError:" + err.Error())
	}
	for _, sqlString := range sqlStrings {
		if _, err := tx.Exec(sqlString); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("sqlExecutionError:" + err.Error())
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("sqlExecutionError:" + err.Error())
	}
	return nil
}

// update item
func (gdb *Gdb) updateItem(sqlString string) (int64, error) {
	r, err := gdb.itemDb.Exec(sqlString)
	if err != nil {
		return -1, fmt.Errorf("sqlExecutionError:" + err.Error())
	}
	e, _ := r.RowsAffected()
	return e, nil
}
