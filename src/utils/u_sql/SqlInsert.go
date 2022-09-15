package u_sql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func (s *Sql) Insert(tableName string, ormStruct interface{}) (id int64, err error) {
	columns := make([]string, 0)
	values := make([]interface{}, 0)
	marks := make([]string, 0)

	orm := s.orm(tableName, ormStruct)
	for _, c := range orm.Columns {
		if c.isPKandAI() || (s.isSQLite() && s.isRowid(c.Field)) {
			continue
		}

		columns = append(columns, c.Field)
		values = append(values, c.Value)
		marks = append(marks, `?`)
	}

	insertSql := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, tableName, strings.Join(columns, `,`), strings.Join(marks, `,`))
	return s.InsertSql(insertSql, values...)
}

func (s *Sql) InsertByUpdate(tableName string, u update) (id int64, err error) {
	if u.data == nil {
		return
	}

	columns := make([]string, 0)
	values := make([]interface{}, 0)
	marks := make([]string, 0)
	for column, value := range u.data {
		columns = append(columns, column)
		values = append(values, value)
		marks = append(marks, `?`)
	}

	insertSql := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, tableName, strings.Join(columns, `,`), strings.Join(marks, `,`))
	return s.InsertSql(insertSql, values...)
}

func (s *Sql) InsertMore(tableName string, ormStructs interface{}) (id int64, err error) {
	valueOf := reflect.ValueOf(ormStructs)
	if valueOf.Kind() != reflect.Slice && valueOf.Kind() != reflect.Array {
		err = errors.New(`insert error, ormStructs is not array`)
		return
	}

	array := make([]interface{}, 0)
	for i := 0; i < valueOf.Len(); i++ {
		array = append(array, valueOf.Index(i).Interface())
	}

	columns := make([]string, 0)
	values := make([]interface{}, 0)
	itemMarks := make([]string, 0)
	marks := make([]string, 0)

	for oi, ormStruct := range array {
		orm := s.orm(tableName, ormStruct)
		for _, c := range orm.Columns {
			if c.isPKandAI() || (s.isSQLite() && s.isRowid(c.Field)) {
				continue
			}

			// Columns
			if oi == 0 {
				columns = append(columns, c.Field)
				itemMarks = append(itemMarks, `?`)
			}

			// Values
			values = append(values, c.Value)
		}

		marks = append(marks, `(`+strings.Join(itemMarks, `,`)+`)`)
	}

	insertSql := fmt.Sprintf(`INSERT INTO %s (%s) VALUES %s`, tableName, strings.Join(columns, `,`), strings.Join(marks, `,`))
	return s.InsertSql(insertSql, values...)
}

func (s *Sql) InsertSql(insertSql string, values ...interface{}) (id int64, err error) {
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
