package uast

import (
	"strconv"
	"strings"
	"time"
)

// Публичные переменные
var DialectMsSQL = &SupportDialect{
	config: &config{
		lengthMaxArray:       128,
		lengthMaxConst:       63,
		lengthMaxFunc:        48,
		lengthMaxIdent:       128,
		lengthMaxLimit:       63,
		lengthMaxParam:       2100,
		lengthMaxQuery:       64 * 1024,
		lengthMaxValueByte:   1024,
		lengthMaxValueString: 128,
		listComparisons:      listComparisonsMssql,
		listFunctions:        listFunctionsMssql,
		listManagement:       listManagementMssql,
		listModifiers:        listModifiersMssql,
		listTypes:            listTypesMssql,
		parensFunction:       true,
		placeholderNumber:    0,
		placeholderStyle:     "@p",
		placeholderType:      true,
		symbolMarkLeft:       "'",
		symbolMarkRight:      "'",
		symbolQuoteLeft:      "[",
		symbolQuoteRight:     "]",
		supportAttrCreateOrder: []modifierService{
			uastModifierAutoIncrement,
			uastModifierNotNull,
			uastModifierDefault,
		},
		supportCascade: map[modifierService]bool{
			uastModifierColumn: false,
			uastModifierIndex:  false,
			uastModifierSchema: false,
			uastModifierTable:  false,
			uastModifierView:   false,
		},
		supportComment: map[modifierService]bool{
			uastModifierColumn: false,
			uastModifierIndex:  false,
			uastModifierSchema: false,
			uastModifierTable:  false,
			uastModifierView:   false,
		},
		supportIfExists: map[modifierService]bool{
			uastModifierColumn: false,
			uastModifierIndex:  false,
			uastModifierSchema: true,
			uastModifierTable:  true,
			uastModifierView:   true,
		},
		supportIfNotExists: map[modifierService]bool{
			uastModifierColumn: false,
			uastModifierIndex:  false,
			uastModifierSchema: true,
			uastModifierTable:  false,
			uastModifierView:   true,
		},
		supportRestartIdentity: false,
		supportReturning:       true,
		supportUpsert:          false,
	},
	name:      "MsSQL",
	strateger: &mssqlStrateger{},
}

// Приватные константы
const (
	// Функции агрегатные
	uastMssqlFunctionGroupConcat functionService = "STRING_AGG"
	uastMssqlFunctionStdDev      functionService = "STDEV"
	uastMssqlFunctionVariance    functionService = "VAR"
	// Функции условий
	uastMssqlFunctionCase functionService = "CASE"
	// Функции конвертации
	uastMssqlFunctionCast       functionService = "CAST"
	uastMssqlFunctionDateFormat functionService = "FORMAT"
	// Функции даты и времени
	uastMssqlFunctionCurDate     functionService = "GETDATE()"
	uastMssqlFunctionCurDateCast functionService = "CAST"
	uastMssqlFunctionCurTime     functionService = "GETDATE()"
	uastMssqlFunctionCurTimeCast functionService = "CAST"
	uastMssqlFunctionDateAdd     functionService = "DATEADD"
	uastMssqlFunctionDateSub     functionService = "DATEADD"
	uastMssqlFunctionDayName     functionService = "DATENAME"
	uastMssqlFunctionHour        functionService = "DATEPART"
	uastMssqlFunctionMinute      functionService = "DATEPART"
	uastMssqlFunctionNow         functionService = "GETDATE"
	uastMssqlFunctionMonthName   functionService = "DATENAME"
	uastMssqlFunctionQuarter     functionService = "DATEPART"
	uastMssqlFunctionSecond      functionService = "DATEPART"
	uastMssqlFunctionTimeAdd     functionService = "DATEADD"
	uastMssqlFunctionTimeSub     functionService = "DATEADD"
	uastMssqlFunctionWeek        functionService = "DATEPART"
	// Функции обмена данными
	uastMssqlFunctionJsonExtract      functionService = ""
	uastMssqlFunctionJsonExtractQuery functionService = "JSON_QUERY"
	uastMssqlFunctionJsonExtractValue functionService = "JSON_VALUE"
	uastMssqlFunctionJsonRemove       functionService = "JSON_MODIFY"
	uastMssqlFunctionJsonSet          functionService = "JSON_MODIFY"
	// Функции математические
	uastMssqlFunctionCeil  functionService = "CEILING"
	uastMssqlFunctionTrunc functionService = "ROUND"
	// Функции строковые
	uastMssqlFunctionLength   functionService = "LEN"
	uastMssqlFunctionPosition functionService = "CHARINDEX"
)

