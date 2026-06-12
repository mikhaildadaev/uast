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
func (validator *baseValidator) validateAlias(value string) error {
	length := len(value)
	if length > validator.config.lengthMaxIdent || length > uastSizeInitByte {
		return ErrOverflowIdentAlias
	}
	if !isSecureString(value, uastFormatAlias) {
		return ErrUnsupportSymbol
	}
	for i := 0; i < length; i++ {
		if value[i] >= 'a' && value[i] <= 'z' {
			validator.contexter.bufferByte[i] = value[i] & 0xDF
		} else {
			validator.contexter.bufferByte[i] = value[i]
		}
	}
	if _, exists := constKeywordUniversal[string(validator.contexter.bufferByte[:length])]; !exists {
		return nil
	}
	return ErrUnsupportIdentAlias
}
func (validator *baseValidator) validateArray(value int) error {
	if value > validator.config.lengthMaxArray {
		return ErrOverflowArray
	}
	return nil
}
func (validator *baseValidator) validateComparison(value transformComparison) error {
	validator.contexter.countMaxComparison++
	if validator.contexter.countMaxComparison > uastCountMaxComparison {
		return ErrExcessMaxComparison
	}
	left := value.getLeft()
	operator := value.getOperator()
	right := value.getRight()
	valueEnd := value.getValueEnd()
	valueStart := value.getValueStart()
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
	validator.contexter.prependCollectionComparison(value)
	return nil
}
func (validator *baseValidator) validateConstant(value any) error {
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
func (validator *baseValidator) validateFunction(value transformFunction) error {
	validator.contexter.countMaxFunction++
	if validator.contexter.countMaxFunction > uastCountMaxFunction {
		return ErrExcessMaxFunction
	}
	distinct := value.getDistinct()
	paramCount := value.getParamCount()
	service := value.getService()
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
		validator.contexter.prependCollectionFunction(value)
		return nil
	}
	return ErrUnsupportFunction
}
func (validator *baseValidator) validateName(value string) error {
	length := len(value)
	if length == 0 {
		return ErrInvalidIdentName
	}
	if length == 1 && value[0] == '*' {
		return nil
	}
	if length > validator.config.lengthMaxIdent || length > uastSizeInitByte {
		return ErrOverflowIdentName
	}
	if !isSecureString(value, uastFormatName) {
		return ErrUnsupportSymbol
	}
	for i := 0; i < length; i++ {
		if value[i] >= 'a' && value[i] <= 'z' {
			validator.contexter.bufferByte[i] = value[i] & 0xDF
		} else {
			validator.contexter.bufferByte[i] = value[i]
		}
	}
	if _, exists := constKeywordUniversal[string(validator.contexter.bufferByte[:length])]; !exists {
		return nil
	}
	return ErrUnsupportIdentName
}
func (validator *baseValidator) validateLiteral(value any) error {
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
		if length > validator.config.lengthMaxLimit {
			return ErrOverflowLiteral
		}
		if !isSecureString(v, uastFormatLiteral) {
			return ErrUnsupportSymbol
		}
		for i := 0; i < length; i++ {
			if v[i] >= 'a' && v[i] <= 'z' {
				validator.contexter.bufferByte[i] = v[i] & 0xDF
			} else {
				validator.contexter.bufferByte[i] = v[i]
			}
		}
		if _, exists := constKeywordUniversal[string(validator.contexter.bufferByte[:length])]; !exists {
			return nil
		}
	case time.Time:
		return nil
	}
	return ErrUnsupportLiteral
}
func (validator *baseValidator) validateOperator(value any) error {
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
func (validator *baseValidator) validateService(value any) error {
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
func (validator *baseValidator) validateSubquery() error {
	validator.contexter.countMaxSubquery++
	if validator.contexter.countMaxSubquery > uastCountMaxSubquery {
		return ErrExcessMaxSubquery
	}
	return nil
}
func (validator *baseValidator) validateValue(value any) error {
	switch v := value.(type) {
	case nil:
		return ErrInvalidValue
	case bool:
		return nil
	case []byte:
		length := len(v)
		if length > validator.config.lengthMaxValueByte {
			return ErrOverflowValueByte
		}
		return nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return nil
	case string:
		length := len(v)
		if length > validator.config.lengthMaxValueString {
			return ErrOverflowValueString
		}
		return nil
	case time.Time:
		return nil
	}
	return ErrUnsupportValue
}
func (validator *baseValidator) validateColumns(columns []markSourceable) error {
	if len(columns) == 0 {
		return nil
	}
	for _, column := range columns {
		if err := column.validate(validator); err != nil {
			return err
		}
	}
	return nil
}
func (validator *baseValidator) validateCommand(command managementService) error {
	return nil
}
func (validator *baseValidator) validateEntity(entity SourceBase) error {
	if entity == nil {
		return ErrInvalidStatement
	}
	switch e := entity.(type) {
	case *sourceIndex:
		return e.validate(validator)
	case *sourceSchema:
		return e.validate(validator)
	case *sourceTable:
		return e.validate(validator)
	case *sourceView:
		return e.validate(validator)
	default:
		return ErrInvalidStatement
	}
}
func (validator *baseValidator) validateFields(fields []markExpressable) error {
	if len(fields) == 0 {
		return ErrInvalidStatementField
	}
	for _, field := range fields {
		if err := field.validate(validator); err != nil {
			return err
		}
	}
	return nil
}
func (validator *baseValidator) validateFrom(from SourceBase) error {
	if from == nil {
		return ErrInvalidStatementFrom
	}
	if err := from.validate(validator); err != nil {
		return err
	}
	return nil
}
func (validator *baseValidator) validateGroupBy(groups []markGroupable) error {
	if len(groups) == 0 {
		return nil
	}
	for _, group := range groups {
		if err := group.validate(validator); err != nil {
			return err
		}
	}
	return nil
}
func (validator *baseValidator) validateHaving(having ExpressionBase) error {
	if having == nil {
		return nil
	}
	if err := having.validate(validator); err != nil {
		return err
	}
	return nil
}
func (validator *baseValidator) validateInto(into SourceBase) error {
	if into == nil {
		return ErrInvalidStatementInto
	}
	if err := into.validate(validator); err != nil {
		return err
	}
	return nil
}
func (validator *baseValidator) validateIsData(data string) error {
	if data == "" {
		return nil
	}
	if err := validator.validateLiteral(data); err != nil {
		return err
	}
	return nil
}
func (validator *baseValidator) validateJoin(joins []*clauseJoin) error {
	if len(joins) == 0 {
		return nil
	}
	for _, join := range joins {
		if err := join.validate(validator); err != nil {
			return err
		}
	}
	return nil
}
func (validator *baseValidator) validatePagination(pagination *clausePagination) error {
	if pagination == nil {
		return nil
	}
	if pagination.valueLimit > validator.config.lengthMaxLimit {
		return ErrOverflowLimit
	}
	if pagination.valueLimit > uastCountMaxLimit {
		return ErrExcessMaxLimit
	}
	if pagination.valueLimit < 0 {
		return ErrInvalidStatementLimit
	}
	if pagination.valueOffset < 0 {
		return ErrInvalidStatementOffset
	}
	return nil
}
func (validator *baseValidator) validateOn(on SourceBase) error {
	if on == nil {
		return nil
	}
	if err := on.validate(validator); err != nil {
		return err
	}
	return nil
}
func (validator *baseValidator) validateOnto(onto SourceBase) error {
	if onto == nil {
		return ErrInvalidStatementOnto
	}
	if err := onto.validate(validator); err != nil {
		return err
	}
	return nil
}
func (validator *baseValidator) validateOrderBy(orders []markOrderable) error {
	if len(orders) == 0 {
		return nil
	}
	for _, order := range orders {
		if err := order.validate(validator); err != nil {
			return err
		}
	}
	return nil
}
func (validator *baseValidator) validateReturning(returnings *clauseReturning) error {
	if returnings == nil {
		return nil
	}
	for _, expression := range returnings.expressions {
		if err := expression.validate(validator); err != nil {
			return err
		}
	}
	return nil
}
func (validator *baseValidator) validateSet(sets []*clauseSet) error {
	if len(sets) == 0 {
		return ErrInvalidStatementSet
	}
	for _, set := range sets {
		if err := set.validate(validator); err != nil {
			return err
		}
	}
	return nil
}
func (validator *baseValidator) validateSource(source statement) error {
	if source == nil {
		return nil
	}
	if err := source.validate(validator); err != nil {
		return err
	}
	return nil
}
func (validator *baseValidator) validateTable(table *TableSource) error {
	if table == nil {
		return ErrInvalidStatementFrom
	}
	if err := table.validate(validator); err != nil {
		return err
	}
	return nil
}
func (validator *baseValidator) validateUnions(unions []*clauseUnions) error {
	if len(unions) == 0 {
		return nil
	}
	validator.contexter.countMaxUnions += len(unions)
	if validator.contexter.countMaxUnions > uastCountMaxUnions {
		return ErrExcessMaxUnions
	}
	for _, union := range unions {
		if err := union.validate(validator); err != nil {
			return err
		}
	}
	return nil
}
func (validator *baseValidator) validateValues(values *clauseValues) error {
	if values == nil {
		return nil
	}
	for _, pair := range values.pairs {
		if err := pair.value.validate(validator); err != nil {
			return err
		}
	}
	if values.upsert != nil {
		if err := values.upsert.validate(validator); err != nil {
			return err
		}
	}
	return nil
}
func (validator *baseValidator) validateWhere(where ExpressionBase) error {
	if where == nil {
		return nil
	}
	if err := where.validate(validator); err != nil {
		return err
	}
	return nil
}
func (validator *baseValidator) validateWith(withs []*clauseWith) error {
	if len(withs) == 0 {
		return nil
	}
	validator.contexter.countMaxWith += len(withs)
	if validator.contexter.countMaxWith > uastCountMaxWith {
		return ErrExcessMaxWith
	}
	seen := make(map[string]bool)
	for _, with := range withs {
		if seen[with.alias] {
			return ErrDuplicateCTE
		}
		seen[with.alias] = true
		if err := with.validate(validator); err != nil {
			return err
		}
	}
	return nil
}
