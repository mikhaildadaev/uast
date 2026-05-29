package uast

import (
	"regexp"
	"strings"
	"testing"
	"time"
)

// Публичные переменные
var Data = struct {
	Column struct {
		Date   *ColumnExpr[time.Time]
		ID     *ColumnExpr[int64]
		Json   *ColumnExpr[string]
		Number *ColumnExpr[int]
		String *ColumnExpr[string]
		Time   *ColumnExpr[time.Time]
	}
	Table *TableSource
}{
	Table: NewTable("data", "d"),
}
var Test = struct {
	Column struct {
		CreateAt *ColumnExpr[time.Time]
		Date     *ColumnExpr[time.Time]
		ID       *ColumnExpr[int64]
		Json     *ColumnExpr[string]
		Name     *ColumnExpr[string]
		Number   *ColumnExpr[int]
		String   *ColumnExpr[string]
		UpdateAt *ColumnExpr[time.Time]
		X        *ColumnExpr[int]
		Y        *ColumnExpr[int]
	}
	Table *TableSource
}{
	Table: NewTable("test", "t"),
}

// Публичные функции
func Test_Core_clauseGroupBy(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.Column.String,
			).
			GroupBy(
				Test.Column.String,
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "GROUP BY `t`.`string`", "GROUP BY")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "GROUP BY [t].[string]", "GROUP BY")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "GROUP BY `t`.`string`", "GROUP BY")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `GROUP BY "t"."string"`, "GROUP BY")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `GROUP BY "t"."string"`, "GROUP BY")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_clauseHaving(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.Column.String,
			).
			Having(
				Greater(Count(Test.Column.ID, false), Value[int64](2)),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "HAVING COUNT(`t`.`id`) > ?", "HAVING")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "HAVING COUNT([t].[id]) > @p1", "HAVING")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "HAVING COUNT(`t`.`id`) > ?", "HAVING")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `HAVING COUNT("t"."id") > $1`, "HAVING")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `HAVING COUNT("t"."id") > ?`, "HAVING")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_clauseJoin(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Data.Table).
			Field(
				Data.Column.ID,
			).
			Join(
				Cross(Test.Table),
				Full(Test.Table, Equal(Test.Column.ID, Data.Column.ID)),
				FullOuter(Test.Table, Equal(Test.Column.ID, Data.Column.ID)),
				Inner(Test.Table, Equal(Test.Column.ID, Data.Column.ID)),
				Left(Test.Table, Equal(Test.Column.ID, Data.Column.ID)),
				LeftOuter(Test.Table, Equal(Test.Column.ID, Data.Column.ID)),
				Right(Test.Table, Equal(Test.Column.ID, Data.Column.ID)),
				RightOuter(Test.Table, Equal(Test.Column.ID, Data.Column.ID)),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "CROSS JOIN `test` AS `t`", "CROSS JOIN")
			assertContains(t, sqlSelectQuery, "FULL JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "FULL JOIN")
			assertContains(t, sqlSelectQuery, "FULL OUTER JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "FULL OUTER JOIN")
			assertContains(t, sqlSelectQuery, "INNER JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "INNER JOIN")
			assertContains(t, sqlSelectQuery, "LEFT JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "LEFT JOIN")
			assertContains(t, sqlSelectQuery, "LEFT OUTER JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "LEFT OUTER JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "RIGHT JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT OUTER JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "RIGHT OUTER JOIN")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "CROSS JOIN [test] AS [t]", "CROSS JOIN")
			assertContains(t, sqlSelectQuery, "FULL JOIN [test] AS [t] ON [t].[id] = [d].[id]", "FULL JOIN")
			assertContains(t, sqlSelectQuery, "FULL OUTER JOIN [test] AS [t] ON [t].[id] = [d].[id]", "FULL OUTER JOIN")
			assertContains(t, sqlSelectQuery, "INNER JOIN [test] AS [t] ON [t].[id] = [d].[id]", "INNER JOIN")
			assertContains(t, sqlSelectQuery, "LEFT JOIN [test] AS [t] ON [t].[id] = [d].[id]", "LEFT JOIN")
			assertContains(t, sqlSelectQuery, "LEFT OUTER JOIN [test] AS [t] ON [t].[id] = [d].[id]", "LEFT OUTER JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT JOIN [test] AS [t] ON [t].[id] = [d].[id]", "RIGHT JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT OUTER JOIN [test] AS [t] ON [t].[id] = [d].[id]", "RIGHT OUTER JOIN")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "CROSS JOIN `test` AS `t`", "CROSS JOIN")
			assertContains(t, sqlSelectQuery, "FULL JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "FULL JOIN")
			assertContains(t, sqlSelectQuery, "FULL OUTER JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "FULL OUTER JOIN")
			assertContains(t, sqlSelectQuery, "INNER JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "INNER JOIN")
			assertContains(t, sqlSelectQuery, "LEFT JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "LEFT JOIN")
			assertContains(t, sqlSelectQuery, "LEFT OUTER JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "LEFT OUTER JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "RIGHT JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT OUTER JOIN `test` AS `t` ON `t`.`id` = `d`.`id`", "RIGHT OUTER JOIN")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `CROSS JOIN "test" AS "t"`, "CROSS JOIN")
			assertContains(t, sqlSelectQuery, `FULL JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "FULL JOIN")
			assertContains(t, sqlSelectQuery, `FULL OUTER JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "FULL OUTER JOIN")
			assertContains(t, sqlSelectQuery, `INNER JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "INNER JOIN")
			assertContains(t, sqlSelectQuery, `LEFT JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "LEFT JOIN")
			assertContains(t, sqlSelectQuery, `LEFT OUTER JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "LEFT OUTER JOIN")
			assertContains(t, sqlSelectQuery, `RIGHT JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "RIGHT JOIN")
			assertContains(t, sqlSelectQuery, `RIGHT OUTER JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "RIGHT OUTER JOIN")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `CROSS JOIN "test" AS "t"`, "CROSS JOIN")
			assertContains(t, sqlSelectQuery, `FULL JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "FULL JOIN")
			assertContains(t, sqlSelectQuery, `FULL OUTER JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "FULL OUTER JOIN")
			assertContains(t, sqlSelectQuery, `INNER JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "INNER JOIN")
			assertContains(t, sqlSelectQuery, `LEFT JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "LEFT JOIN")
			assertContains(t, sqlSelectQuery, `LEFT OUTER JOIN "test" AS "t" ON "t"."id" = "d"."id"`, "LEFT OUTER JOIN")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_clauseOrderBy(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "ORDER BY", "ORDER BY")
			assertContains(t, sqlSelectQuery, "`t`.`string` ASC", "ORDER BY ASC")
			assertContains(t, sqlSelectQuery, "`t`.`string` DESC", "ORDER BY DESC")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "ORDER BY", "ORDER BY")
			assertContains(t, sqlSelectQuery, "[t].[string] ASC", "ORDER BY ASC")
			assertContains(t, sqlSelectQuery, "[t].[string] DESC", "ORDER BY DESC")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "ORDER BY", "ORDER BY")
			assertContains(t, sqlSelectQuery, "`t`.`string` ASC", "ORDER BY ASC")
			assertContains(t, sqlSelectQuery, "`t`.`string` DESC", "ORDER BY DESC")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `ORDER BY`, "ORDER BY")
			assertContains(t, sqlSelectQuery, `"t"."string" ASC`, "ORDER BY ASC")
			assertContains(t, sqlSelectQuery, `"t"."string" DESC`, "ORDER BY DESC")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `ORDER BY`, "ORDER BY")
			assertContains(t, sqlSelectQuery, `"t"."string" ASC`, "ORDER BY ASC")
			assertContains(t, sqlSelectQuery, `"t"."string" DESC`, "ORDER BY DESC")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_clausePagination(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.Column.ID,
			).
			Pagination(10, 0)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "LIMIT ? OFFSET ?", "PAGINATION")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "OFFSET @p1 ROWS FETCH NEXT @p2 ROWS ONLY", "PAGINATION")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "LIMIT ? OFFSET ?", "PAGINATION")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `LIMIT $1 OFFSET $2`, "PAGINATION")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `LIMIT ? OFFSET ?`, "PAGINATION")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_clauseReturning(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtDelete := NewDelete(Test.Table).
			Returning(
				Test.Column.ID,
				Test.Column.String,
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlDeleteQuery, "RETURNING `t`.`id`, `t`.`string`", "RETURNING")
		case DialectMsSQL:
			assertContains(t, sqlDeleteQuery, "OUTPUT [t].[id], [t].[string]", "RETURNING")
		case DialectMySQL:
			// Not support
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `RETURNING "t"."id", "t"."string"`, "RETURNING")
		case DialectSQLite:
			assertContains(t, sqlDeleteQuery, `RETURNING "t"."id", "t"."string"`, "RETURNING")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_Core_clauseSet(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtUpdate := NewUpdate(Test.Table).
			Set(
				Assign(Test.Column.String, Value("active")),
			)
		sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlUpdateQuery, "SET `t`.`string` = ?", "SET")
		case DialectMsSQL:
			assertContains(t, sqlUpdateQuery, "SET [t].[string] = @p1", "SET")
		case DialectMySQL:
			assertContains(t, sqlUpdateQuery, "SET `t`.`string` = ?", "SET")
		case DialectPostgreSQL:
			assertContains(t, sqlUpdateQuery, `SET "t"."string" = $1`, "SET")
		case DialectSQLite:
			assertContains(t, sqlUpdateQuery, `SET "t"."string" = ?`, "SET")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
	})
}
func Test_Core_clauseUnions(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.Column.String,
			).
			Unions(
				Union(NewSelect(Test.Table).
					Field(
						Test.Column.String,
					),
				),
				UnionAll(NewSelect(Test.Table).
					Field(
						Test.Column.String,
					),
				),
				UnionExcept(NewSelect(Test.Table).
					Field(
						Test.Column.String,
					),
				),
				UnionIntersect(NewSelect(Test.Table).
					Field(
						Test.Column.String,
					),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "UNION SELECT `t`.`string` FROM `test` AS `t`", "UNION")
			assertContains(t, sqlSelectQuery, "UNION ALL SELECT `t`.`string` FROM `test` AS `t`", "UNION ALL")
			assertContains(t, sqlSelectQuery, "EXCEPT SELECT `t`.`string` FROM `test` AS `t`", "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, "INTERSECT SELECT `t`.`string` FROM `test` AS `t`", "UNION INTERSECT")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "UNION SELECT [t].[string] FROM [test] AS [t]", "UNION")
			assertContains(t, sqlSelectQuery, "UNION ALL SELECT [t].[string] FROM [test] AS [t]", "UNION ALL")
			assertContains(t, sqlSelectQuery, "EXCEPT SELECT [t].[string] FROM [test] AS [t]", "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, "INTERSECT SELECT [t].[string] FROM [test] AS [t]", "UNION INTERSECT")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "UNION SELECT `t`.`string` FROM `test` AS `t`", "UNION")
			assertContains(t, sqlSelectQuery, "UNION ALL SELECT `t`.`string` FROM `test` AS `t`", "UNION ALL")
			assertContains(t, sqlSelectQuery, "EXCEPT SELECT `t`.`string` FROM `test` AS `t`", "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, "INTERSECT SELECT `t`.`string` FROM `test` AS `t`", "UNION INTERSECT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `UNION SELECT "t"."string" FROM "test" AS "t"`, "UNION")
			assertContains(t, sqlSelectQuery, `UNION ALL SELECT "t"."string" FROM "test" AS "t"`, "UNION ALL")
			assertContains(t, sqlSelectQuery, `EXCEPT SELECT "t"."string" FROM "test" AS "t"`, "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, `INTERSECT SELECT "t"."string" FROM "test" AS "t"`, "UNION INTERSECT")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `UNION SELECT "t"."string" FROM "test" AS "t"`, "UNION")
			assertContains(t, sqlSelectQuery, `UNION ALL SELECT "t"."string" FROM "test" AS "t"`, "UNION ALL")
			assertContains(t, sqlSelectQuery, `EXCEPT SELECT "t"."string" FROM "test" AS "t"`, "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, `INTERSECT SELECT "t"."string" FROM "test" AS "t"`, "UNION INTERSECT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_clauseValues(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtInsert := NewInsert(Test.Table).
			Values(
				Pair(Test.Column.String, Value("ivan")),
				Pair(Test.Column.Number, Value(2)),
			).
			Upsert(
				Pair(Test.Column.String, Value("updated")),
			)
		sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlInsertQuery, "VALUES (?, ?)", "VALUES")
			assertContains(t, sqlInsertQuery, "ON DUPLICATE KEY UPDATE `string` = ?", "UPSERT")
		case DialectMsSQL:
			assertContains(t, sqlInsertQuery, "VALUES (@p1, @p2)", "VALUES")
		case DialectMySQL:
			assertContains(t, sqlInsertQuery, "VALUES (?, ?)", "VALUES")
			assertContains(t, sqlInsertQuery, "ON DUPLICATE KEY UPDATE `string` = ?", "UPSERT")
		case DialectPostgreSQL:
			assertContains(t, sqlInsertQuery, `VALUES ($1, $2)`, "VALUES")
			assertContains(t, sqlInsertQuery, `ON CONFLICT DO UPDATE SET "string" = $3`, "UPSERT")
		case DialectSQLite:
			assertContains(t, sqlInsertQuery, `VALUES (?, ?)`, "VALUES")
			assertContains(t, sqlInsertQuery, `ON CONFLICT DO UPDATE SET "string" = ?`, "UPSERT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
	})
}
func Test_Core_clauseWhere(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "WHERE `t`.`string` = ?", "WHERE")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "WHERE [t].[string] = @p1", "WHERE")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "WHERE `t`.`string` = ?", "WHERE")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `WHERE "t"."string" = $1`, "WHERE")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `WHERE "t"."string" = ?`, "WHERE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_clauseWith(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtWithN := WithN("cte_norecursive", NewSelect(Test.Table).
			Field(
				Test.Column.ID,
				Test.Column.String,
			).
			Where(
				Equal(Test.Column.String, Value("active")),
			),
			"id", "string",
		)
		stmtWithR := WithR("cte_recursive", NewSelect(Test.Table).
			Field(
				Test.Column.ID,
				Test.Column.String,
			).
			Where(
				Equal(Test.Column.String, Value("active")),
			).
			Unions(
				UnionAll(NewSelect(Test.Table).
					Field(
						Test.Column.ID,
						Test.Column.String,
					).
					Join(
						Inner(NewCTE("cte_recursive", "rec"), Equal(Test.Column.ID, Column[int64]("rec", "id"))),
					),
				),
			),
			"id", "string",
		)
		stmtSelect := NewSelect(Test.Table).
			Field(
				Test.Column.ID,
				Test.Column.Number,
			).
			Join(
				Inner(NewCTE("cte_norecursive", "cnr"), Equal(Test.Column.ID, Column[int64]("cnr", "id"))),
			).
			With(
				stmtWithR,
				stmtWithN,
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "WITH", "WITH")
			assertContains(t, sqlSelectQuery, "RECURSIVE `cte_recursive`", "WITH RECURSIVE")
			assertContains(t, sqlSelectQuery, "`cte_norecursive`", "WITH NORECURSIVE")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "WITH", "WITH")
			assertContains(t, sqlSelectQuery, "RECURSIVE [cte_recursive]", "WITH RECURSIVE")
			assertContains(t, sqlSelectQuery, "[cte_norecursive]", "WITH NORECURSIVE")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "WITH", "WITH")
			assertContains(t, sqlSelectQuery, "RECURSIVE `cte_recursive`", "WITH RECURSIVE")
			assertContains(t, sqlSelectQuery, "`cte_norecursive`", "WITH NORECURSIVE")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `WITH`, "WITH")
			assertContains(t, sqlSelectQuery, `RECURSIVE "cte_recursive"`, "WITH RECURSIVE")
			assertContains(t, sqlSelectQuery, `"cte_norecursive"`, "WITH NORECURSIVE")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `WITH`, "WITH")
			assertContains(t, sqlSelectQuery, `RECURSIVE "cte_recursive"`, "WITH RECURSIVE")
			assertContains(t, sqlSelectQuery, `"cte_norecursive"`, "WITH NORECURSIVE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_exprArray(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "(?, ?, ?)", "ARRAY INT")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "(@p1, @p2, @p3)", "ARRAY INT")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "(?, ?, ?)", "ARRAY INT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `($1, $2, $3)`, "ARRAY INT")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `(?, ?, ?)`, "ARRAY INT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_exprBinary(t *testing.T) {
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
		case DialectMariaDB:
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
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "[t].[number] & @p1", "BINARY BITWISE AND")
			assertContains(t, sqlSelectQuery, "[t].[number] | @p1", "BINARY BITWISE OR")
			assertContains(t, sqlSelectQuery, "[t].[number] ^ @p1", "BINARY BITWISE XOR")
			assertContains(t, sqlSelectQuery, "[t].[number] / @p1", "BINARY DIVIDE")
			assertContains(t, sqlSelectQuery, "[t].[number] - @p1", "BINARY MINUS")
			assertContains(t, sqlSelectQuery, "[t].[number] % @p1", "BINARY MODULO")
			assertContains(t, sqlSelectQuery, "[t].[number] * @p1", "BINARY MULTIPLY")
			assertContains(t, sqlSelectQuery, "[t].[number] + @p1", "BINARY PLUS")
			assertContains(t, sqlSelectQuery, "[t].[number] << @p1", "BINARY SHIFT LEFT")
			assertContains(t, sqlSelectQuery, "[t].[number] >> @p1", "BINARY SHIFT RIGHT")
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
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `"t"."number" & ?`, "BINARY BITWISE AND")
			assertContains(t, sqlSelectQuery, `"t"."number" | ?`, "BINARY BITWISE OR")
			assertContains(t, sqlSelectQuery, `"t"."number" ^ ?`, "BINARY BITWISE XOR")
			assertContains(t, sqlSelectQuery, `"t"."number" / ?`, "BINARY DIVIDE")
			assertContains(t, sqlSelectQuery, `"t"."number" - ?`, "BINARY MINUS")
			assertContains(t, sqlSelectQuery, `"t"."number" % ?`, "BINARY MODULO")
			assertContains(t, sqlSelectQuery, `"t"."number" * ?`, "BINARY MULTIPLY")
			assertContains(t, sqlSelectQuery, `"t"."number" + ?`, "BINARY PLUS")
			assertContains(t, sqlSelectQuery, `"t"."number" << ?`, "BINARY SHIFT LEFT")
			assertContains(t, sqlSelectQuery, `"t"."number" >> ?`, "BINARY SHIFT RIGHT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_exprColumn(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "`t`.`id`", "COLUMN ID")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "[t].[id]", "COLUMN ID")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`t`.`id`", "COLUMN ID")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"t"."id"`, "COLUMN ID")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `"t"."id"`, "COLUMN ID")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_exprComparison(t *testing.T) {
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
		case DialectMariaDB:
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
			assertContains(t, sqlSelectQuery, "`t`.`number` <= ?", "COMPARISON LESSEQUAL")
			assertContains(t, sqlSelectQuery, "`t`.`string` LIKE ?", "COMPARISON LIKE")
			assertContains(t, sqlSelectQuery, "`t`.`number` NOT BETWEEN ? AND ?", "COMPARISON NOT BETWEEN")
			assertContains(t, sqlSelectQuery, "`t`.`number` <> ?", "COMPARISON NOT EQUAL")
			assertContains(t, sqlSelectQuery, "NOT EXISTS (SELECT 1 FROM `test` AS `t`)", "COMPARISON NOT EXISTS")
			assertContains(t, sqlSelectQuery, "LOWER(`t`.`string`) NOT LIKE LOWER(?)", "COMPARISON NOT ILIKE")
			assertContains(t, sqlSelectQuery, "`t`.`string` NOT IN (?, ?)", "COMPARISON NOT IN")
			assertContains(t, sqlSelectQuery, "`t`.`string` NOT LIKE ?", "COMPARISON NOT LIKE")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "[t].[number] BETWEEN @p1 AND @p2", "COMPARISON BETWEEN")
			assertContains(t, sqlSelectQuery, "[t].[number] = @p1", "COMPARISON EQUAL")
			assertContains(t, sqlSelectQuery, "EXISTS (SELECT 1 FROM [test] AS [t])", "COMPARISON EXISTS")
			assertContains(t, sqlSelectQuery, "[t].[number] > @p1", "COMPARISON GREATER")
			assertContains(t, sqlSelectQuery, "[t].[number] >= @p1", "COMPARISON GREATEREQUAL")
			assertContains(t, sqlSelectQuery, "LOWER([t].[string]) LIKE LOWER(@p1)", "COMPARISON ILIKE")
			assertContains(t, sqlSelectQuery, "[t].[string] IN (@p1, @p2)", "COMPARISON IN")
			assertContains(t, sqlSelectQuery, "[t].[string] IS NOT NULL", "COMPARISON IS NOT NULL")
			assertContains(t, sqlSelectQuery, "[t].[string] IS NULL", "COMPARISON IS NULL")
			assertContains(t, sqlSelectQuery, "[t].[number] < @p1", "COMPARISON LESS")
			assertContains(t, sqlSelectQuery, "[t].[number] <= @p1", "COMPARISON LESSEQUAL")
			assertContains(t, sqlSelectQuery, "[t].[string] LIKE @p1", "COMPARISON LIKE")
			assertContains(t, sqlSelectQuery, "[t].[number] NOT BETWEEN @p1 AND @p2", "COMPARISON NOT BETWEEN")
			assertContains(t, sqlSelectQuery, "[t].[number] <> @p1", "COMPARISON NOT EQUAL")
			assertContains(t, sqlSelectQuery, "NOT EXISTS (SELECT 1 FROM [test] AS [t])", "COMPARISON NOT EXISTS")
			assertContains(t, sqlSelectQuery, "LOWER([t].[string]) NOT LIKE LOWER(@p1)", "COMPARISON NOT ILIKE")
			assertContains(t, sqlSelectQuery, "[t].[string] NOT IN (@p1, @p2)", "COMPARISON NOT IN")
			assertContains(t, sqlSelectQuery, "[t].[string] NOT LIKE @p1", "COMPARISON NOT LIKE")
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
			assertContains(t, sqlSelectQuery, "`t`.`number` <= ?", "COMPARISON LESSEQUAL")
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
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `"t"."number" BETWEEN ? AND ?`, "COMPARISON BETWEEN")
			assertContains(t, sqlSelectQuery, `"t"."number" = ?`, "COMPARISON EQUAL")
			assertContains(t, sqlSelectQuery, `EXISTS (SELECT 1 FROM "test" AS "t")`, "COMPARISON EXISTS")
			assertContains(t, sqlSelectQuery, `"t"."number" > ?`, "COMPARISON GREATER")
			assertContains(t, sqlSelectQuery, `"t"."number" >= ?`, "COMPARISON GREATEREQUAL")
			assertContains(t, sqlSelectQuery, `LOWER("t"."string") LIKE LOWER(?)`, "COMPARISON ILIKE")
			assertContains(t, sqlSelectQuery, `"t"."string" IN (?, ?)`, "COMPARISON IN")
			assertContains(t, sqlSelectQuery, `"t"."string" IS NOT NULL`, "COMPARISON IS NOT NULL")
			assertContains(t, sqlSelectQuery, `"t"."string" IS NULL`, "COMPARISON IS NULL")
			assertContains(t, sqlSelectQuery, `"t"."number" < ?`, "COMPARISON LESS")
			assertContains(t, sqlSelectQuery, `"t"."number" <= ?`, "COMPARISON LESSEQUAL")
			assertContains(t, sqlSelectQuery, `"t"."string" LIKE ?`, "COMPARISON LIKE")
			assertContains(t, sqlSelectQuery, `"t"."number" NOT BETWEEN ? AND ?`, "COMPARISON NOT BETWEEN")
			assertContains(t, sqlSelectQuery, `"t"."number" <> ?`, "COMPARISON NOT EQUAL")
			assertContains(t, sqlSelectQuery, `NOT EXISTS (SELECT 1 FROM "test" AS "t")`, "COMPARISON NOT EXISTS")
			assertContains(t, sqlSelectQuery, `LOWER("t"."string") NOT LIKE LOWER(?)`, "COMPARISON NOT ILIKE")
			assertContains(t, sqlSelectQuery, `"t"."string" NOT IN (?, ?)`, "COMPARISON NOT IN")
			assertContains(t, sqlSelectQuery, `"t"."string" NOT LIKE ?`, "COMPARISON NOT LIKE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_exprConstant(t *testing.T) {
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
		case DialectMariaDB:
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
		case DialectMsSQL:
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
		case DialectSQLite:
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
func Test_Core_exprFunction(t *testing.T) {
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
				CharLength(Test.Column.String).As("convert_charlength"),
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
		case DialectMariaDB:
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
		case DialectMsSQL:
			// Функции агрегатные
			assertContains(t, sqlSelectQuery, "AVG([t].[number])", "FUNCTION AVG")
			assertContains(t, sqlSelectQuery, "BIT_AND([t].[number])", "FUNCTION BITAND")
			assertContains(t, sqlSelectQuery, "BIT_OR([t].[number])", "FUNCTION BITOR")
			assertContains(t, sqlSelectQuery, "BIT_XOR([t].[number])", "FUNCTION BITXOR")
			assertContains(t, sqlSelectQuery, "COUNT([t].[string])", "FUNCTION COUNT")
			assertContains(t, sqlSelectQuery, "STRING_AGG([t].[string], ',')", "FUNCTION GROUPCONCAT")
			assertContains(t, sqlSelectQuery, "MAX([t].[number])", "FUNCTION MAX")
			assertContains(t, sqlSelectQuery, "MIN([t].[number])", "FUNCTION MIN")
			assertContains(t, sqlSelectQuery, "STDEV([t].[number])", "FUNCTION STDDEV")
			assertContains(t, sqlSelectQuery, "SUM([t].[number])", "FUNCTION SUM")
			assertContains(t, sqlSelectQuery, "VAR([t].[number])", "FUNCTION VARIANCE")
			// Функции аналитические
			assertContains(t, sqlSelectQuery, "FIRST_VALUE([t].[name]) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC", "FUNCTION FIRSTVALUE")
			assertContains(t, sqlSelectQuery, "LAG([t].[number], 2) OVER (PARTITION BY [t].[id] ORDER BY [t].[date] ASC)", "FUNCTION LAG")
			assertContains(t, sqlSelectQuery, "LAST_VALUE([t].[name]) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)", "FUNCTION LASTVALUE")
			assertContains(t, sqlSelectQuery, "LEAD([t].[number], 2) OVER (PARTITION BY [t].[id] ORDER BY [t].[date] ASC)", "FUNCTION LEAD")
			assertContains(t, sqlSelectQuery, "NTH_VALUE([t].[name], 2) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)", "FUNCTION NTHVALUE")
			// Функции условий
			assertContains(t, sqlSelectQuery, "CASE WHEN [t].[number] < @p1 THEN @p2 ELSE @p3 END", "FUNCTION CASE")
			assertContains(t, sqlSelectQuery, "COALESCE([t].[createat], [t].[updateat])", "FUNCTION COALESCE")
			assertContains(t, sqlSelectQuery, "GREATEST([t].[createat], [t].[updateat])", "FUNCTION GREATEST")
			assertContains(t, sqlSelectQuery, "LEAST([t].[createat], [t].[updateat])", "FUNCTION LEAST")
			assertContains(t, sqlSelectQuery, "NULLIF([t].[createat], [t].[updateat])", "FUNCTION NULLIF")
			// Функции конвертации
			assertContains(t, sqlSelectQuery, "CAST([t].[number] AS NVARCHAR)", "FUNCTION CAST")
			assertContains(t, sqlSelectQuery, "CHAR_LENGTH([t].[string])", "FUNCTION CHARLENGTH")
			assertContains(t, sqlSelectQuery, "FORMAT([t].[createat], '%Y-%m-%d')", "FUNCTION DATEFORMAT")
			assertContains(t, sqlSelectQuery, "DEGREES([t].[number])", "FUNCTION DEGREES")
			assertContains(t, sqlSelectQuery, "LEN([t].[string])", "FUNCTION LENGTH")
			assertContains(t, sqlSelectQuery, "CHARINDEX(@p1, [t].[string])", "FUNCTION POSITION")
			assertContains(t, sqlSelectQuery, "RADIANS([t].[number])", "FUNCTION RADIANS")
			// Функции даты и времени
			//assertContains(t, sqlSelectQuery, "CAST(GETDATE() AS DATE)", "FUNCTION CURDATE")
			//assertContains(t, sqlSelectQuery, "CAST(GETDATE() AS TIME)", "FUNCTION CURTIME")
			//assertContains(t, sqlSelectQuery, "DATE_ADD(DAY, 2, [t].[createat])", "FUNCTION DATEADD")
			assertContains(t, sqlSelectQuery, "DATEDIFF([t].[updateat], [t].[createat])", "FUNCTION DATEDIFF")
			//assertContains(t, sqlSelectQuery, "DATE_ADD(DAY, -2, [t].[createat])", "FUNCTION DATESUB")
			assertContains(t, sqlSelectQuery, "DAY([t].[createat])", "FUNCTION DAY")
			assertContains(t, sqlSelectQuery, "DATENAME(WEEKDAY, [t].[createat])", "FUNCTION DAYNAME")
			assertContains(t, sqlSelectQuery, "DATEPART(HOUR, [t].[createat])", "FUNCTION HOUR")
			assertContains(t, sqlSelectQuery, "DATEPART(MINUTE, [t].[createat])", "FUNCTION MINUTE")
			assertContains(t, sqlSelectQuery, "MONTH([t].[createat])", "FUNCTION MONTH")
			assertContains(t, sqlSelectQuery, "DATENAME(MONTH, [t].[createat])", "FUNCTION MONTHNAME")
			assertContains(t, sqlSelectQuery, "GETDATE()", "FUNCTION NOW")
			assertContains(t, sqlSelectQuery, "DATEPART(QUARTER, [t].[createat])", "FUNCTION QUARTER")
			assertContains(t, sqlSelectQuery, "DATEPART(SECOND, [t].[createat])", "FUNCTION SECOND")
			//assertContains(t, sqlSelectQuery, "DATE_ADD(HOUR, 2, [t].[createat])", "FUNCTION TIMEADD")
			assertContains(t, sqlSelectQuery, "TIMEDIFF([t].[updateat], [t].[createat])", "FUNCTION TIMEDIFF")
			//assertContains(t, sqlSelectQuery, "DATE_ADD(HOUR, -2, [t].[createat])", "FUNCTION TIMESUB")
			assertContains(t, sqlSelectQuery, "DATEPART(WEEK, [t].[createat])", "FUNCTION WEEK")
			assertContains(t, sqlSelectQuery, "YEAR([t].[createat])", "FUNCTION YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, "JSON_ARRAY([t].[json], @p1, @p2)", "FUNCTION JSONARRAY")
			assertContains(t, sqlSelectQuery, "JSON_ARRAYAGG([t].[json])", "FUNCTION JSONARRAYAGG")
			//assertContains(t, sqlSelectQuery, "JSON_CONTAINS([t].[json], ?)", "FUNCTION JSONCONTAINS")
			//assertContains(t, sqlSelectQuery, "([t].[json] ->> '$.parent[0].child')", "FUNCTION JSONEXTRACT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECT('key', COUNT([t].[json]))", "FUNCTION JSONOBJECT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECTAGG([t].[json], [t].[number])", "FUNCTION JSONOBJECTAGG")
			//assertContains(t, sqlSelectQuery, "JSON_MODIFY([t].[json], '$.key1', '$.key2')", "FUNCTION JSONREMOVE")
			//assertContains(t, sqlSelectQuery, "JSON_MODIFY([t].[json], '$.key1', ?, '$.key2', ?)", "FUNCTION JSONSET")
			//assertContains(t, sqlSelectQuery, "JSON_TYPE([t].[json])", "FUNCTION JSONTYPE")
			// Функции математические
			assertContains(t, sqlSelectQuery, "ABS([t].[number])", "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, "ACOS([t].[number])", "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, "ASIN([t].[number])", "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, "ATAN([t].[number])", "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, "ATAN2([t].[y], [t].[x])", "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, "CBRT([t].[number])", "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, "CEILING([t].[number])", "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, "COS([t].[number])", "FUNCTION COS")
			assertContains(t, sqlSelectQuery, "EXP([t].[number])", "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, "FLOOR([t].[number])", "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, "LN([t].[number])", "FUNCTION LN")
			assertContains(t, sqlSelectQuery, "LOG([t].[number], @p1)", "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, "MOD([t].[number], @p1)", "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, "PI()", "FUNCTION PI")
			assertContains(t, sqlSelectQuery, "POWER([t].[number], @p1)", "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, "RAND()", "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, "ROUND([t].[number], @p1)", "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, "SIN([t].[number])", "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, "SQRT([t].[number])", "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, "TAN([t].[number])", "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, "ROUND([t].[number], @p1, 1)", "FUNCTION TRUNC")
			// Функции ранжирующие
			assertContains(t, sqlSelectQuery, "CUME_DIST() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)", "FUNCTION CUMEDIST")
			assertContains(t, sqlSelectQuery, "DENSE_RANK() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)", "FUNCTION DENSERANK")
			assertContains(t, sqlSelectQuery, "NTILE(2) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)", "FUNCTION NTILE")
			assertContains(t, sqlSelectQuery, "PERCENT_RANK() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)", "FUNCTION PERCENTRANK")
			assertContains(t, sqlSelectQuery, "RANK() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)", "FUNCTION RANK")
			assertContains(t, sqlSelectQuery, "ROW_NUMBER() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)", "FUNCTION ROWNUMBER")
			// Функции строковые
			assertContains(t, sqlSelectQuery, "CONCAT([t].[string], @p1, @p2)", "FUNCTION CONCAT")
			assertContains(t, sqlSelectQuery, "CONCAT_WS(@p1, [t].[string], @p2, @p3)", "FUNCTION CONCATWS")
			assertContains(t, sqlSelectQuery, "LEFT([t].[string], @p1)", "FUNCTION LEFTSTRING")
			assertContains(t, sqlSelectQuery, "LOWER([t].[string])", "FUNCTION LOWER")
			assertContains(t, sqlSelectQuery, "LPAD([t].[string], @p1, @p2)", "FUNCTION LPAD")
			assertContains(t, sqlSelectQuery, "LTRIM([t].[string])", "FUNCTION LTRIM")
			assertContains(t, sqlSelectQuery, "REPEAT([t].[string], @p1)", "FUNCTION REPEAT")
			assertContains(t, sqlSelectQuery, "REPLACE([t].[string], @p1, @p2)", "FUNCTION REPLACE")
			assertContains(t, sqlSelectQuery, "REVERSE([t].[string])", "FUNCTION REVERSE")
			assertContains(t, sqlSelectQuery, "RIGHT([t].[string], @p1)", "FUNCTION RIGHTSTRING")
			assertContains(t, sqlSelectQuery, "RPAD([t].[string], @p1, @p2)", "FUNCTION RPAD")
			assertContains(t, sqlSelectQuery, "RTRIM([t].[string])", "FUNCTION RTRIM")
			assertContains(t, sqlSelectQuery, "SUBSTRING([t].[string], @p1, @p2)", "FUNCTION SUBSTRING")
			assertContains(t, sqlSelectQuery, "TRIM([t].[string])", "FUNCTION TRIM")
			assertContains(t, sqlSelectQuery, "UPPER([t].[string])", "FUNCTION UPPER")
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
		case DialectSQLite:
			// Функции агрегатные
			assertContains(t, sqlSelectQuery, `AVG("t"."number")`, "FUNCTION AVG")
			assertContains(t, sqlSelectQuery, `BIT_AND("t"."number")`, "FUNCTION BITAND")
			assertContains(t, sqlSelectQuery, `BIT_OR("t"."number")`, "FUNCTION BITOR")
			assertContains(t, sqlSelectQuery, `BIT_XOR("t"."number")`, "FUNCTION BITXOR")
			assertContains(t, sqlSelectQuery, `COUNT("t"."string")`, "FUNCTION COUNT")
			assertContains(t, sqlSelectQuery, `GROUP_CONCAT("t"."string" SEPARATOR ',')`, "FUNCTION GROUPCONCAT")
			assertContains(t, sqlSelectQuery, `MAX("t"."number")`, "FUNCTION MAX")
			assertContains(t, sqlSelectQuery, `MIN("t"."number")`, "FUNCTION MIN")
			assertContains(t, sqlSelectQuery, `STDEV("t"."number")`, "FUNCTION STDDEV")
			assertContains(t, sqlSelectQuery, `SUM("t"."number")`, "FUNCTION SUM")
			assertContains(t, sqlSelectQuery, `VARIANCE("t"."number")`, "FUNCTION VARIANCE")
			// Функции аналитические
			assertContains(t, sqlSelectQuery, `FIRST_VALUE("t"."name") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION FIRSTVALUE")
			assertContains(t, sqlSelectQuery, `LAG("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)`, "FUNCTION LAG")
			assertContains(t, sqlSelectQuery, `LAST_VALUE("t"."name") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)`, "FUNCTION LASTVALUE")
			assertContains(t, sqlSelectQuery, `LEAD("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)`, "FUNCTION LEAD")
			assertContains(t, sqlSelectQuery, `NTH_VALUE("t"."name", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)`, "FUNCTION NTHVALUE")
			// Функции условий
			assertContains(t, sqlSelectQuery, `CASE WHEN "t"."number" < ? THEN ? ELSE ? END`, "FUNCTION CASE")
			assertContains(t, sqlSelectQuery, `COALESCE("t"."createat", "t"."updateat")`, "FUNCTION COALESCE")
			assertContains(t, sqlSelectQuery, `GREATEST("t"."createat", "t"."updateat")`, "FUNCTION GREATEST")
			assertContains(t, sqlSelectQuery, `LEAST("t"."createat", "t"."updateat")`, "FUNCTION LEAST")
			assertContains(t, sqlSelectQuery, `NULLIF("t"."createat", "t"."updateat")`, "FUNCTION NULLIF")
			// Функции конвертации
			assertContains(t, sqlSelectQuery, `CAST("t"."number" AS TEXT) `, "FUNCTION CAST")
			assertContains(t, sqlSelectQuery, `CHAR_LENGTH("t"."string")`, "FUNCTION CHARLENGTH")
			assertContains(t, sqlSelectQuery, `strftime("t"."createat", '%Y-%m-%d')`, "FUNCTION DATEFORMAT")
			assertContains(t, sqlSelectQuery, `DEGREES("t"."number")`, "FUNCTION DEGREES")
			assertContains(t, sqlSelectQuery, `LENGTH("t"."string")`, "FUNCTION LENGTH")
			assertContains(t, sqlSelectQuery, `POSITION(? IN "t"."string")`, "FUNCTION POSITION")
			assertContains(t, sqlSelectQuery, `RADIANS("t"."number")`, "FUNCTION RADIANS")
			// Функции даты и времени
			assertContains(t, sqlSelectQuery, `date('now')`, "FUNCTION CURDATE")
			assertContains(t, sqlSelectQuery, `time('now')`, "FUNCTION CURTIME")
			assertContains(t, sqlSelectQuery, `datetime("t"."createat", '+2 DAY')`, "FUNCTION DATEADD")
			assertContains(t, sqlSelectQuery, `DATEDIFF("t"."updateat", "t"."createat")`, "FUNCTION DATEDIFF")
			assertContains(t, sqlSelectQuery, `datetime("t"."createat", '-2 DAY')`, "FUNCTION DATESUB")
			assertContains(t, sqlSelectQuery, `DAY("t"."createat")`, "FUNCTION DAY")
			assertContains(t, sqlSelectQuery, `strftime('%w', "t"."createat")`, "FUNCTION DAYNAME")
			assertContains(t, sqlSelectQuery, `HOUR("t"."createat")`, "FUNCTION HOUR")
			assertContains(t, sqlSelectQuery, `MINUTE("t"."createat")`, "FUNCTION MINUTE")
			assertContains(t, sqlSelectQuery, `MONTH("t"."createat")`, "FUNCTION MONTH")
			assertContains(t, sqlSelectQuery, `strftime('%m', "t"."createat")`, "FUNCTION MONTHNAME")
			assertContains(t, sqlSelectQuery, `datetime('now')`, "FUNCTION NOW")
			assertContains(t, sqlSelectQuery, `QUARTER("t"."createat")`, "FUNCTION QUARTER")
			assertContains(t, sqlSelectQuery, `SECOND("t"."createat")`, "FUNCTION SECOND")
			assertContains(t, sqlSelectQuery, `time("t"."createat", '+2 HOUR')`, "FUNCTION TIMEADD")
			assertContains(t, sqlSelectQuery, `TIMEDIFF("t"."updateat", "t"."createat")`, "FUNCTION TIMEDIFF")
			assertContains(t, sqlSelectQuery, `time("t"."createat", '-2 HOUR')`, "FUNCTION TIMESUB")
			assertContains(t, sqlSelectQuery, `WEEK("t"."createat")`, "FUNCTION WEEK")
			assertContains(t, sqlSelectQuery, `YEAR("t"."createat")`, "FUNCTION YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, `JSON_ARRAY("t"."json", ?, ?)`, "FUNCTION JSONARRAY")
			assertContains(t, sqlSelectQuery, `JSON_GROUP_ARRAY("t"."json")`, "FUNCTION JSONARRAYAGG")
			assertContains(t, sqlSelectQuery, `JSON_CONTAINS("t"."json", ?)`, "FUNCTION JSONCONTAINS")
			assertContains(t, sqlSelectQuery, `("t"."json" ->> '$.parent[0].child')`, "FUNCTION JSONEXTRACT")
			assertContains(t, sqlSelectQuery, `JSON_OBJECT('key', COUNT("t"."json"))`, "FUNCTION JSONOBJECT")
			assertContains(t, sqlSelectQuery, `JSON_GROUP_OBJECT("t"."json", "t"."number")`, "FUNCTION JSONOBJECTAGG")
			assertContains(t, sqlSelectQuery, `JSON_REMOVE("t"."json", '$.key1', '$.key2')`, "FUNCTION JSONREMOVE")
			assertContains(t, sqlSelectQuery, `JSON_SET("t"."json", '$.key1', ?, '$.key2', ?)`, "FUNCTION JSONSET")
			assertContains(t, sqlSelectQuery, `JSON_TYPE("t"."json")`, "FUNCTION JSONTYPE")
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
			assertContains(t, sqlSelectQuery, `LOG("t"."number", ?)`, "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, `MOD("t"."number", ?)`, "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, `PI()`, "FUNCTION PI")
			assertContains(t, sqlSelectQuery, `POWER("t"."number", ?)`, "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, `RANDOM`, "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, `ROUND("t"."number", ?)`, "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, `SIN("t"."number")`, "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, `SQRT("t"."number")`, "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, `TAN("t"."number")`, "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, `TRUNC("t"."number", ?)`, "FUNCTION TRUNC")
			// Функции ранжирующие
			assertContains(t, sqlSelectQuery, `CUME_DIST() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION CUMEDIST")
			assertContains(t, sqlSelectQuery, `DENSE_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "DFUNCTION ENSERANK")
			assertContains(t, sqlSelectQuery, `NTILE(2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION NTILE")
			assertContains(t, sqlSelectQuery, `PERCENT_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION PERCENTRANK")
			assertContains(t, sqlSelectQuery, `RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION RANK")
			assertContains(t, sqlSelectQuery, `ROW_NUMBER() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)`, "FUNCTION ROWNUMBER")
			// Функции строковые
			assertContains(t, sqlSelectQuery, `CONCAT("t"."string", ?, ?)`, "FUNCTION CONCAT")
			assertContains(t, sqlSelectQuery, `CONCAT_WS(?, "t"."string", ?, ?)`, "FUNCTION CONCATWS")
			assertContains(t, sqlSelectQuery, `LEFT("t"."string", ?)`, "FUNCTION LEFTSTRING")
			assertContains(t, sqlSelectQuery, `LOWER("t"."string")`, "FUNCTION LOWER")
			assertContains(t, sqlSelectQuery, `LPAD("t"."string", ?, ?)`, "FUNCTION LPAD")
			assertContains(t, sqlSelectQuery, `LTRIM("t"."string")`, "FUNCTION LTRIM")
			assertContains(t, sqlSelectQuery, `REPEAT("t"."string", ?)`, "FUNCTION REPEAT")
			assertContains(t, sqlSelectQuery, `REPLACE("t"."string", ?, ?)`, "FUNCTION REPLACE")
			assertContains(t, sqlSelectQuery, `REVERSE("t"."string")`, "FUNCTION REVERSE")
			assertContains(t, sqlSelectQuery, `RIGHT("t"."string", ?)`, "FUNCTION RIGHTSTRING")
			assertContains(t, sqlSelectQuery, `RPAD("t"."string", ?, ?)`, "FUNCTION RPAD")
			assertContains(t, sqlSelectQuery, `RTRIM("t"."string")`, "FUNCTION RTRIM")
			assertContains(t, sqlSelectQuery, `SUBSTRING("t"."string", ?, ?)`, "FUNCTION SUBSTRING")
			assertContains(t, sqlSelectQuery, `TRIM("t"."string")`, "FUNCTION TRIM")
			assertContains(t, sqlSelectQuery, `UPPER("t"."string")`, "FUNCTION UPPER")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_exprLiteral(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "'%Y-%m-%d'", "LITERAL STRING")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "'%Y-%m-%d'", "LITERAL STRING")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "'%Y-%m-%d'", "LITERAL STRING")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `'%Y-%m-%d'`, "LITERAL STRING")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `'%Y-%m-%d'`, "LITERAL STRING")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_exprLogical(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "`t`.`string` = ? AND `t`.`number` > ?", "LOGICAL AND")
			assertContains(t, sqlSelectQuery, "`t`.`string` = ? OR `t`.`number` > ?", "LOGICAL OR")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "[t].[string] = @p1 AND [t].[number] > @p2", "LOGICAL AND")
			assertContains(t, sqlSelectQuery, "[t].[string] = @p1 OR [t].[number] > @p2", "LOGICAL OR")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`t`.`string` = ? AND `t`.`number` > ?", "LOGICAL AND")
			assertContains(t, sqlSelectQuery, "`t`.`string` = ? OR `t`.`number` > ?", "LOGICAL OR")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"t"."string" = $1 AND "t"."number" > $2`, "LOGICAL AND")
			assertContains(t, sqlSelectQuery, `"t"."string" = $1 OR "t"."number" > $2`, "LOGICAL OR")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `"t"."string" = ? AND "t"."number" > ?`, "LOGICAL AND")
			assertContains(t, sqlSelectQuery, `"t"."string" = ? OR "t"."number" > ?`, "LOGICAL OR")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_exprSubquery(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "(SELECT `t`.`id` FROM `test` AS `t`)", "SUBQUERY")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "(SELECT [t].[id] FROM [test] AS [t])", "SUBQUERY")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "(SELECT `t`.`id` FROM `test` AS `t`)", "SUBQUERY")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `(SELECT "t"."id" FROM "test" AS "t")`, "SUBQUERY")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `(SELECT "t"."id" FROM "test" AS "t")`, "SUBQUERY")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Core_exprValue(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "`t`.`string` = ?", "VALUE PLACEHOLDER")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "[t].[string] = @p1", "VALUE PLACEHOLDER")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`t`.`string` = ?", "VALUE PLACEHOLDER")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"t"."string" = $1`, "VALUE PLACEHOLDER")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `"t"."string" = ?`, "VALUE PLACEHOLDER")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_SQL(t *testing.T) {
	t.Run("Delete", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmt := NewDelete(Test.Table).
				Join(
					Inner(Data.Table, Equal(Test.Column.ID, Data.Column.ID)),
				).
				Where(
					And(
						Equal(Test.Column.String, Value("active")),
						ILike(Test.Column.String, Value("%ivan%")),
					),
				)
			query1, _, _ := sql.Build(stmt)
			nextDialect := getNextDialect(supportDialect)
			sql.SetDialect(nextDialect)
			query2, _, _ := sql.Build(stmt)
			sql.SetDialect(supportDialect)
			query3, _, _ := sql.Build(stmt)
			if query1 != query3 {
				t.Errorf("AST мутировал!\n  Ожидалось: %s\n  Получено:  %s", query1, query3)
			}
			t.Logf("Query1 (%s): %s", supportDialect.name, query1)
			t.Logf("Query2 (%s): %s", nextDialect.name, query2)
			t.Logf("Query3 (%s): %s", supportDialect.name, query3)
		})
	})
	t.Run("Insert", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmt := NewInsert(Test.Table).
				Values(
					Pair(Test.Column.String, Value("ivan")),
					Pair(Test.Column.Number, Value(2)),
				).
				Upsert(
					Pair(Test.Column.String, Value("updated")),
				)
			query1, _, _ := sql.Build(stmt)
			nextDialect := getNextDialect(supportDialect)
			sql.SetDialect(nextDialect)
			query2, _, _ := sql.Build(stmt)
			sql.SetDialect(supportDialect)
			query3, _, _ := sql.Build(stmt)
			if query1 != query3 {
				t.Errorf("AST мутировал!\n  Ожидалось: %s\n  Получено:  %s", query1, query3)
			}
			t.Logf("Query1 (%s): %s", supportDialect.name, query1)
			t.Logf("Query2 (%s): %s", nextDialect.name, query2)
			t.Logf("Query3 (%s): %s", supportDialect.name, query3)
		})
	})
	t.Run("Select", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmt := NewSelect(Test.Table).
				Field(
					Avg(Test.Column.Number, false).As("avg_result"),
					Ceil(Test.Column.Number).As("ceil_result"),
					Count(Test.Column.String, false).As("count_result"),
					FirstValue(Test.Column.Name).Over(
						PartitionBy(Test.Column.ID),
						OrderBy(Desc(Test.Column.Number)),
					).As("first_value"),
					Trunc(Test.Column.Number, Value(2)).As("trunc_result"),
				).
				Join(
					Inner(Data.Table, Equal(Test.Column.ID, Data.Column.ID)),
				).
				Where(
					And(
						Equal(Test.Column.String, Value("active")),
						Greater(Test.Column.Number, Value(2)),
						ILike(Test.Column.String, Value("%ivan%")),
					),
				).
				GroupBy(
					Test.Column.ID,
					Test.Column.String,
				).
				Having(
					Greater(Count(Test.Column.ID, false), Value[int64](2)),
				).
				OrderBy(
					Desc(Test.Column.Number),
				)
			query1, _, _ := sql.Build(stmt)
			nextDialect := getNextDialect(supportDialect)
			sql.SetDialect(nextDialect)
			query2, _, _ := sql.Build(stmt)
			sql.SetDialect(supportDialect)
			query3, _, _ := sql.Build(stmt)
			if query1 != query3 {
				t.Errorf("AST мутировал после смены диалекта!\n  Ожидалось: %s\n  Получено:  %s", query1, query3)
			}
			t.Logf("Query1 (%s): %s", supportDialect.name, query1)
			t.Logf("Query2 (%s): %s", nextDialect.name, query2)
			t.Logf("Query3 (%s): %s", supportDialect.name, query3)
		})
	})
	t.Run("Update", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmt := NewUpdate(Test.Table).
				Set(
					Assign(Test.Column.String, Value("updated")),
					Assign(Test.Column.Number, Value(2)),
				).
				Join(
					Inner(Data.Table, Equal(Test.Column.ID, Data.Column.ID)),
				).
				Where(
					And(
						Equal(Test.Column.String, Value("active")),
						ILike(Test.Column.String, Value("%ivan%")),
					),
				)
			query1, _, _ := sql.Build(stmt)
			nextDialect := getNextDialect(supportDialect)
			sql.SetDialect(nextDialect)
			query2, _, _ := sql.Build(stmt)
			sql.SetDialect(supportDialect)
			query3, _, _ := sql.Build(stmt)
			if query1 != query3 {
				t.Errorf("AST мутировал!\n  Ожидалось: %s\n  Получено:  %s", query1, query3)
			}
			t.Logf("Query1 (%s): %s", supportDialect.name, query1)
			t.Logf("Query2 (%s): %s", nextDialect.name, query2)
			t.Logf("Query3 (%s): %s", supportDialect.name, query3)
		})
	})
}
func Test_SQL_Comment(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtComment := NewComment("Test comment").OnColumn(Test.Column.ID)
		sqlCommentQuery, sqlCommentArguments, err := sql.Build(stmtComment)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlCommentQuery, "COMMENT ON COLUMN `t`.`id` IS 'Test comment'", "COMMENT")
		case DialectMsSQL:
			// Not supported
		case DialectMySQL:
			assertContains(t, sqlCommentQuery, "COMMENT ON COLUMN `t`.`id` IS 'Test comment'", "COMMENT")
		case DialectPostgreSQL:
			assertContains(t, sqlCommentQuery, `COMMENT ON COLUMN "t"."id" IS 'Test comment'`, "COMMENT")
		case DialectSQLite:
			assertContains(t, sqlCommentQuery, `COMMENT ON COLUMN "t"."id" IS 'Test comment'`, "COMMENT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCommentArguments, supportDialect.name, sqlCommentQuery)
	})
}
func Test_SQL_Comment_Table(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtComment := NewComment("Test comment").OnTable(Test.Table)
		sqlCommentQuery, sqlCommentArguments, err := sql.Build(stmtComment)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlCommentQuery, "COMMENT ON TABLE `test` AS `t` IS 'Test comment'", "COMMENT")
		case DialectMsSQL:
			// Not supported
		case DialectMySQL:
			assertContains(t, sqlCommentQuery, "COMMENT ON TABLE `test` AS `t` IS 'Test comment'", "COMMENT")
		case DialectPostgreSQL:
			assertContains(t, sqlCommentQuery, `COMMENT ON TABLE "test" AS "t" IS 'Test comment'`, "COMMENT")
		case DialectSQLite:
			assertContains(t, sqlCommentQuery, `COMMENT ON TABLE "test" AS "t" IS 'Test comment'`, "COMMENT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCommentArguments, supportDialect.name, sqlCommentQuery)
	})
}
func Test_SQL_Delete(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlDeleteQuery, "DELETE `t` FROM `test` AS `t`", "DELETE")
		case DialectMsSQL:
			assertContains(t, sqlDeleteQuery, "DELETE [t] FROM [test] AS [t]", "DELETE")
		case DialectMySQL:
			assertContains(t, sqlDeleteQuery, "DELETE `t` FROM `test` AS `t`", "DELETE")
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `DELETE FROM "test" AS "t"`, "DELETE")
		case DialectSQLite:
			assertContains(t, sqlDeleteQuery, `DELETE FROM "test" AS "t"`, "DELETE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
	})
}
func Test_SQL_Drop(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtDrop := NewDrop()
		sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlDropQuery, "DROP", "DROP")
		case DialectMsSQL:
			assertContains(t, sqlDropQuery, "DROP", "DROP")
		case DialectMySQL:
			assertContains(t, sqlDropQuery, "DROP", "DROP")
		case DialectPostgreSQL:
			assertContains(t, sqlDropQuery, `DROP`, "DROP")
		case DialectSQLite:
			assertContains(t, sqlDropQuery, `DROP`, "DROP")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDropArguments, supportDialect.name, sqlDropQuery)
	})
}
func Test_SQL_Insert(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtInsert := NewInsert(Test.Table).
			Values(
				Pair(Test.Column.String, Value("ivan")),
				Pair(Test.Column.Number, Value(2)),
			)
		sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlInsertQuery, "INSERT INTO `test` AS `t` (`string`, `number`)", "INSERT")
		case DialectMsSQL:
			assertContains(t, sqlInsertQuery, "INSERT INTO [test] AS [t] ([string], [number])", "INSERT")
		case DialectMySQL:
			assertContains(t, sqlInsertQuery, "INSERT INTO `test` AS `t` (`string`, `number`)", "INSERT")
		case DialectPostgreSQL:
			assertContains(t, sqlInsertQuery, `INSERT INTO "test" AS "t" ("string", "number")`, "INSERT")
		case DialectSQLite:
			assertContains(t, sqlInsertQuery, `INSERT INTO "test" AS "t" ("string", "number")`, "INSERT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
	})
}
func Test_SQL_Insert_Source(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtInsert := NewInsert(Data.Table).
			Source(NewSelect(Test.Table).
				Field(
					Test.Column.String,
					Test.Column.Number,
				).
				Where(
					Equal(Test.Column.String, Value("active")),
				),
			)
		sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlInsertQuery, "SELECT `t`.`string`, `t`.`number` FROM `test` AS `t` WHERE `t`.`string` = ?", "SOURCE")
		case DialectMsSQL:
			assertContains(t, sqlInsertQuery, "SELECT [t].[string], [t].[number] FROM [test] AS [t] WHERE [t].[string] = @p1", "SOURCE")
		case DialectMySQL:
			assertContains(t, sqlInsertQuery, "SELECT `t`.`string`, `t`.`number` FROM `test` AS `t` WHERE `t`.`string` = ?", "SOURCE")
		case DialectPostgreSQL:
			assertContains(t, sqlInsertQuery, `SELECT "t"."string", "t"."number" FROM "test" AS "t" WHERE "t"."string" = $1`, "SOURCE")
		case DialectSQLite:
			assertContains(t, sqlInsertQuery, `SELECT "t"."string", "t"."number" FROM "test" AS "t" WHERE "t"."string" = ?`, "SOURCE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
	})
}
func Test_SQL_Select(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "SELECT `t`.`string` FROM `test` AS `t`", "SELECT")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "SELECT [t].[string] FROM [test] AS [t]", "SELECT")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "SELECT `t`.`string` FROM `test` AS `t`", "SELECT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `SELECT "t"."string" FROM "test" AS "t"`, "SELECT")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `SELECT "t"."string" FROM "test" AS "t"`, "SELECT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_SQL_Select_Distinct(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "DISTINCT", "DISTINCT")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "DISTINCT", "DISTINCT")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "DISTINCT", "DISTINCT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `DISTINCT`, "DISTINCT")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `DISTINCT`, "DISTINCT")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_SQL_Truncate(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtTruncate := NewTruncate(Test.Table).Cascade().RestartIdentity()
		sqlTruncateQuery, sqlTruncateArguments, err := sql.Build(stmtTruncate)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlTruncateQuery, "TRUNCATE TABLE `test`", "TRUNCATE")
		case DialectMsSQL:
			assertContains(t, sqlTruncateQuery, "TRUNCATE TABLE [test]", "TRUNCATE")
		case DialectMySQL:
			assertContains(t, sqlTruncateQuery, "TRUNCATE TABLE `test`", "TRUNCATE")
		case DialectPostgreSQL:
			assertContains(t, sqlTruncateQuery, `TRUNCATE TABLE "test"`, "TRUNCATE")
		case DialectSQLite:
			assertContains(t, sqlTruncateQuery, `TRUNCATE TABLE "test"`, "TRUNCATE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlTruncateArguments, supportDialect.name, sqlTruncateQuery)
	})
}
func Test_SQL_Update(t *testing.T) {
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
		case DialectMariaDB:
			assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t`", "UPDATE")
		case DialectMsSQL:
			assertContains(t, sqlUpdateQuery, "UPDATE [test] AS [t]", "UPDATE")
		case DialectMySQL:
			assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t`", "UPDATE")
		case DialectPostgreSQL:
			assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t"`, "UPDATE")
		case DialectSQLite:
			assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t"`, "UPDATE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
	})
}
func Test_Transformer_Comparison(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(WithDialect(supportDialect))
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(Test.Column.ID).
			Where(
				And(
					ILike(Test.Column.String, Value("%alex%")),
					And(
						ILike(Test.Column.String, Value("%ivan%")),
						ILike(Test.Column.String, Value("%petr%")),
					),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "(LOWER(`t`.`string`) LIKE LOWER(?) AND LOWER(`t`.`string`) LIKE LOWER(?) AND LOWER(`t`.`string`) LIKE LOWER(?))", "ILIKE")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "(LOWER([t].[string]) LIKE LOWER(@p1) AND LOWER([t].[string]) LIKE LOWER(@p2) AND LOWER([t].[string]) LIKE LOWER(@p3))", "ILIKE")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "(LOWER(`t`.`string`) LIKE LOWER(?) AND LOWER(`t`.`string`) LIKE LOWER(?) AND LOWER(`t`.`string`) LIKE LOWER(?))", "ILIKE")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `("t"."string" ILIKE $1 AND "t"."string" ILIKE $2 AND "t"."string" ILIKE $3)`, "ILIKE")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `(LOWER("t"."string") LIKE LOWER(?) AND LOWER("t"."string") LIKE LOWER(?) AND LOWER("t"."string") LIKE LOWER(?))`, "ILIKE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Transformer_Function(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(WithDialect(supportDialect))
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Field(
				Trunc(Ceil(Test.Column.Number), Value(2)).As("result"),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "TRUNCATE(CEILING(`t`.`number`), ?)", "TRUNC→TRUNCATE, CEIL→CEILING")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "ROUND(CEILING([t].[number]), @p1, 1)", "TRUNC→TRUNCATE, CEIL→CEILING")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "TRUNCATE(CEILING(`t`.`number`), ?)", "TRUNC→TRUNCATE, CEIL→CEILING")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `TRUNC(CEIL("t"."number"), $1)`, "TRUNC→TRUNC, CEIL→CEIL")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `TRUNC(CEIL("t"."number"), ?)`, "TRUNC→TRUNC, CEIL→CEIL")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}

