package types

var MySQLTypeMap = map[string]string{
	"varchar":    "string",
	"char":       "string",
	"text":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"longtext":   "string",
	"int":        "integer",
	"integer":    "integer",
	"tinyint":    "integer",
	"smallint":   "integer",
	"mediumint":  "integer",
	"bigint":     "integer",
	"decimal":    "numeric",
	"numeric":    "numeric",
	"float":      "numeric",
	"double":     "numeric",
	"real":       "numeric",
	"date":       "date",
	"datetime":   "datetime",
	"timestamp":  "datetime",
	"time":       "datetime",
}

var PgTypeMap = map[string]string{
	"int2":      "integer",
	"int4":      "integer",
	"int8":      "integer",
	"smallint":  "integer",
	"integer":   "integer",
	"bigint":    "integer",
	"serial":    "integer",
	"bigserial": "integer",
	"timestamp": "datetime",
	"date":      "datetime",
	"time":      "datetime",
	"bool":      "boolean",
	"boolean":   "boolean",
	"varchar":   "string",
	"char":      "string",
	"text":      "string",
	"numeric":   "decimal",
	"decimal":   "decimal",
	"real":      "decimal",
	"double":    "decimal",
}

var PLSqlTypeMap = map[string]string{
	"BINARY_INTEGER":   "integer",
	"PLS_INTEGER":      "integer",
	"NATURAL":          "integer",
	"BINARY_FLOAT":     "decimal",
	"BINARY_DOUBLE":    "decimal",
	"NATURALN":         "integer",
	"POSITIVE":         "integer",
	"POSITIVEN":        "integer",
	"SIGNTYPE":         "integer",
	"SIMPLE_INTEGER":   "integer",
	"NVARCHAR2":        "string",
	"DEC":              "decimal",
	"INTEGER":          "integer",
	"INT":              "integer",
	"NUMERIC":          "decimal",
	"SMALLINT":         "integer",
	"NUMBER":           "decimal",
	"DECIMAL":          "decimal",
	"DOUBLE PRECISION": "decimal",
	"FLOAT":            "decimal",
	"REAL":             "decimal",
	"NCHAR":            "string",
	// "LONG RAW":                       "",
	"CHAR":      "string",
	"CHARACTER": "string",
	"VARCHAR2":  "string",
	"VARCHAR":   "string",
	"STRING":    "string",
	"RAW":       "",
	"BOOLEAN":   "boolean",
	"DATE":      "datetime",
	// "ROWID":                          "",
	// "UROWID":                         "",
	// "YEAR":                           "",
	// "MONTH":                          "",
	// "DAY":                            "",
	// "HOUR":                           "",
	// "MINUTE":                         "",
	// "SECOND":                         "",
	// "TIMEZONE_HOUR":                  "",
	// "TIMEZONE_MINUTE":                "",
	// "TIMEZONE_REGION":                "",
	// "TIMEZONE_ABBR":                  "",
	"TIMESTAMP": "datetime",
	// "TIMESTAMP_UNCONSTRAINED":        "",
	// "TIMESTAMP_TZ_UNCONSTRAINED":     "",
	// "TIMESTAMP_LTZ_UNCONSTRAINED":    "",
	// "YMINTERVAL_UNCONSTRAINED":       "",
	// "DSINTERVAL_UNCONSTRAINED":       "",
	// "BFILE":                          "",
	// "BLOB":                           "",
	// "CLOB":                           "",
	// "NCLOB":                          "",
	// "MLSLABEL":                       "",
	// "XMLTYPE":                        "",
	// "LONG":                           "",
	// "INTERVAL YEAR TO MONTH":         "",
	// "INTERVAL DAY TO SECOND":         "",
	"TIMESTAMP WITH TIME ZONE":       "datetime",
	"TIMESTAMP WITH LOCAL TIME ZONE": "datetime",
}

var HiveTypeMap = map[string]string{
	"tinyint":   "integer",
	"smallint":  "integer",
	"int":       "integer",
	"bigint":    "integer",
	"boolean":   "boolean",
	"float":     "numeric",
	"real":      "numeric",
	"double":    "numeric",
	"date":      "date",
	"datetime":  "datetime",
	"timestamp": "datetime",
	// | KW_TIMESTAMPLOCALTZ
	// | KW_TIMESTAMPTZ
	// | KW_TIMESTAMP KW_WITH KW_TIME KW_ZONE
	// | KW_INTERVAL KW_YEAR KW_TO KW_MONTH
	// | KW_INTERVAL KW_DAY KW_TO KW_SECOND
	// | KW_BINARY
	"string":  "string",
	"varchar": "string",
	"char":    "string",
	"decimal": "numeric",
}

var SqliteTypeMap = map[string]string{
	"INT":              "integer",
	"INTEGER":          "integer",
	"TINYINT":          "integer",
	"SMALLINT":         "integer",
	"MEDIUMINT":        "integer",
	"BIGINT":           "integer",
	"UNSIGNEDBIGINT":   "integer",
	"INT2":             "integer",
	"INT8":             "integer",
	"CHARACTER":        "string",
	"VARCHAR":          "string",
	"VARYINGCHARACTER": "string",
	"NCHAR":            "string",
	"NATIVECHARACTER":  "string",
	"NVARCHAR":         "string",
	"TEXT":             "string",
	"CLOB":             "string",
	// "BLOB":             "",
	"REAL":            "decimal",
	"DOUBLE":          "decimal",
	"DOUBLEPRECISION": "decimal",
	"FLOAT":           "decimal",
	"NUMERIC":         "decimal",
	"DECIMAL":         "decimal",
	"BOOLEAN":         "integer",
	"DATE":            "string",
	"DATETIME":        "string",
}

var TSqlTypeMap = map[string]string{
	"tinyint":    "integer",
	"int":        "integer",
	"bigint":     "integer",
	"smallint":   "integer",
	"bit":        "integer",
	"decimal":    "decimal",
	"numeric":    "decimal",
	"money":      "decimal",
	"smallmoney": "decimal",
	"float":      "decimal",
	"real":       "decimal",
	"date":       "date",
	"time":       "time",
	"datetime2":  "datetime",
	// "datetimeoffset":   "",
	"datetime": "datetime",
	"char":     "string",
	"varchar":  "string",
	"text":     "string",
	"nchar":    "string",
	"nvarchar": "string",
	"ntext":    "string",
	// "binary":           "",
	// "varbinary":        "",
	// "image":            "",
	// "cursor":           "",
	// "geography":        "",
	// "geometry":         "",
	// "hierarchyid":      "",
	// "json":             "",
	// "rowversion":       "",
	// "sql_variant":      "",
	// "table":            "",
	// "uniqueidentifier": "",
	// "xml":              "",
}