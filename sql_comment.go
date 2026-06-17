package uast

// !!!Внимание, находится в стадии разработки

// Публичные конструкторы
func NewComment(on SourceBase) *stmtComment {
	return &stmtComment{
		command: uastManagementComment,
		on:      on,
	}
}

// Публичные методы
func (stmt *stmtComment) Is(comment string) *stmtComment {
	stmt.comment = comment
	return stmt
}

// Приватные структуры
type stmtComment struct {
	command managementService
	comment string
	on      SourceBase
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
