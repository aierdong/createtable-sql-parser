package visitor

import (
	"errors"
	"fmt"
	parser "github.com/aierdong/createtable-sql-parser/parser/plsql"
	"github.com/aierdong/createtable-sql-parser/types"
	"github.com/antlr4-go/antlr/v4"
	"math"
	"regexp"
	"strings"
)

// OracleVisitor is the visitor for Oracle SQL
// https://github.com/bytebase/plsql-parser

type OracleVisitor struct {
	*parser.BasePlSqlParserVisitor
	Table  *types.AntlrTable
	Column *types.AntlrColumn
	Err    error
}

func ParsePlSql(sql string) (*types.AntlrTable, error) {
	sqls := strings.Split(sql, ";")
	var table *types.AntlrTable

	for _, s := range sqls {
		s += ";"
		if len(s) > 12 && strings.ToUpper(s[:12]) == "CREATE TABLE" {
			var err error
			table, err = parseOracleTable(s)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	if table == nil {
		return nil, errors.New("not dound create table statment")
	}

	for _, s := range sqls {
		s = strings.TrimSpace(s)
		if len(s) > 17 && strings.ToUpper(s[:17]) == "COMMENT ON COLUMN" {
			col, err := parseOracleColumnComment(s)
			if err != nil {
				return nil, err
			}
			if col == nil { // no comment
				continue
			}
			for _, c := range table.Columns {
				if c.Name == col.Name {
					c.Comment = col.Comment
				}
			}
			continue
		}
		if len(s) > 16 && strings.ToUpper(s[:16]) == "COMMENT ON TABLE" {
			t, err := parseOracleTableComment(s)
			if err != nil {
				return nil, err
			}
			table.Comment = t
		}
	}
	return table, nil
}

func (v *OracleVisitor) VisitCreate_table(ctx *parser.Create_tableContext) interface{} {
	if ctx.Table_name() == nil {
		v.Err = errors.New("table name is nil")
		return nil
	} else {
		v.Table.Name = strings.Trim(ctx.Table_name().GetText(), "\"")
	}

	if ctx.Schema_name() != nil {
		v.Table.Database = strings.Trim(ctx.Schema_name().GetText(), "\"")
	}

	if ctx.Relational_table() == nil {
		v.Err = errors.New("relational table is nil")
		return nil
	}

	for _, child := range ctx.Relational_table().AllRelational_property() {
		colDef := child.Column_definition()
		if colDef == nil {
			continue
		}

		col := colDef.Accept(v)
		if v.Err != nil {
			return nil
		}
		if col == nil {
			continue
		}
		v.Table.Columns = append(v.Table.Columns, col.(*types.AntlrColumn))
	}

	if len(v.Table.Columns) == 0 {
		v.Err = errors.New("no column found")
	}

	return nil
}

func (v *OracleVisitor) VisitColumn_definition(ctx *parser.Column_definitionContext) interface{} {
	if ctx.Column_name() == nil {
		v.Err = errors.New("column name is nil")
		return nil
	}
	if ctx.Datatype() == nil {
		v.Err = errors.New("data type is nil")
		return nil
	}
	if ctx.Datatype().INTERVAL() != nil {
		v.Err = errors.New("'INTERVAL ???' type is not supported")
		return nil
	}
	if ctx.Datatype().TIME() != nil {
		v.Err = errors.New("'TIMESTAMP WITH ???' type is not supported")
		return nil
	}

	ret, err := v.parseColumnType(ctx.Datatype().GetText())
	if err != nil {
		v.Err = err
		return nil
	}

	return &types.AntlrColumn{
		Name:         strings.Trim(ctx.Column_name().GetText(), "\""),
		DataType:     ret.DataType,
		StringLength: ret.StringLength,
		Scale:        ret.Scale,
	}
}

func (v *OracleVisitor) VisitComment_on_column(ctx *parser.Comment_on_columnContext) interface{} {
	if ctx.Column_name() == nil || ctx.Quoted_string() == nil {
		return nil
	}
	arr := strings.Split(ctx.Column_name().GetText(), ".")
	v.Column = &types.AntlrColumn{
		Name:    strings.Trim(arr[len(arr)-1], "\""),
		Comment: strings.Trim(ctx.Quoted_string().GetText(), "'"),
	}
	return nil
}

func (v *OracleVisitor) VisitComment_on_table(ctx *parser.Comment_on_tableContext) interface{} {
	if ctx.Quoted_string() == nil {
		return nil
	}
	v.Table.Comment = strings.Trim(ctx.Quoted_string().GetText(), "'")
	return nil
}

// parseColumnType parses the column type definition and returns an AntlrColumn.
func (v *OracleVisitor) parseColumnType(typeStr string) (*types.AntlrColumn, error) {
	originalType, length, scale, err := v.parseTypeString(typeStr)
	if err != nil {
		return nil, err
	}

	column, err := v.mapColumnType(originalType)
	if err != nil {
		return nil, err
	}

	v.setColumnAttributes(column, originalType, length, scale)
	return column, nil
}

// parseTypeString parses the type string and extracts the type, length, and scale.
func (v *OracleVisitor) parseTypeString(typeStr string) (string, int, int, error) {
	re := regexp.MustCompile(`(?i)^(\w+)(?:\((\d+)(?:,(\d+))?\))?$`)
	matches := re.FindStringSubmatch(typeStr)
	if matches == nil {
		return "", 0, 0, fmt.Errorf("invalid type format: %s", typeStr)
	}

	originalType := strings.ToUpper(matches[1])
	length, scale := 0, 0
	if len(matches) >= 3 && matches[2] != "" {
		if _, err := fmt.Sscanf(matches[2], "%d", &length); err != nil {
			return "", 0, 0, fmt.Errorf("invalid type length: %s", matches[2])
		}
	}
	if len(matches) >= 4 && matches[3] != "" {
		if _, err := fmt.Sscanf(matches[3], "%d", &scale); err != nil {
			return "", 0, 0, fmt.Errorf("invalid type scale: %s", matches[3])
		}
	}

	return originalType, length, scale, nil
}

// mapColumnType maps the original type to a simplified type and initializes the column.
func (v *OracleVisitor) mapColumnType(originalType string) (*types.AntlrColumn, error) {
	simplifiedType, exists := types.PLSqlTypeMap[originalType]
	if !exists {
		return nil, fmt.Errorf("unsupported data type: %s", originalType)
	}

	return &types.AntlrColumn{DataType: simplifiedType}, nil
}

// setColumnAttributes sets the column attributes based on the original type, length, and scale.
func (v *OracleVisitor) setColumnAttributes(column *types.AntlrColumn, originalType string, length, scale int) {
	switch originalType {
	case "BINARY_INTEGER", "PLS_INTEGER", "NATURAL", "NATURALN", "POSITIVE",
		"POSITIVEN", "INT", "INTEGER", "SMALLINT":
		column.MaxInteger = math.MaxInt32
	case "SIGNTYPE":
		column.MaxInteger = 1
		column.MinInteger = -1
	case "BINARY_FLOAT", "REAL":
		if scale == 0 {
			column.DataType = "integer"
			column.MaxInteger = getMaxInt32(length)
		} else {
			column.MaxFloat = getMaxFloat32(length)
			column.Scale = scale
		}
	case "BINARY_DOUBLE", "FLOAT", "DOUBLE", "DOUBLE PRECISION", "DEC",
		"NUMBER", "NUMERIC", "DECIMAL":
		if scale == 0 {
			column.DataType = "integer"
			column.MaxInteger = getMaxInt64(length)
		} else {
			column.MaxFloat = getMaxFloat64(length)
			column.Scale = scale
		}
	case "CHAR", "NCHAR", "VARCAHR", "VARCHAR2", "NVARCHAR2", "CHARACTER", "STRING":
		column.StringLength = length
	}
}

func parseOracleColumnComment(sql string) (*types.AntlrColumn, error) {
	lexer := parser.NewPlSqlLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewPlSqlParser(stream)

	tree := p.Comment_on_column()
	visitor := &OracleVisitor{
		BasePlSqlParserVisitor: &parser.BasePlSqlParserVisitor{},
		Column:                 &types.AntlrColumn{},
	}
	tree.Accept(visitor)
	return visitor.Column, visitor.Err
}

func parseOracleTableComment(sql string) (string, error) {
	lexer := parser.NewPlSqlLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewPlSqlParser(stream)

	tree := p.Comment_on_table()
	visitor := &OracleVisitor{
		BasePlSqlParserVisitor: &parser.BasePlSqlParserVisitor{},
		Table:                  &types.AntlrTable{},
	}
	tree.Accept(visitor)
	return visitor.Table.Comment, visitor.Err
}

func parseOracleTable(sql string) (*types.AntlrTable, error) {
	lexer := parser.NewPlSqlLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewPlSqlParser(stream)

	tree := p.Create_table()

	visitor := &OracleVisitor{
		BasePlSqlParserVisitor: &parser.BasePlSqlParserVisitor{},
		Table: &types.AntlrTable{
			Columns: make([]*types.AntlrColumn, 0),
		},
	}
	tree.Accept(visitor)
	return visitor.Table, visitor.Err
}
