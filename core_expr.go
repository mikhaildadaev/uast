package uast

// Публичные интерфейсы
type ExpressionBase interface {
	isExpressionBase()
	render(baseRenderer *baseRenderer) error
	validate(baseValidator *baseValidator) error
}
type ExpressionSafe[T typeScalar] interface {
	ExpressionBase
	markPredicable
	isExpressionSafe(T)
}

// Публичные структуры
type AliasExpr[T typeScalar] = exprAlias[T]
type ColumnExpr[T typeScalar] = exprColumn[T]

// Приватные интерфейсы
type markExpressable interface {
	ExpressionBase
	isColumnable()
	isFieldable()
}
type markPredicable interface {
	ExpressionBase
	isPredicable()
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
	name       string
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
type exprJson struct {
	expressions []ExpressionBase
	operator    compositeOperator
	values      []ExpressionBase
}
type exprLiteral[T typeScalar] struct {
	value T
}
type exprLogical struct {
	expressions     []markPredicable
	operator        logicalOperator
	parentIsLogical bool
}
type exprOperator[T typeString] struct {
	value T
}
type exprPair[T typeScalar] struct {
	name string
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
func (expr *exprAlias[T]) isExpressionBase()  {}
func (expr *exprAlias[T]) isExpressionSafe(T) {}
func (expr *exprAlias[T]) isColumnable()      {}
func (expr *exprAlias[T]) isFieldable()       {}
func (expr *exprAlias[T]) isPredicable()      {}
func (expr *exprAlias[T]) isReturnable()      {}
func (expr *exprAlias[T]) render(baseRenderer *baseRenderer) error {
	if err := expr.expression.render(baseRenderer); err != nil {
		return err
	}
	if expr.aliasName != "" {
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierAs)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderAlias(expr.aliasName)
	}
	return nil
}
func (expr *exprAlias[T]) validate(baseValidator *baseValidator) error {
	if expr.expression != nil {
		return expr.expression.validate(baseValidator)
	}
	if err := baseValidator.validateAlias(expr.aliasName); err != nil {
		return err
	}
	return nil
}
func (expr *exprArray[T]) isExpressionBase()  {}
func (expr *exprArray[T]) isExpressionSafe(T) {}
func (expr *exprArray[T]) isColumnable()      {}
func (expr *exprArray[T]) isFieldable()       {}
func (expr *exprArray[T]) isGroupable()       {}
func (expr *exprArray[T]) isOrderable()       {}
func (expr *exprArray[T]) isPredicable()      {}
func (expr *exprArray[T]) isReturnable()      {}
func (expr *exprArray[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeParenLeft)
	valueCount := len(expr.array) - 1
	for i, value := range expr.array {
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
func (expr *exprArray[T]) validate(baseValidator *baseValidator) error {
	if expr.array == nil {
		return ErrInvalidArray
	}
	if err := baseValidator.validateArray(len(expr.array)); err != nil {
		return err
	}
	for _, value := range expr.array {
		if err := baseValidator.validateValue(any(value)); err != nil {
			return err
		}
	}
	return nil
}
func (expr *exprBinary[T]) isExpressionBase()  {}
func (expr *exprBinary[T]) isExpressionSafe(T) {}
func (expr *exprBinary[T]) isColumnable()      {}
func (expr *exprBinary[T]) isFieldable()       {}
func (expr *exprBinary[T]) isPredicable()      {}
func (expr *exprBinary[T]) isReturnable()      {}
func (expr *exprBinary[T]) render(baseRenderer *baseRenderer) error {
	if err := expr.left.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderOperator(expr.operator)
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	if err := expr.right.render(baseRenderer); err != nil {
		return err
	}
	return nil
}
func (expr *exprBinary[T]) validate(baseValidator *baseValidator) error {
	if expr.left == nil || expr.right == nil {
		return ErrInvalidBinary
	}
	return nil
}
func (expr *exprColumn[T]) isExpressionBase()  {}
func (expr *exprColumn[T]) isExpressionSafe(T) {}
func (expr *exprColumn[T]) isColumnable()      {}
func (expr *exprColumn[T]) isFieldable()       {}
func (expr *exprColumn[T]) isGroupable()       {}
func (expr *exprColumn[T]) isOrderable()       {}
func (expr *exprColumn[T]) isPredicable()      {}
func (expr *exprColumn[T]) isReturnable()      {}
func (expr *exprColumn[T]) render(baseRenderer *baseRenderer) error {
	if expr.tableAlias != "" {
		baseRenderer.renderAlias(expr.tableAlias)
		baseRenderer.renderOperator(uastCompositeSinglePoint)
	}
	baseRenderer.renderName(expr.name)
	return nil
}
func (expr *exprColumn[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(expr.name); err != nil {
		return err
	}
	if err := baseValidator.validateAlias(expr.tableAlias); err != nil {
		return err
	}
	return nil
}
func (expr *exprComparison[T]) isExpressionBase()  {}
func (expr *exprComparison[T]) isExpressionSafe(T) {}
func (expr *exprComparison[T]) isPredicable()      {}
func (expr *exprComparison[T]) render(baseRenderer *baseRenderer) error {
	switch expr.operator {
	case uastComparisonExists, uastComparisonNotExists:
		baseRenderer.renderOperator(expr.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := expr.left.render(baseRenderer); err != nil {
			return err
		}
	case uastComparisonIn, uastComparisonNotIn:
		if err := expr.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(expr.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := expr.right.render(baseRenderer); err != nil {
			return err
		}
	case uastComparisonIsNull, uastComparisonIsNotNull:
		if err := expr.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(expr.operator)
	case uastComparisonLike, uastComparisonNotLike:
		if err := expr.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(expr.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := expr.right.render(baseRenderer); err != nil {
			return err
		}
	case uastComparisonILike, uastComparisonNotILike:
		if err := expr.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(expr.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := expr.right.render(baseRenderer); err != nil {
			return err
		}
	case uastComparisonBetween, uastComparisonNotBetween:
		if err := expr.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(expr.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := expr.valueStart.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(expr.valueGap)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := expr.valueEnd.render(baseRenderer); err != nil {
			return err
		}
	default:
		if err := expr.left.render(baseRenderer); err != nil {
			return err
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(expr.operator)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		if err := expr.right.render(baseRenderer); err != nil {
			return err
		}
	}
	return nil
}
func (expr *exprComparison[T]) transformGetLeft() ExpressionBase {
	return expr.left
}
func (expr *exprComparison[T]) transformGetOperator() comparisonOperator {
	return expr.operator
}
func (expr *exprComparison[T]) transformGetRight() ExpressionBase {
	return expr.right
}
func (expr *exprComparison[T]) transformGetValueEnd() ExpressionBase {
	return expr.valueEnd
}
func (expr *exprComparison[T]) transformGetValueStart() ExpressionBase {
	return expr.valueStart
}
func (expr *exprComparison[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateComparison(expr); err != nil {
		return err
	}
	return nil
}
func (expr *exprComposite[T]) isExpressionBase()  {}
func (expr *exprComposite[T]) isExpressionSafe(T) {}
func (expr *exprComposite[T]) isPredicable()      {}
func (expr *exprComposite[T]) render(baseRenderer *baseRenderer) error {
	last := len(expr.expressions) - 1
	for i, predicate := range expr.expressions {
		if err := predicate.render(baseRenderer); err != nil {
			return err
		}
		if i < last {
			baseRenderer.renderOperator(expr.operator)
		}
	}
	return nil
}
func (expr *exprComposite[T]) validate(baseValidator *baseValidator) error {
	if expr.expressions == nil {
		return ErrInvalidComposite
	}
	return nil
}
func (expr *exprConstant[T]) isExpressionBase()  {}
func (expr *exprConstant[T]) isExpressionSafe(T) {}
func (expr *exprConstant[T]) isColumnable()      {}
func (expr *exprConstant[T]) isFieldable()       {}
func (expr *exprConstant[T]) isGroupable()       {}
func (expr *exprConstant[T]) isOrderable()       {}
func (expr *exprConstant[T]) isPredicable()      {}
func (expr *exprConstant[T]) isReturnable()      {}
func (expr *exprConstant[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderConstant(expr.value)
	return nil
}
func (expr *exprConstant[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateConstant(expr.value); err != nil {
		return err
	}
	return nil
}
func (expr *exprFunction[InLT, InRT, T]) isExpressionBase()  {}
func (expr *exprFunction[InLT, InRT, T]) isExpressionSafe(T) {}
func (expr *exprFunction[InLT, InRT, T]) isColumnable()      {}
func (expr *exprFunction[InLT, InRT, T]) isFieldable()       {}
func (expr *exprFunction[InLT, InRT, T]) isGroupable()       {}
func (expr *exprFunction[InLT, InRT, T]) isOrderable()       {}
func (expr *exprFunction[InLT, InRT, T]) isPredicable()      {}
func (expr *exprFunction[InLT, InRT, T]) isReturnable()      {}
func (expr *exprFunction[InLT, InRT, T]) render(baseRenderer *baseRenderer) error {
	switch expr.process {
	case uastProcessСross:
		baseRenderer.renderService(expr.service)
		length := len(expr.valueArray)
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
			if err := expr.valueArray[i].render(baseRenderer); err != nil {
				return err
			}
		}
		if hasElse {
			baseRenderer.renderOperator(uastCompositeSingleSpace)
			baseRenderer.renderService(uastModifierElse)
			baseRenderer.renderOperator(uastCompositeSingleSpace)
			if err := expr.valueArray[length-1].render(baseRenderer); err != nil {
				return err
			}
		}
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierEnd)
		return nil
	case uastProcessDirect:
		baseRenderer.renderService(expr.service)
		if expr.distinct {
			baseRenderer.renderService(uastModifierDistinct)
			baseRenderer.renderOperator(uastCompositeSingleSpace)
		}
		baseRenderer.renderOperator(uastCompositeParenLeft)
		if expr.left != nil {
			if err := expr.left.render(baseRenderer); err != nil {
				return err
			}
		}
		if expr.operator != "" {
			baseRenderer.renderOperator(expr.operator)
		}
		if expr.right != nil {
			if err := expr.right.render(baseRenderer); err != nil {
				return err
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
	case uastProcessEmpty:
		baseRenderer.renderService(expr.service)
		if baseRenderer.config.parensFunction {
			baseRenderer.renderOperator(uastCompositeParenLeft)
			baseRenderer.renderOperator(uastCompositeParenRight)
		}
		return nil
	case uastProcessInvert:
		baseRenderer.renderService(expr.service)
		if expr.distinct {
			baseRenderer.renderService(uastModifierDistinct)
			baseRenderer.renderOperator(uastCompositeSingleSpace)
		}
		baseRenderer.renderOperator(uastCompositeParenLeft)
		if expr.right != nil {
			if err := expr.right.render(baseRenderer); err != nil {
				return err
			}
		}
		if expr.operator != "" {
			baseRenderer.renderOperator(expr.operator)
		}
		if expr.left != nil {
			if err := expr.left.render(baseRenderer); err != nil {
				return err
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
	case uastProcessJson:
		baseRenderer.renderService(expr.service)
		baseRenderer.renderOperator(uastCompositeParenLeft)
		if expr.left != nil {
			if err := expr.left.render(baseRenderer); err != nil {
				return err
			}
		}
		if expr.operator != "" {
			baseRenderer.renderOperator(expr.operator)
		}
		last := len(expr.json) - 1
		for i, group := range expr.json {
			if err := group.render(baseRenderer); err != nil {
				return err
			}
			if i < last {
				baseRenderer.renderOperator(expr.operator)
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
	case uastProcessWindow:
		baseRenderer.renderService(expr.service)
		baseRenderer.renderOperator(uastCompositeParenLeft)
		if expr.left != nil {
			expr.left.render(baseRenderer)
			if expr.right != nil {
				baseRenderer.renderOperator(uastCompositeCommaSpace)
				expr.right.render(baseRenderer)
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderService(uastModifierOver)
		baseRenderer.renderOperator(uastCompositeSingleSpace)
		baseRenderer.renderOperator(uastCompositeParenLeft)
		if expr.window != nil {
			if len(expr.window.partition) > 0 {
				baseRenderer.renderService(uastManagementPartitionBy)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				for i, expression := range expr.window.partition {
					if i > 0 {
						baseRenderer.renderOperator(uastCompositeCommaSpace)
					}
					expression.render(baseRenderer)
				}
			}
			if len(expr.window.order) > 0 {
				if len(expr.window.partition) > 0 {
					baseRenderer.renderOperator(uastCompositeSingleSpace)
				}
				baseRenderer.renderService(uastManagementOrderBy)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				for i, order := range expr.window.order {
					if i > 0 {
						baseRenderer.renderOperator(uastCompositeCommaSpace)
					}
					order.render(baseRenderer)
				}
			}
			if expr.window.frame != nil {
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderOperator(expr.window.frame.Type)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderService(uastModifierBetween)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderOperator(expr.window.frame.Start)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderService(uastModifierAnd)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderOperator(expr.window.frame.End)
			}
		}
		baseRenderer.renderOperator(uastCompositeParenRight)
	}
	return nil
}
func (expr *exprFunction[InLT, InRT, T]) transformGetDistinct() bool {
	return expr.distinct
}
func (expr *exprFunction[InLT, InRT, T]) transformGetJson() []*exprJson {
	return expr.json
}
func (expr *exprFunction[InLT, InRT, T]) transformGetLeft() ExpressionBase {
	return expr.left
}
func (expr *exprFunction[InLT, InRT, T]) transformGetParamCount() int {
	var countJson int
	var countLeft int
	var countRight int
	var countValueArray int
	var countValueType int
	countJson = len(expr.json)
	switch left := expr.left.(type) {
	case nil:
		countLeft = 0
	case *exprComposite[InLT]:
		countLeft = len(left.expressions)
	default:
		countLeft = 1
	}
	switch right := expr.right.(type) {
	case nil:
		countRight = 0
	case *exprComposite[InRT]:
		countRight = len(right.expressions)
	default:
		countRight = 1
	}
	countValueArray = len(expr.valueArray)
	if expr.valueType != "" {
		countValueType = 1
	}
	return countJson + countLeft + countRight + countValueArray + countValueType
}
func (expr *exprFunction[InLT, InRT, T]) transformGetProcess() processingStage {
	return expr.process
}
func (expr *exprFunction[InLT, InRT, T]) transformGetRight() ExpressionBase {
	return expr.right
}
func (expr *exprFunction[InLT, InRT, T]) transformGetService() functionService {
	return expr.service
}
func (expr *exprFunction[InLT, InRT, T]) transformGetValueArray() []ExpressionBase {
	return expr.valueArray
}
func (expr *exprFunction[InLT, InRT, T]) transformGetValueType() ValueType {
	return expr.valueType
}
func (expr *exprFunction[InLT, InRT, T]) transformSetJson(json []*exprJson) {
	expr.json = json
}
func (expr *exprFunction[InLT, InRT, T]) transformSetLeft(left ExpressionBase) {
	if expression, ok := left.(ExpressionSafe[InLT]); ok {
		expr.left = expression
	}
}
func (expr *exprFunction[InLT, InRT, T]) transformSetOperator(operator compositeOperator) {
	expr.operator = operator
}
func (expr *exprFunction[InLT, InRT, T]) transformSetProcess(process processingStage) {
	expr.process = process
}
func (expr *exprFunction[InLT, InRT, T]) transformSetRight(right ExpressionBase) {
	if expression, ok := right.(ExpressionSafe[InRT]); ok {
		expr.right = expression
	}
}
func (expr *exprFunction[InLT, InRT, T]) transformSetService(service functionService) {
	expr.service = service
}
func (expr *exprFunction[InLT, InRT, T]) transformSetValueArray(valueArray []ExpressionBase) {
	expr.valueArray = valueArray
}
func (expr *exprFunction[InLT, InRT, T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateFunction(expr); err != nil {
		return err
	}
	return nil
}
func (expr *exprJson) isExpressionBase() {}
func (expr *exprJson) render(baseRenderer *baseRenderer) error {
	lengthExpressions := len(expr.expressions) - 1
	for i, expression := range expr.expressions {
		if err := expression.render(baseRenderer); err != nil {
			return err
		}
		if i < lengthExpressions {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	if expr.expressions != nil && expr.values != nil {
		baseRenderer.renderOperator(expr.operator)
	}
	lengthValues := len(expr.values) - 1
	for i, value := range expr.values {
		if err := value.render(baseRenderer); err != nil {
			return err
		}
		if i < lengthValues {
			baseRenderer.renderOperator(uastCompositeCommaSpace)
		}
	}
	return nil
}
func (expr *exprJson) validate(baseValidator *baseValidator) error {
	if expr.expressions == nil {
		return ErrInvalidStatementJson
	}
	return nil
}
func (expr *exprLiteral[T]) isExpressionBase()  {}
func (expr *exprLiteral[T]) isExpressionSafe(T) {}
func (expr *exprLiteral[T]) isColumnable()      {}
func (expr *exprLiteral[T]) isFieldable()       {}
func (expr *exprLiteral[T]) isGroupable()       {}
func (expr *exprLiteral[T]) isOrderable()       {}
func (expr *exprLiteral[T]) isPredicable()      {}
func (expr *exprLiteral[T]) isReturnable()      {}
func (expr *exprLiteral[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderLiteral(expr.value)
	return nil
}
func (expr *exprLiteral[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateLiteral(expr.value); err != nil {
		return err
	}
	return nil
}
func (expr *exprLogical) isExpressionBase()     {}
func (expr *exprLogical) isExpressionSafe(bool) {}
func (expr *exprLogical) isPredicable()         {}
func (expr *exprLogical) render(baseRenderer *baseRenderer) error {
	switch len(expr.expressions) {
	case 1:
		return expr.expressions[0].render(baseRenderer)
	default:
		if !expr.parentIsLogical {
			baseRenderer.renderOperator(uastCompositeParenLeft)
		}
		// !!! Внимание, необходимо доработать скобки на первом уровне
		last := len(expr.expressions) - 1
		for i, predicate := range expr.expressions {
			if logical, ok := predicate.(*exprLogical); ok {
				logical.parentIsLogical = true
			}
			if err := predicate.render(baseRenderer); err != nil {
				return err
			}
			if i < last {
				baseRenderer.renderOperator(uastCompositeSingleSpace)
				baseRenderer.renderOperator(expr.operator)
				baseRenderer.renderOperator(uastCompositeSingleSpace)
			}
		}
		if !expr.parentIsLogical {
			baseRenderer.renderOperator(uastCompositeParenLeft)
		}
	}
	return nil
}
func (expr *exprLogical) validate(baseValidator *baseValidator) error {
	if expr.expressions == nil {
		return ErrInvalidLogical
	}
	for _, predicate := range expr.expressions {
		if err := predicate.validate(baseValidator); err != nil {
			return err
		}
	}
	return nil
}
func (expr *exprOperator[T]) isExpressionBase()  {}
func (expr *exprOperator[T]) isExpressionSafe(T) {}
func (expr *exprOperator[T]) isPredicable()      {}
func (expr *exprOperator[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(expr.value)
	return nil
}
func (expr *exprOperator[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateOperator(expr.value); err != nil {
		return err
	}
	return nil
}
func (expr *exprPair[T]) isExpressionBase()  {}
func (expr *exprPair[T]) isExpressionSafe(T) {}
func (expr *exprPair[T]) isColumnable()      {}
func (expr *exprPair[T]) isFieldable()       {}
func (expr *exprPair[T]) isGroupable()       {}
func (expr *exprPair[T]) isOrderable()       {}
func (expr *exprPair[T]) isPredicable()      {}
func (expr *exprPair[T]) isReturnable()      {}
func (expr *exprPair[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderName(expr.name)
	return nil
}
func (expr *exprPair[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateName(expr.name); err != nil {
		return err
	}
	return nil
}
func (expr *exprService[T]) isExpressionBase()  {}
func (expr *exprService[T]) isExpressionSafe(T) {}
func (expr *exprService[T]) isPredicable()      {}
func (expr *exprService[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderService(expr.value)
	return nil
}
func (expr *exprService[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateService(expr.value); err != nil {
		return err
	}
	return nil
}
func (expr *exprSubquery[T]) isExpressionBase()  {}
func (expr *exprSubquery[T]) isExpressionSafe(T) {}
func (expr *exprSubquery[T]) isColumnable()      {}
func (expr *exprSubquery[T]) isFieldable()       {}
func (expr *exprSubquery[T]) isGroupable()       {}
func (expr *exprSubquery[T]) isOrderable()       {}
func (expr *exprSubquery[T]) isPredicable()      {}
func (expr *exprSubquery[T]) isReturnable()      {}
func (expr *exprSubquery[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderOperator(uastCompositeParenLeft)
	if err := expr.statement.render(baseRenderer); err != nil {
		return err
	}
	baseRenderer.renderOperator(uastCompositeParenRight)
	return nil
}
func (expr *exprSubquery[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateSubquery(); err != nil {
		return err
	}
	if err := expr.statement.validate(baseValidator); err != nil {
		return err
	}
	return nil
}
func (expr *exprValue[T]) isExpressionBase()  {}
func (expr *exprValue[T]) isExpressionSafe(T) {}
func (expr *exprValue[T]) isColumnable()      {}
func (expr *exprValue[T]) isFieldable()       {}
func (expr *exprValue[T]) isGroupable()       {}
func (expr *exprValue[T]) isOrderable()       {}
func (expr *exprValue[T]) isPredicable()      {}
func (expr *exprValue[T]) isReturnable()      {}
func (expr *exprValue[T]) render(baseRenderer *baseRenderer) error {
	baseRenderer.renderValue(expr.value)
	return nil
}
func (expr *exprValue[T]) validate(baseValidator *baseValidator) error {
	if err := baseValidator.validateValue(expr.value); err != nil {
		return err
	}
	return nil
}
