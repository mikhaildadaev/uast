package uast

// Приватные интерфейсы
type processor interface {
	renderer
	transformer
	validator
}
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
	renderColumn(columns []markColumnable) error
	renderDistinct(distinct bool) error
	renderField(fields []markFieldable) error
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
	validateColumn(columns []markColumnable) error
	validateCommand(command managementService) error
	validateField(fields []markFieldable) error
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
