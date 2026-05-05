package uast

// Публичные функции
func WithN(alias string, statement statement, columns ...string) *clauseWith {
	return &clauseWith{
		alias:     alias,
		columns:   columns,
		statement: statement,
	}
}
func WithR(alias string, statement statement, columns ...string) *clauseWith {
	return &clauseWith{
		alias:     alias,
		columns:   columns,
		statement: statement,
		recursive: true,
	}
}

// Приватные структуры
type clauseWith struct {
	alias     string
	columns   []string
	statement statement
	recursive bool
}

// Приватные методы
func (clauseWith *clauseWith) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderAlias(clauseWith.alias)
	if len(clauseWith.columns) > 0 {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(uastCompositeParenLeft)
		columnsCount := len(clauseWith.columns) - 1
		for i, column := range clauseWith.columns {
			baseRenderer.renderName(column)
			if i < columnsCount {
				baseRenderer.renderOperator(uastCompositeCommaSpace)
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastModifierAs)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderOperator(uastCompositeParenLeft)
	if err := clauseWith.statement.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeParenRight)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	return nil
}
func (clauseWith *clauseWith) validate(baseValidator *baseValidator) error {
	if clauseWith.statement == nil {
		return ErrInvalidStatementWith
	}
	if err := clauseWith.statement.validate(baseValidator); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(clauseWith.alias); err != nil {
		return err
	}
	for _, column := range clauseWith.columns {
		if err := baseValidator.validateName(column); err != nil {
			return err
		}
	}
	return nil
}
