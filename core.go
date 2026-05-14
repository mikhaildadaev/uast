package uast

import (
	"errors"
	"time"
)

// Публичные типы
type (
	ValueType string
)

// Публичные константы
const (
	Author  = "Mikhail Dadaev"
	Version = "1.26.11"
)
const (
	// Типы бинарные
	TypeBinary    ValueType = "BINARY"
	TypeVarBinary ValueType = "VARBINARY"
	// Типы даты и времени
	TypeDate      ValueType = "DATE"
	TypeDateTime  ValueType = "DATETIME"
	TypeTime      ValueType = "TIME"
	TypeTimestamp ValueType = "TIMESTAMP"
	// Типы числовые
	TypeBigInt   ValueType = "BIGINT"
	TypeDecimal  ValueType = "DECIMAL"
	TypeDouble   ValueType = "DOUBLE"
	TypeFloat    ValueType = "FLOAT"
	TypeInt      ValueType = "INT"
	TypeSmallInt ValueType = "SMALLINT"
	// Типы строковые
	TypeChar    ValueType = "CHAR"
	TypeString  ValueType = "STRING"
	TypeText    ValueType = "TEXT"
	TypeVarChar ValueType = "VARCHAR"
	// Типы специальные
	TypeArray   ValueType = "ARRAY"
	TypeBoolean ValueType = "BOOLEAN"
	TypeJSON    ValueType = "JSON"
	TypeUUID    ValueType = "UUID"
	TypeXML     ValueType = "XML"
)

// Публичные переменные
var (
	ErrDuplicateCTE                       = errors.New("duplicate cte")
	ErrExcessMaxDepth                     = errors.New("excess max depth")
	ErrExcessMaxComparison                = errors.New("excess max comparison")
	ErrExcessMaxFunction                  = errors.New("excess max function")
	ErrExcessMaxLimit                     = errors.New("excess max limit")
	ErrExcessMaxSubquery                  = errors.New("excess max subquery")
	ErrExcessMaxUnions                    = errors.New("excess max unions")
	ErrExcessMaxWith                      = errors.New("excess max with")
	ErrInjectionConstant                  = errors.New("injection constant")
	ErrInjectionFunction                  = errors.New("injection function")
	ErrInjectionIdentAlias                = errors.New("injection identifier alias")
	ErrInjectionIdentName                 = errors.New("injection identifier name")
	ErrInjectionLiteral                   = errors.New("injection literal")
	ErrInvalidDialect                     = errors.New("invalid dialect")
	ErrInvalidStatement                   = errors.New("invalid statement")
	ErrInvalidStatementField              = errors.New("invalid statement field")
	ErrInvalidStatementFrom               = errors.New("invalid statement from")
	ErrInvalidStatementInto               = errors.New("invalid statement into")
	ErrInvalidStatementJoin               = errors.New("invalid statement join")
	ErrInvalidStatementJoinCross          = errors.New("invalid statement join cross")
	ErrInvalidStatementJson               = errors.New("invalid statement json")
	ErrInvalidStatementGroupBy            = errors.New("invalid statement groupby")
	ErrInvalidStatementLimit              = errors.New("invalid statement limit")
	ErrInvalidStatementOffset             = errors.New("invalid statement offset")
	ErrInvalidStatementOnto               = errors.New("invalid statement onto")
	ErrInvalidStatementOrderBy            = errors.New("invalid statement orderby")
	ErrInvalidStatementReturning          = errors.New("invalid statement returning")
	ErrInvalidStatementSet                = errors.New("invalid statement set")
	ErrInvalidStatementTarget             = errors.New("invalid statement target")
	ErrInvalidStatementUnions             = errors.New("invalid statement unions")
	ErrInvalidStatementWith               = errors.New("invalid statement with")
	ErrInvalidAlias                       = errors.New("invalid alias")
	ErrInvalidArray                       = errors.New("invalid array")
	ErrInvalidBinary                      = errors.New("invalid binary")
	ErrInvalidComparison                  = errors.New("invalid comparison")
	ErrInvalidComparisonBetweenNotBetween = errors.New("invalid comparison between/not between")
	ErrInvalidComparisonExistsNotExists   = errors.New("invalid comparison exists/not exists")
	ErrInvalidComparisonInNotIn           = errors.New("invalid comparison in/not in")
	ErrInvalidComparisonIsNullIsNotNull   = errors.New("invalid comparison is null/is not null")
	ErrInvalidComparisonLikeNotLike       = errors.New("invalid comparison like/not like")
	ErrInvalidComparisonILikeNotILike     = errors.New("invalid comparison ilike/not ilike")
	ErrInvalidComposite                   = errors.New("invalid composite")
	ErrInvalidConstant                    = errors.New("invalid constant")
	ErrInvalidFunction                    = errors.New("invalid function")
	ErrInvalidIdentAlias                  = errors.New("invalid identifier alias")
	ErrInvalidIdentName                   = errors.New("invalid identifier name")
	ErrInvalidJson                        = errors.New("invalid json")
	ErrInvalidLiteral                     = errors.New("invalid literal")
	ErrInvalidLogical                     = errors.New("invalid logical")
	ErrInvalidOperator                    = errors.New("invalid operator")
	ErrInvalidOperatorBinary              = errors.New("invalid operator binary")
	ErrInvalidOperatorComparison          = errors.New("invalid operator comparison")
	ErrInvalidOperatorComposite           = errors.New("invalid operator composite")
	ErrInvalidOperatorJoin                = errors.New("invalid operator join")
	ErrInvalidOperatorLogical             = errors.New("invalid operator logical")
	ErrInvalidOperatorOrder               = errors.New("invalid operator order")
	ErrInvalidOperatorUnion               = errors.New("invalid operator union")
	ErrInvalidService                     = errors.New("invalid service")
	ErrInvalidServiceFunction             = errors.New("invalid service function")
	ErrInvalidServiceManagement           = errors.New("invalid service management")
	ErrInvalidServiceModifier             = errors.New("invalid service modifer")
	ErrInvalidSubquery                    = errors.New("invalid subquery")
	ErrInvalidValue                       = errors.New("invalid value")
	ErrOverflowArray                      = errors.New("overflow array")
	ErrOverflowConstant                   = errors.New("overflow constant")
	ErrOverflowIdentAlias                 = errors.New("overflow identifier alias")
	ErrOverflowIdentName                  = errors.New("overflow identifier name")
	ErrOverflowLimit                      = errors.New("overflow limit")
	ErrOverflowLiteral                    = errors.New("overflow literal")
	ErrOverflowValueByte                  = errors.New("overflow value byte")
	ErrOverflowValueString                = errors.New("overflow value string")
	ErrUnsupportConstant                  = errors.New("unsupport constant")
	ErrUnsupportConstantString            = errors.New("unsupport constant string")
	ErrUnsupportComparison                = errors.New("unsupport comparison")
	ErrUnsupportFunction                  = errors.New("unsupport function")
	ErrUnsupportFunctionDistinct          = errors.New("unsupport function distinct")
	ErrUnsupportFunctionParamMax          = errors.New("unsupport function param max")
	ErrUnsupportFunctionParamMin          = errors.New("unsupport function param min")
	ErrUnsupportIdentAlias                = errors.New("unsupport identifier alias")
	ErrUnsupportIdentName                 = errors.New("unsupport identifier name")
	ErrUnsupportLiteral                   = errors.New("unsupport literal")
	ErrUnsupportOperator                  = errors.New("unsupport operator")
	ErrUnsupportService                   = errors.New("unsupport service")
	ErrUnsupportStatement                 = errors.New("unsupport statement")
	ErrUnsupportSymbol                    = errors.New("unsupport symbol")
	ErrUnsupportValue                     = errors.New("unsupport value")
	ErrUntransformComparison              = errors.New("untransform comparison ")
	ErrUntransformFunction                = errors.New("untransform function")
	ErrUntransformParam                   = errors.New("untransform param")
	ErrUntransformType                    = errors.New("untransform type")
)

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

