package uast

import (
	"strconv"
	"sync/atomic"
)

// Публичные интерфейсы
type SourceBase interface {
	alias() string
	clone() SourceBase
	format() modifierService
	name() string
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
func NewColumn[T typeScalar](columnName string, table *TableSource, valueType ValueType) *sourceColumn[T] {
	return &sourceColumn[T]{
		field:     Column[T](table.aliasName, columnName),
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
func NewView(name, alias string, table *TableSource) *sourceView {
	if alias == "" {
		alias = name
	}
	return &sourceView{
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
func (source *sourceColumn[T]) DefaultValue(value ExpressionBase) *sourceColumn[T] {
	source.defaultValue = value
	return source
}
func (source *sourceColumn[T]) Expr() *exprColumn[T] {
	return source.field
}
func (source *sourceColumn[T]) NotNull() *sourceColumn[T] {
	source.isNotNull = true
	return source
}
func (source *sourceColumn[T]) PrimaryKey() *sourceColumn[T] {
	source.isPrimaryKey = true
	return source
}
func (source *sourceColumn[T]) Unique() *sourceColumn[T] {
	source.isUnique = true
	return source
}

// Приватные интерфейсы
type markSourceable interface {
	SourceBase
	isColumnable()
}
type transformColumn interface {
	SourceBase
	transformGetAutoIncrement() bool
	transformGetDefaultValue() ExpressionBase
	transformGetNotNull() bool
	transformGetPrimaryKey() bool
	transformGetUnique() bool
	transformGetValueType() ValueType
}

// Приватные переменные
var queryCounter atomic.Int64

// Приватные структуры
type sourceColumn[T typeScalar] struct {
	defaultValue    ExpressionBase
	field           *exprColumn[T]
	isAutoIncrement bool
	isNotNull       bool
	isPrimaryKey    bool
	isUnique        bool
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
func (source *sourceColumn[T]) alias() string {
	return source.table.alias()
}
func (source *sourceColumn[T]) clone() SourceBase {
	copy := *source
	return &copy
}
func (source *sourceColumn[T]) format() modifierService {
	return uastModifierColumn
}
func (source *sourceColumn[T]) transformGetDefaultValue() ExpressionBase {
	return source.defaultValue
}
func (source *sourceColumn[T]) transformGetAutoIncrement() bool {
	return source.isAutoIncrement
}
func (source *sourceColumn[T]) transformGetNotNull() bool {
	return source.isNotNull
}
func (source *sourceColumn[T]) transformGetPrimaryKey() bool {
	return source.isPrimaryKey
}
func (source *sourceColumn[T]) transformGetUnique() bool {
	return source.isUnique
}
func (source *sourceColumn[T]) transformGetValueType() ValueType {
	return source.valueType
}
func (source *sourceColumn[T]) isSourceBase() {}
func (source *sourceColumn[T]) isColumnable() {}
func (source *sourceColumn[T]) name() string {
	return source.field.transformGetName()
}
func (source *sourceColumn[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(source.name())
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(baseRenderer.config.listTypes[source.valueType])
	for _, attr := range baseRenderer.config.supportAttrCreateOrder {
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
		case uastModifierPrimaryKey:
			if source.isPrimaryKey {
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderService(uastModifierPrimaryKey)
			}
		case uastModifierUnique:
			if source.isUnique {
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderService(uastModifierUnique)
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
func (source *sourceCte) format() modifierService {
	return uastModifierCTE
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
func (source *sourceIndex) format() modifierService {
	return uastModifierIndex
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
func (source *sourceQuery) format() modifierService {
	return uastModifierQuery
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
func (source *sourceSchema) format() modifierService {
	return uastModifierSchema
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
func (source *sourceTable) format() modifierService {
	return uastModifierTable
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
func (source *sourceView) format() modifierService {
	return uastModifierView
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
