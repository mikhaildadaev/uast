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
	renderCommand(command managementService) error
	renderColumn(columns []markExpressable) error
	renderDistinct(distinct bool) error
	renderField(fields []markExpressable) error
	renderFrom(from SourceBase) error
	renderGroupBy(groups []markGroupable) error
	renderHaving(having ExpressionBase) error
	renderInto(into SourceBase) error
	renderJoin(joins []*clauseJoin) error
	renderLimit(limit *clauseLimit) error
	renderOffset(offset *clauseOffset) error
	renderOnto(onto SourceBase) error
	renderOrderBy(orders []markOrderable) error
	renderReturning(returnings []markReturnable) error
	renderSet(sets []*clauseSet) error
	renderSource(source statement) error
	renderUnions(unions []*clauseUnions) error
	renderValues(values [][]ExpressionBase) error
	renderWhere(where ExpressionBase) error
	renderWith(withs []*clauseWith) error
}
type statementRenderer interface {
	renderDelete(baseRenderer *baseRenderer, stmt *stmtDelete) error
	renderInsert(baseRenderer *baseRenderer, stmt *stmtInsert) error
	renderSelect(baseRenderer *baseRenderer, stmt *stmtSelect) error
	renderUpdate(baseRenderer *baseRenderer, stmt *stmtUpdate) error
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
	validateColumn(columns []markExpressable) error
	validateCommand(command managementService) error
	validateField(fields []markExpressable) error
	validateFrom(from SourceBase) error
	validateGroupBy(groups []markGroupable) error
	validateHaving(having ExpressionBase) error
	validateInto(into SourceBase) error
	validateJoin(joins []*clauseJoin) error
	validateLimit(limit *clauseLimit) error
	validateOffset(offset *clauseOffset) error
	validateOnto(onto SourceBase) error
	validateOrderBy(orders []markOrderable) error
	validateReturning(returnings []markReturnable) error
	validateSet(sets []*clauseSet) error
	validateSource(source statement) error
	validateUnions(unions []*clauseUnions) error
	validateValues(values [][]ExpressionBase) error
	validateWhere(where ExpressionBase) error
	validateWith(withs []*clauseWith) error
}
type statementValidator interface {
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
