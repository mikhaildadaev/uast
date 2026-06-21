---
outline: deep
---

# API / Ядро / Методы

::: info **Информация**
Эта страница документирует методы, доступные для выражений: `As` для назначения псевдонимов и `Over` для добавления оконных спецификаций. Каждый метод показан с рабочим примером кода и ожидаемым выводом SQL.
:::

## exprColumn
### As
Назначает псевдоним выражению-колонке.
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
Назначает псевдоним выражению-функции.
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
Добавляет оконную спецификацию к функции, превращая её в оконную функцию.
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
Назначает псевдоним выражению-подзапросу.
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
