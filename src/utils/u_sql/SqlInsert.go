package u_sql

import (
	"fmt"
	"strings"
)

func (s *Sql) Insert(table string, cv columnValue) (id int64, err error) {
	columns := make([]string, 0)
	values := make([]interface{}, 0)
	placeholders := make([]string, 0)

	for _, col := range cv.columns {
		columns = append(columns, col)
		values = append(values, cv.values[col])
		placeholders = append(placeholders, `?`)
	}

	insertSql := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, table, strings.Join(columns, `,`), strings.Join(placeholders, `,`))

	if !s.task {
		err = s.Open()
		if err != nil {
			return
		}
		defer s.DB.Close()
	}

	stmt, err := s.DB.Prepare(insertSql)
	if err != nil {
		return
	}

	result, err := stmt.Exec(values...)
	if err != nil {
		return
	}

	id, _ = result.LastInsertId()
	return
}

func (s *Sql) InsertMore(table string, columns []string, count int, values ...interface{}) (id int64, err error) {
	placeholder := make([]string, 0)
	for i := 0; i < len(columns); i++ {
		placeholder = append(placeholder, `?`)
	}
	placeholders := make([]string, 0)
	for i := 0; i < count; i++ {
		placeholders = append(placeholders, `(`+strings.Join(placeholder, `,`)+`)`)
	}

	insertSql := fmt.Sprintf(`INSERT INTO %s (%s) VALUES %s`, table, strings.Join(columns, `,`), strings.Join(placeholders, `,`))

	if !s.task {
		err = s.Open()
		if err != nil {
			return
		}
		defer s.DB.Close()
	}

	stmt, err := s.DB.Prepare(insertSql)
	if err != nil {
		return
	}

	result, err := stmt.Exec(values...)
	if err != nil {
		return
	}

	id, _ = result.LastInsertId()
	return
}
