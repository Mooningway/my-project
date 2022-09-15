package u_sql

import (
	"bytes"
	"fmt"
	"strings"
)

type update struct {
	data map[string]interface{}
}

func (u *update) Set(column string, value interface{}) *update {
	u.data[column] = value
	return u
}

func (s *Sql) UpdateById(tableName string, ormStruct interface{}, id interface{}) (affectCount int64, err error) {
	setSql := make([]string, 0)
	values := make([]interface{}, 0)

	orm := s.orm(tableName, ormStruct)
	for _, c := range orm.Columns {
		if c.isPKandAI() || (s.isSQLite() && s.isRowid(c.Field)) {
			continue
		}

		setSql = append(setSql, fmt.Sprintf(`%s = ?`, c.Field))
		values = append(values, c.Value)
	}
	values = append(values, id)

	var updateSql string
	if s.isSQLite() {
		// SQLite
		updateSql = fmt.Sprintf(`UPDATE %s SET %s WHERE rowid = ?`, tableName, strings.Join(setSql, `,`))
	} else {
		// Others
		updateSql = fmt.Sprintf(`UPDATE %s SET %s WHERE id = ?`, tableName, strings.Join(setSql, `,`))
	}
	return s.UpdateSql(updateSql, values...)
}

func (s *Sql) Update(tableName string, u update, q query) (affectCount int64, err error) {
	if u.data == nil {
		return
	}

	setSql := make([]string, 0)
	setValues := make([]interface{}, 0)
	for column, value := range u.data {
		setSql = append(setSql, fmt.Sprintf(`%s = ?`, column))
		setValues = append(setValues, value)
	}

	whereSql, whereValues := q.toSql()

	values := make([]interface{}, 0)
	values = append(values, setValues...)
	if len(whereValues) > 0 {
		values = append(values, whereValues...)
	}

	var updateSql bytes.Buffer
	updateSql.WriteString(fmt.Sprintf(`UPDATE %s SET %s`, tableName, strings.Join(setSql, `,`)))
	if len(whereSql) > 0 {
		updateSql.WriteString(whereSql)
	}
	return s.UpdateSql(updateSql.String(), values...)
}

func (s *Sql) UpdateSql(updateSql string, values ...interface{}) (affectCount int64, err error) {
	if !s.task {
		err = s.Open()
		if err != nil {
			return
		}
		defer s.DB.Close()
	}

	stmt, err := s.DB.Prepare(updateSql)
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