// Приватные переменные
var listComparisonsMssql = map[comparisonOperator]comparisonTransform{
	uastComparisonILike:    mssqlComparisonILike,
	uastComparisonNotILike: mssqlComparisonNotILike,
}
var listFunctionsMssql = map[functionService]functionTransform{
	// Функции агрегатные
	uastFunctionGroupConcat: mssqlFunctionGroupConcat,
	uastFunctionStdDev:      mssqlFunctionStdDev,
	uastFunctionVariance:    mssqlFunctionVariance,
	// Функции условий
	uastFunctionCase: mssqlFunctionCase,
	// Функции конвертации
	uastFunctionCast:       mssqlFunctionCast,
	uastFunctionDateFormat: mssqlFunctionDateFormat,
	// Функции даты и времени
	uastFunctionCurDate:   mssqlFunctionCurDate,
	uastFunctionCurTime:   mssqlFunctionCurTime,
	uastFunctionDateAdd:   mssqlFunctionDateAdd,
	uastFunctionDateSub:   mssqlFunctionDateSub,
	uastFunctionDayName:   mssqlFunctionDayName,
	uastFunctionHour:      mssqlFunctionHour,
	uastFunctionMinute:    mssqlFunctionMinute,
	uastFunctionNow:       mssqlFunctionNow,
	uastFunctionMonthName: mssqlFunctionMonthName,
	uastFunctionQuarter:   mssqlFunctionQuarter,
	uastFunctionSecond:    mssqlFunctionSecond,
	uastFunctionTimeAdd:   mssqlFunctionTimeAdd,
	uastFunctionTimeSub:   mssqlFunctionTimeSub,
	uastFunctionWeek:      mssqlFunctionWeek,
	// Функции обмена данными
	uastFunctionJsonExtract: mssqlFunctionJsonExtract,
	uastFunctionJsonRemove:  mssqlFunctionJsonRemove,
	uastFunctionJsonSet:     mssqlFunctionJsonSet,
	// Функции математические
	uastFunctionCeil:  mssqlFunctionCeil,
	uastFunctionTrunc: mssqlFunctionTrunc,
	// Функции строковые
	uastFunctionLength:   mssqlFunctionLength,
	uastFunctionPosition: mssqlFunctionPosition,
}
var listManagementMssql = map[managementService]managementService{
	uastManagementUpsert: "",
}
var listModifiersMssql = map[modifierService]modifierService{
	uastModifierAutoIncrement: "IDENTITY(1,1)",
}
var listTypesMssql = map[ValueType]typeService{
	// Типы бинарные
	TypeBinary:    "BINARY",
	TypeVarBinary: "VARBINARY",
	// Типы даты и времени
	TypeDate:      "DATE",
	TypeDateTime:  "DATETIME2",
	TypeTime:      "TIME",
	TypeTimestamp: "DATETIME2",
	// Типы числовые
	TypeBigInt:   "BIGINT",
	TypeDecimal:  "DECIMAL",
	TypeDouble:   "FLOAT",
	TypeFloat:    "REAL",
	TypeInt:      "INT",
	TypeSmallInt: "SMALLINT",
	// Типы строковые
	TypeChar:    "CHAR",
	TypeString:  "NVARCHAR",
	TypeText:    "NVARCHAR(MAX)",
	TypeVarChar: "NVARCHAR",
	// Типы специальные
	TypeArray:   "NVARCHAR(MAX)",
	TypeBoolean: "BIT",
	TypeJSON:    "NVARCHAR(MAX)",
	TypeUUID:    "UNIQUEIDENTIFIER",
	TypeXML:     "XML",
}

