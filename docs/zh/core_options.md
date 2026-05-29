---
outline: deep
---

# API / 核心 / 选项

::: info **关于**
本页面涵盖了所有配置选项：`clauseGroupBy`、`clauseHaving`、`clauseJoin`、`clauseOrderBy`、`clausePagination`、`clauseReturning`、`clauseSet`、`clauseUnions`、`clauseValues`、`clauseWhere`、`clauseWith`、`exprArray`、`exprBinary`、`exprColumn`、`exprComparison`、`exprConstant`、`exprFunction`、`exprLiteral`、`exprLogical`、`exprSubquery`、`exprValue`。每个选项都配有可运行的代码示例和预期输出。
:::

## clauseGroupBy
添加 GROUP BY 子句，按指定列或表达式对行进行分组。
```go
groupBy := GroupBy(
	uast.Column[string]("t", "string"),
)
```
Output MariaDB:
```text
GROUP BY `t`.`string`
```
Output MsSQL:
```text
GROUP BY [t].[string]
```
Output MySQL:
```text
GROUP BY `t`.`string`
```
Output PostgreSQL:
```text
GROUP BY "t"."string"
```
Output SQLite:
```text
GROUP BY "t"."string"
```

## clauseHaving
添加 HAVING 子句以过滤分组。与 GROUP BY 一起使用以过滤聚合结果。
```go
having := Having(
	uast.Greater(uast.Count(uast.Column[int64]("t", "id"), false), uast.Value[int64](2)),
)
```
Output MariaDB:
```text
HAVING COUNT(`t`.`id`) > ?
```
Output MsSQL:
```text
HAVING COUNT([t].[id]) > @p1
```
Output MySQL:
```text
HAVING COUNT(`t`.`id`) > ?
```
Output PostgreSQL:
```text
HAVING COUNT("t"."id") > $1
```
Output SQLite:
```text
HAVING COUNT("t"."id") > ?
```

## clauseJoin
### Cross
向语句添加 CROSS JOIN。返回两个表的笛卡尔积。
```go
join := uast.Cross(uast.NewTable("test").As("t"))
```
Output MariaDB:
```text
CROSS JOIN `test` AS `t`
```
Output MsSQL:
```text
CROSS JOIN [test] AS [t]
```
Output MySQL:
```text
CROSS JOIN `test` AS `t`
```
Output PostgreSQL:
```text
CROSS JOIN "test" AS "t"
```
Output SQLite:
```text
CROSS JOIN "test" AS "t"
```

### Full
向语句添加 FULL JOIN。返回两个表中的所有行，不匹配处为 NULL。
```go
join := uast.Full(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
FULL JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
FULL JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
FULL JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
FULL JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
FULL JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```

### FullOuter
向语句添加 FULL OUTER JOIN。返回两个表中的所有行，不匹配处为 NULL。
```go
join := uast.FullOuter(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
FULL OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
FULL OUTER JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
FULL OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
FULL OUTER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
FULL OUTER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```

### Inner
向语句添加 INNER JOIN。返回两个表中具有匹配值的行。
```go
join := uast.Inner(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
INNER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
INNER JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
INNER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
INNER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
INNER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```

### Left
向语句添加 LEFT JOIN。返回左表中的所有行以及右表中的匹配行。
```go
join := uast.Left(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
LEFT JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
LEFT JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
LEFT JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
LEFT JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
LEFT JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```

### LeftOuter
向语句添加 LEFT OUTER JOIN。返回左表中的所有行以及右表中的匹配行。
```go
join := uast.LeftOuter(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
LEFT OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
LEFT OUTER JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
LEFT OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
LEFT OUTER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
LEFT OUTER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```

### Right
向语句添加 RIGHT JOIN。返回右表中的所有行以及左表中的匹配行。SQLite 不支持。
```go
join := uast.Right(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
RIGHT JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
RIGHT JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
RIGHT JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
RIGHT JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
// Not supported
```

### RightOuter
向语句添加 RIGHT OUTER JOIN。返回右表中的所有行以及左表中的匹配行。SQLite 不支持。
```go
join := uast.RightOuter(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
RIGHT OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
RIGHT OUTER JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
RIGHT OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
RIGHT OUTER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
// Not supported
```

## clauseOrderBy
### Asc
指定升序排序（最小的在前，A 到 Z）。用于对查询中的行或窗口函数内的行进行排序。
```go
orderBy := uast.Asc(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` ASC
```
Output MsSQL:
```text
[t].[string] ASC
```
Output MySQL:
```text
`t`.`string` ASC
```
Output PostgreSQL:
```text
"t"."string" ASC
```
Output SQLite:
```text
"t"."string" ASC
```

### Desc
指定降序排序（最大的在前，Z 到 A）。用于对查询中的行或窗口函数内的行进行排序。
```go
orderBy := uast.Desc(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` DESC
```
Output MsSQL:
```text
[t].[string] DESC
```
Output MySQL:
```text
`t`.`string` DESC
```
Output PostgreSQL:
```text
"t"."string" DESC
```
Output SQLite:
```text
"t"."string" DESC
```

## clausePagination
使用 `Pagination(limit, offset)` 为 SELECT `limit` `offset` 指定在返回结果之前要跳过的行数。渲染顺序和语法会自动适应每种方言。
```go
pagination := Pagination(10,0)
```
Output MariaDB:
```text
LIMIT ? OFFSET ?
```
Output MsSQL:
```text
OFFSET @p1 ROWS FETCH NEXT @p2 ROWS ONLY
```
Output MySQL:
```text
LIMIT ? OFFSET ?
```
Output PostgreSQL:
```text
LIMIT $1 OFFSET $2
```
Output SQLite:
```text
LIMIT ? OFFSET ?
```

## clauseReturning
添加 RETURNING 子句以返回修改后的行。MariaDB、PostgreSQL 和 SQLite 支持。MySQL 原生不支持此子句。
```go
returning = Returning(
	uast.Column[int64]("t", "id")
    uast.Column[string]("t", "string")
)
```
Output MariaDB:
```text
RETURNING `t`.`id`, `t`.`string`
```
Output MsSQL:
```text
OUTPUT [t].[id], [t].[string]
```
Output MySQL:
```text
// Not support
```
Output PostgreSQL:
```text
RETURNING "t"."id", "t"."string"
```
Output SQLite:
```text
RETURNING "t"."id", "t"."string"
```

## clauseSet
### Assign
使用 `Assign` 指定列及其新值，将列与值关联。支持多个对以更新多个列。
```go
set := Set(
	uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
)
```
Output MariaDB:
```text
UPDATE `test` AS `t` SET `t`.`string` = ?
```
Output MsSQL:
```text
UPDATE [test] AS [t] SET [t].[string] = @p1
```
Output MySQL:
```text
UPDATE `test` AS `t` SET `t`.`string` = ?
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" SET "t"."string" = $1
```
Output SQLite:
```text
UPDATE "test" AS "t" SET "t"."string" = ?
```

## clauseUnions
### Union
合并多个 SELECT 语句的结果。UNION 返回不重复的行。
```go
unions := uast.Union(uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ),
)
```
Output MariaDB:
```text
UNION SELECT `t`.`string` FROM `test` AS `t` 
```
Output MsSQL:
```text
UNION SELECT [t].[string] FROM [test] AS [t]
```
Output MySQL:
```text
UNION SELECT `t`.`string` FROM `test` AS `t`
```
Output PostgreSQL:
```text
UNION SELECT "t"."string" FROM "test" AS "t"
```
Output SQLite:
```text
UNION SELECT "t"."string" FROM "test" AS "t"
```

### UnionAll
合并多个 SELECT 语句的结果。UNION ALL 返回所有行，包括重复行。
```go
unions := uast.UnionAll(uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ),
)
```
Output MariaDB:
```text
UNION ALL SELECT `t`.`string` FROM `test` AS `t`
```
Output MsSQL:
```text
UNION ALL SELECT [t].[string] FROM [test] AS [t]
```
Output MySQL:
```text
UNION ALL SELECT `t`.`string` FROM `test` AS `t`
```
Output PostgreSQL:
```text
UNION ALL SELECT "t"."string" FROM "test" AS "t"
```
Output SQLite:
```text
UNION ALL SELECT "t"."string" FROM "test" AS "t"
```

### UnionExcept
合并多个 SELECT 语句的结果。EXCEPT 返回第一个查询中存在但第二个查询中不存在的不同行。
```go
unions := uast.UnionExcept(uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ),
)
```
Output MariaDB:
```text
EXCEPT SELECT `t`.`string` FROM `test` AS `t`
```
Output MsSQL:
```text
EXCEPT SELECT [t].[string] FROM [test] AS [t]
```
Output MySQL:
```text
EXCEPT SELECT `t`.`string` FROM `test` AS `t`
```
Output PostgreSQL:
```text
EXCEPT SELECT "t"."string" FROM "test" AS "t"
```
Output SQLite:
```text
EXCEPT SELECT "t"."string" FROM "test" AS "t"
```

### UnionIntersect
合并多个 SELECT 语句的结果。INTERSECT 返回两个查询共有的不同行。
```go
unions := uast.UnionIntersect(uast.NewSelect(uast.NewTable("test").As("t")).
	Field(
		uast.Column[string]("t", "string"),
	),
)
```
Output MariaDB:
```text
INTERSECT SELECT `t`.`string` FROM `test` AS `t`
```
Output MsSQL:
```text
INTERSECT SELECT [t].[string] FROM [test] AS [t]
```
Output MySQL:
```text
INTERSECT SELECT `t`.`string` FROM `test` AS `t`
```
Output PostgreSQL:
```text
INTERSECT SELECT "t"."string" FROM "test" AS "t"
```
Output SQLite:
```text
INTERSECT SELECT "t"."string" FROM "test" AS "t"
```

## clauseValues
### Pair
使用 `Pair` 指定插入的值，将列与值关联。列会自动从对中推断。
```go
values := Values(
    uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
	uast.Pair(uast.Column[int]("t", "number"), uast.Value(2)),
)
```
Output MariaDB:
```text
VALUES (?, ?)
```
Output MsSQL:
```text
VALUES (@p1, @p2)
```
Output MySQL:
```text
VALUES (?, ?)
```
Output PostgreSQL:
```text
VALUES ($1, $2)
```
Output SQLite:
```text
VALUES (?, ?)
```

### Upsert
使用 `Upsert` 向 INSERT ... VALUES 添加 upsert 子句。将列与值关联。
```go
values := Values(
    uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
	uast.Pair(uast.Column[int]("t", "number"), uast.Value(2)),
).
Upsert(
    uast.Pair(uast.Column[string]("t", "string"), uast.Value("updated")),
)
```
Output MariaDB:
```text
VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
```
Output MsSQL:
```text
// Not supported
```
Output MySQL:
```text
VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
```
Output PostgreSQL:
```text
VALUES ($1, $2) ON CONFLICT DO UPDATE SET "string" = $3
```
Output SQLite:
```text
VALUES (?, ?) ON CONFLICT DO UPDATE SET "string" = ?
```

## clauseWhere
添加 WHERE 子句以在分组或聚合之前过滤行。接受比较表达式、逻辑运算符和子查询。
```go
where = Where(
	uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
)
```
Output MariaDB:
```text
WHERE `t`.`string` = ?
```
Output MsSQL:
```text
WHERE [t].[string] = @p1
```
Output MySQL:
```text
WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
WHERE "t"."string" = $1
```
Output SQLite:
```text
WHERE "t"."string" = ?
```

## clauseWith
### Norecursive
使用 `WithN` 向语句添加非递归公共表表达式（CTE）。列通过可变字符串参数别名化。
```go
with := WithN("cte_norecursive", NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[int64]("t", "id"),
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    ),
    "id", "string",
)
```
Output MariaDB:
```text
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ?)
```
Output MsSQL:
```text
WITH [cte_norecursive] ([id], [string]) AS (SELECT [t].[id], [t].[string] FROM [test] AS [t] WHERE [t].[string] = @p1)
```
Output MySQL:
```text
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ?)
```
Output PostgreSQL:
```text
WITH "cte_norecursive" ("id", "string") AS (SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = $1)
```
Output SQLite:
```text
WITH "cte_norecursive" ("id", "string") AS (SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = ?)
```

### Recursive
使用 `WithR` 向语句添加递归公共表表达式（CTE）。需要带有 `UnionAll` 的 `Unions` 子句来定义递归步骤。
```go
with := WithR("cte_recursive", NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[int64]("t", "id"),
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Unions(
        uast.UnionAll(uast.NewSelect(uast.NewTable("test").As("t")).
            Field(
                uast.Column[int64]("t", "id"),
                uast.Column[string]("t", "string"),
            ).
            Join(
                uast.Inner(uast.NewCTE("cte_recursive", "rec"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("rec", "id"))),
            ),
        ),
    ),
    "id", "string",
)
```
Output MariaDB:
```text
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ? UNION ALL SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` INNER JOIN `cte_recursive` AS `rec` ON `t`.`id` = `rec`.`id`)
```
Output MsSQL:
```text
WITH RECURSIVE [cte_recursive] ([id], [string]) AS (SELECT [t].[id], [t].[string] FROM [test] AS [t] WHERE [t].[string] = @p1 UNION ALL SELECT [t].[id], [t].[string] FROM [test] AS [t] INNER JOIN [cte_recursive] AS [rec] ON [t].[id] = [rec].[id])
```
Output MySQL:
```text
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ? UNION ALL SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` INNER JOIN `cte_recursive` AS `rec` ON `t`.`id` = `rec`.`id`)
```
Output PostgreSQL:
```text
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = $1 UNION ALL SELECT "t"."id", "t"."string" FROM "test" AS "t" INNER JOIN "cte_recursive" AS "rec" ON "t"."id" = "rec"."id")
```
Output SQLite:
```text
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = ? UNION ALL SELECT "t"."id", "t"."string" FROM "test" AS "t" INNER JOIN "cte_recursive" AS "rec" ON "t"."id" = "rec"."id")
```

