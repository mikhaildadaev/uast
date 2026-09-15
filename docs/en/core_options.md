---
outline: deep
---

# API / Core / Options

::: info **Info**
This page covers all configuration options: `clauseGroupBy`, `clauseHaving`, `clauseJoin`, `clauseOrderBy`, `clausePagination`, `clauseReturning`, `clauseSet`, `clauseUnions`, `clauseValues`, `clauseWhere`, `clauseWith`, `exprArray`, `exprBinary`, `exprComparison`, `exprConstant`, `exprField`, `exprFunction`, `exprLiteral`, `exprLogical`, `exprSubquery`, `exprValue`. Each option is shown with a working code example and expected **Output.
:::**

## clauseGroupBy
Adds a GROUP BY clause to group rows by specified columns or expressions.
```go
groupBy := GroupBy(
	uast.Field[string]("u", "string"),
)
```
**Output MariaDB:**
```sql
GROUP BY `u`.`string`
```
**Output MsSQL:**
```sql
GROUP BY [u].[string]
```
**Output MySQL:**
```sql
GROUP BY `u`.`string`
```
**Output PostgreSQL:**
```sql
GROUP BY "u"."string"
```
**Output SQLite:**
```sql
GROUP BY "u"."string"
```

## clauseHaving
Adds a HAVING clause to filter groups. Used with GROUP BY to filter aggregated results.
```go
having := Having(
	uast.Greater(uast.Count(uast.Field[int64]("u", "id"), false), uast.Value[int64](2)),
)
```
**Output MariaDB:**
```sql
HAVING COUNT(`u`.`id`) > ?
```
**Output MsSQL:**
```sql
HAVING COUNT([u].[id]) > @p1
```
**Output MySQL:**
```sql
HAVING COUNT(`u`.`id`) > ?
```
**Output PostgreSQL:**
```sql
HAVING COUNT("u"."id") > $1
```
**Output SQLite:**
```sql
HAVING COUNT("u"."id") > ?
```

## clauseJoin
### Cross
Adds a CROSS JOIN to the statement. Returns the Cartesian product of both tables.
```go
join := uast.Cross(uast.NewTable("users").As("u"))
```
**Output MariaDB:**
```sql
CROSS JOIN `users` AS `u`
```
**Output MsSQL:**
```sql
CROSS JOIN [users] AS [u]
```
**Output MySQL:**
```sql
CROSS JOIN `users` AS `u`
```
**Output PostgreSQL:**
```sql
CROSS JOIN "users" AS "u"
```
**Output SQLite:**
```sql
CROSS JOIN "users" AS "u"
```

