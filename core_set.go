package uast

// Публичные функции
func Assign[T typeScalar](field *exprField[T], value ExpressionSafe[T]) *clauseSet {
	return &clauseSet{
		field: field,
		value: value,
	}
}

// Приватные структуры
type clauseSet struct {
	field markExpressable
	value ExpressionBase
}

// Приватные методы
func (clause *clauseSet) clone() *clauseSet {
	copy := *clause
	return &copy
}
func (clause *clauseSet) render(baseRenderer *baseRenderer) error {
	if err := clause.field.render(baseRenderer); err != nil {
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
	if clause.field == nil || clause.value == nil {
		return ErrInvalidStatementSet
	}
	if err := clause.field.validate(baseValidator); err != nil {
		return err
	}
	if err := clause.value.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
