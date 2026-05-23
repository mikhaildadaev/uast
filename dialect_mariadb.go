package uast

import (
	"strconv"
	"time"
)

// Публичные переменные
var DialectMariaDB = &SupportDialect{
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
		listComparisons:      listMariaDBComparisons,
		listFunctions:        listMariaDBFunctions,
		parensFunction:       true,
		placeholderNumber:    -1,
		placeholderStyle:     "?",
		placeholderType:      false,
		symbolMarkLeft:       "'",
		symbolMarkRight:      "'",
		symbolQuoteLeft:      "`",
		symbolQuoteRight:     "`",
	},
	name:      "MariaDB",
	strateger: &mariaDBStrateger{},
}

// Приватные константы
const (
	// Функции агрегатные
	uastMariaDBFunctionGroupConcat functionService = "GROUP_CONCAT"
	// Функции условий
	uastMariaDBFunctionCase functionService = "CASE"
	// Функции конвертации
	uastMariaDBFunctionCast functionService = "CAST"
	// Функции даты и времени
	uastMariaDBFunctionDateAdd functionService = "DATE_ADD"
	uastMariaDBFunctionDateSub functionService = "DATE_SUB"
	uastMariaDBFunctionTimeAdd functionService = "TIME_ADD"
	uastMariaDBFunctionTimeSub functionService = "TIME_SUB"
	// Функции обмена данными
	uastMariaDBFunctionJsonExtract     functionService = ""
	uastMariaDBFunctionJsonExtractCast functionService = "CAST"
	uastMariaDBFunctionJsonRemove      functionService = "JSON_REMOVE"
	uastMariaDBFunctionJsonSet         functionService = "JSON_SET"
	// Функции математические
	uastMariaDBFunctionCeil  functionService = "CEILING"
	uastMariaDBFunctionTrunc functionService = "TRUNCATE"
	// Функции строковые
)
const (
	uastMariaDBManagementUpsert managementService = "ON DUPLICATE KEY UPDATE"
)
const (
	// Типы бинарные
	uastMariaDBTypeBinary    typeService = "BINARY"
	uastMariaDBTypeVarBinary typeService = "VARBINARY"
	// Типы даты и времени
	uastMariaDBTypeDate      typeService = "DATE"
	uastMariaDBTypeDateTime  typeService = "DATETIME"
	uastMariaDBTypeTime      typeService = "TIME"
	uastMariaDBTypeTimestamp typeService = "DATETIME"
	// Типы числовые
	uastMariaDBTypeBigInt   typeService = "SIGNED"
	uastMariaDBTypeDecimal  typeService = "DECIMAL"
	uastMariaDBTypeDouble   typeService = "DECIMAL"
	uastMariaDBTypeFloat    typeService = "DECIMAL"
	uastMariaDBTypeInt      typeService = "SIGNED"
	uastMariaDBTypeSmallInt typeService = "SIGNED"
	// Типы строковые
	uastMariaDBTypeChar    typeService = "CHAR"
	uastMariaDBTypeString  typeService = "VARCHAR"
	uastMariaDBTypeText    typeService = "TEXT"
	uastMariaDBTypeVarChar typeService = "VARCHAR"
	// Типы специальные
	uastMariaDBTypeArray   typeService = "JSON"
	uastMariaDBTypeBoolean typeService = "TINYINT(1)"
	uastMariaDBTypeJson    typeService = "JSON"
	uastMariaDBTypeUUID    typeService = "UUID"
	uastMariaDBTypeXML     typeService = "TEXT"
)

