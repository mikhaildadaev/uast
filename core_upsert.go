package uast

// Приватные структуры
type clauseUpsert struct {
	pairs   []*clausePair
	service managementService
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
