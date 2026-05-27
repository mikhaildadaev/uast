package uast

// Приватные структуры
type clausePagination struct {
	limit  int
	offset int
}

// Приватные методы
func (clause *clausePagination) clone() *clausePagination {
	copy := *clause
	return &copy
}
func (clause *clausePagination) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderPagination(clause)
	return nil
}
func (clause *clausePagination) validate(baseValidator *baseValidator) error {
	if clause == nil {
		return nil
	}
	if clause.limit > uastCountMaxLimit {
		return ErrExcessMaxLimit
	}
	if clause.limit < 0 {
		return ErrInvalidStatementLimit
	}
	if clause.offset < 0 {
		return ErrInvalidStatementOffset
	}
	return nil
}
