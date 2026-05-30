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
stmtColumn := uast.NewComment("Test comment").
    OnColumn(uast.Column[int64]("test", "id"))
stmtTable := uast.NewComment("Test comment").
    OnTable(uast.Table("test"))
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
stmtDefault := uast.NewDelete(uast.Table("test")).
    Where(
        uast.Equal(uast.Column[string]("test", "string"), uast.Value("active")),
    )
```
Output MariaDB:
```text
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output MsSQL:
```text
DELETE [t] FROM [test] AS [t] WHERE [t].[string] = @p1
```
Output MySQL:
```text
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = $1
```
Output SQLite:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = ?
```

## NewInsert
创建一个新的 INSERT 语句实例。接受表源并返回一个可使用 `Returning`、`Source/Values`、`With`进行配置的语句。
```go
stmtDefault := uast.NewInsert(uast.Table("test")).
    Values(
		uast.Pair(uast.Column[string]("test", "string"), uast.Value("ivan")),
		uast.Pair(uast.Column[int]("test", "number"), uast.Value(2)),
	)
stmtSource := uast.NewInsert(uast.Table("test")).
	Source(NewSelect(uast.Table("test")).
		Field(
			uast.Column[string]("test", "string"),
			uast.Column[int]("test", "number"),
		).
		Where(
			uast.Equal(uast.Column[string]("test", "string"), uast.Value("active")),
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
INSERT INTO "test" AS "d" ("string", "number") SELECT "t"."string", "t"."number" FROM "test" AS "t" WHERE "t"."string" = ?
```

## NewSelect
创建一个新的 SELECT 语句实例。接受字段，返回一个可使用 `Distinct`、`GroupBy`、`Having`、`Join`、`OrderBy`、`Pagination`、`Unions`、`Where`、`With` 进行配置的语句。
```go
stmtDefault := uast.NewSelect(uast.Table("test")).
    Field(
        uast.Column[int64]("test", "id"),
    ).
    Where(
		uast.Equal(uast.Column[int]("test", "number"), uast.Value(2)),
	)
stmtDefault := uast.NewSelect(uast.Table("test")).
    Distinct().
    Field(
        uast.Column[int64]("test", "id"),
    ).
    Where(
		uast.Equal(uast.Column[int]("test", "number"), uast.Value(2)),
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

## NewUpdate
创建一个新的 UPDATE 语句实例。接受表源并返回一个可使用 `Join`、`Returning`、`Set`、`Where`、`With` 进行配置的语句。
```go
stmtDefault := uast.NewUpdate(uast.Table("test")).
    Set(
        Pair(uast.Column[string]("test", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Column[int]("test", "number"), uast.Value(2)),
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
