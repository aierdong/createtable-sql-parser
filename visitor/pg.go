package visitor

import (
	"errors"
	"fmt"
	parser "github.com/aierdong/createtable-sql-parser/parser/pg"
	"github.com/aierdong/createtable-sql-parser/types"
	"github.com/antlr4-go/antlr/v4"
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

func ParsePgSql(sql string) (*types.AntlrTable, error) {
	sqls := strings.Split(sql, ";")
	var table *types.AntlrTable

	for _, s := range sqls {
		if len(s) > 12 && strings.ToUpper(s[:12]) == "CREATE TABLE" {
			var err error
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
			col.Name = ele.Accept(v).(string)
			continue
		}
		if ele, ok := child.(*parser.TypenameContext); ok {
			if t := ele.Accept(v); t != nil {
				col.Type = t.(*types.AntlrColumn).Type
				col.Length = t.(*types.AntlrColumn).Length
				col.Scale = t.(*types.AntlrColumn).Scale
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
	col, err := parseDataType(fullname)
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

func parseDataType(dataTypeStr string) (*types.AntlrColumn, error) {
	// 定义正则表达式
	re := regexp.MustCompile(`(?P<Type>\w+)(?:\((?P<Length>\d+)(?:,\s*(?P<Scale>\d+))?\))?`)

	// 匹配字符串
	match := re.FindStringSubmatch(dataTypeStr)
	if match == nil {
		return nil, fmt.Errorf("invalid data type format: %s", dataTypeStr)
	}

	// Extract and map type
	originalType := match[1]
	simplifiedType, exists := types.PgTypeMap[originalType]
	if !exists {
		return nil, fmt.Errorf("unsupported data type: %s", originalType)
	}

	// 提取匹配结果
	result := &types.AntlrColumn{Type: simplifiedType}
	if match[2] != "" {
		length, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, err
		}
		result.Length = length
	}
	if match[3] != "" {
		scale, err := strconv.Atoi(match[3])
		if err != nil {
			return nil, err
		}
		result.Scale = scale
	}

	return result, nil
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
			Columns: make([]*types.AntlrColumn, 0),
		},
	}
	tree.Accept(visitor)
	return visitor.Table, visitor.Err
}
