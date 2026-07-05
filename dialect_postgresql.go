package uast

import (
	"strconv"
	"time"
)

// Публичные переменные
var DialectPostgreSQL = &SupportDialect{
	config: &config{
		lengthMaxArray:       128,
		lengthMaxConst:       63,
		lengthMaxFunc:        48,
		lengthMaxIdent:       63,
		lengthMaxLimit:       63,
		lengthMaxParam:       34464,
		lengthMaxQuery:       1 * 1024 * 1024,
		lengthMaxValueByte:   1024,
		lengthMaxValueString: 128,
		listComparisons:      listComparisonsPostgresql,
		listFunctions:        listFunctionsPostgresql,
		listManagement:       listManagementPostgresql,
		listModifiers:        listModifiersPostgresql,
		listTypes:            listTypesPostgresql,
		orderSupportAttr: []modifierService{
			uastModifierAutoIncrement,
			uastModifierNotNull,
			uastModifierDefault,
		},
		parensFunction:    false,
		placeholderNumber: 0,
		placeholderStyle:  "$",
		placeholderType:   true,
		symbolMarkLeft:    "'",
		symbolMarkRight:   "'",
		symbolQuoteLeft:   `"`,
		symbolQuoteRight:  `"`,
		supportAdd: map[modifierService]bool{
			uastModifierColumn:     true,
			uastModifierConstraint: true,
			uastModifierDefault:    true,
			uastModifierNotNull:    true,
		},
		supportCascade: map[modifierService]bool{
			uastModifierColumn: true,
			uastModifierIndex:  true,
			uastModifierSchema: true,
			uastModifierTable:  true,
			uastModifierView:   true,
		},
		supportComment: map[modifierService]bool{
			uastModifierColumn: true,
			uastModifierIndex:  true,
			uastModifierSchema: true,
			uastModifierTable:  true,
			uastModifierView:   true,
		},
		supportDrop: map[modifierService]bool{
			uastModifierColumn:     true,
			uastModifierConstraint: true,
			uastModifierDefault:    true,
			uastModifierNotNull:    true,
		},
		supportIfExists: map[modifierService]bool{
			uastModifierColumn: true,
			uastModifierIndex:  true,
			uastModifierSchema: true,
			uastModifierTable:  true,
			uastModifierView:   true,
		},
		supportIfNotExists: map[modifierService]bool{
			uastModifierColumn: true,
			uastModifierIndex:  true,
			uastModifierSchema: true,
			uastModifierTable:  true,
			uastModifierView:   true,
		},
		supportOptions: map[modifierService]bool{
			uastModifierRestartIdentity: true,
			uastModifierUpsert:          true,
		},
		supportRename: map[modifierService]bool{
			uastModifierColumn:     true,
			uastModifierConstraint: true,
			uastModifierIndex:      true,
			uastModifierSchema:     true,
			uastModifierTable:      true,
			uastModifierView:       true,
		},
		supportReturning: true,
		supportSet: map[modifierService]bool{
			uastModifierType: true,
		},
	},
	name:      "PostgreSQL",
	strateger: &postgresqlStrateger{},
}

// Приватные константы
const (
	// Функции агрегатные
	uastPostgresqlFunctionGroupConcat functionService = "STRING_AGG"
	uastPostgresqlFunctionStdDev      functionService = "STDDEV_SAMP"
	uastPostgresqlFunctionVariance    functionService = "VAR_SAMP"
	// Функции условий
	uastPostgresqlFunctionCase functionService = "CASE"
	// Функции конвертации
	uastPostgresqlFunctionCast       functionService = "CAST"
	uastPostgresqlFunctionDateFormat functionService = "TO_CHAR"
	// Функции даты и времени
	uastPostgresqlFunctionCurDate   functionService = "CURRENT_DATE"
	uastPostgresqlFunctionCurTime   functionService = "CURRENT_TIME"
	uastPostgresqlFunctionDateAdd   functionService = ""
	uastPostgresqlFunctionDateDiff  functionService = "DATE_PART"
	uastPostgresqlFunctionDateSub   functionService = ""
	uastPostgresqlFunctionDay       functionService = "EXTRACT"
	uastPostgresqlFunctionDayName   functionService = "TO_CHAR"
	uastPostgresqlFunctionHour      functionService = "EXTRACT"
	uastPostgresqlFunctionMinute    functionService = "EXTRACT"
	uastPostgresqlFunctionMonth     functionService = "EXTRACT"
	uastPostgresqlFunctionMonthName functionService = "TO_CHAR"
	uastPostgresqlFunctionNow       functionService = "CURRENT_TIMESTAMP"
	uastPostgresqlFunctionQuarter   functionService = "EXTRACT"
	uastPostgresqlFunctionSecond    functionService = "EXTRACT"
	uastPostgresqlFunctionTimeAdd   functionService = ""
	uastPostgresqlFunctionTimeDiff  functionService = "DATE_PART"
	uastPostgresqlFunctionTimeSub   functionService = ""
	uastPostgresqlFunctionWeek      functionService = "EXTRACT"
	uastPostgresqlFunctionYear      functionService = "EXTRACT"
	// Функции обмена данными
	uastPostgresqlFunctionJsonArrayAgg    functionService = "JSON_AGG"
	uastPostgresqlFunctionJsonContains    functionService = ""
	uastPostgresqlFunctionJsonExtract     functionService = ""
	uastPostgresqlFunctionJsonExtractCast functionService = "CAST"
	uastPostgresqlFunctionJsonObject      functionService = "JSON_BUILD_OBJECT"
	uastPostgresqlFunctionJsonObjectAgg   functionService = "JSON_OBJECT_AGG"
	uastPostgresqlFunctionJsonRemove      functionService = ""
	uastPostgresqlFunctionJsonSet         functionService = "jsonb_set"
	uastPostgresqlFunctionJsonType        functionService = "jsonb_typeof"
	// Функции математические
	uastPostgresqlFunctionRand  functionService = "RANDOM"
	uastPostgresqlFunctionATan2 functionService = "ATAN2"
	// Функции строковые
)

