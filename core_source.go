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
func (sourceCte *sourceCte) alias() string {
	if sourceCte == nil {
		return ""
	}
	return sourceCte.aliasName
}
func (sourceCte *sourceCte) isSourceBase() {}
func (sourceCte *sourceCte) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(sourceCte.cteName)
	if sourceCte.aliasName != "" {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierAs)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderAlias(sourceCte.aliasName)
	}
	return nil
}
func (sourceCte *sourceCte) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(sourceCte.cteName); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(sourceCte.aliasName); err != nil {
		return err
	}
	return nil
}
func (sourceTable *sourceTable) alias() string {
	if sourceTable == nil {
		return ""
	}
	return sourceTable.aliasName
}
func (sourceTable *sourceTable) isSourceBase() {}
func (sourceTable *sourceTable) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(sourceTable.tableName)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastModifierAs)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderAlias(sourceTable.aliasName)
	return nil
}
func (sourceTable *sourceTable) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(sourceTable.tableName); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(sourceTable.aliasName); err != nil {
		return err
	}
	return nil
}
func (sourceQuery *sourceQuery) alias() string {
	if sourceQuery == nil {
		return ""
	}
	return sourceQuery.aliasName
}
func (sourceQuery *sourceQuery) isSourceBase() {}
func (sourceQuery *sourceQuery) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeParenLeft)
	if err := sourceQuery.statement.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeParenRight)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastModifierAs)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderAlias(sourceQuery.aliasName)
	return nil
}
func (sourceQuery *sourceQuery) validate(baseValidator *baseValidator) error {
	if sourceQuery.statement == nil {
		return ErrInvalidSubquery
	}
	if err := sourceQuery.statement.validate(baseValidator); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(sourceQuery.aliasName); err != nil {
		return err
	}
	return nil
}
