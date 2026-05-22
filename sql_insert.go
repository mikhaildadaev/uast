package uast

import "time"

// Публичные методы
func (stmtInsert *stmtInsert) Into(into SourceBase) *stmtInsert {
	stmtInsert.into = into
	return stmtInsert
}
func (stmtInsert *stmtInsert) Returning(returnings ...markReturnable) *stmtInsert {
	stmtInsert.returning = returnings
	return stmtInsert
}
func (stmtInsert *stmtInsert) Source(source *stmtSelect) *stmtInsert {
	columns := make([]markExpressable, len(source.field))
	for i, field := range source.field {
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
	stmtInsert.column = columns
	stmtInsert.source = source
	return stmtInsert
}
func (stmtInsert *stmtInsert) Upsert(pairs ...*clausePair) *stmtInsert {
	if stmtInsert.values != nil {
		stmtInsert.values.upsert = &clauseUpsert{pairs: pairs}
	}
	return stmtInsert
}
func (stmtInsert *stmtInsert) Values(pairs ...*clausePair) *stmtInsert {
	columns := make([]markExpressable, len(pairs))
	for i, pair := range pairs {
		columns[i] = pair.column
	}
	stmtInsert.column = columns
	stmtInsert.values = &clauseValues{pairs: pairs}
	return stmtInsert
}
func (stmtInsert *stmtInsert) With(with ...*clauseWith) *stmtInsert {
	stmtInsert.with = with
	return stmtInsert
}

// Приватные структуры
type stmtInsert struct {
	command   managementService
	column    []markExpressable
	into      SourceBase
	source    statement
	returning []markReturnable
	values    *clauseValues
	with      []*clauseWith
}

// Приватные методы
func (stmtInsert *stmtInsert) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderInsert(baseRenderer, stmtInsert)
}
func (stmtInsert *stmtInsert) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformInsert(baseTransformer, stmtInsert)
}
func (stmtInsert *stmtInsert) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateInsert(baseValidator, stmtInsert)
}
