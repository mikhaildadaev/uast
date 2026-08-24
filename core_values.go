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
func Pair[T typeScalar](field *exprField[T], value ExpressionSafe[T]) *clausePair {
	return &clausePair{
		field: &exprPair[T]{
			name: field.name,
		},
		value: value,
	}
}

// Приватные структуры
type clausePair struct {
	field markExpressable
	value ExpressionBase
}
type clauseUpsert struct {
	pairs   []*clausePair
	service modifierService
}
type clauseValues struct {
	pairs  []*clausePair
	upsert *clauseUpsert
}

// Приватные методы
func (clause *clauseUpsert) clone() *clauseUpsert {
	copy := *clause
	copy.pairs = append([]*clausePair{}, clause.pairs...)
	return &copy
}
func (clause *clauseUpsert) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(clause.service)
	for i, pair := range clause.pairs {
		if i > 0 {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := pair.field.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSpaceEqualSpace)
		if err := pair.value.render(baseRenderer); err != nil {
			return err
		}
	}
	return nil
}
func (clause *clauseUpsert) validate(baseValidator *baseValidator) error {
	if clause.pairs == nil {
		return ErrInvalidStatementSet
	}
	for _, pair := range clause.pairs {
		if err := pair.field.validate(baseValidator); err != nil {
			return err
		}
		if err := pair.value.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
func (clause *clauseValues) clone() *clauseValues {
	copy := *clause
	copy.pairs = append([]*clausePair{}, clause.pairs...)
	if clause.upsert != nil {
		u := *clause.upsert
		copy.upsert = &u
	}
	return &copy
}
func (clause *clauseValues) render(baseRenderer *baseRenderer) error {
	for _, pair := range clause.pairs {
		if err := pair.value.render(baseRenderer); err != nil {
			return err
		}
	}
	return nil
}
func (clause *clauseValues) validate(baseValidator *baseValidator) error {
	if clause.pairs == nil {
		return ErrInvalidStatementValues
	}
	for _, pair := range clause.pairs {
		if pair.field == nil || pair.value == nil {
			return ErrInvalidStatementValues
		}
		if err := pair.field.validate(baseValidator); err != nil {
			return err
		}
		if err := pair.value.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