// Публичные конструкторы
func NewDelete(from SourceBase) *stmtDelete {
	return &stmtDelete{
		command: uastManagementDelete,
		from:    from,
	}
}
func NewInsert(into SourceBase) *stmtInsert {
	return &stmtInsert{
		command: uastManagementInsert,
		into:    into,
	}
}
func NewSelect(from SourceBase) *stmtSelect {
	return &stmtSelect{
		command: uastManagementSelect,
		from:    from,
	}
}
func NewUpdate(onto SourceBase) *stmtUpdate {
	return &stmtUpdate{
		command: uastManagementUpdate,
		onto:    onto,
	}
}

// Публичные функции
func CaseIf[OutT typeScalar](pairs ...casePair[OutT]) []casePair[OutT] {
	return pairs
}
func CaseElse[OutT typeScalar](else_ ExpressionSafe[OutT]) ExpressionSafe[OutT] {
	return else_
}
func CasePair[InT, OutT typeScalar](when ExpressionSafe[InT], then ExpressionSafe[OutT]) casePair[OutT] {
	return casePair[OutT]{
		then: then,
		when: when,
	}
}
func JsonGroup(expressions []ExpressionBase, values ...ExpressionBase) *exprJson {
	return &exprJson{
		expressions: expressions,
		values:      values,
	}
}
func JsonIndex(index int) *exprLiteral[int] {
	return &exprLiteral[int]{
		value: index,
	}
}
func JsonKey(key string) *exprLiteral[string] {
	return &exprLiteral[string]{
		value: key,
	}
}
func JsonPair(key ExpressionSafe[string], value ExpressionBase) *exprJson {
	return &exprJson{
		expressions: []ExpressionBase{key},
		operator:    uastCompositeCommaSpace,
		values:      []ExpressionBase{value},
	}
}
func JsonPath(path ...ExpressionBase) []ExpressionBase {
	return path
}
func Row(values ...ExpressionBase) []ExpressionBase {
	return values
}

// -- Array -- //
func Array[T typeScalar](array ...T) *exprArray[T] {
	return &exprArray[T]{
		array: array,
	}
}

