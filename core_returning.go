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
type markReturnable interface {
	ExpressionBase
	isReturnable()
}

// Приватные структуры
type clauseReturning struct {
	expressions      []markReturnable
	serviceReturning modifierService
}

// Приватные методы
func (clause *clauseReturning) clone() *clauseReturning {
	copy := *clause
	copy.expressions = make([]markReturnable, len(clause.expressions))
	for i, expr := range clause.expressions {
		copy.expressions[i] = expr
	}
	return &copy
}
func (clause *clauseReturning) isExpressionBase() {}
func (clause *clauseReturning) isReturnable()     {}
func (clause *clauseReturning) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(clause.serviceReturning)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	for i, expressions := range clause.expressions {
		if i > 0 {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
		if err := expressions.render(baseRenderer); err != nil {
			return err
		}
	}
	return nil
}
func (clause *clauseReturning) validate(baseValidator *baseValidator) error {
	if clause.expressions == nil {
		return nil
	}
	for _, expression := range clause.expressions {
		if err := expression.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
