package uast

// !!!Внимание, находится в стадии разработки

// Публичные конструкторы
func NewComment(onTable SourceBase) *stmtComment {
	return &stmtComment{
		command: uastManagementComment,
		onTable: onTable,
	}
}

// Публичные методы
func (stmt *stmtComment) OnColumn(onColumn SourceBase) *stmtComment {
	stmt.onColumn = onColumn
	return stmt
}
func (stmt *stmtComment) Is(comment string) *stmtComment {
	stmt.comment = comment
	return stmt
}

// Приватные структуры
type stmtComment struct {
	command  managementService
	comment  string
	onColumn SourceBase
	onTable  SourceBase
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
