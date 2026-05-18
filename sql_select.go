package uast

// Публичные методы
func (stmtSelect *stmtSelect) Distinct() *stmtSelect {
	stmtSelect.distinct = true
	return stmtSelect
}
func (stmtSelect *stmtSelect) Field(fields ...markExpressable) *stmtSelect {
	stmtSelect.field = fields
	return stmtSelect
}
func (stmtSelect *stmtSelect) From(from SourceBase) *stmtSelect {
	stmtSelect.from = from
	return stmtSelect
}
func (stmtSelect *stmtSelect) GroupBy(groupbys ...markGroupable) *stmtSelect {
	stmtSelect.groupBy = groupbys
	return stmtSelect
}
func (stmtSelect *stmtSelect) Having(having markPredicable) *stmtSelect {
	stmtSelect.having = having
	return stmtSelect
}
func (stmtSelect *stmtSelect) Join(joins ...*clauseJoin) *stmtSelect {
	stmtSelect.join = joins
	return stmtSelect
}
func (stmtSelect *stmtSelect) Limit(limit int) *stmtSelect {
	stmtSelect.limit = &clauseLimit{value: limit}
	return stmtSelect
}
func (stmtSelect *stmtSelect) Offset(offset int) *stmtSelect {
	stmtSelect.offset = &clauseOffset{value: offset}
	return stmtSelect
}
func (stmtSelect *stmtSelect) OrderBy(orderbys ...markOrderable) *stmtSelect {
	stmtSelect.orderBy = orderbys
	return stmtSelect
}
func (stmtSelect *stmtSelect) Unions(unions ...*clauseUnions) *stmtSelect {
	stmtSelect.unions = unions
	return stmtSelect
}
func (stmtSelect *stmtSelect) Where(where markPredicable) *stmtSelect {
	stmtSelect.where = where
	return stmtSelect
}
func (stmtSelect *stmtSelect) With(withs ...*clauseWith) *stmtSelect {
	stmtSelect.with = withs
	return stmtSelect
}

// Приватные структуры
type stmtSelect struct {
	command  managementService
	distinct bool
	field    []markExpressable
	from     SourceBase
	groupBy  []markGroupable
	having   markPredicable
	join     []*clauseJoin
	limit    *clauseLimit
	offset   *clauseOffset
	orderBy  []markOrderable
	unions   []*clauseUnions
	where    markPredicable
	with     []*clauseWith
}

// Приватные методы
func (stmtSelect *stmtSelect) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderSelect(baseRenderer, stmtSelect)
}
func (stmtSelect *stmtSelect) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformSelect(baseTransformer, stmtSelect)
}
func (stmtSelect *stmtSelect) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateSelect(baseValidator, stmtSelect)
}
