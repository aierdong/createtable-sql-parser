package types

type DbType = string

const (
	String   DbType = "string"
	Char     DbType = "char"
	Integer  DbType = "integer"
	Numeric  DbType = "numeric"
	Date     DbType = "date"
	Time     DbType = "time"
	DateTime DbType = "datetime"
	Boolean  DbType = "boolean"
)
