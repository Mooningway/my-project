package u_sql

import (
	"database/sql"
	"encoding/json"
)

func (s *Sql) QueryById(sqlQuery string, result interface{}, id int64) (err error) {
	err = s.Open()
	defer s.Db.Close()
	if err != nil {
		return
	}
	return s.QueryByIdExec(sqlQuery, result, id)
}

func (s *Sql) QueryByIdExec(sqlQuery string, result interface{}, id int64) (err error) {
	rows, err := s.Db.Query(sqlQuery, id)
	if err != nil {
		return
	}
	jsonResult, err := handleQueryOne(rows)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonResult, result)
	return
}

func (s *Sql) QuerySlice(sqlQuery string, result interface{}, args ...interface{}) (err error) {
	err = s.Open()
	defer s.Db.Close()
	if err != nil {
		return
	}
	return s.QuerySliceExec(sqlQuery, result, args...)
}

func (s *Sql) QuerySliceExec(sqlQuery string, result interface{}, args ...interface{}) (err error) {
	var rows *sql.Rows
	if len(args) > 0 {
		rows, err = s.Db.Query(sqlQuery, args...)
	} else {
		rows, err = s.Db.Query(sqlQuery)
	}
	if err != nil {
		return
	}
	jsonResult, err := handleQuerySlice(rows)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonResult, result)
	return
}

func (s *Sql) QueryOne(sqlQuery string, result interface{}, args ...interface{}) (err error) {
	err = s.Open()
	defer s.Db.Close()
	if err != nil {
		return
	}
	return s.QueryOneExec(sqlQuery, result, args...)
}

func (s *Sql) QueryOneExec(sqlQuery string, result interface{}, args ...interface{}) (err error) {
	var rows *sql.Rows
	if len(args) > 0 {
		rows, err = s.Db.Query(sqlQuery, args...)
	} else {
		rows, err = s.Db.Query(sqlQuery)
	}
	if err != nil {
		return
	}
	jsonResult, err := handleQueryOne(rows)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonResult, result)
	return
}
