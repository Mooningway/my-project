package u_sql

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

func (s *Sql) CreateTable(tableName string, ormStruct interface{}) error {
	if !s.task {
		err := s.Open()
		if err != nil {
			return err
		}
		defer s.DB.Close()
	}
	sqlString := ``
	if s.isSQLite() {
		// SQLite
		sqlString = s.sqliteCreateTableSql(tableName, ormStruct)
	} else {
		// Others undo
	}
	_, err := s.DB.Exec(sqlString)
	if err != nil {
		out := fmt.Sprintf(`Create table[%s] error: %v`, tableName, err)
		err = errors.New(out)
	}
	return err
}

func (s *Sql) CreateTables(tableMap map[string]interface{}) []error {
	errorArray := make([]error, 0)
	err := s.Open()
	if err != nil {
		errorArray = append(errorArray, err)
		return errorArray
	}
	defer s.DB.Close()

	for tableName, ormStruct := range tableMap {
		if s.isSQLite() {
			// SQLite
			sqlString := s.sqliteCreateTableSql(tableName, ormStruct)
			_, err := s.DB.Exec(sqlString)
			if err != nil {
				out := fmt.Sprintf(`Create table[%s] error: %v`, tableName, err)
				errorArray = append(errorArray, errors.New(out))
			}
		} else {
			// Others undo
		}
	}
	return errorArray
}

func (s *Sql) sqliteCreateTableSql(tableName string, ormStruct interface{}) string {
	orm := s.orm(tableName, ormStruct)

	columns := make([]string, 0)
	primaryKeys := make([]string, 0)
	hasAutoIncrement := false
	for _, c := range orm.Columns {
		if c.Field == SQLITE_ROWID {
			continue
		}

		var columnBuffer bytes.Buffer
		// Field, Field Type
		columnBuffer.WriteString(fmt.Sprintf(`"%s" %s`, c.Field, c.Type))
		// Not Null
		if c.NotNull {
			columnBuffer.WriteString(` ` + SQL_NOT_NULL)
		}
		// Unique
		if c.Unique {
			columnBuffer.WriteString(` ` + SQL_UNIQUE)
		}
		// default value
		if c.ValueDefault != nil {
			columnBuffer.WriteString(fmt.Sprintf(` DEFAULT %v`, c.ValueDefault))
		}
		columns = append(columns, columnBuffer.String())
		// Primary Key
		if c.PrimaryKey {
			primaryKeys = append(primaryKeys, fmt.Sprintf(`"%s"`, c.Field))
		}
		// Auto Increment
		if c.AutoIncrement && orm.PrimaryKeyCount == 1 {
			hasAutoIncrement = true
		}
	}

	// Primary Key and Auto Increment
	primaryKeySql := ``
	if orm.PrimaryKeyCount == 1 {
		if hasAutoIncrement {
			primaryKeySql = fmt.Sprintf(`,PRIMARY KEY(%s %s)`, strings.Join(primaryKeys, `,`), SQLITE_AUTO_INCREMENT)
		} else {
			primaryKeySql = fmt.Sprintf(`,PRIMARY KEY(%s)`, strings.Join(primaryKeys, `,`))
		}
	} else if orm.PrimaryKeyCount > 1 {
		primaryKeySql = fmt.Sprintf(`,PRIMARY KEY(%s)`, strings.Join(primaryKeys, `,`))
	}

	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (%s %s)`, tableName, strings.Join(columns, `,`), primaryKeySql)
}
