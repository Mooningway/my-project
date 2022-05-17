package conf_sql

// Database configuration
import "my-project/src/utils/u_sql"

const (
	driverName string = `sqlite3`
	// DbPath is required
	dbPath string = ``
)

func InitSqlite() u_sql.Sql {
	return u_sql.Sql{DriverName: driverName, DataSourceName: dbPath}
}
