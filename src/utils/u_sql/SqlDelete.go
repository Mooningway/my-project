package u_sql

func (s *Sql) Delete(sqlQuery string, args ...interface{}) (affectCount int64, err error) {
	err = s.Open()
	defer s.Db.Close()
	if err != nil {
		return
	}
	return s.DeleteExec(sqlQuery, args...)
}

func (s *Sql) DeleteExec(sqlQuery string, args ...interface{}) (affectCount int64, err error) {
	stmt, err := s.Db.Prepare(sqlQuery)
	if err != nil {
		return
	}
	result, err := stmt.Exec(args...)
	if err != nil {
		return
	}
	affectCount, err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}
