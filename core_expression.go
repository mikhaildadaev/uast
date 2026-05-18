package uast

// Приватные интерфейсы
type markExpressable interface {
	ExpressionBase
	isFieldable()
}
type markGroupable interface {
	ExpressionBase
	isGroupable()
}
type markOrderable interface {
	ExpressionBase
	isOrderable()
}
type markPredicable interface {
	ExpressionBase
	isPredicable()
}
type markReturnable interface {
	ExpressionBase
	isReturnable()
}
type transformComparison interface {
	ExpressionBase
	transformGetLeft() ExpressionBase
	transformGetOperator() comparisonOperator
	transformGetRight() ExpressionBase
	transformGetValueEnd() ExpressionBase
	transformGetValueStart() ExpressionBase
}
type transformFunction interface {
	ExpressionBase
	transformGetDistinct() bool
	transformGetJson() []*exprJson
	transformGetLeft() ExpressionBase
	transformGetParamCount() int
	transformGetProcess() processingStage
	transformGetRight() ExpressionBase
	transformGetService() functionService
	transformGetValueArray() []ExpressionBase
	transformGetValueType() ValueType
	transformSetJson([]*exprJson)
	transformSetLeft(left ExpressionBase)
	transformSetOperator(operator compositeOperator)
	transformSetProcess(format processingStage)
	transformSetRight(right ExpressionBase)
	transformSetService(service functionService)
	transformSetValueArray(valueArray []ExpressionBase)
}

// Приватные структуры
type exprAlias[T typeScalar] struct {
	aliasName  string
	expression ExpressionSafe[T]
}
type exprArray[T typeScalar] struct {
	array []T
}
type exprBinary[T typeScalar] struct {
	left     ExpressionSafe[T]
	operator binaryOperator
	right    ExpressionSafe[T]
}
type exprColumn[T typeScalar] struct {
	columnName string
	tableAlias string
}
type exprComparison[T typeScalar] struct {
	left       ExpressionSafe[T]
	operator   comparisonOperator
	right      ExpressionSafe[T]
	valueEnd   ExpressionSafe[T]
	valueGap   logicalOperator
	valueStart ExpressionSafe[T]
}
type exprComposite[T typeScalar] struct {
	expressions []ExpressionSafe[T]
	operator    compositeOperator
}
type exprConstant[T typeScalar] struct {
	value T
}
type exprFunction[InLT, InRT, T typeScalar] struct {
	distinct   bool
	json       []*exprJson
	left       ExpressionSafe[InLT]
	operator   compositeOperator
	process    processingStage
	right      ExpressionSafe[InRT]
	service    functionService
	valueArray []ExpressionBase
	valueType  ValueType
	window     *WindowSpec
}
type exprGroupBy struct {
	expression ExpressionBase
}
type exprJson struct {
	expressions []ExpressionBase
	operator    compositeOperator
	values      []ExpressionBase
}
type exprLiteral[T typeScalar] struct {
	value T
}
type exprLogical struct {
	expressions []markPredicable
	operator    logicalOperator
}
type exprOperator[T typeString] struct {
	value T
}
type exprOrderBy struct {
	direction  bool
	expression ExpressionBase
}
type exprService[T typeString] struct {
	value T
}
type exprSubquery[T typeScalar] struct {
	statement statement
}
type exprValue[T typeScalar] struct {
	value T
}

