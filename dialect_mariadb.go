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
		listComparisons:      listComparisonsMariadb,
		listFunctions:        listFunctionsMariadb,
		listModifiers:        listModifiersMariadb,
		listTypes:            listTypesMariadb,
		orderSupportAlter: []modifierService{
			uastModifierType,
			uastModifierDefault,
			uastModifierNotNull,
		},
		orderSupportCreate: []modifierService{
			uastModifierNotNull,
			uastModifierAutoIncrement,
			uastModifierDefault,
		},
		parensFunction:    true,
		placeholderNumber: -1,
		placeholderStyle:  "?",
		placeholderType:   false,
		symbolMarkLeft:    "'",
		symbolMarkRight:   "'",
		symbolQuoteLeft:   "`",
		symbolQuoteRight:  "`",
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
			uastModifierTable:  false,
			uastModifierView:   true,
		},
		supportComment: map[modifierService]bool{
			uastModifierColumn: true,
			uastModifierIndex:  false,
			uastModifierSchema: false,
			uastModifierTable:  true,
			uastModifierView:   false,
		},
		supportDrop: map[modifierService]bool{
			uastModifierColumn:     true,
			uastModifierConstraint: true,
			uastModifierDefault:    true,
			uastModifierNotNull:    true,
		},
		supportIfExists: map[modifierService]bool{
			uastModifierColumn: false,
			uastModifierIndex:  true,
			uastModifierSchema: true,
			uastModifierTable:  true,
			uastModifierView:   true,
		},
		supportIfNotExists: map[modifierService]bool{
			uastModifierColumn: false,
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
			uastModifierIndex:      false,
			uastModifierSchema:     false,
			uastModifierTable:      true,
			uastModifierView:       false,
		},
		supportReturning: true,
		supportSet: map[modifierService]bool{
			uastModifierDefault: true,
			uastModifierNotNull: true,
			uastModifierType:    true,
		},
	},
	name:      "MariaDB",
	strateger: &mariadbStrateger{},
}

// Приватные константы
const (
	// Функции агрегатные
	uastMariadbFunctionGroupConcat functionService = "GROUP_CONCAT"
	// Функции условий
	uastMariadbFunctionCase functionService = "CASE"
	// Функции конвертации
	uastMariadbFunctionCast functionService = "CAST"
	// Функции даты и времени
	uastMariadbFunctionDateAdd functionService = "DATE_ADD"
	uastMariadbFunctionDateSub functionService = "DATE_SUB"
	uastMariadbFunctionTimeAdd functionService = "TIME_ADD"
	uastMariadbFunctionTimeSub functionService = "TIME_SUB"
	// Функции обмена данными
	uastMariadbFunctionJsonExtract     functionService = ""
	uastMariadbFunctionJsonExtractCast functionService = "CAST"
	uastMariadbFunctionJsonRemove      functionService = "JSON_REMOVE"
	uastMariadbFunctionJsonSet         functionService = "JSON_SET"
	// Функции математические
	uastMariadbFunctionCeil  functionService = "CEILING"
	uastMariadbFunctionTrunc functionService = "TRUNCATE"
	// Функции строковые
)

