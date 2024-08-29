package visitor

import (
	"errors"
	"fmt"
	parser "github.com/aierdong/createtable-sql-parser/parser/pg"
	"github.com/aierdong/createtable-sql-parser/types"
	"github.com/antlr4-go/antlr/v4"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type PgVisitor struct {
	*parser.BasePostgreSQLParserVisitor
	Table  *types.AntlrTable
	Column *types.AntlrColumn
	Err    error
}

func ParsePgSql(sql string) (table *types.AntlrTable, err error) {
	defer func() {
		if r := recover(); r != nil {
			table = nil
			err = errors.New(fmt.Sprint("parse sql error: ", r))
		}
	}()

	sqls := strings.Split(sql, ";")

	for _, s := range sqls {
		if len(s) > 12 && strings.ToUpper(s[:12]) == "CREATE TABLE" {
			table, err = parsePgTable(s)
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
			col, err := parsePgColumnComment(s)
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
			t, err := parsePgTableComment(s)
			if err != nil {
				return nil, err
			}
			table.Comment = t
		}
	}

	return table, nil
}

func (v *PgVisitor) VisitCreatestmt(ctx *parser.CreatestmtContext) interface{} {
	if ctx.TABLE() == nil {
		v.Err = errors.New("not a create table statement")
		return nil
	}
	if err := ctx.Qualified_name(0).Accept(v); err != nil {
		return err
	}

	tbl := ctx.Opttableelementlist()
	if tbl == nil || len(tbl.GetChildren()) == 0 {
		v.Err = errors.New("table element list is nil")
		return nil
	}

	for _, child := range tbl.GetChild(0).GetChildren() {
		if ele, ok := child.(*parser.TableelementContext); ok {
			if col, ok := ele.GetChild(0).(*parser.ColumnDefContext); ok {
				if err := col.Accept(v); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (v *PgVisitor) VisitQualified_name(ctx *parser.Qualified_nameContext) interface{} {
	arr := strings.Split(ctx.GetText(), ".")
	if len(arr) == 2 {
		v.Table.Database = strings.Trim(arr[0], "`")
		v.Table.Name = strings.Trim(arr[1], "`")
	} else if len(arr) == 1 {
		v.Table.Name = strings.Trim(arr[0], "`")
	} else {
		v.Err = errors.New("table name error")
	}
	return nil
}

func (v *PgVisitor) VisitColumnDef(ctx *parser.ColumnDefContext) interface{} {
	col := &types.AntlrColumn{}
	for _, child := range ctx.GetChildren() {
		if ele, ok := child.(*parser.ColidContext); ok {
			col.Name = strings.Trim(ele.GetText(), "\"")
			continue
		}
		if ele, ok := child.(*parser.TypenameContext); ok {
			if t := ele.Accept(v); t != nil {
				tc := t.(*types.AntlrColumn)
				col.DataType = tc.DataType
				col.StringLength = tc.StringLength
				col.Scale = tc.Scale
				col.AutoIncrement = tc.AutoIncrement
			}
			continue
		}
	}

	v.Table.Columns = append(v.Table.Columns, col)
	return nil
}

func (v *PgVisitor) VisitColid(ctx *parser.ColidContext) interface{} {
	return strings.Trim(ctx.Identifier().GetText(), "\"")
}

func (v *PgVisitor) VisitTypename(ctx *parser.TypenameContext) interface{} {
	fullname := strings.ToLower(ctx.GetText())
	col, err := v.parseColumnType(fullname)
	if err != nil {
		v.Err = err
		return nil
	}
	return col
}

func (v *PgVisitor) VisitCommentstmt(ctx *parser.CommentstmtContext) interface{} {
	arr := strings.Split(ctx.Any_name().GetText(), ".")

	if ctx.COLUMN() != nil {
		v.Column = &types.AntlrColumn{
			Name:    strings.Trim(arr[len(arr)-1], "\""),
			Comment: strings.Trim(ctx.Comment_text().GetText(), "'"),
		}
	}

	if ctx.Object_type_any_name() != nil {
		if _, ok := ctx.Object_type_any_name().(*parser.Object_type_any_nameContext); ok {
			v.Table.Comment = strings.Trim(ctx.Comment_text().GetText(), "'")
		}
	}

	return nil
}

// extractColumnTypeInfo extracts the column type information using regular expressions.
func (v *PgVisitor) extractColumnTypeInfo(dataType string) (originalType string, length int, scale int, err error) {
	re := regexp.MustCompile(`(?P<DataType>\w+)(?:\((?P<StringLength>\d+)(?:,\s*(?P<Scale>\d+))?\))?`)
	matches := re.FindStringSubmatch(dataType)
	if matches == nil || len(matches) < 2 {
		return "", 0, 0, fmt.Errorf("invalid data type format: %s", dataType)
	}

	originalType = strings.ToLower(matches[1])
	if len(matches) >= 3 && matches[2] != "" {
		length, err = strconv.Atoi(matches[2])
		if err != nil {
			return "", 0, 0, err
		}
	}
	if len(matches) >= 4 && matches[3] != "" {
		scale, err = strconv.Atoi(matches[3])
		if err != nil {
			return "", 0, 0, err
		}
	}
	return originalType, length, scale, nil
}

// mapColumnType maps the original type to a simplified type.
func (v *PgVisitor) mapColumnType(originalType string) (string, error) {
	simplifiedType, exists := types.PgTypeMap[originalType]
	if !exists {
		return "", fmt.Errorf("unsupported data type: %s", originalType)
	}
	return simplifiedType, nil
}

// setColumnAttributes sets the attributes of the column based on its type.
func (v *PgVisitor) setColumnAttributes(column *types.AntlrColumn, originalType string, length int, scale int) {
	switch originalType {
	case "int2", "smallint":
		column.MaxInteger = math.MaxInt16
	case "int4", "integer":
		column.MaxInteger = math.MaxInt32
	case "int8", "bigint":
		column.MaxInteger = math.MaxInt64
	case "serial":
		column.MaxInteger = math.MaxInt32
		column.AutoIncrement = true
	case "bigserial":
		column.MaxInteger = math.MaxInt64
		column.AutoIncrement = true
	case "varchar", "char", "text":
		column.StringLength = If(length > 0 && length < 50, length, 50)
	case "numeric", "decimal", "double":
		column.MaxFloat = getMaxFloat64(length)
		column.Scale = If(scale > 0, scale, 2)
	case "real":
		column.MaxFloat = getMaxFloat32(length)
		column.Scale = If(scale > 0, scale, 2)
	}
}

// parseColumnType parses the column type definition and returns an AntlrColumn.
func (v *PgVisitor) parseColumnType(dataType string) (*types.AntlrColumn, error) {
	originalType, length, scale, err := v.extractColumnTypeInfo(dataType)
	if err != nil {
		return nil, err
	}

	simplifiedType, err := v.mapColumnType(originalType)
	if err != nil {
		return nil, err
	}

	column := &types.AntlrColumn{DataType: simplifiedType}
	v.setColumnAttributes(column, originalType, length, scale)
	return column, nil
}

func parsePgColumnComment(sql string) (*types.AntlrColumn, error) {
	lexer := parser.NewPostgreSQLLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewPostgreSQLParser(stream)
	p.BuildParseTrees = true

	tree := p.Commentstmt()
	visitor := &PgVisitor{
		BasePostgreSQLParserVisitor: &parser.BasePostgreSQLParserVisitor{},
		Column:                      &types.AntlrColumn{},
	}
	tree.Accept(visitor)
	return visitor.Column, visitor.Err
}

func parsePgTableComment(sql string) (string, error) {
	lexer := parser.NewPostgreSQLLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewPostgreSQLParser(stream)
	p.BuildParseTrees = true

	tree := p.Commentstmt()
	visitor := &PgVisitor{
		BasePostgreSQLParserVisitor: &parser.BasePostgreSQLParserVisitor{},
		Table:                       &types.AntlrTable{},
	}
	tree.Accept(visitor)
	return visitor.Table.Comment, visitor.Err
}

func parsePgTable(sql string) (*types.AntlrTable, error) {
	lexer := parser.NewPostgreSQLLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewPostgreSQLParser(stream)
	p.BuildParseTrees = true

	tree := p.Createstmt()

	visitor := &PgVisitor{
		BasePostgreSQLParserVisitor: &parser.BasePostgreSQLParserVisitor{},
		Table: &types.AntlrTable{
			Dialect: types.PostgreSQL,
			Columns: make([]*types.AntlrColumn, 0),
		},
	}
	tree.Accept(visitor)
	return visitor.Table, visitor.Err
}

func getMaxFloat64(length int) float64 {
	if length == 0 {
		length = 18
	}
	maxFloat := math.Pow(10, float64(length)) - 1
	if math.IsInf(maxFloat, 0) {
		return math.MaxFloat64
	}
	return maxFloat
}

func getMaxFloat32(length int) float64 {
	if length == 0 {
		length = 10
	}
	maxFloat := math.Pow(10, float64(length)) - 1
	if math.IsInf(maxFloat, 0) || maxFloat > math.MaxFloat32 {
		return math.MaxFloat32
	}
	return maxFloat
}

func getMaxInt64(length int) int64 {
	if length == 0 {
		length = 19
	}
	maxInt := math.Pow(10, float64(length)) - 1
	if maxInt > float64(math.MaxInt64) {
		return math.MaxInt64
	}
	return int64(maxInt)
}

func getMaxInt32(length int) int64 {
	if length == 0 {
		length = 10
	}
	maxInt := math.Pow(10, float64(length)) - 1
	if maxInt > float64(math.MaxInt32) {
		return math.MaxInt32
	}
	return int64(maxInt)
}
