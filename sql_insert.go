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
	columns := make([]markExpressable, len(source.fields))
	for i, field := range source.fields {
		switch column := field.(type) {
		case *ColumnExpr[string]:
			columns[i] = &exprPair[string]{name: column.name}
		case *ColumnExpr[int]:
			columns[i] = &exprPair[int]{name: column.name}
		case *ColumnExpr[int8]:
			columns[i] = &exprPair[int8]{name: column.name}
		case *ColumnExpr[int16]:
			columns[i] = &exprPair[int16]{name: column.name}
		case *ColumnExpr[int32]:
			columns[i] = &exprPair[int32]{name: column.name}
		case *ColumnExpr[int64]:
			columns[i] = &exprPair[int64]{name: column.name}
		case *ColumnExpr[float32]:
			columns[i] = &exprPair[float32]{name: column.name}
		case *ColumnExpr[float64]:
			columns[i] = &exprPair[float64]{name: column.name}
		case *ColumnExpr[bool]:
			columns[i] = &exprPair[bool]{name: column.name}
		case *ColumnExpr[time.Time]:
			columns[i] = &exprPair[time.Time]{name: column.name}
		case *ColumnExpr[uint]:
			columns[i] = &exprPair[uint]{name: column.name}
		case *ColumnExpr[uint8]:
			columns[i] = &exprPair[uint8]{name: column.name}
		case *ColumnExpr[uint16]:
			columns[i] = &exprPair[uint16]{name: column.name}
		case *ColumnExpr[uint32]:
			columns[i] = &exprPair[uint32]{name: column.name}
		case *ColumnExpr[uint64]:
			columns[i] = &exprPair[uint64]{name: column.name}
		}
	}
	stmt.columns = columns
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
	columns := make([]markExpressable, len(pairs))
	for i, pair := range pairs {
		columns[i] = pair.column
	}
	stmt.columns = columns
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
	columns   []markExpressable
	into      SourceBase
	source    statement
	returning *clauseReturning
	values    *clauseValues
	with      []*clauseWith
}

// Приватные методы
func (stmt *stmtInsert) clone() statement {
	copy := *stmt
	if stmt.columns != nil {
		copy.columns = make([]markExpressable, len(stmt.columns))
		for i, col := range stmt.columns {
			copy.columns[i] = col.clone().(markExpressable)
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
