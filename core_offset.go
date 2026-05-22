package uast

// Приватные структуры
type clauseOffset struct {
	value int
}

// Приватные методы
func (clause *clauseOffset) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderValue(clause.value)
	return nil
}
func (clause *clauseOffset) validate(baseValidator *baseValidator) error {
	if clause.value < 0 {
		return ErrInvalidStatementOffset
	}
	return nil
}
