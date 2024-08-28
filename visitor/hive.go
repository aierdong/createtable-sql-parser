package visitor

import (
	"errors"
	"fmt"
	parser "github.com/aierdong/createtable-sql-parser/parser/hive"
	"github.com/aierdong/createtable-sql-parser/types"
	"github.com/antlr4-go/antlr/v4"
	"regexp"
	"strconv"
	"strings"
)

type HiveVisitor struct {
	*parser.BaseHiveParserVisitor
	Table *types.AntlrTable
	Err   error
}

func ParseHiveSql(sql string) (*types.AntlrTable, error) {
	lexer := parser.NewHiveLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewHiveParser(stream)
	p.BuildParseTrees = true

	tree := p.CreateTableStatement()

	visitor := &HiveVisitor{
		BaseHiveParserVisitor: &parser.BaseHiveParserVisitor{},
		Table: &types.AntlrTable{
			Columns: make([]*types.AntlrColumn, 0),
		},
	}
	tree.Accept(visitor)

	if visitor.Err != nil {
		return nil, visitor.Err
	}
	return visitor.Table, nil
}

func (v *HiveVisitor) VisitCreateTableStatement(ctx *parser.CreateTableStatementContext) interface{} {
	if ctx.TableName() == nil ||
		ctx.ColumnNameTypeOrConstraintList() == nil ||
		len(ctx.ColumnNameTypeOrConstraintList().(*parser.ColumnNameTypeOrConstraintListContext).AllColumnNameTypeOrConstraint()) == 0 {
		v.Err = errors.New("invalid create table statement")
		return nil
	}

	dbName, tableName := v.resolveTableName(ctx.TableName().GetText())
	v.Table.Database = dbName
	v.Table.Name = tableName

	for _, child := range ctx.ColumnNameTypeOrConstraintList().(*parser.ColumnNameTypeOrConstraintListContext).AllColumnNameTypeOrConstraint() {
		colDef := child.ColumnNameTypeConstraint()
		if colDef.Id_() == nil || colDef.ColType() == nil {
			v.Err = errors.New("column definition, name, or field is nil")
			return nil
		}

		// dataType: integer, string..., and length, scala
		column, err := v.parseColumnType(colDef.ColType().GetText())
		if err != nil {
			v.Err = err
			return nil
		}

		v.Table.Columns = append(v.Table.Columns, &types.AntlrColumn{
			Name:    strings.Trim(colDef.Id_().GetText(), "`"),
			Type:    column.Type,
			Length:  column.Length,
			Scale:   column.Scale,
			Comment: column.Comment,
		})
	}
	return nil
}

func (v *HiveVisitor) resolveTableName(tableName string) (database string, table string) {
	parts := strings.Split(tableName, ".")
	table = strings.Trim(parts[len(parts)-1], "`")
	database = "default"
	if len(parts) >= 2 {
		database = strings.Trim(parts[0], "`")
	}
	return database, table
}

func (v *HiveVisitor) parseColumnType(dataType string) (column *types.AntlrColumn, err error) {
	baseType := ""

	// Regular expressions to match different data types and their lengths/scales
	re := regexp.MustCompile(`(?i)(\w+)(?:\((\d+)(?:,(\d+))?\))?`)
	matches := re.FindStringSubmatch(dataType)

	column = &types.AntlrColumn{}
	if len(matches) > 0 {
		baseType = strings.ToLower(matches[1])
		if simplifiedType, exists := types.HiveTypeMap[baseType]; exists {
			column.Type = simplifiedType
		} else {
			column.Type = ""
		}

		if len(matches) > 2 && matches[2] != "" {
			column.Length, _ = strconv.Atoi(matches[2])
		}
		if len(matches) > 3 && matches[3] != "" {
			column.Scale, _ = strconv.Atoi(matches[3])
		}

		if column.Type == "numeric" && column.Scale == 0 {
			column.Length = 2
			column.Scale = 2
		}

		if baseType == "char" {
			if column.Length == 0 {
				return nil, errors.New("char type must have a length")
			} else {
				column.FixLength = true
			}
		}
	}

	if baseType == "" {
		return nil, errors.New("unknown data type")
	}
	if column.Type == "" {
		return nil, fmt.Errorf("unknown data type: %s", baseType)
	}

	return column, nil
}
