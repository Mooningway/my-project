package u_sql

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

func (s *Sql) FindOne(table string, w where, formatData interface{}, columns ...string) (err error) {
	selectSql, values := s.findToSql(true, table, w, columns...)
	if err != nil {
		return
	}
	return s.FindOneSql(selectSql, formatData, values...)
}

func (s *Sql) FindSlice(table string, w where, formatData interface{}, columns ...string) (err error) {
	selectSql, values := s.findToSql(false, table, w, columns...)
	if err != nil {
		return
	}
	return s.FindSliceSql(selectSql, formatData, values...)
}

func (s *Sql) FindOneSql(selectSql string, formatData interface{}, values ...interface{}) (err error) {
	if !s.task {
		err = s.Open()
		if err != nil {
			return
		}
		defer s.DB.Close()
	}

	stmt, err := s.DB.Prepare(selectSql)
	if err != nil {
		return
	}

	var rows *sql.Rows
	if len(values) == 0 {
		rows, err = stmt.Query()
	} else {
		rows, err = stmt.Query(values...)
	}

	jsonBytes, err := handleFindOne(rows)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonBytes, formatData)
	return
}

func (s *Sql) FindSliceSql(selectSql string, formatData interface{}, values ...interface{}) (err error) {
	if !s.task {
		err = s.Open()
		if err != nil {
			return
		}
		defer s.DB.Close()
	}

	stmt, err := s.DB.Prepare(selectSql)
	if err != nil {
		return
	}

	var rows *sql.Rows
	if len(values) == 0 {
		rows, err = stmt.Query()
	} else {
		rows, err = stmt.Query(values...)
	}

	jsonBytes, err := handleFindSlice(rows)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonBytes, formatData)
	return
}

func (s *Sql) findToSql(one bool, table string, w where, columns ...string) (selectSql string, values []interface{}) {
	columnsSql := `*`
	if len(columns) > 0 {
		columnsSql = strings.Join(columns, `,`)
	}

	var sqlBuffer bytes.Buffer
	sqlBuffer.WriteString(fmt.Sprintf(`SELECT %s FROM %s`, columnsSql, table))

	whereSql, values := w.toSql()
	if len(whereSql) > 0 {
		sqlBuffer.WriteString(whereSql)
	}

	if one {
		sqlBuffer.WriteString(` LIMIT 1`)
	}

	selectSql = sqlBuffer.String()
	return
}

func handleFindSlice(rows *sql.Rows) (jsonBytes []byte, err error) {
	tableData, err := handleFindResult(rows)
	if err != nil {
		return
	}
	return json.Marshal(tableData)
}

func handleFindOne(rows *sql.Rows) (jsonBytes []byte, err error) {
	tableData, err := handleFindResult(rows)
	if len(tableData) > 0 {
		jsonBytes, _ = json.Marshal(tableData[0])
	}
	return
}

func handleFindResult(rows *sql.Rows) ([]map[string]interface{}, error) {
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
