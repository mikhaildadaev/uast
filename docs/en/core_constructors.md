---
outline: deep
---

# API / Core / Constructors

::: info **Info**
This page covers the four statement constructors: `NewDelete`, `NewInsert`, `NewSelect`, `NewUpdate`. Each constructor creates a new statement instance that can be configured with methods and built into SQL using `Build()`.
:::

## NewComment
Creates a new COMMENT statement instance. Accepts a comment text and returns a statement that can be configured with `OnColumn` or `OnTable`.
```go
stmtCommentColumn := uast.NewComment("Test comment").
    OnColumn(uast.Field[int64]("t", "id"))
stmtCommentTable := uast.NewComment("Test comment").
    OnTable(uast.NewTable("test", "t"))
```
Output MariaDB:
```text
COMMENT ON COLUMN `t`.`id` IS 'Test comment'
COMMENT ON TABLE `test` AS `t` IS 'Test comment'
```
Output MsSQL:
```text
// Not supported
// Not supported
```
Output MySQL:
```text
COMMENT ON COLUMN `t`.`id` IS 'Test comment'
COMMENT ON TABLE `test` AS `t` IS 'Test comment'
```
Output PostgreSQL:
```text
COMMENT ON COLUMN "t"."id" IS 'Test comment'
COMMENT ON TABLE "test" AS "t" IS 'Test comment'
```
Output SQLite:
```text
COMMENT ON COLUMN "t"."id" IS 'Test comment'
COMMENT ON TABLE "test" AS "t" IS 'Test comment'
```

## NewDelete
Creates a new DELETE statement instance. Accepts a table source and returns a statement that can be configured with `Join`, `Returning`, `Where`, `With`.
```go
stmtDeleteJoin := uast.NewDelete(uast.NewTable("test", "t")).
    Join(
		Inner(uast.NewTable("data", "d"), Equal(uast.Field[int64]("t", "id"), uast.Field[int64]("d", "id"))),
    ).
    Where(
        uast.Equal(uast.Field[string]("t", "string"), uast.Value("active")),
    )
stmtDeleteReturning := uast.NewDelete(uast.NewTable("test", "t")).
    Where(
        uast.Equal(uast.Field[string]("t", "string"), uast.Value("active")),
    ).
    Returning(
		uast.Field[int64]("t", "id"),
		uast.Field[string]("t", "string"),
	)
stmtDeleteWhere := uast.NewDelete(uast.NewTable("test", "t")).
    Where(
        uast.Equal(uast.Field[string]("t", "string"), uast.Value("active")),
    )
stmtDeleteWith := NewDelete(uast.NewTable("test", "t")).
	With(
		uast.WithN("old_users", uast.NewSelect(uast.NewTable("test", "t")).
			Field(
                uast.Field[int64]("t", "id"),
            ).
			Where(
                Less(uast.Field[int]("t", "number"), Value(2)),
            ),
		),
	)
```
Output MariaDB:
```text
DELETE `t` FROM `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id` WHERE `t`.`string` = ?
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ? RETURNING `t`.`id`, `t`.`string`
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?
WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) DELETE `t` FROM `test` AS `t`
```
Output MsSQL:
```text
DELETE [t] FROM [test] AS [t] INNER JOIN [data] AS [d] ON [t].[id] = [d].[id] WHERE [t].[string] = @p1
DELETE [t] FROM [test] AS [t] OUTPUT [t].[id], [t].[string] WHERE [t].[string] = @p1
DELETE [t] FROM [test] AS [t] WHERE [t].[string] = @p1
WITH [old_users] AS (SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] < @p1) DELETE [t] FROM [test] AS [t]
```
Output MySQL:
```text
DELETE `t` FROM `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id` WHERE `t`.`string` = ?
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?
WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) DELETE `t` FROM `test` AS `t`
```
Output PostgreSQL:
```text
DELETE FROM "test" AS "t" USING "data" AS "d" WHERE ("t"."id" = "d"."id" AND "t"."string" = $1)
DELETE FROM "test" AS "t" WHERE "t"."string" = $1 RETURNING "t"."id", "t"."string"
DELETE FROM "test" AS "t" WHERE "t"."string" = $1
WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < $1) DELETE FROM "test" AS "t"
```
Output SQLite:
```text
DELETE FROM "test" AS "t" INNER JOIN "data" AS "d" ON "t"."id" = "d"."id" WHERE "t"."string" = ?
DELETE FROM "test" AS "t" WHERE "t"."string" = ? RETURNING "t"."id", "t"."string"
DELETE FROM "test" AS "t" WHERE "t"."string" = ?
WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < ?) DELETE FROM "test" AS "t"
```

