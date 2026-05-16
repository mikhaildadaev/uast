package uast

import (
	"regexp"
	"testing"
)

// Публичные функции
func Test_Core_Array(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				Array(0, 1, 2),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "(?, ?, ?)", "ARRAY INT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `($1, $2, $3)`, "ARRAY INT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_Binary(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.ID,
			).
			Where(
				And(
					Equal(Test.Number, BitwiseAnd(Test.Number, Value(0b0010))),
					Equal(Test.Number, BitwiseOr(Test.Number, Value(0b0010))),
					Equal(Test.Number, BitwiseXor(Test.Number, Value(0b0010))),
					Equal(Test.Number, Divide(Test.Number, Value(2))),
					Equal(Test.Number, Minus(Test.Number, Value(2))),
					Equal(Test.Number, Modulo(Test.Number, Value(2))),
					Equal(Test.Number, Multiply(Test.Number, Value(2))),
					Equal(Test.Number, Plus(Test.Number, Value(2))),
					Equal(Test.Number, ShiftLeft(Test.Number, Value(2))),
					Equal(Test.Number, ShiftRight(Test.Number, Value(2))),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`t`.`number` & ?", "BINARY BITWISE AND")
			assertContains(t, sqlSelectQuery, "`t`.`number` | ?", "BINARY BITWISE OR")
			assertContains(t, sqlSelectQuery, "`t`.`number` ^ ?", "BINARY BITWISE XOR")
			assertContains(t, sqlSelectQuery, "`t`.`number` / ?", "BINARY DIVIDE")
			assertContains(t, sqlSelectQuery, "`t`.`number` - ?", "BINARY MINUS")
			assertContains(t, sqlSelectQuery, "`t`.`number` % ?", "BINARY MODULO")
			assertContains(t, sqlSelectQuery, "`t`.`number` * ?", "BINARY MULTIPLY")
			assertContains(t, sqlSelectQuery, "`t`.`number` + ?", "BINARY PLUS")
			assertContains(t, sqlSelectQuery, "`t`.`number` << ?", "BINARY SHIFT LEFT")
			assertContains(t, sqlSelectQuery, "`t`.`number` >> ?", "BINARY SHIFT RIGHT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"t"."number" & $1`, "BINARY BITWISE AND")
			assertContains(t, sqlSelectQuery, `"t"."number" | $1`, "BINARY BITWISE OR")
			assertContains(t, sqlSelectQuery, `"t"."number" ^ $1`, "BINARY BITWISE XOR")
			assertContains(t, sqlSelectQuery, `"t"."number" / $1`, "BINARY DIVIDE")
			assertContains(t, sqlSelectQuery, `"t"."number" - $1`, "BINARY MINUS")
			assertContains(t, sqlSelectQuery, `"t"."number" % $1`, "BINARY MODULO")
			assertContains(t, sqlSelectQuery, `"t"."number" * $1`, "BINARY MULTIPLY")
			assertContains(t, sqlSelectQuery, `"t"."number" + $1`, "BINARY PLUS")
			assertContains(t, sqlSelectQuery, `"t"."number" << $1`, "BINARY SHIFT LEFT")
			assertContains(t, sqlSelectQuery, `"t"."number" >> $1`, "BINARY SHIFT RIGHT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_Column(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(Test.ID.As("id"))
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`t`.`id`", "COLUMN ID")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"t"."id"`, "COLUMN ID")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_Comparison(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(Test.ID.As("id")).
			Where(
				And(
					Between(Test.Number, Value(0), Value(2)),
					Equal(Test.Number, Value(2)),
					Exists(Subquery[int](NewSelect(Test.Table).Field(ConstIntOne()))),
					Greater(Test.Number, Value(2)),
					GreaterEqual(Test.Number, Value(2)),
					ILike(Test.String, Value("%ivan%")),
					In(Test.String, Array("active", "pending")),
					IsNotNull(Test.String),
					IsNull(Test.String),
					Less(Test.Number, Value(2)),
					LessEqual(Test.Number, Value(2)),
					Like(Test.String, Value("%ivan%")),
					NotBetween(Test.Number, Value(0), Value(2)),
					NotEqual(Test.Number, Value(2)),
					NotExists(Subquery[int](NewSelect(Test.Table).Field(ConstIntOne()))),
					NotILike(Test.String, Value("%ivan%")),
					NotIn(Test.String, Array("active", "pending")),
					NotLike(Test.String, Value("%ivan%")),
				))
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`t`.`number` BETWEEN ? AND ?", "COMPARISON BETWEEN")
			assertContains(t, sqlSelectQuery, "`t`.`number` = ?", "COMPARISON EQUAL")
			assertContains(t, sqlSelectQuery, "EXISTS (SELECT 1 FROM `test` AS `t`)", "COMPARISON EXISTS")
			assertContains(t, sqlSelectQuery, "`t`.`number` > ?", "COMPARISON GREATER")
			assertContains(t, sqlSelectQuery, "`t`.`number` >= ?", "COMPARISON GREATEREQUAL")
			assertContains(t, sqlSelectQuery, "LOWER(`t`.`string`) LIKE LOWER(?)", "COMPARISON ILIKE")
			assertContains(t, sqlSelectQuery, "`t`.`string` IN (?, ?)", "COMPARISON IN")
			assertContains(t, sqlSelectQuery, "`t`.`string` IS NOT NULL", "COMPARISON IS NOT NULL")
			assertContains(t, sqlSelectQuery, "`t`.`string` IS NULL", "COMPARISON IS NULL")
			assertContains(t, sqlSelectQuery, "`t`.`number` < ?", "COMPARISON LESS")
			assertContains(t, sqlSelectQuery, "`t`.`number` <= ? ", "COMPARISON LESSEQUAL")
			assertContains(t, sqlSelectQuery, "`t`.`string` LIKE ?", "COMPARISON LIKE")
			assertContains(t, sqlSelectQuery, "`t`.`number` NOT BETWEEN ? AND ?", "COMPARISON NOT BETWEEN")
			assertContains(t, sqlSelectQuery, "`t`.`number` <> ?", "COMPARISON NOT EQUAL")
			assertContains(t, sqlSelectQuery, "NOT EXISTS (SELECT 1 FROM `test` AS `t`)", "COMPARISON NOT EXISTS")
			assertContains(t, sqlSelectQuery, "LOWER(`t`.`string`) NOT LIKE LOWER(?)", "COMPARISON NOT ILIKE")
			assertContains(t, sqlSelectQuery, "`t`.`string` NOT IN (?, ?)", "COMPARISON NOT IN")
			assertContains(t, sqlSelectQuery, "`t`.`string` NOT LIKE ?", "COMPARISON NOT LIKE")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"t"."number" BETWEEN $1 AND $2`, "COMPARISON BETWEEN")
			assertContains(t, sqlSelectQuery, `"t"."number" = $1`, "COMPARISON EQUAL")
			assertContains(t, sqlSelectQuery, `EXISTS (SELECT 1 FROM "test" AS "t")`, "COMPARISON EXISTS")
			assertContains(t, sqlSelectQuery, `"t"."number" > $1`, "COMPARISON GREATER")
			assertContains(t, sqlSelectQuery, `"t"."number" >= $1`, "COMPARISON GREATEREQUAL")
			assertContains(t, sqlSelectQuery, `"t"."string" ILIKE $1`, "COMPARISON ILIKE")
			assertContains(t, sqlSelectQuery, `"t"."string" IN ($1, $2)`, "COMPARISON IN")
			assertContains(t, sqlSelectQuery, `"t"."string" IS NOT NULL`, "COMPARISON IS NOT NULL")
			assertContains(t, sqlSelectQuery, `"t"."string" IS NULL`, "COMPARISON IS NULL")
			assertContains(t, sqlSelectQuery, `"t"."number" < $1`, "COMPARISON LESS")
			assertContains(t, sqlSelectQuery, `"t"."number" <= $1`, "COMPARISON LESSEQUAL")
			assertContains(t, sqlSelectQuery, `"t"."string" LIKE $1`, "COMPARISON LIKE")
			assertContains(t, sqlSelectQuery, `"t"."number" NOT BETWEEN $1 AND $2`, "COMPARISON NOT BETWEEN")
			assertContains(t, sqlSelectQuery, `"t"."number" <> $1`, "COMPARISON NOT EQUAL")
			assertContains(t, sqlSelectQuery, `NOT EXISTS (SELECT 1 FROM "test" AS "t")`, "COMPARISON NOT EXISTS")
			assertContains(t, sqlSelectQuery, `"t"."string" NOT ILIKE $1`, "COMPARISON NOT ILIKE")
			assertContains(t, sqlSelectQuery, `"t"."string" NOT IN ($1, $2)`, "COMPARISON NOT IN")
			assertContains(t, sqlSelectQuery, `"t"."string" NOT LIKE $1`, "COMPARISON NOT LIKE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_Constant(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(Test.ID).
			Where(
				And(
					Equal(ConstBoolFalse(), ConstBoolFalse()),
					Equal(ConstBoolTrue(), ConstBoolTrue()),
					Equal(ConstFloat32One(), ConstFloat32One()),
					Equal(ConstFloat64One(), ConstFloat64One()),
					Equal(ConstIntOne(), ConstIntOne()),
					Equal(ConstInt8One(), ConstInt8One()),
					Equal(ConstInt16One(), ConstInt16One()),
					Equal(ConstInt32One(), ConstInt32One()),
					Equal(ConstInt64One(), ConstInt64One()),
					Equal(ConstStringDefault(), ConstStringDefault()),
					Equal(ConstStringNull(), ConstStringNull()),
					Equal(ConstUintOne(), ConstUintOne()),
					Equal(ConstUint8One(), ConstUint8One()),
					Equal(ConstUint16One(), ConstUint16One()),
					Equal(ConstUint32One(), ConstUint32One()),
					Equal(ConstUint64One(), ConstUint64One()),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "FALSE", "CONSTANT BOOlFALSE")
			assertContains(t, sqlSelectQuery, "TRUE", "CONSTANT BOOlTRUE")
			assertContains(t, sqlSelectQuery, "1.0", "CONSTANT FLOAT32ONE")
			assertContains(t, sqlSelectQuery, "1.000000", "CONSTANT FLOAT64ONE")
			assertContains(t, sqlSelectQuery, "1", "CONSTANT INTONE")
			assertContains(t, sqlSelectQuery, "1", "CONSTANT INT8ONE")
			assertContains(t, sqlSelectQuery, "1", "CONSTANT INT16ONE")
			assertContains(t, sqlSelectQuery, "1", "CONSTANT INT32ONE")
			assertContains(t, sqlSelectQuery, "1", "CONSTANT INT64ONE")
			assertContains(t, sqlSelectQuery, "DEFAULT", "CONSTANT STRINGDEFAULT")
			assertContains(t, sqlSelectQuery, "NULL", "CONSTANT NULLSTRINGNULL")
			assertContains(t, sqlSelectQuery, "1", "CONSTANT UINTONE")
			assertContains(t, sqlSelectQuery, "1", "CONSTANT UINT8ONE")
			assertContains(t, sqlSelectQuery, "1", "CONSTANT UINT16ONE")
			assertContains(t, sqlSelectQuery, "1", "CONSTANT UINT32ONE")
			assertContains(t, sqlSelectQuery, "1", "CONSTANT UINT64ONE")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `FALSE`, "CONSTANT BOOlFALSE")
			assertContains(t, sqlSelectQuery, `TRUE`, "CONSTANT BOOlTRUE")
			assertContains(t, sqlSelectQuery, `1.0`, "CONSTANT FLOAT32ONE")
			assertContains(t, sqlSelectQuery, `1.000000`, "CONSTANT FLOAT64ONE")
			assertContains(t, sqlSelectQuery, `1`, "CONSTANT INTONE")
			assertContains(t, sqlSelectQuery, `1`, "CONSTANT INT8ONE")
			assertContains(t, sqlSelectQuery, `1`, "CONSTANT INT16ONE")
			assertContains(t, sqlSelectQuery, `1`, "CONSTANT INT32ONE")
			assertContains(t, sqlSelectQuery, `1`, "CONSTANT INT64ONE")
			assertContains(t, sqlSelectQuery, `DEFAULT`, "CONSTANT STRINGDEFAULT")
			assertContains(t, sqlSelectQuery, `NULL`, "CONSTANT NULLSTRINGNULL")
			assertContains(t, sqlSelectQuery, `1`, "CONSTANT UINTONE")
			assertContains(t, sqlSelectQuery, `1`, "CONSTANT UINT8ONE")
			assertContains(t, sqlSelectQuery, `1`, "CONSTANT UINT16ONE")
			assertContains(t, sqlSelectQuery, `1`, "CONSTANT UINT32ONE")
			assertContains(t, sqlSelectQuery, `1`, "CONSTANT UINT64ONE")

		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_Function(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				// Функции агрегатные
				Avg(Test.Number, false).As("aggregate_avg"),
				BitAnd(Test.Number, false).As("aggregate_bitand"),
				BitOr(Test.Number, false).As("aggregate_bitor"),
				BitXor(Test.Number, false).As("aggregate_bitxor"),
				Count(Test.String, false).As("aggregate_count"),
				GroupConcat(Test.String, false).As("aggregate_groupconcat"),
				Max(Test.Number, false).As("aggregate_max"),
				Min(Test.Number, false).As("aggregate_min"),
				StdDev(Test.Number, false).As("aggregate_stddev"),
				Sum(Test.Number, false).As("aggregate_sum"),
				Variance(Test.Number, false).As("aggregate_variance"),
				// Функции аналитические
				FirstValue(Test.Name).Over(
					PartitionBy(Test.ID),
					OrderBy(Desc(Test.Number)),
				).As("analytical_firstvalue"),
				Lag(Test.Number, 2).Over(
					PartitionBy(Test.ID),
					OrderBy(Asc(Test.Date)),
				).As("analytical_lag"),
				LastValue(Test.Name).Over(
					PartitionBy(Test.ID),
					OrderBy(Asc(Test.Number)),
					RowsBetween("CURRENT ROW", "UNBOUNDED FOLLOWING"),
				).As("analytical_lastvalue"),
				Lead(Test.Number, 2).Over(
					PartitionBy(Test.ID),
					OrderBy(Asc(Test.Date)),
				).As("analytical_lead"),
				NthValue(Test.Name, 2).Over(
					PartitionBy(Test.ID),
					OrderBy(Desc(Test.Number)),
					RowsBetween("UNBOUNDED PRECEDING", "CURRENT ROW"),
				).As("analytical_nthvalue"),
				// Функции условий
				Case(CaseIf(CasePair(Less(Test.Number, Value(2)), Value("old"))), CaseElse(Value("new"))).As("condition_case"),
				Coalesce(Test.CreateAt, Test.UpdateAt).As("condition_coalesce"),
				Greatest(Test.CreateAt, Test.UpdateAt).As("condition_greatest"),
				Least(Test.CreateAt, Test.UpdateAt).As("condition_least"),
				NullIf(Test.CreateAt, Test.UpdateAt).As("condition_if"),
				// Функции конвертации
				Cast(Test.Number, TypeString).As("convert_cast"),
				CharLength(Test.String).As("convert_length"),
				DateFormat(Test.CreateAt, Literal("%Y-%m-%d")).As("convert_dateformat"),
				Degrees(Test.Number).As("convert_degrees"),
				Length(Test.String).As("convert_length"),
				Position(Test.String, Value("old")).As("convert_position"),
				Radians(Test.Number).As("convert_radians"),
				// Функции даты и времени
				CurDate().As("datetime_curdate"),
				CurTime().As("datetime_curtime"),
				DateAdd(Test.CreateAt, Literal("2 DAY")).As("datetime_dateadd"),
				DateDiff(Test.UpdateAt, Test.CreateAt).As("datetime_datediff"),
				DateSub(Test.CreateAt, Literal("2 DAY")).As("datetime_datesub"),
				Day(Test.CreateAt).As("datetime_day"),
				DayName(Test.CreateAt).As("datetime_dayname"),
				Hour(Test.CreateAt).As("datetime_hour"),
				Minute(Test.CreateAt).As("datetime_minute"),
				Month(Test.CreateAt).As("datetime_month"),
				MonthName(Test.CreateAt).As("datetime_monthname"),
				Now().As("datetime_now"),
				Quarter(Test.CreateAt).As("datetime_quarter"),
				Second(Test.CreateAt).As("datetime_second"),
				TimeAdd(Test.CreateAt, Literal("2 HOUR")).As("datetime_timeadd"),
				TimeDiff(Test.UpdateAt, Test.CreateAt).As("datetime_timediff"),
				TimeSub(Test.CreateAt, Literal("2 HOUR")).As("datetime_timesub"),
				Week(Test.CreateAt).As("datetime_week"),
				Year(Test.CreateAt).As("datetime_year"),
				// Функции обмена данными
				JsonArray(Test.Json, Value("val1"), Value("val2")).As("json_jsonarray"),
				JsonArrayAgg(Test.Json).As("json_jsonarrayagg"),
				JsonContains(Test.Json, Value(`{"key":"val"}`)).As("json_jsoncontains"),
				JsonExtract(Test.Json, JsonGroup(JsonPath(JsonKey("parent"), JsonIndex(0), JsonKey("child"))), TypeString).As("json_jsonextract"),
				JsonObject(JsonPair(JsonKey("key"), Count(Test.Json, false))).As("json_jsonobject"),
				JsonObjectAgg(Test.Json, Test.Number).As("json_jsonobjectagg"),
				JsonRemove(Test.Json, JsonGroup(JsonPath(JsonKey("key1"))), JsonGroup(JsonPath(JsonKey("key2")))).As("json_jsonremove"),
				JsonSet(Test.Json, JsonGroup(JsonPath(JsonKey("key1")), Value("val1")), JsonGroup(JsonPath(JsonKey("key2")), Value("val2"))).As("json_jsonset"),
				JsonType(Test.Json).As("json_jsontype"),
				// Функции математические
				Abs(Test.Number).As("math_abs"),
				ACos(Test.Number).As("math_acos"),
				ASin(Test.Number).As("math_asin"),
				ATan(Test.Number).As("math_atan"),
				ATan2(Test.Y, Test.X).As("math_atan2"),
				Cbrt(Test.Number).As("math_cbrt"),
				Ceil(Test.Number).As("math_ceil"),
				Cos(Test.Number).As("math_cos"),
				Exp(Test.Number).As("math_exp"),
				Floor(Test.Number).As("math_floor"),
				Ln(Test.Number).As("math_ln"),
				Log(Test.Number, Value(2)).As("math_log"),
				Mod(Test.Number, Value(2)).As("math_mod"),
				Pi().As("math_pi"),
				Power(Test.Number, Value(2)).As("math_power"),
				Rand().As("math_rand"),
				Round(Test.Number, Value(2)).As("math_round"),
				Sin(Test.Number).As("math_sin"),
				Sqrt(Test.Number).As("math_sqrt"),
				Tan(Test.Number).As("math_tan"),
				Trunc(Test.Number, Value(2)).As("math_trunc"),
				// Функции ранжирующие
				CumeDist().Over(
					PartitionBy(Test.ID),
					OrderBy(Desc(Test.Number)),
				).As("ranking_cumedist"),
				DenseRank().Over(
					PartitionBy(Test.ID),
					OrderBy(Desc(Test.Number)),
				).As("ranking_denserank"),
				NTile(2).Over(
					PartitionBy(Test.ID),
					OrderBy(Desc(Test.Number)),
				).As("ranking_ntile"),
				PercentRank().Over(
					PartitionBy(Test.ID),
					OrderBy(Desc(Test.Number)),
				).As("ranking_percentrank"),
				Rank().Over(
					PartitionBy(Test.ID),
					OrderBy(Desc(Test.Number)),
				).As("ranking_rank"),
				RowNumber().Over(
					PartitionBy(Test.ID),
					OrderBy(Desc(Test.Number)),
				).As("ranking_rownumber"),
				// Функции строковые
				Concat(Test.String, Value("old"), Value("new")).As("string_concat"),
				ConcatWs(Value("_"), Test.String, Value("old"), Value("new")).As("string_concatws"),
				LeftString(Test.String, Value(2)).As("string_lstr"),
				Lower(Test.String).As("string_lower"),
				LPad(Test.String, Value(2), Value(",")).As("string_lpad"),
				LTrim(Test.String).As("string_ltrim"),
				Repeat(Test.String, Value(2)).As("string_repeat"),
				Replace(Test.String, Value("old"), Value("new")).As("string_replace"),
				Reverse(Test.String).As("string_reverse"),
				RightString(Test.String, Value(2)).As("string_rstr"),
				RPad(Test.String, Value(2), Value(",")).As("string_rpad"),
				RTrim(Test.String).As("string_rtrim"),
				SubString(Test.String, Value(0), Value(2)).As("string_substring"),
				Trim(Test.String).As("string_trim"),
				Upper(Test.String).As("string_upper"),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			// Функции агрегатные
			assertContains(t, sqlSelectQuery, "AVG(`t`.`number`)", "FUNCTION AVG")
			assertContains(t, sqlSelectQuery, "BIT_AND(`t`.`number`)", "FUNCTION BITAND")
			assertContains(t, sqlSelectQuery, "BIT_OR(`t`.`number`)", "FUNCTION BITOR")
			assertContains(t, sqlSelectQuery, "BIT_XOR(`t`.`number`)", "FUNCTION BITXOR")
			assertContains(t, sqlSelectQuery, "COUNT(`t`.`string`)", "FUNCTION COUNT")
			assertContains(t, sqlSelectQuery, "GROUP_CONCAT(`t`.`string` SEPARATOR ',')", "FUNCTION GROUPCONCAT")
			assertContains(t, sqlSelectQuery, "MAX(`t`.`number`)", "FUNCTION MAX")
			assertContains(t, sqlSelectQuery, "MIN(`t`.`number`)", "FUNCTION MIN")
			assertContains(t, sqlSelectQuery, "STDDEV(`t`.`number`)", "FUNCTION STDDEV")
			assertContains(t, sqlSelectQuery, "SUM(`t`.`number`)", "FUNCTION SUM")
			assertContains(t, sqlSelectQuery, "VARIANCE(`t`.`number`)", "FUNCTION VARIANCE")
			// Функции аналитические
			assertContains(t, sqlSelectQuery, "FIRST_VALUE(`t`.`name`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC", "FUNCTION FIRSTVALUE")
			assertContains(t, sqlSelectQuery, "LAG(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)", "FUNCTION LAG")
			assertContains(t, sqlSelectQuery, "LAST_VALUE(`t`.`name`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)", "FUNCTION LASTVALUE")
			assertContains(t, sqlSelectQuery, "LEAD(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)", "FUNCTION LEAD")
			assertContains(t, sqlSelectQuery, "NTH_VALUE(`t`.`name`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)", "FUNCTION NTHVALUE")
			// Функции условий
			assertContains(t, sqlSelectQuery, "CASE WHEN `t`.`number` < ? THEN ? ELSE ? END", "FUNCTION CASE")
			assertContains(t, sqlSelectQuery, "COALESCE(`t`.`createat`, `t`.`updateat`)", "FUNCTION COALESCE")
			assertContains(t, sqlSelectQuery, "GREATEST(`t`.`createat`, `t`.`updateat`)", "FUNCTION GREATEST")
			assertContains(t, sqlSelectQuery, "LEAST(`t`.`createat`, `t`.`updateat`)", "FUNCTION LEAST")
			assertContains(t, sqlSelectQuery, "NULLIF(`t`.`createat`, `t`.`updateat`)", "FUNCTION NULLIF")
			// Функции конвертации
			assertContains(t, sqlSelectQuery, "CAST(`t`.`number` AS VARCHAR)", "FUNCTION CAST")
			assertContains(t, sqlSelectQuery, "CHAR_LENGTH(`t`.`string`)", "FUNCTION CHARLENGTH")
			assertContains(t, sqlSelectQuery, "DATE_FORMAT(`t`.`createat`, '%Y-%m-%d')", "FUNCTION DATEFORMAT")
			assertContains(t, sqlSelectQuery, "DEGREES(`t`.`number`)", "FUNCTION DEGREES")
			assertContains(t, sqlSelectQuery, "LENGTH(`t`.`string`)", "FUNCTION LENGTH")
			assertContains(t, sqlSelectQuery, "POSITION(? IN `t`.`string`)", "FUNCTION POSITION")
			assertContains(t, sqlSelectQuery, "RADIANS(`t`.`number`)", "FUNCTION RADIANS")
			// Функции даты и времени
			assertContains(t, sqlSelectQuery, "CURDATE()", "FUNCTION CURDATE")
			assertContains(t, sqlSelectQuery, "CURTIME()", "FUNCTION CURTIME")
			assertContains(t, sqlSelectQuery, "DATE_ADD(`t`.`createat`, INTERVAL '2 DAY')", "FUNCTION DATEADD")
			assertContains(t, sqlSelectQuery, "DATEDIFF(`t`.`updateat`, `t`.`createat`)", "FUNCTION DATEDIFF")
			assertContains(t, sqlSelectQuery, "DATE_SUB(`t`.`createat`, INTERVAL '2 DAY')", "FUNCTION DATESUB")
			assertContains(t, sqlSelectQuery, "DAY(`t`.`createat`)", "FUNCTION DAY")
			assertContains(t, sqlSelectQuery, "DAYNAME(`t`.`createat`)", "FUNCTION DAYNAME")
			assertContains(t, sqlSelectQuery, "HOUR(`t`.`createat`)", "FUNCTION HOUR")
			assertContains(t, sqlSelectQuery, "MINUTE(`t`.`createat`)", "FUNCTION MINUTE")
			assertContains(t, sqlSelectQuery, "MONTH(`t`.`createat`)", "FUNCTION MONTH")
			assertContains(t, sqlSelectQuery, "MONTHNAME(`t`.`createat`)", "FUNCTION MONTHNAME")
			assertContains(t, sqlSelectQuery, "NOW()", "FUNCTION NOW")
			assertContains(t, sqlSelectQuery, "QUARTER(`t`.`createat`)", "FUNCTION QUARTER")
			assertContains(t, sqlSelectQuery, "SECOND(`t`.`createat`)", "FUNCTION SECOND")
			assertContains(t, sqlSelectQuery, "TIME_ADD(`t`.`createat`, INTERVAL '2 HOUR')", "FUNCTION TIMEADD")
			assertContains(t, sqlSelectQuery, "TIMEDIFF(`t`.`updateat`, `t`.`createat`)", "FUNCTION TIMEDIFF")
			assertContains(t, sqlSelectQuery, "TIME_SUB(`t`.`createat`, INTERVAL '2 HOUR')", "FUNCTION TIMESUB")
			assertContains(t, sqlSelectQuery, "WEEK(`t`.`createat`)", "FUNCTION WEEK")
			assertContains(t, sqlSelectQuery, "YEAR(`t`.`createat`)", "FUNCTION YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, "JSON_ARRAY(`t`.`json`, ?, ?)", "FUNCTION JSONARRAY")
			assertContains(t, sqlSelectQuery, "JSON_ARRAYAGG(`t`.`json`)", "FUNCTION JSONARRAYAGG")
			assertContains(t, sqlSelectQuery, "JSON_CONTAINS(`t`.`json`, ?)", "FUNCTION JSONCONTAINS")
			assertContains(t, sqlSelectQuery, "(`t`.`json` ->> '$.parent[0].child')", "FUNCTION JSONEXTRACT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECT('key', COUNT(`t`.`json`))", "FUNCTION JSONOBJECT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECTAGG(`t`.`json`, `t`.`number`)", "FUNCTION JSONOBJECTAGG")
			assertContains(t, sqlSelectQuery, "JSON_REMOVE(`t`.`json`, '$.key1', '$.key2')", "FUNCTION JSONREMOVE")
			assertContains(t, sqlSelectQuery, "JSON_SET(`t`.`json`, '$.key1', ?, '$.key2', ?)", "FUNCTION JSONSET")
			assertContains(t, sqlSelectQuery, "JSON_TYPE(`t`.`json`)", "FUNCTION JSONTYPE")
			// Функции математические
			assertContains(t, sqlSelectQuery, "ABS(`t`.`number`)", "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, "ACOS(`t`.`number`)", "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, "ASIN(`t`.`number`)", "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, "ATAN(`t`.`number`)", "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, "ATAN2(`t`.`y`, `t`.`x`)", "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, "CBRT(`t`.`number`)", "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, "CEILING(`t`.`number`)", "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, "COS(`t`.`number`)", "FUNCTION COS")
			assertContains(t, sqlSelectQuery, "EXP(`t`.`number`)", "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, "FLOOR(`t`.`number`)", "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, "LN(`t`.`number`)", "FUNCTION LN")
			assertContains(t, sqlSelectQuery, "LOG(`t`.`number`, ?)", "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, "MOD(`t`.`number`, ?)", "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, "PI()", "FUNCTION PI")
			assertContains(t, sqlSelectQuery, "POWER(`t`.`number`, ?)", "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, "RAND()", "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, "ROUND(`t`.`number`, ?)", "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, "SIN(`t`.`number`)", "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, "SQRT(`t`.`number`)", "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, "TAN(`t`.`number`)", "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, "TRUNCATE(`t`.`number`, ?)", "FUNCTION TRUNC")
			// Функции ранжирующие
			assertContains(t, sqlSelectQuery, "CUME_DIST() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)", "FUNCTION CUMEDIST")
			assertContains(t, sqlSelectQuery, "DENSE_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)", "FUNCTION DENSERANK")
			assertContains(t, sqlSelectQuery, "NTILE(2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)", "FUNCTION NTILE")
			assertContains(t, sqlSelectQuery, "PERCENT_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)", "FUNCTION PERCENTRANK")
			assertContains(t, sqlSelectQuery, "RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)", "FUNCTION RANK")
			assertContains(t, sqlSelectQuery, "ROW_NUMBER() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)", "FUNCTION ROWNUMBER")
			// Функции строковые
			assertContains(t, sqlSelectQuery, "CONCAT(`t`.`string`, ?, ?)", "FUNCTION CONCAT")
			assertContains(t, sqlSelectQuery, "CONCAT_WS(?, `t`.`string`, ?, ?)", "FUNCTION CONCATWS")
			assertContains(t, sqlSelectQuery, "LEFT(`t`.`string`, ?)", "FUNCTION LEFTSTRING")
			assertContains(t, sqlSelectQuery, "LOWER(`t`.`string`)", "FUNCTION LOWER")
			assertContains(t, sqlSelectQuery, "LPAD(`t`.`string`, ?, ?)", "FUNCTION LPAD")
			assertContains(t, sqlSelectQuery, "LTRIM(`t`.`string`)", "FUNCTION LTRIM")
			assertContains(t, sqlSelectQuery, "REPEAT(`t`.`string`, ?)", "FUNCTION REPEAT")
			assertContains(t, sqlSelectQuery, "REPLACE(`t`.`string`, ?, ?)", "FUNCTION REPLACE")
			assertContains(t, sqlSelectQuery, "REVERSE(`t`.`string`)", "FUNCTION REVERSE")
			assertContains(t, sqlSelectQuery, "RIGHT(`t`.`string`, ?)", "FUNCTION RIGHTSTRING")
			assertContains(t, sqlSelectQuery, "RPAD(`t`.`string`, ?, ?)", "FUNCTION RPAD")
			assertContains(t, sqlSelectQuery, "RTRIM(`t`.`string`)", "FUNCTION RTRIM")
			assertContains(t, sqlSelectQuery, "SUBSTRING(`t`.`string`, ?, ?)", "FUNCTION SUBSTRING")
			assertContains(t, sqlSelectQuery, "TRIM(`t`.`string`)", "FUNCTION TRIM")
			assertContains(t, sqlSelectQuery, "UPPER(`t`.`string`)", "FUNCTION UPPER")
		case DialectPostgreSQL:
			// Функции агрегатные
			assertContains(t, sqlSelectQuery, `AVG("t"."number")`, "FUNCTION AVG")
			assertContains(t, sqlSelectQuery, `BIT_AND("t"."number")`, "FUNCTION BITAND")
			assertContains(t, sqlSelectQuery, `BIT_OR("t"."number")`, "FUNCTION BITOR")
			assertContains(t, sqlSelectQuery, `BIT_XOR("t"."number")`, "FUNCTION BITXOR")
			assertContains(t, sqlSelectQuery, `COUNT("t"."string")`, "FUNCTION COUNT")
			assertContains(t, sqlSelectQuery, `STRING_AGG("t"."string", ',')`, "FUNCTION GROUPCONCAT")
			assertContains(t, sqlSelectQuery, `MAX("t"."number")`, "FUNCTION MAX")
			assertContains(t, sqlSelectQuery, `MIN("t"."number")`, "FUNCTION MIN")
			assertContains(t, sqlSelectQuery, `STDDEV_SAMP("t"."number")`, "FUNCTION STDDEV")
			assertContains(t, sqlSelectQuery, `SUM("t"."number")`, "FUNCTION SUM")
			assertContains(t, sqlSelectQuery, `VAR_SAMP("t"."number")`, "FUNCTION VARIANCE")
			// Функции аналитические
			assertContains(t, sqlSelectQuery, `FIRST_VALUE("t"."name") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION FIRSTVALUE")
			assertContains(t, sqlSelectQuery, `LAG("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)`, "FUNCTION LAG")
			assertContains(t, sqlSelectQuery, `LAST_VALUE("t"."name") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)`, "FUNCTION LASTVALUE")
			assertContains(t, sqlSelectQuery, `LEAD("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)`, "FUNCTION LEAD")
			assertContains(t, sqlSelectQuery, `NTH_VALUE("t"."name", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)`, "FUNCTION NTHVALUE")
			// Функции условий
			assertContains(t, sqlSelectQuery, `CASE WHEN "t"."number" < $1 THEN $2 ELSE $3 END`, "FUNCTION CASE")
			assertContains(t, sqlSelectQuery, `COALESCE("t"."createat", "t"."updateat")`, "FUNCTION COALESCE")
			assertContains(t, sqlSelectQuery, `GREATEST("t"."createat", "t"."updateat")`, "FUNCTION GREATEST")
			assertContains(t, sqlSelectQuery, `LEAST("t"."createat", "t"."updateat")`, "FUNCTION LEAST")
			assertContains(t, sqlSelectQuery, `NULLIF("t"."createat", "t"."updateat")`, "FUNCTION NULLIF")
			// Функции конвертации
			assertContains(t, sqlSelectQuery, `CAST("t"."number" AS VARCHAR) `, "FUNCTION CAST")
			assertContains(t, sqlSelectQuery, `CHAR_LENGTH("t"."string")`, "FUNCTION CHARLENGTH")
			assertContains(t, sqlSelectQuery, `TO_CHAR("t"."createat", '%Y-%m-%d')`, "FUNCTION DATEFORMAT")
			assertContains(t, sqlSelectQuery, `DEGREES("t"."number")`, "FUNCTION DEGREES")
			assertContains(t, sqlSelectQuery, `LENGTH("t"."string")`, "FUNCTION LENGTH")
			assertContains(t, sqlSelectQuery, `POSITION($1 IN "t"."string")`, "FUNCTION POSITION")
			assertContains(t, sqlSelectQuery, `RADIANS("t"."number")`, "FUNCTION RADIANS")
			// Функции даты и времени
			assertContains(t, sqlSelectQuery, `CURRENT_DATE`, "FUNCTION CURDATE")
			assertContains(t, sqlSelectQuery, `CURRENT_TIME`, "FUNCTION CURTIME")
			assertContains(t, sqlSelectQuery, `("t"."createat" + INTERVAL '2 DAY')`, "FUNCTION DATEADD")
			assertContains(t, sqlSelectQuery, `DATE_PART('day', "t"."updateat" - "t"."createat")`, "FUNCTION DATEDIFF")
			assertContains(t, sqlSelectQuery, `("t"."createat" - INTERVAL '2 DAY')`, "FUNCTION DATESUB")
			assertContains(t, sqlSelectQuery, `EXTRACT(DAY FROM "t"."createat")`, "FUNCTION DAY")
			assertContains(t, sqlSelectQuery, `TO_CHAR("t"."createat", 'Day')`, "FUNCTION DAYNAME")
			assertContains(t, sqlSelectQuery, `EXTRACT(HOUR FROM "t"."createat")`, "FUNCTION HOUR")
			assertContains(t, sqlSelectQuery, `EXTRACT(MINUTE FROM "t"."createat")`, "FUNCTION MINUTE")
			assertContains(t, sqlSelectQuery, `EXTRACT(MONTH FROM "t"."createat")`, "FUNCTION MONTH")
			assertContains(t, sqlSelectQuery, `TO_CHAR("t"."createat", 'Month')`, "FUNCTION MONTHNAME")
			assertContains(t, sqlSelectQuery, `CURRENT_TIMESTAMP`, "FUNCTION NOW")
			assertContains(t, sqlSelectQuery, `EXTRACT(QUARTER FROM "t"."createat")`, "FUNCTION QUARTER")
			assertContains(t, sqlSelectQuery, `EXTRACT(SECOND FROM "t"."createat")`, "FUNCTION SECOND")
			assertContains(t, sqlSelectQuery, `("t"."createat" + INTERVAL '2 HOUR')`, "FUNCTION TIMEADD")
			assertContains(t, sqlSelectQuery, `DATE_PART('time', "t"."updateat" - "t"."createat")`, "FUNCTION TIMEDIFF")
			assertContains(t, sqlSelectQuery, `("t"."createat" - INTERVAL '2 HOUR')`, "FUNCTION TIMESUB")
			assertContains(t, sqlSelectQuery, `EXTRACT(WEEK FROM "t"."createat")`, "FUNCTION WEEK")
			assertContains(t, sqlSelectQuery, `EXTRACT(YEAR FROM "t"."createat")`, "FUNCTION YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, `JSON_ARRAY("t"."json", $1, $2)`, "FUNCTION JSONARRAY")
			assertContains(t, sqlSelectQuery, `JSON_AGG("t"."json")`, "FUNCTION JSONARRAYAGG")
			assertContains(t, sqlSelectQuery, `("t"."json" @> $1)`, "FUNCTION JSONCONTAINS")
			assertContains(t, sqlSelectQuery, `("t"."json" #>> '{parent,0,child}')`, "FUNCTION JSONEXTRACT")
			assertContains(t, sqlSelectQuery, `JSON_BUILD_OBJECT('key', COUNT("t"."json"))`, "FUNCTION JSONOBJECT")
			assertContains(t, sqlSelectQuery, `JSON_OBJECT_AGG("t"."json", "t"."number")`, "FUNCTION JSONOBJECTAGG")
			assertContains(t, sqlSelectQuery, `("t"."json" - '{key1}' - '{key2}')`, "FUNCTION JSONREMOVE")
			//assertContains(t, sqlSelectQuery, `jsonb_set`, "FUNCTION JSONSET")
			assertContains(t, sqlSelectQuery, `jsonb_typeof("t"."json")`, "FUNCTION JSONTYPE")
			// Функции математические
			assertContains(t, sqlSelectQuery, `ABS("t"."number")`, "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, `ACOS("t"."number")`, "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, `ASIN("t"."number")`, "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, `ATAN("t"."number")`, "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, `ATAN2("t"."y", "t"."x")`, "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, `CBRT("t"."number")`, "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, `CEIL("t"."number")`, "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, `COS("t"."number")`, "FUNCTION COS")
			assertContains(t, sqlSelectQuery, `EXP("t"."number")`, "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, `FLOOR("t"."number")`, "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, `LN("t"."number")`, "FUNCTION LN")
			assertContains(t, sqlSelectQuery, `LOG("t"."number", $1)`, "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, `MOD("t"."number", $1)`, "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, `PI()`, "FUNCTION PI")
			assertContains(t, sqlSelectQuery, `POWER("t"."number", $1)`, "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, `RANDOM`, "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, `ROUND("t"."number", $1)`, "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, `SIN("t"."number")`, "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, `SQRT("t"."number")`, "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, `TAN("t"."number")`, "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, `TRUNC("t"."number", $1)`, "FUNCTION TRUNC")
			// Функции ранжирующие
			assertContains(t, sqlSelectQuery, `CUME_DIST() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION CUMEDIST")
			assertContains(t, sqlSelectQuery, `DENSE_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "DFUNCTION ENSERANK")
			assertContains(t, sqlSelectQuery, `NTILE(2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION NTILE")
			assertContains(t, sqlSelectQuery, `PERCENT_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION PERCENTRANK")
			assertContains(t, sqlSelectQuery, `RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION RANK")
			assertContains(t, sqlSelectQuery, `ROW_NUMBER() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION ROWNUMBER")
			// Функции строковые
			assertContains(t, sqlSelectQuery, `CONCAT("t"."string", $1, $2)`, "FUNCTION CONCAT")
			assertContains(t, sqlSelectQuery, `CONCAT_WS($1, "t"."string", $2, $3)`, "FUNCTION CONCATWS")
			assertContains(t, sqlSelectQuery, `LEFT("t"."string", $1)`, "FUNCTION LEFTSTRING")
			assertContains(t, sqlSelectQuery, `LOWER("t"."string")`, "FUNCTION LOWER")
			assertContains(t, sqlSelectQuery, `LPAD("t"."string", $1, $2)`, "FUNCTION LPAD")
			assertContains(t, sqlSelectQuery, `LTRIM("t"."string")`, "FUNCTION LTRIM")
			assertContains(t, sqlSelectQuery, `REPEAT("t"."string", $1)`, "FUNCTION REPEAT")
			assertContains(t, sqlSelectQuery, `REPLACE("t"."string", $1, $2)`, "FUNCTION REPLACE")
			assertContains(t, sqlSelectQuery, `REVERSE("t"."string")`, "FUNCTION REVERSE")
			assertContains(t, sqlSelectQuery, `RIGHT("t"."string", $1)`, "FUNCTION RIGHTSTRING")
			assertContains(t, sqlSelectQuery, `RPAD("t"."string", $1, $2)`, "FUNCTION RPAD")
			assertContains(t, sqlSelectQuery, `RTRIM("t"."string")`, "FUNCTION RTRIM")
			assertContains(t, sqlSelectQuery, `SUBSTRING("t"."string", $1, $2)`, "FUNCTION SUBSTRING")
			assertContains(t, sqlSelectQuery, `TRIM("t"."string")`, "FUNCTION TRIM")
			assertContains(t, sqlSelectQuery, `UPPER("t"."string")`, "FUNCTION UPPER")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_Literal(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(Test.ID).
			Where(
				Equal(DateFormat(Test.CreateAt, Literal("%Y-%m-%d")), Value("2026-01-01")),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "'%Y-%m-%d'", "LITERAL STRING")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `'%Y-%m-%d'`, "LITERAL STRING")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_Logical(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(Test.ID.As("id")).
			Where(
				And(
					And(
						Equal(Test.String, Value("active")),
						Greater(Test.Number, Value(2)),
					),
					Or(
						Equal(Test.String, Value("active")),
						Greater(Test.Number, Value(2)),
					),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`t`.`string` = ? AND `t`.`number` > ?", "LOGICAL AND")
			assertContains(t, sqlSelectQuery, "`t`.`string` = ? OR `t`.`number` > ?", "LOGICAL OR")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"t"."string" = $1 AND "t"."number" > $2`, "LOGICAL AND")
			assertContains(t, sqlSelectQuery, `"t"."string" = $1 OR "t"."number" > $2`, "LOGICAL OR")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_Order(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(Test.ID.As("id")).
			OrderBy(
				Asc(Test.String),
				Desc(Test.String),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`t`.`string` ASC", "ORDER ASC")
			assertContains(t, sqlSelectQuery, "`t`.`string` DESC", "ORDER DESC")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"t"."string" ASC`, "ORDER ASC")
			assertContains(t, sqlSelectQuery, `"t"."string" DESC`, "ORDER DESC")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_Subquery(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(Subquery[int64](NewSelect(Test.Table).Field(Test.ID)).As("SUB"))
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "(SELECT `t`.`id` FROM `test` AS `t`)", "SUBQUERT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `(SELECT "t"."id" FROM "test" AS "t")`, "SUBQUERT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_Value(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		var data string = "ivan"
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(Test.ID.As("id")).
			Where(
				Equal(Test.String, Value(data)),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`t`.`string` = ?", "VALUE PLACEHOLDER")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"t"."string" = $1`, "VALUE PLACEHOLDER")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Delete(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtDelete := NewDelete(Test.Table).
			Where(
				Equal(Test.String, Value("active")),
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "DELETE", "DELETE")
			assertContains(t, sqlDeleteQuery, "FROM", "FROM")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `DELETE`, "DELETE")
			assertContains(t, sqlDeleteQuery, `FROM`, "FROM")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_Delete_Join(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtDelete := NewDelete(Test.Table).
			Join(
				Inner(Test1.Table, Equal(Test1.ID, Test.ID)),
				Left(Test2.Table, Equal(Test2.ID, Test.ID)),
			).
			Where(
				Equal(Test.String, Value("active")),
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "JOIN", "JOIN")
			assertContains(t, sqlDeleteQuery, "INNER JOIN", "INNER JOIN")
			assertContains(t, sqlDeleteQuery, "LEFT JOIN", "LEFT JOIN")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `USING`, "USING")
			assertContains(t, sqlDeleteQuery, `"test1" AS "t1", "test2" AS "t2"`, "LIST")
			assertContains(t, sqlDeleteQuery, `"t1"."id" = "t"."id" AND "t2"."id" = "t"."id"`, "CONDITION")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_Delete_Returning(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtDelete := NewDelete(Test.Table).
			Where(
				Equal(Test.String, Value("active")),
			).
			Returning(
				Test.ID,
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			// Not support
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `RETURNING "t"."id"`, "RETURNING")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_Delete_Where(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtDelete := NewDelete(Test.Table).
			Where(
				Equal(Test.String, Value("active")),
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "`t`.`string` = ?", "EQUAL")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `"t"."string" = $1`, "EQUAL")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_Delete_With(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		cte := WithN("old_data", NewSelect(Test.Table).
			Field(Test.ID).
			Where(
				Equal(Test.String, Value("old")),
			),
		)
		stmtDelete := NewDelete(Test.Table).
			Where(
				In(Test.ID, Subquery[int64](NewSelect(Test.Table).Field(Column[int64]("old_data", "id")))),
			).
			With(cte)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "WITH", "WITH")
			assertContains(t, sqlDeleteQuery, "old_data", "CTE")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `WITH`, "WITH")
			assertContains(t, sqlDeleteQuery, `old_data`, "CTE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_Insert(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtInsert := NewInsert(Test.Table).
			Column(Test.String, Test.Number).
			Values(
				Row(
					Value("ivan"),
					Value(2),
				),
			)
		sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlInsertQuery, "INSERT", "INSERT")
			assertContains(t, sqlInsertQuery, "INTO", "INTO")
		case DialectPostgreSQL:
			assertContains(t, sqlInsertQuery, `INSERT`, "INSERT")
			assertContains(t, sqlInsertQuery, `INTO`, "INTO")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
	})
}
func Test_Select(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(Test.String).
			Where(
				Equal(Test.String, Value("active")),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "SELECT", "SELECT")
			assertContains(t, sqlSelectQuery, "FROM", "FROM")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `SELECT`, "SELECT")
			assertContains(t, sqlSelectQuery, `FROM`, "FROM")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Distinct(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).Distinct().
			Field(Test.ID.As("id"))
		sqlSelectQuery, _, _ := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "DISTINCT", "DISTINCT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `DISTINCT`, "DISTINCT")
		}
		t.Logf("dialect: %s sql: %s", supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_GroupBy(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.String,
				Count(Test.ID, false).As("count"),
			).
			GroupBy(
				Test.String,
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "GROUP BY", "GROUP BY")
			assertContains(t, sqlSelectQuery, "`t`.`string`", "LIST")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `GROUP BY`, "GROUP BY")
			assertContains(t, sqlSelectQuery, `"t"."string"`, "LIST")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Having(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		queryMain := NewSelect(Test.Table).
			Field(
				Test.String,
				Count(Test.ID, false).As("count"),
			).
			GroupBy(Test.String).
			Having(
				Greater(Count(Test.ID, false), Value[int64](2)),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "HAVING", "HAVING")
			assertContains(t, sqlSelectQuery, "COUNT(`t`.`id`) > ?", "COUNT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `HAVING`, "HAVING")
			assertContains(t, sqlSelectQuery, `COUNT("t"."id") > $1`, "COUNT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Join(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(Test.ID.As("id")).
			Join(
				Cross(Test1.Table),
				Full(Test1.Table, Equal(Test1.ID, Test.ID)),
				FullOuter(Test1.Table, Equal(Test1.ID, Test.ID)),
				Inner(Test1.Table, Equal(Test1.ID, Test.ID)),
				Left(Test1.Table, Equal(Test1.ID, Test.ID)),
				LeftOuter(Test1.Table, Equal(Test1.ID, Test.ID)),
				Right(Test1.Table, Equal(Test1.ID, Test.ID)),
				RightOuter(Test1.Table, Equal(Test1.ID, Test.ID)),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "JOIN", "JOIN")
			assertContains(t, sqlSelectQuery, "CROSS JOIN `test1` AS `t1`", "CROSS")
			assertContains(t, sqlSelectQuery, "FULL JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "FULL")
			assertContains(t, sqlSelectQuery, "FULL OUTER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "FULL OUTER")
			assertContains(t, sqlSelectQuery, "INNER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "INNER")
			assertContains(t, sqlSelectQuery, "LEFT JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "LEFT")
			assertContains(t, sqlSelectQuery, "LEFT OUTER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "LEFT OUTER")
			assertContains(t, sqlSelectQuery, "RIGHT JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "RIGHT")
			assertContains(t, sqlSelectQuery, "RIGHT OUTER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "RIGHT OUTER")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `JOIN`, "JOIN")
			assertContains(t, sqlSelectQuery, `CROSS JOIN "test1" AS "t1"`, "CROSS")
			assertContains(t, sqlSelectQuery, `FULL JOIN "test1" AS "t1" ON "t1"."id" = "t"."id"`, "FULL")
			assertContains(t, sqlSelectQuery, `FULL OUTER JOIN "test1" AS "t1" ON "t1"."id" = "t"."id"`, "FULL OUTER")
			assertContains(t, sqlSelectQuery, `INNER JOIN "test1" AS "t1" ON "t1"."id" = "t"."id"`, "INNER")
			assertContains(t, sqlSelectQuery, `LEFT JOIN "test1" AS "t1" ON "t1"."id" = "t"."id"`, "LEFT")
			assertContains(t, sqlSelectQuery, `LEFT OUTER JOIN "test1" AS "t1" ON "t1"."id" = "t"."id"`, "LEFT OUTER")
			assertContains(t, sqlSelectQuery, `RIGHT JOIN "test1" AS "t1" ON "t1"."id" = "t"."id"`, "RIGHT")
			assertContains(t, sqlSelectQuery, `RIGHT OUTER JOIN "test1" AS "t1" ON "t1"."id" = "t"."id"`, "RIGHT OUTER")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Limit(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.ID,
			).
			Where(
				Equal(Test.String, Value("active")),
			).
			Limit(10)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "LIMIT ?", "LIMIT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `LIMIT $1`, "LIMIT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Offset(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.ID,
			).
			Where(
				Equal(Test.String, Value("active")),
			).
			Offset(20)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "OFFSET ?", "OFFSET")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `OFFSET $1`, "OFFSET")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_OrderBy(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.ID.As("id"),
			).
			OrderBy(
				Asc(Test.String),
				Desc(Test.ID),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "ORDER BY", "ORDER BY")
			assertContains(t, sqlSelectQuery, "`t`.`string` ASC", "ASC")
			assertContains(t, sqlSelectQuery, "`t`.`id` DESC", "DESC")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `ORDER BY`, "ORDER BY")
			assertContains(t, sqlSelectQuery, `"t"."string" ASC`, "ASC")
			assertContains(t, sqlSelectQuery, `"t"."id" DESC`, "DESC")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Unions(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		// Стандартизировать
		queryUnion := NewSelect(Users.Table).
			Field(
				Users.Name.As("person_name"),
				Users.Email.As("contact_email"),
			).
			Where(
				And(
					Equal(Users.Status, Value("active")),
					Like(Users.Email, Value("%@domain.ltd")),
				),
			)
		queryUnionAll := NewSelect(Products.Table).
			Field(
				Products.Name.As("person_name"),
				Products.Name.As("contact_email"),
			).
			Where(
				Greater(Products.Count, Value(2)),
			)
		queryUnionExcept := NewSelect(Categories.Table).
			Field(
				Categories.Name.As("person_name"),
				Categories.Type.As("contact_email"),
			).
			Where(
				In(Categories.Type, Array("premium", "standard", "basic")),
			)
		queryUnionIntersect := NewSelect(Departments.Table).
			Field(
				Departments.Name.As("person_name"),
				Departments.Name.As("contact_email"),
			).
			Where(
				IsNull(Departments.ParentID),
			)
		stmtSelect := NewSelect(Users.Table).
			Field(
				Users.ID.As("entity_id"),
				Users.Name.As("entity_name"),
				Users.Email.As("entity_email"),
				Users.Status.As("entity_status"),
			).
			Where(
				Equal(Users.Status, Value("active")),
			).
			Unions(
				Union(queryUnion),
				UnionAll(queryUnionAll),
				UnionExcept(queryUnionExcept),
				UnionIntersect(queryUnionIntersect),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "UNION", "UNION")
			assertContains(t, sqlSelectQuery, "UNION ALL", "UNION ALL")
			assertContains(t, sqlSelectQuery, "EXCEPT", "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, "INTERSECT", "UNION INTERSECT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `UNION`, "UNION")
			assertContains(t, sqlSelectQuery, `UNION ALL`, "UNION ALL")
			assertContains(t, sqlSelectQuery, `EXCEPT`, "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, `INTERSECT`, "UNION INTERSECT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Where(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		existsSub := NewSelect(Test.Table).
			Field(Test.ID)
		stmtSelect := NewSelect(Test.Table).
			Field(Test.ID).
			Where(
				And(
					Equal(Test.String, Value("active")),
					Greater(Test.Number, Value(2)),
					Less(Test.Number, Value(2)),
					In(Test.Number, Array(1, 2, 3)),
					Like(Test.String, Value("%ivan")),
					Exists(Subquery[int64](existsSub)),
					NotEqual(Test.String, Value("pending")),
					IsNotNull(Test.String),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "WHERE", "WHERE")
			assertContains(t, sqlSelectQuery, "`t`.`string` = ?", "EQUAL")
			assertContains(t, sqlSelectQuery, "`t`.`number` > ?", "GREATER")
			assertContains(t, sqlSelectQuery, "`t`.`number` < ?", "LESS")
			assertContains(t, sqlSelectQuery, "`t`.`number` IN (?, ?, ?)", "IN")
			assertContains(t, sqlSelectQuery, "`t`.`string` LIKE ?", "LIKE")
			assertContains(t, sqlSelectQuery, "EXISTS", "EXISTS")
			assertContains(t, sqlSelectQuery, "`t`.`string` <> ?", "NOTEQUAL")
			assertContains(t, sqlSelectQuery, "`t`.`string` IS NOT NULL", "IS NOT NULL")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `WHERE`, "WHERE")
			assertContains(t, sqlSelectQuery, `"t"."string" = $1`, "EQUAL")
			assertContains(t, sqlSelectQuery, `"t"."number" > $1`, "GREATER")
			assertContains(t, sqlSelectQuery, `"t"."number" < $1`, "LESS")
			assertContains(t, sqlSelectQuery, `"t"."number" IN ($1, $2, $3)`, "IN")
			assertContains(t, sqlSelectQuery, `"t"."string" LIKE $1`, "LIKE")
			assertContains(t, sqlSelectQuery, `EXISTS`, "EXISTS")
			assertContains(t, sqlSelectQuery, `"t"."string" <> $1`, "NOTEQUAL")
			assertContains(t, sqlSelectQuery, `"t"."string" IS NOT NULL`, "IS NOT NULL")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_With(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		// Стандартизировать
		queryWithN := WithN("WithN", NewSelect(Orders.Table).
			Field(
				Orders.UserID,
				Count(Orders.ID, false).As("order_count"),
				Sum(Orders.Amount, false).As("total_spent"),
			).
			Where(
				Greater(Orders.Amount, Value(2)),
			).
			GroupBy(Orders.UserID),
			"user_id", "order_count", "total_spent",
		)
		queryWithR := WithR("WithR", NewSelect(Departments.Table).
			Field(
				Departments.ID,
				Departments.Name,
				Departments.ParentID,
			).
			Where(IsNull(Departments.ParentID)).
			Unions(
				UnionAll(NewSelect(Departments.Table).
					Field(
						Departments.ID,
						Departments.Name,
						Departments.ParentID,
					).
					Join(Inner(CTE("WithR", "dh"), Equal(Departments.ParentID, Departments.ID)))),
			),
			"dept_id", "dept_name", "parent_id",
		)
		stmtSelect := NewSelect(Users.Table).
			Field(
				Users.ID.As("user_id"),
				Users.Name.As("user_name"),
				Users.Email.As("email"),
				Users.Age.As("age"),
				Users.Status.As("status"),
				Stats.OrderCount.As("premium_order_count"),
				Stats.TotalSpent.As("total_premium_spent"),
				Categories.Name.As("fav_category"),
				Levels.Status.As("user_level_status"),
			).
			Join(
				Inner(CTE("WithN", "us"), Equal(Users.ID, Stats.UserID)),
				Left(Levels.Table,
					And(
						Equal(Users.ID, Levels.UserID),
						Equal(Levels.Status, Value("active")),
					),
				),
				Left(CTE("WithR", "dh"), Equal(Users.ID, Column[int64]("dh", "dept_id"))),
				Left(Categories.Table, Equal(Categories.Type, Users.Status)),
				Left(Query(NewSelect(Categories.Table).
					Field(
						Categories.ID,
						Categories.Name,
						Count(Products.ID, false).As("product_count"),
						Sum(Products.Count, false).As("total_inventory"),
					).
					Join(
						Inner(Products.Table, Equal(Categories.ID, Products.ID)),
					).
					GroupBy(
						Categories.ID,
						Categories.Name,
					), "pbc"), Equal(Categories.ID, Column[int64]("pbc", "id"))),
				Left(Products.Table, Equal(Users.ID, Products.ID)),
			).
			Where(
				And(
					Equal(Users.Status, Value("active")),
					Greater(Stats.OrderCount, Value(2)),
					Or(
						And(
							Greater(Users.Age, Value(2)),
							Less(Users.Age, Value(2)),
						),
						Equal(Users.Age, Value(2)),
					),
					IsNotNull(Users.Email),
					NotLike(Users.Email, Value("%@test.%")),
				),
			).
			GroupBy(
				Users.ID,
				Users.Name,
				Users.Email,
			).
			Having(
				And(
					Greater(Stats.OrderCount, Value(2)),
					Greater(Stats.TotalSpent, Value(2)),
					Or(
						Equal(Levels.Status, Value("active")),
						IsNull(Levels.ID),
					),
				),
			).
			OrderBy(
				Desc(Stats.TotalSpent),
				Desc(Stats.OrderCount),
				Asc(Users.Name),
			).
			Limit(25).
			Offset(0).
			With(queryWithN, queryWithR)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "WITH", "WITH")
			assertContains(t, sqlSelectQuery, "WithN", "WithN")
			assertContains(t, sqlSelectQuery, "WithR", "WithR")
			assertContains(t, sqlSelectQuery, "RECURSIVE", "RECURSIVE")
			assertContains(t, sqlSelectQuery, "UNION ALL", "UNION ALL for RCTE")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `WITH`, "WITH")
			assertContains(t, sqlSelectQuery, `WithN`, "WithN")
			assertContains(t, sqlSelectQuery, `WithR`, "WithR")
			assertContains(t, sqlSelectQuery, `RECURSIVE`, "RECURSIVE")
			assertContains(t, sqlSelectQuery, `UNION ALL`, "UNION ALL for RCTE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Update(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtUpdate := NewUpdate(Test.Table).
			Set(
				Pair(Test.String, Value("active")),
			).
			Where(
				Equal(Test.Number, Value(2)),
			)
		sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlUpdateQuery, "UPDATE", "UPDATE")
			assertContains(t, sqlUpdateQuery, "SET", "SET")
		case DialectPostgreSQL:
			assertContains(t, sqlUpdateQuery, `UPDATE`, "UPDATE")
			assertContains(t, sqlUpdateQuery, `SET`, "SET")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
	})
}

// Приватные функции
func assertContains(t *testing.T, str, substr string, message string) {
	t.Helper()
	escaped := regexp.QuoteMeta(substr)
	re := regexp.MustCompile(`\\\$[0-9]+`)
	pattern := re.ReplaceAllString(escaped, `\$[0-9]+`)
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		t.Fatalf("Reg err: [%v]\n", err)
	}
	if !matched {
		t.Errorf("Req: [%s] / Pat: [%s]\n", message, substr)
	}
}
func testAllDialects(t *testing.T, testFunc func(t *testing.T, supportDialect *SupportDialect)) {
	for _, supportDialect := range listSupportDialects {
		currentDialect := supportDialect
		t.Run(currentDialect.name, func(t *testing.T) {
			testFunc(t, currentDialect)
		})
	}
}
