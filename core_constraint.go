package uast

// Публичные структуры
type ConstraintCheck struct {
	Expression ExpressionBase
	Name       string
}
type ConstraintForeign struct {
	Columns    []SourceBase
	Name       string
	OnDelete   ReferenceAction
	OnUpdate   ReferenceAction
	References []SourceBase
	Table      *sourceTable
}
type ConstraintPrimary struct {
	Columns []SourceBase
	Name    string
}
type ConstraintUnique struct {
	Columns []SourceBase
	Name    string
}
type ForeignRelation struct {
	Column    SourceBase
	Reference SourceBase
}

// Публичные конструкторы
func NewCheck(name string, expression ExpressionBase) *ConstraintCheck {
	return &ConstraintCheck{
		Expression: expression,
		Name:       name,
	}
}
func NewForeignKey(name string, table *sourceTable, onDelete, onUpdate ReferenceAction, relations ...ForeignRelation) *ConstraintForeign {
	columns := make([]SourceBase, len(relations))
	references := make([]SourceBase, len(relations))
	for i, relation := range relations {
		columns[i] = relation.Column
		references[i] = relation.Reference
	}
	return &ConstraintForeign{
		Name:       name,
		Table:      table,
		Columns:    columns,
		References: references,
		OnDelete:   onDelete,
		OnUpdate:   onUpdate,
	}
}
func NewPrimaryKey(name string, columns ...SourceBase) *ConstraintPrimary {
	return &ConstraintPrimary{
		Columns: columns,
		Name:    name,
	}
}

func NewUnique(name string, columns ...SourceBase) *ConstraintUnique {
	return &ConstraintUnique{
		Columns: columns,
		Name:    name,
	}
}

// Публичные функции
func Cascade() ReferenceAction {
	return ActionCascade
}
func NoAction() ReferenceAction {
	return ActionNoAction
}
func Relation(column, reference SourceBase) ForeignRelation {
	return ForeignRelation{
		Column:    column,
		Reference: reference,
	}
}
func Restrict() ReferenceAction {
	return ActionRestrict
}
func SetDefault() ReferenceAction {
	return ActionSetDefault
}
func SetNull() ReferenceAction {
	return ActionSetNull
}
