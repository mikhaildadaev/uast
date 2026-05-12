---
outline: deep
---

# API / Ядро / Методы

::: info **Информация**
Эта страница документирует методы, доступные для выражений: `As` для назначения псевдонимов и `Over` для добавления оконных спецификаций. Каждый метод показан с рабочим примером кода и ожидаемым выводом SQL.
:::

## exprColumn
Методы, доступные для выражений-колонок.

### As
Назначает псевдоним выражению-колонке.
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

## exprFunction
Методы, доступные для выражений-функций.

### As
Назначает псевдоним выражению-функции.
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

### Over
Добавляет оконную спецификацию к функции, превращая её в оконную функцию.
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

## exprSubquery
Методы, доступные для выражений-подзапросов.

### As
Назначает псевдоним выражению-подзапросу.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Column[int64]("t", "id")).From(uast.Table("test"))).As("alias")
```
Output MySQL:
```text
(SELECT `t`.`id` FROM `test` AS `t`) AS `alias`
```
Output PostgreSQL:
```text
(SELECT "t"."id" FROM "test" AS "t") AS "alias"
```
