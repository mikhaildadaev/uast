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
	renderCascade(cascade bool) error
	renderCommand(command managementService) error
	renderColumns(columns []markExpressable) error
	renderDistinct(distinct bool) error
	renderEntity(entity SourceBase, ifExists bool, ifNotExists bool) error
	renderFields(fields []markExpressable) error
	renderFrom(from SourceBase) error
	renderGroupBy(groups []markGroupable) error
	renderHaving(having ExpressionBase) error
	renderIfExists(ifExists bool) error
	renderIfNotExists(ifNotExists bool) error
	renderInto(into SourceBase) error
	renderIsComment(comment string) error
	renderJoin(joins []*clauseJoin) error
	renderOnColumn(column markExpressable) error
	renderOnFields(table *TableSource, fields []markExpressable) error
	renderOnTable(table *TableSource) error
	renderOnto(onto SourceBase) error
	renderOrderBy(orders []markOrderable) error
	renderPagination(pagination *clausePagination) error
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
	transformFields(fields []markExpressable, columns *[]string) error
}
type statementTransformer interface {
	// DDL
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
	validateColumns(columns []markExpressable) error
	validateCommand(command managementService) error
	validateEntity(entity SourceBase) error
	validateFields(fields []markExpressable) error
	validateFrom(from SourceBase) error
	validateGroupBy(groups []markGroupable) error
	validateHaving(having ExpressionBase) error
	validateInto(into SourceBase) error
	validateIsComment(comment string) error
	validateJoin(joins []*clauseJoin) error
	validateOnColumn(column markExpressable) error
	validateOnTable(table *TableSource) error
	validateOnto(onto SourceBase) error
	validateOrderBy(orders []markOrderable) error
	validatePagination(pagination *clausePagination) error
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
