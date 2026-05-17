---
outline: deep
---

# API / SQL / Methods

::: info **Info**
This page documents methods available on statement instances: `Delete`, `Insert`, `Select`, `Update`. Each method configures a specific clause and returns the statement for chaining. Every method is shown with a working code example and expected SQL output.
:::

## stmtDelete
### Join
Adds JOIN clauses to the DELETE statement. Supports INNER, LEFT, RIGHT, FULL, CROSS and their OUTER variants.
```go
stmtDelete := uast.NewDelete(uast.Table("test").As("t")).
    Join(
        uast.Inner(uast.Table("test1"), uast.Equal(uast.Column[int64]("t1", "id"), uast.Column[int64]("t", "id"))),
        uast.Left(uast.Table("test2"), uast.Equal(uast.Column[int64]("t2", "id"), uast.Column[int64]("t", "id"))),
    ).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    )
```
Output MySQL:
```text
DELETE `t` FROM `test` AS `t` INNER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id` LEFT JOIN `test2` AS `t2` ON `t2`.`id` = `t`.`id` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
DELETE FROM "test" AS "t" USING "test1" AS "t1", "test2" AS "t2" WHERE "t1"."id" = "t"."id" AND "t2"."id" = "t"."id" AND "t"."string" = $1
```

### Returning
Adds a RETURNING clause to return deleted rows. Supported by PostgreSQL. MySQL does not support this clause natively.
```go
stmtDelete := uast.NewDelete(uast.Table("test").As("t")).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Returning(
        uast.Column[int64]("t", "id"),
    )
```
Output MySQL:
```text

```
Output PostgreSQL:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = $1 RETURNING "t"."id"
```

### Where
Adds a WHERE clause to filter rows for deletion. Accepts comparison expressions, logical operators, and subqueries.
```go
stmtDelete := uast.NewDelete(uast.Table("test").As("t")).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    )
```
Output MySQL:
```text
DELETE FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = $1
```

### With
Adds a Common Table Expression (CTE) to the DELETE statement using `WithN` (non-recursive) or `WithR` (recursive).
```go
stmtWithN := uast.WithN("cte_nr", uast.NewSelect(uast.Table("test").As("t")).
    Field(
        uast.Column[int64]("t", "id"),
    ).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("old")),
    ),
)
stmtDelete := uast.NewDelete(uast.Table("test").As("t")).
    Where(
        uast.In(uast.Column[int64]("t", "id"), uast.Subquery[int64](uast.NewSelect(uast.Table("test").As("t")).Field(uast.Column[int64]("cte_nr", "id")))),
    ).
    With(
		stmtWithN,
	)
```
Output MySQL:
```text
WITH cte_nr AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`string` = ?) DELETE FROM `test` AS `t` WHERE `t`.`id` IN (SELECT `cte_nr`.`id` FROM `test` AS `t`)
```
Output PostgreSQL:
```text
WITH cte_nr AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."string" = $1) DELETE FROM "test" AS "t" WHERE "t"."id" IN (SELECT "cte_nr"."id" FROM "test" AS "t")
```

## stmtInsert
### Returning
Adds a RETURNING clause to return inserted rows. Supported by PostgreSQL. MySQL does not support this clause natively.
```go
stmtInsert := uast.NewInsert(uast.Table("test").As("t")).
    Values(
        uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
    ).
    Returning(
        uast.Column[int64]("t", "id"),
    )
```
Output MySQL:
```text
// Not supported
```
Output PostgreSQL:
```text
...
```

### Source
Specifies a subquery as the data source for INSERT. Used for `INSERT ... SELECT` statements. When using `Source`, columns are inferred from the subquery fields.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Values
Specifies values for insertion using `Pair` to associate columns with values. Columns are automatically inferred from the pairs.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### With
Adds a Common Table Expression (CTE) to the INSERT statement.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

## stmtSelect
### Distinct
Adds the DISTINCT modifier to remove duplicate rows from the result set.
```go
stmtSelect := NewSelect(uast.Table("test").As("t")).
    Distinct().
	Field(
		uast.Column[int64]("t", "id"),
	)
```
Output MySQL:
```text
SELECT DISTINCT `t`.`id` FROM `test` AS `t`
```
Output PostgreSQL:
```text
SELECT DISTINCT "t"."id" FROM "test" AS "t"
```

