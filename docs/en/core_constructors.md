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
    Where(
        uast.Equal(uast.Column[string]("test", "string"), uast.Value("active")),
    )
```
Output MySQL:
```text
DELETE FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = $1
```

## NewInsert
Creates a new INSERT statement instance. Accepts columns and returns a statement that can be configured with `Into`, `Returning`, `Source`, `Values`, `With`.
```go
statement := uast.NewInsert(uast.Column[string]("test", "string"), uast.Column[int]("test", "number")).
    Into(uast.Table("test")).
    Values(
        uast.Row(
            uast.Value("ivan"), 
            uast.Value(2),
        ),
    )
```
Output MySQL:
```text
INSERT INTO `test` (`test`.`string`, `test`.`number`) VALUES (?, ?)
```
Output PostgreSQL:
```text
INSERT INTO "test" ("test"."string", "test"."number") VALUES ($1, $2)
```

## NewSelect
Creates a new SELECT statement instance. Accepts fields and returns a statement that can be configured with `Distinct`, `From`, `GroupBy`, `Having`, `Join`, `Limit`, `Offset`, `OrderBy`, `Unions`, `Where`, `With`.
```go
statement := uast.NewSelect(uast.Column[string]("test", "string")).
    From(uast.Table("test")).
    Where(
        uast.Equal(uast.Column[string]("test", "string"), uast.Value("active")),
    )
```
Output MySQL:
```text
SELECT `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
SELECT "t"."string" FROM "test" AS "t" WHERE "t"."string" = $1
```

## NewUpdate
Creates a new UPDATE statement instance. Accepts a table source and returns a statement that can be configured with `Join`, `Returning`, `Set`, `Where`, `With`.
```go
statement := uast.NewUpdate(uast.Table("test")).
    Set(
        Assign(
            uast.Column[string]("test", "string"), 
            uast.Value("active"),
        ),
    ).
    Where(
        uast.Equal(uast.Column[int]("test", "number"), uast.Value(2)),
    )
```
Output MySQL:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2
```
