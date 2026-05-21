package uast

// Приватные интерфейсы
type markGroupable interface {
	ExpressionBase
	isGroupable()
}

// Приватные структуры
type clauseGroupBy struct {
	expression ExpressionBase
}

// Приватные методы
func (clauseGroupBy *clauseGroupBy) isExpressionBase() {}
func (clauseGroupBy *clauseGroupBy) isGroupable()      {}
func (clauseGroupBy *clauseGroupBy) render(baseRenderer *baseRenderer) error {
	if err := clauseGroupBy.expression.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clauseGroupBy *clauseGroupBy) validate(baseValidator *baseValidator) error {
	if clauseGroupBy.expression == nil {
		return ErrInvalidStatementGroupBy
	}
	return nil
}
