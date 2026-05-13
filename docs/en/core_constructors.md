---
outline: deep
---

# API / Core / Constructors

::: info **Info**
This page covers the four statement constructors: `NewDelete`, `NewInsert`, `NewSelect`, `NewUpdate`. Each constructor creates a new statement instance that can be configured with methods and built into SQL using `Build()`.
:::

## NewDelete
Creates a new DELETE statement instance. Accepts a table source and returns a statement that can be configured with `Join`, `Returning`, `Where`, `With`.
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
Creates a new INSERT statement instance. Accepts columns and returns a statement that can be configured with `Into`, `Returning`, `Source`, `Values`, `With`.
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
Creates a new SELECT statement instance. Accepts fields and returns a statement that can be configured with `Distinct`, `From`, `GroupBy`, `Having`, `Join`, `Limit`, `Offset`, `OrderBy`, `Unions`, `Where`, `With`.
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
Creates a new UPDATE statement instance. Accepts a table source and returns a statement that can be configured with `Join`, `Returning`, `Set`, `Where`, `With`.
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