### Field
Specifies the fields to select. Accepts columns, functions, subqueries, and aliases.
```go
stmtSelect := NewSelect(uast.Table("test").As("t")).
    Field(
	    uast.Column[string]("t", "string"),
		uast.Count(uast.Column[int64]("t", "id"), false).As("count"),
	).
	GroupBy(
		uast.Column[string]("t", "string"),
	)
```
Output MySQL:
```text
SELECT `t`.`string`, COUNT(`t`.`id`) AS `count` FROM `test` AS `t` GROUP BY `t`.`string`
```
Output PostgreSQL:
```text
SELECT "t"."string", COUNT("t"."id") AS "count" FROM "test" AS "t" GROUP BY "t"."string"
```

### GroupBy
Adds a GROUP BY clause to group rows by specified columns or expressions.
```go
stmtSelect := NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[string]("t", "string"),
		uast.Count(uast.Column[int64]("t", "id"), false).As("count"),
	).
	GroupBy(
		uast.Column[string]("t", "string"),
	)
```
Output MySQL:
```text
SELECT `t`.`string`, COUNT(`t`.`id`) AS `count` FROM `test` AS `t` GROUP BY `t`.`string`
```
Output PostgreSQL:
```text
SELECT "t"."string", COUNT("t"."id") AS "count" FROM "test" AS "t" GROUP BY "t"."string"
```

### Having
Adds a HAVING clause to filter groups. Used with GROUP BY to filter aggregated results.
```go
stmtSelect := NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[string]("t", "string"),
		uast.Count(uast.Column[int64]("t", "id"), false).As("count"),
	).
	GroupBy(
        uast.Column[string]("t", "string"),
    ).
	Having(
		uast.Greater(uast.Count(uast.Column[int64]("t", "id"), false), uast.Value[int64](2)),
	)
```
Output MySQL:
```text
SELECT `t`.`string`, COUNT(`t`.`id`) AS `count` FROM `test` AS `t` GROUP BY `t`.`string` HAVING COUNT(`t`.`id`) > ?
```
Output PostgreSQL:
```text
SELECT "t"."string", COUNT("t"."id") AS "count" FROM "test" AS "t" GROUP BY "t"."string" HAVING COUNT("t"."id") > $1
```

### Join
Adds JOIN clauses to combine rows from multiple tables. Supports all 8 join types.
```go
stmtSelect := NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	Join(
		uast.Cross(uast.Table("test1")),
		uast.Full(uast.Table("test1"), uast.Equal(uast.Column[int64]("t1", "id"), uast.Column[int64]("t", "id"))),
		uast.FullOuter(uast.Table("test1"), uast.Equal(uast.Column[int64]("t1", "id"), uast.Column[int64]("t", "id"))),
		uast.Inner(uast.Table("test1"), uast.Equal(uast.Column[int64]("t1", "id"), uast.Column[int64]("t", "id"))),
		uast.Left(uast.Table("test1"), uast.Equal(uast.Column[int64]("t1", "id"), uast.Column[int64]("t", "id"))),
		uast.LeftOuter(uast.Table("test1"), uast.Equal(uast.Column[int64]("t1", "id"), uast.Column[int64]("t", "id"))),
		uast.Right(uast.Table("test1"), uast.Equal(uast.Column[int64]("t1", "id"), uast.Column[int64]("t", "id"))),
		uast.RightOuter(uast.Table("test1"), uast.Equal(uast.Column[int64]("t1", "id"), uast.Column[int64]("t", "id"))),
	)
```
Output MySQL:
```text
SELECT `t`.`id` FROM `test` AS `t` CROSS JOIN `test1` AS `t1` FULL JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id` FULL OUTER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id` INNER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id` LEFT JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id` LEFT OUTER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id` RIGHT JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id` RIGHT OUTER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id`
```
Output PostgreSQL:
```text
SELECT "t"."id" FROM "test" AS "t" CROSS JOIN "test1" AS "t1" FULL JOIN "test1" AS "t1" ON "t1"."id" = "t"."id" FULL OUTER JOIN "test1" AS "t1" ON "t1"."id" = "t"."id" INNER JOIN "test1" AS "t1" ON "t1"."id" = "t"."id" LEFT JOIN "test1" AS "t1" ON "t1"."id" = "t"."id" LEFT OUTER JOIN "test1" AS "t1" ON "t1"."id" = "t"."id" RIGHT JOIN "test1" AS "t1" ON "t1"."id" = "t"."id" RIGHT OUTER JOIN "test1" AS "t1" ON "t1"."id" = "t"."id"
```

### Limit
Limits the number of rows returned by the query.
```go
stmtSelect := NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	Limit(10)
```
Output MySQL:
```text
SELECT `t`.`id` FROM `test` AS `t` LIMIT ?
```
Output PostgreSQL:
```text
SELECT "t"."id" FROM "test" AS "t" LIMIT $1
```

### Offset
Skips a specified number of rows before returning results. Used for pagination with Limit.
```go
stmtSelect := NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	Offset(20)
```
Output MySQL:
```text
SELECT `t`.`id` FROM `test` AS `t` OFFSET ?
```
Output PostgreSQL:
```text
SELECT "t"."id" FROM "test" AS "t" OFFSET $1
```

### OrderBy
Adds an ORDER BY clause to sort results by specified columns or expressions in ascending or descending order.
```go
stmtSelect := NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	OrderBy(
		uast.Column[string]("t", "string"),
	)
