package u_sql

import (
	"bytes"
	"fmt"
)

func (s *Sql) DeleteById(tableName string, id interface{}) (affectCount int64, err error) {
	var deleteSql string
	if s.isSQLite() {
		// SQLite
		deleteSql = fmt.Sprintf(`DELETE FROM %s WHERE rowid = ?`, tableName)
	} else {
		// Others
		deleteSql = fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, tableName)
	}
	return s.DeleteSql(deleteSql, id)
}

func (s *Sql) Delete(table string, q query) (affectCount int64, err error) {
	var sqlBuffer bytes.Buffer
	sqlBuffer.WriteString(fmt.Sprintf(`DELETE FROM %s`, table))

	whereSql, whereValues := q.toSql()
	if len(whereSql) > 0 {
		sqlBuffer.WriteString(whereSql)
	}

	deleteSql := sqlBuffer.String()
	return s.DeleteSql(deleteSql, whereValues...)
}

func (s *Sql) DeleteSql(deleteSql string, values ...interface{}) (affectCount int64, err error) {
	if !s.task {
		err = s.Open()
		if err != nil {
			return
		}
		defer s.DB.Close()
	}

	stmt, err := s.DB.Prepare(deleteSql)
	if err != nil {
		return
	}

	result, err := stmt.Exec(values...)
	if err != nil {
		return
	}

	affectCount, _ = result.RowsAffected()
	return
}
