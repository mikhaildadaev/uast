package uast

// Приватные интерфейсы
type renderer interface {
	elementRenderer
	componentRenderer
	statementRenderer
}
type elementRenderer interface {
	renderAlias(value string) error
	renderConstant(value any) error
	renderFunction(value string) error
	renderName(value string) error
	renderLiteral(value any) error
	renderOperator(value any) error
	renderService(value any) error
	renderValue(value any) error
}
type componentRenderer interface {
	renderAs() error
	renderCascade(entity SourceBase, cascade bool) error
	renderCommand(command managementService) error
	renderColumns(command managementService, addColumns []markSourceable, addConstraints []ConstraintBase, dropColumns []markSourceable, dropConstraints []ConstraintBase) error
	renderDistinct(distinct bool) error
	renderEntity(entity SourceBase, isReplace bool, ifExists bool, ifNotExists bool) error
	renderFields(fields []markExpressable, isParen bool) error
	renderFrom(from SourceBase) error
	renderGroupBy(groups []markGroupable) error
	renderHaving(having ExpressionBase) error
	renderIfExists(ifExists bool) error
	renderIfNotExists(ifNotExists bool) error
	renderIndex(columns []markSourceable) error
	renderInTo(inTo SourceBase) error
	renderIsData(data string) error
	renderJoin(joins []*clauseJoin) error
	renderOn(on SourceBase) error
	renderOnFrom(onTable SourceBase, onColumn SourceBase) error
	renderOnTo(onTo SourceBase) error
	renderOrderBy(orders []markOrderable) error
	renderPagination(pagination *clausePagination) error
	renderRenameColumn(columnRename *columnRename) error
	renderRenameConstraint(constraintRename *constraintRename) error
	renderRenameTo(name string) error
	renderReplace(replace bool) error
	renderRestartIdentity(restartIdentity bool) error
	renderReturning(returnings *clauseReturning) error
	renderSet(sets []*clauseSet) error
	renderSource(source statement) error
	renderTable(table *TableSource) error
	renderTarget(source SourceBase) error
	renderUnions(unions []*clauseUnions) error
	renderUnique(unique bool) error
	renderUsing(joins []*clauseJoin) error
	renderValues(values *clauseValues) error
	renderWhere(where ExpressionBase) error
	renderWith(withs []*clauseWith) error
}
type statementRenderer interface {
	// DDL
	renderAlter(baseRenderer *baseRenderer, stmtAlter *stmtAlter) error
	renderComment(baseRenderer *baseRenderer, stmtComment *stmtComment) error
	renderCreate(baseRenderer *baseRenderer, stmtCreate *stmtCreate) error
	renderDrop(baseRenderer *baseRenderer, stmtDrop *stmtDrop) error
	renderTruncate(baseRenderer *baseRenderer, stmtTruncate *stmtTruncate) error
	// DML
	renderDelete(baseRenderer *baseRenderer, stmtDelete *stmtDelete) error
	renderInsert(baseRenderer *baseRenderer, stmtInsert *stmtInsert) error
	renderSelect(baseRenderer *baseRenderer, stmtSelect *stmtSelect) error
	renderUpdate(baseRenderer *baseRenderer, stmtUpdate *stmtUpdate) error
}
type transformer interface {
	elementTransformer
	componentTransformer
	statementTransformer
}
type elementTransformer interface {
	transformComparison() error
	transformFunction() error
}
type componentTransformer interface {
}
type statementTransformer interface {
	// DDL
	transformAlter(baseTransformer *baseTransformer, stmtAlter *stmtAlter) error
	transformComment(baseTransformer *baseTransformer, stmtComment *stmtComment) error
	transformCreate(baseTransformer *baseTransformer, stmtCreate *stmtCreate) error
	transformDrop(baseTransformer *baseTransformer, stmtDrop *stmtDrop) error
	transformTruncate(baseTransformer *baseTransformer, stmtTruncate *stmtTruncate) error
	// DML
	transformDelete(baseTransformer *baseTransformer, stmtDelete *stmtDelete) error
	transformInsert(baseTransformer *baseTransformer, stmtInsert *stmtInsert) error
	transformSelect(baseTransformer *baseTransformer, stmtSelect *stmtSelect) error
	transformUpdate(baseTransformer *baseTransformer, stmtUpdate *stmtUpdate) error
}
type validator interface {
	elementValidator
	componentValidator
	statementValidator
}
type elementValidator interface {
	validateAlias(value string) error
	validateArray(value int) error
	validateComparison(value transformComparison) error
	validateFunction(value transformFunction) error
	validateName(value string) error
	validateLiteral(value any) error
	validateOperator(value any) error
	validateService(value any) error
	validateSubquery() error
	validateValue(value any) error
}
type componentValidator interface {
	validateColumns(columns []markSourceable) error
	validateCommand(command managementService) error
	validateEntity(entity SourceBase) error
	validateFields(fields []markExpressable) error
	validateFrom(from SourceBase) error
	validateGroupBy(groups []markGroupable) error
	validateHaving(having ExpressionBase) error
	validateInTo(inTo SourceBase) error
	validateIsData(data string) error
	validateJoin(joins []*clauseJoin) error
	validateOn(on SourceBase) error
	validateOnFrom(onTable, onColumn SourceBase) error
	validateOnTo(onTo SourceBase) error
	validateOrderBy(orders []markOrderable) error
	validatePagination(pagination *clausePagination) error
	validateRenameColumn(columnRename *columnRename) error
	validateRenameConstraint(constraintRename *constraintRename) error
	validateRenameTo(name string) error
	validateReturning(returnings *clauseReturning) error
	validateSet(sets []*clauseSet) error
	validateSource(source statement) error
	validateTable(table *TableSource) error
	validateUnions(unions []*clauseUnions) error
	validateValues(values *clauseValues) error
	validateWhere(where ExpressionBase) error
	validateWith(withs []*clauseWith) error
}
type statementValidator interface {
	// DDL
	validateAlter(baseValidator *baseValidator, stmtAlter *stmtAlter) error
	validateComment(baseValidator *baseValidator, stmtComment *stmtComment) error
	validateCreate(baseValidator *baseValidator, stmtCreate *stmtCreate) error
	validateDrop(baseValidator *baseValidator, stmtDrop *stmtDrop) error
	validateTruncate(baseValidator *baseValidator, stmtTruncate *stmtTruncate) error
	// DML
	validateDelete(baseValidator *baseValidator, stmtDelete *stmtDelete) error
	validateInsert(baseValidator *baseValidator, stmtInsert *stmtInsert) error
	validateSelect(baseValidator *baseValidator, stmtSelect *stmtSelect) error
	validateUpdate(baseValidator *baseValidator, stmtUpdate *stmtUpdate) error
}

// Приватные структуры
type processor struct {
	renderer    renderer
	transformer transformer
	validator   validator
}

// Приватные методы
func (processor processor) createRenderer(config *config, contexter *contexter, strateger strateger) *baseRenderer {
	return newRenderer(config, contexter, strateger)
}
func (processor processor) createTransformer(config *config, contexter *contexter, strateger strateger) *baseTransformer {
	return newTransformer(config, contexter, strateger)
}
func (processor processor) createValidator(config *config, contexter *contexter, strateger strateger) *baseValidator {
	return newValidator(config, contexter, strateger)
}
