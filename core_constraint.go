package uast

// Публичные структуры
type ConstraintCheck struct {
	Name       string
	Expression ExpressionBase
}
type ConstraintForeign struct {
	Name       string
	Table      *sourceTable
	Columns    []string
	RefColumns []string
	OnDelete   ReferenceAction
	OnUpdate   ReferenceAction
}
type ConstraintForeignOption func(*ConstraintForeign)
type ConstraintPrimary struct {
	Name    string
	Columns []string
}
type ConstraintUnique struct {
	Name    string
	Columns []string
}

// Публичные конструкторы
func NewCheck(name string, expression ExpressionBase) *ConstraintCheck {
	return &ConstraintCheck{
		Expression: expression,
		Name:       name,
	}
}
func NewForeignKey(name string, table *sourceTable, columns, refColumns []string, options ...ConstraintForeignOption) *ConstraintForeign {
	foreignKey := &ConstraintForeign{
		Columns:    columns,
		Name:       name,
		RefColumns: refColumns,
		Table:      table,
	}
	for _, option := range options {
		option(foreignKey)
	}
	return foreignKey
}
func NewPrimaryKey(name string, columns ...string) *ConstraintPrimary {
	return &ConstraintPrimary{
		Columns: columns,
		Name:    name,
	}
}

func NewUnique(name string, columns ...string) *ConstraintUnique {
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
func OnDelete(action ReferenceAction) ConstraintForeignOption {
	return func(foreignKey *ConstraintForeign) {
		foreignKey.OnDelete = action
	}
}
func OnUpdate(action ReferenceAction) ConstraintForeignOption {
	return func(foreignKey *ConstraintForeign) {
		foreignKey.OnUpdate = action
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
