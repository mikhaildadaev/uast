package uast

import (
	"regexp"
	"strings"
	"testing"
	"time"
)

// Публичные переменные
var Test struct {
	// Constraint
	Check struct {
		OrdersNumber *CheckConstraint
		UsersNumber  *CheckConstraint
	}
	Foreign struct {
		OrdersUsers *ForeignConstraint
		UsersOrders *ForeignConstraint
	}
	Primary struct {
		OrdersID *PrimaryConstraint
		UsersID  *PrimaryConstraint
	}
	Unique struct {
		OrdersName *UniqueConstraint
		UsersName  *UniqueConstraint
	}
	// Source
	Index struct {
		OrdersID *IndexSource
		UsersID  *IndexSource
	}
	Schema *SchemaSource
	Table  struct {
		Order  *TableSource
		Orders struct {
			Date   *ColumnSource[time.Time]
			ID     *ColumnSource[int64]
			Json   *ColumnSource[string]
			Name   *ColumnSource[string]
			Number *ColumnSource[int]
			String *ColumnSource[string]
			Time   *ColumnSource[time.Time]
			UserID *ColumnSource[int64]
		}
		User  *TableSource
		Users struct {
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
	}
	View struct {
		OrdersGeneral *ViewSource
		UsersGeneral  *ViewSource
	}
}

// Публичные функции
func Test_Core_clauseGroupBy(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(
			WithDialect(supportDialect),
		)
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.String.Expr(),
			).
			GroupBy(
				Test.Table.Users.String.Expr(),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "GROUP BY `u`.`string`", "GROUP BY")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "GROUP BY [u].[string]", "GROUP BY")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "GROUP BY `u`.`string`", "GROUP BY")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `GROUP BY "u"."string"`, "GROUP BY")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `GROUP BY "u"."string"`, "GROUP BY")
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.String.Expr(),
			).
			Having(
				Greater(Count(Test.Table.Users.ID.Expr(), false), Value[int64](2)),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "HAVING COUNT(`u`.`id`) > ?", "HAVING")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "HAVING COUNT([u].[id]) > @p1", "HAVING")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "HAVING COUNT(`u`.`id`) > ?", "HAVING")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `HAVING COUNT("u"."id") > $1`, "HAVING")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `HAVING COUNT("u"."id") > ?`, "HAVING")
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
		stmtSelect := NewSelect(Test.Table.Order).
			Fields(
				Test.Table.Orders.ID.Expr(),
			).
			Join(
				Cross(Test.Table.User),
				Full(Test.Table.User, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				FullOuter(Test.Table.User, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				Inner(Test.Table.User, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				Left(Test.Table.User, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				LeftOuter(Test.Table.User, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				Right(Test.Table.User, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				RightOuter(Test.Table.User, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "CROSS JOIN `users` AS `u`", "CROSS JOIN")
			assertContains(t, sqlSelectQuery, "FULL JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "FULL JOIN")
			assertContains(t, sqlSelectQuery, "FULL OUTER JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "FULL OUTER JOIN")
			assertContains(t, sqlSelectQuery, "INNER JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "INNER JOIN")
			assertContains(t, sqlSelectQuery, "LEFT JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "LEFT JOIN")
			assertContains(t, sqlSelectQuery, "LEFT OUTER JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "LEFT OUTER JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "RIGHT JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT OUTER JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "RIGHT OUTER JOIN")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "CROSS JOIN [users] AS [u]", "CROSS JOIN")
			assertContains(t, sqlSelectQuery, "FULL JOIN [users] AS [u] ON [u].[id] = [o].[id]", "FULL JOIN")
			assertContains(t, sqlSelectQuery, "FULL OUTER JOIN [users] AS [u] ON [u].[id] = [o].[id]", "FULL OUTER JOIN")
			assertContains(t, sqlSelectQuery, "INNER JOIN [users] AS [u] ON [u].[id] = [o].[id]", "INNER JOIN")
			assertContains(t, sqlSelectQuery, "LEFT JOIN [users] AS [u] ON [u].[id] = [o].[id]", "LEFT JOIN")
			assertContains(t, sqlSelectQuery, "LEFT OUTER JOIN [users] AS [u] ON [u].[id] = [o].[id]", "LEFT OUTER JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT JOIN [users] AS [u] ON [u].[id] = [o].[id]", "RIGHT JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT OUTER JOIN [users] AS [u] ON [u].[id] = [o].[id]", "RIGHT OUTER JOIN")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "CROSS JOIN `users` AS `u`", "CROSS JOIN")
			assertContains(t, sqlSelectQuery, "FULL JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "FULL JOIN")
			assertContains(t, sqlSelectQuery, "FULL OUTER JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "FULL OUTER JOIN")
			assertContains(t, sqlSelectQuery, "INNER JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "INNER JOIN")
			assertContains(t, sqlSelectQuery, "LEFT JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "LEFT JOIN")
			assertContains(t, sqlSelectQuery, "LEFT OUTER JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "LEFT OUTER JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "RIGHT JOIN")
			assertContains(t, sqlSelectQuery, "RIGHT OUTER JOIN `users` AS `u` ON `u`.`id` = `o`.`id`", "RIGHT OUTER JOIN")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `CROSS JOIN "users" AS "u"`, "CROSS JOIN")
			assertContains(t, sqlSelectQuery, `FULL JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "FULL JOIN")
			assertContains(t, sqlSelectQuery, `FULL OUTER JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "FULL OUTER JOIN")
			assertContains(t, sqlSelectQuery, `INNER JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "INNER JOIN")
			assertContains(t, sqlSelectQuery, `LEFT JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "LEFT JOIN")
			assertContains(t, sqlSelectQuery, `LEFT OUTER JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "LEFT OUTER JOIN")
			assertContains(t, sqlSelectQuery, `RIGHT JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "RIGHT JOIN")
			assertContains(t, sqlSelectQuery, `RIGHT OUTER JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "RIGHT OUTER JOIN")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `CROSS JOIN "users" AS "u"`, "CROSS JOIN")
			assertContains(t, sqlSelectQuery, `FULL JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "FULL JOIN")
			assertContains(t, sqlSelectQuery, `FULL OUTER JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "FULL OUTER JOIN")
			assertContains(t, sqlSelectQuery, `INNER JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "INNER JOIN")
			assertContains(t, sqlSelectQuery, `LEFT JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "LEFT JOIN")
			assertContains(t, sqlSelectQuery, `LEFT OUTER JOIN "users" AS "u" ON "u"."id" = "o"."id"`, "LEFT OUTER JOIN")
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
			).
			OrderBy(
				Asc(Test.Table.Users.String.Expr()),
				Desc(Test.Table.Users.String.Expr()),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "ORDER BY", "ORDER BY")
			assertContains(t, sqlSelectQuery, "`u`.`string` ASC", "ORDER BY ASC")
			assertContains(t, sqlSelectQuery, "`u`.`string` DESC", "ORDER BY DESC")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "ORDER BY", "ORDER BY")
			assertContains(t, sqlSelectQuery, "[u].[string] ASC", "ORDER BY ASC")
			assertContains(t, sqlSelectQuery, "[u].[string] DESC", "ORDER BY DESC")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "ORDER BY", "ORDER BY")
			assertContains(t, sqlSelectQuery, "`u`.`string` ASC", "ORDER BY ASC")
			assertContains(t, sqlSelectQuery, "`u`.`string` DESC", "ORDER BY DESC")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `ORDER BY`, "ORDER BY")
			assertContains(t, sqlSelectQuery, `"u"."string" ASC`, "ORDER BY ASC")
			assertContains(t, sqlSelectQuery, `"u"."string" DESC`, "ORDER BY DESC")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `ORDER BY`, "ORDER BY")
			assertContains(t, sqlSelectQuery, `"u"."string" ASC`, "ORDER BY ASC")
			assertContains(t, sqlSelectQuery, `"u"."string" DESC`, "ORDER BY DESC")
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
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
		stmtDelete := NewDelete(Test.Table.User).
			Returning(
				Test.Table.Users.ID.Expr(),
				Test.Table.Users.String.Expr(),
			)
		sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlDeleteQuery, "RETURNING `u`.`id`, `u`.`string`", "RETURNING")
		case DialectMsSQL:
			assertContains(t, sqlDeleteQuery, "OUTPUT [u].[id], [u].[string]", "RETURNING")
		case DialectMySQL:
			// Not supported - RETURNING
		case DialectPostgreSQL:
			assertContains(t, sqlDeleteQuery, `RETURNING "u"."id", "u"."string"`, "RETURNING")
		case DialectSQLite:
			assertContains(t, sqlDeleteQuery, `RETURNING "u"."id", "u"."string"`, "RETURNING")
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
		stmtUpdate := NewUpdate(Test.Table.User).
			Set(
				Assign(Test.Table.Users.String.Expr(), Value("active")),
			)
		sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlUpdateQuery, "SET `u`.`string` = ?", "SET")
		case DialectMsSQL:
			assertContains(t, sqlUpdateQuery, "SET [u].[string] = @p1", "SET")
		case DialectMySQL:
			assertContains(t, sqlUpdateQuery, "SET `u`.`string` = ?", "SET")
		case DialectPostgreSQL:
			assertContains(t, sqlUpdateQuery, `SET "u"."string" = $1`, "SET")
		case DialectSQLite:
			assertContains(t, sqlUpdateQuery, `SET "u"."string" = ?`, "SET")
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.String.Expr(),
			).
			Unions(
				Union(NewSelect(Test.Table.User).
					Fields(
						Test.Table.Users.String.Expr(),
					),
				),
				UnionAll(NewSelect(Test.Table.User).
					Fields(
						Test.Table.Users.String.Expr(),
					),
				),
				UnionExcept(NewSelect(Test.Table.User).
					Fields(
						Test.Table.Users.String.Expr(),
					),
				),
				UnionIntersect(NewSelect(Test.Table.User).
					Fields(
						Test.Table.Users.String.Expr(),
					),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "UNION SELECT `u`.`string` FROM `users` AS `u`", "UNION")
			assertContains(t, sqlSelectQuery, "UNION ALL SELECT `u`.`string` FROM `users` AS `u`", "UNION ALL")
			assertContains(t, sqlSelectQuery, "EXCEPT SELECT `u`.`string` FROM `users` AS `u`", "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, "INTERSECT SELECT `u`.`string` FROM `users` AS `u`", "UNION INTERSECT")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "UNION SELECT [u].[string] FROM [users] AS [u]", "UNION")
			assertContains(t, sqlSelectQuery, "UNION ALL SELECT [u].[string] FROM [users] AS [u]", "UNION ALL")
			assertContains(t, sqlSelectQuery, "EXCEPT SELECT [u].[string] FROM [users] AS [u]", "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, "INTERSECT SELECT [u].[string] FROM [users] AS [u]", "UNION INTERSECT")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "UNION SELECT `u`.`string` FROM `users` AS `u`", "UNION")
			assertContains(t, sqlSelectQuery, "UNION ALL SELECT `u`.`string` FROM `users` AS `u`", "UNION ALL")
			assertContains(t, sqlSelectQuery, "EXCEPT SELECT `u`.`string` FROM `users` AS `u`", "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, "INTERSECT SELECT `u`.`string` FROM `users` AS `u`", "UNION INTERSECT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `UNION SELECT "u"."string" FROM "users" AS "u"`, "UNION")
			assertContains(t, sqlSelectQuery, `UNION ALL SELECT "u"."string" FROM "users" AS "u"`, "UNION ALL")
			assertContains(t, sqlSelectQuery, `EXCEPT SELECT "u"."string" FROM "users" AS "u"`, "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, `INTERSECT SELECT "u"."string" FROM "users" AS "u"`, "UNION INTERSECT")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `UNION SELECT "u"."string" FROM "users" AS "u"`, "UNION")
			assertContains(t, sqlSelectQuery, `UNION ALL SELECT "u"."string" FROM "users" AS "u"`, "UNION ALL")
			assertContains(t, sqlSelectQuery, `EXCEPT SELECT "u"."string" FROM "users" AS "u"`, "UNION EXCEPT")
			assertContains(t, sqlSelectQuery, `INTERSECT SELECT "u"."string" FROM "users" AS "u"`, "UNION INTERSECT")
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
		stmtInsert := NewInsert(Test.Table.User).
			Values(
				Pair(Test.Table.Users.String.Expr(), Value("ivan")),
				Pair(Test.Table.Users.Number.Expr(), Value(2)),
			).
			Upsert(
				Pair(Test.Table.Users.String.Expr(), Value("updated")),
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
			).
			Where(
				Equal(Test.Table.Users.String.Expr(), Value("active")),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "WHERE `u`.`string` = ?", "WHERE")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "WHERE [u].[string] = @p1", "WHERE")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "WHERE `u`.`string` = ?", "WHERE")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `WHERE "u"."string" = $1`, "WHERE")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `WHERE "u"."string" = ?`, "WHERE")
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
		stmtWithN := WithN("cte_norecursive", NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
				Test.Table.Users.String.Expr(),
			).
			Where(
				Equal(Test.Table.Users.String.Expr(), Value("active")),
			),
			"id", "string",
		)
		stmtWithR := WithR("cte_recursive", NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
				Test.Table.Users.String.Expr(),
			).
			Where(
				Equal(Test.Table.Users.String.Expr(), Value("active")),
			).
			Unions(
				UnionAll(NewSelect(Test.Table.User).
					Fields(
						Test.Table.Users.ID.Expr(),
						Test.Table.Users.String.Expr(),
					).
					Join(
						Inner(NewCTE("cte_recursive", "rec"), Equal(Test.Table.Users.ID.Expr(), Field[int64]("rec", "id"))),
					),
				),
			),
			"id", "string",
		)
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
				Test.Table.Users.Number.Expr(),
			).
			Join(
				Inner(NewCTE("cte_norecursive", "cnr"), Equal(Test.Table.Users.ID.Expr(), Field[int64]("cnr", "id"))),
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
		stmtSelect := NewSelect(Test.Table.User).
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
			).
			Where(
				And(
					Equal(Test.Table.Users.Number.Expr(), BitwiseAnd(Test.Table.Users.Number.Expr(), Value(0b0010))),
					Equal(Test.Table.Users.Number.Expr(), BitwiseOr(Test.Table.Users.Number.Expr(), Value(0b0010))),
					Equal(Test.Table.Users.Number.Expr(), BitwiseXor(Test.Table.Users.Number.Expr(), Value(0b0010))),
					Equal(Test.Table.Users.Number.Expr(), Divide(Test.Table.Users.Number.Expr(), Value(2))),
					Equal(Test.Table.Users.Number.Expr(), Minus(Test.Table.Users.Number.Expr(), Value(2))),
					Equal(Test.Table.Users.Number.Expr(), Modulo(Test.Table.Users.Number.Expr(), Value(2))),
					Equal(Test.Table.Users.Number.Expr(), Multiply(Test.Table.Users.Number.Expr(), Value(2))),
					Equal(Test.Table.Users.Number.Expr(), Plus(Test.Table.Users.Number.Expr(), Value(2))),
					Equal(Test.Table.Users.Number.Expr(), ShiftLeft(Test.Table.Users.Number.Expr(), Value(2))),
					Equal(Test.Table.Users.Number.Expr(), ShiftRight(Test.Table.Users.Number.Expr(), Value(2))),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "`u`.`number` & ?", "BINARY BITWISE AND")
			assertContains(t, sqlSelectQuery, "`u`.`number` | ?", "BINARY BITWISE OR")
			assertContains(t, sqlSelectQuery, "`u`.`number` ^ ?", "BINARY BITWISE XOR")
			assertContains(t, sqlSelectQuery, "`u`.`number` / ?", "BINARY DIVIDE")
			assertContains(t, sqlSelectQuery, "`u`.`number` - ?", "BINARY MINUS")
			assertContains(t, sqlSelectQuery, "`u`.`number` % ?", "BINARY MODULO")
			assertContains(t, sqlSelectQuery, "`u`.`number` * ?", "BINARY MULTIPLY")
			assertContains(t, sqlSelectQuery, "`u`.`number` + ?", "BINARY PLUS")
			assertContains(t, sqlSelectQuery, "`u`.`number` << ?", "BINARY SHIFT LEFT")
			assertContains(t, sqlSelectQuery, "`u`.`number` >> ?", "BINARY SHIFT RIGHT")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "[u].[number] & @p1", "BINARY BITWISE AND")
			assertContains(t, sqlSelectQuery, "[u].[number] | @p1", "BINARY BITWISE OR")
			assertContains(t, sqlSelectQuery, "[u].[number] ^ @p1", "BINARY BITWISE XOR")
			assertContains(t, sqlSelectQuery, "[u].[number] / @p1", "BINARY DIVIDE")
			assertContains(t, sqlSelectQuery, "[u].[number] - @p1", "BINARY MINUS")
			assertContains(t, sqlSelectQuery, "[u].[number] % @p1", "BINARY MODULO")
			assertContains(t, sqlSelectQuery, "[u].[number] * @p1", "BINARY MULTIPLY")
			assertContains(t, sqlSelectQuery, "[u].[number] + @p1", "BINARY PLUS")
			assertContains(t, sqlSelectQuery, "[u].[number] << @p1", "BINARY SHIFT LEFT")
			assertContains(t, sqlSelectQuery, "[u].[number] >> @p1", "BINARY SHIFT RIGHT")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`u`.`number` & ?", "BINARY BITWISE AND")
			assertContains(t, sqlSelectQuery, "`u`.`number` | ?", "BINARY BITWISE OR")
			assertContains(t, sqlSelectQuery, "`u`.`number` ^ ?", "BINARY BITWISE XOR")
			assertContains(t, sqlSelectQuery, "`u`.`number` / ?", "BINARY DIVIDE")
			assertContains(t, sqlSelectQuery, "`u`.`number` - ?", "BINARY MINUS")
			assertContains(t, sqlSelectQuery, "`u`.`number` % ?", "BINARY MODULO")
			assertContains(t, sqlSelectQuery, "`u`.`number` * ?", "BINARY MULTIPLY")
			assertContains(t, sqlSelectQuery, "`u`.`number` + ?", "BINARY PLUS")
			assertContains(t, sqlSelectQuery, "`u`.`number` << ?", "BINARY SHIFT LEFT")
			assertContains(t, sqlSelectQuery, "`u`.`number` >> ?", "BINARY SHIFT RIGHT")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"u"."number" & $1`, "BINARY BITWISE AND")
			assertContains(t, sqlSelectQuery, `"u"."number" | $1`, "BINARY BITWISE OR")
			assertContains(t, sqlSelectQuery, `"u"."number" ^ $1`, "BINARY BITWISE XOR")
			assertContains(t, sqlSelectQuery, `"u"."number" / $1`, "BINARY DIVIDE")
			assertContains(t, sqlSelectQuery, `"u"."number" - $1`, "BINARY MINUS")
			assertContains(t, sqlSelectQuery, `"u"."number" % $1`, "BINARY MODULO")
			assertContains(t, sqlSelectQuery, `"u"."number" * $1`, "BINARY MULTIPLY")
			assertContains(t, sqlSelectQuery, `"u"."number" + $1`, "BINARY PLUS")
			assertContains(t, sqlSelectQuery, `"u"."number" << $1`, "BINARY SHIFT LEFT")
			assertContains(t, sqlSelectQuery, `"u"."number" >> $1`, "BINARY SHIFT RIGHT")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `"u"."number" & ?`, "BINARY BITWISE AND")
			assertContains(t, sqlSelectQuery, `"u"."number" | ?`, "BINARY BITWISE OR")
			assertContains(t, sqlSelectQuery, `"u"."number" ^ ?`, "BINARY BITWISE XOR")
			assertContains(t, sqlSelectQuery, `"u"."number" / ?`, "BINARY DIVIDE")
			assertContains(t, sqlSelectQuery, `"u"."number" - ?`, "BINARY MINUS")
			assertContains(t, sqlSelectQuery, `"u"."number" % ?`, "BINARY MODULO")
			assertContains(t, sqlSelectQuery, `"u"."number" * ?`, "BINARY MULTIPLY")
			assertContains(t, sqlSelectQuery, `"u"."number" + ?`, "BINARY PLUS")
			assertContains(t, sqlSelectQuery, `"u"."number" << ?`, "BINARY SHIFT LEFT")
			assertContains(t, sqlSelectQuery, `"u"."number" >> ?`, "BINARY SHIFT RIGHT")
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "`u`.`id`", "COLUMN ID")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "[u].[id]", "COLUMN ID")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`u`.`id`", "COLUMN ID")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"u"."id"`, "COLUMN ID")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `"u"."id"`, "COLUMN ID")
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
			).
			Where(
				And(
					Between(Test.Table.Users.Number.Expr(), Value(0), Value(2)),
					Equal(Test.Table.Users.Number.Expr(), Value(2)),
					Exists(Subquery[int](NewSelect(Test.Table.User).Fields(ConstIntOne()))),
					Greater(Test.Table.Users.Number.Expr(), Value(2)),
					GreaterEqual(Test.Table.Users.Number.Expr(), Value(2)),
					ILike(Test.Table.Users.String.Expr(), Value("%ivan%")),
					In(Test.Table.Users.String.Expr(), Array("active", "pending")),
					IsNotNull(Test.Table.Users.String.Expr()),
					IsNull(Test.Table.Users.String.Expr()),
					Less(Test.Table.Users.Number.Expr(), Value(2)),
					LessEqual(Test.Table.Users.Number.Expr(), Value(2)),
					Like(Test.Table.Users.String.Expr(), Value("%ivan%")),
					NotBetween(Test.Table.Users.Number.Expr(), Value(0), Value(2)),
					NotEqual(Test.Table.Users.Number.Expr(), Value(2)),
					NotExists(Subquery[int](NewSelect(Test.Table.User).Fields(ConstIntOne()))),
					NotILike(Test.Table.Users.String.Expr(), Value("%ivan%")),
					NotIn(Test.Table.Users.String.Expr(), Array("active", "pending")),
					NotLike(Test.Table.Users.String.Expr(), Value("%ivan%")),
				))
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "`u`.`number` BETWEEN ? AND ?", "COMPARISON BETWEEN")
			assertContains(t, sqlSelectQuery, "`u`.`number` = ?", "COMPARISON EQUAL")
			assertContains(t, sqlSelectQuery, "EXISTS (SELECT 1 FROM `users` AS `u`)", "COMPARISON EXISTS")
			assertContains(t, sqlSelectQuery, "`u`.`number` > ?", "COMPARISON GREATER")
			assertContains(t, sqlSelectQuery, "`u`.`number` >= ?", "COMPARISON GREATEREQUAL")
			assertContains(t, sqlSelectQuery, "LOWER(`u`.`string`) LIKE LOWER(?)", "COMPARISON ILIKE")
			assertContains(t, sqlSelectQuery, "`u`.`string` IN (?, ?)", "COMPARISON IN")
			assertContains(t, sqlSelectQuery, "`u`.`string` IS NOT NULL", "COMPARISON IS NOT NULL")
			assertContains(t, sqlSelectQuery, "`u`.`string` IS NULL", "COMPARISON IS NULL")
			assertContains(t, sqlSelectQuery, "`u`.`number` < ?", "COMPARISON LESS")
			assertContains(t, sqlSelectQuery, "`u`.`number` <= ?", "COMPARISON LESSEQUAL")
			assertContains(t, sqlSelectQuery, "`u`.`string` LIKE ?", "COMPARISON LIKE")
			assertContains(t, sqlSelectQuery, "`u`.`number` NOT BETWEEN ? AND ?", "COMPARISON NOT BETWEEN")
			assertContains(t, sqlSelectQuery, "`u`.`number` <> ?", "COMPARISON NOT EQUAL")
			assertContains(t, sqlSelectQuery, "NOT EXISTS (SELECT 1 FROM `users` AS `u`)", "COMPARISON NOT EXISTS")
			assertContains(t, sqlSelectQuery, "LOWER(`u`.`string`) NOT LIKE LOWER(?)", "COMPARISON NOT ILIKE")
			assertContains(t, sqlSelectQuery, "`u`.`string` NOT IN (?, ?)", "COMPARISON NOT IN")
			assertContains(t, sqlSelectQuery, "`u`.`string` NOT LIKE ?", "COMPARISON NOT LIKE")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "[u].[number] BETWEEN @p1 AND @p2", "COMPARISON BETWEEN")
			assertContains(t, sqlSelectQuery, "[u].[number] = @p1", "COMPARISON EQUAL")
			assertContains(t, sqlSelectQuery, "EXISTS (SELECT 1 FROM [users] AS [u])", "COMPARISON EXISTS")
			assertContains(t, sqlSelectQuery, "[u].[number] > @p1", "COMPARISON GREATER")
			assertContains(t, sqlSelectQuery, "[u].[number] >= @p1", "COMPARISON GREATEREQUAL")
			assertContains(t, sqlSelectQuery, "LOWER([u].[string]) LIKE LOWER(@p1)", "COMPARISON ILIKE")
			assertContains(t, sqlSelectQuery, "[u].[string] IN (@p1, @p2)", "COMPARISON IN")
			assertContains(t, sqlSelectQuery, "[u].[string] IS NOT NULL", "COMPARISON IS NOT NULL")
			assertContains(t, sqlSelectQuery, "[u].[string] IS NULL", "COMPARISON IS NULL")
			assertContains(t, sqlSelectQuery, "[u].[number] < @p1", "COMPARISON LESS")
			assertContains(t, sqlSelectQuery, "[u].[number] <= @p1", "COMPARISON LESSEQUAL")
			assertContains(t, sqlSelectQuery, "[u].[string] LIKE @p1", "COMPARISON LIKE")
			assertContains(t, sqlSelectQuery, "[u].[number] NOT BETWEEN @p1 AND @p2", "COMPARISON NOT BETWEEN")
			assertContains(t, sqlSelectQuery, "[u].[number] <> @p1", "COMPARISON NOT EQUAL")
			assertContains(t, sqlSelectQuery, "NOT EXISTS (SELECT 1 FROM [users] AS [u])", "COMPARISON NOT EXISTS")
			assertContains(t, sqlSelectQuery, "LOWER([u].[string]) NOT LIKE LOWER(@p1)", "COMPARISON NOT ILIKE")
			assertContains(t, sqlSelectQuery, "[u].[string] NOT IN (@p1, @p2)", "COMPARISON NOT IN")
			assertContains(t, sqlSelectQuery, "[u].[string] NOT LIKE @p1", "COMPARISON NOT LIKE")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`u`.`number` BETWEEN ? AND ?", "COMPARISON BETWEEN")
			assertContains(t, sqlSelectQuery, "`u`.`number` = ?", "COMPARISON EQUAL")
			assertContains(t, sqlSelectQuery, "EXISTS (SELECT 1 FROM `users` AS `u`)", "COMPARISON EXISTS")
			assertContains(t, sqlSelectQuery, "`u`.`number` > ?", "COMPARISON GREATER")
			assertContains(t, sqlSelectQuery, "`u`.`number` >= ?", "COMPARISON GREATEREQUAL")
			assertContains(t, sqlSelectQuery, "LOWER(`u`.`string`) LIKE LOWER(?)", "COMPARISON ILIKE")
			assertContains(t, sqlSelectQuery, "`u`.`string` IN (?, ?)", "COMPARISON IN")
			assertContains(t, sqlSelectQuery, "`u`.`string` IS NOT NULL", "COMPARISON IS NOT NULL")
			assertContains(t, sqlSelectQuery, "`u`.`string` IS NULL", "COMPARISON IS NULL")
			assertContains(t, sqlSelectQuery, "`u`.`number` < ?", "COMPARISON LESS")
			assertContains(t, sqlSelectQuery, "`u`.`number` <= ?", "COMPARISON LESSEQUAL")
			assertContains(t, sqlSelectQuery, "`u`.`string` LIKE ?", "COMPARISON LIKE")
			assertContains(t, sqlSelectQuery, "`u`.`number` NOT BETWEEN ? AND ?", "COMPARISON NOT BETWEEN")
			assertContains(t, sqlSelectQuery, "`u`.`number` <> ?", "COMPARISON NOT EQUAL")
			assertContains(t, sqlSelectQuery, "NOT EXISTS (SELECT 1 FROM `users` AS `u`)", "COMPARISON NOT EXISTS")
			assertContains(t, sqlSelectQuery, "LOWER(`u`.`string`) NOT LIKE LOWER(?)", "COMPARISON NOT ILIKE")
			assertContains(t, sqlSelectQuery, "`u`.`string` NOT IN (?, ?)", "COMPARISON NOT IN")
			assertContains(t, sqlSelectQuery, "`u`.`string` NOT LIKE ?", "COMPARISON NOT LIKE")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"u"."number" BETWEEN $1 AND $2`, "COMPARISON BETWEEN")
			assertContains(t, sqlSelectQuery, `"u"."number" = $1`, "COMPARISON EQUAL")
			assertContains(t, sqlSelectQuery, `EXISTS (SELECT 1 FROM "users" AS "u")`, "COMPARISON EXISTS")
			assertContains(t, sqlSelectQuery, `"u"."number" > $1`, "COMPARISON GREATER")
			assertContains(t, sqlSelectQuery, `"u"."number" >= $1`, "COMPARISON GREATEREQUAL")
			assertContains(t, sqlSelectQuery, `"u"."string" ILIKE $1`, "COMPARISON ILIKE")
			assertContains(t, sqlSelectQuery, `"u"."string" IN ($1, $2)`, "COMPARISON IN")
			assertContains(t, sqlSelectQuery, `"u"."string" IS NOT NULL`, "COMPARISON IS NOT NULL")
			assertContains(t, sqlSelectQuery, `"u"."string" IS NULL`, "COMPARISON IS NULL")
			assertContains(t, sqlSelectQuery, `"u"."number" < $1`, "COMPARISON LESS")
			assertContains(t, sqlSelectQuery, `"u"."number" <= $1`, "COMPARISON LESSEQUAL")
			assertContains(t, sqlSelectQuery, `"u"."string" LIKE $1`, "COMPARISON LIKE")
			assertContains(t, sqlSelectQuery, `"u"."number" NOT BETWEEN $1 AND $2`, "COMPARISON NOT BETWEEN")
			assertContains(t, sqlSelectQuery, `"u"."number" <> $1`, "COMPARISON NOT EQUAL")
			assertContains(t, sqlSelectQuery, `NOT EXISTS (SELECT 1 FROM "users" AS "u")`, "COMPARISON NOT EXISTS")
			assertContains(t, sqlSelectQuery, `"u"."string" NOT ILIKE $1`, "COMPARISON NOT ILIKE")
			assertContains(t, sqlSelectQuery, `"u"."string" NOT IN ($1, $2)`, "COMPARISON NOT IN")
			assertContains(t, sqlSelectQuery, `"u"."string" NOT LIKE $1`, "COMPARISON NOT LIKE")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `"u"."number" BETWEEN ? AND ?`, "COMPARISON BETWEEN")
			assertContains(t, sqlSelectQuery, `"u"."number" = ?`, "COMPARISON EQUAL")
			assertContains(t, sqlSelectQuery, `EXISTS (SELECT 1 FROM "users" AS "u")`, "COMPARISON EXISTS")
			assertContains(t, sqlSelectQuery, `"u"."number" > ?`, "COMPARISON GREATER")
			assertContains(t, sqlSelectQuery, `"u"."number" >= ?`, "COMPARISON GREATEREQUAL")
			assertContains(t, sqlSelectQuery, `LOWER("u"."string") LIKE LOWER(?)`, "COMPARISON ILIKE")
			assertContains(t, sqlSelectQuery, `"u"."string" IN (?, ?)`, "COMPARISON IN")
			assertContains(t, sqlSelectQuery, `"u"."string" IS NOT NULL`, "COMPARISON IS NOT NULL")
			assertContains(t, sqlSelectQuery, `"u"."string" IS NULL`, "COMPARISON IS NULL")
			assertContains(t, sqlSelectQuery, `"u"."number" < ?`, "COMPARISON LESS")
			assertContains(t, sqlSelectQuery, `"u"."number" <= ?`, "COMPARISON LESSEQUAL")
			assertContains(t, sqlSelectQuery, `"u"."string" LIKE ?`, "COMPARISON LIKE")
			assertContains(t, sqlSelectQuery, `"u"."number" NOT BETWEEN ? AND ?`, "COMPARISON NOT BETWEEN")
			assertContains(t, sqlSelectQuery, `"u"."number" <> ?`, "COMPARISON NOT EQUAL")
			assertContains(t, sqlSelectQuery, `NOT EXISTS (SELECT 1 FROM "users" AS "u")`, "COMPARISON NOT EXISTS")
			assertContains(t, sqlSelectQuery, `LOWER("u"."string") NOT LIKE LOWER(?)`, "COMPARISON NOT ILIKE")
			assertContains(t, sqlSelectQuery, `"u"."string" NOT IN (?, ?)`, "COMPARISON NOT IN")
			assertContains(t, sqlSelectQuery, `"u"."string" NOT LIKE ?`, "COMPARISON NOT LIKE")
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				// Функции агрегатные
				Avg(Test.Table.Users.Number.Expr(), false).As("aggregate_avg"),
				BitAnd(Test.Table.Users.Number.Expr(), false).As("aggregate_bitand"),
				BitOr(Test.Table.Users.Number.Expr(), false).As("aggregate_bitor"),
				BitXor(Test.Table.Users.Number.Expr(), false).As("aggregate_bitxor"),
				Count(Test.Table.Users.String.Expr(), false).As("aggregate_count"),
				GroupConcat(Test.Table.Users.String.Expr(), false).As("aggregate_groupconcat"),
				Max(Test.Table.Users.Number.Expr(), false).As("aggregate_max"),
				Min(Test.Table.Users.Number.Expr(), false).As("aggregate_min"),
				StdDev(Test.Table.Users.Number.Expr(), false).As("aggregate_stddev"),
				Sum(Test.Table.Users.Number.Expr(), false).As("aggregate_sum"),
				Variance(Test.Table.Users.Number.Expr(), false).As("aggregate_variance"),
				// Функции аналитические
				FirstValue(Test.Table.Users.Name.Expr()).Over(
					PartitionBy(Test.Table.Users.ID.Expr()),
					OrderBy(Desc(Test.Table.Users.Number.Expr())),
				).As("analytical_firstvalue"),
				Lag(Test.Table.Users.Number.Expr(), 2).Over(
					PartitionBy(Test.Table.Users.ID.Expr()),
					OrderBy(Asc(Test.Table.Users.Date.Expr())),
				).As("analytical_lag"),
				LastValue(Test.Table.Users.Name.Expr()).Over(
					PartitionBy(Test.Table.Users.ID.Expr()),
					OrderBy(Asc(Test.Table.Users.Number.Expr())),
					RowsBetween("CURRENT ROW", "UNBOUNDED FOLLOWING"),
				).As("analytical_lastvalue"),
				Lead(Test.Table.Users.Number.Expr(), 2).Over(
					PartitionBy(Test.Table.Users.ID.Expr()),
					OrderBy(Asc(Test.Table.Users.Date.Expr())),
				).As("analytical_lead"),
				NthValue(Test.Table.Users.Name.Expr(), 2).Over(
					PartitionBy(Test.Table.Users.ID.Expr()),
					OrderBy(Desc(Test.Table.Users.Number.Expr())),
					RowsBetween("UNBOUNDED PRECEDING", "CURRENT ROW"),
				).As("analytical_nthvalue"),
				// Функции условий
				Case(CaseIf(CasePair(Less(Test.Table.Users.Number.Expr(), Value(2)), Value("old"))), CaseElse(Value("new"))).As("condition_case"),
				Coalesce(Test.Table.Users.CreateAt.Expr(), Test.Table.Users.UpdateAt.Expr()).As("condition_coalesce"),
				Greatest(Test.Table.Users.CreateAt.Expr(), Test.Table.Users.UpdateAt.Expr()).As("condition_greauser"),
				Least(Test.Table.Users.CreateAt.Expr(), Test.Table.Users.UpdateAt.Expr()).As("condition_least"),
				NullIf(Test.Table.Users.CreateAt.Expr(), Test.Table.Users.UpdateAt.Expr()).As("condition_if"),
				// Функции конвертации
				Cast(Test.Table.Users.Number.Expr(), TypeString).As("convert_cast"),
				CharLength(Test.Table.Users.String.Expr()).As("convert_charlength"),
				DateFormat(Test.Table.Users.CreateAt.Expr(), Literal("%Y-%m-%d")).As("convert_dateformat"),
				Degrees(Test.Table.Users.Number.Expr()).As("convert_degrees"),
				Length(Test.Table.Users.String.Expr()).As("convert_length"),
				Position(Test.Table.Users.String.Expr(), Value("old")).As("convert_position"),
				Radians(Test.Table.Users.Number.Expr()).As("convert_radians"),
				// Функции даты и времени
				CurDate().As("datetime_curdate"),
				CurTime().As("datetime_curtime"),
				DateAdd(Test.Table.Users.CreateAt.Expr(), Literal("2 DAY")).As("datetime_dateadd"),
				DateDiff(Test.Table.Users.UpdateAt.Expr(), Test.Table.Users.CreateAt.Expr()).As("datetime_datediff"),
				DateSub(Test.Table.Users.CreateAt.Expr(), Literal("2 DAY")).As("datetime_datesub"),
				Day(Test.Table.Users.CreateAt.Expr()).As("datetime_day"),
				DayName(Test.Table.Users.CreateAt.Expr()).As("datetime_dayname"),
				Hour(Test.Table.Users.CreateAt.Expr()).As("datetime_hour"),
				Minute(Test.Table.Users.CreateAt.Expr()).As("datetime_minute"),
				Month(Test.Table.Users.CreateAt.Expr()).As("datetime_month"),
				MonthName(Test.Table.Users.CreateAt.Expr()).As("datetime_monthname"),
				Now().As("datetime_now"),
				Quarter(Test.Table.Users.CreateAt.Expr()).As("datetime_quarter"),
				Second(Test.Table.Users.CreateAt.Expr()).As("datetime_second"),
				TimeAdd(Test.Table.Users.CreateAt.Expr(), Literal("2 HOUR")).As("datetime_timeadd"),
				TimeDiff(Test.Table.Users.UpdateAt.Expr(), Test.Table.Users.CreateAt.Expr()).As("datetime_timediff"),
				TimeSub(Test.Table.Users.CreateAt.Expr(), Literal("2 HOUR")).As("datetime_timesub"),
				Week(Test.Table.Users.CreateAt.Expr()).As("datetime_week"),
				Year(Test.Table.Users.CreateAt.Expr()).As("datetime_year"),
				// Функции обмена данными
				JsonArray(Test.Table.Users.Json.Expr(), Value("val1"), Value("val2")).As("json_jsonarray"),
				JsonArrayAgg(Test.Table.Users.Json.Expr()).As("json_jsonarrayagg"),
				JsonContains(Test.Table.Users.Json.Expr(), Value(`{"key":"val"}`)).As("json_jsoncontains"),
				JsonExtract(Test.Table.Users.Json.Expr(), JsonGroup(JsonPath(JsonKey("parent"), JsonIndex(0), JsonKey("child"))), TypeString).As("json_jsonextract"),
				JsonObject(JsonPair(JsonKey("key"), Count(Test.Table.Users.Json.Expr(), false))).As("json_jsonobject"),
				JsonObjectAgg(Test.Table.Users.Json.Expr(), Test.Table.Users.Number.Expr()).As("json_jsonobjectagg"),
				JsonRemove(Test.Table.Users.Json.Expr(), JsonGroup(JsonPath(JsonKey("key1"))), JsonGroup(JsonPath(JsonKey("key2")))).As("json_jsonremove"),
				JsonSet(Test.Table.Users.Json.Expr(), JsonGroup(JsonPath(JsonKey("key1")), Value("val1")), JsonGroup(JsonPath(JsonKey("key2")), Value("val2"))).As("json_jsonset"),
				JsonType(Test.Table.Users.Json.Expr()).As("json_jsontype"),
				// Функции математические
				Abs(Test.Table.Users.X.Expr()).As("math_abs"),
				ACos(Test.Table.Users.X.Expr()).As("math_acos"),
				ASin(Test.Table.Users.X.Expr()).As("math_asin"),
				ATan(Test.Table.Users.X.Expr()).As("math_atan"),
				ATan2(Test.Table.Users.Y.Expr(), Test.Table.Users.X.Expr()).As("math_atan2"),
				Cbrt(Test.Table.Users.X.Expr()).As("math_cbrt"),
				Ceil(Test.Table.Users.X.Expr()).As("math_ceil"),
				Cos(Test.Table.Users.X.Expr()).As("math_cos"),
				Exp(Test.Table.Users.X.Expr()).As("math_exp"),
				Floor(Test.Table.Users.X.Expr()).As("math_floor"),
				Ln(Test.Table.Users.X.Expr()).As("math_ln"),
				Log(Test.Table.Users.X.Expr(), Value(2)).As("math_log"),
				Mod(Test.Table.Users.X.Expr(), Value(2)).As("math_mod"),
				Pi().As("math_pi"),
				Power(Test.Table.Users.X.Expr(), Value(2)).As("math_power"),
				Rand().As("math_rand"),
				Round(Test.Table.Users.X.Expr(), Value(2)).As("math_round"),
				Sin(Test.Table.Users.X.Expr()).As("math_sin"),
				Sqrt(Test.Table.Users.X.Expr()).As("math_sqrt"),
				Tan(Test.Table.Users.X.Expr()).As("math_tan"),
				Trunc(Test.Table.Users.X.Expr(), Value(2)).As("math_trunc"),
				// Функции ранжирующие
				CumeDist().Over(
					PartitionBy(Test.Table.Users.ID.Expr()),
					OrderBy(Desc(Test.Table.Users.Number.Expr())),
				).As("ranking_cumedist"),
				DenseRank().Over(
					PartitionBy(Test.Table.Users.ID.Expr()),
					OrderBy(Desc(Test.Table.Users.Number.Expr())),
				).As("ranking_denserank"),
				NTile(2).Over(
					PartitionBy(Test.Table.Users.ID.Expr()),
					OrderBy(Desc(Test.Table.Users.Number.Expr())),
				).As("ranking_ntile"),
				PercentRank().Over(
					PartitionBy(Test.Table.Users.ID.Expr()),
					OrderBy(Desc(Test.Table.Users.Number.Expr())),
				).As("ranking_percentrank"),
				Rank().Over(
					PartitionBy(Test.Table.Users.ID.Expr()),
					OrderBy(Desc(Test.Table.Users.Number.Expr())),
				).As("ranking_rank"),
				RowNumber().Over(
					PartitionBy(Test.Table.Users.ID.Expr()),
					OrderBy(Desc(Test.Table.Users.Number.Expr())),
				).As("ranking_rownumber"),
				// Функции строковые
				Concat(Test.Table.Users.String.Expr(), Value("old"), Value("new")).As("string_concat"),
				ConcatWs(Value("_"), Test.Table.Users.String.Expr(), Value("old"), Value("new")).As("string_concatws"),
				LeftString(Test.Table.Users.String.Expr(), Value(2)).As("string_lstr"),
				Lower(Test.Table.Users.String.Expr()).As("string_lower"),
				LPad(Test.Table.Users.String.Expr(), Value(2), Value(",")).As("string_lpad"),
				LTrim(Test.Table.Users.String.Expr()).As("string_ltrim"),
				Repeat(Test.Table.Users.String.Expr(), Value(2)).As("string_repeat"),
				Replace(Test.Table.Users.String.Expr(), Value("old"), Value("new")).As("string_replace"),
				Reverse(Test.Table.Users.String.Expr()).As("string_reverse"),
				RightString(Test.Table.Users.String.Expr(), Value(2)).As("string_rstr"),
				RPad(Test.Table.Users.String.Expr(), Value(2), Value(",")).As("string_rpad"),
				RTrim(Test.Table.Users.String.Expr()).As("string_rtrim"),
				SubString(Test.Table.Users.String.Expr(), Value(0), Value(2)).As("string_substring"),
				Trim(Test.Table.Users.String.Expr()).As("string_trim"),
				Upper(Test.Table.Users.String.Expr()).As("string_upper"),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			// Функции агрегатные
			assertContains(t, sqlSelectQuery, "AVG(`u`.`number`)", "FUNCTION AVG")
			assertContains(t, sqlSelectQuery, "BIT_AND(`u`.`number`)", "FUNCTION BITAND")
			assertContains(t, sqlSelectQuery, "BIT_OR(`u`.`number`)", "FUNCTION BITOR")
			assertContains(t, sqlSelectQuery, "BIT_XOR(`u`.`number`)", "FUNCTION BITXOR")
			assertContains(t, sqlSelectQuery, "COUNT(`u`.`string`)", "FUNCTION COUNT")
			assertContains(t, sqlSelectQuery, "GROUP_CONCAT(`u`.`string` SEPARATOR ',')", "FUNCTION GROUPCONCAT")
			assertContains(t, sqlSelectQuery, "MAX(`u`.`number`)", "FUNCTION MAX")
			assertContains(t, sqlSelectQuery, "MIN(`u`.`number`)", "FUNCTION MIN")
			assertContains(t, sqlSelectQuery, "STDDEV(`u`.`number`)", "FUNCTION STDDEV")
			assertContains(t, sqlSelectQuery, "SUM(`u`.`number`)", "FUNCTION SUM")
			assertContains(t, sqlSelectQuery, "VARIANCE(`u`.`number`)", "FUNCTION VARIANCE")
			// Функции аналитические
			assertContains(t, sqlSelectQuery, "FIRST_VALUE(`u`.`name`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC", "FUNCTION FIRSTVALUE")
			assertContains(t, sqlSelectQuery, "LAG(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)", "FUNCTION LAG")
			assertContains(t, sqlSelectQuery, "LAST_VALUE(`u`.`name`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)", "FUNCTION LASTVALUE")
			assertContains(t, sqlSelectQuery, "LEAD(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)", "FUNCTION LEAD")
			assertContains(t, sqlSelectQuery, "NTH_VALUE(`u`.`name`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)", "FUNCTION NTHVALUE")
			// Функции условий
			assertContains(t, sqlSelectQuery, "CASE WHEN `u`.`number` < ? THEN ? ELSE ? END", "FUNCTION CASE")
			assertContains(t, sqlSelectQuery, "COALESCE(`u`.`createat`, `u`.`updateat`)", "FUNCTION COALESCE")
			assertContains(t, sqlSelectQuery, "GREATEST(`u`.`createat`, `u`.`updateat`)", "FUNCTION GREATEST")
			assertContains(t, sqlSelectQuery, "LEAST(`u`.`createat`, `u`.`updateat`)", "FUNCTION LEAST")
			assertContains(t, sqlSelectQuery, "NULLIF(`u`.`createat`, `u`.`updateat`)", "FUNCTION NULLIF")
			// Функции конвертации
			assertContains(t, sqlSelectQuery, "CAST(`u`.`number` AS VARCHAR)", "FUNCTION CAST")
			assertContains(t, sqlSelectQuery, "CHAR_LENGTH(`u`.`string`)", "FUNCTION CHARLENGTH")
			assertContains(t, sqlSelectQuery, "DATE_FORMAT(`u`.`createat`, '%Y-%m-%d')", "FUNCTION DATEFORMAT")
			assertContains(t, sqlSelectQuery, "DEGREES(`u`.`number`)", "FUNCTION DEGREES")
			assertContains(t, sqlSelectQuery, "LENGTH(`u`.`string`)", "FUNCTION LENGTH")
			assertContains(t, sqlSelectQuery, "POSITION(? IN `u`.`string`)", "FUNCTION POSITION")
			assertContains(t, sqlSelectQuery, "RADIANS(`u`.`number`)", "FUNCTION RADIANS")
			// Функции даты и времени
			assertContains(t, sqlSelectQuery, "CURDATE()", "FUNCTION CURDATE")
			assertContains(t, sqlSelectQuery, "CURTIME()", "FUNCTION CURTIME")
			assertContains(t, sqlSelectQuery, "DATE_ADD(`u`.`createat`, INTERVAL '2 DAY')", "FUNCTION DATEADD")
			assertContains(t, sqlSelectQuery, "DATEDIFF(`u`.`updateat`, `u`.`createat`)", "FUNCTION DATEDIFF")
			assertContains(t, sqlSelectQuery, "DATE_SUB(`u`.`createat`, INTERVAL '2 DAY')", "FUNCTION DATESUB")
			assertContains(t, sqlSelectQuery, "DAY(`u`.`createat`)", "FUNCTION DAY")
			assertContains(t, sqlSelectQuery, "DAYNAME(`u`.`createat`)", "FUNCTION DAYNAME")
			assertContains(t, sqlSelectQuery, "HOUR(`u`.`createat`)", "FUNCTION HOUR")
			assertContains(t, sqlSelectQuery, "MINUTE(`u`.`createat`)", "FUNCTION MINUTE")
			assertContains(t, sqlSelectQuery, "MONTH(`u`.`createat`)", "FUNCTION MONTH")
			assertContains(t, sqlSelectQuery, "MONTHNAME(`u`.`createat`)", "FUNCTION MONTHNAME")
			assertContains(t, sqlSelectQuery, "NOW()", "FUNCTION NOW")
			assertContains(t, sqlSelectQuery, "QUARTER(`u`.`createat`)", "FUNCTION QUARTER")
			assertContains(t, sqlSelectQuery, "SECOND(`u`.`createat`)", "FUNCTION SECOND")
			assertContains(t, sqlSelectQuery, "TIME_ADD(`u`.`createat`, INTERVAL '2 HOUR')", "FUNCTION TIMEADD")
			assertContains(t, sqlSelectQuery, "TIMEDIFF(`u`.`updateat`, `u`.`createat`)", "FUNCTION TIMEDIFF")
			assertContains(t, sqlSelectQuery, "TIME_SUB(`u`.`createat`, INTERVAL '2 HOUR')", "FUNCTION TIMESUB")
			assertContains(t, sqlSelectQuery, "WEEK(`u`.`createat`)", "FUNCTION WEEK")
			assertContains(t, sqlSelectQuery, "YEAR(`u`.`createat`)", "FUNCTION YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, "JSON_ARRAY(`u`.`json`, ?, ?)", "FUNCTION JSONARRAY")
			assertContains(t, sqlSelectQuery, "JSON_ARRAYAGG(`u`.`json`)", "FUNCTION JSONARRAYAGG")
			assertContains(t, sqlSelectQuery, "JSON_CONTAINS(`u`.`json`, ?)", "FUNCTION JSONCONTAINS")
			assertContains(t, sqlSelectQuery, "(`u`.`json` ->> '$.parent[0].child')", "FUNCTION JSONEXTRACT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECT('key', COUNT(`u`.`json`))", "FUNCTION JSONOBJECT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECTAGG(`u`.`json`, `u`.`number`)", "FUNCTION JSONOBJECTAGG")
			assertContains(t, sqlSelectQuery, "JSON_REMOVE(`u`.`json`, '$.key1', '$.key2')", "FUNCTION JSONREMOVE")
			assertContains(t, sqlSelectQuery, "JSON_SET(`u`.`json`, '$.key1', ?, '$.key2', ?)", "FUNCTION JSONSET")
			assertContains(t, sqlSelectQuery, "JSON_TYPE(`u`.`json`)", "FUNCTION JSONTYPE")
			// Функции математические
			assertContains(t, sqlSelectQuery, "ABS(`u`.`x`)", "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, "ACOS(`u`.`x`)", "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, "ASIN(`u`.`x`)", "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, "ATAN(`u`.`x`)", "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, "ATAN2(`u`.`y`, `u`.`x`)", "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, "CBRT(`u`.`x`)", "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, "CEILING(`u`.`x`)", "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, "COS(`u`.`x`)", "FUNCTION COS")
			assertContains(t, sqlSelectQuery, "EXP(`u`.`x`)", "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, "FLOOR(`u`.`x`)", "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, "LN(`u`.`x`)", "FUNCTION LN")
			assertContains(t, sqlSelectQuery, "LOG(`u`.`x`, ?)", "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, "MOD(`u`.`x`, ?)", "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, "PI()", "FUNCTION PI")
			assertContains(t, sqlSelectQuery, "POWER(`u`.`x`, ?)", "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, "RAND()", "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, "ROUND(`u`.`x`, ?)", "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, "SIN(`u`.`x`)", "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, "SQRT(`u`.`x`)", "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, "TAN(`u`.`x`)", "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, "TRUNCATE(`u`.`x`, ?)", "FUNCTION TRUNC")
			// Функции ранжирующие
			assertContains(t, sqlSelectQuery, "CUME_DIST() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION CUMEDIST")
			assertContains(t, sqlSelectQuery, "DENSE_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION DENSERANK")
			assertContains(t, sqlSelectQuery, "NTILE(2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION NTILE")
			assertContains(t, sqlSelectQuery, "PERCENT_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION PERCENTRANK")
			assertContains(t, sqlSelectQuery, "RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION RANK")
			assertContains(t, sqlSelectQuery, "ROW_NUMBER() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION ROWNUMBER")
			// Функции строковые
			assertContains(t, sqlSelectQuery, "CONCAT(`u`.`string`, ?, ?)", "FUNCTION CONCAT")
			assertContains(t, sqlSelectQuery, "CONCAT_WS(?, `u`.`string`, ?, ?)", "FUNCTION CONCATWS")
			assertContains(t, sqlSelectQuery, "LEFT(`u`.`string`, ?)", "FUNCTION LEFTSTRING")
			assertContains(t, sqlSelectQuery, "LOWER(`u`.`string`)", "FUNCTION LOWER")
			assertContains(t, sqlSelectQuery, "LPAD(`u`.`string`, ?, ?)", "FUNCTION LPAD")
			assertContains(t, sqlSelectQuery, "LTRIM(`u`.`string`)", "FUNCTION LTRIM")
			assertContains(t, sqlSelectQuery, "REPEAT(`u`.`string`, ?)", "FUNCTION REPEAT")
			assertContains(t, sqlSelectQuery, "REPLACE(`u`.`string`, ?, ?)", "FUNCTION REPLACE")
			assertContains(t, sqlSelectQuery, "REVERSE(`u`.`string`)", "FUNCTION REVERSE")
			assertContains(t, sqlSelectQuery, "RIGHT(`u`.`string`, ?)", "FUNCTION RIGHTSTRING")
			assertContains(t, sqlSelectQuery, "RPAD(`u`.`string`, ?, ?)", "FUNCTION RPAD")
			assertContains(t, sqlSelectQuery, "RTRIM(`u`.`string`)", "FUNCTION RTRIM")
			assertContains(t, sqlSelectQuery, "SUBSTRING(`u`.`string`, ?, ?)", "FUNCTION SUBSTRING")
			assertContains(t, sqlSelectQuery, "TRIM(`u`.`string`)", "FUNCTION TRIM")
			assertContains(t, sqlSelectQuery, "UPPER(`u`.`string`)", "FUNCTION UPPER")
		case DialectMsSQL:
			// Функции агрегатные
			assertContains(t, sqlSelectQuery, "AVG([u].[number])", "FUNCTION AVG")
			assertContains(t, sqlSelectQuery, "BIT_AND([u].[number])", "FUNCTION BITAND")
			assertContains(t, sqlSelectQuery, "BIT_OR([u].[number])", "FUNCTION BITOR")
			assertContains(t, sqlSelectQuery, "BIT_XOR([u].[number])", "FUNCTION BITXOR")
			assertContains(t, sqlSelectQuery, "COUNT([u].[string])", "FUNCTION COUNT")
			assertContains(t, sqlSelectQuery, "STRING_AGG([u].[string], ',')", "FUNCTION GROUPCONCAT")
			assertContains(t, sqlSelectQuery, "MAX([u].[number])", "FUNCTION MAX")
			assertContains(t, sqlSelectQuery, "MIN([u].[number])", "FUNCTION MIN")
			assertContains(t, sqlSelectQuery, "STDEV([u].[number])", "FUNCTION STDDEV")
			assertContains(t, sqlSelectQuery, "SUM([u].[number])", "FUNCTION SUM")
			assertContains(t, sqlSelectQuery, "VAR([u].[number])", "FUNCTION VARIANCE")
			// Функции аналитические
			assertContains(t, sqlSelectQuery, "FIRST_VALUE([u].[name]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC", "FUNCTION FIRSTVALUE")
			assertContains(t, sqlSelectQuery, "LAG([u].[number], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[date] ASC)", "FUNCTION LAG")
			assertContains(t, sqlSelectQuery, "LAST_VALUE([u].[name]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)", "FUNCTION LASTVALUE")
			assertContains(t, sqlSelectQuery, "LEAD([u].[number], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[date] ASC)", "FUNCTION LEAD")
			assertContains(t, sqlSelectQuery, "NTH_VALUE([u].[name], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)", "FUNCTION NTHVALUE")
			// Функции условий
			assertContains(t, sqlSelectQuery, "CASE WHEN [u].[number] < @p1 THEN @p2 ELSE @p3 END", "FUNCTION CASE")
			assertContains(t, sqlSelectQuery, "COALESCE([u].[createat], [u].[updateat])", "FUNCTION COALESCE")
			assertContains(t, sqlSelectQuery, "GREATEST([u].[createat], [u].[updateat])", "FUNCTION GREATEST")
			assertContains(t, sqlSelectQuery, "LEAST([u].[createat], [u].[updateat])", "FUNCTION LEAST")
			assertContains(t, sqlSelectQuery, "NULLIF([u].[createat], [u].[updateat])", "FUNCTION NULLIF")
			// Функции конвертации
			assertContains(t, sqlSelectQuery, "CAST([u].[number] AS NVARCHAR)", "FUNCTION CAST")
			assertContains(t, sqlSelectQuery, "CHAR_LENGTH([u].[string])", "FUNCTION CHARLENGTH")
			assertContains(t, sqlSelectQuery, "FORMAT([u].[createat], '%Y-%m-%d')", "FUNCTION DATEFORMAT")
			assertContains(t, sqlSelectQuery, "DEGREES([u].[number])", "FUNCTION DEGREES")
			assertContains(t, sqlSelectQuery, "LEN([u].[string])", "FUNCTION LENGTH")
			assertContains(t, sqlSelectQuery, "CHARINDEX(@p1, [u].[string])", "FUNCTION POSITION")
			assertContains(t, sqlSelectQuery, "RADIANS([u].[number])", "FUNCTION RADIANS")
			// Функции даты и времени
			assertContains(t, sqlSelectQuery, "CAST(GETDATE() AS DATE)", "FUNCTION CURDATE")
			assertContains(t, sqlSelectQuery, "CAST(GETDATE() AS TIME)", "FUNCTION CURTIME")
			assertContains(t, sqlSelectQuery, "DATEADD(DAY, 2, [u].[createat])", "FUNCTION DATEADD")
			assertContains(t, sqlSelectQuery, "DATEDIFF([u].[updateat], [u].[createat])", "FUNCTION DATEDIFF")
			assertContains(t, sqlSelectQuery, "DATEADD(DAY, -2, [u].[createat])", "FUNCTION DATESUB")
			assertContains(t, sqlSelectQuery, "DAY([u].[createat])", "FUNCTION DAY")
			assertContains(t, sqlSelectQuery, "DATENAME(WEEKDAY, [u].[createat])", "FUNCTION DAYNAME")
			assertContains(t, sqlSelectQuery, "DATEPART(HOUR, [u].[createat])", "FUNCTION HOUR")
			assertContains(t, sqlSelectQuery, "DATEPART(MINUTE, [u].[createat])", "FUNCTION MINUTE")
			assertContains(t, sqlSelectQuery, "MONTH([u].[createat])", "FUNCTION MONTH")
			assertContains(t, sqlSelectQuery, "DATENAME(MONTH, [u].[createat])", "FUNCTION MONTHNAME")
			assertContains(t, sqlSelectQuery, "GETDATE()", "FUNCTION NOW")
			assertContains(t, sqlSelectQuery, "DATEPART(QUARTER, [u].[createat])", "FUNCTION QUARTER")
			assertContains(t, sqlSelectQuery, "DATEPART(SECOND, [u].[createat])", "FUNCTION SECOND")
			assertContains(t, sqlSelectQuery, "DATEADD(HOUR, 2, [u].[createat])", "FUNCTION TIMEADD")
			assertContains(t, sqlSelectQuery, "TIMEDIFF([u].[updateat], [u].[createat])", "FUNCTION TIMEDIFF")
			assertContains(t, sqlSelectQuery, "DATEADD(HOUR, -2, [u].[createat])", "FUNCTION TIMESUB")
			assertContains(t, sqlSelectQuery, "DATEPART(WEEK, [u].[createat])", "FUNCTION WEEK")
			assertContains(t, sqlSelectQuery, "YEAR([u].[createat])", "FUNCTION YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, "JSON_ARRAY([u].[json], @p1, @p2)", "FUNCTION JSONARRAY")
			assertContains(t, sqlSelectQuery, "JSON_ARRAYAGG([u].[json])", "FUNCTION JSONARRAYAGG")
			// Not supported - FUNCTION JSONCONTAINS
			assertContains(t, sqlSelectQuery, "JSON_VALUE([u].[json], '$.parent[0].child')", "FUNCTION JSONEXTRACT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECT('key', COUNT([u].[json]))", "FUNCTION JSONOBJECT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECTAGG([u].[json], [u].[number])", "FUNCTION JSONOBJECTAGG")
			assertContains(t, sqlSelectQuery, "JSON_MODIFY(JSON_MODIFY([u].[json], '$.key1', NULL), '$.key2', NULL)", "FUNCTION JSONREMOVE")
			assertContains(t, sqlSelectQuery, "JSON_MODIFY(JSON_MODIFY([u].[json], '$.key1', @p1), '$.key2', @p2)", "FUNCTION JSONSET")
			// Not supported - FUNCTION JSONTYPE
			// Функции математические
			assertContains(t, sqlSelectQuery, "ABS([u].[x])", "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, "ACOS([u].[x])", "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, "ASIN([u].[x])", "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, "ATAN([u].[x])", "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, "ATAN2([u].[y], [u].[x])", "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, "CBRT([u].[x])", "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, "CEILING([u].[x])", "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, "COS([u].[x])", "FUNCTION COS")
			assertContains(t, sqlSelectQuery, "EXP([u].[x])", "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, "FLOOR([u].[x])", "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, "LN([u].[x])", "FUNCTION LN")
			assertContains(t, sqlSelectQuery, "LOG([u].[x], @p1)", "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, "MOD([u].[x], @p1)", "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, "PI()", "FUNCTION PI")
			assertContains(t, sqlSelectQuery, "POWER([u].[x], @p1)", "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, "RAND()", "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, "ROUND([u].[x], @p1)", "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, "SIN([u].[x])", "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, "SQRT([u].[x])", "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, "TAN([u].[x])", "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, "ROUND([u].[x], @p1, 1)", "FUNCTION TRUNC")
			// Функции ранжирующие
			assertContains(t, sqlSelectQuery, "CUME_DIST() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)", "FUNCTION CUMEDIST")
			assertContains(t, sqlSelectQuery, "DENSE_RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)", "FUNCTION DENSERANK")
			assertContains(t, sqlSelectQuery, "NTILE(2) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)", "FUNCTION NTILE")
			assertContains(t, sqlSelectQuery, "PERCENT_RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)", "FUNCTION PERCENTRANK")
			assertContains(t, sqlSelectQuery, "RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)", "FUNCTION RANK")
			assertContains(t, sqlSelectQuery, "ROW_NUMBER() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)", "FUNCTION ROWNUMBER")
			// Функции строковые
			assertContains(t, sqlSelectQuery, "CONCAT([u].[string], @p1, @p2)", "FUNCTION CONCAT")
			assertContains(t, sqlSelectQuery, "CONCAT_WS(@p1, [u].[string], @p2, @p3)", "FUNCTION CONCATWS")
			assertContains(t, sqlSelectQuery, "LEFT([u].[string], @p1)", "FUNCTION LEFTSTRING")
			assertContains(t, sqlSelectQuery, "LOWER([u].[string])", "FUNCTION LOWER")
			assertContains(t, sqlSelectQuery, "LPAD([u].[string], @p1, @p2)", "FUNCTION LPAD")
			assertContains(t, sqlSelectQuery, "LTRIM([u].[string])", "FUNCTION LTRIM")
			assertContains(t, sqlSelectQuery, "REPEAT([u].[string], @p1)", "FUNCTION REPEAT")
			assertContains(t, sqlSelectQuery, "REPLACE([u].[string], @p1, @p2)", "FUNCTION REPLACE")
			assertContains(t, sqlSelectQuery, "REVERSE([u].[string])", "FUNCTION REVERSE")
			assertContains(t, sqlSelectQuery, "RIGHT([u].[string], @p1)", "FUNCTION RIGHTSTRING")
			assertContains(t, sqlSelectQuery, "RPAD([u].[string], @p1, @p2)", "FUNCTION RPAD")
			assertContains(t, sqlSelectQuery, "RTRIM([u].[string])", "FUNCTION RTRIM")
			assertContains(t, sqlSelectQuery, "SUBSTRING([u].[string], @p1, @p2)", "FUNCTION SUBSTRING")
			assertContains(t, sqlSelectQuery, "TRIM([u].[string])", "FUNCTION TRIM")
			assertContains(t, sqlSelectQuery, "UPPER([u].[string])", "FUNCTION UPPER")
		case DialectMySQL:
			// Функции агрегатные
			assertContains(t, sqlSelectQuery, "AVG(`u`.`number`)", "FUNCTION AVG")
			assertContains(t, sqlSelectQuery, "BIT_AND(`u`.`number`)", "FUNCTION BITAND")
			assertContains(t, sqlSelectQuery, "BIT_OR(`u`.`number`)", "FUNCTION BITOR")
			assertContains(t, sqlSelectQuery, "BIT_XOR(`u`.`number`)", "FUNCTION BITXOR")
			assertContains(t, sqlSelectQuery, "COUNT(`u`.`string`)", "FUNCTION COUNT")
			assertContains(t, sqlSelectQuery, "GROUP_CONCAT(`u`.`string` SEPARATOR ',')", "FUNCTION GROUPCONCAT")
			assertContains(t, sqlSelectQuery, "MAX(`u`.`number`)", "FUNCTION MAX")
			assertContains(t, sqlSelectQuery, "MIN(`u`.`number`)", "FUNCTION MIN")
			assertContains(t, sqlSelectQuery, "STDDEV(`u`.`number`)", "FUNCTION STDDEV")
			assertContains(t, sqlSelectQuery, "SUM(`u`.`number`)", "FUNCTION SUM")
			assertContains(t, sqlSelectQuery, "VARIANCE(`u`.`number`)", "FUNCTION VARIANCE")
			// Функции аналитические
			assertContains(t, sqlSelectQuery, "FIRST_VALUE(`u`.`name`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC", "FUNCTION FIRSTVALUE")
			assertContains(t, sqlSelectQuery, "LAG(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)", "FUNCTION LAG")
			assertContains(t, sqlSelectQuery, "LAST_VALUE(`u`.`name`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)", "FUNCTION LASTVALUE")
			assertContains(t, sqlSelectQuery, "LEAD(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)", "FUNCTION LEAD")
			assertContains(t, sqlSelectQuery, "NTH_VALUE(`u`.`name`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)", "FUNCTION NTHVALUE")
			// Функции условий
			assertContains(t, sqlSelectQuery, "CASE WHEN `u`.`number` < ? THEN ? ELSE ? END", "FUNCTION CASE")
			assertContains(t, sqlSelectQuery, "COALESCE(`u`.`createat`, `u`.`updateat`)", "FUNCTION COALESCE")
			assertContains(t, sqlSelectQuery, "GREATEST(`u`.`createat`, `u`.`updateat`)", "FUNCTION GREATEST")
			assertContains(t, sqlSelectQuery, "LEAST(`u`.`createat`, `u`.`updateat`)", "FUNCTION LEAST")
			assertContains(t, sqlSelectQuery, "NULLIF(`u`.`createat`, `u`.`updateat`)", "FUNCTION NULLIF")
			// Функции конвертации
			assertContains(t, sqlSelectQuery, "CAST(`u`.`number` AS VARCHAR)", "FUNCTION CAST")
			assertContains(t, sqlSelectQuery, "CHAR_LENGTH(`u`.`string`)", "FUNCTION CHARLENGTH")
			assertContains(t, sqlSelectQuery, "DATE_FORMAT(`u`.`createat`, '%Y-%m-%d')", "FUNCTION DATEFORMAT")
			assertContains(t, sqlSelectQuery, "DEGREES(`u`.`number`)", "FUNCTION DEGREES")
			assertContains(t, sqlSelectQuery, "LENGTH(`u`.`string`)", "FUNCTION LENGTH")
			assertContains(t, sqlSelectQuery, "POSITION(? IN `u`.`string`)", "FUNCTION POSITION")
			assertContains(t, sqlSelectQuery, "RADIANS(`u`.`number`)", "FUNCTION RADIANS")
			// Функции даты и времени
			assertContains(t, sqlSelectQuery, "CURDATE()", "FUNCTION CURDATE")
			assertContains(t, sqlSelectQuery, "CURTIME()", "FUNCTION CURTIME")
			assertContains(t, sqlSelectQuery, "DATE_ADD(`u`.`createat`, INTERVAL '2 DAY')", "FUNCTION DATEADD")
			assertContains(t, sqlSelectQuery, "DATEDIFF(`u`.`updateat`, `u`.`createat`)", "FUNCTION DATEDIFF")
			assertContains(t, sqlSelectQuery, "DATE_SUB(`u`.`createat`, INTERVAL '2 DAY')", "FUNCTION DATESUB")
			assertContains(t, sqlSelectQuery, "DAY(`u`.`createat`)", "FUNCTION DAY")
			assertContains(t, sqlSelectQuery, "DAYNAME(`u`.`createat`)", "FUNCTION DAYNAME")
			assertContains(t, sqlSelectQuery, "HOUR(`u`.`createat`)", "FUNCTION HOUR")
			assertContains(t, sqlSelectQuery, "MINUTE(`u`.`createat`)", "FUNCTION MINUTE")
			assertContains(t, sqlSelectQuery, "MONTH(`u`.`createat`)", "FUNCTION MONTH")
			assertContains(t, sqlSelectQuery, "MONTHNAME(`u`.`createat`)", "FUNCTION MONTHNAME")
			assertContains(t, sqlSelectQuery, "NOW()", "FUNCTION NOW")
			assertContains(t, sqlSelectQuery, "QUARTER(`u`.`createat`)", "FUNCTION QUARTER")
			assertContains(t, sqlSelectQuery, "SECOND(`u`.`createat`)", "FUNCTION SECOND")
			assertContains(t, sqlSelectQuery, "TIME_ADD(`u`.`createat`, INTERVAL '2 HOUR')", "FUNCTION TIMEADD")
			assertContains(t, sqlSelectQuery, "TIMEDIFF(`u`.`updateat`, `u`.`createat`)", "FUNCTION TIMEDIFF")
			assertContains(t, sqlSelectQuery, "TIME_SUB(`u`.`createat`, INTERVAL '2 HOUR')", "FUNCTION TIMESUB")
			assertContains(t, sqlSelectQuery, "WEEK(`u`.`createat`)", "FUNCTION WEEK")
			assertContains(t, sqlSelectQuery, "YEAR(`u`.`createat`)", "FUNCTION YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, "JSON_ARRAY(`u`.`json`, ?, ?)", "FUNCTION JSONARRAY")
			assertContains(t, sqlSelectQuery, "JSON_ARRAYAGG(`u`.`json`)", "FUNCTION JSONARRAYAGG")
			assertContains(t, sqlSelectQuery, "JSON_CONTAINS(`u`.`json`, ?)", "FUNCTION JSONCONTAINS")
			assertContains(t, sqlSelectQuery, "(`u`.`json` ->> '$.parent[0].child')", "FUNCTION JSONEXTRACT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECT('key', COUNT(`u`.`json`))", "FUNCTION JSONOBJECT")
			assertContains(t, sqlSelectQuery, "JSON_OBJECTAGG(`u`.`json`, `u`.`number`)", "FUNCTION JSONOBJECTAGG")
			assertContains(t, sqlSelectQuery, "JSON_REMOVE(`u`.`json`, '$.key1', '$.key2')", "FUNCTION JSONREMOVE")
			assertContains(t, sqlSelectQuery, "JSON_SET(`u`.`json`, '$.key1', ?, '$.key2', ?)", "FUNCTION JSONSET")
			assertContains(t, sqlSelectQuery, "JSON_TYPE(`u`.`json`)", "FUNCTION JSONTYPE")
			// Функции математические
			assertContains(t, sqlSelectQuery, "ABS(`u`.`x`)", "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, "ACOS(`u`.`x`)", "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, "ASIN(`u`.`x`)", "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, "ATAN(`u`.`x`)", "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, "ATAN2(`u`.`y`, `u`.`x`)", "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, "CBRT(`u`.`x`)", "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, "CEILING(`u`.`x`)", "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, "COS(`u`.`x`)", "FUNCTION COS")
			assertContains(t, sqlSelectQuery, "EXP(`u`.`x`)", "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, "FLOOR(`u`.`x`)", "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, "LN(`u`.`x`)", "FUNCTION LN")
			assertContains(t, sqlSelectQuery, "LOG(`u`.`x`, ?)", "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, "MOD(`u`.`x`, ?)", "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, "PI()", "FUNCTION PI")
			assertContains(t, sqlSelectQuery, "POWER(`u`.`x`, ?)", "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, "RAND()", "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, "ROUND(`u`.`x`, ?)", "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, "SIN(`u`.`x`)", "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, "SQRT(`u`.`x`)", "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, "TAN(`u`.`x`)", "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, "TRUNCATE(`u`.`x`, ?)", "FUNCTION TRUNC")
			// Функции ранжирующие
			assertContains(t, sqlSelectQuery, "CUME_DIST() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION CUMEDIST")
			assertContains(t, sqlSelectQuery, "DENSE_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION DENSERANK")
			assertContains(t, sqlSelectQuery, "NTILE(2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION NTILE")
			assertContains(t, sqlSelectQuery, "PERCENT_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION PERCENTRANK")
			assertContains(t, sqlSelectQuery, "RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION RANK")
			assertContains(t, sqlSelectQuery, "ROW_NUMBER() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)", "FUNCTION ROWNUMBER")
			// Функции строковые
			assertContains(t, sqlSelectQuery, "CONCAT(`u`.`string`, ?, ?)", "FUNCTION CONCAT")
			assertContains(t, sqlSelectQuery, "CONCAT_WS(?, `u`.`string`, ?, ?)", "FUNCTION CONCATWS")
			assertContains(t, sqlSelectQuery, "LEFT(`u`.`string`, ?)", "FUNCTION LEFTSTRING")
			assertContains(t, sqlSelectQuery, "LOWER(`u`.`string`)", "FUNCTION LOWER")
			assertContains(t, sqlSelectQuery, "LPAD(`u`.`string`, ?, ?)", "FUNCTION LPAD")
			assertContains(t, sqlSelectQuery, "LTRIM(`u`.`string`)", "FUNCTION LTRIM")
			assertContains(t, sqlSelectQuery, "REPEAT(`u`.`string`, ?)", "FUNCTION REPEAT")
			assertContains(t, sqlSelectQuery, "REPLACE(`u`.`string`, ?, ?)", "FUNCTION REPLACE")
			assertContains(t, sqlSelectQuery, "REVERSE(`u`.`string`)", "FUNCTION REVERSE")
			assertContains(t, sqlSelectQuery, "RIGHT(`u`.`string`, ?)", "FUNCTION RIGHTSTRING")
			assertContains(t, sqlSelectQuery, "RPAD(`u`.`string`, ?, ?)", "FUNCTION RPAD")
			assertContains(t, sqlSelectQuery, "RTRIM(`u`.`string`)", "FUNCTION RTRIM")
			assertContains(t, sqlSelectQuery, "SUBSTRING(`u`.`string`, ?, ?)", "FUNCTION SUBSTRING")
			assertContains(t, sqlSelectQuery, "TRIM(`u`.`string`)", "FUNCTION TRIM")
			assertContains(t, sqlSelectQuery, "UPPER(`u`.`string`)", "FUNCTION UPPER")
		case DialectPostgreSQL:
			// Функции агрегатные
			assertContains(t, sqlSelectQuery, `AVG("u"."number")`, "FUNCTION AVG")
			assertContains(t, sqlSelectQuery, `BIT_AND("u"."number")`, "FUNCTION BITAND")
			assertContains(t, sqlSelectQuery, `BIT_OR("u"."number")`, "FUNCTION BITOR")
			assertContains(t, sqlSelectQuery, `BIT_XOR("u"."number")`, "FUNCTION BITXOR")
			assertContains(t, sqlSelectQuery, `COUNT("u"."string")`, "FUNCTION COUNT")
			assertContains(t, sqlSelectQuery, `STRING_AGG("u"."string", ',')`, "FUNCTION GROUPCONCAT")
			assertContains(t, sqlSelectQuery, `MAX("u"."number")`, "FUNCTION MAX")
			assertContains(t, sqlSelectQuery, `MIN("u"."number")`, "FUNCTION MIN")
			assertContains(t, sqlSelectQuery, `STDDEV_SAMP("u"."number")`, "FUNCTION STDDEV")
			assertContains(t, sqlSelectQuery, `SUM("u"."number")`, "FUNCTION SUM")
			assertContains(t, sqlSelectQuery, `VAR_SAMP("u"."number")`, "FUNCTION VARIANCE")
			// Функции аналитические
			assertContains(t, sqlSelectQuery, `FIRST_VALUE("u"."name") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION FIRSTVALUE")
			assertContains(t, sqlSelectQuery, `LAG("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)`, "FUNCTION LAG")
			assertContains(t, sqlSelectQuery, `LAST_VALUE("u"."name") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)`, "FUNCTION LASTVALUE")
			assertContains(t, sqlSelectQuery, `LEAD("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)`, "FUNCTION LEAD")
			assertContains(t, sqlSelectQuery, `NTH_VALUE("u"."name", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)`, "FUNCTION NTHVALUE")
			// Функции условий
			assertContains(t, sqlSelectQuery, `CASE WHEN "u"."number" < $1 THEN $2 ELSE $3 END`, "FUNCTION CASE")
			assertContains(t, sqlSelectQuery, `COALESCE("u"."createat", "u"."updateat")`, "FUNCTION COALESCE")
			assertContains(t, sqlSelectQuery, `GREATEST("u"."createat", "u"."updateat")`, "FUNCTION GREATEST")
			assertContains(t, sqlSelectQuery, `LEAST("u"."createat", "u"."updateat")`, "FUNCTION LEAST")
			assertContains(t, sqlSelectQuery, `NULLIF("u"."createat", "u"."updateat")`, "FUNCTION NULLIF")
			// Функции конвертации
			assertContains(t, sqlSelectQuery, `CAST("u"."number" AS VARCHAR) `, "FUNCTION CAST")
			assertContains(t, sqlSelectQuery, `CHAR_LENGTH("u"."string")`, "FUNCTION CHARLENGTH")
			assertContains(t, sqlSelectQuery, `TO_CHAR("u"."createat", '%Y-%m-%d')`, "FUNCTION DATEFORMAT")
			assertContains(t, sqlSelectQuery, `DEGREES("u"."number")`, "FUNCTION DEGREES")
			assertContains(t, sqlSelectQuery, `LENGTH("u"."string")`, "FUNCTION LENGTH")
			assertContains(t, sqlSelectQuery, `POSITION($1 IN "u"."string")`, "FUNCTION POSITION")
			assertContains(t, sqlSelectQuery, `RADIANS("u"."number")`, "FUNCTION RADIANS")
			// Функции даты и времени
			assertContains(t, sqlSelectQuery, `CURRENT_DATE`, "FUNCTION CURDATE")
			assertContains(t, sqlSelectQuery, `CURRENT_TIME`, "FUNCTION CURTIME")
			assertContains(t, sqlSelectQuery, `("u"."createat" + INTERVAL '2 DAY')`, "FUNCTION DATEADD")
			assertContains(t, sqlSelectQuery, `DATE_PART('day', "u"."updateat" - "u"."createat")`, "FUNCTION DATEDIFF")
			assertContains(t, sqlSelectQuery, `("u"."createat" - INTERVAL '2 DAY')`, "FUNCTION DATESUB")
			assertContains(t, sqlSelectQuery, `EXTRACT(DAY FROM "u"."createat")`, "FUNCTION DAY")
			assertContains(t, sqlSelectQuery, `TO_CHAR("u"."createat", 'Day')`, "FUNCTION DAYNAME")
			assertContains(t, sqlSelectQuery, `EXTRACT(HOUR FROM "u"."createat")`, "FUNCTION HOUR")
			assertContains(t, sqlSelectQuery, `EXTRACT(MINUTE FROM "u"."createat")`, "FUNCTION MINUTE")
			assertContains(t, sqlSelectQuery, `EXTRACT(MONTH FROM "u"."createat")`, "FUNCTION MONTH")
			assertContains(t, sqlSelectQuery, `TO_CHAR("u"."createat", 'Month')`, "FUNCTION MONTHNAME")
			assertContains(t, sqlSelectQuery, `CURRENT_TIMESTAMP`, "FUNCTION NOW")
			assertContains(t, sqlSelectQuery, `EXTRACT(QUARTER FROM "u"."createat")`, "FUNCTION QUARTER")
			assertContains(t, sqlSelectQuery, `EXTRACT(SECOND FROM "u"."createat")`, "FUNCTION SECOND")
			assertContains(t, sqlSelectQuery, `("u"."createat" + INTERVAL '2 HOUR')`, "FUNCTION TIMEADD")
			assertContains(t, sqlSelectQuery, `DATE_PART('time', "u"."updateat" - "u"."createat")`, "FUNCTION TIMEDIFF")
			assertContains(t, sqlSelectQuery, `("u"."createat" - INTERVAL '2 HOUR')`, "FUNCTION TIMESUB")
			assertContains(t, sqlSelectQuery, `EXTRACT(WEEK FROM "u"."createat")`, "FUNCTION WEEK")
			assertContains(t, sqlSelectQuery, `EXTRACT(YEAR FROM "u"."createat")`, "FUNCTION YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, `JSON_ARRAY("u"."json", $1, $2)`, "FUNCTION JSONARRAY")
			assertContains(t, sqlSelectQuery, `JSON_AGG("u"."json")`, "FUNCTION JSONARRAYAGG")
			assertContains(t, sqlSelectQuery, `("u"."json" @> $1)`, "FUNCTION JSONCONTAINS")
			assertContains(t, sqlSelectQuery, `("u"."json" #>> '{parent,0,child}')`, "FUNCTION JSONEXTRACT")
			assertContains(t, sqlSelectQuery, `JSON_BUILD_OBJECT('key', COUNT("u"."json"))`, "FUNCTION JSONOBJECT")
			assertContains(t, sqlSelectQuery, `JSON_OBJECT_AGG("u"."json", "u"."number")`, "FUNCTION JSONOBJECTAGG")
			assertContains(t, sqlSelectQuery, `("u"."json" - '{key1}' - '{key2}')`, "FUNCTION JSONREMOVE")
			assertContains(t, sqlSelectQuery, `jsonb_set(jsonb_set("u"."json", '{key1}', $1), '{key2}', $2)`, "FUNCTION JSONSET")
			assertContains(t, sqlSelectQuery, `jsonb_typeof("u"."json")`, "FUNCTION JSONTYPE")
			// Функции математические
			assertContains(t, sqlSelectQuery, `ABS("u"."x")`, "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, `ACOS("u"."x")`, "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, `ASIN("u"."x")`, "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, `ATAN("u"."x")`, "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, `ATAN2("u"."y", "u"."x")`, "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, `CBRT("u"."x")`, "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, `CEIL("u"."x")`, "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, `COS("u"."x")`, "FUNCTION COS")
			assertContains(t, sqlSelectQuery, `EXP("u"."x")`, "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, `FLOOR("u"."x")`, "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, `LN("u"."x")`, "FUNCTION LN")
			assertContains(t, sqlSelectQuery, `LOG("u"."x", $1)`, "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, `MOD("u"."x", $1)`, "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, `PI()`, "FUNCTION PI")
			assertContains(t, sqlSelectQuery, `POWER("u"."x", $1)`, "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, `RANDOM`, "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, `ROUND("u"."x", $1)`, "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, `SIN("u"."x")`, "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, `SQRT("u"."x")`, "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, `TAN("u"."x")`, "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, `TRUNC("u"."x", $1)`, "FUNCTION TRUNC")
			// Функции ранжирующие
			assertContains(t, sqlSelectQuery, `CUME_DIST() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION CUMEDIST")
			assertContains(t, sqlSelectQuery, `DENSE_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "DFUNCTION ENSERANK")
			assertContains(t, sqlSelectQuery, `NTILE(2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION NTILE")
			assertContains(t, sqlSelectQuery, `PERCENT_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION PERCENTRANK")
			assertContains(t, sqlSelectQuery, `RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION RANK")
			assertContains(t, sqlSelectQuery, `ROW_NUMBER() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION ROWNUMBER")
			// Функции строковые
			assertContains(t, sqlSelectQuery, `CONCAT("u"."string", $1, $2)`, "FUNCTION CONCAT")
			assertContains(t, sqlSelectQuery, `CONCAT_WS($1, "u"."string", $2, $3)`, "FUNCTION CONCATWS")
			assertContains(t, sqlSelectQuery, `LEFT("u"."string", $1)`, "FUNCTION LEFTSTRING")
			assertContains(t, sqlSelectQuery, `LOWER("u"."string")`, "FUNCTION LOWER")
			assertContains(t, sqlSelectQuery, `LPAD("u"."string", $1, $2)`, "FUNCTION LPAD")
			assertContains(t, sqlSelectQuery, `LTRIM("u"."string")`, "FUNCTION LTRIM")
			assertContains(t, sqlSelectQuery, `REPEAT("u"."string", $1)`, "FUNCTION REPEAT")
			assertContains(t, sqlSelectQuery, `REPLACE("u"."string", $1, $2)`, "FUNCTION REPLACE")
			assertContains(t, sqlSelectQuery, `REVERSE("u"."string")`, "FUNCTION REVERSE")
			assertContains(t, sqlSelectQuery, `RIGHT("u"."string", $1)`, "FUNCTION RIGHTSTRING")
			assertContains(t, sqlSelectQuery, `RPAD("u"."string", $1, $2)`, "FUNCTION RPAD")
			assertContains(t, sqlSelectQuery, `RTRIM("u"."string")`, "FUNCTION RTRIM")
			assertContains(t, sqlSelectQuery, `SUBSTRING("u"."string", $1, $2)`, "FUNCTION SUBSTRING")
			assertContains(t, sqlSelectQuery, `TRIM("u"."string")`, "FUNCTION TRIM")
			assertContains(t, sqlSelectQuery, `UPPER("u"."string")`, "FUNCTION UPPER")
		case DialectSQLite:
			// Функции агрегатные
			assertContains(t, sqlSelectQuery, `AVG("u"."number")`, "FUNCTION AVG")
			assertContains(t, sqlSelectQuery, `BIT_AND("u"."number")`, "FUNCTION BITAND")
			assertContains(t, sqlSelectQuery, `BIT_OR("u"."number")`, "FUNCTION BITOR")
			assertContains(t, sqlSelectQuery, `BIT_XOR("u"."number")`, "FUNCTION BITXOR")
			assertContains(t, sqlSelectQuery, `COUNT("u"."string")`, "FUNCTION COUNT")
			assertContains(t, sqlSelectQuery, `GROUP_CONCAT("u"."string" SEPARATOR ',')`, "FUNCTION GROUPCONCAT")
			assertContains(t, sqlSelectQuery, `MAX("u"."number")`, "FUNCTION MAX")
			assertContains(t, sqlSelectQuery, `MIN("u"."number")`, "FUNCTION MIN")
			assertContains(t, sqlSelectQuery, `STDEV("u"."number")`, "FUNCTION STDDEV")
			assertContains(t, sqlSelectQuery, `SUM("u"."number")`, "FUNCTION SUM")
			assertContains(t, sqlSelectQuery, `VARIANCE("u"."number")`, "FUNCTION VARIANCE")
			// Функции аналитические
			assertContains(t, sqlSelectQuery, `FIRST_VALUE("u"."name") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION FIRSTVALUE")
			assertContains(t, sqlSelectQuery, `LAG("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)`, "FUNCTION LAG")
			assertContains(t, sqlSelectQuery, `LAST_VALUE("u"."name") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)`, "FUNCTION LASTVALUE")
			assertContains(t, sqlSelectQuery, `LEAD("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)`, "FUNCTION LEAD")
			assertContains(t, sqlSelectQuery, `NTH_VALUE("u"."name", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)`, "FUNCTION NTHVALUE")
			// Функции условий
			assertContains(t, sqlSelectQuery, `CASE WHEN "u"."number" < ? THEN ? ELSE ? END`, "FUNCTION CASE")
			assertContains(t, sqlSelectQuery, `COALESCE("u"."createat", "u"."updateat")`, "FUNCTION COALESCE")
			assertContains(t, sqlSelectQuery, `GREATEST("u"."createat", "u"."updateat")`, "FUNCTION GREATEST")
			assertContains(t, sqlSelectQuery, `LEAST("u"."createat", "u"."updateat")`, "FUNCTION LEAST")
			assertContains(t, sqlSelectQuery, `NULLIF("u"."createat", "u"."updateat")`, "FUNCTION NULLIF")
			// Функции конвертации
			assertContains(t, sqlSelectQuery, `CAST("u"."number" AS TEXT) `, "FUNCTION CAST")
			assertContains(t, sqlSelectQuery, `CHAR_LENGTH("u"."string")`, "FUNCTION CHARLENGTH")
			assertContains(t, sqlSelectQuery, `STRFTIME("u"."createat", '%Y-%m-%d')`, "FUNCTION DATEFORMAT")
			assertContains(t, sqlSelectQuery, `DEGREES("u"."number")`, "FUNCTION DEGREES")
			assertContains(t, sqlSelectQuery, `LENGTH("u"."string")`, "FUNCTION LENGTH")
			assertContains(t, sqlSelectQuery, `POSITION(? IN "u"."string")`, "FUNCTION POSITION")
			assertContains(t, sqlSelectQuery, `RADIANS("u"."number")`, "FUNCTION RADIANS")
			// Функции даты и времени
			assertContains(t, sqlSelectQuery, `DATE('now')`, "FUNCTION CURDATE")
			assertContains(t, sqlSelectQuery, `TIME('now')`, "FUNCTION CURTIME")
			assertContains(t, sqlSelectQuery, `DATETIME("u"."createat", '+2 DAY')`, "FUNCTION DATEADD")
			assertContains(t, sqlSelectQuery, `DATEDIFF("u"."updateat", "u"."createat")`, "FUNCTION DATEDIFF")
			assertContains(t, sqlSelectQuery, `DATETIME("u"."createat", '-2 DAY')`, "FUNCTION DATESUB")
			assertContains(t, sqlSelectQuery, `DAY("u"."createat")`, "FUNCTION DAY")
			assertContains(t, sqlSelectQuery, `STRFTIME('%w', "u"."createat")`, "FUNCTION DAYNAME")
			assertContains(t, sqlSelectQuery, `HOUR("u"."createat")`, "FUNCTION HOUR")
			assertContains(t, sqlSelectQuery, `MINUTE("u"."createat")`, "FUNCTION MINUTE")
			assertContains(t, sqlSelectQuery, `MONTH("u"."createat")`, "FUNCTION MONTH")
			assertContains(t, sqlSelectQuery, `STRFTIME('%m', "u"."createat")`, "FUNCTION MONTHNAME")
			assertContains(t, sqlSelectQuery, `DATETIME('now')`, "FUNCTION NOW")
			assertContains(t, sqlSelectQuery, `QUARTER("u"."createat")`, "FUNCTION QUARTER")
			assertContains(t, sqlSelectQuery, `SECOND("u"."createat")`, "FUNCTION SECOND")
			assertContains(t, sqlSelectQuery, `TIME("u"."createat", '+2 HOUR')`, "FUNCTION TIMEADD")
			assertContains(t, sqlSelectQuery, `TIMEDIFF("u"."updateat", "u"."createat")`, "FUNCTION TIMEDIFF")
			assertContains(t, sqlSelectQuery, `TIME("u"."createat", '-2 HOUR')`, "FUNCTION TIMESUB")
			assertContains(t, sqlSelectQuery, `WEEK("u"."createat")`, "FUNCTION WEEK")
			assertContains(t, sqlSelectQuery, `YEAR("u"."createat")`, "FUNCTION YEAR")
			// Функции обмена данными
			assertContains(t, sqlSelectQuery, `JSON_ARRAY("u"."json", ?, ?)`, "FUNCTION JSONARRAY")
			assertContains(t, sqlSelectQuery, `JSON_GROUP_ARRAY("u"."json")`, "FUNCTION JSONARRAYAGG")
			assertContains(t, sqlSelectQuery, `JSON_CONTAINS("u"."json", ?)`, "FUNCTION JSONCONTAINS")
			assertContains(t, sqlSelectQuery, `("u"."json" ->> '$.parent[0].child')`, "FUNCTION JSONEXTRACT")
			assertContains(t, sqlSelectQuery, `JSON_OBJECT('key', COUNT("u"."json"))`, "FUNCTION JSONOBJECT")
			assertContains(t, sqlSelectQuery, `JSON_GROUP_OBJECT("u"."json", "u"."number")`, "FUNCTION JSONOBJECTAGG")
			assertContains(t, sqlSelectQuery, `JSON_REMOVE("u"."json", '$.key1', '$.key2')`, "FUNCTION JSONREMOVE")
			assertContains(t, sqlSelectQuery, `JSON_SET("u"."json", '$.key1', ?, '$.key2', ?)`, "FUNCTION JSONSET")
			assertContains(t, sqlSelectQuery, `JSON_TYPE("u"."json")`, "FUNCTION JSONTYPE")
			// Функции математические
			assertContains(t, sqlSelectQuery, `ABS("u"."x")`, "FUNCTION ABS")
			assertContains(t, sqlSelectQuery, `ACOS("u"."x")`, "FUNCTION ACOS")
			assertContains(t, sqlSelectQuery, `ASIN("u"."x")`, "FUNCTION ASIN")
			assertContains(t, sqlSelectQuery, `ATAN("u"."x")`, "FUNCTION ATAN")
			assertContains(t, sqlSelectQuery, `ATAN2("u"."y", "u"."x")`, "FUNCTION ATAN2")
			assertContains(t, sqlSelectQuery, `CBRT("u"."x")`, "FUNCTION CBRT")
			assertContains(t, sqlSelectQuery, `CEIL("u"."x")`, "FUNCTION CEIL")
			assertContains(t, sqlSelectQuery, `COS("u"."x")`, "FUNCTION COS")
			assertContains(t, sqlSelectQuery, `EXP("u"."x")`, "FUNCTION EXP")
			assertContains(t, sqlSelectQuery, `FLOOR("u"."x")`, "FUNCTION FLOOR")
			assertContains(t, sqlSelectQuery, `LN("u"."x")`, "FUNCTION LN")
			assertContains(t, sqlSelectQuery, `LOG("u"."x", ?)`, "FUNCTION LOG")
			assertContains(t, sqlSelectQuery, `MOD("u"."x", ?)`, "FUNCTION MOD")
			assertContains(t, sqlSelectQuery, `PI()`, "FUNCTION PI")
			assertContains(t, sqlSelectQuery, `POWER("u"."x", ?)`, "FUNCTION POWER")
			assertContains(t, sqlSelectQuery, `RANDOM`, "FUNCTION RAND")
			assertContains(t, sqlSelectQuery, `ROUND("u"."x", ?)`, "FUNCTION ROUND")
			assertContains(t, sqlSelectQuery, `SIN("u"."x")`, "FUNCTION SIN")
			assertContains(t, sqlSelectQuery, `SQRT("u"."x")`, "FUNCTION SQRT")
			assertContains(t, sqlSelectQuery, `TAN("u"."x")`, "FUNCTION TAN")
			assertContains(t, sqlSelectQuery, `TRUNC("u"."x", ?)`, "FUNCTION TRUNC")
			// Функции ранжирующие
			assertContains(t, sqlSelectQuery, `CUME_DIST() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION CUMEDIST")
			assertContains(t, sqlSelectQuery, `DENSE_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "DFUNCTION ENSERANK")
			assertContains(t, sqlSelectQuery, `NTILE(2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION NTILE")
			assertContains(t, sqlSelectQuery, `PERCENT_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION PERCENTRANK")
			assertContains(t, sqlSelectQuery, `RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION RANK")
			assertContains(t, sqlSelectQuery, `ROW_NUMBER() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)`, "FUNCTION ROWNUMBER")
			// Функции строковые
			assertContains(t, sqlSelectQuery, `CONCAT("u"."string", ?, ?)`, "FUNCTION CONCAT")
			assertContains(t, sqlSelectQuery, `CONCAT_WS(?, "u"."string", ?, ?)`, "FUNCTION CONCATWS")
			assertContains(t, sqlSelectQuery, `LEFT("u"."string", ?)`, "FUNCTION LEFTSTRING")
			assertContains(t, sqlSelectQuery, `LOWER("u"."string")`, "FUNCTION LOWER")
			assertContains(t, sqlSelectQuery, `LPAD("u"."string", ?, ?)`, "FUNCTION LPAD")
			assertContains(t, sqlSelectQuery, `LTRIM("u"."string")`, "FUNCTION LTRIM")
			assertContains(t, sqlSelectQuery, `REPEAT("u"."string", ?)`, "FUNCTION REPEAT")
			assertContains(t, sqlSelectQuery, `REPLACE("u"."string", ?, ?)`, "FUNCTION REPLACE")
			assertContains(t, sqlSelectQuery, `REVERSE("u"."string")`, "FUNCTION REVERSE")
			assertContains(t, sqlSelectQuery, `RIGHT("u"."string", ?)`, "FUNCTION RIGHTSTRING")
			assertContains(t, sqlSelectQuery, `RPAD("u"."string", ?, ?)`, "FUNCTION RPAD")
			assertContains(t, sqlSelectQuery, `RTRIM("u"."string")`, "FUNCTION RTRIM")
			assertContains(t, sqlSelectQuery, `SUBSTRING("u"."string", ?, ?)`, "FUNCTION SUBSTRING")
			assertContains(t, sqlSelectQuery, `TRIM("u"."string")`, "FUNCTION TRIM")
			assertContains(t, sqlSelectQuery, `UPPER("u"."string")`, "FUNCTION UPPER")
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
			).
			Where(
				Equal(DateFormat(Test.Table.Users.CreateAt.Expr(), Literal("%Y-%m-%d")), Value("2026-01-01")),
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
			).
			Where(
				And(
					And(
						Equal(Test.Table.Users.String.Expr(), Value("active")),
						Greater(Test.Table.Users.Number.Expr(), Value(2)),
					),
					Or(
						Equal(Test.Table.Users.String.Expr(), Value("active")),
						Greater(Test.Table.Users.Number.Expr(), Value(2)),
					),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "`u`.`string` = ? AND `u`.`number` > ?", "LOGICAL AND")
			assertContains(t, sqlSelectQuery, "`u`.`string` = ? OR `u`.`number` > ?", "LOGICAL OR")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "[u].[string] = @p1 AND [u].[number] > @p2", "LOGICAL AND")
			assertContains(t, sqlSelectQuery, "[u].[string] = @p1 OR [u].[number] > @p2", "LOGICAL OR")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`u`.`string` = ? AND `u`.`number` > ?", "LOGICAL AND")
			assertContains(t, sqlSelectQuery, "`u`.`string` = ? OR `u`.`number` > ?", "LOGICAL OR")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"u"."string" = $1 AND "u"."number" > $2`, "LOGICAL AND")
			assertContains(t, sqlSelectQuery, `"u"."string" = $1 OR "u"."number" > $2`, "LOGICAL OR")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `"u"."string" = ? AND "u"."number" > ?`, "LOGICAL AND")
			assertContains(t, sqlSelectQuery, `"u"."string" = ? OR "u"."number" > ?`, "LOGICAL OR")
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Subquery[int64](NewSelect(Test.Table.User).Fields(Test.Table.Users.ID.Expr())).As("SUB"),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "(SELECT `u`.`id` FROM `users` AS `u`)", "SUBQUERY")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "(SELECT [u].[id] FROM [users] AS [u])", "SUBQUERY")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "(SELECT `u`.`id` FROM `users` AS `u`)", "SUBQUERY")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `(SELECT "u"."id" FROM "users" AS "u")`, "SUBQUERY")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `(SELECT "u"."id" FROM "users" AS "u")`, "SUBQUERY")
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
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Test.Table.Users.ID.Expr(),
			).
			Where(
				Equal(Test.Table.Users.String.Expr(), Value(data)),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "`u`.`string` = ?", "VALUE PLACEHOLDER")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "[u].[string] = @p1", "VALUE PLACEHOLDER")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "`u`.`string` = ?", "VALUE PLACEHOLDER")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `"u"."string" = $1`, "VALUE PLACEHOLDER")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `"u"."string" = ?`, "VALUE PLACEHOLDER")
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
			stmt := NewDelete(Test.Table.User).
				Join(
					Inner(Test.Table.Order, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				).
				Where(
					And(
						Equal(Test.Table.Users.String.Expr(), Value("active")),
						ILike(Test.Table.Users.String.Expr(), Value("%ivan%")),
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
			stmt := NewInsert(Test.Table.User).
				Values(
					Pair(Test.Table.Users.String.Expr(), Value("ivan")),
					Pair(Test.Table.Users.Number.Expr(), Value(2)),
				).
				Upsert(
					Pair(Test.Table.Users.String.Expr(), Value("updated")),
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
			stmt := NewSelect(Test.Table.User).
				Fields(
					Avg(Test.Table.Users.Number.Expr(), false).As("avg_result"),
					Ceil(Test.Table.Users.Number.Expr()).As("ceil_result"),
					Count(Test.Table.Users.String.Expr(), false).As("count_result"),
					FirstValue(Test.Table.Users.Name.Expr()).Over(
						PartitionBy(Test.Table.Users.ID.Expr()),
						OrderBy(Desc(Test.Table.Users.Number.Expr())),
					).As("first_value"),
					Trunc(Test.Table.Users.Number.Expr(), Value(2)).As("trunc_result"),
				).
				Join(
					Inner(Test.Table.Order, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				).
				Where(
					And(
						Equal(Test.Table.Users.String.Expr(), Value("active")),
						Greater(Test.Table.Users.Number.Expr(), Value(2)),
						ILike(Test.Table.Users.String.Expr(), Value("%ivan%")),
					),
				).
				GroupBy(
					Test.Table.Users.ID.Expr(),
					Test.Table.Users.String.Expr(),
				).
				Having(
					Greater(Count(Test.Table.Users.ID.Expr(), false), Value[int64](2)),
				).
				OrderBy(
					Desc(Test.Table.Users.Number.Expr()),
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
			stmt := NewUpdate(Test.Table.User).
				Set(
					Assign(Test.Table.Users.String.Expr(), Value("updated")),
					Assign(Test.Table.Users.Number.Expr(), Value(2)),
				).
				Join(
					Inner(Test.Table.Order, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				).
				Where(
					And(
						Equal(Test.Table.Users.String.Expr(), Value("active")),
						ILike(Test.Table.Users.String.Expr(), Value("%ivan%")),
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
func Test_SQL_Alter(t *testing.T) {
	// !!!Внимание, находится в стадии разработки
	t.Run("Index", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			// ALTER INDEX ... RENAME TO ...
			switch supportDialect {
			case DialectMariaDB:
				// Not supported - RENAME
			case DialectMsSQL:
				// Not supported - RENAME
			case DialectMySQL:
				// Not supported - RENAME
			case DialectPostgreSQL:
				//assertContains(t, sqlCommentQuery, `RENAME`, "RENAME")
			case DialectSQLite:
				// Not supported - RENAME
			}
			//t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCommentArguments, supportDialect.name, sqlCommentQuery)
		})
	})
	t.Run("Schema", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			// ALTER SCHEMA ... RENAME TO ...
			switch supportDialect {
			case DialectMariaDB:
				// Not supported - RENAME
			case DialectMsSQL:
				//assertContains(t, sqlCommentQuery, "RENAME", "RENAME")
			case DialectMySQL:
				// Not supported - RENAME
			case DialectPostgreSQL:
				//assertContains(t, sqlCommentQuery, `RENAME`, "RENAME")
			case DialectSQLite:
				// Not supported - RENAME
			}
			//t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCommentArguments, supportDialect.name, sqlCommentQuery)
		})
	})
	t.Run("Table", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			// ADD/DROP COLUMN, ADD/DROP CONSTRAINT, RENAME
			stmtAlter := NewAlter(Test.Table.User).
				AddColumns(
					Test.Table.Users.String,
					Test.Table.Users.Date,
				).
				AddConstraints(
					Test.Check.OrdersNumber,
					Test.Primary.OrdersID,
					Test.Unique.OrdersName,
				).
				DropColumns(
					Test.Table.Users.ID,
					Test.Table.Users.Name,
				).
				DropConstraints(
					Test.Check.UsersNumber,
					Test.Foreign.UsersOrders,
					Test.Primary.UsersID,
					Test.Unique.UsersName,
				)
			sqlAlterQuery, sqlAlterArguments, err := sql.Build(stmtAlter)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlAlterQuery, "ALTER TABLE `users` ADD COLUMN `string` VARCHAR, ADD COLUMN `date` DATE, ADD CONSTRAINT `ck_orders_number` CHECK(`o`.`number` > ?), ADD CONSTRAINT `pk_orders_id` PRIMARY KEY(`id`), ADD CONSTRAINT `un_orders_name` UNIQUE(`string`), DROP COLUMN `id`, DROP COLUMN `name`, DROP CONSTRAINT `ck_users_number`, DROP CONSTRAINT `fk_users_orders`, DROP CONSTRAINT `pk_users_id`, DROP CONSTRAINT `un_users_name`", "ALTER TABLE")
			case DialectMsSQL:
				assertContains(t, sqlAlterQuery, "ALTER TABLE [users] ADD COLUMN [string] NVARCHAR, ADD COLUMN [date] DATE, ADD CONSTRAINT [ck_orders_number] CHECK([o].[number] > @p1), ADD CONSTRAINT [pk_orders_id] PRIMARY KEY([id]), ADD CONSTRAINT [un_orders_name] UNIQUE([string]), DROP COLUMN [id], DROP COLUMN [name], DROP CONSTRAINT [ck_users_number], DROP CONSTRAINT [fk_users_orders], DROP CONSTRAINT [pk_users_id], DROP CONSTRAINT [un_users_name]", "ALTER TABLE")
			case DialectMySQL:
				assertContains(t, sqlAlterQuery, "ALTER TABLE `users` ADD COLUMN `string` VARCHAR, ADD COLUMN `date` DATE, ADD CONSTRAINT `ck_orders_number` CHECK(`o`.`number` > ?), ADD CONSTRAINT `pk_orders_id` PRIMARY KEY(`id`), ADD CONSTRAINT `un_orders_name` UNIQUE(`string`), DROP COLUMN `id`, DROP COLUMN `name`, DROP CONSTRAINT `ck_users_number`, DROP CONSTRAINT `fk_users_orders`, DROP CONSTRAINT `pk_users_id`, DROP CONSTRAINT `un_users_name`", "ALTER TABLE")
			case DialectPostgreSQL:
				assertContains(t, sqlAlterQuery, `ALTER TABLE "users" ADD COLUMN "string" VARCHAR, ADD COLUMN "date" DATE, ADD CONSTRAINT "ck_orders_number" CHECK("o"."number" > $1), ADD CONSTRAINT "pk_orders_id" PRIMARY KEY("id"), ADD CONSTRAINT "un_orders_name" UNIQUE("string"), DROP COLUMN "id", DROP COLUMN "name", DROP CONSTRAINT "ck_users_number", DROP CONSTRAINT "fk_users_orders", DROP CONSTRAINT "pk_users_id", DROP CONSTRAINT "un_users_name"`, "ALTER TABLE")
			case DialectSQLite:
				assertContains(t, sqlAlterQuery, `ALTER TABLE "users" ADD COLUMN "string" TEXT, ADD COLUMN "date" TEXT`, "ALTER TABLE")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlAlterArguments, supportDialect.name, sqlAlterQuery)
		})
	})
	t.Run("View", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			// ALTER VIEW ... RENAME TO ...
			switch supportDialect {
			case DialectMariaDB:
				// Not supported - RENAME
			case DialectMsSQL:
				//assertContains(t, sqlCommentQuery, "RENAME", "RENAME")
			case DialectMySQL:
				// Not supported - RENAME
			case DialectPostgreSQL:
				//assertContains(t, sqlCommentQuery, `RENAME`, "RENAME")
			case DialectSQLite:
				// Not supported - RENAME
			}
			//t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCommentArguments, supportDialect.name, sqlCommentQuery)
		})
	})
}
func Test_SQL_Comment(t *testing.T) {
	// !!!Внимание, находится в стадии разработки
	t.Run("Column", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtComment := NewComment(Test.Table.Users.ID).
				Is("user comment")
			sqlCommentQuery, sqlCommentArguments, err := sql.Build(stmtComment)
			switch supportDialect {
			case DialectMariaDB:
				//assertContains(t, sqlCommentQuery, "COMMENT ON COLUMN `users`.`id` IS 'user comment'", "COMMENT")
			case DialectMsSQL:
				// Not supported - COMMENT
			case DialectMySQL:
				//assertContains(t, sqlCommentQuery, "COMMENT ON COLUMN `users`.`id` IS 'user comment'", "COMMENT")
			case DialectPostgreSQL:
				//assertContains(t, sqlCommentQuery, `COMMENT ON COLUMN "users"."id" IS 'user comment'`, "COMMENT")
			case DialectSQLite:
				//assertContains(t, sqlCommentQuery, `COMMENT ON COLUMN "users"."id" IS 'user comment'`, "COMMENT")
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
			stmtComment := NewComment(Test.Table.User).Is("user comment")
			sqlCommentQuery, sqlCommentArguments, err := sql.Build(stmtComment)
			switch supportDialect {
			case DialectMariaDB:
				//assertContains(t, sqlCommentQuery, "COMMENT ON TABLE `users` IS 'user comment'", "COMMENT")
			case DialectMsSQL:
				// Not supported - COMMENT
			case DialectMySQL:
				//assertContains(t, sqlCommentQuery, "COMMENT ON TABLE `users` IS 'user comment'", "COMMENT")
			case DialectPostgreSQL:
				//assertContains(t, sqlCommentQuery, `COMMENT ON TABLE "users" IS 'user comment'`, "COMMENT")
			case DialectSQLite:
				//assertContains(t, sqlCommentQuery, `COMMENT ON TABLE "users" IS 'user comment'`, "COMMENT")
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
			stmtCreate := NewCreate(Test.Index.UsersID).
				IfNotExists().
				IsUnique().
				On(Test.Table.User).
				Columns(
					Test.Table.Users.String,
					Test.Table.Users.Number,
				)
			sqlCreateQuery, sqlCreateArguments, err := sql.Build(stmtCreate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlCreateQuery, "CREATE UNIQUE INDEX IF NOT EXISTS `users_id` ON `users` (`string`, `number`)", "CREATE INDEX")
			case DialectMsSQL:
				assertContains(t, sqlCreateQuery, "CREATE UNIQUE INDEX [users_id] ON [users] ([string], [number])", "CREATE INDEX")
			case DialectMySQL:
				assertContains(t, sqlCreateQuery, "CREATE UNIQUE INDEX `users_id` ON `users` (`string`, `number`)", "CREATE INDEX")
			case DialectPostgreSQL:
				assertContains(t, sqlCreateQuery, `CREATE UNIQUE INDEX IF NOT EXISTS "users_id" ON "users" ("string", "number")`, "CREATE INDEX")
			case DialectSQLite:
				assertContains(t, sqlCreateQuery, `CREATE UNIQUE INDEX IF NOT EXISTS "users_id" ON "users" ("string", "number")`, "CREATE INDEX")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCreateArguments, supportDialect.name, sqlCreateQuery)
		})
	})
	t.Run("Schema", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtCreate := NewCreate(Test.Schema).
				IfNotExists()
			sqlCreateQuery, sqlCreateArguments, err := sql.Build(stmtCreate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlCreateQuery, "CREATE SCHEMA IF NOT EXISTS `test`", "CREATE SCHEMA")
			case DialectMsSQL:
				assertContains(t, sqlCreateQuery, "CREATE SCHEMA IF NOT EXISTS [test]", "CREATE SCHEMA")
			case DialectMySQL:
				assertContains(t, sqlCreateQuery, "CREATE SCHEMA `test`", "CREATE SCHEMA")
			case DialectPostgreSQL:
				assertContains(t, sqlCreateQuery, `CREATE SCHEMA IF NOT EXISTS "test"`, "CREATE SCHEMA")
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
			stmtCreate := NewCreate(Test.Table.User).
				Constraints(
					Test.Check.UsersNumber,
					Test.Foreign.UsersOrders,
					Test.Primary.UsersID,
					Test.Unique.UsersName,
				).
				IfNotExists().
				Columns(
					Test.Table.Users.ID.AutoIncrement(),
					Test.Table.Users.Name.NotNull(),
				)
			sqlCreateQuery, sqlCreateArguments, err := sql.Build(stmtCreate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlCreateQuery, "CREATE TABLE IF NOT EXISTS `users` (`id` SIGNED AUTO_INCREMENT, `name` VARCHAR NOT NULL, CONSTRAINT `ck_users_number` CHECK(`u`.`number` > ?), CONSTRAINT `fk_users_orders` FOREIGN KEY(`data_id`, `name`) REFERENCES `orders`(`id`, `string`) ON DELETE CASCADE ON UPDATE RESTRICT, CONSTRAINT `pk_users_id` PRIMARY KEY(`id`), CONSTRAINT `un_users_name` UNIQUE(`name`))", "CREATE TABLE")
			case DialectMsSQL:
				assertContains(t, sqlCreateQuery, "CREATE TABLE [users] ([id] BIGINT IDENTITY(1,1), [name] NVARCHAR NOT NULL, CONSTRAINT [ck_users_number] CHECK([u].[number] > @p1), CONSTRAINT [fk_users_orders] FOREIGN KEY([data_id], [name]) REFERENCES [orders]([id], [string]) ON DELETE CASCADE ON UPDATE RESTRICT, CONSTRAINT [pk_users_id] PRIMARY KEY([id]), CONSTRAINT [un_users_name] UNIQUE([name]))", "CREATE TABLE")
			case DialectMySQL:
				assertContains(t, sqlCreateQuery, "CREATE TABLE IF NOT EXISTS `users` (`id` SIGNED AUTO_INCREMENT, `name` VARCHAR NOT NULL, CONSTRAINT `ck_users_number` CHECK(`u`.`number` > ?), CONSTRAINT `fk_users_orders` FOREIGN KEY(`data_id`, `name`) REFERENCES `orders`(`id`, `string`) ON DELETE CASCADE ON UPDATE RESTRICT, CONSTRAINT `pk_users_id` PRIMARY KEY(`id`), CONSTRAINT `un_users_name` UNIQUE(`name`))", "CREATE TABLE")
			case DialectPostgreSQL:
				assertContains(t, sqlCreateQuery, `CREATE TABLE IF NOT EXISTS "users" ("id" BIGINT GENERATED BY DEFAULT AS IDENTITY, "name" VARCHAR NOT NULL, CONSTRAINT "ck_users_number" CHECK("u"."number" > $1), CONSTRAINT "fk_users_orders" FOREIGN KEY("data_id", "name") REFERENCES "orders"("id", "string") ON DELETE CASCADE ON UPDATE RESTRICT, CONSTRAINT "pk_users_id" PRIMARY KEY("id"), CONSTRAINT "un_users_name" UNIQUE("name"))`, "CREATE TABLE")
			case DialectSQLite:
				assertContains(t, sqlCreateQuery, `CREATE TABLE IF NOT EXISTS "users" ("id" INTEGER AUTOINCREMENT, "name" TEXT NOT NULL, CONSTRAINT "ck_users_number" CHECK("u"."number" > ?), CONSTRAINT "fk_users_orders" FOREIGN KEY("data_id", "name") REFERENCES "orders"("id", "string") ON DELETE CASCADE ON UPDATE RESTRICT, CONSTRAINT "pk_users_id" PRIMARY KEY("id"), CONSTRAINT "un_users_name" UNIQUE("name"))`, "CREATE TABLE")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlCreateArguments, supportDialect.name, sqlCreateQuery)
		})
	})
	t.Run("View", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtCreate := NewCreate(Test.View.UsersGeneral).
				IsReplace().
				Source(NewSelect(Test.Table.User).
					Fields(Test.Table.Users.ID.Expr(), Test.Table.Users.String.Expr()),
				)
			sqlCreateQuery, sqlCreateArguments, err := sql.Build(stmtCreate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlCreateQuery, "CREATE OR REPLACE VIEW `users_general` AS SELECT `u`.`id`, `u`.`string` FROM `users` AS `u`", "CREATE VIEW")
			case DialectMsSQL:
				assertContains(t, sqlCreateQuery, "CREATE OR REPLACE VIEW [users_general] AS SELECT [u].[id], [u].[string] FROM [users] AS [u]", "CREATE VIEW")
			case DialectMySQL:
				assertContains(t, sqlCreateQuery, "CREATE OR REPLACE VIEW `users_general` AS SELECT `u`.`id`, `u`.`string` FROM `users` AS `u`", "CREATE VIEW")
			case DialectPostgreSQL:
				assertContains(t, sqlCreateQuery, `CREATE OR REPLACE VIEW "users_general" AS SELECT "u"."id", "u"."string" FROM "users" AS "u"`, "CREATE VIEW")
			case DialectSQLite:
				assertContains(t, sqlCreateQuery, `CREATE OR REPLACE VIEW "users_general" AS SELECT "u"."id", "u"."string" FROM "users" AS "u"`, "CREATE VIEW")
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
			stmtDelete := NewDelete(Test.Table.User).
				Join(
					Inner(Test.Table.Order, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				).
				Where(
					Equal(Test.Table.Users.String.Expr(), Value("active")),
				)
			sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDeleteQuery, "DELETE `u` FROM `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id` WHERE `u`.`string` = ?", "DELETE JOIN")
			case DialectMsSQL:
				assertContains(t, sqlDeleteQuery, "DELETE [u] FROM [users] AS [u] INNER JOIN [orders] AS [o] ON [u].[id] = [o].[id] WHERE [u].[string] = @p1", "DELETE JOIN")
			case DialectMySQL:
				assertContains(t, sqlDeleteQuery, "DELETE `u` FROM `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id` WHERE `u`.`string` = ?", "DELETE JOIN")
			case DialectPostgreSQL:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "users" AS "u" USING "orders" AS "o" WHERE ("u"."id" = "o"."id" AND "u"."string" = $1)`, "DELETE JOIN")
			case DialectSQLite:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "users" AS "u" INNER JOIN "orders" AS "o" ON "u"."id" = "o"."id" WHERE "u"."string" = ?`, "DELETE JOIN")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
		})
	})
	t.Run("Returning", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtDelete := NewDelete(Test.Table.User).
				Where(
					Equal(Test.Table.Users.String.Expr(), Value("active")),
				).
				Returning(
					Test.Table.Users.ID.Expr(),
					Test.Table.Users.String.Expr(),
				)
			sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDeleteQuery, "DELETE `u` FROM `users` AS `u` WHERE `u`.`string` = ? RETURNING `u`.`id`, `u`.`string`", "DELETE RETURNING")
			case DialectMsSQL:
				assertContains(t, sqlDeleteQuery, "DELETE [u] FROM [users] AS [u] OUTPUT [u].[id], [u].[string] WHERE [u].[string] = @p1", "DELETE RETURNING")
			case DialectMySQL:
				assertContains(t, sqlDeleteQuery, "DELETE `u` FROM `users` AS `u` WHERE `u`.`string` = ?", "DELETE WITHOUT RETURNING")
			case DialectPostgreSQL:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "users" AS "u" WHERE "u"."string" = $1 RETURNING "u"."id", "u"."string"`, "DELETE RETURNING")
			case DialectSQLite:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "users" AS "u" WHERE "u"."string" = ? RETURNING "u"."id", "u"."string"`, "DELETE RETURNING")
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
			stmtDelete := NewDelete(Test.Table.User).
				Where(
					Equal(Test.Table.Users.String.Expr(), Value("active")),
				)
			sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDeleteQuery, "DELETE `u` FROM `users` AS `u` WHERE `u`.`string` = ?", "DELETE")
			case DialectMsSQL:
				assertContains(t, sqlDeleteQuery, "DELETE [u] FROM [users] AS [u] WHERE [u].[string] = @p1", "DELETE")
			case DialectMySQL:
				assertContains(t, sqlDeleteQuery, "DELETE `u` FROM `users` AS `u` WHERE `u`.`string` = ?", "DELETE")
			case DialectPostgreSQL:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "users" AS "u" WHERE "u"."string" = $1`, "DELETE")
			case DialectSQLite:
				assertContains(t, sqlDeleteQuery, `DELETE FROM "users" AS "u" WHERE "u"."string" = ?`, "DELETE")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
		})
	})
	t.Run("With", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtDelete := NewDelete(Test.Table.User).
				With(
					WithN("old_users", NewSelect(Test.Table.User).
						Fields(
							Test.Table.Users.ID.Expr(),
						).
						Where(
							Less(Test.Table.Users.Number.Expr(), Value(2)),
						),
					),
				)
			sqlDeleteQuery, sqlDeleteArguments, err := sql.Build(stmtDelete)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDeleteQuery, "WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) DELETE `u` FROM `users` AS `u`", "WITH")
			case DialectMsSQL:
				assertContains(t, sqlDeleteQuery, "WITH [old_users] AS (SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] < @p1) DELETE [u] FROM [users] AS [u]", "WITH")
			case DialectMySQL:
				assertContains(t, sqlDeleteQuery, "WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) DELETE `u` FROM `users` AS `u`", "WITH")
			case DialectPostgreSQL:
				assertContains(t, sqlDeleteQuery, `WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < $1) DELETE FROM "users" AS "u"`, "WITH")
			case DialectSQLite:
				assertContains(t, sqlDeleteQuery, `WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < ?) DELETE FROM "users" AS "u"`, "WITH")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDeleteArguments, supportDialect.name, sqlDeleteQuery)
		})
	})
}
func Test_SQL_Drop(t *testing.T) {
	t.Run("CascadeIndex", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtDrop := NewDrop(Test.Index.UsersID).IsCascade()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP INDEX `users_id` CASCADE", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP INDEX [users_id]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP INDEX `users_id`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP INDEX "users_id" CASCADE`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP INDEX "users_id"`, "DROP")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDropArguments, supportDialect.name, sqlDropQuery)
		})
	})
	t.Run("CascadeSchema", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtDrop := NewDrop(Test.Schema).IsCascade()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP SCHEMA `test` CASCADE", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP SCHEMA [test]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP SCHEMA `test`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP SCHEMA "test" CASCADE`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP SCHEMA "test"`, "DROP")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDropArguments, supportDialect.name, sqlDropQuery)
		})
	})
	t.Run("CascadeTable", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtDrop := NewDrop(Test.Table.User).IsCascade()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP TABLE `users`", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP TABLE [users]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP TABLE `users`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP TABLE "users" CASCADE`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP TABLE "users"`, "DROP")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlDropArguments, supportDialect.name, sqlDropQuery)
		})
	})
	t.Run("CascadeView", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(
				WithDialect(supportDialect),
			)
			defer sql.Close()
			stmtDrop := NewDrop(Test.View.UsersGeneral).IsCascade()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP VIEW `users_general` CASCADE", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP VIEW [users_general]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP VIEW `users_general`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP VIEW "users_general" CASCADE`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP VIEW "users_general"`, "DROP")
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
			stmtDrop := NewDrop(Test.Index.UsersID).IfExists()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP INDEX IF EXISTS `users_id`", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP INDEX [users_id]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP INDEX `users_id`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP INDEX IF EXISTS "users_id"`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP INDEX IF EXISTS "users_id"`, "DROP")
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
			stmtDrop := NewDrop(Test.Schema).IfExists()
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
			stmtDrop := NewDrop(Test.Table.User).IfExists()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP TABLE IF EXISTS `users`", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP TABLE IF EXISTS [users]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP TABLE IF EXISTS `users`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP TABLE IF EXISTS "users"`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP TABLE IF EXISTS "users"`, "DROP")
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
			stmtDrop := NewDrop(Test.View.UsersGeneral).IfExists()
			sqlDropQuery, sqlDropArguments, err := sql.Build(stmtDrop)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlDropQuery, "DROP VIEW IF EXISTS `users_general`", "DROP")
			case DialectMsSQL:
				assertContains(t, sqlDropQuery, "DROP VIEW IF EXISTS [users_general]", "DROP")
			case DialectMySQL:
				assertContains(t, sqlDropQuery, "DROP VIEW IF EXISTS `users_general`", "DROP")
			case DialectPostgreSQL:
				assertContains(t, sqlDropQuery, `DROP VIEW IF EXISTS "users_general"`, "DROP")
			case DialectSQLite:
				assertContains(t, sqlDropQuery, `DROP VIEW IF EXISTS "users_general"`, "DROP")
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
			stmtInsert := NewInsert(Test.Table.User).
				Values(
					Pair(Test.Table.Users.String.Expr(), Value("ivan")),
					Pair(Test.Table.Users.Number.Expr(), Value(2)),
				).
				Returning(
					Test.Table.Users.ID.Expr(),
					Test.Table.Users.String.Expr(),
				)
			sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlInsertQuery, "INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?) RETURNING `u`.`id`, `u`.`string`", "INSERT RETURNING")
			case DialectMsSQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO [users] AS [u] ([string], [number]) OUTPUT [u].[id], [u].[string] VALUES (@p1, @p2)", "INSERT RETURNING")
			case DialectMySQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?)", "INSERT WITHOUT RETURNING")
			case DialectPostgreSQL:
				assertContains(t, sqlInsertQuery, `INSERT INTO "users" AS "u" ("string", "number") VALUES ($1, $2) RETURNING "u"."id", "u"."string"`, "INSERT RETURNING")
			case DialectSQLite:
				assertContains(t, sqlInsertQuery, `INSERT INTO "users" AS "u" ("string", "number") VALUES (?, ?) RETURNING "u"."id", "u"."string"`, "INSERT RETURNING")
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
			stmtInsert := NewInsert(Test.Table.User).
				Source(NewSelect(Test.Table.User).
					Fields(
						Test.Table.Users.String.Expr(),
						Test.Table.Users.Number.Expr(),
					).
					Where(
						Equal(Test.Table.Users.String.Expr(), Value("active")),
					),
				)
			sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlInsertQuery, "INSERT INTO `users` AS `u` (`string`, `number`) SELECT `u`.`string`, `u`.`number` FROM `users` AS `u` WHERE `u`.`string` = ?", "INSERT SOURCE")
			case DialectMsSQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO [users] AS [u] ([string], [number]) SELECT [u].[string], [u].[number] FROM [users] AS [u] WHERE [u].[string] = @p1", "INSERT SOURCE")
			case DialectMySQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO `users` AS `u` (`string`, `number`) SELECT `u`.`string`, `u`.`number` FROM `users` AS `u` WHERE `u`.`string` = ?", "INSERT SOURCE")
			case DialectPostgreSQL:
				assertContains(t, sqlInsertQuery, `INSERT INTO "users" AS "u" ("string", "number") SELECT "u"."string", "u"."number" FROM "users" AS "u" WHERE "u"."string" = $1`, "INSERT SOURCE")
			case DialectSQLite:
				assertContains(t, sqlInsertQuery, `INSERT INTO "users" AS "u" ("string", "number") SELECT "u"."string", "u"."number" FROM "users" AS "u" WHERE "u"."string" = ?`, "INSERT SOURCE")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
		})
	})
	t.Run("Values", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtInsert := NewInsert(Test.Table.User).
				Values(
					Pair(Test.Table.Users.String.Expr(), Value("ivan")),
					Pair(Test.Table.Users.Number.Expr(), Value(2)),
				).
				Upsert(
					Pair(Test.Table.Users.String.Expr(), Value("updated")),
				)
			sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlInsertQuery, "INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?", "INSERT VALUES WITH UPSERT")
			case DialectMsSQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO [users] AS [u] ([string], [number]) VALUES (@p1, @p2)", "INSERT VALUES WITHOUT UPSERT")
			case DialectMySQL:
				assertContains(t, sqlInsertQuery, "INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?", "INSERT VALUES WITH UPSERT")
			case DialectPostgreSQL:
				assertContains(t, sqlInsertQuery, `INSERT INTO "users" AS "u" ("string", "number") VALUES ($1, $2) ON CONFLICT DO UPDATE SET "string" = $3`, "INSERT UVALUES WITH PSERT")
			case DialectSQLite:
				assertContains(t, sqlInsertQuery, `INSERT INTO "users" AS "u" ("string", "number") VALUES (?, ?) ON CONFLICT DO UPDATE SET "string" = ?`, "INSERT VALUES WITH UPSERT")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlInsertArguments, supportDialect.name, sqlInsertQuery)
		})
	})
	t.Run("With", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtInsert := NewInsert(Test.Table.User).
				Values(
					Pair(Test.Table.Users.String.Expr(), Value("ivan")),
					Pair(Test.Table.Users.Number.Expr(), Value(2)),
				).
				With(
					WithN("old_users", NewSelect(Test.Table.User).
						Fields(
							Test.Table.Users.ID.Expr(),
						).
						Where(
							Less(Test.Table.Users.Number.Expr(), Value(2)),
						),
					),
				)
			sqlInsertQuery, sqlInsertArguments, err := sql.Build(stmtInsert)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlInsertQuery, "WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?)", "INSERT WITH")
			case DialectMsSQL:
				assertContains(t, sqlInsertQuery, "WITH [old_users] AS (SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] < @p1) INSERT INTO [users] AS [u] ([string], [number]) VALUES (@p2, @p3)", "INSERT WITH")
			case DialectMySQL:
				assertContains(t, sqlInsertQuery, "WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?)", "INSERT WITH")
			case DialectPostgreSQL:
				assertContains(t, sqlInsertQuery, `WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < $1) INSERT INTO "users" AS "u" ("string", "number") VALUES ($2, $3)`, "INSERT WITH")
			case DialectSQLite:
				assertContains(t, sqlInsertQuery, `WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < ?) INSERT INTO "users" AS "u" ("string", "number") VALUES (?, ?)`, "INSERT WITH")
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
			stmtSelect := NewSelect(Test.Table.User).
				Distinct().
				Fields(
					Test.Table.Users.ID.Expr(),
				).
				Where(
					Equal(Test.Table.Users.Number.Expr(), Value(2)),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT DISTINCT `u`.`id` FROM `users` AS `u`", "SELECT DISTINCT")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT DISTINCT [u].[id] FROM [users] AS [u]", "SELECT DISTINCT")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT DISTINCT `u`.`id` FROM `users` AS `u`", "SELECT DISTINCT")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT DISTINCT "u"."id" FROM "users" AS "u"`, "SELECT DISTINCT")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT DISTINCT "u"."id" FROM "users" AS "u"`, "SELECT DISTINCT")
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
			stmtSelect := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.ID.Expr(),
				).
				Where(
					Equal(Test.Table.Users.Number.Expr(), Value(2)),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` = ?", "SELECT")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] = @p1", "SELECT")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` = ?", "SELECT")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" = $1`, "SELECT")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" = ?`, "SELECT")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("GroupBy", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.String.Expr(),
					Count(Test.Table.Users.ID.Expr(), false).As("cnt"),
				).
				GroupBy(
					Test.Table.Users.String.Expr(),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`string`, COUNT(`u`.`id`) AS `cnt` FROM `users` AS `u` GROUP BY `u`.`string`", "SELECT GROUP BY")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [u].[string], COUNT([u].[id]) AS [cnt] FROM [users] AS [u] GROUP BY [u].[string]", "SELECT GROUP BY")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`string`, COUNT(`u`.`id`) AS `cnt` FROM `users` AS `u` GROUP BY `u`.`string`", "SELECT GROUP BY")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "u"."string", COUNT("u"."id") AS "cnt" FROM "users" AS "u" GROUP BY "u"."string"`, "SELECT GROUP BY")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "u"."string", COUNT("u"."id") AS "cnt" FROM "users" AS "u" GROUP BY "u"."string"`, "SELECT GROUP BY")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("Having", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.String.Expr(),
					Count(Test.Table.Users.ID.Expr(), false).As("cnt"),
				).
				GroupBy(
					Test.Table.Users.String.Expr(),
				).
				Having(
					Greater(Count(Test.Table.Users.ID.Expr(), false), Value[int64](2)),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`string`, COUNT(`u`.`id`) AS `cnt` FROM `users` AS `u` GROUP BY `u`.`string` HAVING COUNT(`u`.`id`) > ?", "SELECT HAVING")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [u].[string], COUNT([u].[id]) AS [cnt] FROM [users] AS [u] GROUP BY [u].[string] HAVING COUNT([u].[id]) > @p1", "SELECT HAVING")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`string`, COUNT(`u`.`id`) AS `cnt` FROM `users` AS `u` GROUP BY `u`.`string` HAVING COUNT(`u`.`id`) > ?", "SELECT HAVING")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "u"."string", COUNT("u"."id") AS "cnt" FROM "users" AS "u" GROUP BY "u"."string" HAVING COUNT("u"."id") > $1`, "SELECT HAVING")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "u"."string", COUNT("u"."id") AS "cnt" FROM "users" AS "u" GROUP BY "u"."string" HAVING COUNT("u"."id") > ?`, "SELECT HAVING")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("Join", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.ID.Expr(),
					Test.Table.Orders.String.Expr(),
				).
				Join(
					Inner(Test.Table.Order, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`id`, `o`.`string` FROM `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id`", "SELECT JOIN")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [u].[id], [o].[string] FROM [users] AS [u] INNER JOIN [orders] AS [o] ON [u].[id] = [o].[id]", "SELECT JOIN")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`id`, `o`.`string` FROM `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id`", "SELECT JOIN")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "u"."id", "o"."string" FROM "users" AS "u" INNER JOIN "orders" AS "o" ON "u"."id" = "o"."id"`, "SELECT JOIN")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "u"."id", "o"."string" FROM "users" AS "u" INNER JOIN "orders" AS "o" ON "u"."id" = "o"."id"`, "SELECT JOIN")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("OrderBy", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.ID.Expr(),
				).
				OrderBy(
					Desc(Test.Table.Users.Number.Expr()),
					Asc(Test.Table.Users.String.Expr()),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`id` FROM `users` AS `u` ORDER BY `u`.`number` DESC, `u`.`string` ASC", "SELECT ORDER BY")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [u].[id] FROM [users] AS [u] ORDER BY [u].[number] DESC, [u].[string] ASC", "SELECT ORDER BY")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`id` FROM `users` AS `u` ORDER BY `u`.`number` DESC, `u`.`string` ASC", "SELECT ORDER BY")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "u"."id" FROM "users" AS "u" ORDER BY "u"."number" DESC, "u"."string" ASC`, "SELECT ORDER BY")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "u"."id" FROM "users" AS "u" ORDER BY "u"."number" DESC, "u"."string" ASC`, "SELECT ORDER BY")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("Pagination", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.ID.Expr(),
				).
				Pagination(10, 20)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`id` FROM `users` AS `u` LIMIT ? OFFSET ?", "SELECT PAGINATION")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [u].[id] FROM [users] AS [u] ORDER BY 1 ASC OFFSET @p1 ROWS FETCH NEXT @p2 ROWS ONLY", "SELECT PAGINATION")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`id` FROM `users` AS `u` LIMIT ? OFFSET ?", "SELECT PAGINATION")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "u"."id" FROM "users" AS "u" LIMIT $1 OFFSET $2`, "SELECT PAGINATION")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "u"."id" FROM "users" AS "u" LIMIT ? OFFSET ?`, "SELECT PAGINATION")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("Unions", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.String.Expr(),
				).
				Unions(
					UnionAll(NewSelect(Test.Table.Order).
						Fields(
							Test.Table.Orders.String.Expr(),
						),
					),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`string` FROM `users` AS `u` UNION ALL SELECT `o`.`string` FROM `orders` AS `o`", "SELECT UNION ALL")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [u].[string] FROM [users] AS [u] UNION ALL SELECT [o].[string] FROM [orders] AS [o]", "SELECT UNION ALL")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`string` FROM `users` AS `u` UNION ALL SELECT `o`.`string` FROM `orders` AS `o`", "SELECT UNION ALL")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "u"."string" FROM "users" AS "u" UNION ALL SELECT "o"."string" FROM "orders" AS "o"`, "SELECT UNION ALL")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "u"."string" FROM "users" AS "u" UNION ALL SELECT "o"."string" FROM "orders" AS "o"`, "SELECT UNION ALL")
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
			stmtSelect := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.ID.Expr(),
				).
				Where(
					Equal(Test.Table.Users.Number.Expr(), Value(2)),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` = ?", "SELECT")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] = @p1", "SELECT")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` = ?", "SELECT")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" = $1`, "SELECT")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" = ?`, "SELECT")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
		})
	})
	t.Run("With", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtSelect := NewSelect(NewCTE("cte_user", "ct")).
				Fields(
					Field[int64]("ct", "id"),
				).
				With(
					WithN("cte_user", NewSelect(Test.Table.User).
						Fields(
							Test.Table.Users.ID.Expr(),
						).
						Where(
							Greater(Test.Table.Users.Number.Expr(), Value(2)),
						),
					),
				)
			sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlSelectQuery, "WITH `cte_user` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` > ?) SELECT `ct`.`id` FROM `cte_user` AS `ct`", "SELECT WITH")
			case DialectMsSQL:
				assertContains(t, sqlSelectQuery, "WITH [cte_user] AS (SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] > @p1) SELECT [ct].[id] FROM [cte_user] AS [ct]", "SELECT WITH")
			case DialectMySQL:
				assertContains(t, sqlSelectQuery, "WITH `cte_user` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` > ?) SELECT `ct`.`id` FROM `cte_user` AS `ct`", "SELECT WITH")
			case DialectPostgreSQL:
				assertContains(t, sqlSelectQuery, `WITH "cte_user" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" > $1) SELECT "ct"."id" FROM "cte_user" AS "ct"`, "SELECT WITH")
			case DialectSQLite:
				assertContains(t, sqlSelectQuery, `WITH "cte_user" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" > ?) SELECT "ct"."id" FROM "cte_user" AS "ct"`, "SELECT WITH")
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
			stmtTruncate := NewTruncate(Test.Table.User)
			sqlTruncateQuery, sqlTruncateArguments, err := sql.Build(stmtTruncate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlTruncateQuery, "TRUNCATE TABLE `users`", "TRUNCATE")
			case DialectMsSQL:
				assertContains(t, sqlTruncateQuery, "TRUNCATE TABLE [users]", "TRUNCATE")
			case DialectMySQL:
				assertContains(t, sqlTruncateQuery, "TRUNCATE TABLE `users`", "TRUNCATE")
			case DialectPostgreSQL:
				assertContains(t, sqlTruncateQuery, `TRUNCATE TABLE "users"`, "TRUNCATE")
			case DialectSQLite:
				assertContains(t, sqlTruncateQuery, `TRUNCATE TABLE "users"`, "TRUNCATE")
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
			stmtTruncate := NewTruncate(Test.Table.User).
				IsCascade()
			sqlTruncateQuery, sqlTruncateArguments, err := sql.Build(stmtTruncate)
			switch supportDialect {
			case DialectMariaDB:
				// Not supported - CASCADE
			case DialectMsSQL:
				// Not supported - CASCADE
			case DialectMySQL:
				// Not supported - CASCADE
			case DialectPostgreSQL:
				assertContains(t, sqlTruncateQuery, `TRUNCATE TABLE "users" CASCADE`, "TRUNCATE CASCADE")
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
			stmtTruncate := NewTruncate(Test.Table.User).
				IsRestartIdentity()
			sqlTruncateQuery, sqlTruncateArguments, err := sql.Build(stmtTruncate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlTruncateQuery, "TRUNCATE TABLE `users` RESTART IDENTITY", "TRUNCATE RESTART IDENTITY")
			case DialectMsSQL:
				// Not supported - RESTART IDENTITY
			case DialectMySQL:
				// Not supported - RESTART IDENTITY
			case DialectPostgreSQL:
				assertContains(t, sqlTruncateQuery, `TRUNCATE TABLE "users" RESTART IDENTITY`, "TRUNCATE RESTART IDENTITY")
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
			stmtUpdate := NewUpdate(Test.Table.User).
				Set(
					Assign(Test.Table.Users.String.Expr(), Value("active")),
				).
				Join(
					Inner(Test.Table.Order, Equal(Test.Table.Users.ID.Expr(), Test.Table.Orders.ID.Expr())),
				).
				Where(
					Equal(Test.Table.Orders.String.Expr(), Value("active")),
				)
			sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlUpdateQuery, "UPDATE `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id` SET `u`.`string` = ? WHERE `o`.`string` = ?", "UPDATE JOIN")
			case DialectMsSQL:
				assertContains(t, sqlUpdateQuery, "UPDATE [users] AS [u] INNER JOIN [orders] AS [o] ON [u].[id] = [o].[id] SET [u].[string] = @p1 WHERE [o].[string] = @p2", "UPDATE JOIN")
			case DialectMySQL:
				assertContains(t, sqlUpdateQuery, "UPDATE `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id` SET `u`.`string` = ? WHERE `o`.`string` = ?", "UPDATE JOIN")
			case DialectPostgreSQL:
				assertContains(t, sqlUpdateQuery, `UPDATE "users" AS "u" INNER JOIN "orders" AS "o" ON "u"."id" = "o"."id" SET "u"."string" = $1 WHERE "o"."string" = $2`, "UPDATE JOIN")
			case DialectSQLite:
				assertContains(t, sqlUpdateQuery, `UPDATE "users" AS "u" INNER JOIN "orders" AS "o" ON "u"."id" = "o"."id" SET "u"."string" = ? WHERE "o"."string" = ?`, "UPDATE JOIN")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
		})
	})
	t.Run("Returning", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtUpdate := NewUpdate(Test.Table.User).
				Set(
					Assign(Test.Table.Users.String.Expr(), Value("active")),
				).
				Where(
					Equal(Test.Table.Users.Number.Expr(), Value(2)),
				).
				Returning(
					Test.Table.Users.ID.Expr(),
					Test.Table.Users.String.Expr(),
				)
			sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlUpdateQuery, "UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ? RETURNING `u`.`id`, `u`.`string`", "UPDATE RETURNING")
			case DialectMsSQL:
				assertContains(t, sqlUpdateQuery, "UPDATE [users] AS [u] OUTPUT [u].[id], [u].[string] SET [u].[string] = @p1 WHERE [u].[number] = @p2", "UPDATE RETURNING")
			case DialectMySQL:
				assertContains(t, sqlUpdateQuery, "UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ?", "UPDATE WITHOUT RETURNING")
			case DialectPostgreSQL:
				assertContains(t, sqlUpdateQuery, `UPDATE "users" AS "u" SET "u"."string" = $1 WHERE "u"."number" = $2 RETURNING "u"."id", "u"."string"`, "UPDATE RETURNING")
			case DialectSQLite:
				assertContains(t, sqlUpdateQuery, `UPDATE "users" AS "u" SET "u"."string" = ? WHERE "u"."number" = ? RETURNING "u"."id", "u"."string"`, "UPDATE RETURNING")
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
			stmtUpdate := NewUpdate(Test.Table.User).
				Set(
					Assign(Test.Table.Users.String.Expr(), Value("active")),
				).
				Where(
					Equal(Test.Table.Users.Number.Expr(), Value(2)),
				)
			sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlUpdateQuery, "UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ?", "UPDATE")
			case DialectMsSQL:
				assertContains(t, sqlUpdateQuery, "UPDATE [users] AS [u] SET [u].[string] = @p1 WHERE [u].[number] = @p2", "UPDATE")
			case DialectMySQL:
				assertContains(t, sqlUpdateQuery, "UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ?", "UPDATE")
			case DialectPostgreSQL:
				assertContains(t, sqlUpdateQuery, `UPDATE "users" AS "u" SET "u"."string" = $1 WHERE "u"."number" = $2`, "UPDATE")
			case DialectSQLite:
				assertContains(t, sqlUpdateQuery, `UPDATE "users" AS "u" SET "u"."string" = ? WHERE "u"."number" = ?`, "UPDATE")
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
			stmtUpdate := NewUpdate(Test.Table.User).
				Set(
					Assign(Test.Table.Users.String.Expr(), Value("active")),
				).
				Where(
					Equal(Test.Table.Users.Number.Expr(), Value(2)),
				)
			sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlUpdateQuery, "UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ?", "UPDATE")
			case DialectMsSQL:
				assertContains(t, sqlUpdateQuery, "UPDATE [users] AS [u] SET [u].[string] = @p1 WHERE [u].[number] = @p2", "UPDATE")
			case DialectMySQL:
				assertContains(t, sqlUpdateQuery, "UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ?", "UPDATE")
			case DialectPostgreSQL:
				assertContains(t, sqlUpdateQuery, `UPDATE "users" AS "u" SET "u"."string" = $1 WHERE "u"."number" = $2`, "UPDATE")
			case DialectSQLite:
				assertContains(t, sqlUpdateQuery, `UPDATE "users" AS "u" SET "u"."string" = ? WHERE "u"."number" = ?`, "UPDATE")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
		})
	})
	t.Run("With", func(t *testing.T) {
		testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
			sql := NewSQL(WithDialect(supportDialect))
			defer sql.Close()
			stmtUpdate := NewUpdate(Test.Table.User).
				Set(
					Assign(Test.Table.Users.String.Expr(), Value("updated")),
				).
				With(
					WithN("old_users", NewSelect(Test.Table.User).
						Fields(
							Test.Table.Users.ID.Expr(),
						).
						Where(
							Less(Test.Table.Users.Number.Expr(), Value(2)),
						),
					),
				)
			sqlUpdateQuery, sqlUpdateArguments, err := sql.Build(stmtUpdate)
			switch supportDialect {
			case DialectMariaDB:
				assertContains(t, sqlUpdateQuery, "WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) UPDATE `users` AS `u` SET `u`.`string` = ?", "UPDATE WITH")
			case DialectMsSQL:
				assertContains(t, sqlUpdateQuery, "WITH [old_users] AS (SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] < @p1) UPDATE [users] AS [u] SET [u].[string] = @p2", "UPDATE WITH")
			case DialectMySQL:
				assertContains(t, sqlUpdateQuery, "WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) UPDATE `users` AS `u` SET `u`.`string` = ?", "UPDATE WITH")
			case DialectPostgreSQL:
				assertContains(t, sqlUpdateQuery, `WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < $1) UPDATE "users" AS "u" SET "u"."string" = $2`, "UPDATE WITH")
			case DialectSQLite:
				assertContains(t, sqlUpdateQuery, `WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < ?) UPDATE "users" AS "u" SET "u"."string" = ?`, "UPDATE WITH")
			}
			t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlUpdateArguments, supportDialect.name, sqlUpdateQuery)
		})
	})
}
func Test_Transformer_Comparison(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(WithDialect(supportDialect))
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table.User).
			Fields(Test.Table.Users.ID.Expr()).
			Where(
				And(
					ILike(Test.Table.Users.String.Expr(), Value("%alex%")),
					And(
						ILike(Test.Table.Users.String.Expr(), Value("%ivan%")),
						ILike(Test.Table.Users.String.Expr(), Value("%petr%")),
					),
				),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "(LOWER(`u`.`string`) LIKE LOWER(?) AND LOWER(`u`.`string`) LIKE LOWER(?) AND LOWER(`u`.`string`) LIKE LOWER(?))", "ILIKE")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "(LOWER([u].[string]) LIKE LOWER(@p1) AND LOWER([u].[string]) LIKE LOWER(@p2) AND LOWER([u].[string]) LIKE LOWER(@p3))", "ILIKE")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "(LOWER(`u`.`string`) LIKE LOWER(?) AND LOWER(`u`.`string`) LIKE LOWER(?) AND LOWER(`u`.`string`) LIKE LOWER(?))", "ILIKE")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `("u"."string" ILIKE $1 AND "u"."string" ILIKE $2 AND "u"."string" ILIKE $3)`, "ILIKE")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `(LOWER("u"."string") LIKE LOWER(?) AND LOWER("u"."string") LIKE LOWER(?) AND LOWER("u"."string") LIKE LOWER(?))`, "ILIKE")
		}
		t.Logf("\nerr: [%s] \nsdn: %s \nsqa: [%s] \nsql: [%s]", err, sqlSelectArguments, supportDialect.name, sqlSelectQuery)
	})
}
func Test_Transformer_Function(t *testing.T) {
	testAllDialects(t, func(t *testing.T, supportDialect *SupportDialect) {
		sql := NewSQL(WithDialect(supportDialect))
		defer sql.Close()
		stmtSelect := NewSelect(Test.Table.User).
			Fields(
				Trunc(Ceil(Test.Table.Users.Number.Expr()), Value(2)).As("result"),
			)
		sqlSelectQuery, sqlSelectArguments, err := sql.Build(stmtSelect)
		switch supportDialect {
		case DialectMariaDB:
			assertContains(t, sqlSelectQuery, "TRUNCATE(CEILING(`u`.`number`), ?)", "TRUNC→TRUNCATE, CEIL→CEILING")
		case DialectMsSQL:
			assertContains(t, sqlSelectQuery, "ROUND(CEILING([u].[number]), @p1, 1)", "TRUNC→TRUNCATE, CEIL→CEILING")
		case DialectMySQL:
			assertContains(t, sqlSelectQuery, "TRUNCATE(CEILING(`u`.`number`), ?)", "TRUNC→TRUNCATE, CEIL→CEILING")
		case DialectPostgreSQL:
			assertContains(t, sqlSelectQuery, `TRUNC(CEIL("u"."number"), $1)`, "TRUNC→TRUNC, CEIL→CEIL")
		case DialectSQLite:
			assertContains(t, sqlSelectQuery, `TRUNC(CEIL("u"."number"), ?)`, "TRUNC→TRUNC, CEIL→CEIL")
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
	// Schema
	Test.Schema = NewSchema("test")
	// Table
	Test.Table.Order = NewTable("orders", "o")
	Test.Table.User = NewTable("users", "u")
	// Column
	Test.Table.Orders.Date = NewColumn[time.Time]("date", Test.Table.Order, TypeDate)
	Test.Table.Orders.ID = NewColumn[int64]("id", Test.Table.Order, TypeBigInt)
	Test.Table.Orders.Json = NewColumn[string]("json", Test.Table.Order, TypeJSON)
	Test.Table.Orders.Name = NewColumn[string]("string", Test.Table.Order, TypeVarChar)
	Test.Table.Orders.Number = NewColumn[int]("number", Test.Table.Order, TypeInt)
	Test.Table.Orders.String = NewColumn[string]("string", Test.Table.Order, TypeVarChar)
	Test.Table.Orders.Time = NewColumn[time.Time]("time", Test.Table.Order, TypeTime)
	Test.Table.Orders.UserID = NewColumn[int64]("data_id", Test.Table.Order, TypeBigInt)
	Test.Table.Users.CreateAt = NewColumn[time.Time]("createat", Test.Table.User, TypeTimestamp)
	Test.Table.Users.DataID = NewColumn[int64]("data_id", Test.Table.User, TypeBigInt)
	Test.Table.Users.Date = NewColumn[time.Time]("date", Test.Table.User, TypeDate)
	Test.Table.Users.ID = NewColumn[int64]("id", Test.Table.User, TypeBigInt)
	Test.Table.Users.Json = NewColumn[string]("json", Test.Table.User, TypeJSON)
	Test.Table.Users.Name = NewColumn[string]("name", Test.Table.User, TypeVarChar)
	Test.Table.Users.Number = NewColumn[int]("number", Test.Table.User, TypeInt)
	Test.Table.Users.String = NewColumn[string]("string", Test.Table.User, TypeVarChar)
	Test.Table.Users.UpdateAt = NewColumn[time.Time]("updateat", Test.Table.User, TypeTimestamp)
	Test.Table.Users.X = NewColumn[int]("x", Test.Table.User, TypeInt)
	Test.Table.Users.Y = NewColumn[int]("y", Test.Table.User, TypeInt)
	// Index
	Test.Index.OrdersID = NewIndex("orders_id", Test.Table.Order)
	Test.Index.UsersID = NewIndex("users_id", Test.Table.User)
	// View
	Test.View.OrdersGeneral = NewView("orders_general", "og", Test.Table.User)
	Test.View.UsersGeneral = NewView("users_general", "ug", Test.Table.User)
	// Constraint
	Test.Check.OrdersNumber = NewCheck("ck_orders_number", Greater(Test.Table.Orders.Number.Expr(), Value(0)))
	Test.Check.UsersNumber = NewCheck("ck_users_number", Greater(Test.Table.Users.Number.Expr(), Value(0)))
	Test.Foreign.OrdersUsers = NewForeignKey("fk_orders_users", Test.Table.User, Cascade(), Restrict(), Relation(Test.Table.Orders.UserID, Test.Table.Users.ID), Relation(Test.Table.Orders.Name, Test.Table.Users.String))
	Test.Foreign.UsersOrders = NewForeignKey("fk_users_orders", Test.Table.Order, Cascade(), Restrict(), Relation(Test.Table.Users.DataID, Test.Table.Orders.ID), Relation(Test.Table.Users.Name, Test.Table.Orders.String))
	Test.Primary.OrdersID = NewPrimaryKey("pk_orders_id", Test.Table.Orders.ID)
	Test.Primary.UsersID = NewPrimaryKey("pk_users_id", Test.Table.Users.ID)
	Test.Unique.OrdersName = NewUnique("un_orders_name", Test.Table.Orders.Name)
	Test.Unique.UsersName = NewUnique("un_users_name", Test.Table.Users.Name)
}

func testAllDialects(t *testing.T, userFunc func(t *testing.T, supportDialect *SupportDialect)) {
	for _, supportDialect := range listSupportDialects {
		currentDialect := supportDialect
		t.Run(currentDialect.name, func(t *testing.T) {
			userFunc(t, currentDialect)
		})
	}
}
