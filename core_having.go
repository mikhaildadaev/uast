package uast

// Приватные структуры
type clauseHaving struct {
	expression ExpressionBase
}

// Приватные методы
func (clauseHaving *clauseHaving) isExpressionBase() {}
func (clauseHaving *clauseHaving) render(baseRenderer *baseRenderer) error {
	return clauseHaving.expression.render(baseRenderer)
}
func (clauseHaving *clauseHaving) validate(baseValidator *baseValidator) error {
	return clauseHaving.expression.validate(baseValidator)
}
