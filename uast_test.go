package uast

import (
	"regexp"
	"testing"
)

// Публичные функции
func Test_Delete(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		queryWith := WithN("user_stats", NewSelect(
			Orders.UserID,
			Count(Orders.ID, false).As("order_count"),
		).
			From(Orders.Table).
			Where(Greater(Orders.Amount, Value(100))).
			GroupBy(Orders.UserID),
			"user_id", "order_count")
		stmtDelete := NewDelete(Users.Table).
			Join(
				Inner(CTE("user_stats", "us"), Equal(Users.ID, Stats.UserID)),
				Left(Levels.Table, Equal(Users.ID, Levels.UserID)),
			).
			Where(
				And(
					Equal(Users.Status, Value("active")),
					Or(
						IsNull(Levels.ID),
						Equal(Levels.Status, Value("expired")),
					),
				),
			).
			Returning(
				Users.ID.As("user_id"),
				Users.Name.As("user_name"),
				Stats.OrderCount.As("total_orders"),
			).
			With(queryWith)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "DELETE", "DELETE")
			assertContains(t, sqlDeleteQuery, "FROM", "FROM")
			assertContains(t, sqlDeleteQuery, "JOIN", "JOIN")
			assertContains(t, sqlDeleteQuery, "WHERE", "WHERE")
			assertContains(t, sqlDeleteQuery, "WITH", "WITH")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, "DELETE", "DELETE")
			assertContains(t, sqlDeleteQuery, "FROM", "FROM")
			assertContains(t, sqlDeleteQuery, "RETURNING", "RETURNING")
			assertContains(t, sqlDeleteQuery, "USING", "USING")
			assertContains(t, sqlDeleteQuery, "WHERE", "WHERE")
			assertContains(t, sqlDeleteQuery, "WITH", "WITH")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_Delete_Join(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		stmtDelete := NewDelete(Users.Table).
			Join(
				Inner(Orders.Table, Equal(Users.ID, Orders.UserID)),
				Left(Levels.Table, Equal(Users.ID, Levels.UserID)),
			).
			Where(
				Equal(Orders.Status, Value("cancelled")),
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "INNER JOIN", "INNER JOIN")
			assertContains(t, sqlDeleteQuery, "LEFT JOIN", "LEFT JOIN")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `"orders" AS "o", "level" AS "l"`, "USING LIST")
			assertContains(t, sqlDeleteQuery, `"u"."id" = "o"."user_id" AND "u"."id" = "l"."user_id"`, "JOIN CONDITION")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_Delete_Returning(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		stmtDelete := NewDelete(Users.Table).
			Where(Equal(Users.Status, Value("inactive"))).
			Returning(
				Users.ID.As("user_id"),
				Users.Name.As("user_name"),
				Users.Email.As("user_email"),
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			// Not support
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `"user_id"`, "RETURNING id")
			assertContains(t, sqlDeleteQuery, `"user_name"`, "RETURNING name")
			assertContains(t, sqlDeleteQuery, `"user_email"`, "RETURNING email")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_Delete_Where(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		existsSub := NewSelect(Levels.ID).
			From(Levels.Table).
			Where(
				Equal(Levels.UserID, Users.ID),
			)
		stmtDelete := NewDelete(Users.Table).
			Where(
				And(
					Equal(Users.Status, Value("active")),
					Greater(Users.Age, Value(18)),
					Less(Users.Age, Value(65)),
					In(Users.DepartmentID, Array[int64](1, 2, 3)),
					Like(Users.Email, Value("%@company.com")),
					Exists(Subquery[int](existsSub)),
					NotEqual(Users.Status, Value("deleted")),
					IsNotNull(Users.Email),
				),
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "`u`.`status` = ?", "EQUAL")
			assertContains(t, sqlDeleteQuery, "`u`.`age` > ?", "GREATER")
			assertContains(t, sqlDeleteQuery, "`u`.`age` < ?", "LESS")
			assertContains(t, sqlDeleteQuery, "`u`.`department_id` IN (?, ?, ?)", "IN")
			assertContains(t, sqlDeleteQuery, "`u`.`email` LIKE ?", "LIKE")
			assertContains(t, sqlDeleteQuery, "EXISTS", "EXISTS")
			assertContains(t, sqlDeleteQuery, "`u`.`status` <> ?", "NOTEQUAL")
			assertContains(t, sqlDeleteQuery, "`u`.`email` IS NOT NULL", "IS NOT NULL")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `"u"."status" = $1`, "EQUAL")
			assertContains(t, sqlDeleteQuery, `"u"."age" > $1`, "GREATER")
			assertContains(t, sqlDeleteQuery, `"u"."age" < $1`, "LESS")
			assertContains(t, sqlDeleteQuery, `"u"."department_id" IN ($1, $2, $3)`, "IN")
			assertContains(t, sqlDeleteQuery, `"u"."email" LIKE $1`, "LIKE")
			assertContains(t, sqlDeleteQuery, `EXISTS`, "EXISTS")
			assertContains(t, sqlDeleteQuery, `"u"."status" <> $1`, "NOTEQUAL")
			assertContains(t, sqlDeleteQuery, `"u"."email" IS NOT NULL`, "IS NOT NULL")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_Delete_With(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		cte := WithN("old_users", NewSelect(Users.ID).
			From(Users.Table).
			Where(Less(Users.Age, Value(18))),
		)
		stmtDelete := NewDelete(Users.Table).
			Where(In(Users.ID, Subquery[int64](NewSelect(Column[int64]("old_users", "id"))))).
			With(cte)

		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "WITH", "WITH")
			assertContains(t, sqlDeleteQuery, "old_users", "CTE")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, "WITH", "WITH")
			assertContains(t, sqlDeleteQuery, "old_users", "CTE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_Insert(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		queryWith := WithN("user_stats", NewSelect(
			Orders.UserID,
			Count(Orders.ID, false).As("order_count"),
		).
			From(Orders.Table).
			Where(
				Greater(Orders.Amount, Value(100)),
			).
			GroupBy(Orders.UserID),
			"user_id", "order_count",
		)
		stmtInsert := NewInsert(Users.Name, Users.Age).
			Into(Users.Table).
			Source(NewSelect(Value("Test User"), Value(35)).
				Where(
					Exists(Subquery[int64](NewSelect(Stats.UserID).
						From(CTE("user_stats", "us")).
						Where(
							Greater(Stats.OrderCount, Value(3)),
						),
					)),
				),
			).
			Returning(
				Users.ID.As("user_id"),
				Users.Name.As("user_name"),
			).
			With(queryWith)
		sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
		switch supportDialect {
		case DialectMySQL:
			//assertContains(t, sqlInsertQuery, "INSERT", "INSERT")
			//assertContains(t, sqlInsertQuery, "INTO", "INTO")
			//assertContains(t, sqlInsertQuery, "SOURCE", "SOURCE")
			//assertContains(t, sqlInsertQuery, "WITH", "WITH")
		case DialectPostgreSQL:
			//assertContains(t, sqlInsertQuery, "INSERT", "INSERT")
			//assertContains(t, sqlInsertQuery, "INTO", "INTO")
			//assertContains(t, sqlInsertQuery, "RETURNING", "RETURNING")
			//assertContains(t, sqlInsertQuery, "SOURCE", "SOURCE")
			//assertContains(t, sqlInsertQuery, "WITH", "WITH")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
	})
}
func Test_Select(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		queryWith := WithN("user_stats", NewSelect(
			Orders.UserID,
			Count(Orders.ID, false).As("order_count"),
			Sum(Orders.Amount, false).As("total_spent"),
		).
			From(Orders.Table).
			Where(Greater(Orders.Amount, Value(100))).
			GroupBy(Orders.UserID),
			"user_id", "order_count", "total_spent",
		)
		stmtSelect := NewSelect(
			Users.ID.As("user_id"),
			Users.Name.As("user_name"),
			Stats.OrderCount.As("premium_orders"),
			Stats.TotalSpent.As("total_premium_spent"),
			Count(Products.ID, false).As("count_product"),
		).
			From(Users.Table).
			Join(
				Left(CTE("user_stats", "us"), Equal(Users.ID, Stats.UserID)),
				Left(Levels.Table,
					And(
						Equal(Users.ID, Levels.UserID),
						Equal(Levels.Status, Value("active")),
					),
				),
			).
			Where(
				Or(
					Equal(Users.Status, Value("active")),
					IsNotNull(Levels.ID),
				),
			).
			GroupBy(
				Users.ID,
				Users.Name,
				Stats.OrderCount,
				Stats.TotalSpent,
			).
			Having(
				Or(
					IsNotNull(Users.ID),
					Greater(Stats.OrderCount, Value(0)),
				),
			).
			OrderBy(
				Desc(Stats.TotalSpent),
				Asc(Users.Name),
			).
			Limit(50).
			Offset(10).
			With(queryWith)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "FROM", "FROM")
			assertContains(t, sqlSelectQuery, "GROUP BY", "GROUP BY")
			assertContains(t, sqlSelectQuery, "HAVING", "HAVING")
			assertContains(t, sqlSelectQuery, "JOIN", "JOIN")
			assertContains(t, sqlSelectQuery, "LIMIT", "LIMIT")
			assertContains(t, sqlSelectQuery, "OFFSET", "OFFSET")
			assertContains(t, sqlSelectQuery, "ORDER BY", "ORDER BY")
			assertContains(t, sqlSelectQuery, "SELECT", "SELECT")
			assertContains(t, sqlSelectQuery, "WHERE", "WHERE")
			assertContains(t, sqlSelectQuery, "WITH", "WITH")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, "FROM", "FROM")
			assertContains(t, sqlSelectQuery, "GROUP BY", "GROUP BY")
			assertContains(t, sqlSelectQuery, "HAVING", "HAVING")
			assertContains(t, sqlSelectQuery, "JOIN", "JOIN")
			assertContains(t, sqlSelectQuery, "LIMIT", "LIMIT")
			assertContains(t, sqlSelectQuery, "OFFSET", "OFFSET")
			assertContains(t, sqlSelectQuery, "ORDER BY", "ORDER BY")
			assertContains(t, sqlSelectQuery, "SELECT", "SELECT")
			assertContains(t, sqlSelectQuery, "WHERE", "WHERE")
			assertContains(t, sqlSelectQuery, "WITH", "WITH")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Comparison(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		stmtSelect := NewSelect(Users.ID.As("user_data_id")).
			From(Users.Table).
			Where(
				And(
					Between(Users.Age, Value(18), Value(65)),
					Equal(Users.Age, Value(18)),
					Exists(Subquery[int](NewSelect(ConstIntOne()).From(Users.Table))),
					Greater(Users.Age, Value(18)),
					GreaterEqual(Users.Age, Value(18)),
					ILike(Users.Email, Value("john%")),
					In(Users.Status, Array("active", "pending")),
					IsNotNull(Users.ID),
					IsNull(Users.ID),
					Less(Users.Age, Value(18)),
					LessEqual(Users.Age, Value(18)),
					Like(Users.Email, Value("%@gmail.com")),
					NotBetween(Users.Age, Value(0), Value(17)),
					NotEqual(Users.Age, Value(18)),
					NotExists(Subquery[int](NewSelect(ConstIntOne()).From(Users.Table))),
					NotILike(Users.Email, Value("admin%")),
					NotIn(Users.Status, Array("banned", "deleted")),
					NotLike(Users.Email, Value("%@test.%")),
				))
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`u`.`age` BETWEEN ? AND ?", "BETWEEN")
			assertContains(t, sqlSelectQuery, "`u`.`age` = ?", "EQUAL")
			assertContains(t, sqlSelectQuery, "EXISTS (SELECT 1 FROM `users` AS `u`)", "EXISTS")
			assertContains(t, sqlSelectQuery, "`u`.`age` > ?", "GREATER")
			assertContains(t, sqlSelectQuery, "`u`.`age` >= ?", "GREATEREQUAL")
			assertContains(t, sqlSelectQuery, "LOWER(`u`.`email`) LIKE LOWER(?)", "ILIKE")
			assertContains(t, sqlSelectQuery, "`u`.`status` IN (?, ?)", "IN")
			assertContains(t, sqlSelectQuery, "`u`.`id` IS NOT NULL", "IS NOT NULL")
			assertContains(t, sqlSelectQuery, "`u`.`id` IS NULL", "IS NULL")
			assertContains(t, sqlSelectQuery, "`u`.`age` < ?", "LESS")
			assertContains(t, sqlSelectQuery, "`u`.`age` <= ? ", "LESSEQUAL")
			assertContains(t, sqlSelectQuery, "`u`.`email` LIKE ?", "LIKE")
			assertContains(t, sqlSelectQuery, "`u`.`age` NOT BETWEEN ? AND ?", "NOT BETWEEN")
			assertContains(t, sqlSelectQuery, "`u`.`age` <> ?", "NOTEQUAL")
			assertContains(t, sqlSelectQuery, "NOT EXISTS (SELECT 1 FROM `users` AS `u`)", "NOT EXISTS")
			assertContains(t, sqlSelectQuery, "LOWER(`u`.`email`) NOT LIKE LOWER(?)", "NOT ILIKE")
			assertContains(t, sqlSelectQuery, "`u`.`status` NOT IN (?, ?)", "NOT IN")
			assertContains(t, sqlSelectQuery, "`u`.`email` NOT LIKE ?", "NOT LIKE")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"u"."age" BETWEEN $1 AND $2`, "BETWEEN")
			assertContains(t, sqlSelectQuery, `"u"."age" = $1`, "EQUAL")
			assertContains(t, sqlSelectQuery, `EXISTS (SELECT 1 FROM "users" AS "u")`, "EXISTS")
			assertContains(t, sqlSelectQuery, `"u"."age" > $1`, "GREATER")
			assertContains(t, sqlSelectQuery, `"u"."age" >= $1`, "GREATEREQUAL")
			assertContains(t, sqlSelectQuery, `"u"."email" ILIKE $1`, "ILIKE")
			assertContains(t, sqlSelectQuery, `"u"."status" IN ($1, $2)`, "IN")
			assertContains(t, sqlSelectQuery, `"u"."id" IS NOT NULL`, "IS NOT NULL")
			assertContains(t, sqlSelectQuery, `"u"."id" IS NULL`, "IS NULL")
			assertContains(t, sqlSelectQuery, `"u"."age" < $1`, "LESS")
			assertContains(t, sqlSelectQuery, `"u"."age" <= $1`, "LESSEQUAL")
			assertContains(t, sqlSelectQuery, `"u"."email" LIKE $1`, "LIKE")
			assertContains(t, sqlSelectQuery, `"u"."age" NOT BETWEEN $1 AND $2`, "NOT BETWEEN")
			assertContains(t, sqlSelectQuery, `"u"."age" <> $1`, "NOTEQUAL")
			assertContains(t, sqlSelectQuery, `NOT EXISTS (SELECT 1 FROM "users" AS "u")`, "NOT EXISTS")
			assertContains(t, sqlSelectQuery, `"u"."email" NOT ILIKE $1`, "NOT ILIKE")
			assertContains(t, sqlSelectQuery, `"u"."status" NOT IN ($1, $2)`, "NOT IN")
			assertContains(t, sqlSelectQuery, `"u"."email" NOT LIKE $1`, "NOT LIKE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Distinct(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		stmtSelect := NewSelect().Distinct().
			Field(Users.ID.As("user_id")).
			From(Users.Table)
		sqlSelectQuery, _, _ := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "DISTINCT", "DISTINCT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, "DISTINCT", "DISTINCT")
		}
		t.Logf("dialect: %s sql: %s", supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Function(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		stmtSelect := NewSelect(
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
			DateAdd(Test.CreateAt, Literal("4 DAY")).As("datetime_dateadd"),
			DateDiff(Test.UpdateAt, Test.CreateAt).As("datetime_datediff"),
			DateSub(Test.CreateAt, Literal("4 DAY")).As("datetime_datesub"),
			Day(Test.CreateAt).As("datetime_day"),
			DayName(Test.CreateAt).As("datetime_dayname"),
			Hour(Test.CreateAt).As("datetime_hour"),
			Minute(Test.CreateAt).As("datetime_minute"),
			Month(Test.CreateAt).As("datetime_month"),
			MonthName(Test.CreateAt).As("datetime_monthname"),
			Now().As("datetime_now"),
			Quarter(Test.CreateAt).As("datetime_quarter"),
			Second(Test.CreateAt).As("datetime_second"),
			TimeAdd(Test.CreateAt, Literal("4 HOUR")).As("datetime_timeadd"),
			TimeDiff(Test.UpdateAt, Test.CreateAt).As("datetime_timediff"),
			TimeSub(Test.CreateAt, Literal("4 HOUR")).As("datetime_timesub"),
			Week(Test.CreateAt).As("datetime_week"),
			Year(Test.CreateAt).As("datetime_year"),
			// Функции обмена данными
			JsonArray(Users.Data, Value("test"), Value("test")).As("json_jsonarray"),
			JsonArrayAgg(Users.Data).As("json_jsonarrayagg"),
			JsonContains(Users.Data, Value(`{"theme":"dark"}`)).As("json_jsoncontains"),
			JsonExtract(Users.Data, JsonGroup(JsonPath(JsonKey("col"), JsonIndex(0), JsonKey("name"))), TypeString).As("json_jsonextract"),
			JsonObject(JsonPair(JsonKey("users"), Count(Users.Data, false))).As("json_jsonobject"),
			JsonObjectAgg(Users.Data, Test.Number).As("json_jsonobjectagg"),
			JsonRemove(Users.Data, JsonGroup(JsonPath(JsonKey("temp"))), JsonGroup(JsonPath(JsonKey("session")))).As("json_jsonremove"),
			JsonSet(Users.Data, JsonGroup(JsonPath(JsonKey("temp")), Value(0)), JsonGroup(JsonPath(JsonKey("session")), Value("active"))).As("json_jsonset"),
			JsonType(Users.Data).As("json_jsontype"),
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
			Log(Test.Number, Value(3)).As("math_log"),
			Mod(Test.Number, Value(3)).As("math_mod"),
			Pi().As("math_pi"),
			Power(Test.Number, Value(3)).As("math_power"),
			Rand().As("math_rand"),
			Round(Test.Number, Value(3)).As("math_round"),
			Sin(Test.Number).As("math_sin"),
			Sqrt(Test.Number).As("math_sqrt"),
			Tan(Test.Number).As("math_tan"),
			Trunc(Test.Number, Value(3)).As("math_trunc"),
			// Функции строковые
			Concat(Test.String, Value("old"), Value("new")).As("string_concat"),
			ConcatWs(Value("_"), Test.String, Value("old"), Value("new")).As("string_concatws"),
			LeftString(Test.String, Value(2)).As("string_lstr"),
			Lower(Test.String).As("string_lower"),
			LPad(Test.String, Value(2), Value(",")).As("string_lpad"),
			LTrim(Test.String).As("string_ltrim"),
			Repeat(Test.String, Value(3)).As("string_repeat"),
			Replace(Test.String, Value("old"), Value("new")).As("string_replace"),
			Reverse(Test.String).As("string_reverse"),
			RightString(Test.String, Value(2)).As("string_rstr"),
			RPad(Test.String, Value(2), Value(",")).As("string_rpad"),
			RTrim(Test.String).As("string_rtrim"),
			SubString(Test.String, Value(0), Value(2)).As("string_substring"),
			Trim(Test.String).As("string_trim"),
			Upper(Test.String).As("string_upper"),
			// Функции ранжирующие
			RowNumber().Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("ranking_rownumber"),
			Rank().Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("ranking_rank"),
			DenseRank().Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("ranking_denserank"),
			PercentRank().Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("ranking_percentrank"),
			CumeDist().Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("ranking_cumedist"),
			NTile(4).Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("ranking_ntile"),
		).
			From(Test.Table)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			// Функции агрегатные
			assertContains(t, sqlSelectQuery, "AVG(`t`.`number`)", "AVG")
			assertContains(t, sqlSelectQuery, "BIT_AND(`t`.`number`)", "BITAND")
			assertContains(t, sqlSelectQuery, "BIT_OR(`t`.`number`)", "BITOR")
			assertContains(t, sqlSelectQuery, "BIT_XOR(`t`.`number`)", "BITXOR")
			assertContains(t, sqlSelectQuery, "COUNT(`t`.`string`)", "COUNT")
			assertContains(t, sqlSelectQuery, "GROUP_CONCAT(`t`.`string` SEPARATOR ',')", "GROUPCONCAT")
			assertContains(t, sqlSelectQuery, "MAX(`t`.`number`)", "MAX")
			assertContains(t, sqlSelectQuery, "MIN(`t`.`number`)", "MIN")
			assertContains(t, sqlSelectQuery, "STDDEV(`t`.`number`)", "STDDEV")
			assertContains(t, sqlSelectQuery, "SUM(`t`.`number`)", "SUM")
			assertContains(t, sqlSelectQuery, "VARIANCE(`t`.`number`)", "VARIANCE")
			// Функции аналитические
			assertContains(t, sqlSelectQuery, "FIRST_VALUE(`t`.`name`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC", "FIRSTVALUE")
			assertContains(t, sqlSelectQuery, "LAG(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)", "LAG")
			assertContains(t, sqlSelectQuery, "LAST_VALUE(`t`.`name`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)", "LASTVALUE")
			assertContains(t, sqlSelectQuery, "LEAD(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)", "LEAD")
			assertContains(t, sqlSelectQuery, "NTH_VALUE(`t`.`name`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)", "NTHVALUE")
			// Функции условий
			assertContains(t, sqlSelectQuery, "CASE WHEN `t`.`number` < ? THEN ? ELSE ? END", "CASE")
			assertContains(t, sqlSelectQuery, "COALESCE(`t`.`createat`, `t`.`updateat`)", "COALESCE")
			assertContains(t, sqlSelectQuery, "GREATEST(`t`.`createat`, `t`.`updateat`)", "GREATEST")
			assertContains(t, sqlSelectQuery, "LEAST(`t`.`createat`, `t`.`updateat`)", "LEAST")
			assertContains(t, sqlSelectQuery, "NULLIF(`t`.`createat`, `t`.`updateat`)", "NULLIF")
			// Функции конвертации
			assertContains(t, sqlSelectQuery, "CAST(`t`.`number` AS CHAR)", "CAST")
			assertContains(t, sqlSelectQuery, "CHAR_LENGTH(`t`.`string`)", "CHARLENGTH")
			assertContains(t, sqlSelectQuery, "DATE_FORMAT(`t`.`createat`, '%Y-%m-%d')", "DATEFORMAT")
			assertContains(t, sqlSelectQuery, "DEGREES(`t`.`number`)", "DEGREES")
			assertContains(t, sqlSelectQuery, "LENGTH(`t`.`string`)", "LENGTH")
			assertContains(t, sqlSelectQuery, "POSITION(? IN `t`.`string`)", "POSITION")
			assertContains(t, sqlSelectQuery, "RADIANS(`t`.`number`)", "RADIANS")
			// Функции даты и времени
			assertContains(t, sqlSelectQuery, "CURDATE()", "CURDATE")
			assertContains(t, sqlSelectQuery, "CURTIME()", "CURTIME")
			assertContains(t, sqlSelectQuery, "DATE_ADD(`t`.`createat`, INTERVAL '4 DAY')", "DATEADD")
			assertContains(t, sqlSelectQuery, "DATEDIFF(`t`.`updateat`, `t`.`createat`)", "DATEDIFF")
			assertContains(t, sqlSelectQuery, "DATE_SUB(`t`.`createat`, INTERVAL '4 DAY')", "DATESUB")
			assertContains(t, sqlSelectQuery, "DAY(`t`.`createat`)", "DAY")
			assertContains(t, sqlSelectQuery, "DAYNAME(`t`.`createat`)", "DAYNAME")
			assertContains(t, sqlSelectQuery, "HOUR(`t`.`createat`)", "HOUR")
			assertContains(t, sqlSelectQuery, "MINUTE(`t`.`createat`)", "MINUTE")
			assertContains(t, sqlSelectQuery, "MONTH(`t`.`createat`)", "MONTH")
			assertContains(t, sqlSelectQuery, "MONTHNAME(`t`.`createat`)", "MONTHNAME")
			assertContains(t, sqlSelectQuery, "NOW()", "NOW")
			assertContains(t, sqlSelectQuery, "QUARTER(`t`.`createat`)", "QUARTER")
			assertContains(t, sqlSelectQuery, "SECOND(`t`.`createat`)", "SECOND")
			assertContains(t, sqlSelectQuery, "TIME_ADD(`t`.`createat`, INTERVAL '4 HOUR')", "TIMEADD")
			assertContains(t, sqlSelectQuery, "TIMEDIFF(`t`.`updateat`, `t`.`createat`)", "TIMEDIFF")
			assertContains(t, sqlSelectQuery, "TIME_SUB(`t`.`createat`, INTERVAL '4 HOUR')", "TIMESUB")
			assertContains(t, sqlSelectQuery, "WEEK(`t`.`createat`)", "WEEK")
			assertContains(t, sqlSelectQuery, "YEAR(`t`.`createat`)", "YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, "JSON_ARRAY(`u`.`data`, ?, ?)", "JSONARRAY")
			assertContains(t, sqlSelectQuery, "JSON_ARRAYAGG(`u`.`data`)", "JSONARRAYAGG")
			assertContains(t, sqlSelectQuery, "JSON_CONTAINS(`u`.`data`, ?)", "JSONCONTAINS")
			assertContains(t, sqlSelectQuery, "(`u`.`data` ->> '$.col[0].name')", "JSONEXTRACT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECT('users', COUNT(`u`.`data`))", "JSONOBJECT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECTAGG(`u`.`data`, `t`.`number`)", "JSONOBJECTAGG")
			assertContains(t, sqlSelectQuery, "JSON_REMOVE(`u`.`data`, '$.temp', '$.session')", "JSONREMOVE")
			assertContains(t, sqlSelectQuery, "JSON_SET(`u`.`data`, '$.temp', ?, '$.session', ?)", "JSONSET")
			assertContains(t, sqlSelectQuery, "JSON_TYPE(`u`.`data`)", "JSONTYPE")
			// Функции математические
			assertContains(t, sqlSelectQuery, "ABS(`t`.`number`)", "ABS")
			assertContains(t, sqlSelectQuery, "ACOS(`t`.`number`)", "ACOS")
			assertContains(t, sqlSelectQuery, "ASIN(`t`.`number`)", "ASIN")
			assertContains(t, sqlSelectQuery, "ATAN(`t`.`number`)", "ATAN")
			assertContains(t, sqlSelectQuery, "ATAN2(`t`.`y`, `t`.`x`)", "ATAN2")
			assertContains(t, sqlSelectQuery, "CBRT(`t`.`number`)", "CBRT")
			assertContains(t, sqlSelectQuery, "CEILING(`t`.`number`)", "CEIL")
			assertContains(t, sqlSelectQuery, "COS(`t`.`number`)", "COS")
			assertContains(t, sqlSelectQuery, "EXP(`t`.`number`)", "EXP")
			assertContains(t, sqlSelectQuery, "FLOOR(`t`.`number`)", "FLOOR")
			assertContains(t, sqlSelectQuery, "LN(`t`.`number`)", "LN")
			assertContains(t, sqlSelectQuery, "LOG(`t`.`number`, ?)", "LOG")
			assertContains(t, sqlSelectQuery, "MOD(`t`.`number`, ?)", "MOD")
			assertContains(t, sqlSelectQuery, "PI()", "PI")
			assertContains(t, sqlSelectQuery, "POWER(`t`.`number`, ?)", "POWER")
			assertContains(t, sqlSelectQuery, "RAND()", "RAND")
			assertContains(t, sqlSelectQuery, "ROUND(`t`.`number`, ?)", "ROUND")
			assertContains(t, sqlSelectQuery, "SIN(`t`.`number`)", "SIN")
			assertContains(t, sqlSelectQuery, "SQRT(`t`.`number`)", "SQRT")
			assertContains(t, sqlSelectQuery, "TAN(`t`.`number`)", "TAN")
			assertContains(t, sqlSelectQuery, "TRUNCATE(`t`.`number`, ?)", "TRUNC")
			// Функции ранжирующие
			assertContains(t, sqlSelectQuery, "ROW_NUMBER() OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "ROWNUMBER")
			assertContains(t, sqlSelectQuery, "RANK() OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "RANK")
			assertContains(t, sqlSelectQuery, "DENSE_RANK() OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "DENSERANK")
			assertContains(t, sqlSelectQuery, "PERCENT_RANK() OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "PERCENTRANK")
			assertContains(t, sqlSelectQuery, "CUME_DIST() OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "CUMEDIST")
			assertContains(t, sqlSelectQuery, "NTILE(4) OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "NTILE")
			// Функции строковые
			assertContains(t, sqlSelectQuery, "CONCAT(`t`.`string`, ?, ?)", "CONCAT")
			assertContains(t, sqlSelectQuery, "CONCAT_WS(?, `t`.`string`, ?, ?)", "CONCATWS")
			assertContains(t, sqlSelectQuery, "LEFT(`t`.`string`, ?)", "LEFTSTRING")
			assertContains(t, sqlSelectQuery, "LOWER(`t`.`string`)", "LOWER")
			assertContains(t, sqlSelectQuery, "LPAD(`t`.`string`, ?, ?)", "LPAD")
			assertContains(t, sqlSelectQuery, "LTRIM(`t`.`string`)", "LTRIM")
			assertContains(t, sqlSelectQuery, "REPEAT(`t`.`string`, ?)", "REPEAT")
			assertContains(t, sqlSelectQuery, "REPLACE(`t`.`string`, ?, ?)", "REPLACE")
			assertContains(t, sqlSelectQuery, "REVERSE(`t`.`string`)", "REVERSE")
			assertContains(t, sqlSelectQuery, "RIGHT(`t`.`string`, ?)", "RIGHTSTRING")
			assertContains(t, sqlSelectQuery, "RPAD(`t`.`string`, ?, ?)", "RPAD")
			assertContains(t, sqlSelectQuery, "RTRIM(`t`.`string`)", "RTRIM")
			assertContains(t, sqlSelectQuery, "SUBSTRING(`t`.`string`, ?, ?)", "SUBSTRING")
			assertContains(t, sqlSelectQuery, "TRIM(`t`.`string`)", "TRIM")
			assertContains(t, sqlSelectQuery, "UPPER(`t`.`string`)", "UPPER")
		case DialectPostgreSQL:
			// Функции агрегатные
			assertContains(t, sqlSelectQuery, `AVG("t"."number")`, "AVG")
			assertContains(t, sqlSelectQuery, `BIT_AND("t"."number")`, "BITAND")
			assertContains(t, sqlSelectQuery, `BIT_OR("t"."number")`, "BITOR")
			assertContains(t, sqlSelectQuery, `BIT_XOR("t"."number")`, "BITXOR")
			assertContains(t, sqlSelectQuery, `COUNT("t"."string")`, "COUNT")
			assertContains(t, sqlSelectQuery, `STRING_AGG("t"."string", ',')`, "GROUPCONCAT")
			assertContains(t, sqlSelectQuery, `MAX("t"."number")`, "MAX")
			assertContains(t, sqlSelectQuery, `MIN("t"."number")`, "MIN")
			assertContains(t, sqlSelectQuery, `STDDEV_SAMP("t"."number")`, "STDDEV")
			assertContains(t, sqlSelectQuery, `SUM("t"."number")`, "SUM")
			assertContains(t, sqlSelectQuery, `VAR_SAMP("t"."number")`, "VARIANCE")
			// Функции аналитические
			assertContains(t, sqlSelectQuery, `FIRST_VALUE("t"."name") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FIRSTVALUE")
			assertContains(t, sqlSelectQuery, `LAG("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)`, "LAG")
			assertContains(t, sqlSelectQuery, `LAST_VALUE("t"."name") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)`, "LASTVALUE")
			assertContains(t, sqlSelectQuery, `LEAD("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)`, "LEAD")
			assertContains(t, sqlSelectQuery, `NTH_VALUE("t"."name", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)`, "NTHVALUE")
			// Функции условий
			assertContains(t, sqlSelectQuery, `CASE WHEN "t"."number" < $1 THEN $2 ELSE $3 END`, "CASE")
			assertContains(t, sqlSelectQuery, `COALESCE("t"."createat", "t"."updateat")`, "COALESCE")
			assertContains(t, sqlSelectQuery, `GREATEST("t"."createat", "t"."updateat")`, "GREATEST")
			assertContains(t, sqlSelectQuery, `LEAST("t"."createat", "t"."updateat")`, "LEAST")
			assertContains(t, sqlSelectQuery, `NULLIF("t"."createat", "t"."updateat")`, "NULLIF")
			// Функции конвертации
			assertContains(t, sqlSelectQuery, `CAST("t"."number" AS VARCHAR) `, "CAST")
			assertContains(t, sqlSelectQuery, `CHAR_LENGTH("t"."string")`, "CHARLENGTH")
			assertContains(t, sqlSelectQuery, `TO_CHAR("t"."createat", '%Y-%m-%d')`, "DATEFORMAT")
			assertContains(t, sqlSelectQuery, `DEGREES("t"."number")`, "DEGREES")
			assertContains(t, sqlSelectQuery, `LENGTH("t"."string")`, "LENGTH")
			assertContains(t, sqlSelectQuery, `POSITION($1 IN "t"."string")`, "POSITION")
			assertContains(t, sqlSelectQuery, `RADIANS("t"."number")`, "RADIANS")
			// Функции даты и времени
			assertContains(t, sqlSelectQuery, `CURRENT_DATE`, "CURDATE")
			assertContains(t, sqlSelectQuery, `CURRENT_TIME`, "CURTIME")
			assertContains(t, sqlSelectQuery, `("t"."createat" + INTERVAL '4 DAY')`, "DATEADD")
			assertContains(t, sqlSelectQuery, `DATE_PART('day', "t"."updateat" - "t"."createat")`, "DATEDIFF")
			assertContains(t, sqlSelectQuery, `("t"."createat" - INTERVAL '4 DAY')`, "DATESUB")
			assertContains(t, sqlSelectQuery, `EXTRACT(DAY FROM "t"."createat")`, "DAY")
			assertContains(t, sqlSelectQuery, `TO_CHAR("t"."createat", 'Day')`, "DAYNAME")
			assertContains(t, sqlSelectQuery, `EXTRACT(HOUR FROM "t"."createat")`, "HOUR")
			assertContains(t, sqlSelectQuery, `EXTRACT(MINUTE FROM "t"."createat")`, "MINUTE")
			assertContains(t, sqlSelectQuery, `EXTRACT(MONTH FROM "t"."createat")`, "MONTH")
			assertContains(t, sqlSelectQuery, `TO_CHAR("t"."createat", 'Month')`, "MONTHNAME")
			assertContains(t, sqlSelectQuery, `CURRENT_TIMESTAMP`, "NOW")
			assertContains(t, sqlSelectQuery, `EXTRACT(QUARTER FROM "t"."createat")`, "QUARTER")
			assertContains(t, sqlSelectQuery, `EXTRACT(SECOND FROM "t"."createat")`, "SECOND")
			assertContains(t, sqlSelectQuery, `("t"."createat" + INTERVAL '4 HOUR')`, "TIMEADD")
			assertContains(t, sqlSelectQuery, `DATE_PART('time', "t"."updateat" - "t"."createat")`, "TIMEDIFF")
			assertContains(t, sqlSelectQuery, `("t"."createat" - INTERVAL '4 HOUR')`, "TIMESUB")
			assertContains(t, sqlSelectQuery, `EXTRACT(WEEK FROM "t"."createat")`, "WEEK")
			assertContains(t, sqlSelectQuery, `EXTRACT(YEAR FROM "t"."createat")`, "YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, `JSON_ARRAY("u"."data", $1, $2)`, "JSONARRAY")
			assertContains(t, sqlSelectQuery, `JSON_AGG("u"."data")`, "JSONARRAYAGG")
			assertContains(t, sqlSelectQuery, `("u"."data" @> $1)`, "JSONCONTAINS")
			assertContains(t, sqlSelectQuery, `("u"."data" #>> '{col,0,name}')`, "JSONEXTRACT")
			assertContains(t, sqlSelectQuery, `JSON_BUILD_OBJECT('users', COUNT("u"."data"))`, "JSONOBJECT")
			assertContains(t, sqlSelectQuery, `JSON_OBJECT_AGG("u"."data", "t"."number")`, "JSONOBJECTAGG")
			assertContains(t, sqlSelectQuery, `("u"."data" - '{temp}' - '{session}')`, "JSONREMOVE")
			//assertContains(t, sqlSelectQuery, `jsonb_set`, "JSONSET")
			assertContains(t, sqlSelectQuery, `jsonb_typeof("u"."data")`, "JSONTYPE")
			// Функции математические
			assertContains(t, sqlSelectQuery, `ABS("t"."number")`, "ABS")
			assertContains(t, sqlSelectQuery, `ACOS("t"."number")`, "ACOS")
			assertContains(t, sqlSelectQuery, `ASIN("t"."number")`, "ASIN")
			assertContains(t, sqlSelectQuery, `ATAN("t"."number")`, "ATAN")
			assertContains(t, sqlSelectQuery, `ATAN2("t"."y", "t"."x")`, "ATAN2")
			assertContains(t, sqlSelectQuery, `CBRT("t"."number")`, "CBRT")
			assertContains(t, sqlSelectQuery, `CEIL("t"."number")`, "CEIL")
			assertContains(t, sqlSelectQuery, `COS("t"."number")`, "COS")
			assertContains(t, sqlSelectQuery, `EXP("t"."number")`, "EXP")
			assertContains(t, sqlSelectQuery, `FLOOR("t"."number")`, "FLOOR")
			assertContains(t, sqlSelectQuery, `LN("t"."number")`, "LN")
			assertContains(t, sqlSelectQuery, `LOG("t"."number", $1)`, "LOG")
			assertContains(t, sqlSelectQuery, `MOD("t"."number", $1)`, "MOD")
			assertContains(t, sqlSelectQuery, `PI()`, "PI")
			assertContains(t, sqlSelectQuery, `POWER("t"."number", $1)`, "POWER")
			assertContains(t, sqlSelectQuery, `RANDOM`, "RAND")
			assertContains(t, sqlSelectQuery, `ROUND("t"."number", $1)`, "ROUND")
			assertContains(t, sqlSelectQuery, `SIN("t"."number")`, "SIN")
			assertContains(t, sqlSelectQuery, `SQRT("t"."number")`, "SQRT")
			assertContains(t, sqlSelectQuery, `TAN("t"."number")`, "TAN")
			assertContains(t, sqlSelectQuery, `TRUNC("t"."number", $1)`, "TRUNC")
			// Функции ранжирующие
			assertContains(t, sqlSelectQuery, `ROW_NUMBER() OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "ROWNUMBER")
			assertContains(t, sqlSelectQuery, `RANK() OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "RANK")
			assertContains(t, sqlSelectQuery, `DENSE_RANK() OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "DENSERANK")
			assertContains(t, sqlSelectQuery, `PERCENT_RANK() OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "PERCENTRANK")
			assertContains(t, sqlSelectQuery, `CUME_DIST() OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "CUMEDIST")
			assertContains(t, sqlSelectQuery, `NTILE(4) OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "NTILE")
			// Функции строковые
			assertContains(t, sqlSelectQuery, `CONCAT("t"."string", $1, $2)`, "CONCAT")
			assertContains(t, sqlSelectQuery, `CONCAT_WS($1, "t"."string", $2, $3)`, "CONCATWS")
			assertContains(t, sqlSelectQuery, `LEFT("t"."string", $1)`, "LEFTSTRING")
			assertContains(t, sqlSelectQuery, `LOWER("t"."string")`, "LOWER")
			assertContains(t, sqlSelectQuery, `LPAD("t"."string", $1, $2)`, "LPAD")
			assertContains(t, sqlSelectQuery, `LTRIM("t"."string")`, "LTRIM")
			assertContains(t, sqlSelectQuery, `REPEAT("t"."string", $1)`, "REPEAT")
			assertContains(t, sqlSelectQuery, `REPLACE("t"."string", $1, $2)`, "REPLACE")
			assertContains(t, sqlSelectQuery, `REVERSE("t"."string")`, "REVERSE")
			assertContains(t, sqlSelectQuery, `RIGHT("t"."string", $1)`, "RIGHTSTRING")
			assertContains(t, sqlSelectQuery, `RPAD("t"."string", $1, $2)`, "RPAD")
			assertContains(t, sqlSelectQuery, `RTRIM("t"."string")`, "RTRIM")
			assertContains(t, sqlSelectQuery, `SUBSTRING("t"."string", $1, $2)`, "SUBSTRING")
			assertContains(t, sqlSelectQuery, `TRIM("t"."string")`, "TRIM")
			assertContains(t, sqlSelectQuery, `UPPER("t"."string")`, "UPPER")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_GroupBy(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		stmtSelect := NewSelect(
			Users.DepartmentID,
			Users.Status,
			Count(Users.ID, false).As("user_count"),
			Sum(Users.Age, false).As("total_age"),
			Avg(Users.Age, false).As("avg_age"),
			Min(Users.Age, false).As("min_age"),
			Max(Users.Age, false).As("max_age"),
		).
			From(Users.Table).
			Where(
				And(
					GreaterEqual(Users.Age, Value(18)),
					NotEqual(Users.Status, Value("deleted")),
				),
			).
			GroupBy(
				Users.DepartmentID,
				Users.Status,
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`u`.`department_id`, `u`.`status`", "LIST")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"u"."department_id", "u"."status"`, "LIST")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Having(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		queryMain := NewSelect(
			Users.DepartmentID,
			Count(Users.ID, false).As("user_count"),
			Sum(Users.Age, false).As("total_age"),
			Avg(Users.Age, false).As("avg_age"),
		).
			From(Users.Table).
			GroupBy(Users.DepartmentID).
			Having(
				And(
					Greater(Count(Users.ID, false), Value[int64](5)),
					Greater(Sum(Users.Age, false), Value(100)),
					Between(Avg(Users.Age, false), Value(25), Value(40)),
					In(Users.DepartmentID, Array[int64](1, 2, 3)),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "COUNT(`u`.`id`) > ?", "COUNT")
			assertContains(t, sqlSelectQuery, "SUM(`u`.`age`) > ?", "SUM")
			assertContains(t, sqlSelectQuery, "AVG(`u`.`age`) BETWEEN ? AND ?", "BETWEEN")
			assertContains(t, sqlSelectQuery, "`u`.`department_id` IN (?, ?, ?)", "IN")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `COUNT("u"."id") > $1`, "COUNT")
			assertContains(t, sqlSelectQuery, `SUM("u"."age") > $1`, "SUM")
			assertContains(t, sqlSelectQuery, `AVG("u"."age") BETWEEN $1 AND $2`, "BETWEEN")
			assertContains(t, sqlSelectQuery, `"u"."department_id" IN ($1, $2, $3)`, "IN")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Join(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		stmtSelect := NewSelect(Users.ID.As("user_id")).
			From(Users.Table).
			Join(
				Cross(Orders.Table),
				Full(Orders.Table, Equal(Users.ID, Orders.UserID)),
				FullOuter(Orders.Table, Equal(Users.ID, Orders.UserID)),
				Inner(Orders.Table, Equal(Users.ID, Orders.UserID)),
				Left(Orders.Table, Equal(Users.ID, Orders.UserID)),
				LeftOuter(Orders.Table, Equal(Users.ID, Orders.UserID)),
				Right(Orders.Table, Equal(Users.ID, Orders.UserID)),
				RightOuter(Orders.Table, Equal(Users.ID, Orders.UserID)),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "CROSS JOIN `orders` AS `o`", "CROSS")
			assertContains(t, sqlSelectQuery, "FULL JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "FULL")
			assertContains(t, sqlSelectQuery, "FULL OUTER JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "FULL OUTER")
			assertContains(t, sqlSelectQuery, "INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "INNER")
			assertContains(t, sqlSelectQuery, "LEFT JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "LEFT")
			assertContains(t, sqlSelectQuery, "LEFT OUTER JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "LEFT OUTER")
			assertContains(t, sqlSelectQuery, "RIGHT JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "RIGHT")
			assertContains(t, sqlSelectQuery, "RIGHT OUTER JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "RIGHT OUTER")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `CROSS JOIN "orders" AS "o"`, "CROSS")
			assertContains(t, sqlSelectQuery, `FULL JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "FULL")
			assertContains(t, sqlSelectQuery, `FULL OUTER JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "FULL OUTER")
			assertContains(t, sqlSelectQuery, `INNER JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "INNER")
			assertContains(t, sqlSelectQuery, `LEFT JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "LEFT")
			assertContains(t, sqlSelectQuery, `LEFT OUTER JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "LEFT OUTER")
			assertContains(t, sqlSelectQuery, `RIGHT JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "RIGHT")
			assertContains(t, sqlSelectQuery, `RIGHT OUTER JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "RIGHT OUTER")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Limit(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		stmtSelect := NewSelect(
			Users.ID,
			Users.Name,
			Users.Age,
		).
			From(Users.Table).
			Where(Equal(Users.Status, Value("active"))).
			Limit(10)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "LIMIT ?", "LIMIT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, "LIMIT $1", "LIMIT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Logical(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		stmtSelect := NewSelect(Users.ID.As("user_id")).
			From(Users.Table).
			Where(
				And(
					And(
						Equal(Users.ID, Value[int64](0)),
						Equal(Users.ID, Value[int64](10)),
					),
					Or(
						Equal(Users.ID, Value[int64](20)),
						Equal(Users.ID, Value[int64](100)),
					),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`u`.`id` = ? AND `u`.`id` = ?", "AND")
			assertContains(t, sqlSelectQuery, "`u`.`id` = ? OR `u`.`id` = ?", "OR")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"u"."id" = $1 AND "u"."id" = $2`, "AND")
			assertContains(t, sqlSelectQuery, `"u"."id" = $3 OR "u"."id" = $4`, "OR")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Offset(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		stmtSelect := NewSelect(Users.ID, Users.Name, Users.Age).
			From(Users.Table).
			Where(Equal(Users.Status, Value("active"))).
			Offset(20)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "OFFSET ?", "OFFSET")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, "OFFSET $1", "OFFSET")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_OrderBy(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		stmtSelect := NewSelect(Users.ID.As("user_id"), Users.Name.As("user_name")).
			From(Users.Table).
			OrderBy(
				Asc(Users.Name),
				Desc(Users.ID),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`u`.`name` ASC", "ASC")
			assertContains(t, sqlSelectQuery, "`u`.`id` DESC", "DESC")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"u"."name" ASC`, "ASC")
			assertContains(t, sqlSelectQuery, `"u"."id" DESC`, "DESC")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Subquery(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		querySub := NewSelect(Count(Orders.ID, false)).
			From(Orders.Table).
			Where(
				Equal(Orders.UserID, Users.ID),
			)
		queryInSub := NewSelect(Orders.UserID).
			From(Orders.Table).
			Where(
				Greater(Orders.Amount, Value(1000)),
			)
		queryExistsSub := NewSelect(Orders.ID).
			From(Orders.Table).
			Where(
				Equal(Orders.UserID, Users.ID),
			)
		stmtSelect := NewSelect(
			Users.ID,
			Users.Name,
			Subquery[int64](querySub).As("SUB"),
		).
			From(Users.Table).
			Where(
				And(
					In(Users.ID, Subquery[int64](queryInSub)),
					Exists(Subquery[int64](queryExistsSub)),
				),
			)

		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "AS", "AS (subquery1)")
			assertContains(t, sqlSelectQuery, "IN", "IN (subquery2)")
			assertContains(t, sqlSelectQuery, "EXISTS", "EXISTS (subquery3)")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, "AS", "AS (subquery1)")
			assertContains(t, sqlSelectQuery, "IN", "IN (subquery2)")
			assertContains(t, sqlSelectQuery, "EXISTS", "EXISTS (subquery3)")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Unions(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		queryUnion := NewSelect(
			Users.Name.As("person_name"),
			Users.Email.As("contact_email"),
		).
			From(Users.Table).
			Where(
				And(
					Equal(Users.Status, Value("active")),
					Like(Users.Email, Value("%@domain.ltd")),
				),
			)
		queryUnionAll := NewSelect(
			Products.Name.As("person_name"),
			Products.Name.As("contact_email"),
		).
			From(Products.Table).
			Where(
				Greater(Products.Count, Value(10)),
			)
		queryUnionExcept := NewSelect(
			Categories.Name.As("person_name"),
			Categories.Type.As("contact_email"),
		).
			From(Categories.Table).
			Where(
				In(Categories.Type, Array("premium", "standard", "basic")),
			)
		queryUnionIntersect := NewSelect(
			Departments.Name.As("person_name"),
			Departments.Name.As("contact_email"),
		).
			From(Departments.Table).
			Where(
				IsNull(Departments.ParentID),
			)
		stmtSelect := NewSelect(
			Users.ID.As("entity_id"),
			Users.Name.As("entity_name"),
			Users.Email.As("entity_email"),
			Users.Status.As("entity_status"),
		).
			From(Users.Table).
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
			assertContains(t, sqlSelectQuery, "UNION", "UNION")
			assertContains(t, sqlSelectQuery, "UNION ALL", "UNION ALL")
			assertContains(t, sqlSelectQuery, "EXCEPT", "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, "INTERSECT", "UNION INTERSECT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_Where(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		existsSub := NewSelect(Levels.ID).
			From(Levels.Table).
			Where(
				Equal(Levels.UserID, Users.ID),
			)
		stmtSelect := NewSelect(Users.ID).
			From(Users.Table).
			Where(
				And(
					Equal(Users.Status, Value("active")),
					Greater(Users.Age, Value(18)),
					Less(Users.Age, Value(65)),
					In(Users.DepartmentID, Array[int64](1, 2, 3)),
					Like(Users.Email, Value("%@company.com")),
					Exists(Subquery[int](existsSub)),
					NotEqual(Users.Status, Value("deleted")),
					IsNotNull(Users.Email),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`u`.`status` = ?", "EQUAL")
			assertContains(t, sqlSelectQuery, "`u`.`age` > ?", "GREATER")
			assertContains(t, sqlSelectQuery, "`u`.`age` < ?", "LESS")
			assertContains(t, sqlSelectQuery, "`u`.`department_id` IN (?, ?, ?)", "IN")
			assertContains(t, sqlSelectQuery, "`u`.`email` LIKE ?", "LIKE")
			assertContains(t, sqlSelectQuery, "EXISTS", "EXISTS")
			assertContains(t, sqlSelectQuery, "`u`.`status` <> ?", "NOTEQUAL")
			assertContains(t, sqlSelectQuery, "`u`.`email` IS NOT NULL", "IS NOT NULL")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"u"."status" = $1`, "EQUAL")
			assertContains(t, sqlSelectQuery, `"u"."age" > $1`, "GREATER")
			assertContains(t, sqlSelectQuery, `"u"."age" < $1`, "LESS")
			assertContains(t, sqlSelectQuery, `"u"."department_id" IN ($1, $2, $3)`, "IN")
			assertContains(t, sqlSelectQuery, `"u"."email" LIKE $1`, "LIKE")
			assertContains(t, sqlSelectQuery, `EXISTS`, "EXISTS")
			assertContains(t, sqlSelectQuery, `"u"."status" <> $1`, "NOTEQUAL")
			assertContains(t, sqlSelectQuery, `"u"."email" IS NOT NULL`, "IS NOT NULL")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Select_With(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		queryWithN := WithN("WithN", NewSelect(
			Orders.UserID,
			Count(Orders.ID, false).As("order_count"),
			Sum(Orders.Amount, false).As("total_spent"),
		).
			From(Orders.Table).
			Where(Greater(Orders.Amount, Value(1000))).
			GroupBy(Orders.UserID),
			"user_id", "order_count", "total_spent",
		)
		queryWithR := WithR("WithR", NewSelect(
			Departments.ID,
			Departments.Name,
			Departments.ParentID,
		).
			From(Departments.Table).
			Where(IsNull(Departments.ParentID)).
			Unions(
				UnionAll(NewSelect(
					Departments.ID,
					Departments.Name,
					Departments.ParentID,
				).
					From(Departments.Table).
					Join(Inner(CTE("WithR", "dh"), Equal(Departments.ParentID, Departments.ID)))),
			),
			"dept_id", "dept_name", "parent_id",
		)
		stmtSelect := NewSelect(
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
			From(Users.Table).
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
				Left(Query(NewSelect(
					Categories.ID,
					Categories.Name,
					Count(Products.ID, false).As("product_count"),
					Sum(Products.Count, false).As("total_inventory"),
				).
					From(Categories.Table).
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
					Greater(Stats.OrderCount, Value(0)),
					Or(
						And(
							Greater(Users.Age, Value(18)),
							Less(Users.Age, Value(65)),
						),
						Equal(Users.Age, Value(0)),
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
					Greater(Stats.OrderCount, Value(5)),
					Greater(Stats.TotalSpent, Value(5000)),
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
			assertContains(t, sqlSelectQuery, "WithN", "WithN")
			assertContains(t, sqlSelectQuery, "WithR", "WithR")
			assertContains(t, sqlSelectQuery, "RECURSIVE", "RECURSIVE")
			assertContains(t, sqlSelectQuery, "UNION ALL", "UNION ALL for RCTE")
			assertContains(t, sqlSelectQuery, "WITH", "WITH")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, "WithN", "WithN")
			assertContains(t, sqlSelectQuery, "WithR", "WithR")
			assertContains(t, sqlSelectQuery, "RECURSIVE", "RECURSIVE")
			assertContains(t, sqlSelectQuery, "UNION ALL", "UNION ALL for RCTE")
			assertContains(t, sqlSelectQuery, "WITH", "WITH")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Update(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql, _ := NewSQL(supportDialect)
		defer sql.Close()
		queryWith := WithN("user_stats", NewSelect(
			Orders.UserID,
			Count(Orders.ID, false).As("order_count"),
		).
			From(Orders.Table).
			Where(Greater(Orders.Amount, Value(100))).
			GroupBy(Orders.UserID),
			"user_id", "order_count")
		stmtUpdate := NewUpdate(Users.Table).
			Where(
				And(
					Equal(Users.Status, Value("active")),
					Or(
						IsNull(Levels.ID),
						Equal(Levels.Status, Value("expired")),
					),
				),
			).
			Returning(
				Users.ID.As("user_id"),
				Users.Name.As("user_name"),
				Stats.OrderCount.As("total_orders"),
			).
			With(queryWith)
		sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
		switch supportDialect {
		case DialectMySQL:
			//assertContains(t, sqlUpdateQuery, "UPDATE", "UPDATE")
			//assertContains(t, sqlUpdateQuery, "FROM", "FROM")
			//assertContains(t, sqlUpdateQuery, "SET", "SET")
			//assertContains(t, sqlUpdateQuery, "WHERE", "WHERE")
			//assertContains(t, sqlUpdateQuery, "WITH", "WITH")
		case DialectPostgreSQL:
			//assertContains(t, sqlUpdateQuery, "UPDATE", "UPDATE")
			//assertContains(t, sqlUpdateQuery, "FROM", "FROM")
			//assertContains(t, sqlUpdateQuery, "RETURNING", "RETURNING")
			//assertContains(t, sqlUpdateQuery, "SET", "SET")
			//assertContains(t, sqlUpdateQuery, "USING", "USING")
			//assertContains(t, sqlUpdateQuery, "WHERE", "WHERE")
			//assertContains(t, sqlUpdateQuery, "WITH", "WITH")
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
