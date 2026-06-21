---
outline: deep
---

# API / Core / Options

::: info **Info**
This page covers all configuration options: `clauseGroupBy`, `clauseHaving`, `clauseJoin`, `clauseOrderBy`, `clausePagination`, `clauseReturning`, `clauseSet`, `clauseUnions`, `clauseValues`, `clauseWhere`, `clauseWith`, `exprArray`, `exprBinary`, `exprComparison`, `exprConstant`, `exprField`, `exprFunction`, `exprLiteral`, `exprLogical`, `exprSubquery`, `exprValue`. Each option is shown with a working code example and expected output.
:::

## clauseGroupBy
Adds a GROUP BY clause to group rows by specified columns or expressions.
```go
groupBy := GroupBy(
	uast.Field[string]("u", "string"),
)
```
Output MariaDB:
```text
GROUP BY `u`.`string`
```
Output MsSQL:
```text
GROUP BY [u].[string]
```
Output MySQL:
```text
GROUP BY `u`.`string`
```
Output PostgreSQL:
```text
GROUP BY "u"."string"
```
Output SQLite:
```text
GROUP BY "u"."string"
```

## clauseHaving
Adds a HAVING clause to filter groups. Used with GROUP BY to filter aggregated results.
```go
having := Having(
	uast.Greater(uast.Count(uast.Field[int64]("u", "id"), false), uast.Value[int64](2)),
)
```
Output MariaDB:
```text
HAVING COUNT(`u`.`id`) > ?
```
Output MsSQL:
```text
HAVING COUNT([u].[id]) > @p1
```
Output MySQL:
```text
HAVING COUNT(`u`.`id`) > ?
```
Output PostgreSQL:
```text
HAVING COUNT("u"."id") > $1
```
Output SQLite:
```text
HAVING COUNT("u"."id") > ?
```

## clauseJoin
### Cross
Adds a CROSS JOIN to the statement. Returns the Cartesian product of both tables.
```go
join := uast.Cross(uast.NewTable("users").As("u"))
```
Output MariaDB:
```text
CROSS JOIN `users` AS `u`
```
Output MsSQL:
```text
CROSS JOIN [users] AS [u]
```
Output MySQL:
```text
CROSS JOIN `users` AS `u`
```
Output PostgreSQL:
```text
CROSS JOIN "users" AS "u"
```
Output SQLite:
```text
CROSS JOIN "users" AS "u"
```