// Приватные переменные
var listComparisonsPostgresql = map[comparisonOperator]comparisonTransform{}
var listFunctionsPostgresql = map[functionService]functionTransform{
	// Функции агрегатные
	uastFunctionGroupConcat: postgresqlFunctionGroupConcat,
	uastFunctionStdDev:      postgresqlFunctionStdDev,
	uastFunctionVariance:    postgresqlFunctionVariance,
	// Функции условий
	uastFunctionCase: postgresqlFunctionCase,
	// Функции конвертации
	uastFunctionCast:       postgresqlFunctionCast,
	uastFunctionDateFormat: postgresqlFunctionDateFormat,
	// Функции даты и времени
	uastFunctionCurDate:   postgresqlFunctionCurDate,
	uastFunctionCurTime:   postgresqlFunctionCurTime,
	uastFunctionDateAdd:   postgresqlFunctionDateAdd,
	uastFunctionDateDiff:  postgresqlFunctionDateDiff,
	uastFunctionDateSub:   postgresqlFunctionDateSub,
	uastFunctionDay:       postgresqlFunctionDay,
	uastFunctionDayName:   postgresqlFunctionDayName,
	uastFunctionHour:      postgresqlFunctionHour,
	uastFunctionMinute:    postgresqlFunctionMinute,
	uastFunctionMonth:     postgresqlFunctionMonth,
	uastFunctionMonthName: postgresqlFunctionMonthName,
	uastFunctionNow:       postgresqlFunctionNow,
	uastFunctionQuarter:   postgresqlFunctionQuarter,
	uastFunctionSecond:    postgresqlFunctionSecond,
	uastFunctionTimeAdd:   postgresqlFunctionTimeAdd,
	uastFunctionTimeDiff:  postgresqlFunctionTimeDiff,
	uastFunctionTimeSub:   postgresqlFunctionTimeSub,
	uastFunctionWeek:      postgresqlFunctionWeek,
	uastFunctionYear:      postgresqlFunctionYear,
	// Функции обмена данными
	uastFunctionJsonArrayAgg:  postgresqlFunctionJsonArrayAgg,
	uastFunctionJsonContains:  postgresqlFunctionJsonContains,
	uastFunctionJsonExtract:   postgresqlFunctionJsonExtract,
	uastFunctionJsonObject:    postgresqlFunctionJsonObject,
	uastFunctionJsonObjectAgg: postgresqlFunctionJsonObjectAgg,
	uastFunctionJsonRemove:    postgresqlFunctionJsonRemove,
	uastFunctionJsonSet:       postgresqlFunctionJsonSet,
	uastFunctionJsonType:      postgresqlFunctionJsonType,
	// Функции математические
	uastFunctionRand:  postgresqlFunctionRand,
	uastFunctionATan2: postgresqlFunctionATan2,
	// Функции строковые
}
var listManagementPostgresql = map[managementService]managementService{}
var listModifiersPostgresql = map[modifierService]modifierService{
	uastModifierAutoIncrement: "GENERATED BY DEFAULT AS IDENTITY",
	uastModifierUpsert:        "ON CONFLICT DO UPDATE SET",
}
var listTypesPostgresql = map[ValueType]typeService{
	// Типы бинарные
	TypeBinary:    "BYTEA",
	TypeVarBinary: "BYTEA",
	// Типы даты и времени
	TypeDate:      "DATE",
	TypeDateTime:  "TIMESTAMP",
	TypeTime:      "TIME",
	TypeTimestamp: "TIMESTAMP",
	// Типы числовые
	TypeBigInt:   "BIGINT",
	TypeDecimal:  "DECIMAL",
	TypeDouble:   "DOUBLE PRECISION",
	TypeFloat:    "REAL",
	TypeInt:      "INTEGER",
	TypeSmallInt: "SMALLINT",
	// Типы строковые
	TypeChar:    "CHAR",
	TypeString:  "VARCHAR",
	TypeText:    "TEXT",
	TypeVarChar: "VARCHAR",
	// Типы специальные
	TypeArray:   "ARRAY",
	TypeBoolean: "BOOLEAN",
	TypeJSON:    "JSONB",
	TypeUUID:    "UUID",
	TypeXML:     "XML",
}

