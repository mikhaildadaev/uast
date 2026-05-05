package uast

// Приватные структуры
type clauseOffset struct {
	value int
}

// Приватные методы
func (clauseOffset *clauseOffset) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderValue(clauseOffset.value)
	return nil
}
func (clauseOffset *clauseOffset) validate(baseValidator *baseValidator) error {
	if clauseOffset.value < 0 {
		return ErrInvalidStatementOffset
	}
	return nil
}