// -- Binary -- //
func BitwiseAnd[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprBinary[T]{
		left:     left,
		operator: uastBinaryBitwiseAnd,
		right:    right,
	}
}
func BitwiseOr[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprBinary[T]{
		left:     left,
		operator: uastBinaryBitwiseOr,
		right:    right,
	}
}
func BitwiseXor[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprBinary[T]{
		left:     left,
		operator: uastBinaryBitwiseXor,
		right:    right,
	}
}
func Divide[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprBinary[T]{
		left:     left,
		operator: uastBinaryDivide,
		right:    right,
	}
}
func Minus[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprBinary[T]{
		left:     left,
		operator: uastBinaryMinus,
		right:    right,
	}
}
func Modulo[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprBinary[T]{
		left:     left,
		operator: uastBinaryModulo,
		right:    right,
	}
}
func Multiply[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprBinary[T]{
		left:     left,
		operator: uastBinaryMultiply,
		right:    right,
	}
}
func Plus[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprBinary[T]{
		left:     left,
		operator: uastBinaryPlus,
		right:    right,
	}
}
func ShiftLeft[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprBinary[T]{
		left:     left,
		operator: uastBinaryShiftLeft,
		right:    right,
	}
}
func ShiftRight[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprBinary[T]{
		left:     left,
		operator: uastBinaryShiftRight,
		right:    right,
	}
}

// -- Column -- //
func Column[T typeScalar](tableAlias, columnName string) *exprColumn[T] {
	return &exprColumn[T]{
		columnName: columnName,
		tableAlias: tableAlias,
	}
}

// -- Comparison -- //
func Between[T typeScalar](left, valueEnd, valueStart ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:       left,
		operator:   uastComparisonBetween,
		valueEnd:   valueEnd,
		valueGap:   uastLogicalAnd,
		valueStart: valueStart,
	}
}
func Equal[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonEqual,
		right:    right,
	}
}
func Exists[T typeScalar](left ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonExists,
	}
}
func Greater[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonGreater,
		right:    right,
	}
}
func GreaterEqual[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonGreaterEqual,
		right:    right,
	}
}
func ILike[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonILike,
		right:    right,
	}
}
func In[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonIn,
		right:    right,
	}
}
func IsNotNull[T typeScalar](left ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonIsNotNull,
	}
}
func IsNull[T typeScalar](left ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonIsNull,
	}
}
func Less[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonLess,
		right:    right,
	}
}
func LessEqual[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonLessEqual,
		right:    right,
	}
}
func Like[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonLike,
		right:    right,
	}
}
func NotBetween[T typeScalar](left, valueEnd, valueStart ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:       left,
		operator:   uastComparisonNotBetween,
		valueEnd:   valueEnd,
		valueGap:   uastLogicalAnd,
		valueStart: valueStart,
	}
}
func NotEqual[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonNotEqual,
		right:    right,
	}
}
func NotExists[T typeScalar](left ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonNotExists,
	}
}
func NotILike[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonNotILike,
		right:    right,
	}
}
func NotIn[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonNotIn,
		right:    right,
	}
}
func NotLike[T typeScalar](left, right ExpressionSafe[T]) ExpressionSafe[T] {
	return &exprComparison[T]{
		left:     left,
		operator: uastComparisonNotLike,
		right:    right,
	}
}

// -- Constant -- //
func ConstBoolFalse() *exprConstant[bool] {
	return &exprConstant[bool]{
		value: false,
	}
}
func ConstBoolTrue() *exprConstant[bool] {
	return &exprConstant[bool]{
		value: true,
	}
}
func ConstFloat32One() *exprConstant[float32] {
	return &exprConstant[float32]{
		value: 1.0,
	}
}
func ConstFloat64One() *exprConstant[float64] {
	return &exprConstant[float64]{
		value: 1.0,
	}
}
func ConstIntOne() *exprConstant[int] {
	return &exprConstant[int]{
		value: 1,
	}
}
func ConstInt8One() *exprConstant[int8] {
	return &exprConstant[int8]{
		value: 1,
	}
}
func ConstInt16One() *exprConstant[int16] {
	return &exprConstant[int16]{
		value: 1,
	}
}
func ConstInt32One() *exprConstant[int32] {
	return &exprConstant[int32]{
		value: 1,
	}
}
func ConstInt64One() *exprConstant[int64] {
	return &exprConstant[int64]{
		value: 1,
	}
}
func ConstStringDefault() *exprConstant[string] {
	return &exprConstant[string]{
		value: uastConstStringDefault,
	}
}
func ConstStringNull() *exprConstant[string] {
	return &exprConstant[string]{
		value: uastConstStringNull,
	}
}
func ConstUintOne() *exprConstant[uint] {
	return &exprConstant[uint]{
		value: 1,
	}
}
func ConstUint8One() *exprConstant[uint8] {
	return &exprConstant[uint8]{
		value: 1,
	}
}
func ConstUint16One() *exprConstant[uint16] {
	return &exprConstant[uint16]{
		value: 1,
	}
}
func ConstUint32One() *exprConstant[uint32] {
	return &exprConstant[uint32]{
		value: 1,
	}
}
func ConstUint64One() *exprConstant[uint64] {
	return &exprConstant[uint64]{
		value: 1,
	}
}

