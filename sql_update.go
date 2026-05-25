package uast

// Публичные конструкторы
func NewUpdate(onto SourceBase) *stmtUpdate {
	return &stmtUpdate{
		command: uastManagementUpdate,
		onto:    onto,
	}
}

// Публичные методы
func (stmt *stmtUpdate) Join(joins ...*clauseJoin) *stmtUpdate {
	stmt.join = joins
	return stmt
}
func (stmt *stmtUpdate) Onto(onto SourceBase) *stmtUpdate {
	stmt.onto = onto
	return stmt
}
func (stmt *stmtUpdate) Returning(returnings ...markReturnable) *stmtUpdate {
	stmt.returning = returnings
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
	onto      SourceBase
	join      []*clauseJoin
	returning []markReturnable
	set       []*clauseSet
	where     markPredicable
	with      []*clauseWith
}

// Приватные методы
func (stmt *stmtUpdate) clone() statement {
	copy := *stmt
	copy.set = append([]*clauseSet{}, stmt.set...)
	copy.join = append([]*clauseJoin{}, stmt.join...)
	copy.returning = append([]markReturnable{}, stmt.returning...)
	copy.with = append([]*clauseWith{}, stmt.with...)
	if stmt.where != nil {
		copy.where = stmt.where.clone().(markPredicable)
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
