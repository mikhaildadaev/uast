package uast

import (
	"strconv"
	"time"
)

// Публичные переменные
var DialectMySQL = &SupportDialect{
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
		listComparisons:      listMysqlComparisons,
		listFunctions:        listMysqlFunctions,
		parensFunction:       true,
		placeholderNumber:    -1,
		placeholderStyle:     "?",
		placeholderType:      false,
		symbolMarkLeft:       "'",
		symbolMarkRight:      "'",
		symbolQuoteLeft:      "`",
		symbolQuoteRight:     "`",
	},
	name:      "MySQL",
	strateger: &mysqlStrateger{},
}

// Приватные константы
const (
	// Функции агрегатные
	uastMysqlFunctionGroupConcat functionService = "GROUP_CONCAT"
	// Функции условий
	uastMysqlFunctionCase functionService = "CASE"
	// Функции конвертации
	uastMysqlFunctionCast functionService = "CAST"
	// Функции даты и времени
	uastMysqlFunctionDateAdd functionService = "DATE_ADD"
	uastMysqlFunctionDateSub functionService = "DATE_SUB"
	uastMysqlFunctionTimeAdd functionService = "TIME_ADD"
	uastMysqlFunctionTimeSub functionService = "TIME_SUB"
	// Функции обмена данными
	uastMysqlFunctionJsonExtract     functionService = ""
	uastMysqlFunctionJsonExtractCast functionService = "CAST"
	uastMysqlFunctionJsonRemove      functionService = "JSON_REMOVE"
	uastMysqlFunctionJsonSet         functionService = "JSON_SET"
	// Функции математические
	uastMysqlFunctionCeil  functionService = "CEILING"
	uastMysqlFunctionTrunc functionService = "TRUNCATE"
	// Функции строковые
)
const (
	uastMySQLManagementUpsert managementService = "ON DUPLICATE KEY UPDATE"
)
const (
	// Типы бинарные
	uastMysqlTypeBinary    typeService = "BINARY"
	uastMysqlTypeVarBinary typeService = "VARBINARY"
	// Типы даты и времени
	uastMysqlTypeDate      typeService = "DATE"
	uastMysqlTypeDateTime  typeService = "DATETIME"
	uastMysqlTypeTime      typeService = "TIME"
	uastMysqlTypeTimestamp typeService = "DATETIME"
	// Типы числовые
	uastMysqlTypeBigInt   typeService = "SIGNED"
	uastMysqlTypeDecimal  typeService = "DECIMAL"
	uastMysqlTypeDouble   typeService = "DECIMAL"
	uastMysqlTypeFloat    typeService = "DECIMAL"
	uastMysqlTypeInt      typeService = "SIGNED"
	uastMysqlTypeSmallInt typeService = "SIGNED"
	// Типы строковые
	uastMysqlTypeChar    typeService = "CHAR"
	uastMysqlTypeString  typeService = "VARCHAR"
	uastMysqlTypeText    typeService = "TEXT"
	uastMysqlTypeVarChar typeService = "VARCHAR"
	// Типы специальные
	uastMysqlTypeArray   typeService = "JSON"
	uastMysqlTypeBoolean typeService = "TINYINT(1)"
	uastMysqlTypeJson    typeService = "JSON"
	uastMysqlTypeUUID    typeService = "CHAR(36)"
	uastMysqlTypeXML     typeService = "TEXT"
)