// -- Function [aggregate] -- //
func Avg[T typeNumeric](number ExpressionSafe[T], distinct bool) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		distinct: distinct,
		left:     number,
		process:  uastProcessDirect,
		service:  uastFunctionAvg,
	}
}
func BitAnd[T typeNumeric](number ExpressionSafe[T], distinct bool) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		distinct: distinct,
		left:     number,
		process:  uastProcessDirect,
		service:  uastFunctionBitAnd,
	}
}
func BitOr[T typeNumeric](number ExpressionSafe[T], distinct bool) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		distinct: distinct,
		left:     number,
		process:  uastProcessDirect,
		service:  uastFunctionBitOr,
	}
}
func BitXor[T typeNumeric](number ExpressionSafe[T], distinct bool) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		distinct: distinct,
		left:     number,
		process:  uastProcessDirect,
		service:  uastFunctionBitXor,
	}
}
func Count[T typeScalar](number ExpressionSafe[T], distinct bool) *exprFunction[T, T, int64] {
	return &exprFunction[T, T, int64]{
		distinct: distinct,
		left:     number,
		process:  uastProcessDirect,
		service:  uastFunctionCount,
	}
}
func GroupConcat(name ExpressionSafe[string], distinct bool) *exprFunction[string, string, string] {
	return &exprFunction[string, string, string]{
		distinct: distinct,
		left:     name,
		operator: uastCompositeSingleSpace,
		process:  uastProcessDirect,
		service:  uastFunctionGroupConcat,
	}
}
func Max[T typeScalar](number ExpressionSafe[T], distinct bool) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		distinct: distinct,
		left:     number,
		process:  uastProcessDirect,
		service:  uastFunctionMax,
	}
}
func Min[T typeScalar](number ExpressionSafe[T], distinct bool) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		distinct: distinct,
		left:     number,
		process:  uastProcessDirect,
		service:  uastFunctionMin,
	}
}
func StdDev[T typeNumeric](number ExpressionSafe[T], distinct bool) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		distinct: distinct,
		left:     number,
		process:  uastProcessDirect,
		service:  uastFunctionStdDev,
	}
}
func Sum[T typeNumeric](number ExpressionSafe[T], distinct bool) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		distinct: distinct,
		left:     number,
		process:  uastProcessDirect,
		service:  uastFunctionSum,
	}
}
func Variance[T typeNumeric](number ExpressionSafe[T], distinct bool) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		distinct: distinct,
		left:     number,
		process:  uastProcessDirect,
		service:  uastFunctionVariance,
	}
}

// -- Function [analytical] -- //
func FirstValue[T typeScalar](expr ExpressionSafe[T]) *exprFunction[T, string, T] {
	return &exprFunction[T, string, T]{
		left:    expr,
		service: uastFunctionFirstValue,
		process: uastProcessWindow,
	}
}
func Lag[T typeScalar](expr ExpressionSafe[T], offset int) *exprFunction[T, int, T] {
	return &exprFunction[T, int, T]{
		left: expr,
		right: &exprLiteral[int]{
			value: offset,
		},
		service: uastFunctionLag,
		process: uastProcessWindow,
	}
}
func LastValue[T typeScalar](expr ExpressionSafe[T]) *exprFunction[T, string, T] {
	return &exprFunction[T, string, T]{
		left:    expr,
		service: uastFunctionLastValue,
		process: uastProcessWindow,
	}
}
func Lead[T typeScalar](expr ExpressionSafe[T], offset int) *exprFunction[T, int, T] {
	return &exprFunction[T, int, T]{
		left: expr,
		right: &exprLiteral[int]{
			value: offset,
		},
		service: uastFunctionLead,
		process: uastProcessWindow,
	}
}
func NthValue[T typeScalar](expr ExpressionSafe[T], n int) *exprFunction[T, int, T] {
	return &exprFunction[T, int, T]{
		left: expr,
		right: &exprLiteral[int]{
			value: n,
		},
		service: uastFunctionNthValue,
		process: uastProcessWindow,
	}
}

// -- Function [condition] -- //
func Case[OutT typeScalar](pairs []casePair[OutT], else_ ...ExpressionSafe[OutT]) *exprFunction[string, string, OutT] {
	valueArray := make([]ExpressionBase, 0, len(pairs)*2+1)
	for _, pair := range pairs {
		valueArray = append(valueArray, pair.when, pair.then)
	}
	if len(else_) > 0 {
		valueArray = append(valueArray, else_[0])
	}
	return &exprFunction[string, string, OutT]{
		process:    uastProcessСross,
		service:    uastFunctionCase,
		valueArray: valueArray,
	}
}
func Coalesce[T typeScalar](expressions ...ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left: &exprComposite[T]{
			expressions: expressions,
			operator:    uastCompositeCommaSpace,
		},
		process: uastProcessDirect,
		service: uastFunctionCoalesce,
	}
}
func Greatest[T typeScalar](expressions ...ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left: &exprComposite[T]{
			expressions: expressions,
			operator:    uastCompositeCommaSpace,
		},
		process: uastProcessDirect,
		service: uastFunctionGreatest,
	}
}
func Least[T typeScalar](expressions ...ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left: &exprComposite[T]{
			expressions: expressions,
			operator:    uastCompositeCommaSpace,
		},
		process: uastProcessDirect,
		service: uastFunctionLeast,
	}
}
func NullIf[T typeScalar](expressionFirst, expressionSecond ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left: &exprComposite[T]{
			expressions: []ExpressionSafe[T]{
				expressionFirst,
				expressionSecond,
			},
			operator: uastCompositeCommaSpace,
		},
		process: uastProcessDirect,
		service: uastFunctionNullIf,
	}
}

