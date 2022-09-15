package u_sql

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	SQL_SQLITE string = `SQLITE`

	SQL_NOT_NULL          string = `NOT NULL`
	SQL_UNIQUE            string = `UNIQUE`
	SQLITE_AUTO_INCREMENT string = `AUTOINCREMENT`

	SQLITE_ROWID   string = `rowid`
	SQLITE_TEXT    string = `TEXT`
	SQLITE_INTEGER string = `INTEGER`
	SQLITE_NUMERIC string = `NUMERIC`
	SQLITE_BLOB    string = `BLOB`
	SQLITE_REAL    string = `REAL`
)

type Sql struct {
	DriverName     string
	DataSourceName string
	task           bool
	DB             *sql.DB
	DbType         DbType
}

type DbType int

const (
	SQLite DbType = 0 + iota
	MySQL
)

var sqlTypes = []string{`sqlite`, `mysql`}

func (dt DbType) String() string {
	if dt >= SQLite && dt <= MySQL {
		return sqlTypes[dt]
	}
	return `DB type not support`
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

func (s *Sql) Task(task func() error) (err error) {
	err = s.Open()
	if err != nil {
		return
	}
	defer s.DB.Close()
	return task()
}

func (s *Sql) NewQuery() *query {
	return &query{}
}

func (s *Sql) NewUpdate() *update {
	return &update{data: map[string]interface{}{}}
}

func (s *Sql) isSQLite() bool {
	return s.DbType == SQLite
}

func (s *Sql) isRowid(field string) bool {
	return field == SQLITE_ROWID
}

func (s *Sql) orm(tableName string, ormStruct interface{}) Orm {
	typeOf := reflect.TypeOf(ormStruct)
	valueOf := reflect.ValueOf(ormStruct)
	primaryKeyCount := 0

	for i := 0; i < typeOf.NumField(); i++ {
		fieldIndex := typeOf.Field(i).Tag.Get(`index`)
		if strings.Contains(strings.ToLower(fieldIndex), `pk`) {
			primaryKeyCount = primaryKeyCount + 1
		}
	}

	columns := make([]Column, 0)
	for i := 0; i < typeOf.NumField(); i++ {
		col := Column{}
		valueOfInterface := valueOf.Field(i).Interface()

		// Field
		field := typeOf.Field(i).Tag.Get(`json`)
		if field == `` {
			field = snakeString(field)
		}
		col.Field = field

		// Field type
		fieldType := typeOf.Field(i).Tag.Get(`type`)
		if fieldType == `` {
			if s.isSQLite() {
				// SQLite
				switch valueOfInterface.(type) {
				case string:
					fieldType = SQLITE_TEXT
				case int, int8, int16, int32, int64:
					fieldType = SQLITE_INTEGER
				case float32, float64:
					fieldType = SQLITE_NUMERIC
				}
			} else {
				// Others undo
			}
		}
		col.Type = fieldType

		// Index
		fieldIndex := typeOf.Field(i).Tag.Get(`index`)

		// Primary Key & Auto Increment
		if strings.Contains(strings.ToLower(fieldIndex), `pk`) {
			col.PrimaryKey = true
			if strings.Contains(strings.ToLower(fieldIndex), `ai`) && primaryKeyCount == 1 {
				col.AutoIncrement = true
			}
		}

		// Not Null
		if strings.Contains(strings.ToLower(fieldIndex), `nn`) {
			col.NotNull = true
		}

		// Unique
		if strings.Contains(strings.ToLower(fieldIndex), `un`) {
			col.Unique = true
		}

		// Default value
		fieldDefault := typeOf.Field(i).Tag.Get(`default`)
		if fieldDefault != `` {
			switch valueOfInterface.(type) {
			case string:
				col.ValueDefault = fmt.Sprintf(`"%s"`, fieldDefault)
			default:
				col.ValueDefault = fmt.Sprintf(`%v`, fieldDefault)
			}
		}

		// Value
		col.Value = valueOfInterface

		columns = append(columns, col)
	}

	return Orm{Table: tableName, Columns: columns, PrimaryKeyCount: primaryKeyCount}
}

func snakeString(val string) string {
	valLen := len(val)
	hasUnderLine := false
	valBytes := make([]byte, 0)

	for i := 0; i < valLen; i++ {
		v := val[i]
		if v == '_' {
			hasUnderLine = true
		}
		if i > 0 && v >= 'A' && v <= 'Z' && !hasUnderLine {
			valBytes = append(valBytes, '_')
			hasUnderLine = false
		}
		valBytes = append(valBytes, v)
	}
	return strings.ToLower(string(valBytes))
}
