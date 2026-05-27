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
	stmt.returning = &clauseReturning{
		expressions:      returnings,
		serviceReturning: uastManagementReturning,
	}
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
	returning *clauseReturning
	where     markPredicable
	with      []*clauseWith
}

// Приватные методы
func (stmt *stmtDelete) clone() statement {
	copy := *stmt
	if stmt.from != nil {
		copy.from = stmt.from.clone()
	}
	if stmt.join != nil {
		copy.join = make([]*clauseJoin, len(stmt.join))
		for i, j := range stmt.join {
			copy.join[i] = j.clone()
		}
	}
	if stmt.returning != nil {
		copy.returning = stmt.returning.clone()
	}
	if stmt.where != nil {
		copy.where = stmt.where.clone().(markPredicable)
	}
	if stmt.with != nil {
		copy.with = make([]*clauseWith, len(stmt.with))
		for i, w := range stmt.with {
			copy.with[i] = w.clone()
		}
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
