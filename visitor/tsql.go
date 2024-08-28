package visitor

import (
	"errors"
	"fmt"
	parser "github.com/aierdong/createtable-sql-parser/parser/tsql"
	"github.com/aierdong/createtable-sql-parser/types"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"strconv"
	"strings"
)

// TSQLVisitor is the visitor for T-SQL
// https://github.com/bytebase/tsql-parser

type MssqlVisitor struct {
	*parser.BaseTSqlParserVisitor
	Table  *types.AntlrTable
	Column *types.AntlrColumn
	Err    error
}

func ParseTSql(sql string) (*types.AntlrTable, error) {
	sqls := strings.Split(sql, ";")
	var table *types.AntlrTable

	for _, s := range sqls {
		s += ";"
		if len(s) > 12 && strings.ToUpper(s[:12]) == "CREATE TABLE" {
			var err error
			table, err = parseTSqlTable(s)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	if table == nil {
		return nil, errors.New("not found create table statement")
	}

	for _, s := range sqls {
		s = strings.TrimSpace(s)
		if len(s) > 5 && strings.ToUpper(s[:5]) == "EXEC " {
			tbl, col, err := parseTSqlComment(s)
			if err != nil {
				log.Fatal(err)
			}
			if tbl.Comment != "" {
				table.Comment = tbl.Comment
			}

			if col.Name != "" {
				for _, c := range table.Columns {
					if strings.ToLower(c.Name) == strings.ToLower(col.Name) {
						c.Comment = col.Comment
					}
				}
			}
		}
	}

	return table, nil
}

func (v *MssqlVisitor) VisitCreate_table(ctx *parser.Create_tableContext) interface{} {
	if ctx.Table_name() == nil {
		v.Err = errors.New("table name is nil")
		return nil
	}

	fullTableName := ctx.Table_name().GetText()
	parts := strings.Split(fullTableName, ".")
	if len(parts) > 1 {
		v.Table.Database = strings.Trim(parts[len(parts)-2], "\"")
	}
	v.Table.Name = strings.Trim(parts[len(parts)-1], "\"")

	if ctx.Column_def_table_constraints() == nil {
		v.Err = errors.New("column def table constraints is nil")
		return nil
	}

	for _, child := range ctx.Column_def_table_constraints().AllColumn_def_table_constraint() {
		colDef := child.Column_definition()
		if colDef == nil {
			continue
		}

		col := colDef.Accept(v)
		if v.Err != nil || col == nil {
			continue
		}

		v.Table.Columns = append(v.Table.Columns, col.(*types.AntlrColumn))
	}

	if len(v.Table.Columns) == 0 {
		v.Err = errors.New("no column found")
	}

	return nil
}
func (v *MssqlVisitor) VisitColumn_definition(ctx *parser.Column_definitionContext) interface{} {
	if ctx.Id_() == nil || ctx.Data_type() == nil {
		v.Err = errors.New("column name or data type is nil")
		return nil
	}

	ret := ctx.Data_type().Accept(v)
	if v.Err != nil {
		return nil
	}

	col := ret.(*types.AntlrColumn)
	return &types.AntlrColumn{
		Name:   strings.Trim(ctx.Id_().GetText(), "[]"),
		Type:   col.Type,
		Length: col.Length,
		Scale:  col.Scale,
	}
}

func (v *MssqlVisitor) VisitData_type(ctx *parser.Data_typeContext) interface{} {
	var originalType string
	if ctx.Id_() != nil {
		originalType = strings.ToLower(ctx.Id_().GetText())
		if ctx.Id_().Keyword() != nil {
			originalType = strings.ToLower(ctx.Id_().Keyword().GetText())
		}
	}

	if originalType == "" {
		v.Err = errors.New("data type is empty")
		return nil
	}

	simplifiedType, exists := types.TSqlTypeMap[originalType]
	if !exists {
		v.Err = fmt.Errorf("unsupported data type: %s", originalType)
		return nil
	}

	col := &types.AntlrColumn{
		Type: simplifiedType,
	}

	switch originalType {
	case "money", "smallmoney", "float", "real":
		col.Scale = 4
	case "bit":
		col.Length = 1
	default:
		if len(ctx.AllDECIMAL()) > 0 {
			if length, err := strconv.Atoi(ctx.DECIMAL(0).GetText()); err == nil {
				col.Length = length
			}
		}
		if len(ctx.AllDECIMAL()) > 1 {
			if scale, err := strconv.Atoi(ctx.DECIMAL(1).GetText()); err == nil {
				col.Scale = scale
			}
		}
		if originalType == "char" || originalType == "nchar" {
			if col.Length == 0 {
				v.Err = errors.New("char type must have a length")
				return nil
			} else {
				col.FixLength = true
			}
		}
	}

	return col
}

func (v *MssqlVisitor) VisitExecute_statement(ctx *parser.Execute_statementContext) interface{} {
	body := ctx.Execute_body()
	if !v.isValidBody(body) {
		return nil
	}

	procName := v.getProcName(body)
	if procName != "sp_addextendedproperty" && procName != "sp_updateextendedproperty" {
		return nil
	}

	args := body.Execute_statement_arg()
	if args == nil {
		return nil
	}

	m := v.parseArgs(args)
	if m == nil {
		return nil
	}

	if strings.ToLower(m["level2type"]) == "column" {
		v.Column = &types.AntlrColumn{
			Name:    strings.Trim(m["level2name"], "'"),
			Comment: m["value"],
		}
		return nil
	}

	if strings.ToLower(m["level1type"]) == "table" {
		v.Table.Comment = m["value"]
		return nil
	}

	return nil
}

// Check if the body and its nested structures are valid
func (v *MssqlVisitor) isValidBody(body parser.IExecute_bodyContext) bool {
	return body != nil &&
		body.Func_proc_name_server_database_schema() != nil &&
		body.Func_proc_name_server_database_schema().Func_proc_name_database_schema() != nil
}

// Retrieve the procedure name from the body
func (v *MssqlVisitor) getProcName(body parser.IExecute_bodyContext) string {
	proc := body.Func_proc_name_server_database_schema().Func_proc_name_database_schema().AllId_()
	return proc[len(proc)-1].GetText()
}

// Parse the arguments and return them as a map
func (v *MssqlVisitor) parseArgs(args parser.IExecute_statement_argContext) map[string]string {
	m := make(map[string]string)
	if args.AllExecute_statement_arg_named() == nil || len(args.AllExecute_statement_arg_named()) == 0 {
		stmt := args.Execute_statement_arg_unnamed()
		if stmt == nil || strings.ToLower(strings.Trim(stmt.GetValue().GetText(), "N'")) != "ms_description" {
			return nil
		}
		allArgs := strings.Split(args.Execute_statement_arg(0).GetText(), ",")
		for i, arg := range allArgs {
			m[getTSqlNonamedArgName(i)] = strings.Trim(arg, "N'")
		}
	} else {
		for _, arg := range args.AllExecute_statement_arg_named() {
			key := strings.Trim(arg.GetName().GetText(), "@")
			value := strings.Trim(arg.GetValue().GetText(), "N'")
			m[key] = value
			if strings.ToLower(key) == "name" && strings.ToLower(value) != "ms_description" {
				return nil
			}
		}
	}
	return m
}

func parseTSqlComment(sql string) (*types.AntlrTable, *types.AntlrColumn, error) {
	lexer := parser.NewTSqlLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewTSqlParser(stream)

	tree := p.Execute_statement()
	visitor := &MssqlVisitor{
		BaseTSqlParserVisitor: &parser.BaseTSqlParserVisitor{},
		Table:                 &types.AntlrTable{},
		Column:                &types.AntlrColumn{},
	}
	tree.Accept(visitor)
	return visitor.Table, visitor.Column, visitor.Err
}

func parseTSqlTable(sql string) (*types.AntlrTable, error) {
	lexer := parser.NewTSqlLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewTSqlParser(stream)

	tree := p.Create_table()

	visitor := &MssqlVisitor{
		BaseTSqlParserVisitor: &parser.BaseTSqlParserVisitor{},
		Table: &types.AntlrTable{
			Columns: make([]*types.AntlrColumn, 0),
		},
	}
	tree.Accept(visitor)
	return visitor.Table, visitor.Err
}

func getTSqlNonamedArgName(i int) string {
	switch i {
	case 0:
		return "value"
	case 1:
		return "level0type"
	case 2:
		return "level0name"
	case 3:
		return "level1type"
	case 4:
		return "level1name"
	case 5:
		return "level2type"
	case 6:
		return "level2name"
	default:
		return ""
	}
}
