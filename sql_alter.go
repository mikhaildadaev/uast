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
func (stmt *stmtAlter) AddConstraints(constraints ...ConstraintBase) *stmtAlter {
	stmt.addConstraints = append(stmt.addConstraints, constraints...)
	return stmt
}
func (stmt *stmtAlter) DropColumns(columns ...markSourceable) *stmtAlter {
	stmt.dropColumns = append(stmt.dropColumns, columns...)
	return stmt
}
func (stmt *stmtAlter) DropConstraints(constraints ...ConstraintBase) *stmtAlter {
	stmt.dropConstraints = append(stmt.dropConstraints, constraints...)
	return stmt
}
func (stmt *stmtAlter) RenameColumn(column SourceBase, name string) *stmtAlter {
	stmt.renameColumn = &columnRename{
		column: column,
		name:   name,
	}
	return stmt
}
func (stmt *stmtAlter) RenameConstraint(constraint ConstraintBase, name string) *stmtAlter {
	stmt.renameConstraint = &constraintRename{
		constraint: constraint,
		name:       name,
	}
	return stmt
}
func (stmt *stmtAlter) RenameTo(name string) *stmtAlter {
	stmt.renameTo = name
	return stmt
}
func (stmt *stmtAlter) SetColumns(oprations ...columnModifiable) *stmtAlter {
	stmt.setColumns = append(stmt.setColumns, oprations...)
	return stmt
}

// Приватные структуры
type stmtAlter struct {
	command          managementService
	addColumns       []markSourceable
	addConstraints   []ConstraintBase
	dropColumns      []markSourceable
	dropConstraints  []ConstraintBase
	entity           SourceBase
	ifExists         bool
	ifNotExists      bool
	on               SourceBase
	renameColumn     *columnRename
	renameConstraint *constraintRename
	renameTo         string
	setColumns       []columnModifiable
}

// Приватные методы
func (stmt *stmtAlter) clone() statement {
	copy := *stmt
	if len(stmt.addColumns) > 0 {
		copy.addColumns = make([]markSourceable, len(stmt.addColumns))
		for i, col := range stmt.addColumns {
			copy.addColumns[i] = col.clone().(markSourceable)
		}
	}
	if len(stmt.addConstraints) > 0 {
		copy.addConstraints = make([]ConstraintBase, len(stmt.addConstraints))
		for i, constraint := range stmt.addConstraints {
			copy.addConstraints[i] = constraint.clone()
		}
	}
	if len(stmt.dropColumns) > 0 {
		copy.dropColumns = make([]markSourceable, len(stmt.dropColumns))
		for i, col := range stmt.dropColumns {
			copy.dropColumns[i] = col.clone().(markSourceable)
		}
	}
	if len(stmt.dropConstraints) > 0 {
		copy.dropConstraints = make([]ConstraintBase, len(stmt.dropConstraints))
		for i, constraint := range stmt.dropConstraints {
			copy.dropConstraints[i] = constraint.clone()
		}
	}
	if stmt.entity != nil {
		copy.entity = stmt.entity.clone()
	}
	if stmt.on != nil {
		copy.on = stmt.on.clone()
	}
	if stmt.renameColumn != nil {
		copy.renameColumn = &columnRename{
			column: stmt.renameColumn.column.clone(),
			name:   stmt.renameColumn.name,
		}
	}
	if stmt.renameConstraint != nil {
		copy.renameConstraint = &constraintRename{
			constraint: stmt.renameConstraint.constraint.clone(),
			name:       stmt.renameConstraint.name,
		}
	}
	if len(stmt.setColumns) > 0 {
		copy.setColumns = make([]columnModifiable, len(stmt.setColumns))
		for i, operation := range stmt.setColumns {
			var clonedColumn SourceBase
			if operation.column != nil {
				clonedColumn = operation.column.clone()
			}
			var clonedValue ExpressionBase
			if operation.value != nil {
				clonedValue = operation.value.clone()
			}
			copy.setColumns[i] = columnModifiable{
				column:    clonedColumn,
				operation: operation.operation,
				value:     clonedValue,
				valueType: operation.valueType,
			}
		}
	}
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
