package uast

// Публичные конструкторы
func NewSelect(from SourceBase) *stmtSelect {
	return &stmtSelect{
		command: uastManagementSelect,
		from:    from,
	}
}

// Публичные методы
func (stmt *stmtSelect) Distinct() *stmtSelect {
	stmt.distinct = true
	return stmt
}
func (stmt *stmtSelect) Field(fields ...markExpressable) *stmtSelect {
	stmt.fields = fields
	return stmt
}
func (stmt *stmtSelect) From(from SourceBase) *stmtSelect {
	stmt.from = from
	return stmt
}
func (stmt *stmtSelect) GroupBy(groupbys ...markGroupable) *stmtSelect {
	stmt.groupBy = groupbys
	return stmt
}
func (stmt *stmtSelect) Having(having markPredicable) *stmtSelect {
	stmt.having = having
	return stmt
}
func (stmt *stmtSelect) Join(joins ...*clauseJoin) *stmtSelect {
	stmt.join = joins
	return stmt
}
func (stmt *stmtSelect) Limit(limit int) *stmtSelect {
	stmt.limit = &clauseLimit{value: limit}
	return stmt
}
func (stmt *stmtSelect) Offset(offset int) *stmtSelect {
	stmt.offset = &clauseOffset{value: offset}
	return stmt
}
func (stmt *stmtSelect) OrderBy(orderbys ...markOrderable) *stmtSelect {
	stmt.orderBy = orderbys
	return stmt
}
func (stmt *stmtSelect) Unions(unions ...*clauseUnions) *stmtSelect {
	stmt.unions = unions
	return stmt
}
func (stmt *stmtSelect) Where(where markPredicable) *stmtSelect {
	stmt.where = where
	return stmt
}
func (stmt *stmtSelect) With(withs ...*clauseWith) *stmtSelect {
	stmt.with = withs
	return stmt
}

// Приватные структуры
type stmtSelect struct {
	command  managementService
	distinct bool
	fields   []markExpressable
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
func (stmt *stmtSelect) clone() statement {
	copy := *stmt
	copy.fields = append([]markExpressable{}, stmt.fields...)
	if stmt.from != nil {
		copy.from = stmt.from.clone()
	}
	copy.groupBy = append([]markGroupable{}, stmt.groupBy...)
	if stmt.having != nil {
		copy.having = stmt.having.clone().(markPredicable)
	}
	copy.join = append([]*clauseJoin{}, stmt.join...)
	copy.orderBy = append([]markOrderable{}, stmt.orderBy...)
	copy.unions = append([]*clauseUnions{}, stmt.unions...)
	if stmt.where != nil {
		copy.where = stmt.where.clone().(markPredicable)
	}
	copy.with = append([]*clauseWith{}, stmt.with...)
	return &copy
}
func (stmt *stmtSelect) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderSelect(baseRenderer, stmt)
}
func (stmt *stmtSelect) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformSelect(baseTransformer, stmt)
}
func (stmt *stmtSelect) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateSelect(baseValidator, stmt)
}
