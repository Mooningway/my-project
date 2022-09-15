package u_sql

import (
	"bytes"
	"fmt"
	"strings"
)

type query struct {
	where    []whereCondition
	orderBy  []orderByCondition
	pageNum  int64
	pageSize int64
}

const (
	and  = `AND`
	or   = `OR`
	eq   = `=`
	ne   = `!=`
	like = `LIKE`
	lt   = `<`
	lte  = `<=`
	gt   = `>`
	gte  = `>=`
	asc  = `ASC`
	desc = `DESC`
)

type whereCondition struct {
	condition  string
	column     string
	expression string
	value      interface{}
}

type orderByCondition struct {
	column string
	order  string
}

// Query

func (q *query) toSql() (whereSql string, values []interface{}) {
	var sqlBuffer bytes.Buffer

	// Where
	wheres := q.where
	if len(wheres) > 0 {
		for ci, cv := range wheres {
			if ci == 0 {
				sqlBuffer.WriteString(` WHERE ` + cv.column + ` ` + cv.expression + ` ?`)
			} else {
				sqlBuffer.WriteString(` ` + cv.condition + ` ` + cv.column + ` ` + cv.expression + ` ?`)
			}
			values = append(values, cv.value)
		}
	}

	// Order by
	orderBys := q.orderBy
	if len(orderBys) > 0 {
		orderSql := make([]string, 0)
		for _, cv := range orderBys {
			orderSql = append(orderSql, cv.column+` `+cv.order)
		}

		sqlBuffer.WriteString(` ORDER BY `)
		sqlBuffer.WriteString(strings.Join(orderSql, `,`))
	}

	// Limit
	if q.pageNum > 0 && q.pageSize > 0 {
		sqlBuffer.WriteString(` LIMIT ?, ?`)
		values = append(values, (q.pageNum-1)*q.pageSize, q.pageSize)
	}

	return sqlBuffer.String(), values
}

// Query - where & and

func (q *query) AndEq(column string, value interface{}) *query {
	q.and(column, value, eq)
	return q
}

func (q *query) AndNe(column string, value interface{}) *query {
	q.and(column, value, ne)
	return q
}

func (q *query) AndLike(column string, value interface{}) *query {
	q.and(column, `%`+fmt.Sprintf(`%v`, value)+`%`, like)
	return q
}

func (q *query) AndLt(column string, value interface{}) *query {
	q.and(column, value, lt)
	return q
}

func (q *query) AndLte(column string, value interface{}) *query {
	q.and(column, value, lte)
	return q
}

func (q *query) AndGt(column string, value interface{}) *query {
	q.and(column, value, gt)
	return q
}

func (q *query) AndGte(column string, value interface{}) *query {
	q.and(column, value, gte)
	return q
}

func (q *query) and(column string, value interface{}, expression string) {
	q.where = append(q.where, whereCondition{condition: and, column: column, expression: expression, value: value})
}

// Query - where & or

func (q *query) OrEq(column string, value interface{}) *query {
	q.or(column, value, eq)
	return q
}

func (q *query) OrNe(column string, value interface{}) *query {
	q.or(column, value, ne)
	return q
}

func (q *query) OrLike(column string, value interface{}) *query {
	q.or(column, `%`+fmt.Sprintf(`%v`, value)+`%`, like)
	return q
}

func (q *query) OrLt(column string, value interface{}) *query {
	q.or(column, value, lt)
	return q
}

func (q *query) OrLte(column string, value interface{}) *query {
	q.or(column, value, lte)
	return q
}

func (q *query) OrGt(column string, value interface{}) *query {
	q.or(column, value, gt)
	return q
}

func (q *query) OrGte(column string, value interface{}) *query {
	q.or(column, value, gte)
	return q
}

func (q *query) or(column string, value interface{}, expression string) {
	q.where = append(q.where, whereCondition{condition: or, column: column, expression: expression, value: value})
}

// Order By

func (q *query) Asc(column string) *query {
	q.orderBy = append(q.orderBy, orderByCondition{column: column, order: asc})
	return q
}

func (q *query) Desc(column string) *query {
	q.orderBy = append(q.orderBy, orderByCondition{column: column, order: desc})
	return q
}

// Page
func (q *query) Page(pageNum, pageSize int64) *query {
	q.pageNum = pageNum
	q.pageSize = pageSize
	return q
}
