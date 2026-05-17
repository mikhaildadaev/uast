package uast

import (
	"strconv"
	"sync/atomic"
)

// Публичные интерфейсы
type SourceBase interface {
	Alias() string
	isSourceBase()
	render(baseRenderer *baseRenderer) error
	validate(baseValidator *baseValidator) error
}

// Публичные структуры
type CteSource struct {
	aliasName string
	cteName   string
}
type TableSource struct {
	aliasName string
	tableName string
}
type QuerySource struct {
	aliasName string
	statement statement
}

// Публичные конструкторы
func NewCTE(name string, alias string) *CteSource {
	if alias == "" {
		alias = name
	}
	return &CteSource{
		aliasName: alias,
		cteName:   name,
	}
}
func NewQuery(statement statement, alias string) *QuerySource {
	if alias == "" {
		alias = "query_" + strconv.FormatInt(queryCounter.Add(1), 10)
	}
	return &QuerySource{
		aliasName: alias,
		statement: statement,
	}
}
func NewTable(name, alias string) *TableSource {
	if alias == "" {
		alias = name
	}
	return &TableSource{
		aliasName: alias,
		tableName: name,
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

// Приватные переменные
var queryCounter atomic.Int64

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
