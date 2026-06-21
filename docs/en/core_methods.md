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
column := uast.Field[string]("u", "string").As("alias")
```
Output MariaDB:
```text
`u`.`string` AS `alias`
```
Output MsSQL:
```text
[u].[string] AS [alias]
```
Output MySQL:
```text
`u`.`string` AS `alias`
```
Output PostgreSQL:
```text
"u"."string" AS "alias"
```
Output SQLite:
```text
"u"."string" AS "alias"
```

## exprFunction
### As
Assigns an alias to a function expression.
```go
function := uast.Avg(uast.Field[int]("u", "number"), false).As("alias")
```
Output MariaDB:
```text
AVG(`u`.`number`) AS `alias`
```
Output MsSQL:
```text
AVG([u].[number]) AS [alias]
```
Output MySQL:
```text
AVG(`u`.`number`) AS `alias`
```
Output PostgreSQL:
```text
AVG("u"."number") AS "alias"
```
Output SQLite:
```text
AVG("u"."number") AS "alias"
```

### Over
Adds a window specification to a function, transforming it into a window function.
```go
function := uast.Avg(uast.Field[int]("u", "number"), false).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
AVG(`u`.`number`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
AVG([u].[number]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
AVG(`u`.`number`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
AVG("u"."number") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
AVG("u"."number") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

## exprSubquery
### As
Assigns an alias to a subquery expression.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Field[int64]("u", "id")).From(uast.NewTable("users", "u"))).As("alias")
```
Output MariaDB:
```text
(SELECT `u`.`id` FROM `users` AS `u`) AS `alias`
```
Output MsSQL:
```text
(SELECT [u].[id] FROM [users] AS [u]) AS [alias]
```
Output MySQL:
```text
(SELECT `u`.`id` FROM `users` AS `u`) AS `alias`
```
Output PostgreSQL:
```text
(SELECT "u"."id" FROM "users" AS "u") AS "alias"
```
Output SQLite:
```text
(SELECT "u"."id" FROM "users" AS "u") AS "alias"
```