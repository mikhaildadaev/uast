package uast

// Публичные функции
func Assign(column markColumnable, value ExpressionBase) *clausePair {
	return &clausePair{
		column: column,
		value:  value,
	}
}

// Приватные структуры
type clausePair struct {
	column markColumnable
	value  ExpressionBase
}

// Приватные методы
func (clausePair *clausePair) render(baseRenderer *baseRenderer) error {
	if err := clausePair.column.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderOperator(uastComparisonEqual)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := clausePair.value.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clausePair *clausePair) validate(baseValidator *baseValidator) error {
	if clausePair.column == nil || clausePair.value == nil {
		return ErrInvalidStatementSet
	}
	if err := clausePair.column.validate(baseValidator); err != nil {
		return err
	}
	if err := clausePair.value.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
