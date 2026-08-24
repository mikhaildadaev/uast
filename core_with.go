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
func WithN(alias string, statement statement, columns ...string) *clauseWith {
	return &clauseWith{
		alias:     alias,
		columns:   columns,
		statement: statement,
	}
}
func WithR(alias string, statement statement, columns ...string) *clauseWith {
	return &clauseWith{
		alias:     alias,
		columns:   columns,
		statement: statement,
		recursive: true,
	}
}

// Приватные структуры
type clauseWith struct {
	alias     string
	columns   []string
	statement statement
	recursive bool
}

// Приватные методы
func (clause *clauseWith) clone() *clauseWith {
	copy := *clause
	copy.statement = clause.statement.clone()
	return &copy
}
func (clause *clauseWith) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderAlias(clause.alias)
	if len(clause.columns) > 0 {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(uastCompositeParenLeft)
		columnsCount := len(clause.columns) - 1
		for i, column := range clause.columns {
			baseRenderer.renderName(column)
			if i < columnsCount {
				baseRenderer.renderOperator(uastCompositeCommaSpace)
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastModifierAs)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderOperator(uastCompositeParenLeft)
	if err := clause.statement.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeParenRight)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	return nil
}
func (clause *clauseWith) validate(baseValidator *baseValidator) error {
	if clause.statement == nil {
		return ErrInvalidStatementWith
	}
	if err := clause.statement.validate(baseValidator); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(clause.alias); err != nil {
		return err
	}
	for _, column := range clause.columns {
		if err := baseValidator.validateName(column); err != nil {
			return err
		}
	}
	return nil
}
