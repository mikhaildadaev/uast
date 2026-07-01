package uast

// Публичные конструкторы
func NewUpdate(onTo SourceBase) *stmtUpdate {
	return &stmtUpdate{
		command: uastManagementUpdate,
		onTo:    onTo,
	}
}

// Публичные методы
func (stmt *stmtUpdate) Join(joins ...*clauseJoin) *stmtUpdate {
	stmt.join = joins
	return stmt
}
func (stmt *stmtUpdate) Onto(onTo SourceBase) *stmtUpdate {
	stmt.onTo = onTo
	return stmt
}
func (stmt *stmtUpdate) Returning(returnings ...markReturnable) *stmtUpdate {
	stmt.returning = &clauseReturning{
		expressions:      returnings,
		serviceReturning: uastManagementReturning,
	}
	return stmt
}
func (stmt *stmtUpdate) Set(sets ...*clauseSet) *stmtUpdate {
	stmt.set = sets
	return stmt
}
func (stmt *stmtUpdate) Where(where markPredicable) *stmtUpdate {
	stmt.where = where
	return stmt
}
func (stmt *stmtUpdate) With(with ...*clauseWith) *stmtUpdate {
	stmt.with = with
	return stmt
}

// Приватные структуры
type stmtUpdate struct {
	command   managementService
	join      []*clauseJoin
	onTo      SourceBase
	returning *clauseReturning
	set       []*clauseSet
	where     markPredicable
	with      []*clauseWith
}

// Приватные методы
func (stmt *stmtUpdate) clone() statement {
	copy := *stmt
	if stmt.join != nil {
		copy.join = make([]*clauseJoin, len(stmt.join))
		for i, j := range stmt.join {
			copy.join[i] = j.clone()
		}
	}
	if stmt.onTo != nil {
		copy.onTo = stmt.onTo.clone()
	}
	if stmt.returning != nil {
		copy.returning = stmt.returning.clone()
	}
	if stmt.set != nil {
		copy.set = make([]*clauseSet, len(stmt.set))
		for i, s := range stmt.set {
			copy.set[i] = s.clone()
		}
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
func (stmt *stmtUpdate) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderUpdate(baseRenderer, stmt)
}
func (stmt *stmtUpdate) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformUpdate(baseTransformer, stmt)
}
func (stmt *stmtUpdate) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateUpdate(baseValidator, stmt)
}
