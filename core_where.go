package uast

// Приватные структуры
type clauseWhere struct {
	expression ExpressionBase
}

// Приватные методы
func (clause *clauseWhere) isExpressionBase() {}
func (clause *clauseWhere) isPredicable()     {}
func (clause *clauseWhere) render(baseRenderer *baseRenderer) error {
	return clause.expression.render(baseRenderer)
}
func (clause *clauseWhere) validate(baseValidator *baseValidator) error {
	return clause.expression.validate(baseValidator)
}
