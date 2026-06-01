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
Creates a new DELETE statement instance. Accepts a table source and returns a statement that can be configured with `Join`, `Returning`, `Where`, `With`.
```go
stmtDeleteDefault := uast.NewDelete(uast.NewTable("test", "t")).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    )
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
```
Output MariaDB:
```text
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?
DELETE `t` FROM `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id` WHERE `t`.`string` = ?
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ? RETURNING `t`.`id`, `t`.`string`
```
Output MsSQL:
```text
DELETE [t] FROM [test] AS [t] WHERE [t].[string] = @p1
DELETE [t] FROM [test] AS [t] INNER JOIN [data] AS [d] ON [t].[id] = [d].[id] WHERE [t].[string] = @p1
DELETE [t] FROM [test] AS [t] OUTPUT [t].[id], [t].[string] WHERE [t].[string] = @p1
```
Output MySQL:
```text
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?
DELETE `t` FROM `test` AS `t` INNER JOIN `data` AS `d` ON `t`.`id` = `d`.`id` WHERE `t`.`string` = ?
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = $1
DELETE FROM "test" AS "t" USING "data" AS "d" WHERE ("t"."id" = "d"."id" AND "t"."string" = $1)
DELETE FROM "test" AS "t" WHERE "t"."string" = $1 RETURNING "t"."id", "t"."string"
```
Output SQLite:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = ?
DELETE FROM "test" AS "t" INNER JOIN "data" AS "d" ON "t"."id" = "d"."id" WHERE "t"."string" = ?
DELETE FROM "test" AS "t" WHERE "t"."string" = ? RETURNING "t"."id", "t"."string"
```

## NewDrop
Creates a new DROP statement instance. Accepts a `Index/Schema/Table/View` source and returns a statement that can be configured with `Cascade()` or `IfExists()`.
```go
stmtDropCascade := uast.NewDrop(uast.NewTable("test", "t")).
    Cascade()
stmtDropIndexIfExists := uast.NewDrop(uast.NewIndex("test")).
    IfExists()
stmtDropSchemaIfExists := uast.NewDrop(uast.NewSchema("test")).
    IfExists()
stmtDropTableIfExists := uast.NewDrop(uast.NewTable("test", "t")).
    IfExists()
stmtDropViewIfExists := uast.NewDrop(uast.NewView("test", "t")).
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
stmtInsertDefault := uast.NewInsert(uast.NewTable("test", "t")).
    Values(
		uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
		uast.Pair(uast.Column[int]("t", "number"), uast.Value(2)),
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
```
Output MariaDB:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
INSERT INTO `test` AS `t` (`string`, `number`) SELECT `t`.`string`, `t`.`number` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output MsSQL:
```text
INSERT INTO [test] AS [t] ([string], [number]) VALUES (@p1, @p2)
INSERT INTO [test] AS [t] ([string], [number]) SELECT [t].[string], [t].[number] FROM [test] AS [t] WHERE [t].[string] = @p1
```
Output MySQL:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
INSERT INTO `test` AS `t` (`string`, `number`) SELECT `t`.`string`, `t`.`number` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES ($1, $2)
INSERT INTO "test" AS "t" ("string", "number") SELECT "t"."string", "t"."number" FROM "test" AS "t" WHERE "t"."string" = $1
```
Output SQLite:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES (?, ?)
INSERT INTO "test" AS "t" ("string", "number") SELECT "t"."string", "t"."number" FROM "test" AS "t" WHERE "t"."string" = ?
```

## NewSelect
Creates a new SELECT statement instance. Accepts a table source and returns a statement that can be configured with `Distinct`, `GroupBy`, `Having`, `Join`, `OrderBy`, `Pagination`, `Unions`, `Where`, `With`.
```go
stmtSelectDefault := uast.NewSelect(uast.NewTable("test", "t")).
    Field(
        uast.Column[int64]("t", "id"),
    ).
    Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	)
stmtSelectDistinct := uast.NewSelect(uast.NewTable("test", "t")).
    Distinct().
    Field(
        uast.Column[int64]("t", "id"),
    ).
    Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	)
```
Output MariaDB:
```text
SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
SELECT DISTINCT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
```
Output MsSQL:
```text
SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] = @p1
SELECT DISTINCT [t].[id] FROM [test] AS [t] WHERE [t].[number] = @p1
```
Output MySQL:
```text
SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
SELECT DISTINCT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
```
Output PostgreSQL:
```text
SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = $1
SELECT DISTINCT "t"."id" FROM "test" AS "t" WHERE "t"."number" = $1
```
Output SQLite:
```text
SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = ?
SELECT DISTINCT "t"."id" FROM "test" AS "t" WHERE "t"."number" = ?
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
stmtUpdateDefault := uast.NewUpdate(uast.NewTable("test", "t")).
    Set(
        Pair(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
    )
```
Output MariaDB:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
```
Output MsSQL:
```text
UPDATE [test] AS [t] SET [t].[string] = @p1 WHERE [t].[number] = @p2
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