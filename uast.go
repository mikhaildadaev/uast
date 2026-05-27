// Copyright [2026] [Mikhail Dadaev]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package uast

import "time"

// Приватные типы
type (
	binaryOperator     string
	comparisonOperator string
	compositeOperator  string
	logicalOperator    string
	joinOperator       string
	orderOperator      string
	unionOperator      string
	functionService    string
	managementService  string
	modifierService    string
	typeService        string
	formattingStage    string
	processingStage    string
)

// Приватные константы
const (
	uastConstFloat32One    float32 = 1.0
	uastConstFloat64One    float64 = 1.000000
	uastConstIntOne        int     = 1
	uastConstInt8One       int8    = 1
	uastConstInt16One      int16   = 1
	uastConstInt32One      int32   = 1
	uastConstInt64One      int64   = 1
	uastConstStringDefault string  = "DEFAULT"
	uastConstStringFalse   string  = "FALSE"
	uastConstStringNull    string  = "NULL"
	uastConstStringTrue    string  = "TRUE"
	uastConstUintOne       uint    = 1
	uastConstUint8One      uint8   = 1
	uastConstUint16One     uint16  = 1
	uastConstUint32One     uint32  = 1
	uastConstUint64One     uint64  = 1
)
const (
	uastCountMaxDepth      = 16
	uastCountMaxComparison = 48
	uastCountMaxFunction   = 128
	uastCountMaxLimit      = 64
	uastCountMaxSubquery   = 32
	uastCountMaxUnions     = 16
	uastCountMaxWith       = 8
	uastSizeInitByte       = 128
	uastSizeInitComparison = 8
	uastSizeInitExpr       = 8
	uastSizeInitFunction   = 8
	uastSizeInitQuery      = 256
	uastSizeInitValue      = 16
)
const (
	// Бинарные операторы
	uastBinaryBitwiseAnd binaryOperator = "&"
	uastBinaryBitwiseOr  binaryOperator = "|"
	uastBinaryBitwiseXor binaryOperator = "^"
	uastBinaryDivide     binaryOperator = "/"
	uastBinaryMinus      binaryOperator = "-"
	uastBinaryModulo     binaryOperator = "%"
	uastBinaryMultiply   binaryOperator = "*"
	uastBinaryPlus       binaryOperator = "+"
	uastBinaryShiftLeft  binaryOperator = "<<"
	uastBinaryShiftRight binaryOperator = ">>"
)
const (
	// Сравнительные операторы
	uastComparisonBetween      comparisonOperator = "BETWEEN"
	uastComparisonEqual        comparisonOperator = "="
	uastComparisonExists       comparisonOperator = "EXISTS"
	uastComparisonGreater      comparisonOperator = ">"
	uastComparisonGreaterEqual comparisonOperator = ">="
	uastComparisonILike        comparisonOperator = "ILIKE"
	uastComparisonIn           comparisonOperator = "IN"
	uastComparisonIsNotNull    comparisonOperator = "IS NOT NULL"
	uastComparisonIsNull       comparisonOperator = "IS NULL"
	uastComparisonLess         comparisonOperator = "<"
	uastComparisonLessEqual    comparisonOperator = "<="
	uastComparisonLike         comparisonOperator = "LIKE"
	uastComparisonNotBetween   comparisonOperator = "NOT BETWEEN"
	uastComparisonNotEqual     comparisonOperator = "<>"
	uastComparisonNotExists    comparisonOperator = "NOT EXISTS"
	uastComparisonNotILike     comparisonOperator = "NOT ILIKE"
	uastComparisonNotIn        comparisonOperator = "NOT IN"
	uastComparisonNotLike      comparisonOperator = "NOT LIKE"
)
const (
	// Разделительные операторы
	uastCompositeBraceLeft                    compositeOperator = "{"
	uastCompositeBraceRight                   compositeOperator = "}"
	uastCompositeBrackLeft                    compositeOperator = "["
	uastCompositeBrackRight                   compositeOperator = "]"
	uastCompositeCommaSpace                   compositeOperator = ", "
	uastCompositeDollarPoint                  compositeOperator = "$."
	uastCompositeParenLeft                    compositeOperator = "("
	uastCompositeParenRight                   compositeOperator = ")"
	uastCompositeSpaceAtGreaterSpace          compositeOperator = " @> "
	uastCompositeSpaceEqualSpace              compositeOperator = " = "
	uastCompositeSpaceMinusDoubleGreaterSpace compositeOperator = " ->> "
	uastCompositeSpaceMinusGreaterSpace       compositeOperator = " -> "
	uastCompositeSpaceMinusSpace              compositeOperator = " - "
	uastCompositeSpaceSignDoubleGreaterSpace  compositeOperator = " #>> "
	uastCompositeSpaceSignGreaterSpace        compositeOperator = " #> "
	uastCompositeSingleComma                  compositeOperator = ","
	uastCompositeSinglePoint                  compositeOperator = "."
	uastCompositeSingleSpace                  compositeOperator = " "
)
const (
	// Функции агрегатные
	uastFunctionAvg         functionService = "AVG"
	uastFunctionBitAnd      functionService = "BIT_AND"
	uastFunctionBitOr       functionService = "BIT_OR"
	uastFunctionBitXor      functionService = "BIT_XOR"
	uastFunctionCount       functionService = "COUNT"
	uastFunctionGroupConcat functionService = "GROUP_CONCAT"
	uastFunctionMax         functionService = "MAX"
	uastFunctionMin         functionService = "MIN"
	uastFunctionStdDev      functionService = "STDDEV"
	uastFunctionSum         functionService = "SUM"
	uastFunctionVariance    functionService = "VARIANCE"
	// Функции аналитические
	uastFunctionFirstValue functionService = "FIRST_VALUE"
	uastFunctionLag        functionService = "LAG"
	uastFunctionLastValue  functionService = "LAST_VALUE"
	uastFunctionLead       functionService = "LEAD"
	uastFunctionNthValue   functionService = "NTH_VALUE"
	// Функции условий
	uastFunctionCase     functionService = "CASE"
	uastFunctionCoalesce functionService = "COALESCE"
	uastFunctionGreatest functionService = "GREATEST"
	uastFunctionLeast    functionService = "LEAST"
	uastFunctionNullIf   functionService = "NULLIF"
	// Функции конвертации
	uastFunctionCast       functionService = "CAST"
	uastFunctionCharLength functionService = "CHAR_LENGTH"
	uastFunctionDateFormat functionService = "DATE_FORMAT"
	uastFunctionDegrees    functionService = "DEGREES"
	uastFunctionLength     functionService = "LENGTH"
	uastFunctionPosition   functionService = "POSITION"
	uastFunctionRadians    functionService = "RADIANS"
	// Функции даты и времени
	uastFunctionCurDate   functionService = "CURDATE"
	uastFunctionCurTime   functionService = "CURTIME"
	uastFunctionDateAdd   functionService = "DATE_ADD"
	uastFunctionDateDiff  functionService = "DATEDIFF"
	uastFunctionDateSub   functionService = "DATE_SUB"
	uastFunctionDay       functionService = "DAY"
	uastFunctionDayName   functionService = "DAYNAME"
	uastFunctionHour      functionService = "HOUR"
	uastFunctionMinute    functionService = "MINUTE"
	uastFunctionMonth     functionService = "MONTH"
	uastFunctionMonthName functionService = "MONTHNAME"
	uastFunctionNow       functionService = "NOW"
	uastFunctionQuarter   functionService = "QUARTER"
	uastFunctionSecond    functionService = "SECOND"
	uastFunctionTimeAdd   functionService = "TIME_ADD"
	uastFunctionTimeDiff  functionService = "TIMEDIFF"
	uastFunctionTimeSub   functionService = "TIME_SUB"
	uastFunctionWeek      functionService = "WEEK"
	uastFunctionYear      functionService = "YEAR"
	// Функции обмена данными
	uastFunctionJsonArray     functionService = "JSON_ARRAY"
	uastFunctionJsonArrayAgg  functionService = "JSON_ARRAYAGG"
	uastFunctionJsonContains  functionService = "JSON_CONTAINS"
	uastFunctionJsonExtract   functionService = "JSON_EXTRACT"
	uastFunctionJsonObject    functionService = "JSON_OBJECT"
	uastFunctionJsonObjectAgg functionService = "JSON_OBJECTAGG"
	uastFunctionJsonRemove    functionService = "JSON_REMOVE"
	uastFunctionJsonSet       functionService = "JSON_SET"
	uastFunctionJsonType      functionService = "JSON_TYPE"
	// Функции математические
	uastFunctionAbs   functionService = "ABS"
	uastFunctionACos  functionService = "ACOS"
	uastFunctionASin  functionService = "ASIN"
	uastFunctionATan  functionService = "ATAN"
	uastFunctionATan2 functionService = "ATAN2"
	uastFunctionCbrt  functionService = "CBRT"
	uastFunctionCeil  functionService = "CEIL"
	uastFunctionCos   functionService = "COS"
	uastFunctionExp   functionService = "EXP"
	uastFunctionFloor functionService = "FLOOR"
	uastFunctionLn    functionService = "LN"
	uastFunctionLog   functionService = "LOG"
	uastFunctionMod   functionService = "MOD"
	uastFunctionPi    functionService = "PI"
	uastFunctionPower functionService = "POWER"
	uastFunctionRand  functionService = "RAND"
	uastFunctionRound functionService = "ROUND"
	uastFunctionSin   functionService = "SIN"
	uastFunctionSqrt  functionService = "SQRT"
	uastFunctionTan   functionService = "TAN"
	uastFunctionTrunc functionService = "TRUNC"
	// Функции строковые
	uastFunctionConcat      functionService = "CONCAT"
	uastFunctionConcatWs    functionService = "CONCAT_WS"
	uastFunctionLeftString  functionService = "LEFT"
	uastFunctionLower       functionService = "LOWER"
	uastFunctionLPad        functionService = "LPAD"
	uastFunctionLTrim       functionService = "LTRIM"
	uastFunctionRepeat      functionService = "REPEAT"
	uastFunctionReplace     functionService = "REPLACE"
	uastFunctionReverse     functionService = "REVERSE"
	uastFunctionRightString functionService = "RIGHT"
	uastFunctionRPad        functionService = "RPAD"
	uastFunctionRTrim       functionService = "RTRIM"
	uastFunctionSubString   functionService = "SUBSTRING"
	uastFunctionTrim        functionService = "TRIM"
	uastFunctionUpper       functionService = "UPPER"
	// Функции ранжирующие
	uastFunctionCumeDist    functionService = "CUME_DIST"
	uastFunctionDenseRank   functionService = "DENSE_RANK"
	uastFunctionNTile       functionService = "NTILE"
	uastFunctionPercentRank functionService = "PERCENT_RANK"
	uastFunctionRank        functionService = "RANK"
	uastFunctionRowNumber   functionService = "ROW_NUMBER"
)
const (
	// Обьединяющие операторы
	uastJoinCross      joinOperator = "CROSS JOIN"
	uastJoinFull       joinOperator = "FULL JOIN"
	uastJoinFullOuter  joinOperator = "FULL OUTER JOIN"
	uastJoinInner      joinOperator = "INNER JOIN"
	uastJoinLeft       joinOperator = "LEFT JOIN"
	uastJoinLeftOuter  joinOperator = "LEFT OUTER JOIN"
	uastJoinRight      joinOperator = "RIGHT JOIN"
	uastJoinRightOuter joinOperator = "RIGHT OUTER JOIN"
)
const (
	// Соединительные операторы
	uastLogicalAnd logicalOperator = "AND"
	uastLogicalOr  logicalOperator = "OR"
)
const (
	// Сортировочные операторы
	uastOrderAsc  orderOperator = "ASC"
	uastOrderDesc orderOperator = "DESC"
)
const (
	// Комбинирующие операторы
	uastUnion          unionOperator = "UNION"
	uastUnionAll       unionOperator = "UNION ALL"
	uastUnionExcept    unionOperator = "EXCEPT"
	uastUnionIntersect unionOperator = "INTERSECT"
)
const (
	// Конструкции информации
	uastManagementDatabase    managementService = "DATABASE"
	uastManagementCurrentUser managementService = "CURRENT_USER"
	uastManagementSessionUser managementService = "SESSION_USER"
	uastManagementSystemUser  managementService = "SYSTEM_USER"
	uastManagementUser        managementService = "USER"
	uastManagementVersion     managementService = "VERSION"
	// Конструкции времени
	uastManagementBenchmark managementService = "BENCHMARK"
	uastManagementSleep     managementService = "SLEEP"
	uastManagementWaitFor   managementService = "WAITFOR"
	uastManagementDelay     managementService = "DELAY"
	uastManagementTimeout   managementService = "TIMEOUT"
	// Конструкции управления
	uastManagementAlter       managementService = "ALTER"
	uastManagementComment     managementService = "COMMENT"
	uastManagementCreate      managementService = "CREATE"
	uastManagementDrop        managementService = "DROP"
	uastManagementDelete      managementService = "DELETE"
	uastManagementInsert      managementService = "INSERT"
	uastManagementSelect      managementService = "SELECT"
	uastManagementUpdate      managementService = "UPDATE"
	uastManagementTruncate    managementService = "TRUNCATE"
	uastManagementFetchNext   managementService = "FETCH NEXT"
	uastManagementFrom        managementService = "FROM"
	uastManagementInto        managementService = "INTO"
	uastManagementSet         managementService = "SET"
	uastManagementTo          managementService = "TO"
	uastManagementGroupBy     managementService = "GROUP BY"
	uastManagementHaving      managementService = "HAVING"
	uastManagementJoin        managementService = "JOIN"
	uastManagementLimit       managementService = "LIMIT"
	uastManagementOffset      managementService = "OFFSET"
	uastManagementOrderBy     managementService = "ORDER BY"
	uastManagementOutput      managementService = "OUTPUT"
	uastManagementPartitionBy managementService = "PARTITION BY"
	uastManagementReturning   managementService = "RETURNING"
	uastManagementUsing       managementService = "USING"
	uastManagementValues      managementService = "VALUES"
	uastManagementWhere       managementService = "WHERE"
	uastManagementWith        managementService = "WITH"
)
const (
	// Конструкции модификаторов
	uastModifierAnd             modifierService = "AND"
	uastModifierAs              modifierService = "AS"
	uastModifierCascade         modifierService = "CASCADE"
	uastModifierColumn          modifierService = "COLUMN"
	uastModifierBetween         modifierService = "BETWEEN"
	uastModifierDistinct        modifierService = "DISTINCT"
	uastModifierElse            modifierService = "ELSE"
	uastModifierEnd             modifierService = "END"
	uastModifierIfExists        modifierService = "IF EXISTS"
	uastModifierIn              modifierService = "IN"
	uastModifierInterval        modifierService = "INTERVAL"
	uastModifierIs              modifierService = "IS"
	uastModifierMonth           modifierService = "MONTH"
	uastModifierOn              modifierService = "ON"
	uastModifierOver            modifierService = "OVER"
	uastModifierRestartIdentity modifierService = "RESTART IDENTITY"
	uastModifierRecursive       modifierService = "RECURSIVE"
	uastModifierRows            modifierService = "ROWS"
	uastModifierRowsOnly        modifierService = "ROWS ONLY"
	uastModifierSeparator       modifierService = "SEPARATOR"
	uastModifierTable           modifierService = "TABLE"
	uastModifierThen            modifierService = "THEN"
	uastModifierWeekday         modifierService = "WEEKDAY"
	uastModifierWhen            modifierService = "WHEN"
)
const (
	// Уровень форматирования
	uastFormatAlias   formattingStage = "ALIAS"
	uastFormatName    formattingStage = "NAME"
	uastFormatLiteral formattingStage = "LITERAL"
	// Уровень обработки
	uastProcessСross  processingStage = "CROSS"
	uastProcessDirect processingStage = "DIRECT"
	uastProcessEmpty  processingStage = "EMPTY"
	uastProcessInvert processingStage = "INVERT"
	uastProcessJson   processingStage = "JSON"
	uastProcessWindow processingStage = "WINDOW"
)