// Приватные переменные
var listMariaDBComparisons = map[comparisonOperator]comparisonTransform{
	uastComparisonILike:    MariaDBComparisonILike,
	uastComparisonNotILike: MariaDBComparisonNotILike,
}
var listMariaDBFunctions = map[functionService]functionTransform{
	// Функции агрегатные
	uastFunctionGroupConcat: MariaDBFunctionGroupConcat,
	// Функции условий
	uastFunctionCase: MariaDBFunctionCase,
	// Функции конвертации
	uastFunctionCast: MariaDBFunctionCast,
	// Функции даты и времени
	uastFunctionDateAdd: MariaDBFunctionDateAdd,
	uastFunctionDateSub: MariaDBFunctionDateSub,
	uastFunctionTimeAdd: MariaDBFunctionTimeAdd,
	uastFunctionTimeSub: MariaDBFunctionTimeSub,
	// Функции обмена данными
	uastFunctionJsonExtract: MariaDBFunctionJsonExtract,
	uastFunctionJsonRemove:  MariaDBFunctionJsonRemove,
	uastFunctionJsonSet:     MariaDBFunctionJsonSet,
	// Функции математические
	uastFunctionCeil:  MariaDBFunctionCeil,
	uastFunctionTrunc: MariaDBFunctionTrunc,
	// Функции строковые
}
var listMariaDBType = map[ValueType]typeService{
	// Типы бинарные
	TypeBinary:    uastMariaDBTypeBinary,
	TypeVarBinary: uastMariaDBTypeVarBinary,
	// Типы даты и времени
	TypeDate:      uastMariaDBTypeDate,
	TypeDateTime:  uastMariaDBTypeDateTime,
	TypeTime:      uastMariaDBTypeTime,
	TypeTimestamp: uastMariaDBTypeTimestamp,
	// Типы числовые
	TypeBigInt:   uastMariaDBTypeBigInt,
	TypeDecimal:  uastMariaDBTypeDecimal,
	TypeDouble:   uastMariaDBTypeDouble,
	TypeFloat:    uastMariaDBTypeFloat,
	TypeInt:      uastMariaDBTypeInt,
	TypeSmallInt: uastMariaDBTypeSmallInt,
	// Типы строковые
	TypeChar:    uastMariaDBTypeChar,
	TypeString:  uastMariaDBTypeString,
	TypeText:    uastMariaDBTypeText,
	TypeVarChar: uastMariaDBTypeVarChar,
	// Типы специальные
	TypeArray:   uastMariaDBTypeArray,
	TypeBoolean: uastMariaDBTypeBoolean,
	TypeJSON:    uastMariaDBTypeJson,
	TypeUUID:    uastMariaDBTypeUUID,
	TypeXML:     uastMariaDBTypeXML,
}

// Приватные структуры
type mariaDBStrateger struct{}

// Приватные функции
func MariaDBComparisonILike(baseTransformer *baseTransformer, expr transformComparison) error {
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
func MariaDBComparisonNotILike(baseTransformer *baseTransformer, expr transformComparison) error {
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
func MariaDBFunctionGroupConcat(baseTransformer *baseTransformer, expr transformFunction) error {
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
	function.service = uastMariaDBFunctionGroupConcat
	return nil
}
func MariaDBFunctionCase(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastMariaDBFunctionCase)
	return nil
}
func MariaDBFunctionCast(baseTransformer *baseTransformer, expr transformFunction) error {
	valueType := expr.transformGetValueType()
	typeService, exists := listMariaDBType[valueType]
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
	expr.transformSetService(uastMariaDBFunctionCast)
	return nil
}
func MariaDBFunctionDateAdd(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := function.right
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastModifierInterval),
			param,
		},
		operator: uastCompositeSingleSpace,
	}
	function.service = uastMariaDBFunctionDateAdd
	return nil
}
func MariaDBFunctionDateSub(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := function.right
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastModifierInterval),
			param,
		},
		operator: uastCompositeSingleSpace,
	}
	function.service = uastMariaDBFunctionDateSub
	return nil
}
func MariaDBFunctionTimeAdd(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := function.right
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastModifierInterval),
			param,
		},
		operator: uastCompositeSingleSpace,
	}
	function.service = uastMariaDBFunctionTimeAdd
	return nil
}
func MariaDBFunctionTimeSub(baseTransformer *baseTransformer, expr transformFunction) error {
	function, exists := expr.(*exprFunction[time.Time, string, time.Time])
	if !exists {
		return ErrUntransformFunction
	}
	param := function.right
	function.right = &exprComposite[string]{
		expressions: []ExpressionSafe[string]{
			serviceString(uastModifierInterval),
			param,
		},
		operator: uastCompositeSingleSpace,
	}
	function.service = uastMariaDBFunctionTimeSub
	return nil
}
func MariaDBFunctionJsonExtract(baseTransformer *baseTransformer, expr transformFunction) error {
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
	typeService, exists := listMariaDBType[valueType]
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
		expr.transformSetService(uastMariaDBFunctionJsonExtract)
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
		expr.transformSetService(uastMariaDBFunctionJsonExtract)
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
		expr.transformSetService(uastMariaDBFunctionJsonExtractCast)
	}
	return nil
}
func MariaDBFunctionJsonRemove(baseTransformer *baseTransformer, expr transformFunction) error {
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
	expr.transformSetService(uastMariaDBFunctionJsonRemove)
	return nil
}
func MariaDBFunctionJsonSet(baseTransformer *baseTransformer, expr transformFunction) error {
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
	expr.transformSetService(uastMariaDBFunctionJsonSet)
	return nil
}
func MariaDBFunctionCeil(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastMariaDBFunctionCeil)
	return nil
}
func MariaDBFunctionTrunc(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastMariaDBFunctionTrunc)
	return nil
}
func MariaDBTarget(baseRenderer *baseRenderer, stmtDelete *stmtDelete) error {
	if stmtDelete.from == nil {
		return ErrInvalidStatementTarget
	}
	var targetAlias string
	switch source := stmtDelete.from.(type) {
	case *CteSource:
		targetAlias = source.aliasName
	case *TableSource:
		targetAlias = source.aliasName
	case *QuerySource:
		targetAlias = source.aliasName
	default:
		return ErrInvalidAlias
	}
	baseRenderer.renderOperator(uastCompositeSingleSpace)
	baseRenderer.renderAlias(targetAlias)
	return nil
}

