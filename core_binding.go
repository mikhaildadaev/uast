package uast

// Публичные функции
func Assign(column markExpressable, value ExpressionBase) *clauseAssign {
	return &clauseAssign{
		column: column,
		value:  value,
	}
}
func Pair(column markExpressable, value ExpressionBase) *clausePair {
	return &clausePair{
		column: column,
		value:  value,
	}
}

// Приватные структуры
type clauseAssign struct {
	column markExpressable
	value  ExpressionBase
}
type clausePair struct {
	column markExpressable
	value  ExpressionBase
}

// Приватные методы
func (clauseAssign *clauseAssign) render(baseRenderer *baseRenderer) error {
	if err := clauseAssign.column.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderOperator(uastComparisonEqual)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := clauseAssign.value.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clauseAssign *clauseAssign) validate(baseValidator *baseValidator) error {
	if clauseAssign.column == nil || clauseAssign.value == nil {
		return ErrInvalidStatementSet
	}
	if err := clauseAssign.column.validate(baseValidator); err != nil {
		return err
	}
	if err := clauseAssign.value.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
func (clausePair *clausePair) render(baseRenderer *baseRenderer) error {
	if err := clausePair.value.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clausePair *clausePair) validate(baseValidator *baseValidator) error {
	if clausePair.column == nil || clausePair.value == nil {
		return ErrInvalidStatementValues
	}
	if err := clausePair.column.validate(baseValidator); err != nil {
		return err
	}
	if err := clausePair.value.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
