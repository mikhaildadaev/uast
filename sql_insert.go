package uast

import "time"

// Публичные конструкторы
func NewInsert(into SourceBase) *stmtInsert {
	return &stmtInsert{
		command: uastManagementInsert,
		into:    into,
	}
}

// Публичные методы
func (stmt *stmtInsert) Into(into SourceBase) *stmtInsert {
	stmt.into = into
	return stmt
}
func (stmt *stmtInsert) Returning(returnings ...markReturnable) *stmtInsert {
	stmt.returning = &clauseReturning{
		expressions:      returnings,
		serviceReturning: uastManagementReturning,
	}
	return stmt
}
func (stmt *stmtInsert) Source(source *stmtSelect) *stmtInsert {
	fields := make([]markExpressable, len(source.fields))
	for i, field := range source.fields {
		switch column := field.(type) {
		case *ColumnExpr[string]:
			fields[i] = &exprPair[string]{name: column.name}
		case *ColumnExpr[int]:
			fields[i] = &exprPair[int]{name: column.name}
		case *ColumnExpr[int8]:
			fields[i] = &exprPair[int8]{name: column.name}
		case *ColumnExpr[int16]:
			fields[i] = &exprPair[int16]{name: column.name}
		case *ColumnExpr[int32]:
			fields[i] = &exprPair[int32]{name: column.name}
		case *ColumnExpr[int64]:
			fields[i] = &exprPair[int64]{name: column.name}
		case *ColumnExpr[float32]:
			fields[i] = &exprPair[float32]{name: column.name}
		case *ColumnExpr[float64]:
			fields[i] = &exprPair[float64]{name: column.name}
		case *ColumnExpr[bool]:
			fields[i] = &exprPair[bool]{name: column.name}
		case *ColumnExpr[time.Time]:
			fields[i] = &exprPair[time.Time]{name: column.name}
		case *ColumnExpr[uint]:
			fields[i] = &exprPair[uint]{name: column.name}
		case *ColumnExpr[uint8]:
			fields[i] = &exprPair[uint8]{name: column.name}
		case *ColumnExpr[uint16]:
			fields[i] = &exprPair[uint16]{name: column.name}
		case *ColumnExpr[uint32]:
			fields[i] = &exprPair[uint32]{name: column.name}
		case *ColumnExpr[uint64]:
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
		fields[i] = pair.column
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
	into      SourceBase
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
