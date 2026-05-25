package uast

// Публичные конструкторы
func NewComment(comment string) *stmtComment {
	return &stmtComment{
		command: uastManagementComment,
		comment: comment,
	}
}

// Публичные методы
func (stmt *stmtComment) OnColumn(column markExpressable) *stmtComment {
	stmt.column = column
	return stmt
}
func (stmt *stmtComment) OnTable(table *TableSource) *stmtComment {
	stmt.table = table
	return stmt
}

// Приватные структуры
type stmtComment struct {
	command managementService
	comment string
	column  markExpressable
	table   *TableSource
}

// Приватные методы
func (stmt *stmtComment) clone() statement {
	copy := *stmt
	return &copy
}
func (stmt *stmtComment) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderComment(baseRenderer, stmt)
}
func (stmt *stmtComment) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformComment(baseTransformer, stmt)
}
func (stmt *stmtComment) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateComment(baseValidator, stmt)
}
