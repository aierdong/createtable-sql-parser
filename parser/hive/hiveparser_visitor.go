// Code generated from E:/learning/sqlparser/g4/HiveParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // HiveParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by HiveParser.
type HiveParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by HiveParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by HiveParser#explainStatement.
	VisitExplainStatement(ctx *ExplainStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#explainOption.
	VisitExplainOption(ctx *ExplainOptionContext) interface{}

	// Visit a parse tree produced by HiveParser#vectorizationOnly.
	VisitVectorizationOnly(ctx *VectorizationOnlyContext) interface{}

	// Visit a parse tree produced by HiveParser#vectorizatonDetail.
	VisitVectorizatonDetail(ctx *VectorizatonDetailContext) interface{}

	// Visit a parse tree produced by HiveParser#execStatement.
	VisitExecStatement(ctx *ExecStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#loadStatement.
	VisitLoadStatement(ctx *LoadStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#replicationClause.
	VisitReplicationClause(ctx *ReplicationClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#replDumpStatement.
	VisitReplDumpStatement(ctx *ReplDumpStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#replDbPolicy.
	VisitReplDbPolicy(ctx *ReplDbPolicyContext) interface{}

	// Visit a parse tree produced by HiveParser#replLoadStatement.
	VisitReplLoadStatement(ctx *ReplLoadStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#replConfigs.
	VisitReplConfigs(ctx *ReplConfigsContext) interface{}

	// Visit a parse tree produced by HiveParser#replConfigsList.
	VisitReplConfigsList(ctx *ReplConfigsListContext) interface{}

	// Visit a parse tree produced by HiveParser#replTableLevelPolicy.
	VisitReplTableLevelPolicy(ctx *ReplTableLevelPolicyContext) interface{}

	// Visit a parse tree produced by HiveParser#replStatusStatement.
	VisitReplStatusStatement(ctx *ReplStatusStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#ddlStatement.
	VisitDdlStatement(ctx *DdlStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#ifExists.
	VisitIfExists(ctx *IfExistsContext) interface{}

	// Visit a parse tree produced by HiveParser#restrictOrCascade.
	VisitRestrictOrCascade(ctx *RestrictOrCascadeContext) interface{}

	// Visit a parse tree produced by HiveParser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) interface{}

	// Visit a parse tree produced by HiveParser#force.
	VisitForce(ctx *ForceContext) interface{}

	// Visit a parse tree produced by HiveParser#rewriteEnabled.
	VisitRewriteEnabled(ctx *RewriteEnabledContext) interface{}

	// Visit a parse tree produced by HiveParser#rewriteDisabled.
	VisitRewriteDisabled(ctx *RewriteDisabledContext) interface{}

	// Visit a parse tree produced by HiveParser#storedAsDirs.
	VisitStoredAsDirs(ctx *StoredAsDirsContext) interface{}

	// Visit a parse tree produced by HiveParser#orReplace.
	VisitOrReplace(ctx *OrReplaceContext) interface{}

	// Visit a parse tree produced by HiveParser#createDatabaseStatement.
	VisitCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dbLocation.
	VisitDbLocation(ctx *DbLocationContext) interface{}

	// Visit a parse tree produced by HiveParser#dbManagedLocation.
	VisitDbManagedLocation(ctx *DbManagedLocationContext) interface{}

	// Visit a parse tree produced by HiveParser#dbProperties.
	VisitDbProperties(ctx *DbPropertiesContext) interface{}

	// Visit a parse tree produced by HiveParser#dbPropertiesList.
	VisitDbPropertiesList(ctx *DbPropertiesListContext) interface{}

	// Visit a parse tree produced by HiveParser#dbConnectorName.
	VisitDbConnectorName(ctx *DbConnectorNameContext) interface{}

	// Visit a parse tree produced by HiveParser#switchDatabaseStatement.
	VisitSwitchDatabaseStatement(ctx *SwitchDatabaseStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dropDatabaseStatement.
	VisitDropDatabaseStatement(ctx *DropDatabaseStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#databaseComment.
	VisitDatabaseComment(ctx *DatabaseCommentContext) interface{}

	// Visit a parse tree produced by HiveParser#truncateTableStatement.
	VisitTruncateTableStatement(ctx *TruncateTableStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dropTableStatement.
	VisitDropTableStatement(ctx *DropTableStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#inputFileFormat.
	VisitInputFileFormat(ctx *InputFileFormatContext) interface{}

	// Visit a parse tree produced by HiveParser#tabTypeExpr.
	VisitTabTypeExpr(ctx *TabTypeExprContext) interface{}

	// Visit a parse tree produced by HiveParser#partTypeExpr.
	VisitPartTypeExpr(ctx *PartTypeExprContext) interface{}

	// Visit a parse tree produced by HiveParser#tabPartColTypeExpr.
	VisitTabPartColTypeExpr(ctx *TabPartColTypeExprContext) interface{}

	// Visit a parse tree produced by HiveParser#descStatement.
	VisitDescStatement(ctx *DescStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#analyzeStatement.
	VisitAnalyzeStatement(ctx *AnalyzeStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#from_in.
	VisitFrom_in(ctx *From_inContext) interface{}

	// Visit a parse tree produced by HiveParser#db_schema.
	VisitDb_schema(ctx *Db_schemaContext) interface{}

	// Visit a parse tree produced by HiveParser#showStatement.
	VisitShowStatement(ctx *ShowStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#showTablesFilterExpr.
	VisitShowTablesFilterExpr(ctx *ShowTablesFilterExprContext) interface{}

	// Visit a parse tree produced by HiveParser#lockStatement.
	VisitLockStatement(ctx *LockStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#lockDatabase.
	VisitLockDatabase(ctx *LockDatabaseContext) interface{}

	// Visit a parse tree produced by HiveParser#lockMode.
	VisitLockMode(ctx *LockModeContext) interface{}

	// Visit a parse tree produced by HiveParser#unlockStatement.
	VisitUnlockStatement(ctx *UnlockStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#unlockDatabase.
	VisitUnlockDatabase(ctx *UnlockDatabaseContext) interface{}

	// Visit a parse tree produced by HiveParser#createRoleStatement.
	VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dropRoleStatement.
	VisitDropRoleStatement(ctx *DropRoleStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#grantPrivileges.
	VisitGrantPrivileges(ctx *GrantPrivilegesContext) interface{}

	// Visit a parse tree produced by HiveParser#revokePrivileges.
	VisitRevokePrivileges(ctx *RevokePrivilegesContext) interface{}

	// Visit a parse tree produced by HiveParser#grantRole.
	VisitGrantRole(ctx *GrantRoleContext) interface{}

	// Visit a parse tree produced by HiveParser#revokeRole.
	VisitRevokeRole(ctx *RevokeRoleContext) interface{}

	// Visit a parse tree produced by HiveParser#showRoleGrants.
	VisitShowRoleGrants(ctx *ShowRoleGrantsContext) interface{}

	// Visit a parse tree produced by HiveParser#showRoles.
	VisitShowRoles(ctx *ShowRolesContext) interface{}

	// Visit a parse tree produced by HiveParser#showCurrentRole.
	VisitShowCurrentRole(ctx *ShowCurrentRoleContext) interface{}

	// Visit a parse tree produced by HiveParser#setRole.
	VisitSetRole(ctx *SetRoleContext) interface{}

	// Visit a parse tree produced by HiveParser#showGrants.
	VisitShowGrants(ctx *ShowGrantsContext) interface{}

	// Visit a parse tree produced by HiveParser#showRolePrincipals.
	VisitShowRolePrincipals(ctx *ShowRolePrincipalsContext) interface{}

	// Visit a parse tree produced by HiveParser#privilegeIncludeColObject.
	VisitPrivilegeIncludeColObject(ctx *PrivilegeIncludeColObjectContext) interface{}

	// Visit a parse tree produced by HiveParser#privilegeObject.
	VisitPrivilegeObject(ctx *PrivilegeObjectContext) interface{}

	// Visit a parse tree produced by HiveParser#privObject.
	VisitPrivObject(ctx *PrivObjectContext) interface{}

	// Visit a parse tree produced by HiveParser#privObjectCols.
	VisitPrivObjectCols(ctx *PrivObjectColsContext) interface{}

	// Visit a parse tree produced by HiveParser#privilegeList.
	VisitPrivilegeList(ctx *PrivilegeListContext) interface{}

	// Visit a parse tree produced by HiveParser#privlegeDef.
	VisitPrivlegeDef(ctx *PrivlegeDefContext) interface{}

	// Visit a parse tree produced by HiveParser#privilegeType.
	VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#principalSpecification.
	VisitPrincipalSpecification(ctx *PrincipalSpecificationContext) interface{}

	// Visit a parse tree produced by HiveParser#principalName.
	VisitPrincipalName(ctx *PrincipalNameContext) interface{}

	// Visit a parse tree produced by HiveParser#withGrantOption.
	VisitWithGrantOption(ctx *WithGrantOptionContext) interface{}

	// Visit a parse tree produced by HiveParser#grantOptionFor.
	VisitGrantOptionFor(ctx *GrantOptionForContext) interface{}

	// Visit a parse tree produced by HiveParser#adminOptionFor.
	VisitAdminOptionFor(ctx *AdminOptionForContext) interface{}

	// Visit a parse tree produced by HiveParser#withAdminOption.
	VisitWithAdminOption(ctx *WithAdminOptionContext) interface{}

	// Visit a parse tree produced by HiveParser#metastoreCheck.
	VisitMetastoreCheck(ctx *MetastoreCheckContext) interface{}

	// Visit a parse tree produced by HiveParser#resourceList.
	VisitResourceList(ctx *ResourceListContext) interface{}

	// Visit a parse tree produced by HiveParser#resource.
	VisitResource(ctx *ResourceContext) interface{}

	// Visit a parse tree produced by HiveParser#resourceType.
	VisitResourceType(ctx *ResourceTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#createFunctionStatement.
	VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dropFunctionStatement.
	VisitDropFunctionStatement(ctx *DropFunctionStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#reloadFunctionsStatement.
	VisitReloadFunctionsStatement(ctx *ReloadFunctionsStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#createMacroStatement.
	VisitCreateMacroStatement(ctx *CreateMacroStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dropMacroStatement.
	VisitDropMacroStatement(ctx *DropMacroStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#createViewStatement.
	VisitCreateViewStatement(ctx *CreateViewStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#viewPartition.
	VisitViewPartition(ctx *ViewPartitionContext) interface{}

	// Visit a parse tree produced by HiveParser#viewOrganization.
	VisitViewOrganization(ctx *ViewOrganizationContext) interface{}

	// Visit a parse tree produced by HiveParser#viewClusterSpec.
	VisitViewClusterSpec(ctx *ViewClusterSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#viewComplexSpec.
	VisitViewComplexSpec(ctx *ViewComplexSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#viewDistSpec.
	VisitViewDistSpec(ctx *ViewDistSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#viewSortSpec.
	VisitViewSortSpec(ctx *ViewSortSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#dropViewStatement.
	VisitDropViewStatement(ctx *DropViewStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#createMaterializedViewStatement.
	VisitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dropMaterializedViewStatement.
	VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#createScheduledQueryStatement.
	VisitCreateScheduledQueryStatement(ctx *CreateScheduledQueryStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dropScheduledQueryStatement.
	VisitDropScheduledQueryStatement(ctx *DropScheduledQueryStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#alterScheduledQueryStatement.
	VisitAlterScheduledQueryStatement(ctx *AlterScheduledQueryStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#alterScheduledQueryChange.
	VisitAlterScheduledQueryChange(ctx *AlterScheduledQueryChangeContext) interface{}

	// Visit a parse tree produced by HiveParser#scheduleSpec.
	VisitScheduleSpec(ctx *ScheduleSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#executedAsSpec.
	VisitExecutedAsSpec(ctx *ExecutedAsSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#definedAsSpec.
	VisitDefinedAsSpec(ctx *DefinedAsSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#showFunctionIdentifier.
	VisitShowFunctionIdentifier(ctx *ShowFunctionIdentifierContext) interface{}

	// Visit a parse tree produced by HiveParser#showStmtIdentifier.
	VisitShowStmtIdentifier(ctx *ShowStmtIdentifierContext) interface{}

	// Visit a parse tree produced by HiveParser#tableComment.
	VisitTableComment(ctx *TableCommentContext) interface{}

	// Visit a parse tree produced by HiveParser#createTablePartitionSpec.
	VisitCreateTablePartitionSpec(ctx *CreateTablePartitionSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#createTablePartitionColumnTypeSpec.
	VisitCreateTablePartitionColumnTypeSpec(ctx *CreateTablePartitionColumnTypeSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#createTablePartitionColumnSpec.
	VisitCreateTablePartitionColumnSpec(ctx *CreateTablePartitionColumnSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#partitionTransformSpec.
	VisitPartitionTransformSpec(ctx *PartitionTransformSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameTransformConstraint.
	VisitColumnNameTransformConstraint(ctx *ColumnNameTransformConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#partitionTransformType.
	VisitPartitionTransformType(ctx *PartitionTransformTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#tableBuckets.
	VisitTableBuckets(ctx *TableBucketsContext) interface{}

	// Visit a parse tree produced by HiveParser#tableImplBuckets.
	VisitTableImplBuckets(ctx *TableImplBucketsContext) interface{}

	// Visit a parse tree produced by HiveParser#tableSkewed.
	VisitTableSkewed(ctx *TableSkewedContext) interface{}

	// Visit a parse tree produced by HiveParser#rowFormat.
	VisitRowFormat(ctx *RowFormatContext) interface{}

	// Visit a parse tree produced by HiveParser#recordReader.
	VisitRecordReader(ctx *RecordReaderContext) interface{}

	// Visit a parse tree produced by HiveParser#recordWriter.
	VisitRecordWriter(ctx *RecordWriterContext) interface{}

	// Visit a parse tree produced by HiveParser#rowFormatSerde.
	VisitRowFormatSerde(ctx *RowFormatSerdeContext) interface{}

	// Visit a parse tree produced by HiveParser#rowFormatDelimited.
	VisitRowFormatDelimited(ctx *RowFormatDelimitedContext) interface{}

	// Visit a parse tree produced by HiveParser#tableRowFormat.
	VisitTableRowFormat(ctx *TableRowFormatContext) interface{}

	// Visit a parse tree produced by HiveParser#tablePropertiesPrefixed.
	VisitTablePropertiesPrefixed(ctx *TablePropertiesPrefixedContext) interface{}

	// Visit a parse tree produced by HiveParser#tableProperties.
	VisitTableProperties(ctx *TablePropertiesContext) interface{}

	// Visit a parse tree produced by HiveParser#tablePropertiesList.
	VisitTablePropertiesList(ctx *TablePropertiesListContext) interface{}

	// Visit a parse tree produced by HiveParser#keyValueProperty.
	VisitKeyValueProperty(ctx *KeyValuePropertyContext) interface{}

	// Visit a parse tree produced by HiveParser#keyProperty.
	VisitKeyProperty(ctx *KeyPropertyContext) interface{}

	// Visit a parse tree produced by HiveParser#tableRowFormatFieldIdentifier.
	VisitTableRowFormatFieldIdentifier(ctx *TableRowFormatFieldIdentifierContext) interface{}

	// Visit a parse tree produced by HiveParser#tableRowFormatCollItemsIdentifier.
	VisitTableRowFormatCollItemsIdentifier(ctx *TableRowFormatCollItemsIdentifierContext) interface{}

	// Visit a parse tree produced by HiveParser#tableRowFormatMapKeysIdentifier.
	VisitTableRowFormatMapKeysIdentifier(ctx *TableRowFormatMapKeysIdentifierContext) interface{}

	// Visit a parse tree produced by HiveParser#tableRowFormatLinesIdentifier.
	VisitTableRowFormatLinesIdentifier(ctx *TableRowFormatLinesIdentifierContext) interface{}

	// Visit a parse tree produced by HiveParser#tableRowNullFormat.
	VisitTableRowNullFormat(ctx *TableRowNullFormatContext) interface{}

	// Visit a parse tree produced by HiveParser#tableFileFormat.
	VisitTableFileFormat(ctx *TableFileFormatContext) interface{}

	// Visit a parse tree produced by HiveParser#tableLocation.
	VisitTableLocation(ctx *TableLocationContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameTypeList.
	VisitColumnNameTypeList(ctx *ColumnNameTypeListContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameTypeOrConstraintList.
	VisitColumnNameTypeOrConstraintList(ctx *ColumnNameTypeOrConstraintListContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameColonTypeList.
	VisitColumnNameColonTypeList(ctx *ColumnNameColonTypeListContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameList.
	VisitColumnNameList(ctx *ColumnNameListContext) interface{}

	// Visit a parse tree produced by HiveParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by HiveParser#extColumnName.
	VisitExtColumnName(ctx *ExtColumnNameContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameOrderList.
	VisitColumnNameOrderList(ctx *ColumnNameOrderListContext) interface{}

	// Visit a parse tree produced by HiveParser#columnParenthesesList.
	VisitColumnParenthesesList(ctx *ColumnParenthesesListContext) interface{}

	// Visit a parse tree produced by HiveParser#enableValidateSpecification.
	VisitEnableValidateSpecification(ctx *EnableValidateSpecificationContext) interface{}

	// Visit a parse tree produced by HiveParser#enableSpecification.
	VisitEnableSpecification(ctx *EnableSpecificationContext) interface{}

	// Visit a parse tree produced by HiveParser#validateSpecification.
	VisitValidateSpecification(ctx *ValidateSpecificationContext) interface{}

	// Visit a parse tree produced by HiveParser#enforcedSpecification.
	VisitEnforcedSpecification(ctx *EnforcedSpecificationContext) interface{}

	// Visit a parse tree produced by HiveParser#relySpecification.
	VisitRelySpecification(ctx *RelySpecificationContext) interface{}

	// Visit a parse tree produced by HiveParser#createConstraint.
	VisitCreateConstraint(ctx *CreateConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#alterConstraintWithName.
	VisitAlterConstraintWithName(ctx *AlterConstraintWithNameContext) interface{}

	// Visit a parse tree produced by HiveParser#tableLevelConstraint.
	VisitTableLevelConstraint(ctx *TableLevelConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#pkUkConstraint.
	VisitPkUkConstraint(ctx *PkUkConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#checkConstraint.
	VisitCheckConstraint(ctx *CheckConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#createForeignKey.
	VisitCreateForeignKey(ctx *CreateForeignKeyContext) interface{}

	// Visit a parse tree produced by HiveParser#alterForeignKeyWithName.
	VisitAlterForeignKeyWithName(ctx *AlterForeignKeyWithNameContext) interface{}

	// Visit a parse tree produced by HiveParser#skewedValueElement.
	VisitSkewedValueElement(ctx *SkewedValueElementContext) interface{}

	// Visit a parse tree produced by HiveParser#skewedColumnValuePairList.
	VisitSkewedColumnValuePairList(ctx *SkewedColumnValuePairListContext) interface{}

	// Visit a parse tree produced by HiveParser#skewedColumnValuePair.
	VisitSkewedColumnValuePair(ctx *SkewedColumnValuePairContext) interface{}

	// Visit a parse tree produced by HiveParser#skewedColumnValues.
	VisitSkewedColumnValues(ctx *SkewedColumnValuesContext) interface{}

	// Visit a parse tree produced by HiveParser#skewedColumnValue.
	VisitSkewedColumnValue(ctx *SkewedColumnValueContext) interface{}

	// Visit a parse tree produced by HiveParser#skewedValueLocationElement.
	VisitSkewedValueLocationElement(ctx *SkewedValueLocationElementContext) interface{}

	// Visit a parse tree produced by HiveParser#orderSpecification.
	VisitOrderSpecification(ctx *OrderSpecificationContext) interface{}

	// Visit a parse tree produced by HiveParser#nullOrdering.
	VisitNullOrdering(ctx *NullOrderingContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameOrder.
	VisitColumnNameOrder(ctx *ColumnNameOrderContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameCommentList.
	VisitColumnNameCommentList(ctx *ColumnNameCommentListContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameComment.
	VisitColumnNameComment(ctx *ColumnNameCommentContext) interface{}

	// Visit a parse tree produced by HiveParser#orderSpecificationRewrite.
	VisitOrderSpecificationRewrite(ctx *OrderSpecificationRewriteContext) interface{}

	// Visit a parse tree produced by HiveParser#columnRefOrder.
	VisitColumnRefOrder(ctx *ColumnRefOrderContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameType.
	VisitColumnNameType(ctx *ColumnNameTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameTypeOrConstraint.
	VisitColumnNameTypeOrConstraint(ctx *ColumnNameTypeOrConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#tableConstraint.
	VisitTableConstraint(ctx *TableConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameTypeConstraint.
	VisitColumnNameTypeConstraint(ctx *ColumnNameTypeConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#columnConstraint.
	VisitColumnConstraint(ctx *ColumnConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#foreignKeyConstraint.
	VisitForeignKeyConstraint(ctx *ForeignKeyConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#colConstraint.
	VisitColConstraint(ctx *ColConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#alterColumnConstraint.
	VisitAlterColumnConstraint(ctx *AlterColumnConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#alterForeignKeyConstraint.
	VisitAlterForeignKeyConstraint(ctx *AlterForeignKeyConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#alterColConstraint.
	VisitAlterColConstraint(ctx *AlterColConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#columnConstraintType.
	VisitColumnConstraintType(ctx *ColumnConstraintTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#defaultVal.
	VisitDefaultVal(ctx *DefaultValContext) interface{}

	// Visit a parse tree produced by HiveParser#tableConstraintType.
	VisitTableConstraintType(ctx *TableConstraintTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#constraintOptsCreate.
	VisitConstraintOptsCreate(ctx *ConstraintOptsCreateContext) interface{}

	// Visit a parse tree produced by HiveParser#constraintOptsAlter.
	VisitConstraintOptsAlter(ctx *ConstraintOptsAlterContext) interface{}

	// Visit a parse tree produced by HiveParser#columnNameColonType.
	VisitColumnNameColonType(ctx *ColumnNameColonTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#colType.
	VisitColType(ctx *ColTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#colTypeList.
	VisitColTypeList(ctx *ColTypeListContext) interface{}

	// Visit a parse tree produced by HiveParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by HiveParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#listType.
	VisitListType(ctx *ListTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#structType.
	VisitStructType(ctx *StructTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#unionType.
	VisitUnionType(ctx *UnionTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#setOperator.
	VisitSetOperator(ctx *SetOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#queryStatementExpression.
	VisitQueryStatementExpression(ctx *QueryStatementExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#queryStatementExpressionBody.
	VisitQueryStatementExpressionBody(ctx *QueryStatementExpressionBodyContext) interface{}

	// Visit a parse tree produced by HiveParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#cteStatement.
	VisitCteStatement(ctx *CteStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#fromStatement.
	VisitFromStatement(ctx *FromStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#singleFromStatement.
	VisitSingleFromStatement(ctx *SingleFromStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#regularBody.
	VisitRegularBody(ctx *RegularBodyContext) interface{}

	// Visit a parse tree produced by HiveParser#atomSelectStatement.
	VisitAtomSelectStatement(ctx *AtomSelectStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#selectStatement.
	VisitSelectStatement(ctx *SelectStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#setOpSelectStatement.
	VisitSetOpSelectStatement(ctx *SetOpSelectStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#selectStatementWithCTE.
	VisitSelectStatementWithCTE(ctx *SelectStatementWithCTEContext) interface{}

	// Visit a parse tree produced by HiveParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by HiveParser#insertClause.
	VisitInsertClause(ctx *InsertClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#destination.
	VisitDestination(ctx *DestinationContext) interface{}

	// Visit a parse tree produced by HiveParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#columnAssignmentClause.
	VisitColumnAssignmentClause(ctx *ColumnAssignmentClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#precedencePlusExpressionOrDefault.
	VisitPrecedencePlusExpressionOrDefault(ctx *PrecedencePlusExpressionOrDefaultContext) interface{}

	// Visit a parse tree produced by HiveParser#setColumnsClause.
	VisitSetColumnsClause(ctx *SetColumnsClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#sqlTransactionStatement.
	VisitSqlTransactionStatement(ctx *SqlTransactionStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#startTransactionStatement.
	VisitStartTransactionStatement(ctx *StartTransactionStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#transactionMode.
	VisitTransactionMode(ctx *TransactionModeContext) interface{}

	// Visit a parse tree produced by HiveParser#transactionAccessMode.
	VisitTransactionAccessMode(ctx *TransactionAccessModeContext) interface{}

	// Visit a parse tree produced by HiveParser#isolationLevel.
	VisitIsolationLevel(ctx *IsolationLevelContext) interface{}

	// Visit a parse tree produced by HiveParser#levelOfIsolation.
	VisitLevelOfIsolation(ctx *LevelOfIsolationContext) interface{}

	// Visit a parse tree produced by HiveParser#commitStatement.
	VisitCommitStatement(ctx *CommitStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#rollbackStatement.
	VisitRollbackStatement(ctx *RollbackStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#setAutoCommitStatement.
	VisitSetAutoCommitStatement(ctx *SetAutoCommitStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#abortTransactionStatement.
	VisitAbortTransactionStatement(ctx *AbortTransactionStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#abortCompactionStatement.
	VisitAbortCompactionStatement(ctx *AbortCompactionStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#mergeStatement.
	VisitMergeStatement(ctx *MergeStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#whenClauses.
	VisitWhenClauses(ctx *WhenClausesContext) interface{}

	// Visit a parse tree produced by HiveParser#whenNotMatchedClause.
	VisitWhenNotMatchedClause(ctx *WhenNotMatchedClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#whenMatchedAndClause.
	VisitWhenMatchedAndClause(ctx *WhenMatchedAndClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#whenMatchedThenClause.
	VisitWhenMatchedThenClause(ctx *WhenMatchedThenClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#updateOrDelete.
	VisitUpdateOrDelete(ctx *UpdateOrDeleteContext) interface{}

	// Visit a parse tree produced by HiveParser#killQueryStatement.
	VisitKillQueryStatement(ctx *KillQueryStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#compactionId.
	VisitCompactionId(ctx *CompactionIdContext) interface{}

	// Visit a parse tree produced by HiveParser#compactionPool.
	VisitCompactionPool(ctx *CompactionPoolContext) interface{}

	// Visit a parse tree produced by HiveParser#compactionType.
	VisitCompactionType(ctx *CompactionTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#compactionStatus.
	VisitCompactionStatus(ctx *CompactionStatusContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatement.
	VisitAlterStatement(ctx *AlterStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#alterTableStatementSuffix.
	VisitAlterTableStatementSuffix(ctx *AlterTableStatementSuffixContext) interface{}

	// Visit a parse tree produced by HiveParser#alterTblPartitionStatementSuffix.
	VisitAlterTblPartitionStatementSuffix(ctx *AlterTblPartitionStatementSuffixContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementPartitionKeyType.
	VisitAlterStatementPartitionKeyType(ctx *AlterStatementPartitionKeyTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#alterViewStatementSuffix.
	VisitAlterViewStatementSuffix(ctx *AlterViewStatementSuffixContext) interface{}

	// Visit a parse tree produced by HiveParser#alterMaterializedViewStatementSuffix.
	VisitAlterMaterializedViewStatementSuffix(ctx *AlterMaterializedViewStatementSuffixContext) interface{}

	// Visit a parse tree produced by HiveParser#alterMaterializedViewSuffixRewrite.
	VisitAlterMaterializedViewSuffixRewrite(ctx *AlterMaterializedViewSuffixRewriteContext) interface{}

	// Visit a parse tree produced by HiveParser#alterMaterializedViewSuffixRebuild.
	VisitAlterMaterializedViewSuffixRebuild(ctx *AlterMaterializedViewSuffixRebuildContext) interface{}

	// Visit a parse tree produced by HiveParser#alterDatabaseStatementSuffix.
	VisitAlterDatabaseStatementSuffix(ctx *AlterDatabaseStatementSuffixContext) interface{}

	// Visit a parse tree produced by HiveParser#alterDatabaseSuffixProperties.
	VisitAlterDatabaseSuffixProperties(ctx *AlterDatabaseSuffixPropertiesContext) interface{}

	// Visit a parse tree produced by HiveParser#alterDatabaseSuffixSetOwner.
	VisitAlterDatabaseSuffixSetOwner(ctx *AlterDatabaseSuffixSetOwnerContext) interface{}

	// Visit a parse tree produced by HiveParser#alterDatabaseSuffixSetLocation.
	VisitAlterDatabaseSuffixSetLocation(ctx *AlterDatabaseSuffixSetLocationContext) interface{}

	// Visit a parse tree produced by HiveParser#alterDatabaseSuffixSetManagedLocation.
	VisitAlterDatabaseSuffixSetManagedLocation(ctx *AlterDatabaseSuffixSetManagedLocationContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixRename.
	VisitAlterStatementSuffixRename(ctx *AlterStatementSuffixRenameContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixAddCol.
	VisitAlterStatementSuffixAddCol(ctx *AlterStatementSuffixAddColContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixAddConstraint.
	VisitAlterStatementSuffixAddConstraint(ctx *AlterStatementSuffixAddConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixUpdateColumns.
	VisitAlterStatementSuffixUpdateColumns(ctx *AlterStatementSuffixUpdateColumnsContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixDropConstraint.
	VisitAlterStatementSuffixDropConstraint(ctx *AlterStatementSuffixDropConstraintContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixRenameCol.
	VisitAlterStatementSuffixRenameCol(ctx *AlterStatementSuffixRenameColContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixUpdateStatsCol.
	VisitAlterStatementSuffixUpdateStatsCol(ctx *AlterStatementSuffixUpdateStatsColContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixUpdateStats.
	VisitAlterStatementSuffixUpdateStats(ctx *AlterStatementSuffixUpdateStatsContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementChangeColPosition.
	VisitAlterStatementChangeColPosition(ctx *AlterStatementChangeColPositionContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixAddPartitions.
	VisitAlterStatementSuffixAddPartitions(ctx *AlterStatementSuffixAddPartitionsContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixAddPartitionsElement.
	VisitAlterStatementSuffixAddPartitionsElement(ctx *AlterStatementSuffixAddPartitionsElementContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixTouch.
	VisitAlterStatementSuffixTouch(ctx *AlterStatementSuffixTouchContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixArchive.
	VisitAlterStatementSuffixArchive(ctx *AlterStatementSuffixArchiveContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixUnArchive.
	VisitAlterStatementSuffixUnArchive(ctx *AlterStatementSuffixUnArchiveContext) interface{}

	// Visit a parse tree produced by HiveParser#partitionLocation.
	VisitPartitionLocation(ctx *PartitionLocationContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixDropPartitions.
	VisitAlterStatementSuffixDropPartitions(ctx *AlterStatementSuffixDropPartitionsContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixProperties.
	VisitAlterStatementSuffixProperties(ctx *AlterStatementSuffixPropertiesContext) interface{}

	// Visit a parse tree produced by HiveParser#alterViewSuffixProperties.
	VisitAlterViewSuffixProperties(ctx *AlterViewSuffixPropertiesContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixSerdeProperties.
	VisitAlterStatementSuffixSerdeProperties(ctx *AlterStatementSuffixSerdePropertiesContext) interface{}

	// Visit a parse tree produced by HiveParser#tablePartitionPrefix.
	VisitTablePartitionPrefix(ctx *TablePartitionPrefixContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixFileFormat.
	VisitAlterStatementSuffixFileFormat(ctx *AlterStatementSuffixFileFormatContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixClusterbySortby.
	VisitAlterStatementSuffixClusterbySortby(ctx *AlterStatementSuffixClusterbySortbyContext) interface{}

	// Visit a parse tree produced by HiveParser#alterTblPartitionStatementSuffixSkewedLocation.
	VisitAlterTblPartitionStatementSuffixSkewedLocation(ctx *AlterTblPartitionStatementSuffixSkewedLocationContext) interface{}

	// Visit a parse tree produced by HiveParser#skewedLocations.
	VisitSkewedLocations(ctx *SkewedLocationsContext) interface{}

	// Visit a parse tree produced by HiveParser#skewedLocationsList.
	VisitSkewedLocationsList(ctx *SkewedLocationsListContext) interface{}

	// Visit a parse tree produced by HiveParser#skewedLocationMap.
	VisitSkewedLocationMap(ctx *SkewedLocationMapContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixLocation.
	VisitAlterStatementSuffixLocation(ctx *AlterStatementSuffixLocationContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixSkewedby.
	VisitAlterStatementSuffixSkewedby(ctx *AlterStatementSuffixSkewedbyContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixExchangePartition.
	VisitAlterStatementSuffixExchangePartition(ctx *AlterStatementSuffixExchangePartitionContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixRenamePart.
	VisitAlterStatementSuffixRenamePart(ctx *AlterStatementSuffixRenamePartContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixStatsPart.
	VisitAlterStatementSuffixStatsPart(ctx *AlterStatementSuffixStatsPartContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixMergeFiles.
	VisitAlterStatementSuffixMergeFiles(ctx *AlterStatementSuffixMergeFilesContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixBucketNum.
	VisitAlterStatementSuffixBucketNum(ctx *AlterStatementSuffixBucketNumContext) interface{}

	// Visit a parse tree produced by HiveParser#blocking.
	VisitBlocking(ctx *BlockingContext) interface{}

	// Visit a parse tree produced by HiveParser#compactPool.
	VisitCompactPool(ctx *CompactPoolContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixCompact.
	VisitAlterStatementSuffixCompact(ctx *AlterStatementSuffixCompactContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixSetOwner.
	VisitAlterStatementSuffixSetOwner(ctx *AlterStatementSuffixSetOwnerContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixSetPartSpec.
	VisitAlterStatementSuffixSetPartSpec(ctx *AlterStatementSuffixSetPartSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#alterStatementSuffixExecute.
	VisitAlterStatementSuffixExecute(ctx *AlterStatementSuffixExecuteContext) interface{}

	// Visit a parse tree produced by HiveParser#fileFormat.
	VisitFileFormat(ctx *FileFormatContext) interface{}

	// Visit a parse tree produced by HiveParser#alterDataConnectorStatementSuffix.
	VisitAlterDataConnectorStatementSuffix(ctx *AlterDataConnectorStatementSuffixContext) interface{}

	// Visit a parse tree produced by HiveParser#alterDataConnectorSuffixProperties.
	VisitAlterDataConnectorSuffixProperties(ctx *AlterDataConnectorSuffixPropertiesContext) interface{}

	// Visit a parse tree produced by HiveParser#alterDataConnectorSuffixSetOwner.
	VisitAlterDataConnectorSuffixSetOwner(ctx *AlterDataConnectorSuffixSetOwnerContext) interface{}

	// Visit a parse tree produced by HiveParser#alterDataConnectorSuffixSetUrl.
	VisitAlterDataConnectorSuffixSetUrl(ctx *AlterDataConnectorSuffixSetUrlContext) interface{}

	// Visit a parse tree produced by HiveParser#likeTableOrFile.
	VisitLikeTableOrFile(ctx *LikeTableOrFileContext) interface{}

	// Visit a parse tree produced by HiveParser#createTableStatement.
	VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#createDataConnectorStatement.
	VisitCreateDataConnectorStatement(ctx *CreateDataConnectorStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dataConnectorComment.
	VisitDataConnectorComment(ctx *DataConnectorCommentContext) interface{}

	// Visit a parse tree produced by HiveParser#dataConnectorUrl.
	VisitDataConnectorUrl(ctx *DataConnectorUrlContext) interface{}

	// Visit a parse tree produced by HiveParser#dataConnectorType.
	VisitDataConnectorType(ctx *DataConnectorTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#dcProperties.
	VisitDcProperties(ctx *DcPropertiesContext) interface{}

	// Visit a parse tree produced by HiveParser#dropDataConnectorStatement.
	VisitDropDataConnectorStatement(ctx *DropDataConnectorStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#tableAllColumns.
	VisitTableAllColumns(ctx *TableAllColumnsContext) interface{}

	// Visit a parse tree produced by HiveParser#tableOrColumn.
	VisitTableOrColumn(ctx *TableOrColumnContext) interface{}

	// Visit a parse tree produced by HiveParser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) interface{}

	// Visit a parse tree produced by HiveParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by HiveParser#aliasList.
	VisitAliasList(ctx *AliasListContext) interface{}

	// Visit a parse tree produced by HiveParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#fromSource.
	VisitFromSource(ctx *FromSourceContext) interface{}

	// Visit a parse tree produced by HiveParser#atomjoinSource.
	VisitAtomjoinSource(ctx *AtomjoinSourceContext) interface{}

	// Visit a parse tree produced by HiveParser#joinSource.
	VisitJoinSource(ctx *JoinSourceContext) interface{}

	// Visit a parse tree produced by HiveParser#joinSourcePart.
	VisitJoinSourcePart(ctx *JoinSourcePartContext) interface{}

	// Visit a parse tree produced by HiveParser#uniqueJoinSource.
	VisitUniqueJoinSource(ctx *UniqueJoinSourceContext) interface{}

	// Visit a parse tree produced by HiveParser#uniqueJoinExpr.
	VisitUniqueJoinExpr(ctx *UniqueJoinExprContext) interface{}

	// Visit a parse tree produced by HiveParser#uniqueJoinToken.
	VisitUniqueJoinToken(ctx *UniqueJoinTokenContext) interface{}

	// Visit a parse tree produced by HiveParser#joinToken.
	VisitJoinToken(ctx *JoinTokenContext) interface{}

	// Visit a parse tree produced by HiveParser#lateralView.
	VisitLateralView(ctx *LateralViewContext) interface{}

	// Visit a parse tree produced by HiveParser#tableAlias.
	VisitTableAlias(ctx *TableAliasContext) interface{}

	// Visit a parse tree produced by HiveParser#tableBucketSample.
	VisitTableBucketSample(ctx *TableBucketSampleContext) interface{}

	// Visit a parse tree produced by HiveParser#splitSample.
	VisitSplitSample(ctx *SplitSampleContext) interface{}

	// Visit a parse tree produced by HiveParser#tableSample.
	VisitTableSample(ctx *TableSampleContext) interface{}

	// Visit a parse tree produced by HiveParser#tableSource.
	VisitTableSource(ctx *TableSourceContext) interface{}

	// Visit a parse tree produced by HiveParser#asOfClause.
	VisitAsOfClause(ctx *AsOfClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#uniqueJoinTableSource.
	VisitUniqueJoinTableSource(ctx *UniqueJoinTableSourceContext) interface{}

	// Visit a parse tree produced by HiveParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by HiveParser#viewName.
	VisitViewName(ctx *ViewNameContext) interface{}

	// Visit a parse tree produced by HiveParser#subQuerySource.
	VisitSubQuerySource(ctx *SubQuerySourceContext) interface{}

	// Visit a parse tree produced by HiveParser#partitioningSpec.
	VisitPartitioningSpec(ctx *PartitioningSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#partitionTableFunctionSource.
	VisitPartitionTableFunctionSource(ctx *PartitionTableFunctionSourceContext) interface{}

	// Visit a parse tree produced by HiveParser#partitionedTableFunction.
	VisitPartitionedTableFunction(ctx *PartitionedTableFunctionContext) interface{}

	// Visit a parse tree produced by HiveParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#searchCondition.
	VisitSearchCondition(ctx *SearchConditionContext) interface{}

	// Visit a parse tree produced by HiveParser#valuesSource.
	VisitValuesSource(ctx *ValuesSourceContext) interface{}

	// Visit a parse tree produced by HiveParser#valuesClause.
	VisitValuesClause(ctx *ValuesClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#valuesTableConstructor.
	VisitValuesTableConstructor(ctx *ValuesTableConstructorContext) interface{}

	// Visit a parse tree produced by HiveParser#valueRowConstructor.
	VisitValueRowConstructor(ctx *ValueRowConstructorContext) interface{}

	// Visit a parse tree produced by HiveParser#firstValueRowConstructor.
	VisitFirstValueRowConstructor(ctx *FirstValueRowConstructorContext) interface{}

	// Visit a parse tree produced by HiveParser#virtualTableSource.
	VisitVirtualTableSource(ctx *VirtualTableSourceContext) interface{}

	// Visit a parse tree produced by HiveParser#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#all_distinct.
	VisitAll_distinct(ctx *All_distinctContext) interface{}

	// Visit a parse tree produced by HiveParser#selectList.
	VisitSelectList(ctx *SelectListContext) interface{}

	// Visit a parse tree produced by HiveParser#selectTrfmClause.
	VisitSelectTrfmClause(ctx *SelectTrfmClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#selectItem.
	VisitSelectItem(ctx *SelectItemContext) interface{}

	// Visit a parse tree produced by HiveParser#trfmClause.
	VisitTrfmClause(ctx *TrfmClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#selectExpression.
	VisitSelectExpression(ctx *SelectExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#selectExpressionList.
	VisitSelectExpressionList(ctx *SelectExpressionListContext) interface{}

	// Visit a parse tree produced by HiveParser#window_clause.
	VisitWindow_clause(ctx *Window_clauseContext) interface{}

	// Visit a parse tree produced by HiveParser#window_defn.
	VisitWindow_defn(ctx *Window_defnContext) interface{}

	// Visit a parse tree produced by HiveParser#window_specification.
	VisitWindow_specification(ctx *Window_specificationContext) interface{}

	// Visit a parse tree produced by HiveParser#window_frame.
	VisitWindow_frame(ctx *Window_frameContext) interface{}

	// Visit a parse tree produced by HiveParser#window_range_expression.
	VisitWindow_range_expression(ctx *Window_range_expressionContext) interface{}

	// Visit a parse tree produced by HiveParser#window_value_expression.
	VisitWindow_value_expression(ctx *Window_value_expressionContext) interface{}

	// Visit a parse tree produced by HiveParser#window_frame_start_boundary.
	VisitWindow_frame_start_boundary(ctx *Window_frame_start_boundaryContext) interface{}

	// Visit a parse tree produced by HiveParser#window_frame_boundary.
	VisitWindow_frame_boundary(ctx *Window_frame_boundaryContext) interface{}

	// Visit a parse tree produced by HiveParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#groupby_expression.
	VisitGroupby_expression(ctx *Groupby_expressionContext) interface{}

	// Visit a parse tree produced by HiveParser#groupByEmpty.
	VisitGroupByEmpty(ctx *GroupByEmptyContext) interface{}

	// Visit a parse tree produced by HiveParser#rollupStandard.
	VisitRollupStandard(ctx *RollupStandardContext) interface{}

	// Visit a parse tree produced by HiveParser#rollupOldSyntax.
	VisitRollupOldSyntax(ctx *RollupOldSyntaxContext) interface{}

	// Visit a parse tree produced by HiveParser#groupingSetExpression.
	VisitGroupingSetExpression(ctx *GroupingSetExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#groupingSetExpressionMultiple.
	VisitGroupingSetExpressionMultiple(ctx *GroupingSetExpressionMultipleContext) interface{}

	// Visit a parse tree produced by HiveParser#groupingExpressionSingle.
	VisitGroupingExpressionSingle(ctx *GroupingExpressionSingleContext) interface{}

	// Visit a parse tree produced by HiveParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#qualifyClause.
	VisitQualifyClause(ctx *QualifyClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#havingCondition.
	VisitHavingCondition(ctx *HavingConditionContext) interface{}

	// Visit a parse tree produced by HiveParser#expressionsInParenthesis.
	VisitExpressionsInParenthesis(ctx *ExpressionsInParenthesisContext) interface{}

	// Visit a parse tree produced by HiveParser#expressionsNotInParenthesis.
	VisitExpressionsNotInParenthesis(ctx *ExpressionsNotInParenthesisContext) interface{}

	// Visit a parse tree produced by HiveParser#expressionPart.
	VisitExpressionPart(ctx *ExpressionPartContext) interface{}

	// Visit a parse tree produced by HiveParser#expressionOrDefault.
	VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{}

	// Visit a parse tree produced by HiveParser#firstExpressionsWithAlias.
	VisitFirstExpressionsWithAlias(ctx *FirstExpressionsWithAliasContext) interface{}

	// Visit a parse tree produced by HiveParser#expressionWithAlias.
	VisitExpressionWithAlias(ctx *ExpressionWithAliasContext) interface{}

	// Visit a parse tree produced by HiveParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) interface{}

	// Visit a parse tree produced by HiveParser#columnRefOrderInParenthesis.
	VisitColumnRefOrderInParenthesis(ctx *ColumnRefOrderInParenthesisContext) interface{}

	// Visit a parse tree produced by HiveParser#columnRefOrderNotInParenthesis.
	VisitColumnRefOrderNotInParenthesis(ctx *ColumnRefOrderNotInParenthesisContext) interface{}

	// Visit a parse tree produced by HiveParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#clusterByClause.
	VisitClusterByClause(ctx *ClusterByClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#partitionByClause.
	VisitPartitionByClause(ctx *PartitionByClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#distributeByClause.
	VisitDistributeByClause(ctx *DistributeByClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#sortByClause.
	VisitSortByClause(ctx *SortByClauseContext) interface{}

	// Visit a parse tree produced by HiveParser#trimFunction.
	VisitTrimFunction(ctx *TrimFunctionContext) interface{}

	// Visit a parse tree produced by HiveParser#function_.
	VisitFunction_(ctx *Function_Context) interface{}

	// Visit a parse tree produced by HiveParser#null_treatment.
	VisitNull_treatment(ctx *Null_treatmentContext) interface{}

	// Visit a parse tree produced by HiveParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by HiveParser#castExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#caseExpression.
	VisitCaseExpression(ctx *CaseExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#whenExpression.
	VisitWhenExpression(ctx *WhenExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#floorExpression.
	VisitFloorExpression(ctx *FloorExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#floorDateQualifiers.
	VisitFloorDateQualifiers(ctx *FloorDateQualifiersContext) interface{}

	// Visit a parse tree produced by HiveParser#extractExpression.
	VisitExtractExpression(ctx *ExtractExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#timeQualifiers.
	VisitTimeQualifiers(ctx *TimeQualifiersContext) interface{}

	// Visit a parse tree produced by HiveParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by HiveParser#prepareStmtParam.
	VisitPrepareStmtParam(ctx *PrepareStmtParamContext) interface{}

	// Visit a parse tree produced by HiveParser#parameterIdx.
	VisitParameterIdx(ctx *ParameterIdxContext) interface{}

	// Visit a parse tree produced by HiveParser#stringLiteralSequence.
	VisitStringLiteralSequence(ctx *StringLiteralSequenceContext) interface{}

	// Visit a parse tree produced by HiveParser#charSetStringLiteral.
	VisitCharSetStringLiteral(ctx *CharSetStringLiteralContext) interface{}

	// Visit a parse tree produced by HiveParser#dateLiteral.
	VisitDateLiteral(ctx *DateLiteralContext) interface{}

	// Visit a parse tree produced by HiveParser#timestampLiteral.
	VisitTimestampLiteral(ctx *TimestampLiteralContext) interface{}

	// Visit a parse tree produced by HiveParser#timestampLocalTZLiteral.
	VisitTimestampLocalTZLiteral(ctx *TimestampLocalTZLiteralContext) interface{}

	// Visit a parse tree produced by HiveParser#intervalValue.
	VisitIntervalValue(ctx *IntervalValueContext) interface{}

	// Visit a parse tree produced by HiveParser#intervalLiteral.
	VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{}

	// Visit a parse tree produced by HiveParser#intervalExpression.
	VisitIntervalExpression(ctx *IntervalExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#intervalQualifiers.
	VisitIntervalQualifiers(ctx *IntervalQualifiersContext) interface{}

	// Visit a parse tree produced by HiveParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#atomExpression.
	VisitAtomExpression(ctx *AtomExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceFieldExpression.
	VisitPrecedenceFieldExpression(ctx *PrecedenceFieldExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceUnaryOperator.
	VisitPrecedenceUnaryOperator(ctx *PrecedenceUnaryOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceUnaryPrefixExpression.
	VisitPrecedenceUnaryPrefixExpression(ctx *PrecedenceUnaryPrefixExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceBitwiseXorOperator.
	VisitPrecedenceBitwiseXorOperator(ctx *PrecedenceBitwiseXorOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceBitwiseXorExpression.
	VisitPrecedenceBitwiseXorExpression(ctx *PrecedenceBitwiseXorExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceStarOperator.
	VisitPrecedenceStarOperator(ctx *PrecedenceStarOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceStarExpression.
	VisitPrecedenceStarExpression(ctx *PrecedenceStarExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedencePlusOperator.
	VisitPrecedencePlusOperator(ctx *PrecedencePlusOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedencePlusExpression.
	VisitPrecedencePlusExpression(ctx *PrecedencePlusExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceConcatenateOperator.
	VisitPrecedenceConcatenateOperator(ctx *PrecedenceConcatenateOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceConcatenateExpression.
	VisitPrecedenceConcatenateExpression(ctx *PrecedenceConcatenateExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceAmpersandOperator.
	VisitPrecedenceAmpersandOperator(ctx *PrecedenceAmpersandOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceAmpersandExpression.
	VisitPrecedenceAmpersandExpression(ctx *PrecedenceAmpersandExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceBitwiseOrOperator.
	VisitPrecedenceBitwiseOrOperator(ctx *PrecedenceBitwiseOrOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceBitwiseOrExpression.
	VisitPrecedenceBitwiseOrExpression(ctx *PrecedenceBitwiseOrExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceRegexpOperator.
	VisitPrecedenceRegexpOperator(ctx *PrecedenceRegexpOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceSimilarOperator.
	VisitPrecedenceSimilarOperator(ctx *PrecedenceSimilarOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#subQueryExpression.
	VisitSubQueryExpression(ctx *SubQueryExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceSimilarExpression.
	VisitPrecedenceSimilarExpression(ctx *PrecedenceSimilarExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceSimilarExpressionMain.
	VisitPrecedenceSimilarExpressionMain(ctx *PrecedenceSimilarExpressionMainContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceSimilarExpressionPart.
	VisitPrecedenceSimilarExpressionPart(ctx *PrecedenceSimilarExpressionPartContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceSimilarExpressionAtom.
	VisitPrecedenceSimilarExpressionAtom(ctx *PrecedenceSimilarExpressionAtomContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceSimilarExpressionQuantifierPredicate.
	VisitPrecedenceSimilarExpressionQuantifierPredicate(ctx *PrecedenceSimilarExpressionQuantifierPredicateContext) interface{}

	// Visit a parse tree produced by HiveParser#quantifierType.
	VisitQuantifierType(ctx *QuantifierTypeContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceSimilarExpressionIn.
	VisitPrecedenceSimilarExpressionIn(ctx *PrecedenceSimilarExpressionInContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceSimilarExpressionPartNot.
	VisitPrecedenceSimilarExpressionPartNot(ctx *PrecedenceSimilarExpressionPartNotContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceDistinctOperator.
	VisitPrecedenceDistinctOperator(ctx *PrecedenceDistinctOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceEqualOperator.
	VisitPrecedenceEqualOperator(ctx *PrecedenceEqualOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceEqualExpression.
	VisitPrecedenceEqualExpression(ctx *PrecedenceEqualExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#isCondition.
	VisitIsCondition(ctx *IsConditionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceUnarySuffixExpression.
	VisitPrecedenceUnarySuffixExpression(ctx *PrecedenceUnarySuffixExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceNotOperator.
	VisitPrecedenceNotOperator(ctx *PrecedenceNotOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceNotExpression.
	VisitPrecedenceNotExpression(ctx *PrecedenceNotExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceAndOperator.
	VisitPrecedenceAndOperator(ctx *PrecedenceAndOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceAndExpression.
	VisitPrecedenceAndExpression(ctx *PrecedenceAndExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceOrOperator.
	VisitPrecedenceOrOperator(ctx *PrecedenceOrOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#precedenceOrExpression.
	VisitPrecedenceOrExpression(ctx *PrecedenceOrExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#booleanValue.
	VisitBooleanValue(ctx *BooleanValueContext) interface{}

	// Visit a parse tree produced by HiveParser#booleanValueTok.
	VisitBooleanValueTok(ctx *BooleanValueTokContext) interface{}

	// Visit a parse tree produced by HiveParser#tableOrPartition.
	VisitTableOrPartition(ctx *TableOrPartitionContext) interface{}

	// Visit a parse tree produced by HiveParser#partitionSpec.
	VisitPartitionSpec(ctx *PartitionSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#partitionVal.
	VisitPartitionVal(ctx *PartitionValContext) interface{}

	// Visit a parse tree produced by HiveParser#partitionSelectorSpec.
	VisitPartitionSelectorSpec(ctx *PartitionSelectorSpecContext) interface{}

	// Visit a parse tree produced by HiveParser#partitionSelectorVal.
	VisitPartitionSelectorVal(ctx *PartitionSelectorValContext) interface{}

	// Visit a parse tree produced by HiveParser#partitionSelectorOperator.
	VisitPartitionSelectorOperator(ctx *PartitionSelectorOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#subQuerySelectorOperator.
	VisitSubQuerySelectorOperator(ctx *SubQuerySelectorOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#sysFuncNames.
	VisitSysFuncNames(ctx *SysFuncNamesContext) interface{}

	// Visit a parse tree produced by HiveParser#descFuncNames.
	VisitDescFuncNames(ctx *DescFuncNamesContext) interface{}

	// Visit a parse tree produced by HiveParser#id_.
	VisitId_(ctx *Id_Context) interface{}

	// Visit a parse tree produced by HiveParser#functionIdentifier.
	VisitFunctionIdentifier(ctx *FunctionIdentifierContext) interface{}

	// Visit a parse tree produced by HiveParser#principalIdentifier.
	VisitPrincipalIdentifier(ctx *PrincipalIdentifierContext) interface{}

	// Visit a parse tree produced by HiveParser#nonReserved.
	VisitNonReserved(ctx *NonReservedContext) interface{}

	// Visit a parse tree produced by HiveParser#sql11ReservedKeywordsUsedAsFunctionName.
	VisitSql11ReservedKeywordsUsedAsFunctionName(ctx *Sql11ReservedKeywordsUsedAsFunctionNameContext) interface{}

	// Visit a parse tree produced by HiveParser#hint.
	VisitHint(ctx *HintContext) interface{}

	// Visit a parse tree produced by HiveParser#hintList.
	VisitHintList(ctx *HintListContext) interface{}

	// Visit a parse tree produced by HiveParser#hintItem.
	VisitHintItem(ctx *HintItemContext) interface{}

	// Visit a parse tree produced by HiveParser#hintName.
	VisitHintName(ctx *HintNameContext) interface{}

	// Visit a parse tree produced by HiveParser#hintArgs.
	VisitHintArgs(ctx *HintArgsContext) interface{}

	// Visit a parse tree produced by HiveParser#hintArgName.
	VisitHintArgName(ctx *HintArgNameContext) interface{}

	// Visit a parse tree produced by HiveParser#prepareStatement.
	VisitPrepareStatement(ctx *PrepareStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#executeStatement.
	VisitExecuteStatement(ctx *ExecuteStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#executeParamList.
	VisitExecuteParamList(ctx *ExecuteParamListContext) interface{}

	// Visit a parse tree produced by HiveParser#resourcePlanDdlStatements.
	VisitResourcePlanDdlStatements(ctx *ResourcePlanDdlStatementsContext) interface{}

	// Visit a parse tree produced by HiveParser#rpAssign.
	VisitRpAssign(ctx *RpAssignContext) interface{}

	// Visit a parse tree produced by HiveParser#rpAssignList.
	VisitRpAssignList(ctx *RpAssignListContext) interface{}

	// Visit a parse tree produced by HiveParser#rpUnassign.
	VisitRpUnassign(ctx *RpUnassignContext) interface{}

	// Visit a parse tree produced by HiveParser#rpUnassignList.
	VisitRpUnassignList(ctx *RpUnassignListContext) interface{}

	// Visit a parse tree produced by HiveParser#createResourcePlanStatement.
	VisitCreateResourcePlanStatement(ctx *CreateResourcePlanStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#withReplace.
	VisitWithReplace(ctx *WithReplaceContext) interface{}

	// Visit a parse tree produced by HiveParser#activate.
	VisitActivate(ctx *ActivateContext) interface{}

	// Visit a parse tree produced by HiveParser#enable.
	VisitEnable(ctx *EnableContext) interface{}

	// Visit a parse tree produced by HiveParser#disable.
	VisitDisable(ctx *DisableContext) interface{}

	// Visit a parse tree produced by HiveParser#unmanaged.
	VisitUnmanaged(ctx *UnmanagedContext) interface{}

	// Visit a parse tree produced by HiveParser#alterResourcePlanStatement.
	VisitAlterResourcePlanStatement(ctx *AlterResourcePlanStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#globalWmStatement.
	VisitGlobalWmStatement(ctx *GlobalWmStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#replaceResourcePlanStatement.
	VisitReplaceResourcePlanStatement(ctx *ReplaceResourcePlanStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dropResourcePlanStatement.
	VisitDropResourcePlanStatement(ctx *DropResourcePlanStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#poolPath.
	VisitPoolPath(ctx *PoolPathContext) interface{}

	// Visit a parse tree produced by HiveParser#triggerExpression.
	VisitTriggerExpression(ctx *TriggerExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#triggerExpressionStandalone.
	VisitTriggerExpressionStandalone(ctx *TriggerExpressionStandaloneContext) interface{}

	// Visit a parse tree produced by HiveParser#triggerOrExpression.
	VisitTriggerOrExpression(ctx *TriggerOrExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#triggerAndExpression.
	VisitTriggerAndExpression(ctx *TriggerAndExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#triggerAtomExpression.
	VisitTriggerAtomExpression(ctx *TriggerAtomExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#triggerLiteral.
	VisitTriggerLiteral(ctx *TriggerLiteralContext) interface{}

	// Visit a parse tree produced by HiveParser#comparisionOperator.
	VisitComparisionOperator(ctx *ComparisionOperatorContext) interface{}

	// Visit a parse tree produced by HiveParser#triggerActionExpression.
	VisitTriggerActionExpression(ctx *TriggerActionExpressionContext) interface{}

	// Visit a parse tree produced by HiveParser#triggerActionExpressionStandalone.
	VisitTriggerActionExpressionStandalone(ctx *TriggerActionExpressionStandaloneContext) interface{}

	// Visit a parse tree produced by HiveParser#createTriggerStatement.
	VisitCreateTriggerStatement(ctx *CreateTriggerStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#alterTriggerStatement.
	VisitAlterTriggerStatement(ctx *AlterTriggerStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dropTriggerStatement.
	VisitDropTriggerStatement(ctx *DropTriggerStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#poolAssign.
	VisitPoolAssign(ctx *PoolAssignContext) interface{}

	// Visit a parse tree produced by HiveParser#poolAssignList.
	VisitPoolAssignList(ctx *PoolAssignListContext) interface{}

	// Visit a parse tree produced by HiveParser#createPoolStatement.
	VisitCreatePoolStatement(ctx *CreatePoolStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#alterPoolStatement.
	VisitAlterPoolStatement(ctx *AlterPoolStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dropPoolStatement.
	VisitDropPoolStatement(ctx *DropPoolStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#createMappingStatement.
	VisitCreateMappingStatement(ctx *CreateMappingStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#alterMappingStatement.
	VisitAlterMappingStatement(ctx *AlterMappingStatementContext) interface{}

	// Visit a parse tree produced by HiveParser#dropMappingStatement.
	VisitDropMappingStatement(ctx *DropMappingStatementContext) interface{}
}