// -- Function [convert] -- //
func Cast[T typeScalar](value ExpressionSafe[T], valueType ValueType) *exprFunction[T, string, string] {
	return &exprFunction[T, string, string]{
		left:      value,
		operator:  uastCompositeSingleSpace,
		process:   uastProcessDirect,
		service:   uastFunctionCast,
		valueType: valueType,
	}
}
func CharLength(str ExpressionSafe[string]) *exprFunction[string, string, int] {
	return &exprFunction[string, string, int]{
		left:    str,
		process: uastProcessDirect,
		service: uastFunctionCharLength,
	}
}
func DateFormat(datetime ExpressionSafe[time.Time], mask ExpressionSafe[string]) *exprFunction[time.Time, string, string] {
	return &exprFunction[time.Time, string, string]{
		distinct: false,
		left:     datetime,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    mask,
		service:  uastFunctionDateFormat,
	}
}
func Degrees[T typeNumeric](angle ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    angle,
		process: uastProcessDirect,
		service: uastFunctionDegrees,
	}
}
func Length(str ExpressionSafe[string]) *exprFunction[string, string, int] {
	return &exprFunction[string, string, int]{
		left:    str,
		process: uastProcessDirect,
		service: uastFunctionLength,
	}
}
func Position(str, subStr ExpressionSafe[string]) *exprFunction[string, string, int] {
	return &exprFunction[string, string, int]{
		left: &exprComposite[string]{
			expressions: []ExpressionSafe[string]{
				subStr,
				serviceString(uastModifierIn),
				str,
			},
			operator: uastCompositeSingleSpace,
		},
		process: uastProcessDirect,
		service: uastFunctionPosition,
	}
}
func Radians[T typeNumeric](angle ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    angle,
		process: uastProcessDirect,
		service: uastFunctionRadians,
	}
}

// -- Function [date and time] -- //
func CurDate() *exprFunction[time.Time, time.Time, time.Time] {
	return &exprFunction[time.Time, time.Time, time.Time]{
		process: uastProcessEmpty,
		service: uastFunctionCurDate,
	}
}
func CurTime() *exprFunction[time.Time, time.Time, time.Time] {
	return &exprFunction[time.Time, time.Time, time.Time]{
		process: uastProcessEmpty,
		service: uastFunctionCurTime,
	}
}
func DateAdd(datetime ExpressionSafe[time.Time], interval ExpressionSafe[string]) *exprFunction[time.Time, string, time.Time] {
	return &exprFunction[time.Time, string, time.Time]{
		left:     datetime,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    interval,
		service:  uastFunctionDateAdd,
	}
}
func DateDiff(datetimeEnd, datetimeStart ExpressionSafe[time.Time]) *exprFunction[time.Time, string, int] {
	return &exprFunction[time.Time, string, int]{
		left: &exprComposite[time.Time]{
			expressions: []ExpressionSafe[time.Time]{
				datetimeEnd,
				datetimeStart,
			},
			operator: uastCompositeCommaSpace,
		},
		process: uastProcessDirect,
		service: uastFunctionDateDiff,
	}
}
func DateSub(datetime ExpressionSafe[time.Time], interval ExpressionSafe[string]) *exprFunction[time.Time, string, time.Time] {
	return &exprFunction[time.Time, string, time.Time]{
		left:     datetime,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    interval,
		service:  uastFunctionDateSub,
	}
}
func Day(datetime ExpressionSafe[time.Time]) *exprFunction[time.Time, string, int] {
	return &exprFunction[time.Time, string, int]{
		left:    datetime,
		process: uastProcessDirect,
		service: uastFunctionDay,
	}
}
func DayName(datetime ExpressionSafe[time.Time]) *exprFunction[time.Time, string, string] {
	return &exprFunction[time.Time, string, string]{
		left:    datetime,
		process: uastProcessDirect,
		service: uastFunctionDayName,
	}
}
func Hour(datetime ExpressionSafe[time.Time]) *exprFunction[time.Time, string, int] {
	return &exprFunction[time.Time, string, int]{
		left:    datetime,
		process: uastProcessDirect,
		service: uastFunctionHour,
	}
}
func Minute(datetime ExpressionSafe[time.Time]) *exprFunction[time.Time, string, int] {
	return &exprFunction[time.Time, string, int]{
		left:    datetime,
		process: uastProcessDirect,
		service: uastFunctionMinute,
	}
}
func Month(datetime ExpressionSafe[time.Time]) *exprFunction[time.Time, string, int] {
	return &exprFunction[time.Time, string, int]{
		left:    datetime,
		process: uastProcessDirect,
		service: uastFunctionMonth,
	}
}
func MonthName(datetime ExpressionSafe[time.Time]) *exprFunction[time.Time, string, string] {
	return &exprFunction[time.Time, string, string]{
		left:    datetime,
		process: uastProcessDirect,
		service: uastFunctionMonthName,
	}
}
func Now() *exprFunction[time.Time, time.Time, time.Time] {
	return &exprFunction[time.Time, time.Time, time.Time]{
		process: uastProcessEmpty,
		service: uastFunctionNow,
	}
}
func Quarter(datetime ExpressionSafe[time.Time]) *exprFunction[time.Time, string, int] {
	return &exprFunction[time.Time, string, int]{
		left:    datetime,
		process: uastProcessDirect,
		service: uastFunctionQuarter,
	}
}
func Second(datetime ExpressionSafe[time.Time]) *exprFunction[time.Time, string, int] {
	return &exprFunction[time.Time, string, int]{
		left:    datetime,
		process: uastProcessDirect,
		service: uastFunctionSecond,
	}
}
func TimeAdd(datetime ExpressionSafe[time.Time], interval ExpressionSafe[string]) *exprFunction[time.Time, string, time.Time] {
	return &exprFunction[time.Time, string, time.Time]{
		left:     datetime,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    interval,
		service:  uastFunctionTimeAdd,
	}
}
func TimeDiff(datetimeEnd, datetimeStart ExpressionSafe[time.Time]) *exprFunction[time.Time, string, int] {
	return &exprFunction[time.Time, string, int]{
		left: &exprComposite[time.Time]{
			expressions: []ExpressionSafe[time.Time]{
				datetimeEnd,
				datetimeStart,
			},
			operator: uastCompositeCommaSpace,
		},
		process: uastProcessDirect,
		service: uastFunctionTimeDiff,
	}
}
func TimeSub(datetime ExpressionSafe[time.Time], interval ExpressionSafe[string]) *exprFunction[time.Time, string, time.Time] {
	return &exprFunction[time.Time, string, time.Time]{
		left:     datetime,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    interval,
		service:  uastFunctionTimeSub,
	}
}
func Week(datetime ExpressionSafe[time.Time]) *exprFunction[time.Time, string, int] {
	return &exprFunction[time.Time, string, int]{
		left:    datetime,
		process: uastProcessDirect,
		service: uastFunctionWeek,
	}
}
func Year(datetime ExpressionSafe[time.Time]) *exprFunction[time.Time, string, int] {
	return &exprFunction[time.Time, string, int]{
		left:    datetime,
		process: uastProcessDirect,
		service: uastFunctionYear,
	}
}

