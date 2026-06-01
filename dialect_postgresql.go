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
		parensFunction:       false,
		placeholderNumber:    0,
		placeholderStyle:     "$",
		placeholderType:      true,
		symbolMarkLeft:       "'",
		symbolMarkRight:      "'",
		symbolQuoteLeft:      `"`,
		symbolQuoteRight:     `"`,
		supportCascade:       true,
		supportIfExists: map[string]bool{
			"INDEX":  true,
			"SCHEMA": true,
			"TABLE":  true,
			"VIEW":   true,
		},
		supportIfNotExists: map[string]bool{
			"INDEX":  true,
			"SCHEMA": true,
			"TABLE":  true,
			"VIEW":   true,
		},
		supportRestartIdentity: true,
		supportReturning:       true,
		supportUpsert:          true,
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
const (
	uastPostgresqlManagementUpsert managementService = "ON CONFLICT DO UPDATE SET"
)
const (
	// Типы бинарные
	uastPostgresqlTypeBinary    typeService = "BYTEA"
	uastPostgresqlTypeVarBinary typeService = "BYTEA"
	// Типы даты и времени
	uastPostgresqlTypeDate      typeService = "DATE"
	uastPostgresqlTypeDateTime  typeService = "TIMESTAMP"
	uastPostgresqlTypeTime      typeService = "TIME"
	uastPostgresqlTypeTimestamp typeService = "TIMESTAMP"
	// Типы числовые
	uastPostgresqlTypeBigInt   typeService = "BIGINT"
	uastPostgresqlTypeDecimal  typeService = "DECIMAL"
	uastPostgresqlTypeDouble   typeService = "DOUBLE PRECISION"
	uastPostgresqlTypeFloat    typeService = "REAL"
	uastPostgresqlTypeInt      typeService = "INTEGER"
	uastPostgresqlTypeSmallInt typeService = "SMALLINT"
	// Типы строковые
	uastPostgresqlTypeChar    typeService = "CHAR"
	uastPostgresqlTypeString  typeService = "VARCHAR"
	uastPostgresqlTypeText    typeService = "TEXT"
	uastPostgresqlTypeVarChar typeService = "VARCHAR"
	// Типы специальные
	uastPostgresqlTypeArray   typeService = "ARRAY"
	uastPostgresqlTypeBoolean typeService = "BOOLEAN"
	uastPostgresqlTypeJson    typeService = "JSONB"
	uastPostgresqlTypeUUID    typeService = "UUID"
	uastPostgresqlTypeXML     typeService = "XML"
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
var listTypePostgresql = map[ValueType]typeService{
	// Типы бинарные
	TypeBinary:    uastPostgresqlTypeBinary,
	TypeVarBinary: uastPostgresqlTypeVarBinary,
	// Типы даты и времени
	TypeDate:      uastPostgresqlTypeDate,
	TypeDateTime:  uastPostgresqlTypeDateTime,
	TypeTime:      uastPostgresqlTypeTime,
	TypeTimestamp: uastPostgresqlTypeTimestamp,
	// Типы числовые
	TypeBigInt:   uastPostgresqlTypeBigInt,
	TypeDecimal:  uastPostgresqlTypeDecimal,
	TypeDouble:   uastPostgresqlTypeDouble,
	TypeFloat:    uastPostgresqlTypeFloat,
	TypeInt:      uastPostgresqlTypeInt,
	TypeSmallInt: uastPostgresqlTypeSmallInt,
	// Типы строковые
	TypeChar:    uastPostgresqlTypeChar,
	TypeString:  uastPostgresqlTypeString,
	TypeText:    uastPostgresqlTypeText,
	TypeVarChar: uastPostgresqlTypeVarChar,
	// Типы специальные
	TypeArray:   uastPostgresqlTypeArray,
	TypeBoolean: uastPostgresqlTypeBoolean,
	TypeJSON:    uastPostgresqlTypeJson,
	TypeUUID:    uastPostgresqlTypeUUID,
	TypeXML:     uastPostgresqlTypeXML,
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
	expr.transformSetService(uastPostgresqlFunctionStdDev)
	return nil
}
func postgresqlFunctionVariance(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastPostgresqlFunctionVariance)
	return nil
}
func postgresqlFunctionCase(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastPostgresqlFunctionCase)
	return nil
}
func postgresqlFunctionCast(baseTransformer *baseTransformer, expr transformFunction) error {
	valueType := expr.transformGetValueType()
	typeService, exists := listTypePostgresql[valueType]
	if !exists {
		return ErrUntransformType
	}
	expr.transformSetRight(&exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastModifierAs),
			serviceString(typeService),
		},
		operator: uastCompositeSingleSpace,
	})
	expr.transformSetService(uastPostgresqlFunctionCast)
	return nil
}
func postgresqlFunctionDateFormat(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastPostgresqlFunctionDateFormat)
	return nil
}
func postgresqlFunctionCurDate(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetProcess(uastProcessEmpty)
	expr.transformSetService(uastPostgresqlFunctionCurDate)
	return nil
}
func postgresqlFunctionCurTime(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetProcess(uastProcessEmpty)
	expr.transformSetService(uastPostgresqlFunctionCurTime)
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
	expr.transformSetProcess(uastProcessEmpty)
	expr.transformSetService(uastPostgresqlFunctionNow)
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
	expr.transformSetService(uastPostgresqlFunctionJsonArrayAgg)
	return nil
}
func postgresqlFunctionJsonContains(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetOperator(uastCompositeSpaceAtGreaterSpace)
	expr.transformSetService(uastPostgresqlFunctionJsonContains)
	return nil
}
func postgresqlFunctionJsonExtract(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.transformGetJson()
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
	valueType := expr.transformGetValueType()
	typeService, exists := listTypePostgresql[valueType]
	if !exists {
		return ErrUntransformType
	}
	switch valueType {
	case TypeJSON:
		expr.transformSetJson([]*exprJson{
			{
				expressions: []ExpressionBase{
					&exprLiteral[string]{
						value: path,
					},
				},
				operator: uastCompositeSingleSpace,
			},
		})
		expr.transformSetOperator(uastCompositeSpaceSignGreaterSpace)
		expr.transformSetService(uastPostgresqlFunctionJsonExtract)
	case TypeString:
		expr.transformSetJson([]*exprJson{
			{
				expressions: []ExpressionBase{
					&exprLiteral[string]{
						value: path,
					},
				},
				operator: uastCompositeSingleSpace,
			},
		})
		expr.transformSetOperator(uastCompositeSpaceSignDoubleGreaterSpace)
		expr.transformSetService(uastPostgresqlFunctionJsonExtract)
	default:
		expr.transformSetJson([]*exprJson{
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
		expr.transformSetOperator(uastCompositeSpaceSignDoubleGreaterSpace)
		expr.transformSetService(uastPostgresqlFunctionJsonExtractCast)
	}
	return nil
}
func postgresqlFunctionJsonObject(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastPostgresqlFunctionJsonObject)
	return nil
}
func postgresqlFunctionJsonObjectAgg(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastPostgresqlFunctionJsonObjectAgg)
	return nil
}
func postgresqlFunctionJsonRemove(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.transformGetJson()
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
	expr.transformSetJson(groups)
	expr.transformSetOperator(uastCompositeSpaceMinusSpace)
	expr.transformSetService(uastPostgresqlFunctionJsonRemove)
	return nil
}
func postgresqlFunctionJsonSet(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.transformGetJson()
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
		expr.transformSetJson([]*exprJson{groups[last-1]})
	} else {
		expr.transformSetJson([]*exprJson{groups[0]})
	}
	expr.transformSetOperator(uastCompositeCommaSpace)
	expr.transformSetProcess(uastProcessJsonRecurcive)
	expr.transformSetService(uastPostgresqlFunctionJsonSet)
	return nil
}
func postgresqlFunctionJsonType(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastPostgresqlFunctionJsonType)
	return nil
}
func postgresqlFunctionRand(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetProcess(uastProcessEmpty)
	expr.transformSetService(uastPostgresqlFunctionRand)
	return nil
}
func postgresqlFunctionATan2(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetProcess(uastProcessInvert)
	expr.transformSetService(uastPostgresqlFunctionATan2)
	return nil
}

