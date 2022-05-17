package u_sql

func (s *Sql) Insert(sqlQuery string, args ...interface{}) (id int64, err error) {
	err = s.Open()
	defer s.Db.Close()
	if err != nil {
		return
	}
	return s.InsertExec(sqlQuery, args...)
}

func (s *Sql) InsertExec(sqlQuery string, args ...interface{}) (id int64, err error) {
	stmt, err := s.Db.Prepare(sqlQuery)
	if err != nil {
		return
	}
	result, err := stmt.Exec(args...)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
