package uast

import (
	"strconv"
	"time"
)

// Приватные интерфейсы
type renderer interface {
	elementRenderer
	componentRenderer
	statementRenderer
}
type elementRenderer interface {
	renderAlias(value string) error
	renderConstant(value any) error
	renderFunction(value string) error
	renderName(value string) error
	renderLiteral(value any) error
	renderOperator(value any) error
	renderService(value any) error
	renderValue(value any) error
}
type componentRenderer interface {
	renderCommand(command managementService) error
	renderColumn(columns []markColumnable) error
	renderDistinct(distinct bool) error
	renderField(fields []markFieldable) error
	renderFrom(from SourceBase) error
	renderGroupBy(groups []markGroupable) error
	renderHaving(having ExpressionBase) error
	renderInto(into SourceBase) error
	renderJoin(joins []*clauseJoin) error
	renderLimit(limit *clauseLimit) error
	renderOffset(offset *clauseOffset) error
	renderOnto(onto SourceBase) error
	renderOrderBy(orders []markOrderable) error
	renderReturning(returnings []markReturnable) error
	renderSet(sets []*clauseSet) error
	renderSource(source statement) error
	renderUnions(unions []*clauseUnions) error
	renderValues(values [][]ExpressionBase) error
	renderWhere(where ExpressionBase) error
	renderWith(withs []*clauseWith) error
}
type statementRenderer interface {
	renderDelete(baseRenderer *baseRenderer, stmtDelete *stmtDelete) error
	renderInsert(baseRenderer *baseRenderer, stmtInsert *stmtInsert) error
	renderSelect(baseRenderer *baseRenderer, stmtSelect *stmtSelect) error
	renderUpdate(baseRenderer *baseRenderer, stmtUpdate *stmtUpdate) error
}

// Приватные структуры
type baseRenderer struct {
	config    *config
	contexter *contexter
	strateger strateger
}

// Приватные конструкторы
func newRenderer(config *config, contexter *contexter, strateger strateger) *baseRenderer {
	return &baseRenderer{
		config:    config,
		contexter: contexter,
		strateger: strateger,
	}
}