## NewDrop
Creates a new DROP statement instance. Accepts a `Index/Schema/Table/View` source and returns a statement that can be configured with `Cascade()` or `IfExists()`.
```go
stmtDropCascade := uast.NewDrop(uast.NewTable("test", "t")).
    Cascade()
stmtDropIfExistsIndex := uast.NewDrop(uast.NewIndex("test")).
    IfExists()
stmtDropIfExistsSchema := uast.NewDrop(uast.NewSchema("test")).
    IfExists()
stmtDropIfExistsTable := uast.NewDrop(uast.NewTable("test", "t")).
    IfExists()
stmtDropIfExistsView := uast.NewDrop(uast.NewView("test", "t")).
    IfExists()
```
Output MariaDB:
```text
DROP TABLE `test` CASCADE
DROP INDEX IF EXISTS `test`
DROP SCHEMA IF EXISTS `test`
DROP TABLE IF EXISTS `test`
DROP VIEW IF EXISTS `test`
```
Output MsSQL:
```text
DROP TABLE [test]
DROP INDEX [test]
DROP SCHEMA IF EXISTS [test]
DROP TABLE IF EXISTS [test]
DROP VIEW IF EXISTS [test]
```
Output MySQL:
```text
DROP TABLE `test`
DROP INDEX IF EXISTS `test`
DROP SCHEMA `test`
DROP TABLE IF EXISTS `test`
DROP VIEW IF EXISTS `test`
```
Output PostgreSQL:
```text
DROP TABLE "test" CASCADE
DROP INDEX IF EXISTS "test"
DROP SCHEMA IF EXISTS "test"
DROP TABLE IF EXISTS "test"
DROP VIEW IF EXISTS "test"
```
Output SQLite:
```text
DROP TABLE "test"
DROP INDEX IF EXISTS "test"
DROP SCHEMA "test"
DROP TABLE IF EXISTS "test"
DROP VIEW IF EXISTS "test"
```