// Приватные методы
func (strateger *mariaDBStrateger) renderDelete(baseRenderer *baseRenderer, stmtDelete *stmtDelete) error {
	if err := baseRenderer.renderWith(stmtDelete.with); err != nil {
		return err
	}
	if err := baseRenderer.renderCommand(stmtDelete.command); err != nil {
		return err
	}
	if err := MariaDBTarget(baseRenderer, stmtDelete); err != nil {
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
func (strateger *mariaDBStrateger) renderInsert(baseRenderer *baseRenderer, stmtInsert *stmtInsert) error {
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
func (strateger *mariaDBStrateger) renderSelect(baseRenderer *baseRenderer, stmtSelect *stmtSelect) error {
	if err := baseRenderer.renderWith(stmtSelect.with); err != nil {
		return err
	}
	if err := baseRenderer.renderCommand(stmtSelect.command); err != nil {
		return err
	}
	if err := baseRenderer.renderDistinct(stmtSelect.distinct); err != nil {
		return err
	}
	if err := baseRenderer.renderField(stmtSelect.field); err != nil {
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
	if err := baseRenderer.renderLimit(stmtSelect.limit); err != nil {
		return err
	}
	if err := baseRenderer.renderOffset(stmtSelect.offset); err != nil {
		return err
	}
	if err := baseRenderer.renderUnions(stmtSelect.unions); err != nil {
		return err
	}
	return nil
}
func (strateger *mariaDBStrateger) renderTruncate(baseRenderer *baseRenderer, stmtTruncate *stmtTruncate) error {
	// !!! Внимание, находится в стадии разработки
	if err := baseRenderer.renderCommand(stmtTruncate.command); err != nil {
		return err
	}
	if err := baseRenderer.renderTable(stmtTruncate.table); err != nil {
		return err
	}
	//if err := baseRenderer.renderIdentity(stmtTruncate.identity); err != nil {
	//	return err
	//}
	//if err := baseRenderer.renderCascade(stmtTruncate.cascade); err != nil {
	//	return err
	//}
	return nil
}
func (strateger *mariaDBStrateger) renderUpdate(baseRenderer *baseRenderer, stmtUpdate *stmtUpdate) error {
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
func (strateger *mariaDBStrateger) transformDelete(baseTransformer *baseTransformer, stmtDelete *stmtDelete) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *mariaDBStrateger) transformInsert(baseTransformer *baseTransformer, stmtInsert *stmtInsert) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	if stmtInsert.values != nil && stmtInsert.values.upsert != nil {
		stmtInsert.values.upsert.service = uastMariaDBManagementUpsert
	}
	return nil
}
func (strateger *mariaDBStrateger) transformSelect(baseTransformer *baseTransformer, stmtSelect *stmtSelect) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *mariaDBStrateger) transformTruncate(baseTransformer *baseTransformer, stmtTruncate *stmtTruncate) error {
	// !!! Внимание, находится в стадии разработки
	return nil
}
func (strateger *mariaDBStrateger) transformUpdate(baseTransformer *baseTransformer, stmtUpdate *stmtUpdate) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *mariaDBStrateger) validateDelete(baseValidator *baseValidator, stmtDelete *stmtDelete) error {
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
func (strateger *mariaDBStrateger) validateInsert(baseValidator *baseValidator, stmtInsert *stmtInsert) error {
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
func (strateger *mariaDBStrateger) validateSelect(baseValidator *baseValidator, stmtSelect *stmtSelect) error {
	if err := baseValidator.validateWith(stmtSelect.with); err != nil {
		return err
	}
	if err := baseValidator.validateField(stmtSelect.field); err != nil {
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
	if err := baseValidator.validateLimit(stmtSelect.limit); err != nil {
		return err
	}
	if err := baseValidator.validateOffset(stmtSelect.offset); err != nil {
		return err
	}
	if err := baseValidator.validateUnions(stmtSelect.unions); err != nil {
		return err
	}
	return nil
}
func (strateger *mariaDBStrateger) validateTruncate(baseValidator *baseValidator, stmtTruncate *stmtTruncate) error {
	// !!! Внимание, находится в стадии разработки
	if err := baseValidator.validateTable(stmtTruncate.table); err != nil {
		return err
	}
	return nil
}
func (strateger *mariaDBStrateger) validateUpdate(baseValidator *baseValidator, stmtUpdate *stmtUpdate) error {
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
