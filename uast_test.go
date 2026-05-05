package uast

import (
	"regexp"
	"testing"
)

// Публичные функции
func Test_Delete(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryWith := WithN("user_stats", builder.Select(
			Orders.UserID,
			Count(Orders.ID, false).As("order_count"),
		).
			From(Orders.Table).
			Where(Greater(Orders.Amount, Value(100))).
			GroupBy(Orders.UserID),
			"user_id", "order_count")
		queryMain := builder.Delete(Users.Table).
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
		sql, _, _ := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "DELETE", "DELETE")
			assertContains(t, sql, "FROM", "FROM")
			assertContains(t, sql, "JOIN", "JOIN")
			assertContains(t, sql, "WHERE", "WHERE")
			assertContains(t, sql, "WITH", "WITH")
		case DialectPostgreSQL:
			assertContains(t, sql, "DELETE", "DELETE")
			assertContains(t, sql, "FROM", "FROM")
			assertContains(t, sql, "RETURNING", "RETURNING")
			assertContains(t, sql, "USING", "USING")
			assertContains(t, sql, "WHERE", "WHERE")
			assertContains(t, sql, "WITH", "WITH")
		}
	})
}
func Test_Delete_Join(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Delete(Users.Table).
			Join(
				Inner(Orders.Table, Equal(Users.ID, Orders.UserID)),
				Left(Levels.Table, Equal(Users.ID, Levels.UserID)),
			).
			Where(
				Equal(Orders.Status, Value("cancelled")),
			)
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "INNER JOIN", "INNER JOIN")
			assertContains(t, sql, "LEFT JOIN", "LEFT JOIN")
		case DialectPostgreSQL:
			assertContains(t, sql, `"orders" AS "o", "level" AS "l"`, "USING LIST")
			assertContains(t, sql, `"u"."id" = "o"."user_id" AND "u"."id" = "l"."user_id"`, "JOIN CONDITION")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Delete_Returning(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Delete(Users.Table).
			Where(Equal(Users.Status, Value("inactive"))).
			Returning(
				Users.ID.As("user_id"),
				Users.Name.As("user_name"),
				Users.Email.As("user_email"),
			)
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			// Not support
		case DialectPostgreSQL:
			assertContains(t, sql, `"user_id"`, "RETURNING id")
			assertContains(t, sql, `"user_name"`, "RETURNING name")
			assertContains(t, sql, `"user_email"`, "RETURNING email")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Delete_Where(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		existsSub := builder.Select(Levels.ID).
			From(Levels.Table).
			Where(
				Equal(Levels.UserID, Users.ID),
			)
		queryMain := builder.Delete(Users.Table).
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
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "`u`.`status` = ?", "EQUAL")
			assertContains(t, sql, "`u`.`age` > ?", "GREATER")
			assertContains(t, sql, "`u`.`age` < ?", "LESS")
			assertContains(t, sql, "`u`.`department_id` IN (?, ?, ?)", "IN")
			assertContains(t, sql, "`u`.`email` LIKE ?", "LIKE")
			assertContains(t, sql, "EXISTS", "EXISTS")
			assertContains(t, sql, "`u`.`status` <> ?", "NOTEQUAL")
			assertContains(t, sql, "`u`.`email` IS NOT NULL", "IS NOT NULL")
		case DialectPostgreSQL:
			assertContains(t, sql, `"u"."status" = $1`, "EQUAL")
			assertContains(t, sql, `"u"."age" > $1`, "GREATER")
			assertContains(t, sql, `"u"."age" < $1`, "LESS")
			assertContains(t, sql, `"u"."department_id" IN ($1, $2, $3)`, "IN")
			assertContains(t, sql, `"u"."email" LIKE $1`, "LIKE")
			assertContains(t, sql, `EXISTS`, "EXISTS")
			assertContains(t, sql, `"u"."status" <> $1`, "NOTEQUAL")
			assertContains(t, sql, `"u"."email" IS NOT NULL`, "IS NOT NULL")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Delete_With(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		cte := WithN("old_users", builder.Select(Users.ID).
			From(Users.Table).
			Where(Less(Users.Age, Value(18))),
		)
		queryMain := builder.Delete(Users.Table).
			Where(In(Users.ID, Subquery[int64](builder.Select(Column[int64]("old_users", "id"))))).
			With(cte)

		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "WITH", "WITH")
			assertContains(t, sql, "old_users", "CTE")
		case DialectPostgreSQL:
			assertContains(t, sql, "WITH", "WITH")
			assertContains(t, sql, "old_users", "CTE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Insert(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryWith := WithN("user_stats", builder.Select(
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
		queryMain := builder.Insert(Users.Name, Users.Age).
			Into(Users.Table).
			Source(builder.Select(Value("Test User"), Value(35)).
				Where(
					Exists(Subquery[int64](builder.Select(Stats.UserID).
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
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			//assertContains(t, sql, "INSERT", "INSERT")
			//assertContains(t, sql, "INTO", "INTO")
			//assertContains(t, sql, "SOURCE", "SOURCE")
			//assertContains(t, sql, "WITH", "WITH")
		case DialectPostgreSQL:
			//assertContains(t, sql, "INSERT", "INSERT")
			//assertContains(t, sql, "INTO", "INTO")
			//assertContains(t, sql, "RETURNING", "RETURNING")
			//assertContains(t, sql, "SOURCE", "SOURCE")
			//assertContains(t, sql, "WITH", "WITH")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryWith := WithN("user_stats", builder.Select(
			Orders.UserID,
			Count(Orders.ID, false).As("order_count"),
			Sum(Orders.Amount, false).As("total_spent"),
		).
			From(Orders.Table).
			Where(Greater(Orders.Amount, Value(100))).
			GroupBy(Orders.UserID),
			"user_id", "order_count", "total_spent",
		)
		queryMain := builder.Select(
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
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "FROM", "FROM")
			assertContains(t, sql, "GROUP BY", "GROUP BY")
			assertContains(t, sql, "HAVING", "HAVING")
			assertContains(t, sql, "JOIN", "JOIN")
			assertContains(t, sql, "LIMIT", "LIMIT")
			assertContains(t, sql, "OFFSET", "OFFSET")
			assertContains(t, sql, "ORDER BY", "ORDER BY")
			assertContains(t, sql, "SELECT", "SELECT")
			assertContains(t, sql, "WHERE", "WHERE")
			assertContains(t, sql, "WITH", "WITH")
		case DialectPostgreSQL:
			assertContains(t, sql, "FROM", "FROM")
			assertContains(t, sql, "GROUP BY", "GROUP BY")
			assertContains(t, sql, "HAVING", "HAVING")
			assertContains(t, sql, "JOIN", "JOIN")
			assertContains(t, sql, "LIMIT", "LIMIT")
			assertContains(t, sql, "OFFSET", "OFFSET")
			assertContains(t, sql, "ORDER BY", "ORDER BY")
			assertContains(t, sql, "SELECT", "SELECT")
			assertContains(t, sql, "WHERE", "WHERE")
			assertContains(t, sql, "WITH", "WITH")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_Comparison(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Select(Users.ID.As("user_data_id")).
			From(Users.Table).
			Where(
				And(
					Between(Users.Age, Value(18), Value(65)),
					Equal(Users.Age, Value(18)),
					Exists(Subquery[int](builder.Select(ConstIntOne()).From(Users.Table))),
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
					NotExists(Subquery[int](builder.Select(ConstIntOne()).From(Users.Table))),
					NotILike(Users.Email, Value("admin%")),
					NotIn(Users.Status, Array("banned", "deleted")),
					NotLike(Users.Email, Value("%@test.%")),
				))
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "`u`.`age` BETWEEN ? AND ?", "BETWEEN")
			assertContains(t, sql, "`u`.`age` = ?", "EQUAL")
			assertContains(t, sql, "EXISTS (SELECT 1 FROM `users` AS `u`)", "EXISTS")
			assertContains(t, sql, "`u`.`age` > ?", "GREATER")
			assertContains(t, sql, "`u`.`age` >= ?", "GREATEREQUAL")
			assertContains(t, sql, "LOWER(`u`.`email`) LIKE LOWER(?)", "ILIKE")
			assertContains(t, sql, "`u`.`status` IN (?, ?)", "IN")
			assertContains(t, sql, "`u`.`id` IS NOT NULL", "IS NOT NULL")
			assertContains(t, sql, "`u`.`id` IS NULL", "IS NULL")
			assertContains(t, sql, "`u`.`age` < ?", "LESS")
			assertContains(t, sql, "`u`.`age` <= ? ", "LESSEQUAL")
			assertContains(t, sql, "`u`.`email` LIKE ?", "LIKE")
			assertContains(t, sql, "`u`.`age` NOT BETWEEN ? AND ?", "NOT BETWEEN")
			assertContains(t, sql, "`u`.`age` <> ?", "NOTEQUAL")
			assertContains(t, sql, "NOT EXISTS (SELECT 1 FROM `users` AS `u`)", "NOT EXISTS")
			assertContains(t, sql, "LOWER(`u`.`email`) NOT LIKE LOWER(?)", "NOT ILIKE")
			assertContains(t, sql, "`u`.`status` NOT IN (?, ?)", "NOT IN")
			assertContains(t, sql, "`u`.`email` NOT LIKE ?", "NOT LIKE")
		case DialectPostgreSQL:
			assertContains(t, sql, `"u"."age" BETWEEN $1 AND $2`, "BETWEEN")
			assertContains(t, sql, `"u"."age" = $1`, "EQUAL")
			assertContains(t, sql, `EXISTS (SELECT 1 FROM "users" AS "u")`, "EXISTS")
			assertContains(t, sql, `"u"."age" > $1`, "GREATER")
			assertContains(t, sql, `"u"."age" >= $1`, "GREATEREQUAL")
			assertContains(t, sql, `"u"."email" ILIKE $1`, "ILIKE")
			assertContains(t, sql, `"u"."status" IN ($1, $2)`, "IN")
			assertContains(t, sql, `"u"."id" IS NOT NULL`, "IS NOT NULL")
			assertContains(t, sql, `"u"."id" IS NULL`, "IS NULL")
			assertContains(t, sql, `"u"."age" < $1`, "LESS")
			assertContains(t, sql, `"u"."age" <= $1`, "LESSEQUAL")
			assertContains(t, sql, `"u"."email" LIKE $1`, "LIKE")
			assertContains(t, sql, `"u"."age" NOT BETWEEN $1 AND $2`, "NOT BETWEEN")
			assertContains(t, sql, `"u"."age" <> $1`, "NOTEQUAL")
			assertContains(t, sql, `NOT EXISTS (SELECT 1 FROM "users" AS "u")`, "NOT EXISTS")
			assertContains(t, sql, `"u"."email" NOT ILIKE $1`, "NOT ILIKE")
			assertContains(t, sql, `"u"."status" NOT IN ($1, $2)`, "NOT IN")
			assertContains(t, sql, `"u"."email" NOT LIKE $1`, "NOT LIKE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_Distinct(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Select().Distinct().
			Field(Users.ID.As("user_id")).
			From(Users.Table)
		sql, _, _ := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "DISTINCT", "DISTINCT")
		case DialectPostgreSQL:
			assertContains(t, sql, "DISTINCT", "DISTINCT")
		}
		t.Logf("dialect: %s sql: %s", supportDialect.name, sql)
	})
}
func Test_Select_Function(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Select(
			// Функции агрегатные
			Avg(Users.Age, false).As("age_avg"),
			BitAnd(Users.Age, false).As("age_bitand"),
			BitOr(Users.Age, false).As("age_bitor"),
			BitXor(Users.Age, false).As("age_bitxor"),
			Count(Users.ID, false).As("id_count"),
			GroupConcat(Users.Status, false).As("status_group_concat"),
			Max(Users.Age, false).As("age_max"),
			Min(Users.Age, false).As("age_min"),
			StdDev(Users.Age, false).As("age_stddev"),
			Sum(Users.Age, false).As("age_sum"),
			Variance(Users.Age, false).As("age_variance"),
			// Функции условий
			Case(CaseIf(CasePair(Less(Users.Age, Value(18)), Value("young"))), CaseElse(Value("adult"))).As("age_case"),
			Coalesce(Users.CreateAt, Users.UpdateAt).As("date_coalesce"),
			Greatest(Users.CreateAt, Users.UpdateAt).As("date_greatest"),
			Least(Users.CreateAt, Users.UpdateAt).As("date_least"),
			NullIf(Users.CreateAt, Users.UpdateAt).As("date_if"),
			// Функции конвертации
			Cast(Users.Age, TypeString).As("age_cast"),
			CharLength(Users.Status).As("status_length"),
			DateFormat(Users.CreateAt, Literal("%Y-%m-%d")).As("createat_dateformat"),
			Degrees(Users.Age).As("age_degrees"),
			Length(Users.Status).As("status_length"),
			Position(Users.Status, Value("test")).As("status_position"),
			Radians(Users.Age).As("age_radians"),
			// Функции даты и времени
			CurDate().As("curdate"),
			CurTime().As("curtime"),
			DateAdd(Users.CreateAt, Literal("7 DAY")).As("createat_dateadd"),
			DateDiff(Users.UpdateAt, Users.CreateAt).As("updateat_datediff"),
			DateSub(Users.CreateAt, Literal("7 DAY")).As("createat_datesub"),
			Day(Users.CreateAt).As("createat_day"),
			DayName(Users.CreateAt).As("createat_dayname"),
			Hour(Users.CreateAt).As("createat_hour"),
			Minute(Users.CreateAt).As("createat_minute"),
			Month(Users.CreateAt).As("createat_month"),
			MonthName(Users.CreateAt).As("createat_monthname"),
			Now().As("now"),
			Quarter(Users.CreateAt).As("createat_quarter"),
			Second(Users.CreateAt).As("createat_second"),
			TimeAdd(Users.CreateAt, Literal("4 HOUR")).As("createat_timeadd"),
			TimeDiff(Users.UpdateAt, Users.CreateAt).As("updateat_timediff"),
			TimeSub(Users.CreateAt, Literal("4 HOUR")).As("createat_timesub"),
			Week(Users.CreateAt).As("createat_week"),
			Year(Users.CreateAt).As("createat_year"),
			// Функции обмена данными
			JsonArray(Users.Data, Value("test"), Value("test")).As("data_jsonarray"),
			JsonArrayAgg(Users.Data).As("data_jsonarrayagg"),
			JsonContains(Users.Data, Value(`{"theme":"dark"}`)).As("data_jsoncontains"),
			JsonExtract(Users.Data, JsonGroup(JsonPath(JsonKey("col"), JsonIndex(0), JsonKey("name"))), TypeString).As("data_jsonextract"),
			JsonObject(JsonPair(JsonKey("users"), Count(Users.Data, false))).As("data_jsonobject"),
			JsonObjectAgg(Users.Data, Users.Age).As("data_jsonobjectagg"),
			JsonRemove(Users.Data, JsonGroup(JsonPath(JsonKey("temp"))), JsonGroup(JsonPath(JsonKey("session")))).As("data_jsonremove"),
			JsonSet(Users.Data, JsonGroup(JsonPath(JsonKey("temp")), Value(0)), JsonGroup(JsonPath(JsonKey("session")), Value("active"))).As("data_jsonset"),
			JsonType(Users.Data).As("jsontype"),
			// Функции математические
			Abs(Users.Age).As("age_abs"),
			ACos(Users.Age).As("age_acos"),
			ASin(Users.Age).As("age_asin"),
			ATan(Users.Age).As("age_atan"),
			ATan2(Users.Age, Users.Age).As("age_atan2"),
			Cbrt(Users.Age).As("age_cbrt"),
			Ceil(Users.Age).As("age_ceil"),
			Cos(Users.Age).As("age_cos"),
			Exp(Users.Age).As("age_exp"),
			Floor(Users.Age).As("age_floor"),
			Ln(Users.Age).As("age_ln"),
			Log(Users.Age, Value(3)).As("age_log"),
			Mod(Users.Age, Value(3)).As("age_mod"),
			Pi().As("pi"),
			Power(Users.Age, Value(3)).As("age_power"),
			Rand().As("rand"),
			Round(Users.Age, Value(2)).As("age_round"),
			Sin(Users.Age).As("age_sin"),
			Sqrt(Users.Age).As("age_sqrt"),
			Tan(Users.Age).As("age_tan"),
			Trunc(Users.Age, Value(2)).As("age_trunc"),
			// Функции строковые
			Concat(Users.Status, Users.Name, Users.Email).As("status_concat"),
			ConcatWs(Value("_"), Users.Status, Users.Name, Users.Email).As("status_concat_ws"),
			LeftString(Users.Status, Value(2)).As("status_lstr"),
			Lower(Users.Status).As("status_lower"),
			LPad(Users.Status, Value(10), Value(",")).As("status_lpad"),
			LTrim(Users.Status).As("status_ltrim"),
			Repeat(Users.Status, Value(3)).As("status_repeat"),
			Replace(Users.Status, Value("old"), Value("new")).As("status_replace"),
			Reverse(Users.Status).As("status_reverse"),
			RightString(Users.Status, Value(2)).As("status_rstr"),
			RPad(Users.Status, Value(10), Value(",")).As("status_rpad"),
			RTrim(Users.Status).As("status_rtrim"),
			SubString(Users.Status, Value(0), Value(2)).As("status_substring"),
			Trim(Users.Status).As("status_trim"),
			Upper(Users.Status).As("status_upper"),
			// Функции ранжирующие
			RowNumber().Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("rank_row_number"),
			Rank().Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("rank_rank"),
			DenseRank().Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("rank_dense_rank"),
			PercentRank().Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("rank_percent_rank"),
			CumeDist().Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("rank_cume_dist"),
			NTile(4).Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
			).As("rank_ntile"),
			// Функции аналитические
			FirstValue(Users.Name).Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
				RowsBetween("UNBOUNDED PRECEDING", "CURRENT ROW"),
			).As("analyt_first_value"),
			LastValue(Users.Name).Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
				RowsBetween("UNBOUNDED PRECEDING", "CURRENT ROW")).As("analyt_last_value"),
			Lag(Users.Salary, 1).Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Asc(Users.HireDate)),
			).As("analyt_lag"),
			Lead(Users.Salary, 1).Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Asc(Users.HireDate)),
			).As("analyt_lead"),
			NthValue(Users.Name, 2).Over(
				PartitionBy(Users.DepartmentID),
				OrderBy(Desc(Users.Salary)),
				RowsBetween("UNBOUNDED PRECEDING", "CURRENT ROW")).As("analyt_nth_value"),
		).
			From(Users.Table)
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			// Функции агрегатные
			assertContains(t, sql, "AVG(`u`.`age`)", "AVG")
			assertContains(t, sql, "BIT_AND(`u`.`age`)", "BIT_AND")
			assertContains(t, sql, "BIT_OR(`u`.`age`)", "BIT_OR")
			assertContains(t, sql, "BIT_XOR(`u`.`age`)", "BIT_XOR")
			assertContains(t, sql, "COUNT(`u`.`id`)", "COUNT")
			assertContains(t, sql, "GROUP_CONCAT(`u`.`status` SEPARATOR ',')", "GROUPCONCAT")
			assertContains(t, sql, "MAX(`u`.`age`)", "MAX")
			assertContains(t, sql, "MIN(`u`.`age`)", "MIN")
			assertContains(t, sql, "STDDEV(`u`.`age`)", "STDDEV")
			assertContains(t, sql, "SUM(`u`.`age`)", "SUM")
			assertContains(t, sql, "VARIANCE(`u`.`age`)", "VARIANCE")
			// Функции условий
			assertContains(t, sql, "CASE WHEN `u`.`age` < ? THEN ? ELSE ? END", "CASE")
			assertContains(t, sql, "COALESCE(`u`.`createat`, `u`.`updateat`)", "COALESCE")
			assertContains(t, sql, "GREATEST(`u`.`createat`, `u`.`updateat`)", "GREATEST")
			assertContains(t, sql, "LEAST(`u`.`createat`, `u`.`updateat`)", "LEAST")
			assertContains(t, sql, "NULLIF(`u`.`createat`, `u`.`updateat`)", "NULLIF")
			// Функции конвертации
			assertContains(t, sql, "CAST(`u`.`age` AS CHAR)", "CAST")
			assertContains(t, sql, "CHAR_LENGTH(`u`.`status`)", "CHARLENGTH")
			assertContains(t, sql, "DATE_FORMAT(`u`.`createat`, '%Y-%m-%d')", "DATEFORMAT")
			assertContains(t, sql, "DEGREES(`u`.`age`)", "DEGREES")
			assertContains(t, sql, "LENGTH(`u`.`status`)", "LENGTH")
			assertContains(t, sql, "POSITION(? IN `u`.`status`)", "POSITION")
			assertContains(t, sql, "RADIANS(`u`.`age`)", "RADIANS")
			// Функции даты и времени
			assertContains(t, sql, "CURDATE()", "CURDATE")
			assertContains(t, sql, "CURTIME()", "CURTIME")
			assertContains(t, sql, "DATE_ADD(`u`.`createat`, INTERVAL '7 DAY')", "DATEADD")
			assertContains(t, sql, "DATEDIFF(`u`.`updateat`, `u`.`createat`)", "DATEDIFF")
			assertContains(t, sql, "DATE_SUB(`u`.`createat`, INTERVAL '7 DAY')", "DATESUB")
			assertContains(t, sql, "DAY(`u`.`createat`)", "DAY")
			assertContains(t, sql, "DAYNAME(`u`.`createat`)", "DAYNAME")
			assertContains(t, sql, "HOUR(`u`.`createat`)", "HOUR")
			assertContains(t, sql, "MINUTE(`u`.`createat`)", "MINUTE")
			assertContains(t, sql, "MONTH(`u`.`createat`)", "MONTH")
			assertContains(t, sql, "MONTHNAME(`u`.`createat`)", "MONTHNAME")
			assertContains(t, sql, "NOW()", "NOW")
			assertContains(t, sql, "QUARTER(`u`.`createat`)", "QUARTER")
			assertContains(t, sql, "SECOND(`u`.`createat`)", "SECOND")
			assertContains(t, sql, "TIME_ADD(`u`.`createat`, INTERVAL '4 HOUR')", "TIMEADD")
			assertContains(t, sql, "TIMEDIFF(`u`.`updateat`, `u`.`createat`)", "TIMEDIFF")
			assertContains(t, sql, "TIME_SUB(`u`.`createat`, INTERVAL '4 HOUR')", "TIMESUB")
			assertContains(t, sql, "WEEK(`u`.`createat`)", "WEEK")
			assertContains(t, sql, "YEAR(`u`.`createat`)", "YEAR")
			// Функции обмена данными
			assertContains(t, sql, "JSON_ARRAY(`u`.`data`, ?, ?)", "JSON_ARRAY")
			assertContains(t, sql, "JSON_ARRAYAGG(`u`.`data`)", "JSON_ARRAYAGG")
			assertContains(t, sql, "JSON_CONTAINS(`u`.`data`, ?)", "JSON_CONTAINS")
			assertContains(t, sql, "(`u`.`data` ->> '$.col[0].name')", "JSON_EXTRACT")
			assertContains(t, sql, "JSON_OBJECT('users', COUNT(`u`.`data`))", "JSON_OBJECT")
			assertContains(t, sql, "JSON_OBJECTAGG(`u`.`data`, `u`.`age`)", "JSON_OBJECTAGG")
			assertContains(t, sql, "JSON_REMOVE(`u`.`data`, '$.temp', '$.session')", "JSON_REMOVE")
			assertContains(t, sql, "JSON_SET(`u`.`data`, '$.temp', ?, '$.session', ?)", "JSON_SET")
			assertContains(t, sql, "JSON_TYPE(`u`.`data`)", "JSON_TYPE")
			// Функции математические
			assertContains(t, sql, "ABS(`u`.`age`)", "ABS")
			assertContains(t, sql, "ACOS(`u`.`age`)", "ACOS")
			assertContains(t, sql, "ASIN(`u`.`age`)", "ASIN")
			assertContains(t, sql, "ATAN(`u`.`age`)", "ATAN")
			assertContains(t, sql, "ATAN2(`u`.`age`, `u`.`age`)", "ATAN2")
			assertContains(t, sql, "CBRT(`u`.`age`)", "CBRT")
			assertContains(t, sql, "CEILING(`u`.`age`)", "CEIL")
			assertContains(t, sql, "COS(`u`.`age`)", "COS")
			assertContains(t, sql, "EXP(`u`.`age`)", "EXP")
			assertContains(t, sql, "FLOOR(`u`.`age`)", "FLOOR")
			assertContains(t, sql, "LN(`u`.`age`)", "LN")
			assertContains(t, sql, "LOG(`u`.`age`, ?)", "LOG")
			assertContains(t, sql, "MOD(`u`.`age`, ?)", "MOD")
			assertContains(t, sql, "PI()", "PI")
			assertContains(t, sql, "POWER(`u`.`age`, ?)", "POWER")
			assertContains(t, sql, "RAND()", "RAND")
			assertContains(t, sql, "ROUND(`u`.`age`, ?)", "ROUND")
			assertContains(t, sql, "SIN(`u`.`age`)", "SIN")
			assertContains(t, sql, "SQRT(`u`.`age`)", "SQRT")
			assertContains(t, sql, "TAN(`u`.`age`)", "TAN")
			assertContains(t, sql, "TRUNCATE(`u`.`age`, ?)", "TRUNC")
			// Функции строковые
			assertContains(t, sql, "CONCAT(`u`.`status`, `u`.`name`, `u`.`email`)", "CONCAT")
			assertContains(t, sql, "CONCAT_WS(?, `u`.`status`, `u`.`name`, `u`.`email`)", "CONCATWS")
			assertContains(t, sql, "LEFT(`u`.`status`, ?)", "LEFTSTRING")
			assertContains(t, sql, "LOWER(`u`.`status`)", "LOWER")
			assertContains(t, sql, "LPAD(`u`.`status`, ?, ?)", "LPAD")
			assertContains(t, sql, "LTRIM(`u`.`status`)", "LTRIM")
			assertContains(t, sql, "REPEAT(`u`.`status`, ?)", "REPEAT")
			assertContains(t, sql, "REPLACE(`u`.`status`, ?, ?)", "REPLACE")
			assertContains(t, sql, "REVERSE(`u`.`status`)", "REVERSE")
			assertContains(t, sql, "RIGHT(`u`.`status`, ?)", "RIGHTSTRING")
			assertContains(t, sql, "RPAD(`u`.`status`, ?, ?)", "RPAD")
			assertContains(t, sql, "RTRIM(`u`.`status`)", "RTRIM")
			assertContains(t, sql, "SUBSTRING(`u`.`status`, ?, ?)", "SUBSTRING")
			assertContains(t, sql, "TRIM(`u`.`status`)", "TRIM")
			assertContains(t, sql, "UPPER(`u`.`status`)", "UPPER")
			// Функции ранжирующие
			assertContains(t, sql, "ROW_NUMBER() OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "ROW_NUMBER")
			assertContains(t, sql, "RANK() OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "RANK")
			assertContains(t, sql, "DENSE_RANK() OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "DENSE_RANK")
			assertContains(t, sql, "PERCENT_RANK() OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "PERCENT_RANK")
			assertContains(t, sql, "CUME_DIST() OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "CUME_DIST")
			assertContains(t, sql, "NTILE(4) OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC)", "NTILE")
			// Функции аналитические
			assertContains(t, sql, "FIRST_VALUE(`u`.`name`) OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)", "FIRST_VALUE")
			assertContains(t, sql, "LAST_VALUE(`u`.`name`) OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)", "LAST_VALUE")
			assertContains(t, sql, "LAG(`u`.`salary`, 1) OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`hire_date` ASC)", "LAG")
			assertContains(t, sql, "LEAD(`u`.`salary`, 1) OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`hire_date` ASC)", "LEAD")
			assertContains(t, sql, "NTH_VALUE(`u`.`name`, 2) OVER (PARTITION BY `u`.`department_id` ORDER BY `u`.`salary` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)", "NTH_VALUE")
		case DialectPostgreSQL:
			// Функции агрегатные
			assertContains(t, sql, `AVG("u"."age")`, "AVG")
			assertContains(t, sql, `BIT_AND("u"."age")`, "BIT_AND")
			assertContains(t, sql, `BIT_OR("u"."age")`, "BIT_OR")
			assertContains(t, sql, `BIT_XOR("u"."age")`, "BIT_XOR")
			assertContains(t, sql, `COUNT("u"."id")`, "COUNT")
			assertContains(t, sql, `STRING_AGG("u"."status", ',')`, "GROUPCONCAT")
			assertContains(t, sql, `MAX("u"."age")`, "MAX")
			assertContains(t, sql, `MIN("u"."age")`, "MIN")
			assertContains(t, sql, `STDDEV_SAMP("u"."age")`, "STDDEV")
			assertContains(t, sql, `SUM("u"."age")`, "SUM")
			assertContains(t, sql, `VAR_SAMP("u"."age")`, "VARIANCE")
			// Функции условий
			assertContains(t, sql, `CASE WHEN "u"."age" < $1 THEN $2 ELSE $3 END`, "CASE")
			assertContains(t, sql, `COALESCE("u"."createat", "u"."updateat")`, "COALESCE")
			assertContains(t, sql, `GREATEST("u"."createat", "u"."updateat")`, "GREATEST")
			assertContains(t, sql, `LEAST("u"."createat", "u"."updateat")`, "LEAST")
			assertContains(t, sql, `NULLIF("u"."createat", "u"."updateat")`, "NULLIF")
			// Функции конвертации
			assertContains(t, sql, `CAST("u"."age" AS VARCHAR) `, "CAST")
			assertContains(t, sql, `CHAR_LENGTH("u"."status")`, "CHARLENGTH")
			assertContains(t, sql, `TO_CHAR("u"."createat", '%Y-%m-%d')`, "DATE_FORMAT")
			assertContains(t, sql, `DEGREES("u"."age")`, "DEGREES")
			assertContains(t, sql, `LENGTH("u"."status")`, "LENGTH")
			assertContains(t, sql, `POSITION($1 IN "u"."status")`, "POSITION")
			assertContains(t, sql, `RADIANS("u"."age")`, "RADIANS")
			// Функции даты и времени
			assertContains(t, sql, `CURRENT_DATE`, "CURDATE")
			assertContains(t, sql, `CURRENT_TIME`, "CURTIME")
			assertContains(t, sql, `("u"."createat" + INTERVAL '7 DAY')`, "DATEADD")
			assertContains(t, sql, `DATE_PART('day', "u"."updateat" - "u"."createat")`, "DATEDIFF")
			assertContains(t, sql, `("u"."createat" - INTERVAL '7 DAY')`, "DATESUB")
			assertContains(t, sql, `EXTRACT(DAY FROM "u"."createat")`, "DAY")
			assertContains(t, sql, `TO_CHAR("u"."createat", 'Day')`, "DAYNAME")
			assertContains(t, sql, `EXTRACT(HOUR FROM "u"."createat")`, "HOUR")
			assertContains(t, sql, `EXTRACT(MINUTE FROM "u"."createat")`, "MINUTE")
			assertContains(t, sql, `EXTRACT(MONTH FROM "u"."createat")`, "MONTH")
			assertContains(t, sql, `TO_CHAR("u"."createat", 'Month')`, "MONTHNAME")
			assertContains(t, sql, `CURRENT_TIMESTAMP`, "NOW")
			assertContains(t, sql, `EXTRACT(QUARTER FROM "u"."createat")`, "QUARTER")
			assertContains(t, sql, `EXTRACT(SECOND FROM "u"."createat")`, "SECOND")
			assertContains(t, sql, `("u"."createat" + INTERVAL '4 HOUR')`, "TIMEADD")
			assertContains(t, sql, `DATE_PART('time', "u"."updateat" - "u"."createat")`, "TIMEDIFF")
			assertContains(t, sql, `("u"."createat" - INTERVAL '4 HOUR')`, "TIMESUB")
			assertContains(t, sql, `EXTRACT(WEEK FROM "u"."createat")`, "WEEK")
			assertContains(t, sql, `EXTRACT(YEAR FROM "u"."createat")`, "YEAR")
			// Функции обмена данными
			assertContains(t, sql, `JSON_ARRAY("u"."data", $1, $2)`, "JSON_ARRAY")
			assertContains(t, sql, `JSON_AGG("u"."data")`, "JSON_ARRAYAGG")
			assertContains(t, sql, `("u"."data" @> $1)`, "JSON_CONTAINS")
			assertContains(t, sql, `("u"."data" #>> '{col,0,name}')`, "JSON_EXTRACT")
			assertContains(t, sql, `JSON_BUILD_OBJECT('users', COUNT("u"."data"))`, "JSON_OBJECT")
			assertContains(t, sql, `JSON_OBJECT_AGG("u"."data", "u"."age")`, "JSON_OBJECTAGG")
			assertContains(t, sql, `("u"."data" - '{temp}' - '{session}')`, "JSON_REMOVE")
			//assertContains(t, sql, `jsonb_set`, "JSON_SET")
			assertContains(t, sql, `jsonb_typeof("u"."data")`, "JSON_TYPE")
			// Функции математические
			assertContains(t, sql, `ABS("u"."age")`, "ABS")
			assertContains(t, sql, `ACOS("u"."age")`, "ACOS")
			assertContains(t, sql, `ASIN("u"."age")`, "ASIN")
			assertContains(t, sql, `ATAN("u"."age")`, "ATAN")
			assertContains(t, sql, `ATAN2("u"."age", "u"."age")`, "ATAN2")
			assertContains(t, sql, `CBRT("u"."age")`, "CBRT")
			assertContains(t, sql, `CEIL("u"."age")`, "CEIL")
			assertContains(t, sql, `COS("u"."age")`, "COS")
			assertContains(t, sql, `EXP("u"."age")`, "EXP")
			assertContains(t, sql, `FLOOR("u"."age")`, "FLOOR")
			assertContains(t, sql, `LN("u"."age")`, "LN")
			assertContains(t, sql, `LOG("u"."age", $1)`, "LOG")
			assertContains(t, sql, `MOD("u"."age", $1)`, "MOD")
			assertContains(t, sql, `PI()`, "PI")
			assertContains(t, sql, `POWER("u"."age", $1)`, "POWER")
			assertContains(t, sql, `RANDOM`, "RAND")
			assertContains(t, sql, `ROUND("u"."age", $1)`, "ROUND")
			assertContains(t, sql, `SIN("u"."age")`, "SIN")
			assertContains(t, sql, `SQRT("u"."age")`, "SQRT")
			assertContains(t, sql, `TAN("u"."age")`, "TAN")
			assertContains(t, sql, `TRUNC("u"."age", $1)`, "TRUNC")
			// Функции строковые
			assertContains(t, sql, `CONCAT("u"."status", "u"."name", "u"."email")`, "CONCAT")
			assertContains(t, sql, `CONCAT_WS($1, "u"."status", "u"."name", "u"."email")`, "CONCATWS")
			assertContains(t, sql, `LEFT("u"."status", $1)`, "LEFTSTRING")
			assertContains(t, sql, `LOWER("u"."status")`, "LOWER")
			assertContains(t, sql, `LPAD("u"."status", $1, $2)`, "LPAD")
			assertContains(t, sql, `LTRIM("u"."status")`, "LTRIM")
			assertContains(t, sql, `REPEAT("u"."status", $1)`, "REPEAT")
			assertContains(t, sql, `REPLACE("u"."status", $1, $2)`, "REPLACE")
			assertContains(t, sql, `REVERSE("u"."status")`, "REVERSE")
			assertContains(t, sql, `RIGHT("u"."status", $1)`, "RIGHTSTRING")
			assertContains(t, sql, `RPAD("u"."status", $1, $2)`, "RPAD")
			assertContains(t, sql, `RTRIM("u"."status")`, "RTRIM")
			assertContains(t, sql, `SUBSTRING("u"."status", $1, $2)`, "SUBSTRING")
			assertContains(t, sql, `TRIM("u"."status")`, "TRIM")
			assertContains(t, sql, `UPPER("u"."status")`, "UPPER")
			// Функции ранжирующие
			assertContains(t, sql, `ROW_NUMBER() OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "ROW_NUMBER")
			assertContains(t, sql, `RANK() OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "RANK")
			assertContains(t, sql, `DENSE_RANK() OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "DENSE_RANK")
			assertContains(t, sql, `PERCENT_RANK() OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "PERCENT_RANK")
			assertContains(t, sql, `CUME_DIST() OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "CUME_DIST")
			assertContains(t, sql, `NTILE(4) OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC)`, "NTILE")
			// Функции аналитические
			assertContains(t, sql, `FIRST_VALUE("u"."name") OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)`, "FIRST_VALUE")
			assertContains(t, sql, `LAST_VALUE("u"."name") OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)`, "LAST_VALUE")
			assertContains(t, sql, `LAG("u"."salary", 1) OVER (PARTITION BY "u"."department_id" ORDER BY "u"."hire_date" ASC)`, "LAG")
			assertContains(t, sql, `LEAD("u"."salary", 1) OVER (PARTITION BY "u"."department_id" ORDER BY "u"."hire_date" ASC)`, "LEAD")
			assertContains(t, sql, `NTH_VALUE("u"."name", 2) OVER (PARTITION BY "u"."department_id" ORDER BY "u"."salary" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)`, "NTH_VALUE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_GroupBy(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Select(
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
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "`u`.`department_id`, `u`.`status`", "LIST")
		case DialectPostgreSQL:
			assertContains(t, sql, `"u"."department_id", "u"."status"`, "LIST")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_Having(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Select(
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
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "COUNT(`u`.`id`) > ?", "COUNT")
			assertContains(t, sql, "SUM(`u`.`age`) > ?", "SUM")
			assertContains(t, sql, "AVG(`u`.`age`) BETWEEN ? AND ?", "BETWEEN")
			assertContains(t, sql, "`u`.`department_id` IN (?, ?, ?)", "IN")
		case DialectPostgreSQL:
			assertContains(t, sql, `COUNT("u"."id") > $1`, "COUNT")
			assertContains(t, sql, `SUM("u"."age") > $1`, "SUM")
			assertContains(t, sql, `AVG("u"."age") BETWEEN $1 AND $2`, "BETWEEN")
			assertContains(t, sql, `"u"."department_id" IN ($1, $2, $3)`, "IN")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_Join(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Select(Users.ID.As("user_id")).
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
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "CROSS JOIN `orders` AS `o`", "CROSS")
			assertContains(t, sql, "FULL JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "FULL")
			assertContains(t, sql, "FULL OUTER JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "FULL OUTER")
			assertContains(t, sql, "INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "INNER")
			assertContains(t, sql, "LEFT JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "LEFT")
			assertContains(t, sql, "LEFT OUTER JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "LEFT OUTER")
			assertContains(t, sql, "RIGHT JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "RIGHT")
			assertContains(t, sql, "RIGHT OUTER JOIN `orders` AS `o` ON `u`.`id` = `o`.`user_id`", "RIGHT OUTER")
		case DialectPostgreSQL:
			assertContains(t, sql, `CROSS JOIN "orders" AS "o"`, "CROSS")
			assertContains(t, sql, `FULL JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "FULL")
			assertContains(t, sql, `FULL OUTER JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "FULL OUTER")
			assertContains(t, sql, `INNER JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "INNER")
			assertContains(t, sql, `LEFT JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "LEFT")
			assertContains(t, sql, `LEFT OUTER JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "LEFT OUTER")
			assertContains(t, sql, `RIGHT JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "RIGHT")
			assertContains(t, sql, `RIGHT OUTER JOIN "orders" AS "o" ON "u"."id" = "o"."user_id"`, "RIGHT OUTER")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_Limit(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Select(
			Users.ID,
			Users.Name,
			Users.Age,
		).
			From(Users.Table).
			Where(Equal(Users.Status, Value("active"))).
			Limit(10)
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "LIMIT ?", "LIMIT")
		case DialectPostgreSQL:
			assertContains(t, sql, "LIMIT $1", "LIMIT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_Logical(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Select(Users.ID.As("user_id")).
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
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "`u`.`id` = ? AND `u`.`id` = ?", "AND")
			assertContains(t, sql, "`u`.`id` = ? OR `u`.`id` = ?", "OR")
		case DialectPostgreSQL:
			assertContains(t, sql, `"u"."id" = $1 AND "u"."id" = $2`, "AND")
			assertContains(t, sql, `"u"."id" = $3 OR "u"."id" = $4`, "OR")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_Offset(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Select(
			Users.ID,
			Users.Name,
			Users.Age,
		).
			From(Users.Table).
			Where(Equal(Users.Status, Value("active"))).
			Offset(20)
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "OFFSET ?", "OFFSET")
		case DialectPostgreSQL:
			assertContains(t, sql, "OFFSET $1", "OFFSET")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_OrderBy(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryMain := builder.Select(Users.ID.As("user_id"), Users.Name.As("user_name")).
			From(Users.Table).
			OrderBy(
				Asc(Users.Name),
				Desc(Users.ID),
			)
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "`u`.`name` ASC", "ASC")
			assertContains(t, sql, "`u`.`id` DESC", "DESC")
		case DialectPostgreSQL:
			assertContains(t, sql, `"u"."name" ASC`, "ASC")
			assertContains(t, sql, `"u"."id" DESC`, "DESC")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_Subquery(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		querySub := builder.Select(Count(Orders.ID, false)).
			From(Orders.Table).
			Where(
				Equal(Orders.UserID, Users.ID),
			)
		queryInSub := builder.Select(Orders.UserID).
			From(Orders.Table).
			Where(
				Greater(Orders.Amount, Value(1000)),
			)
		queryExistsSub := builder.Select(Orders.ID).
			From(Orders.Table).
			Where(
				Equal(Orders.UserID, Users.ID),
			)
		queryMain := builder.Select(
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

		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "AS", "AS (subquery1)")
			assertContains(t, sql, "IN", "IN (subquery2)")
			assertContains(t, sql, "EXISTS", "EXISTS (subquery3)")
		case DialectPostgreSQL:
			assertContains(t, sql, "AS", "AS (subquery1)")
			assertContains(t, sql, "IN", "IN (subquery2)")
			assertContains(t, sql, "EXISTS", "EXISTS (subquery3)")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_Unions(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryUnion := builder.Select(
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
		queryUnionAll := builder.Select(
			Products.Name.As("person_name"),
			Products.Name.As("contact_email"),
		).
			From(Products.Table).
			Where(
				Greater(Products.Count, Value(10)),
			)
		queryUnionExcept := builder.Select(
			Categories.Name.As("person_name"),
			Categories.Type.As("contact_email"),
		).
			From(Categories.Table).
			Where(
				In(Categories.Type, Array("premium", "standard", "basic")),
			)
		queryUnionIntersect := builder.Select(
			Departments.Name.As("person_name"),
			Departments.Name.As("contact_email"),
		).
			From(Departments.Table).
			Where(
				IsNull(Departments.ParentID),
			)
		queryMain := builder.Select(
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
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "UNION", "UNION")
			assertContains(t, sql, "UNION ALL", "UNION ALL")
			assertContains(t, sql, "EXCEPT", "UNION EXCEPT")
			assertContains(t, sql, "INTERSECT", "UNION INTERSECT")
		case DialectPostgreSQL:
			assertContains(t, sql, "UNION", "UNION")
			assertContains(t, sql, "UNION ALL", "UNION ALL")
			assertContains(t, sql, "EXCEPT", "UNION EXCEPT")
			assertContains(t, sql, "INTERSECT", "UNION INTERSECT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_Where(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		existsSub := builder.Select(Levels.ID).
			From(Levels.Table).
			Where(
				Equal(Levels.UserID, Users.ID),
			)
		queryMain := builder.Select(Users.ID).
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
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "`u`.`status` = ?", "EQUAL")
			assertContains(t, sql, "`u`.`age` > ?", "GREATER")
			assertContains(t, sql, "`u`.`age` < ?", "LESS")
			assertContains(t, sql, "`u`.`department_id` IN (?, ?, ?)", "IN")
			assertContains(t, sql, "`u`.`email` LIKE ?", "LIKE")
			assertContains(t, sql, "EXISTS", "EXISTS")
			assertContains(t, sql, "`u`.`status` <> ?", "NOTEQUAL")
			assertContains(t, sql, "`u`.`email` IS NOT NULL", "IS NOT NULL")
		case DialectPostgreSQL:
			assertContains(t, sql, `"u"."status" = $1`, "EQUAL")
			assertContains(t, sql, `"u"."age" > $1`, "GREATER")
			assertContains(t, sql, `"u"."age" < $1`, "LESS")
			assertContains(t, sql, `"u"."department_id" IN ($1, $2, $3)`, "IN")
			assertContains(t, sql, `"u"."email" LIKE $1`, "LIKE")
			assertContains(t, sql, `EXISTS`, "EXISTS")
			assertContains(t, sql, `"u"."status" <> $1`, "NOTEQUAL")
			assertContains(t, sql, `"u"."email" IS NOT NULL`, "IS NOT NULL")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Select_With(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryWithN := WithN("WithN", builder.Select(
			Orders.UserID,
			Count(Orders.ID, false).As("order_count"),
			Sum(Orders.Amount, false).As("total_spent"),
		).
			From(Orders.Table).
			Where(Greater(Orders.Amount, Value(1000))).
			GroupBy(Orders.UserID),
			"user_id", "order_count", "total_spent",
		)
		queryWithR := WithR("WithR", builder.Select(
			Departments.ID,
			Departments.Name,
			Departments.ParentID,
		).
			From(Departments.Table).
			Where(IsNull(Departments.ParentID)).
			Unions(
				UnionAll(builder.Select(
					Departments.ID,
					Departments.Name,
					Departments.ParentID,
				).
					From(Departments.Table).
					Join(Inner(CTE("WithR", "dh"), Equal(Departments.ParentID, Departments.ID)))),
			),
			"dept_id", "dept_name", "parent_id",
		)
		queryMain := builder.Select(
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
				Left(Query(builder.Select(
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
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			assertContains(t, sql, "WithN", "WithN")
			assertContains(t, sql, "WithR", "WithR")
			assertContains(t, sql, "RECURSIVE", "RECURSIVE")
			assertContains(t, sql, "UNION ALL", "UNION ALL for RCTE")
			assertContains(t, sql, "WITH", "WITH")
		case DialectPostgreSQL:
			assertContains(t, sql, "WithN", "WithN")
			assertContains(t, sql, "WithR", "WithR")
			assertContains(t, sql, "RECURSIVE", "RECURSIVE")
			assertContains(t, sql, "UNION ALL", "UNION ALL for RCTE")
			assertContains(t, sql, "WITH", "WITH")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
	})
}
func Test_Update(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		builder, _ := NewBuilder(supportDialect)
		defer builder.Close()
		queryWith := WithN("user_stats", builder.Select(
			Orders.UserID,
			Count(Orders.ID, false).As("order_count"),
		).
			From(Orders.Table).
			Where(Greater(Orders.Amount, Value(100))).
			GroupBy(Orders.UserID),
			"user_id", "order_count")
		queryMain := builder.Update(Users.Table).
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
		sql, sqa, err := builder.Build(queryMain)
		switch supportDialect {
		case DialectMySQL:
			//assertContains(t, sql, "UPDATE", "UPDATE")
			//assertContains(t, sql, "FROM", "FROM")
			//assertContains(t, sql, "SET", "SET")
			//assertContains(t, sql, "WHERE", "WHERE")
			//assertContains(t, sql, "WITH", "WITH")
		case DialectPostgreSQL:
			//assertContains(t, sql, "UPDATE", "UPDATE")
			//assertContains(t, sql, "FROM", "FROM")
			//assertContains(t, sql, "RETURNING", "RETURNING")
			//assertContains(t, sql, "SET", "SET")
			//assertContains(t, sql, "USING", "USING")
			//assertContains(t, sql, "WHERE", "WHERE")
			//assertContains(t, sql, "WITH", "WITH")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqa, supportDialect.name, sql)
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
