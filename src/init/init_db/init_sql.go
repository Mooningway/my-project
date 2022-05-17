package init_db

const sql_create_exrate_code string = `
CREATE TABLE if not exists "exrate_code" (
	"name"	TEXT,
	"code"	TEXT,
	"sort"	INTEGER
)`

const sql_create_exrate_rate string = `
CREATE TABLE if not exists "exrate_rate" (
	"date_unix"	INTEGER,
	"code"		TEXT,
	"rates"		BLOB
)`