### Full
Adds a FULL JOIN to the statement. Returns all rows from both tables, with NULLs where there is no match.
```go
join := uast.Full(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
FULL JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
FULL JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
FULL JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
FULL JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
FULL JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### FullOuter
Adds a FULL OUTER JOIN to the statement. Returns all rows from both tables, with NULLs where there is no match.
```go
join := uast.FullOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
FULL OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
FULL OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
FULL OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
FULL OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
FULL OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Inner
Adds an INNER JOIN to the statement. Returns rows that have matching values in both tables.
```go
join := uast.Inner(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
INNER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
INNER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
INNER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
INNER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
INNER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Left
Adds a LEFT JOIN to the statement. Returns all rows from the left table, and matching rows from the right table.
```go
join := uast.Left(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
LEFT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
LEFT JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
LEFT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
LEFT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
LEFT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### LeftOuter
Adds a LEFT OUTER JOIN to the statement. Returns all rows from the left table, and matching rows from the right table.
```go
join := uast.LeftOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
LEFT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
LEFT OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
LEFT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
LEFT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
LEFT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Right
Adds a RIGHT JOIN to the statement. Returns all rows from the right table, and matching rows from the left table. Not supported by SQLite.
```go
join := uast.Right(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
RIGHT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
RIGHT JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
RIGHT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
RIGHT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
// Not supported
```

### RightOuter
Adds a RIGHT OUTER JOIN to the statement. Returns all rows from the right table, and matching rows from the left table. Not supported by SQLite.
```go
join := uast.RightOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
RIGHT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
RIGHT OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
RIGHT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
RIGHT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
// Not supported
```

## clauseOrderBy
### Asc
Specifies ascending sort order (smallest first, A-to-Z). Used for sorting rows in a query or within a window function.
```go
orderBy := uast.Asc(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
`u`.`string` ASC
```
Output MsSQL:
```text
[u].[string] ASC
```
Output MySQL:
```text
`u`.`string` ASC
```
Output PostgreSQL:
```text
"u"."string" ASC
```
Output SQLite:
```text
"u"."string" ASC
```

### Desc
Specifies descending sort order (largest first, Z-to-A). Used for sorting rows in a query or within a window function.
```go
orderBy := uast.Desc(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
`u`.`string` DESC
```
Output MsSQL:
```text
[u].[string] DESC
```
Output MySQL:
```text
`u`.`string` DESC
```
Output PostgreSQL:
```text
"u"."string" DESC
```
Output SQLite:
```text
"u"."string" DESC
```

## clausePagination
Specifies pagination for a SELECT statement using `Pagination(limit, offset)`. The `limit` sets the maximum number of rows to return. The `offset` specifies the number of rows to skip before returning results. The rendering order and syntax adapts to each dialect automatically.
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
Adds a RETURNING clause to return modified rows. Supported by MariaDB, PostgreSQL, and SQLite. MySQL does not support this clause natively.
```go
returning = Returning(
	uast.Field[int64]("u", "id")
    uast.Field[string]("u", "string")
)
```
Output MariaDB:
```text
RETURNING `u`.`id`, `u`.`string`
```
Output MsSQL:
```text
OUTPUT [u].[id], [u].[string]
```
Output MySQL:
```text
// Not support
```
Output PostgreSQL:
```text
RETURNING "u"."id", "u"."string"
```
Output SQLite:
```text
RETURNING "u"."id", "u"."string"
```

## clauseSet
### Assign
Specifies columns and their new values using `Assign` to associate columns with values. Supports multiple pairs for updating multiple columns.
```go
set := Set(
	uast.Assign(uast.Field[string]("u", "string"), uast.Value("active")),
)
```
Output MariaDB:
```text
UPDATE `users` AS `u` SET `u`.`string` = ?
```
Output MsSQL:
```text
UPDATE [users] AS [u] SET [u].[string] = @p1
```
Output MySQL:
```text
UPDATE `users` AS `u` SET `u`.`string` = ?
```
Output PostgreSQL:
```text
UPDATE "users" AS "u" SET "u"."string" = $1
```
Output SQLite:
```text
UPDATE "users" AS "u" SET "u"."string" = ?
```

## clauseUnions
### Union
Combines results from multiple SELECT statements. UNION returns distinct rows.
```go
unions := uast.Union(uast.NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[string]("u", "string"),
    ),
)
```
Output MariaDB:
```text
UNION SELECT `u`.`string` FROM `users` AS `u` 
```
Output MsSQL:
```text
UNION SELECT [u].[string] FROM [users] AS [u]
```
Output MySQL:
```text
UNION SELECT `u`.`string` FROM `users` AS `u`
```
Output PostgreSQL:
```text
UNION SELECT "u"."string" FROM "users" AS "u"
```
Output SQLite:
```text
UNION SELECT "u"."string" FROM "users" AS "u"
```

### UnionAll
Combines results from multiple SELECT statements. UNION ALL returns all rows, including duplicates.
```go
unions := uast.UnionAll(uast.NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[string]("u", "string"),
    ),
)
```
Output MariaDB:
```text
UNION ALL SELECT `u`.`string` FROM `users` AS `u`
```
Output MsSQL:
```text
UNION ALL SELECT [u].[string] FROM [users] AS [u]
```
Output MySQL:
```text
UNION ALL SELECT `u`.`string` FROM `users` AS `u`
```
Output PostgreSQL:
```text
UNION ALL SELECT "u"."string" FROM "users" AS "u"
```
Output SQLite:
```text
UNION ALL SELECT "u"."string" FROM "users" AS "u"
```

### UnionExcept
Combines results from multiple SELECT statements. EXCEPT returns distinct rows from the first query that are not in the second.
```go
unions := uast.UnionExcept(uast.NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[string]("u", "string"),
    ),
)
```
Output MariaDB:
```text
EXCEPT SELECT `u`.`string` FROM `users` AS `u`
```
Output MsSQL:
```text
EXCEPT SELECT [u].[string] FROM [users] AS [u]
```
Output MySQL:
```text
EXCEPT SELECT `u`.`string` FROM `users` AS `u`
```
Output PostgreSQL:
```text
EXCEPT SELECT "u"."string" FROM "users" AS "u"
```
Output SQLite:
```text
EXCEPT SELECT "u"."string" FROM "users" AS "u"
```

### UnionIntersect
Combines results from multiple SELECT statements. INTERSECT returns distinct rows that are common to both queries.
```go
unions := uast.UnionIntersect(uast.NewSelect(uast.NewTable("users").As("u")).
	Field(
		uast.Field[string]("u", "string"),
	),
)
```
Output MariaDB:
```text
INTERSECT SELECT `u`.`string` FROM `users` AS `u`
```
Output MsSQL:
```text
INTERSECT SELECT [u].[string] FROM [users] AS [u]
```
Output MySQL:
```text
INTERSECT SELECT `u`.`string` FROM `users` AS `u`
```
Output PostgreSQL:
```text
INTERSECT SELECT "u"."string" FROM "users" AS "u"
```
Output SQLite:
```text
INTERSECT SELECT "u"."string" FROM "users" AS "u"
```

## clauseValues
### Pair
Specifies values for insertion using `Pair` to associate columns with values. Columns are automatically inferred from the pairs.
```go
values := Values(
    uast.Pair(uast.Field[string]("u", "string"), uast.Value("ivan")),
	uast.Pair(uast.Field[int]("u", "number"), uast.Value(2)),
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
Adds an upsert clause to INSERT ... VALUES using `Upsert`. Associates columns with values.
```go
values := Values(
    uast.Pair(uast.Field[string]("u", "string"), uast.Value("ivan")),
	uast.Pair(uast.Field[int]("u", "number"), uast.Value(2)),
).
Upsert(
    uast.Pair(uast.Field[string]("u", "string"), uast.Value("updated")),
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
Adds a WHERE clause to filter rows before grouping or aggregation. Accepts comparison expressions, logical operators, and subqueries.
```go
where = Where(
	uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
)
```
Output MariaDB:
```text
WHERE `u`.`string` = ?
```
Output MsSQL:
```text
WHERE [u].[string] = @p1
```
Output MySQL:
```text
WHERE `u`.`string` = ?
```
Output PostgreSQL:
```text
WHERE "u"."string" = $1
```
Output SQLite:
```text
WHERE "u"."string" = ?
```

## clauseWith
### Norecursive
Adds a norecursive Common Table Expression (CTE) to the statement using `WithN`. Columns are aliased via the variadic string arguments.
```go
with := WithN("cte_norecursive", NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[int64]("u", "id"),
        uast.Field[string]("u", "string"),
    ).
    Where(
        uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    ),
    "id", "string",
)
```
Output MariaDB:
```text
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ?)
```
Output MsSQL:
```text
WITH [cte_norecursive] ([id], [string]) AS (SELECT [u].[id], [u].[string] FROM [users] AS [u] WHERE [u].[string] = @p1)
```
Output MySQL:
```text
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ?)
```
Output PostgreSQL:
```text
WITH "cte_norecursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = $1)
```
Output SQLite:
```text
WITH "cte_norecursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = ?)
```

### Recursive
Adds a recursive Common Table Expression (CTE) to the statement using `WithR`. Requires a `Unions` clause with `UnionAll` to define the recursive step.
```go
with := WithR("cte_recursive", NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[int64]("u", "id"),
        uast.Field[string]("u", "string"),
    ).
    Where(
        uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    ).
    Unions(
        uast.UnionAll(uast.NewSelect(uast.NewTable("users").As("u")).
            Field(
                uast.Field[int64]("u", "id"),
                uast.Field[string]("u", "string"),
            ).
            Join(
                uast.Inner(uast.NewCTE("cte_recursive", "rec"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("rec", "id"))),
            ),
        ),
    ),
    "id", "string",
)
```
Output MariaDB:
```text
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ? UNION ALL SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` INNER JOIN `cte_recursive` AS `rec` ON `u`.`id` = `rec`.`id`)
```
Output MsSQL:
```text
WITH RECURSIVE [cte_recursive] ([id], [string]) AS (SELECT [u].[id], [u].[string] FROM [users] AS [u] WHERE [u].[string] = @p1 UNION ALL SELECT [u].[id], [u].[string] FROM [users] AS [u] INNER JOIN [cte_recursive] AS [rec] ON [u].[id] = [rec].[id])
```
Output MySQL:
```text
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ? UNION ALL SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` INNER JOIN `cte_recursive` AS `rec` ON `u`.`id` = `rec`.`id`)
```
Output PostgreSQL:
```text
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = $1 UNION ALL SELECT "u"."id", "u"."string" FROM "users" AS "u" INNER JOIN "cte_recursive" AS "rec" ON "u"."id" = "rec"."id")
```
Output SQLite:
```text
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = ? UNION ALL SELECT "u"."id", "u"."string" FROM "users" AS "u" INNER JOIN "cte_recursive" AS "rec" ON "u"."id" = "rec"."id")
```

## exprArray
### Array
Constructs an array expression for use in SQL queries.
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
Performs a bitwise AND operation between two expressions.
```go
binary := uast.BitwiseAnd(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`u`.`number` & ?
```
Output MsSQL:
```text
[u].[number] & @p1
```
Output MySQL:
```text
`u`.`number` & ?
```
Output PostgreSQL:
```text
"u"."number" & $1
```
Output SQLite:
```text
"u"."number" & ?
```

### BitwiseOr
Performs a bitwise OR operation between two expressions.
```go
binary := uast.BitwiseOr(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`u`.`number` | ?
```
Output MsSQL:
```text
[u].[number] | @p1
```
Output MySQL:
```text
`u`.`number` | ?
```
Output PostgreSQL:
```text
"u"."number" | $1
```
Output SQLite:
```text
"u"."number" | ?
```

### BitwiseXor
Performs a bitwise XOR operation between two expressions.
```go
binary := uast.BitwiseXor(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`u`.`number` ^ ?
```
Output MsSQL:
```text
[u].[number] ^ @p1
```
Output MySQL:
```text
`u`.`number` ^ ?
```
Output PostgreSQL:
```text
"u"."number" ^ $1
```
Output SQLite:
```text
"u"."number" ^ ?
```

### Divide
Divides the left expression by the right expression.
```go
binary := uast.Divide(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` / ?
```
Output MsSQL:
```text
[u].[number] / @p1
```
Output MySQL:
```text
`u`.`number` / ?
```
Output PostgreSQL:
```text
"u"."number" / $1
```
Output SQLite:
```text
"u"."number" / ?
```

### Minus
Subtracts the right expression from the left expression.
```go
binary := uast.Minus(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` - ?
```
Output MsSQL:
```text
[u].[number] - @p1
```
Output MySQL:
```text
`u`.`number` - ?
```
Output PostgreSQL:
```text
"u"."number" - $1
```
Output SQLite:
```text
"u"."number" - ?
```

### Modulo
Returns the remainder of dividing the left expression by the right expression.
```go
binary := uast.Modulo(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` % ?
```
Output MsSQL:
```text
[u].[number] % @p1
```
Output MySQL:
```text
`u`.`number` % ?
```
Output PostgreSQL:
```text
"u"."number" % $1
```
Output SQLite:
```text
"u"."number" % ?
```

### Multiply
Multiplies the left expression by the right expression.
```go
binary := uast.Multiply(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` * ?
```
Output MsSQL:
```text
[u].[number] * @p1
```
Output MySQL:
```text
`u`.`number` * ?
```
Output PostgreSQL:
```text
"u"."number" * $1
```
Output SQLite:
```text
"u"."number" * ?
```

### Plus
Adds the left expression to the right expression.
```go
binary := uast.Plus(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` + ?
```
Output MsSQL:
```text
[u].[number] + @p1
```
Output MySQL:
```text
`u`.`number` + ?
```
Output PostgreSQL:
```text
"u"."number" + $1
```
Output SQLite:
```text
"u"."number" + ?
```

### ShiftLeft
Performs a bitwise left shift on the left expression by the number of bits specified in the right expression.
```go
binary := uast.ShiftLeft(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` << ?
```
Output MsSQL:
```text
[u].[number] << @p1
```
Output MySQL:
```text
`u`.`number` << ?
```
Output PostgreSQL:
```text
"u"."number" << $1
```
Output SQLite:
```text
"u"."number" << ?
```

### ShiftRight
Performs a bitwise right shift on the left expression by the number of bits specified in the right expression.
```go
binary := uast.ShiftRight(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` >> ?
```
Output MsSQL:
```text
[u].[number] >> @p1
```
Output MySQL:
```text
`u`.`number` >> ?
```
Output PostgreSQL:
```text
"u"."number" >> $1
```
Output SQLite:
```text
"u"."number" >> ?
```

## exprComparison
### Between
Checks if the left expression falls within the range defined by `valueStart` and `valueEnd` (inclusive).
```go
comparison := uast.Between(uast.Field[int]("u", "number"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` BETWEEN ? AND ?
```
Output MsSQL:
```text
[u].[number] BETWEEN @p1 AND @p2
```
Output MySQL:
```text
`u`.`number` BETWEEN ? AND ?
```
Output PostgreSQL:
```text
"u"."number" BETWEEN $1 AND $2
```
Output SQLite:
```text
"u"."number" BETWEEN ? AND ?
```

### Equal
Compares two expressions for equality (`=`).
```go
comparison := uast.Equal(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` = ?
```
Output MsSQL:
```text
[u].[number] = @p1
```
Output MySQL:
```text
`u`.`number` = ?
```
Output PostgreSQL:
```text
"u"."number" = $1
```
Output SQLite:
```text
"u"."number" = ?
```

### Exists
Checks if the subquery returns any rows. Returns `true` if at least one row exists.
```go
comparison := uast.Exists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("users").As("u"))))
```
Output MariaDB:
```text
EXISTS (SELECT 1 FROM `users` AS `u`)
```
Output MsSQL:
```text
EXISTS (SELECT 1 FROM [users] AS [u])
```
Output MySQL:
```text
EXISTS (SELECT 1 FROM `users` AS `u`)
```
Output PostgreSQL:
```text
EXISTS (SELECT 1 FROM "users" AS "u")
```
Output SQLite:
```text
EXISTS (SELECT 1 FROM "users" AS "u")
```

### Greater
Compares if the left expression is greater than the right expression (`>`).
```go
comparison := uast.Greater(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` > ?
```
Output MsSQL:
```text
[u].[number] > @p1
```
Output MySQL:
```text
`u`.`number` > ?
```
Output PostgreSQL:
```text
"u"."number" > $1
```
Output SQLite:
```text
"u"."number" > ?
```

### GreaterEqual
Compares if the left expression is greater than or equal to the right expression (`>=`).
```go
comparison := uast.GreaterEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` >= ?
```
Output MsSQL:
```text
[u].[number] >= @p1
```
Output MySQL:
```text
`u`.`number` >= ?
```
Output PostgreSQL:
```text
"u"."number" >= $1
```
Output SQLite:
```text
"u"."number" >= ?
```

### ILike
Performs a case-insensitive pattern matching comparison. The right expression should contain a pattern with `%` (any sequence) and `_` (single character) wildcards.
```go
comparison := uast.ILike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
LOWER(`u`.`string`) LIKE LOWER(?)
```
Output MsSQL:
```text
LOWER([u].[string]) LIKE LOWER(@p1)
```
Output MySQL:
```text
LOWER(`u`.`string`) LIKE LOWER(?)
```
Output PostgreSQL:
```text
"u"."string" ILIKE $1
```
Output SQLite:
```text
LOWER("u"."string") LIKE LOWER(?)
```

### In
Checks if the left expression matches any value contained within the right expression (typically a subquery or array).
```go
comparison := uast.In(uast.Field[string]("u", "string"), uast.Array("active", "pending"))
```
Output MariaDB:
```text
`u`.`string` IN (?, ?)
```
Output MsSQL:
```text
[u].[string] IN (@p1, @p2)
```
Output MySQL:
```text
`u`.`string` IN (?, ?)
```
Output PostgreSQL:
```text
"u"."string" IN ($1, $2)
```
Output SQLite:
```text
"u"."string" IN (?, ?)
```

### IsNotNull
Checks if the expression is not `NULL`.
```go
comparison := uast.IsNotNull(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
`u`.`string` IS NOT NULL
```
Output MsSQL:
```text
[u].[string] IS NOT NULL
```
Output MySQL:
```text
`u`.`string` IS NOT NULL
```
Output PostgreSQL:
```text
"u"."string" IS NOT NULL
```
Output SQLite:
```text
"u"."string" IS NOT NULL
```

### IsNull
Checks if the expression is `NULL`.
```go
comparison := uast.IsNull(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
`u`.`string` IS NULL
```
Output MsSQL:
```text
[u].[string] IS NULL
```
Output MySQL:
```text
`u`.`string` IS NULL
```
Output PostgreSQL:
```text
"u"."string" IS NULL
```
Output SQLite:
```text
"u"."string" IS NULL
```

### Less
Compares if the left expression is less than the right expression (`<`).
```go
comparison := uast.Less(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` < ?
```
Output MsSQL:
```text
[u].[number] < @p1
```
Output MySQL:
```text
`u`.`number` < ?
```
Output PostgreSQL:
```text
"u"."number" < $1
```
Output SQLite:
```text
"u"."number" < ?
```

### LessEqual
Compares if the left expression is less than or equal to the right expression (`<=`).
```go
comparison := uast.LessEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` <= ?
```
Output MsSQL:
```text
[u].[number] <= @p1
```
Output MySQL:
```text
`u`.`number` <= ?
```
Output PostgreSQL:
```text
"u"."number" <= $1
```
Output SQLite:
```text
"u"."number" <= ?
```

### Like
Performs a case-sensitive pattern matching comparison. The right expression should contain a pattern with `%` and `_` wildcards.
```go
comparison := uast.Like(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
`u`.`string` LIKE ?
```
Output MsSQL:
```text
[u].[number] LIKE @p1
```
Output MySQL:
```text
`u`.`string` LIKE ?
```
Output PostgreSQL:
```text
"u"."string" LIKE $1
```
Output SQLite:
```text
"u"."string" LIKE ?
```

### NotBetween
Checks if the left expression falls outside the range defined by `valueStart` and `valueEnd`.
```go
comparison := uast.NotBetween(uast.Field[int]("u", "number"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` NOT BETWEEN ? AND ?
```
Output MsSQL:
```text
[u].[number] NOT BETWEEN @p1 AND @p2
```
Output MySQL:
```text
`u`.`number` NOT BETWEEN ? AND ?
```
Output PostgreSQL:
```text
"u"."number" NOT BETWEEN $1 AND $2
```
Output SQLite:
```text
"u"."number" NOT BETWEEN ? AND ?
```

### NotEqual
Compares two expressions for inequality (`!=` or `<>`).
```go
comparison := uast.NotEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` != ?
```
Output MsSQL:
```text
[u].[number] != @p1
```
Output MySQL:
```text
`u`.`number` != ?
```
Output PostgreSQL:
```text
"u"."number" != $1
```
Output SQLite:
```text
"u"."number" != ?
```

### NotExists
Checks if the subquery returns no rows. Returns `true` if the subquery result is empty.
```go
comparison := uast.NotExists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("users").As("u"))))
```
Output MariaDB:
```text
NOT EXISTS (SELECT 1 FROM `users` AS `u`)
```
Output MsSQL:
```text
NOT EXISTS (SELECT 1 FROM [users] AS [u])
```
Output MySQL:
```text
NOT EXISTS (SELECT 1 FROM `users` AS `u`)
```
Output PostgreSQL:
```text
NOT EXISTS (SELECT 1 FROM "users" AS "u")
```
Output SQLite:
```text
NOT EXISTS (SELECT 1 FROM "users" AS "u")
```

### NotILike
Performs a negated case-insensitive pattern matching comparison.
```go
comparison := uast.NotILike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
LOWER(`u`.`string`) NOT LIKE LOWER(?)
```
Output MsSQL:
```text
LOWER([u].[string]) NOT LIKE LOWER(@p1)
```
Output MySQL:
```text
LOWER(`u`.`string`) NOT LIKE LOWER(?)
```
Output PostgreSQL:
```text
"u"."string" NOT ILIKE $1
```
Output SQLite:
```text
LOWER("u"."string") NOT LIKE LOWER(?)
```

### NotIn
Checks if the left expression does not match any value contained within the right expression.
```go
comparison := uast.NotIn(uast.Field[string]("u", "string"), uast.Array("active", "pending"))
```
Output MariaDB:
```text
`u`.`string` NOT IN (?, ?)
```
Output MsSQL:
```text
[u].[string] NOT IN (@p1, @p2)
```
Output MySQL:
```text
`u`.`string` NOT IN (?, ?)
```
Output PostgreSQL:
```text
"u"."string" NOT IN ($1, $2)
```
Output SQLite:
```text
"u"."string" NOT IN (?, ?)
```

### NotLike
Performs a negated case-sensitive pattern matching comparison.
```go
comparison := uast.NotLike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
`u`.`string` NOT LIKE ?
```
Output MsSQL:
```text
[u].[string] NOT LIKE @p1
```
Output MySQL:
```text
`u`.`string` NOT LIKE ?
```
Output PostgreSQL:
```text
"u"."string" NOT LIKE $1
```
Output SQLite:
```text
"u"."string" NOT LIKE ?
```

## exprConstant
### ConstBoolFalse
Returns a constant boolean `FALSE` expression.
```go
constant := uast.ConstBoolFalse()
```
Output:
```text
FALSE
```

### ConstBoolTrue
Returns a constant boolean `TRUE` expression.
```go
constant := uast.ConstBoolTrue()
```
Output:
```text
TRUE
```

### ConstFloat32One
Returns a constant `float32` value of `1.0`. 
```go
constant := uast.ConstFloat32One()
```
Output:
```text
1.0
```

### ConstFloat64One
Returns a constant `float64` value of `1.000000`.
```go
constant := uast.ConstFloat64One()
```
Output:
```text
1.000000
```

### ConstIntOne
Returns a constant `int` value of `1`.
```go
constant := uast.ConstIntOne()
```
Output:
```text
1
```

### ConstInt8One
Returns a constant `int8` value of `1`.
```go
constant := uast.ConstInt8One()
```
Output:
```text
1
```

### ConstInt16One
Returns a constant `int16` value of `1`.
```go
constant := uast.ConstInt16One()
```
Output:
```text
1
```

### ConstInt32One
Returns a constant `int32` value of `1`.
```go
constant := uast.ConstInt32One()
```
Output:
```text
1
```

### ConstInt64One
Returns a constant `int64` value of `1`.
```go
constant := uast.ConstInt64One()
```
Output:
```text
1
```

### ConstStringDefault
Returns a constant `string` value of `DEFAULT`.
```go
constant := uast.ConstStringDefault()
```
Output:
```text
DEFAULT
```

### ConstStringNull
Returns a constant `string` value of `NULL`.
```go
constant := uast.ConstStringNull()
```
Output:
```text
NULL
```

### ConstUintOne
Returns a constant `uint` value of `1`.
```go
constant := uast.ConstUintOne()
```
Output:
```text
1
```

### ConstUint8One
Returns a constant `uint8` value of `1`.
```go
constant := uast.ConstUint8One()
```
Output:
```text
1
```

### ConstUint16One
Returns a constant `uint16` value of `1`.
```go
constant := uast.ConstUint16One()
```
Output:
```text
1
```

### ConstUint32One
Returns a constant `uint32` value of `1`.
```go
constant := uast.ConstUint32One()
```
Output:
```text
1
```

### ConstUint64One
Returns a constant `uint64` value of `1`.
```go
constant := uast.ConstUint64One()
```
Output:
```text
1
```

## exprField
### Field
Creates a reference to a table column, optionally qualified with a table alias. This is the primary way to reference database columns in expressions.
```go
field := uast.Field[string]("u", "string")
```
Output MariaDB:
```text
`u`.`string`
```
Output MsSQL:
```text
[u].[string]
```
Output MySQL:
```text
`u`.`string`
```
Output PostgreSQL:
```text
"u"."string"
```
Output SQLite:
```text
"u"."string"
```

## exprFunction
### Aggregate
#### Avg
Returns the average (arithmetic mean) of all non-NULL values in the expression. If `distinct` is `true`, the average is calculated over distinct values only.
```go
function := uast.Avg(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Avg(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
AVG(`u`.`number`)
AVG(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
AVG([u].[number])
AVG(DISTINCT [u].[number])
```
Output MySQL:
```text
AVG(`u`.`number`)
AVG(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
AVG("u"."number")
AVG(DISTINCT "u"."number")
```
Output SQLite:
```text
AVG("u"."number")
AVG(DISTINCT "u"."number")
```

#### BitAnd
Returns the bitwise AND of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitAnd(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitAnd(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
BIT_AND(`u`.`number`)
BIT_AND(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
BIT_AND([u].[number])
BIT_AND(DISTINCT [u].[number])
```
Output MySQL:
```text
BIT_AND(`u`.`number`)
BIT_AND(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
BIT_AND("u"."number")
BIT_AND(DISTINCT "u"."number")
```
Output SQLite:
```text
BIT_AND("u"."number")
BIT_AND(DISTINCT "u"."number")
```

#### BitOr
Returns the bitwise OR of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitOr(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitOr(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
BIT_OR(`u`.`number`)
BIT_OR(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
BIT_OR([u].[number])
BIT_OR(DISTINCT [u].[number])
```
Output MySQL:
```text
BIT_OR(`u`.`number`)
BIT_OR(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
BIT_OR("u"."number")
BIT_OR(DISTINCT "u"."number")
```
Output SQLite:
```text
BIT_OR("u"."number")
BIT_OR(DISTINCT "u"."number")
```

#### BitXor
Returns the bitwise XOR of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitXor(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitXor(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
BIT_XOR(`u`.`number`)
BIT_XOR(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
BIT_XOR([u].[number])
BIT_XOR(DISTINCT [u].[number])
```
Output MySQL:
```text
BIT_XOR(`u`.`number`)
BIT_XOR(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
BIT_XOR("u"."number")
BIT_XOR(DISTINCT "u"."number")
```
Output SQLite:
```text
BIT_XOR("u"."number")
BIT_XOR(DISTINCT "u"."number")
```

#### Count
Returns the number of rows matching the query, or the number of non-NULL values if an expression is provided. When `distinct` is `true`, counts only distinct values.
```go
function := uast.Count(uast.Field[string]("u", "string"), false)
functionWithDistinct := uast.Count(uast.Field[string]("u", "string"), true)
```
Output MariaDB:
```text
COUNT(`u`.`string`)
COUNT(DISTINCT `u`.`string`)
```
Output MsSQL:
```text
COUNT([u].[string])
COUNT(DISTINCT [u].[string])
```
Output MySQL:
```text
COUNT(`u`.`string`)
COUNT(DISTINCT `u`.`string`)
```
Output PostgreSQL:
```text
COUNT("u"."string")
COUNT(DISTINCT "u"."string")
```
Output SQLite:
```text
COUNT("u"."string")
COUNT(DISTINCT "u"."string")
```

#### GroupConcat
Concatenates values from a group into a single string, separated by a default delimiter (typically a comma). The `distinct` flag removes duplicates before concatenation.
```go
function := uast.GroupConcat(uast.Field[string]("u", "string"), false)
functionWithDistinct := uast.GroupConcat(uast.Field[string]("u", "string"), true)
```
Output MariaDB:
```text
GROUP_CONCAT(`u`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `u`.`string` SEPARATOR ',')
```
Output MsSQL:
```text
GROUP_CONCAT([u].[string], ',')
GROUP_CONCAT(DISTINCT [u].[string], ',')
```
Output MySQL:
```text
GROUP_CONCAT(`u`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `u`.`string` SEPARATOR ',')
```
Output PostgreSQL:
```text
STRING_AGG("u"."string", ',')
STRING_AGG(DISTINCT "u"."string", ',')
```
Output SQLite:
```text
GROUP_CONCAT("u"."string" SEPARATOR ',')
GROUP_CONCAT(DISTINCT "u"."string" SEPARATOR ',')
```

#### Max
Returns the maximum value of the expression across all rows in the group.
```go
function := uast.Max(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Max(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
MAX(`u`.`number`)
MAX(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
MAX([u].[number])
MAX(DISTINCT [u].[number])
```
Output MySQL:
```text
MAX(`u`.`number`)
MAX(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
MAX("u"."number")
MAX(DISTINCT "u"."number")
```
Output SQLite:
```text
MAX("u"."number")
MAX(DISTINCT "u"."number")
```

#### Min
Returns the minimum value of the expression across all rows in the group.
```go
function := uast.Min(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Min(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
MIN(`u`.`number`)
MIN(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
MIN([u].[number])
MIN(DISTINCT [u].[number])
```
Output MySQL:
```text
MIN(`u`.`number`)
MIN(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
MIN("u"."number")
MIN(DISTINCT "u"."number")
```
Output SQLite:
```text
MIN("u"."number")
MIN(DISTINCT "u"."number")
```

#### StdDev
Returns the population standard deviation of the expression.
```go
function := uast.StdDev(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.StdDev(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
STDDEV(`u`.`number`)
STDDEV(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
STDEV([u].[number])
STDEV(DISTINCT [u].[number])
```
Output MySQL:
```text
STDDEV(`u`.`number`)
STDDEV(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
STDDEV_SAMP("u"."number")
STDDEV_SAMP(DISTINCT "u"."number")
```
Output SQLite:
```text
STDEV("u"."number")
STDEV(DISTINCT "u"."number")
```

#### Sum
Returns the sum of all values in the expression. If `distinct` is `true`, sums only distinct values.
```go
function := uast.Sum(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Sum(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
SUM(`u`.`number`)
SUM(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
SUM([u].[number])
SUM(DISTINCT [u].[number])
```
Output MySQL:
```text
SUM(`u`.`number`)
SUM(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
SUM("u"."number")
SUM(DISTINCT "u"."number")
```
Output SQLite:
```text
SUM("u"."number")
SUM(DISTINCT "u"."number")
```

#### Variance
Returns the population variance of the expression.
```go
function := uast.Variance(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Variance(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
VARIANCE(`u`.`number`)
VARIANCE(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
VAR([u].[number])
VAR(DISTINCT [u].[number])
```
Output MySQL:
```text
VARIANCE(`u`.`number`)
VARIANCE(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
VAR_SAMP("u"."number")
VAR_SAMP(DISTINCT "u"."number")
```
Output SQLite:
```text
VARIANCE("u"."number")
VARIANCE(DISTINCT "u"."number")
```

### Analytical
#### FirstValue
Returns the value of the expression from the first row of the window frame. Requires an `OVER` clause with window specification.
```go
function := uast.FirstValue(uast.Field[string]("u", "string")).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
FIRST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
FIRST_VALUE([u].[string]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
FIRST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
FIRST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
FIRST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### Lag
Returns the value of the expression from a row that is `offset` rows before the current row within the partition.
```go
function := uast.Lag(uast.Field[int]("u", "number"), 2).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Asc(uast.Field[time.Time]("u", "date"))),
)
```
Output MariaDB:
```text
LAG(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
Output MsSQL:
```text
LAG([u].[number], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[date] ASC)
```
Output MySQL:
```text
LAG(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
Output PostgreSQL:
```text
LAG("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```
Output SQLite:
```text
LAG("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```

#### LastValue
Returns the value of the expression from the last row of the window frame.
```go
function := uast.LastValue(uast.Field[string]("u", "string")).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Asc(uast.Field[int]("u", "number"))),
    uast.RowsBetween("CURRENT ROW", "UNBOUNDED FOLLOWING"),
)
```
Output MariaDB:
```text
LAST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output MsSQL:
```text
LAST_VALUE([u].[string]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output MySQL:
```text
LAST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output PostgreSQL:
```text
LAST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output SQLite:
```text
LAST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```

#### Lead
Returns the value of the expression from a row that is `offset` rows after the current row within the partition.
```go
function := uast.Lead(uast.Field[int]("u", "number"), 2).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Asc(uast.Field[time.Time]("u", "date"))),
)
```
Output MariaDB:
```text
LEAD(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
Output MsSQL:
```text
LEAD([u].[number], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[date] ASC)
```
Output MySQL:
```text
LEAD(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
Output PostgreSQL:
```text
LEAD("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```
Output SQLite:
```text
LEAD("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```

#### NthValue
Returns the value of the expression from the `n-th` row of the window frame.
```go
function := uast.NthValue(uast.Field[string]("u", "string"), 2).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
    uast.RowsBetween("UNBOUNDED PRECEDING", "CURRENT ROW"),
)
```
Output MariaDB:
```text
NTH_VALUE(`u`.`string`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output MsSQL:
```text
NTH_VALUE([u].[string], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output MySQL:
```text
NTH_VALUE(`u`.`string`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output PostgreSQL:
```text
NTH_VALUE("u"."string", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output SQLite:
```text
NTH_VALUE("u"."string", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```

### Condition
#### Case
Evaluates a list of `WHEN`-`THEN` pairs and returns the `THEN` expression for the first true WHEN. If no condition is true, returns the `ELSE` expression if provided, or `NULL`.
```go
pairs := uast.CaseIf(
    uast.CasePair(
        uast.Less(uast.Field[int]("u", "number"), uast.Value(2)),
        uast.Value("old"),
    ),
)
elseExpr := uast.CaseElse(uast.Value("new"))
function := uast.Case(pairs, elseExpr)
```
Output MariaDB:
```text
CASE WHEN `u`.`number` < ? THEN ? ELSE ? END
```
Output MsSQL:
```text
CASE WHEN [u].[number] < @p1 THEN @p2 ELSE @p3 END
```
Output MySQL:
```text
CASE WHEN `u`.`number` < ? THEN ? ELSE ? END
```
Output PostgreSQL:
```text
CASE WHEN "u"."number" < $1 THEN $2 ELSE $3 END
```
Output SQLite:
```text
CASE WHEN "u"."number" < ? THEN ? ELSE ? END
```

#### Coalesce
Returns the first non-NULL expression from the provided list. Useful for providing fallback values.
```go
function := uast.Coalesce(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
Output MariaDB:
```text
COALESCE(`u`.`createat`, `u`.`updateat`)
```
Output MsSQL:
```text
COALESCE([u].[createat], [u].[updateat])
```
Output MySQL:
```text
COALESCE(`u`.`createat`, `u`.`updateat`)
```
Output PostgreSQL:
```text
COALESCE("u"."createat", "u"."updateat")
```
Output SQLite:
```text
COALESCE("u"."createat", "u"."updateat")
```

#### Greatest
Returns the largest value from the provided list of expressions.
```go
function := uast.Greatest(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
Output MariaDB:
```text
GREATEST(`u`.`createat`, `u`.`updateat`)
```
Output MsSQL:
```text
GREATEST([u].[createat], [u].[updateat])
```
Output MySQL:
```text
GREATEST(`u`.`createat`, `u`.`updateat`)
```
Output PostgreSQL:
```text
GREATEST("u"."createat", "u"."updateat")
```
Output SQLite:
```text
GREATEST("u"."createat", "u"."updateat")
```

#### Least
Returns the smallest value from the provided list of expressions.
```go
function := uast.Least(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
Output MariaDB:
```text
LEAST(`u`.`createat`, `u`.`updateat`)
```
Output MsSQL:
```text
LEAST([u].[createat], [u].[updateat])
```
Output MySQL:
```text
LEAST(`u`.`createat`, `u`.`updateat`)
```
Output PostgreSQL:
```text
LEAST("u"."createat", "u"."updateat")
```
Output SQLite:
```text
LEAST("u"."createat", "u"."updateat")
```

#### NullIf
Returns `NULL` if the two expressions are equal; otherwise returns the first expression.
```go
function := uast.NullIf(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
Output MariaDB:
```text
NULLIF(`u`.`createat`, `u`.`updateat`)
```
Output MsSQL:
```text
NULLIF([u].[createat], [u].[updateat])
```
Output MySQL:
```text
NULLIF(`u`.`createat`, `u`.`updateat`)
```
Output PostgreSQL:
```text
NULLIF("u"."createat", "u"."updateat")
```
Output SQLite:
```text
NULLIF("u"."createat", "u"."updateat")
```

### Convert
#### Cast
Converts an expression to a specified data type.
```go
function := uast.Cast(uast.Field[int]("u", "number"), uast.TypeString)
```
Output MariaDB:
```text
CAST(`u`.`number` AS CHAR)
```
Output MsSQL:
```text
CAST([u].[number] AS NVARCHAR)
```
Output MySQL:
```text
CAST(`u`.`number` AS CHAR)
```
Output PostgreSQL:
```text
CAST("u"."number" AS VARCHAR)
```
Output SQLite:
```text
CAST("u"."number" AS TEXT)
```

#### CharLength
Returns the number of characters in a string expression.
```go
function := uast.CharLength(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
CHAR_LENGTH(`u`.`string`)
```
Output MsSQL:
```text
CHAR_LENGTH([u].[string])
```
Output MySQL:
```text
CHAR_LENGTH(`u`.`string`)
```
Output PostgreSQL:
```text
CHAR_LENGTH("u"."string")
```
Output SQLite:
```text
CHAR_LENGTH("u"."string")
```

#### DateFormat
Formats a datetime expression according to a specified format mask.
```go
function := uast.DateFormat(uast.Field[time.Time]("u", "createat"), uast.Value("%Y-%m-%d"))
```
Output MariaDB:
```text
DATE_FORMAT(`u`.`createat`, '%Y-%m-%d')
```
Output MsSQL:
```text
FORMAT([u].[createat], '%Y-%m-%d')
```
Output MySQL:
```text
DATE_FORMAT(`u`.`createat`, '%Y-%m-%d')
```
Output PostgreSQL:
```text
TO_CHAR("u"."createat", '%Y-%m-%d')
```
Output SQLite:
```text
STRFTIME("u"."createat", '%Y-%m-%d')
```

#### Degrees
Converts an angle from radians to degrees.
```go
function := uast.Degrees(uast.Field[int]("u", "number"))
```
Output MariaDB:
```text
DEGREES(`u`.`number`)
```
Output MsSQL:
```text
DEGREES([u].[number])
```
Output MySQL:
```text
DEGREES(`u`.`number`)
```
Output PostgreSQL:
```text
DEGREES("u"."number")
```
Output SQLite:
```text
DEGREES("u"."number")
```

#### Length
Returns the byte length of a string expression.
```go
function := uast.Length(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
LENGTH(`u`.`string`)
```
Output MsSQL:
```text
LEN([u].[string])
```
Output MySQL:
```text
LENGTH(`u`.`string`)
```
Output PostgreSQL:
```text
LENGTH("u"."string")
```
Output SQLite:
```text
LENGTH("u"."string")
```

#### Position
Returns the starting position of the first occurrence of a substring within a string.
```go
function := uast.Position(uast.Field[string]("u", "string"), uast.Value("old"))
```
Output MariaDB:
```text
POSITION(? IN `u`.`string`)
```
Output MsSQL:
```text
CHARINDEX(@p1, [u].[string])
```
Output MySQL:
```text
POSITION(? IN `u`.`string`)
```
Output PostgreSQL:
```text
POSITION($1 IN "u"."string")
```
Output SQLite:
```text
POSITION(? IN "u"."string")
```

#### Radians
Converts an angle from degrees to radians.
```go
function := uast.Radians(uast.Field[int]("u", "number"))
```
Output MariaDB:
```text
RADIANS(`u`.`number`)
```
Output MsSQL:
```text
RADIANS([u].[number])
```
Output MySQL:
```text
RADIANS(`u`.`number`)
```
Output PostgreSQL:
```text
RADIANS("u"."number")
```
Output SQLite:
```text
RADIANS("u"."number")
```

### Date and time
#### CurDate
Returns the current date (without time).
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
Returns the current time (without date).
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
Adds a time/date interval to a datetime expression and returns the resulting datetime.
```go
function := uast.DateAdd(uast.Field[time.Time]("u", "createat"), uast.Value("2 DAY"))
```
Output MariaDB:
```text
DATE_ADD(`u`.`createat`, INTERVAL 2 DAY)
```
Output MsSQL:
```text
DATEADD(DAY, 2, [u].[createat])
```
Output MySQL:
```text
DATE_ADD(`u`.`createat`, INTERVAL 2 DAY)
```
Output PostgreSQL:
```text
("u"."createat" + INTERVAL '2 DAY')
```
Output SQLite:
```text
DATETIME("u"."createat", '+2 DAY')
```

#### DateDiff
Returns the difference in days between two datetime expressions (`datetimeEnd` - `datetimeStart`).
```go
function := uast.DateDiff(uast.Field[time.Time]("u", "updateat"), uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
DATEDIFF(`u`.`updateat`, `u`.`createat`)
```
Output MsSQL:
```text
DATEDIFF([u].[updateat], [u].[createat])
```
Output MySQL:
```text
DATEDIFF(`u`.`updateat`, `u`.`createat`)
```
Output PostgreSQL:
```text
DATE_PART('day', "u"."updateat" - "u"."createat")
```
Output SQLite:
```text
DATEDIFF("u"."updateat", "u"."createat")
```

#### DateSub
Subtracts a time/date interval from a datetime expression and returns the resulting datetime.
```go
function := uast.DateSub(uast.Field[time.Time]("u", "createat"), uast.Value("2 DAY"))
```
Output MariaDB:
```text
DATE_SUB(`u`.`createat`, INTERVAL 2 DAY)
```
Output MsSQL:
```text
DATEADD(DAY, -2, [u].[createat])
```
Output MySQL:
```text
DATE_SUB(`u`.`createat`, INTERVAL 2 DAY)
```
Output PostgreSQL:
```text
("u"."createat" - INTERVAL '2 DAY')
```
Output SQLite:
```text
DATETIME("u"."createat", '-2 DAY')
```

#### Day
Extracts the day of the month (1–31) from a datetime expression.
```go
function := uast.Day(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
DAY(`u`.`createat`)
```
Output MsSQL:
```text
DAY([u].[createat])
```
Output MySQL:
```text
DAY(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(DAY FROM "u"."createat")
```
Output SQLite:
```text
DAY("u"."createat")
```

#### DayName
Returns the name of the weekday (e.g., 'Monday', 'Tuesday') for a given datetime expression.
```go
function := uast.DayName(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
DAYNAME(`u`.`createat`)
```
Output MsSQL:
```text
DATENAME(WEEKDAY, [u].[createat])
```
Output MySQL:
```text
DAYNAME(`u`.`createat`)
```
Output PostgreSQL:
```text
TO_CHAR("u"."createat", 'Day')
```
Output SQLite:
```text
STRFTIME('%w', "u"."createat")
```

#### Hour
Extracts the hour (0–23) from a datetime expression.
```go
function := uast.Hour(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
HOUR(`u`.`createat`)
```
Output MsSQL:
```text
DATEPART(HOUR, [u].[createat])
```
Output MySQL:
```text
HOUR(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(HOUR FROM "u"."createat")
```
Output SQLite:
```text
HOUR("u"."createat")
```

#### Minute
Extracts the minute (0–59) from a datetime expression.
```go
function := uast.Minute(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
MINUTE(`u`.`createat`)
```
Output MsSQL:
```text
DATEPART(MINUTE, [u].[createat])
```
Output MySQL:
```text
MINUTE(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(MINUTE FROM "u"."createat")
```
Output SQLite:
```text
MINUTE("u"."createat")
```

#### Month
Extracts the month (1–12) from a datetime expression.
```go
function := uast.Month(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
MONTH(`u`.`createat`)
```
Output MsSQL:
```text
MONTH([u].[createat])
```
Output MySQL:
```text
MONTH(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(MONTH FROM "u"."createat")
```
Output SQLite:
```text
MONTH("u"."createat")
```

#### MonthName
Returns the name of the month (e.g., 'January', 'February') for a given datetime expression.
```go
function := uast.MonthName(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
MONTHNAME(`u`.`createat`)
```
Output MsSQL:
```text
DATENAME(MONTH, [u].[createat])
```
Output MySQL:
```text
MONTHNAME(`u`.`createat`)
```
Output PostgreSQL:
```text
TO_CHAR("u"."createat", 'Month')
```
Output SQLite:
```text
STRFTIME('%m', "u"."createat")
```

#### Now
Returns the current date and time.
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
Extracts the quarter (1–4) from a datetime expression.
```go
function := uast.Quarter(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
QUARTER(`u`.`createat`)
```
Output MsSQL:
```text
DATEPART(QUARTER, [u].[createat])
```
Output MySQL:
```text
QUARTER(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(QUARTER FROM "u"."createat")
```
Output SQLite:
```text
QUARTER("u"."createat")
```

#### Second
Extracts the second (0–59) from a datetime expression.
```go
function := uast.Second(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
SECOND(`u`.`createat`)
```
Output MsSQL:
```text
DATEPART(SECOND, [u].[createat])
```
Output MySQL:
```text
SECOND(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(SECOND FROM "u"."createat")
```
Output SQLite:
```text
SECOND("u"."createat")
```

#### TimeAdd
Adds a time interval to a time/datetime expression and returns the resulting time.
```go
function := uast.TimeAdd(uast.Field[time.Time]("u", "createat"), uast.Value("2 HOUR"))
```
Output MariaDB:
```text
TIME_ADD(`u`.`createat`, '2 HOUR')
```
Output MsSQL:
```text
DATEADD(HOUR, 2, [u].[createat])
```
Output MySQL:
```text
TIME_ADD(`u`.`createat`, '2 HOUR')
```
Output PostgreSQL:
```text
("u"."createat" + INTERVAL '2 HOUR')
```
Output SQLite:
```text
TIME("u"."createat", '+2 HOUR')
```

#### TimeDiff
Returns the difference between two time/datetime expressions (`timeEnd` - `timeStart`).
```go
function := uast.TimeDiff(uast.Field[time.Time]("u", "updateat"), uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
TIMEDIFF(`u`.`updateat`, `u`.`createat`)
```
Output MsSQL:
```text
TIMEDIFF([u].[updateat], [u].[createat])
```
Output MySQL:
```text
TIMEDIFF(`u`.`updateat`, `u`.`createat`)
```
Output PostgreSQL:
```text
DATE_PART('time', "u"."updateat" - "u"."createat")
```
Output SQLite:
```text
TIMEDIFF("u"."updateat", "u"."createat")
```

#### TimeSub
Subtracts a time interval from a time/datetime expression and returns the resulting time.
```go
function := uast.TimeSub(uast.Field[time.Time]("u", "createat"), uast.Value("2 HOUR"))
```
Output MariaDB:
```text
TIME_SUB(`u`.`createat`, '2 HOUR')
```
Output MsSQL:
```text
DATEADD(HOUR, -2, [u].[createat])
```
Output MySQL:
```text
TIME_SUB(`u`.`createat`, '2 HOUR')
```
Output PostgreSQL:
```text
("u"."createat" - INTERVAL '2 HOUR')
```
Output SQLite:
```text
TIME("u"."createat", '-2 HOUR')
```

#### Week
Extracts the week number (1–53) from a datetime expression.
```go
function := uast.Week(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
WEEK(`u`.`createat`)
```
Output MsSQL:
```text
DATEPART(WEEK, [u].[createat])
```
Output MySQL:
```text
WEEK(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(WEEK FROM "u"."createat")
```
Output SQLite:
```text
WEEK("u"."createat")
```

#### Year
Extracts the year from a datetime expression.
```go
function := uast.Year(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
YEAR(`u`.`createat`)
```
Output MsSQL:
```text
YEAR([u].[createat])
```
Output MySQL:
```text
YEAR(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(YEAR FROM "u"."createat")
```
Output SQLite:
```text
YEAR("u"."createat")
```

### Json
#### JsonArray
Creates a JSON array from the given expression and optional additional values.
```go
function := uast.JsonArray(
    uast.Field[string]("u", "json"), 
    uast.Value("val1"), 
    uast.Value("val2"),
)
```
Output MariaDB:
```text
JSON_ARRAY(`u`.`json`, ?, ?)
```
Output MsSQL:
```text
JSON_ARRAY([u].[json], @p1, @p2)
```
Output MySQL:
```text
JSON_ARRAY(`u`.`json`, ?, ?)
```
Output PostgreSQL:
```text
JSON_ARRAY("u"."json", $1, $2)
```
Output SQLite:
```text
JSON_ARRAY("u"."json", ?, ?)
```

#### JsonArrayAgg
Aggregates values from a group into a JSON array.
```go
function := uast.JsonArrayAgg(
    uast.Field[string]("u", "json"),
)
```
Output MariaDB:
```text
JSON_ARRAYAGG(`u`.`json`)
```
Output MsSQL:
```text
JSON_ARRAYAGG([u].[json])
```
Output MySQL:
```text
JSON_ARRAYAGG(`u`.`json`)
```
Output PostgreSQL:
```text
JSON_AGG("u"."json")
```
Output SQLite:
```text
JSON_GROUP_ARRAY("u"."json")
```

#### JsonContains
Checks whether a JSON document contains a specified value.
```go
function := uast.JsonContains(
    uast.Field[string]("u", "json"),
    uast.Value(`{"key":"val"}`),
)
```
Output MariaDB:
```text
JSON_CONTAINS(`u`.`json`, '{"key":"val"}')
```
Output MsSQL:
```text
// Not supported
```
Output MySQL:
```text
JSON_CONTAINS(`u`.`json`, '{"key":"val"}')
```
Output PostgreSQL:
```text
("u"."json" @> '{"key":"val"}')
```
Output SQLite:
```text
JSON_CONTAINS("u"."json", '{"key":"val"}')
```

#### JsonExtract
Extracts a value from a JSON document at the specified path. The `json` parameter is built with `JsonPath` and optional `JsonKey`/`JsonIndex`.
```go
function := JsonExtract(
    uast.Field[string]("u", "json"), 
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
(`u`.`json` ->> '$.parent[0].child')
```
Output MsSQL:
```text
JSON_VALUE([u].[json], '$.parent[0].child')
```
Output MySQL:
```text
(`u`.`json` ->> '$.parent[0].child')
```
Output PostgreSQL:
```text
("u"."json" #>> '{parent,0,child}')
```
Output SQLite:
```text
("u"."json" ->> '$.parent[0].child')
```

#### JsonObject
Builds a JSON object from key-value pairs.
```go
function := uast.JsonObject(
    uast.JsonPair(
        uast.JsonKey("key"), 
        uast.Count(uast.Field[string]("u", "json"), false),
    ),
)
```
Output MariaDB:
```text
JSON_OBJECT('key', COUNT(`u`.`json`))
```
Output MsSQL:
```text
JSON_OBJECT('key', COUNT([u].[json]))
```
Output MySQL:
```text
JSON_OBJECT('key', COUNT(`u`.`json`))
```
Output PostgreSQL:
```text
JSON_BUILD_OBJECT('key', COUNT("u"."json"))
```
Output SQLite:
```text
JSON_OBJECT('key', COUNT("u"."json"))
```

#### JsonObjectAgg
Aggregates key-value pairs from a group into a single JSON object.
```go
function := uast.JsonObjectAgg(
    uast.Field[string]("u", "json"),
    uast.Field[int]("u", "number"),
)
```
Output MariaDB:
```text
JSON_OBJECTAGG(`u`.`json`, `u`.`number`)
```
Output MsSQL:
```text
JSON_OBJECTAGG([u].[json], [u].[number])
```
Output MySQL:
```text
JSON_OBJECTAGG(`u`.`json`, `u`.`number`)
```
Output PostgreSQL:
```text
JSON_OBJECT_AGG("u"."json", "u"."number")
```
Output SQLite:
```text
JSON_GROUP_OBJECT("u"."json", "u"."number")
```

#### JsonRemove
Removes a value from a JSON document at the specified path(s).
```go
function := uast.JsonRemove(
    uast.Field[string]("u", "json"),
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
JSON_REMOVE(`u`.`json`, '$.key1', '$.key2')
```
Output MsSQL:
```text
JSON_MODIFY(JSON_MODIFY([u].[json], '$.key1', NULL), '$.key2', NULL)
```
Output MySQL:
```text
JSON_REMOVE(`u`.`json`, '$.key1', '$.key2')
```
Output PostgreSQL:
```text
("u"."json" - '{key1}' - '{key2}')
```
Output SQLite:
```text
JSON_REMOVE("u"."json", '$.key1', '$.key2')
```

#### JsonSet
Sets a value in a JSON document at the specified path(s). Creates the path if it does not exist.
```go
function := uast.JsonSet(
    uast.Field[string]("u", "json"),
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
JSON_SET(`u`.`json`, '$.key1', ?, '$.key2', ?)
```
Output MsSQL:
```text
JSON_MODIFY(JSON_MODIFY([u].[json], '$.key1', @p1), '$.key2', @p2)
```
Output MySQL:
```text
JSON_SET(`u`.`json`, '$.key1', ?, '$.key2', ?)
```
Output PostgreSQL:
```text
jsonb_set(jsonb_set("u"."json", '{key1}', $1), '{key2}', $2)
```
Output SQLite:
```text
JSON_SET("u"."json", '$.key1', ?, '$.key2', ?)
```

#### JsonType
Returns the JSON type of a JSON value (e.g., 'OBJECT', 'ARRAY', 'STRING', 'INTEGER', 'NULL').
```go
function := uast.JsonType(uast.Field[string]("u", "json"))
```
Output MariaDB:
```text
JSON_TYPE(`u`.`json`)
```
Output MsSQL:
```text
// Not supported
```
Output MySQL:
```text
JSON_TYPE(`u`.`json`)
```
Output PostgreSQL:
```text
jsonb_typeof("u"."json")
```
Output SQLite:
```text
JSON_TYPE("u"."json")
```

### Math
#### Abs
Returns the absolute (non-negative) value of a numeric expression.
```go
function := uast.Abs(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
ABS(`u`.`x`)
```
Output MsSQL:
```text
ABS([u].[x])
```
Output MySQL:
```text
ABS(`u`.`x`)
```
Output PostgreSQL:
```text
ABS("u"."x")
```
Output SQLite:
```text
ABS("u"."x")
```

#### ACos
Returns the arc cosine (inverse cosine) of the expression, in radians.
```go
function := uast.ACos(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
ACOS(`u`.`x`)
```
Output MsSQL:
```text
ACOS([u].[x])
```
Output MySQL:
```text
ACOS(`u`.`x`)
```
Output PostgreSQL:
```text
ACOS("u"."x")
```
Output SQLite:
```text
ACOS("u"."x")
```

#### ASin
Returns the arc sine (inverse sine) of the expression, in radians.
```go
function := uast.ASin(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
ASIN(`u`.`x`)
```
Output MsSQL:
```text
ASIN([u].[x])
```
Output MySQL:
```text
ASIN(`u`.`x`)
```
Output PostgreSQL:
```text
ASIN("u"."x")
```
Output SQLite:
```text
ASIN("u"."x")
```

#### ATan
Returns the arc tangent (inverse tangent) of the expression, in radians.
```go
function := uast.ATan(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
ATAN(`u`.`x`)
```
Output MsSQL:
```text
ATAN([u].[x])
```
Output MySQL:
```text
ATAN(`u`.`x`)
```
Output PostgreSQL:
```text
ATAN("u"."x")
```
Output SQLite:
```text
ATAN("u"."x")
```

#### ATan2
Returns the arc tangent of the quotient of its two arguments (`y`/`x`), using their signs to determine the quadrant.
```go
function := uast.ATan2(uast.Field[int]("u", "y"), uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
ATAN2(`u`.`y`, `u`.`x`)
```
Output MsSQL:
```text
ATAN2([u].[y], [u].[x])
```
Output MySQL:
```text
ATAN2(`u`.`y`, `u`.`x`)
```
Output PostgreSQL:
```text
ATAN2("u"."y", "u"."x")
```
Output SQLite:
```text
ATAN2("u"."y", "u"."x")
```

#### Cbrt
Returns the cube root of a numeric expression.
```go
function := uast.Cbrt(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
CBRT(`u`.`x`)
```
Output MsSQL:
```text
CBRT([u].[x])
```
Output MySQL:
```text
CBRT(`u`.`x`)
```
Output PostgreSQL:
```text
CBRT("u"."x")
```
Output SQLite:
```text
CBRT("u"."x")
```

#### Ceil
Returns the smallest integer value not less than the argument (rounds up).
```go
function := uast.Ceil(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
CEILING(`u`.`x`)
```
Output MsSQL:
```text
CEILING([u].[x])
```
Output MySQL:
```text
CEILING(`u`.`x`)
```
Output PostgreSQL:
```text
CEIL("u"."x")
```
Output SQLite:
```text
CEIL("u"."x")
```

#### Cos
Returns the cosine of the expression, where the expression is in radians.
```go
function := uast.Cos(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
COS(`u`.`x`)
```
Output MsSQL:
```text
COS([u].[x])
```
Output MySQL:
```text
COS(`u`.`x`)
```
Output PostgreSQL:
```text
COS("u"."x")
```
Output SQLite:
```text
COS("u"."x")
```

#### Exp
Returns `e` (Euler's number, ~2.71828) raised to the power of the expression.
```go
function := uast.Exp(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
EXP(`u`.`x`)
```
Output MsSQL:
```text
EXP([u].[x])
```
Output MySQL:
```text
EXP(`u`.`x`)
```
Output PostgreSQL:
```text
EXP("u"."x")
```
Output SQLite:
```text
EXP("u"."x")
```

#### Floor
Returns the largest integer value not greater than the argument (rounds down).
```go
function := uast.Floor(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
FLOOR(`u`.`x`)
```
Output MsSQL:
```text
FLOOR([u].[x])
```
Output MySQL:
```text
FLOOR(`u`.`x`)
```
Output PostgreSQL:
```text
FLOOR("u"."x")
```
Output SQLite:
```text
FLOOR("u"."x")
```

#### Ln
Returns the natural logarithm (base `e`) of the expression.
```go
function := uast.Ln(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
LN(`u`.`x`)
```
Output MsSQL:
```text
LN([u].[x])
```
Output MySQL:
```text
LN(`u`.`x`)
```
Output PostgreSQL:
```text
LN("u"."x")
```
Output SQLite:
```text
LN("u"."x")
```

#### Log
Returns the logarithm of the expression to the specified base.
```go
function := uast.Log(uast.Field[int]("u", "x"), uast.Value(2))
```
Output MariaDB:
```text
LOG(`u`.`x`, ?)
```
Output MsSQL:
```text
LOG([u].[x], @p1)
```
Output MySQL:
```text
LOG(`u`.`x`, ?)
```
Output PostgreSQL:
```text
LOG("u"."x", $1)
```
Output SQLite:
```text
LOG("u"."x", ?)
```

#### Mod
Returns the remainder (modulo) of the division of the first expression by the second.
```go
function := uast.Mod(uast.Field[int]("u", "x"), uast.Value(2))
```
Output MariaDB:
```text
MOD(`u`.`x`, ?)
```
Output MsSQL:
```text
MOD([u].[x], @p1)
```
Output MySQL:
```text
MOD(`u`.`x`, ?)
```
Output PostgreSQL:
```text
MOD("u"."x", $1)
```
Output SQLite:
```text
MOD("u"."x", ?)
```

#### Pi
Returns the mathematical constant `π` (~3.14159).
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
Returns the expression raised to the power of the exponent.
```go
function := uast.Power(uast.Field[int]("u", "x"), uast.Value(2))
```
Output MariaDB:
```text
POWER(`u`.`x`, ?)
```
Output MsSQL:
```text
POWER([u].[x], @p1)
```
Output MySQL:
```text
POWER(`u`.`x`, ?)
```
Output PostgreSQL:
```text
POWER("u"."x", $1)
```
Output SQLite:
```text
POWER("u"."x", ?)
```

#### Rand
Returns a random floating-point value in the range [0, 1].
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
Rounds the expression to the specified number of decimal places.
```go
function := uast.Round(uast.Field[int]("u", "x"), uast.Value(2))
```
Output MariaDB:
```text
ROUND(`u`.`x`, ?)
```
Output MsSQL:
```text
ROUND([u].[x], @p1)
```
Output MySQL:
```text
ROUND(`u`.`x`, ?)
```
Output PostgreSQL:
```text
ROUND("u"."x", $1)
```
Output SQLite:
```text
ROUND("u"."x", ?)
```

#### Sin
Returns the sine of the expression, where the expression is in radians.
```go
function := uast.Sin(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
SIN(`u`.`x`)
```
Output MsSQL:
```text
SIN([u].[x])
```
Output MySQL:
```text
SIN(`u`.`x`)
```
Output PostgreSQL:
```text
SIN("u"."x")
```
Output SQLite:
```text
SIN("u"."x")
```

#### Sqrt
Returns the square root of the expression.
```go
function := uast.Sqrt(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
SQRT(`u`.`x`)
```
Output MsSQL:
```text
SQRT([u].[x])
```
Output MySQL:
```text
SQRT(`u`.`x`)
```
Output PostgreSQL:
```text
SQRT("u"."x")
```
Output SQLite:
```text
SQRT("u"."x")
```

#### Tan
Returns the tangent of the expression, where the expression is in radians.
```go
function := uast.Tan(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
TAN(`u`.`x`)
```
Output MsSQL:
```text
TAN([u].[x])
```
Output MySQL:
```text
TAN(`u`.`x`)
```
Output PostgreSQL:
```text
TAN("u"."x")
```
Output SQLite:
```text
TAN("u"."x")
```

#### Trunc
Truncates the numeric expression to the specified number of decimal places (without rounding).
```go
function := uast.Trunc(uast.Field[int]("u", "x"), uast.Value(2))
```
Output MariaDB:
```text
TRUNCATE(`u`.`x`, ?)
```
Output MsSQL:
```text
ROUND([u].[x], @p1, 1)
```
Output MySQL:
```text
TRUNCATE(`u`.`x`, ?)
```
Output PostgreSQL:
```text
TRUNC("u"."x", $1)
```
Output SQLite:
```text
TRUNC("u"."x", ?)
```

### Ranking
#### CumeDist
Returns the cumulative distribution of a value within a partition (the ratio of rows that come before or are peers with the current row). Must be used with an `OVER` clause.
```go
function := uast.CumeDist().Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
CUME_DIST() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
CUME_DIST() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
CUME_DIST() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
CUME_DIST() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
CUME_DIST() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### DenseRank
Returns the rank of a row without gaps. Rows with equal values receive the same rank, and the next rank is the immediate next integer. Requires `OVER`.
```go
function := uast.DenseRank().Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
DENSE_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
DENSE_RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
DENSE_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
DENSE_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
DENSE_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### NTile
Divides the rows within a partition into `n` approximately equal groups and returns the group number (1 through `n`) for each row.
```go
function := uast.NTile(2).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
NTILE(2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
NTILE(2) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
NTILE(2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
NTILE(2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
NTILE(2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### PercentRank
Returns the percentile rank of a row within a partition (range 0 to 1). Rank of first row is always 0. Requires `OVER`.
```go
function := uast.PercentRank().Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
PERCENT_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
PERCENT_RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
PERCENT_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
PERCENT_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
PERCENT_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### Rank
Returns the rank of a row with gaps. Equal values receive the same rank, and the next distinct value skips ahead. Requires `OVER`.
```go
function := uast.Rank().Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### RowNumber
Assigns a unique sequential integer to each row within the partition, starting from 1. Order determines the numbering sequence.
```go
function := uast.RowNumber().Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
ROW_NUMBER() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
ROW_NUMBER() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
ROW_NUMBER() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
ROW_NUMBER() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
ROW_NUMBER() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

### String
#### Concat
Concatenates two or more string expressions into a single string. `NULL` arguments are treated as empty strings in most dialects.
```go
function := uast.Concat(uast.Field[string]("u", "string"), uast.Value("old"), uast.Value("new"))
```
Output MariaDB:
```text
CONCAT(`u`.`string`, ?, ?)
```
Output MsSQL:
```text
CONCAT([u].[string], @p1, @p2)
```
Output MySQL:
```text
CONCAT(`u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
CONCAT("u"."string", $1, $2)
```
Output SQLite:
```text
CONCAT("u"."string", ?, ?)
```

#### ConcatWs
Concatenates two or more string expressions with a specified separator between them. Skips `NULL` arguments.
```go
function := uast.ConcatWs(uast.Value("_"), uast.Field[string]("u", "string"), uast.Value("old"),uast.Value("new"))
```
Output MariaDB:
```text
CONCAT_WS(?, `u`.`string`, ?, ?)
```
Output MsSQL:
```text
CONCAT_WS(@p1, [u].[string], @p2, @p3)
```
Output MySQL:
```text
CONCAT_WS(?, `u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
CONCAT_WS($1, "u"."string", $2, $3)
```
Output SQLite:
```text
CONCAT_WS(?, "u"."string", ?, ?)
```

#### LeftString
Returns the leftmost `count` characters from a string expression.
```go
function := uast.LeftString(uast.Field[string]("u", "string"), uast.Value(2))
```
Output MariaDB:
```text
LEFT(`u`.`string`, ?)
```
Output MsSQL:
```text
LEFT([u].[string], @p1)
```
Output MySQL:
```text
LEFT(`u`.`string`, ?)
```
Output PostgreSQL:
```text
LEFT("u"."string", $1)
```
Output SQLite:
```text
LEFT("u"."string", ?)
```

#### Lower
Converts a string expression to lowercase.
```go
function := uast.Lower(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
LOWER(`u`.`string`)
```
Output MsSQL:
```text
LOWER([u].[string])
```
Output MySQL:
```text
LOWER(`u`.`string`)
```
Output PostgreSQL:
```text
LOWER("u"."string")
```
Output SQLite:
```text
LOWER("u"."string")
```

#### LPad
Left-pads a string expression with the specified separator to a total length of `count` characters.
```go
function := uast.LPad(uast.Field[string]("u", "string"), uast.Value(2), uast.Value(","))
```
Output MariaDB:
```text
LPAD(`u`.`string`, ?, ?)
```
Output MsSQL:
```text
LPAD([u].[string], @p1, @p2)
```
Output MySQL:
```text
LPAD(`u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
LPAD("u"."string", $1, $2)
```
Output SQLite:
```text
LPAD("u"."string", ?, ?)
```

#### LTrim
Removes leading spaces from a string expression.
```go
function := uast.LTrim(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
LTRIM(`u`.`string`)
```
Output MsSQL:
```text
LTRIM([u].[string])
```
Output MySQL:
```text
LTRIM(`u`.`string`)
```
Output PostgreSQL:
```text
LTRIM("u"."string")
```
Output SQLite:
```text
LTRIM("u"."string")
```

#### Repeat
Repeats a string expression `count` times.
```go
function := uast.Repeat(uast.Field[string]("u", "string"), uast.Value(2))
```
Output MariaDB:
```text
REPEAT(`u`.`string`, ?)
```
Output MsSQL:
```text
REPEAT([u].[string], @p1)
```
Output MySQL:
```text
REPEAT(`u`.`string`, ?)
```
Output PostgreSQL:
```text
REPEAT("u"."string", $1)
```
Output SQLite:
```text
REPEAT("u"."string", ?)
```

#### Replace
Replaces all occurrences of a substring in a string with a new substring.
```go
function := uast.Replace(uast.Field[string]("u", "string"), uast.Value("old"), uast.Value("new"))
```
Output MariaDB:
```text
REPLACE(`u`.`string`, ?, ?)
```
Output MsSQL:
```text
REPLACE([u].[string], @p1, @p2)
```
Output MySQL:
```text
REPLACE(`u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
REPLACE("u"."string", $1, $2)
```
Output SQLite:
```text
REPLACE("u"."string", ?, ?)
```

#### Reverse
Reverses the characters in a string expression.
```go
function := uast.Reverse(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
REVERSE(`u`.`string`)
```
Output MsSQL:
```text
REVERSE([u].[string])
```
Output MySQL:
```text
REVERSE(`u`.`string`)
```
Output PostgreSQL:
```text
REVERSE("u"."string")
```
Output SQLite:
```text
REVERSE("u"."string")
```

#### RightString
Returns the rightmost `count` characters from a string expression.
```go
function := uast.RightString(uast.Field[string]("u", "string"), uast.Value(2))
```
Output MariaDB:
```text
RIGHT(`u`.`string`, ?)
```
Output MsSQL:
```text
RIGHT([u].[string], @p1)
```
Output MySQL:
```text
RIGHT(`u`.`string`, ?)
```
Output PostgreSQL:
```text
RIGHT("u"."string", $1)
```
Output SQLite:
```text
RIGHT("u"."string", ?)
```

#### RPad
Right-pads a string expression with the specified separator to a total length of `count` characters.
```go
function := uast.RPad(uast.Field[string]("u", "string"), uast.Value(2), uast.Value(","))
```
Output MariaDB:
```text
RPAD(`u`.`string`, ?, ?)
```
Output MsSQL:
```text
RPAD([u].[string], @p1, @p2)
```
Output MySQL:
```text
RPAD(`u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
RPAD("u"."string", $1, $2)
```
Output SQLite:
```text
RPAD("u"."string", ?, ?)
```

#### RTrim
Removes trailing spaces from a string expression.
```go
function := uast.RTrim(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
RTRIM(`u`.`string`)
```
Output MsSQL:
```text
RTRIM([u].[string])
```
Output MySQL:
```text
RTRIM(`u`.`string`)
```
Output PostgreSQL:
```text
RTRIM("u"."string")
```
Output SQLite:
```text
RTRIM("u"."string")
```

#### SubString
Extracts a substring from a string expression starting at `startPos` (1-based) for `lengthStr` characters.
```go
function := uast.SubString(uast.Field[string]("u", "string"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
SUBSTRING(`u`.`string`, ?, ?)
```
Output MsSQL:
```text
SUBSTRING([u].[string], @p1, @p2)
```
Output MySQL:
```text
SUBSTRING(`u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
SUBSTRING("u"."string", $1, $2)
```
Output SQLite:
```text
SUBSTRING("u"."string", ?, ?)
```

#### Trim
Removes both leading and trailing spaces from a string expression.
```go
function := uast.Trim(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
TRIM(`u`.`string`)
```
Output MsSQL:
```text
TRIM([u].[string])
```
Output MySQL:
```text
TRIM(`u`.`string`)
```
Output PostgreSQL:
```text
TRIM("u"."string")
```
Output SQLite:
```text
TRIM("u"."string")
```

#### Upper
Converts a string expression to uppercase.
```go
function := uast.Upper(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
UPPER(`u`.`string`)
```
Output MsSQL:
```text
UPPER([u].[string])
```
Output MySQL:
```text
UPPER(`u`.`string`)
```
Output PostgreSQL:
```text
UPPER("u"."string")
```
Output SQLite:
```text
UPPER("u"."string")
```

## exprLiteral
### Literal
Embeds a raw literal value directly into the generated SQL string (not parameterized). Use with caution — values are written as-is. Prefer `Value` for user-supplied data.
```go
literal := uast.Literal("%Y-%m-%d")
```
Output:
```text
'%Y-%m-%d'
```

## exprLogical
### And
Combines multiple conditions with a logical `AND`. All conditions must be true for the combined expression to be true.
```go
logical := uast.And(
    uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    uast.Greater(uast.Field[int]("u", "number"), uast.Value(2)),
)
```
Output MariaDB:
```text
(`u`.`string` = ? AND `u`.`number` > ?)
```
Output MsSQL:
```text
([u].[string] = @p1 AND [u].[number] > @p2)
```
Output MySQL:
```text
(`u`.`string` = ? AND `u`.`number` > ?)
```
Output PostgreSQL:
```text
("u"."string" = $1 AND "u"."number" > $2)
```
Output SQLite:
```text
("u"."string" = ? AND "u"."number" > ?)
```

### Or
Combines multiple conditions with a logical `OR`. At least one condition must be true for the combined expression to be true.
```go
logical := uast.Or(
    uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    uast.Greater(uast.Field[int]("u", "number"), uast.Value(2)),
)
```
Output MariaDB:
```text
(`u`.`string` = ? OR `u`.`number` > ?)
```
Output MsSQL:
```text
([u].[string] = @p1 OR [u].[number] > @p2)
```
Output MySQL:
```text
(`u`.`string` = ? OR `u`.`number` > ?)
```
Output PostgreSQL:
```text
("u"."string" = $1 OR "u"."number" > $2)
```
Output SQLite:
```text
("u"."string" = ? OR "u"."number" > ?)
```

## exprSubquery
### Subquery
Wraps a `SELECT` statement as a typed expression that can be used in comparisons (`In`, `Exists`, `Equal`, etc.) or as a column in a `SELECT` clause. The generic parameter `u` specifies the scalar type of the single column returned by the subquery.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Field[int64]("u", "id")).From(uast.NewTable("users").As("u")))
```
Output MariaDB:
```text
(SELECT `u`.`id` FROM `users` AS `u`)
```
Output MsSQL:
```text
(SELECT [u].[id] FROM [users] AS [u])
```
Output MySQL:
```text
(SELECT `u`.`id` FROM `users` AS `u`)
```
Output PostgreSQL:
```text
(SELECT "u"."id" FROM "users" AS "u")
```
Output SQLite:
```text
(SELECT "u"."id" FROM "users" AS "u")
```

## exprValue
### Value
Wraps a Go value as a parameterized expression. The value is NOT inserted into the SQL string directly — instead, a placeholder (`?`, `$1`, etc.) is generated and the value is appended to the arguments slice returned by `Build()`. This is the safe way to pass user-supplied data and prevents SQL injection. 
Supported types: `bool`, `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `string`, `time.Time`.
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