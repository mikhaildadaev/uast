package uast

import (
	"time"
)

// Приватные структуры
type baseValidator struct {
	config    *config
	contexter *contexter
	strateger strateger
}

// Приватные конструкторы
func newValidator(config *config, contexter *contexter, strateger strateger) *baseValidator {
	return &baseValidator{
		config:    config,
		contexter: contexter,
		strateger: strateger,
	}
}

// Приватные методы
func (baseValidator *baseValidator) validateAlias(value string) error {
	length := len(value)
	if length > baseValidator.config.lengthMaxIdent || length > uastSizeInitByte {
		return ErrOverflowIdentAlias
	}
	if !isSecureString(value, uastFormatAlias) {
		return ErrUnsupportSymbol
	}
	for i := 0; i < length; i++ {
		if value[i] >= 'a' && value[i] <= 'z' {
			baseValidator.contexter.bufferByte[i] = value[i] & 0xDF
		} else {
			baseValidator.contexter.bufferByte[i] = value[i]
		}
	}
	if _, exists := constKeywordUniversal[string(baseValidator.contexter.bufferByte[:length])]; !exists {
		return nil
	}
	return ErrUnsupportIdentAlias
}
func (baseValidator *baseValidator) validateArray(value int) error {
	if value > baseValidator.config.lengthMaxArray {
		return ErrOverflowArray
	}
	return nil
}
func (baseValidator *baseValidator) validateComparison(value transformComparison) error {
	baseValidator.contexter.countMaxComparison++
	if baseValidator.contexter.countMaxComparison > uastCountMaxComparison {
		return ErrExcessMaxComparison
	}
	left := value.transformGetLeft()
	operator := value.transformGetOperator()
	right := value.transformGetRight()
	valueEnd := value.transformGetValueEnd()
	valueStart := value.transformGetValueStart()
	switch operator {
	case uastComparisonBetween, uastComparisonNotBetween:
		if left == nil || valueEnd == nil || valueStart == nil {
			return ErrInvalidComparisonBetweenNotBetween
		}
	case uastComparisonExists, uastComparisonNotExists:
		if left == nil {
			return ErrInvalidComparisonExistsNotExists
		}
	case uastComparisonIn, uastComparisonNotIn:
		if left == nil || right == nil {
			return ErrInvalidComparisonInNotIn
		}
	case uastComparisonIsNull, uastComparisonIsNotNull:
		if left == nil {
			return ErrInvalidComparisonIsNullIsNotNull
		}
	case uastComparisonILike, uastComparisonNotILike:
		if left == nil || right == nil {
			return ErrInvalidComparisonILikeNotILike
		}
	case uastComparisonLike, uastComparisonNotLike:
		if left == nil || right == nil {
			return ErrInvalidComparisonLikeNotLike
		}
	}
	baseValidator.contexter.appendCollectionComparison(value)
	return nil
}
func (baseValidator *baseValidator) validateConstant(value any) error {
	switch v := value.(type) {
	case nil:
		return nil
	case bool:
		return nil
	case float32:
		if v == uastConstFloat32One {
			return nil
		}
	case float64:
		if v == uastConstFloat64One {
			return nil
		}
	case int:
		if v == uastConstIntOne {
			return nil
		}
	case int8:
		if v == uastConstInt8One {
			return nil
		}
	case int16:
		if v == uastConstInt16One {
			return nil
		}
	case int32:
		if v == uastConstInt32One {
			return nil
		}
	case int64:
		if v == uastConstInt64One {
			return nil
		}
	case string:
		if v == uastConstStringDefault {
			return nil
		}
	case uint:
		if v == uastConstUintOne {
			return nil
		}
	case uint8:
		if v == uastConstUint8One {
			return nil
		}
	case uint16:
		if v == uastConstUint16One {
			return nil
		}
	case uint32:
		if v == uastConstUint32One {
			return nil
		}
	case uint64:
		if v == uastConstUint64One {
			return nil
		}
	}
	return ErrUnsupportConstant
}
func (baseValidator *baseValidator) validateFunction(value transformFunction) error {
	baseValidator.contexter.countMaxFunction++
	if baseValidator.contexter.countMaxFunction > uastCountMaxFunction {
		return ErrExcessMaxFunction
	}
	distinct := value.transformGetDistinct()
	paramCount := value.transformGetParamCount()
	service := value.transformGetService()
	if i, exists := constFunctionParameters[functionService(service)]; exists {
		if distinct && !i.distinct {
			return ErrUnsupportFunctionDistinct
		}
		if i.min != -1 && paramCount < i.min {
			return ErrUnsupportFunctionParamMin
		}
		if i.max != -1 && paramCount > i.max {
			return ErrUnsupportFunctionParamMax
		}
		// !!! Внимание - Проверить порядок добавления (дети - родитель)
		baseValidator.contexter.appendCollectionFunction(value)
		return nil
	}
	return ErrUnsupportFunction
}
func (baseValidator *baseValidator) validateName(value string) error {
	length := len(value)
	if length == 0 {
		return ErrInvalidIdentName
	}
	if length == 1 && value[0] == '*' {
		return nil
	}
	if length > baseValidator.config.lengthMaxIdent || length > uastSizeInitByte {
		return ErrOverflowIdentName
	}
	if !isSecureString(value, uastFormatName) {
		return ErrUnsupportSymbol
	}
	for i := 0; i < length; i++ {
		if value[i] >= 'a' && value[i] <= 'z' {
			baseValidator.contexter.bufferByte[i] = value[i] & 0xDF
		} else {
			baseValidator.contexter.bufferByte[i] = value[i]
		}
	}
	if _, exists := constKeywordUniversal[string(baseValidator.contexter.bufferByte[:length])]; !exists {
		return nil
	}
	return ErrUnsupportIdentName
}
func (baseValidator *baseValidator) validateLiteral(value any) error {
	switch v := value.(type) {
	case nil:
		return nil
	case bool:
		return nil
	case float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return nil
	case string:
		length := len(v)
		if length == 0 {
			return ErrInvalidLiteral
		}
		if length > baseValidator.config.lengthMaxLimit {
			return ErrOverflowLiteral
		}
		if !isSecureString(v, uastFormatLiteral) {
			return ErrUnsupportSymbol
		}
		for i := 0; i < length; i++ {
			if v[i] >= 'a' && v[i] <= 'z' {
				baseValidator.contexter.bufferByte[i] = v[i] & 0xDF
			} else {
				baseValidator.contexter.bufferByte[i] = v[i]
			}
		}
		if _, exists := constKeywordUniversal[string(baseValidator.contexter.bufferByte[:length])]; !exists {
			return nil
		}
	case time.Time:
		return nil
	}
	return ErrUnsupportLiteral
}
func (baseValidator *baseValidator) validateOperator(value any) error {
	switch v := value.(type) {
	case nil:
		return ErrInvalidOperator
	case binaryOperator:
		if constBinaryOperators[v] == "" {
			return ErrInvalidOperatorBinary
		}
		return nil
	case comparisonOperator:
		if constComparisonOperators[v] == "" {
			return ErrInvalidOperatorComparison
		}
		return nil
	case compositeOperator:
		if constCompositeOperators[v] == "" {
			return ErrInvalidOperatorComposite
		}
		return nil
	case joinOperator:
		if constJoinOperators[v] == "" {
			return ErrInvalidOperatorJoin
		}
		return nil
	case logicalOperator:
		if constLogicalOperators[v] == "" {
			return ErrInvalidOperatorLogical
		}
		return nil
	case orderOperator:
		if constOrderOperators[v] == "" {
			return ErrInvalidOperatorOrder
		}
		return nil
	case unionOperator:
		if constUnionOperators[v] == "" {
			return ErrInvalidOperatorUnion
		}
		return nil
	}
	return ErrUnsupportOperator
}
func (baseValidator *baseValidator) validateService(value any) error {
	switch v := value.(type) {
	case nil:
		return ErrInvalidService
	case functionService:
		if constFunctionServices[v] == "" {
			return ErrInvalidServiceFunction
		}
		return nil
	case managementService:
		if constManagementServices[v] == "" {
			return ErrInvalidServiceManagement
		}
		return nil
	case modifierService:
		if constModifierServices[v] == "" {
			return ErrInvalidServiceModifier
		}
		return nil
	}
	return ErrUnsupportService
}
func (baseValidator *baseValidator) validateSubquery() error {
	baseValidator.contexter.countMaxSubquery++
	if baseValidator.contexter.countMaxSubquery > uastCountMaxSubquery {
		return ErrExcessMaxSubquery
	}
	return nil
}
func (baseValidator *baseValidator) validateValue(value any) error {
	switch v := value.(type) {
	case nil:
		return ErrInvalidValue
	case bool:
		return nil
	case []byte:
		length := len(v)
		if length > baseValidator.config.lengthMaxValueByte {
			return ErrOverflowValueByte
		}
		return nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return nil
	case string:
		length := len(v)
		if length > baseValidator.config.lengthMaxValueString {
			return ErrOverflowValueString
		}
		return nil
	case time.Time:
		return nil
	}
	return ErrUnsupportValue
}
func (baseValidator *baseValidator) validateColumn(columns []markColumnable) error {
	if columns == nil {
		return nil
	}
	for _, column := range columns {
		if err := column.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
func (baseValidator *baseValidator) validateCommand(command managementService) error {
	return nil
}
func (baseValidator *baseValidator) validateField(fields []markFieldable) error {
	if fields == nil {
		return ErrInvalidStatementField
	}
	for _, field := range fields {
		if err := field.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
func (baseValidator *baseValidator) validateFrom(from SourceBase) error {
	if from == nil {
		return ErrInvalidStatementFrom
	}
	if err := from.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
func (baseValidator *baseValidator) validateGroupBy(groups []markGroupable) error {
	if groups == nil {
		return nil
	}
	for _, group := range groups {
		if err := group.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
func (baseValidator *baseValidator) validateHaving(having ExpressionBase) error {
	if having == nil {
		return nil
	}
	if err := having.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
func (baseValidator *baseValidator) validateInto(into SourceBase) error {
	if into == nil {
		return ErrInvalidStatementInto
	}
	if err := into.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
func (baseValidator *baseValidator) validateJoin(joins []*clauseJoin) error {
	if joins == nil {
		return nil
	}
	for _, join := range joins {
		if err := join.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
func (baseValidator *baseValidator) validateLimit(limit *clauseLimit) error {
	if limit == nil {
		return nil
	}
	if limit.value > baseValidator.config.lengthMaxLimit {
		return ErrOverflowLimit
	}
	if err := limit.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
func (baseValidator *baseValidator) validateOffset(offset *clauseOffset) error {
	if offset == nil {
		return nil
	}
	if err := offset.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
func (baseValidator *baseValidator) validateOnto(onto SourceBase) error {
	if onto == nil {
		return ErrInvalidStatementOnto
	}
	if err := onto.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
func (baseValidator *baseValidator) validateOrderBy(orders []markOrderable) error {
	if orders == nil {
		return nil
	}
	for _, order := range orders {
		if err := order.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
func (baseValidator *baseValidator) validateReturning(returnings []markReturnable) error {
	if returnings == nil {
		return nil
	}
	for _, returning := range returnings {
		if err := returning.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
func (baseValidator *baseValidator) validateSet(sets []*clauseSet) error {
	if sets == nil {
		return ErrInvalidStatementSet
	}
	for _, set := range sets {
		if err := set.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
func (baseValidator *baseValidator) validateSource(source statement) error {
	if source == nil {
		return nil
	}
	if err := source.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
func (baseValidator *baseValidator) validateUnions(unions []*clauseUnions) error {
	if unions == nil {
		return nil
	}
	baseValidator.contexter.countMaxUnions += len(unions)
	if baseValidator.contexter.countMaxUnions > uastCountMaxUnions {
		return ErrExcessMaxUnions
	}
	for _, union := range unions {
		if err := union.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
func (baseValidator *baseValidator) validateValues(values [][]ExpressionBase) error {
	if values == nil {
		return nil
	}
	for _, part := range values {
		for _, value := range part {
			if err := value.validate(baseValidator); err != nil {
				return err
			}
		}
	}
	return nil
}
func (baseValidator *baseValidator) validateWhere(where ExpressionBase) error {
	if where == nil {
		return nil
	}
	if err := where.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
func (baseValidator *baseValidator) validateWith(withs []*clauseWith) error {
	if withs == nil {
		return nil
	}
	baseValidator.contexter.countMaxWith += len(withs)
	if baseValidator.contexter.countMaxWith > uastCountMaxWith {
		return ErrExcessMaxWith
	}
	seen := make(map[string]bool)
	for _, with := range withs {
		if seen[with.alias] {
			return ErrDuplicateCTE
		}
		seen[with.alias] = true
		if err := with.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
