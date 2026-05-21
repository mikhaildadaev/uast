package uast

// Приватные структуры
type clauseWhere struct {
	expression ExpressionBase
}

// Приватные методы
func (clauseWhere *clauseWhere) isExpressionBase() {}
func (clauseWhere *clauseWhere) isPredicable()     {}
func (clauseWhere *clauseWhere) render(baseRenderer *baseRenderer) error {
	return clauseWhere.expression.render(baseRenderer)
}
func (clauseWhere *clauseWhere) validate(baseValidator *baseValidator) error {
	return clauseWhere.expression.validate(baseValidator)
}
