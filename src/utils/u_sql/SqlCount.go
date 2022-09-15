package u_sql

import (
	"bytes"
	"database/sql"
	"fmt"
)

func (s *Sql) Count(tableName string, q query) (count int64, err error) {
	var sqlBuffer bytes.Buffer
	sqlBuffer.WriteString(fmt.Sprintf(`SELECT COUNT(*) FROM %s`, tableName))

	whereSql, whereValues := q.toSql()
	if len(whereSql) > 0 {
		sqlBuffer.WriteString(whereSql)
	}

	countSql := sqlBuffer.String()

	return s.CountSql(countSql, whereValues...)
}

func (s *Sql) CountSql(countSql string, values ...interface{}) (count int64, err error) {
	if !s.task {
		err = s.Open()
		if err != nil {
			return
		}
		defer s.DB.Close()
	}

	stmt, err := s.DB.Prepare(countSql)
	if err != nil {
		return
	}

	var rows *sql.Rows
	if len(values) == 0 {
		rows, err = stmt.Query()
	} else {
		rows, err = stmt.Query(values...)
	}
	if err != nil {
		return
	}

	for rows.Next() {
		rows.Scan(&count)
	}
	return
}
