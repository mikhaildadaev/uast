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
    OnColumn(uast.Field[int64]("u", "id"))
stmtCommentTable := uast.NewComment("Test comment").
    OnTable(uast.NewTable("users", "u"))
```
Output MariaDB:
```text
COMMENT ON COLUMN `u`.`id` IS 'Test comment'
COMMENT ON TABLE `users` AS `u` IS 'Test comment'
```
Output MsSQL:
```text
// Not supported
// Not supported
```
Output MySQL:
```text
COMMENT ON COLUMN `u`.`id` IS 'Test comment'
COMMENT ON TABLE `users` AS `u` IS 'Test comment'
```
Output PostgreSQL:
```text
COMMENT ON COLUMN "u"."id" IS 'Test comment'
COMMENT ON TABLE "users" AS "u" IS 'Test comment'
```
Output SQLite:
```text
COMMENT ON COLUMN "u"."id" IS 'Test comment'
COMMENT ON TABLE "users" AS "u" IS 'Test comment'
```

## NewDelete
Creates a new DELETE statement instance. Accepts a table source and returns a statement that can be configured with `Join`, `Returning`, `Where`, `With`.
```go
stmtDeleteJoin := uast.NewDelete(uast.NewTable("users", "u")).
    Join(
		Inner(uast.NewTable("data", "d"), Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("d", "id"))),
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
Creates a new DROP statement instance. Accepts a `Index/Schema/Table/View` source and returns a statement that can be configured with `Cascade()` or `IfExists()`.
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
stmtDropIfExistsSchema := uast.NewDrop(uast.NewSchema("users")).
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
DROP SCHEMA "users" CASCADE
DROP TABLE "users" CASCADE
DROP VIEW "users_general" CASCADE
DROP INDEX IF EXISTS "users"
DROP SCHEMA IF EXISTS "users"
DROP TABLE IF EXISTS "users"
DROP VIEW IF EXISTS "users"
```
Output SQLite:
```text
DROP INDEX "users_id"
DROP SCHEMA "users"
DROP TABLE "users"
DROP VIEW "users_general"
DROP INDEX IF EXISTS "users"
DROP SCHEMA "users"
DROP TABLE IF EXISTS "users"
DROP VIEW IF EXISTS "users"
```

## NewInsert
Creates a new INSERT statement instance. Accepts a table source and returns a statement that can be configured with `Returning`, `Source/Values`, `With`.
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
Creates a new SELECT statement instance. Accepts a table source and returns a statement that can be configured with `Distinct`, `Field`, `GroupBy`, `Having`, `Join`, `OrderBy`, `Pagination`, `Unions`, `Where`, `With`.
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
		uast.Field[string]("d", "string"),
	).
	Join(
		uast.Inner(uast.NewTable("data", "d"), Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("d", "id"))),
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
		uast.UnionAll(uast.NewSelect(uast.NewTable("data", "d")).
			Field(
				uast.Field[string]("d", "string"),
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
Creates a new TRUNCATE statement instance. Accepts a table source and returns a statement that can be configured with `Cascade()` or `RestartIdentity()`.
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
Creates a new UPDATE statement instance. Accepts a table source and returns a statement that can be configured with `Join`, `Returning`, `Set`, `Where`, `With`.
```go
stmtUpdateJoin := uast.NewUpdate(uast.NewTable("users", "u")).
    Join(
		uast.Inner(uast.NewTable("data", "d"), Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("d", "id"))),
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