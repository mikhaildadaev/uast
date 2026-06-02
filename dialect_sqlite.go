package uast

import (
	"strconv"
	"time"
)

// Публичные переменные
var DialectSQLite = &SupportDialect{
	config: &config{
		lengthMaxArray:       128,
		lengthMaxConst:       63,
		lengthMaxFunc:        48,
		lengthMaxIdent:       64,
		lengthMaxLimit:       63,
		lengthMaxParam:       65535,
		lengthMaxQuery:       64 * 1024,
		lengthMaxValueByte:   1024,
		lengthMaxValueString: 128,
		listComparisons:      listComparisonsSQLite,
		listFunctions:        listFunctionsSQLite,
		parensFunction:       false,
		placeholderNumber:    0,
		placeholderStyle:     "?",
		placeholderType:      false,
		symbolMarkLeft:       "'",
		symbolMarkRight:      "'",
		symbolQuoteLeft:      "\"",
		symbolQuoteRight:     "\"",
		supportCascade:       false,
		supportIfExists: map[string]bool{
			"INDEX":  true,
			"SCHEMA": false,
			"TABLE":  true,
			"VIEW":   true,
		},
		supportIfNotExists: map[string]bool{
			"INDEX":  true,
			"SCHEMA": false,
			"TABLE":  true,
			"VIEW":   true,
		},
		supportRestartIdentity: false,
		supportReturning:       true,
		supportUpsert:          true,
	},
	name:      "SQLite",
	strateger: &sqliteStrateger{},
}

// Приватные константы
const (
	// Функции агрегатные
	uastSQLiteFunctionGroupConcat functionService = "GROUP_CONCAT"
	uastSQLiteFunctionStdDev      functionService = "STDEV"
	// Функции условий
	uastSQLiteFunctionCast functionService = "CAST"
	// Функции конвертации
	uastSQLiteFunctionDateFormat functionService = "STRFTIME"
	// Функции даты и времени
	uastSQLiteFunctionCurDate   functionService = "DATE"
	uastSQLiteFunctionCurTime   functionService = "TIME"
	uastSQLiteFunctionDateAdd   functionService = "DATETIME"
	uastSQLiteFunctionDateSub   functionService = "DATETIME"
	uastSQLiteFunctionDayName   functionService = "STRFTIME"
	uastSQLiteFunctionNow       functionService = "DATETIME"
	uastSQLiteFunctionMonthName functionService = "STRFTIME"
	uastSQLiteFunctionTimeAdd   functionService = "TIME"
	uastSQLiteFunctionTimeSub   functionService = "TIME"
	// Функции обмена данными
	uastSQLiteFunctionJsonArrayAgg    functionService = "JSON_GROUP_ARRAY"
	uastSQLiteFunctionJsonExtract     functionService = ""
	uastSQLiteFunctionJsonExtractCast functionService = "CAST"
	uastSQLiteFunctionJsonObject      functionService = "JSON_OBJECT"
	uastSQLiteFunctionJsonObjectAgg   functionService = "JSON_GROUP_OBJECT"
	uastSQLiteFunctionJsonRemove      functionService = "JSON_REMOVE"
	uastSQLiteFunctionJsonSet         functionService = "JSON_SET"
	// Функции математические
	uastSQLiteFunctionRand functionService = "RANDOM"
	// Функции строковые
)
const (
	uastSQLiteManagementUpsert managementService = "ON CONFLICT DO UPDATE SET"
)
const (
	// Типы бинарные
	uastSQLiteTypeBinary    typeService = "BLOB"
	uastSQLiteTypeVarBinary typeService = "BLOB"
	// Типы даты и времени
	uastSQLiteTypeDate      typeService = "TEXT"
	uastSQLiteTypeDateTime  typeService = "TEXT"
	uastSQLiteTypeTime      typeService = "TEXT"
	uastSQLiteTypeTimestamp typeService = "TEXT"
	// Типы числовые
	uastSQLiteTypeBigInt   typeService = "INTEGER"
	uastSQLiteTypeDecimal  typeService = "REAL"
	uastSQLiteTypeDouble   typeService = "REAL"
	uastSQLiteTypeFloat    typeService = "REAL"
	uastSQLiteTypeInt      typeService = "INTEGER"
	uastSQLiteTypeSmallInt typeService = "INTEGER"
	// Типы строковые
	uastSQLiteTypeChar    typeService = "TEXT"
	uastSQLiteTypeString  typeService = "TEXT"
	uastSQLiteTypeText    typeService = "TEXT"
	uastSQLiteTypeVarChar typeService = "TEXT"
	// Типы специальные
	uastSQLiteTypeArray   typeService = "TEXT"
	uastSQLiteTypeBoolean typeService = "INTEGER"
	uastSQLiteTypeJson    typeService = "TEXT"
	uastSQLiteTypeUUID    typeService = "TEXT"
	uastSQLiteTypeXML     typeService = "TEXT"
)