// Приватные методы
func (exprAlias *exprAlias[T]) isExpressionBase()  {}
func (exprAlias *exprAlias[T]) isExpressionSafe(T) {}
func (exprAlias *exprAlias[T]) isFieldable()       {}
func (exprAlias *exprAlias[T]) isPredicable()      {}
func (exprAlias *exprAlias[T]) isReturnable()      {}
func (exprAlias *exprAlias[T]) render(baseRenderer *baseRenderer) error {
	if err := exprAlias.expression.render(baseRenderer); err != nil {
		return err
	}
	if exprAlias.aliasName != "" {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierAs)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderAlias(exprAlias.aliasName)
	}
	return nil
}
func (exprAlias *exprAlias[T]) validate(baseValidator *baseValidator) error {
	if exprAlias.expression != nil {
		return exprAlias.expression.validate(baseValidator)
	}
	if err := baseValidator.validateAlias(exprAlias.aliasName); err != nil {
		return err
	}
	return nil
}
func (exprArray *exprArray[T]) isExpressionBase()  {}
func (exprArray *exprArray[T]) isExpressionSafe(T) {}
func (exprArray *exprArray[T]) isColumnable()      {}
func (exprArray *exprArray[T]) isFieldable()       {}
func (exprArray *exprArray[T]) isGroupable()       {}
func (exprArray *exprArray[T]) isOrderable()       {}
func (exprArray *exprArray[T]) isPredicable()      {}
func (exprArray *exprArray[T]) isReturnable()      {}
func (exprArray *exprArray[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeParenLeft)
	valueCount := len(exprArray.array) - 1
	for i, value := range exprArray.array {
		exprValue := &exprValue[T]{
			value: value,
		}
		if err := exprValue.render(baseRenderer); err != nil {
			return err
		}
		if i < valueCount {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	baseRenderer.renderOperator(uastCompositeParenRight)
	return nil
}
func (exprArray *exprArray[T]) validate(baseValidator *baseValidator) error {
	if exprArray.array == nil {
		return ErrInvalidArray
	}
	if err := baseValidator.validateArray(len(exprArray.array)); err != nil {
		return err
	}
	for _, value := range exprArray.array {
		if err := baseValidator.validateValue(any(value)); err != nil {
			return err
		}
	}
	return nil
}
func (exprBinary *exprBinary[T]) isExpressionBase()  {}
func (exprBinary *exprBinary[T]) isExpressionSafe(T) {}
func (exprBinary *exprBinary[T]) isFieldable()       {}
func (exprBinary *exprBinary[T]) isPredicable()      {}
func (exprBinary *exprBinary[T]) isReturnable()      {}
func (exprBinary *exprBinary[T]) render(baseRenderer *baseRenderer) error {
	if err := exprBinary.left.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderOperator(exprBinary.operator)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := exprBinary.right.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (exprBinary *exprBinary[T]) validate(baseValidator *baseValidator) error {
	if exprBinary.left == nil || exprBinary.right == nil {
		return ErrInvalidBinary
	}
	return nil
}
func (exprColumn *exprColumn[T]) isExpressionBase()  {}
func (exprColumn *exprColumn[T]) isExpressionSafe(T) {}
func (exprColumn *exprColumn[T]) isColumnable()      {}
func (exprColumn *exprColumn[T]) isFieldable()       {}
func (exprColumn *exprColumn[T]) isGroupable()       {}
func (exprColumn *exprColumn[T]) isOrderable()       {}
func (exprColumn *exprColumn[T]) isPredicable()      {}
func (exprColumn *exprColumn[T]) isReturnable()      {}
func (exprColumn *exprColumn[T]) render(baseRenderer *baseRenderer) error {
	if exprColumn.tableAlias != "" {
		baseRenderer.renderAlias(exprColumn.tableAlias)
		baseRenderer.renderOperator(uastCompositeSinglePoint)
	}
	baseRenderer.renderName(exprColumn.columnName)
	return nil
}
func (exprColumn *exprColumn[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(exprColumn.columnName); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(exprColumn.tableAlias); err != nil {
		return err
	}
	return nil
}
func (exprComparison *exprComparison[T]) isExpressionBase()  {}
func (exprComparison *exprComparison[T]) isExpressionSafe(T) {}
func (exprComparison *exprComparison[T]) isPredicable()      {}
func (exprComparison *exprComparison[T]) render(baseRenderer *baseRenderer) error {
	switch exprComparison.operator {
	case uastComparisonExists, uastComparisonNotExists:
		baseRenderer.renderOperator(exprComparison.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := exprComparison.left.render(baseRenderer); err != nil {
			return err
		}
	case uastComparisonIn, uastComparisonNotIn:
		if err := exprComparison.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(exprComparison.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := exprComparison.right.render(baseRenderer); err != nil {
			return err
		}
	case uastComparisonIsNull, uastComparisonIsNotNull:
		if err := exprComparison.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(exprComparison.operator)
	case uastComparisonLike, uastComparisonNotLike:
		if err := exprComparison.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(exprComparison.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := exprComparison.right.render(baseRenderer); err != nil {
			return err
		}
	case uastComparisonILike, uastComparisonNotILike:
		if err := exprComparison.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(exprComparison.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := exprComparison.right.render(baseRenderer); err != nil {
			return err
		}
	case uastComparisonBetween, uastComparisonNotBetween:
		if err := exprComparison.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(exprComparison.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := exprComparison.valueStart.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(exprComparison.valueGap)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := exprComparison.valueEnd.render(baseRenderer); err != nil {
			return err
		}
	default:
		if err := exprComparison.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(exprComparison.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := exprComparison.right.render(baseRenderer); err != nil {
			return err
		}
	}
	return nil
}
func (exprComparison *exprComparison[T]) transformGetLeft() ExpressionBase {
	return exprComparison.left
}
func (exprComparison *exprComparison[T]) transformGetOperator() comparisonOperator {
	return exprComparison.operator
}
func (exprComparison *exprComparison[T]) transformGetRight() ExpressionBase {
	return exprComparison.right
}
func (exprComparison *exprComparison[T]) transformGetValueEnd() ExpressionBase {
	return exprComparison.valueEnd
}
func (exprComparison *exprComparison[T]) transformGetValueStart() ExpressionBase {
	return exprComparison.valueStart
}
func (exprComparison *exprComparison[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateComparison(exprComparison); err != nil {
		return err
	}
	return nil
}
func (exprComposite *exprComposite[T]) isExpressionBase()  {}
func (exprComposite *exprComposite[T]) isExpressionSafe(T) {}
func (exprComposite *exprComposite[T]) isPredicable()      {}
func (exprComposite *exprComposite[T]) render(baseRenderer *baseRenderer) error {
	last := len(exprComposite.expressions) - 1
	for i, predicate := range exprComposite.expressions {
		if err := predicate.render(baseRenderer); err != nil {
			return err
		}
		if i < last {
			baseRenderer.renderOperator(exprComposite.operator)
		}
	}
	return nil
}
func (exprComposite *exprComposite[T]) validate(baseValidator *baseValidator) error {
	if exprComposite.expressions == nil {
		return ErrInvalidComposite
	}
	return nil
}
func (exprConstant *exprConstant[T]) isExpressionBase()  {}
func (exprConstant *exprConstant[T]) isExpressionSafe(T) {}
func (exprConstant *exprConstant[T]) isColumnable()      {}
func (exprConstant *exprConstant[T]) isFieldable()       {}
func (exprConstant *exprConstant[T]) isGroupable()       {}
func (exprConstant *exprConstant[T]) isOrderable()       {}
func (exprConstant *exprConstant[T]) isPredicable()      {}
func (exprConstant *exprConstant[T]) isReturnable()      {}
func (exprConstant *exprConstant[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderConstant(exprConstant.value)
	return nil
}
func (exprConstant *exprConstant[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateConstant(exprConstant.value); err != nil {
		return err
	}
	return nil
}
func (exprFunction *exprFunction[InLT, InRT, T]) isExpressionBase()  {}
func (exprFunction *exprFunction[InLT, InRT, T]) isExpressionSafe(T) {}
func (exprFunction *exprFunction[InLT, InRT, T]) isColumnable()      {}
func (exprFunction *exprFunction[InLT, InRT, T]) isFieldable()       {}
func (exprFunction *exprFunction[InLT, InRT, T]) isGroupable()       {}
func (exprFunction *exprFunction[InLT, InRT, T]) isOrderable()       {}
func (exprFunction *exprFunction[InLT, InRT, T]) isPredicable()      {}
func (exprFunction *exprFunction[InLT, InRT, T]) isReturnable()      {}
func (exprFunction *exprFunction[InLT, InRT, T]) render(baseRenderer *baseRenderer) error {
	switch exprFunction.process {
	case uastProcessСross:
		baseRenderer.renderService(exprFunction.service)
		length := len(exprFunction.valueArray)
		hasElse := length%2 == 1
		endIdx := length
		if hasElse {
			endIdx = length - 1
		}
		for i := 0; i < endIdx; i++ {
			if i%2 == 0 {
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderService(uastModifierWhen)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
			} else {
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderService(uastModifierThen)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
			}
			if err := exprFunction.valueArray[i].render(baseRenderer); err != nil {
				return err
			}
		}
		if hasElse {
			baseRenderer.renderOperator(uastCompositeSingleSpace)
			baseRenderer.renderService(uastModifierElse)
			baseRenderer.renderOperator(uastCompositeSingleSpace)
			if err := exprFunction.valueArray[length-1].render(baseRenderer); err != nil {
				return err
			}
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierEnd)
		return nil
	case uastProcessDirect:
		baseRenderer.renderService(exprFunction.service)
		if exprFunction.distinct {
			baseRenderer.renderService(uastModifierDistinct)
			baseRenderer.renderOperator(uastCompositeSingleSpace)
		}
		baseRenderer.renderOperator(uastCompositeParenLeft)
		if exprFunction.left != nil {
			if err := exprFunction.left.render(baseRenderer); err != nil {
				return err
			}
		}
		if exprFunction.operator != "" {
			baseRenderer.renderOperator(exprFunction.operator)
		}
		if exprFunction.right != nil {
			if err := exprFunction.right.render(baseRenderer); err != nil {
				return err
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
	case uastProcessEmpty:
		baseRenderer.renderService(exprFunction.service)
		if baseRenderer.config.parensFunction {
			baseRenderer.renderOperator(uastCompositeParenLeft)
			baseRenderer.renderOperator(uastCompositeParenRight)
		}
		return nil
	case uastProcessInvert:
		baseRenderer.renderService(exprFunction.service)
		if exprFunction.distinct {
			baseRenderer.renderService(uastModifierDistinct)
			baseRenderer.renderOperator(uastCompositeSingleSpace)
		}
		baseRenderer.renderOperator(uastCompositeParenLeft)
		if exprFunction.right != nil {
			if err := exprFunction.right.render(baseRenderer); err != nil {
				return err
			}
		}
		if exprFunction.operator != "" {
			baseRenderer.renderOperator(exprFunction.operator)
		}
		if exprFunction.left != nil {
			if err := exprFunction.left.render(baseRenderer); err != nil {
				return err
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
	case uastProcessJson:
		baseRenderer.renderService(exprFunction.service)
		baseRenderer.renderOperator(uastCompositeParenLeft)
		if exprFunction.left != nil {
			if err := exprFunction.left.render(baseRenderer); err != nil {
				return err
			}
		}
		if exprFunction.operator != "" {
			baseRenderer.renderOperator(exprFunction.operator)
		}
		last := len(exprFunction.json) - 1
		for i, group := range exprFunction.json {
			if err := group.render(baseRenderer); err != nil {
				return err
			}
			if i < last {
				baseRenderer.renderOperator(exprFunction.operator)
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
	case uastProcessWindow:
		baseRenderer.renderService(exprFunction.service)
		baseRenderer.renderOperator(uastCompositeParenLeft)
		if exprFunction.left != nil {
			exprFunction.left.render(baseRenderer)
			if exprFunction.right != nil {
				baseRenderer.renderOperator(uastCompositeCommaSpace)
				exprFunction.right.render(baseRenderer)
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierOver)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(uastCompositeParenLeft)
		if exprFunction.window != nil {
			if len(exprFunction.window.partition) > 0 {
				baseRenderer.renderService(uastManagementPartitionBy)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				for i, expr := range exprFunction.window.partition {
					if i > 0 {
						baseRenderer.renderOperator(uastCompositeCommaSpace)
					}
					expr.render(baseRenderer)
				}
			}
			if len(exprFunction.window.order) > 0 {
				if len(exprFunction.window.partition) > 0 {
					baseRenderer.renderOperator(uastCompositeSingleSpace)
				}
				baseRenderer.renderService(uastManagementOrderBy)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				for i, order := range exprFunction.window.order {
					if i > 0 {
						baseRenderer.renderOperator(uastCompositeCommaSpace)
					}
					order.render(baseRenderer)
				}
			}
			if exprFunction.window.frame != nil {
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderOperator(exprFunction.window.frame.Type)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderService(uastModifierBetween)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderOperator(exprFunction.window.frame.Start)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderService(uastModifierAnd)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderOperator(exprFunction.window.frame.End)
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
	}
	return nil
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformGetDistinct() bool {
	return exprFunction.distinct
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformGetJson() []*exprJson {
	return exprFunction.json
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformGetLeft() ExpressionBase {
	return exprFunction.left
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformGetParamCount() int {
	var countJson int
	var countLeft int
	var countRight int
	var countValueArray int
	var countValueType int
	countJson = len(exprFunction.json)
	switch left := exprFunction.left.(type) {
	case nil:
		countLeft = 0
	case *exprComposite[InLT]:
		countLeft = len(left.expressions)
	default:
		countLeft = 1
	}
	switch right := exprFunction.right.(type) {
	case nil:
		countRight = 0
	case *exprComposite[InRT]:
		countRight = len(right.expressions)
	default:
		countRight = 1
	}
	countValueArray = len(exprFunction.valueArray)
	if exprFunction.valueType != "" {
		countValueType = 1
	}
	return countJson + countLeft + countRight + countValueArray + countValueType
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformGetProcess() processingStage {
	return exprFunction.process
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformGetRight() ExpressionBase {
	return exprFunction.right
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformGetService() functionService {
	return exprFunction.service
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformGetValueArray() []ExpressionBase {
	return exprFunction.valueArray
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformGetValueType() ValueType {
	return exprFunction.valueType
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformSetJson(json []*exprJson) {
	exprFunction.json = json
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformSetLeft(left ExpressionBase) {
	if expression, ok := left.(ExpressionSafe[InLT]); ok {
		exprFunction.left = expression
	}
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformSetOperator(operator compositeOperator) {
	exprFunction.operator = operator
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformSetProcess(process processingStage) {
	exprFunction.process = process
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformSetRight(right ExpressionBase) {
	if expression, ok := right.(ExpressionSafe[InRT]); ok {
		exprFunction.right = expression
	}
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformSetService(service functionService) {
	exprFunction.service = service
}
func (exprFunction *exprFunction[InLT, InRT, T]) transformSetValueArray(valueArray []ExpressionBase) {
	exprFunction.valueArray = valueArray
}
func (exprFunction *exprFunction[InLT, InRT, T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateFunction(exprFunction); err != nil {
		return err
	}
	return nil
}
func (exprGroupBy *exprGroupBy) isExpressionBase() {}
func (exprGroupBy *exprGroupBy) isGroupable()      {}
func (exprGroupBy *exprGroupBy) render(baseRenderer *baseRenderer) error {
	if err := exprGroupBy.expression.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (exprGroupBy *exprGroupBy) validate(baseValidator *baseValidator) error {
	if exprGroupBy.expression == nil {
		return ErrInvalidStatementGroupBy
	}
	return nil
}
func (exprJson *exprJson) isExpressionBase() {}
func (exprJson *exprJson) render(baseRenderer *baseRenderer) error {
	lengthExpressions := len(exprJson.expressions) - 1
	for i, expression := range exprJson.expressions {
		if err := expression.render(baseRenderer); err != nil {
			return err
		}
		if i < lengthExpressions {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	if exprJson.expressions != nil && exprJson.values != nil {
		baseRenderer.renderOperator(exprJson.operator)
	}
	lengthValues := len(exprJson.values) - 1
	for i, value := range exprJson.values {
		if err := value.render(baseRenderer); err != nil {
			return err
		}
		if i < lengthValues {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (exprJson *exprJson) validate(baseValidator *baseValidator) error {
	if exprJson.expressions == nil {
		return ErrInvalidStatementJson
	}
	return nil
}
func (exprLiteral *exprLiteral[T]) isExpressionBase()  {}
func (exprLiteral *exprLiteral[T]) isExpressionSafe(T) {}
func (exprLiteral *exprLiteral[T]) isColumnable()      {}
func (exprLiteral *exprLiteral[T]) isFieldable()       {}
func (exprLiteral *exprLiteral[T]) isGroupable()       {}
func (exprLiteral *exprLiteral[T]) isOrderable()       {}
func (exprLiteral *exprLiteral[T]) isPredicable()      {}
func (exprLiteral *exprLiteral[T]) isReturnable()      {}
func (exprLiteral *exprLiteral[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderLiteral(exprLiteral.value)
	return nil
}
func (exprLiteral *exprLiteral[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateLiteral(exprLiteral.value); err != nil {
		return err
	}
	return nil
}
func (exprLogical *exprLogical) isExpressionBase()     {}
func (exprLiteral *exprLogical) isExpressionSafe(bool) {}
func (exprLogical *exprLogical) isPredicable()         {}
func (exprLogical *exprLogical) render(baseRenderer *baseRenderer) error {
	switch len(exprLogical.expressions) {
	case 1:
		return exprLogical.expressions[0].render(baseRenderer)
	default:
		// !!! Внимание - Доработать скобки на первом уровне
		baseRenderer.renderOperator(uastCompositeParenLeft)
		last := len(exprLogical.expressions) - 1
		for i, predicate := range exprLogical.expressions {
			if err := predicate.render(baseRenderer); err != nil {
				return err
			}
			if i < last {
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderOperator(exprLogical.operator)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
	}
	return nil
}
func (exprLogical *exprLogical) validate(baseValidator *baseValidator) error {
	if exprLogical.expressions == nil {
		return ErrInvalidLogical
	}
	for _, predicate := range exprLogical.expressions {
		if err := predicate.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
func (exprOperator *exprOperator[T]) isExpressionBase()  {}
func (exprOperator *exprOperator[T]) isExpressionSafe(T) {}
func (exprOperator *exprOperator[T]) isPredicable()      {}
func (exprOperator *exprOperator[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(exprOperator.value)
	return nil
}
func (exprOperator *exprOperator[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateOperator(exprOperator.value); err != nil {
		return err
	}
	return nil
}
func (exprOrderBy *exprOrderBy) isExpressionBase() {}
func (exprOrderBy *exprOrderBy) isOrderable()      {}
func (exprOrderBy *exprOrderBy) render(baseRenderer *baseRenderer) error {
	if err := exprOrderBy.expression.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if exprOrderBy.direction {
		baseRenderer.renderOperator(uastOrderDesc)
	} else {
		baseRenderer.renderOperator(uastOrderAsc)
	}
	return nil
}
func (exprOrderBy *exprOrderBy) validate(baseValidator *baseValidator) error {
	if exprOrderBy == nil {
		return ErrInvalidStatementOrderBy
	}
	return nil
}
func (exprService *exprService[T]) isExpressionBase()  {}
func (exprService *exprService[T]) isExpressionSafe(T) {}
func (exprService *exprService[T]) isPredicable()      {}
func (exprService *exprService[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderService(exprService.value)
	return nil
}
func (exprService *exprService[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateService(exprService.value); err != nil {
		return err
	}
	return nil
}
func (exprSubquery *exprSubquery[T]) isExpressionBase()  {}
func (exprSubquery *exprSubquery[T]) isExpressionSafe(T) {}
func (exprSubquery *exprSubquery[T]) isColumnable()      {}
func (exprSubquery *exprSubquery[T]) isFieldable()       {}
func (exprSubquery *exprSubquery[T]) isGroupable()       {}
func (exprSubquery *exprSubquery[T]) isOrderable()       {}
func (exprSubquery *exprSubquery[T]) isPredicable()      {}
func (exprSubquery *exprSubquery[T]) isReturnable()      {}
func (exprSubquery *exprSubquery[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeParenLeft)
	if err := exprSubquery.statement.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeParenRight)
	return nil
}
func (exprSubquery *exprSubquery[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateSubquery(); err != nil {
		return err
	}
	if err := exprSubquery.statement.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
func (exprValue *exprValue[T]) isExpressionBase()  {}
func (exprValue *exprValue[T]) isExpressionSafe(T) {}
func (exprValue *exprValue[T]) isColumnable()      {}
func (exprValue *exprValue[T]) isFieldable()       {}
func (exprValue *exprValue[T]) isGroupable()       {}
func (exprValue *exprValue[T]) isOrderable()       {}
func (exprValue *exprValue[T]) isPredicable()      {}
func (exprValue *exprValue[T]) isReturnable()      {}
func (exprValue *exprValue[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderValue(exprValue.value)
	return nil
}
func (exprValue *exprValue[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateValue(exprValue.value); err != nil {
		return err
	}
	return nil
}
