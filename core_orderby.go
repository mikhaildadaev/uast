package uast

// Приватные интерфейсы
type markOrderable interface {
	ExpressionBase
	isOrderable()
}

// Приватные структуры
type clauseOrderBy struct {
	direction  bool
	expression ExpressionBase
}

// Приватные методы
func (clauseOrderBy *clauseOrderBy) isExpressionBase() {}
func (clauseOrderBy *clauseOrderBy) isOrderable()      {}
func (clauseOrderBy *clauseOrderBy) render(baseRenderer *baseRenderer) error {
	if err := clauseOrderBy.expression.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if clauseOrderBy.direction {
		baseRenderer.renderOperator(uastOrderDesc)
	} else {
		baseRenderer.renderOperator(uastOrderAsc)
	}
	return nil
}
func (clauseOrderBy *clauseOrderBy) validate(baseValidator *baseValidator) error {
	if clauseOrderBy == nil {
		return ErrInvalidStatementOrderBy
	}
	return nil
}