// Приватные переменные
var listComparisonsSQLite = map[comparisonOperator]comparisonTransform{
	uastComparisonILike:    sqliteComparisonILike,
	uastComparisonNotILike: sqliteComparisonNotILike,
}
var listFunctionsSQLite = map[functionService]functionTransform{
	// Функции агрегатные
	uastFunctionGroupConcat: sqliteFunctionGroupConcat,
	uastFunctionStdDev:      sqliteFunctionStdDev,
	// Функции условий
	uastFunctionCast: sqliteFunctionCast,
	// Функции конвертации
	uastFunctionDateFormat: sqliteFunctionDateFormat,
	// Функции даты и времени
	uastFunctionCurDate:   sqliteFunctionCurDate,
	uastFunctionCurTime:   sqliteFunctionCurTime,
	uastFunctionDateAdd:   sqliteFunctionDateAdd,
	uastFunctionDateSub:   sqliteFunctionDateSub,
	uastFunctionDayName:   sqliteFunctionDayName,
	uastFunctionNow:       sqliteFunctionNow,
	uastFunctionMonthName: sqliteFunctionMonthName,
	uastFunctionTimeAdd:   sqliteFunctionTimeAdd,
	uastFunctionTimeSub:   sqliteFunctionTimeSub,
	// Функции обмена данными
	uastFunctionJsonArrayAgg:  sqliteFunctionJsonArrayAgg,
	uastFunctionJsonExtract:   sqliteFunctionJsonExtract,
	uastFunctionJsonObject:    sqliteFunctionJsonObject,
	uastFunctionJsonObjectAgg: sqliteFunctionJsonObjectAgg,
	uastFunctionJsonRemove:    sqliteFunctionJsonRemove,
	uastFunctionJsonSet:       sqliteFunctionJsonSet,
	// Функции математические
	uastFunctionRand: sqliteFunctionRand,
	// Функции строковые
}
var listTypeSQLite = map[ValueType]typeService{
	// Типы бинарные
	TypeBinary:    uastSQLiteTypeBinary,
	TypeVarBinary: uastSQLiteTypeVarBinary,
	// Типы даты и времени
	TypeDate:      uastSQLiteTypeDate,
	TypeDateTime:  uastSQLiteTypeDateTime,
	TypeTime:      uastSQLiteTypeTime,
	TypeTimestamp: uastSQLiteTypeTimestamp,
	// Типы числовые
	TypeBigInt:   uastSQLiteTypeBigInt,
	TypeDecimal:  uastSQLiteTypeDecimal,
	TypeDouble:   uastSQLiteTypeDouble,
	TypeFloat:    uastSQLiteTypeFloat,
	TypeInt:      uastSQLiteTypeInt,
	TypeSmallInt: uastSQLiteTypeSmallInt,
	// Типы строковые
	TypeChar:    uastSQLiteTypeChar,
	TypeString:  uastSQLiteTypeString,
	TypeText:    uastSQLiteTypeText,
	TypeVarChar: uastSQLiteTypeVarChar,
	// Типы специальные
	TypeArray:   uastSQLiteTypeArray,
	TypeBoolean: uastSQLiteTypeBoolean,
	TypeJSON:    uastSQLiteTypeJson,
	TypeUUID:    uastSQLiteTypeUUID,
	TypeXML:     uastSQLiteTypeXML,
}

// Приватные структуры
type sqliteStrateger struct{}

