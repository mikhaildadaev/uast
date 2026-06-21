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
stmtCommentColumn := uast.NewComment("Comment").
    OnColumn(uast.Field[int64]("u", "id"))
stmtCommentTable := uast.NewComment("Comment").
    OnTable(uast.NewTable("users", "u"))
```
Output MariaDB:
```text
COMMENT ON COLUMN `u`.`id` IS 'Comment'
COMMENT ON TABLE `users` AS `u` IS 'Comment'
```
Output MsSQL:
```text
// Not supported
// Not supported
```
Output MySQL:
```text
COMMENT ON COLUMN `u`.`id` IS 'Comment'
COMMENT ON TABLE `users` AS `u` IS 'Comment'
```
Output PostgreSQL:
```text
COMMENT ON COLUMN "u"."id" IS 'Comment'
COMMENT ON TABLE "users" AS "u" IS 'Comment'
```
Output SQLite:
```text
COMMENT ON COLUMN "u"."id" IS 'Comment'
COMMENT ON TABLE "users" AS "u" IS 'Comment'
```

## NewDelete
创建一个新的 DELETE 语句实例。接受一个表源，返回一个可使用 `Join`、`Returning`、`Where`、`With` 进行配置的语句。
```go
stmtDeleteJoin := uast.NewDelete(uast.NewTable("users", "u")).
    Join(
		uast.Inner(uast.NewTable("orders", "o"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("o", "id"))),
    ).
    Where(
        uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    )
stmtDeleteReturning := uast.NewDelete(uast.NewTable("users", "u")).
    Where(
        uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    ).
    Returning(
		uast.Field[int64]("u", "id"),
		uast.Field[string]("u", "string"),
	)
stmtDeleteWhere := uast.NewDelete(uast.NewTable("users", "u")).
    Where(
        uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    )
stmtDeleteWith := NewDelete(uast.NewTable("users", "u")).
	With(
		uast.WithN("old_users", uast.NewSelect(uast.NewTable("users", "u")).
			Field(
                uast.Field[int64]("u", "id"),
            ).
			Where(
                Less(uast.Field[int]("u", "number"), Value(2)),
            ),
		),
	)
```
Output MariaDB:
```text
DELETE `u` FROM `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id` WHERE `u`.`string` = ?
DELETE `u` FROM `users` AS `u` WHERE `u`.`string` = ? RETURNING `u`.`id`, `u`.`string`
DELETE `u` FROM `users` AS `u` WHERE `u`.`string` = ?
WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) DELETE `u` FROM `users` AS `u`
```
Output MsSQL:
```text
DELETE [u] FROM [users] AS [u] INNER JOIN [orders] AS [o] ON [u].[id] = [o].[id] WHERE [u].[string] = @p1
DELETE [u] FROM [users] AS [u] OUTPUT [u].[id], [u].[string] WHERE [u].[string] = @p1
DELETE [u] FROM [users] AS [u] WHERE [u].[string] = @p1
WITH [old_users] AS (SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] < @p1) DELETE [u] FROM [users] AS [u]
```
Output MySQL:
```text
DELETE `u` FROM `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id` WHERE `u`.`string` = ?
DELETE `u` FROM `users` AS `u` WHERE `u`.`string` = ?
DELETE `u` FROM `users` AS `u` WHERE `u`.`string` = ?
WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) DELETE `u` FROM `users` AS `u`
```
Output PostgreSQL:
```text
DELETE FROM "users" AS "u" USING "orders" AS "o" WHERE ("u"."id" = "o"."id" AND "u"."string" = $1)
DELETE FROM "users" AS "u" WHERE "u"."string" = $1 RETURNING "u"."id", "u"."string"
DELETE FROM "users" AS "u" WHERE "u"."string" = $1
WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < $1) DELETE FROM "users" AS "u"
```
Output SQLite:
```text
DELETE FROM "users" AS "u" INNER JOIN "orders" AS "o" ON "u"."id" = "o"."id" WHERE "u"."string" = ?
DELETE FROM "users" AS "u" WHERE "u"."string" = ? RETURNING "u"."id", "u"."string"
DELETE FROM "users" AS "u" WHERE "u"."string" = ?
WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < ?) DELETE FROM "users" AS "u"
```

## NewDrop
创建一个新的 DROP 语句实例。接受一个 `Index/Schema/Table/View` 源，并返回一个可使用 `Cascade()` 或 `IfExists()` 配置的语句。
```go
stmtDropCascadeIndex := uast.NewDrop(uast.NewIndex("users_id", uast.NewTable("users", "u"))).
    Cascade()
