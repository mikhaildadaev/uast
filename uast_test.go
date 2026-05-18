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
				Test.Column.ID,
			).
			Where(
				And(
					Equal(Test.Column.Number, BitwiseAnd(Test.Column.Number, Value(0b0010))),
					Equal(Test.Column.Number, BitwiseOr(Test.Column.Number, Value(0b0010))),
					Equal(Test.Column.Number, BitwiseXor(Test.Column.Number, Value(0b0010))),
					Equal(Test.Column.Number, Divide(Test.Column.Number, Value(2))),
					Equal(Test.Column.Number, Minus(Test.Column.Number, Value(2))),
					Equal(Test.Column.Number, Modulo(Test.Column.Number, Value(2))),
					Equal(Test.Column.Number, Multiply(Test.Column.Number, Value(2))),
					Equal(Test.Column.Number, Plus(Test.Column.Number, Value(2))),
					Equal(Test.Column.Number, ShiftLeft(Test.Column.Number, Value(2))),
					Equal(Test.Column.Number, ShiftRight(Test.Column.Number, Value(2))),
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
			Field(
				Test.Column.ID,
			)
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
			Field(
				Test.Column.ID,
			).
			Where(
				And(
					Between(Test.Column.Number, Value(0), Value(2)),
					Equal(Test.Column.Number, Value(2)),
					Exists(Subquery[int](NewSelect(Test.Table).Field(ConstIntOne()))),
					Greater(Test.Column.Number, Value(2)),
					GreaterEqual(Test.Column.Number, Value(2)),
					ILike(Test.Column.String, Value("%ivan%")),
					In(Test.Column.String, Array("active", "pending")),
					IsNotNull(Test.Column.String),
					IsNull(Test.Column.String),
					Less(Test.Column.Number, Value(2)),
					LessEqual(Test.Column.Number, Value(2)),
					Like(Test.Column.String, Value("%ivan%")),
					NotBetween(Test.Column.Number, Value(0), Value(2)),
					NotEqual(Test.Column.Number, Value(2)),
					NotExists(Subquery[int](NewSelect(Test.Table).Field(ConstIntOne()))),
					NotILike(Test.Column.String, Value("%ivan%")),
					NotIn(Test.Column.String, Array("active", "pending")),
					NotLike(Test.Column.String, Value("%ivan%")),
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
			Field(
				Test.Column.ID,
			).
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
				Avg(Test.Column.Number, false).As("aggregate_avg"),
				BitAnd(Test.Column.Number, false).As("aggregate_bitand"),
				BitOr(Test.Column.Number, false).As("aggregate_bitor"),
				BitXor(Test.Column.Number, false).As("aggregate_bitxor"),
				Count(Test.Column.String, false).As("aggregate_count"),
				GroupConcat(Test.Column.String, false).As("aggregate_groupconcat"),
				Max(Test.Column.Number, false).As("aggregate_max"),
				Min(Test.Column.Number, false).As("aggregate_min"),
				StdDev(Test.Column.Number, false).As("aggregate_stddev"),
				Sum(Test.Column.Number, false).As("aggregate_sum"),
				Variance(Test.Column.Number, false).As("aggregate_variance"),
				// Функции аналитические
				FirstValue(Test.Column.Name).Over(
					PartitionBy(Test.Column.ID),
					OrderBy(Desc(Test.Column.Number)),
				).As("analytical_firstvalue"),
				Lag(Test.Column.Number, 2).Over(
					PartitionBy(Test.Column.ID),
					OrderBy(Asc(Test.Column.Date)),
				).As("analytical_lag"),
				LastValue(Test.Column.Name).Over(
					PartitionBy(Test.Column.ID),
					OrderBy(Asc(Test.Column.Number)),
					RowsBetween("CURRENT ROW", "UNBOUNDED FOLLOWING"),
				).As("analytical_lastvalue"),
				Lead(Test.Column.Number, 2).Over(
					PartitionBy(Test.Column.ID),
					OrderBy(Asc(Test.Column.Date)),
				).As("analytical_lead"),
				NthValue(Test.Column.Name, 2).Over(
					PartitionBy(Test.Column.ID),
					OrderBy(Desc(Test.Column.Number)),
					RowsBetween("UNBOUNDED PRECEDING", "CURRENT ROW"),
				).As("analytical_nthvalue"),
				// Функции условий
				Case(CaseIf(CasePair(Less(Test.Column.Number, Value(2)), Value("old"))), CaseElse(Value("new"))).As("condition_case"),
				Coalesce(Test.Column.CreateAt, Test.Column.UpdateAt).As("condition_coalesce"),
				Greatest(Test.Column.CreateAt, Test.Column.UpdateAt).As("condition_greatest"),
				Least(Test.Column.CreateAt, Test.Column.UpdateAt).As("condition_least"),
				NullIf(Test.Column.CreateAt, Test.Column.UpdateAt).As("condition_if"),
				// Функции конвертации
				Cast(Test.Column.Number, TypeString).As("convert_cast"),
				CharLength(Test.Column.String).As("convert_length"),
				DateFormat(Test.Column.CreateAt, Literal("%Y-%m-%d")).As("convert_dateformat"),
				Degrees(Test.Column.Number).As("convert_degrees"),
				Length(Test.Column.String).As("convert_length"),
				Position(Test.Column.String, Value("old")).As("convert_position"),
				Radians(Test.Column.Number).As("convert_radians"),
				// Функции даты и времени
				CurDate().As("datetime_curdate"),
				CurTime().As("datetime_curtime"),
				DateAdd(Test.Column.CreateAt, Literal("2 DAY")).As("datetime_dateadd"),
				DateDiff(Test.Column.UpdateAt, Test.Column.CreateAt).As("datetime_datediff"),
				DateSub(Test.Column.CreateAt, Literal("2 DAY")).As("datetime_datesub"),
				Day(Test.Column.CreateAt).As("datetime_day"),
				DayName(Test.Column.CreateAt).As("datetime_dayname"),
				Hour(Test.Column.CreateAt).As("datetime_hour"),
				Minute(Test.Column.CreateAt).As("datetime_minute"),
				Month(Test.Column.CreateAt).As("datetime_month"),
				MonthName(Test.Column.CreateAt).As("datetime_monthname"),
				Now().As("datetime_now"),
				Quarter(Test.Column.CreateAt).As("datetime_quarter"),
				Second(Test.Column.CreateAt).As("datetime_second"),
				TimeAdd(Test.Column.CreateAt, Literal("2 HOUR")).As("datetime_timeadd"),
				TimeDiff(Test.Column.UpdateAt, Test.Column.CreateAt).As("datetime_timediff"),
				TimeSub(Test.Column.CreateAt, Literal("2 HOUR")).As("datetime_timesub"),
				Week(Test.Column.CreateAt).As("datetime_week"),
				Year(Test.Column.CreateAt).As("datetime_year"),
				// Функции обмена данными
				JsonArray(Test.Column.Json, Value("val1"), Value("val2")).As("json_jsonarray"),
				JsonArrayAgg(Test.Column.Json).As("json_jsonarrayagg"),
				JsonContains(Test.Column.Json, Value(`{"key":"val"}`)).As("json_jsoncontains"),
				JsonExtract(Test.Column.Json, JsonGroup(JsonPath(JsonKey("parent"), JsonIndex(0), JsonKey("child"))), TypeString).As("json_jsonextract"),
				JsonObject(JsonPair(JsonKey("key"), Count(Test.Column.Json, false))).As("json_jsonobject"),
				JsonObjectAgg(Test.Column.Json, Test.Column.Number).As("json_jsonobjectagg"),
				JsonRemove(Test.Column.Json, JsonGroup(JsonPath(JsonKey("key1"))), JsonGroup(JsonPath(JsonKey("key2")))).As("json_jsonremove"),
				JsonSet(Test.Column.Json, JsonGroup(JsonPath(JsonKey("key1")), Value("val1")), JsonGroup(JsonPath(JsonKey("key2")), Value("val2"))).As("json_jsonset"),
				JsonType(Test.Column.Json).As("json_jsontype"),
				// Функции математические
				Abs(Test.Column.Number).As("math_abs"),
				ACos(Test.Column.Number).As("math_acos"),
				ASin(Test.Column.Number).As("math_asin"),
				ATan(Test.Column.Number).As("math_atan"),
				ATan2(Test.Column.Y, Test.Column.X).As("math_atan2"),
				Cbrt(Test.Column.Number).As("math_cbrt"),
				Ceil(Test.Column.Number).As("math_ceil"),
				Cos(Test.Column.Number).As("math_cos"),
				Exp(Test.Column.Number).As("math_exp"),
				Floor(Test.Column.Number).As("math_floor"),
				Ln(Test.Column.Number).As("math_ln"),
				Log(Test.Column.Number, Value(2)).As("math_log"),
				Mod(Test.Column.Number, Value(2)).As("math_mod"),
				Pi().As("math_pi"),
				Power(Test.Column.Number, Value(2)).As("math_power"),
				Rand().As("math_rand"),
				Round(Test.Column.Number, Value(2)).As("math_round"),
				Sin(Test.Column.Number).As("math_sin"),
				Sqrt(Test.Column.Number).As("math_sqrt"),
				Tan(Test.Column.Number).As("math_tan"),
				Trunc(Test.Column.Number, Value(2)).As("math_trunc"),
				// Функции ранжирующие
				CumeDist().Over(
					PartitionBy(Test.Column.ID),
					OrderBy(Desc(Test.Column.Number)),
				).As("ranking_cumedist"),
				DenseRank().Over(
					PartitionBy(Test.Column.ID),
					OrderBy(Desc(Test.Column.Number)),
				).As("ranking_denserank"),
				NTile(2).Over(
					PartitionBy(Test.Column.ID),
					OrderBy(Desc(Test.Column.Number)),
				).As("ranking_ntile"),
				PercentRank().Over(
					PartitionBy(Test.Column.ID),
					OrderBy(Desc(Test.Column.Number)),
				).As("ranking_percentrank"),
				Rank().Over(
					PartitionBy(Test.Column.ID),
					OrderBy(Desc(Test.Column.Number)),
				).As("ranking_rank"),
				RowNumber().Over(
					PartitionBy(Test.Column.ID),
					OrderBy(Desc(Test.Column.Number)),
				).As("ranking_rownumber"),
				// Функции строковые
				Concat(Test.Column.String, Value("old"), Value("new")).As("string_concat"),
				ConcatWs(Value("_"), Test.Column.String, Value("old"), Value("new")).As("string_concatws"),
				LeftString(Test.Column.String, Value(2)).As("string_lstr"),
				Lower(Test.Column.String).As("string_lower"),
				LPad(Test.Column.String, Value(2), Value(",")).As("string_lpad"),
				LTrim(Test.Column.String).As("string_ltrim"),
				Repeat(Test.Column.String, Value(2)).As("string_repeat"),
				Replace(Test.Column.String, Value("old"), Value("new")).As("string_replace"),
				Reverse(Test.Column.String).As("string_reverse"),
				RightString(Test.Column.String, Value(2)).As("string_rstr"),
				RPad(Test.Column.String, Value(2), Value(",")).As("string_rpad"),
				RTrim(Test.Column.String).As("string_rtrim"),
				SubString(Test.Column.String, Value(0), Value(2)).As("string_substring"),
				Trim(Test.Column.String).As("string_trim"),
				Upper(Test.Column.String).As("string_upper"),
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
			Field(
				Test.Column.ID,
			).
			Where(
				Equal(DateFormat(Test.Column.CreateAt, Literal("%Y-%m-%d")), Value("2026-01-01")),
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
			Field(
				Test.Column.ID,
			).
			Where(
				And(
					And(
						Equal(Test.Column.String, Value("active")),
						Greater(Test.Column.Number, Value(2)),
					),
					Or(
						Equal(Test.Column.String, Value("active")),
						Greater(Test.Column.Number, Value(2)),
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
			Field(
				Test.Column.ID,
			).
			OrderBy(
				Asc(Test.Column.String),
				Desc(Test.Column.String),
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
			Field(
				Subquery[int64](NewSelect(Test.Table).Field(Test.Column.ID)).As("SUB"),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "(SELECT `t`.`id` FROM `test` AS `t`)", "SUBQUERY")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `(SELECT "t"."id" FROM "test" AS "t")`, "SUBQUERY")
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
			Field(
				Test.Column.ID,
			).
			Where(
				Equal(Test.Column.String, Value(data)),
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
				Equal(Test.Column.String, Value("active")),
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "DELETE `t` FROM `test` AS `t`", "DELETE")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `DELETE FROM "test" AS "t"`, "DELETE")
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
				Inner(Test1.Table, Equal(Test1.Column.ID, Test.Column.ID)),
				Left(Test2.Table, Equal(Test2.Column.ID, Test.Column.ID)),
			).
			Where(
				Equal(Test.Column.String, Value("active")),
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "INNER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "INNER JOIN")
			assertContains(t, sqlDeleteQuery, "LEFT JOIN `test2` AS `t2` ON `t2`.`id` = `t`.`id`", "LEFT JOIN")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `USING "test1" AS "t1", "test2" AS "t2"`, "USING LIST")
			assertContains(t, sqlDeleteQuery, `"t1"."id" = "t"."id" AND "t2"."id" = "t"."id"`, "USING CONDITION")
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
				Equal(Test.Column.String, Value("active")),
			).
			Returning(
				Test.Column.ID,
				Test.Column.String,
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			// Not support
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `RETURNING "t"."id", "t"."string"`, "RETURNING")
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
				Equal(Test.Column.String, Value("active")),
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "WHERE `t`.`string` = ?", "WHERE")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `WHERE "t"."string" = $1`, "WHERE")
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
		stmtWithN := WithN("cte_nr", NewSelect(Test.Table).
			Field(
				Test.Column.ID,
			).
			Where(
				Equal(Test.Column.String, Value("old")),
			),
		)
		stmtDelete := NewDelete(Test.Table).
			Where(
				In(Test.Column.ID, Subquery[int64](NewSelect(Test.Table).Field(Column[int64]("cte_nr", "id")))),
			).
			With(
				stmtWithN,
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "WITH `cte_nr`", "WITH")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `WITH "cte_nr"`, "WITH")
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
			Column(Test.Column.String, Test.Column.Number).
			Values(
				Row(
					Value("ivan"),
					Value(2),
				),
			)
		sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlInsertQuery, "INSERT INTO `test` AS `t`", "INSERT")
		case DialectPostgreSQL:
			assertContains(t, sqlInsertQuery, `INSERT INTO "test" AS "t"`, "INSERT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
	})
}
func Test_Insert_Returning(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		// Изменить реализацию
		stmtInsert := NewInsert(Test.Table).
			Column(Test.Column.String, Test.Column.Number).
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
		case DialectPostgreSQL:
			assertContains(t, sqlInsertQuery, `INSERT`, "INSERT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
	})
}
func Test_Insert_Source(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		// Изменить реализацию
		stmtInsert := NewInsert(Test.Table).
			Column(Test.Column.String, Test.Column.Number).
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
		case DialectPostgreSQL:
			assertContains(t, sqlInsertQuery, `INSERT`, "INSERT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
	})
}
func Test_Insert_Values(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		// Изменить реализацию
		stmtInsert := NewInsert(Test.Table).
			Column(Test.Column.String, Test.Column.Number).
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
		case DialectPostgreSQL:
			assertContains(t, sqlInsertQuery, `INSERT`, "INSERT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
	})
}
func Test_Insert_With(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		// Изменить реализацию
		stmtInsert := NewInsert(Test.Table).
			Column(Test.Column.String, Test.Column.Number).
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
		case DialectPostgreSQL:
			assertContains(t, sqlInsertQuery, `INSERT`, "INSERT")
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
			Field(
				Test.Column.String,
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "SELECT `t`.`string` FROM `test` AS `t`", "SELECT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `SELECT "t"."string" FROM "test" AS "t"`, "SELECT")
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
		stmtSelect := NewSelect(Test.Table).
			Distinct().
			Field(
				Test.Column.ID,
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "DISTINCT", "DISTINCT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `DISTINCT`, "DISTINCT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
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
				Test.Column.String,
				Count(Test.Column.ID, false).As("count"),
			).
			GroupBy(
				Test.Column.String,
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "GROUP BY `t`.`string`", "GROUP BY")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `GROUP BY "t"."string"`, "GROUP BY")
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
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.Column.String,
				Count(Test.Column.ID, false).As("count"),
			).
			GroupBy(
				Test.Column.String,
			).
			Having(
				Greater(Count(Test.Column.ID, false), Value[int64](2)),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "HAVING COUNT(`t`.`id`) > ?", "COUNT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `HAVING COUNT("t"."id") > $1`, "COUNT")
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
			Field(
				Test.Column.ID,
			).
			Join(
				Cross(Test1.Table),
				Full(Test1.Table, Equal(Test1.Column.ID, Test.Column.ID)),
				FullOuter(Test1.Table, Equal(Test1.Column.ID, Test.Column.ID)),
				Inner(Test1.Table, Equal(Test1.Column.ID, Test.Column.ID)),
				Left(Test1.Table, Equal(Test1.Column.ID, Test.Column.ID)),
				LeftOuter(Test1.Table, Equal(Test1.Column.ID, Test.Column.ID)),
				Right(Test1.Table, Equal(Test1.Column.ID, Test.Column.ID)),
				RightOuter(Test1.Table, Equal(Test1.Column.ID, Test.Column.ID)),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "CROSS JOIN `test1` AS `t1`", "CROSS")
			assertContains(t, sqlSelectQuery, "FULL JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "FULL")
			assertContains(t, sqlSelectQuery, "FULL OUTER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "FULL OUTER")
			assertContains(t, sqlSelectQuery, "INNER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "INNER")
			assertContains(t, sqlSelectQuery, "LEFT JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "LEFT")
			assertContains(t, sqlSelectQuery, "LEFT OUTER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "LEFT OUTER")
			assertContains(t, sqlSelectQuery, "RIGHT JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "RIGHT")
			assertContains(t, sqlSelectQuery, "RIGHT OUTER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "RIGHT OUTER")
		case DialectPostgreSQL:
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
				Test.Column.ID,
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
				Test.Column.ID,
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
				Test.Column.ID,
			).
			OrderBy(
				Test.Column.String,
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "ORDER BY `t`.`string`", "ORDER BY")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `ORDER BY "t"."string"`, "ORDER BY")
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
		stmtWithR := WithR("cte_re", NewSelect(Test.Table).
			Field(
				Test.Column.ID,
			).
			Where(
				Equal(Test.Column.Number, Value(0)),
			).
			Unions(
				UnionAll(NewSelect(Test.Table).
					Field(
						Test.Column.ID,
					).
					Join(
						Inner(NewCTE("cte_re", "ctere"), Equal(Test.Column.ID, Column[int64]("ctere", "id"))),
					),
				),
			),
		)
		stmtUnion := NewSelect(Test.Table).
			Field(
				Test.Column.String,
			)
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.Column.String,
			).
			Unions(
				Union(stmtUnion),
				UnionAll(stmtUnion),
				UnionExcept(stmtUnion),
				UnionIntersect(stmtUnion),
			).
			With(
				stmtWithR,
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "WITH RECURSIVE `cte_re`", "WITH")
			assertContains(t, sqlSelectQuery, "UNION", "UNION")
			assertContains(t, sqlSelectQuery, "UNION ALL", "UNION ALL")
			assertContains(t, sqlSelectQuery, "EXCEPT", "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, "INTERSECT", "UNION INTERSECT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `WITH RECURSIVE "cte_re"`, "WITH")
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
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.Column.ID,
			).
			Where(
				Equal(Test.Column.String, Value("active")),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "WHERE `t`.`string` = ?", "WHERE")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `WHERE "t"."string" = $1`, "EQUAL")
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
		stmtWithN := WithN("cte_nr", NewSelect(Test.Table).
			Field(
				Test.Column.ID,
				Test.Column.String,
			).
			Where(
				Equal(Test.Column.String, Value("active")),
			),
			"id", "string",
		)
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.Column.ID,
				Test.Column.Number,
			).
			Join(
				Inner(NewCTE("cte_nr", "ctenr"), Equal(Test.Column.ID, Column[int64]("ctenr", "id"))),
			).
			With(
				stmtWithN,
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "WITH `cte_nr`", "WITH")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `WITH "cte_nr"`, "WITH")
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
				Assign(Test.Column.String, Value("active")),
			).
			Where(
				Equal(Test.Column.Number, Value(2)),
			)
		sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t`", "UPDATE")
		case DialectPostgreSQL:
			assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t"`, "UPDATE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
	})
}
func Test_Update_Join(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtUpdate := NewUpdate(Test.Table).
			Set(
				Assign(Test.Column.String, Value("active")),
			).
			Join(
				Inner(Test1.Table, Equal(Test1.Column.ID, Test.Column.ID)),
			).
			Where(Equal(Test1.Column.String, Value("active")))
		sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t` INNER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`", "UPDATE")
		case DialectPostgreSQL:
			assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t" INNER JOIN "test1" AS "t1" ON "t1"."id" = "t"."id"`, "UPDATE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
	})
}
func Test_Update_Returning(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtUpdate := NewUpdate(Test.Table).
			Set(
				Assign(Test.Column.String, Value("active")),
			).
			Where(
				Equal(Test.Column.Number, Value(2)),
			).
			Returning(
				Test.Column.ID,
			)
		sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
		switch supportDialect {
		case DialectMySQL:
			// Not support
		case DialectPostgreSQL:
			assertContains(t, sqlUpdateQuery, `RETURNING "t"."id"`, "RETURNING")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
	})
}
func Test_Update_Set(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtUpdate := NewUpdate(Test.Table).
			Set(
				Assign(Test.Column.String, Value("active")),
			).
			Where(
				Equal(Test.Column.Number, Value(2)),
			)
		sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlUpdateQuery, "SET `t`.`string` = ?", "SET")
		case DialectPostgreSQL:
			assertContains(t, sqlUpdateQuery, `SET "t"."string" = $1`, "SET")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
	})
}
func Test_Update_Where(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtUpdate := NewUpdate(Test.Table).
			Set(
				Assign(Test.Column.String, Value("active")),
			).
			Where(
				Equal(Test.Column.Number, Value(2)),
			)
		sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlUpdateQuery, "WHERE `t`.`number` = ?", "WHERE")
		case DialectPostgreSQL:
			assertContains(t, sqlUpdateQuery, `WHERE "t"."number" = $2`, "WHERE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
	})
}
func Test_Update_With(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtWithN := WithN("cte_nr", NewSelect(Test.Table).
			Field(
				Test.Column.ID,
			).
			Where(
				Equal(Test.Column.String, Value("pending")),
			),
		)
		stmtUpdate := NewUpdate(Test.Table).
			Set(
				Assign(Test.Column.String, Value("active")),
			).
			Where(
				In(Test.Column.ID, Subquery[int64](NewSelect(Test.Table).Field(Column[int64]("cte_nr", "id")))),
			).
			With(
				stmtWithN,
			)
		sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlUpdateQuery, "WITH `cte_nr`", "WITH")
		case DialectPostgreSQL:
			assertContains(t, sqlUpdateQuery, `WITH "cte_nr"`, "WITH")
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
