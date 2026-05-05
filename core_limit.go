package uast

// Приватные структуры
type clauseLimit struct {
	value int
}

// Приватные методы
func (clauseLimit *clauseLimit) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderValue(clauseLimit.value)
	return nil
}
func (clauseLimit *clauseLimit) validate(baseValidator *baseValidator) error {
	if clauseLimit.value < 0 {
		return ErrInvalidStatementLimit
	}
	if clauseLimit.value > uastCountMaxLimit {
		return ErrExcessMaxLimit
	}
	return nil
}
