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

func ParseSqliteSql(sql string) (*types.AntlrTable, error) {
	lexer := parser.NewSQLiteLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewSQLiteParser(stream)
	p.BuildParseTrees = true

	tree := p.Create_table_stmt()

	visitor := &SqliteVisitor{
		BaseSQLiteParserVisitor: &parser.BaseSQLiteParserVisitor{},
		Table:                   &types.AntlrTable{},
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
	v.Table.Columns = make([]*types.AntlrColumn, 0)
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
			length = 60
		}
		scale := 0
		if simplifiedType == "decimal" {
			scale = 2
		}

		v.Table.Columns = append(v.Table.Columns, &types.AntlrColumn{
			Name:   strings.Trim(col.Column_name().GetText(), "`\"[]"),
			Type:   simplifiedType,
			Length: length,
			Scale:  scale,
		})
	}

	return nil
}