// Приватные структуры
type mssqlStrateger struct{}

// Приватные функции
func mssqlComparisonILike(baseTransformer *baseTransformer, expr transformComparison) error {
	if comparison, ok := expr.(*exprComparison[string]); ok {
		comparison.left = &exprFunction[string, string, string]{
			left:    comparison.left,
			process: uastProcessDirect,
			service: uastFunctionLower,
		}
		comparison.operator = uastComparisonLike
		comparison.right = &exprFunction[string, string, string]{
			left:    comparison.right,
			process: uastProcessDirect,
			service: uastFunctionLower,
		}
		return nil
	}
	return ErrUntransformComparison
}
func mssqlComparisonNotILike(baseTransformer *baseTransformer, expr transformComparison) error {
	if comparison, ok := expr.(*exprComparison[string]); ok {
		comparison.left = &exprFunction[string, string, string]{
			left:    comparison.left,
			process: uastProcessDirect,
			service: uastFunctionLower,
		}
		comparison.operator = uastComparisonNotLike
		comparison.right = &exprFunction[string, string, string]{
			left:    comparison.right,
			process: uastProcessDirect,
			service: uastFunctionLower,
		}
		return nil
	}
	return ErrUntransformComparison
}
func mssqlFunctionGroupConcat(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[string, string, string])
	if !exists {
		return ErrUntransformFunction
	}
	if !exists {
		return ErrUntransformParam
	}
	function.operator = uastCompositeCommaSpace
	function.right = &exprLiteral[string]{value: ","}
	function.service = uastMssqlFunctionGroupConcat
	return nil
}
func mssqlFunctionStdDev(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastMssqlFunctionStdDev)
	return nil
}
func mssqlFunctionVariance(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastMssqlFunctionVariance)
	return nil
}
func mssqlFunctionCase(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastMssqlFunctionCase)
	return nil
}
func mssqlFunctionCast(baseTransformer *baseTransformer, expr transformFunction) error {
	valueType := expr.getValueType()
	typeService, exists := listTypesMssql[valueType]
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
	expr.setService(uastMssqlFunctionCast)
	return nil
}
func mssqlFunctionDateFormat(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastMssqlFunctionDateFormat)
	return nil
}
func mssqlFunctionCurDate(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setLeft(&exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastMssqlFunctionCurDate),
			serviceString(uastModifierAs),
			serviceString(listTypesMysql[TypeDate]),
		},
		operator: uastCompositeSingleSpace,
	})
	expr.setProcess(uastProcessDirect)
	expr.setService(uastMssqlFunctionCurDateCast)
	return nil
}
func mssqlFunctionCurTime(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setLeft(&exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastMssqlFunctionCurTime),
			serviceString(uastModifierAs),
			serviceString(listTypesMysql[TypeTime]),
		},
		operator: uastCompositeSingleSpace,
	})
	expr.setProcess(uastProcessDirect)
	expr.setService(uastMssqlFunctionCurTimeCast)
	return nil
}
func mssqlFunctionDateAdd(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	number, interval, ok := strings.Cut(function.right.(*exprLiteral[string]).value, " ")
	if !ok {
		return ErrUntransformParam
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(interval),
			serviceString(number),
		},
		operator: uastCompositeCommaSpace,
	}
	function.process = uastProcessInvert
	function.service = uastMssqlFunctionDateAdd
	return nil
}
func mssqlFunctionDateSub(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	number, interval, ok := strings.Cut(function.right.(*exprLiteral[string]).value, " ")
	if !ok {
		return ErrUntransformParam
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(interval),
			serviceString(string(uastBinaryMinus) + number),
		},
		operator: uastCompositeCommaSpace,
	}
	function.process = uastProcessInvert
	function.service = uastMssqlFunctionDateSub
	return nil
}
func mssqlFunctionDayName(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, string])
	if !exists {
		return ErrUntransformFunction
	}
	function.operator = uastCompositeCommaSpace
	function.process = uastProcessInvert
	function.right = serviceString(uastModifierWeekday)
	function.service = uastMssqlFunctionDayName
	return nil
}
func mssqlFunctionHour(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = serviceString(uastFunctionHour)
	function.operator = uastCompositeCommaSpace
	function.process = uastProcessInvert
	function.service = uastMssqlFunctionHour
	return nil
}
func mssqlFunctionMinute(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = serviceString(uastFunctionMinute)
	function.operator = uastCompositeCommaSpace
	function.process = uastProcessInvert
	function.service = uastMssqlFunctionMinute
	return nil
}
func mssqlFunctionMonthName(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, string])
	if !exists {
		return ErrUntransformFunction
	}
	function.operator = uastCompositeCommaSpace
	function.process = uastProcessInvert
	function.right = serviceString(uastModifierMonth)
	function.service = uastMssqlFunctionMonthName
	return nil
}
func mssqlFunctionNow(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastMssqlFunctionNow)
	return nil
}
func mssqlFunctionQuarter(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = serviceString(uastFunctionQuarter)
	function.operator = uastCompositeCommaSpace
	function.process = uastProcessInvert
	function.service = uastMssqlFunctionQuarter
	return nil
}
func mssqlFunctionSecond(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = serviceString(uastFunctionSecond)
	function.operator = uastCompositeCommaSpace
	function.process = uastProcessInvert
	function.service = uastMssqlFunctionSecond
	return nil
}
func mssqlFunctionTimeAdd(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	number, interval, ok := strings.Cut(function.right.(*exprLiteral[string]).value, " ")
	if !ok {
		return ErrUntransformParam
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(interval),
			serviceString(number),
		},
		operator: uastCompositeCommaSpace,
	}
	function.process = uastProcessInvert
	function.service = uastMssqlFunctionTimeAdd
	return nil
}
func mssqlFunctionTimeSub(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	number, interval, ok := strings.Cut(function.right.(*exprLiteral[string]).value, " ")
	if !ok {
		return ErrUntransformParam
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(interval),
			serviceString(string(uastBinaryMinus) + number),
		},
		operator: uastCompositeCommaSpace,
	}
	function.process = uastProcessInvert
	function.service = uastMssqlFunctionTimeSub
	return nil
}
func mssqlFunctionWeek(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, int])
	if !exists {
		return ErrUntransformFunction
	}
	function.right = serviceString(uastFunctionWeek)
	function.operator = uastCompositeCommaSpace
	function.process = uastProcessInvert
	function.service = uastMssqlFunctionWeek
	return nil
}
func mssqlFunctionJsonExtract(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.getJson()
	path := string(uastCompositeDollarPoint)
	for j, expression := range json[0].expressions {
		switch e := expression.(type) {
		case *exprLiteral[int]:
			path += string(uastCompositeBrackLeft) + strconv.Itoa(e.value) + string(uastCompositeBrackRight)
		case *exprLiteral[string]:
			if j > 0 {
				path += string(uastCompositeSinglePoint)
			}
			path += e.value
		default:
			return ErrInvalidLiteral
		}
	}
	valueType := expr.getValueType()
	typeService, exists := listTypesMysql[valueType]
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
		expr.setOperator(uastCompositeCommaSpace)
		expr.setService(uastMssqlFunctionJsonExtractQuery)
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
		expr.setOperator(uastCompositeCommaSpace)
		expr.setService(uastMssqlFunctionJsonExtractValue)
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
		expr.setOperator(uastCompositeCommaSpace)
		expr.setService(uastMssqlFunctionJsonExtractValue)
	}
	return nil
}
func mssqlFunctionJsonRemove(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.getJson()
	last := len(json)
	groups := make([]*exprJson, last)
	for i, group := range json {
		path := string(uastCompositeDollarPoint)
		for j, expression := range group.expressions {
			if j > 0 {
				path += string(uastCompositeSinglePoint)
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
		groups[i] = &exprJson{
			expressions: []ExpressionBase{
				&exprLiteral[string]{value: path},
			},
			operator: uastCompositeCommaSpace,
			values: []ExpressionBase{
				ConstStringNull(),
			},
		}
	}
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
	expr.setService(uastMssqlFunctionJsonRemove)
	return nil
}
func mssqlFunctionJsonSet(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.getJson()
	last := len(json)
	groups := make([]*exprJson, last)
	for i, group := range json {
		path := string(uastCompositeDollarPoint)
		for j, expression := range group.expressions {
			if j > 0 {
				path += string(uastCompositeSinglePoint)
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
		groups[i] = &exprJson{
			expressions: []ExpressionBase{
				&exprLiteral[string]{value: path},
			},
			operator: uastCompositeCommaSpace,
			values:   group.values,
		}
	}
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
	expr.setService(uastMssqlFunctionJsonSet)
	return nil
}
func mssqlFunctionCeil(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastMssqlFunctionCeil)
	return nil
}
func mssqlFunctionTrunc(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[int, int, int])
	if !exists {
		return ErrUntransformFunction
	}
	param := function.right
	function.right = &exprComposite[int]{
		expressions: []ExpressionSafe[int]{
			param,
			&exprConstant[int]{value: 1},
		},
		operator: uastCompositeCommaSpace,
	}
	expr.setService(uastMssqlFunctionTrunc)
	return nil
}
func mssqlFunctionLength(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastMssqlFunctionLength)
	return nil
}
func mssqlFunctionPosition(baseTransformer *baseTransformer, expr transformFunction) error {
	if composite, ok := expr.getLeft().(*exprComposite[string]); ok {
		expr.setLeft(&exprComposite[string]{
			expressions: []ExpressionSafe[string]{
				composite.expressions[0],
				composite.expressions[2],
			},
			operator: uastCompositeCommaSpace,
		})
	}
	expr.setService(uastMssqlFunctionPosition)
	return nil
}

// Приватные методы
func (strateger *mssqlStrateger) renderAlter(baseRenderer *baseRenderer, stmtAlter *stmtAlter) error {
	// !!!Внимание, находится в стадии разработки
	if err := baseRenderer.renderCommand(stmtAlter.command); err != nil {
		return err
	}
	if err := baseRenderer.renderEntity(stmtAlter.entity, false, false, stmtAlter.ifExists, stmtAlter.ifNotExists); err != nil {
		return err
	}
	switch stmtAlter.entity.(type) {
	case *sourceIndex:
	case *sourceSchema:
	case *sourceTable:
		//if err := baseRenderer.renderColumns(stmtAlter.columns, stmtAlter.constraints); err != nil {
		//	return err
		//}
	case *sourceView:
	}
	return nil
}
func (strateger *mssqlStrateger) renderComment(baseRenderer *baseRenderer, stmtComment *stmtComment) error {
	// !!!Внимание, находится в стадии разработки
	return nil
}
func (strateger *mssqlStrateger) renderCreate(baseRenderer *baseRenderer, stmtCreate *stmtCreate) error {
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
		if err := baseRenderer.renderColumns(stmtCreate.columns, stmtCreate.constraints); err != nil {
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
func (strateger *mssqlStrateger) renderDelete(baseRenderer *baseRenderer, stmtDelete *stmtDelete) error {
	if err := baseRenderer.renderWith(stmtDelete.with); err != nil {
		return err
	}
	if err := baseRenderer.renderCommand(stmtDelete.command); err != nil {
		return err
	}
	if err := baseRenderer.renderTarget(stmtDelete.from); err != nil {
		return err
	}
	if err := baseRenderer.renderFrom(stmtDelete.from); err != nil {
		return err
	}
	if err := baseRenderer.renderReturning(stmtDelete.returning); err != nil {
		return err
	}
	if err := baseRenderer.renderJoin(stmtDelete.join); err != nil {
		return err
	}
	if err := baseRenderer.renderWhere(stmtDelete.where); err != nil {
		return err
	}
	return nil
}
func (strateger *mssqlStrateger) renderDrop(baseRenderer *baseRenderer, stmtDrop *stmtDrop) error {
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
func (strateger *mssqlStrateger) renderInsert(baseRenderer *baseRenderer, stmtInsert *stmtInsert) error {
	if err := baseRenderer.renderWith(stmtInsert.with); err != nil {
		return err
	}
	if err := baseRenderer.renderCommand(stmtInsert.command); err != nil {
		return err
	}
	if err := baseRenderer.renderInto(stmtInsert.into); err != nil {
		return err
	}
	if err := baseRenderer.renderFields(stmtInsert.fields, true); err != nil {
		return err
	}
	if err := baseRenderer.renderReturning(stmtInsert.returning); err != nil {
		return err
	}
	if err := baseRenderer.renderSource(stmtInsert.source); err != nil {
		return err
	}
	if err := baseRenderer.renderValues(stmtInsert.values); err != nil {
		return err
	}
	return nil
}
func (strateger *mssqlStrateger) renderSelect(baseRenderer *baseRenderer, stmtSelect *stmtSelect) error {
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
func (strateger *mssqlStrateger) renderTruncate(baseRenderer *baseRenderer, stmtTruncate *stmtTruncate) error {
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
func (strateger *mssqlStrateger) renderUpdate(baseRenderer *baseRenderer, stmtUpdate *stmtUpdate) error {
	if err := baseRenderer.renderWith(stmtUpdate.with); err != nil {
		return err
	}
	if err := baseRenderer.renderCommand(stmtUpdate.command); err != nil {
		return err
	}
	if err := baseRenderer.renderOnto(stmtUpdate.onto); err != nil {
		return err
	}
	if err := baseRenderer.renderReturning(stmtUpdate.returning); err != nil {
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
	return nil
}
func (strateger *mssqlStrateger) transformAlter(baseTransformer *baseTransformer, stmtAlter *stmtAlter) error {
	// !!!Внимание, находится в стадии разработки
	return nil
}
func (strateger *mssqlStrateger) transformComment(baseTransformer *baseTransformer, stmtComment *stmtComment) error {
	// !!!Внимание, находится в стадии разработки
	return nil
}
func (strateger *mssqlStrateger) transformCreate(baseTransformer *baseTransformer, stmtCreate *stmtCreate) error {
	return nil
}
func (strateger *mssqlStrateger) transformDelete(baseTransformer *baseTransformer, stmtDelete *stmtDelete) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	if stmtDelete.returning != nil {
		stmtDelete.returning.serviceReturning = uastManagementOutput
	}
	return nil
}
func (strateger *mssqlStrateger) transformDrop(baseTransformer *baseTransformer, stmtDrop *stmtDrop) error {
	return nil
}
func (strateger *mssqlStrateger) transformInsert(baseTransformer *baseTransformer, stmtInsert *stmtInsert) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	if stmtInsert.values != nil && stmtInsert.values.upsert != nil {
		if service, exists := listManagementMssql[uastManagementUpsert]; exists {
			stmtInsert.values.upsert.service = service
		}
	}
	if stmtInsert.returning != nil {
		stmtInsert.returning.serviceReturning = uastManagementOutput
	}
	return nil
}
func (strateger *mssqlStrateger) transformSelect(baseTransformer *baseTransformer, stmtSelect *stmtSelect) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	if stmtSelect.pagination != nil {
		if len(stmtSelect.orderBy) == 0 {
			stmtSelect.orderBy = []markOrderable{
				&clauseOrderBy{
					direction:  false,
					expression: &exprLiteral[int]{value: 1},
				},
			}
		}
		stmtSelect.pagination.reverse = true
		stmtSelect.pagination.serviceLimit = uastManagementFetchNext
		stmtSelect.pagination.serviceOffset = uastManagementOffset
		stmtSelect.pagination.suffixLimit = uastModifierRowsOnly
		stmtSelect.pagination.suffixOffset = uastModifierRows
	}
	return nil
}
func (strateger *mssqlStrateger) transformTruncate(baseTransformer *baseTransformer, stmtTruncate *stmtTruncate) error {
	return nil
}
func (strateger *mssqlStrateger) transformUpdate(baseTransformer *baseTransformer, stmtUpdate *stmtUpdate) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	if stmtUpdate.returning != nil {
		stmtUpdate.returning.serviceReturning = uastManagementOutput
	}
	return nil
}
func (strateger *mssqlStrateger) validateAlter(baseValidator *baseValidator, stmtAlter *stmtAlter) error {
	// !!!Внимание, находится в стадии разработки
	if err := baseValidator.validateEntity(stmtAlter.entity); err != nil {
		return err
	}
	return nil
}
func (strateger *mssqlStrateger) validateComment(baseValidator *baseValidator, stmtComment *stmtComment) error {
	// !!!Внимание, находится в стадии разработки
	return ErrUnsupportStatement
}
func (strateger *mssqlStrateger) validateCreate(baseValidator *baseValidator, stmtCreate *stmtCreate) error {
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
func (strateger *mssqlStrateger) validateDelete(baseValidator *baseValidator, stmtDelete *stmtDelete) error {
	if err := baseValidator.validateWith(stmtDelete.with); err != nil {
		return err
	}
	if err := baseValidator.validateFrom(stmtDelete.from); err != nil {
		return err
	}
	if err := baseValidator.validateReturning(stmtDelete.returning); err != nil {
		return err
	}
	if err := baseValidator.validateJoin(stmtDelete.join); err != nil {
		return err
	}
	if err := baseValidator.validateWhere(stmtDelete.where); err != nil {
		return err
	}
	return nil
}
func (strateger *mssqlStrateger) validateDrop(baseValidator *baseValidator, stmtDrop *stmtDrop) error {
	if err := baseValidator.validateEntity(stmtDrop.entity); err != nil {
		return err
	}
	return nil
}
func (strateger *mssqlStrateger) validateInsert(baseValidator *baseValidator, stmtInsert *stmtInsert) error {
	if (stmtInsert.source == nil && stmtInsert.values == nil) || (stmtInsert.source != nil && stmtInsert.values != nil) {
		return ErrInvalidStatement
	}
	if err := baseValidator.validateWith(stmtInsert.with); err != nil {
		return err
	}
	if err := baseValidator.validateInto(stmtInsert.into); err != nil {
		return err
	}
	if err := baseValidator.validateFields(stmtInsert.fields); err != nil {
		return err
	}
	if err := baseValidator.validateReturning(stmtInsert.returning); err != nil {
		return err
	}
	if err := baseValidator.validateSource(stmtInsert.source); err != nil {
		return err
	}
	if err := baseValidator.validateValues(stmtInsert.values); err != nil {
		return err
	}
	return nil
}
func (strateger *mssqlStrateger) validateSelect(baseValidator *baseValidator, stmtSelect *stmtSelect) error {
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
func (strateger *mssqlStrateger) validateTruncate(baseValidator *baseValidator, stmtTruncate *stmtTruncate) error {
	if err := baseValidator.validateTable(stmtTruncate.table); err != nil {
		return err
	}
	return nil
}
func (strateger *mssqlStrateger) validateUpdate(baseValidator *baseValidator, stmtUpdate *stmtUpdate) error {
	if err := baseValidator.validateWith(stmtUpdate.with); err != nil {
		return err
	}
	if err := baseValidator.validateOnto(stmtUpdate.onto); err != nil {
		return err
	}
	if err := baseValidator.validateReturning(stmtUpdate.returning); err != nil {
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
	return nil
}
