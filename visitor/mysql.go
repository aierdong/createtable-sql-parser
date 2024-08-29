package visitor

import (
	"errors"
	"fmt"
	parser "github.com/aierdong/createtable-sql-parser/parser/mysql"
	"github.com/aierdong/createtable-sql-parser/types"
	"github.com/antlr4-go/antlr/v4"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type MySQLVisitor struct {
	*parser.BaseMySQLParserVisitor
	Table *types.AntlrTable
	Err   error
}

func ParseMySql(sql string) (table *types.AntlrTable, err error) {
	defer func() {
		if r := recover(); r != nil {
			table = nil
			err = errors.New(fmt.Sprint("parse sql error: ", r))
		}
	}()

	lexer := parser.NewMySQLLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewMySQLParser(stream)
	p.BuildParseTrees = true

	tree := p.CreateStatement()

	visitor := &MySQLVisitor{
		BaseMySQLParserVisitor: &parser.BaseMySQLParserVisitor{},
		Table: &types.AntlrTable{
			Dialect: types.MySQL,
			Columns: make([]*types.AntlrColumn, 0),
		},
	}
	tree.Accept(visitor)

	if visitor.Err != nil {
		return nil, visitor.Err
	}
	return visitor.Table, nil
}

func (v *MySQLVisitor) VisitCreateStatement(ctx *parser.CreateStatementContext) interface{} {
	if ctx.CreateTable() == nil {
		v.Err = fmt.Errorf("not found a valid create table statement")
		return nil
	}
	return ctx.CreateTable().Accept(v)
}

func (v *MySQLVisitor) VisitCreateTable(ctx *parser.CreateTableContext) interface{} {
	if ctx.TableName() == nil {
		v.Err = fmt.Errorf("table name is nil")
		return nil
	}

	arr := strings.Split(ctx.TableName().GetText(), ".")
	if len(arr) == 2 {
		v.Table.Database = strings.Trim(arr[0], "`")
	}
	v.Table.Name = strings.Trim(arr[len(arr)-1], "`")

	if ctx.TableElementList() == nil {
		v.Err = fmt.Errorf("table element list is nil")
		return nil
	}
	ctx.TableElementList().Accept(v)

	if ctx.CreateTableOptions() != nil {
		ctx.CreateTableOptions().Accept(v)
	}

	return nil
}

func (v *MySQLVisitor) VisitTableElementList(ctx *parser.TableElementListContext) interface{} {
	for _, child := range ctx.AllTableElement() {
		colDef := child.ColumnDefinition()
		if colDef == nil {
			continue
		}
		if colDef.ColumnName() == nil || colDef.FieldDefinition() == nil || colDef.FieldDefinition().DataType() == nil {
			v.Err = fmt.Errorf("column definition, name, or field is nil")
			return nil
		}

		dataType := v.getDataType(colDef)

		// dataType: integer, string..., and length, scala
		column, err := v.parseColumnType(dataType)
		if err != nil {
			v.Err = err
			return nil
		}

		// column comment
		for _, att := range colDef.FieldDefinition().AllColumnAttribute() {
			if att.COMMENT_SYMBOL() != nil && att.TextLiteral() != nil {
				column.Comment = strings.Trim(att.TextLiteral().GetText(), "'")
			}
			if att.AUTO_INCREMENT_SYMBOL() != nil {
				column.AutoIncrement = true
			}
		}

		v.Table.Columns = append(v.Table.Columns, &types.AntlrColumn{
			Name:          strings.Trim(colDef.ColumnName().GetText(), "`"),
			DataType:      column.DataType,
			StringLength:  column.StringLength,
			Scale:         column.Scale,
			Comment:       column.Comment,
			AutoIncrement: column.AutoIncrement,
		})
	}
	return nil
}

func (v *MySQLVisitor) getDataType(colDef parser.IColumnDefinitionContext) string {
	dataType := colDef.FieldDefinition().DataType()
	dataTypeStr := dataType.GetText()
	if dataType.FieldOptions() != nil {
		fopt := dataType.FieldOptions().GetText()
		if strings.HasSuffix(dataTypeStr, fopt) {
			dataTypeStr = strings.TrimSuffix(dataTypeStr, fopt)
		}
	}
	return dataTypeStr
}

func (v *MySQLVisitor) VisitCreateTableOptions(ctx *parser.CreateTableOptionsContext) interface{} {
	for _, child := range ctx.AllCreateTableOption() {
		if child.COMMENT_SYMBOL() != nil && child.TextStringLiteral() != nil {
			v.Table.Comment = strings.Trim(child.TextStringLiteral().GetText(), "'")
			return nil
		}
	}
	return nil
}

func (v *MySQLVisitor) parseColumnType(dataType string) (column *types.AntlrColumn, err error) {
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

func (v *MySQLVisitor) extractColumnTypeInfo(dataType string) (originalType string, length int, scale int, err error) {
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

func (v *MySQLVisitor) mapColumnType(originalType string) (string, error) {
	if simplifiedType, exists := types.MySQLTypeMap[originalType]; exists {
		return simplifiedType, nil
	}
	return "", fmt.Errorf("unknown integer type: %s", originalType)
}

func (v *MySQLVisitor) setColumnAttributes(column *types.AntlrColumn, originalType string, length int, scale int) {
	switch originalType {
	case "char", "varchar", "string", "text", "tinytext", "mediumtext", "longtext":
		column.StringLength = If(length > 0 && length < 50, length, 50)
	case "tinyint":
		column.MaxInteger = math.MaxInt8
	case "smallint":
		column.MaxInteger = math.MaxInt16
	case "mediumint":
		column.MaxInteger = 1<<23 - 1 // 8388607
	case "int", "integer":
		column.MaxInteger = math.MaxInt32
	case "bigint":
		column.MaxInteger = math.MaxInt64
	case "float", "real":
		column.MaxFloat = getMaxFloat32(length)
		column.Scale = If(scale > 0, scale, 2)
	case "double", "decimal", "numeric":
		column.MaxFloat = getMaxFloat64(length)
		column.Scale = If(scale > 0, scale, 2)
	}
}
