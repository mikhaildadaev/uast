package uast

// Приватные структуры
type clausePagination struct {
	process       bool
	serviceLimit  managementService
	serviceOffset managementService
	valueLimit    int
	valueOffset   int
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
	baseValidator.validatePagination(clause)
	return nil
}
