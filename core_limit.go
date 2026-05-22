package uast

// Приватные структуры
type clauseLimit struct {
	value int
}

// Приватные методы
func (clause *clauseLimit) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderValue(clause.value)
	return nil
}
func (clause *clauseLimit) validate(baseValidator *baseValidator) error {
	if clause.value < 0 {
		return ErrInvalidStatementLimit
	}
	if clause.value > uastCountMaxLimit {
		return ErrExcessMaxLimit
	}
	return nil
}
