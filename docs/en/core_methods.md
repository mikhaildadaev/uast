---
outline: deep
---

# API / Core / Methods

::: info **Info**
This page documents methods available on expressions: `As` for assigning aliases and `Over` for adding window specifications. Each method is shown with a working code example and expected SQL output.
:::

## exprColumn
### As
Assigns an alias to a column expression.
```go
column := uast.Column[string]("t", "string").As("alias")
```
Output MySQL:
```text
`t`.`string` AS `alias`
```
Output PostgreSQL:
```text
"t"."string" AS "alias"
```
Output SQLite:
```text
"t"."string" AS "alias"
```

## exprFunction
### As
Assigns an alias to a function expression.
```go
function := uast.Avg(uast.Column[int]("t", "number"), false).As("alias")
```
Output MySQL:
```text
AVG(`t`.`number`) AS `alias`
```
Output PostgreSQL:
```text
AVG("t"."number") AS "alias"
```
Output SQLite:
```text
AVG("t"."number") AS "alias"
```

### Over
Adds a window specification to a function, transforming it into a window function.
```go
function := uast.Avg(uast.Column[int]("t", "number"), false).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MySQL:
```text
AVG(`t`.`number`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
AVG("t"."number") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
AVG("t"."number") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

## exprSubquery
### As
Assigns an alias to a subquery expression.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Column[int64]("t", "id")).From(uast.NewTable("test").As("t"))).As("alias")
```
Output MySQL:
```text
(SELECT `t`.`id` FROM `test` AS `t`) AS `alias`
```
Output PostgreSQL:
```text
(SELECT "t"."id" FROM "test" AS "t") AS "alias"
```
Output SQLite:
```text
(SELECT "t"."id" FROM "test" AS "t") AS "alias"
```