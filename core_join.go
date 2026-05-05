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
func (clauseJoin *clauseJoin) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(clauseJoin.operator)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := clauseJoin.source.render(baseRenderer); err != nil {
		return err
	}
	if clauseJoin.expression != nil {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierOn)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := clauseJoin.expression.render(baseRenderer); err != nil {
			return err
		}
	}
	return nil
}
func (clauseJoin *clauseJoin) validate(baseValidator *baseValidator) error {
	if clauseJoin.source == nil {
		return ErrInvalidStatementJoin
	}
	if (clauseJoin.operator == uastJoinCross && clauseJoin.expression != nil) || (clauseJoin.operator != uastJoinCross && clauseJoin.expression == nil) {
		return ErrInvalidStatementJoinCross
	}
	return nil
}
