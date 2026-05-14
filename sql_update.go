package uast

// Публичные методы
func (stmtUpdate *stmtUpdate) Onto(onto SourceBase) *stmtUpdate {
	stmtUpdate.onto = onto
	return stmtUpdate
}
func (stmtUpdate *stmtUpdate) Join(joins ...*clauseJoin) *stmtUpdate {
	stmtUpdate.join = joins
	return stmtUpdate
}
func (stmtUpdate *stmtUpdate) Returning(returnings ...markReturnable) *stmtUpdate {
	stmtUpdate.returning = returnings
	return stmtUpdate
}
func (stmtUpdate *stmtUpdate) Set(sets ...*clausePair) *stmtUpdate {
	stmtUpdate.set = sets
	return stmtUpdate
}
func (stmtUpdate *stmtUpdate) Where(where markPredicable) *stmtUpdate {
	stmtUpdate.where = where
	return stmtUpdate
}
func (stmtUpdate *stmtUpdate) With(with ...*clauseWith) *stmtUpdate {
	stmtUpdate.with = with
	return stmtUpdate
}

// Приватные структуры
type stmtUpdate struct {
	command   managementService
	onto      SourceBase
	join      []*clauseJoin
	returning []markReturnable
	set       []*clausePair
	where     markPredicable
	with      []*clauseWith
}

// Приватные методы
func (stmtUpdate *stmtUpdate) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderUpdate(baseRenderer, stmtUpdate)
}
func (stmtUpdate *stmtUpdate) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformUpdate(baseTransformer, stmtUpdate)
}
func (stmtUpdate *stmtUpdate) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateUpdate(baseValidator, stmtUpdate)
}
