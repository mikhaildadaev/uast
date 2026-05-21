package uast

// Публичные функции
func Assign[T typeScalar](column *exprColumn[T], value ExpressionSafe[T]) *clauseSet {
	return &clauseSet{
		column: column,
		value:  value,
	}
}

// Приватные структуры
type clauseSet struct {
	column markExpressable
	value  ExpressionBase
}

// Приватные методы
func (clauseSet *clauseSet) render(baseRenderer *baseRenderer) error {
	if err := clauseSet.column.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderOperator(uastComparisonEqual)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := clauseSet.value.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clauseSet *clauseSet) validate(baseValidator *baseValidator) error {
	if clauseSet.column == nil || clauseSet.value == nil {
		return ErrInvalidStatementSet
	}
	if err := clauseSet.column.validate(baseValidator); err != nil {
		return err
	}
	if err := clauseSet.value.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
