package uast

// Публичные функции
func Union(statement statement) *clauseUnions {
	return &clauseUnions{
		operator:  uastUnion,
		statement: statement,
	}
}
func UnionAll(statement statement) *clauseUnions {
	return &clauseUnions{
		operator:  uastUnionAll,
		statement: statement,
	}
}
func UnionExcept(statement statement) *clauseUnions {
	return &clauseUnions{
		operator:  uastUnionExcept,
		statement: statement,
	}
}
func UnionIntersect(statement statement) *clauseUnions {
	return &clauseUnions{
		operator:  uastUnionIntersect,
		statement: statement,
	}
}

// Приватные структуры
type clauseUnions struct {
	operator  unionOperator
	statement statement
}

// Приватные методы
func (clauseUnions *clauseUnions) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(clauseUnions.operator)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := clauseUnions.statement.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (clauseUnions *clauseUnions) validate(baseValidator *baseValidator) error {
	if clauseUnions.statement == nil {
		return ErrInvalidStatementUnions
	}
	if err := clauseUnions.statement.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