// Приватные переменные
var listMysqlComparisons = map[comparisonOperator]comparisonTransform{
	uastComparisonILike:    mysqlComparisonILike,
	uastComparisonNotILike: mysqlComparisonNotILike,
}
var listMysqlFunctions = map[functionService]functionTransform{
	// Функции агрегатные
	uastFunctionGroupConcat: mysqlFunctionGroupConcat,
	// Функции условий
	uastFunctionCase: mysqlFunctionCase,
	// Функции конвертации
	uastFunctionCast: mysqlFunctionCast,
	// Функции даты и времени
	uastFunctionDateAdd: mysqlFunctionDateAdd,
	uastFunctionDateSub: mysqlFunctionDateSub,
	uastFunctionTimeAdd: mysqlFunctionTimeAdd,
	uastFunctionTimeSub: mysqlFunctionTimeSub,
	// Функции обмена данными
	uastFunctionJsonExtract: mysqlFunctionJsonExtract,
	uastFunctionJsonRemove:  mysqlFunctionJsonRemove,
	uastFunctionJsonSet:     mysqlFunctionJsonSet,
	// Функции математические
	uastFunctionCeil:  mysqlFunctionCeil,
	uastFunctionTrunc: mysqlFunctionTrunc,
	// Функции строковые
}
var listMysqlType = map[ValueType]typeService{
	// Типы бинарные
	TypeBinary:    uastMysqlTypeBinary,
	TypeVarBinary: uastMysqlTypeVarBinary,
	// Типы даты и времени
	TypeDate:      uastMysqlTypeDate,
	TypeDateTime:  uastMysqlTypeDateTime,
	TypeTime:      uastMysqlTypeTime,
	TypeTimestamp: uastMysqlTypeTimestamp,
	// Типы числовые
	TypeBigInt:   uastMysqlTypeBigInt,
	TypeDecimal:  uastMysqlTypeDecimal,
	TypeDouble:   uastMysqlTypeDouble,
	TypeFloat:    uastMysqlTypeFloat,
	TypeInt:      uastMysqlTypeInt,
	TypeSmallInt: uastMysqlTypeSmallInt,
	// Типы строковые
	TypeChar:    uastMysqlTypeChar,
	TypeString:  uastMysqlTypeString,
	TypeText:    uastMysqlTypeText,
	TypeVarChar: uastMysqlTypeVarChar,
	// Типы специальные
	TypeArray:   uastMysqlTypeArray,
	TypeBoolean: uastMysqlTypeBoolean,
	TypeJSON:    uastMysqlTypeJson,
	TypeUUID:    uastMysqlTypeUUID,
	TypeXML:     uastMysqlTypeXML,
}

// Приватные структуры
type mysqlStrateger struct{}

