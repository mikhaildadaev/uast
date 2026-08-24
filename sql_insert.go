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

import "time"

// Публичные конструкторы
func NewInsert(inTo SourceBase) *stmtInsert {
	return &stmtInsert{
		command: uastManagementInsert,
		inTo:    inTo,
	}
}

// Публичные методы
func (stmt *stmtInsert) Into(inTo SourceBase) *stmtInsert {
	stmt.inTo = inTo
	return stmt
}
func (stmt *stmtInsert) Returning(returnings ...markReturnable) *stmtInsert {
	stmt.returning = &clauseReturning{
		expressions:      returnings,
		serviceReturning: uastModifierReturning,
	}
	return stmt
}
func (stmt *stmtInsert) Source(source *stmtSelect) *stmtInsert {
	fields := make([]markExpressable, len(source.fields))
	for i, field := range source.fields {
		switch column := field.(type) {
		case *FieldExpr[string]:
			fields[i] = &exprPair[string]{name: column.name}
		case *FieldExpr[int]:
			fields[i] = &exprPair[int]{name: column.name}
		case *FieldExpr[int8]:
			fields[i] = &exprPair[int8]{name: column.name}
		case *FieldExpr[int16]:
			fields[i] = &exprPair[int16]{name: column.name}
		case *FieldExpr[int32]:
			fields[i] = &exprPair[int32]{name: column.name}
		case *FieldExpr[int64]:
			fields[i] = &exprPair[int64]{name: column.name}
		case *FieldExpr[float32]:
			fields[i] = &exprPair[float32]{name: column.name}
		case *FieldExpr[float64]:
			fields[i] = &exprPair[float64]{name: column.name}
		case *FieldExpr[bool]:
			fields[i] = &exprPair[bool]{name: column.name}
		case *FieldExpr[time.Time]:
			fields[i] = &exprPair[time.Time]{name: column.name}
		case *FieldExpr[uint]:
			fields[i] = &exprPair[uint]{name: column.name}
		case *FieldExpr[uint8]:
			fields[i] = &exprPair[uint8]{name: column.name}
		case *FieldExpr[uint16]:
			fields[i] = &exprPair[uint16]{name: column.name}
		case *FieldExpr[uint32]:
			fields[i] = &exprPair[uint32]{name: column.name}
		case *FieldExpr[uint64]:
			fields[i] = &exprPair[uint64]{name: column.name}
		}
	}
	stmt.fields = fields
	stmt.source = source
	return stmt
}
func (stmt *stmtInsert) Upsert(pairs ...*clausePair) *stmtInsert {
	if stmt.values != nil {
		stmt.values.upsert = &clauseUpsert{pairs: pairs}
	}
	return stmt
}
func (stmt *stmtInsert) Values(pairs ...*clausePair) *stmtInsert {
	fields := make([]markExpressable, len(pairs))
	for i, pair := range pairs {
		fields[i] = pair.field
	}
	stmt.fields = fields
	stmt.values = &clauseValues{pairs: pairs}
	return stmt
}
func (stmt *stmtInsert) With(with ...*clauseWith) *stmtInsert {
	stmt.with = with
	return stmt
}

// Приватные структуры
type stmtInsert struct {
	command   managementService
	fields    []markExpressable
	inTo      SourceBase
	source    statement
	returning *clauseReturning
	values    *clauseValues
	with      []*clauseWith
}

// Приватные методы
func (stmt *stmtInsert) clone() statement {
	copy := *stmt
	if stmt.fields != nil {
		copy.fields = make([]markExpressable, len(stmt.fields))
		for i, col := range stmt.fields {
			copy.fields[i] = col.clone().(markExpressable)
		}
	}
	if stmt.returning != nil {
		copy.returning = stmt.returning.clone()
	}
	if stmt.source != nil {
		copy.source = stmt.source.clone()
	}
	if stmt.values != nil {
		copy.values = stmt.values.clone()
	}
	if stmt.with != nil {
		copy.with = make([]*clauseWith, len(stmt.with))
		for i, w := range stmt.with {
			copy.with[i] = w.clone()
		}
	}
	return &copy
}
func (stmt *stmtInsert) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderInsert(baseRenderer, stmt)
}
func (stmt *stmtInsert) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformInsert(baseTransformer, stmt)
}
func (stmt *stmtInsert) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateInsert(baseValidator, stmt)
}
