package uast

// Публичные функции
func Pair[T typeScalar](column *exprColumn[T], value ExpressionSafe[T]) *clauseValues {
	return &clauseValues{
		column: &exprPair[T]{
			name: column.name},
		value: value,
	}
}

// Приватные структуры
type clauseValues struct {
	column markExpressable
	value  ExpressionBase
}

// Приватные методы
func (clauseValues *clauseValues) render(baseRenderer *baseRenderer) error {
	if err := clauseValues.value.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clauseValues *clauseValues) validate(baseValidator *baseValidator) error {
	if clauseValues.column == nil || clauseValues.value == nil {
		return ErrInvalidStatementValues
	}
	if err := clauseValues.column.validate(baseValidator); err != nil {
		return err
	}
	if err := clauseValues.value.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
