package visitor

import (
	"errors"
	"fmt"
	parser "github.com/aierdong/createtable-sql-parser/parser/sqlite"
	"github.com/aierdong/createtable-sql-parser/types"
	"github.com/antlr4-go/antlr/v4"
	"strings"
)

type SqliteVisitor struct {
	*parser.BaseSQLiteParserVisitor
	Table  *types.AntlrTable
	Column *types.AntlrColumn
	Err    error
}

func ParseSqliteSql(sql string) (table *types.AntlrTable, err error) {
	defer func() {
		if r := recover(); r != nil {
			table = nil
			err = errors.New(fmt.Sprint("parse sql error: ", r))
		}
	}()

	lexer := parser.NewSQLiteLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewSQLiteParser(stream)
	p.BuildParseTrees = true

	tree := p.Create_table_stmt()

	visitor := &SqliteVisitor{
		BaseSQLiteParserVisitor: &parser.BaseSQLiteParserVisitor{},
		Table: &types.AntlrTable{
			Dialect: types.SQLite3,
			Columns: make([]*types.AntlrColumn, 0),
		},
	}
	tree.Accept(visitor)

	if visitor.Err != nil {
		return nil, visitor.Err
	}

	return visitor.Table, nil
}

func (v *SqliteVisitor) VisitCreate_table_stmt(ctx *parser.Create_table_stmtContext) interface{} {
	if ctx.Table_name() == nil {
		v.Err = errors.New("table name is nil")
		return nil
	}
	if len(ctx.AllColumn_def()) == 0 {
		v.Err = errors.New("column def is nil")
		return nil
	}

	v.Table.Name = strings.Trim(ctx.Table_name().GetText(), "`\"[]")

	for _, col := range ctx.AllColumn_def() {
		if col.Column_name() == nil || col.Type_name() == nil {
			v.Err = errors.New("column name or type name is nil")
			return nil
		}

		originalType := strings.Trim(col.Type_name().GetText(), "`\"[]")

		simplifiedType, exists := types.SqliteTypeMap[originalType]
		if !exists {
			v.Err = fmt.Errorf("unsupported data type: %s", originalType)
			return nil
		}

		length := 0
		if simplifiedType == "string" {
			length = 50
		}
		scale := 0
		if simplifiedType == "decimal" {
			length = 31
			scale = 2
		}

		v.Table.Columns = append(v.Table.Columns, &types.AntlrColumn{
			Name:         strings.Trim(col.Column_name().GetText(), "`\"[]"),
			DataType:     simplifiedType,
			StringLength: length,
			Scale:        scale,
		})
	}

	return nil
}
