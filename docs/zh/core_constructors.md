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
    Where(uast.Equal(uast.Column[string]("test", "status"), uast.Value("active")))
```
Output MySQL:
```text
DELETE FROM `test` WHERE `test`.`status` = ?
```
Output PostgreSQL:
```text
DELETE FROM "test" WHERE "test"."status" = $1
```

## NewInsert
创建一个新的 INSERT 语句实例。接受列，返回一个可使用 `Into`、`Returning`、`Source`、`Values`、`With`进行配置的语句。
```go
statement := uast.NewInsert(uast.Column[string]("test", "name"), uast.Column[int]("test", "age")).
    Into(uast.Table("test")).
    Values(uast.Value("ivan"), uast.Value(2))
```
Output MySQL:
```text
INSERT INTO `test` (`test`.`name`, `test`.`age`) VALUES (?, ?)
```
Output PostgreSQL:
```text
INSERT INTO "test" ("test"."name", "test"."age") VALUES ($1, $2)
```

## NewSelect
创建一个新的 SELECT 语句实例。接受字段，返回一个可使用 `Distinct`、`From`、`GroupBy`、`Having`、`Join`、`Limit`、`Offset`、`OrderBy`、`Unions`、`Where`、`With` 进行配置的语句。
```go
statement := uast.NewSelect(uast.Column[string]("test", "email")).
    From(uast.Table("test")).
    Where(uast.Equal(uast.Column[string]("test", "status"), uast.Value("active")))
```
Output MySQL:
```text
SELECT `test`.`email` FROM `test` WHERE `test`.`status` = ?
```
Output PostgreSQL:
```text
SELECT "test"."email" FROM "test" WHERE "test"."status" = $1
```

## NewUpdate
创建一个新的 UPDATE 语句实例。接受一个表源，返回一个可使用 `Join`、`Returning`、`Set`、`Where`、`With` 进行配置的语句。
```go
statement := uast.NewUpdate(uast.Table("test")).
    Set(uast.Set(uast.Column[string]("test", "status"), uast.Value("active"))).
    Where(uast.Equal(uast.Column[int]("test", "id"), uast.Value(2)))
```
Output MySQL:
```text
UPDATE `test` SET `test`.`status` = ? WHERE `test`.`id` = ?
```
Output PostgreSQL:
```text
UPDATE "test" SET "test"."status" = $1 WHERE "test"."id" = $2
```
