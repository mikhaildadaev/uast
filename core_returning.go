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
func (clauseReturning *clauseReturning) isExpressionBase() {}
func (clauseReturning *clauseReturning) isReturnable()     {}
func (clauseReturning *clauseReturning) render(baseRenderer *baseRenderer) error {
	return clauseReturning.expression.render(baseRenderer)
}
func (clauseReturning *clauseReturning) validate(baseValidator *baseValidator) error {
	return clauseReturning.expression.validate(baseValidator)
}
