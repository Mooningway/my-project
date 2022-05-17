package u_sql

func (s *Sql) Update(sqlQuery string, args ...interface{}) (affectCount int64, err error) {
	err = s.Open()
	defer s.Db.Close()
	if err != nil {
		return
	}
	return s.UpdateExec(sqlQuery, args...)
}

func (s *Sql) UpdateExec(sqlQuery string, args ...interface{}) (affectCount int64, err error) {
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
