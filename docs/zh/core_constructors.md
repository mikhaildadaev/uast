---
outline: deep
---

# API / 核心 / 构造函数

::: info **关于**
本页面涵盖四个语句构造函数：`NewDelete`、`NewInsert`、`NewSelect`、`NewUpdate`。每个构造函数创建一个新的语句实例，可使用方法进行配置，并通过 `Build()` 编译为 SQL。
:::

## NewDelete
创建一个新的 DELETE 语句实例。接受一个表源，返回一个可使用 `Join`、`Returning`、`Where`、`With` 进行配置的语句。
```go
statement := uast.NewDelete(uast.Table("test")).
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
statement := uast.NewInsert(uast.Table("test")).
    Column(
        uast.Column[string]("test", "string"), 
        uast.Column[int]("test", "number"),
    ).
    Values(
        uast.Row(
            uast.Value("ivan"), 
            uast.Value(2),
        ),
    )
```
Output MariaDB:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
```
Output MsSQL:
```text
INSERT INTO [test] AS [t] ([string], [number]) VALUES (@p1, @p2)
```
Output MySQL:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
```
Output PostgreSQL:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES ($1, $2)
```
Output SQLite:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES (?, ?)
```

## NewSelect
创建一个新的 SELECT 语句实例。接受字段，返回一个可使用 `Distinct`、`GroupBy`、`Having`、`Join`、`OrderBy`、`Pagination`、`Unions`、`Where`、`With` 进行配置的语句。
```go
statement := uast.NewSelect(uast.Table("test")).
    Field(
        uast.Column[string]("test", "string"),
    ).
    Where(
        uast.Equal(uast.Column[string]("test", "string"), uast.Value("active")),
    )
```
Output MariaDB:
```text
SELECT `t`.`string` FROM `test` AS `t`
```
Output MsSQL:
```text
SELECT [t].[string] FROM [test] AS [t]
```
Output MySQL:
```text
SELECT `t`.`string` FROM `test` AS `t`
```
Output PostgreSQL:
```text
SELECT "t"."string" FROM "test" AS "t"
```
Output SQLite:
```text
SELECT "t"."string" FROM "test" AS "t"
```

## NewUpdate
创建一个新的 UPDATE 语句实例。接受表源并返回一个可使用 `Join`、`Returning`、`Set`、`Where`、`With` 进行配置的语句。
```go
statement := uast.NewUpdate(uast.Table("test")).
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