stmtDropCascadeSchema := uast.NewDrop(uast.NewSchema("test")).
    Cascade()
stmtDropCascadeTable := uast.NewDrop(uast.NewTable("users", "u")).
    Cascade()
stmtDropCascadeView := uast.NewDrop(uast.NewView("users_general", "ug", uast.NewTable("users", "u"))).
    Cascade()
stmtDropIfExistsIndex := uast.NewDrop(uast.NewIndex("users")).
    IfExists()
stmtDropIfExistsSchema := uast.NewDrop(uast.NewSchema("test")).
    IfExists()
stmtDropIfExistsTable := uast.NewDrop(uast.NewTable("users", "u")).
    IfExists()
stmtDropIfExistsView := uast.NewDrop(uast.NewView("users", "u")).
    IfExists()
```
Output MariaDB:
```text
DROP INDEX `users_id` CASCADE
DROP SCHEMA `test` CASCADE
DROP TABLE `users`
DROP VIEW `users_general` CASCADE
DROP INDEX IF EXISTS `users`
DROP SCHEMA IF EXISTS `test`
DROP TABLE IF EXISTS `users`
DROP VIEW IF EXISTS `users`
```
Output MsSQL:
```text
DROP INDEX [users_id]
DROP SCHEMA [test]
DROP TABLE [users]
DROP VIEW [users_general]
DROP INDEX [users]
DROP SCHEMA IF EXISTS [test]
DROP TABLE IF EXISTS [users]
DROP VIEW IF EXISTS [users]
```
Output MySQL:
```text
DROP INDEX `users_id`
DROP SCHEMA `test`
DROP TABLE `users`
DROP VIEW `users_general`
DROP INDEX IF EXISTS `users`
DROP SCHEMA `test`
DROP TABLE IF EXISTS `users`
DROP VIEW IF EXISTS `users`
```
Output PostgreSQL:
```text
DROP INDEX "users_id" CASCADE
DROP SCHEMA "test" CASCADE
DROP TABLE "users" CASCADE
DROP VIEW "users_general" CASCADE
DROP INDEX IF EXISTS "users"
DROP SCHEMA IF EXISTS "test"
DROP TABLE IF EXISTS "users"
DROP VIEW IF EXISTS "users"
```
Output SQLite:
```text
DROP INDEX "users_id"
DROP SCHEMA "test"
DROP TABLE "users"
DROP VIEW "users_general"
DROP INDEX IF EXISTS "users"
DROP SCHEMA "test"
DROP TABLE IF EXISTS "users"
DROP VIEW IF EXISTS "users"
```

## NewInsert
创建一个新的 INSERT 语句实例。接受表源并返回一个可使用 `Returning`、`Source/Values`、`With`进行配置的语句。
```go
stmtInsertReturning := uast.NewInsert(uast.NewTable("users", "u")).
    Values(
		uast.Pair(uast.Field[string]("u", "string"), uast.Value("ivan")),
		uast.Pair(uast.Field[int]("u", "number"), uast.Value(2)),
	).
	Returning(
		uast.Field[int64]("u", "id"),
		uast.Field[string]("u", "string"),
	)
stmtInsertSource := uast.NewInsert(uast.NewTable("users", "u")).
	Source(NewSelect(uast.NewTable("users", "u")).
		Field(
			uast.Field[string]("u", "string"),
			uast.Field[int]("u", "number"),
		).
		Where(
			uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
		),
	)
stmtInsertValues := uast.NewInsert(uast.NewTable("users", "u")).
    Values(
		uast.Pair(uast.Field[string]("u", "string"), uast.Value("ivan")),
		uast.Pair(uast.Field[int]("u", "number"), uast.Value(2)),
	).
    Upsert(
		uast.Pair(uast.Field[string]("u", "string"), uast.Value("updated")),
	)