// Приватные функции
func mysqlComparisonILike(baseTransformer *baseTransformer, expr transformComparison) error {
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
func mysqlComparisonNotILike(baseTransformer *baseTransformer, expr transformComparison) error {
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
func mysqlFunctionGroupConcat(baseTransformer *baseTransformer, expr transformFunction) error {
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
	function.service = uastMysqlFunctionGroupConcat
	return nil
}
func mysqlFunctionCase(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastMysqlFunctionCase)
	return nil
}
func mysqlFunctionCast(baseTransformer *baseTransformer, expr transformFunction) error {
	valueType := expr.transformGetValueType()
	typeService, exists := listMysqlType[valueType]
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
	expr.transformSetService(uastMysqlFunctionCast)
	return nil
}
func mysqlFunctionDateAdd(baseTransformer *baseTransformer, expr transformFunction) error {
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
	function.service = uastMysqlFunctionDateAdd
	return nil
}
func mysqlFunctionDateSub(baseTransformer *baseTransformer, expr transformFunction) error {
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
	function.service = uastMysqlFunctionDateSub
	return nil
}
func mysqlFunctionTimeAdd(baseTransformer *baseTransformer, expr transformFunction) error {
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
	function.service = uastMysqlFunctionTimeAdd
	return nil
}
func mysqlFunctionTimeSub(baseTransformer *baseTransformer, expr transformFunction) error {
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
	function.service = uastMysqlFunctionTimeSub
	return nil
}
func mysqlFunctionJsonExtract(baseTransformer *baseTransformer, expr transformFunction) error {
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
	typeService, exists := listMysqlType[valueType]
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
		expr.transformSetService(uastMysqlFunctionJsonExtract)
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
		expr.transformSetService(uastMysqlFunctionJsonExtract)
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
		expr.transformSetService(uastMysqlFunctionJsonExtractCast)
	}
	return nil
}
func mysqlFunctionJsonRemove(baseTransformer *baseTransformer, expr transformFunction) error {
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
	expr.transformSetService(uastMysqlFunctionJsonRemove)
	return nil
}
func mysqlFunctionJsonSet(baseTransformer *baseTransformer, expr transformFunction) error {
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
	expr.transformSetService(uastMysqlFunctionJsonSet)
	return nil
}
func mysqlFunctionCeil(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastMysqlFunctionCeil)
	return nil
}
func mysqlFunctionTrunc(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.transformSetService(uastMysqlFunctionTrunc)
	return nil
}
func mysqlTarget(baseRenderer *baseRenderer, stmtDelete *stmtDelete) error {
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
func (strateger *mysqlStrateger) renderComment(baseRenderer *baseRenderer, stmtComment *stmtComment) error {
	if err := baseRenderer.renderCommand(stmtComment.command); err != nil {
		return err
	}
	if err := baseRenderer.renderOnColumn(stmtComment.column, stmtComment.comment); err != nil {
		return err
	}
	if err := baseRenderer.renderOnTable(stmtComment.table, stmtComment.comment); err != nil {
		return err
	}
	return nil
}
func (strateger *mysqlStrateger) renderDelete(baseRenderer *baseRenderer, stmtDelete *stmtDelete) error {
	if err := baseRenderer.renderWith(stmtDelete.with); err != nil {
		return err
	}
	if err := baseRenderer.renderCommand(stmtDelete.command); err != nil {
		return err
	}
	if err := mysqlTarget(baseRenderer, stmtDelete); err != nil {
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
	return nil
}
func (strateger *mysqlStrateger) renderInsert(baseRenderer *baseRenderer, stmtInsert *stmtInsert) error {
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
	return nil
}
func (strateger *mysqlStrateger) renderSelect(baseRenderer *baseRenderer, stmtSelect *stmtSelect) error {
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
func (strateger *mysqlStrateger) renderTruncate(baseRenderer *baseRenderer, stmtTruncate *stmtTruncate) error {
	if err := baseRenderer.renderCommand(stmtTruncate.command); err != nil {
		return err
	}
	if err := baseRenderer.renderTable(stmtTruncate.table); err != nil {
		return err
	}
	return nil
}
func (strateger *mysqlStrateger) renderUpdate(baseRenderer *baseRenderer, stmtUpdate *stmtUpdate) error {
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
	return nil
}
func (strateger *mysqlStrateger) transformComment(baseTransformer *baseTransformer, stmtComment *stmtComment) error {
	return nil
}
func (strateger *mysqlStrateger) transformDelete(baseTransformer *baseTransformer, stmtDelete *stmtDelete) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *mysqlStrateger) transformInsert(baseTransformer *baseTransformer, stmtInsert *stmtInsert) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	if stmtInsert.values != nil && stmtInsert.values.upsert != nil {
		stmtInsert.values.upsert.service = uastMySQLManagementUpsert
	}
	return nil
}
func (strateger *mysqlStrateger) transformSelect(baseTransformer *baseTransformer, stmtSelect *stmtSelect) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *mysqlStrateger) transformTruncate(baseTransformer *baseTransformer, stmtTruncate *stmtTruncate) error {
	return nil
}
func (strateger *mysqlStrateger) transformUpdate(baseTransformer *baseTransformer, stmtUpdate *stmtUpdate) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *mysqlStrateger) validateComment(baseValidator *baseValidator, stmtComment *stmtComment) error {
	if stmtComment.column == nil && stmtComment.table == nil {
		return ErrInvalidStatement
	}
	if err := baseValidator.validateOnColumn(stmtComment.column, stmtComment.comment); err != nil {
		return err
	}
	if err := baseValidator.validateOnTable(stmtComment.table, stmtComment.comment); err != nil {
		return err
	}
	return nil
}
func (strateger *mysqlStrateger) validateDelete(baseValidator *baseValidator, stmtDelete *stmtDelete) error {
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
	return nil
}
func (strateger *mysqlStrateger) validateInsert(baseValidator *baseValidator, stmtInsert *stmtInsert) error {
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
	return nil
}
func (strateger *mysqlStrateger) validateSelect(baseValidator *baseValidator, stmtSelect *stmtSelect) error {
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
func (strateger *mysqlStrateger) validateTruncate(baseValidator *baseValidator, stmtTruncate *stmtTruncate) error {
	if err := baseValidator.validateTable(stmtTruncate.table); err != nil {
		return err
	}
	return nil
}
func (strateger *mysqlStrateger) validateUpdate(baseValidator *baseValidator, stmtUpdate *stmtUpdate) error {
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
	return nil
}