// -- Function [json] -- //
func JsonArray[InLT typeScalar](data ExpressionSafe[InLT], values ...ExpressionBase) *exprFunction[InLT, string, string] {
	return &exprFunction[InLT, string, string]{
		json: []*exprJson{
			{
				operator: uastCompositeCommaSpace,
				values:   values,
			},
		},
		left:     data,
		operator: uastCompositeCommaSpace,
		process:  uastProcessJson,
		service:  uastFunctionJsonArray,
	}
}
func JsonArrayAgg[InLT typeScalar](data ExpressionSafe[InLT]) *exprFunction[InLT, string, string] {
	return &exprFunction[InLT, string, string]{
		left:    data,
		process: uastProcessJson,
		service: uastFunctionJsonArrayAgg,
	}
}
func JsonContains[InLT, InRT typeScalar](data ExpressionSafe[InLT], expression ExpressionSafe[InRT]) *exprFunction[InLT, InRT, bool] {
	return &exprFunction[InLT, InRT, bool]{
		json: []*exprJson{
			{
				expressions: []ExpressionBase{
					expression,
				},
				operator: uastCompositeCommaSpace,
			},
		},
		left:     data,
		operator: uastCompositeCommaSpace,
		process:  uastProcessJson,
		service:  uastFunctionJsonContains,
	}
}
func JsonExtract[InLT typeScalar](data ExpressionSafe[InLT], json *exprJson, valueType ValueType) *exprFunction[InLT, string, string] {
	return &exprFunction[InLT, string, string]{
		json: []*exprJson{
			json,
		},
		left:      data,
		operator:  uastCompositeCommaSpace,
		process:   uastProcessJson,
		service:   uastFunctionJsonExtract,
		valueType: valueType,
	}
}
func JsonObject(json ...*exprJson) *exprFunction[string, string, string] {
	return &exprFunction[string, string, string]{
		json:    json,
		process: uastProcessJson,
		service: uastFunctionJsonObject,
	}
}
func JsonObjectAgg[InLT, InRT typeScalar](key ExpressionSafe[InLT], value ExpressionSafe[InRT]) *exprFunction[InLT, InRT, string] {
	return &exprFunction[InLT, InRT, string]{
		json: []*exprJson{
			{
				expressions: []ExpressionBase{key},
				operator:    uastCompositeCommaSpace,
				values:      []ExpressionBase{value},
			},
		},
		process: uastProcessJson,
		service: uastFunctionJsonObjectAgg,
	}
}
func JsonRemove[InLT typeScalar](data ExpressionSafe[InLT], json ...*exprJson) *exprFunction[InLT, string, InLT] {
	return &exprFunction[InLT, string, InLT]{
		json:     json,
		left:     data,
		operator: uastCompositeCommaSpace,
		process:  uastProcessJson,
		service:  uastFunctionJsonRemove,
	}
}
func JsonSet[InLT typeScalar](data ExpressionSafe[InLT], json ...*exprJson) *exprFunction[InLT, string, InLT] {
	return &exprFunction[InLT, string, InLT]{
		json:     json,
		left:     data,
		operator: uastCompositeCommaSpace,
		process:  uastProcessJson,
		service:  uastFunctionJsonSet,
	}
}
func JsonType[InLT typeScalar](data ExpressionSafe[InLT]) *exprFunction[InLT, string, string] {
	return &exprFunction[InLT, string, string]{
		left:    data,
		process: uastProcessJson,
		service: uastFunctionJsonType,
	}
}