// Приватные функции
func sqliteComparisonILike(baseTransformer *baseTransformer, expr transformComparison) error {
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
func sqliteComparisonNotILike(baseTransformer *baseTransformer, expr transformComparison) error {
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
func sqliteFunctionGroupConcat(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[string, string, string])
	if !exists {
		return ErrUntransformFunction
	}
	if !exists {
		return ErrUntransformParam
	}
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastModifierSeparator),
			&exprLiteral[string]{value: ","},
		},
		operator: uastCompositeSingleSpace,
	}
	function.service = uastSQLiteFunctionGroupConcat
	return nil
}
func sqliteFunctionStdDev(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastSQLiteFunctionStdDev)
	return nil
}
func sqliteFunctionCast(baseTransformer *baseTransformer, expr transformFunction) error {
	valueType := expr.transformGetValueType()
	typeService, exists := listTypeSQLite[valueType]
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
	expr.transformSetService(uastSQLiteFunctionCast)
	return nil
}
func sqliteFunctionDateFormat(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastSQLiteFunctionDateFormat)
	return nil
}
func sqliteFunctionCurDate(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastSQLiteFunctionCurDate)
	expr.transformSetLeft(&exprLiteral[string]{
		value: "now",
	})
	expr.transformSetProcess(uastProcessDirect)
	return nil
}
func sqliteFunctionCurTime(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastSQLiteFunctionCurTime)
	expr.transformSetLeft(&exprLiteral[string]{
		value: "now",
	})
	expr.transformSetProcess(uastProcessDirect)
	return nil
}
func sqliteFunctionDateAdd(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := string(uastBinaryPlus) + function.right.(*exprLiteral[string]).value
	function.right = &exprLiteral[string]{
		value: param,
	}
	function.service = uastSQLiteFunctionDateAdd
	return nil
}
func sqliteFunctionDateSub(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := string(uastBinaryMinus) + function.right.(*exprLiteral[string]).value
	function.right = &exprLiteral[string]{
		value: param,
	}
	function.service = uastSQLiteFunctionDateSub
	return nil
}
func sqliteFunctionDayName(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetOperator(uastCompositeCommaSpace)
	expr.transformSetProcess(uastProcessInvert)
	expr.transformSetRight(&exprLiteral[string]{
		value: "%w",
	})
	expr.transformSetService(uastSQLiteFunctionDayName)
	return nil
}
func sqliteFunctionNow(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastSQLiteFunctionNow)
	expr.transformSetLeft(&exprLiteral[string]{
		value: "now",
	})
	expr.transformSetProcess(uastProcessDirect)
	return nil
}
func sqliteFunctionMonthName(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetOperator(uastCompositeCommaSpace)
	expr.transformSetProcess(uastProcessInvert)
	expr.transformSetRight(&exprLiteral[string]{
		value: "%m",
	})
	expr.transformSetService(uastSQLiteFunctionMonthName)
	return nil
}
func sqliteFunctionTimeAdd(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := string(uastBinaryPlus) + function.right.(*exprLiteral[string]).value
	function.right = &exprLiteral[string]{
		value: param,
	}
	function.service = uastSQLiteFunctionTimeAdd
	return nil
}
func sqliteFunctionTimeSub(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := string(uastBinaryMinus) + function.right.(*exprLiteral[string]).value
	function.right = &exprLiteral[string]{
		value: param,
	}
	function.service = uastSQLiteFunctionTimeSub
	return nil
}
func sqliteFunctionJsonArrayAgg(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastSQLiteFunctionJsonArrayAgg)
	return nil
}
func sqliteFunctionJsonExtract(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.transformGetJson()
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
	valueType := expr.transformGetValueType()
	typeService, exists := listTypeSQLite[valueType]
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
		expr.transformSetOperator(uastCompositeSpaceMinusGreaterSpace)
		expr.transformSetService(uastSQLiteFunctionJsonExtract)
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
		expr.transformSetOperator(uastCompositeSpaceMinusDoubleGreaterSpace)
		expr.transformSetService(uastSQLiteFunctionJsonExtract)
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
		expr.transformSetOperator(uastCompositeSpaceMinusDoubleGreaterSpace)
		expr.transformSetService(uastSQLiteFunctionJsonExtractCast)
	}
	return nil
}
func sqliteFunctionJsonObject(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastSQLiteFunctionJsonObject)
	return nil
}
func sqliteFunctionJsonObjectAgg(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastSQLiteFunctionJsonObjectAgg)
	return nil
}
func sqliteFunctionJsonRemove(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.transformGetJson()
	groups := make([]*exprJson, len(json))
	for i, group := range json {
		path := string(uastCompositeDollarPoint)
		for j, expression := range group.expressions {
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
		groups[i] = &exprJson{
			expressions: []ExpressionBase{
				&exprLiteral[string]{
					value: path,
				},
			},
			operator: uastCompositeCommaSpace,
		}
	}
	expr.transformSetJson(groups)
	expr.transformSetOperator(uastCompositeCommaSpace)
	expr.transformSetService(uastSQLiteFunctionJsonRemove)
	return nil
}
func sqliteFunctionJsonSet(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.transformGetJson()
	groups := make([]*exprJson, len(json))
	for i, group := range json {
		path := string(uastCompositeDollarPoint)
		for j, expression := range group.expressions {
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
		groups[i] = &exprJson{
			expressions: []ExpressionBase{
				&exprLiteral[string]{
					value: path,
				},
			},
			operator: uastCompositeCommaSpace,
			values:   group.values,
		}
	}
	expr.transformSetJson(groups)
	expr.transformSetOperator(uastCompositeCommaSpace)
	expr.transformSetService(uastSQLiteFunctionJsonSet)
	return nil
}
func sqliteFunctionRand(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastSQLiteFunctionRand)
	expr.transformSetProcess(uastProcessEmpty)
	return nil
}

// Приватные методы
func (strateger *sqliteStrateger) renderComment(baseRenderer *baseRenderer, stmtComment *stmtComment) error {
	if err := baseRenderer.renderCommand(stmtComment.command); err != nil {
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
func (strateger *sqliteStrateger) renderCreate(baseRenderer *baseRenderer, stmtCreate *stmtCreate) error {
	// !!! Внимание, находится в стадии разработки
	if err := baseRenderer.renderCommand(stmtCreate.command); err != nil {
		return err
	}
	switch stmtCreate.entity.(type) {
	case *sourceIndex:
		if err := baseRenderer.renderUnique(stmtCreate.unique); err != nil {
			return err
		}
		if err := baseRenderer.renderEntity(stmtCreate.entity, false, stmtCreate.ifNotExists); err != nil {
			return err
		}
		if err := baseRenderer.renderOn(stmtCreate.table); err != nil {
			return err
		}
	case *sourceSchema:
	case *sourceTable:
		if err := baseRenderer.renderEntity(stmtCreate.entity, false, stmtCreate.ifNotExists); err != nil {
			return err
		}
		if err := baseRenderer.renderColumns(stmtCreate.columns); err != nil {
			return err
		}
	case *sourceView:
		if err := baseRenderer.renderReplace(stmtCreate.replace); err != nil {
			return err
		}
		if err := baseRenderer.renderEntity(stmtCreate.entity, false, stmtCreate.ifNotExists); err != nil {
			return err
		}
		if err := baseRenderer.renderAs(); err != nil {
			return err
		}
		if err := baseRenderer.renderSource(stmtCreate.source); err != nil {
			return err
		}
	}
	return nil
}
func (strateger *sqliteStrateger) renderDelete(baseRenderer *baseRenderer, stmtDelete *stmtDelete) error {
	if err := baseRenderer.renderWith(stmtDelete.with); err != nil {
		return err
	}
	if err := baseRenderer.renderCommand(stmtDelete.command); err != nil {
		return err
	}
	if err := baseRenderer.renderFrom(stmtDelete.from); err != nil {
		return err
	}
	if err := baseRenderer.renderJoin(stmtDelete.join); err != nil {
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
func (strateger *sqliteStrateger) renderDrop(baseRenderer *baseRenderer, stmtDrop *stmtDrop) error {
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
func (strateger *sqliteStrateger) renderInsert(baseRenderer *baseRenderer, stmtInsert *stmtInsert) error {
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
func (strateger *sqliteStrateger) renderSelect(baseRenderer *baseRenderer, stmtSelect *stmtSelect) error {
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
func (strateger *sqliteStrateger) renderTruncate(baseRenderer *baseRenderer, stmtTruncate *stmtTruncate) error {
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
func (strateger *sqliteStrateger) renderUpdate(baseRenderer *baseRenderer, stmtUpdate *stmtUpdate) error {
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
func (strateger *sqliteStrateger) transformComment(baseTransformer *baseTransformer, stmtComment *stmtComment) error {
	return nil
}
func (strateger *sqliteStrateger) transformCreate(baseTransformer *baseTransformer, stmtCreate *stmtCreate) error {
	// !!! Внимание, находится в стадии разработки
	return nil
}
func (strateger *sqliteStrateger) transformDelete(baseTransformer *baseTransformer, stmtDelete *stmtDelete) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *sqliteStrateger) transformDrop(baseTransformer *baseTransformer, stmtDrop *stmtDrop) error {
	return nil
}
func (strateger *sqliteStrateger) transformInsert(baseTransformer *baseTransformer, stmtInsert *stmtInsert) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	if stmtInsert.values != nil && stmtInsert.values.upsert != nil {
		stmtInsert.values.upsert.service = uastSQLiteManagementUpsert
	}
	return nil
}
func (strateger *sqliteStrateger) transformSelect(baseTransformer *baseTransformer, stmtSelect *stmtSelect) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	if stmtSelect.join != nil {
		joins := make([]*clauseJoin, 0, len(stmtSelect.join))
		for _, join := range stmtSelect.join {
			if join.operator != uastJoinRight && join.operator != uastJoinRightOuter {
				joins = append(joins, join)
			}
		}
		stmtSelect.join = joins
	}
	return nil
}
func (strateger *sqliteStrateger) transformTruncate(baseTransformer *baseTransformer, stmtTruncate *stmtTruncate) error {
	return nil
}
func (strateger *sqliteStrateger) transformUpdate(baseTransformer *baseTransformer, stmtUpdate *stmtUpdate) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *sqliteStrateger) validateComment(baseValidator *baseValidator, stmtComment *stmtComment) error {
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
func (strateger *sqliteStrateger) validateCreate(baseValidator *baseValidator, stmtCreate *stmtCreate) error {
	// !!! Внимание, находится в стадии разработки
	switch stmtCreate.entity.(type) {
	case *sourceIndex:
		if err := baseValidator.validateEntity(stmtCreate.entity); err != nil {
			return err
		}
		if err := baseValidator.validateColumns(stmtCreate.columns); err != nil {
			return err
		}
	case *sourceSchema:
	case *sourceTable:
		if err := baseValidator.validateEntity(stmtCreate.entity); err != nil {
			return err
		}
		if err := baseValidator.validateColumns(stmtCreate.columns); err != nil {
			return err
		}
	case *sourceView:
		if err := baseValidator.validateEntity(stmtCreate.entity); err != nil {
			return err
		}
		if err := baseValidator.validateSource(stmtCreate.source); err != nil {
			return err
		}
	default:
		return ErrInvalidStatement
	}
	return nil
}
func (strateger *sqliteStrateger) validateDelete(baseValidator *baseValidator, stmtDelete *stmtDelete) error {
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
func (strateger *sqliteStrateger) validateDrop(baseValidator *baseValidator, stmtDrop *stmtDrop) error {
	if err := baseValidator.validateEntity(stmtDrop.entity); err != nil {
		return err
	}
	return nil
}
func (strateger *sqliteStrateger) validateInsert(baseValidator *baseValidator, stmtInsert *stmtInsert) error {
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
func (strateger *sqliteStrateger) validateSelect(baseValidator *baseValidator, stmtSelect *stmtSelect) error {
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
func (strateger *sqliteStrateger) validateTruncate(baseValidator *baseValidator, stmtTruncate *stmtTruncate) error {
	if err := baseValidator.validateTable(stmtTruncate.table); err != nil {
		return err
	}
	return nil
}
func (strateger *sqliteStrateger) validateUpdate(baseValidator *baseValidator, stmtUpdate *stmtUpdate) error {
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