stmtInsertWith := NewInsert(uast.NewTable("users", "u")).
	Values(
		uast.Pair(uast.Field[string]("u", "string"), uast.Value("ivan")),
		uast.Pair(uast.Field[int]("u", "number"), uast.Value(2)),
	).
	With(
		uast.WithN("old_users", uast.NewSelect(uast.NewTable("users", "u")).
			Field(
                uast.Field[int64]("u", "id"),
		    ).
			Where(
				uast.Less(uast.Field[int]("u", "number"), uast.Value(2)),
			),
		),
	)
```
Output MariaDB:
```text
INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?) RETURNING `u`.`id`, `u`.`string`
INSERT INTO `users` AS `u` (`string`, `number`) SELECT `u`.`string`, `u`.`number` FROM `users` AS `u` WHERE `u`.`string` = ?
INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?)
```
Output MsSQL:
```text
INSERT INTO [users] AS [u] ([string], [number]) OUTPUT [u].[id], [u].[string] VALUES (@p1, @p2)
INSERT INTO [users] AS [u] ([string], [number]) SELECT [u].[string], [u].[number] FROM [users] AS [u] WHERE [u].[string] = @p1
INSERT INTO [users] AS [u] ([string], [number]) VALUES (@p1, @p2)
WITH [old_users] AS (SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] < @p1) INSERT INTO [users] AS [u] ([string], [number]) VALUES (@p2, @p3)
```
Output MySQL:
```text
INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?)
INSERT INTO `users` AS `u` (`string`, `number`) SELECT `u`.`string`, `u`.`number` FROM `users` AS `u` WHERE `u`.`string` = ?
INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) INSERT INTO `users` AS `u` (`string`, `number`) VALUES (?, ?)
```
Output PostgreSQL:
```text
INSERT INTO "users" AS "u" ("string", "number") VALUES ($1, $2) RETURNING "u"."id", "u"."string"
INSERT INTO "users" AS "u" ("string", "number") SELECT "u"."string", "u"."number" FROM "users" AS "u" WHERE "u"."string" = $1
INSERT INTO "users" AS "u" ("string", "number") VALUES ($1, $2) ON CONFLICT DO UPDATE SET "string" = $3
WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < $1) INSERT INTO "users" AS "u" ("string", "number") VALUES ($2, $3)
```
Output SQLite:
```text
INSERT INTO "users" AS "u" ("string", "number") VALUES (?, ?) RETURNING "u"."id", "u"."string"
INSERT INTO "users" AS "u" ("string", "number") SELECT "u"."string", "u"."number" FROM "users" AS "u" WHERE "u"."string" = ?
INSERT INTO "users" AS "u" ("string", "number") VALUES (?, ?) ON CONFLICT DO UPDATE SET "string" = ?
WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < ?) INSERT INTO "users" AS "u" ("string", "number") VALUES (?, ?)
```

## NewSelect
创建一个新的 SELECT 语句实例。接受字段，返回一个可使用 `Distinct`、`Field`、`GroupBy`、`Having`、`Join`、`OrderBy`、`Pagination`、`Unions`、`Where`、`With` 进行配置的语句。
```go
stmtSelectDistinct := uast.NewSelect(uast.NewTable("users", "u")).
    Distinct().
    Field(
        uast.Field[int64]("u", "id"),
    ).
    Where(
		uast.Equal(uast.Field[int]("u", "number"), uast.Value(2)),
	)
stmtSelectField := uast.NewSelect(uast.NewTable("users", "u")).
    Field(
        uast.Field[int64]("u", "id"),
    ).
    Where(
		uast.Equal(uast.Field[int]("u", "number"), uast.Value(2)),
	)
stmtSelectGroupBy := uast.NewSelect(uast.NewTable("users", "u")).
	Field(
		uast.Field[string]("u", "string"),
		uast.Count(uast.Field[int64]("u", "id"), false).As("cnt"),
	).
	GroupBy(
		uast.Field[string]("u", "string"),
	)
stmtSelectHaving := uast.NewSelect(uast.NewTable("users", "u")).
	Field(
		uast.Field[string]("u", "string"),
		uast.Count(uast.Field[int64]("u", "id"), false).As("cnt"),
	).
	GroupBy(
        uast.Field[string]("u", "string"),
    ).
	Having(
		uast.Greater(uast.Count(uast.Field[int64]("u", "id"), false), uast.Value[int64](2)),
	)
