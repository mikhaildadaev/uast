package uast

// !!! Внимание, находится в стадии разработки

// Публичные конструкторы
func NewCreate(entity SourceBase) *stmtCreate {
	return &stmtCreate{
		command: uastManagementCreate,
		entity:  entity,
	}
}

// Публичные методы
func (stmt *stmtCreate) Columns(columns ...markSourceable) *stmtCreate {
	stmt.columns = columns
	return stmt
}
func (stmt *stmtCreate) IfNotExists() *stmtCreate {
	stmt.ifNotExists = true
	return stmt
}
func (stmt *stmtCreate) Replace() *stmtCreate {
	stmt.replace = true
	return stmt
}
func (stmt *stmtCreate) Source(source statement) *stmtCreate {
	stmt.source = source
	return stmt
}
func (stmt *stmtCreate) On(on SourceBase) *stmtCreate {
	stmt.on = on
	return stmt
}
func (stmt *stmtCreate) Unique() *stmtCreate {
	stmt.unique = true
	return stmt
}

// Приватные структуры
type stmtCreate struct {
	command     managementService
	columns     []markSourceable
	entity      SourceBase
	ifNotExists bool
	replace     bool
	on          SourceBase
	source      statement
	unique      bool
}

// Приватные методы
func (stmt *stmtCreate) clone() statement {
	copy := *stmt
	return &copy
}
func (stmt *stmtCreate) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderCreate(baseRenderer, stmt)
}
func (stmt *stmtCreate) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformCreate(baseTransformer, stmt)
}
func (stmt *stmtCreate) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateCreate(baseValidator, stmt)
}
