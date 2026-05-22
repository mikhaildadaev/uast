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
column := uast.Column[string]("t", "string").As("alias")
```
Output MariaDB:
```text
`t`.`string` AS `alias`
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
为函数表达式分配别名。
```go
function := uast.Avg(uast.Column[int]("t", "number"), false).As("alias")
```
Output MariaDB:
```text
AVG(`t`.`number`) AS `alias`
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
为函数添加窗口规范，将其转换为窗口函数。
```go
function := uast.Avg(uast.Column[int]("t", "number"), false).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
AVG(`t`.`number`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
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
为子查询表达式分配别名。
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Column[int64]("t", "id")).From(uast.Table("test"))).As("alias")
```
Output MariaDB:
```text
(SELECT `t`.`id` FROM `test` AS `t`) AS `alias`
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
