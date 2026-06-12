package uast

// Публичные интерфейсы
type ConstraintBase interface {
	clone() ConstraintBase
	isConstraint()
	getName() string
	render(baseRenderer *baseRenderer) error
	validate(baseValidator *baseValidator) error
}

// Публичные структуры
type CheckConstraint = constraintCheck
type ForeignConstraint = constraintForeign
type ForeignRelation struct {
	Column    SourceBase
	Reference SourceBase
}
type PrimaryConstraint = constraintPrimary
type UniqueConstraint = constraintUnique

// Публичные конструкторы
func NewCheck(name string, expression ExpressionBase) *CheckConstraint {
	return &CheckConstraint{
		expression: expression,
		name:       name,
	}
}
func NewForeignKey(name string, table *sourceTable, onDelete, onUpdate ReferenceAction, relations ...ForeignRelation) *ForeignConstraint {
	columns := make([]SourceBase, len(relations))
	references := make([]SourceBase, len(relations))
	for i, relation := range relations {
		columns[i] = relation.Column
		references[i] = relation.Reference
	}
	return &ForeignConstraint{
		columns:    columns,
		name:       name,
		onDelete:   onDelete,
		onUpdate:   onUpdate,
		references: references,
		table:      table,
	}
}
func NewPrimaryKey(name string, columns ...SourceBase) *PrimaryConstraint {
	return &PrimaryConstraint{
		columns: columns,
		name:    name,
	}
}
func NewUnique(name string, columns ...SourceBase) *UniqueConstraint {
	return &UniqueConstraint{
		columns: columns,
		name:    name,
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

// Приватные структуры
type constraintCheck struct {
	expression ExpressionBase
	name       string
}
type constraintForeign struct {
	columns    []SourceBase
	name       string
	onDelete   ReferenceAction
	onUpdate   ReferenceAction
	references []SourceBase
	table      *sourceTable
}
type constraintPrimary struct {
	columns []SourceBase
	name    string
}
type constraintUnique struct {
	columns []SourceBase
	name    string
}

// Приватные методы
func (constraint *constraintCheck) clone() ConstraintBase {
	return &constraintCheck{
		expression: constraint.expression,
		name:       constraint.name,
	}
}
func (constraint *constraintCheck) getName() string {
	return constraint.name
}
func (constraint *constraintCheck) isConstraint() {}
func (constraint *constraintCheck) render(renderer *baseRenderer) error {
	renderer.renderService(uastModifierConstraint)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderName(constraint.name)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastModifierCheck)
	renderer.renderOperator(uastCompositeParenLeft)
	if err := constraint.expression.render(renderer); err != nil {
		return err
	}
	renderer.renderOperator(uastCompositeParenRight)
	return nil
}
func (constraint *constraintCheck) validate(validator *baseValidator) error {
	if err := validator.validateName(constraint.name); err != nil {
		return err
	}
	if constraint.expression == nil {
		return ErrInvalidConstraintCheck
	}
	return constraint.expression.validate(validator)
}
func (constraint *constraintForeign) clone() ConstraintBase {
	columns := make([]SourceBase, len(constraint.columns))
	references := make([]SourceBase, len(constraint.references))
	copy(columns, constraint.columns)
	copy(references, constraint.references)
	return &constraintForeign{
		columns:    columns,
		name:       constraint.name,
		onDelete:   constraint.onDelete,
		onUpdate:   constraint.onUpdate,
		references: references,
		table:      constraint.table,
	}
}
func (constraint *constraintForeign) getName() string {
	return constraint.name
}
func (constraint *constraintForeign) isConstraint() {}
func (constraint *constraintForeign) render(renderer *baseRenderer) error {
	renderer.renderService(uastModifierConstraint)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderName(constraint.name)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastModifierForeignKey)
	renderer.renderOperator(uastCompositeParenLeft)
	for i, column := range constraint.columns {
		if i > 0 {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
		renderer.renderName(column.getName())
	}
	renderer.renderOperator(uastCompositeParenRight)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastModifierReferences)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderName(constraint.table.getName())
	renderer.renderOperator(uastCompositeParenLeft)
	for i, reference := range constraint.references {
		if i > 0 {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
		renderer.renderName(reference.getName())
	}
	renderer.renderOperator(uastCompositeParenRight)
	if constraint.onDelete != "" {
		renderer.renderOperator(uastCompositeSingleSpace)
		renderer.renderService(uastModifierOnDelete)
		renderer.renderOperator(uastCompositeSingleSpace)
		renderer.renderService(modifierService(constraint.onDelete))
	}
	if constraint.onUpdate != "" {
		renderer.renderOperator(uastCompositeSingleSpace)
		renderer.renderService(uastModifierOnUpdate)
		renderer.renderOperator(uastCompositeSingleSpace)
		renderer.renderService(modifierService(constraint.onUpdate))
	}
	return nil
}
func (constraint *constraintForeign) validate(validator *baseValidator) error {
	if err := validator.validateName(constraint.name); err != nil {
		return err
	}
	if constraint.table == nil || len(constraint.columns) == 0 || len(constraint.columns) != len(constraint.references) {
		return ErrInvalidConstraintForeignKey
	}
	for i, column := range constraint.columns {
		if err := validator.validateName(column.getName()); err != nil {
			return err
		}
		if err := validator.validateName(constraint.references[i].getName()); err != nil {
			return err
		}
	}
	return nil
}
func (constraint *constraintPrimary) clone() ConstraintBase {
	columns := make([]SourceBase, len(constraint.columns))
	copy(columns, constraint.columns)
	return &constraintPrimary{
		columns: columns,
		name:    constraint.name,
	}
}
func (constraint *constraintPrimary) getName() string {
	return constraint.name
}
func (constraint *constraintPrimary) isConstraint() {}
func (constraint *constraintPrimary) render(renderer *baseRenderer) error {
	renderer.renderService(uastModifierConstraint)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderName(constraint.name)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastModifierPrimaryKey)
	renderer.renderOperator(uastCompositeParenLeft)
	for i, column := range constraint.columns {
		if i > 0 {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
		renderer.renderName(column.getName())
	}
	renderer.renderOperator(uastCompositeParenRight)
	return nil
}
func (constraint *constraintPrimary) validate(validator *baseValidator) error {
	if len(constraint.columns) == 0 {
		return ErrInvalidConstraintPrimaryKey
	}
	if err := validator.validateName(constraint.name); err != nil {
		return err
	}
	for _, column := range constraint.columns {
		if err := validator.validateName(column.getName()); err != nil {
			return err
		}
	}
	return nil
}
func (constraint *constraintUnique) clone() ConstraintBase {
	columns := make([]SourceBase, len(constraint.columns))
	copy(columns, constraint.columns)
	return &constraintUnique{
		columns: columns,
		name:    constraint.name,
	}
}
func (constraint *constraintUnique) getName() string {
	return constraint.name
}
func (constraint *constraintUnique) isConstraint() {}
func (constraint *constraintUnique) render(renderer *baseRenderer) error {
	renderer.renderService(uastModifierConstraint)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderName(constraint.name)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastModifierUnique)
	renderer.renderOperator(uastCompositeParenLeft)
	for i, column := range constraint.columns {
		if i > 0 {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
		renderer.renderName(column.getName())
	}
	renderer.renderOperator(uastCompositeParenRight)
	return nil
}
func (constraint *constraintUnique) validate(validator *baseValidator) error {
	if len(constraint.columns) == 0 {
		return ErrInvalidConstraintUnique
	}
	if err := validator.validateName(constraint.name); err != nil {
		return err
	}
	for _, column := range constraint.columns {
		if err := validator.validateName(column.getName()); err != nil {
			return err
		}
	}
	return nil
}
