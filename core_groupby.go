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

// Приватные интерфейсы
type markGroupable interface {
	ExpressionBase
	isGroupable()
}

// Приватные структуры
type clauseGroupBy struct {
	expression ExpressionBase
}

// Приватные методы
func (clause *clauseGroupBy) clone() ExpressionBase {
	copy := *clause
	return &copy
}
func (clause *clauseGroupBy) isExpressionBase() {}
func (clause *clauseGroupBy) isGroupable()      {}
func (clause *clauseGroupBy) render(baseRenderer *baseRenderer) error {
	if err := clause.expression.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clause *clauseGroupBy) validate(baseValidator *baseValidator) error {
	if clause.expression == nil {
		return ErrInvalidStatementGroupBy
	}
	return nil
}
