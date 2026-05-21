---
outline: deep
---

# API / SQL / Methods

::: info **Info**
This page documents methods available on statement instances: `Delete`, `Insert`, `Select`, `Update`. Each method configures a specific clause and returns the statement for chaining. Every method is shown with a working code example and expected SQL output.
:::

## stmtDelete
### Returning
Adds a RETURNING clause to return deleted rows. Supported by PostgreSQL. MySQL does not support this clause natively.
```go
stmtDelete := uast.NewDelete(uast.NewTable("test").As("t")).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Returning(
        uast.Column[int64]("t", "id"),
    )
```
Output MariaDB:
```text
DELETE FROM `test` AS `t` WHERE `t`.`string` = ? RETURNING `t`.`id`
```
Output MySQL:
```text
// Not supported
```
Output PostgreSQL:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = $1 RETURNING "t"."id"
```
Output SQLite:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = ? RETURNING "t"."id"
```

### Where
Adds a WHERE clause to filter rows for deletion. Accepts comparison expressions, logical operators, and subqueries.
```go
stmtDelete := uast.NewDelete(uast.NewTable("test").As("t")).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    )
```
Output MariaDB:
```text
DELETE FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output MySQL:
```text
DELETE FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = $1
```
Output SQLite:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = ?
```

### With
Adds a Common Table Expression (CTE) to the DELETE statement using `WithN` (non-recursive) or `WithR` (recursive).
```go
stmtWithN := uast.WithN("cte_nr", uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[int64]("t", "id"),
    ).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("old")),
    ),
)
stmtDelete := uast.NewDelete(uast.NewTable("test").As("t")).
    Where(
        uast.In(uast.Column[int64]("t", "id"), uast.Subquery[int64](uast.NewSelect(uast.NewTable("test").As("t")).Field(uast.Column[int64]("cte_nr", "id")))),
    ).
    With(
		stmtWithN,
	)
```
Output MariaDB:
```text
WITH cte_nr AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`string` = ?) DELETE FROM `test` AS `t` WHERE `t`.`id` IN (SELECT `cte_nr`.`id` FROM `test` AS `t`)
```
Output MySQL:
```text
WITH cte_nr AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`string` = ?) DELETE FROM `test` AS `t` WHERE `t`.`id` IN (SELECT `cte_nr`.`id` FROM `test` AS `t`)
```
Output PostgreSQL:
```text
WITH cte_nr AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."string" = $1) DELETE FROM "test" AS "t" WHERE "t"."id" IN (SELECT "cte_nr"."id" FROM "test" AS "t")
```
Output SQLite:
```text
WITH cte_nr AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."string" = ?) DELETE FROM "test" AS "t" WHERE "t"."id" IN (SELECT "cte_nr"."id" FROM "test" AS "t")
```

## stmtInsert
### Returning
Adds a RETURNING clause to return inserted rows. Supported by PostgreSQL. MySQL does not support this clause natively.
```go
stmtInsert := uast.NewInsert(uast.NewTable("test").As("t")).
    Values(
        uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
		uast.Pair(uast.Column[int]("t", "number"), uast.Value(2)),
    ).
    Returning(
        uast.Column[int64]("t", "id"),
    )
```
Output MariaDB:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?) RETURNING `t`.`id`
```
Output MySQL:
```text
// Not supported
```
Output PostgreSQL:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES ($1, $2) RETURNING "t"."id"
```
Output SQLite:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES (?, ?) RETURNING "t"."id"
```

### Source
Specifies a subquery as the data source for INSERT. Used for `INSERT ... SELECT` statements. When using `Source`, columns are inferred from the subquery fields.
```go
stmtInsert := NewInsert(uast.NewTable("test").As("t")).
	Source(uast.NewSelect(uast.NewTable("test1").As("t1")).
		Field(
			uast.Column[string]("t1", "string"),
			uast.Column[int]("t1", "number"),
		).
		Where(
			uast.Equal(uast.Column[string]("t1", "string"), uast.Value("active")),
		),
	)
