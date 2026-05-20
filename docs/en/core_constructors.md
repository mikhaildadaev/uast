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
statement := uast.NewDelete(uast.NewTable("test").As("t")).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    )
```
Output MariaDB:
```text
DELETE FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output MySQL:
```text
DELETE FROM `test` AS `t` WHERE `t`.`string` = ?
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
Creates a new INSERT statement instance. Accepts a table source and returns a statement that can be configured with `Returning`, `Source`, `Values`, `With`.
```go
statement := uast.NewInsert(uast.NewTable("test").As("t")).
    Values(
        uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
        uast.Pair(uast.Column[int]("t", "number"), uast.Value(2)),
    )
```
Output MariaDB:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
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
Creates a new SELECT statement instance. Accepts a table source and returns a statement that can be configured with `Distinct`, `GroupBy`, `Having`, `Join`, `Limit`, `Offset`, `OrderBy`, `Unions`, `Where`, `With`.
```go
statement := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    )
```
Output MariaDB:
```text
SELECT `t`.`string` FROM `test` AS `t`
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
Creates a new UPDATE statement instance. Accepts a table source and returns a statement that can be configured with `Join`, `Returning`, `Set`, `Where`, `With`.
```go
statement := uast.NewUpdate(uast.NewTable("test").As("t")).
    Set(
        uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
    )
```
Output MariaDB:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
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