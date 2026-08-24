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

// Приватные структуры
type clausePagination struct {
	reverse       bool
	serviceLimit  modifierService
	serviceOffset modifierService
	suffixLimit   modifierService
	suffixOffset  modifierService
	valueLimit    int
	valueOffset   int
}

// Приватные методы
func (clause *clausePagination) clone() *clausePagination {
	copy := *clause
	return &copy
}
func (clause *clausePagination) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderPagination(clause)
	return nil
}
func (clause *clausePagination) validate(baseValidator *baseValidator) error {
	baseValidator.validatePagination(clause)
	return nil
}
