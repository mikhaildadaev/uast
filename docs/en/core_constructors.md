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
stmtDeleteJoin := NewDelete(Test.Tables.Users).
	Join(
		uast.Inner(Test.Tables.Orders, uast.Equal(Test.Users.ID.Expr(), Test.Orders.ID.Expr())),
	).
	Where(
		uast.Equal(Test.Users.String.Expr(), uast.Value("active")),
	)
stmtDeleteReturning := uast.NewDelete(Test.Tables.Users).
	Where(
		uast.Equal(Test.Users.String.Expr(), uast.Value("active")),
	).
	Returning(
		Test.Users.ID.Expr(),
		Test.Users.String.Expr(),
	)
stmtDeleteWhere := uast.NewDelete(Test.Tables.Users).
	Where(
		uast.Equal(Test.Users.String.Expr(), uast.Value("active")),
	)
stmtDeleteWith := NewDelete(Test.Tables.Users).
	With(
		uast.WithN("old_users", uast.NewSelect(Test.Tables.Users).
			Fields(
				Test.Users.ID.Expr(),
			).
			Where(
				uast.Less(Test.Users.Number.Expr(), uast.Value(2)),
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
stmtDropCascadeIndex := uast.NewDrop(Test.Index.UsersID).
    Cascade()
stmtDropCascadeSchema := uast.NewDrop(Test.Schema).
    Cascade()
stmtDropCascadeTable := uast.NewDrop(Test.Tables.Users).
    Cascade()
stmtDropCascadeView := uast.NewDrop(Test.View.UsersGeneral).
    Cascade()
stmtDropIfExistsIndex := uast.NewDrop(Test.Index.UsersID).
    IfExists()
stmtDropIfExistsSchema := uast.NewDrop(Test.Schema).
    IfExists()
stmtDropIfExistsTable := uast.NewDrop(Test.Tables.Users).
    IfExists()
stmtDropIfExistsView := uast.NewDrop(Test.View.UsersGeneral).
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
Creates a new INSERT statement instance. Accepts a table source and returns a statement that can be configured with `Returning`, `Source/Values`, `With`.
```go
stmtInsertReturning := uast.NewInsert(Test.Tables.Users).
	Values(
		uast.Pair(Test.Users.String.Expr(), uast.Value("ivan")),
		uast.Pair(Test.Users.Number.Expr(), uast.Value(2)),
	).
	Returning(
		Test.Users.ID.Expr(),
		Test.Users.String.Expr(),
	)
stmtInsertSource := uast.NewInsert(Test.Tables.Users).
	Source(uast.NewSelect(Test.Tables.Users).
		Fields(
			Test.Users.String.Expr(),
			Test.Users.Number.Expr(),
		).
		Where(
			uast.Equal(Test.Users.String.Expr(), uast.Value("active")),
		),
	)
stmtInsertValues := uast.NewInsert(Test.Tables.Users).
	Values(
		uast.Pair(Test.Users.String.Expr(), uast.Value("ivan")),
		uast.Pair(Test.Users.Number.Expr(), uast.Value(2)),
	).
	Upsert(
		uast.Pair(Test.Users.String.Expr(), uast.Value("updated")),
	)
stmtInsertWith := uast.NewInsert(Test.Tables.Users).
	Values(
		uast.Pair(Test.Users.String.Expr(), uast.Value("ivan")),
		uast.Pair(Test.Users.Number.Expr(), uast.Value(2)),
	).
	With(
		uast.WithN("old_users", uast.NewSelect(Test.Tables.Users).
			Fields(
				Test.Users.ID.Expr(),
			).
			Where(
				uast.Less(Test.Users.Number.Expr(), uast.Value(2)),
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
stmtSelectDistinct := uast.NewSelect(Test.Tables.Users).
	Distinct().
	Fields(
		Test.Users.ID.Expr(),
	).
	Where(
		uast.Equal(Test.Users.Number.Expr(), uast.Value(2)),
	)
stmtSelectField := uast.NewSelect(Test.Tables.Users).
	Fields(
		Test.Users.ID.Expr(),
	).
	Where(
		uast.Equal(Test.Users.Number.Expr(), uast.Value(2)),
	)
stmtSelectGroupBy := uast.NewSelect(Test.Tables.Users).
	Fields(
		Test.Users.String.Expr(),
		uast.Count(Test.Users.ID.Expr(), false).As("cnt"),
	).
	GroupBy(
		Test.Users.String.Expr(),
	)
stmtSelectHaving := uast.NewSelect(Test.Tables.Users).
	Fields(
		Test.Users.String.Expr(),
		uast.Count(Test.Users.ID.Expr(), false).As("cnt"),
	).
	GroupBy(
		Test.Users.String.Expr(),
	).
	Having(
		uast.Greater(uast.Count(Test.Users.ID.Expr(), false), uast.Value[int64](2)),
	)
stmtSelectJoin := uast.NewSelect(Test.Tables.Users).
	Fields(
		Test.Users.ID.Expr(),
		Test.Orders.String.Expr(),
	).
	Join(
		uast.Inner(Test.Tables.Orders, uast.Equal(Test.Users.ID.Expr(), Test.Orders.ID.Expr())),
	)
stmtSelectOrderBy := uast.NewSelect(Test.Tables.Users).
	Fields(
		Test.Users.ID.Expr(),
	).
	OrderBy(
		uast.Desc(Test.Users.Number.Expr()),
		uast.Asc(Test.Users.String.Expr()),
	)
stmtSelectPagination := uast.NewSelect(Test.Tables.Users).
	Fields(
		Test.Users.ID.Expr(),
	).
	Pagination(10, 20)
stmtSelectUnions := uast.NewSelect(Test.Tables.Users).
	Fields(
		Test.Users.String.Expr(),
	).
	Unions(
		uast.UnionAll(uast.NewSelect(Test.Tables.Orders).
			Fields(
				Test.Orders.String.Expr(),
			),
		),
	)
stmtSelectWhere := uast.NewSelect(Test.Tables.Users).
	Fields(
		Test.Users.ID.Expr(),
	).
	Where(
		uast.Equal(Test.Users.Number.Expr(), uast.Value(2)),
	)
stmtSelectWith := uast.NewSelect(uast.NewCTE("cte_user", "ct")).
	Fields(
		uast.Field[int64]("ct", "id"),
	).
	With(
		uast.WithN("cte_user", uast.NewSelect(Test.Tables.Users).
			Fields(
				Test.Users.ID.Expr(),
			).
			Where(
				uast.Greater(Test.Users.Number.Expr(), uast.Value(2)),
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
stmtTruncateDefault := uast.NewTruncate(Test.Tables.Users)
stmtTruncateCascade := uast.NewTruncate(Test.Tables.Users).
    Cascade()
stmtTruncateRestartIdentity := uast.NewTruncate(Test.Tables.Users).
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
stmtUpdateJoin := uast.NewUpdate(Test.Tables.Users).
	Set(
		uast.Assign(Test.Users.String.Expr(), uast.Value("active")),
	).
	Join(
		uast.Inner(Test.Tables.Orders, uast.Equal(Test.Users.ID.Expr(), Test.Orders.ID.Expr())),
	).
	Where(
		uast.Equal(Test.Orders.String.Expr(), uast.Value("active")),
	)
stmtUpdateReturning := uast.NewUpdate(Test.Tables.Users).
	Set(
		uast.Assign(Test.Users.String.Expr(), uast.Value("active")),
	).
	Where(
		uast.Equal(Test.Users.Number.Expr(), uast.Value(2)),
	).
	Returning(
		Test.Users.ID.Expr(),
		Test.Users.String.Expr(),
	)
stmtUpdateSet := uast.NewUpdate(Test.Tables.Users).
	Set(
		uast.Assign(Test.Users.String.Expr(), uast.Value("active")),
	).
	Where(
		uast.Equal(Test.Users.Number.Expr(), uast.Value(2)),
	)
stmtUpdateWhere := uast.NewUpdate(Test.Tables.Users).
	Set(
		uast.Assign(Test.Users.String.Expr(), uast.Value("active")),
	).
	Where(
		uast.Equal(Test.Users.Number.Expr(), uast.Value(2)),
	)
stmtUpdateWith := uast.NewUpdate(Test.Tables.Users).
	Set(
		uast.Assign(Test.Users.String.Expr(), uast.Value("updated")),
	).
	With(
		uast.WithN("old_users", uast.NewSelect(Test.Tables.Users).
			Fields(
				Test.Users.ID.Expr(),
			).
			Where(
				uast.Less(Test.Users.Number.Expr(), uast.Value(2)),
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