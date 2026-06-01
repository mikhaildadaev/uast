package uast

import (
	"strconv"
	"sync/atomic"
)

// Публичные интерфейсы
type SourceBase interface {
	alias() string
	clone() SourceBase
	name() string
	isSourceBase()
	render(baseRenderer *baseRenderer) error
	validate(baseValidator *baseValidator) error
}

// Публичные структуры
type CteSource = sourceCte
type IndexSource = sourceIndex
type QuerySource = sourceQuery
type SchemaSource = sourceSchema
type TableSource = sourceTable
type ViewSource = sourceView

// Публичные конструкторы
func NewCTE(name string, alias string) *CteSource {
	if alias == "" {
		alias = name
	}
	return &CteSource{
		aliasName: alias,
		cteName:   name,
		withAlias: true,
	}
}
func NewIndex(name string) *IndexSource {
	return &IndexSource{
		indexName: name,
	}
}
func NewQuery(statement statement, alias string) *QuerySource {
	if alias == "" {
		alias = "query_" + strconv.FormatInt(queryCounter.Add(1), 10)
	}
	return &QuerySource{
		aliasName: alias,
		statement: statement,
		withAlias: true,
	}
}
func NewSchema(name string) *SchemaSource {
	return &SchemaSource{
		schemaName: name,
	}
}
func NewTable(name, alias string) *TableSource {
	if alias == "" {
		alias = name
	}
	return &TableSource{
		aliasName: alias,
		tableName: name,
		withAlias: true,
	}
}
func NewView(name, alias string) *sourceView {
	if alias == "" {
		alias = name
	}
	return &sourceView{
		aliasName: alias,
		viewName:  name,
		withAlias: true,
	}
}

// Приватные переменные
var queryCounter atomic.Int64

// Приватные структуры
type sourceCte struct {
	aliasName string
	cteName   string
	withAlias bool
}
type sourceIndex struct {
	indexName string
}
type sourceQuery struct {
	aliasName string
	statement statement
	withAlias bool
}
type sourceSchema struct {
	schemaName string
}
type sourceTable struct {
	aliasName string
	tableName string
	withAlias bool
}
type sourceView struct {
	aliasName string
	viewName  string
	withAlias bool
}

// Приватные методы
func (source *sourceCte) alias() string {
	if source == nil {
		return ""
	}
	return source.aliasName
}
func (source *sourceCte) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceCte) isSourceBase() {}
func (source *sourceCte) name() string {
	return source.cteName
}
func (source *sourceCte) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(source.cteName)
	if source.withAlias {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierAs)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderAlias(source.aliasName)
	}
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
func (source *sourceIndex) alias() string {
	return ""
}
func (source *sourceIndex) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceIndex) name() string {
	return source.indexName
}
func (source *sourceIndex) isSourceBase() {}
func (source *sourceIndex) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(source.indexName)
	return nil
}
func (source *sourceIndex) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(source.indexName); err != nil {
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
func (source *sourceQuery) clone() SourceBase {
	copy := *source
	copy.statement = source.statement.clone()
	return &copy
}
func (source *sourceQuery) isSourceBase() {}
func (source *sourceQuery) name() string {
	return ""
}
func (source *sourceQuery) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeParenLeft)
	if err := source.statement.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeParenRight)
	if source.withAlias {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierAs)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderAlias(source.aliasName)
	}
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
func (source *sourceSchema) alias() string {
	return ""
}
func (source *sourceSchema) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceSchema) name() string {
	return source.schemaName
}
func (source *sourceSchema) isSourceBase() {}
func (source *sourceSchema) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(source.schemaName)
	return nil
}
func (source *sourceSchema) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(source.schemaName); err != nil {
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
func (source *sourceTable) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceTable) name() string {
	return source.tableName
}
func (source *sourceTable) isSourceBase() {}
func (source *sourceTable) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(source.tableName)
	if source.withAlias {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierAs)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderAlias(source.aliasName)
	}
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
func (source *sourceView) alias() string {
	if source == nil {
		return ""
	}
	return source.aliasName
}
func (source *sourceView) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceView) name() string {
	return source.viewName
}
func (source *sourceView) isSourceBase() {}
func (source *sourceView) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(source.viewName)
	if source.withAlias {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierAs)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderAlias(source.aliasName)
	}
	return nil
}
func (source *sourceView) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(source.viewName); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(source.aliasName); err != nil {
		return err
	}
	return nil
}