```
Output MySQL:
```text
SELECT `t`.`id` FROM `test` AS `t` ORDER BY `t`.`string`
```
Output PostgreSQL:
```text
SELECT "t"."id" FROM "test" AS "t" ORDER BY "t"."string"
```

### Unions
Combines results from multiple SELECT statements using UNION, UNION ALL, EXCEPT, or INTERSECT.
```go
stmtWithR := WithR("cte_re", NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(0)),
	).
	Unions(
		uast.UnionAll(uast.NewSelect(uast.Table("test").As("t")).
			Field(
				uast.Column[int64]("t", "id"),
			).
			Join(
				uast.Inner(uast.CTE("cte_re", "ctere"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("ctere", "id"))),
			),
		),
	),
)
stmtUnion := NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[string]("t", "string"),
	)
stmtSelect := NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[string]("t", "string"),
	).
	Unions(
		uast.Union(stmtUnion),
		uast.UnionAll(stmtUnion),
		uast.UnionExcept(stmtUnion),
		uast.UnionIntersect(stmtUnion),
	).
	With(
		stmtWithR,
	)
```
Output MySQL:
```text
WITH RECURSIVE `cte_re` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ? UNION ALL SELECT `t`.`id` FROM `test` AS `t` INNER JOIN `cte_re` AS `ctere` ON `t`.`id` = `ctere`.`id`) SELECT `t`.`string` FROM `test` AS `t` UNION SELECT `t`.`string` FROM `test` AS `t` UNION ALL SELECT `t`.`string` FROM `test` AS `t` EXCEPT SELECT `t`.`string` FROM `test` AS `t` INTERSECT SELECT `t`.`string` FROM `test` AS `t`
```
Output PostgreSQL:
```text
WITH RECURSIVE "cte_re" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = $1 UNION ALL SELECT "t"."id" FROM "test" AS "t" INNER JOIN "cte_re" AS "ctere" ON "t"."id" = "ctere"."id") SELECT "t"."string" FROM "test" AS "t" UNION SELECT "t"."string" FROM "test" AS "t" UNION ALL SELECT "t"."string" FROM "test" AS "t" EXCEPT SELECT "t"."string" FROM "test" AS "t" INTERSECT SELECT "t"."string" FROM "test" AS "t"
```

### Where
Adds a WHERE clause to filter rows before grouping or aggregation.
```go
stmtSelect := NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	Where(
		uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
	)
