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
func (stmt *stmtTruncate) IsCascade() *stmtTruncate {
	stmt.isCascade = true
	return stmt
}
func (stmt *stmtTruncate) IsRestartIdentity() *stmtTruncate {
	stmt.isRestartIdentity = true
	return stmt
}

// Приватные структуры
type stmtTruncate struct {
	command           managementService
	table             *TableSource
	isCascade         bool
	isRestartIdentity bool
}

// Приватные методы
func (stmt *stmtTruncate) clone() statement {
	copy := *stmt
	if stmt.table != nil {
		copy.table = stmt.table.clone().(*TableSource)
	}
	return &copy
}
func (stmt *stmtTruncate) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderTruncate(baseRenderer, stmt)
}
func (stmt *stmtTruncate) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformTruncate(baseTransformer, stmt)
}
func (stmt *stmtTruncate) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateTruncate(baseValidator, stmt)
}
