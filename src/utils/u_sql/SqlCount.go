package u_sql

import "database/sql"

func (s *Sql) Count(sqlQuery string, args ...interface{}) (count int64, err error) {
	err = s.Open()
	defer s.Db.Close()
	if err != nil {
		return
	}
	return s.CountExec(sqlQuery, args...)
}

func (s *Sql) CountExec(sqlQuery string, args ...interface{}) (count int64, err error) {
	var rows *sql.Rows
	if len(args) > 0 {
		rows, err = s.Db.Query(sqlQuery, args...)
	} else {
		rows, err = s.Db.Query(sqlQuery)
	}
	if err != nil {
		return
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	return
}