```
Output MySQL:
```text
SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
SELECT "t"."id" FROM "test" AS "t" WHERE "t"."string" = $1
```

### With
Adds a Common Table Expression (CTE) to the SELECT statement.
```go
stmtWithN := WithN("cte_nr", NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
		uast.Column[string]("t", "string"),
	).
	Where(
		uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
	),
	"id", "string",
)
stmtSelect := NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
		uast.Column[int]("t", "number"),
	).
	Join(
		uast.Inner(uast.CTE("cte_nr", "ctenr"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("ctenr", "id"))),
	).
	With(
		stmtWithN,
	)
```
Output MySQL:
```text
WITH `cte_nr` (`id`, `string`) AS (SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ?) SELECT `t`.`id`, `t`.`number` FROM `test` AS `t` INNER JOIN `cte_nr` AS `ctenr` ON `t`.`id` = `ctenr`.`id`
```
Output PostgreSQL:
```text
WITH "cte_nr" ("id", "string") AS (SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = $1) SELECT "t"."id", "t"."number" FROM "test" AS "t" INNER JOIN "cte_nr" AS "ctenr" ON "t"."id" = "ctenr"."id"
```

## stmtUpdate
### Join
Adds JOIN clauses to the UPDATE statement for updating rows based on related table data.
```go
stmtUpdate := NewUpdate(uast.Table("test").As("t")).
	Set(
		uast.Pair(uast.Column[string]("t", "string"), uast.Value("active")),
	).
	Join(
		uast.Inner(uast.Column[string]("t1", "string"), uast.Equal(uast.Column[int64]("t1", "id"), uast.Column[int64]("t", "id"))),
	).
	Where(
		uast.Equal(uast.Column[string]("t1", "string"), uast.Value("active")),
	)
```
Output MySQL:
```text
UPDATE `test` AS `t` INNER JOIN `test1` AS `t1` ON `t1`.`id` = `t`.`id` SET `t`.`string` = ? WHERE `t1`.`string` = ?
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" INNER JOIN "test1" AS "t1" ON "t1"."id" = "t"."id" SET "t"."string" = $1 WHERE "t1"."string" = $2
```

### Returning
Adds a RETURNING clause to return updated rows. Supported by PostgreSQL.
```go
stmtUpdate := NewUpdate(uast.Table("test").As("t")).
	Set(
		uast.Pair(uast.Column[string]("t", "string"), uast.Value("active")),
	).
	Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	).
	Returning(
		uast.Column[int64]("t", "id"),
	)
```
Output MySQL:
```text
// Not supported
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2 RETURNING "t"."id"
```

### Set
Specifies columns and their new values using `Pair` to associate columns with values. Supports multiple pairs for updating multiple columns.
```go
stmtUpdate := NewUpdate(uast.Table("test").As("t")).
	Set(
		uast.Pair(uast.Column[string]("t", "string"), uast.Value("active")),
	).
	Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	)
```
Output MySQL:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2
```

### Where
Adds a WHERE clause to filter rows for updating.
```go
stmtUpdate := NewUpdate(uast.Table("test").As("t")).
	Set(
		uast.Pair(uast.Column[string]("t", "string"), uast.Value("active")),
	).
	Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	)
```
Output MySQL:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2
```

### With
Adds a Common Table Expression (CTE) to the UPDATE statement.
```go
stmtWithN := WithN("cte_nr", NewSelect(uast.Table("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	Where(
		uast.Equal(uast.Column[string]("t", "string"), uast.Value("pending")),
	),
)
stmtUpdate := NewUpdate(uast.Table("test").As("t")).
	Set(
		uast.Pair(uast.Column[string]("t", "string"), uast.Value("active")),
	).
	Where(
		uast.In(uast.Column[int64]("t", "id"), Subquery[int64](NewSelect(uast.Table("test").As("t")).Field(Column[int64]("cte_nr", "id")))),
	).
	With(
		stmtWithN,
	)
```
Output MySQL:
```text
WITH `cte_nr` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`string` = ?) UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`id` IN (SELECT `cte_nr`.`id` FROM `test` AS `t`)
```
Output PostgreSQL:
```text
WITH "cte_nr" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."string" = $1) UPDATE "test" AS "t" SET "t"."string" = $2 WHERE "t"."id" IN (SELECT "cte_nr"."id" FROM "test" AS "t")
```