## exprArray
### Array
构造用于 SQL 查询的数组表达式。
```go
array := uast.Array(0, 1, 2)
```
Output MariaDB:
```text
ARRAY[?, ?, ?]
```
Output MsSQL:
```text
ARRAY[@p1, @p2, @p3]
```
Output MySQL:
```text
ARRAY[?, ?, ?]
```
Output PostgreSQL:
```text
ARRAY[$1, $2, $3]
```
Output SQLite:
```text
ARRAY[?, ?, ?]
```

## exprBinary
### BitwiseAnd
对两个表达式执行按位与运算。
```go
binary := uast.BitwiseAnd(uast.Column[int]("t", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`t`.`number` & ?
```
Output MsSQL:
```text
[t].[number] & @p1
```
Output MySQL:
```text
`t`.`number` & ?
```
Output PostgreSQL:
```text
"t"."number" & $1
```
Output SQLite:
```text
"t"."number" & ?
```

### BitwiseOr
对两个表达式执行按位或运算。
```go
binary := uast.BitwiseOr(uast.Column[int]("t", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`t`.`number` | ?
```
Output MsSQL:
```text
[t].[number] | @p1
```
Output MySQL:
```text
`t`.`number` | ?
```
Output PostgreSQL:
```text
"t"."number" | $1
```
Output SQLite:
```text
"t"."number" | ?
```

### BitwiseXor
对两个表达式执行按位异或运算。
```go
binary := uast.BitwiseXor(uast.Column[int]("t", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`t`.`number` ^ ?
```
Output MsSQL:
```text
[t].[number] ^ @p1
```
Output MySQL:
```text
`t`.`number` ^ ?
```
Output PostgreSQL:
```text
"t"."number" ^ $1
```
Output SQLite:
```text
"t"."number" ^ ?
```