// Приватные переменные
var (
	constFunctionServices   map[functionService]string
	constFunctionParameters map[functionService]struct {
		distinct bool
		min, max int
	}
	constBinaryOperators     map[binaryOperator]string
	constComparisonOperators map[comparisonOperator]string
	constCompositeOperators  map[compositeOperator]string
	constJoinOperators       map[joinOperator]string
	constLogicalOperators    map[logicalOperator]string
	constOrderOperators      map[orderOperator]string
	constUnionOperators      map[unionOperator]string
	constManagementServices  map[managementService]string
	constModifierServices    map[modifierService]string
	constKeywordUniversal    map[string]struct{}
)
var listBinaryOperators = []binaryOperator{uastBinaryBitwiseAnd, uastBinaryBitwiseOr, uastBinaryBitwiseXor, uastBinaryDivide, uastBinaryMinus, uastBinaryModulo, uastBinaryMultiply, uastBinaryPlus, uastBinaryShiftLeft, uastBinaryShiftRight}
var listComparisonOperators = []comparisonOperator{uastComparisonBetween, uastComparisonEqual, uastComparisonExists, uastComparisonGreater, uastComparisonGreaterEqual, uastComparisonILike, uastComparisonIn, uastComparisonIsNotNull, uastComparisonIsNull, uastComparisonLess, uastComparisonLessEqual, uastComparisonLike, uastComparisonNotBetween, uastComparisonNotEqual, uastComparisonNotExists, uastComparisonNotILike, uastComparisonNotIn, uastComparisonNotLike}
var listCompositeOperators = []compositeOperator{uastCompositeCommaSpace, uastCompositeParenLeft, uastCompositeParenRight, uastCompositeSinglePoint, uastCompositeSingleSpace}
var listJoinOperators = []joinOperator{uastJoinCross, uastJoinFull, uastJoinFullOuter, uastJoinInner, uastJoinLeft, uastJoinLeftOuter, uastJoinRight, uastJoinRightOuter}
var listLogicalOperators = []logicalOperator{uastLogicalAnd, uastLogicalOr}
var listOrderOperators = []orderOperator{uastOrderAsc, uastOrderDesc}
var listUnionOperators = []unionOperator{uastUnion, uastUnionAll, uastUnionExcept, uastUnionIntersect}
var listFunctionServices = []functionService{
	// Функции агрегатные
	uastFunctionAvg, uastFunctionBitAnd, uastFunctionBitOr, uastFunctionBitXor, uastFunctionCount, uastFunctionGroupConcat, uastFunctionMax, uastFunctionMin, uastFunctionStdDev, uastFunctionSum, uastFunctionVariance,
	// Функции аналитические
	uastFunctionFirstValue, uastFunctionLag, uastFunctionLastValue, uastFunctionLead, uastFunctionNthValue,
	// Функции условий
	uastFunctionCase, uastFunctionCoalesce, uastFunctionGreatest, uastFunctionLeast, uastFunctionNullIf,
	// Функции конвертации
	uastFunctionCast, uastFunctionCharLength, uastFunctionDateFormat, uastFunctionDegrees, uastFunctionLength, uastFunctionPosition, uastFunctionRadians,
	// Функции даты и времени
	uastFunctionCurDate, uastFunctionCurTime, uastFunctionDateAdd, uastFunctionDateDiff, uastFunctionDateSub, uastFunctionDay, uastFunctionDayName, uastFunctionHour, uastFunctionMinute, uastFunctionMonth, uastFunctionMonthName, uastFunctionNow, uastFunctionQuarter, uastFunctionSecond, uastFunctionTimeAdd, uastFunctionTimeDiff, uastFunctionTimeSub, uastFunctionWeek, uastFunctionYear,
	// Функции обмена данными
	uastFunctionJsonArray, uastFunctionJsonArrayAgg, uastFunctionJsonContains, uastFunctionJsonExtract, uastFunctionJsonObject, uastFunctionJsonObjectAgg, uastFunctionJsonRemove, uastFunctionJsonSet, uastFunctionJsonType,
	// Функции математические
	uastFunctionAbs, uastFunctionACos, uastFunctionASin, uastFunctionATan, uastFunctionATan2, uastFunctionCbrt, uastFunctionCeil, uastFunctionCos, uastFunctionExp, uastFunctionFloor, uastFunctionLn, uastFunctionLog, uastFunctionMod, uastFunctionPi, uastFunctionPower, uastFunctionRand, uastFunctionRound, uastFunctionSin, uastFunctionSqrt, uastFunctionTan, uastFunctionTrunc,
	// Функции строковые
	uastFunctionConcat, uastFunctionConcatWs, uastFunctionLeftString, uastFunctionLower, uastFunctionLPad, uastFunctionLTrim, uastFunctionRepeat, uastFunctionReplace, uastFunctionReverse, uastFunctionRightString, uastFunctionRPad, uastFunctionRTrim, uastFunctionSubString, uastFunctionTrim, uastFunctionUpper,
	// Функции ранжирующие
	uastFunctionCumeDist, uastFunctionDenseRank, uastFunctionNTile, uastFunctionPercentRank, uastFunctionRank, uastFunctionRowNumber,
}
var listManagementServices = []managementService{uastManagementDatabase, uastManagementCurrentUser, uastManagementSessionUser, uastManagementSystemUser, uastManagementUser, uastManagementVersion, uastManagementBenchmark, uastManagementDelay, uastManagementSleep, uastManagementTimeout, uastManagementWaitFor, uastManagementAlter, uastManagementComment, uastManagementCreate, uastManagementDrop, uastManagementDelete, uastManagementInsert, uastManagementSelect, uastManagementUpdate, uastManagementTruncate, uastManagementFetchNext, uastManagementFrom, uastManagementInto, uastManagementSet, uastManagementTo, uastManagementGroupBy, uastManagementHaving, uastManagementJoin, uastManagementLimit, uastManagementOffset, uastManagementOrderBy, uastManagementPartitionBy, uastManagementReturning, uastManagementUsing, uastManagementValues, uastManagementWhere, uastManagementWith}
var listModifierServices = []modifierService{uastModifierAnd, uastModifierAs, uastModifierCascade, uastModifierColumn, uastModifierBetween, uastModifierDistinct, uastModifierElse, uastModifierEnd, uastModifierIfExists, uastModifierIn, uastModifierInterval, uastModifierIs, uastModifierMonth, uastModifierOn, uastModifierOver, uastModifierRestartIdentity, uastModifierRecursive, uastModifierRows, uastModifierRowsOnly, uastModifierSeparator, uastModifierTable, uastModifierThen, uastModifierWeekday, uastModifierWhen}
var listSupportDialects = []*SupportDialect{DialectMariaDB, DialectMsSQL, DialectMySQL, DialectPostgreSQL, DialectSQLite}
var listSymbolSafeAlias = [256]bool{
	'+': true, '-': true, '_': true,
	'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
	'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true, 'H': true, 'I': true, 'J': true, 'K': true, 'L': true, 'M': true, 'N': true, 'O': true, 'P': true, 'Q': true, 'R': true, 'S': true, 'T': true, 'U': true, 'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true,
	'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true, 'h': true, 'i': true, 'j': true, 'k': true, 'l': true, 'm': true, 'n': true, 'o': true, 'p': true, 'q': true, 'r': true, 's': true, 't': true, 'u': true, 'v': true, 'w': true, 'x': true, 'y': true, 'z': true,
}
var listSymbolSafeName = [256]bool{
	'_': true,
	'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
	'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true, 'H': true, 'I': true, 'J': true, 'K': true, 'L': true, 'M': true, 'N': true, 'O': true, 'P': true, 'Q': true, 'R': true, 'S': true, 'T': true, 'U': true, 'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true,
	'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true, 'h': true, 'i': true, 'j': true, 'k': true, 'l': true, 'm': true, 'n': true, 'o': true, 'p': true, 'q': true, 'r': true, 's': true, 't': true, 'u': true, 'v': true, 'w': true, 'x': true, 'y': true, 'z': true,
}
var listSymbolSafeLiteral = [256]bool{
	' ': true, '%': true, '+': true, ',': true, '-': true, '.': true, '/': true, ':': true, '_': true,
	'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
	'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true, 'H': true, 'I': true, 'J': true, 'K': true, 'L': true, 'M': true, 'N': true, 'O': true, 'P': true, 'Q': true, 'R': true, 'S': true, 'T': true, 'U': true, 'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true,
	'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true, 'h': true, 'i': true, 'j': true, 'k': true, 'l': true, 'm': true, 'n': true, 'o': true, 'p': true, 'q': true, 'r': true, 's': true, 't': true, 'u': true, 'v': true, 'w': true, 'x': true, 'y': true, 'z': true,
}

