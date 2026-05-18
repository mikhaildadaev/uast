package uast

// Публичные методы
func (stmtInsert *stmtInsert) Into(into SourceBase) *stmtInsert {
	stmtInsert.into = into
	return stmtInsert
}
func (stmtInsert *stmtInsert) Returning(returnings ...markReturnable) *stmtInsert {
	stmtInsert.returning = returnings
	return stmtInsert
}
func (stmtInsert *stmtInsert) Source(source *stmtSelect) *stmtInsert {
	stmtInsert.source = source
	stmtInsert.column = source.field
	return stmtInsert
}
func (stmtInsert *stmtInsert) Values(pairs ...*clausePair) *stmtInsert {
	columns := make([]markExpressable, len(pairs))
	values := make([]ExpressionBase, len(pairs))
	for i, pair := range pairs {
		columns[i] = pair.column
		values[i] = pair.value
	}
	stmtInsert.column = columns
	stmtInsert.values = [][]ExpressionBase{values}
	return stmtInsert
}
func (stmtInsert *stmtInsert) With(with ...*clauseWith) *stmtInsert {
	stmtInsert.with = with
	return stmtInsert
}

// Приватные структуры
type stmtInsert struct {
	command   managementService
	column    []markExpressable
	into      SourceBase
	source    statement
	returning []markReturnable
	values    [][]ExpressionBase
	with      []*clauseWith
}

// Приватные методы
func (stmtInsert *stmtInsert) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderInsert(baseRenderer, stmtInsert)
}
func (stmtInsert *stmtInsert) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformInsert(baseTransformer, stmtInsert)
}
func (stmtInsert *stmtInsert) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateInsert(baseValidator, stmtInsert)
}