// -- Function [math] -- //
func Abs[T typeNumeric](numeric ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    numeric,
		process: uastProcessDirect,
		service: uastFunctionAbs,
	}
}
func ACos[T typeNumeric](angle ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    angle,
		process: uastProcessDirect,
		service: uastFunctionACos,
	}
}
func ASin[T typeNumeric](angle ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    angle,
		process: uastProcessDirect,
		service: uastFunctionASin,
	}
}
func ATan[T typeNumeric](angle ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    angle,
		process: uastProcessDirect,
		service: uastFunctionATan,
	}
}
func ATan2[T typeNumeric](y, x ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left: &exprComposite[T]{
			expressions: []ExpressionSafe[T]{
				y,
				x,
			},
			operator: uastCompositeCommaSpace,
		},
		process: uastProcessDirect,
		service: uastFunctionATan2,
	}
}
func Cbrt[T typeNumeric](numeric ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    numeric,
		process: uastProcessDirect,
		service: uastFunctionCbrt,
	}
}
func Ceil[T typeNumeric](numeric ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    numeric,
		process: uastProcessDirect,
		service: uastFunctionCeil,
	}
}
func Cos[T typeNumeric](angle ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    angle,
		process: uastProcessDirect,
		service: uastFunctionCos,
	}
}
func Exp[T typeNumeric](numeric ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    numeric,
		process: uastProcessDirect,
		service: uastFunctionExp,
	}
}
func Floor[T typeNumeric](numeric ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    numeric,
		process: uastProcessDirect,
		service: uastFunctionFloor,
	}
}
func Ln[T typeNumeric](numeric ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    numeric,
		process: uastProcessDirect,
		service: uastFunctionLn,
	}
}
func Log[T typeNumeric](numeric ExpressionSafe[T], base ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:     numeric,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    base,
		service:  uastFunctionLog,
	}
}
func Mod[T typeNumeric](numeric ExpressionSafe[T], divisor ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:     numeric,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    divisor,
		service:  uastFunctionMod,
	}
}
func Pi() *exprFunction[float64, float64, float64] {
	return &exprFunction[float64, float64, float64]{
		process: uastProcessDirect,
		service: uastFunctionPi,
	}
}
func Power[T typeNumeric](numeric ExpressionSafe[T], exponent ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:     numeric,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    exponent,
		service:  uastFunctionPower,
	}
}
func Rand() *exprFunction[int, int, int] {
	return &exprFunction[int, int, int]{
		process: uastProcessEmpty,
		service: uastFunctionRand,
	}
}
func Round[T typeNumeric](numeric ExpressionSafe[T], precision ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:     numeric,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    precision,
		service:  uastFunctionRound,
	}
}
func Sin[T typeNumeric](angle ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    angle,
		process: uastProcessDirect,
		service: uastFunctionSin,
	}
}
func Sqrt[T typeNumeric](numeric ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    numeric,
		process: uastProcessDirect,
		service: uastFunctionSqrt,
	}
}
func Tan[T typeNumeric](angle ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:    angle,
		process: uastProcessDirect,
		service: uastFunctionTan,
	}
}
func Trunc[T typeNumeric](numeric ExpressionSafe[T], places ExpressionSafe[T]) *exprFunction[T, T, T] {
	return &exprFunction[T, T, T]{
		left:     numeric,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    places,
		service:  uastFunctionTrunc,
	}
}

// -- Function [ranking] -- //
func CumeDist() *exprFunction[string, string, float64] {
	return &exprFunction[string, string, float64]{
		service: uastFunctionCumeDist,
		process: uastProcessWindow,
	}
}
func DenseRank() *exprFunction[string, string, int64] {
	return &exprFunction[string, string, int64]{
		service: uastFunctionDenseRank,
		process: uastProcessWindow,
	}
}
func NTile(n int) *exprFunction[int, string, int64] {
	return &exprFunction[int, string, int64]{
		left: &exprLiteral[int]{
			value: n,
		},
		service: uastFunctionNTile,
		process: uastProcessWindow,
	}
}
func PercentRank() *exprFunction[string, string, float64] {
	return &exprFunction[string, string, float64]{
		service: uastFunctionPercentRank,
		process: uastProcessWindow,
	}
}
func Rank() *exprFunction[string, string, int64] {
	return &exprFunction[string, string, int64]{
		service: uastFunctionRank,
		process: uastProcessWindow,
	}
}
func RowNumber() *exprFunction[string, string, int64] {
	return &exprFunction[string, string, int64]{
		service: uastFunctionRowNumber,
		process: uastProcessWindow,
	}
}

