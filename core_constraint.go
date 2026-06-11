package uast

// Публичные интерфейсы
type Constraint interface {
	clone() Constraint
	isConstraint()
	render(baseRenderer *baseRenderer) error
	validate(baseValidator *baseValidator) error
}

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
		Columns:    columns,
		Name:       name,
		OnDelete:   onDelete,
		OnUpdate:   onUpdate,
		References: references,
		Table:      table,
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

// Приватные методы
func (сonstraintCheck *ConstraintCheck) clone() Constraint {
	return &ConstraintCheck{
		Expression: сonstraintCheck.Expression,
		Name:       сonstraintCheck.Name,
	}
}
func (сonstraintCheck *ConstraintCheck) isConstraint() {}
func (сonstraintCheck *ConstraintCheck) render(renderer *baseRenderer) error {
	renderer.renderService(uastModifierConstraint)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderName(сonstraintCheck.Name)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastModifierCheck)
	renderer.renderOperator(uastCompositeParenLeft)
	if err := сonstraintCheck.Expression.render(renderer); err != nil {
		return err
	}
	renderer.renderOperator(uastCompositeParenRight)
	return nil
}
func (constraintCheck *ConstraintCheck) validate(validator *baseValidator) error {
	if err := validator.validateName(constraintCheck.Name); err != nil {
		return err
	}
	if constraintCheck.Expression == nil {
		return ErrInvalidConstraintCheck
	}
	return constraintCheck.Expression.validate(validator)
}
func (constraintForeign *ConstraintForeign) clone() Constraint {
	columns := make([]SourceBase, len(constraintForeign.Columns))
	references := make([]SourceBase, len(constraintForeign.References))
	copy(columns, constraintForeign.Columns)
	copy(references, constraintForeign.References)
	return &ConstraintForeign{
		Columns:    columns,
		Name:       constraintForeign.Name,
		OnDelete:   constraintForeign.OnDelete,
		OnUpdate:   constraintForeign.OnUpdate,
		References: references,
		Table:      constraintForeign.Table,
	}
}
func (constraintForeign *ConstraintForeign) isConstraint() {}
func (constraintForeign *ConstraintForeign) render(renderer *baseRenderer) error {
	renderer.renderService(uastModifierConstraint)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderName(constraintForeign.Name)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastModifierForeignKey)
	renderer.renderOperator(uastCompositeParenLeft)
	for i, column := range constraintForeign.Columns {
		if i > 0 {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
		renderer.renderName(column.name())
	}
	renderer.renderOperator(uastCompositeParenRight)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastModifierReferences)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderName(constraintForeign.Table.name())
	renderer.renderOperator(uastCompositeParenLeft)
	for i, ref := range constraintForeign.References {
		if i > 0 {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
		renderer.renderName(ref.name())
	}
	renderer.renderOperator(uastCompositeParenRight)

	if constraintForeign.OnDelete != "" {
		renderer.renderOperator(uastCompositeSingleSpace)
		renderer.renderService(uastModifierOnDelete)
		renderer.renderOperator(uastCompositeSingleSpace)
		renderer.renderService(modifierService(constraintForeign.OnDelete))
	}
	if constraintForeign.OnUpdate != "" {
		renderer.renderOperator(uastCompositeSingleSpace)
		renderer.renderService(uastModifierOnUpdate)
		renderer.renderOperator(uastCompositeSingleSpace)
		renderer.renderService(modifierService(constraintForeign.OnUpdate))
	}
	return nil
}
func (constraintForeign *ConstraintForeign) validate(validator *baseValidator) error {
	if err := validator.validateName(constraintForeign.Name); err != nil {
		return err
	}
	if constraintForeign.Table == nil || len(constraintForeign.Columns) == 0 || len(constraintForeign.Columns) != len(constraintForeign.References) {
		return ErrInvalidConstraintForeignKey
	}
	for i, column := range constraintForeign.Columns {
		if err := validator.validateName(column.name()); err != nil {
			return err
		}
		if err := validator.validateName(constraintForeign.References[i].name()); err != nil {
			return err
		}
	}
	return nil
}
func (constraintPrimary *ConstraintPrimary) clone() Constraint {
	columns := make([]SourceBase, len(constraintPrimary.Columns))
	copy(columns, constraintPrimary.Columns)
	return &ConstraintPrimary{
		Columns: columns,
		Name:    constraintPrimary.Name,
	}
}
func (constraintPrimary *ConstraintPrimary) isConstraint() {}
func (constraintPrimary *ConstraintPrimary) render(renderer *baseRenderer) error {
	renderer.renderService(uastModifierConstraint)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderName(constraintPrimary.Name)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastModifierPrimaryKey)
	renderer.renderOperator(uastCompositeParenLeft)
	for i, column := range constraintPrimary.Columns {
		if i > 0 {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
		renderer.renderName(column.name())
	}
	renderer.renderOperator(uastCompositeParenRight)
	return nil
}
func (constraintPrimary *ConstraintPrimary) validate(validator *baseValidator) error {
	if len(constraintPrimary.Columns) == 0 {
		return ErrInvalidConstraintPrimaryKey
	}
	if err := validator.validateName(constraintPrimary.Name); err != nil {
		return err
	}
	for _, column := range constraintPrimary.Columns {
		if err := validator.validateName(column.name()); err != nil {
			return err
		}
	}
	return nil
}
func (constraintUnique *ConstraintUnique) clone() Constraint {
	columns := make([]SourceBase, len(constraintUnique.Columns))
	copy(columns, constraintUnique.Columns)
	return &ConstraintUnique{
		Columns: columns,
		Name:    constraintUnique.Name,
	}
}
func (constraintUnique *ConstraintUnique) isConstraint() {}
func (constraintUnique *ConstraintUnique) render(renderer *baseRenderer) error {
	renderer.renderService(uastModifierConstraint)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderName(constraintUnique.Name)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastModifierUnique)
	renderer.renderOperator(uastCompositeParenLeft)
	for i, column := range constraintUnique.Columns {
		if i > 0 {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
		renderer.renderName(column.name())
	}
	renderer.renderOperator(uastCompositeParenRight)
	return nil
}
func (constraintUnique *ConstraintUnique) validate(validator *baseValidator) error {
	if len(constraintUnique.Columns) == 0 {
		return ErrInvalidConstraintUnique
	}
	if err := validator.validateName(constraintUnique.Name); err != nil {
		return err
	}
	for _, column := range constraintUnique.Columns {
		if err := validator.validateName(column.name()); err != nil {
			return err
		}
	}
	return nil
}
