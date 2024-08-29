package types

type Dialect string

const (
	MySQL      Dialect = "mysql"
	PostgreSQL Dialect = "postgres"
	Oracle     Dialect = "oracle"
	SQLServer  Dialect = "sqlserver"
	SQLite3    Dialect = "sqlite3"
	Hive       Dialect = "hive"
)