// Приватные переменные
var listComparisonsMariadb = map[comparisonOperator]comparisonTransform{
	uastComparisonILike:    mariadbComparisonILike,
	uastComparisonNotILike: mariadbComparisonNotILike,
}
var listFunctionsMariadb = map[functionService]functionTransform{
	// Функции агрегатные
	uastFunctionGroupConcat: mariadbFunctionGroupConcat,
	// Функции условий
	uastFunctionCase: mariadbFunctionCase,
	// Функции конвертации
	uastFunctionCast: mariadbFunctionCast,
	// Функции даты и времени
	uastFunctionDateAdd: mariadbFunctionDateAdd,
	uastFunctionDateSub: mariadbFunctionDateSub,
	uastFunctionTimeAdd: mariadbFunctionTimeAdd,
	uastFunctionTimeSub: mariadbFunctionTimeSub,
	// Функции обмена данными
	uastFunctionJsonExtract: mariadbFunctionJsonExtract,
	uastFunctionJsonRemove:  mariadbFunctionJsonRemove,
	uastFunctionJsonSet:     mariadbFunctionJsonSet,
	// Функции математические
	uastFunctionCeil:  mariadbFunctionCeil,
	uastFunctionTrunc: mariadbFunctionTrunc,
	// Функции строковые
}
var listModifiersMariadb = map[modifierService]modifierService{
	uastModifierAutoIncrement: "AUTO_INCREMENT",
	uastModifierDefault:       "SET DEFAULT",
	uastModifierDefaultAction: uastModifierAlter,
	uastModifierNotNull:       uastModifierNotNull,
	uastModifierNotNullAction: uastModifierModify,
	uastModifierType:          "",
	uastModifierTypeAction:    uastModifierModify,
	uastModifierUpsert:        "ON DUPLICATE KEY UPDATE",
}
var listTypesMariadb = map[ValueType]typeService{
	// Типы бинарные
	TypeBinary:    "BINARY",
	TypeVarBinary: "VARBINARY",
	// Типы даты и времени
	TypeDate:      "DATE",
	TypeDateTime:  "DATETIME",
	TypeTime:      "TIME",
	TypeTimestamp: "DATETIME",
	// Типы числовые
	TypeBigInt:   "SIGNED",
	TypeDecimal:  "DECIMAL",
	TypeDouble:   "DECIMAL",
	TypeFloat:    "DECIMAL",
	TypeInt:      "SIGNED",
	TypeSmallInt: "SIGNED",
	// Типы строковые
	TypeChar:    "CHAR",
	TypeString:  "VARCHAR",
	TypeText:    "TEXT",
	TypeVarChar: "VARCHAR",
	// Типы специальные
	TypeArray:   "JSON",
	TypeBoolean: "TINYINT(1)",
	TypeJSON:    "JSON",
	TypeUUID:    "UUID",
	TypeXML:     "TEXT",
}

// Приватные структуры
type mariadbStrateger struct{}

// Приватные функции
func mariadbComparisonILike(baseTransformer *baseTransformer, expr transformComparison) error {
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
func mariadbComparisonNotILike(baseTransformer *baseTransformer, expr transformComparison) error {
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
func mariadbFunctionGroupConcat(baseTransformer *baseTransformer, expr transformFunction) error {
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
	function.service = uastMariadbFunctionGroupConcat
	return nil
}
func mariadbFunctionCase(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastMariadbFunctionCase)
	return nil
}
func mariadbFunctionCast(baseTransformer *baseTransformer, expr transformFunction) error {
	valueType := expr.getValueType()
	typeService, exists := listTypesMariadb[valueType]
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
	expr.setService(uastMariadbFunctionCast)
	return nil
}
func mariadbFunctionDateAdd(baseTransformer *baseTransformer, expr transformFunction) error {
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
	function.service = uastMariadbFunctionDateAdd
	return nil
}
func mariadbFunctionDateSub(baseTransformer *baseTransformer, expr transformFunction) error {
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
	function.service = uastMariadbFunctionDateSub
	return nil
}
func mariadbFunctionTimeAdd(baseTransformer *baseTransformer, expr transformFunction) error {
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
	function.service = uastMariadbFunctionTimeAdd
	return nil
}
func mariadbFunctionTimeSub(baseTransformer *baseTransformer, expr transformFunction) error {
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
	function.service = uastMariadbFunctionTimeSub
	return nil
}
func mariadbFunctionJsonExtract(baseTransformer *baseTransformer, expr transformFunction) error {
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
	typeService, exists := listTypesMariadb[valueType]
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
		expr.setOperator(uastCompositeSpaceMinusGreaterSpace)
		expr.setService(uastMariadbFunctionJsonExtract)
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
		expr.setOperator(uastCompositeSpaceMinusDoubleGreaterSpace)
		expr.setService(uastMariadbFunctionJsonExtract)
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
		expr.setOperator(uastCompositeSpaceMinusDoubleGreaterSpace)
		expr.setService(uastMariadbFunctionJsonExtractCast)
	}
	return nil
}
func mariadbFunctionJsonRemove(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.getJson()
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
	expr.setJson(groups)
	expr.setOperator(uastCompositeCommaSpace)
	expr.setService(uastMariadbFunctionJsonRemove)
	return nil
}
func mariadbFunctionJsonSet(baseTransformer *baseTransformer, expr transformFunction) error {
	json := expr.getJson()
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
	expr.setJson(groups)
	expr.setOperator(uastCompositeCommaSpace)
	expr.setService(uastMariadbFunctionJsonSet)
	return nil
}
func mariadbFunctionCeil(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastMariadbFunctionCeil)
	return nil
}
func mariadbFunctionTrunc(baseTransformer *baseTransformer, expr transformFunction) error {
	expr.setService(uastMariadbFunctionTrunc)
	return nil
}

