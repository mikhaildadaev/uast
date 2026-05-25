package uast

// Публичные конструкторы
func NewDelete(from SourceBase) *stmtDelete {
	return &stmtDelete{
		command: uastManagementDelete,
		from:    from,
	}
}

// Публичные методы
func (stmt *stmtDelete) From(from SourceBase) *stmtDelete {
	stmt.from = from
	return stmt
}
func (stmt *stmtDelete) Join(joins ...*clauseJoin) *stmtDelete {
	stmt.join = joins
	return stmt
}
func (stmt *stmtDelete) Returning(returnings ...markReturnable) *stmtDelete {
	stmt.returning = returnings
	return stmt
}
func (stmt *stmtDelete) Where(where markPredicable) *stmtDelete {
	stmt.where = where
	return stmt
}
func (stmt *stmtDelete) With(with ...*clauseWith) *stmtDelete {
	stmt.with = with
	return stmt
}

// Приватные структуры
type stmtDelete struct {
	command   managementService
	from      SourceBase
	join      []*clauseJoin
	returning []markReturnable
	where     markPredicable
	with      []*clauseWith
}

// Приватные методы
func (stmt *stmtDelete) clone() statement {
	copy := *stmt
	if stmt.from != nil {
		copy.from = stmt.from.clone()
	}
	copy.join = append([]*clauseJoin{}, stmt.join...)
	copy.returning = append([]markReturnable{}, stmt.returning...)
	copy.with = append([]*clauseWith{}, stmt.with...)
	if stmt.where != nil {
		copy.where = stmt.where.clone().(markPredicable)
	}
	return &copy
}
func (stmt *stmtDelete) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderDelete(baseRenderer, stmt)
}
func (stmt *stmtDelete) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformDelete(baseTransformer, stmt)
}
func (stmt *stmtDelete) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateDelete(baseValidator, stmt)
}
