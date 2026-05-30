package uast

// Публичные конструкторы
func NewDrop(entity EntityTarget) *stmtDrop {
	return &stmtDrop{
		command: uastManagementDrop,
		entity:  entity,
	}
}

// Публичные методы
func (stmt *stmtDrop) Cascade() *stmtDrop {
	stmt.cascade = true
	return stmt
}
func (stmt *stmtDrop) IfExists() *stmtDrop {
	stmt.ifExists = true
	return stmt
}

// Приватные структуры
type stmtDrop struct {
	command  managementService
	cascade  bool
	entity   EntityTarget
	ifExists bool
}

// Приватные методы
func (stmt *stmtDrop) clone() statement {
	copy := *stmt
	return &copy
}
func (stmt *stmtDrop) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderDrop(baseRenderer, stmt)
}
func (stmt *stmtDrop) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformDrop(baseTransformer, stmt)
}
func (stmt *stmtDrop) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateDrop(baseValidator, stmt)
}
