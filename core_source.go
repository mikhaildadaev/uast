package uast

import (
	"strconv"
	"sync/atomic"
)

// Публичные интерфейсы
type SourceBase interface {
	clone() SourceBase
	getAlias() string
	getFormat() modifierService
	getName() string
	isSourceBase()
	render(baseRenderer *baseRenderer) error
	validate(baseValidator *baseValidator) error
}
type SourceSafe[T typeScalar] interface {
	SourceBase
	isSourceSafe(T)
}

// Публичные структуры
type ColumnSource[T typeScalar] = sourceColumn[T]
type CteSource = sourceCte
type IndexSource = sourceIndex
type QuerySource = sourceQuery
type SchemaSource = sourceSchema
type TableSource = sourceTable
type ViewSource = sourceView

// Публичные конструкторы
func NewColumn[T typeScalar](columnName string, table *TableSource, valueType ValueType) *ColumnSource[T] {
	return &ColumnSource[T]{
		field:     Field[T](table.aliasName, columnName),
		table:     table,
		valueType: valueType,
	}
}
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
func NewIndex(name string, table *TableSource) *IndexSource {
	return &IndexSource{
		indexName: name,
		table:     table,
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
func NewView(name, alias string, table *TableSource) *ViewSource {
	if alias == "" {
		alias = name
	}
	return &ViewSource{
		aliasName: alias,
		table:     table,
		viewName:  name,
		withAlias: true,
	}
}

// Публичные методы
func (source *sourceColumn[T]) AutoIncrement() *sourceColumn[T] {
	source.isAutoIncrement = true
	return source
}
func (source *sourceColumn[T]) Default(value T) *sourceColumn[T] {
	source.defaultValue = Value(value)
	return source
}
func (source *sourceColumn[T]) Expr() *exprField[T] {
	return source.field
}
func (source *sourceColumn[T]) NotNull() *sourceColumn[T] {
	source.isNotNull = true
	return source
}
func (source *sourceColumn[T]) SetDefault(value T) columnModifiable {
	return columnModifiable{
		column:    source,
		operation: uastModifierDefault,
		value:     Value(value),
	}
}
func (source *sourceColumn[T]) SetNotNull() columnModifiable {
	return columnModifiable{
		column:    source,
		operation: uastModifierNotNull,
	}
}
func (source *sourceColumn[T]) SetType(valueType ValueType) columnModifiable {
	return columnModifiable{
		column:    source,
		operation: uastModifierType,
		valueType: valueType,
	}
}
func (source *sourceIndex) Unique() *sourceIndex {
	source.isUnique = true
	return source
}

// Приватные интерфейсы
type markSourceable interface {
	SourceBase
	isColumnable()
}
type registerStatement interface {
	register(statement)
}

// Приватные переменные
var queryCounter atomic.Int64

// Приватные структуры
type columnModifiable struct {
	column    SourceBase
	operation modifierService
	value     ExpressionBase
	valueType ValueType
}
type columnRename struct {
	column SourceBase
	name   string
}
type sourceColumn[T typeScalar] struct {
	defaultValue    ExpressionBase
	field           *exprField[T]
	isAutoIncrement bool
	isNotNull       bool
	stmt            *stmtCreate
	table           *sourceTable
	valueType       ValueType
}
type sourceCte struct {
	aliasName string
	cteName   string
	withAlias bool
}
type sourceIndex struct {
	indexName string
	isUnique  bool
	table     *sourceTable
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
	table     *sourceTable
	viewName  string
	withAlias bool
}

// Приватные методы
func (source *sourceColumn[T]) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceColumn[T]) getAlias() string {
	return source.table.getAlias()
}
func (source *sourceColumn[T]) getFormat() modifierService {
	return uastModifierColumn
}
func (source *sourceColumn[T]) getName() string {
	return source.field.getName()
}
func (source *sourceColumn[T]) isSourceBase() {}
func (source *sourceColumn[T]) isColumnable() {}
func (source *sourceColumn[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(source.getName())
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(baseRenderer.config.listTypes[source.valueType])
	for _, attr := range baseRenderer.config.orderSupportCreate {
		switch attr {
		case uastModifierAutoIncrement:
			if source.isAutoIncrement {
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderService(baseRenderer.config.listModifiers[uastModifierAutoIncrement])
			}
		case uastModifierNotNull:
			if source.isNotNull {
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderService(uastModifierNotNull)
			}
		case uastModifierDefault:
			if source.defaultValue != nil {
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderService(uastModifierDefault)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				source.defaultValue.render(baseRenderer)
			}
		}
	}
	return nil
}
func (col *sourceColumn[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(col.table.tableName); err != nil {
		return err
	}
	return nil
}
func (source *sourceCte) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceCte) getAlias() string {
	if source == nil {
		return ""
	}
	return source.aliasName
}
func (source *sourceCte) getFormat() modifierService {
	return uastModifierCTE
}
func (source *sourceCte) getName() string {
	return source.cteName
}
func (source *sourceCte) isSourceBase() {}
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
func (source *sourceIndex) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceIndex) getAlias() string {
	return ""
}
func (source *sourceIndex) getFormat() modifierService {
	return uastModifierIndex
}
func (source *sourceIndex) getName() string {
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
func (source *sourceQuery) clone() SourceBase {
	copy := *source
	copy.statement = source.statement.clone()
	return &copy
}
func (source *sourceQuery) getAlias() string {
	if source == nil {
		return ""
	}
	return source.aliasName
}
func (source *sourceQuery) getFormat() modifierService {
	return uastModifierQuery
}
func (source *sourceQuery) getName() string {
	return ""
}
func (source *sourceQuery) isSourceBase() {}
func (source *sourceColumn[T]) register(stmt *stmtCreate) {
	source.stmt = stmt
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
func (source *sourceSchema) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceSchema) getAlias() string {
	return ""
}
func (source *sourceSchema) getFormat() modifierService {
	return uastModifierSchema
}
func (source *sourceSchema) getName() string {
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
func (source *sourceTable) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceTable) getAlias() string {
	if source == nil {
		return ""
	}
	return source.aliasName
}
func (source *sourceTable) getFormat() modifierService {
	return uastModifierTable
}
func (source *sourceTable) getName() string {
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
func (source *sourceView) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceView) getAlias() string {
	if source == nil {
		return ""
	}
	return source.aliasName
}
func (source *sourceView) getFormat() modifierService {
	return uastModifierView
}
func (source *sourceView) getName() string {
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