// Приватные методы
func (strateger *mariadbStrateger) renderAlter(baseRenderer *baseRenderer, stmtAlter *stmtAlter) error {
	// !!!Внимание, находится в стадии разработки
	if err := baseRenderer.renderCommand(stmtAlter.command); err != nil {
		return err
	}
	if err := baseRenderer.renderEntity(stmtAlter.entity, false, false, stmtAlter.ifExists, stmtAlter.ifNotExists); err != nil {
		return err
	}
	switch stmtAlter.entity.(type) {
	case *sourceIndex:
		if stmtAlter.renameTo != "" {
			if err := baseRenderer.renderRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	case *sourceSchema:
		if stmtAlter.renameTo != "" {
			if err := baseRenderer.renderRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	case *sourceTable:
		if err := baseRenderer.renderTableModifyData(stmtAlter.addColumns, stmtAlter.addConstraints, stmtAlter.dropColumns, stmtAlter.dropConstraints, stmtAlter.setColumns); err != nil {
			return err
		}
		if err := baseRenderer.renderTableRenameData(stmtAlter.renameColumn, stmtAlter.renameConstraint); err != nil {
			return err
		}
		if stmtAlter.renameTo != "" {
			if err := baseRenderer.renderRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	case *sourceView:
		if stmtAlter.renameTo != "" {
			if err := baseRenderer.renderRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	}
	return nil
}
func (strateger *mariadbStrateger) renderComment(baseRenderer *baseRenderer, stmtComment *stmtComment) error {
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
func (strateger *mariadbStrateger) renderCreate(baseRenderer *baseRenderer, stmtCreate *stmtCreate) error {
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
		if err := baseRenderer.renderTableNewData(stmtCreate.columns, stmtCreate.constraints); err != nil {
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
func (strateger *mariadbStrateger) renderDelete(baseRenderer *baseRenderer, stmtDelete *stmtDelete) error {
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
func (strateger *mariadbStrateger) renderDrop(baseRenderer *baseRenderer, stmtDrop *stmtDrop) error {
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
func (strateger *mariadbStrateger) renderInsert(baseRenderer *baseRenderer, stmtInsert *stmtInsert) error {
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
func (strateger *mariadbStrateger) renderSelect(baseRenderer *baseRenderer, stmtSelect *stmtSelect) error {
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
func (strateger *mariadbStrateger) renderTruncate(baseRenderer *baseRenderer, stmtTruncate *stmtTruncate) error {
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
func (strateger *mariadbStrateger) renderUpdate(baseRenderer *baseRenderer, stmtUpdate *stmtUpdate) error {
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
func (strateger *mariadbStrateger) transformAlter(baseTransformer *baseTransformer, stmtAlter *stmtAlter) error {
	return nil
}
func (strateger *mariadbStrateger) transformComment(baseTransformer *baseTransformer, stmtComment *stmtComment) error {
	return nil
}
func (strateger *mariadbStrateger) transformCreate(baseTransformer *baseTransformer, stmtCreate *stmtCreate) error {
	return nil
}
func (strateger *mariadbStrateger) transformDelete(baseTransformer *baseTransformer, stmtDelete *stmtDelete) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *mariadbStrateger) transformDrop(baseTransformer *baseTransformer, stmtDrop *stmtDrop) error {
	return nil
}
func (strateger *mariadbStrateger) transformInsert(baseTransformer *baseTransformer, stmtInsert *stmtInsert) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	if stmtInsert.values != nil && stmtInsert.values.upsert != nil {
		if service, exists := listModifiersMariadb[uastModifierUpsert]; exists {
			stmtInsert.values.upsert.service = service
		}
	}
	return nil
}
func (strateger *mariadbStrateger) transformSelect(baseTransformer *baseTransformer, stmtSelect *stmtSelect) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *mariadbStrateger) transformTruncate(baseTransformer *baseTransformer, stmtTruncate *stmtTruncate) error {
	return nil
}
func (strateger *mariadbStrateger) transformUpdate(baseTransformer *baseTransformer, stmtUpdate *stmtUpdate) error {
	if err := baseTransformer.transformComparison(); err != nil {
		return err
	}
	if err := baseTransformer.transformFunction(); err != nil {
		return err
	}
	return nil
}
func (strateger *mariadbStrateger) validateAlter(baseValidator *baseValidator, stmtAlter *stmtAlter) error {
	// !!!Внимание, находится в стадии разработки
	if err := baseValidator.validateEntity(stmtAlter.entity); err != nil {
		return err
	}
	switch stmtAlter.entity.(type) {
	case *sourceIndex:
		if stmtAlter.renameTo != "" {
			if !baseValidator.config.supportRename[uastModifierIndex] {
				return ErrUnsupportEntityIndex
			}
			if err := baseValidator.validateRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	case *sourceSchema:
		if stmtAlter.renameTo != "" {
			if !baseValidator.config.supportRename[uastModifierSchema] {
				return ErrUnsupportEntitySchema
			}
			if err := baseValidator.validateRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
	case *sourceTable:
		if stmtAlter.renameTo != "" {
			if !baseValidator.config.supportRename[uastModifierTable] {
				return ErrUnsupportEntityTable
			}
			if err := baseValidator.validateRenameTo(stmtAlter.renameTo); err != nil {
				return err
			}
		}
		if stmtAlter.renameColumn != nil && stmtAlter.renameConstraint != nil {
			return ErrUnsupportEntityTable
		} else {
			if err := baseValidator.validateTableRenameData(stmtAlter.renameColumn, stmtAlter.renameConstraint); err != nil {
				return err
			}
		}
	case *sourceView:
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
func (strateger *mariadbStrateger) validateComment(baseValidator *baseValidator, stmtComment *stmtComment) error {
	if err := baseValidator.validateOnFrom(stmtComment.onTable, stmtComment.onColumn); err != nil {
		return err
	}
	if err := baseValidator.validateIsData(stmtComment.comment); err != nil {
		return err
	}
	return nil
}
func (strateger *mariadbStrateger) validateCreate(baseValidator *baseValidator, stmtCreate *stmtCreate) error {
	if err := baseValidator.validateEntity(stmtCreate.entity); err != nil {
		return err
	}
	switch stmtCreate.entity.(type) {
	case *sourceIndex:
		if err := baseValidator.validateIndex(stmtCreate.columns); err != nil {
			return err
		}
	case *sourceSchema:
	case *sourceTable:
		if err := baseValidator.validateTableNewData(stmtCreate.columns, stmtCreate.constraints); err != nil {
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
func (strateger *mariadbStrateger) validateDelete(baseValidator *baseValidator, stmtDelete *stmtDelete) error {
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
func (strateger *mariadbStrateger) validateDrop(baseValidator *baseValidator, stmtDrop *stmtDrop) error {
	if err := baseValidator.validateEntity(stmtDrop.entity); err != nil {
		return err
	}
	return nil
}
func (strateger *mariadbStrateger) validateInsert(baseValidator *baseValidator, stmtInsert *stmtInsert) error {
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
func (strateger *mariadbStrateger) validateSelect(baseValidator *baseValidator, stmtSelect *stmtSelect) error {
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
func (strateger *mariadbStrateger) validateTruncate(baseValidator *baseValidator, stmtTruncate *stmtTruncate) error {
	if err := baseValidator.validateTable(stmtTruncate.table); err != nil {
		return err
	}
	return nil
}
func (strateger *mariadbStrateger) validateUpdate(baseValidator *baseValidator, stmtUpdate *stmtUpdate) error {
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
