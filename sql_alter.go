package uast

// !!! Внимание, находится в стадии разработки

// Публичные конструкторы
func NewAlter(entity SourceBase) *stmtAlter {
	return &stmtAlter{
		command: uastManagementAlter,
		entity:  entity,
	}
}

// Публичные методы
func (stmt *stmtAlter) AddColumns(columns ...markSourceable) *stmtAlter {
	stmt.addColumns = columns
	for _, column := range columns {
		if col, ok := column.(registerStatement); ok {
			col.register(stmt)
		}
	}
	return stmt
}
func (stmt *stmtAlter) AddConstraints(constraints ...Constraint) *stmtAlter {
	stmt.addConstraints = append(stmt.addConstraints, constraints...)
	return stmt
}
func (stmt *stmtAlter) DropColumns(columns ...string) *stmtAlter {
	stmt.dropColumns = columns
	return stmt
}
func (stmt *stmtAlter) DropConstraints(names ...string) *stmtAlter {
	stmt.dropConstraints = append(stmt.dropConstraints, names...)
	return stmt
}

// Приватные структуры
type stmtAlter struct {
	command         managementService
	addColumns      []markSourceable
	addConstraints  []Constraint
	dropConstraints []string
	dropColumns     []string
	entity          SourceBase
	ifExists        bool
	ifNotExists     bool
	on              SourceBase
}

// Приватные методы
func (stmt *stmtAlter) clone() statement {
	copy := *stmt
	return &copy
}
func (stmt *stmtAlter) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderAlter(baseRenderer, stmt)
}
func (stmt *stmtAlter) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformAlter(baseTransformer, stmt)
}
func (stmt *stmtAlter) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateAlter(baseValidator, stmt)
}