// Приватные интерфейсы
type strateger interface {
	statementRenderer
	statementTransformer
	statementValidator
}
type typeNumeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}
type typeScalar interface {
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~bool | time.Time
}
type typeString interface {
	~string
}

// Приватные структуры
type config struct {
	lengthMaxArray       int
	lengthMaxConst       int
	lengthMaxFunc        int
	lengthMaxIdent       int
	lengthMaxLimit       int
	lengthMaxParam       int
	lengthMaxQuery       int
	lengthMaxValueByte   int
	lengthMaxValueString int
	listComparisons      map[comparisonOperator]comparisonTransform
	listFunctions        map[functionService]functionTransform
	parensFunction       bool
	placeholderNumber    int
	placeholderStyle     string
	placeholderType      bool
	symbolMarkLeft       string
	symbolMarkRight      string
	symbolQuoteLeft      string
	symbolQuoteRight     string
}
type casePair[T typeScalar] struct {
	then ExpressionSafe[T]
	when ExpressionBase
}

// Приватные функции
func init() {
	initOperatorMaps()
	initServiceMaps()
	initUniversalMaps()
}
func initOperatorMaps() {
	constBinaryOperators = make(map[binaryOperator]string, len(listBinaryOperators))
	for _, keyword := range listBinaryOperators {
		constBinaryOperators[keyword] = string(keyword)
	}
	constComparisonOperators = make(map[comparisonOperator]string, len(listComparisonOperators))
	for _, keyword := range listComparisonOperators {
		constComparisonOperators[keyword] = string(keyword)
	}
	constCompositeOperators = make(map[compositeOperator]string, len(listCompositeOperators))
	for _, keyword := range listCompositeOperators {
		constCompositeOperators[keyword] = string(keyword)
	}
	constJoinOperators = make(map[joinOperator]string, len(listJoinOperators))
	for _, keyword := range listJoinOperators {
		constJoinOperators[keyword] = string(keyword)
	}
	constLogicalOperators = make(map[logicalOperator]string, len(listLogicalOperators))
	for _, keyword := range listLogicalOperators {
		constLogicalOperators[keyword] = string(keyword)
	}
	constOrderOperators = make(map[orderOperator]string, len(listOrderOperators))
	for _, keyword := range listOrderOperators {
		constOrderOperators[keyword] = string(keyword)
	}
	constUnionOperators = make(map[unionOperator]string, len(listUnionOperators))
	for _, keyword := range listUnionOperators {
		constUnionOperators[keyword] = string(keyword)
	}
}
func initServiceMaps() {
	constFunctionServices = make(map[functionService]string, len(listFunctionServices))
	for _, keyword := range listFunctionServices {
		constFunctionServices[keyword] = string(keyword)
	}
	constFunctionParameters = map[functionService]struct {
		distinct bool
		min, max int
	}{
		// Функции агрегатные
		uastFunctionAvg: {true, 1, 1}, uastFunctionBitAnd: {true, 1, 1}, uastFunctionBitOr: {true, 1, 1}, uastFunctionBitXor: {true, 1, 1}, uastFunctionCount: {true, 0, 1}, uastFunctionGroupConcat: {true, 1, 1}, uastFunctionMax: {true, 1, 1}, uastFunctionMin: {true, 1, 1}, uastFunctionStdDev: {true, 1, 1}, uastFunctionSum: {true, 1, 1}, uastFunctionVariance: {true, 1, 1},
		// Функции аналитические
		uastFunctionFirstValue: {false, 1, 1}, uastFunctionLag: {false, 1, 3}, uastFunctionLastValue: {false, 1, 1}, uastFunctionLead: {false, 1, 3}, uastFunctionNthValue: {false, 2, 2},
		// Функции условий
		uastFunctionCase: {false, 2, -1}, uastFunctionCoalesce: {false, 1, -1}, uastFunctionGreatest: {false, 1, -1}, uastFunctionLeast: {false, 1, -1}, uastFunctionNullIf: {false, 2, 2},
		// Функции конвертации
		uastFunctionCast: {false, 2, 2}, uastFunctionCharLength: {false, 1, 1}, uastFunctionDateFormat: {false, 2, 2}, uastFunctionDegrees: {false, 1, 1}, uastFunctionLength: {false, 1, 1}, uastFunctionPosition: {false, 3, 3}, uastFunctionRadians: {false, 1, 1},
		// Функции даты и времени
		uastFunctionCurDate: {false, 0, 0}, uastFunctionCurTime: {false, 0, 0}, uastFunctionDateAdd: {false, 2, 2}, uastFunctionDateDiff: {false, 2, 2}, uastFunctionDateSub: {false, 2, 2}, uastFunctionDay: {false, 1, 1}, uastFunctionDayName: {false, 1, 1}, uastFunctionHour: {false, 1, 1}, uastFunctionMinute: {false, 1, 1}, uastFunctionMonth: {false, 1, 1}, uastFunctionMonthName: {false, 1, 1}, uastFunctionNow: {false, 0, 0}, uastFunctionQuarter: {false, 1, 1}, uastFunctionSecond: {false, 1, 1}, uastFunctionTimeAdd: {false, 2, 2}, uastFunctionTimeDiff: {false, 2, 2}, uastFunctionTimeSub: {false, 2, 2}, uastFunctionWeek: {false, 1, 1}, uastFunctionYear: {false, 1, 1},
		// Функции обмена данными
		uastFunctionJsonArray: {false, 1, -1}, uastFunctionJsonArrayAgg: {false, 1, 1}, uastFunctionJsonContains: {false, 2, 2}, uastFunctionJsonExtract: {false, 3, -1}, uastFunctionJsonObject: {false, 1, -1}, uastFunctionJsonObjectAgg: {false, 1, 1}, uastFunctionJsonRemove: {false, 2, -1}, uastFunctionJsonSet: {false, 3, -1}, uastFunctionJsonType: {false, 1, 1},
		// Функции математические
		uastFunctionAbs: {false, 1, 1}, uastFunctionACos: {false, 1, 1}, uastFunctionASin: {false, 1, 1}, uastFunctionATan: {false, 1, 1}, uastFunctionATan2: {false, 2, 2}, uastFunctionCbrt: {false, 1, 1}, uastFunctionCeil: {false, 1, 1}, uastFunctionCos: {false, 1, 1}, uastFunctionExp: {false, 1, 1}, uastFunctionFloor: {false, 1, 1}, uastFunctionLn: {false, 1, 1}, uastFunctionLog: {false, 2, 2}, uastFunctionMod: {false, 2, 2}, uastFunctionPi: {false, 0, 0}, uastFunctionPower: {false, 2, 2}, uastFunctionRand: {false, 0, 0}, uastFunctionRound: {false, 2, 2}, uastFunctionSin: {false, 1, 1}, uastFunctionSqrt: {false, 1, 1}, uastFunctionTan: {false, 1, 1}, uastFunctionTrunc: {false, 2, 2},
		// Функции строковые
		uastFunctionConcat: {false, 1, -1}, uastFunctionConcatWs: {false, 2, -1}, uastFunctionLeftString: {false, 2, 2}, uastFunctionLower: {false, 1, 1}, uastFunctionLPad: {false, 3, 3}, uastFunctionLTrim: {false, 1, 1}, uastFunctionRepeat: {false, 2, 2}, uastFunctionReplace: {false, 3, 3}, uastFunctionReverse: {false, 1, 1}, uastFunctionRightString: {false, 2, 2}, uastFunctionRPad: {false, 3, 3}, uastFunctionRTrim: {false, 1, 1}, uastFunctionSubString: {false, 3, 3}, uastFunctionTrim: {false, 1, 1}, uastFunctionUpper: {false, 1, 1},
		// Функции ранжирующие
		uastFunctionCumeDist: {false, 0, 0}, uastFunctionDenseRank: {false, 0, 0}, uastFunctionNTile: {false, 1, 1}, uastFunctionPercentRank: {false, 0, 0}, uastFunctionRank: {false, 0, 0}, uastFunctionRowNumber: {false, 0, 0},
	}
	constManagementServices = make(map[managementService]string, len(listManagementServices))
	for _, keyword := range listManagementServices {
		constManagementServices[keyword] = string(keyword)
	}
	constModifierServices = make(map[modifierService]string, len(listModifierServices))
	for _, keyword := range listModifierServices {
		constModifierServices[keyword] = string(keyword)
	}
}
func initUniversalMaps() {
	constKeywordUniversal = make(map[string]struct{})
	for keyword := range constBinaryOperators {
		constKeywordUniversal[string(keyword)] = struct{}{}
	}
	for keyword := range constComparisonOperators {
		constKeywordUniversal[string(keyword)] = struct{}{}
	}
	for keyword := range constLogicalOperators {
		constKeywordUniversal[string(keyword)] = struct{}{}
	}
	for keyword := range constJoinOperators {
		constKeywordUniversal[string(keyword)] = struct{}{}
	}
	for keyword := range constOrderOperators {
		constKeywordUniversal[string(keyword)] = struct{}{}
	}
	for keyword := range constUnionOperators {
		constKeywordUniversal[string(keyword)] = struct{}{}
	}
	for keyword := range constFunctionServices {
		constKeywordUniversal[string(keyword)] = struct{}{}
	}
	for keyword := range constManagementServices {
		constKeywordUniversal[string(keyword)] = struct{}{}
	}
	for keyword := range constModifierServices {
		constKeywordUniversal[string(keyword)] = struct{}{}
	}
}
func isSecureString(value string, format formattingStage) bool {
	length := len(value)
	// Метки символьных инъекций (плейсхолдеры, спецсимволы)
	switch format {
	case uastFormatAlias:
		for i := 0; i < length; i++ {
			symbol := value[i]
			if symbol >= 128 || !listSymbolSafeAlias[symbol] {
				return false
			}
		}
	case uastFormatName:
		for i := 0; i < length; i++ {
			symbol := value[i]
			if symbol >= 128 || !listSymbolSafeName[symbol] {
				return false
			}
		}
	case uastFormatLiteral:
		for i := 0; i < length; i++ {
			symbol := value[i]
			if symbol >= 128 || !listSymbolSafeLiteral[symbol] {
				return false
			}
		}
	}
	// Метки строковых инъекций (комментарии, спецстроки)
	for i := 0; i < length-1; i++ {
		symbolEven := value[i]
		symbolOdd := value[i+1]
		if (symbolEven == '/' && symbolOdd == '*') || (symbolEven == '*' && symbolOdd == '/') || (symbolEven == '-' && symbolOdd == '-') {
			return false
		}
	}
	return true
}
func operatorString[T typeString](value T) ExpressionSafe[string] {
	return &exprOperator[string]{
		value: string(value),
	}
}
func serviceString[T typeString](value T) ExpressionSafe[string] {
	return &exprService[string]{
		value: string(value),
	}
}
