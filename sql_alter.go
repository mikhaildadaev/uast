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
func (stmt *stmtAlter) AddConstraints(constraints ...Constraint) *stmtAlter {
	for _, constraint := range constraints {
		switch data := constraint.(type) {
		case *ConstraintCheck:
			stmt.constraintChecksAdd = append(stmt.constraintChecksAdd, data)
		case *ConstraintForeign:
			stmt.constraintForeignsAdd = append(stmt.constraintForeignsAdd, data)
		case *ConstraintPrimary:
			stmt.constraintPrimarysAdd = append(stmt.constraintPrimarysAdd, data)
		case *ConstraintUnique:
			stmt.constraintUniquesAdd = append(stmt.constraintUniquesAdd, data)
		}
	}
	return stmt
}
func (stmt *stmtAlter) DropConstraints(constraints ...Constraint) *stmtAlter {
	for _, constraint := range constraints {
		switch data := constraint.(type) {
		case *ConstraintCheck:
			stmt.constraintChecksDrop = append(stmt.constraintChecksDrop, data)
		case *ConstraintForeign:
			stmt.constraintForeignsDrop = append(stmt.constraintForeignsDrop, data)
		case *ConstraintPrimary:
			stmt.constraintPrimarysDrop = append(stmt.constraintPrimarysDrop, data)
		case *ConstraintUnique:
			stmt.constraintUniquesDrop = append(stmt.constraintUniquesDrop, data)
		}
	}
	return stmt
}

// Приватные структуры
type stmtAlter struct {
	command                managementService
	columns                []markSourceable
	constraintChecksAdd    []*ConstraintCheck
	constraintChecksDrop   []*ConstraintCheck
	constraintForeignsAdd  []*ConstraintForeign
	constraintForeignsDrop []*ConstraintForeign
	constraintPrimarysAdd  []*ConstraintPrimary
	constraintPrimarysDrop []*ConstraintPrimary
	constraintUniquesAdd   []*ConstraintUnique
	constraintUniquesDrop  []*ConstraintUnique
	entity                 SourceBase
	ifExists               bool
	ifNotExists            bool
	on                     SourceBase
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
