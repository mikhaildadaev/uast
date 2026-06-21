---
outline: deep
---

# API / 核心 / 方法

::: info **关于**
本页面记录了可用于表达式的方法： `As` 用于分配别名，`Over` 用于添加窗口规范。每个方法都配有可运行的代码示例和预期 SQL 输出。
:::

## exprColumn
### As
为列表达式分配别名。
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
为函数表达式分配别名。
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
为函数添加窗口规范，将其转换为窗口函数。
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
为子查询表达式分配别名。
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