## NewInsert
Creates a new INSERT statement instance. Accepts a table source and returns a statement that can be configured with `Returning`, `Source/Values`, `With`.
```go
stmtInsertReturning := uast.NewInsert(uast.NewTable("test", "t")).
    Values(
		uast.Pair(uast.Field[string]("t", "string"), uast.Value("ivan")),
		uast.Pair(uast.Field[int]("t", "number"), uast.Value(2)),
	).
	Returning(
		uast.Field[int64]("t", "id"),
		uast.Field[string]("t", "string"),
	)
stmtInsertSource := uast.NewInsert(uast.NewTable("test", "t")).
	Source(NewSelect(uast.NewTable("test", "t")).
		Field(
			uast.Field[string]("t", "string"),
			uast.Field[int]("t", "number"),
		).
		Where(
			uast.Equal(uast.Field[string]("t", "string"), uast.Value("active")),
		),
	)
stmtInsertValues := uast.NewInsert(uast.NewTable("test", "t")).
    Values(
		uast.Pair(uast.Field[string]("t", "string"), uast.Value("ivan")),
		uast.Pair(uast.Field[int]("t", "number"), uast.Value(2)),
	).
    Upsert(
		uast.Pair(uast.Field[string]("t", "string"), uast.Value("updated")),
	)
stmtInsertWith := NewInsert(uast.NewTable("test", "t")).
	Values(
		uast.Pair(uast.Field[string]("t", "string"), uast.Value("ivan")),
		uast.Pair(uast.Field[int]("t", "number"), uast.Value(2)),
	).
	With(
		uast.WithN("old_users", uast.NewSelect(uast.NewTable("test", "t")).
			Field(
                uast.Field[int64]("t", "id"),
		    ).
			Where(
				uast.Less(uast.Field[int]("t", "number"), uast.Value(2)),
			),
		),
	)
```
Output MariaDB:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?) RETURNING `t`.`id`, `t`.`string`
INSERT INTO `test` AS `t` (`string`, `number`) SELECT `t`.`string`, `t`.`number` FROM `test` AS `t` WHERE `t`.`string` = ?
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
```
Output MsSQL:
```text
INSERT INTO [test] AS [t] ([string], [number]) OUTPUT [t].[id], [t].[string] VALUES (@p1, @p2)
INSERT INTO [test] AS [t] ([string], [number]) SELECT [t].[string], [t].[number] FROM [test] AS [t] WHERE [t].[string] = @p1
INSERT INTO [test] AS [t] ([string], [number]) VALUES (@p1, @p2)
WITH [old_users] AS (SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] < @p1) INSERT INTO [test] AS [t] ([string], [number]) VALUES (@p2, @p3)
```
Output MySQL:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
INSERT INTO `test` AS `t` (`string`, `number`) SELECT `t`.`string`, `t`.`number` FROM `test` AS `t` WHERE `t`.`string` = ?
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
```
Output PostgreSQL:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES ($1, $2) RETURNING "t"."id", "t"."string"
INSERT INTO "test" AS "t" ("string", "number") SELECT "t"."string", "t"."number" FROM "test" AS "t" WHERE "t"."string" = $1
INSERT INTO "test" AS "t" ("string", "number") VALUES ($1, $2) ON CONFLICT DO UPDATE SET "string" = $3
WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < $1) INSERT INTO "test" AS "t" ("string", "number") VALUES ($2, $3)
```
Output SQLite:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES (?, ?) RETURNING "t"."id", "t"."string"
INSERT INTO "test" AS "t" ("string", "number") SELECT "t"."string", "t"."number" FROM "test" AS "t" WHERE "t"."string" = ?
INSERT INTO "test" AS "t" ("string", "number") VALUES (?, ?) ON CONFLICT DO UPDATE SET "string" = ?
WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < ?) INSERT INTO "test" AS "t" ("string", "number") VALUES (?, ?)
```

