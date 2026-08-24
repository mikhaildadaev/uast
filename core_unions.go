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
func Union(statement statement) *clauseUnions {
	return &clauseUnions{
		operator:  uastUnion,
		statement: statement,
	}
}
func UnionAll(statement statement) *clauseUnions {
	return &clauseUnions{
		operator:  uastUnionAll,
		statement: statement,
	}
}
func UnionExcept(statement statement) *clauseUnions {
	return &clauseUnions{
		operator:  uastUnionExcept,
		statement: statement,
	}
}
func UnionIntersect(statement statement) *clauseUnions {
	return &clauseUnions{
		operator:  uastUnionIntersect,
		statement: statement,
	}
}

// Приватные структуры
type clauseUnions struct {
	operator  unionOperator
	statement statement
}

// Приватные методы
func (clause *clauseUnions) clone() *clauseUnions {
	copy := *clause
	copy.statement = clause.statement.clone()
	return &copy
}
func (clause *clauseUnions) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(clause.operator)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := clause.statement.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clause *clauseUnions) validate(baseValidator *baseValidator) error {
	if clause.statement == nil {
		return ErrInvalidStatementUnions
	}
	if err := clause.statement.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