```
Output MariaDB:
```text
INSERT INTO `test` AS `t` (`string`, `number`) SELECT `t1`.`string`, `t1`.`number` FROM `test1` AS `t1` WHERE `t1`.`string` = ?
```
Output MySQL:
```text
INSERT INTO `test` AS `t` (`string`, `number`) SELECT `t1`.`string`, `t1`.`number` FROM `test1` AS `t1` WHERE `t1`.`string` = ?
```
Output PostgreSQL:
```text
INSERT INTO "test" AS "t" ("string", "number") SELECT "t1"."string", "t1"."number" FROM "test1" AS "t1" WHERE "t1"."string" = $1
```
Output SQLite:
```text
INSERT INTO "test" AS "t" ("string", "number") SELECT "t1"."string", "t1"."number" FROM "test1" AS "t1" WHERE "t1"."string" = ?
```

### With
Adds a Common Table Expression (CTE) to the INSERT statement.
```go
stmtWithN := WithN("cte_nr", NewSelect(uast.NewTable("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
		uast.Column[string]("t", "string"),
	).
	Where(
		uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
	),
	"id", "string",
)
stmtInsert := NewInsert(uast.NewTable("test").As("t")).
	Source(
		NewSelect(NewCTE("cte_nr", "ctenr")).
			Field(
				uast.Column[int64]("ctenr", "id"),
				uast.Column[string]("ctenr", "string"),
			),
	).
	With(
		stmtWithN,
	)
```
Output MariaDB:
```text
WITH `cte_nr` (`id`, `string`) AS (SELECT `t1`.`id`, `t1`.`string` FROM `test1` AS `t1` WHERE `t1`.`string` = ?) INSERT INTO `test` AS `t` (`id`, `string`) SELECT `cte_nr`.`id`, `cte_nr`.`string` FROM `cte_nr` AS `ctenr`
```
Output MySQL:
```text
WITH `cte_nr` (`id`, `string`) AS (SELECT `t1`.`id`, `t1`.`string` FROM `test1` AS `t1` WHERE `t1`.`string` = ?) INSERT INTO `test` AS `t` (`id`, `string`) SELECT `cte_nr`.`id`, `cte_nr`.`string` FROM `cte_nr` AS `ctenr`
```
Output PostgreSQL:
```text
WITH "cte_nr" ("id", "string") AS (SELECT "t1"."id", "t1"."string" FROM "test1" AS "t1" WHERE "t1"."string" = $1) INSERT INTO "test" AS "t" ("id", "string") SELECT "cte_nr"."id", "cte_nr"."string" FROM "cte_nr" AS "ctenr"
```
Output SQLite:
```text
WITH "cte_nr" ("id", "string") AS (SELECT "t1"."id", "t1"."string" FROM "test1" AS "t1" WHERE "t1"."string" = ?) INSERT INTO "test" AS "t" ("id", "string") SELECT "cte_nr"."id", "cte_nr"."string" FROM "cte_nr" AS "ctenr"
```

## stmtSelect
### Distinct
Adds the DISTINCT modifier to remove duplicate rows from the result set.
```go
stmtSelect := NewSelect(uast.NewTable("test").As("t")).
    Distinct().
	Field(
		uast.Column[int64]("t", "id"),
	)
```
Output MariaDB:
```text
SELECT DISTINCT `t`.`id` FROM `test` AS `t`
```
Output MySQL:
```text
SELECT DISTINCT `t`.`id` FROM `test` AS `t`
```
Output PostgreSQL:
```text
SELECT DISTINCT "t"."id" FROM "test" AS "t"
```
Output SQLite:
```text
SELECT DISTINCT "t"."id" FROM "test" AS "t"
```

### Where
Adds a WHERE clause to filter rows before grouping or aggregation.
```go
stmtSelect := NewSelect(uast.NewTable("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	Where(
		uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
	)
```
Output MariaDB:
```text
SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output MySQL:
```text
SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
SELECT "t"."id" FROM "test" AS "t" WHERE "t"."string" = $1
```
Output SQLite:
```text
SELECT "t"."id" FROM "test" AS "t" WHERE "t"."string" = ?
```

### With
Adds a Common Table Expression (CTE) to the SELECT statement.
```go
stmtWithN := WithN("cte_nr", NewSelect(uast.NewTable("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
		uast.Column[string]("t", "string"),
	).
	Where(
		uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
	),
	"id", "string",
)
stmtSelect := NewSelect(uast.NewTable("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
		uast.Column[int]("t", "number"),
	).
	Join(
		uast.Inner(uast.NewCTE("cte_nr", "ctenr"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("ctenr", "id"))),
	).
	With(
		stmtWithN,
	)
```
Output MariaDB:
```text
WITH `cte_nr` (`id`, `string`) AS (SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ?) SELECT `t`.`id`, `t`.`number` FROM `test` AS `t` INNER JOIN `cte_nr` AS `ctenr` ON `t`.`id` = `ctenr`.`id`
```
Output MySQL:
```text
WITH `cte_nr` (`id`, `string`) AS (SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ?) SELECT `t`.`id`, `t`.`number` FROM `test` AS `t` INNER JOIN `cte_nr` AS `ctenr` ON `t`.`id` = `ctenr`.`id`
```
Output PostgreSQL:
```text
WITH "cte_nr" ("id", "string") AS (SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = $1) SELECT "t"."id", "t"."number" FROM "test" AS "t" INNER JOIN "cte_nr" AS "ctenr" ON "t"."id" = "ctenr"."id"
```
Output SQLite:
```text
WITH "cte_nr" ("id", "string") AS (SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = ?) SELECT "t"."id", "t"."number" FROM "test" AS "t" INNER JOIN "cte_nr" AS "ctenr" ON "t"."id" = "ctenr"."id"
```

## stmtUpdate
### Returning
Adds a RETURNING clause to return updated rows. Supported by PostgreSQL.
```go
stmtUpdate := NewUpdate(uast.NewTable("test").As("t")).
	Set(
		uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
	).
	Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	).
	Returning(
		uast.Column[int64]("t", "id"),
	)
```
Output MariaDB:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ? RETURNING `t`.`id`
```
Output MySQL:
```text
// Not supported
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2 RETURNING "t"."id"
```
Output SQLite:
```text
UPDATE "test" AS "t" SET "t"."string" = ? WHERE "t"."number" = ? RETURNING "t"."id"
```

### Where
Adds a WHERE clause to filter rows for updating.
```go
stmtUpdate := NewUpdate(uast.NewTable("test").As("t")).
	Set(
		uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
	).
	Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	)
```
Output MariaDB:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
```
Output MySQL:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2
```
Output SQLite:
```text
UPDATE "test" AS "t" SET "t"."string" = ? WHERE "t"."number" = ?
```

### With
Adds a Common Table Expression (CTE) to the UPDATE statement.
```go
stmtWithN := WithN("cte_nr", NewSelect(uast.NewTable("test").As("t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	Where(
		uast.Equal(uast.Column[string]("t", "string"), uast.Value("pending")),
	),
)
stmtUpdate := NewUpdate(uast.NewTable("test").As("t")).
	Set(
		uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
	).
	Where(
		uast.In(uast.Column[int64]("t", "id"), Subquery[int64](NewSelect(uast.NewTable("test").As("t")).Field(Column[int64]("cte_nr", "id")))),
	).
	With(
		stmtWithN,
	)
```
Output MariaDB:
```text
WITH `cte_nr` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`string` = ?) UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`id` IN (SELECT `cte_nr`.`id` FROM `test` AS `t`)
```
Output MySQL:
```text
WITH `cte_nr` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`string` = ?) UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`id` IN (SELECT `cte_nr`.`id` FROM `test` AS `t`)
```
Output PostgreSQL:
```text
WITH "cte_nr" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."string" = $1) UPDATE "test" AS "t" SET "t"."string" = $2 WHERE "t"."id" IN (SELECT "cte_nr"."id" FROM "test" AS "t")
```
Output SQLite:
```text
WITH "cte_nr" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."string" = ?) UPDATE "test" AS "t" SET "t"."string" = ? WHERE "t"."id" IN (SELECT "cte_nr"."id" FROM "test" AS "t")
```
