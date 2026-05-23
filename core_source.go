package uast

import (
	"strconv"
	"sync/atomic"
)

// Публичные интерфейсы
type SourceBase interface {
	alias() string
	isSourceBase()
	render(baseRenderer *baseRenderer) error
	validate(baseValidator *baseValidator) error
}

// Публичные структуры
type CteSource = sourceCte
type TableSource = sourceTable
type QuerySource = sourceQuery

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

// Приватные переменные
var queryCounter atomic.Int64

// Приватные структуры
type sourceCte struct {
	aliasName string
	cteName   string
}
type sourceTable struct {
	aliasName string
	tableName string
}
type sourceQuery struct {
	aliasName string
	statement statement
}

// Приватные методы
func (source *sourceCte) alias() string {
	if source == nil {
		return ""
	}
	return source.aliasName
}
func (source *sourceCte) isSourceBase() {}
func (source *sourceCte) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(source.cteName)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastModifierAs)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderAlias(source.aliasName)
	return nil
}
func (source *sourceCte) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(source.cteName); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(source.aliasName); err != nil {
		return err
	}
	return nil
}
func (source *sourceTable) alias() string {
	if source == nil {
		return ""
	}
	return source.aliasName
}
func (source *sourceTable) isSourceBase() {}
func (source *sourceTable) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(source.tableName)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastModifierAs)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderAlias(source.aliasName)
	return nil
}
func (source *sourceTable) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(source.tableName); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(source.aliasName); err != nil {
		return err
	}
	return nil
}
func (source *sourceQuery) alias() string {
	if source == nil {
		return ""
	}
	return source.aliasName
}
func (source *sourceQuery) isSourceBase() {}
func (source *sourceQuery) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeParenLeft)
	if err := source.statement.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeParenRight)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastModifierAs)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderAlias(source.aliasName)
	return nil
}
func (source *sourceQuery) validate(baseValidator *baseValidator) error {
	if source.statement == nil {
		return ErrInvalidSubquery
	}
	if err := source.statement.validate(baseValidator); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(source.aliasName); err != nil {
		return err
	}
	return nil
}
