package uast

// Приватные структуры
type clauseHaving struct {
	expression ExpressionBase
}

// Приватные методы
func (clause *clauseHaving) clone() *clauseHaving {
	copy := *clause
	return &copy
}
func (clause *clauseHaving) isExpressionBase() {}
func (clause *clauseHaving) render(baseRenderer *baseRenderer) error {
	return clause.expression.render(baseRenderer)
}
func (clause *clauseHaving) validate(baseValidator *baseValidator) error {
	return clause.expression.validate(baseValidator)
}
