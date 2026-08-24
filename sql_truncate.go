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
func NewTruncate(table *TableSource) *stmtTruncate {
	copy := *table
	copy.withAlias = false
	return &stmtTruncate{
		command: uastManagementTruncate,
		table:   &copy,
	}
}

// Публичные методы
func (stmt *stmtTruncate) IsCascade() *stmtTruncate {
	stmt.isCascade = true
	return stmt
}
func (stmt *stmtTruncate) IsRestartIdentity() *stmtTruncate {
	stmt.isRestartIdentity = true
	return stmt
}

// Приватные структуры
type stmtTruncate struct {
	command           managementService
	table             *TableSource
	isCascade         bool
	isRestartIdentity bool
}

// Приватные методы
func (stmt *stmtTruncate) clone() statement {
	copy := *stmt
	if stmt.table != nil {
		copy.table = stmt.table.clone().(*TableSource)
	}
	return &copy
}
func (stmt *stmtTruncate) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderTruncate(baseRenderer, stmt)
}
func (stmt *stmtTruncate) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformTruncate(baseTransformer, stmt)
}
func (stmt *stmtTruncate) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateTruncate(baseValidator, stmt)
}