// Приватные методы
func (baseRenderer *baseRenderer) renderAlias(value string) error {
	baseRenderer.contexter.bufferQuery.WriteString(baseRenderer.config.symbolQuoteLeft)
	baseRenderer.contexter.bufferQuery.WriteString(value)
	baseRenderer.contexter.bufferQuery.WriteString(baseRenderer.config.symbolQuoteRight)
	return nil
}
func (baseRenderer *baseRenderer) renderConstant(value any) error {
	switch v := value.(type) {
	case nil:
		baseRenderer.contexter.bufferQuery.WriteString(uastConstStringNull)
		return nil
	case bool:
		if v {
			baseRenderer.contexter.bufferQuery.WriteString(uastConstStringTrue)
			return nil
		}
		baseRenderer.contexter.bufferQuery.WriteString(uastConstStringFalse)
		return nil
	case float32:
		if v == uastConstFloat32One {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatFloat(float64(v), 'f', -1, 32))
			return nil
		}
	case float64:
		if v == uastConstFloat64One {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
			return nil
		}
	case int:
		if v == uastConstIntOne {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.Itoa(v))
			return nil
		}
	case int8:
		if v == uastConstInt8One {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
			return nil
		}
	case int16:
		if v == uastConstInt16One {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
			return nil
		}
	case int32:
		if v == uastConstInt32One {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
			return nil
		}
	case int64:
		if v == uastConstInt64One {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatInt(v, 10))
			return nil
		}
	case string:
		if v == uastConstStringDefault {
			baseRenderer.contexter.bufferQuery.WriteString(v)
			return nil
		}
	case uint:
		if v == uastConstUintOne {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
			return nil
		}
	case uint8:
		if v == uastConstUint8One {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
			return nil
		}
	case uint16:
		if v == uastConstUint16One {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
			return nil
		}
	case uint32:
		if v == uastConstUint32One {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
			return nil
		}
	case uint64:
		if v == uastConstUint64One {
			baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatUint(v, 10))
			return nil
		}
	}
	return ErrUnsupportConstant
}
func (baseRenderer *baseRenderer) renderFunction(value string) error {
	baseRenderer.contexter.bufferQuery.WriteString(value)
	return nil
}
func (baseRenderer *baseRenderer) renderName(value string) error {
	baseRenderer.contexter.bufferQuery.WriteString(baseRenderer.config.symbolQuoteLeft)
	baseRenderer.contexter.bufferQuery.WriteString(value)
	baseRenderer.contexter.bufferQuery.WriteString(baseRenderer.config.symbolQuoteRight)
	return nil
}
func (baseRenderer *baseRenderer) renderLiteral(value any) error {
	switch v := value.(type) {
	case nil:
		baseRenderer.contexter.bufferQuery.WriteString(uastConstStringNull)
		return nil
	case bool:
		if v {
			baseRenderer.contexter.bufferQuery.WriteString(uastConstStringTrue)
			return nil
		}
		baseRenderer.contexter.bufferQuery.WriteString(uastConstStringFalse)
		return nil
	case float32:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatFloat(float64(v), 'f', -1, 32))
		return nil
	case float64:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		return nil
	case int:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.Itoa(v))
		return nil
	case int8:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
		return nil
	case int16:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
		return nil
	case int32:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
		return nil
	case int64:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatInt(v, 10))
		return nil
	case string:
		baseRenderer.contexter.bufferQuery.WriteString(baseRenderer.config.symbolMarkLeft)
		baseRenderer.contexter.bufferQuery.WriteString(v)
		baseRenderer.contexter.bufferQuery.WriteString(baseRenderer.config.symbolMarkRight)
		return nil
	case time.Time:
		baseRenderer.contexter.bufferQuery.WriteString(baseRenderer.config.symbolMarkLeft)
		baseRenderer.contexter.bufferQuery.WriteString(v.Format("2006-01-02 15:04:05"))
		baseRenderer.contexter.bufferQuery.WriteString(baseRenderer.config.symbolMarkRight)
		return nil
	case uint:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
		return nil
	case uint8:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
		return nil
	case uint16:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
		return nil
	case uint32:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
		return nil
	case uint64:
		baseRenderer.contexter.bufferQuery.WriteString(strconv.FormatUint(v, 10))
		return nil
	}
	return nil
}
func (baseRenderer *baseRenderer) renderOperator(value any) error {
	switch v := value.(type) {
	case string:
		baseRenderer.contexter.bufferQuery.WriteString(v)
		return nil
	case binaryOperator:
		baseRenderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case comparisonOperator:
		baseRenderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case compositeOperator:
		baseRenderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case joinOperator:
		baseRenderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case logicalOperator:
		baseRenderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case orderOperator:
		baseRenderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case unionOperator:
		baseRenderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	}
	return ErrUnsupportOperator
}
func (baseRenderer *baseRenderer) renderService(value any) error {
	switch v := value.(type) {
	case string:
		baseRenderer.contexter.bufferQuery.WriteString(v)
		return nil
	case functionService:
		baseRenderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case managementService:
		baseRenderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case modifierService:
		baseRenderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case typeService:
		baseRenderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	}
	return ErrUnsupportService
}
func (baseRenderer *baseRenderer) renderValue(value any) error {
	baseRenderer.contexter.bufferValue = append(baseRenderer.contexter.bufferValue, value)
	baseRenderer.contexter.bufferQuery.WriteString(baseRenderer.config.placeholderStyle)
	if baseRenderer.config.placeholderType {
		baseRenderer.contexter.bufferQuery.WriteString(strconv.Itoa(baseRenderer.config.placeholderNumber + len(baseRenderer.contexter.bufferValue)))
	}
	return nil
}
func (baseRenderer *baseRenderer) renderCommand(command managementService) error {
	baseRenderer.renderService(command)
	return nil
}
func (baseRenderer *baseRenderer) renderDistinct(distinct bool) error {
	if distinct {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierDistinct)
	}
	return nil
}
func (baseRenderer *baseRenderer) renderColumn(columns []markColumnable) error {
	if columns == nil {
		return nil
	}
	columnsCount := len(columns) - 1
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderOperator(uastCompositeParenLeft)
	for i, column := range columns {
		if err := column.render(baseRenderer); err != nil {
			return err
		}
		if i < columnsCount {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	baseRenderer.renderOperator(uastCompositeParenRight)
	return nil
}
func (baseRenderer *baseRenderer) renderField(fields []markFieldable) error {
	if fields == nil {
		return nil
	}
	fieldsCount := len(fields) - 1
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	for i, field := range fields {
		if err := field.render(baseRenderer); err != nil {
			return err
		}
		if i < fieldsCount {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (baseRenderer *baseRenderer) renderFrom(from SourceBase) error {
	if from == nil {
		return nil
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastManagementFrom)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := from.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (baseRenderer *baseRenderer) renderGroupBy(groups []markGroupable) error {
	if groups == nil {
		return nil
	}
	groupsCount := len(groups) - 1
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastManagementGroupBy)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	for i, group := range groups {
		if err := group.render(baseRenderer); err != nil {
			return err
		}
		if i < groupsCount {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (baseRenderer *baseRenderer) renderHaving(having ExpressionBase) error {
	if having == nil {
		return nil
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastManagementHaving)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := having.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (baseRenderer *baseRenderer) renderInto(into SourceBase) error {
	if into == nil {
		return nil
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastManagementInto)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := into.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (baseRenderer *baseRenderer) renderJoin(joins []*clauseJoin) error {
	if joins == nil {
		return nil
	}
	joinsCount := len(joins) - 1
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	for i, join := range joins {
		if err := join.render(baseRenderer); err != nil {
			return err
		}
		if i < joinsCount {
			baseRenderer.renderOperator(uastCompositeSingleSpace)
		}
	}
	return nil
}
func (baseRenderer *baseRenderer) renderLimit(limit *clauseLimit) error {
	if limit == nil {
		return nil
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastManagementLimit)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := limit.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (baseRenderer *baseRenderer) renderOffset(offset *clauseOffset) error {
	if offset == nil {
		return nil
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastManagementOffset)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := offset.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (baseRenderer *baseRenderer) renderOnto(onto SourceBase) error {
	if onto == nil {
		return nil
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := onto.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (baseRenderer *baseRenderer) renderOrderBy(orders []markOrderable) error {
	if orders == nil {
		return nil
	}
	ordersCount := len(orders) - 1
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastManagementOrderBy)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	for i, order := range orders {
		if err := order.render(baseRenderer); err != nil {
			return err
		}
		if i < ordersCount {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (baseRenderer *baseRenderer) renderReturning(returnings []markReturnable) error {
	if returnings == nil {
		return nil
	}
	returningsCount := len(returnings) - 1
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastManagementReturning)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	for i, returning := range returnings {
		if err := returning.render(baseRenderer); err != nil {
			return err
		}
		if i < returningsCount {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (baseRenderer *baseRenderer) renderSet(sets []*clauseSet) error {
	if sets == nil {
		return nil
	}
	setsCount := len(sets) - 1
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastManagementSet)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	for i, set := range sets {
		if err := set.render(baseRenderer); err != nil {
			return err
		}
		if i < setsCount {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (baseRenderer *baseRenderer) renderSource(source statement) error {
	if source == nil {
		return nil
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := source.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (baseRenderer *baseRenderer) renderUnions(unions []*clauseUnions) error {
	if unions == nil {
		return nil
	}
	unionsCount := len(unions) - 1
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	for i, union := range unions {
		if err := union.render(baseRenderer); err != nil {
			return err
		}
		if i < unionsCount {
			baseRenderer.renderOperator(uastCompositeSingleSpace)
		}
	}
	return nil
}
func (baseRenderer *baseRenderer) renderValues(values [][]ExpressionBase) error {
	if values == nil {
		return nil
	}
	valuesCount := len(values) - 1
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastManagementValues)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	for i, part := range values {
		baseRenderer.renderOperator(uastCompositeParenLeft)
		partCount := len(part) - 1
		for j, value := range part {
			if err := value.render(baseRenderer); err != nil {
				return err
			}
			if j < partCount {
				baseRenderer.renderOperator(uastCompositeCommaSpace)
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
		if i < valuesCount {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (baseRenderer *baseRenderer) renderWhere(where ExpressionBase) error {
	if where == nil {
		return nil
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderService(uastManagementWhere)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := where.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (baseRenderer *baseRenderer) renderWith(withs []*clauseWith) error {
	if withs == nil {
		return nil
	}
	baseRenderer.renderService(uastManagementWith)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	for _, with := range withs {
		if with.recursive {
			baseRenderer.renderService(uastModifierRecursive)
			baseRenderer.renderOperator(uastCompositeSingleSpace)
			break
		}
	}
	withCount := len(withs) - 1
	for i, with := range withs {
		if err := with.render(baseRenderer); err != nil {
			return err
		}
		if i < withCount {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
