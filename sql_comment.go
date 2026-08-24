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
func NewComment(onTable SourceBase) *stmtComment {
	return &stmtComment{
		command: uastManagementComment,
		onTable: onTable,
	}
}

// Публичные методы
func (stmt *stmtComment) OnColumn(onColumn SourceBase) *stmtComment {
	stmt.onColumn = onColumn
	return stmt
}
func (stmt *stmtComment) Is(comment string) *stmtComment {
	stmt.comment = comment
	return stmt
}

// Приватные структуры
type stmtComment struct {
	command  managementService
	comment  string
	onColumn SourceBase
	onTable  SourceBase
}

// Приватные методы
func (stmt *stmtComment) clone() statement {
	copy := *stmt
	if stmt.onColumn != nil {
		copy.onColumn = stmt.onColumn.clone()
	}
	if stmt.onTable != nil {
		copy.onTable = stmt.onTable.clone()
	}
	return &copy
}
func (stmt *stmtComment) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderComment(baseRenderer, stmt)
}
func (stmt *stmtComment) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformComment(baseTransformer, stmt)
}
func (stmt *stmtComment) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateComment(baseValidator, stmt)
}
