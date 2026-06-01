---
outline: deep
---

# API / 核心 / 构造函数

::: info **关于**
本页面涵盖四个语句构造函数：`NewDelete`、`NewInsert`、`NewSelect`、`NewUpdate`。每个构造函数创建一个新的语句实例，可使用方法进行配置，并通过 `Build()` 编译为 SQL。
:::

## NewComment
创建一个新的 COMMENT 语句实例。接受注释文本，并返回一个可以使用 `OnColumn` 或 `OnTable` 进行配置的语句。
```go
stmtCommentColumn := uast.NewComment("Test comment").
    OnColumn(uast.Column[int64]("t", "id"))
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
创建一个新的 DELETE 语句实例。接受一个表源，返回一个可使用 `Join`、`Returning`、`Where`、`With` 进行配置的语句。
```go
stmtDeleteJoin := uast.NewDelete(uast.NewTable("test", "t")).
    Join(
		Inner(uast.NewTable("data", "d"), Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("d", "id"))),
    ).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    )
stmtDeleteReturning := uast.NewDelete(uast.NewTable("test", "t")).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Returning(
		uast.Column[int64]("t", "id"),
		uast.Column[string]("t", "string"),
	)
stmtDeleteWhere := uast.NewDelete(uast.NewTable("test", "t")).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    )
stmtDeleteWith := NewDelete(uast.NewTable("test", "t")).
	With(
		uast.WithN("old_users", uast.NewSelect(uast.NewTable("test", "t")).
			Field(
                uast.Column[int64]("t", "id"),
            ).
			Where(
                Less(uast.Column[int]("t", "number"), Value(2)),
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
创建一个新的 DROP 语句实例。接受一个 `Index/Schema/Table/View` 源，并返回一个可使用 `Cascade()` 或 `IfExists()` 配置的语句。
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
创建一个新的 INSERT 语句实例。接受表源并返回一个可使用 `Returning`、`Source/Values`、`With`进行配置的语句。
```go
stmtInsertReturning := uast.NewInsert(uast.NewTable("test", "t")).
    Values(
		uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
		uast.Pair(uast.Column[int]("t", "number"), uast.Value(2)),
	).
	Returning(
		uast.Column[int64]("t", "id"),
		uast.Column[string]("t", "string"),
	)
stmtInsertSource := uast.NewInsert(uast.NewTable("test", "t")).
	Source(NewSelect(uast.NewTable("test", "t")).
		Field(
			uast.Column[string]("t", "string"),
			uast.Column[int]("t", "number"),
		).
		Where(
			uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
		),
	)
stmtInsertValues := uast.NewInsert(uast.NewTable("test", "t")).
    Values(
		uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
		uast.Pair(uast.Column[int]("t", "number"), uast.Value(2)),
	).
    Upsert(
		uast.Pair(uast.Column[string]("t", "string"), uast.Value("updated")),
	)
stmtInsertWith := NewInsert(uast.NewTable("test", "t")).
	Values(
		uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
		uast.Pair(uast.Column[int]("t", "number"), uast.Value(2)),
	).
	With(
		uast.WithN("old_users", uast.NewSelect(uast.NewTable("test", "t")).
			Field(
                uast.Column[int64]("t", "id"),
		    ).
			Where(
				uast.Less(uast.Column[int]("t", "number"), uast.Value(2)),
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
创建一个新的 SELECT 语句实例。接受字段，返回一个可使用 `Distinct`、`Field`、`GroupBy`、`Having`、`Join`、`OrderBy`、`Pagination`、`Unions`、`Where`、`With` 进行配置的语句。
```go
stmtSelectDistinct := uast.NewSelect(uast.NewTable("test", "t")).
    Distinct().
    Field(
        uast.Column[int64]("t", "id"),
    ).
    Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	)
stmtSelectField := uast.NewSelect(uast.NewTable("test", "t")).
    Field(
        uast.Column[int64]("t", "id"),
    ).
    Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	)
stmtSelectGroupBy := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Column[string]("t", "string"),
		uast.Count(uast.Column[int64]("t", "id"), false).As("cnt"),
	).
	GroupBy(
		uast.Column[string]("t", "string"),
	)
stmtSelectHaving := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Column[string]("t", "string"),
		uast.Count(uast.Column[int64]("t", "id"), false).As("cnt"),
	).
	GroupBy(
        uast.Column[string]("t", "string"),
    ).
	Having(
		uast.Greater(uast.Count(uast.Column[int64]("t", "id"), false), uast.Value[int64](2)),
	)
stmtSelectJoin := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Column[int64]("t", "id"),
		uast.Column[string]("d", "string"),
	).
	Join(
		uast.Inner(uast.NewTable("data", "d"), Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("d", "id"))),
	)
stmtSelectOrderBy := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	OrderBy(
		uast.Desc(uast.Column[int]("t", "number")),
		uast.Asc(uast.Column[string]("t", "string")),
	)
stmtSelectPagination := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	Pagination(10, 20)
stmtSelectUnions := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Column[string]("t", "string"),
	).
	Unions(
		uast.UnionAll(uast.NewSelect(uast.NewTable("data", "d")).
			Field(
				uast.Column[string]("d", "string"),
			),
		),
	)
stmtSelectWhere := uast.NewSelect(uast.NewTable("test", "t")).
	Field(
		uast.Column[int64]("t", "id"),
	).
	Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	)
stmtSelectWith := uast.NewSelect(uast.NewCTE("cte_test", "ct")).
	Field(
		uast.Column[int64]("ct", "id"),
	).
	With(
		uast.WithN("cte_test", uast.NewSelect(uast.NewTable("test", "t")).
			Field(
                uast.Column[int64]("t", "id"),
            ).
			Where(
                uast.Greater(uast.Column[int]("t", "number"), uast.Value(2)),
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
创建一个新的 TRUNCATE 语句实例。接受表源，并返回一个可以使用 `Cascade()` 或 `RestartIdentity()` 进行配置的语句。
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
创建一个新的 UPDATE 语句实例。接受表源并返回一个可使用 `Join`、`Returning`、`Set`、`Where`、`With` 进行配置的语句。
```go
stmtUpdateJoin := uast.NewUpdate(uast.NewTable("test", "t")).
    Join(
		uast.Inner(uast.NewTable("data", "d"), Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("d", "id"))),
    ).
    Set(
        uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
    ).
stmtUpdateReturning := uast.NewUpdate(uast.NewTable("test", "t")).
    Set(
        uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
    ).
    Returning(
        uast.Column[int64]("t", "id"),
        uast.Column[string]("t", "string")
    )
stmtUpdateSet := uast.NewUpdate(uast.NewTable("test", "t")).
    Set(
        uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
    )
stmtUpdateWhere := uast.NewUpdate(uast.NewTable("test", "t")).
    Set(
        uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
    )
stmtUpdateWith := NewUpdate(uast.NewTable("test", "t")).
	Set(
		uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
	).
	With(
		uast.WithN("old_users", uast.NewSelect(uast.NewTable("test", "t")).
			Field(
                uast.Column[int64]("t", "id"),
            ).
			Where(
                uast.Less(uast.Column[int]("t", "number"), uast.Value(2)),
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