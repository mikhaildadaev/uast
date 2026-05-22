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
type clauseUpsert struct {
	pairs   []*clausePair
	service managementService
}
type clauseValues struct {
	pairs  []*clausePair
	upsert *clauseUpsert
}

// Приватные методы
func (clauseUpsert *clauseUpsert) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(clauseUpsert.service)
	for i, pair := range clauseUpsert.pairs {
		if i > 0 {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := pair.column.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSpaceEqualSpace)
		if err := pair.value.render(baseRenderer); err != nil {
			return err
		}
	}
	return nil
}
func (clauseUpsert *clauseUpsert) validate(baseValidator *baseValidator) error {
	if clauseUpsert.pairs == nil {
		return ErrInvalidStatementSet
	}
	for _, pair := range clauseUpsert.pairs {
		if err := pair.column.validate(baseValidator); err != nil {
			return err
		}
		if err := pair.value.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
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
