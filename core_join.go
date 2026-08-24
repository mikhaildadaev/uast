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
func Cross(source SourceBase) *clauseJoin {
	return &clauseJoin{
		operator: uastJoinCross,
		source:   source,
	}
}
func Full(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinFull,
		source:     source,
	}
}
func FullOuter(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinFullOuter,
		source:     source,
	}
}
func Inner(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinInner,
		source:     source,
	}
}
func Left(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinLeft,
		source:     source,
	}
}
func LeftOuter(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinLeftOuter,
		source:     source,
	}
}
func Right(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinRight,
		source:     source,
	}
}
func RightOuter(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinRightOuter,
		source:     source,
	}
}

// Приватные структуры
type clauseJoin struct {
	expression markPredicable
	operator   joinOperator
	source     SourceBase
}

// Приватные методы
func (clause *clauseJoin) clone() *clauseJoin {
	copy := *clause
	if clause.expression != nil {
		copy.expression = clause.expression.clone().(markPredicable)
	}
	if clause.source != nil {
		copy.source = clause.source.clone()
	}
	return &copy
}
func (clause *clauseJoin) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(clause.operator)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := clause.source.render(baseRenderer); err != nil {
		return err
	}
	if clause.expression != nil {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierOn)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := clause.expression.render(baseRenderer); err != nil {
			return err
		}
	}
	return nil
}
func (clause *clauseJoin) validate(baseValidator *baseValidator) error {
	if clause.source == nil {
		return ErrInvalidStatementJoin
	}
	if (clause.operator == uastJoinCross && clause.expression != nil) || (clause.operator != uastJoinCross && clause.expression == nil) {
		return ErrInvalidStatementJoinCross
	}
	return nil
}