## NewSelect
Creates a new SELECT statement instance. Accepts a table source and returns a statement that can be configured with `Distinct`, `Field`, `GroupBy`, `Having`, `Join`, `OrderBy`, `Pagination`, `Unions`, `Where`, `With`.
```go
stmtSelectDistinct := uast.NewSelect(uast.NewTable("test", "t")).
    Distinct().
    Field(
        uast.Field[int64]("t", "id"),
    ).
    Where(
		uast.Equal(uast.Field[int]("t", "number"), uast.Value(2)),
	)
stmtSelectField := uast.NewSelect(uast.NewTable("test", "t")).
    Field(
        uast.Field[int64]("t", "id"),
    ).
    Where(
		uast.Equal(uast.Field[int]("t", "number"), uast.Value(2)),
	)
stmtSelectGroupBy := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Field[string]("t", "string"),
		uast.Count(uast.Field[int64]("t", "id"), false).As("cnt"),
	).
	GroupBy(
		uast.Field[string]("t", "string"),
	)
stmtSelectHaving := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Field[string]("t", "string"),
		uast.Count(uast.Field[int64]("t", "id"), false).As("cnt"),
	).
	GroupBy(
        uast.Field[string]("t", "string"),
    ).
	Having(
		uast.Greater(uast.Count(uast.Field[int64]("t", "id"), false), uast.Value[int64](2)),
	)
stmtSelectJoin := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Field[int64]("t", "id"),
		uast.Field[string]("d", "string"),
	).
	Join(
		uast.Inner(uast.NewTable("data", "d"), Equal(uast.Field[int64]("t", "id"), uast.Field[int64]("d", "id"))),
	)
stmtSelectOrderBy := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Field[int64]("t", "id"),
	).
	OrderBy(
		uast.Desc(uast.Field[int]("t", "number")),
		uast.Asc(uast.Field[string]("t", "string")),
	)
stmtSelectPagination := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Field[int64]("t", "id"),
	).
	Pagination(10, 20)
stmtSelectUnions := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Field[string]("t", "string"),
	).
	Unions(
		uast.UnionAll(uast.NewSelect(uast.NewTable("data", "d")).
			Field(
				uast.Field[string]("d", "string"),
			),
		),
	)
stmtSelectWhere := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Field[int64]("t", "id"),
	).
	Where(
		uast.Equal(uast.Field[int]("t", "number"), uast.Value(2)),
	)
stmtSelectWith := uast.NewSelect(uast.NewCTE("cte_test", "ct")).
	Field(
		uast.Field[int64]("ct", "id"),
	).
	With(
		uast.WithN("cte_test", uast.NewSelect(uast.NewTable("test", "t")).
			Field(
                uast.Field[int64]("t", "id"),
            ).
			Where(
                uast.Greater(uast.Field[int]("t", "number"), uast.Value(2)),
            ),
		),
	)
```
Output MariaDB:
```text
SELECT DISTINCT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
SELECT `t`.`string`, COUNT(`t`.`id`) AS `cnt` FROM `test` AS `t` GROUP BY `t`.`string`
SELECT `t`.`string`, COUNT(`t`.`id`) AS `cnt` FROM `test` AS `t` GROUP BY `t`.`string` HAVING COUNT(`t`.`id`) > ?
SELECT `t`.`id`, `d`.`string` FROM `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id`
SELECT `t`.`id` FROM `test` AS `t` ORDER BY `t`.`number` DESC, `t`.`string` ASC
SELECT `t`.`id` FROM `test` AS `t` LIMIT ? OFFSET ?
SELECT `t`.`string` FROM `test` AS `t` UNION ALL SELECT `d`.`string` FROM `data` AS `d`
SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
WITH `cte_test` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` > ?) SELECT `ct`.`id` FROM `cte_test` AS `ct`
```
Output MsSQL:
```text
SELECT DISTINCT [t].[id] FROM [test] AS [t] WHERE [t].[number] = @p1
SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] = @p1
SELECT [t].[string], COUNT([t].[id]) AS [cnt] FROM [test] AS [t] GROUP BY [t].[string]
SELECT [t].[string], COUNT([t].[id]) AS [cnt] FROM [test] AS [t] GROUP BY [t].[string] HAVING COUNT([t].[id]) > @p1
SELECT [t].[id], [d].[string] FROM [test] AS [t] INNER JOIN [data] AS [d] ON [t].[id] = [d].[id]
SELECT [t].[id] FROM [test] AS [t] ORDER BY [t].[number] DESC, [t].[string] ASC
SELECT [t].[id] FROM [test] AS [t] ORDER BY 1 ASC OFFSET @p1 ROWS FETCH NEXT @p2 ROWS ONLY
SELECT [t].[string] FROM [test] AS [t] UNION ALL SELECT [d].[string] FROM [data] AS [d]
SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] = @p1
WITH [cte_test] AS (SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] > @p1) SELECT [ct].[id] FROM [cte_test] AS [ct]
```
Output MySQL:
```text
SELECT DISTINCT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
SELECT `t`.`string`, COUNT(`t`.`id`) AS `cnt` FROM `test` AS `t` GROUP BY `t`.`string`
SELECT `t`.`string`, COUNT(`t`.`id`) AS `cnt` FROM `test` AS `t` GROUP BY `t`.`string` HAVING COUNT(`t`.`id`) > ?
SELECT `t`.`id`, `d`.`string` FROM `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id`
SELECT `t`.`id` FROM `test` AS `t` ORDER BY `t`.`number` DESC, `t`.`string` ASC
SELECT `t`.`id` FROM `test` AS `t` LIMIT ? OFFSET ?
SELECT `t`.`string` FROM `test` AS `t` UNION ALL SELECT `d`.`string` FROM `data` AS `d`
SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
WITH `cte_test` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` > ?) SELECT `ct`.`id` FROM `cte_test` AS `ct`
```
Output PostgreSQL:
```text
SELECT DISTINCT "t"."id" FROM "test" AS "t" WHERE "t"."number" = $1
SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = $1
SELECT "t"."string", COUNT("t"."id") AS "cnt" FROM "test" AS "t" GROUP BY "t"."string"
SELECT "t"."string", COUNT("t"."id") AS "cnt" FROM "test" AS "t" GROUP BY "t"."string" HAVING COUNT("t"."id") > $1
SELECT "t"."id", "d"."string" FROM "test" AS "t" INNER JOIN "data" AS "d" ON "t"."id" = "d"."id"
SELECT "t"."id" FROM "test" AS "t" ORDER BY "t"."number" DESC, "t"."string" ASC
SELECT "t"."id" FROM "test" AS "t" LIMIT $1 OFFSET $2
SELECT "t"."string" FROM "test" AS "t" UNION ALL SELECT "d"."string" FROM "data" AS "d"
SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = $1
WITH "cte_test" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" > $1) SELECT "ct"."id" FROM "cte_test" AS "ct"
```
Output SQLite:
```text
SELECT DISTINCT "t"."id" FROM "test" AS "t" WHERE "t"."number" = ?
SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = ?
SELECT "t"."string", COUNT("t"."id") AS "cnt" FROM "test" AS "t" GROUP BY "t"."string"
SELECT "t"."string", COUNT("t"."id") AS "cnt" FROM "test" AS "t" GROUP BY "t"."string" HAVING COUNT("t"."id") > ?
SELECT "t"."id", "d"."string" FROM "test" AS "t" INNER JOIN "data" AS "d" ON "t"."id" = "d"."id"
SELECT "t"."id" FROM "test" AS "t" ORDER BY "t"."number" DESC, "t"."string" ASC
SELECT "t"."id" FROM "test" AS "t" LIMIT ? OFFSET ?
SELECT "t"."string" FROM "test" AS "t" UNION ALL SELECT "d"."string" FROM "data" AS "d"
SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = ?
WITH "cte_test" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" > ?) SELECT "ct"."id" FROM "cte_test" AS "ct"
```

