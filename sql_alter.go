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
//
//	func (stmt *stmtAlter) AddColumn(column *sourceColumn) *stmtAlter {
//		stmt.action = uastManagementAdd
//		stmt.column = column
//		return stmt
//	}
//
//	func (stmt *stmtAlter) AddConstraintCheck(expr ExpressionBase) *stmtAlter {
//		stmt.action = uastManagementAdd
//		stmt.constraint = &constraintData{
//			constraintType: uastModifierCheck,
//			check:          expr,
//		}
//		return stmt
//	}
//
//	func (stmt *stmtAlter) AddConstraintForeignKey(columns ...markSourceable) *stmtAlter {
//		stmt.action = uastManagementAdd
//		stmt.constraint = &constraintData{
//			constraintType: uastModifierForeignKey,
//			columns:        columns,
//		}
//		return stmt
//	}
//
//	func (stmt *stmtAlter) AddConstraintPrimaryKey(columns ...markSourceable) *stmtAlter {
//		stmt.action = uastManagementAdd
//		stmt.constraint = &constraintData{
//			constraintType: uastModifierPrimaryKey,
//			columns:        columns,
//		}
//		return stmt
//	}
//
//	func (stmt *stmtAlter) AddConstraintUnique(columns ...markSourceable) *stmtAlter {
//		stmt.action = uastManagementAdd
//		stmt.constraint = &constraintData{
//			constraintType: uastModifierUnique,
//			columns:        columns,
//		}
//		return stmt
//	}
//
//	func (stmt *stmtAlter) DropColumn(name string) *stmtAlter {
//		stmt.action = uastManagementDrop
//		stmt.oldName = name
//		return stmt
//	}
//
//	func (stmt *stmtAlter) DropConstraintCheck(name string) *stmtAlter {
//		stmt.action = uastManagementDrop
//		stmt.constraint = &constraintData{
//			constraintType: uastModifierCheck,
//			name:           name,
//		}
//		return stmt
//	}
//
//	func (stmt *stmtAlter) DropConstraintForeignKey(name string) *stmtAlter {
//		stmt.action = uastManagementDrop
//		stmt.constraint = &constraintData{
//			constraintType: uastModifierForeignKey,
//			name:           name,
//		}
//		return stmt
//	}
//
//	func (stmt *stmtAlter) DropConstraintPrimaryKey() *stmtAlter {
//		stmt.action = uastManagementDrop
//		stmt.constraint = &constraintData{
//			constraintType: uastModifierPrimaryKey,
//		}
//		return stmt
//	}
//
//	func (stmt *stmtAlter) DropConstraintUnique(name string) *stmtAlter {
//		stmt.action = uastManagementDrop
//		stmt.constraint = &constraintData{
//			constraintType: uastModifierUnique,
//			name:           name,
//		}
//		return stmt
//	}
//
//	func (stmt *stmtAlter) ModifyColumn(column *sourceColumn) *stmtAlter {
//		stmt.action = uastManagementAlter
//		stmt.column = column
//		return stmt
//	}
func (stmt *stmtAlter) RenameColumn(oldName, newName string) *stmtAlter {
	stmt.action = uastManagementRename
	stmt.oldName = oldName
	stmt.newName = newName
	return stmt
}

// Приватные структуры
type stmtAlter struct {
	action  managementService
	command managementService
	entity  SourceBase
	//column      *sourceColumn
	ifExists    bool
	ifNotExists bool
	oldName     string
	newName     string
	//constraint  *constraintData
	on         SourceBase
	references *sourceTable
	refColumns []string
	onDelete   modifierService
	onUpdate   modifierService
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
