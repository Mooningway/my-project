package u_sql

import (
	"bytes"
	"database/sql"
	"fmt"
)

func (s *Sql) Count(table string, w where) (count int64, err error) {
	var sqlBuffer bytes.Buffer
	sqlBuffer.WriteString(fmt.Sprintf(`SELECT COUNT(*) FROM %s`, table))

	whereSql, whereValues := w.toSql()
	if len(whereSql) > 0 {
		sqlBuffer.WriteString(whereSql)
	}

	countSql := sqlBuffer.String()

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
	if len(whereValues) == 0 {
		rows, err = stmt.Query()
	} else {
		rows, err = stmt.Query(whereValues...)
	}
	if err != nil {
		return
	}

	for rows.Next() {
		rows.Scan(&count)
	}
	return
}
