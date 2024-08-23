// Code generated from E:/learning/sqlparser/g4/HiveParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // HiveParser

import "github.com/antlr4-go/antlr/v4"

type BaseHiveParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseHiveParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExplainStatement(ctx *ExplainStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExplainOption(ctx *ExplainOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitVectorizationOnly(ctx *VectorizationOnlyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitVectorizatonDetail(ctx *VectorizatonDetailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExecStatement(ctx *ExecStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitLoadStatement(ctx *LoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitReplicationClause(ctx *ReplicationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitReplDumpStatement(ctx *ReplDumpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitReplDbPolicy(ctx *ReplDbPolicyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitReplLoadStatement(ctx *ReplLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitReplConfigs(ctx *ReplConfigsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitReplConfigsList(ctx *ReplConfigsListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitReplTableLevelPolicy(ctx *ReplTableLevelPolicyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitReplStatusStatement(ctx *ReplStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDdlStatement(ctx *DdlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitIfExists(ctx *IfExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRestrictOrCascade(ctx *RestrictOrCascadeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitForce(ctx *ForceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRewriteEnabled(ctx *RewriteEnabledContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRewriteDisabled(ctx *RewriteDisabledContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitStoredAsDirs(ctx *StoredAsDirsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitOrReplace(ctx *OrReplaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDbLocation(ctx *DbLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDbManagedLocation(ctx *DbManagedLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDbProperties(ctx *DbPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDbPropertiesList(ctx *DbPropertiesListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDbConnectorName(ctx *DbConnectorNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSwitchDatabaseStatement(ctx *SwitchDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropDatabaseStatement(ctx *DropDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDatabaseComment(ctx *DatabaseCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTruncateTableStatement(ctx *TruncateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropTableStatement(ctx *DropTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitInputFileFormat(ctx *InputFileFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTabTypeExpr(ctx *TabTypeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartTypeExpr(ctx *PartTypeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTabPartColTypeExpr(ctx *TabPartColTypeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDescStatement(ctx *DescStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAnalyzeStatement(ctx *AnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFrom_in(ctx *From_inContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDb_schema(ctx *Db_schemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitShowStatement(ctx *ShowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitShowTablesFilterExpr(ctx *ShowTablesFilterExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitLockStatement(ctx *LockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitLockDatabase(ctx *LockDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitLockMode(ctx *LockModeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitUnlockStatement(ctx *UnlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitUnlockDatabase(ctx *UnlockDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropRoleStatement(ctx *DropRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitGrantPrivileges(ctx *GrantPrivilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRevokePrivileges(ctx *RevokePrivilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitGrantRole(ctx *GrantRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRevokeRole(ctx *RevokeRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitShowRoleGrants(ctx *ShowRoleGrantsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitShowRoles(ctx *ShowRolesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitShowCurrentRole(ctx *ShowCurrentRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSetRole(ctx *SetRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitShowGrants(ctx *ShowGrantsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitShowRolePrincipals(ctx *ShowRolePrincipalsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrivilegeIncludeColObject(ctx *PrivilegeIncludeColObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrivilegeObject(ctx *PrivilegeObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrivObject(ctx *PrivObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrivObjectCols(ctx *PrivObjectColsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrivilegeList(ctx *PrivilegeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrivlegeDef(ctx *PrivlegeDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrincipalSpecification(ctx *PrincipalSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrincipalName(ctx *PrincipalNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWithGrantOption(ctx *WithGrantOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitGrantOptionFor(ctx *GrantOptionForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAdminOptionFor(ctx *AdminOptionForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWithAdminOption(ctx *WithAdminOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitMetastoreCheck(ctx *MetastoreCheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitResourceList(ctx *ResourceListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitResource(ctx *ResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitResourceType(ctx *ResourceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropFunctionStatement(ctx *DropFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitReloadFunctionsStatement(ctx *ReloadFunctionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateMacroStatement(ctx *CreateMacroStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropMacroStatement(ctx *DropMacroStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateViewStatement(ctx *CreateViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitViewPartition(ctx *ViewPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitViewOrganization(ctx *ViewOrganizationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitViewClusterSpec(ctx *ViewClusterSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitViewComplexSpec(ctx *ViewComplexSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitViewDistSpec(ctx *ViewDistSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitViewSortSpec(ctx *ViewSortSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropViewStatement(ctx *DropViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateScheduledQueryStatement(ctx *CreateScheduledQueryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropScheduledQueryStatement(ctx *DropScheduledQueryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterScheduledQueryStatement(ctx *AlterScheduledQueryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterScheduledQueryChange(ctx *AlterScheduledQueryChangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitScheduleSpec(ctx *ScheduleSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExecutedAsSpec(ctx *ExecutedAsSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDefinedAsSpec(ctx *DefinedAsSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitShowFunctionIdentifier(ctx *ShowFunctionIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitShowStmtIdentifier(ctx *ShowStmtIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableComment(ctx *TableCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateTablePartitionSpec(ctx *CreateTablePartitionSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateTablePartitionColumnTypeSpec(ctx *CreateTablePartitionColumnTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateTablePartitionColumnSpec(ctx *CreateTablePartitionColumnSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitionTransformSpec(ctx *PartitionTransformSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameTransformConstraint(ctx *ColumnNameTransformConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitionTransformType(ctx *PartitionTransformTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableBuckets(ctx *TableBucketsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableImplBuckets(ctx *TableImplBucketsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableSkewed(ctx *TableSkewedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRowFormat(ctx *RowFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRecordReader(ctx *RecordReaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRecordWriter(ctx *RecordWriterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRowFormatSerde(ctx *RowFormatSerdeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRowFormatDelimited(ctx *RowFormatDelimitedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableRowFormat(ctx *TableRowFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTablePropertiesPrefixed(ctx *TablePropertiesPrefixedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableProperties(ctx *TablePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTablePropertiesList(ctx *TablePropertiesListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitKeyValueProperty(ctx *KeyValuePropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitKeyProperty(ctx *KeyPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableRowFormatFieldIdentifier(ctx *TableRowFormatFieldIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableRowFormatCollItemsIdentifier(ctx *TableRowFormatCollItemsIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableRowFormatMapKeysIdentifier(ctx *TableRowFormatMapKeysIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableRowFormatLinesIdentifier(ctx *TableRowFormatLinesIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableRowNullFormat(ctx *TableRowNullFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableFileFormat(ctx *TableFileFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableLocation(ctx *TableLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameTypeList(ctx *ColumnNameTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameTypeOrConstraintList(ctx *ColumnNameTypeOrConstraintListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameColonTypeList(ctx *ColumnNameColonTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameList(ctx *ColumnNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExtColumnName(ctx *ExtColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameOrderList(ctx *ColumnNameOrderListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnParenthesesList(ctx *ColumnParenthesesListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitEnableValidateSpecification(ctx *EnableValidateSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitEnableSpecification(ctx *EnableSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitValidateSpecification(ctx *ValidateSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitEnforcedSpecification(ctx *EnforcedSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRelySpecification(ctx *RelySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateConstraint(ctx *CreateConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterConstraintWithName(ctx *AlterConstraintWithNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableLevelConstraint(ctx *TableLevelConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPkUkConstraint(ctx *PkUkConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCheckConstraint(ctx *CheckConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateForeignKey(ctx *CreateForeignKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterForeignKeyWithName(ctx *AlterForeignKeyWithNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSkewedValueElement(ctx *SkewedValueElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSkewedColumnValuePairList(ctx *SkewedColumnValuePairListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSkewedColumnValuePair(ctx *SkewedColumnValuePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSkewedColumnValues(ctx *SkewedColumnValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSkewedColumnValue(ctx *SkewedColumnValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSkewedValueLocationElement(ctx *SkewedValueLocationElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitOrderSpecification(ctx *OrderSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitNullOrdering(ctx *NullOrderingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameOrder(ctx *ColumnNameOrderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameCommentList(ctx *ColumnNameCommentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameComment(ctx *ColumnNameCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitOrderSpecificationRewrite(ctx *OrderSpecificationRewriteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnRefOrder(ctx *ColumnRefOrderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameType(ctx *ColumnNameTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameTypeOrConstraint(ctx *ColumnNameTypeOrConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableConstraint(ctx *TableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameTypeConstraint(ctx *ColumnNameTypeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnConstraint(ctx *ColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitForeignKeyConstraint(ctx *ForeignKeyConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColConstraint(ctx *ColConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterColumnConstraint(ctx *AlterColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterForeignKeyConstraint(ctx *AlterForeignKeyConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterColConstraint(ctx *AlterColConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnConstraintType(ctx *ColumnConstraintTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDefaultVal(ctx *DefaultValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableConstraintType(ctx *TableConstraintTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitConstraintOptsCreate(ctx *ConstraintOptsCreateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitConstraintOptsAlter(ctx *ConstraintOptsAlterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnNameColonType(ctx *ColumnNameColonTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColType(ctx *ColTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColTypeList(ctx *ColTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitListType(ctx *ListTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitUnionType(ctx *UnionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSetOperator(ctx *SetOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitQueryStatementExpression(ctx *QueryStatementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitQueryStatementExpressionBody(ctx *QueryStatementExpressionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCteStatement(ctx *CteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFromStatement(ctx *FromStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSingleFromStatement(ctx *SingleFromStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRegularBody(ctx *RegularBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAtomSelectStatement(ctx *AtomSelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSelectStatement(ctx *SelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSetOpSelectStatement(ctx *SetOpSelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSelectStatementWithCTE(ctx *SelectStatementWithCTEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitInsertClause(ctx *InsertClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDestination(ctx *DestinationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnAssignmentClause(ctx *ColumnAssignmentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedencePlusExpressionOrDefault(ctx *PrecedencePlusExpressionOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSetColumnsClause(ctx *SetColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSqlTransactionStatement(ctx *SqlTransactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitStartTransactionStatement(ctx *StartTransactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTransactionMode(ctx *TransactionModeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTransactionAccessMode(ctx *TransactionAccessModeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitIsolationLevel(ctx *IsolationLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitLevelOfIsolation(ctx *LevelOfIsolationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCommitStatement(ctx *CommitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRollbackStatement(ctx *RollbackStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSetAutoCommitStatement(ctx *SetAutoCommitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAbortTransactionStatement(ctx *AbortTransactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAbortCompactionStatement(ctx *AbortCompactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitMergeStatement(ctx *MergeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWhenClauses(ctx *WhenClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWhenNotMatchedClause(ctx *WhenNotMatchedClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWhenMatchedAndClause(ctx *WhenMatchedAndClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWhenMatchedThenClause(ctx *WhenMatchedThenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitUpdateOrDelete(ctx *UpdateOrDeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitKillQueryStatement(ctx *KillQueryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCompactionId(ctx *CompactionIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCompactionPool(ctx *CompactionPoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCompactionType(ctx *CompactionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCompactionStatus(ctx *CompactionStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatement(ctx *AlterStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterTableStatementSuffix(ctx *AlterTableStatementSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterTblPartitionStatementSuffix(ctx *AlterTblPartitionStatementSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementPartitionKeyType(ctx *AlterStatementPartitionKeyTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterViewStatementSuffix(ctx *AlterViewStatementSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterMaterializedViewStatementSuffix(ctx *AlterMaterializedViewStatementSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterMaterializedViewSuffixRewrite(ctx *AlterMaterializedViewSuffixRewriteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterMaterializedViewSuffixRebuild(ctx *AlterMaterializedViewSuffixRebuildContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterDatabaseStatementSuffix(ctx *AlterDatabaseStatementSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterDatabaseSuffixProperties(ctx *AlterDatabaseSuffixPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterDatabaseSuffixSetOwner(ctx *AlterDatabaseSuffixSetOwnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterDatabaseSuffixSetLocation(ctx *AlterDatabaseSuffixSetLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterDatabaseSuffixSetManagedLocation(ctx *AlterDatabaseSuffixSetManagedLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixRename(ctx *AlterStatementSuffixRenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixAddCol(ctx *AlterStatementSuffixAddColContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixAddConstraint(ctx *AlterStatementSuffixAddConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixUpdateColumns(ctx *AlterStatementSuffixUpdateColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixDropConstraint(ctx *AlterStatementSuffixDropConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixRenameCol(ctx *AlterStatementSuffixRenameColContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixUpdateStatsCol(ctx *AlterStatementSuffixUpdateStatsColContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixUpdateStats(ctx *AlterStatementSuffixUpdateStatsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementChangeColPosition(ctx *AlterStatementChangeColPositionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixAddPartitions(ctx *AlterStatementSuffixAddPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixAddPartitionsElement(ctx *AlterStatementSuffixAddPartitionsElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixTouch(ctx *AlterStatementSuffixTouchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixArchive(ctx *AlterStatementSuffixArchiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixUnArchive(ctx *AlterStatementSuffixUnArchiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitionLocation(ctx *PartitionLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixDropPartitions(ctx *AlterStatementSuffixDropPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixProperties(ctx *AlterStatementSuffixPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterViewSuffixProperties(ctx *AlterViewSuffixPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixSerdeProperties(ctx *AlterStatementSuffixSerdePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTablePartitionPrefix(ctx *TablePartitionPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixFileFormat(ctx *AlterStatementSuffixFileFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixClusterbySortby(ctx *AlterStatementSuffixClusterbySortbyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterTblPartitionStatementSuffixSkewedLocation(ctx *AlterTblPartitionStatementSuffixSkewedLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSkewedLocations(ctx *SkewedLocationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSkewedLocationsList(ctx *SkewedLocationsListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSkewedLocationMap(ctx *SkewedLocationMapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixLocation(ctx *AlterStatementSuffixLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixSkewedby(ctx *AlterStatementSuffixSkewedbyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixExchangePartition(ctx *AlterStatementSuffixExchangePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixRenamePart(ctx *AlterStatementSuffixRenamePartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixStatsPart(ctx *AlterStatementSuffixStatsPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixMergeFiles(ctx *AlterStatementSuffixMergeFilesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixBucketNum(ctx *AlterStatementSuffixBucketNumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitBlocking(ctx *BlockingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCompactPool(ctx *CompactPoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixCompact(ctx *AlterStatementSuffixCompactContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixSetOwner(ctx *AlterStatementSuffixSetOwnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixSetPartSpec(ctx *AlterStatementSuffixSetPartSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterStatementSuffixExecute(ctx *AlterStatementSuffixExecuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFileFormat(ctx *FileFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterDataConnectorStatementSuffix(ctx *AlterDataConnectorStatementSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterDataConnectorSuffixProperties(ctx *AlterDataConnectorSuffixPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterDataConnectorSuffixSetOwner(ctx *AlterDataConnectorSuffixSetOwnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterDataConnectorSuffixSetUrl(ctx *AlterDataConnectorSuffixSetUrlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitLikeTableOrFile(ctx *LikeTableOrFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateDataConnectorStatement(ctx *CreateDataConnectorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDataConnectorComment(ctx *DataConnectorCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDataConnectorUrl(ctx *DataConnectorUrlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDataConnectorType(ctx *DataConnectorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDcProperties(ctx *DcPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropDataConnectorStatement(ctx *DropDataConnectorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableAllColumns(ctx *TableAllColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableOrColumn(ctx *TableOrColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDefaultValue(ctx *DefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAliasList(ctx *AliasListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFromSource(ctx *FromSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAtomjoinSource(ctx *AtomjoinSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitJoinSource(ctx *JoinSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitJoinSourcePart(ctx *JoinSourcePartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitUniqueJoinSource(ctx *UniqueJoinSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitUniqueJoinExpr(ctx *UniqueJoinExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitUniqueJoinToken(ctx *UniqueJoinTokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitJoinToken(ctx *JoinTokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitLateralView(ctx *LateralViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableAlias(ctx *TableAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableBucketSample(ctx *TableBucketSampleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSplitSample(ctx *SplitSampleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableSample(ctx *TableSampleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableSource(ctx *TableSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAsOfClause(ctx *AsOfClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitUniqueJoinTableSource(ctx *UniqueJoinTableSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitViewName(ctx *ViewNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSubQuerySource(ctx *SubQuerySourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitioningSpec(ctx *PartitioningSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitionTableFunctionSource(ctx *PartitionTableFunctionSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitionedTableFunction(ctx *PartitionedTableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSearchCondition(ctx *SearchConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitValuesSource(ctx *ValuesSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitValuesClause(ctx *ValuesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitValuesTableConstructor(ctx *ValuesTableConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitValueRowConstructor(ctx *ValueRowConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFirstValueRowConstructor(ctx *FirstValueRowConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitVirtualTableSource(ctx *VirtualTableSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSelectClause(ctx *SelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAll_distinct(ctx *All_distinctContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSelectList(ctx *SelectListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSelectTrfmClause(ctx *SelectTrfmClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSelectItem(ctx *SelectItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTrfmClause(ctx *TrfmClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSelectExpression(ctx *SelectExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSelectExpressionList(ctx *SelectExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWindow_clause(ctx *Window_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWindow_defn(ctx *Window_defnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWindow_specification(ctx *Window_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWindow_frame(ctx *Window_frameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWindow_range_expression(ctx *Window_range_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWindow_value_expression(ctx *Window_value_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWindow_frame_start_boundary(ctx *Window_frame_start_boundaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWindow_frame_boundary(ctx *Window_frame_boundaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitGroupby_expression(ctx *Groupby_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitGroupByEmpty(ctx *GroupByEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRollupStandard(ctx *RollupStandardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRollupOldSyntax(ctx *RollupOldSyntaxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitGroupingSetExpression(ctx *GroupingSetExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitGroupingSetExpressionMultiple(ctx *GroupingSetExpressionMultipleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitGroupingExpressionSingle(ctx *GroupingExpressionSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitQualifyClause(ctx *QualifyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitHavingCondition(ctx *HavingConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExpressionsInParenthesis(ctx *ExpressionsInParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExpressionsNotInParenthesis(ctx *ExpressionsNotInParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExpressionPart(ctx *ExpressionPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFirstExpressionsWithAlias(ctx *FirstExpressionsWithAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExpressionWithAlias(ctx *ExpressionWithAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnRefOrderInParenthesis(ctx *ColumnRefOrderInParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitColumnRefOrderNotInParenthesis(ctx *ColumnRefOrderNotInParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitClusterByClause(ctx *ClusterByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitionByClause(ctx *PartitionByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDistributeByClause(ctx *DistributeByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSortByClause(ctx *SortByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTrimFunction(ctx *TrimFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFunction_(ctx *Function_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitNull_treatment(ctx *Null_treatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCastExpression(ctx *CastExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCaseExpression(ctx *CaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWhenExpression(ctx *WhenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFloorExpression(ctx *FloorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFloorDateQualifiers(ctx *FloorDateQualifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExtractExpression(ctx *ExtractExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTimeQualifiers(ctx *TimeQualifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrepareStmtParam(ctx *PrepareStmtParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitParameterIdx(ctx *ParameterIdxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitStringLiteralSequence(ctx *StringLiteralSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCharSetStringLiteral(ctx *CharSetStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDateLiteral(ctx *DateLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTimestampLiteral(ctx *TimestampLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTimestampLocalTZLiteral(ctx *TimestampLocalTZLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitIntervalValue(ctx *IntervalValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitIntervalExpression(ctx *IntervalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitIntervalQualifiers(ctx *IntervalQualifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAtomExpression(ctx *AtomExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceFieldExpression(ctx *PrecedenceFieldExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceUnaryOperator(ctx *PrecedenceUnaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceUnaryPrefixExpression(ctx *PrecedenceUnaryPrefixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceBitwiseXorOperator(ctx *PrecedenceBitwiseXorOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceBitwiseXorExpression(ctx *PrecedenceBitwiseXorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceStarOperator(ctx *PrecedenceStarOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceStarExpression(ctx *PrecedenceStarExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedencePlusOperator(ctx *PrecedencePlusOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedencePlusExpression(ctx *PrecedencePlusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceConcatenateOperator(ctx *PrecedenceConcatenateOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceConcatenateExpression(ctx *PrecedenceConcatenateExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceAmpersandOperator(ctx *PrecedenceAmpersandOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceAmpersandExpression(ctx *PrecedenceAmpersandExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceBitwiseOrOperator(ctx *PrecedenceBitwiseOrOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceBitwiseOrExpression(ctx *PrecedenceBitwiseOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceRegexpOperator(ctx *PrecedenceRegexpOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceSimilarOperator(ctx *PrecedenceSimilarOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSubQueryExpression(ctx *SubQueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceSimilarExpression(ctx *PrecedenceSimilarExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceSimilarExpressionMain(ctx *PrecedenceSimilarExpressionMainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceSimilarExpressionPart(ctx *PrecedenceSimilarExpressionPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceSimilarExpressionAtom(ctx *PrecedenceSimilarExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceSimilarExpressionQuantifierPredicate(ctx *PrecedenceSimilarExpressionQuantifierPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitQuantifierType(ctx *QuantifierTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceSimilarExpressionIn(ctx *PrecedenceSimilarExpressionInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceSimilarExpressionPartNot(ctx *PrecedenceSimilarExpressionPartNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceDistinctOperator(ctx *PrecedenceDistinctOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceEqualOperator(ctx *PrecedenceEqualOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceEqualExpression(ctx *PrecedenceEqualExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitIsCondition(ctx *IsConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceUnarySuffixExpression(ctx *PrecedenceUnarySuffixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceNotOperator(ctx *PrecedenceNotOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceNotExpression(ctx *PrecedenceNotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceAndOperator(ctx *PrecedenceAndOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceAndExpression(ctx *PrecedenceAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceOrOperator(ctx *PrecedenceOrOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrecedenceOrExpression(ctx *PrecedenceOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitBooleanValueTok(ctx *BooleanValueTokContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTableOrPartition(ctx *TableOrPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitionSpec(ctx *PartitionSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitionVal(ctx *PartitionValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitionSelectorSpec(ctx *PartitionSelectorSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitionSelectorVal(ctx *PartitionSelectorValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPartitionSelectorOperator(ctx *PartitionSelectorOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSubQuerySelectorOperator(ctx *SubQuerySelectorOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSysFuncNames(ctx *SysFuncNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDescFuncNames(ctx *DescFuncNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitId_(ctx *Id_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitFunctionIdentifier(ctx *FunctionIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrincipalIdentifier(ctx *PrincipalIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitNonReserved(ctx *NonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitSql11ReservedKeywordsUsedAsFunctionName(ctx *Sql11ReservedKeywordsUsedAsFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitHint(ctx *HintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitHintList(ctx *HintListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitHintItem(ctx *HintItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitHintName(ctx *HintNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitHintArgs(ctx *HintArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitHintArgName(ctx *HintArgNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPrepareStatement(ctx *PrepareStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExecuteStatement(ctx *ExecuteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitExecuteParamList(ctx *ExecuteParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitResourcePlanDdlStatements(ctx *ResourcePlanDdlStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRpAssign(ctx *RpAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRpAssignList(ctx *RpAssignListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRpUnassign(ctx *RpUnassignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitRpUnassignList(ctx *RpUnassignListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateResourcePlanStatement(ctx *CreateResourcePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitWithReplace(ctx *WithReplaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitActivate(ctx *ActivateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitEnable(ctx *EnableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDisable(ctx *DisableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitUnmanaged(ctx *UnmanagedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterResourcePlanStatement(ctx *AlterResourcePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitGlobalWmStatement(ctx *GlobalWmStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitReplaceResourcePlanStatement(ctx *ReplaceResourcePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropResourcePlanStatement(ctx *DropResourcePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPoolPath(ctx *PoolPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTriggerExpression(ctx *TriggerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTriggerExpressionStandalone(ctx *TriggerExpressionStandaloneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTriggerOrExpression(ctx *TriggerOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTriggerAndExpression(ctx *TriggerAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTriggerAtomExpression(ctx *TriggerAtomExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTriggerLiteral(ctx *TriggerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitComparisionOperator(ctx *ComparisionOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTriggerActionExpression(ctx *TriggerActionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitTriggerActionExpressionStandalone(ctx *TriggerActionExpressionStandaloneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateTriggerStatement(ctx *CreateTriggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterTriggerStatement(ctx *AlterTriggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropTriggerStatement(ctx *DropTriggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPoolAssign(ctx *PoolAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitPoolAssignList(ctx *PoolAssignListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreatePoolStatement(ctx *CreatePoolStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterPoolStatement(ctx *AlterPoolStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropPoolStatement(ctx *DropPoolStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitCreateMappingStatement(ctx *CreateMappingStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitAlterMappingStatement(ctx *AlterMappingStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHiveParserVisitor) VisitDropMappingStatement(ctx *DropMappingStatementContext) interface{} {
	return v.VisitChildren(ctx)
}