// Приватные методы
func (strateger *postgresqlStrateger) renderComment(baseRenderer *baseRenderer, stmtComment *stmtComment) error {
	if err := baseRenderer.renderCommand(stmtComment.command); err != nil {
		return err
	}
	if err := baseRenderer.renderOnColumn(stmtComment.column); err != nil {
		return err
	}
	if err := baseRenderer.renderOnTable(stmtComment.table); err != nil {
		return err
	}
	if err := baseRenderer.renderIsComment(stmtComment.comment); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) renderCreate(baseRenderer *baseRenderer, stmtCreate *stmtCreate) error {
	// !!! Внимание, находится в стадии разработки
	if err := baseRenderer.renderCommand(stmtCreate.command); err != nil {
		return err
	}
	if err := baseRenderer.renderReplace(stmtCreate.replace); err != nil {
		return err
	}
	if err := baseRenderer.renderUnique(stmtCreate.unique); err != nil {
		return err
	}
	if err := baseRenderer.renderEntity(stmtCreate.entity, false, stmtCreate.ifNotExists); err != nil {
		return err
	}
	if err := baseRenderer.renderOnColumns(stmtCreate.table, stmtCreate.columns); err != nil {
		return err
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
	if err := baseRenderer.renderEntity(stmtDrop.entity, stmtDrop.ifExists, false); err != nil {
		return err
	}
	if err := baseRenderer.renderCascade(stmtDrop.cascade); err != nil {
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
	if err := baseRenderer.renderInto(stmtInsert.into); err != nil {
		return err
	}
	if err := baseRenderer.renderColumns(stmtInsert.columns); err != nil {
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
	if err := baseRenderer.renderFields(stmtSelect.fields); err != nil {
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
	if err := baseRenderer.renderCascade(stmtTruncate.cascade); err != nil {
		return err
	}
	if err := baseRenderer.renderRestartIdentity(stmtTruncate.restartIdentity); err != nil {
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
	if err := baseRenderer.renderOnto(stmtUpdate.onto); err != nil {
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
func (strateger *postgresqlStrateger) transformComment(baseTransformer *baseTransformer, stmtComment *stmtComment) error {
	return nil
}
func (strateger *postgresqlStrateger) transformCreate(baseTransformer *baseTransformer, stmtCreate *stmtCreate) error {
	// !!! Внимание, находится в стадии разработки
	if err := baseTransformer.transformColumns(stmtCreate.fields, &stmtCreate.columns); err != nil {
		return err
	}
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
		stmtInsert.values.upsert.service = uastPostgresqlManagementUpsert
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
func (strateger *postgresqlStrateger) validateComment(baseValidator *baseValidator, stmtComment *stmtComment) error {
	if stmtComment.column == nil && stmtComment.table == nil {
		return ErrInvalidStatement
	}
	if err := baseValidator.validateOnColumn(stmtComment.column); err != nil {
		return err
	}
	if err := baseValidator.validateOnTable(stmtComment.table); err != nil {
		return err
	}
	if err := baseValidator.validateIsComment(stmtComment.comment); err != nil {
		return err
	}
	return nil
}
func (strateger *postgresqlStrateger) validateCreate(baseValidator *baseValidator, stmtCreate *stmtCreate) error {
	// !!! Внимание, находится в стадии разработки
	if err := baseValidator.validateEntity(stmtCreate.entity); err != nil {
		return err
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
	if err := baseValidator.validateWith(stmtInsert.with); err != nil {
		return err
	}
	if err := baseValidator.validateInto(stmtInsert.into); err != nil {
		return err
	}
	if err := baseValidator.validateColumns(stmtInsert.columns); err != nil {
		return err
	}
	if (stmtInsert.source == nil && stmtInsert.values == nil) || (stmtInsert.source != nil && stmtInsert.values != nil) {
		return ErrInvalidStatement
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
	if err := baseValidator.validateOnto(stmtUpdate.onto); err != nil {
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