stmtSelectJoin := uast.NewSelect(uast.NewTable("users", "u")).
	Field(
		uast.Field[int64]("u", "id"),
		uast.Field[string]("o", "string"),
	).
	Join(
		uast.Inner(uast.NewTable("orders", "o"), Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("o", "id"))),
	)
stmtSelectOrderBy := uast.NewSelect(uast.NewTable("users", "u")).
	Field(
		uast.Field[int64]("u", "id"),
	).
	OrderBy(
		uast.Desc(uast.Field[int]("u", "number")),
		uast.Asc(uast.Field[string]("u", "string")),
	)
stmtSelectPagination := uast.NewSelect(uast.NewTable("users", "u")).
	Field(
		uast.Field[int64]("u", "id"),
	).
	Pagination(10, 20)
stmtSelectUnions := uast.NewSelect(uast.NewTable("users", "u")).
	Field(
		uast.Field[string]("u", "string"),
	).
	Unions(
		uast.UnionAll(uast.NewSelect(uast.NewTable("orders", "o")).
			Field(
				uast.Field[string]("o", "string"),
			),
		),
	)
stmtSelectWhere := uast.NewSelect(uast.NewTable("users", "u")).
	Field(
		uast.Field[int64]("u", "id"),
	).
	Where(
		uast.Equal(uast.Field[int]("u", "number"), uast.Value(2)),
	)
stmtSelectWith := uast.NewSelect(uast.NewCTE("cte_test", "ct")).
	Field(
		uast.Field[int64]("ct", "id"),
	).
	With(
		uast.WithN("cte_test", uast.NewSelect(uast.NewTable("users", "u")).
			Field(
                uast.Field[int64]("u", "id"),
            ).
			Where(
                uast.Greater(uast.Field[int]("u", "number"), uast.Value(2)),
            ),
		),
	)
