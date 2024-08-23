package visitor

import (
	"errors"
	"fmt"
	parser "github.com/aierdong/createtable-sql-parser/parser/mysql"
	"github.com/aierdong/createtable-sql-parser/types"
	"github.com/antlr4-go/antlr/v4"
	"regexp"
	"strconv"
	"strings"
)

type MySQLVisitor struct {
	*parser.BaseMySQLParserVisitor
	Table *types.AntlrTable
	Err   error
}

func ParseMySql(sql string) (*types.AntlrTable, error) {
	lexer := parser.NewMySQLLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewMySQLParser(stream)
	p.BuildParseTrees = true

	tree := p.CreateStatement()

	visitor := &MySQLVisitor{
		BaseMySQLParserVisitor: &parser.BaseMySQLParserVisitor{},
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

	ctx.TableElementList().Accept(v)
	ctx.CreateTableOptions().Accept(v)

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
		simpleDataType, err := v.parseColumnType(dataType)
		if err != nil {
			v.Err = err
			return nil
		}

		// column comment
		for _, att := range colDef.FieldDefinition().AllColumnAttribute() {
			if att.COMMENT_SYMBOL() != nil && att.TextLiteral() != nil {
				simpleDataType.Comment = strings.Trim(att.TextLiteral().GetText(), "'")
			}
			if att.AUTO_INCREMENT_SYMBOL() != nil {
				simpleDataType.AutoIncrement = true
			}
		}

		v.Table.Columns = append(v.Table.Columns, &types.AntlrColumn{
			Name:          strings.Trim(colDef.ColumnName().GetText(), "`"),
			Type:          simpleDataType.Type,
			Length:        simpleDataType.Length,
			Scale:         simpleDataType.Scale,
			Comment:       simpleDataType.Comment,
			AutoIncrement: simpleDataType.AutoIncrement,
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

func (v *MySQLVisitor) parseColumnType(dataType string) (simpleDatType types.AntlrColumn, err error) {
	baseType := ""

	// Regular expressions to match different data types and their lengths/scales
	re := regexp.MustCompile(`(?i)(\w+)(?:\((\d+)(?:,(\d+))?\))?`)
	matches := re.FindStringSubmatch(dataType)

	if len(matches) > 0 {
		baseType = strings.ToLower(matches[1])
		if simplifiedType, exists := types.MySQLTypeMap[baseType]; exists {
			simpleDatType.Type = simplifiedType
		} else {
			simpleDatType.Type = ""
		}

		if len(matches) > 2 && matches[2] != "" {
			simpleDatType.Length, _ = strconv.Atoi(matches[2])
		}
		if len(matches) > 3 && matches[3] != "" {
			simpleDatType.Scale, _ = strconv.Atoi(matches[3])
		}

		if simpleDatType.Type == "numeric" && simpleDatType.Scale == 0 {
			simpleDatType.Length = 2
			simpleDatType.Scale = 2
		}
	}

	if baseType == "" {
		return types.AntlrColumn{}, errors.New("unknown data type")
	}
	if simpleDatType.Type == "" {
		return types.AntlrColumn{}, fmt.Errorf("unknown data type: %s", baseType)
	}

	return simpleDatType, nil
}
