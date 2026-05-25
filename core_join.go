package uast

// Публичные функции
func Cross(source SourceBase) *clauseJoin {
	return &clauseJoin{
		operator: uastJoinCross,
		source:   source,
	}
}
func Full(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinFull,
		source:     source,
	}
}
func FullOuter(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinFullOuter,
		source:     source,
	}
}
func Inner(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinInner,
		source:     source,
	}
}
func Left(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinLeft,
		source:     source,
	}
}
func LeftOuter(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinLeftOuter,
		source:     source,
	}
}
func Right(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinRight,
		source:     source,
	}
}
func RightOuter(source SourceBase, expression markPredicable) *clauseJoin {
	return &clauseJoin{
		expression: expression,
		operator:   uastJoinRightOuter,
		source:     source,
	}
}

// Приватные структуры
type clauseJoin struct {
	expression markPredicable
	operator   joinOperator
	source     SourceBase
}

// Приватные методы
func (clause *clauseJoin) clone() *clauseJoin {
	copy := *clause
	return &copy
}
func (clause *clauseJoin) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(clause.operator)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := clause.source.render(baseRenderer); err != nil {
		return err
	}
	if clause.expression != nil {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierOn)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := clause.expression.render(baseRenderer); err != nil {
			return err
		}
	}
	return nil
}
func (clause *clauseJoin) validate(baseValidator *baseValidator) error {
	if clause.source == nil {
		return ErrInvalidStatementJoin
	}
	if (clause.operator == uastJoinCross && clause.expression != nil) || (clause.operator != uastJoinCross && clause.expression == nil) {
		return ErrInvalidStatementJoinCross
	}
	return nil
}