// Приватные структуры
type postgresqlStrateger struct{}

// Приватные функции
func postgresqlFunctionGroupConcat(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[string, string, string])
	if !exists {
		return ErrUntransformFunction
	}
	if !exists {
		return ErrUntransformParam
	}
	function.operator = uastCompositeCommaSpace
	function.right = &exprLiteral[string]{value: ","}
	function.service = uastPostgresqlFunctionGroupConcat
	return nil
}
func postgresqlFunctionStdDev(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastPostgresqlFunctionStdDev)
	return nil
}
func postgresqlFunctionVariance(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastPostgresqlFunctionVariance)
	return nil
}
func postgresqlFunctionCase(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastPostgresqlFunctionCase)
	return nil
}
func postgresqlFunctionCast(baseTransformer *baseTransformer, expr transformFunction) error {
	valueType := expr.getValueType()
	typeService, exists := listTypesPostgresql[valueType]
	if !exists {
		return ErrUntransformType
	}
	expr.setRight(&exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastModifierAs),
			serviceString(typeService),
		},
		operator: uastCompositeSingleSpace,
	})
	expr.setService(uastPostgresqlFunctionCast)
	return nil
}
func postgresqlFunctionDateFormat(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastPostgresqlFunctionDateFormat)
	return nil
}
func postgresqlFunctionCurDate(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setProcess(uastProcessEmpty)
	expr.setService(uastPostgresqlFunctionCurDate)
	return nil
}
func postgresqlFunctionCurTime(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setProcess(uastProcessEmpty)
	expr.setService(uastPostgresqlFunctionCurTime)
	return nil
}
func postgresqlFunctionDateAdd(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := function.right
	function.operator = uastCompositeSingleSpace
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			operatorString(uastBinaryPlus),
			serviceString(uastModifierInterval),
			param,
		},
		operator: uastCompositeSingleSpace,
	}
	function.service = uastPostgresqlFunctionDateAdd
	return nil
}
func postgresqlFunctionDateDiff(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	param, exists := function.left.(*exprComposite[time.Time])
	if !exists {
		return ErrUntransformParam
	}
	function.left = &exprBinary[time.Time]{
		left:     param.expressions[0],
		operator: uastBinaryMinus,
		right:    param.expressions[1],
	}
	function.operator = uastCompositeCommaSpace
	function.process = uastProcessInvert
	function.right = &exprLiteral[string]{value: "day"}
	function.service = uastPostgresqlFunctionDateDiff
	return nil
}
func postgresqlFunctionDateSub(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := function.right
	function.operator = uastCompositeSingleSpace
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			operatorString(uastBinaryMinus),
			serviceString(uastModifierInterval),
			param,
		},
		operator: uastCompositeSingleSpace,
	}
	function.service = uastPostgresqlFunctionDateSub
	return nil
}
func postgresqlFunctionDay(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastFunctionDay),
			serviceString(uastManagementFrom),
		},
		operator: uastCompositeSingleSpace,
	}
	function.operator = uastCompositeSingleSpace
	function.process = uastProcessInvert
	function.service = uastPostgresqlFunctionDay
	return nil
}
func postgresqlFunctionDayName(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, string])
	if !exists {
		return ErrUntransformFunction
	}
	function.operator = uastCompositeCommaSpace
	function.right = &exprLiteral[string]{value: "Day"}
	function.service = uastPostgresqlFunctionDayName
	return nil
}
func postgresqlFunctionHour(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastFunctionHour),
			serviceString(uastManagementFrom),
		},
		operator: uastCompositeSingleSpace,
	}
	function.operator = uastCompositeSingleSpace
	function.process = uastProcessInvert
	function.service = uastPostgresqlFunctionHour
	return nil
}
func postgresqlFunctionMinute(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastFunctionMinute),
			serviceString(uastManagementFrom),
		},
		operator: uastCompositeSingleSpace,
	}
	function.operator = uastCompositeSingleSpace
	function.process = uastProcessInvert
	function.service = uastPostgresqlFunctionMinute
	return nil
}
func postgresqlFunctionMonth(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastFunctionMonth),
			serviceString(uastManagementFrom),
		},
		operator: uastCompositeSingleSpace,
	}
	function.operator = uastCompositeSingleSpace
	function.process = uastProcessInvert
	function.service = uastPostgresqlFunctionMonth
	return nil
}
func postgresqlFunctionMonthName(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, string])
	if !exists {
		return ErrUntransformFunction
	}
	function.operator = uastCompositeCommaSpace
	function.right = &exprLiteral[string]{value: "Month"}
	function.service = uastPostgresqlFunctionMonthName
	return nil
}
func postgresqlFunctionNow(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setProcess(uastProcessEmpty)
	expr.setService(uastPostgresqlFunctionNow)
	return nil
}
func postgresqlFunctionQuarter(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastFunctionQuarter),
			serviceString(uastManagementFrom),
		},
		operator: uastCompositeSingleSpace,
	}
	function.operator = uastCompositeSingleSpace
	function.process = uastProcessInvert
	function.service = uastPostgresqlFunctionQuarter
	return nil
}
func postgresqlFunctionSecond(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastFunctionSecond),
			serviceString(uastManagementFrom),
		},
		operator: uastCompositeSingleSpace,
	}
	function.operator = uastCompositeSingleSpace
	function.process = uastProcessInvert
	function.service = uastPostgresqlFunctionSecond
	return nil
}
func postgresqlFunctionTimeAdd(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := function.right
	function.operator = uastCompositeSingleSpace
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			operatorString(uastBinaryPlus),
			serviceString(uastModifierInterval),
			param,
		},
		operator: uastCompositeSingleSpace,
	}
	function.service = uastPostgresqlFunctionTimeAdd
	return nil
}
func postgresqlFunctionTimeDiff(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	param, exists := function.left.(*exprComposite[time.Time])
	if !exists {
		return ErrUntransformParam
	}
	function.left = &exprBinary[time.Time]{
		left:     param.expressions[0],
		operator: uastBinaryMinus,
		right:    param.expressions[1],
	}
	function.operator = uastCompositeCommaSpace
	function.process = uastProcessInvert
	function.right = &exprLiteral[string]{value: "time"}
	function.service = uastPostgresqlFunctionTimeDiff
	return nil
}
func postgresqlFunctionTimeSub(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := function.right
	function.operator = uastCompositeSingleSpace
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			operatorString(uastBinaryMinus),
			serviceString(uastModifierInterval),
			param,
		},
		operator: uastCompositeSingleSpace,
	}
	function.service = uastPostgresqlFunctionTimeSub
	return nil
}
func postgresqlFunctionWeek(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastFunctionWeek),
			serviceString(uastManagementFrom),
		},
		operator: uastCompositeSingleSpace,
	}
	function.operator = uastCompositeSingleSpace
	function.process = uastProcessInvert
	function.service = uastPostgresqlFunctionWeek
	return nil
}
func postgresqlFunctionYear(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastFunctionYear),
			serviceString(uastManagementFrom),
		},
		operator: uastCompositeSingleSpace,
	}
	function.operator = uastCompositeSingleSpace
	function.process = uastProcessInvert
	function.service = uastPostgresqlFunctionYear
	return nil
}
func postgresqlFunctionJsonArrayAgg(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastPostgresqlFunctionJsonArrayAgg)
	return nil
}
func postgresqlFunctionJsonContains(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setOperator(uastCompositeSpaceAtGreaterSpace)
	expr.setService(uastPostgresqlFunctionJsonContains)
	return nil
}
func postgresqlFunctionJsonExtract(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.getJson()
	path := string(uastCompositeBraceLeft)
	for j, expression := range json[0].expressions {
		if j > 0 {
			path += string(uastCompositeSingleComma)
		}
		switch e := expression.(type) {
		case *exprLiteral[int]:
			path += strconv.Itoa(e.value)
		case *exprLiteral[string]:
			path += e.value
		default:
			return ErrInvalidLiteral
		}
	}
	path += string(uastCompositeBraceRight)
	valueType := expr.getValueType()
	typeService, exists := listTypesPostgresql[valueType]
	if !exists {
		return ErrUntransformType
	}
	switch valueType {
	case TypeJSON:
		expr.setJson([]*exprJson{
			{
				expressions: []ExpressionBase{
					&exprLiteral[string]{
						value: path,
					},
				},
				operator: uastCompositeSingleSpace,
			},
		})
		expr.setOperator(uastCompositeSpaceSignGreaterSpace)
		expr.setService(uastPostgresqlFunctionJsonExtract)
	case TypeString:
		expr.setJson([]*exprJson{
			{
				expressions: []ExpressionBase{
					&exprLiteral[string]{
						value: path,
					},
				},
				operator: uastCompositeSingleSpace,
			},
		})
		expr.setOperator(uastCompositeSpaceSignDoubleGreaterSpace)
		expr.setService(uastPostgresqlFunctionJsonExtract)
	default:
		expr.setJson([]*exprJson{
			{
				expressions: []ExpressionBase{
					&exprLiteral[string]{
						value: path,
					},
					serviceString(uastModifierAs),
					serviceString(typeService),
				},
				operator: uastCompositeSingleSpace,
			},
		})
		expr.setOperator(uastCompositeSpaceSignDoubleGreaterSpace)
		expr.setService(uastPostgresqlFunctionJsonExtractCast)
	}
	return nil
}
func postgresqlFunctionJsonObject(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastPostgresqlFunctionJsonObject)
	return nil
}
func postgresqlFunctionJsonObjectAgg(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastPostgresqlFunctionJsonObjectAgg)
	return nil
}
func postgresqlFunctionJsonRemove(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.getJson()
	groups := make([]*exprJson, len(json))
	for i, group := range json {
		path := string(uastCompositeBraceLeft)
		for j, expression := range group.expressions {
			if j > 0 {
				path += string(uastCompositeSingleComma)
			}
			switch e := expression.(type) {
			case *exprLiteral[int]:
				path += strconv.Itoa(e.value)
			case *exprLiteral[string]:
				path += e.value
			default:
				return ErrInvalidLiteral
			}
		}
		path += string(uastCompositeBraceRight)
		groups[i] = &exprJson{
			expressions: []ExpressionBase{
				&exprLiteral[string]{
					value: path,
				},
			},
			operator: uastCompositeSpaceMinusSpace,
		}
	}
	expr.setJson(groups)
	expr.setOperator(uastCompositeSpaceMinusSpace)
	expr.setService(uastPostgresqlFunctionJsonRemove)
	return nil
}
func postgresqlFunctionJsonSet(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.getJson()
	groups := make([]*exprJson, len(json))
	for i, group := range json {
		path := string(uastCompositeBraceLeft)
		for j, expression := range group.expressions {
			if j > 0 {
				path += string(uastCompositeSingleComma)
			}
			switch e := expression.(type) {
			case *exprLiteral[int]:
				path += strconv.Itoa(e.value)
			case *exprLiteral[string]:
				path += e.value
			default:
				return ErrInvalidLiteral
			}
		}
		path += string(uastCompositeBraceRight)
		groups[i] = &exprJson{
			expressions: []ExpressionBase{
				&exprLiteral[string]{value: path},
			},
			operator: uastCompositeCommaSpace,
			values:   group.values,
		}
	}
	last := len(json)
	if last >= 2 {
		for i := last - 1; i > 0; i-- {
			groups[i].children = []*exprJson{groups[i-1]}
		}
		expr.setJson([]*exprJson{groups[last-1]})
	} else {
		expr.setJson([]*exprJson{groups[0]})
	}
	expr.setOperator(uastCompositeCommaSpace)
	expr.setProcess(uastProcessJsonRecurcive)
	expr.setService(uastPostgresqlFunctionJsonSet)
	return nil
}
func postgresqlFunctionJsonType(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastPostgresqlFunctionJsonType)
	return nil
}
func postgresqlFunctionRand(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setProcess(uastProcessEmpty)
	expr.setService(uastPostgresqlFunctionRand)
	return nil
}
func postgresqlFunctionATan2(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setProcess(uastProcessInvert)
	expr.setService(uastPostgresqlFunctionATan2)
	return nil
}

