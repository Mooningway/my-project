package u_sql

import (
	"fmt"
	"strings"
)

func (s *Sql) Update(table string, cv columnValue, w where) (affectCount int64, err error) {
	setValues := make([]interface{}, 0)

	setSql := make([]string, 0)
	for _, col := range cv.columns {
		setSqlTemp := fmt.Sprintf(`%s = ?`, col)
		setSql = append(setSql, setSqlTemp)
		setValues = append(setValues, cv.values[col])
	}

	whereSql, whereValues := w.toSql()

	values := make([]interface{}, 0)
	values = append(values, setValues...)
	if len(whereValues) > 0 {
		values = append(values, whereValues...)
	}

	var updateSql string
	if len(whereValues) > 0 {
		updateSql = fmt.Sprintf(`UPDATE %s SET %s %s`, table, strings.Join(setSql, `,`), whereSql)
	} else {
		updateSql = fmt.Sprintf(`UPDATE %s SET %s`, table, strings.Join(setSql, `,`))
	}

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
