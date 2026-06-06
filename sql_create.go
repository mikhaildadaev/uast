package uast

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
	for _, column := range columns {
		if col, ok := column.(registerStatement); ok {
			col.register(stmt)
		}
	}
	return stmt
}
func (stmt *stmtCreate) Constraints(constraints ...any) *stmtCreate {
	for _, constraint := range constraints {
		switch data := constraint.(type) {
		case *ConstraintCheck:
			stmt.constraintChecks = append(stmt.constraintChecks, data)
		case *ConstraintForeign:
			stmt.constraintForeigns = append(stmt.constraintForeigns, data)
		case *ConstraintPrimary:
			stmt.constraintPrimarys = append(stmt.constraintPrimarys, data)
		case *ConstraintUnique:
			stmt.constraintUniques = append(stmt.constraintUniques, data)
		}
	}
	return stmt
}
func (stmt *stmtCreate) IfNotExists() *stmtCreate {
	stmt.ifNotExists = true
	return stmt
}
func (stmt *stmtCreate) IsReplace() *stmtCreate {
	stmt.isReplace = true
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

// Приватные структуры
type stmtCreate struct {
	command            managementService
	columns            []markSourceable
	constraintChecks   []*ConstraintCheck
	constraintForeigns []*ConstraintForeign
	constraintPrimarys []*ConstraintPrimary
	constraintUniques  []*ConstraintUnique
	entity             SourceBase
	ifNotExists        bool
	isReplace          bool
	on                 SourceBase
	source             statement
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