// Приватные функции
func assertContains(t *testing.T, str, substr string, message string) {
	t.Helper()
	re := regexp.MustCompile(`@p\d+|\$\d+`)
	strNormalized := re.ReplaceAllString(str, `?`)
	substrNormalized := re.ReplaceAllString(substr, `?`)
	if !strings.Contains(strNormalized, substrNormalized) {
		t.Errorf("Req: [%s] / Pat: [%s]\n  Str: [%s]", message, substr, str)
	}
}
func getNextDialect(dialect *SupportDialect) *SupportDialect {
	switch dialect {
	case DialectMariaDB:
		return DialectMySQL
	case DialectMySQL:
		return DialectPostgreSQL
	case DialectPostgreSQL:
		return DialectSQLite
	case DialectSQLite:
		return DialectMariaDB
	}
	return DialectPostgreSQL
}
func init() {
	// Data
	Data.Column.Date = Column[time.Time](Data.Table.aliasName, "date")
	Data.Column.ID = Column[int64](Data.Table.aliasName, "id")
	Data.Column.Json = Column[string](Data.Table.aliasName, "json")
	Data.Column.Number = Column[int](Data.Table.aliasName, "number")
	Data.Column.String = Column[string](Data.Table.aliasName, "string")
	Data.Column.Time = Column[time.Time](Data.Table.aliasName, "time")
	// Test
	Test.Column.CreateAt = Column[time.Time](Test.Table.aliasName, "createat")
	Test.Column.Date = Column[time.Time](Test.Table.aliasName, "date")
	Test.Column.ID = Column[int64](Test.Table.aliasName, "id")
	Test.Column.Json = Column[string](Test.Table.aliasName, "json")
	Test.Column.Name = Column[string](Test.Table.aliasName, "name")
	Test.Column.Number = Column[int](Test.Table.aliasName, "number")
	Test.Column.String = Column[string](Test.Table.aliasName, "string")
	Test.Column.UpdateAt = Column[time.Time](Test.Table.aliasName, "updateat")
	Test.Column.X = Column[int](Test.Table.aliasName, "x")
	Test.Column.Y = Column[int](Test.Table.aliasName, "y")
}
func testAllDialects(t *testing.T, testFunc func(t *testing.T, supportDialect *SupportDialect)) {
	for _, supportDialect := range listSupportDialects {
		currentDialect := supportDialect
		t.Run(currentDialect.name, func(t *testing.T) {
			testFunc(t, currentDialect)
		})
	}
}
