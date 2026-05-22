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
func (clause *clauseOrderBy) isExpressionBase() {}
func (clause *clauseOrderBy) isOrderable()      {}
func (clause *clauseOrderBy) render(baseRenderer *baseRenderer) error {
	if err := clause.expression.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if clause.direction {
		baseRenderer.renderOperator(uastOrderDesc)
	} else {
		baseRenderer.renderOperator(uastOrderAsc)
	}
	return nil
}
func (clause *clauseOrderBy) validate(baseValidator *baseValidator) error {
	if clause == nil {
		return ErrInvalidStatementOrderBy
	}
	return nil
}