## NewTruncate
Creates a new TRUNCATE statement instance. Accepts a table source and returns a statement that can be configured with `Cascade()` or `RestartIdentity()`.
```go
stmtTruncateDefault := uast.NewTruncate(uast.NewTable("test", "t"))
stmtTruncateCascade := uast.NewTruncate(uast.NewTable("test", "t")).
    Cascade()
stmtTruncateRestartIdentity := uast.NewTruncate(uast.NewTable("test", "t")).
    RestartIdentity()
```
Output MariaDB:
```text
TRUNCATE TABLE `test`
TRUNCATE TABLE `test` CASCADE
TRUNCATE TABLE `test` RESTART IDENTITY
```
Output MsSQL:
```text
TRUNCATE TABLE [test]
// Not supported
// Not supported
```
Output MySQL:
```text
TRUNCATE TABLE `test`
// Not supported
// Not supported
```
Output PostgreSQL:
```text
TRUNCATE TABLE "test"
TRUNCATE TABLE "test" CASCADE
TRUNCATE TABLE "test" RESTART IDENTITY
```
Output SQLite:
```text
TRUNCATE TABLE "test"
// Not supported
// Not supported
```

## NewUpdate
Creates a new UPDATE statement instance. Accepts a table source and returns a statement that can be configured with `Join`, `Returning`, `Set`, `Where`, `With`.
```go
stmtUpdateJoin := uast.NewUpdate(uast.NewTable("test", "t")).
    Join(
		uast.Inner(uast.NewTable("data", "d"), Equal(uast.Field[int64]("t", "id"), uast.Field[int64]("d", "id"))),
    ).
    Set(
        uast.Assign(uast.Field[string]("t", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Field[int]("t", "number"), uast.Value(2)),
    ).
stmtUpdateReturning := uast.NewUpdate(uast.NewTable("test", "t")).
    Set(
        uast.Assign(uast.Field[string]("t", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Field[int]("t", "number"), uast.Value(2)),
    ).
    Returning(
        uast.Field[int64]("t", "id"),
        uast.Field[string]("t", "string")
    )
stmtUpdateSet := uast.NewUpdate(uast.NewTable("test", "t")).
    Set(
        uast.Assign(uast.Field[string]("t", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Field[int]("t", "number"), uast.Value(2)),
    )
stmtUpdateWhere := uast.NewUpdate(uast.NewTable("test", "t")).
    Set(
        uast.Assign(uast.Field[string]("t", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Field[int]("t", "number"), uast.Value(2)),
    )
stmtUpdateWith := NewUpdate(uast.NewTable("test", "t")).
	Set(
		uast.Assign(uast.Field[string]("t", "string"), uast.Value("active")),
	).
	With(
		uast.WithN("old_users", uast.NewSelect(uast.NewTable("test", "t")).
			Field(
                uast.Field[int64]("t", "id"),
            ).
			Where(
                uast.Less(uast.Field[int]("t", "number"), uast.Value(2)),
            ),
		),
	)
```
Output MariaDB:
```text
UPDATE `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id` SET `t`.`string` = ? WHERE `d`.`string` = ?
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ? RETURNING `t`.`id`, `t`.`string`
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) UPDATE `test` AS `t` SET `t`.`string` = ?
```
Output MsSQL:
```text
UPDATE [test] AS [t] INNER JOIN [data] AS [d] ON [t].[id] = [d].[id] SET [t].[string] = @p1 WHERE [d].[string] = @p2
UPDATE [test] AS [t] OUTPUT [t].[id], [t].[string] SET [t].[string] = @p1 WHERE [t].[number] = @p2
UPDATE [test] AS [t] SET [t].[string] = @p1 WHERE [t].[number] = @p2
UPDATE [test] AS [t] SET [t].[string] = @p1 WHERE [t].[number] = @p2
WITH [old_users] AS (SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] < @p1) UPDATE [test] AS [t] SET [t].[string] = @p2
```
Output MySQL:
```text
UPDATE `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id` SET `t`.`string` = ? WHERE `d`.`string` = ?
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
WITH `old_users` AS (SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` < ?) UPDATE `test` AS `t` SET `t`.`string` = ?
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" INNER JOIN "data" AS "d" ON "t"."id" = "d"."id" SET "t"."string" = $1 WHERE "d"."string" = $2
UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2 RETURNING "t"."id", "t"."string"
UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2
UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2
WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < $1) UPDATE "test" AS "t" SET "t"."string" = $2
```
Output SQLite:
```text
UPDATE "test" AS "t" INNER JOIN "data" AS "d" ON "t"."id" = "d"."id" SET "t"."string" = ? WHERE "d"."string" = ?
UPDATE "test" AS "t" SET "t"."string" = ? WHERE "t"."number" = ? RETURNING "t"."id", "t"."string"
UPDATE "test" AS "t" SET "t"."string" = ? WHERE "t"."number" = ?
UPDATE "test" AS "t" SET "t"."string" = ? WHERE "t"."number" = ?
WITH "old_users" AS (SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" < ?) UPDATE "test" AS "t" SET "t"."string" = ?
```