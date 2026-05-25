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
func (clause *clauseSet) clone() *clauseSet {
	copy := *clause
	return &copy
}
func (clause *clauseSet) render(baseRenderer *baseRenderer) error {
	if err := clause.column.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderOperator(uastComparisonEqual)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := clause.value.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clause *clauseSet) validate(baseValidator *baseValidator) error {
	if clause.column == nil || clause.value == nil {
		return ErrInvalidStatementSet
	}
	if err := clause.column.validate(baseValidator); err != nil {
		return err
	}
	if err := clause.value.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
