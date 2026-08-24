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
type markOrderable interface {
	ExpressionBase
	isOrderable()
}

// Приватные структуры
type clauseOrderBy struct {
	direction  bool
	expression ExpressionBase
}

// Приватные методы
func (clause *clauseOrderBy) clone() ExpressionBase {
	copy := *clause
	return &copy
}
func (clause *clauseOrderBy) isExpressionBase() {}
func (clause *clauseOrderBy) isOrderable()      {}
func (clause *clauseOrderBy) render(baseRenderer *baseRenderer) error {
	if err := clause.expression.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if clause.direction {
		baseRenderer.renderOperator(uastOrderDesc)
	} else {
		baseRenderer.renderOperator(uastOrderAsc)
	}
	return nil
}
func (clause *clauseOrderBy) validate(baseValidator *baseValidator) error {
	if clause == nil {
		return ErrInvalidStatementOrderBy
	}
	return nil
}
