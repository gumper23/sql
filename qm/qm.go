package qm

import (
	"database/sql"
	"fmt"
	"time"
)

// QueryRows Executes query on db, returns an array of maps of the resultset in colname->value format.
// Column names are in the second return value.
func QueryRows(db *sql.DB, query string) ([]map[string]string, []string, error) {
	var resultsets []map[string]string
	var cols []string

	rows, err := db.Query(query)
	if err != nil {
		return resultsets, cols, err
	}
	defer rows.Close()

	cols, err = rows.Columns()
	if err != nil {
		return resultsets, cols, err
	}

	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
	}

	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			return resultsets, cols, err
		}
		resultset := make(map[string]string)
		for i := 0; i < len(vals); i++ {
			value := vals[i].(*interface{})
			switch v := (*value).(type) {
			case nil:
				resultset[cols[i]] = "NULL"
			case bool:
				if v {
					resultset[cols[i]] = "true"
				} else {
					resultset[cols[i]] = "false"
				}
			case []byte:
				resultset[cols[i]] = string(v)
			case time.Time:
				resultset[cols[i]] = v.Format("2006-01-02 15:04:05.999")
			case int64:
				resultset[cols[i]] = fmt.Sprintf("%d", v)
			default:
				resultset[cols[i]] = v.(string)
			}
		}
		resultsets = append(resultsets, resultset)
	}

	if err = rows.Err(); err != nil {
		return resultsets, cols, err
	}

	return resultsets, cols, nil
}

// QueryRow returns the first row from QueryRows().
func QueryRow(db *sql.DB, query string) (map[string]string, []string, error) {
	var result map[string]string
	var cols []string

	resultsets, cols, err := QueryRows(db, query)
	if err != nil {
		return result, cols, err
	}

	if len(resultsets) == 0 {
		return result, cols, sql.ErrNoRows
	}

	return resultsets[0], cols, nil
}
