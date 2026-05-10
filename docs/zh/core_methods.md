---
outline: deep
---

# API / 核心 / 方法

::: info **关于**
This page documents methods available on expressions: `As` for assigning aliases and `Over` for adding window specifications. Each method is shown with a working code example and expected SQL output.
:::

## exprColumn
Methods available on column expressions.

### As
Assigns an alias to a column expression.
```go
column := uast.Column[string]("t", "string").As("alias")
```
Output:
```text
"t"."string" AS "alias"
```

## exprFunction
Methods available on function expressions.

### As
Assigns an alias to a function expression.
```go
function := uast.Avg(uast.Column[int]("t", "number"), false).As("alias")
```
Output:
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
Output:
```text
AVG("t"."number") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

## exprSubquery
Methods available on subquery expressions.

### As
Assigns an alias to a subquery expression.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Column[int64]("t", "id")).From(uast.Table("test"))).As("alias")
```
Output:
```text
(SELECT "t"."id" FROM "test" AS "t") AS "alias"
```