```
Output MariaDB:
```text
SELECT DISTINCT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` = ?
SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` = ?
SELECT `u`.`string`, COUNT(`u`.`id`) AS `cnt` FROM `users` AS `u` GROUP BY `u`.`string`
SELECT `u`.`string`, COUNT(`u`.`id`) AS `cnt` FROM `users` AS `u` GROUP BY `u`.`string` HAVING COUNT(`u`.`id`) > ?
SELECT `u`.`id`, `o`.`string` FROM `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id`
SELECT `u`.`id` FROM `users` AS `u` ORDER BY `u`.`number` DESC, `u`.`string` ASC
SELECT `u`.`id` FROM `users` AS `u` LIMIT ? OFFSET ?
SELECT `u`.`string` FROM `users` AS `u` UNION ALL SELECT `o`.`string` FROM `orders` AS `o`
SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` = ?
WITH `cte_test` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` > ?) SELECT `ct`.`id` FROM `cte_test` AS `ct`
```
Output MsSQL:
```text
SELECT DISTINCT [u].[id] FROM [users] AS [u] WHERE [u].[number] = @p1
SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] = @p1
SELECT [u].[string], COUNT([u].[id]) AS [cnt] FROM [users] AS [u] GROUP BY [u].[string]
SELECT [u].[string], COUNT([u].[id]) AS [cnt] FROM [users] AS [u] GROUP BY [u].[string] HAVING COUNT([u].[id]) > @p1
SELECT [u].[id], [o].[string] FROM [users] AS [u] INNER JOIN [orders] AS [o] ON [u].[id] = [o].[id]
SELECT [u].[id] FROM [users] AS [u] ORDER BY [u].[number] DESC, [u].[string] ASC
SELECT [u].[id] FROM [users] AS [u] ORDER BY 1 ASC OFFSET @p1 ROWS FETCH NEXT @p2 ROWS ONLY
SELECT [u].[string] FROM [users] AS [u] UNION ALL SELECT [o].[string] FROM [orders] AS [o]
SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] = @p1
WITH [cte_test] AS (SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] > @p1) SELECT [ct].[id] FROM [cte_test] AS [ct]
```
Output MySQL:
```text
SELECT DISTINCT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` = ?
SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` = ?
SELECT `u`.`string`, COUNT(`u`.`id`) AS `cnt` FROM `users` AS `u` GROUP BY `u`.`string`
SELECT `u`.`string`, COUNT(`u`.`id`) AS `cnt` FROM `users` AS `u` GROUP BY `u`.`string` HAVING COUNT(`u`.`id`) > ?
SELECT `u`.`id`, `o`.`string` FROM `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id`
SELECT `u`.`id` FROM `users` AS `u` ORDER BY `u`.`number` DESC, `u`.`string` ASC
SELECT `u`.`id` FROM `users` AS `u` LIMIT ? OFFSET ?
SELECT `u`.`string` FROM `users` AS `u` UNION ALL SELECT `o`.`string` FROM `orders` AS `o`
SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` = ?
WITH `cte_test` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` > ?) SELECT `ct`.`id` FROM `cte_test` AS `ct`
```
Output PostgreSQL:
```text
SELECT DISTINCT "u"."id" FROM "users" AS "u" WHERE "u"."number" = $1
SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" = $1
SELECT "u"."string", COUNT("u"."id") AS "cnt" FROM "users" AS "u" GROUP BY "u"."string"
SELECT "u"."string", COUNT("u"."id") AS "cnt" FROM "users" AS "u" GROUP BY "u"."string" HAVING COUNT("u"."id") > $1
SELECT "u"."id", "o"."string" FROM "users" AS "u" INNER JOIN "orders" AS "o" ON "u"."id" = "o"."id"
SELECT "u"."id" FROM "users" AS "u" ORDER BY "u"."number" DESC, "u"."string" ASC
SELECT "u"."id" FROM "users" AS "u" LIMIT $1 OFFSET $2
SELECT "u"."string" FROM "users" AS "u" UNION ALL SELECT "o"."string" FROM "orders" AS "o"
SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" = $1
WITH "cte_test" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" > $1) SELECT "ct"."id" FROM "cte_test" AS "ct"
```
Output SQLite:
```text
SELECT DISTINCT "u"."id" FROM "users" AS "u" WHERE "u"."number" = ?
SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" = ?
SELECT "u"."string", COUNT("u"."id") AS "cnt" FROM "users" AS "u" GROUP BY "u"."string"
SELECT "u"."string", COUNT("u"."id") AS "cnt" FROM "users" AS "u" GROUP BY "u"."string" HAVING COUNT("u"."id") > ?
SELECT "u"."id", "o"."string" FROM "users" AS "u" INNER JOIN "orders" AS "o" ON "u"."id" = "o"."id"
SELECT "u"."id" FROM "users" AS "u" ORDER BY "u"."number" DESC, "u"."string" ASC
SELECT "u"."id" FROM "users" AS "u" LIMIT ? OFFSET ?
SELECT "u"."string" FROM "users" AS "u" UNION ALL SELECT "o"."string" FROM "orders" AS "o"
SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" = ?
WITH "cte_test" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" > ?) SELECT "ct"."id" FROM "cte_test" AS "ct"
```

## NewTruncate
创建一个新的 TRUNCATE 语句实例。接受表源，并返回一个可以使用 `Cascade()` 或 `RestartIdentity()` 进行配置的语句。
```go
stmtTruncateDefault := uast.NewTruncate(uast.NewTable("users", "u"))
stmtTruncateCascade := uast.NewTruncate(uast.NewTable("users", "u")).
    Cascade()
stmtTruncateRestartIdentity := uast.NewTruncate(uast.NewTable("users", "u")).
    RestartIdentity()
