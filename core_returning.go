package uast

// Приватные интерфейсы
type markReturnable interface {
	ExpressionBase
	isReturnable()
}

// Приватные структуры
type clauseReturning struct {
	expression ExpressionBase
}

// Приватные методы
func (clause *clauseReturning) clone() ExpressionBase {
	copy := *clause
	return &copy
}
func (clause *clauseReturning) isExpressionBase() {}
func (clause *clauseReturning) isReturnable()     {}
func (clause *clauseReturning) render(baseRenderer *baseRenderer) error {
	return clause.expression.render(baseRenderer)
}
func (clause *clauseReturning) validate(baseValidator *baseValidator) error {
	return clause.expression.validate(baseValidator)
}