// Приватные методы
func (strateger *postgresqlStrateger) renderAlter(baseRenderer *baseRenderer, stmtAlter *stmtAlter) error {
	// !!!Внимание, находится в стадии разработки
	if err := baseRenderer.renderCommand(stmtAlter.command); err != nil {
		return err
	}
	if err := baseRenderer.renderEntity(stmtAlter.entity, false, false, stmtAlter.ifExists, stmtAlter.ifNotExists); err != nil {
		return err
	}
	switch stmtAlter.entity.(type) {
	case *sourceIndex:
		if stmtAlter.renameColumn != nil {
			if err := baseRenderer.renderRenameColumn(stmtAlter.renameColumn); err != nil {
				return err
			}
		}
		if stmtAlter.renameConstraint != nil {
			if err := baseRenderer.renderRenameConstraint(stmtAlter.renameConstraint); err != nil {
				return err
			}
		}
		if stmtAlter.renameTo != "" {
			if err := baseRenderer.renderRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	case *sourceSchema:
		if stmtAlter.renameColumn != nil {
			if err := baseRenderer.renderRenameColumn(stmtAlter.renameColumn); err != nil {
				return err
			}
		}
		if stmtAlter.renameConstraint != nil {
			if err := baseRenderer.renderRenameConstraint(stmtAlter.renameConstraint); err != nil {
				return err
			}
		}
		if stmtAlter.renameTo != "" {
			if err := baseRenderer.renderRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	case *sourceTable:
		if err := baseRenderer.renderColumns(stmtAlter.command, stmtAlter.addColumns, stmtAlter.addConstraints, stmtAlter.dropColumns, stmtAlter.dropConstraints); err != nil {
			return err
		}
		if stmtAlter.renameColumn != nil {
			if err := baseRenderer.renderRenameColumn(stmtAlter.renameColumn); err != nil {
				return err
			}
		}
		if stmtAlter.renameConstraint != nil {
			if err := baseRenderer.renderRenameConstraint(stmtAlter.renameConstraint); err != nil {
				return err
			}
		}
		if stmtAlter.renameTo != "" {
			if err := baseRenderer.renderRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	case *sourceView:
		if stmtAlter.renameColumn != nil {
			if err := baseRenderer.renderRenameColumn(stmtAlter.renameColumn); err != nil {
				return err
			}
		}
		if stmtAlter.renameConstraint != nil {
			if err := baseRenderer.renderRenameConstraint(stmtAlter.renameConstraint); err != nil {
				return err
			}
		}
		if stmtAlter.renameTo != "" {
			if err := baseRenderer.renderRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	}
	return nil
}
func (strateger *postgresqlStrateger) renderComment(baseRenderer *baseRenderer, stmtComment *stmtComment) error {
	if err := baseRenderer.renderCommand(stmtComment.command); err != nil {
		return err
	}
	if err := baseRenderer.renderOnFrom(stmtComment.onTable, stmtComment.onColumn); err != nil {
		return err
	}
	if err := baseRenderer.renderIsData(stmtComment.comment); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) renderCreate(baseRenderer *baseRenderer, stmtCreate *stmtCreate) error {
	if err := baseRenderer.renderCommand(stmtCreate.command); err != nil {
		return err
	}
	if err := baseRenderer.renderEntity(stmtCreate.entity, stmtCreate.isReplace, stmtCreate.isUnique, false, stmtCreate.ifNotExists); err != nil {
		return err
	}
	switch stmtCreate.entity.(type) {
	case *sourceIndex:
		if err := baseRenderer.renderOn(stmtCreate.on); err != nil {
			return err
		}
		if err := baseRenderer.renderIndex(stmtCreate.columns); err != nil {
			return err
		}
	case *sourceSchema:
	case *sourceTable:
		if err := baseRenderer.renderColumns(stmtCreate.command, stmtCreate.columns, stmtCreate.constraints, nil, nil); err != nil {
			return err
		}
	case *sourceView:
		if err := baseRenderer.renderAs(); err != nil {
			return err
		}
		if err := baseRenderer.renderSource(stmtCreate.source); err != nil {
			return err
		}
	}
	return nil
}
func (strateger *postgresqlStrateger) renderDelete(baseRenderer *baseRenderer, stmtDelete *stmtDelete) error {
	if err := baseRenderer.renderWith(stmtDelete.with); err != nil {
		return err
	}
	if err := baseRenderer.renderCommand(stmtDelete.command); err != nil {
		return err
	}
	if err := baseRenderer.renderFrom(stmtDelete.from); err != nil {
		return err
	}
	if err := baseRenderer.renderUsing(stmtDelete.join); err != nil {
		return err
	}
	if err := baseRenderer.renderWhere(stmtDelete.where); err != nil {
		return err
	}
	if err := baseRenderer.renderReturning(stmtDelete.returning); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) renderDrop(baseRenderer *baseRenderer, stmtDrop *stmtDrop) error {
	if err := baseRenderer.renderCommand(stmtDrop.command); err != nil {
		return err
	}
	if err := baseRenderer.renderEntity(stmtDrop.entity, false, false, stmtDrop.ifExists, false); err != nil {
		return err
	}
	if err := baseRenderer.renderCascade(stmtDrop.entity, stmtDrop.isCascade); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) renderInsert(baseRenderer *baseRenderer, stmtInsert *stmtInsert) error {
	if err := baseRenderer.renderWith(stmtInsert.with); err != nil {
		return err
	}
	if err := baseRenderer.renderCommand(stmtInsert.command); err != nil {
		return err
	}
	if err := baseRenderer.renderInTo(stmtInsert.inTo); err != nil {
		return err
	}
	if err := baseRenderer.renderFields(stmtInsert.fields, true); err != nil {
		return err
	}
	if err := baseRenderer.renderSource(stmtInsert.source); err != nil {
		return err
	}
	if err := baseRenderer.renderValues(stmtInsert.values); err != nil {
		return err
	}
	if err := baseRenderer.renderReturning(stmtInsert.returning); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) renderSelect(baseRenderer *baseRenderer, stmtSelect *stmtSelect) error {
	if err := baseRenderer.renderWith(stmtSelect.with); err != nil {
		return err
	}
	if err := baseRenderer.renderCommand(stmtSelect.command); err != nil {
		return err
	}
	if err := baseRenderer.renderDistinct(stmtSelect.distinct); err != nil {
		return err
	}
	if err := baseRenderer.renderFields(stmtSelect.fields, false); err != nil {
		return err
	}
	if err := baseRenderer.renderFrom(stmtSelect.from); err != nil {
		return err
	}
	if err := baseRenderer.renderJoin(stmtSelect.join); err != nil {
		return err
	}
	if err := baseRenderer.renderWhere(stmtSelect.where); err != nil {
		return err
	}
	if err := baseRenderer.renderGroupBy(stmtSelect.groupBy); err != nil {
		return err
	}
	if err := baseRenderer.renderHaving(stmtSelect.having); err != nil {
		return err
	}
	if err := baseRenderer.renderOrderBy(stmtSelect.orderBy); err != nil {
		return err
	}
	if err := baseRenderer.renderPagination(stmtSelect.pagination); err != nil {
		return err
	}
	if err := baseRenderer.renderUnions(stmtSelect.unions); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) renderTruncate(baseRenderer *baseRenderer, stmtTruncate *stmtTruncate) error {
	if err := baseRenderer.renderCommand(stmtTruncate.command); err != nil {
		return err
	}
	if err := baseRenderer.renderTable(stmtTruncate.table); err != nil {
		return err
	}
	if err := baseRenderer.renderCascade(stmtTruncate.table, stmtTruncate.isCascade); err != nil {
		return err
	}
	if err := baseRenderer.renderRestartIdentity(stmtTruncate.isRestartIdentity); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) renderUpdate(baseRenderer *baseRenderer, stmtUpdate *stmtUpdate) error {
	if err := baseRenderer.renderWith(stmtUpdate.with); err != nil {
		return err
	}
	if err := baseRenderer.renderCommand(stmtUpdate.command); err != nil {
		return err
	}
	if err := baseRenderer.renderOnTo(stmtUpdate.onTo); err != nil {
		return err
	}
	if err := baseRenderer.renderJoin(stmtUpdate.join); err != nil {
		return err
	}
	if err := baseRenderer.renderSet(stmtUpdate.set); err != nil {
		return err
	}
	if err := baseRenderer.renderWhere(stmtUpdate.where); err != nil {
		return err
	}
	if err := baseRenderer.renderReturning(stmtUpdate.returning); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) transformAlter(baseTransformer *baseTransformer, stmtAlter *stmtAlter) error {
	// !!!Внимание, находится в стадии разработки
	return nil
}
func (strateger *postgresqlStrateger) transformComment(baseTransformer *baseTransformer, stmtComment *stmtComment) error {
	return nil
}
func (strateger *postgresqlStrateger) transformCreate(baseTransformer *baseTransformer, stmtCreate *stmtCreate) error {
	return nil
}
func (strateger *postgresqlStrateger) transformDelete(baseTransformer *baseTransformer, stmtDelete *stmtDelete) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	if len(stmtDelete.join) > 0 {
		conditions := make([]markPredicable, 0)
		for _, join := range stmtDelete.join {
			if join.operator != uastJoinCross && join.expression != nil {
				conditions = append(conditions, join.expression)
			}
		}
		if stmtDelete.where != nil {
			conditions = append(conditions, stmtDelete.where)
		}
		stmtDelete.where = And(conditions...)
	}
	return nil
}
func (strateger *postgresqlStrateger) transformDrop(baseTransformer *baseTransformer, stmtDrop *stmtDrop) error {
	return nil
}
func (strateger *postgresqlStrateger) transformInsert(baseTransformer *baseTransformer, stmtInsert *stmtInsert) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	if stmtInsert.values != nil && stmtInsert.values.upsert != nil {
		if service, exists := listModifiersPostgresql[uastModifierUpsert]; exists {
			stmtInsert.values.upsert.service = service
		}
	}
	return nil
}
func (strateger *postgresqlStrateger) transformSelect(baseTransformer *baseTransformer, stmtSelect *stmtSelect) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) transformTruncate(baseTransformer *baseTransformer, stmtTruncate *stmtTruncate) error {
	return nil
}
func (strateger *postgresqlStrateger) transformUpdate(baseTransformer *baseTransformer, stmtUpdate *stmtUpdate) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) validateAlter(baseValidator *baseValidator, stmtAlter *stmtAlter) error {
	// !!!Внимание, находится в стадии разработки
	if err := baseValidator.validateEntity(stmtAlter.entity); err != nil {
		return err
	}
	switch stmtAlter.entity.(type) {
	case *sourceIndex:
		if stmtAlter.renameColumn != nil {
			if !baseValidator.config.supportRename[uastModifierColumn] {
				return ErrUnsupportEntityIndex
			}
			if err := baseValidator.validateRenameColumn(stmtAlter.renameColumn); err != nil {
				return err
			}
		}
		if stmtAlter.renameConstraint != nil {
			if !baseValidator.config.supportRename[uastModifierConstraint] {
				return ErrUnsupportEntityIndex
			}
			if err := baseValidator.validateRenameConstraint(stmtAlter.renameConstraint); err != nil {
				return err
			}
		}
		if stmtAlter.renameTo != "" {
			if !baseValidator.config.supportRename[uastModifierIndex] {
				return ErrUnsupportEntityIndex
			}
			if err := baseValidator.validateRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	case *sourceSchema:
		if stmtAlter.renameColumn != nil {
			if !baseValidator.config.supportRename[uastModifierColumn] {
				return ErrUnsupportEntitySchema
			}
			if err := baseValidator.validateRenameColumn(stmtAlter.renameColumn); err != nil {
				return err
			}
		}
		if stmtAlter.renameConstraint != nil {
			if !baseValidator.config.supportRename[uastModifierConstraint] {
				return ErrUnsupportEntitySchema
			}
			if err := baseValidator.validateRenameConstraint(stmtAlter.renameConstraint); err != nil {
				return err
			}
		}
		if stmtAlter.renameTo != "" {
			if !baseValidator.config.supportRename[uastModifierSchema] {
				return ErrUnsupportEntitySchema
			}
			if err := baseValidator.validateRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	case *sourceTable:
		if stmtAlter.renameColumn != nil && stmtAlter.renameConstraint != nil {
			return ErrUnsupportEntityTable
		}
		if stmtAlter.renameColumn != nil {
			if !baseValidator.config.supportRename[uastModifierColumn] {
				return ErrUnsupportEntityTable
			}
			if err := baseValidator.validateRenameColumn(stmtAlter.renameColumn); err != nil {
				return err
			}
		}
		if stmtAlter.renameConstraint != nil {
			if !baseValidator.config.supportRename[uastModifierConstraint] {
				return ErrUnsupportEntityTable
			}
			if err := baseValidator.validateRenameConstraint(stmtAlter.renameConstraint); err != nil {
				return err
			}
		}
		if stmtAlter.renameTo != "" {
			if !baseValidator.config.supportRename[uastModifierTable] {
				return ErrUnsupportEntityTable
			}
			if err := baseValidator.validateRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	case *sourceView:
		if stmtAlter.renameColumn != nil {
			if !baseValidator.config.supportRename[uastModifierColumn] {
				return ErrUnsupportEntityView
			}
			if err := baseValidator.validateRenameColumn(stmtAlter.renameColumn); err != nil {
				return err
			}
		}
		if stmtAlter.renameConstraint != nil {
			if !baseValidator.config.supportRename[uastModifierConstraint] {
				return ErrUnsupportEntityView
			}
			if err := baseValidator.validateRenameConstraint(stmtAlter.renameConstraint); err != nil {
				return err
			}
		}
		if stmtAlter.renameTo != "" {
			if !baseValidator.config.supportRename[uastModifierView] {
				return ErrUnsupportEntityView
			}
			if err := baseValidator.validateRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	}
	return nil
}
func (strateger *postgresqlStrateger) validateComment(baseValidator *baseValidator, stmtComment *stmtComment) error {
	if err := baseValidator.validateOnFrom(stmtComment.onTable, stmtComment.onColumn); err != nil {
		return err
	}
	if err := baseValidator.validateIsData(stmtComment.comment); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) validateCreate(baseValidator *baseValidator, stmtCreate *stmtCreate) error {
	if err := baseValidator.validateEntity(stmtCreate.entity); err != nil {
		return err
	}
	switch stmtCreate.entity.(type) {
	case *sourceIndex:
		if err := baseValidator.validateColumns(stmtCreate.columns); err != nil {
			return err
		}
	case *sourceSchema:
	case *sourceTable:
		if err := baseValidator.validateColumns(stmtCreate.columns); err != nil {
			return err
		}
	case *sourceView:
		if err := baseValidator.validateSource(stmtCreate.source); err != nil {
			return err
		}
	default:
		return ErrInvalidStatement
	}
	return nil
}
func (strateger *postgresqlStrateger) validateDelete(baseValidator *baseValidator, stmtDelete *stmtDelete) error {
	if err := baseValidator.validateWith(stmtDelete.with); err != nil {
		return err
	}
	if err := baseValidator.validateFrom(stmtDelete.from); err != nil {
		return err
	}
	if err := baseValidator.validateJoin(stmtDelete.join); err != nil {
		return err
	}
	if err := baseValidator.validateWhere(stmtDelete.where); err != nil {
		return err
	}
	if err := baseValidator.validateReturning(stmtDelete.returning); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) validateDrop(baseValidator *baseValidator, stmtDrop *stmtDrop) error {
	if err := baseValidator.validateEntity(stmtDrop.entity); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) validateInsert(baseValidator *baseValidator, stmtInsert *stmtInsert) error {
	if (stmtInsert.source == nil && stmtInsert.values == nil) || (stmtInsert.source != nil && stmtInsert.values != nil) {
		return ErrInvalidStatement
	}
	if err := baseValidator.validateWith(stmtInsert.with); err != nil {
		return err
	}
	if err := baseValidator.validateInTo(stmtInsert.inTo); err != nil {
		return err
	}
	if err := baseValidator.validateFields(stmtInsert.fields); err != nil {
		return err
	}
	if err := baseValidator.validateSource(stmtInsert.source); err != nil {
		return err
	}
	if err := baseValidator.validateValues(stmtInsert.values); err != nil {
		return err
	}
	if err := baseValidator.validateReturning(stmtInsert.returning); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) validateSelect(baseValidator *baseValidator, stmtSelect *stmtSelect) error {
	if err := baseValidator.validateWith(stmtSelect.with); err != nil {
		return err
	}
	if err := baseValidator.validateFields(stmtSelect.fields); err != nil {
		return err
	}
	if err := baseValidator.validateFrom(stmtSelect.from); err != nil {
		return err
	}
	if err := baseValidator.validateJoin(stmtSelect.join); err != nil {
		return err
	}
	if err := baseValidator.validateWhere(stmtSelect.where); err != nil {
		return err
	}
	if err := baseValidator.validateGroupBy(stmtSelect.groupBy); err != nil {
		return err
	}
	if err := baseValidator.validateHaving(stmtSelect.having); err != nil {
		return err
	}
	if err := baseValidator.validateOrderBy(stmtSelect.orderBy); err != nil {
		return err
	}
	if err := baseValidator.validatePagination(stmtSelect.pagination); err != nil {
		return err
	}
	if err := baseValidator.validateUnions(stmtSelect.unions); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) validateTruncate(baseValidator *baseValidator, stmtTruncate *stmtTruncate) error {
	if err := baseValidator.validateTable(stmtTruncate.table); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) validateUpdate(baseValidator *baseValidator, stmtUpdate *stmtUpdate) error {
	if err := baseValidator.validateWith(stmtUpdate.with); err != nil {
		return err
	}
	if err := baseValidator.validateOnTo(stmtUpdate.onTo); err != nil {
		return err
	}
	if err := baseValidator.validateJoin(stmtUpdate.join); err != nil {
		return err
	}
	if err := baseValidator.validateSet(stmtUpdate.set); err != nil {
		return err
	}
	if err := baseValidator.validateWhere(stmtUpdate.where); err != nil {
		return err
	}
	if err := baseValidator.validateReturning(stmtUpdate.returning); err != nil {
		return err
	}
	return nil
}
