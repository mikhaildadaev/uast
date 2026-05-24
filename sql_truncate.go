package uast

// Публичные конструкторы
func NewTruncate(table *TableSource) *stmtTruncate {
	copy := *table
	copy.withAlias = false
	return &stmtTruncate{
		command: uastManagementTruncate,
		table:   &copy,
	}
}

// Публичные методы
func (stmt *stmtTruncate) Cascade() *stmtTruncate {
	stmt.cascade = true
	return stmt
}
func (stmt *stmtTruncate) RestartIdentity() *stmtTruncate {
	stmt.restartIdentity = true
	return stmt
}

// Приватные структуры
type stmtTruncate struct {
	command         managementService
	table           *TableSource
	cascade         bool
	restartIdentity bool
}

// Приватные методы
func (stmt *stmtTruncate) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderTruncate(baseRenderer, stmt)
}
func (stmt *stmtTruncate) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformTruncate(baseTransformer, stmt)
}
func (stmt *stmtTruncate) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateTruncate(baseValidator, stmt)
}