### Divide
左表达式除以右表达式。
```go
binary := uast.Divide(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` / ?
```
Output MsSQL:
```text
[t].[number] / @p1
```
Output MySQL:
```text
`t`.`number` / ?
```
Output PostgreSQL:
```text
"t"."number" / $1
```
Output SQLite:
```text
"t"."number" / ?
```

### Minus
左表达式减去右表达式。
```go
binary := uast.Minus(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` - ?
```
Output MsSQL:
```text
[t].[number] - @p1
```
Output MySQL:
```text
`t`.`number` - ?
```
Output PostgreSQL:
```text
"t"."number" - $1
```
Output SQLite:
```text
"t"."number" - ?
```

### Modulo
返回左表达式除以右表达式的余数。
```go
binary := uast.Modulo(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` % ?
```
Output MsSQL:
```text
[t].[number] % @p1
```
Output MySQL:
```text
`t`.`number` % ?
```
Output PostgreSQL:
```text
"t"."number" % $1
```
Output SQLite:
```text
"t"."number" % ?
```

### Multiply
左表达式乘以右表达式。
```go
binary := uast.Multiply(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` * ?
```
Output MsSQL:
```text
[t].[number] * @p1
```
Output MySQL:
```text
`t`.`number` * ?
```
Output PostgreSQL:
```text
"t"."number" * $1
```
Output SQLite:
```text
"t"."number" * ?
```

### Plus
左表达式加上右表达式。
```go
binary := uast.Plus(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` + ?
```
Output MsSQL:
```text
[t].[number] + @p1
```
Output MySQL:
```text
`t`.`number` + ?
```
Output PostgreSQL:
```text
"t"."number" + $1
```
Output SQLite:
```text
"t"."number" + ?
```

### ShiftLeft
对左表达式执行按位左移，移动位数由右表达式指定。
```go
binary := uast.ShiftLeft(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` << ?
```
Output MsSQL:
```text
[t].[number] << @p1
```
Output MySQL:
```text
`t`.`number` << ?
```
Output PostgreSQL:
```text
"t"."number" << $1
```
Output SQLite:
```text
"t"."number" << ?
```

### ShiftRight
对左表达式执行按位右移，移动位数由右表达式指定。
```go
binary := uast.ShiftRight(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` >> ?
```
Output MsSQL:
```text
[t].[number] >> @p1
```
Output MySQL:
```text
`t`.`number` >> ?
```
Output PostgreSQL:
```text
"t"."number" >> $1
```
Output SQLite:
```text
"t"."number" >> ?
```

## exprColumn
### Column
创建对表列的引用，可选择用表别名限定。这是在表达式中引用数据库列的主要方式。
```go
column := uast.Column[string]("t", "string")
```
Output MariaDB:
```text
`t`.`string`
```
Output MsSQL:
```text
[t].[string]
```
Output MySQL:
```text
`t`.`string`
```
Output PostgreSQL:
```text
"t"."string"
```
Output SQLite:
```text
"t"."string"
```

## exprComparison
### Between
检查左表达式是否落在 `valueStart` 和 `valueEnd` 定义的范围内（含）。
```go
comparison := uast.Between(uast.Column[int]("t", "number"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` BETWEEN ? AND ?
```
Output MsSQL:
```text
[t].[number] BETWEEN @p1 AND @p2
```
Output MySQL:
```text
`t`.`number` BETWEEN ? AND ?
```
Output PostgreSQL:
```text
"t"."number" BETWEEN $1 AND $2
```
Output SQLite:
```text
"t"."number" BETWEEN ? AND ?
```

### Equal
比较两个表达式是否相等（`=`）。
```go
comparison := uast.Equal(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` = ?
```
Output MsSQL:
```text
[t].[number] = @p1
```
Output MySQL:
```text
`t`.`number` = ?
```
Output PostgreSQL:
```text
"t"."number" = $1
```
Output SQLite:
```text
"t"."number" = ?
```

### Exists
检查子查询是否返回任何行。如果至少存在一行则返回 `true`。
```go
comparison := uast.Exists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("test").As("t"))))
```
Output MariaDB:
```text
EXISTS (SELECT 1 FROM `test` AS `t`)
```
Output MsSQL:
```text
EXISTS (SELECT 1 FROM [test] AS [t])
```
Output MySQL:
```text
EXISTS (SELECT 1 FROM `test` AS `t`)
```
Output PostgreSQL:
```text
EXISTS (SELECT 1 FROM "test" AS "t")
```
Output SQLite:
```text
EXISTS (SELECT 1 FROM "test" AS "t")
```

### Greater
比较左表达式是否大于右表达式（`>`）。
```go
comparison := uast.Greater(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` > ?
```
Output MsSQL:
```text
[t].[number] > @p1
```
Output MySQL:
```text
`t`.`number` > ?
```
Output PostgreSQL:
```text
"t"."number" > $1
```
Output SQLite:
```text
"t"."number" > ?
```

### GreaterEqual
比较左表达式是否大于或等于右表达式（`>=`）。
```go
comparison := uast.GreaterEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` >= ?
```
Output MsSQL:
```text
[t].[number] >= @p1
```
Output MySQL:
```text
`t`.`number` >= ?
```
Output PostgreSQL:
```text
"t"."number" >= $1
```
Output SQLite:
```text
"t"."number" >= ?
```

### ILike
执行不区分大小写的模式匹配比较。右表达式应包含带有 `%`（任意序列）和 `_`（单个字符）通配符的模式。
```go
comparison := uast.ILike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
LOWER(`t`.`string`) LIKE LOWER(?)
```
Output MsSQL:
```text
LOWER([t].[string]) LIKE LOWER(@p1)
```
Output MySQL:
```text
LOWER(`t`.`string`) LIKE LOWER(?)
```
Output PostgreSQL:
```text
"t"."string" ILIKE $1
```
Output SQLite:
```text
LOWER("t"."string") LIKE LOWER(?)
```

### In
检查左表达式是否匹配右表达式中包含的任何值（通常是子查询或数组）。
```go
comparison := uast.In(uast.Column[string]("t", "string"), uast.Array("active", "pending"))
```
Output MariaDB:
```text
`t`.`string` IN (?, ?)
```
Output MsSQL:
```text
[t].[string] IN (@p1, @p2)
```
Output MySQL:
```text
`t`.`string` IN (?, ?)
```
Output PostgreSQL:
```text
"t"."string" IN ($1, $2)
```
Output SQLite:
```text
"t"."string" IN (?, ?)
```

### IsNotNull
检查表达式是否不为 `NULL`。
```go
comparison := uast.IsNotNull(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` IS NOT NULL
```
Output MsSQL:
```text
[t].[string] IS NOT NULL
```
Output MySQL:
```text
`t`.`string` IS NOT NULL
```
Output PostgreSQL:
```text
"t"."string" IS NOT NULL
```
Output SQLite:
```text
"t"."string" IS NOT NULL
```

### IsNull
检查表达式是否为 `NULL`。
```go
comparison := uast.IsNull(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` IS NULL
```
Output MsSQL:
```text
[t].[string] IS NULL
```
Output MySQL:
```text
`t`.`string` IS NULL
```
Output PostgreSQL:
```text
"t"."string" IS NULL
```
Output SQLite:
```text
"t"."string" IS NULL
```

### Less
比较左表达式是否小于右表达式（`<`）。
```go
comparison := uast.Less(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` < ?
```
Output MsSQL:
```text
[t].[number] < @p1
```
Output MySQL:
```text
`t`.`number` < ?
```
Output PostgreSQL:
```text
"t"."number" < $1
```
Output SQLite:
```text
"t"."number" < ?
```

### LessEqual
比较左表达式是否小于或等于右表达式（`<=`）。
```go
comparison := uast.LessEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` <= ?
```
Output MsSQL:
```text
[t].[number] <= @p1
```
Output MySQL:
```text
`t`.`number` <= ?
```
Output PostgreSQL:
```text
"t"."number" <= $1
```
Output SQLite:
```text
"t"."number" <= ?
```

### Like
执行区分大小写的模式匹配比较。右表达式应包含带有 `%` 和 `_` 通配符的模式。
```go
comparison := uast.Like(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
`t`.`string` LIKE ?
```
Output MsSQL:
```text
[t].[number] LIKE @p1
```
Output MySQL:
```text
`t`.`string` LIKE ?
```
Output PostgreSQL:
```text
"t"."string" LIKE $1
```
Output SQLite:
```text
"t"."string" LIKE ?
```

### NotBetween
检查左表达式是否落在 `valueStart` 和 `valueEnd` 定义的范围之外。
```go
comparison := uast.NotBetween(uast.Column[int]("t", "number"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` NOT BETWEEN ? AND ?
```
Output MsSQL:
```text
[t].[number] NOT BETWEEN @p1 AND @p2
```
Output MySQL:
```text
`t`.`number` NOT BETWEEN ? AND ?
```
Output PostgreSQL:
```text
"t"."number" NOT BETWEEN $1 AND $2
```
Output SQLite:
```text
"t"."number" NOT BETWEEN ? AND ?
```

### NotEqual
比较两个表达式是否不相等（`!=` 或 `<>`）。
```go
comparison := uast.NotEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` != ?
```
Output MsSQL:
```text
[t].[number] != @p1
```
Output MySQL:
```text
`t`.`number` != ?
```
Output PostgreSQL:
```text
"t"."number" != $1
```
Output SQLite:
```text
"t"."number" != ?
```

### NotExists
检查子查询是否不返回任何行。如果子查询结果为空则返回 `true`。
```go
comparison := uast.NotExists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("test").As("t"))))
```
Output MariaDB:
```text
NOT EXISTS (SELECT 1 FROM `test` AS `t`)
```
Output MsSQL:
```text
NOT EXISTS (SELECT 1 FROM [test] AS [t])
```
Output MySQL:
```text
NOT EXISTS (SELECT 1 FROM `test` AS `t`)
```
Output PostgreSQL:
```text
NOT EXISTS (SELECT 1 FROM "test" AS "t")
```
Output SQLite:
```text
NOT EXISTS (SELECT 1 FROM "test" AS "t")
```

### NotILike
执行否定的不区分大小写的模式匹配比较。
```go
comparison := uast.NotILike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
LOWER(`t`.`string`) NOT LIKE LOWER(?)
```
Output MsSQL:
```text
LOWER([t].[string]) NOT LIKE LOWER(@p1)
```
Output MySQL:
```text
LOWER(`t`.`string`) NOT LIKE LOWER(?)
```
Output PostgreSQL:
```text
"t"."string" NOT ILIKE $1
```
Output SQLite:
```text
LOWER("t"."string") NOT LIKE LOWER(?)
```

### NotIn
检查左表达式是否不匹配右表达式中包含的任何值。
```go
comparison := uast.NotIn(uast.Column[string]("t", "string"), uast.Array("active", "pending"))
```
Output MariaDB:
```text
`t`.`string` NOT IN (?, ?)
```
Output MsSQL:
```text
[t].[string] NOT IN (@p1, @p2)
```
Output MySQL:
```text
`t`.`string` NOT IN (?, ?)
```
Output PostgreSQL:
```text
"t"."string" NOT IN ($1, $2)
```
Output SQLite:
```text
"t"."string" NOT IN (?, ?)
```

### NotLike
执行否定的区分大小写的模式匹配比较。
```go
comparison := uast.NotLike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
`t`.`string` NOT LIKE ?
```
Output MsSQL:
```text
[t].[string] NOT LIKE @p1
```
Output MySQL:
```text
`t`.`string` NOT LIKE ?
```
Output PostgreSQL:
```text
"t"."string" NOT LIKE $1
```
Output SQLite:
```text
"t"."string" NOT LIKE ?
```

## exprConstant
### ConstBoolFalse
返回常量布尔 `FALSE` 表达式。
```go
constant := uast.ConstBoolFalse()
```
Output:
```text
FALSE
```

### ConstBoolTrue
返回常量布尔 `TRUE` 表达式。
```go
constant := uast.ConstBoolTrue()
```
Output:
```text
TRUE
```

### ConstFloat32One
返回值为 `1.0` 的 `float32` 常量。
```go
constant := uast.ConstFloat32One()
```
Output:
```text
1.0
```

### ConstFloat64One
返回值为 `1.000000` 的 `float64` 常量。
```go
constant := uast.ConstFloat64One()
```
Output:
```text
1.000000
```

### ConstIntOne
返回值为 `1` 的 `int` 常量。
```go
constant := uast.ConstIntOne()
```
Output:
```text
1
```

### ConstInt8One
返回值为 `1` 的 `int8` 常量。
```go
constant := uast.ConstInt8One()
```
Output:
```text
1
```

### ConstInt16One
返回值为 `1` 的 `int16` 常量。
```go
constant := uast.ConstInt16One()
```
Output:
```text
1
```

### ConstInt32One
返回值为 `1` 的 `int32` 常量。
```go
constant := uast.ConstInt32One()
```
Output:
```text
1
```

### ConstInt64One
返回值为 `1` 的 `int64` 常量。
```go
constant := uast.ConstInt64One()
```
Output:
```text
1
```

### ConstStringDefault
返回值为 `DEFAULT` 的 `string` 常量。
```go
constant := uast.ConstStringDefault()
```
Output:
```text
DEFAULT
```

### ConstStringNull
返回值为 `NULL` 的 `string` 常量。
```go
constant := uast.ConstStringNull()
```
Output:
```text
NULL
```

### ConstUintOne
返回值为 `1` 的 `uint` 常量。
```go
constant := uast.ConstUintOne()
```
Output:
```text
1
```

### ConstUint8One
返回值为 `1` 的 `uint8` 常量。
```go
constant := uast.ConstUint8One()
```
Output:
```text
1
```

### ConstUint16One
返回值为 `1` 的 `uint16` 常量。
```go
constant := uast.ConstUint16One()
```
Output:
```text
1
```

### ConstUint32One
返回值为 `1` 的 `uint32` 常量。
```go
constant := uast.ConstUint32One()
```
Output:
```text
1
```

### ConstUint64One
返回值为 `1` 的 `uint64` 常量。
```go
constant := uast.ConstUint64One()
```
Output:
```text
1
```

## exprFunction
### Aggregate
#### Avg
返回表达式中所有非 NULL 值的平均值（算术平均）。如果 `distinct` 为 `true`，则仅对不重复的值计算平均值。
```go
function := uast.Avg(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Avg(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
AVG(`t`.`number`)
AVG(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
AVG([t].[number])
AVG(DISTINCT [t].[number])
```
Output MySQL:
```text
AVG(`t`.`number`)
AVG(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
AVG("t"."number")
AVG(DISTINCT "t"."number")
```
Output SQLite:
```text
AVG("t"."number")
AVG(DISTINCT "t"."number")
```

#### BitAnd
返回表达式中所有位的按位与。仅对整数类型有意义。
```go
function := uast.BitAnd(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitAnd(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
BIT_AND(`t`.`number`)
BIT_AND(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
BIT_AND([t].[number])
BIT_AND(DISTINCT [t].[number])
```
Output MySQL:
```text
BIT_AND(`t`.`number`)
BIT_AND(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
BIT_AND("t"."number")
BIT_AND(DISTINCT "t"."number")
```
Output SQLite:
```text
BIT_AND("t"."number")
BIT_AND(DISTINCT "t"."number")
```

#### BitOr
返回表达式中所有位的按位或。仅对整数类型有意义。
```go
function := uast.BitOr(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitOr(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
BIT_OR(`t`.`number`)
BIT_OR(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
BIT_OR([t].[number])
BIT_OR(DISTINCT [t].[number])
```
Output MySQL:
```text
BIT_OR(`t`.`number`)
BIT_OR(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
BIT_OR("t"."number")
BIT_OR(DISTINCT "t"."number")
```
Output SQLite:
```text
BIT_OR("t"."number")
BIT_OR(DISTINCT "t"."number")
```

#### BitXor
返回表达式中所有位的按位异或。仅对整数类型有意义。
```go
function := uast.BitXor(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitXor(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
BIT_XOR(`t`.`number`)
BIT_XOR(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
BIT_XOR([t].[number])
BIT_XOR(DISTINCT [t].[number])
```
Output MySQL:
```text
BIT_XOR(`t`.`number`)
BIT_XOR(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
BIT_XOR("t"."number")
BIT_XOR(DISTINCT "t"."number")
```
Output SQLite:
```text
BIT_XOR("t"."number")
BIT_XOR(DISTINCT "t"."number")
```

#### Count
返回匹配查询的行数，如果提供了表达式，则返回非 NULL 值的数量。当 `distinct` 为 `true` 时，仅计数不重复的值。
```go
function := uast.Count(uast.Column[string]("t", "string"), false)
functionWithDistinct := uast.Count(uast.Column[string]("t", "string"), true)
```
Output MariaDB:
```text
COUNT(`t`.`string`)
COUNT(DISTINCT `t`.`string`)
```
Output MsSQL:
```text
COUNT([t].[string])
COUNT(DISTINCT [t].[string])
```
Output MySQL:
```text
COUNT(`t`.`string`)
COUNT(DISTINCT `t`.`string`)
```
Output PostgreSQL:
```text
COUNT("t"."string")
COUNT(DISTINCT "t"."string")
```
Output SQLite:
```text
COUNT("t"."string")
COUNT(DISTINCT "t"."string")
```

#### GroupConcat
将组中的值连接成一个字符串，用默认分隔符（通常是逗号）分隔。`distinct` 标志在连接前删除重复项。
```go
function := uast.GroupConcat(uast.Column[string]("t", "string"), false)
functionWithDistinct := uast.GroupConcat(uast.Column[string]("t", "string"), true)
```
Output MariaDB:
```text
GROUP_CONCAT(`t`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `t`.`string` SEPARATOR ',')
```
Output MsSQL:
```text
GROUP_CONCAT([t].[string], ',')
GROUP_CONCAT(DISTINCT [t].[string], ',')
```
Output MySQL:
```text
GROUP_CONCAT(`t`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `t`.`string` SEPARATOR ',')
```
Output PostgreSQL:
```text
STRING_AGG("t"."string", ',')
STRING_AGG(DISTINCT "t"."string", ',')
```
Output SQLite:
```text
GROUP_CONCAT("t"."string" SEPARATOR ',')
GROUP_CONCAT(DISTINCT "t"."string" SEPARATOR ',')
```

#### Max
返回组中所有行的表达式的最大值。
```go
function := uast.Max(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Max(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
MAX(`t`.`number`)
MAX(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
MAX([t].[number])
MAX(DISTINCT [t].[number])
```
Output MySQL:
```text
MAX(`t`.`number`)
MAX(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
MAX("t"."number")
MAX(DISTINCT "t"."number")
```
Output SQLite:
```text
MAX("t"."number")
MAX(DISTINCT "t"."number")
```

#### Min
返回组中所有行的表达式的最小值。
```go
function := uast.Min(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Min(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
MIN(`t`.`number`)
MIN(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
MIN([t].[number])
MIN(DISTINCT [t].[number])
```
Output MySQL:
```text
MIN(`t`.`number`)
MIN(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
MIN("t"."number")
MIN(DISTINCT "t"."number")
```
Output SQLite:
```text
MIN("t"."number")
MIN(DISTINCT "t"."number")
```

#### StdDev
返回表达式的总体标准差。
```go
function := uast.StdDev(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.StdDev(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
STDDEV(`t`.`number`)
STDDEV(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
STDEV([t].[number])
STDEV(DISTINCT [t].[number])
```
Output MySQL:
```text
STDDEV(`t`.`number`)
STDDEV(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
STDDEV_SAMP("t"."number")
STDDEV_SAMP(DISTINCT "t"."number")
```
Output SQLite:
```text
STDEV("t"."number")
STDEV(DISTINCT "t"."number")
```

#### Sum
返回表达式中所有值的总和。如果 `distinct` 为 `true`，则仅对不重复的值求和。
```go
function := uast.Sum(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Sum(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
SUM(`t`.`number`)
SUM(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
SUM([t].[number])
SUM(DISTINCT [t].[number])
```
Output MySQL:
```text
SUM(`t`.`number`)
SUM(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
SUM("t"."number")
SUM(DISTINCT "t"."number")
```
Output SQLite:
```text
SUM("t"."number")
SUM(DISTINCT "t"."number")
```

#### Variance
返回表达式的总体方差。
```go
function := uast.Variance(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Variance(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
VARIANCE(`t`.`number`)
VARIANCE(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
VAR([t].[number])
VAR(DISTINCT [t].[number])
```
Output MySQL:
```text
VARIANCE(`t`.`number`)
VARIANCE(DISTINCT "t"."number")
```
Output PostgreSQL:
```text
VAR_SAMP("t"."number")
VAR_SAMP(DISTINCT "t"."number")
```
Output SQLite:
```text
VARIANCE("t"."number")
VARIANCE(DISTINCT "t"."number")
```

### Analytical
#### FirstValue
返回窗口框架第一行的表达式值。需要带有窗口规范的 `OVER` 子句。
```go
function := uast.FirstValue(uast.Column[string]("t", "string")).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
FIRST_VALUE(`t`.`string`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
FIRST_VALUE([t].[string]) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
FIRST_VALUE(`t`.`string`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
FIRST_VALUE("t"."string") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
FIRST_VALUE("t"."string") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### Lag
返回分区内当前行之前偏移 `offset` 行的表达式值。
```go
function := uast.Lag(uast.Column[int]("t", "number"), 2).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Asc(uast.Column[time.Time]("t", "date"))),
)
```
Output MariaDB:
```text
LAG(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)
```
Output MsSQL:
```text
LAG([t].[number], 2) OVER (PARTITION BY [t].[id] ORDER BY [t].[date] ASC)
```
Output MySQL:
```text
LAG(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)
```
Output PostgreSQL:
```text
LAG("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)
```
Output SQLite:
```text
LAG("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)
```

#### LastValue
返回窗口框架最后一行的表达式值。
```go
function := uast.LastValue(uast.Column[string]("t", "string")).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Asc(uast.Column[int]("t", "number"))),
    uast.RowsBetween("CURRENT ROW", "UNBOUNDED FOLLOWING"),
)
```
Output MariaDB:
```text
LAST_VALUE(`t`.`string`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output MsSQL:
```text
LAST_VALUE([t].[string]) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output MySQL:
```text
LAST_VALUE(`t`.`string`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output PostgreSQL:
```text
LAST_VALUE("t"."string") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output SQLite:
```text
LAST_VALUE("t"."string") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```

#### Lead
返回分区内当前行之后偏移 `offset` 行的表达式值。
```go
function := uast.Lead(uast.Column[int]("t", "number"), 2).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Asc(uast.Column[time.Time]("t", "date"))),
)
```
Output MariaDB:
```text
LEAD(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)
```
Output MsSQL:
```text
LEAD([t].[number], 2) OVER (PARTITION BY [t].[id] ORDER BY [t].[date] ASC)
```
Output MySQL:
```text
LEAD(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)
```
Output PostgreSQL:
```text
LEAD("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)
```
Output SQLite:
```text
LEAD("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)
```

#### NthValue
返回窗口框架第 `n` 行的表达式值。
```go
function := uast.NthValue(uast.Column[string]("t", "string"), 2).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
    uast.RowsBetween("UNBOUNDED PRECEDING", "CURRENT ROW"),
)
```
Output MariaDB:
```text
NTH_VALUE(`t`.`string`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output MsSQL:
```text
NTH_VALUE([t].[string], 2) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output MySQL:
```text
NTH_VALUE(`t`.`string`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output PostgreSQL:
```text
NTH_VALUE("t"."string", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output SQLite:
```text
NTH_VALUE("t"."string", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```

### Condition
#### Case
计算 `WHEN`-`THEN` 对列表，并为第一个为真的 `WHEN` 返回 `THEN` 表达式。如果没有条件为真，则返回 `ELSE` 表达式（如果提供），否则返回 `NULL`。
```go
pairs := uast.CaseIf(
    uast.CasePair(
        uast.Less(uast.Column[int]("t", "number"), uast.Value(2)),
        uast.Value("old"),
    ),
)
elseExpr := uast.CaseElse(uast.Value("new"))
function := uast.Case(pairs, elseExpr)
```
Output MariaDB:
```text
CASE WHEN `t`.`number` < ? THEN ? ELSE ? END
```
Output MsSQL:
```text
CASE WHEN [t].[number] < @p1 THEN @p2 ELSE @p3 END
```
Output MySQL:
```text
CASE WHEN `t`.`number` < ? THEN ? ELSE ? END
```
Output PostgreSQL:
```text
CASE WHEN "t"."number" < $1 THEN $2 ELSE $3 END
```
Output SQLite:
```text
CASE WHEN "t"."number" < ? THEN ? ELSE ? END
```

#### Coalesce
返回提供列表中的第一个非 NULL 表达式。用于提供回退值。
```go
function := uast.Coalesce(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
COALESCE(`t`.`createat`, `t`.`updateat`)
```
Output MsSQL:
```text
COALESCE([t].[createat], [t].[updateat])
```
Output MySQL:
```text
COALESCE(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
COALESCE("t"."createat", "t"."updateat")
```
Output SQLite:
```text
COALESCE("t"."createat", "t"."updateat")
```

#### Greatest
返回提供表达式列表中的最大值。
```go
function := uast.Greatest(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
GREATEST(`t`.`createat`, `t`.`updateat`)
```
Output MsSQL:
```text
GREATEST([t].[createat], [t].[updateat])
```
Output MySQL:
```text
GREATEST(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
GREATEST("t"."createat", "t"."updateat")
```
Output SQLite:
```text
GREATEST("t"."createat", "t"."updateat")
```

#### Least
返回提供表达式列表中的最小值。
```go
function := uast.Least(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
LEAST(`t`.`createat`, `t`.`updateat`)
```
Output MsSQL:
```text
LEAST([t].[createat], [t].[updateat])
```
Output MySQL:
```text
LEAST(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
LEAST("t"."createat", "t"."updateat")
```
Output SQLite:
```text
LEAST("t"."createat", "t"."updateat")
```

#### NullIf
如果两个表达式相等，则返回 `NULL`；否则返回第一个表达式。
```go
function := uast.NullIf(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
NULLIF(`t`.`createat`, `t`.`updateat`)
```
Output MsSQL:
```text
NULLIF([t].[createat], [t].[updateat])
```
Output MySQL:
```text
NULLIF(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
NULLIF("t"."createat", "t"."updateat")
```
Output SQLite:
```text
NULLIF("t"."createat", "t"."updateat")
```

### Convert
#### Cast
将表达式转换为指定的数据类型。
```go
function := uast.Cast(uast.Column[int]("t", "number"), uast.TypeString)
```
Output MariaDB:
```text
CAST(`t`.`number` AS CHAR)
```
Output MsSQL:
```text
CAST([t].[number] AS NVARCHAR)
```
Output MySQL:
```text
CAST(`t`.`number` AS CHAR)
```
Output PostgreSQL:
```text
CAST("t"."number" AS VARCHAR)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

#### CharLength
返回字符串表达式中的字符数。
```go
function := uast.CharLength(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
CHAR_LENGTH(`t`.`string`)
```
Output MsSQL:
```text
CHAR_LENGTH([t].[string])
```
Output MySQL:
```text
CHAR_LENGTH(`t`.`string`)
```
Output PostgreSQL:
```text
CHAR_LENGTH("t"."string")
```
Output SQLite:
```text
CHAR_LENGTH("t"."string")
```

#### DateFormat
根据指定的格式掩码格式化日期时间表达式。
```go
function := uast.DateFormat(uast.Column[time.Time]("t", "createat"), uast.Value("%Y-%m-%d"))
```
Output MariaDB:
```text
DATE_FORMAT(`t`.`createat`, '%Y-%m-%d')
```
Output MsSQL:
```text
FORMAT([t].[createat], '%Y-%m-%d')
```
Output MySQL:
```text
DATE_FORMAT(`t`.`createat`, '%Y-%m-%d')
```
Output PostgreSQL:
```text
TO_CHAR("t"."createat", '%Y-%m-%d')
```
Output SQLite:
```text
STRFTIME("t"."createat", '%Y-%m-%d')
```

#### Degrees
将角度从弧度转换为度数。
```go
function := uast.Degrees(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
DEGREES(`t`.`number`)
```
Output MsSQL:
```text
DEGREES([t].[number])
```
Output MySQL:
```text
DEGREES(`t`.`number`)
```
Output PostgreSQL:
```text
DEGREES("t"."number")
```
Output SQLite:
```text
DEGREES("t"."number")
```

#### Length
返回字符串表达式的字节长度。
```go
function := uast.Length(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
LENGTH(`t`.`string`)
```
Output MsSQL:
```text
LEN([t].[string])
```
Output MySQL:
```text
LENGTH(`t`.`string`)
```
Output PostgreSQL:
```text
LENGTH("t"."string")
```
Output SQLite:
```text
LENGTH("t"."string")
```

#### Position
返回子字符串在字符串中首次出现的起始位置。
```go
function := uast.Position(uast.Column[string]("t", "string"), uast.Value("old"))
```
Output MariaDB:
```text
POSITION(? IN `t`.`string`)
```
Output MsSQL:
```text
CHARINDEX(@p1, [t].[string])
```
Output MySQL:
```text
POSITION(? IN `t`.`string`)
```
Output PostgreSQL:
```text
POSITION($1 IN "t"."string")
```
Output SQLite:
```text
POSITION(? IN "t"."string")
```

#### Radians
将角度从度数转换为弧度。
```go
function := uast.Radians(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
RADIANS(`t`.`number`)
```
Output MsSQL:
```text
RADIANS([t].[number])
```
Output MySQL:
```text
RADIANS(`t`.`number`)
```
Output PostgreSQL:
```text
RADIANS("t"."number")
```
Output SQLite:
```text
RADIANS("t"."number")
```

### Date and time
#### CurDate
返回当前日期（不含时间）。
```go
function := uast.CurDate()
```
Output MariaDB:
```text
CURDATE()
```
Output MsSQL:
```text
CAST(GETDATE() AS DATE)
```
Output MySQL:
```text
CURDATE()
```
Output PostgreSQL:
```text
CURRENT_DATE
```
Output SQLite:
```text
DATE('now')
```

#### CurTime
返回当前时间（不含日期）。
```go
function := uast.CurTime()
```
Output MariaDB:
```text
CURTIME()
```
Output MsSQL:
```text
CAST(GETDATE() AS TIME)
```
Output MySQL:
```text
CURTIME()
```
Output PostgreSQL:
```text
CURRENT_TIME
```
Output SQLite:
```text
TIME('now')
```

#### DateAdd
将时间/日期间隔添加到日期时间表达式，并返回结果日期时间。
```go
function := uast.DateAdd(uast.Column[time.Time]("t", "createat"), uast.Value("2 DAY"))
```
Output MariaDB:
```text
DATE_ADD(`t`.`createat`, INTERVAL 2 DAY)
```
Output MsSQL:
```text
DATEADD(DAY, 2, [t].[createat])
```
Output MySQL:
```text
DATE_ADD(`t`.`createat`, INTERVAL 2 DAY)
```
Output PostgreSQL:
```text
("t"."createat" + INTERVAL '2 DAY')
```
Output SQLite:
```text
DATETIME("t"."createat", '+2 DAY')
```

#### DateDiff
返回两个日期时间表达式之间的天数差（`datetimeEnd` - `datetimeStart`）。
```go
function := uast.DateDiff(uast.Column[time.Time]("t", "updateat"), uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
DATEDIFF(`t`.`updateat`, `t`.`createat`)
```
Output MsSQL:
```text
DATEDIFF([t].[updateat], [t].[createat])
```
Output MySQL:
```text
DATEDIFF(`t`.`updateat`, `t`.`createat`)
```
Output PostgreSQL:
```text
DATE_PART('day', "t"."updateat" - "t"."createat")
```
Output SQLite:
```text
DATEDIFF("t"."updateat", "t"."createat")
```

#### DateSub
从日期时间表达式中减去时间/日期间隔，并返回结果日期时间。
```go
function := uast.DateSub(uast.Column[time.Time]("t", "createat"), uast.Value("2 DAY"))
```
Output MariaDB:
```text
DATE_SUB(`t`.`createat`, INTERVAL 2 DAY)
```
Output MsSQL:
```text
DATEADD(DAY, -2, [t].[createat])
```
Output MySQL:
```text
DATE_SUB(`t`.`createat`, INTERVAL 2 DAY)
```
Output PostgreSQL:
```text
("t"."createat" - INTERVAL '2 DAY')
```
Output SQLite:
```text
DATETIME("t"."createat", '-2 DAY')
```

#### Day
从日期时间表达式中提取月份中的日期（1–31）。
```go
function := uast.Day(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
DAY(`t`.`createat`)
```
Output MsSQL:
```text
DAY([t].[createat])
```
Output MySQL:
```text
DAY(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(DAY FROM "t"."createat")
```
Output SQLite:
```text
DAY("t"."createat")
```

#### DayName
返回给定日期时间表达式的星期名称（例如 'Monday', 'Tuesday'）。
```go
function := uast.DayName(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
DAYNAME(`t`.`createat`)
```
Output MsSQL:
```text
DATENAME(WEEKDAY, [t].[createat])
```
Output MySQL:
```text
DAYNAME(`t`.`createat`)
```
Output PostgreSQL:
```text
TO_CHAR("t"."createat", 'Day')
```
Output SQLite:
```text
STRFTIME('%w', "t"."createat")
```

#### Hour
从日期时间表达式中提取小时（0–23）。
```go
function := uast.Hour(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
HOUR(`t`.`createat`)
```
Output MsSQL:
```text
DATEPART(HOUR, [t].[createat])
```
Output MySQL:
```text
HOUR(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(HOUR FROM "t"."createat")
```
Output SQLite:
```text
HOUR("t"."createat")
```

#### Minute
从日期时间表达式中提取分钟（0–59）。
```go
function := uast.Minute(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
MINUTE(`t`.`createat`)
```
Output MsSQL:
```text
DATEPART(MINUTE, [t].[createat])
```
Output MySQL:
```text
MINUTE(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(MINUTE FROM "t"."createat")
```
Output SQLite:
```text
MINUTE("t"."createat")
```

#### Month
从日期时间表达式中提取月份（1–12）。
```go
function := uast.Month(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
MONTH(`t`.`createat`)
```
Output MsSQL:
```text
MONTH([t].[createat])
```
Output MySQL:
```text
MONTH(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(MONTH FROM "t"."createat")
```
Output SQLite:
```text
MONTH("t"."createat")
```

#### MonthName
返回给定日期时间表达式的月份名称（例如 'January', 'February'）。
```go
function := uast.MonthName(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
MONTHNAME(`t`.`createat`)
```
Output MsSQL:
```text
DATENAME(MONTH, [t].[createat])
```
Output MySQL:
```text
MONTHNAME(`t`.`createat`)
```
Output PostgreSQL:
```text
TO_CHAR("t"."createat", 'Month')
```
Output SQLite:
```text
STRFTIME('%m', "t"."createat")
```

#### Now
返回当前日期和时间。
```go
function := uast.Now()
```
Output MariaDB:
```text
NOW()
```
Output MsSQL:
```text
GETDATE()
```
Output MySQL:
```text
NOW()
```
Output PostgreSQL:
```text
CURRENT_TIMESTAMP
```
Output SQLite:
```text
DATETIME('now')
```

#### Quarter
从日期时间表达式中提取季度（1–4）。
```go
function := uast.Quarter(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
QUARTER(`t`.`createat`)
```
Output MsSQL:
```text
DATEPART(QUARTER, [t].[createat])
```
Output MySQL:
```text
QUARTER(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(QUARTER FROM "t"."createat")
```
Output SQLite:
```text
QUARTER("t"."createat")
```

#### Second
从日期时间表达式中提取秒数（0–59）。
```go
function := uast.Second(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
SECOND(`t`.`createat`)
```
Output MsSQL:
```text
DATEPART(SECOND, [t].[createat])
```
Output MySQL:
```text
SECOND(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(SECOND FROM "t"."createat")
```
Output SQLite:
```text
SECOND("t"."createat")
```

#### TimeAdd
将时间间隔添加到时间/日期时间表达式，并返回结果时间。
```go
function := uast.TimeAdd(uast.Column[time.Time]("t", "createat"), uast.Value("2 HOUR"))
```
Output MariaDB:
```text
TIME_ADD(`t`.`createat`, '2 HOUR')
```
Output MsSQL:
```text
DATEADD(HOUR, 2, [t].[createat])
```
Output MySQL:
```text
TIME_ADD(`t`.`createat`, '2 HOUR')
```
Output PostgreSQL:
```text
("t"."createat" + INTERVAL '2 HOUR')
```
Output SQLite:
```text
TIME("t"."createat", '+2 HOUR')
```

#### TimeDiff
返回两个时间/日期时间表达式之间的差值（`timeEnd` - `timeStart`）。
```go
function := uast.TimeDiff(uast.Column[time.Time]("t", "updateat"), uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
TIMEDIFF(`t`.`updateat`, `t`.`createat`)
```
Output MsSQL:
```text
TIMEDIFF([t].[updateat], [t].[createat])
```
Output MySQL:
```text
TIMEDIFF(`t`.`updateat`, `t`.`createat`)
```
Output PostgreSQL:
```text
DATE_PART('time', "t"."updateat" - "t"."createat")
```
Output SQLite:
```text
TIMEDIFF("t"."updateat", "t"."createat")
```

#### TimeSub
从时间/日期时间表达式中减去时间间隔，并返回结果时间。
```go
function := uast.TimeSub(uast.Column[time.Time]("t", "createat"), uast.Value("2 HOUR"))
```
Output MariaDB:
```text
TIME_SUB(`t`.`createat`, '2 HOUR')
```
Output MsSQL:
```text
DATEADD(HOUR, -2, [t].[createat])
```
Output MySQL:
```text
TIME_SUB(`t`.`createat`, '2 HOUR')
```
Output PostgreSQL:
```text
("t"."createat" - INTERVAL '2 HOUR')
```
Output SQLite:
```text
TIME("t"."createat", '-2 HOUR')
```

#### Week
从日期时间表达式中提取周数（1–53）。
```go
function := uast.Week(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
WEEK(`t`.`createat`)
```
Output MsSQL:
```text
DATEPART(WEEK, [t].[createat])
```
Output MySQL:
```text
WEEK(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(WEEK FROM "t"."createat")
```
Output SQLite:
```text
WEEK("t"."createat")
```

#### Year
从日期时间表达式中提取年份。
```go
function := uast.Year(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
YEAR(`t`.`createat`)
```
Output MsSQL:
```text
YEAR([t].[createat])
```
Output MySQL:
```text
YEAR(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(YEAR FROM "t"."createat")
```
Output SQLite:
```text
YEAR("t"."createat")
```

### Json
#### JsonArray
从给定表达式和可选附加值创建 JSON 数组。
```go
function := uast.JsonArray(
    uast.Column[string]("t", "json"), 
    uast.Value("val1"), 
    uast.Value("val2"),
)
```
Output MariaDB:
```text
JSON_ARRAY(`t`.`json`, ?, ?)
```
Output MsSQL:
```text
JSON_ARRAY([t].[json], @p1, @p2)
```
Output MySQL:
```text
JSON_ARRAY(`t`.`json`, ?, ?)
```
Output PostgreSQL:
```text
JSON_ARRAY("t"."json", $1, $2)
```
Output SQLite:
```text
JSON_ARRAY("t"."json", ?, ?)
```

#### JsonArrayAgg
将组中的值聚合到 JSON 数组中。
```go
function := uast.JsonArrayAgg(
    uast.Column[string]("t", "json"),
)
```
Output MariaDB:
```text
JSON_ARRAYAGG(`t`.`json`)
```
Output MsSQL:
```text
JSON_ARRAYAGG([t].[json])
```
Output MySQL:
```text
JSON_ARRAYAGG(`t`.`json`)
```
Output PostgreSQL:
```text
JSON_AGG("t"."json")
```
Output SQLite:
```text
JSON_GROUP_ARRAY("t"."json")
```

#### JsonContains
检查 JSON 文档是否包含指定值。
```go
function := uast.JsonContains(
    uast.Column[string]("t", "json"),
    uast.Value(`{"key":"val"}`),
)
```
Output MariaDB:
```text
JSON_CONTAINS(`t`.`json`, '{"key":"val"}')
```
Output MsSQL:
```text
// In development
```
Output MySQL:
```text
JSON_CONTAINS(`t`.`json`, '{"key":"val"}')
```
Output PostgreSQL:
```text
("t"."json" @> '{"key":"val"}')
```
Output SQLite:
```text
JSON_CONTAINS("t"."json", '{"key":"val"}')
```

#### JsonExtract
在指定路径从 JSON 文档中提取值。`json` 参数使用 `JsonPath` 和可选的 `JsonKey`/`JsonIndex` 构建。
```go
function := JsonExtract(
    uast.Column[string]("t", "json"), 
    uast.JsonGroup(
        uast.JsonPath(
            uast.JsonKey("parent"), 
            uast.JsonIndex(0), 
            uast.JsonKey("child"),
        ),
    ),
    uast.TypeString,
)
```
Output MariaDB:
```text
(`t`.`json` ->> '$.parent[0].child')
```
Output MsSQL:
```text
JSON_VALUE([t].[json], '$.parent[0].child')
```
Output MySQL:
```text
(`t`.`json` ->> '$.parent[0].child')
```
Output PostgreSQL:
```text
("t"."json" #>> '{parent,0,child}')
```
Output SQLite:
```text
("t"."json" ->> '$.parent[0].child')
```

#### JsonObject
从键值对构建 JSON 对象。
```go
function := uast.JsonObject(
    uast.JsonPair(
        uast.JsonKey("key"), 
        uast.Count(uast.Column[string]("t", "json"), false),
    ),
)
```
Output MariaDB:
```text
JSON_OBJECT('key', COUNT(`t`.`json`))
```
Output MsSQL:
```text
JSON_OBJECT('key', COUNT([t].[json]))
```
Output MySQL:
```text
JSON_OBJECT('key', COUNT(`t`.`json`))
```
Output PostgreSQL:
```text
JSON_BUILD_OBJECT('key', COUNT("t"."json"))
```
Output SQLite:
```text
JSON_OBJECT('key', COUNT("t"."json"))
```

#### JsonObjectAgg
将组中的键值对聚合到单个 JSON 对象中。
```go
function := uast.JsonObjectAgg(
    uast.Column[string]("t", "json"),
    uast.Column[int]("t", "number"),
)
```
Output MariaDB:
```text
JSON_OBJECTAGG(`t`.`json`, `t`.`number`)
```
Output MsSQL:
```text
JSON_OBJECTAGG([t].[json], [t].[number])
```
Output MySQL:
```text
JSON_OBJECTAGG(`t`.`json`, `t`.`number`)
```
Output PostgreSQL:
```text
JSON_OBJECT_AGG("t"."json", "t"."number")
```
Output SQLite:
```text
JSON_GROUP_OBJECT("t"."json", "t"."number")
```

#### JsonRemove
在指定路径从 JSON 文档中删除值。
```go
function := uast.JsonRemove(
    uast.Column[string]("t", "json"),
    uast.JsonGroup(
        uast.JsonPath(
            uast.JsonKey("key1"),
        ),
    ), 
    uast.JsonGroup(
        uast.JsonPath(
            uast.JsonKey("key2"),
        ),
    ),
)
```
Output MariaDB:
```text
JSON_REMOVE(`t`.`json`, '$.key1', '$.key2')
```
Output MsSQL:
```text
JSON_MODIFY(JSON_MODIFY([t].[json], '$.key1', NULL), '$.key2', NULL)
```
Output MySQL:
```text
JSON_REMOVE(`t`.`json`, '$.key1', '$.key2')
```
Output PostgreSQL:
```text
("t"."json" - '{key1}' - '{key2}')
```
Output SQLite:
```text
JSON_REMOVE("t"."json", '$.key1', '$.key2')
```

#### JsonSet
在指定路径设置 JSON 文档中的值。如果路径不存在则创建。
```go
function := uast.JsonSet(
    uast.Column[string]("t", "json"),
    uast.JsonGroup(
        uast.JsonPath(
            uast.JsonKey("key1"),
        ), 
        uast.Value("val1"),
    ),
    uast.JsonGroup(
        uast.JsonPath(
            uast.JsonKey("key2"),
        ),
        uast.Value("val2"),
    ),
)
```
Output MariaDB:
```text
JSON_SET(`t`.`json`, '$.key1', ?, '$.key2', ?)
```
Output MsSQL:
```text
JSON_MODIFY(JSON_MODIFY([t].[json], '$.key1', @p1), '$.key2', @p2)
```
Output MySQL:
```text
JSON_SET(`t`.`json`, '$.key1', ?, '$.key2', ?)
```
Output PostgreSQL:
```text
jsonb_set(jsonb_set("t"."json", '{key1}', $1), '{key2}', $2)
```
Output SQLite:
```text
JSON_SET("t"."json", '$.key1', ?, '$.key2', ?)
```

#### JsonType
返回 JSON 值的类型（例如 'OBJECT', 'ARRAY', 'STRING', 'INTEGER', 'NULL'）。
```go
function := uast.JsonType(uast.Column[string]("t", "json"))
```
Output MariaDB:
```text
JSON_TYPE(`t`.`json`)
```
Output MsSQL:
```text
// In development
```
Output MySQL:
```text
JSON_TYPE(`t`.`json`)
```
Output PostgreSQL:
```text
jsonb_typeof("t"."json")
```
Output SQLite:
```text
JSON_TYPE("t"."json")
```

### Math
#### Abs
返回数值表达式的绝对值（非负值）。
```go
function := uast.Abs(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ABS(`t`.`number`)
```
Output MsSQL:
```text
ABS([t].[number])
```
Output MySQL:
```text
ABS(`t`.`number`)
```
Output PostgreSQL:
```text
ABS("t"."number")
```
Output SQLite:
```text
ABS("t"."number")
```

#### ACos
返回表达式的反余弦（逆余弦），以弧度为单位。
```go
function := uast.ACos(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ACOS(`t`.`number`)
```
Output MsSQL:
```text
ACOS([t].[number])
```
Output MySQL:
```text
ACOS(`t`.`number`)
```
Output PostgreSQL:
```text
ACOS("t"."number")
```
Output SQLite:
```text
ACOS("t"."number")
```

#### ASin
返回表达式的反正弦（逆正弦），以弧度为单位。
```go
function := uast.ASin(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ASIN(`t`.`number`)
```
Output MsSQL:
```text
ASIN([t].[number])
```
Output MySQL:
```text
ASIN(`t`.`number`)
```
Output PostgreSQL:
```text
ASIN("t"."number")
```
Output SQLite:
```text
ASIN("t"."number")
```

#### ATan
返回表达式的反正切（逆正切），以弧度为单位。
```go
function := uast.ATan(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ATAN(`t`.`number`)
```
Output MsSQL:
```text
ATAN([t].[number])
```
Output MySQL:
```text
ATAN(`t`.`number`)
```
Output PostgreSQL:
```text
ATAN("t"."number")
```
Output SQLite:
```text
ATAN("t"."number")
```

#### ATan2
返回两个参数（`y`/`x`）商的反正切，使用它们的符号确定象限。
```go
function := uast.ATan2(uast.Column[int]("t", "y"), uast.Column[int]("t", "x"))
```
Output MariaDB:
```text
ATAN2(`t`.`y`, `t`.`x`)
```
Output MsSQL:
```text
ATAN2([t].[y], [t].[x])
```
Output MySQL:
```text
ATAN2(`t`.`y`, `t`.`x`)
```
Output PostgreSQL:
```text
ATAN2("t"."y", "t"."x")
```
Output SQLite:
```text
ATAN2("t"."y", "t"."x")
```

#### Cbrt
返回数值表达式的立方根。
```go
function := uast.Cbrt(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
CBRT(`t`.`number`)
```
Output MsSQL:
```text
CBRT([t].[number])
```
Output MySQL:
```text
CBRT(`t`.`number`)
```
Output PostgreSQL:
```text
CBRT("t"."number")
```
Output SQLite:
```text
CBRT("t"."number")
```

#### Ceil
返回不小于参数的最小整数值（向上取整）。
```go
function := uast.Ceil(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
CEILING(`t`.`number`)
```
Output MsSQL:
```text
CEILING([t].[number])
```
Output MySQL:
```text
CEILING(`t`.`number`)
```
Output PostgreSQL:
```text
CEIL("t"."number")
```
Output SQLite:
```text
CEIL("t"."number")
```

#### Cos
返回表达式的余弦，以弧度为单位。
```go
function := uast.Cos(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
COS(`t`.`number`)
```
Output MsSQL:
```text
COS([t].[number])
```
Output MySQL:
```text
COS(`t`.`number`)
```
Output PostgreSQL:
```text
COS("t"."number")
```
Output SQLite:
```text
COS("t"."number")
```

#### Exp
返回 `e`（欧拉数，~2.71828）的表达式次幂。
```go
function := uast.Exp(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
EXP(`t`.`number`)
```
Output MsSQL:
```text
EXP([t].[number])
```
Output MySQL:
```text
EXP(`t`.`number`)
```
Output PostgreSQL:
```text
EXP("t"."number")
```
Output SQLite:
```text
EXP("t"."number")
```

#### Floor
返回不大于参数的最大整数值（向下取整）。
```go
function := uast.Floor(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
FLOOR(`t`.`number`)
```
Output MsSQL:
```text
FLOOR([t].[number])
```
Output MySQL:
```text
FLOOR(`t`.`number`)
```
Output PostgreSQL:
```text
FLOOR("t"."number")
```
Output SQLite:
```text
FLOOR("t"."number")
```

#### Ln
返回表达式的自然对数（以 `e` 为底）。
```go
function := uast.Ln(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
LN(`t`.`number`)
```
Output MsSQL:
```text
LN([t].[number])
```
Output MySQL:
```text
LN(`t`.`number`)
```
Output PostgreSQL:
```text
LN("t"."number")
```
Output SQLite:
```text
LN("t"."number")
```

#### Log
返回表达式对指定底数的对数。
```go
function := uast.Log(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
LOG(`t`.`number`, ?)
```
Output MsSQL:
```text
LOG([t].[number], @p1)
```
Output MySQL:
```text
LOG(`t`.`number`, ?)
```
Output PostgreSQL:
```text
LOG("t"."number", $1)
```
Output SQLite:
```text
LOG("t"."number", ?)
```

#### Mod
返回第一个表达式除以第二个表达式的余数（模）。
```go
function := uast.Mod(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
MOD(`t`.`number`, ?)
```
Output MsSQL:
```text
MOD([t].[number], @p1)
```
Output MySQL:
```text
MOD(`t`.`number`, ?)
```
Output PostgreSQL:
```text
MOD("t"."number", $1)
```
Output SQLite:
```text
MOD("t"."number", ?)
```

#### Pi
返回数学常数 `π`（~3.14159）。
```go
function := uast.Pi()
```
Output MariaDB:
```text
PI()
```
Output MsSQL:
```text
PI()
```
Output MySQL:
```text
PI()
```
Output PostgreSQL:
```text
PI()
```
Output SQLite:
```text
PI()
```

#### Power
返回表达式对 exponent 次幂的值。
```go
function := uast.Power(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
POWER(`t`.`number`, ?)
```
Output MsSQL:
```text
POWER([t].[number], @p1)
```
Output MySQL:
```text
POWER(`t`.`number`, ?)
```
Output PostgreSQL:
```text
POWER("t"."number", $1)
```
Output SQLite:
```text
POWER("t"."number", ?)
```

#### Rand
返回范围 [0, 1] 内的随机浮点值。
```go
function := uast.Rand()
```
Output MariaDB:
```text
RAND()
```
Output MsSQL:
```text
RAND()
```
Output MySQL:
```text
RAND()
```
Output PostgreSQL:
```text
RANDOM()
```
Output SQLite:
```text
RANDOM()
```

#### Round
将表达式四舍五入到指定的小数位数。
```go
function := uast.Round(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
ROUND(`t`.`number`, ?)
```
Output MsSQL:
```text
ROUND([t].[number], @p1)
```
Output MySQL:
```text
ROUND(`t`.`number`, ?)
```
Output PostgreSQL:
```text
ROUND("t"."number", $1)
```
Output SQLite:
```text
ROUND("t"."number", ?)
```

#### Sin
返回表达式的正弦，以弧度为单位。
```go
function := uast.Sin(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
SIN(`t`.`number`)
```
Output MsSQL:
```text
SIN([t].[number])
```
Output MySQL:
```text
SIN(`t`.`number`)
```
Output PostgreSQL:
```text
SIN("t"."number")
```
Output SQLite:
```text
SIN("t"."number")
```

#### Sqrt
返回表达式的平方根。
```go
function := uast.Sqrt(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
SQRT(`t`.`number`)
```
Output MsSQL:
```text
SQRT([t].[number])
```
Output MySQL:
```text
SQRT(`t`.`number`)
```
Output PostgreSQL:
```text
SQRT("t"."number")
```
Output SQLite:
```text
SQRT("t"."number")
```

#### Tan
返回表达式的正切，以弧度为单位。
```go
function := uast.Tan(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
TAN(`t`.`number`)
```
Output MsSQL:
```text
TAN([t].[number])
```
Output MySQL:
```text
TAN(`t`.`number`)
```
Output PostgreSQL:
```text
TAN("t"."number")
```
Output SQLite:
```text
TAN("t"."number")
```

#### Trunc
将数值表达式截断到指定的小数位数（不进行四舍五入）。
```go
function := uast.Trunc(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
TRUNCATE(`t`.`number`, ?)
```
Output MsSQL:
```text
ROUND([t].[number], @p1, 1)
```
Output MySQL:
```text
TRUNCATE(`t`.`number`, ?)
```
Output PostgreSQL:
```text
TRUNC("t"."number", $1)
```
Output SQLite:
```text
TRUNC("t"."number", ?)
```

### Ranking
#### CumeDist
返回分区内值的累积分布（在当前行之前或与当前行相等的行数比率）。必须与 `OVER` 子句一起使用。
```go
function := uast.CumeDist().Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
CUME_DIST() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
CUME_DIST() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
CUME_DIST() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
CUME_DIST() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
CUME_DIST() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### DenseRank
返回无间隙的行排名。相等值的行获得相同的排名，下一个排名是紧接着的下一个整数。需要 `OVER`。
```go
function := uast.DenseRank().Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
DENSE_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
DENSE_RANK() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
DENSE_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
DENSE_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
DENSE_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### NTile
将分区内的行划分为 `n` 个近似相等的组，并返回每行的组号（1 到 `n`）。
```go
function := uast.NTile(2).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
NTILE(2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
NTILE(2) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
NTILE(2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
NTILE(2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
NTILE(2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### PercentRank
返回分区内行的百分位排名（范围 0 到 1）。第一行的排名始终为 0。需要 `OVER`。
```go
function := uast.PercentRank().Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
PERCENT_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
PERCENT_RANK() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
PERCENT_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
PERCENT_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
PERCENT_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### Rank
返回有间隙的行排名。相等值获得相同排名，下一个不同值跳过排名。需要 `OVER`。
```go
function := uast.Rank().Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
RANK() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### RowNumber
为分区内的每一行分配一个唯一的序号，从 1 开始。顺序决定编号序列。
```go
function := uast.RowNumber().Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
ROW_NUMBER() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
ROW_NUMBER() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
ROW_NUMBER() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
ROW_NUMBER() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
ROW_NUMBER() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

### String
#### Concat
将两个或多个字符串表达式连接成一个字符串。`NULL` 参数在大多数方言中被视为空字符串。
```go
function := uast.Concat(uast.Column[string]("t", "string"), uast.Value("old"), uast.Value("new"))
```
Output MariaDB:
```text
CONCAT(`t`.`string`, ?, ?)
```
Output MsSQL:
```text
CONCAT([t].[string], @p1, @p2)
```
Output MySQL:
```text
CONCAT(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
CONCAT("t"."string", $1, $2)
```
Output SQLite:
```text
CONCAT("t"."string", ?, ?)
```

#### ConcatWs
用指定的分隔符连接两个或多个字符串表达式。跳过 `NULL` 参数。
```go
function := uast.ConcatWs(uast.Value("_"), uast.Column[string]("t", "string"), uast.Value("old"),uast.Value("new"))
```
Output MariaDB:
```text
CONCAT_WS(?, `t`.`string`, ?, ?)
```
Output MsSQL:
```text
CONCAT_WS(@p1, [t].[string], @p2, @p3)
```
Output MySQL:
```text
CONCAT_WS(?, `t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
CONCAT_WS($1, "t"."string", $2, $3)
```
Output SQLite:
```text
CONCAT_WS(?, "t"."string", ?, ?)
```

#### LeftString
返回字符串表达式最左边的 `count` 个字符。
```go
function := uast.LeftString(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MariaDB:
```text
LEFT(`t`.`string`, ?)
```
Output MsSQL:
```text
LEFT([t].[string], @p1)
```
Output MySQL:
```text
LEFT(`t`.`string`, ?)
```
Output PostgreSQL:
```text
LEFT("t"."string", $1)
```
Output SQLite:
```text
LEFT("t"."string", ?)
```

#### Lower
将字符串表达式转换为小写。
```go
function := uast.Lower(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
LOWER(`t`.`string`)
```
Output MsSQL:
```text
LOWER([t].[string])
```
Output MySQL:
```text
LOWER(`t`.`string`)
```
Output PostgreSQL:
```text
LOWER("t"."string")
```
Output SQLite:
```text
LOWER("t"."string")
```

#### LPad
用指定的分隔符在左侧填充字符串表达式，使其总长度达到 `count` 个字符。
```go
function := uast.LPad(uast.Column[string]("t", "string"), uast.Value(2), uast.Value(","))
```
Output MariaDB:
```text
LPAD(`t`.`string`, ?, ?)
```
Output MsSQL:
```text
LPAD([t].[string], @p1, @p2)
```
Output MySQL:
```text
LPAD(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
LPAD("t"."string", $1, $2)
```
Output SQLite:
```text
LPAD("t"."string", ?, ?)
```

#### LTrim
删除字符串表达式的前导空格。
```go
function := uast.LTrim(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
LTRIM(`t`.`string`)
```
Output MsSQL:
```text
LTRIM([t].[string])
```
Output MySQL:
```text
LTRIM(`t`.`string`)
```
Output PostgreSQL:
```text
LTRIM("t"."string")
```
Output SQLite:
```text
LTRIM("t"."string")
```

#### Repeat
将字符串表达式重复 `count` 次。
```go
function := uast.Repeat(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MariaDB:
```text
REPEAT(`t`.`string`, ?)
```
Output MsSQL:
```text
REPEAT([t].[string], @p1)
```
Output MySQL:
```text
REPEAT(`t`.`string`, ?)
```
Output PostgreSQL:
```text
REPEAT("t"."string", $1)
```
Output SQLite:
```text
REPEAT("t"."string", ?)
```

#### Replace
将字符串中所有出现的子字符串替换为新的子字符串。
```go
function := uast.Replace(uast.Column[string]("t", "string"), uast.Value("old"), uast.Value("new"))
```
Output MariaDB:
```text
REPLACE(`t`.`string`, ?, ?)
```
Output MsSQL:
```text
REPLACE([t].[string], @p1, @p2)
```
Output MySQL:
```text
REPLACE(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
REPLACE("t"."string", $1, $2)
```
Output SQLite:
```text
REPLACE("t"."string", ?, ?)
```

#### Reverse
反转字符串表达式中的字符。
```go
function := uast.Reverse(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
REVERSE(`t`.`string`)
```
Output MsSQL:
```text
REVERSE([t].[string])
```
Output MySQL:
```text
REVERSE(`t`.`string`)
```
Output PostgreSQL:
```text
REVERSE("t"."string")
```
Output SQLite:
```text
REVERSE("t"."string")
```

#### RightString
返回字符串表达式最右边的 `count` 个字符。
```go
function := uast.RightString(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MariaDB:
```text
RIGHT(`t`.`string`, ?)
```
Output MsSQL:
```text
RIGHT([t].[string], @p1)
```
Output MySQL:
```text
RIGHT(`t`.`string`, ?)
```
Output PostgreSQL:
```text
RIGHT("t"."string", $1)
```
Output SQLite:
```text
RIGHT("t"."string", ?)
```

#### RPad
用指定的分隔符在右侧填充字符串表达式，使其总长度达到 `count` 个字符。
```go
function := uast.RPad(uast.Column[string]("t", "string"), uast.Value(2), uast.Value(","))
```
Output MariaDB:
```text
RPAD(`t`.`string`, ?, ?)
```
Output MsSQL:
```text
RPAD([t].[string], @p1, @p2)
```
Output MySQL:
```text
RPAD(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
RPAD("t"."string", $1, $2)
```
Output SQLite:
```text
RPAD("t"."string", ?, ?)
```

#### RTrim
删除字符串表达式的尾部空格。
```go
function := uast.RTrim(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
RTRIM(`t`.`string`)
```
Output MsSQL:
```text
RTRIM([t].[string])
```
Output MySQL:
```text
RTRIM(`t`.`string`)
```
Output PostgreSQL:
```text
RTRIM("t"."string")
```
Output SQLite:
```text
RTRIM("t"."string")
```

#### SubString
从字符串表达式中提取子字符串，从 `startPos`（基于 1）开始，长度为 `lengthStr` 个字符。
```go
function := uast.SubString(uast.Column[string]("t", "string"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
SUBSTRING(`t`.`string`, ?, ?)
```
Output MsSQL:
```text
SUBSTRING([t].[string], @p1, @p2)
```
Output MySQL:
```text
SUBSTRING(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
SUBSTRING("t"."string", $1, $2)
```
Output SQLite:
```text
SUBSTRING("t"."string", ?, ?)
```

#### Trim
删除字符串表达式的前导和尾部空格。
```go
function := uast.Trim(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
TRIM(`t`.`string`)
```
Output MsSQL:
```text
TRIM([t].[string])
```
Output MySQL:
```text
TRIM(`t`.`string`)
```
Output PostgreSQL:
```text
TRIM("t"."string")
```
Output SQLite:
```text
TRIM("t"."string")
```

#### Upper
将字符串表达式转换为大写。
```go
function := uast.Upper(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
UPPER(`t`.`string`)
```
Output MsSQL:
```text
UPPER([t].[string])
```
Output MySQL:
```text
UPPER(`t`.`string`)
```
Output PostgreSQL:
```text
UPPER("t"."string")
```
Output SQLite:
```text
UPPER("t"."string")
```

## exprLiteral
### Literal
将原始字面量值直接嵌入生成的 SQL 字符串（不参数化）。请谨慎使用 — 值按原样写入。对于用户提供的数据，请优先使用 `Value`。
```go
literal := uast.Literal("%Y-%m-%d")
```
Output:
```text
'%Y-%m-%d'
```

## exprLogical
### And
通过逻辑 `AND` 组合多个条件。要使组合表达式为真，所有条件都必须为真。
```go
logical := uast.And(
    uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    uast.Greater(uast.Column[int]("t", "number"), uast.Value(2)),
)
```
Output MariaDB:
```text
(`t`.`string` = ? AND `t`.`number` > ?)
```
Output MsSQL:
```text
([t].[string] = @p1 AND [t].[number] > @p2)
```
Output MySQL:
```text
(`t`.`string` = ? AND `t`.`number` > ?)
```
Output PostgreSQL:
```text
("t"."string" = $1 AND "t"."number" > $2)
```
Output SQLite:
```text
("t"."string" = ? AND "t"."number" > ?)
```

### Or
通过逻辑 `OR` 组合多个条件。要使组合表达式为真，至少有一个条件为真。
```go
logical := uast.Or(
    uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    uast.Greater(uast.Column[int]("t", "number"), uast.Value(2)),
)
```
Output MariaDB:
```text
(`t`.`string` = ? OR `t`.`number` > ?)
```
Output MsSQL:
```text
([t].[string] = @p1 OR [t].[number] > @p2)
```
Output MySQL:
```text
(`t`.`string` = ? OR `t`.`number` > ?)
```
Output PostgreSQL:
```text
("t"."string" = $1 OR "t"."number" > $2)
```
Output SQLite:
```text
("t"."string" = ? OR "t"."number" > ?)
```

## exprSubquery
### Subquery
将 `SELECT` 语句包装为类型化表达式，可用于比较（`In`、`Exists`、`Equal` 等）或作为 `SELECT` 子句中的列。泛型参数 `T` 指定子查询返回的单个列的标量类型。
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Column[int64]("t", "id")).From(uast.NewTable("test").As("t")))
```
Output MariaDB:
```text
(SELECT `t`.`id` FROM `test` AS `t`)
```
Output MsSQL:
```text
(SELECT [t].[id] FROM [test] AS [t])
```
Output MySQL:
```text
(SELECT `t`.`id` FROM `test` AS `t`)
```
Output PostgreSQL:
```text
(SELECT "t"."id" FROM "test" AS "t")
```
Output SQLite:
```text
(SELECT "t"."id" FROM "test" AS "t")
```

## exprValue
### Value
将 Go 值包装为参数化表达式。该值不会直接插入 SQL 字符串中 — 而是生成一个占位符（`?`、`$1` 等），并将该值附加到 `Build()` 返回的参数切片中。这是传递用户提供的数据并防止 SQL 注入的安全方式。
支持的类型：`bool`、`float32`、`float64`、`int`、`int8`、`int16`、`int32`、`int64`、`uint`、`uint8`、`uint16`、`uint32`、`uint64`、`string`、`time.Time`。
```go
var data string = "ivan"
value := uast.Value(data)
```
Output MariaDB:
```text
?
```
Output MsSQL:
```text
@p1
```
Output MySQL:
```text
?
```
Output PostgreSQL:
```text
$1
```
Output SQLite:
```text
?
```