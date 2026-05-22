package uast

// Публичные функции
func Pair[T typeScalar](column *exprColumn[T], value ExpressionSafe[T]) *clausePair {
	return &clausePair{
		column: &exprPair[T]{
			name: column.name,
		},
		value: value,
	}
}

// Приватные структуры
type clausePair struct {
	column markExpressable
	value  ExpressionBase
}
type clauseValues struct {
	pairs  []*clausePair
	upsert *clauseUpsert
}

// Приватные методы
func (clauseValues *clauseValues) render(baseRenderer *baseRenderer) error {
	for _, pair := range clauseValues.pairs {
		if err := pair.value.render(baseRenderer); err != nil {
			return err
		}
	}
	return nil
}
func (clauseValues *clauseValues) validate(baseValidator *baseValidator) error {
	if clauseValues.pairs == nil {
		return ErrInvalidStatementValues
	}
	for _, pair := range clauseValues.pairs {
		if pair.column == nil || pair.value == nil {
			return ErrInvalidStatementValues
		}
		if err := pair.column.validate(baseValidator); err != nil {
			return err
		}
		if err := pair.value.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
