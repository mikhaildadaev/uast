package uast

// Публичные методы
func (stmtDelete *stmtDelete) From(from SourceBase) *stmtDelete {
	stmtDelete.from = from
	return stmtDelete
}
func (stmtDelete *stmtDelete) Join(joins ...*clauseJoin) *stmtDelete {
	stmtDelete.join = joins
	return stmtDelete
}
func (stmtDelete *stmtDelete) Returning(returnings ...markReturnable) *stmtDelete {
	stmtDelete.returning = returnings
	return stmtDelete
}
func (stmtDelete *stmtDelete) Where(where markPredicable) *stmtDelete {
	stmtDelete.where = where
	return stmtDelete
}
func (stmtDelete *stmtDelete) With(with ...*clauseWith) *stmtDelete {
	stmtDelete.with = with
	return stmtDelete
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
func (stmtDelete *stmtDelete) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderDelete(baseRenderer, stmtDelete)
}
func (stmtDelete *stmtDelete) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformDelete(baseTransformer, stmtDelete)
}
func (stmtDelete *stmtDelete) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateDelete(baseValidator, stmtDelete)
}
