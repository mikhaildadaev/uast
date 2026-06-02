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
func (stmt *stmtSelect) Fields(fields ...markExpressable) *stmtSelect {
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
func (stmt *stmtSelect) OrderBy(orderbys ...markOrderable) *stmtSelect {
	stmt.orderBy = orderbys
	return stmt
}
func (stmt *stmtSelect) Pagination(limit, offset int) *stmtSelect {
	stmt.pagination = &clausePagination{
		serviceLimit:  uastManagementLimit,
		serviceOffset: uastManagementOffset,
		valueLimit:    limit,
		valueOffset:   offset,
	}
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
	command    managementService
	distinct   bool
	fields     []markExpressable
	from       SourceBase
	groupBy    []markGroupable
	having     markPredicable
	join       []*clauseJoin
	pagination *clausePagination
	orderBy    []markOrderable
	unions     []*clauseUnions
	where      markPredicable
	with       []*clauseWith
}

// Приватные методы
func (stmt *stmtSelect) clone() statement {
	copy := *stmt
	if stmt.fields != nil {
		copy.fields = make([]markExpressable, len(stmt.fields))
		for i, f := range stmt.fields {
			copy.fields[i] = f.clone().(markExpressable)
		}
	}
	if stmt.from != nil {
		copy.from = stmt.from.clone()
	}
	if stmt.groupBy != nil {
		copy.groupBy = make([]markGroupable, len(stmt.groupBy))
		for i, g := range stmt.groupBy {
			copy.groupBy[i] = g.clone().(markGroupable)
		}
	}
	if stmt.having != nil {
		copy.having = stmt.having.clone().(markPredicable)
	}
	if stmt.join != nil {
		copy.join = make([]*clauseJoin, len(stmt.join))
		for i, j := range stmt.join {
			copy.join[i] = j.clone()
		}
	}
	if stmt.orderBy != nil {
		copy.orderBy = make([]markOrderable, len(stmt.orderBy))
		for i, o := range stmt.orderBy {
			copy.orderBy[i] = o.clone().(markOrderable)
		}
	}
	if stmt.unions != nil {
		copy.unions = make([]*clauseUnions, len(stmt.unions))
		for i, u := range stmt.unions {
			copy.unions[i] = u.clone()
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
func (stmt *stmtSelect) render(baseRenderer *baseRenderer) error {
	return baseRenderer.strateger.renderSelect(baseRenderer, stmt)
}
func (stmt *stmtSelect) transform(baseTransformer *baseTransformer) error {
	return baseTransformer.strateger.transformSelect(baseTransformer, stmt)
}
func (stmt *stmtSelect) validate(baseValidator *baseValidator) error {
	return baseValidator.strateger.validateSelect(baseValidator, stmt)
}
