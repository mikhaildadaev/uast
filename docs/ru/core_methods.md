---
outline: deep
---

# API / Core / Методы

::: info **Информация**
Эта страница документирует методы, доступные для выражений: `As` для назначения псевдонимов и `Over` для добавления оконных спецификаций. Каждый метод показан с рабочим примером кода и ожидаемым выводом SQL.
:::**

## exprColumn
### As
Назначает псевдоним выражению-колонке.
```go
column := uast.Field[string]("u", "string").As("alias")
```
**Output MariaDB:**
```sql
`u`.`string` AS `alias`
```
**Output MsSQL:**
```sql
[u].[string] AS [alias]
```
**Output MySQL:**
```sql
`u`.`string` AS `alias`
```
**Output PostgreSQL:**
```sql
"u"."string" AS "alias"
```
**Output SQLite:**
```sql
"u"."string" AS "alias"
```

## exprFunction
### As
Назначает псевдоним выражению-функции.
```go
function := uast.Avg(uast.Field[int]("u", "number"), false).As("alias")
```
**Output MariaDB:**
```sql
AVG(`u`.`number`) AS `alias`
```
**Output MsSQL:**
```sql
AVG([u].[number]) AS [alias]
```
**Output MySQL:**
```sql
AVG(`u`.`number`) AS `alias`
```
**Output PostgreSQL:**
```sql
AVG("u"."number") AS "alias"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
AVG(`u`.`number`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
AVG([u].[number]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
AVG(`u`.`number`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
AVG("u"."number") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
AVG("u"."number") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

## exprSubquery
### As
Назначает псевдоним выражению-подзапросу.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Field[int64]("u", "id")).From(uast.NewTable("users", "u"))).As("alias")
```
**Output MariaDB:**
```sql
(SELECT `u`.`id` FROM `users` AS `u`) AS `alias`
```
**Output MsSQL:**
```sql
(SELECT [u].[id] FROM [users] AS [u]) AS [alias]
```
**Output MySQL:**
```sql
(SELECT `u`.`id` FROM `users` AS `u`) AS `alias`
```
**Output PostgreSQL:**
```sql
(SELECT "u"."id" FROM "users" AS "u") AS "alias"
```
**Output SQLite:**
```sql
(SELECT "u"."id" FROM "users" AS "u") AS "alias"
```