// -- Function [string] -- //
func Concat(strs ...ExpressionSafe[string]) *exprFunction[string, string, string] {
	return &exprFunction[string, string, string]{
		left: &exprComposite[string]{
			expressions: strs,
			operator:    uastCompositeCommaSpace,
		},
		process: uastProcessDirect,
		service: uastFunctionConcat,
	}
}
func ConcatWs(separator ExpressionSafe[string], strs ...ExpressionSafe[string]) *exprFunction[string, string, string] {
	return &exprFunction[string, string, string]{
		left:     separator,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right: &exprComposite[string]{
			expressions: strs,
			operator:    uastCompositeCommaSpace,
		},
		service: uastFunctionConcatWs,
	}
}
func LeftString(str ExpressionSafe[string], count ExpressionSafe[int]) *exprFunction[string, int, string] {
	return &exprFunction[string, int, string]{
		left:     str,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    count,
		service:  uastFunctionLeftString,
	}
}
func Lower(str ExpressionSafe[string]) *exprFunction[string, string, string] {
	return &exprFunction[string, string, string]{
		left:    str,
		process: uastProcessDirect,
		service: uastFunctionLower,
	}
}
func LPad(str ExpressionSafe[string], count ExpressionSafe[int], separator ExpressionSafe[string]) *exprFunction[string, int, string] {
	return &exprFunction[string, int, string]{
		left: &exprComposite[string]{
			expressions: []ExpressionSafe[string]{
				str,
				separator,
			},
			operator: uastCompositeCommaSpace,
		},
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    count,
		service:  uastFunctionLPad,
	}
}
func LTrim(str ExpressionSafe[string]) *exprFunction[string, string, string] {
	return &exprFunction[string, string, string]{
		left:    str,
		process: uastProcessDirect,
		service: uastFunctionLTrim,
	}
}
func Repeat(str ExpressionSafe[string], count ExpressionSafe[int]) *exprFunction[string, int, string] {
	return &exprFunction[string, int, string]{
		left:     str,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    count,
		service:  uastFunctionRepeat,
	}
}
func Replace(str ExpressionSafe[string], strOld, strNew ExpressionSafe[string]) *exprFunction[string, string, string] {
	return &exprFunction[string, string, string]{
		left:     str,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right: &exprComposite[string]{
			expressions: []ExpressionSafe[string]{
				strOld,
				strNew,
			},
			operator: uastCompositeCommaSpace,
		},
		service: uastFunctionReplace,
	}
}
func Reverse(str ExpressionSafe[string]) *exprFunction[string, string, string] {
	return &exprFunction[string, string, string]{
		left:    str,
		process: uastProcessDirect,
		service: uastFunctionReverse,
	}
}
func RightString(str ExpressionSafe[string], count ExpressionSafe[int]) *exprFunction[string, int, string] {
	return &exprFunction[string, int, string]{
		left:     str,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    count,
		service:  uastFunctionRightString,
	}
}
func RPad(str ExpressionSafe[string], count ExpressionSafe[int], separator ExpressionSafe[string]) *exprFunction[string, int, string] {
	return &exprFunction[string, int, string]{
		left: &exprComposite[string]{
			expressions: []ExpressionSafe[string]{
				str,
				separator,
			},
			operator: uastCompositeCommaSpace,
		},
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right:    count,
		service:  uastFunctionRPad,
	}
}
func RTrim(str ExpressionSafe[string]) *exprFunction[string, string, string] {
	return &exprFunction[string, string, string]{
		left:    str,
		process: uastProcessDirect,
		service: uastFunctionRTrim,
	}
}
func SubString(str ExpressionSafe[string], startPos, lengthStr ExpressionSafe[int]) *exprFunction[string, int, string] {
	return &exprFunction[string, int, string]{
		left:     str,
		operator: uastCompositeCommaSpace,
		process:  uastProcessDirect,
		right: &exprComposite[int]{
			expressions: []ExpressionSafe[int]{
				startPos,
				lengthStr,
			},
			operator: uastCompositeCommaSpace,
		},
		service: uastFunctionSubString,
	}
}
func Trim(str ExpressionSafe[string]) *exprFunction[string, string, string] {
	return &exprFunction[string, string, string]{
		left:    str,
		process: uastProcessDirect,
		service: uastFunctionTrim,
	}
}
func Upper(str ExpressionSafe[string]) *exprFunction[string, string, string] {
	return &exprFunction[string, string, string]{
		left:    str,
		process: uastProcessDirect,
		service: uastFunctionUpper,
	}
}

// -- Literal -- //
func Literal[T typeString](value T) *exprLiteral[T] {
	return &exprLiteral[T]{
		value: value,
	}
}

// -- Logical -- //
func And(expressions ...markPredicable) *exprLogical {
	return &exprLogical{
		expressions: expressions,
		operator:    uastLogicalAnd,
	}
}
func Or(expressions ...markPredicable) *exprLogical {
	return &exprLogical{
		expressions: expressions,
		operator:    uastLogicalOr,
	}
}

// -- Order -- //
func Asc(expression markOrderable) *exprOrderBy {
	return &exprOrderBy{
		direction:  false,
		expression: expression,
	}
}
func Desc(expression markOrderable) *exprOrderBy {
	return &exprOrderBy{
		direction:  true,
		expression: expression,
	}
}

// -- Subquery -- //
func Subquery[T typeScalar](statement statement) *exprSubquery[T] {
	return &exprSubquery[T]{
		statement: statement,
	}
}

// -- Value -- //
func Value[T typeScalar](value T) *exprValue[T] {
	return &exprValue[T]{
		value: value,
	}
}

// Публичные методы
func (expression *exprColumn[T]) As(aliasName string) *exprAlias[T] {
	return &exprAlias[T]{
		aliasName:  aliasName,
		expression: expression,
	}
}
func (expression *exprFunction[InLT, InRT, T]) As(aliasName string) *exprAlias[T] {
	return &exprAlias[T]{
		aliasName:  aliasName,
		expression: expression,
	}
}
func (expression *exprFunction[InLT, InRT, T]) Over(options ...WindowOption) *exprFunction[InLT, InRT, T] {
	spec := &WindowSpec{}
	for _, option := range options {
		option(spec)
	}
	copy := *expression
	copy.window = spec
	copy.process = uastProcessWindow
	return &copy
}
func (expression *exprSubquery[T]) As(aliasName string) *exprAlias[T] {
	return &exprAlias[T]{
		aliasName:  aliasName,
		expression: expression,
	}
}
