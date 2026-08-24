// UAST (Universal Abstract Syntax Tree)
// Copyright (C) 2026 Mikhail Dadaev
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package uast

// Публичные конструкторы
func NewDrop(entity SourceBase) *stmtDrop {
	return &stmtDrop{
		command: uastManagementDrop,
		entity:  entity,
	}
}

// Публичные методы
func (stmt *stmtDrop) IfExists() *stmtDrop {
	stmt.ifExists = true
	return stmt
}
func (stmt *stmtDrop) IsCascade() *stmtDrop {
	stmt.isCascade = true
	return stmt
}

// Приватные структуры
type stmtDrop struct {
	command   managementService
	entity    SourceBase
	ifExists  bool
	isCascade bool
}

// Приватные методы
func (stmt *stmtDrop) clone() statement {
	copy := *stmt
	if stmt.entity != nil {
		copy.entity = stmt.entity.clone()
	}
	return &copy
}
func (stmt *stmtDrop) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderDrop(baseRenderer, stmt)
}
func (stmt *stmtDrop) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformDrop(baseTransformer, stmt)
}
func (stmt *stmtDrop) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateDrop(baseValidator, stmt)
}
