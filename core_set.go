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

// Публичные функции
func Assign[T typeScalar](field *exprField[T], value ExpressionSafe[T]) *clauseSet {
	return &clauseSet{
		field: field,
		value: value,
	}
}

// Приватные структуры
type clauseSet struct {
	field markExpressable
	value ExpressionBase
}

// Приватные методы
func (clause *clauseSet) clone() *clauseSet {
	copy := *clause
	return &copy
}
func (clause *clauseSet) render(baseRenderer *baseRenderer) error {
	if err := clause.field.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderOperator(uastComparisonEqual)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := clause.value.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clause *clauseSet) validate(baseValidator *baseValidator) error {
	if clause.field == nil || clause.value == nil {
		return ErrInvalidStatementSet
	}
	if err := clause.field.validate(baseValidator); err != nil {
		return err
	}
	if err := clause.value.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
