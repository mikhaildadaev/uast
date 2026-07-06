package uast

// Приватные интерфейсы
type markReturnable interface {
	ExpressionBase
	isReturnable()
}

// Приватные структуры
type clauseReturning struct {
	expressions      []markReturnable
	serviceReturning modifierService
}

// Приватные методы
func (clause *clauseReturning) clone() *clauseReturning {
	copy := *clause
	copy.expressions = make([]markReturnable, len(clause.expressions))
	for i, expr := range clause.expressions {
		copy.expressions[i] = expr
	}
	return &copy
}
func (clause *clauseReturning) isExpressionBase() {}
func (clause *clauseReturning) isReturnable()     {}
func (clause *clauseReturning) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(clause.serviceReturning)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	for i, expressions := range clause.expressions {
		if i > 0 {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
		if err := expressions.render(baseRenderer); err != nil {
			return err
		}
	}
	return nil
}
func (clause *clauseReturning) validate(baseValidator *baseValidator) error {
	if clause.expressions == nil {
		return nil
	}
	for _, expression := range clause.expressions {
		if err := expression.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