### Full
Adds a FULL JOIN to the statement. Returns all rows from both tables, with NULLs where there is no match.
```go
join := uast.Full(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
FULL JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
FULL JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
FULL JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
FULL JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
FULL JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### FullOuter
Adds a FULL OUTER JOIN to the statement. Returns all rows from both tables, with NULLs where there is no match.
```go
join := uast.FullOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
FULL OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
FULL OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
FULL OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
FULL OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
FULL OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Inner
Adds an INNER JOIN to the statement. Returns rows that have matching values in both tables.
```go
join := uast.Inner(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
INNER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
INNER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
INNER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
INNER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
INNER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Left
Adds a LEFT JOIN to the statement. Returns all rows from the left table, and matching rows from the right table.
```go
join := uast.Left(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
LEFT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
LEFT JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
LEFT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
LEFT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
LEFT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### LeftOuter
Adds a LEFT OUTER JOIN to the statement. Returns all rows from the left table, and matching rows from the right table.
```go
join := uast.LeftOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
LEFT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
LEFT OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
LEFT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
LEFT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
LEFT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Right
Adds a RIGHT JOIN to the statement. Returns all rows from the right table, and matching rows from the left table. Not supported by SQLite.
```go
join := uast.Right(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
RIGHT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
RIGHT JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
RIGHT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
RIGHT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
// Not supported
```

### RightOuter
Adds a RIGHT OUTER JOIN to the statement. Returns all rows from the right table, and matching rows from the left table. Not supported by SQLite.
```go
join := uast.RightOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
RIGHT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
RIGHT OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
RIGHT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
RIGHT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
// Not supported
```

## clauseOrderBy
### Asc
Specifies ascending sort order (smallest first, A-to-Z). Used for sorting rows in a query or within a window function.
```go
orderBy := uast.Asc(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
`u`.`string` ASC
```
**Output MsSQL:**
```sql
[u].[string] ASC
```
**Output MySQL:**
```sql
`u`.`string` ASC
```
**Output PostgreSQL:**
```sql
"u"."string" ASC
```
**Output SQLite:**
```sql
"u"."string" ASC
```

### Desc
Specifies descending sort order (largest first, Z-to-A). Used for sorting rows in a query or within a window function.
```go
orderBy := uast.Desc(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
`u`.`string` DESC
```
**Output MsSQL:**
```sql
[u].[string] DESC
```
**Output MySQL:**
```sql
`u`.`string` DESC
```
**Output PostgreSQL:**
```sql
"u"."string" DESC
```
**Output SQLite:**
```sql
"u"."string" DESC
```

## clausePagination
Specifies pagination for a SELECT statement using `Pagination(limit, offset)`. The `limit` sets the maximum number of rows to return. The `offset` specifies the number of rows to skip before returning results. The rendering order and syntax adapts to each dialect automatically.
```go
pagination := Pagination(10,0)
```
**Output MariaDB:**
```sql
LIMIT ? OFFSET ?
```
**Output MsSQL:**
```sql
OFFSET @p1 ROWS FETCH NEXT @p2 ROWS ONLY
```
**Output MySQL:**
```sql
LIMIT ? OFFSET ?
```
**Output PostgreSQL:**
```sql
LIMIT $1 OFFSET $2
```
**Output SQLite:**
```sql
LIMIT ? OFFSET ?
```

## clauseReturning
Adds a RETURNING clause to return modified rows. Supported by MariaDB, PostgreSQL, and SQLite. MySQL does not support this clause natively.
```go
returning = Returning(
	uast.Field[int64]("u", "id"),
    uast.Field[string]("u", "string"),
)
```
**Output MariaDB:**
```sql
RETURNING `u`.`id`, `u`.`string`
```
**Output MsSQL:**
```sql
**Output [u].[id], [u].[string]
```
**Output MySQL:**
```sql
// Not support
```
**Output PostgreSQL:**
```sql
RETURNING "u"."id", "u"."string"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
UPDATE `users` AS `u` SET `u`.`string` = ?
```
**Output MsSQL:**
```sql
UPDATE [users] AS [u] SET [u].[string] = @p1
```
**Output MySQL:**
```sql
UPDATE `users` AS `u` SET `u`.`string` = ?
```
**Output PostgreSQL:**
```sql
UPDATE "users" AS "u" SET "u"."string" = $1
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
UNION SELECT `u`.`string` FROM `users` AS `u` 
```
**Output MsSQL:**
```sql
UNION SELECT [u].[string] FROM [users] AS [u]
```
**Output MySQL:**
```sql
UNION SELECT `u`.`string` FROM `users` AS `u`
```
**Output PostgreSQL:**
```sql
UNION SELECT "u"."string" FROM "users" AS "u"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
UNION ALL SELECT `u`.`string` FROM `users` AS `u`
```
**Output MsSQL:**
```sql
UNION ALL SELECT [u].[string] FROM [users] AS [u]
```
**Output MySQL:**
```sql
UNION ALL SELECT `u`.`string` FROM `users` AS `u`
```
**Output PostgreSQL:**
```sql
UNION ALL SELECT "u"."string" FROM "users" AS "u"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
EXCEPT SELECT `u`.`string` FROM `users` AS `u`
```
**Output MsSQL:**
```sql
EXCEPT SELECT [u].[string] FROM [users] AS [u]
```
**Output MySQL:**
```sql
EXCEPT SELECT `u`.`string` FROM `users` AS `u`
```
**Output PostgreSQL:**
```sql
EXCEPT SELECT "u"."string" FROM "users" AS "u"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
INTERSECT SELECT `u`.`string` FROM `users` AS `u`
```
**Output MsSQL:**
```sql
INTERSECT SELECT [u].[string] FROM [users] AS [u]
```
**Output MySQL:**
```sql
INTERSECT SELECT `u`.`string` FROM `users` AS `u`
```
**Output PostgreSQL:**
```sql
INTERSECT SELECT "u"."string" FROM "users" AS "u"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
VALUES (?, ?)
```
**Output MsSQL:**
```sql
VALUES (@p1, @p2)
```
**Output MySQL:**
```sql
VALUES (?, ?)
```
**Output PostgreSQL:**
```sql
VALUES ($1, $2)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
```
**Output MsSQL:**
```sql
// Not supported
```
**Output MySQL:**
```sql
VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
```
**Output PostgreSQL:**
```sql
VALUES ($1, $2) ON CONFLICT DO UPDATE SET "string" = $3
```
**Output SQLite:**
```sql
VALUES (?, ?) ON CONFLICT DO UPDATE SET "string" = ?
```

## clauseWhere
Adds a WHERE clause to filter rows before grouping or aggregation. Accepts comparison expressions, logical operators, and subqueries.
```go
where = Where(
	uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
)
```
**Output MariaDB:**
```sql
WHERE `u`.`string` = ?
```
**Output MsSQL:**
```sql
WHERE [u].[string] = @p1
```
**Output MySQL:**
```sql
WHERE `u`.`string` = ?
```
**Output PostgreSQL:**
```sql
WHERE "u"."string" = $1
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ?)
```
**Output MsSQL:**
```sql
WITH [cte_norecursive] ([id], [string]) AS (SELECT [u].[id], [u].[string] FROM [users] AS [u] WHERE [u].[string] = @p1)
```
**Output MySQL:**
```sql
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ?)
```
**Output PostgreSQL:**
```sql
WITH "cte_norecursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = $1)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ? UNION ALL SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` INNER JOIN `cte_recursive` AS `rec` ON `u`.`id` = `rec`.`id`)
```
**Output MsSQL:**
```sql
WITH RECURSIVE [cte_recursive] ([id], [string]) AS (SELECT [u].[id], [u].[string] FROM [users] AS [u] WHERE [u].[string] = @p1 UNION ALL SELECT [u].[id], [u].[string] FROM [users] AS [u] INNER JOIN [cte_recursive] AS [rec] ON [u].[id] = [rec].[id])
```
**Output MySQL:**
```sql
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ? UNION ALL SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` INNER JOIN `cte_recursive` AS `rec` ON `u`.`id` = `rec`.`id`)
```
**Output PostgreSQL:**
```sql
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = $1 UNION ALL SELECT "u"."id", "u"."string" FROM "users" AS "u" INNER JOIN "cte_recursive" AS "rec" ON "u"."id" = "rec"."id")
```
**Output SQLite:**
```sql
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = ? UNION ALL SELECT "u"."id", "u"."string" FROM "users" AS "u" INNER JOIN "cte_recursive" AS "rec" ON "u"."id" = "rec"."id")
```

## exprArray
### Array
Constructs an array expression for use in SQL queries.
```go
array := uast.Array(0, 1, 2)
```
**Output MariaDB:**
```sql
ARRAY[?, ?, ?]
```
**Output MsSQL:**
```sql
ARRAY[@p1, @p2, @p3]
```
**Output MySQL:**
```sql
ARRAY[?, ?, ?]
```
**Output PostgreSQL:**
```sql
ARRAY[$1, $2, $3]
```
**Output SQLite:**
```sql
ARRAY[?, ?, ?]
```

## exprBinary
### BitwiseAnd
Performs a bitwise AND operation between two expressions.
```go
binary := uast.BitwiseAnd(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
**Output MariaDB:**
```sql
`u`.`number` & ?
```
**Output MsSQL:**
```sql
[u].[number] & @p1
```
**Output MySQL:**
```sql
`u`.`number` & ?
```
**Output PostgreSQL:**
```sql
"u"."number" & $1
```
**Output SQLite:**
```sql
"u"."number" & ?
```

### BitwiseOr
Performs a bitwise OR operation between two expressions.
```go
binary := uast.BitwiseOr(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
**Output MariaDB:**
```sql
`u`.`number` | ?
```
**Output MsSQL:**
```sql
[u].[number] | @p1
```
**Output MySQL:**
```sql
`u`.`number` | ?
```
**Output PostgreSQL:**
```sql
"u"."number" | $1
```
**Output SQLite:**
```sql
"u"."number" | ?
```

### BitwiseXor
Performs a bitwise XOR operation between two expressions.
```go
binary := uast.BitwiseXor(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
**Output MariaDB:**
```sql
`u`.`number` ^ ?
```
**Output MsSQL:**
```sql
[u].[number] ^ @p1
```
**Output MySQL:**
```sql
`u`.`number` ^ ?
```
**Output PostgreSQL:**
```sql
"u"."number" ^ $1
```
**Output SQLite:**
```sql
"u"."number" ^ ?
```

### Divide
Divides the left expression by the right expression.
```go
binary := uast.Divide(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` / ?
```
**Output MsSQL:**
```sql
[u].[number] / @p1
```
**Output MySQL:**
```sql
`u`.`number` / ?
```
**Output PostgreSQL:**
```sql
"u"."number" / $1
```
**Output SQLite:**
```sql
"u"."number" / ?
```

### Minus
Subtracts the right expression from the left expression.
```go
binary := uast.Minus(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` - ?
```
**Output MsSQL:**
```sql
[u].[number] - @p1
```
**Output MySQL:**
```sql
`u`.`number` - ?
```
**Output PostgreSQL:**
```sql
"u"."number" - $1
```
**Output SQLite:**
```sql
"u"."number" - ?
```

### Modulo
Returns the remainder of dividing the left expression by the right expression.
```go
binary := uast.Modulo(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` % ?
```
**Output MsSQL:**
```sql
[u].[number] % @p1
```
**Output MySQL:**
```sql
`u`.`number` % ?
```
**Output PostgreSQL:**
```sql
"u"."number" % $1
```
**Output SQLite:**
```sql
"u"."number" % ?
```

### Multiply
Multiplies the left expression by the right expression.
```go
binary := uast.Multiply(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` * ?
```
**Output MsSQL:**
```sql
[u].[number] * @p1
```
**Output MySQL:**
```sql
`u`.`number` * ?
```
**Output PostgreSQL:**
```sql
"u"."number" * $1
```
**Output SQLite:**
```sql
"u"."number" * ?
```

### Plus
Adds the left expression to the right expression.
```go
binary := uast.Plus(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` + ?
```
**Output MsSQL:**
```sql
[u].[number] + @p1
```
**Output MySQL:**
```sql
`u`.`number` + ?
```
**Output PostgreSQL:**
```sql
"u"."number" + $1
```
**Output SQLite:**
```sql
"u"."number" + ?
```

### ShiftLeft
Performs a bitwise left shift on the left expression by the number of bits specified in the right expression.
```go
binary := uast.ShiftLeft(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` << ?
```
**Output MsSQL:**
```sql
[u].[number] << @p1
```
**Output MySQL:**
```sql
`u`.`number` << ?
```
**Output PostgreSQL:**
```sql
"u"."number" << $1
```
**Output SQLite:**
```sql
"u"."number" << ?
```

### ShiftRight
Performs a bitwise right shift on the left expression by the number of bits specified in the right expression.
```go
binary := uast.ShiftRight(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` >> ?
```
**Output MsSQL:**
```sql
[u].[number] >> @p1
```
**Output MySQL:**
```sql
`u`.`number` >> ?
```
**Output PostgreSQL:**
```sql
"u"."number" >> $1
```
**Output SQLite:**
```sql
"u"."number" >> ?
```

## exprComparison
### Between
Checks if the left expression falls within the range defined by `valueStart` and `valueEnd` (inclusive).
```go
comparison := uast.Between(uast.Field[int]("u", "number"), uast.Value(0), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` BETWEEN ? AND ?
```
**Output MsSQL:**
```sql
[u].[number] BETWEEN @p1 AND @p2
```
**Output MySQL:**
```sql
`u`.`number` BETWEEN ? AND ?
```
**Output PostgreSQL:**
```sql
"u"."number" BETWEEN $1 AND $2
```
**Output SQLite:**
```sql
"u"."number" BETWEEN ? AND ?
```

### Equal
Compares two expressions for equality (`=`).
```go
comparison := uast.Equal(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` = ?
```
**Output MsSQL:**
```sql
[u].[number] = @p1
```
**Output MySQL:**
```sql
`u`.`number` = ?
```
**Output PostgreSQL:**
```sql
"u"."number" = $1
```
**Output SQLite:**
```sql
"u"."number" = ?
```

### Exists
Checks if the subquery returns any rows. Returns `true` if at least one row exists.
```go
comparison := uast.Exists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("users").As("u"))))
```
**Output MariaDB:**
```sql
EXISTS (SELECT 1 FROM `users` AS `u`)
```
**Output MsSQL:**
```sql
EXISTS (SELECT 1 FROM [users] AS [u])
```
**Output MySQL:**
```sql
EXISTS (SELECT 1 FROM `users` AS `u`)
```
**Output PostgreSQL:**
```sql
EXISTS (SELECT 1 FROM "users" AS "u")
```
**Output SQLite:**
```sql
EXISTS (SELECT 1 FROM "users" AS "u")
```

### Greater
Compares if the left expression is greater than the right expression (`>`).
```go
comparison := uast.Greater(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` > ?
```
**Output MsSQL:**
```sql
[u].[number] > @p1
```
**Output MySQL:**
```sql
`u`.`number` > ?
```
**Output PostgreSQL:**
```sql
"u"."number" > $1
```
**Output SQLite:**
```sql
"u"."number" > ?
```

### GreaterEqual
Compares if the left expression is greater than or equal to the right expression (`>=`).
```go
comparison := uast.GreaterEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` >= ?
```
**Output MsSQL:**
```sql
[u].[number] >= @p1
```
**Output MySQL:**
```sql
`u`.`number` >= ?
```
**Output PostgreSQL:**
```sql
"u"."number" >= $1
```
**Output SQLite:**
```sql
"u"."number" >= ?
```

### ILike
Performs a case-insensitive pattern matching comparison. The right expression should contain a pattern with `%` (any sequence) and `_` (single character) wildcards.
```go
comparison := uast.ILike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
**Output MariaDB:**
```sql
LOWER(`u`.`string`) LIKE LOWER(?)
```
**Output MsSQL:**
```sql
LOWER([u].[string]) LIKE LOWER(@p1)
```
**Output MySQL:**
```sql
LOWER(`u`.`string`) LIKE LOWER(?)
```
**Output PostgreSQL:**
```sql
"u"."string" ILIKE $1
```
**Output SQLite:**
```sql
LOWER("u"."string") LIKE LOWER(?)
```

### In
Checks if the left expression matches any value contained within the right expression (typically a subquery or array).
```go
comparison := uast.In(uast.Field[string]("u", "string"), uast.Array("active", "pending"))
```
**Output MariaDB:**
```sql
`u`.`string` IN (?, ?)
```
**Output MsSQL:**
```sql
[u].[string] IN (@p1, @p2)
```
**Output MySQL:**
```sql
`u`.`string` IN (?, ?)
```
**Output PostgreSQL:**
```sql
"u"."string" IN ($1, $2)
```
**Output SQLite:**
```sql
"u"."string" IN (?, ?)
```

### IsNotNull
Checks if the expression is not `NULL`.
```go
comparison := uast.IsNotNull(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
`u`.`string` IS NOT NULL
```
**Output MsSQL:**
```sql
[u].[string] IS NOT NULL
```
**Output MySQL:**
```sql
`u`.`string` IS NOT NULL
```
**Output PostgreSQL:**
```sql
"u"."string" IS NOT NULL
```
**Output SQLite:**
```sql
"u"."string" IS NOT NULL
```

### IsNull
Checks if the expression is `NULL`.
```go
comparison := uast.IsNull(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
`u`.`string` IS NULL
```
**Output MsSQL:**
```sql
[u].[string] IS NULL
```
**Output MySQL:**
```sql
`u`.`string` IS NULL
```
**Output PostgreSQL:**
```sql
"u"."string" IS NULL
```
**Output SQLite:**
```sql
"u"."string" IS NULL
```

### Less
Compares if the left expression is less than the right expression (`<`).
```go
comparison := uast.Less(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` < ?
```
**Output MsSQL:**
```sql
[u].[number] < @p1
```
**Output MySQL:**
```sql
`u`.`number` < ?
```
**Output PostgreSQL:**
```sql
"u"."number" < $1
```
**Output SQLite:**
```sql
"u"."number" < ?
```

### LessEqual
Compares if the left expression is less than or equal to the right expression (`<=`).
```go
comparison := uast.LessEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` <= ?
```
**Output MsSQL:**
```sql
[u].[number] <= @p1
```
**Output MySQL:**
```sql
`u`.`number` <= ?
```
**Output PostgreSQL:**
```sql
"u"."number" <= $1
```
**Output SQLite:**
```sql
"u"."number" <= ?
```

### Like
Performs a case-sensitive pattern matching comparison. The right expression should contain a pattern with `%` and `_` wildcards.
```go
comparison := uast.Like(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
**Output MariaDB:**
```sql
`u`.`string` LIKE ?
```
**Output MsSQL:**
```sql
[u].[number] LIKE @p1
```
**Output MySQL:**
```sql
`u`.`string` LIKE ?
```
**Output PostgreSQL:**
```sql
"u"."string" LIKE $1
```
**Output SQLite:**
```sql
"u"."string" LIKE ?
```

### NotBetween
Checks if the left expression falls outside the range defined by `valueStart` and `valueEnd`.
```go
comparison := uast.NotBetween(uast.Field[int]("u", "number"), uast.Value(0), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` NOT BETWEEN ? AND ?
```
**Output MsSQL:**
```sql
[u].[number] NOT BETWEEN @p1 AND @p2
```
**Output MySQL:**
```sql
`u`.`number` NOT BETWEEN ? AND ?
```
**Output PostgreSQL:**
```sql
"u"."number" NOT BETWEEN $1 AND $2
```
**Output SQLite:**
```sql
"u"."number" NOT BETWEEN ? AND ?
```

### NotEqual
Compares two expressions for inequality (`!=` or `<>`).
```go
comparison := uast.NotEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` != ?
```
**Output MsSQL:**
```sql
[u].[number] != @p1
```
**Output MySQL:**
```sql
`u`.`number` != ?
```
**Output PostgreSQL:**
```sql
"u"."number" != $1
```
**Output SQLite:**
```sql
"u"."number" != ?
```

### NotExists
Checks if the subquery returns no rows. Returns `true` if the subquery result is empty.
```go
comparison := uast.NotExists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("users").As("u"))))
```
**Output MariaDB:**
```sql
NOT EXISTS (SELECT 1 FROM `users` AS `u`)
```
**Output MsSQL:**
```sql
NOT EXISTS (SELECT 1 FROM [users] AS [u])
```
**Output MySQL:**
```sql
NOT EXISTS (SELECT 1 FROM `users` AS `u`)
```
**Output PostgreSQL:**
```sql
NOT EXISTS (SELECT 1 FROM "users" AS "u")
```
**Output SQLite:**
```sql
NOT EXISTS (SELECT 1 FROM "users" AS "u")
```

### NotILike
Performs a negated case-insensitive pattern matching comparison.
```go
comparison := uast.NotILike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
**Output MariaDB:**
```sql
LOWER(`u`.`string`) NOT LIKE LOWER(?)
```
**Output MsSQL:**
```sql
LOWER([u].[string]) NOT LIKE LOWER(@p1)
```
**Output MySQL:**
```sql
LOWER(`u`.`string`) NOT LIKE LOWER(?)
```
**Output PostgreSQL:**
```sql
"u"."string" NOT ILIKE $1
```
**Output SQLite:**
```sql
LOWER("u"."string") NOT LIKE LOWER(?)
```

### NotIn
Checks if the left expression does not match any value contained within the right expression.
```go
comparison := uast.NotIn(uast.Field[string]("u", "string"), uast.Array("active", "pending"))
```
**Output MariaDB:**
```sql
`u`.`string` NOT IN (?, ?)
```
**Output MsSQL:**
```sql
[u].[string] NOT IN (@p1, @p2)
```
**Output MySQL:**
```sql
`u`.`string` NOT IN (?, ?)
```
**Output PostgreSQL:**
```sql
"u"."string" NOT IN ($1, $2)
```
**Output SQLite:**
```sql
"u"."string" NOT IN (?, ?)
```

### NotLike
Performs a negated case-sensitive pattern matching comparison.
```go
comparison := uast.NotLike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
**Output MariaDB:**
```sql
`u`.`string` NOT LIKE ?
```
**Output MsSQL:**
```sql
[u].[string] NOT LIKE @p1
```
**Output MySQL:**
```sql
`u`.`string` NOT LIKE ?
```
**Output PostgreSQL:**
```sql
"u"."string" NOT LIKE $1
```
**Output SQLite:**
```sql
"u"."string" NOT LIKE ?
```

## exprConstant
### ConstBoolFalse
Returns a constant boolean `FALSE` expression.
```go
constant := uast.ConstBoolFalse()
```
**Output SQL:**
```sql
FALSE
```

### ConstBoolTrue
Returns a constant boolean `TRUE` expression.
```go
constant := uast.ConstBoolTrue()
```
**Output SQL:**
```sql
TRUE
```

### ConstFloat32One
Returns a constant `float32` value of `1.0`. 
```go
constant := uast.ConstFloat32One()
```
**Output SQL:**
```sql
1.0
```

### ConstFloat64One
Returns a constant `float64` value of `1.000000`.
```go
constant := uast.ConstFloat64One()
```
**Output SQL:**
```sql
1.000000
```

### ConstIntOne
Returns a constant `int` value of `1`.
```go
constant := uast.ConstIntOne()
```
**Output SQL:**
```sql
1
```

### ConstInt8One
Returns a constant `int8` value of `1`.
```go
constant := uast.ConstInt8One()
```
**Output SQL:**
```sql
1
```

### ConstInt16One
Returns a constant `int16` value of `1`.
```go
constant := uast.ConstInt16One()
```
**Output SQL:**
```sql
1
```

### ConstInt32One
Returns a constant `int32` value of `1`.
```go
constant := uast.ConstInt32One()
```
**Output SQL:**
```sql
1
```

### ConstInt64One
Returns a constant `int64` value of `1`.
```go
constant := uast.ConstInt64One()
```
**Output SQL:**
```sql
1
```

### ConstStringDefault
Returns a constant `string` value of `DEFAULT`.
```go
constant := uast.ConstStringDefault()
```
**Output SQL:**
```sql
DEFAULT
```

### ConstStringNull
Returns a constant `string` value of `NULL`.
```go
constant := uast.ConstStringNull()
```
**Output SQL:**
```sql
NULL
```

### ConstUintOne
Returns a constant `uint` value of `1`.
```go
constant := uast.ConstUintOne()
```
**Output SQL:**
```sql
1
```

### ConstUint8One
Returns a constant `uint8` value of `1`.
```go
constant := uast.ConstUint8One()
```
**Output SQL:**
```sql
1
```

### ConstUint16One
Returns a constant `uint16` value of `1`.
```go
constant := uast.ConstUint16One()
```
**Output SQL:**
```sql
1
```

### ConstUint32One
Returns a constant `uint32` value of `1`.
```go
constant := uast.ConstUint32One()
```
**Output SQL:**
```sql
1
```

### ConstUint64One
Returns a constant `uint64` value of `1`.
```go
constant := uast.ConstUint64One()
```
**Output SQL:**
```sql
1
```

## exprField
### Field
Creates a reference to a table column, optionally qualified with a table alias. This is the primary way to reference database columns in expressions.
```go
field := uast.Field[string]("u", "string")
```
**Output MariaDB:**
```sql
`u`.`string`
```
**Output MsSQL:**
```sql
[u].[string]
```
**Output MySQL:**
```sql
`u`.`string`
```
**Output PostgreSQL:**
```sql
"u"."string"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
AVG(`u`.`number`)
AVG(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
AVG([u].[number])
AVG(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
AVG(`u`.`number`)
AVG(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
AVG("u"."number")
AVG(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
AVG("u"."number")
AVG(DISTINCT "u"."number")
```

#### BitAnd
Returns the bitwise AND of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitAnd(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitAnd(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
BIT_AND(`u`.`number`)
BIT_AND(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
BIT_AND([u].[number])
BIT_AND(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
BIT_AND(`u`.`number`)
BIT_AND(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
BIT_AND("u"."number")
BIT_AND(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
BIT_AND("u"."number")
BIT_AND(DISTINCT "u"."number")
```

#### BitOr
Returns the bitwise OR of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitOr(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitOr(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
BIT_OR(`u`.`number`)
BIT_OR(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
BIT_OR([u].[number])
BIT_OR(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
BIT_OR(`u`.`number`)
BIT_OR(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
BIT_OR("u"."number")
BIT_OR(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
BIT_OR("u"."number")
BIT_OR(DISTINCT "u"."number")
```

#### BitXor
Returns the bitwise XOR of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitXor(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitXor(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
BIT_XOR(`u`.`number`)
BIT_XOR(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
BIT_XOR([u].[number])
BIT_XOR(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
BIT_XOR(`u`.`number`)
BIT_XOR(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
BIT_XOR("u"."number")
BIT_XOR(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
BIT_XOR("u"."number")
BIT_XOR(DISTINCT "u"."number")
```

#### Count
Returns the number of rows matching the query, or the number of non-NULL values if an expression is provided. When `distinct` is `true`, counts only distinct values.
```go
function := uast.Count(uast.Field[string]("u", "string"), false)
functionWithDistinct := uast.Count(uast.Field[string]("u", "string"), true)
```
**Output MariaDB:**
```sql
COUNT(`u`.`string`)
COUNT(DISTINCT `u`.`string`)
```
**Output MsSQL:**
```sql
COUNT([u].[string])
COUNT(DISTINCT [u].[string])
```
**Output MySQL:**
```sql
COUNT(`u`.`string`)
COUNT(DISTINCT `u`.`string`)
```
**Output PostgreSQL:**
```sql
COUNT("u"."string")
COUNT(DISTINCT "u"."string")
```
**Output SQLite:**
```sql
COUNT("u"."string")
COUNT(DISTINCT "u"."string")
```

#### GroupConcat
Concatenates values from a group into a single string, separated by a default delimiter (typically a comma). The `distinct` flag removes duplicates before concatenation.
```go
function := uast.GroupConcat(uast.Field[string]("u", "string"), false)
functionWithDistinct := uast.GroupConcat(uast.Field[string]("u", "string"), true)
```
**Output MariaDB:**
```sql
GROUP_CONCAT(`u`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `u`.`string` SEPARATOR ',')
```
**Output MsSQL:**
```sql
GROUP_CONCAT([u].[string], ',')
GROUP_CONCAT(DISTINCT [u].[string], ',')
```
**Output MySQL:**
```sql
GROUP_CONCAT(`u`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `u`.`string` SEPARATOR ',')
```
**Output PostgreSQL:**
```sql
STRING_AGG("u"."string", ',')
STRING_AGG(DISTINCT "u"."string", ',')
```
**Output SQLite:**
```sql
GROUP_CONCAT("u"."string" SEPARATOR ',')
GROUP_CONCAT(DISTINCT "u"."string" SEPARATOR ',')
```

#### Max
Returns the maximum value of the expression across all rows in the group.
```go
function := uast.Max(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Max(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
MAX(`u`.`number`)
MAX(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
MAX([u].[number])
MAX(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
MAX(`u`.`number`)
MAX(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
MAX("u"."number")
MAX(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
MAX("u"."number")
MAX(DISTINCT "u"."number")
```

#### Min
Returns the minimum value of the expression across all rows in the group.
```go
function := uast.Min(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Min(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
MIN(`u`.`number`)
MIN(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
MIN([u].[number])
MIN(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
MIN(`u`.`number`)
MIN(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
MIN("u"."number")
MIN(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
MIN("u"."number")
MIN(DISTINCT "u"."number")
```

#### StdDev
Returns the population standard deviation of the expression.
```go
function := uast.StdDev(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.StdDev(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
STDDEV(`u`.`number`)
STDDEV(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
STDEV([u].[number])
STDEV(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
STDDEV(`u`.`number`)
STDDEV(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
STDDEV_SAMP("u"."number")
STDDEV_SAMP(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
STDEV("u"."number")
STDEV(DISTINCT "u"."number")
```

#### Sum
Returns the sum of all values in the expression. If `distinct` is `true`, sums only distinct values.
```go
function := uast.Sum(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Sum(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
SUM(`u`.`number`)
SUM(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
SUM([u].[number])
SUM(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
SUM(`u`.`number`)
SUM(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
SUM("u"."number")
SUM(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
SUM("u"."number")
SUM(DISTINCT "u"."number")
```

#### Variance
Returns the population variance of the expression.
```go
function := uast.Variance(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Variance(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
VARIANCE(`u`.`number`)
VARIANCE(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
VAR([u].[number])
VAR(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
VARIANCE(`u`.`number`)
VARIANCE(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
VAR_SAMP("u"."number")
VAR_SAMP(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
FIRST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
FIRST_VALUE([u].[string]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
FIRST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
FIRST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
LAG(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
**Output MsSQL:**
```sql
LAG([u].[number], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[date] ASC)
```
**Output MySQL:**
```sql
LAG(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
**Output PostgreSQL:**
```sql
LAG("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
LAST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
**Output MsSQL:**
```sql
LAST_VALUE([u].[string]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
**Output MySQL:**
```sql
LAST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
**Output PostgreSQL:**
```sql
LAST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
LEAD(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
**Output MsSQL:**
```sql
LEAD([u].[number], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[date] ASC)
```
**Output MySQL:**
```sql
LEAD(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
**Output PostgreSQL:**
```sql
LEAD("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
NTH_VALUE(`u`.`string`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
**Output MsSQL:**
```sql
NTH_VALUE([u].[string], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
**Output MySQL:**
```sql
NTH_VALUE(`u`.`string`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
**Output PostgreSQL:**
```sql
NTH_VALUE("u"."string", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
CASE WHEN `u`.`number` < ? THEN ? ELSE ? END
```
**Output MsSQL:**
```sql
CASE WHEN [u].[number] < @p1 THEN @p2 ELSE @p3 END
```
**Output MySQL:**
```sql
CASE WHEN `u`.`number` < ? THEN ? ELSE ? END
```
**Output PostgreSQL:**
```sql
CASE WHEN "u"."number" < $1 THEN $2 ELSE $3 END
```
**Output SQLite:**
```sql
CASE WHEN "u"."number" < ? THEN ? ELSE ? END
```

#### Coalesce
Returns the first non-NULL expression from the provided list. Useful for providing fallback values.
```go
function := uast.Coalesce(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
**Output MariaDB:**
```sql
COALESCE(`u`.`createat`, `u`.`updateat`)
```
**Output MsSQL:**
```sql
COALESCE([u].[createat], [u].[updateat])
```
**Output MySQL:**
```sql
COALESCE(`u`.`createat`, `u`.`updateat`)
```
**Output PostgreSQL:**
```sql
COALESCE("u"."createat", "u"."updateat")
```
**Output SQLite:**
```sql
COALESCE("u"."createat", "u"."updateat")
```

#### Greatest
Returns the largest value from the provided list of expressions.
```go
function := uast.Greatest(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
**Output MariaDB:**
```sql
GREATEST(`u`.`createat`, `u`.`updateat`)
```
**Output MsSQL:**
```sql
GREATEST([u].[createat], [u].[updateat])
```
**Output MySQL:**
```sql
GREATEST(`u`.`createat`, `u`.`updateat`)
```
**Output PostgreSQL:**
```sql
GREATEST("u"."createat", "u"."updateat")
```
**Output SQLite:**
```sql
GREATEST("u"."createat", "u"."updateat")
```

#### Least
Returns the smallest value from the provided list of expressions.
```go
function := uast.Least(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
**Output MariaDB:**
```sql
LEAST(`u`.`createat`, `u`.`updateat`)
```
**Output MsSQL:**
```sql
LEAST([u].[createat], [u].[updateat])
```
**Output MySQL:**
```sql
LEAST(`u`.`createat`, `u`.`updateat`)
```
**Output PostgreSQL:**
```sql
LEAST("u"."createat", "u"."updateat")
```
**Output SQLite:**
```sql
LEAST("u"."createat", "u"."updateat")
```

#### NullIf
Returns `NULL` if the two expressions are equal; otherwise returns the first expression.
```go
function := uast.NullIf(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
**Output MariaDB:**
```sql
NULLIF(`u`.`createat`, `u`.`updateat`)
```
**Output MsSQL:**
```sql
NULLIF([u].[createat], [u].[updateat])
```
**Output MySQL:**
```sql
NULLIF(`u`.`createat`, `u`.`updateat`)
```
**Output PostgreSQL:**
```sql
NULLIF("u"."createat", "u"."updateat")
```
**Output SQLite:**
```sql
NULLIF("u"."createat", "u"."updateat")
```

### Convert
#### Cast
Converts an expression to a specified data type.
```go
function := uast.Cast(uast.Field[int]("u", "number"), uast.TypeString)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS CHAR)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS NVARCHAR)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS CHAR)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS VARCHAR)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

#### CharLength
Returns the number of characters in a string expression.
```go
function := uast.CharLength(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
CHAR_LENGTH(`u`.`string`)
```
**Output MsSQL:**
```sql
CHAR_LENGTH([u].[string])
```
**Output MySQL:**
```sql
CHAR_LENGTH(`u`.`string`)
```
**Output PostgreSQL:**
```sql
CHAR_LENGTH("u"."string")
```
**Output SQLite:**
```sql
CHAR_LENGTH("u"."string")
```

#### DateFormat
Formats a datetime expression according to a specified format mask.
```go
function := uast.DateFormat(uast.Field[time.Time]("u", "createat"), uast.Value("%Y-%m-%d"))
```
**Output MariaDB:**
```sql
DATE_FORMAT(`u`.`createat`, '%Y-%m-%d')
```
**Output MsSQL:**
```sql
FORMAT([u].[createat], '%Y-%m-%d')
```
**Output MySQL:**
```sql
DATE_FORMAT(`u`.`createat`, '%Y-%m-%d')
```
**Output PostgreSQL:**
```sql
TO_CHAR("u"."createat", '%Y-%m-%d')
```
**Output SQLite:**
```sql
STRFTIME("u"."createat", '%Y-%m-%d')
```

#### Degrees
Converts an angle from radians to degrees.
```go
function := uast.Degrees(uast.Field[int]("u", "number"))
```
**Output MariaDB:**
```sql
DEGREES(`u`.`number`)
```
**Output MsSQL:**
```sql
DEGREES([u].[number])
```
**Output MySQL:**
```sql
DEGREES(`u`.`number`)
```
**Output PostgreSQL:**
```sql
DEGREES("u"."number")
```
**Output SQLite:**
```sql
DEGREES("u"."number")
```

#### Length
Returns the byte length of a string expression.
```go
function := uast.Length(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
LENGTH(`u`.`string`)
```
**Output MsSQL:**
```sql
LEN([u].[string])
```
**Output MySQL:**
```sql
LENGTH(`u`.`string`)
```
**Output PostgreSQL:**
```sql
LENGTH("u"."string")
```
**Output SQLite:**
```sql
LENGTH("u"."string")
```

#### Position
Returns the starting position of the first occurrence of a substring within a string.
```go
function := uast.Position(uast.Field[string]("u", "string"), uast.Value("old"))
```
**Output MariaDB:**
```sql
POSITION(? IN `u`.`string`)
```
**Output MsSQL:**
```sql
CHARINDEX(@p1, [u].[string])
```
**Output MySQL:**
```sql
POSITION(? IN `u`.`string`)
```
**Output PostgreSQL:**
```sql
POSITION($1 IN "u"."string")
```
**Output SQLite:**
```sql
POSITION(? IN "u"."string")
```

#### Radians
Converts an angle from degrees to radians.
```go
function := uast.Radians(uast.Field[int]("u", "number"))
```
**Output MariaDB:**
```sql
RADIANS(`u`.`number`)
```
**Output MsSQL:**
```sql
RADIANS([u].[number])
```
**Output MySQL:**
```sql
RADIANS(`u`.`number`)
```
**Output PostgreSQL:**
```sql
RADIANS("u"."number")
```
**Output SQLite:**
```sql
RADIANS("u"."number")
```

### Date and time
#### CurDate
Returns the current date (without time).
```go
function := uast.CurDate()
```
**Output MariaDB:**
```sql
CURDATE()
```
**Output MsSQL:**
```sql
CAST(GETDATE() AS DATE)
```
**Output MySQL:**
```sql
CURDATE()
```
**Output PostgreSQL:**
```sql
CURRENT_DATE
```
**Output SQLite:**
```sql
DATE('now')
```

#### CurTime
Returns the current time (without date).
```go
function := uast.CurTime()
```
**Output MariaDB:**
```sql
CURTIME()
```
**Output MsSQL:**
```sql
CAST(GETDATE() AS TIME)
```
**Output MySQL:**
```sql
CURTIME()
```
**Output PostgreSQL:**
```sql
CURRENT_TIME
```
**Output SQLite:**
```sql
TIME('now')
```

#### DateAdd
Adds a time/date interval to a datetime expression and returns the resulting datetime.
```go
function := uast.DateAdd(uast.Field[time.Time]("u", "createat"), uast.Value("2 DAY"))
```
**Output MariaDB:**
```sql
DATE_ADD(`u`.`createat`, INTERVAL 2 DAY)
```
**Output MsSQL:**
```sql
DATEADD(DAY, 2, [u].[createat])
```
**Output MySQL:**
```sql
DATE_ADD(`u`.`createat`, INTERVAL 2 DAY)
```
**Output PostgreSQL:**
```sql
("u"."createat" + INTERVAL '2 DAY')
```
**Output SQLite:**
```sql
DATETIME("u"."createat", '+2 DAY')
```

#### DateDiff
Returns the difference in days between two datetime expressions (`datetimeEnd` - `datetimeStart`).
```go
function := uast.DateDiff(uast.Field[time.Time]("u", "updateat"), uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
DATEDIFF(`u`.`updateat`, `u`.`createat`)
```
**Output MsSQL:**
```sql
DATEDIFF([u].[updateat], [u].[createat])
```
**Output MySQL:**
```sql
DATEDIFF(`u`.`updateat`, `u`.`createat`)
```
**Output PostgreSQL:**
```sql
DATE_PART('day', "u"."updateat" - "u"."createat")
```
**Output SQLite:**
```sql
DATEDIFF("u"."updateat", "u"."createat")
```

#### DateSub
Subtracts a time/date interval from a datetime expression and returns the resulting datetime.
```go
function := uast.DateSub(uast.Field[time.Time]("u", "createat"), uast.Value("2 DAY"))
```
**Output MariaDB:**
```sql
DATE_SUB(`u`.`createat`, INTERVAL 2 DAY)
```
**Output MsSQL:**
```sql
DATEADD(DAY, -2, [u].[createat])
```
**Output MySQL:**
```sql
DATE_SUB(`u`.`createat`, INTERVAL 2 DAY)
```
**Output PostgreSQL:**
```sql
("u"."createat" - INTERVAL '2 DAY')
```
**Output SQLite:**
```sql
DATETIME("u"."createat", '-2 DAY')
```

#### Day
Extracts the day of the month (1–31) from a datetime expression.
```go
function := uast.Day(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
DAY(`u`.`createat`)
```
**Output MsSQL:**
```sql
DAY([u].[createat])
```
**Output MySQL:**
```sql
DAY(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(DAY FROM "u"."createat")
```
**Output SQLite:**
```sql
DAY("u"."createat")
```

#### DayName
Returns the name of the weekday (e.g., 'Monday', 'Tuesday') for a given datetime expression.
```go
function := uast.DayName(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
DAYNAME(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATENAME(WEEKDAY, [u].[createat])
```
**Output MySQL:**
```sql
DAYNAME(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
TO_CHAR("u"."createat", 'Day')
```
**Output SQLite:**
```sql
STRFTIME('%w', "u"."createat")
```

#### Hour
Extracts the hour (0–23) from a datetime expression.
```go
function := uast.Hour(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
HOUR(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATEPART(HOUR, [u].[createat])
```
**Output MySQL:**
```sql
HOUR(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(HOUR FROM "u"."createat")
```
**Output SQLite:**
```sql
HOUR("u"."createat")
```

#### Minute
Extracts the minute (0–59) from a datetime expression.
```go
function := uast.Minute(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
MINUTE(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATEPART(MINUTE, [u].[createat])
```
**Output MySQL:**
```sql
MINUTE(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(MINUTE FROM "u"."createat")
```
**Output SQLite:**
```sql
MINUTE("u"."createat")
```

#### Month
Extracts the month (1–12) from a datetime expression.
```go
function := uast.Month(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
MONTH(`u`.`createat`)
```
**Output MsSQL:**
```sql
MONTH([u].[createat])
```
**Output MySQL:**
```sql
MONTH(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(MONTH FROM "u"."createat")
```
**Output SQLite:**
```sql
MONTH("u"."createat")
```

#### MonthName
Returns the name of the month (e.g., 'January', 'February') for a given datetime expression.
```go
function := uast.MonthName(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
MONTHNAME(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATENAME(MONTH, [u].[createat])
```
**Output MySQL:**
```sql
MONTHNAME(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
TO_CHAR("u"."createat", 'Month')
```
**Output SQLite:**
```sql
STRFTIME('%m', "u"."createat")
```

#### Now
Returns the current date and time.
```go
function := uast.Now()
```
**Output MariaDB:**
```sql
NOW()
```
**Output MsSQL:**
```sql
GETDATE()
```
**Output MySQL:**
```sql
NOW()
```
**Output PostgreSQL:**
```sql
CURRENT_TIMESTAMP
```
**Output SQLite:**
```sql
DATETIME('now')
```

#### Quarter
Extracts the quarter (1–4) from a datetime expression.
```go
function := uast.Quarter(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
QUARTER(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATEPART(QUARTER, [u].[createat])
```
**Output MySQL:**
```sql
QUARTER(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(QUARTER FROM "u"."createat")
```
**Output SQLite:**
```sql
QUARTER("u"."createat")
```

#### Second
Extracts the second (0–59) from a datetime expression.
```go
function := uast.Second(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
SECOND(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATEPART(SECOND, [u].[createat])
```
**Output MySQL:**
```sql
SECOND(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(SECOND FROM "u"."createat")
```
**Output SQLite:**
```sql
SECOND("u"."createat")
```

#### TimeAdd
Adds a time interval to a time/datetime expression and returns the resulting time.
```go
function := uast.TimeAdd(uast.Field[time.Time]("u", "createat"), uast.Value("2 HOUR"))
```
**Output MariaDB:**
```sql
TIME_ADD(`u`.`createat`, '2 HOUR')
```
**Output MsSQL:**
```sql
DATEADD(HOUR, 2, [u].[createat])
```
**Output MySQL:**
```sql
TIME_ADD(`u`.`createat`, '2 HOUR')
```
**Output PostgreSQL:**
```sql
("u"."createat" + INTERVAL '2 HOUR')
```
**Output SQLite:**
```sql
TIME("u"."createat", '+2 HOUR')
```

#### TimeDiff
Returns the difference between two time/datetime expressions (`timeEnd` - `timeStart`).
```go
function := uast.TimeDiff(uast.Field[time.Time]("u", "updateat"), uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
TIMEDIFF(`u`.`updateat`, `u`.`createat`)
```
**Output MsSQL:**
```sql
TIMEDIFF([u].[updateat], [u].[createat])
```
**Output MySQL:**
```sql
TIMEDIFF(`u`.`updateat`, `u`.`createat`)
```
**Output PostgreSQL:**
```sql
DATE_PART('time', "u"."updateat" - "u"."createat")
```
**Output SQLite:**
```sql
TIMEDIFF("u"."updateat", "u"."createat")
```

#### TimeSub
Subtracts a time interval from a time/datetime expression and returns the resulting time.
```go
function := uast.TimeSub(uast.Field[time.Time]("u", "createat"), uast.Value("2 HOUR"))
```
**Output MariaDB:**
```sql
TIME_SUB(`u`.`createat`, '2 HOUR')
```
**Output MsSQL:**
```sql
DATEADD(HOUR, -2, [u].[createat])
```
**Output MySQL:**
```sql
TIME_SUB(`u`.`createat`, '2 HOUR')
```
**Output PostgreSQL:**
```sql
("u"."createat" - INTERVAL '2 HOUR')
```
**Output SQLite:**
```sql
TIME("u"."createat", '-2 HOUR')
```

#### Week
Extracts the week number (1–53) from a datetime expression.
```go
function := uast.Week(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
WEEK(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATEPART(WEEK, [u].[createat])
```
**Output MySQL:**
```sql
WEEK(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(WEEK FROM "u"."createat")
```
**Output SQLite:**
```sql
WEEK("u"."createat")
```

#### Year
Extracts the year from a datetime expression.
```go
function := uast.Year(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
YEAR(`u`.`createat`)
```
**Output MsSQL:**
```sql
YEAR([u].[createat])
```
**Output MySQL:**
```sql
YEAR(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(YEAR FROM "u"."createat")
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_ARRAY(`u`.`json`, ?, ?)
```
**Output MsSQL:**
```sql
JSON_ARRAY([u].[json], @p1, @p2)
```
**Output MySQL:**
```sql
JSON_ARRAY(`u`.`json`, ?, ?)
```
**Output PostgreSQL:**
```sql
JSON_ARRAY("u"."json", $1, $2)
```
**Output SQLite:**
```sql
JSON_ARRAY("u"."json", ?, ?)
```

#### JsonArrayAgg
Aggregates values from a group into a JSON array.
```go
function := uast.JsonArrayAgg(
    uast.Field[string]("u", "json"),
)
```
**Output MariaDB:**
```sql
JSON_ARRAYAGG(`u`.`json`)
```
**Output MsSQL:**
```sql
JSON_ARRAYAGG([u].[json])
```
**Output MySQL:**
```sql
JSON_ARRAYAGG(`u`.`json`)
```
**Output PostgreSQL:**
```sql
JSON_AGG("u"."json")
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_CONTAINS(`u`.`json`, '{"key":"val"}')
```
**Output MsSQL:**
```sql
// Not supported
```
**Output MySQL:**
```sql
JSON_CONTAINS(`u`.`json`, '{"key":"val"}')
```
**Output PostgreSQL:**
```sql
("u"."json" @> '{"key":"val"}')
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
(`u`.`json` ->> '$.parent[0].child')
```
**Output MsSQL:**
```sql
JSON_VALUE([u].[json], '$.parent[0].child')
```
**Output MySQL:**
```sql
(`u`.`json` ->> '$.parent[0].child')
```
**Output PostgreSQL:**
```sql
("u"."json" #>> '{parent,0,child}')
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_OBJECT('key', COUNT(`u`.`json`))
```
**Output MsSQL:**
```sql
JSON_OBJECT('key', COUNT([u].[json]))
```
**Output MySQL:**
```sql
JSON_OBJECT('key', COUNT(`u`.`json`))
```
**Output PostgreSQL:**
```sql
JSON_BUILD_OBJECT('key', COUNT("u"."json"))
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_OBJECTAGG(`u`.`json`, `u`.`number`)
```
**Output MsSQL:**
```sql
JSON_OBJECTAGG([u].[json], [u].[number])
```
**Output MySQL:**
```sql
JSON_OBJECTAGG(`u`.`json`, `u`.`number`)
```
**Output PostgreSQL:**
```sql
JSON_OBJECT_AGG("u"."json", "u"."number")
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_REMOVE(`u`.`json`, '$.key1', '$.key2')
```
**Output MsSQL:**
```sql
JSON_MODIFY(JSON_MODIFY([u].[json], '$.key1', NULL), '$.key2', NULL)
```
**Output MySQL:**
```sql
JSON_REMOVE(`u`.`json`, '$.key1', '$.key2')
```
**Output PostgreSQL:**
```sql
("u"."json" - '{key1}' - '{key2}')
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_SET(`u`.`json`, '$.key1', ?, '$.key2', ?)
```
**Output MsSQL:**
```sql
JSON_MODIFY(JSON_MODIFY([u].[json], '$.key1', @p1), '$.key2', @p2)
```
**Output MySQL:**
```sql
JSON_SET(`u`.`json`, '$.key1', ?, '$.key2', ?)
```
**Output PostgreSQL:**
```sql
jsonb_set(jsonb_set("u"."json", '{key1}', $1), '{key2}', $2)
```
**Output SQLite:**
```sql
JSON_SET("u"."json", '$.key1', ?, '$.key2', ?)
```

#### JsonType
Returns the JSON type of a JSON value (e.g., 'OBJECT', 'ARRAY', 'STRING', 'INTEGER', 'NULL').
```go
function := uast.JsonType(uast.Field[string]("u", "json"))
```
**Output MariaDB:**
```sql
JSON_TYPE(`u`.`json`)
```
**Output MsSQL:**
```sql
// Not supported
```
**Output MySQL:**
```sql
JSON_TYPE(`u`.`json`)
```
**Output PostgreSQL:**
```sql
jsonb_typeof("u"."json")
```
**Output SQLite:**
```sql
JSON_TYPE("u"."json")
```

### Math
#### Abs
Returns the absolute (non-negative) value of a numeric expression.
```go
function := uast.Abs(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
ABS(`u`.`x`)
```
**Output MsSQL:**
```sql
ABS([u].[x])
```
**Output MySQL:**
```sql
ABS(`u`.`x`)
```
**Output PostgreSQL:**
```sql
ABS("u"."x")
```
**Output SQLite:**
```sql
ABS("u"."x")
```

#### ACos
Returns the arc cosine (inverse cosine) of the expression, in radians.
```go
function := uast.ACos(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
ACOS(`u`.`x`)
```
**Output MsSQL:**
```sql
ACOS([u].[x])
```
**Output MySQL:**
```sql
ACOS(`u`.`x`)
```
**Output PostgreSQL:**
```sql
ACOS("u"."x")
```
**Output SQLite:**
```sql
ACOS("u"."x")
```

#### ASin
Returns the arc sine (inverse sine) of the expression, in radians.
```go
function := uast.ASin(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
ASIN(`u`.`x`)
```
**Output MsSQL:**
```sql
ASIN([u].[x])
```
**Output MySQL:**
```sql
ASIN(`u`.`x`)
```
**Output PostgreSQL:**
```sql
ASIN("u"."x")
```
**Output SQLite:**
```sql
ASIN("u"."x")
```

#### ATan
Returns the arc tangent (inverse tangent) of the expression, in radians.
```go
function := uast.ATan(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
ATAN(`u`.`x`)
```
**Output MsSQL:**
```sql
ATAN([u].[x])
```
**Output MySQL:**
```sql
ATAN(`u`.`x`)
```
**Output PostgreSQL:**
```sql
ATAN("u"."x")
```
**Output SQLite:**
```sql
ATAN("u"."x")
```

#### ATan2
Returns the arc tangent of the quotient of its two arguments (`y`/`x`), using their signs to determine the quadrant.
```go
function := uast.ATan2(uast.Field[int]("u", "y"), uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
ATAN2(`u`.`y`, `u`.`x`)
```
**Output MsSQL:**
```sql
ATAN2([u].[y], [u].[x])
```
**Output MySQL:**
```sql
ATAN2(`u`.`y`, `u`.`x`)
```
**Output PostgreSQL:**
```sql
ATAN2("u"."y", "u"."x")
```
**Output SQLite:**
```sql
ATAN2("u"."y", "u"."x")
```

#### Cbrt
Returns the cube root of a numeric expression.
```go
function := uast.Cbrt(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
CBRT(`u`.`x`)
```
**Output MsSQL:**
```sql
CBRT([u].[x])
```
**Output MySQL:**
```sql
CBRT(`u`.`x`)
```
**Output PostgreSQL:**
```sql
CBRT("u"."x")
```
**Output SQLite:**
```sql
CBRT("u"."x")
```

#### Ceil
Returns the smallest integer value not less than the argument (rounds up).
```go
function := uast.Ceil(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
CEILING(`u`.`x`)
```
**Output MsSQL:**
```sql
CEILING([u].[x])
```
**Output MySQL:**
```sql
CEILING(`u`.`x`)
```
**Output PostgreSQL:**
```sql
CEIL("u"."x")
```
**Output SQLite:**
```sql
CEIL("u"."x")
```

#### Cos
Returns the cosine of the expression, where the expression is in radians.
```go
function := uast.Cos(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
COS(`u`.`x`)
```
**Output MsSQL:**
```sql
COS([u].[x])
```
**Output MySQL:**
```sql
COS(`u`.`x`)
```
**Output PostgreSQL:**
```sql
COS("u"."x")
```
**Output SQLite:**
```sql
COS("u"."x")
```

#### Exp
Returns `e` (Euler's number, ~2.71828) raised to the power of the expression.
```go
function := uast.Exp(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
EXP(`u`.`x`)
```
**Output MsSQL:**
```sql
EXP([u].[x])
```
**Output MySQL:**
```sql
EXP(`u`.`x`)
```
**Output PostgreSQL:**
```sql
EXP("u"."x")
```
**Output SQLite:**
```sql
EXP("u"."x")
```

#### Floor
Returns the largest integer value not greater than the argument (rounds down).
```go
function := uast.Floor(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
FLOOR(`u`.`x`)
```
**Output MsSQL:**
```sql
FLOOR([u].[x])
```
**Output MySQL:**
```sql
FLOOR(`u`.`x`)
```
**Output PostgreSQL:**
```sql
FLOOR("u"."x")
```
**Output SQLite:**
```sql
FLOOR("u"."x")
```

#### Ln
Returns the natural logarithm (base `e`) of the expression.
```go
function := uast.Ln(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
LN(`u`.`x`)
```
**Output MsSQL:**
```sql
LN([u].[x])
```
**Output MySQL:**
```sql
LN(`u`.`x`)
```
**Output PostgreSQL:**
```sql
LN("u"."x")
```
**Output SQLite:**
```sql
LN("u"."x")
```

#### Log
Returns the logarithm of the expression to the specified base.
```go
function := uast.Log(uast.Field[int]("u", "x"), uast.Value(2))
```
**Output MariaDB:**
```sql
LOG(`u`.`x`, ?)
```
**Output MsSQL:**
```sql
LOG([u].[x], @p1)
```
**Output MySQL:**
```sql
LOG(`u`.`x`, ?)
```
**Output PostgreSQL:**
```sql
LOG("u"."x", $1)
```
**Output SQLite:**
```sql
LOG("u"."x", ?)
```

#### Mod
Returns the remainder (modulo) of the division of the first expression by the second.
```go
function := uast.Mod(uast.Field[int]("u", "x"), uast.Value(2))
```
**Output MariaDB:**
```sql
MOD(`u`.`x`, ?)
```
**Output MsSQL:**
```sql
MOD([u].[x], @p1)
```
**Output MySQL:**
```sql
MOD(`u`.`x`, ?)
```
**Output PostgreSQL:**
```sql
MOD("u"."x", $1)
```
**Output SQLite:**
```sql
MOD("u"."x", ?)
```

#### Pi
Returns the mathematical constant `π` (~3.14159).
```go
function := uast.Pi()
```
**Output MariaDB:**
```sql
PI()
```
**Output MsSQL:**
```sql
PI()
```
**Output MySQL:**
```sql
PI()
```
**Output PostgreSQL:**
```sql
PI()
```
**Output SQLite:**
```sql
PI()
```

#### Power
Returns the expression raised to the power of the exponent.
```go
function := uast.Power(uast.Field[int]("u", "x"), uast.Value(2))
```
**Output MariaDB:**
```sql
POWER(`u`.`x`, ?)
```
**Output MsSQL:**
```sql
POWER([u].[x], @p1)
```
**Output MySQL:**
```sql
POWER(`u`.`x`, ?)
```
**Output PostgreSQL:**
```sql
POWER("u"."x", $1)
```
**Output SQLite:**
```sql
POWER("u"."x", ?)
```

#### Rand
Returns a random floating-point value in the range [0, 1].
```go
function := uast.Rand()
```
**Output MariaDB:**
```sql
RAND()
```
**Output MsSQL:**
```sql
RAND()
```
**Output MySQL:**
```sql
RAND()
```
**Output PostgreSQL:**
```sql
RANDOM()
```
**Output SQLite:**
```sql
RANDOM()
```

#### Round
Rounds the expression to the specified number of decimal places.
```go
function := uast.Round(uast.Field[int]("u", "x"), uast.Value(2))
```
**Output MariaDB:**
```sql
ROUND(`u`.`x`, ?)
```
**Output MsSQL:**
```sql
ROUND([u].[x], @p1)
```
**Output MySQL:**
```sql
ROUND(`u`.`x`, ?)
```
**Output PostgreSQL:**
```sql
ROUND("u"."x", $1)
```
**Output SQLite:**
```sql
ROUND("u"."x", ?)
```

#### Sin
Returns the sine of the expression, where the expression is in radians.
```go
function := uast.Sin(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
SIN(`u`.`x`)
```
**Output MsSQL:**
```sql
SIN([u].[x])
```
**Output MySQL:**
```sql
SIN(`u`.`x`)
```
**Output PostgreSQL:**
```sql
SIN("u"."x")
```
**Output SQLite:**
```sql
SIN("u"."x")
```

#### Sqrt
Returns the square root of the expression.
```go
function := uast.Sqrt(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
SQRT(`u`.`x`)
```
**Output MsSQL:**
```sql
SQRT([u].[x])
```
**Output MySQL:**
```sql
SQRT(`u`.`x`)
```
**Output PostgreSQL:**
```sql
SQRT("u"."x")
```
**Output SQLite:**
```sql
SQRT("u"."x")
```

#### Tan
Returns the tangent of the expression, where the expression is in radians.
```go
function := uast.Tan(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
TAN(`u`.`x`)
```
**Output MsSQL:**
```sql
TAN([u].[x])
```
**Output MySQL:**
```sql
TAN(`u`.`x`)
```
**Output PostgreSQL:**
```sql
TAN("u"."x")
```
**Output SQLite:**
```sql
TAN("u"."x")
```

#### Trunc
Truncates the numeric expression to the specified number of decimal places (without rounding).
```go
function := uast.Trunc(uast.Field[int]("u", "x"), uast.Value(2))
```
**Output MariaDB:**
```sql
TRUNCATE(`u`.`x`, ?)
```
**Output MsSQL:**
```sql
ROUND([u].[x], @p1, 1)
```
**Output MySQL:**
```sql
TRUNCATE(`u`.`x`, ?)
```
**Output PostgreSQL:**
```sql
TRUNC("u"."x", $1)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
CUME_DIST() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
CUME_DIST() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
CUME_DIST() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
CUME_DIST() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
DENSE_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
DENSE_RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
DENSE_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
DENSE_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
NTILE(2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
NTILE(2) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
NTILE(2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
NTILE(2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
PERCENT_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
PERCENT_RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
PERCENT_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
PERCENT_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
ROW_NUMBER() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
ROW_NUMBER() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
ROW_NUMBER() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
ROW_NUMBER() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
ROW_NUMBER() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

### String
#### Concat
Concatenates two or more string expressions into a single string. `NULL` arguments are treated as empty strings in most dialects.
```go
function := uast.Concat(uast.Field[string]("u", "string"), uast.Value("old"), uast.Value("new"))
```
**Output MariaDB:**
```sql
CONCAT(`u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
CONCAT([u].[string], @p1, @p2)
```
**Output MySQL:**
```sql
CONCAT(`u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
CONCAT("u"."string", $1, $2)
```
**Output SQLite:**
```sql
CONCAT("u"."string", ?, ?)
```

#### ConcatWs
Concatenates two or more string expressions with a specified separator between them. Skips `NULL` arguments.
```go
function := uast.ConcatWs(uast.Value("_"), uast.Field[string]("u", "string"), uast.Value("old"),uast.Value("new"))
```
**Output MariaDB:**
```sql
CONCAT_WS(?, `u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
CONCAT_WS(@p1, [u].[string], @p2, @p3)
```
**Output MySQL:**
```sql
CONCAT_WS(?, `u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
CONCAT_WS($1, "u"."string", $2, $3)
```
**Output SQLite:**
```sql
CONCAT_WS(?, "u"."string", ?, ?)
```

#### LeftString
Returns the leftmost `count` characters from a string expression.
```go
function := uast.LeftString(uast.Field[string]("u", "string"), uast.Value(2))
```
**Output MariaDB:**
```sql
LEFT(`u`.`string`, ?)
```
**Output MsSQL:**
```sql
LEFT([u].[string], @p1)
```
**Output MySQL:**
```sql
LEFT(`u`.`string`, ?)
```
**Output PostgreSQL:**
```sql
LEFT("u"."string", $1)
```
**Output SQLite:**
```sql
LEFT("u"."string", ?)
```

#### Lower
Converts a string expression to lowercase.
```go
function := uast.Lower(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
LOWER(`u`.`string`)
```
**Output MsSQL:**
```sql
LOWER([u].[string])
```
**Output MySQL:**
```sql
LOWER(`u`.`string`)
```
**Output PostgreSQL:**
```sql
LOWER("u"."string")
```
**Output SQLite:**
```sql
LOWER("u"."string")
```

#### LPad
Left-pads a string expression with the specified separator to a total length of `count` characters.
```go
function := uast.LPad(uast.Field[string]("u", "string"), uast.Value(2), uast.Value(","))
```
**Output MariaDB:**
```sql
LPAD(`u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
LPAD([u].[string], @p1, @p2)
```
**Output MySQL:**
```sql
LPAD(`u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
LPAD("u"."string", $1, $2)
```
**Output SQLite:**
```sql
LPAD("u"."string", ?, ?)
```

#### LTrim
Removes leading spaces from a string expression.
```go
function := uast.LTrim(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
LTRIM(`u`.`string`)
```
**Output MsSQL:**
```sql
LTRIM([u].[string])
```
**Output MySQL:**
```sql
LTRIM(`u`.`string`)
```
**Output PostgreSQL:**
```sql
LTRIM("u"."string")
```
**Output SQLite:**
```sql
LTRIM("u"."string")
```

#### Repeat
Repeats a string expression `count` times.
```go
function := uast.Repeat(uast.Field[string]("u", "string"), uast.Value(2))
```
**Output MariaDB:**
```sql
REPEAT(`u`.`string`, ?)
```
**Output MsSQL:**
```sql
REPEAT([u].[string], @p1)
```
**Output MySQL:**
```sql
REPEAT(`u`.`string`, ?)
```
**Output PostgreSQL:**
```sql
REPEAT("u"."string", $1)
```
**Output SQLite:**
```sql
REPEAT("u"."string", ?)
```

#### Replace
Replaces all occurrences of a substring in a string with a new substring.
```go
function := uast.Replace(uast.Field[string]("u", "string"), uast.Value("old"), uast.Value("new"))
```
**Output MariaDB:**
```sql
REPLACE(`u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
REPLACE([u].[string], @p1, @p2)
```
**Output MySQL:**
```sql
REPLACE(`u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
REPLACE("u"."string", $1, $2)
```
**Output SQLite:**
```sql
REPLACE("u"."string", ?, ?)
```

#### Reverse
Reverses the characters in a string expression.
```go
function := uast.Reverse(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
REVERSE(`u`.`string`)
```
**Output MsSQL:**
```sql
REVERSE([u].[string])
```
**Output MySQL:**
```sql
REVERSE(`u`.`string`)
```
**Output PostgreSQL:**
```sql
REVERSE("u"."string")
```
**Output SQLite:**
```sql
REVERSE("u"."string")
```

#### RightString
Returns the rightmost `count` characters from a string expression.
```go
function := uast.RightString(uast.Field[string]("u", "string"), uast.Value(2))
```
**Output MariaDB:**
```sql
RIGHT(`u`.`string`, ?)
```
**Output MsSQL:**
```sql
RIGHT([u].[string], @p1)
```
**Output MySQL:**
```sql
RIGHT(`u`.`string`, ?)
```
**Output PostgreSQL:**
```sql
RIGHT("u"."string", $1)
```
**Output SQLite:**
```sql
RIGHT("u"."string", ?)
```

#### RPad
Right-pads a string expression with the specified separator to a total length of `count` characters.
```go
function := uast.RPad(uast.Field[string]("u", "string"), uast.Value(2), uast.Value(","))
```
**Output MariaDB:**
```sql
RPAD(`u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
RPAD([u].[string], @p1, @p2)
```
**Output MySQL:**
```sql
RPAD(`u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
RPAD("u"."string", $1, $2)
```
**Output SQLite:**
```sql
RPAD("u"."string", ?, ?)
```

#### RTrim
Removes trailing spaces from a string expression.
```go
function := uast.RTrim(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
RTRIM(`u`.`string`)
```
**Output MsSQL:**
```sql
RTRIM([u].[string])
```
**Output MySQL:**
```sql
RTRIM(`u`.`string`)
```
**Output PostgreSQL:**
```sql
RTRIM("u"."string")
```
**Output SQLite:**
```sql
RTRIM("u"."string")
```

#### SubString
Extracts a substring from a string expression starting at `startPos` (1-based) for `lengthStr` characters.
```go
function := uast.SubString(uast.Field[string]("u", "string"), uast.Value(0), uast.Value(2))
```
**Output MariaDB:**
```sql
SUBSTRING(`u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
SUBSTRING([u].[string], @p1, @p2)
```
**Output MySQL:**
```sql
SUBSTRING(`u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
SUBSTRING("u"."string", $1, $2)
```
**Output SQLite:**
```sql
SUBSTRING("u"."string", ?, ?)
```

#### Trim
Removes both leading and trailing spaces from a string expression.
```go
function := uast.Trim(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
TRIM(`u`.`string`)
```
**Output MsSQL:**
```sql
TRIM([u].[string])
```
**Output MySQL:**
```sql
TRIM(`u`.`string`)
```
**Output PostgreSQL:**
```sql
TRIM("u"."string")
```
**Output SQLite:**
```sql
TRIM("u"."string")
```

#### Upper
Converts a string expression to uppercase.
```go
function := uast.Upper(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
UPPER(`u`.`string`)
```
**Output MsSQL:**
```sql
UPPER([u].[string])
```
**Output MySQL:**
```sql
UPPER(`u`.`string`)
```
**Output PostgreSQL:**
```sql
UPPER("u"."string")
```
**Output SQLite:**
```sql
UPPER("u"."string")
```

## exprLiteral
### Literal
Embeds a raw literal value directly into the generated SQL string (not parameterized). Use with caution — values are written as-is. Prefer `Value` for user-supplied data.
```go
literal := uast.Literal("%Y-%m-%d")
```
**Output SQL:**
```sql
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
**Output MariaDB:**
```sql
(`u`.`string` = ? AND `u`.`number` > ?)
```
**Output MsSQL:**
```sql
([u].[string] = @p1 AND [u].[number] > @p2)
```
**Output MySQL:**
```sql
(`u`.`string` = ? AND `u`.`number` > ?)
```
**Output PostgreSQL:**
```sql
("u"."string" = $1 AND "u"."number" > $2)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
(`u`.`string` = ? OR `u`.`number` > ?)
```
**Output MsSQL:**
```sql
([u].[string] = @p1 OR [u].[number] > @p2)
```
**Output MySQL:**
```sql
(`u`.`string` = ? OR `u`.`number` > ?)
```
**Output PostgreSQL:**
```sql
("u"."string" = $1 OR "u"."number" > $2)
```
**Output SQLite:**
```sql
("u"."string" = ? OR "u"."number" > ?)
```

## exprSubquery
### Subquery
Wraps a `SELECT` statement as a typed expression that can be used in comparisons (`In`, `Exists`, `Equal`, etc.) or as a column in a `SELECT` clause. The generic parameter `u` specifies the scalar type of the single column returned by the subquery.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Field[int64]("u", "id")).From(uast.NewTable("users").As("u")))
```
**Output MariaDB:**
```sql
(SELECT `u`.`id` FROM `users` AS `u`)
```
**Output MsSQL:**
```sql
(SELECT [u].[id] FROM [users] AS [u])
```
**Output MySQL:**
```sql
(SELECT `u`.`id` FROM `users` AS `u`)
```
**Output PostgreSQL:**
```sql
(SELECT "u"."id" FROM "users" AS "u")
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
?
```
**Output MsSQL:**
```sql
@p1
```
**Output MySQL:**
```sql
?
```
**Output PostgreSQL:**
```sql
$1
```
**Output SQLite:**
```sql
?
```