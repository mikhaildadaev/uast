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
func (clause *clauseGroupBy) isExpressionBase() {}
func (clause *clauseGroupBy) isGroupable()      {}
func (clause *clauseGroupBy) render(baseRenderer *baseRenderer) error {
	if err := clause.expression.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clause *clauseGroupBy) validate(baseValidator *baseValidator) error {
	if clause.expression == nil {
		return ErrInvalidStatementGroupBy
	}
	return nil
}