```
Output MariaDB:
```text
TRUNCATE TABLE `users`
TRUNCATE TABLE `users` CASCADE
TRUNCATE TABLE `users` RESTART IDENTITY
```
Output MsSQL:
```text
TRUNCATE TABLE [users]
// Not supported
// Not supported
```
Output MySQL:
```text
TRUNCATE TABLE `users`
// Not supported
// Not supported
```
Output PostgreSQL:
```text
TRUNCATE TABLE "users"
TRUNCATE TABLE "users" CASCADE
TRUNCATE TABLE "users" RESTART IDENTITY
```
Output SQLite:
```text
TRUNCATE TABLE "users"
// Not supported
// Not supported
```

## NewUpdate
创建一个新的 UPDATE 语句实例。接受表源并返回一个可使用 `Join`、`Returning`、`Set`、`Where`、`With` 进行配置的语句。
```go
stmtUpdateJoin := uast.NewUpdate(uast.NewTable("users", "u")).
    Join(
		uast.Inner(uast.NewTable("orders", "o"), Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("o", "id"))),
    ).
    Set(
        uast.Assign(uast.Field[string]("u", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Field[int]("u", "number"), uast.Value(2)),
    ).
stmtUpdateReturning := uast.NewUpdate(uast.NewTable("users", "u")).
    Set(
        uast.Assign(uast.Field[string]("u", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Field[int]("u", "number"), uast.Value(2)),
    ).
    Returning(
        uast.Field[int64]("u", "id"),
        uast.Field[string]("u", "string")
    )
stmtUpdateSet := uast.NewUpdate(uast.NewTable("users", "u")).
    Set(
        uast.Assign(uast.Field[string]("u", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Field[int]("u", "number"), uast.Value(2)),
    )
stmtUpdateWhere := uast.NewUpdate(uast.NewTable("users", "u")).
    Set(
        uast.Assign(uast.Field[string]("u", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Field[int]("u", "number"), uast.Value(2)),
    )
stmtUpdateWith := NewUpdate(uast.NewTable("users", "u")).
	Set(
		uast.Assign(uast.Field[string]("u", "string"), uast.Value("active")),
	).
	With(
		uast.WithN("old_users", uast.NewSelect(uast.NewTable("users", "u")).
			Field(
                uast.Field[int64]("u", "id"),
            ).
			Where(
                uast.Less(uast.Field[int]("u", "number"), uast.Value(2)),
            ),
		),
	)
```
Output MariaDB:
```text
UPDATE `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id` SET `u`.`string` = ? WHERE `o`.`string` = ?
UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ? RETURNING `u`.`id`, `u`.`string`
UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ?
UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ?
WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) UPDATE `users` AS `u` SET `u`.`string` = ?
```
Output MsSQL:
```text
UPDATE [users] AS [u] INNER JOIN [orders] AS [o] ON [u].[id] = [o].[id] SET [u].[string] = @p1 WHERE [o].[string] = @p2
UPDATE [users] AS [u] OUTPUT [u].[id], [u].[string] SET [u].[string] = @p1 WHERE [u].[number] = @p2
UPDATE [users] AS [u] SET [u].[string] = @p1 WHERE [u].[number] = @p2
UPDATE [users] AS [u] SET [u].[string] = @p1 WHERE [u].[number] = @p2
WITH [old_users] AS (SELECT [u].[id] FROM [users] AS [u] WHERE [u].[number] < @p1) UPDATE [users] AS [u] SET [u].[string] = @p2
```
Output MySQL:
```text
UPDATE `users` AS `u` INNER JOIN `orders` AS `o` ON `u`.`id` = `o`.`id` SET `u`.`string` = ? WHERE `o`.`string` = ?
UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ?
UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ?
UPDATE `users` AS `u` SET `u`.`string` = ? WHERE `u`.`number` = ?
WITH `old_users` AS (SELECT `u`.`id` FROM `users` AS `u` WHERE `u`.`number` < ?) UPDATE `users` AS `u` SET `u`.`string` = ?
```
Output PostgreSQL:
```text
UPDATE "users" AS "u" INNER JOIN "orders" AS "o" ON "u"."id" = "o"."id" SET "u"."string" = $1 WHERE "o"."string" = $2
UPDATE "users" AS "u" SET "u"."string" = $1 WHERE "u"."number" = $2 RETURNING "u"."id", "u"."string"
UPDATE "users" AS "u" SET "u"."string" = $1 WHERE "u"."number" = $2
UPDATE "users" AS "u" SET "u"."string" = $1 WHERE "u"."number" = $2
WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < $1) UPDATE "users" AS "u" SET "u"."string" = $2
```
Output SQLite:
```text
UPDATE "users" AS "u" INNER JOIN "orders" AS "o" ON "u"."id" = "o"."id" SET "u"."string" = ? WHERE "o"."string" = ?
UPDATE "users" AS "u" SET "u"."string" = ? WHERE "u"."number" = ? RETURNING "u"."id", "u"."string"
UPDATE "users" AS "u" SET "u"."string" = ? WHERE "u"."number" = ?
UPDATE "users" AS "u" SET "u"."string" = ? WHERE "u"."number" = ?
WITH "old_users" AS (SELECT "u"."id" FROM "users" AS "u" WHERE "u"."number" < ?) UPDATE "users" AS "u" SET "u"."string" = ?
```