package uast

import (
	"strconv"
	"time"
)

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
func (renderer *baseRenderer) renderAlias(value string) error {
	renderer.contexter.bufferQuery.WriteString(renderer.config.symbolQuoteLeft)
	renderer.contexter.bufferQuery.WriteString(value)
	renderer.contexter.bufferQuery.WriteString(renderer.config.symbolQuoteRight)
	return nil
}
func (renderer *baseRenderer) renderConstant(value any) error {
	switch v := value.(type) {
	case nil:
		return nil
	case bool:
		if v {
			renderer.contexter.bufferQuery.WriteString(uastConstStringTrue)
			return nil
		}
		renderer.contexter.bufferQuery.WriteString(uastConstStringFalse)
		return nil
	case float32:
		if v == uastConstFloat32One {
			renderer.contexter.bufferQuery.WriteString(strconv.FormatFloat(float64(v), 'f', 1, 32))
			return nil
		}
	case float64:
		if v == uastConstFloat64One {
			renderer.contexter.bufferQuery.WriteString(strconv.FormatFloat(v, 'f', 6, 64))
			return nil
		}
	case int:
		if v == uastConstIntOne {
			renderer.contexter.bufferQuery.WriteString(strconv.Itoa(v))
			return nil
		}
	case int8:
		if v == uastConstInt8One {
			renderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
			return nil
		}
	case int16:
		if v == uastConstInt16One {
			renderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
			return nil
		}
	case int32:
		if v == uastConstInt32One {
			renderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
			return nil
		}
	case int64:
		if v == uastConstInt64One {
			renderer.contexter.bufferQuery.WriteString(strconv.FormatInt(v, 10))
			return nil
		}
	case string:
		if v == uastConstStringDefault || v == uastConstStringNull {
			renderer.contexter.bufferQuery.WriteString(v)
			return nil
		}
	case uint:
		if v == uastConstUintOne {
			renderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
			return nil
		}
	case uint8:
		if v == uastConstUint8One {
			renderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
			return nil
		}
	case uint16:
		if v == uastConstUint16One {
			renderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
			return nil
		}
	case uint32:
		if v == uastConstUint32One {
			renderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
			return nil
		}
	case uint64:
		if v == uastConstUint64One {
			renderer.contexter.bufferQuery.WriteString(strconv.FormatUint(v, 10))
			return nil
		}
	}
	return ErrUnsupportConstant
}
func (renderer *baseRenderer) renderFunction(value string) error {
	renderer.contexter.bufferQuery.WriteString(value)
	return nil
}
func (renderer *baseRenderer) renderName(value string) error {
	renderer.contexter.bufferQuery.WriteString(renderer.config.symbolQuoteLeft)
	renderer.contexter.bufferQuery.WriteString(value)
	renderer.contexter.bufferQuery.WriteString(renderer.config.symbolQuoteRight)
	return nil
}
func (renderer *baseRenderer) renderLiteral(value any) error {
	switch v := value.(type) {
	case nil:
		return nil
	case bool:
		if v {
			renderer.contexter.bufferQuery.WriteString(uastConstStringTrue)
			return nil
		}
		renderer.contexter.bufferQuery.WriteString(uastConstStringFalse)
		return nil
	case float32:
		renderer.contexter.bufferQuery.WriteString(strconv.FormatFloat(float64(v), 'f', 1, 32))
		return nil
	case float64:
		renderer.contexter.bufferQuery.WriteString(strconv.FormatFloat(v, 'f', 6, 64))
		return nil
	case int:
		renderer.contexter.bufferQuery.WriteString(strconv.Itoa(v))
		return nil
	case int8:
		renderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
		return nil
	case int16:
		renderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
		return nil
	case int32:
		renderer.contexter.bufferQuery.WriteString(strconv.FormatInt(int64(v), 10))
		return nil
	case int64:
		renderer.contexter.bufferQuery.WriteString(strconv.FormatInt(v, 10))
		return nil
	case string:
		renderer.contexter.bufferQuery.WriteString(renderer.config.symbolMarkLeft)
		renderer.contexter.bufferQuery.WriteString(v)
		renderer.contexter.bufferQuery.WriteString(renderer.config.symbolMarkRight)
		return nil
	case time.Time:
		renderer.contexter.bufferQuery.WriteString(renderer.config.symbolMarkLeft)
		renderer.contexter.bufferQuery.WriteString(v.Format("2006-01-02 15:04:05"))
		renderer.contexter.bufferQuery.WriteString(renderer.config.symbolMarkRight)
		return nil
	case uint:
		renderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
		return nil
	case uint8:
		renderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
		return nil
	case uint16:
		renderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
		return nil
	case uint32:
		renderer.contexter.bufferQuery.WriteString(strconv.FormatUint(uint64(v), 10))
		return nil
	case uint64:
		renderer.contexter.bufferQuery.WriteString(strconv.FormatUint(v, 10))
		return nil
	}
	return nil
}
func (renderer *baseRenderer) renderOperator(value any) error {
	switch v := value.(type) {
	case string:
		renderer.contexter.bufferQuery.WriteString(v)
		return nil
	case binaryOperator:
		renderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case comparisonOperator:
		renderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case compositeOperator:
		renderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case joinOperator:
		renderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case logicalOperator:
		renderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case orderOperator:
		renderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case unionOperator:
		renderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	}
	return ErrUnsupportOperator
}
func (renderer *baseRenderer) renderService(value any) error {
	switch v := value.(type) {
	case string:
		renderer.contexter.bufferQuery.WriteString(v)
		return nil
	case functionService:
		renderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case managementService:
		renderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case modifierService:
		renderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	case typeService:
		renderer.contexter.bufferQuery.WriteString(string(v))
		return nil
	}
	return ErrUnsupportService
}
func (renderer *baseRenderer) renderValue(value any) error {
	renderer.contexter.bufferValue = append(renderer.contexter.bufferValue, value)
	renderer.contexter.bufferQuery.WriteString(renderer.config.placeholderStyle)
	if renderer.config.placeholderType {
		renderer.contexter.bufferQuery.WriteString(strconv.Itoa(renderer.config.placeholderNumber + len(renderer.contexter.bufferValue)))
	}
	return nil
}
func (renderer *baseRenderer) renderCascade(cascade bool) error {
	if cascade {
		renderer.renderOperator(uastCompositeSingleSpace)
		renderer.renderService(uastModifierCascade)
	}
	return nil
}
func (renderer *baseRenderer) renderCommand(command managementService) error {
	renderer.renderService(command)
	return nil
}
func (renderer *baseRenderer) renderDistinct(distinct bool) error {
	if distinct {
		renderer.renderOperator(uastCompositeSingleSpace)
		renderer.renderService(uastModifierDistinct)
	}
	return nil
}
func (renderer *baseRenderer) renderColumns(columns []markExpressable) error {
	if columns == nil {
		return nil
	}
	columnsCount := len(columns) - 1
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderOperator(uastCompositeParenLeft)
	for i, column := range columns {
		if err := column.render(renderer); err != nil {
			return err
		}
		if i < columnsCount {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	renderer.renderOperator(uastCompositeParenRight)
	return nil
}
func (renderer *baseRenderer) renderFields(fields []markExpressable) error {
	if fields == nil {
		return nil
	}
	fieldsCount := len(fields) - 1
	renderer.renderOperator(uastCompositeSingleSpace)
	for i, field := range fields {
		if err := field.render(renderer); err != nil {
			return err
		}
		if i < fieldsCount {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (renderer *baseRenderer) renderFrom(from SourceBase) error {
	if from == nil {
		return nil
	}
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementFrom)
	renderer.renderOperator(uastCompositeSingleSpace)
	if err := from.render(renderer); err != nil {
		return err
	}
	return nil
}
func (renderer *baseRenderer) renderGroupBy(groups []markGroupable) error {
	if groups == nil {
		return nil
	}
	groupsCount := len(groups) - 1
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementGroupBy)
	renderer.renderOperator(uastCompositeSingleSpace)
	for i, group := range groups {
		if err := group.render(renderer); err != nil {
			return err
		}
		if i < groupsCount {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (renderer *baseRenderer) renderHaving(having ExpressionBase) error {
	if having == nil {
		return nil
	}
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementHaving)
	renderer.renderOperator(uastCompositeSingleSpace)
	if err := having.render(renderer); err != nil {
		return err
	}
	return nil
}
func (renderer *baseRenderer) renderInto(into SourceBase) error {
	if into == nil {
		return nil
	}
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementInto)
	renderer.renderOperator(uastCompositeSingleSpace)
	if err := into.render(renderer); err != nil {
		return err
	}
	return nil
}
func (renderer *baseRenderer) renderJoin(joins []*clauseJoin) error {
	if joins == nil {
		return nil
	}
	joinsCount := len(joins) - 1
	renderer.renderOperator(uastCompositeSingleSpace)
	for i, join := range joins {
		if err := join.render(renderer); err != nil {
			return err
		}
		if i < joinsCount {
			renderer.renderOperator(uastCompositeSingleSpace)
		}
	}
	return nil
}
func (renderer *baseRenderer) renderLimit(limit *clauseLimit) error {
	if limit == nil {
		return nil
	}
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementLimit)
	renderer.renderOperator(uastCompositeSingleSpace)
	if err := limit.render(renderer); err != nil {
		return err
	}
	return nil
}
func (renderer *baseRenderer) renderOffset(offset *clauseOffset) error {
	if offset == nil {
		return nil
	}
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementOffset)
	renderer.renderOperator(uastCompositeSingleSpace)
	if err := offset.render(renderer); err != nil {
		return err
	}
	return nil
}
func (renderer *baseRenderer) renderOnto(onto SourceBase) error {
	if onto == nil {
		return nil
	}
	renderer.renderOperator(uastCompositeSingleSpace)
	if err := onto.render(renderer); err != nil {
		return err
	}
	return nil
}
func (renderer *baseRenderer) renderOrderBy(orders []markOrderable) error {
	if orders == nil {
		return nil
	}
	ordersCount := len(orders) - 1
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementOrderBy)
	renderer.renderOperator(uastCompositeSingleSpace)
	for i, order := range orders {
		if err := order.render(renderer); err != nil {
			return err
		}
		if i < ordersCount {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (renderer *baseRenderer) renderReindex(reindex bool) error {
	if reindex {
		renderer.renderOperator(uastCompositeSingleSpace)
		renderer.renderService(uastModifierReindex)
	}
	return nil
}
func (renderer *baseRenderer) renderReturning(returnings []markReturnable) error {
	if returnings == nil {
		return nil
	}
	returningsCount := len(returnings) - 1
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementReturning)
	renderer.renderOperator(uastCompositeSingleSpace)
	for i, returning := range returnings {
		if err := returning.render(renderer); err != nil {
			return err
		}
		if i < returningsCount {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (renderer *baseRenderer) renderSet(sets []*clauseSet) error {
	if sets == nil {
		return nil
	}
	setsCount := len(sets) - 1
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementSet)
	renderer.renderOperator(uastCompositeSingleSpace)
	for i, set := range sets {
		if err := set.render(renderer); err != nil {
			return err
		}
		if i < setsCount {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (renderer *baseRenderer) renderSource(source statement) error {
	if source == nil {
		return nil
	}
	renderer.renderOperator(uastCompositeSingleSpace)
	if err := source.render(renderer); err != nil {
		return err
	}
	return nil
}
func (renderer *baseRenderer) renderTable(table *TableSource) error {
	if table == nil {
		return nil
	}
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementTable)
	renderer.renderOperator(uastCompositeSingleSpace)
	if err := table.render(renderer); err != nil {
		return err
	}
	return nil
}
func (renderer *baseRenderer) renderUnions(unions []*clauseUnions) error {
	if unions == nil {
		return nil
	}
	unionsCount := len(unions) - 1
	renderer.renderOperator(uastCompositeSingleSpace)
	for i, union := range unions {
		if err := union.render(renderer); err != nil {
			return err
		}
		if i < unionsCount {
			renderer.renderOperator(uastCompositeSingleSpace)
		}
	}
	return nil
}
func (renderer *baseRenderer) renderValues(values *clauseValues) error {
	if values == nil {
		return nil
	}
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementValues)
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderOperator(uastCompositeParenLeft)
	for i, pair := range values.pairs {
		if i > 0 {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
		if err := pair.value.render(renderer); err != nil {
			return err
		}
	}
	renderer.renderOperator(uastCompositeParenRight)
	if values.upsert != nil {
		values.upsert.render(renderer)
	}
	return nil
}
func (renderer *baseRenderer) renderWhere(where ExpressionBase) error {
	if where == nil {
		return nil
	}
	renderer.renderOperator(uastCompositeSingleSpace)
	renderer.renderService(uastManagementWhere)
	renderer.renderOperator(uastCompositeSingleSpace)
	if err := where.render(renderer); err != nil {
		return err
	}
	return nil
}
func (renderer *baseRenderer) renderWith(withs []*clauseWith) error {
	if withs == nil {
		return nil
	}
	renderer.renderService(uastManagementWith)
	renderer.renderOperator(uastCompositeSingleSpace)
	for _, with := range withs {
		if with.recursive {
			renderer.renderService(uastModifierRecursive)
			renderer.renderOperator(uastCompositeSingleSpace)
			break
		}
	}
	withCount := len(withs) - 1
	for i, with := range withs {
		if err := with.render(renderer); err != nil {
			return err
		}
		if i < withCount {
			renderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
