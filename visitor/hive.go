package visitor

import (
	"errors"
	"fmt"
	parser "github.com/aierdong/createtable-sql-parser/parser/hive"
	"github.com/aierdong/createtable-sql-parser/types"
	"github.com/antlr4-go/antlr/v4"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type HiveVisitor struct {
	*parser.BaseHiveParserVisitor
	Table *types.AntlrTable
	Err   error
}

func ParseHiveSql(sql string) (table *types.AntlrTable, err error) {
	defer func() {
		if r := recover(); r != nil {
			table = nil
			err = errors.New(fmt.Sprint("parse sql error: ", r))
		}
	}()

	lexer := parser.NewHiveLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewHiveParser(stream)
	p.BuildParseTrees = true

	tree := p.CreateTableStatement()

	visitor := &HiveVisitor{
		BaseHiveParserVisitor: &parser.BaseHiveParserVisitor{},
		Table: &types.AntlrTable{
			Dialect: types.Hive,
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
		column.Name = strings.Trim(colDef.Id_().GetText(), "`")

		v.Table.Columns = append(v.Table.Columns, column)
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

// parseColumnType parses the column type definition and returns an AntlrColumn.
func (v *HiveVisitor) parseColumnType(dataType string) (column *types.AntlrColumn, err error) {
	originalType, length, scale, err := v.extractColumnTypeInfo(dataType)
	if err != nil {
		return nil, err
	}

	column = &types.AntlrColumn{}
	column.DataType, err = v.mapColumnType(originalType)
	if err != nil {
		return nil, err
	}

	v.setColumnAttributes(column, originalType, length, scale)
	return column, nil
}

// extractColumnTypeInfo extracts the column type information using regular expressions.
func (v *HiveVisitor) extractColumnTypeInfo(dataType string) (originalType string, length int, scale int, err error) {
	re := regexp.MustCompile(`(?i)(\w+)(?:\((\d+)(?:,(\d+))?\))?`)
	matches := re.FindStringSubmatch(dataType)
	if len(matches) == 0 {
		return "", 0, 0, errors.New("unknown column define: " + dataType)
	}

	originalType = strings.ToLower(matches[1])
	if len(matches) > 2 && matches[2] != "" {
		length, _ = strconv.Atoi(matches[2])
	}
	if len(matches) > 3 && matches[3] != "" {
		scale, _ = strconv.Atoi(matches[3])
	}
	return originalType, length, scale, nil
}

// mapColumnType maps the original type to a simplified type.
func (v *HiveVisitor) mapColumnType(originalType string) (string, error) {
	if simplifiedType, exists := types.HiveTypeMap[originalType]; exists {
		return simplifiedType, nil
	}
	return "", fmt.Errorf("unknown integer type: %s", originalType)
}

// setColumnAttributes sets the attributes of the column based on its type.
func (v *HiveVisitor) setColumnAttributes(column *types.AntlrColumn, originalType string, length int, scale int) {
	switch originalType {
	case "char", "varchar", "string":
		column.StringLength = If(length > 0 && length < 50, length, 50)
	case "tinyint":
		column.MaxInteger = math.MaxInt8
	case "smallint":
		column.MaxInteger = math.MaxInt16
	case "int", "integer":
		column.MaxInteger = math.MaxInt32
	case "bigint":
		column.MaxInteger = math.MaxInt64
	case "float":
		column.MaxFloat = getMaxFloat32(length)
		column.Scale = If(scale > 0, scale, 2)
	case "double", "decimal":
		column.MaxFloat = getMaxFloat64(length)
		column.Scale = If(scale > 0, scale, 2)
	}
}

func If[T any](express bool, a, b T) T {
	if express {
		return a
	}
	return b
}
