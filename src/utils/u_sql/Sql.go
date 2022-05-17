package u_sql

import (
	"database/sql"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

type Sql struct {
	DriverName     string
	DataSourceName string
	Db             *sql.DB
}

func (s *Sql) Open() (err error) {
	db, err := sql.Open(s.DriverName, s.DataSourceName)
	if err != nil {
		return
	}
	s.Db = db
	return
}

func (s *Sql) Exec(biz func() (err error)) (err error) {
	err = s.Open()
	if err != nil {
		return
	}
	defer s.Db.Close()
	return biz()
}

// Handle query result

func handleQuerySlice(rows *sql.Rows) (jsonBytes []byte, err error) {
	tableData, err := handleQueryResult(rows)
	if err != nil {
		return
	}
	return json.Marshal(tableData)
}

func handleQueryOne(rows *sql.Rows) (jsonBytes []byte, err error) {
	tableData, err := handleQueryResult(rows)
	if len(tableData) > 0 {
		jsonBytes, _ = json.Marshal(tableData[0])
	}
	return
}

func handleQueryResult(rows *sql.Rows) ([]map[string]interface{}, error) {
	tableData := make([]map[string]interface{}, 0)
	columns, err := rows.Columns()
	if err != nil {
		return tableData, err
	}

	columnsCount := len(columns)
	values := make([]interface{}, columnsCount)
	props := make([]interface{}, columnsCount)
	for rows.Next() {
		for i := 0; i < columnsCount; i++ {
			props[i] = &values[i]
		}
		rows.Scan(props...)

		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			valByte, ok := val.([]byte)

			if ok {
				v = string(valByte)
			} else {
				v = val
			}
			entry[col] = v
		}

		tableData = append(tableData, entry)
	}
	return tableData, nil
}
