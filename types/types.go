package types

type AntlrColumn struct {
	Name          string
	DataType      string
	StringLength  int
	MaxInteger    int64   // for integer datatype, max value
	MinInteger    int64   // for integer datatype, only oracle 'SIGNTYPE' has min value, or 'bit' for tsql
	MaxFloat      float64 // for float datatype, max value
	Scale         int
	Comment       string
	AutoIncrement bool
}

type AntlrTable struct {
	Database string
	Name     string
	Columns  []*AntlrColumn
	Comment  string
}
