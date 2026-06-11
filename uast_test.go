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
		Date   *ColumnSource[time.Time]
		ID     *ColumnSource[int64]
		Json   *ColumnSource[string]
		Number *ColumnSource[int]
		String *ColumnSource[string]
		Time   *ColumnSource[time.Time]
	}
	Table *TableSource
}{
	Table: NewTable("data", "d"),
}
var Test = struct {
	Column struct {
		CreateAt *ColumnSource[time.Time]
		DataID   *ColumnSource[int64]
		Date     *ColumnSource[time.Time]
		ID       *ColumnSource[int64]
		Json     *ColumnSource[string]
		Name     *ColumnSource[string]
		Number   *ColumnSource[int]
		String   *ColumnSource[string]
		UpdateAt *ColumnSource[time.Time]
		X        *ColumnSource[int]
		Y        *ColumnSource[int]
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
			Fields(
				Test.Column.String.Expr(),
			).
			GroupBy(
				Test.Column.String.Expr(),
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
			Fields(
				Test.Column.String.Expr(),
			).
			Having(
				Greater(Count(Test.Column.ID.Expr(), false), Value[int64](2)),
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
			Fields(
				Data.Column.ID.Expr(),
			).
			Join(
				Cross(Test.Table),
				Full(Test.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				FullOuter(Test.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				Inner(Test.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				Left(Test.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				LeftOuter(Test.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				Right(Test.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				RightOuter(Test.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
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
			// Not supported - RIGHT JOIN
			// Not supported - RIGHT OUTER JOIN
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
			Fields(
				Test.Column.ID.Expr(),
			).
			OrderBy(
				Asc(Test.Column.String.Expr()),
				Desc(Test.Column.String.Expr()),
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
			Fields(
				Test.Column.ID.Expr(),
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
				Test.Column.ID.Expr(),
				Test.Column.String.Expr(),
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlDeleteQuery, "RETURNING `t`.`id`, `t`.`string`", "RETURNING")
		case DialectMsSQL:
			assertContains(t, sqlDeleteQuery, "OUTPUT [t].[id], [t].[string]", "RETURNING")
		case DialectMySQL:
			// Not supported - RETURNING
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
				Assign(Test.Column.String.Expr(), Value("active")),
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
			Fields(
				Test.Column.String.Expr(),
			).
			Unions(
				Union(NewSelect(Test.Table).
					Fields(
						Test.Column.String.Expr(),
					),
				),
				UnionAll(NewSelect(Test.Table).
					Fields(
						Test.Column.String.Expr(),
					),
				),
				UnionExcept(NewSelect(Test.Table).
					Fields(
						Test.Column.String.Expr(),
					),
				),
				UnionIntersect(NewSelect(Test.Table).
					Fields(
						Test.Column.String.Expr(),
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
				Pair(Test.Column.String.Expr(), Value("ivan")),
				Pair(Test.Column.Number.Expr(), Value(2)),
			).
			Upsert(
				Pair(Test.Column.String.Expr(), Value("updated")),
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
			Fields(
				Test.Column.ID.Expr(),
			).
			Where(
				Equal(Test.Column.String.Expr(), Value("active")),
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
			Fields(
				Test.Column.ID.Expr(),
				Test.Column.String.Expr(),
			).
			Where(
				Equal(Test.Column.String.Expr(), Value("active")),
			),
			"id", "string",
		)
		stmtWithR := WithR("cte_recursive", NewSelect(Test.Table).
			Fields(
				Test.Column.ID.Expr(),
				Test.Column.String.Expr(),
			).
			Where(
				Equal(Test.Column.String.Expr(), Value("active")),
			).
			Unions(
				UnionAll(NewSelect(Test.Table).
					Fields(
						Test.Column.ID.Expr(),
						Test.Column.String.Expr(),
					).
					Join(
						Inner(NewCTE("cte_recursive", "rec"), Equal(Test.Column.ID.Expr(), Field[int64]("rec", "id"))),
					),
				),
			),
			"id", "string",
		)
		stmtSelect := NewSelect(Test.Table).
			Fields(
				Test.Column.ID.Expr(),
				Test.Column.Number.Expr(),
			).
			Join(
				Inner(NewCTE("cte_norecursive", "cnr"), Equal(Test.Column.ID.Expr(), Field[int64]("cnr", "id"))),
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
			Fields(
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
			Fields(
				Test.Column.ID.Expr(),
			).
			Where(
				And(
					Equal(Test.Column.Number.Expr(), BitwiseAnd(Test.Column.Number.Expr(), Value(0b0010))),
					Equal(Test.Column.Number.Expr(), BitwiseOr(Test.Column.Number.Expr(), Value(0b0010))),
					Equal(Test.Column.Number.Expr(), BitwiseXor(Test.Column.Number.Expr(), Value(0b0010))),
					Equal(Test.Column.Number.Expr(), Divide(Test.Column.Number.Expr(), Value(2))),
					Equal(Test.Column.Number.Expr(), Minus(Test.Column.Number.Expr(), Value(2))),
					Equal(Test.Column.Number.Expr(), Modulo(Test.Column.Number.Expr(), Value(2))),
					Equal(Test.Column.Number.Expr(), Multiply(Test.Column.Number.Expr(), Value(2))),
					Equal(Test.Column.Number.Expr(), Plus(Test.Column.Number.Expr(), Value(2))),
					Equal(Test.Column.Number.Expr(), ShiftLeft(Test.Column.Number.Expr(), Value(2))),
					Equal(Test.Column.Number.Expr(), ShiftRight(Test.Column.Number.Expr(), Value(2))),
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
			Fields(
				Test.Column.ID.Expr(),
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
			Fields(
				Test.Column.ID.Expr(),
			).
			Where(
				And(
					Between(Test.Column.Number.Expr(), Value(0), Value(2)),
					Equal(Test.Column.Number.Expr(), Value(2)),
					Exists(Subquery[int](NewSelect(Test.Table).Fields(ConstIntOne()))),
					Greater(Test.Column.Number.Expr(), Value(2)),
					GreaterEqual(Test.Column.Number.Expr(), Value(2)),
					ILike(Test.Column.String.Expr(), Value("%ivan%")),
					In(Test.Column.String.Expr(), Array("active", "pending")),
					IsNotNull(Test.Column.String.Expr()),
					IsNull(Test.Column.String.Expr()),
					Less(Test.Column.Number.Expr(), Value(2)),
					LessEqual(Test.Column.Number.Expr(), Value(2)),
					Like(Test.Column.String.Expr(), Value("%ivan%")),
					NotBetween(Test.Column.Number.Expr(), Value(0), Value(2)),
					NotEqual(Test.Column.Number.Expr(), Value(2)),
					NotExists(Subquery[int](NewSelect(Test.Table).Fields(ConstIntOne()))),
					NotILike(Test.Column.String.Expr(), Value("%ivan%")),
					NotIn(Test.Column.String.Expr(), Array("active", "pending")),
					NotLike(Test.Column.String.Expr(), Value("%ivan%")),
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
			Fields(
				Test.Column.ID.Expr(),
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
			Fields(
				// Функции агрегатные
				Avg(Test.Column.Number.Expr(), false).As("aggregate_avg"),
				BitAnd(Test.Column.Number.Expr(), false).As("aggregate_bitand"),
				BitOr(Test.Column.Number.Expr(), false).As("aggregate_bitor"),
				BitXor(Test.Column.Number.Expr(), false).As("aggregate_bitxor"),
				Count(Test.Column.String.Expr(), false).As("aggregate_count"),
				GroupConcat(Test.Column.String.Expr(), false).As("aggregate_groupconcat"),
				Max(Test.Column.Number.Expr(), false).As("aggregate_max"),
				Min(Test.Column.Number.Expr(), false).As("aggregate_min"),
				StdDev(Test.Column.Number.Expr(), false).As("aggregate_stddev"),
				Sum(Test.Column.Number.Expr(), false).As("aggregate_sum"),
				Variance(Test.Column.Number.Expr(), false).As("aggregate_variance"),
				// Функции аналитические
				FirstValue(Test.Column.Name.Expr()).Over(
					PartitionBy(Test.Column.ID.Expr()),
					OrderBy(Desc(Test.Column.Number.Expr())),
				).As("analytical_firstvalue"),
				Lag(Test.Column.Number.Expr(), 2).Over(
					PartitionBy(Test.Column.ID.Expr()),
					OrderBy(Asc(Test.Column.Date.Expr())),
				).As("analytical_lag"),
				LastValue(Test.Column.Name.Expr()).Over(
					PartitionBy(Test.Column.ID.Expr()),
					OrderBy(Asc(Test.Column.Number.Expr())),
					RowsBetween("CURRENT ROW", "UNBOUNDED FOLLOWING"),
				).As("analytical_lastvalue"),
				Lead(Test.Column.Number.Expr(), 2).Over(
					PartitionBy(Test.Column.ID.Expr()),
					OrderBy(Asc(Test.Column.Date.Expr())),
				).As("analytical_lead"),
				NthValue(Test.Column.Name.Expr(), 2).Over(
					PartitionBy(Test.Column.ID.Expr()),
					OrderBy(Desc(Test.Column.Number.Expr())),
					RowsBetween("UNBOUNDED PRECEDING", "CURRENT ROW"),
				).As("analytical_nthvalue"),
				// Функции условий
				Case(CaseIf(CasePair(Less(Test.Column.Number.Expr(), Value(2)), Value("old"))), CaseElse(Value("new"))).As("condition_case"),
				Coalesce(Test.Column.CreateAt.Expr(), Test.Column.UpdateAt.Expr()).As("condition_coalesce"),
				Greatest(Test.Column.CreateAt.Expr(), Test.Column.UpdateAt.Expr()).As("condition_greatest"),
				Least(Test.Column.CreateAt.Expr(), Test.Column.UpdateAt.Expr()).As("condition_least"),
				NullIf(Test.Column.CreateAt.Expr(), Test.Column.UpdateAt.Expr()).As("condition_if"),
				// Функции конвертации
				Cast(Test.Column.Number.Expr(), TypeString).As("convert_cast"),
				CharLength(Test.Column.String.Expr()).As("convert_charlength"),
				DateFormat(Test.Column.CreateAt.Expr(), Literal("%Y-%m-%d")).As("convert_dateformat"),
				Degrees(Test.Column.Number.Expr()).As("convert_degrees"),
				Length(Test.Column.String.Expr()).As("convert_length"),
				Position(Test.Column.String.Expr(), Value("old")).As("convert_position"),
				Radians(Test.Column.Number.Expr()).As("convert_radians"),
				// Функции даты и времени
				CurDate().As("datetime_curdate"),
				CurTime().As("datetime_curtime"),
				DateAdd(Test.Column.CreateAt.Expr(), Literal("2 DAY")).As("datetime_dateadd"),
				DateDiff(Test.Column.UpdateAt.Expr(), Test.Column.CreateAt.Expr()).As("datetime_datediff"),
				DateSub(Test.Column.CreateAt.Expr(), Literal("2 DAY")).As("datetime_datesub"),
				Day(Test.Column.CreateAt.Expr()).As("datetime_day"),
				DayName(Test.Column.CreateAt.Expr()).As("datetime_dayname"),
				Hour(Test.Column.CreateAt.Expr()).As("datetime_hour"),
				Minute(Test.Column.CreateAt.Expr()).As("datetime_minute"),
				Month(Test.Column.CreateAt.Expr()).As("datetime_month"),
				MonthName(Test.Column.CreateAt.Expr()).As("datetime_monthname"),
				Now().As("datetime_now"),
				Quarter(Test.Column.CreateAt.Expr()).As("datetime_quarter"),
				Second(Test.Column.CreateAt.Expr()).As("datetime_second"),
				TimeAdd(Test.Column.CreateAt.Expr(), Literal("2 HOUR")).As("datetime_timeadd"),
				TimeDiff(Test.Column.UpdateAt.Expr(), Test.Column.CreateAt.Expr()).As("datetime_timediff"),
				TimeSub(Test.Column.CreateAt.Expr(), Literal("2 HOUR")).As("datetime_timesub"),
				Week(Test.Column.CreateAt.Expr()).As("datetime_week"),
				Year(Test.Column.CreateAt.Expr()).As("datetime_year"),
				// Функции обмена данными
				JsonArray(Test.Column.Json.Expr(), Value("val1"), Value("val2")).As("json_jsonarray"),
				JsonArrayAgg(Test.Column.Json.Expr()).As("json_jsonarrayagg"),
				JsonContains(Test.Column.Json.Expr(), Value(`{"key":"val"}`)).As("json_jsoncontains"),
				JsonExtract(Test.Column.Json.Expr(), JsonGroup(JsonPath(JsonKey("parent"), JsonIndex(0), JsonKey("child"))), TypeString).As("json_jsonextract"),
				JsonObject(JsonPair(JsonKey("key"), Count(Test.Column.Json.Expr(), false))).As("json_jsonobject"),
				JsonObjectAgg(Test.Column.Json.Expr(), Test.Column.Number.Expr()).As("json_jsonobjectagg"),
				JsonRemove(Test.Column.Json.Expr(), JsonGroup(JsonPath(JsonKey("key1"))), JsonGroup(JsonPath(JsonKey("key2")))).As("json_jsonremove"),
				JsonSet(Test.Column.Json.Expr(), JsonGroup(JsonPath(JsonKey("key1")), Value("val1")), JsonGroup(JsonPath(JsonKey("key2")), Value("val2"))).As("json_jsonset"),
				JsonType(Test.Column.Json.Expr()).As("json_jsontype"),
				// Функции математические
				Abs(Test.Column.X.Expr()).As("math_abs"),
				ACos(Test.Column.X.Expr()).As("math_acos"),
				ASin(Test.Column.X.Expr()).As("math_asin"),
				ATan(Test.Column.X.Expr()).As("math_atan"),
				ATan2(Test.Column.Y.Expr(), Test.Column.X.Expr()).As("math_atan2"),
				Cbrt(Test.Column.X.Expr()).As("math_cbrt"),
				Ceil(Test.Column.X.Expr()).As("math_ceil"),
				Cos(Test.Column.X.Expr()).As("math_cos"),
				Exp(Test.Column.X.Expr()).As("math_exp"),
				Floor(Test.Column.X.Expr()).As("math_floor"),
				Ln(Test.Column.X.Expr()).As("math_ln"),
				Log(Test.Column.X.Expr(), Value(2)).As("math_log"),
				Mod(Test.Column.X.Expr(), Value(2)).As("math_mod"),
				Pi().As("math_pi"),
				Power(Test.Column.X.Expr(), Value(2)).As("math_power"),
				Rand().As("math_rand"),
				Round(Test.Column.X.Expr(), Value(2)).As("math_round"),
				Sin(Test.Column.X.Expr()).As("math_sin"),
				Sqrt(Test.Column.X.Expr()).As("math_sqrt"),
				Tan(Test.Column.X.Expr()).As("math_tan"),
				Trunc(Test.Column.X.Expr(), Value(2)).As("math_trunc"),
				// Функции ранжирующие
				CumeDist().Over(
					PartitionBy(Test.Column.ID.Expr()),
					OrderBy(Desc(Test.Column.Number.Expr())),
				).As("ranking_cumedist"),
				DenseRank().Over(
					PartitionBy(Test.Column.ID.Expr()),
					OrderBy(Desc(Test.Column.Number.Expr())),
				).As("ranking_denserank"),
				NTile(2).Over(
					PartitionBy(Test.Column.ID.Expr()),
					OrderBy(Desc(Test.Column.Number.Expr())),
				).As("ranking_ntile"),
				PercentRank().Over(
					PartitionBy(Test.Column.ID.Expr()),
					OrderBy(Desc(Test.Column.Number.Expr())),
				).As("ranking_percentrank"),
				Rank().Over(
					PartitionBy(Test.Column.ID.Expr()),
					OrderBy(Desc(Test.Column.Number.Expr())),
				).As("ranking_rank"),
				RowNumber().Over(
					PartitionBy(Test.Column.ID.Expr()),
					OrderBy(Desc(Test.Column.Number.Expr())),
				).As("ranking_rownumber"),
				// Функции строковые
				Concat(Test.Column.String.Expr(), Value("old"), Value("new")).As("string_concat"),
				ConcatWs(Value("_"), Test.Column.String.Expr(), Value("old"), Value("new")).As("string_concatws"),
				LeftString(Test.Column.String.Expr(), Value(2)).As("string_lstr"),
				Lower(Test.Column.String.Expr()).As("string_lower"),
				LPad(Test.Column.String.Expr(), Value(2), Value(",")).As("string_lpad"),
				LTrim(Test.Column.String.Expr()).As("string_ltrim"),
				Repeat(Test.Column.String.Expr(), Value(2)).As("string_repeat"),
				Replace(Test.Column.String.Expr(), Value("old"), Value("new")).As("string_replace"),
				Reverse(Test.Column.String.Expr()).As("string_reverse"),
				RightString(Test.Column.String.Expr(), Value(2)).As("string_rstr"),
				RPad(Test.Column.String.Expr(), Value(2), Value(",")).As("string_rpad"),
				RTrim(Test.Column.String.Expr()).As("string_rtrim"),
				SubString(Test.Column.String.Expr(), Value(0), Value(2)).As("string_substring"),
				Trim(Test.Column.String.Expr()).As("string_trim"),
				Upper(Test.Column.String.Expr()).As("string_upper"),
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
			assertContains(t, sqlSelectQuery, "ABS(`t`.`x`)", "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, "ACOS(`t`.`x`)", "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, "ASIN(`t`.`x`)", "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, "ATAN(`t`.`x`)", "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, "ATAN2(`t`.`y`, `t`.`x`)", "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, "CBRT(`t`.`x`)", "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, "CEILING(`t`.`x`)", "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, "COS(`t`.`x`)", "FUNCTION COS")
			assertContains(t, sqlSelectQuery, "EXP(`t`.`x`)", "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, "FLOOR(`t`.`x`)", "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, "LN(`t`.`x`)", "FUNCTION LN")
			assertContains(t, sqlSelectQuery, "LOG(`t`.`x`, ?)", "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, "MOD(`t`.`x`, ?)", "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, "PI()", "FUNCTION PI")
			assertContains(t, sqlSelectQuery, "POWER(`t`.`x`, ?)", "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, "RAND()", "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, "ROUND(`t`.`x`, ?)", "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, "SIN(`t`.`x`)", "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, "SQRT(`t`.`x`)", "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, "TAN(`t`.`x`)", "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, "TRUNCATE(`t`.`x`, ?)", "FUNCTION TRUNC")
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
			assertContains(t, sqlSelectQuery, "CAST(GETDATE() AS DATE)", "FUNCTION CURDATE")
			assertContains(t, sqlSelectQuery, "CAST(GETDATE() AS TIME)", "FUNCTION CURTIME")
			assertContains(t, sqlSelectQuery, "DATEADD(DAY, 2, [t].[createat])", "FUNCTION DATEADD")
			assertContains(t, sqlSelectQuery, "DATEDIFF([t].[updateat], [t].[createat])", "FUNCTION DATEDIFF")
			assertContains(t, sqlSelectQuery, "DATEADD(DAY, -2, [t].[createat])", "FUNCTION DATESUB")
			assertContains(t, sqlSelectQuery, "DAY([t].[createat])", "FUNCTION DAY")
			assertContains(t, sqlSelectQuery, "DATENAME(WEEKDAY, [t].[createat])", "FUNCTION DAYNAME")
			assertContains(t, sqlSelectQuery, "DATEPART(HOUR, [t].[createat])", "FUNCTION HOUR")
			assertContains(t, sqlSelectQuery, "DATEPART(MINUTE, [t].[createat])", "FUNCTION MINUTE")
			assertContains(t, sqlSelectQuery, "MONTH([t].[createat])", "FUNCTION MONTH")
			assertContains(t, sqlSelectQuery, "DATENAME(MONTH, [t].[createat])", "FUNCTION MONTHNAME")
			assertContains(t, sqlSelectQuery, "GETDATE()", "FUNCTION NOW")
			assertContains(t, sqlSelectQuery, "DATEPART(QUARTER, [t].[createat])", "FUNCTION QUARTER")
			assertContains(t, sqlSelectQuery, "DATEPART(SECOND, [t].[createat])", "FUNCTION SECOND")
			assertContains(t, sqlSelectQuery, "DATEADD(HOUR, 2, [t].[createat])", "FUNCTION TIMEADD")
			assertContains(t, sqlSelectQuery, "TIMEDIFF([t].[updateat], [t].[createat])", "FUNCTION TIMEDIFF")
			assertContains(t, sqlSelectQuery, "DATEADD(HOUR, -2, [t].[createat])", "FUNCTION TIMESUB")
			assertContains(t, sqlSelectQuery, "DATEPART(WEEK, [t].[createat])", "FUNCTION WEEK")
			assertContains(t, sqlSelectQuery, "YEAR([t].[createat])", "FUNCTION YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, "JSON_ARRAY([t].[json], @p1, @p2)", "FUNCTION JSONARRAY")
			assertContains(t, sqlSelectQuery, "JSON_ARRAYAGG([t].[json])", "FUNCTION JSONARRAYAGG")
			// Not supported - FUNCTION JSONCONTAINS
			assertContains(t, sqlSelectQuery, "JSON_VALUE([t].[json], '$.parent[0].child')", "FUNCTION JSONEXTRACT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECT('key', COUNT([t].[json]))", "FUNCTION JSONOBJECT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECTAGG([t].[json], [t].[number])", "FUNCTION JSONOBJECTAGG")
			assertContains(t, sqlSelectQuery, "JSON_MODIFY(JSON_MODIFY([t].[json], '$.key1', NULL), '$.key2', NULL)", "FUNCTION JSONREMOVE")
			assertContains(t, sqlSelectQuery, "JSON_MODIFY(JSON_MODIFY([t].[json], '$.key1', @p1), '$.key2', @p2)", "FUNCTION JSONSET")
			// Not supported - FUNCTION JSONTYPE
			// Функции математические
			assertContains(t, sqlSelectQuery, "ABS([t].[x])", "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, "ACOS([t].[x])", "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, "ASIN([t].[x])", "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, "ATAN([t].[x])", "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, "ATAN2([t].[y], [t].[x])", "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, "CBRT([t].[x])", "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, "CEILING([t].[x])", "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, "COS([t].[x])", "FUNCTION COS")
			assertContains(t, sqlSelectQuery, "EXP([t].[x])", "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, "FLOOR([t].[x])", "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, "LN([t].[x])", "FUNCTION LN")
			assertContains(t, sqlSelectQuery, "LOG([t].[x], @p1)", "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, "MOD([t].[x], @p1)", "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, "PI()", "FUNCTION PI")
			assertContains(t, sqlSelectQuery, "POWER([t].[x], @p1)", "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, "RAND()", "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, "ROUND([t].[x], @p1)", "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, "SIN([t].[x])", "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, "SQRT([t].[x])", "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, "TAN([t].[x])", "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, "ROUND([t].[x], @p1, 1)", "FUNCTION TRUNC")
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
			assertContains(t, sqlSelectQuery, "ABS(`t`.`x`)", "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, "ACOS(`t`.`x`)", "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, "ASIN(`t`.`x`)", "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, "ATAN(`t`.`x`)", "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, "ATAN2(`t`.`y`, `t`.`x`)", "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, "CBRT(`t`.`x`)", "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, "CEILING(`t`.`x`)", "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, "COS(`t`.`x`)", "FUNCTION COS")
			assertContains(t, sqlSelectQuery, "EXP(`t`.`x`)", "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, "FLOOR(`t`.`x`)", "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, "LN(`t`.`x`)", "FUNCTION LN")
			assertContains(t, sqlSelectQuery, "LOG(`t`.`x`, ?)", "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, "MOD(`t`.`x`, ?)", "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, "PI()", "FUNCTION PI")
			assertContains(t, sqlSelectQuery, "POWER(`t`.`x`, ?)", "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, "RAND()", "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, "ROUND(`t`.`x`, ?)", "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, "SIN(`t`.`x`)", "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, "SQRT(`t`.`x`)", "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, "TAN(`t`.`x`)", "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, "TRUNCATE(`t`.`x`, ?)", "FUNCTION TRUNC")
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
			assertContains(t, sqlSelectQuery, `jsonb_set(jsonb_set("t"."json", '{key1}', $1), '{key2}', $2)`, "FUNCTION JSONSET")
			assertContains(t, sqlSelectQuery, `jsonb_typeof("t"."json")`, "FUNCTION JSONTYPE")
			// Функции математические
			assertContains(t, sqlSelectQuery, `ABS("t"."x")`, "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, `ACOS("t"."x")`, "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, `ASIN("t"."x")`, "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, `ATAN("t"."x")`, "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, `ATAN2("t"."y", "t"."x")`, "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, `CBRT("t"."x")`, "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, `CEIL("t"."x")`, "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, `COS("t"."x")`, "FUNCTION COS")
			assertContains(t, sqlSelectQuery, `EXP("t"."x")`, "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, `FLOOR("t"."x")`, "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, `LN("t"."x")`, "FUNCTION LN")
			assertContains(t, sqlSelectQuery, `LOG("t"."x", $1)`, "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, `MOD("t"."x", $1)`, "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, `PI()`, "FUNCTION PI")
			assertContains(t, sqlSelectQuery, `POWER("t"."x", $1)`, "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, `RANDOM`, "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, `ROUND("t"."x", $1)`, "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, `SIN("t"."x")`, "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, `SQRT("t"."x")`, "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, `TAN("t"."x")`, "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, `TRUNC("t"."x", $1)`, "FUNCTION TRUNC")
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
			assertContains(t, sqlSelectQuery, `STRFTIME("t"."createat", '%Y-%m-%d')`, "FUNCTION DATEFORMAT")
			assertContains(t, sqlSelectQuery, `DEGREES("t"."number")`, "FUNCTION DEGREES")
			assertContains(t, sqlSelectQuery, `LENGTH("t"."string")`, "FUNCTION LENGTH")
			assertContains(t, sqlSelectQuery, `POSITION(? IN "t"."string")`, "FUNCTION POSITION")
			assertContains(t, sqlSelectQuery, `RADIANS("t"."number")`, "FUNCTION RADIANS")
			// Функции даты и времени
			assertContains(t, sqlSelectQuery, `DATE('now')`, "FUNCTION CURDATE")
			assertContains(t, sqlSelectQuery, `TIME('now')`, "FUNCTION CURTIME")
			assertContains(t, sqlSelectQuery, `DATETIME("t"."createat", '+2 DAY')`, "FUNCTION DATEADD")
			assertContains(t, sqlSelectQuery, `DATEDIFF("t"."updateat", "t"."createat")`, "FUNCTION DATEDIFF")
			assertContains(t, sqlSelectQuery, `DATETIME("t"."createat", '-2 DAY')`, "FUNCTION DATESUB")
			assertContains(t, sqlSelectQuery, `DAY("t"."createat")`, "FUNCTION DAY")
			assertContains(t, sqlSelectQuery, `STRFTIME('%w', "t"."createat")`, "FUNCTION DAYNAME")
			assertContains(t, sqlSelectQuery, `HOUR("t"."createat")`, "FUNCTION HOUR")
			assertContains(t, sqlSelectQuery, `MINUTE("t"."createat")`, "FUNCTION MINUTE")
			assertContains(t, sqlSelectQuery, `MONTH("t"."createat")`, "FUNCTION MONTH")
			assertContains(t, sqlSelectQuery, `STRFTIME('%m', "t"."createat")`, "FUNCTION MONTHNAME")
			assertContains(t, sqlSelectQuery, `DATETIME('now')`, "FUNCTION NOW")
			assertContains(t, sqlSelectQuery, `QUARTER("t"."createat")`, "FUNCTION QUARTER")
			assertContains(t, sqlSelectQuery, `SECOND("t"."createat")`, "FUNCTION SECOND")
			assertContains(t, sqlSelectQuery, `TIME("t"."createat", '+2 HOUR')`, "FUNCTION TIMEADD")
			assertContains(t, sqlSelectQuery, `TIMEDIFF("t"."updateat", "t"."createat")`, "FUNCTION TIMEDIFF")
			assertContains(t, sqlSelectQuery, `TIME("t"."createat", '-2 HOUR')`, "FUNCTION TIMESUB")
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
			assertContains(t, sqlSelectQuery, `ABS("t"."x")`, "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, `ACOS("t"."x")`, "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, `ASIN("t"."x")`, "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, `ATAN("t"."x")`, "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, `ATAN2("t"."y", "t"."x")`, "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, `CBRT("t"."x")`, "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, `CEIL("t"."x")`, "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, `COS("t"."x")`, "FUNCTION COS")
			assertContains(t, sqlSelectQuery, `EXP("t"."x")`, "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, `FLOOR("t"."x")`, "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, `LN("t"."x")`, "FUNCTION LN")
			assertContains(t, sqlSelectQuery, `LOG("t"."x", ?)`, "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, `MOD("t"."x", ?)`, "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, `PI()`, "FUNCTION PI")
			assertContains(t, sqlSelectQuery, `POWER("t"."x", ?)`, "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, `RANDOM`, "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, `ROUND("t"."x", ?)`, "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, `SIN("t"."x")`, "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, `SQRT("t"."x")`, "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, `TAN("t"."x")`, "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, `TRUNC("t"."x", ?)`, "FUNCTION TRUNC")
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
			Fields(
				Test.Column.ID.Expr(),
			).
			Where(
				Equal(DateFormat(Test.Column.CreateAt.Expr(), Literal("%Y-%m-%d")), Value("2026-01-01")),
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
			Fields(
				Test.Column.ID.Expr(),
			).
			Where(
				And(
					And(
						Equal(Test.Column.String.Expr(), Value("active")),
						Greater(Test.Column.Number.Expr(), Value(2)),
					),
					Or(
						Equal(Test.Column.String.Expr(), Value("active")),
						Greater(Test.Column.Number.Expr(), Value(2)),
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
			Fields(
				Subquery[int64](NewSelect(Test.Table).Fields(Test.Column.ID.Expr())).As("SUB"),
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
			Fields(
				Test.Column.ID.Expr(),
			).
			Where(
				Equal(Test.Column.String.Expr(), Value(data)),
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
					Inner(Data.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				).
				Where(
					And(
						Equal(Test.Column.String.Expr(), Value("active")),
						ILike(Test.Column.String.Expr(), Value("%ivan%")),
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
					Pair(Test.Column.String.Expr(), Value("ivan")),
					Pair(Test.Column.Number.Expr(), Value(2)),
				).
				Upsert(
					Pair(Test.Column.String.Expr(), Value("updated")),
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
				Fields(
					Avg(Test.Column.Number.Expr(), false).As("avg_result"),
					Ceil(Test.Column.Number.Expr()).As("ceil_result"),
					Count(Test.Column.String.Expr(), false).As("count_result"),
					FirstValue(Test.Column.Name.Expr()).Over(
						PartitionBy(Test.Column.ID.Expr()),
						OrderBy(Desc(Test.Column.Number.Expr())),
					).As("first_value"),
					Trunc(Test.Column.Number.Expr(), Value(2)).As("trunc_result"),
				).
				Join(
					Inner(Data.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				).
				Where(
					And(
						Equal(Test.Column.String.Expr(), Value("active")),
						Greater(Test.Column.Number.Expr(), Value(2)),
						ILike(Test.Column.String.Expr(), Value("%ivan%")),
					),
				).
				GroupBy(
					Test.Column.ID.Expr(),
					Test.Column.String.Expr(),
				).
				Having(
					Greater(Count(Test.Column.ID.Expr(), false), Value[int64](2)),
				).
				OrderBy(
					Desc(Test.Column.Number.Expr()),
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
					Assign(Test.Column.String.Expr(), Value("updated")),
					Assign(Test.Column.Number.Expr(), Value(2)),
				).
				Join(
					Inner(Data.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				).
				Where(
					And(
						Equal(Test.Column.String.Expr(), Value("active")),
						ILike(Test.Column.String.Expr(), Value("%ivan%")),
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
	t.Run("Column", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtComment := NewComment(Test.Column.ID).
				Is("Test comment")
			sqlCommentQuery, sqlCommentArguments, err := sql.Build(stmtComment)
			switch supportDialect {
			case DialectMariaDB:
				//assertContains(t, sqlCommentQuery, "COMMENT ON COLUMN `test`.`id` IS 'Test comment'", "COMMENT")
			case DialectMsSQL:
				// Not supported - COMMENT
			case DialectMySQL:
				//assertContains(t, sqlCommentQuery, "COMMENT ON COLUMN `test`.`id` IS 'Test comment'", "COMMENT")
			case DialectPostgreSQL:
				//assertContains(t, sqlCommentQuery, `COMMENT ON COLUMN "test"."id" IS 'Test comment'`, "COMMENT")
			case DialectSQLite:
				//assertContains(t, sqlCommentQuery, `COMMENT ON COLUMN "test"."id" IS 'Test comment'`, "COMMENT")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCommentArguments, supportDialect.name, sqlCommentQuery)
		})
	})
	t.Run("Table", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtComment := NewComment(Test.Table).Is("Test comment")
			sqlCommentQuery, sqlCommentArguments, err := sql.Build(stmtComment)
			switch supportDialect {
			case DialectMariaDB:
				//assertContains(t, sqlCommentQuery, "COMMENT ON TABLE `test` IS 'Test comment'", "COMMENT")
			case DialectMsSQL:
				// Not supported - COMMENT
			case DialectMySQL:
				//assertContains(t, sqlCommentQuery, "COMMENT ON TABLE `test` IS 'Test comment'", "COMMENT")
			case DialectPostgreSQL:
				//assertContains(t, sqlCommentQuery, `COMMENT ON TABLE "test" IS 'Test comment'`, "COMMENT")
			case DialectSQLite:
				//assertContains(t, sqlCommentQuery, `COMMENT ON TABLE "test" IS 'Test comment'`, "COMMENT")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCommentArguments, supportDialect.name, sqlCommentQuery)
		})
	})
}
func Test_SQL_Create(t *testing.T) {
	t.Run("Index", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtCreate := NewCreate(NewIndex("idx_test", Test.Table)).
				IfNotExists().
				IsUnique().
				On(Test.Table).
				Columns(
					Test.Column.String,
					Test.Column.Number,
				)
			sqlCreateQuery, sqlCreateArguments, err := sql.Build(stmtCreate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlCreateQuery, "CREATE UNIQUE INDEX IF NOT EXISTS `idx_test` ON `test` (`string`, `number`)", "CREATE INDEX")
			case DialectMsSQL:
				assertContains(t, sqlCreateQuery, "CREATE UNIQUE INDEX [idx_test] ON [test] ([string], [number])", "CREATE INDEX")
			case DialectMySQL:
				assertContains(t, sqlCreateQuery, "CREATE UNIQUE INDEX `idx_test` ON `test` (`string`, `number`)", "CREATE INDEX")
			case DialectPostgreSQL:
				assertContains(t, sqlCreateQuery, `CREATE UNIQUE INDEX IF NOT EXISTS "idx_test" ON "test" ("string", "number")`, "CREATE INDEX")
			case DialectSQLite:
				assertContains(t, sqlCreateQuery, `CREATE UNIQUE INDEX IF NOT EXISTS "idx_test" ON "test" ("string", "number")`, "CREATE INDEX")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCreateArguments, supportDialect.name, sqlCreateQuery)
		})
	})
	t.Run("Schema", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtCreate := NewCreate(NewSchema("public")).
				IfNotExists()
			sqlCreateQuery, sqlCreateArguments, err := sql.Build(stmtCreate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlCreateQuery, "CREATE SCHEMA IF NOT EXISTS `public`", "CREATE SCHEMA")
			case DialectMsSQL:
				assertContains(t, sqlCreateQuery, "CREATE SCHEMA IF NOT EXISTS [public]", "CREATE SCHEMA")
			case DialectMySQL:
				assertContains(t, sqlCreateQuery, "CREATE SCHEMA `public`", "CREATE SCHEMA")
			case DialectPostgreSQL:
				assertContains(t, sqlCreateQuery, `CREATE SCHEMA IF NOT EXISTS "public"`, "CREATE SCHEMA")
			case DialectSQLite:
				// Not supported
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCreateArguments, supportDialect.name, sqlCreateQuery)
		})
	})
	t.Run("Table", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtCreate := NewCreate(Test.Table).
				Constraints(
					NewCheck("ck_id", Greater(Test.Column.ID.Expr(), Value[int64](0))),
					NewForeignKey("fk_test_data", Data.Table, Cascade(), Restrict(),
						Relation(Test.Column.DataID, Data.Column.ID),
						Relation(Test.Column.Name, Data.Column.String),
					),
					NewPrimaryKey("pk_orders", Test.Column.ID),
					NewUnique("uk_orders", Test.Column.Name),
				).
				IfNotExists().
				Columns(
					Test.Column.ID.AutoIncrement(),
					Test.Column.Name.NotNull(),
				)
			sqlCreateQuery, sqlCreateArguments, err := sql.Build(stmtCreate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlCreateQuery, "CREATE TABLE IF NOT EXISTS `test` (`id` SIGNED AUTO_INCREMENT, `name` VARCHAR NOT NULL, CONSTRAINT `ck_id` CHECK(`t`.`id` > ?), CONSTRAINT `fk_test_data` FOREIGN KEY(`data_id`, `name`) REFERENCES `data`(`id`, `string`) ON DELETE CASCADE ON UPDATE RESTRICT, CONSTRAINT `pk_orders` PRIMARY KEY(`id`), CONSTRAINT `uk_orders` UNIQUE(`name`))", "CREATE TABLE")
			case DialectMsSQL:
				assertContains(t, sqlCreateQuery, "CREATE TABLE [test] ([id] BIGINT IDENTITY(1,1), [name] NVARCHAR NOT NULL, CONSTRAINT [ck_id] CHECK([t].[id] > @p1), CONSTRAINT [fk_test_data] FOREIGN KEY([data_id], [name]) REFERENCES [data]([id], [string]) ON DELETE CASCADE ON UPDATE RESTRICT, CONSTRAINT [pk_orders] PRIMARY KEY([id]), CONSTRAINT [uk_orders] UNIQUE([name]))", "CREATE TABLE")
			case DialectMySQL:
				assertContains(t, sqlCreateQuery, "CREATE TABLE IF NOT EXISTS `test` (`id` SIGNED AUTO_INCREMENT, `name` VARCHAR NOT NULL, CONSTRAINT `ck_id` CHECK(`t`.`id` > ?), CONSTRAINT `fk_test_data` FOREIGN KEY(`data_id`, `name`) REFERENCES `data`(`id`, `string`) ON DELETE CASCADE ON UPDATE RESTRICT, CONSTRAINT `pk_orders` PRIMARY KEY(`id`), CONSTRAINT `uk_orders` UNIQUE(`name`))", "CREATE TABLE")
			case DialectPostgreSQL:
				assertContains(t, sqlCreateQuery, `CREATE TABLE IF NOT EXISTS "test" ("id" BIGINT GENERATED BY DEFAULT AS IDENTITY, "name" VARCHAR NOT NULL, CONSTRAINT "ck_id" CHECK("t"."id" > $1), CONSTRAINT "fk_test_data" FOREIGN KEY("data_id", "name") REFERENCES "data"("id", "string") ON DELETE CASCADE ON UPDATE RESTRICT, CONSTRAINT "pk_orders" PRIMARY KEY("id"), CONSTRAINT "uk_orders" UNIQUE("name"))`, "CREATE TABLE")
			case DialectSQLite:
				assertContains(t, sqlCreateQuery, `CREATE TABLE IF NOT EXISTS "test" ("id" INTEGER AUTOINCREMENT, "name" TEXT NOT NULL, CONSTRAINT "ck_id" CHECK("t"."id" > ?), CONSTRAINT "fk_test_data" FOREIGN KEY("data_id", "name") REFERENCES "data"("id", "string") ON DELETE CASCADE ON UPDATE RESTRICT, CONSTRAINT "pk_orders" PRIMARY KEY("id"), CONSTRAINT "uk_orders" UNIQUE("name"))`, "CREATE TABLE")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCreateArguments, supportDialect.name, sqlCreateQuery)
		})
	})
	t.Run("View", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtCreate := NewCreate(NewView("test_view", "tv", Test.Table)).
				IsReplace().
				Source(NewSelect(NewTable("test", "t")).
					Fields(Test.Column.ID.Expr(), Test.Column.String.Expr()),
				)
			sqlCreateQuery, sqlCreateArguments, err := sql.Build(stmtCreate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlCreateQuery, "CREATE OR REPLACE VIEW `test_view` AS SELECT `t`.`id`, `t`.`string` FROM `test` AS `t`", "CREATE VIEW")
			case DialectMsSQL:
				assertContains(t, sqlCreateQuery, "CREATE OR REPLACE VIEW [test_view] AS SELECT [t].[id], [t].[string] FROM [test] AS [t]", "CREATE VIEW")
			case DialectMySQL:
				assertContains(t, sqlCreateQuery, "CREATE OR REPLACE VIEW `test_view` AS SELECT `t`.`id`, `t`.`string` FROM `test` AS `t`", "CREATE VIEW")
			case DialectPostgreSQL:
				assertContains(t, sqlCreateQuery, `CREATE OR REPLACE VIEW "test_view" AS SELECT "t"."id", "t"."string" FROM "test" AS "t"`, "CREATE VIEW")
			case DialectSQLite:
				assertContains(t, sqlCreateQuery, `CREATE OR REPLACE VIEW "test_view" AS SELECT "t"."id", "t"."string" FROM "test" AS "t"`, "CREATE VIEW")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCreateArguments, supportDialect.name, sqlCreateQuery)
		})
	})
}
func Test_SQL_Delete(t *testing.T) {
	t.Run("Join", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtDelete := NewDelete(Test.Table).
				Join(
					Inner(Data.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				).
				Where(
					Equal(Test.Column.String.Expr(), Value("active")),
				)
			sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDeleteQuery, "DELETE `t` FROM `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id` WHERE `t`.`string` = ?", "DELETE JOIN")
			case DialectMsSQL:
				assertContains(t, sqlDeleteQuery, "DELETE [t] FROM [test] AS [t] INNER JOIN [data] AS [d] ON [t].[id] = [d].[id] WHERE [t].[string] = @p1", "DELETE JOIN")
			case DialectMySQL:
				assertContains(t, sqlDeleteQuery, "DELETE `t` FROM `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id` WHERE `t`.`string` = ?", "DELETE JOIN")
			case DialectPostgreSQL:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "test" AS "t" USING "data" AS "d" WHERE ("t"."id" = "d"."id" AND "t"."string" = $1)`, "DELETE JOIN")
			case DialectSQLite:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "test" AS "t" INNER JOIN "data" AS "d" ON "t"."id" = "d"."id" WHERE "t"."string" = ?`, "DELETE JOIN")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
		})
	})
	t.Run("Returning", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtDelete := NewDelete(Test.Table).
				Where(
					Equal(Test.Column.String.Expr(), Value("active")),
				).
				Returning(
					Test.Column.ID.Expr(),
					Test.Column.String.Expr(),
				)
			sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDeleteQuery, "DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ? RETURNING `t`.`id`, `t`.`string`", "DELETE RETURNING")
			case DialectMsSQL:
				assertContains(t, sqlDeleteQuery, "DELETE [t] FROM [test] AS [t] OUTPUT [t].[id], [t].[string] WHERE [t].[string] = @p1", "DELETE RETURNING")
			case DialectMySQL:
				assertContains(t, sqlDeleteQuery, "DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?", "DELETE WITHOUT RETURNING")
			case DialectPostgreSQL:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "test" AS "t" WHERE "t"."string" = $1 RETURNING "t"."id", "t"."string"`, "DELETE RETURNING")
			case DialectSQLite:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "test" AS "t" WHERE "t"."string" = ? RETURNING "t"."id", "t"."string"`, "DELETE RETURNING")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
		})
	})
	t.Run("Where", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtDelete := NewDelete(Test.Table).
				Where(
					Equal(Test.Column.String.Expr(), Value("active")),
				)
			sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDeleteQuery, "DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?", "DELETE")
			case DialectMsSQL:
				assertContains(t, sqlDeleteQuery, "DELETE [t] FROM [test] AS [t] WHERE [t].[string] = @p1", "DELETE")
			case DialectMySQL:
				assertContains(t, sqlDeleteQuery, "DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?", "DELETE")
			case DialectPostgreSQL:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "test" AS "t" WHERE "t"."string" = $1`, "DELETE")
			case DialectSQLite:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "test" AS "t" WHERE "t"."string" = ?`, "DELETE")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
		})
	})
	t.Run("With", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtDelete := NewDelete(Test.Table).
				With(
					WithN("old_users", NewSelect(Test.Table).
						Fields(
							Test.Column.ID.Expr(),
						).
						Where(
							Less(Test.Column.Number.Expr(), Value(2)),
						),
					),
				)
			sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDeleteQuery, "WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) DELETE `t` FROM `test` AS `t`", "WITH")
			case DialectMsSQL:
				assertContains(t, sqlDeleteQuery, "WITH [old_users] AS (SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] < @p1) DELETE [t] FROM [test] AS [t]", "WITH")
			case DialectMySQL:
				assertContains(t, sqlDeleteQuery, "WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) DELETE `t` FROM `test` AS `t`", "WITH")
			case DialectPostgreSQL:
				assertContains(t, sqlDeleteQuery, `WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < $1) DELETE FROM "test" AS "t"`, "WITH")
			case DialectSQLite:
				assertContains(t, sqlDeleteQuery, `WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < ?) DELETE FROM "test" AS "t"`, "WITH")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
		})
	})
}
func Test_SQL_Drop(t *testing.T) {
	t.Run("Cascade", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtDrop := NewDrop(Test.Table).IsCascade()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP TABLE `test` CASCADE", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP TABLE [test]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP TABLE `test`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP TABLE "test" CASCADE`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP TABLE "test"`, "DROP")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDropArguments, supportDialect.name, sqlDropQuery)
		})
	})
	t.Run("IfExistsIndex", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtDrop := NewDrop(NewIndex("test", Test.Table)).IfExists()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP INDEX IF EXISTS `test`", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP INDEX [test]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP INDEX `test`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP INDEX IF EXISTS "test"`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP INDEX IF EXISTS "test"`, "DROP")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDropArguments, supportDialect.name, sqlDropQuery)
		})
	})
	t.Run("IfExistsSchema", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtDrop := NewDrop(NewSchema("test")).IfExists()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP SCHEMA IF EXISTS `test`", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP SCHEMA IF EXISTS [test]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP SCHEMA `test`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP SCHEMA IF EXISTS "test"`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP SCHEMA "test"`, "DROP")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDropArguments, supportDialect.name, sqlDropQuery)
		})
	})
	t.Run("IfExistsTable", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtDrop := NewDrop(Test.Table).IfExists()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP TABLE IF EXISTS `test`", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP TABLE IF EXISTS [test]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP TABLE IF EXISTS `test`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP TABLE IF EXISTS "test"`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP TABLE IF EXISTS "test"`, "DROP")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDropArguments, supportDialect.name, sqlDropQuery)
		})
	})
	t.Run("IfExistsView", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtDrop := NewDrop(NewView("test", "t", Test.Table)).IfExists()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP VIEW IF EXISTS `test`", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP VIEW IF EXISTS [test]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP VIEW IF EXISTS `test`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP VIEW IF EXISTS "test"`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP VIEW IF EXISTS "test"`, "DROP")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDropArguments, supportDialect.name, sqlDropQuery)
		})
	})
}
func Test_SQL_Insert(t *testing.T) {
	t.Run("Returning", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtInsert := NewInsert(Test.Table).
				Values(
					Pair(Test.Column.String.Expr(), Value("ivan")),
					Pair(Test.Column.Number.Expr(), Value(2)),
				).
				Returning(
					Test.Column.ID.Expr(),
					Test.Column.String.Expr(),
				)
			sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlInsertQuery, "INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?) RETURNING `t`.`id`, `t`.`string`", "INSERT RETURNING")
			case DialectMsSQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO [test] AS [t] ([string], [number]) OUTPUT [t].[id], [t].[string] VALUES (@p1, @p2)", "INSERT RETURNING")
			case DialectMySQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)", "INSERT WITHOUT RETURNING")
			case DialectPostgreSQL:
				assertContains(t, sqlInsertQuery, `INSERT INTO "test" AS "t" ("string", "number") VALUES ($1, $2) RETURNING "t"."id", "t"."string"`, "INSERT RETURNING")
			case DialectSQLite:
				assertContains(t, sqlInsertQuery, `INSERT INTO "test" AS "t" ("string", "number") VALUES (?, ?) RETURNING "t"."id", "t"."string"`, "INSERT RETURNING")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
		})
	})
	t.Run("Source", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtInsert := NewInsert(Test.Table).
				Source(NewSelect(Test.Table).
					Fields(
						Test.Column.String.Expr(),
						Test.Column.Number.Expr(),
					).
					Where(
						Equal(Test.Column.String.Expr(), Value("active")),
					),
				)
			sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlInsertQuery, "INSERT INTO `test` AS `t` (`string`, `number`) SELECT `t`.`string`, `t`.`number` FROM `test` AS `t` WHERE `t`.`string` = ?", "INSERT SOURCE")
			case DialectMsSQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO [test] AS [t] ([string], [number]) SELECT [t].[string], [t].[number] FROM [test] AS [t] WHERE [t].[string] = @p1", "INSERT SOURCE")
			case DialectMySQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO `test` AS `t` (`string`, `number`) SELECT `t`.`string`, `t`.`number` FROM `test` AS `t` WHERE `t`.`string` = ?", "INSERT SOURCE")
			case DialectPostgreSQL:
				assertContains(t, sqlInsertQuery, `INSERT INTO "test" AS "t" ("string", "number") SELECT "t"."string", "t"."number" FROM "test" AS "t" WHERE "t"."string" = $1`, "INSERT SOURCE")
			case DialectSQLite:
				assertContains(t, sqlInsertQuery, `INSERT INTO "test" AS "t" ("string", "number") SELECT "t"."string", "t"."number" FROM "test" AS "t" WHERE "t"."string" = ?`, "INSERT SOURCE")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
		})
	})
	t.Run("Values", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtInsert := NewInsert(Test.Table).
				Values(
					Pair(Test.Column.String.Expr(), Value("ivan")),
					Pair(Test.Column.Number.Expr(), Value(2)),
				).
				Upsert(
					Pair(Test.Column.String.Expr(), Value("updated")),
				)
			sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlInsertQuery, "INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?", "INSERT VALUES WITH UPSERT")
			case DialectMsSQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO [test] AS [t] ([string], [number]) VALUES (@p1, @p2)", "INSERT VALUES WITHOUT UPSERT")
			case DialectMySQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?", "INSERT VALUES WITH UPSERT")
			case DialectPostgreSQL:
				assertContains(t, sqlInsertQuery, `INSERT INTO "test" AS "t" ("string", "number") VALUES ($1, $2) ON CONFLICT DO UPDATE SET "string" = $3`, "INSERT UVALUES WITH PSERT")
			case DialectSQLite:
				assertContains(t, sqlInsertQuery, `INSERT INTO "test" AS "t" ("string", "number") VALUES (?, ?) ON CONFLICT DO UPDATE SET "string" = ?`, "INSERT VALUES WITH UPSERT")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
		})
	})
	t.Run("With", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtInsert := NewInsert(Test.Table).
				Values(
					Pair(Test.Column.String.Expr(), Value("ivan")),
					Pair(Test.Column.Number.Expr(), Value(2)),
				).
				With(
					WithN("old_users", NewSelect(Test.Table).
						Fields(
							Test.Column.ID.Expr(),
						).
						Where(
							Less(Test.Column.Number.Expr(), Value(2)),
						),
					),
				)
			sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlInsertQuery, "WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)", "INSERT WITH")
			case DialectMsSQL:
				assertContains(t, sqlInsertQuery, "WITH [old_users] AS (SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] < @p1) INSERT INTO [test] AS [t] ([string], [number]) VALUES (@p2, @p3)", "INSERT WITH")
			case DialectMySQL:
				assertContains(t, sqlInsertQuery, "WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)", "INSERT WITH")
			case DialectPostgreSQL:
				assertContains(t, sqlInsertQuery, `WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < $1) INSERT INTO "test" AS "t" ("string", "number") VALUES ($2, $3)`, "INSERT WITH")
			case DialectSQLite:
				assertContains(t, sqlInsertQuery, `WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < ?) INSERT INTO "test" AS "t" ("string", "number") VALUES (?, ?)`, "INSERT WITH")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
		})
	})
}
func Test_SQL_Select(t *testing.T) {
	t.Run("Distinct", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table).
				Distinct().
				Fields(
					Test.Column.ID.Expr(),
				).
				Where(
					Equal(Test.Column.Number.Expr(), Value(2)),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT DISTINCT `t`.`id` FROM `test` AS `t`", "SELECT DISTINCT")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT DISTINCT [t].[id] FROM [test] AS [t]", "SELECT DISTINCT")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT DISTINCT `t`.`id` FROM `test` AS `t`", "SELECT DISTINCT")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT DISTINCT "t"."id" FROM "test" AS "t"`, "SELECT DISTINCT")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT DISTINCT "t"."id" FROM "test" AS "t"`, "SELECT DISTINCT")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("Fields", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table).
				Fields(
					Test.Column.ID.Expr(),
				).
				Where(
					Equal(Test.Column.Number.Expr(), Value(2)),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?", "SELECT")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] = @p1", "SELECT")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?", "SELECT")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = $1`, "SELECT")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = ?`, "SELECT")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("GroupBy", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table).
				Fields(
					Test.Column.String.Expr(),
					Count(Test.Column.ID.Expr(), false).As("cnt"),
				).
				GroupBy(
					Test.Column.String.Expr(),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`string`, COUNT(`t`.`id`) AS `cnt` FROM `test` AS `t` GROUP BY `t`.`string`", "SELECT GROUP BY")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [t].[string], COUNT([t].[id]) AS [cnt] FROM [test] AS [t] GROUP BY [t].[string]", "SELECT GROUP BY")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`string`, COUNT(`t`.`id`) AS `cnt` FROM `test` AS `t` GROUP BY `t`.`string`", "SELECT GROUP BY")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "t"."string", COUNT("t"."id") AS "cnt" FROM "test" AS "t" GROUP BY "t"."string"`, "SELECT GROUP BY")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "t"."string", COUNT("t"."id") AS "cnt" FROM "test" AS "t" GROUP BY "t"."string"`, "SELECT GROUP BY")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("Having", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table).
				Fields(
					Test.Column.String.Expr(),
					Count(Test.Column.ID.Expr(), false).As("cnt"),
				).
				GroupBy(
					Test.Column.String.Expr(),
				).
				Having(
					Greater(Count(Test.Column.ID.Expr(), false), Value[int64](2)),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`string`, COUNT(`t`.`id`) AS `cnt` FROM `test` AS `t` GROUP BY `t`.`string` HAVING COUNT(`t`.`id`) > ?", "SELECT HAVING")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [t].[string], COUNT([t].[id]) AS [cnt] FROM [test] AS [t] GROUP BY [t].[string] HAVING COUNT([t].[id]) > @p1", "SELECT HAVING")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`string`, COUNT(`t`.`id`) AS `cnt` FROM `test` AS `t` GROUP BY `t`.`string` HAVING COUNT(`t`.`id`) > ?", "SELECT HAVING")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "t"."string", COUNT("t"."id") AS "cnt" FROM "test" AS "t" GROUP BY "t"."string" HAVING COUNT("t"."id") > $1`, "SELECT HAVING")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "t"."string", COUNT("t"."id") AS "cnt" FROM "test" AS "t" GROUP BY "t"."string" HAVING COUNT("t"."id") > ?`, "SELECT HAVING")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("Join", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table).
				Fields(
					Test.Column.ID.Expr(),
					Data.Column.String.Expr(),
				).
				Join(
					Inner(Data.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`id`, `d`.`string` FROM `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id`", "SELECT JOIN")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [t].[id], [d].[string] FROM [test] AS [t] INNER JOIN [data] AS [d] ON [t].[id] = [d].[id]", "SELECT JOIN")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`id`, `d`.`string` FROM `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id`", "SELECT JOIN")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "t"."id", "d"."string" FROM "test" AS "t" INNER JOIN "data" AS "d" ON "t"."id" = "d"."id"`, "SELECT JOIN")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "t"."id", "d"."string" FROM "test" AS "t" INNER JOIN "data" AS "d" ON "t"."id" = "d"."id"`, "SELECT JOIN")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("OrderBy", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table).
				Fields(
					Test.Column.ID.Expr(),
				).
				OrderBy(
					Desc(Test.Column.Number.Expr()),
					Asc(Test.Column.String.Expr()),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`id` FROM `test` AS `t` ORDER BY `t`.`number` DESC, `t`.`string` ASC", "SELECT ORDER BY")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [t].[id] FROM [test] AS [t] ORDER BY [t].[number] DESC, [t].[string] ASC", "SELECT ORDER BY")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`id` FROM `test` AS `t` ORDER BY `t`.`number` DESC, `t`.`string` ASC", "SELECT ORDER BY")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "t"."id" FROM "test" AS "t" ORDER BY "t"."number" DESC, "t"."string" ASC`, "SELECT ORDER BY")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "t"."id" FROM "test" AS "t" ORDER BY "t"."number" DESC, "t"."string" ASC`, "SELECT ORDER BY")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("Pagination", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table).
				Fields(
					Test.Column.ID.Expr(),
				).
				Pagination(10, 20)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`id` FROM `test` AS `t` LIMIT ? OFFSET ?", "SELECT PAGINATION")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [t].[id] FROM [test] AS [t] ORDER BY 1 ASC OFFSET @p1 ROWS FETCH NEXT @p2 ROWS ONLY", "SELECT PAGINATION")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`id` FROM `test` AS `t` LIMIT ? OFFSET ?", "SELECT PAGINATION")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "t"."id" FROM "test" AS "t" LIMIT $1 OFFSET $2`, "SELECT PAGINATION")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "t"."id" FROM "test" AS "t" LIMIT ? OFFSET ?`, "SELECT PAGINATION")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("Unions", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table).
				Fields(
					Test.Column.String.Expr(),
				).
				Unions(
					UnionAll(NewSelect(Data.Table).
						Fields(
							Data.Column.String.Expr(),
						),
					),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`string` FROM `test` AS `t` UNION ALL SELECT `d`.`string` FROM `data` AS `d`", "SELECT UNION ALL")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [t].[string] FROM [test] AS [t] UNION ALL SELECT [d].[string] FROM [data] AS [d]", "SELECT UNION ALL")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`string` FROM `test` AS `t` UNION ALL SELECT `d`.`string` FROM `data` AS `d`", "SELECT UNION ALL")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "t"."string" FROM "test" AS "t" UNION ALL SELECT "d"."string" FROM "data" AS "d"`, "SELECT UNION ALL")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "t"."string" FROM "test" AS "t" UNION ALL SELECT "d"."string" FROM "data" AS "d"`, "SELECT UNION ALL")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("Where", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table).
				Fields(
					Test.Column.ID.Expr(),
				).
				Where(
					Equal(Test.Column.Number.Expr(), Value(2)),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?", "SELECT")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] = @p1", "SELECT")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?", "SELECT")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = $1`, "SELECT")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = ?`, "SELECT")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("With", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(NewCTE("cte_test", "ct")).
				Fields(
					Field[int64]("ct", "id"),
				).
				With(
					WithN("cte_test", NewSelect(Test.Table).
						Fields(
							Test.Column.ID.Expr(),
						).
						Where(
							Greater(Test.Column.Number.Expr(), Value(2)),
						),
					),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "WITH `cte_test` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` > ?) SELECT `ct`.`id` FROM `cte_test` AS `ct`", "SELECT WITH")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "WITH [cte_test] AS (SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] > @p1) SELECT [ct].[id] FROM [cte_test] AS [ct]", "SELECT WITH")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "WITH `cte_test` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` > ?) SELECT `ct`.`id` FROM `cte_test` AS `ct`", "SELECT WITH")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `WITH "cte_test" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" > $1) SELECT "ct"."id" FROM "cte_test" AS "ct"`, "SELECT WITH")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `WITH "cte_test" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" > ?) SELECT "ct"."id" FROM "cte_test" AS "ct"`, "SELECT WITH")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
}
func Test_SQL_Truncate(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtTruncate := NewTruncate(Test.Table)
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
	})
	t.Run("Cascade", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtTruncate := NewTruncate(Test.Table).
				IsCascade()
			sqlTruncateQuery, sqlTruncateArguments, err := sql.Build(stmtTruncate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlTruncateQuery, "TRUNCATE TABLE `test` CASCADE", "TRUNCATE CASCADE")
			case DialectMsSQL:
				// Not supported - CASCADE
			case DialectMySQL:
				// Not supported - CASCADE
			case DialectPostgreSQL:
				assertContains(t, sqlTruncateQuery, `TRUNCATE TABLE "test" CASCADE`, "TRUNCATE CASCADE")
			case DialectSQLite:
				// Not supported - CASCADE
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlTruncateArguments, supportDialect.name, sqlTruncateQuery)
		})
	})
	t.Run("RestartIdentity", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtTruncate := NewTruncate(Test.Table).
				IsRestartIdentity()
			sqlTruncateQuery, sqlTruncateArguments, err := sql.Build(stmtTruncate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlTruncateQuery, "TRUNCATE TABLE `test` RESTART IDENTITY", "TRUNCATE RESTART IDENTITY")
			case DialectMsSQL:
				// Not supported - RESTART IDENTITY
			case DialectMySQL:
				// Not supported - RESTART IDENTITY
			case DialectPostgreSQL:
				assertContains(t, sqlTruncateQuery, `TRUNCATE TABLE "test" RESTART IDENTITY`, "TRUNCATE RESTART IDENTITY")
			case DialectSQLite:
				// Not supported - RESTART IDENTITY
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlTruncateArguments, supportDialect.name, sqlTruncateQuery)
		})
	})
}
func Test_SQL_Update(t *testing.T) {
	t.Run("Join", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtUpdate := NewUpdate(Test.Table).
				Set(
					Assign(Test.Column.String.Expr(), Value("active")),
				).
				Join(
					Inner(Data.Table, Equal(Test.Column.ID.Expr(), Data.Column.ID.Expr())),
				).
				Where(
					Equal(Data.Column.String.Expr(), Value("active")),
				)
			sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id` SET `t`.`string` = ? WHERE `d`.`string` = ?", "UPDATE JOIN")
			case DialectMsSQL:
				assertContains(t, sqlUpdateQuery, "UPDATE [test] AS [t] INNER JOIN [data] AS [d] ON [t].[id] = [d].[id] SET [t].[string] = @p1 WHERE [d].[string] = @p2", "UPDATE JOIN")
			case DialectMySQL:
				assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id` SET `t`.`string` = ? WHERE `d`.`string` = ?", "UPDATE JOIN")
			case DialectPostgreSQL:
				assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t" INNER JOIN "data" AS "d" ON "t"."id" = "d"."id" SET "t"."string" = $1 WHERE "d"."string" = $2`, "UPDATE JOIN")
			case DialectSQLite:
				assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t" INNER JOIN "data" AS "d" ON "t"."id" = "d"."id" SET "t"."string" = ? WHERE "d"."string" = ?`, "UPDATE JOIN")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
		})
	})
	t.Run("Returning", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtUpdate := NewUpdate(Test.Table).
				Set(
					Assign(Test.Column.String.Expr(), Value("active")),
				).
				Where(
					Equal(Test.Column.Number.Expr(), Value(2)),
				).
				Returning(
					Test.Column.ID.Expr(),
					Test.Column.String.Expr(),
				)
			sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ? RETURNING `t`.`id`, `t`.`string`", "UPDATE RETURNING")
			case DialectMsSQL:
				assertContains(t, sqlUpdateQuery, "UPDATE [test] AS [t] OUTPUT [t].[id], [t].[string] SET [t].[string] = @p1 WHERE [t].[number] = @p2", "UPDATE RETURNING")
			case DialectMySQL:
				assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?", "UPDATE WITHOUT RETURNING")
			case DialectPostgreSQL:
				assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2 RETURNING "t"."id", "t"."string"`, "UPDATE RETURNING")
			case DialectSQLite:
				assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t" SET "t"."string" = ? WHERE "t"."number" = ? RETURNING "t"."id", "t"."string"`, "UPDATE RETURNING")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
		})
	})
	t.Run("Set", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtUpdate := NewUpdate(Test.Table).
				Set(
					Assign(Test.Column.String.Expr(), Value("active")),
				).
				Where(
					Equal(Test.Column.Number.Expr(), Value(2)),
				)
			sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?", "UPDATE")
			case DialectMsSQL:
				assertContains(t, sqlUpdateQuery, "UPDATE [test] AS [t] SET [t].[string] = @p1 WHERE [t].[number] = @p2", "UPDATE")
			case DialectMySQL:
				assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?", "UPDATE")
			case DialectPostgreSQL:
				assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2`, "UPDATE")
			case DialectSQLite:
				assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t" SET "t"."string" = ? WHERE "t"."number" = ?`, "UPDATE")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
		})
	})
	t.Run("Where", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtUpdate := NewUpdate(Test.Table).
				Set(
					Assign(Test.Column.String.Expr(), Value("active")),
				).
				Where(
					Equal(Test.Column.Number.Expr(), Value(2)),
				)
			sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?", "UPDATE")
			case DialectMsSQL:
				assertContains(t, sqlUpdateQuery, "UPDATE [test] AS [t] SET [t].[string] = @p1 WHERE [t].[number] = @p2", "UPDATE")
			case DialectMySQL:
				assertContains(t, sqlUpdateQuery, "UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?", "UPDATE")
			case DialectPostgreSQL:
				assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2`, "UPDATE")
			case DialectSQLite:
				assertContains(t, sqlUpdateQuery, `UPDATE "test" AS "t" SET "t"."string" = ? WHERE "t"."number" = ?`, "UPDATE")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
		})
	})
	t.Run("With", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtUpdate := NewUpdate(Test.Table).
				Set(
					Assign(Test.Column.String.Expr(), Value("updated")),
				).
				With(
					WithN("old_users", NewSelect(Test.Table).
						Fields(
							Test.Column.ID.Expr(),
						).
						Where(
							Less(Test.Column.Number.Expr(), Value(2)),
						),
					),
				)
			sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlUpdateQuery, "WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) UPDATE `test` AS `t` SET `t`.`string` = ?", "UPDATE WITH")
			case DialectMsSQL:
				assertContains(t, sqlUpdateQuery, "WITH [old_users] AS (SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] < @p1) UPDATE [test] AS [t] SET [t].[string] = @p2", "UPDATE WITH")
			case DialectMySQL:
				assertContains(t, sqlUpdateQuery, "WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) UPDATE `test` AS `t` SET `t`.`string` = ?", "UPDATE WITH")
			case DialectPostgreSQL:
				assertContains(t, sqlUpdateQuery, `WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < $1) UPDATE "test" AS "t" SET "t"."string" = $2`, "UPDATE WITH")
			case DialectSQLite:
				assertContains(t, sqlUpdateQuery, `WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < ?) UPDATE "test" AS "t" SET "t"."string" = ?`, "UPDATE WITH")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
		})
	})
}
func Test_Transformer_Comparison(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(WithDialect(supportDialect))
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table).
			Fields(Test.Column.ID.Expr()).
			Where(
				And(
					ILike(Test.Column.String.Expr(), Value("%alex%")),
					And(
						ILike(Test.Column.String.Expr(), Value("%ivan%")),
						ILike(Test.Column.String.Expr(), Value("%petr%")),
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
			Fields(
				Trunc(Ceil(Test.Column.Number.Expr()), Value(2)).As("result"),
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
	Data.Column.Date = NewColumn[time.Time]("date", Data.Table, TypeDate)
	Data.Column.ID = NewColumn[int64]("id", Data.Table, TypeBigInt)
	Data.Column.Json = NewColumn[string]("json", Data.Table, TypeJSON)
	Data.Column.Number = NewColumn[int]("number", Data.Table, TypeInt)
	Data.Column.String = NewColumn[string]("string", Data.Table, TypeVarChar)
	Data.Column.Time = NewColumn[time.Time]("time", Data.Table, TypeTime)
	// Test
	Test.Column.CreateAt = NewColumn[time.Time]("createat", Test.Table, TypeTimestamp)
	Test.Column.DataID = NewColumn[int64]("data_id", Test.Table, TypeBigInt)
	Test.Column.Date = NewColumn[time.Time]("date", Test.Table, TypeDate)
	Test.Column.ID = NewColumn[int64]("id", Test.Table, TypeBigInt)
	Test.Column.Json = NewColumn[string]("json", Test.Table, TypeJSON)
	Test.Column.Name = NewColumn[string]("name", Test.Table, TypeVarChar)
	Test.Column.Number = NewColumn[int]("number", Test.Table, TypeInt)
	Test.Column.String = NewColumn[string]("string", Test.Table, TypeVarChar)
	Test.Column.UpdateAt = NewColumn[time.Time]("updateat", Test.Table, TypeTimestamp)
	Test.Column.X = NewColumn[int]("x", Test.Table, TypeInt)
	Test.Column.Y = NewColumn[int]("y", Test.Table, TypeInt)
}
func testAllDialects(t *testing.T, testFunc func(t *testing.T, supportDialect *SupportDialect)) {
	for _, supportDialect := range listSupportDialects {
		currentDialect := supportDialect
		t.Run(currentDialect.name, func(t *testing.T) {
			testFunc(t, currentDialect)
		})
	}
}
