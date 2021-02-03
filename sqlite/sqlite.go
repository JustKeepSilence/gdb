/*
creatTime: 2020/11/10 14:51
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

//const sqlitePath = "./db/ldb.db" // SQLite path

type insertResult struct {
	Err error
}

// SQLite errors
type sqliteConnectionError struct {
	ErrorInfo string
}

func (ct sqliteConnectionError) Error() string {
	return ct.ErrorInfo
}

type sqliteExecutionError struct {
	ErrorInfo string
}

func (se sqliteExecutionError) Error() string {
	return se.ErrorInfo
}

type sqliteRowsError struct {
	ErrorInfo string
}

func (sr sqliteRowsError) Error() string {
	return sr.ErrorInfo
}

type sqliteTransactionError struct {
	ErrorInfo string
}

func (st sqliteTransactionError) Error() string {
	return st.ErrorInfo
}

// connect SQLite, if not exist then create
func connectSqlite(sqlitePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqlitePath)
	if db == nil {
		return nil, sqliteConnectionError{"sqliteConnectionError: null sqlite pointer"}
	}
	if err != nil {
		return nil, sqliteConnectionError{"sqliteConnectionError: " + err.Error()}
	}
	return db, nil
}

// query data from SQLite, the format of return value is: [{columnName1: value1, columnName2: value2}]
func Query(sqlitePath, queryString string) ([]map[string]string, error) {
	db, err := connectSqlite(sqlitePath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(queryString)
	if err != nil {
		return nil, sqliteExecutionError{"sqliteExecutionError: " + err.Error()}
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
			return nil, sqliteRowsError{"sqliteRowsError: " + err.Error()}
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
func InsertItems(sqlitePath, insertString string, rowValues ...[]string) error {
	db, err := connectSqlite(sqlitePath)
	if err != nil {
		return sqliteConnectionError{"sqliteConnectionError: " + err.Error()}
	}
	tx, err := db.Begin()
	if err != nil {
		return sqliteTransactionError{"sqliteTransactionError: " + err.Error()}
	}
	stmt, err := tx.Prepare(insertString)
	if err != nil {
		return sqliteTransactionError{"sqliteTransactionError: " + err.Error()}
	}
	var resultChan = make(chan insertResult, len(rowValues)) // Create an error type channel to handle errors
	var errorFlag bool                                       // Whether an error occurred, the default value is false
	var errorMessage string
	var wg sync.WaitGroup
	for index, rowValue := range rowValues {
		wg.Add(1)
		go func(insertedRows []string, st *sql.Stmt, j int) {
			defer wg.Done()
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
				resultChan <- insertResult{sqliteExecutionError{"sqliteExecutionError: " + err.Error() + "\n"}}
				return
			} else {
				resultChan <- insertResult{nil}
				return
			}
		}(rowValue, stmt, index)
	}
	go func() {
		wg.Wait()
		close(resultChan) // close channel
	}()
	// The for loop is used in the pipeline. When the channel is closed and there is no value to receive,
	// it jumps out of the loop, so you need to close the channel in the goroutine
	for rc := range resultChan {
		if rc.Err != nil {
			errorFlag = true // errors
			errorMessage += rc.Err.Error()
		}
	}
	if !errorFlag {
		// There is no error in the whole execution process
		if err := tx.Commit(); err != nil {
			defer tx.Rollback()
			defer db.Close()
			defer stmt.Close()
			return sqliteTransactionError{"sqliteTransactionError: " + err.Error()}
		} else {
			defer tx.Rollback()
			defer db.Close()
			defer stmt.Close()
			return nil
		}
	} else {
		// errors
		defer tx.Rollback()
		defer db.Close()
		defer stmt.Close()
		return fmt.Errorf(errorMessage)
	}
}

// update items with transaction
func UpdateItems(sqlitePath string, sqlStrings ...string) error {
	db, err := connectSqlite(sqlitePath)
	if err != nil {
		return sqliteConnectionError{"sqliteConnectionError: " + err.Error()}
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return sqliteTransactionError{"sqliteTransactionError: " + err.Error()}
	}
	for _, sqlString := range sqlStrings {
		if _, err := tx.Exec(sqlString, nil); err != nil {
			tx.Rollback()
			return sqliteExecutionError{"sqliteExecutionError: " + err.Error()}
		}
	}
	if err := tx.Commit(); err != nil {
		return sqliteExecutionError{"sqliteExecutionError: " + err.Error()}
	}
	return nil
}

// update item
func UpdateItem(sqlitePath, sqlString string) (int64, error) {
	db, err := connectSqlite(sqlitePath)
	if err != nil {
		return -1, sqliteConnectionError{"sqliteConnectionError: " + err.Error()}
	}
	defer db.Close()
	r, err := db.Exec(sqlString, nil)
	if err != nil {
		return -1, sqliteConnectionError{"sqliteConnectionError: " + err.Error()}
	}
	e, _ := r.RowsAffected()
	return e, nil
}
