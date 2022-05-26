package u_sql

import (
	"bytes"
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Sql struct {
	DriverName     string
	DataSourceName string
	task           bool
	DB             *sql.DB
}

type columnValue struct {
	columns []string
	values  map[string]interface{}
}

type where struct {
	conditions []condition
	orders     []order
}

type condition struct {
	condition  string
	column     string
	expression string
	value      interface{}
}

type order struct {
	column string
	sort   string
}

// Connection
// Open sql connection
func (s *Sql) Open() (err error) {
	db, err := sql.Open(s.DriverName, s.DataSourceName)
	if err != nil {
		return
	}
	s.DB = db
	return
}

func (s *Sql) NewColumn() *columnValue {
	return &columnValue{values: map[string]interface{}{}}
}

func (c *columnValue) Set(column string, value interface{}) *columnValue {
	c.columns = append(c.columns, column)
	c.values[column] = value
	return c
}

func (s *Sql) NewWHere() *where {
	return &where{}
}

func (s *Sql) Task(task func() error) (err error) {
	s.task = true
	err = s.Open()
	if err != nil {
		return
	}
	defer s.DB.Close()
	return task()
}

func (w *where) AndEq(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `AND`, column: column, expression: `=`, value: value})
	return w
}

func (w *where) AndNq(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `AND`, column: column, expression: `!=`, value: value})
	return w
}

func (w *where) AndLike(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `AND`, column: column, expression: `LIKE`, value: value})
	return w
}

func (w *where) AndLt(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `AND`, column: column, expression: `<`, value: value})
	return w
}

func (w *where) AndLte(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `AND`, column: column, expression: `<=`, value: value})
	return w
}

func (w *where) AndGt(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `AND`, column: column, expression: `>=`, value: value})
	return w
}

func (w *where) AndGte(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `AND`, column: column, expression: `>=`, value: value})
	return w
}

func (w *where) OrEq(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `OR`, column: column, expression: `=`, value: value})
	return w
}

func (w *where) OrNq(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `OR`, column: column, expression: `!=`, value: value})
	return w
}

func (w *where) OrLike(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `OR`, column: column, expression: `LIKE`, value: value})
	return w
}

func (w *where) OrLt(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `OR`, column: column, expression: `<`, value: value})
	return w
}

func (w *where) OrLte(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `OR`, column: column, expression: `<=`, value: value})
	return w
}

func (w *where) OrGt(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `OR`, column: column, expression: `>`, value: value})
	return w
}

func (w *where) OrGte(column string, value interface{}) *where {
	w.conditions = append(w.conditions, condition{condition: `OR`, column: column, expression: `>=`, value: value})
	return w
}

func (w *where) Asc(column string) *where {
	w.orders = append(w.orders, order{column: column, sort: `ASC`})
	return w
}

func (w *where) Desc(column string) *where {
	w.orders = append(w.orders, order{column: column, sort: `DESC`})
	return w
}

func (w *where) toSql() (whereSql string, values []interface{}) {
	var sqlBuffer bytes.Buffer

	if len(w.conditions) > 0 {
		for ci, cv := range w.conditions {
			if ci == 0 {
				sqlBuffer.WriteString(` WHERE ` + cv.column + ` ` + cv.expression + ` ?`)
			} else {
				sqlBuffer.WriteString(` ` + cv.condition + ` ` + cv.column + ` ` + cv.expression + ` ?`)
			}
			values = append(values, cv.value)
		}
	}

	orderSql := make([]string, 0)
	if len(w.orders) > 0 {
		for _, ov := range w.orders {
			orderSql = append(orderSql, ov.column+` `+ov.sort)
		}
	}
	if len(orderSql) > 0 {
		sqlBuffer.WriteString(` ORDER BY `)
		sqlBuffer.WriteString(strings.Join(orderSql, `,`))
	}

	whereSql = sqlBuffer.String()
	return
}
