package uast

// Публичные интерфейсы
type SourceBase interface {
	Alias() string
	isSourceBase()
	render(baseRenderer *baseRenderer) error
	validate(baseValidator *baseValidator) error
}

// Публичные структуры
type CteSource struct {
	cteName   string
	aliasName string
}
type TableSource struct {
	tableName string
	aliasName string
}
type QuerySource struct {
	statement statement
	aliasName string
}

// Публичные конструкторы
func NewTable(tableName, tableAlias string) *TableSource {
	if tableAlias == "" {
		tableAlias = tableName
	}
	return &TableSource{
		tableName: tableName,
		aliasName: tableAlias,
	}
}

// Публичные функции
func CTE(cteName string, aliasName string) SourceBase {
	return &CteSource{
		aliasName: aliasName,
		cteName:   cteName,
	}
}
func Query(statement statement, aliasName string) SourceBase {
	return &QuerySource{
		aliasName: aliasName,
		statement: statement,
	}
}
func Table(tableName, aliasName string) SourceBase {
	return &TableSource{
		aliasName: aliasName,
		tableName: tableName,
	}
}

// Публичные методы
func (cteSource *CteSource) Alias() string {
	if cteSource == nil {
		return ""
	}
	return cteSource.aliasName
}
func (tableSource *TableSource) Alias() string {
	if tableSource == nil {
		return ""
	}
	return tableSource.aliasName
}
func (querySource *QuerySource) Alias() string {
	if querySource == nil {
		return ""
	}
	return querySource.aliasName
}

// Приватные методы
func (cteSource *CteSource) isSourceBase() {}
func (cteSource *CteSource) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(cteSource.cteName)
	if cteSource.aliasName != "" {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierAs)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderAlias(cteSource.aliasName)
	}
	return nil
}
func (cteSource *CteSource) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(cteSource.cteName); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(cteSource.aliasName); err != nil {
		return err
	}
	return nil
}
func (tableSource *TableSource) isSourceBase() {}
func (tableSource *TableSource) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(tableSource.tableName)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastModifierAs)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderAlias(tableSource.aliasName)
	return nil
}
func (tableSource *TableSource) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(tableSource.tableName); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(tableSource.aliasName); err != nil {
		return err
	}
	return nil
}
func (querySource *QuerySource) isSourceBase() {}
func (querySource *QuerySource) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeParenLeft)
	if err := querySource.statement.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeParenRight)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastModifierAs)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderAlias(querySource.aliasName)
	return nil
}
func (querySource *QuerySource) validate(baseValidator *baseValidator) error {
	if querySource.statement == nil {
		return ErrInvalidSubquery
	}
	if err := querySource.statement.validate(baseValidator); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(querySource.aliasName); err != nil {
		return err
	}
	return nil
}
