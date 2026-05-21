---
outline: deep
---

# API / Core / Options

::: info **Info**
This page covers all configuration options: `exprGroupBy`, `exprHaving`, `clauseJoin`, `clauseLimit`, `clauseOffset`, `clauseOrderBy`, `clauseReturning`, `clauseSet`, `clauseUnions`, `clauseValues`, `clauseWhere`, `clauseWith`, `exprArray`, `exprBinary`, `exprColumn`, `exprComparison`, `exprConstant`, `exprFunction`, `exprLiteral`, `exprLogical`, `exprSubquery`, `exprValue`. Each option is shown with a working code example and expected output.
:::

## exprGroupBy
Adds a GROUP BY clause to group rows by specified columns or expressions.
```go
groupBy := GroupBy(
	uast.Column[string]("t", "string"),
)
```
Output MariaDB:
```text
GROUP BY `t`.`string`
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

## exprHaving
Adds a HAVING clause to filter groups. Used with GROUP BY to filter aggregated results.
```go
having := Having(
	uast.Greater(uast.Count(uast.Column[int64]("t", "id"), false), uast.Value[int64](2)),
)
```
Output MariaDB:
```text
HAVING COUNT(`t`.`id`) > ?
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
Adds a CROSS JOIN to the statement. Returns the Cartesian product of both tables.
```go
join := uast.Cross(uast.Test.Table)
```
Output MariaDB:
```text
CROSS JOIN `test` AS `t`
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
Adds a FULL JOIN to the statement. Returns all rows from both tables, with NULLs where there is no match.
```go
join := uast.Full(uast.Test.Table, uast.Equal(uast.Test.Column.ID, uast.Test1.Column.ID))
```
Output MariaDB:
```text
FULL JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
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
Adds a FULL OUTER JOIN to the statement. Returns all rows from both tables, with NULLs where there is no match.
```go
join := uast.FullOuter(uast.Test.Table, uast.Equal(uast.Test.Column.ID, uast.Test1.Column.ID))
```
Output MariaDB:
```text
FULL OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
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
Adds an INNER JOIN to the statement. Returns rows that have matching values in both tables.
```go
join := uast.Inner(uast.Test.Table, uast.Equal(uast.Test.Column.ID, uast.Test1.Column.ID))
```
Output MariaDB:
```text
INNER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
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
Adds a LEFT JOIN to the statement. Returns all rows from the left table, and matching rows from the right table.
```go
join := uast.Left(uast.Test.Table, uast.Equal(uast.Test.Column.ID, uast.Test1.Column.ID))
```
Output MariaDB:
```text
LEFT JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
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
Adds a LEFT OUTER JOIN to the statement. Returns all rows from the left table, and matching rows from the right table.
```go
join := uast.LeftOuter(uast.Test.Table, uast.Equal(uast.Test.Column.ID, uast.Test1.Column.ID))
```
Output MariaDB:
```text
LEFT OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
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
Adds a RIGHT JOIN to the statement. Returns all rows from the right table, and matching rows from the left table. Not supported by SQLite.
```go
join := uast.Right(uast.Test.Table, uast.Equal(uast.Test.Column.ID, uast.Test1.Column.ID))
```
Output MariaDB:
```text
RIGHT JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
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
Adds a RIGHT OUTER JOIN to the statement. Returns all rows from the right table, and matching rows from the left table. Not supported by SQLite.
```go
join := uast.RightOuter(uast.Test.Table, uast.Equal(uast.Test.Column.ID, uast.Test1.Column.ID))
```
Output MariaDB:
```text
RIGHT OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
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

## clauseLimit
Limits the number of rows returned by the query.
```go
limit := Limit(10)
```
Output MariaDB:
```text
LIMIT ?
```
Output MySQL:
```text
LIMIT ?
```
Output PostgreSQL:
```text
LIMIT $1
```
Output SQLite:
```text
LIMIT ?
```

## clauseOffset
Skips a specified number of rows before returning results. Used for pagination with Limit.
```go
offset := Offset(20)
```
Output MariaDB:
```text
OFFSET ?
```
Output MySQL:
```text
OFFSET ?
```
Output PostgreSQL:
```text
OFFSET $1
```
Output SQLite:
```text
OFFSET ?
```

## clauseOrderBy
### Asc
Specifies ascending sort order (smallest first, A-to-Z). Used for sorting rows in a query or within a window function.
```go
orderBy := uast.Asc(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` ASC
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
Specifies descending sort order (largest first, Z-to-A). Used for sorting rows in a query or within a window function.
```go
orderBy := uast.Desc(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` DESC
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

## exprReturning
Adds a RETURNING clause to return modified rows. Supported by MariaDB, PostgreSQL, and SQLite. MySQL does not support this clause natively.
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
Output MySQL:
```text
// Not support
```
Output PostgreSQL:
```text
RETURNING `t`.`id`, `t`.`string`
```
Output SQLite:
```text
RETURNING `t`.`id`, `t`.`string`
```

## clauseSet
Specifies columns and their new values using `Assign` to associate columns with values. Supports multiple pairs for updating multiple columns.
```go
set := Set(
	uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
)
```
Output MariaDB:
```text
UPDATE `test` AS `t` SET `t`.`string` = ?
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
Combines results from multiple SELECT statements. UNION returns distinct rows.
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
Combines results from multiple SELECT statements. UNION ALL returns all rows, including duplicates.
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
Combines results from multiple SELECT statements. EXCEPT returns distinct rows from the first query that are not in the second.
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
Combines results from multiple SELECT statements. INTERSECT returns distinct rows that are common to both queries.
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
Specifies values for insertion using `Pair` to associate columns with values. Columns are automatically inferred from the pairs.
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

## clauseWhere
Adds a WHERE clause to filter rows before grouping or aggregation. Accepts comparison expressions, logical operators, and subqueries.
```go
where = Where(
	uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
)
```
Output MariaDB:
```text
WHERE `t`.`string` = ?
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
### Non-Recursive
...
### Recursive
...

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
binary := uast.BitwiseAnd(uast.Column[int]("t", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`t`.`number` & ?
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
Performs a bitwise OR operation between two expressions.
```go
binary := uast.BitwiseOr(uast.Column[int]("t", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`t`.`number` | ?
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
Performs a bitwise XOR operation between two expressions.
```go
binary := uast.BitwiseXor(uast.Column[int]("t", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`t`.`number` ^ ?
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
Divides the left expression by the right expression.
```go
binary := uast.Divide(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` / ?
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
Subtracts the right expression from the left expression.
```go
binary := uast.Minus(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` - ?
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
Returns the remainder of dividing the left expression by the right expression.
```go
binary := uast.Modulo(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` % ?
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
Multiplies the left expression by the right expression.
```go
binary := uast.Multiply(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` * ?
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
Adds the left expression to the right expression.
```go
binary := uast.Plus(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` + ?
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
Performs a bitwise left shift on the left expression by the number of bits specified in the right expression.
```go
binary := uast.ShiftLeft(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` << ?
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
Performs a bitwise right shift on the left expression by the number of bits specified in the right expression.
```go
binary := uast.ShiftRight(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` >> ?
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
Creates a reference to a table column, optionally qualified with a table alias. This is the primary way to reference database columns in expressions.
```go
column := uast.Column[string]("t", "string")
```
Output MariaDB:
```text
`t`.`string`
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
Checks if the left expression falls within the range defined by `valueStart` and `valueEnd` (inclusive).
```go
comparison := uast.Between(uast.Column[int]("t", "number"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` BETWEEN ? AND ?
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
Compares two expressions for equality (`=`).
```go
comparison := uast.Equal(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` = ?
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
Checks if the subquery returns any rows. Returns `true` if at least one row exists.
```go
comparison := uast.Exists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("test").As("t"))))
```
Output MariaDB:
```text
EXISTS (SELECT 1 FROM `test` AS `t`)
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
Compares if the left expression is greater than the right expression (`>`).
```go
comparison := uast.Greater(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` > ?
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
Compares if the left expression is greater than or equal to the right expression (`>=`).
```go
comparison := uast.GreaterEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` >= ?
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
Performs a case-insensitive pattern matching comparison. The right expression should contain a pattern with `%` (any sequence) and `_` (single character) wildcards.
```go
comparison := uast.ILike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
LOWER(`t`.`string`) LIKE LOWER(?)
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
Checks if the left expression matches any value contained within the right expression (typically a subquery or array).
```go
comparison := uast.In(uast.Column[string]("t", "string"), uast.Array("active", "pending"))
```
Output MariaDB:
```text
`t`.`string` IN (?, ?)
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
Checks if the expression is not `NULL`.
```go
comparison := uast.IsNotNull(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` IS NOT NULL
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
Checks if the expression is `NULL`.
```go
comparison := uast.IsNull(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` IS NULL
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
Compares if the left expression is less than the right expression (`<`).
```go
comparison := uast.Less(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` < ?
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
Compares if the left expression is less than or equal to the right expression (`<=`).
```go
comparison := uast.LessEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` <= ?
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
Performs a case-sensitive pattern matching comparison. The right expression should contain a pattern with `%` and `_` wildcards.
```go
comparison := uast.Like(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
`t`.`string` LIKE ?
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
Checks if the left expression falls outside the range defined by `valueStart` and `valueEnd`.
```go
comparison := uast.NotBetween(uast.Column[int]("t", "number"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` NOT BETWEEN ? AND ?
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
Compares two expressions for inequality (`!=` or `<>`).
```go
comparison := uast.NotEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` != ?
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
Checks if the subquery returns no rows. Returns `true` if the subquery result is empty.
```go
comparison := uast.NotExists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("test").As("t"))))
```
Output MariaDB:
```text
NOT EXISTS (SELECT 1 FROM `test` AS `t`)
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
Performs a negated case-insensitive pattern matching comparison.
```go
comparison := uast.NotILike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
LOWER(`t`.`string`) NOT LIKE LOWER(?)
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
Checks if the left expression does not match any value contained within the right expression.
```go
comparison := uast.NotIn(uast.Column[string]("t", "string"), uast.Array("active", "pending"))
```
Output MariaDB:
```text
`t`.`string` NOT IN (?, ?)
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
Performs a negated case-sensitive pattern matching comparison.
```go
comparison := uast.NotLike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
`t`.`string` NOT LIKE ?
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

## exprFunction
### Aggregate
#### Avg
Returns the average (arithmetic mean) of all non-NULL values in the expression. If `distinct` is `true`, the average is calculated over distinct values only.
```go
function := uast.Avg(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Avg(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
AVG(`t`.`number`)
AVG(DISTINCT `t`.`number`)
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
Returns the bitwise AND of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitAnd(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitAnd(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
BIT_AND(`t`.`number`)
BIT_AND(DISTINCT `t`.`number`)
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
Returns the bitwise OR of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitOr(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitOr(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
BIT_OR(`t`.`number`)
BIT_OR(DISTINCT `t`.`number`)
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
Returns the bitwise XOR of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitXor(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitXor(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
BIT_XOR(`t`.`number`)
BIT_XOR(DISTINCT `t`.`number`)
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
Returns the number of rows matching the query, or the number of non-NULL values if an expression is provided. When `distinct` is `true`, counts only distinct values.
```go
function := uast.Count(uast.Column[string]("t", "string"), false)
functionWithDistinct := uast.Count(uast.Column[string]("t", "string"), true)
```
Output MariaDB:
```text
COUNT(`t`.`string`)
COUNT(DISTINCT `t`.`string`)
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
Concatenates values from a group into a single string, separated by a default delimiter (typically a comma). The `distinct` flag removes duplicates before concatenation.
```go
function := uast.GroupConcat(uast.Column[string]("t", "string"), false)
functionWithDistinct := uast.GroupConcat(uast.Column[string]("t", "string"), true)
```
Output MariaDB:
```text
GROUP_CONCAT(`t`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `t`.`string` SEPARATOR ',')
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
Returns the maximum value of the expression across all rows in the group.
```go
function := uast.Max(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Max(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
MAX(`t`.`number`)
MAX(DISTINCT `t`.`number`)
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
Returns the minimum value of the expression across all rows in the group.
```go
function := uast.Min(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Min(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
MIN(`t`.`number`)
MIN(DISTINCT `t`.`number`)
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
Returns the population standard deviation of the expression.
```go
function := uast.StdDev(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.StdDev(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
STDDEV(`t`.`number`)
STDDEV(DISTINCT `t`.`number`)
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
Returns the sum of all values in the expression. If `distinct` is `true`, sums only distinct values.
```go
function := uast.Sum(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Sum(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
SUM(`t`.`number`)
SUM(DISTINCT `t`.`number`)
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
Returns the population variance of the expression.
```go
function := uast.Variance(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Variance(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
VARIANCE(`t`.`number`)
VARIANCE(DISTINCT "t"."number")
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
Returns the value of the expression from the first row of the window frame. Requires an `OVER` clause with window specification.
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
Returns the value of the expression from a row that is `offset` rows before the current row within the partition.
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
Returns the value of the expression from the last row of the window frame.
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
Returns the value of the expression from a row that is `offset` rows after the current row within the partition.
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
Returns the value of the expression from the `n-th` row of the window frame.
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
Evaluates a list of `WHEN`-`THEN` pairs and returns the `THEN` expression for the first true WHEN. If no condition is true, returns the `ELSE` expression if provided, or `NULL`.
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
Returns the first non-NULL expression from the provided list. Useful for providing fallback values.
```go
function := uast.Coalesce(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
COALESCE(`t`.`createat`, `t`.`updateat`)
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
Returns the largest value from the provided list of expressions.
```go
function := uast.Greatest(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
GREATEST(`t`.`createat`, `t`.`updateat`)
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
Returns the smallest value from the provided list of expressions.
```go
function := uast.Least(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
LEAST(`t`.`createat`, `t`.`updateat`)
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
Returns `NULL` if the two expressions are equal; otherwise returns the first expression.
```go
function := uast.NullIf(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
NULLIF(`t`.`createat`, `t`.`updateat`)
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
Converts an expression to a specified data type.
```go
function := uast.Cast(uast.Column[int]("t", "number"), uast.TypeString)
```
Output MariaDB:
```text
CAST(`t`.`number` AS CHAR)
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
Returns the number of characters in a string expression.
```go
function := uast.CharLength(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
CHAR_LENGTH(`t`.`string`)
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
Formats a datetime expression according to a specified format mask.
```go
function := uast.DateFormat(uast.Column[time.Time]("t", "createat"), uast.Value("%Y-%m-%d"))
```
Output MariaDB:
```text
DATE_FORMAT(`t`.`createat`, '%Y-%m-%d')
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
strftime("t"."createat", '%Y-%m-%d')
```

#### Degrees
Converts an angle from radians to degrees.
```go
function := uast.Degrees(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
DEGREES(`t`.`number`)
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
Returns the byte length of a string expression.
```go
function := uast.Length(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
LENGTH(`t`.`string`)
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
Returns the starting position of the first occurrence of a substring within a string.
```go
function := uast.Position(uast.Column[string]("t", "string"), uast.Value("old"))
```
Output MariaDB:
```text
POSITION(? IN `t`.`string`)
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
Converts an angle from degrees to radians.
```go
function := uast.Radians(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
RADIANS(`t`.`number`)
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
Returns the current date (without time).
```go
function := uast.CurDate()
```
Output MariaDB:
```text
CURDATE()
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
date('now')
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
time('now')
```

#### DateAdd
Adds a time/date interval to a datetime expression and returns the resulting datetime.
```go
function := uast.DateAdd(uast.Column[time.Time]("t", "createat"), uast.Value("2 DAY"))
```
Output MariaDB:
```text
DATE_ADD(`t`.`createat`, INTERVAL 2 DAY)
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
datetime("t"."createat",  '+2 DAY')
```

#### DateDiff
Returns the difference in days between two datetime expressions (`datetimeEnd` - `datetimeStart`).
```go
function := uast.DateDiff(uast.Column[time.Time]("t", "updateat"), uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
DATEDIFF(`t`.`updateat`, `t`.`createat`)
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
Subtracts a time/date interval from a datetime expression and returns the resulting datetime.
```go
function := uast.DateSub(uast.Column[time.Time]("t", "createat"), uast.Value("2 DAY"))
```
Output MariaDB:
```text
DATE_SUB(`t`.`createat`, INTERVAL 2 DAY)
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
datetime("t"."createat", '-2 DAY')
```

#### Day
Extracts the day of the month (1–31) from a datetime expression.
```go
function := uast.Day(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
DAY(`t`.`createat`)
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
Returns the name of the weekday (e.g., 'Monday', 'Tuesday') for a given datetime expression.
```go
function := uast.DayName(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
DAYNAME(`t`.`createat`)
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
strftime('%w', "t"."createat")
```

#### Hour
Extracts the hour (0–23) from a datetime expression.
```go
function := uast.Hour(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
HOUR(`t`.`createat`)
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
Extracts the minute (0–59) from a datetime expression.
```go
function := uast.Minute(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
MINUTE(`t`.`createat`)
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
Extracts the month (1–12) from a datetime expression.
```go
function := uast.Month(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
MONTH(`t`.`createat`)
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
Returns the name of the month (e.g., 'January', 'February') for a given datetime expression.
```go
function := uast.MonthName(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
MONTHNAME(`t`.`createat`)
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
strftime('%m', "t"."createat")
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
datetime('now')
```

#### Quarter
Extracts the quarter (1–4) from a datetime expression.
```go
function := uast.Quarter(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
QUARTER(`t`.`createat`)
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
Extracts the second (0–59) from a datetime expression.
```go
function := uast.Second(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
SECOND(`t`.`createat`)
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
Adds a time interval to a time/datetime expression and returns the resulting time.
```go
function := uast.TimeAdd(uast.Column[time.Time]("t", "createat"), uast.Value("2 HOUR"))
```
Output MariaDB:
```text
TIME_ADD(`t`.`createat`, '2 HOUR')
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
time("t"."createat", '+2 HOUR')
```

#### TimeDiff
Returns the difference between two time/datetime expressions (`timeEnd` - `timeStart`).
```go
function := uast.TimeDiff(uast.Column[time.Time]("t", "updateat"), uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
TIMEDIFF(`t`.`updateat`, `t`.`createat`)
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
Subtracts a time interval from a time/datetime expression and returns the resulting time.
```go
function := uast.TimeSub(uast.Column[time.Time]("t", "createat"), uast.Value("2 HOUR"))
```
Output MariaDB:
```text
TIME_SUB(`t`.`createat`, '2 HOUR')
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
time("t"."createat", '-2 HOUR')
```

#### Week
Extracts the week number (1–53) from a datetime expression.
```go
function := uast.Week(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
WEEK(`t`.`createat`)
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
Extracts the year from a datetime expression.
```go
function := uast.Year(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
YEAR(`t`.`createat`)
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
Creates a JSON array from the given expression and optional additional values.
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
Aggregates values from a group into a JSON array.
```go
function := uast.JsonArrayAgg(
    uast.Column[string]("t", "json"),
)
```
Output MariaDB:
```text
JSON_ARRAYAGG(`t`.`json`)
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
Checks whether a JSON document contains a specified value.
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
Extracts a value from a JSON document at the specified path. The `json` parameter is built with `JsonPath` and optional `JsonKey`/`JsonIndex`.
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
Builds a JSON object from key-value pairs.
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
JSON_OBJECT('key', COUNT(`t`.`json`))
```

#### JsonObjectAgg
Aggregates key-value pairs from a group into a single JSON object.
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
Removes a value from a JSON document at the specified path(s).
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
Sets a value in a JSON document at the specified path(s). Creates the path if it does not exist.
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
Output MySQL:
```text
JSON_SET(`t`.`json`, '$.key1', ?, '$.key2', ?)
```
Output PostgreSQL:
```text
-
```
Output SQLite:
```text
JSON_SET("t"."json", '$.key1', ?, '$.key2', ?)
```

#### JsonType
Returns the JSON type of a JSON value (e.g., 'OBJECT', 'ARRAY', 'STRING', 'INTEGER', 'NULL').
```go
function := uast.JsonType(uast.Column[string]("t", "json"))
```
Output MariaDB:
```text
JSON_TYPE(`t`.`json`)
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
Returns the absolute (non-negative) value of a numeric expression.
```go
function := uast.Abs(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ABS(`t`.`number`)
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
Returns the arc cosine (inverse cosine) of the expression, in radians.
```go
function := uast.ACos(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ACOS(`t`.`number`)
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
Returns the arc sine (inverse sine) of the expression, in radians.
```go
function := uast.ASin(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ASIN(`t`.`number`)
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
Returns the arc tangent (inverse tangent) of the expression, in radians.
```go
function := uast.ATan(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ATAN(`t`.`number`)
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
Returns the arc tangent of the quotient of its two arguments (`y`/`x`), using their signs to determine the quadrant.
```go
function := uast.ATan2(uast.Column[int]("t", "y"), uast.Column[int]("t", "x"))
```
Output MariaDB:
```text
ATAN2(`t`.`y`, `t`.`x`)
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
Returns the cube root of a numeric expression.
```go
function := uast.Cbrt(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
CBRT(`t`.`number`)
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
Returns the smallest integer value not less than the argument (rounds up).
```go
function := uast.Ceil(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
CEILING(`t`.`number`)
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
Returns the cosine of the expression, where the expression is in radians.
```go
function := uast.Cos(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
COS(`t`.`number`)
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
Returns `e` (Euler's number, ~2.71828) raised to the power of the expression.
```go
function := uast.Exp(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
EXP(`t`.`number`)
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
Returns the largest integer value not greater than the argument (rounds down).
```go
function := uast.Floor(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
FLOOR(`t`.`number`)
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
Returns the natural logarithm (base `e`) of the expression.
```go
function := uast.Ln(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
LN(`t`.`number`)
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
Returns the logarithm of the expression to the specified base.
```go
function := uast.Log(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
LOG(`t`.`number`, ?)
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
Returns the remainder (modulo) of the division of the first expression by the second.
```go
function := uast.Mod(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
MOD(`t`.`number`, ?)
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
Returns the mathematical constant `p` (~3.14159).
```go
function := uast.Pi()
```
Output MariaDB:
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
function := uast.Power(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
POWER(`t`.`number`, ?)
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
Returns a random floating-point value in the range [0, 1].
```go
function := uast.Rand()
```
Output MariaDB:
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
function := uast.Round(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
ROUND(`t`.`number`, ?)
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
Returns the sine of the expression, where the expression is in radians.
```go
function := uast.Sin(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
SIN(`t`.`number`)
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
Returns the square root of the expression.
```go
function := uast.Sqrt(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
SQRT(`t`.`number`)
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
Returns the tangent of the expression, where the expression is in radians.
```go
function := uast.Tan(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
TAN(`t`.`number`)
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
Truncates the numeric expression to the specified number of decimal places (without rounding).
```go
function := uast.Trunc(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
TRUNCATE(`t`.`number`, ?)
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
Returns the cumulative distribution of a value within a partition (the ratio of rows that come before or are peers with the current row). Must be used with an `OVER` clause.
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
Returns the rank of a row without gaps. Rows with equal values receive the same rank, and the next rank is the immediate next integer. Requires `OVER`.
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
Divides the rows within a partition into `n` approximately equal groups and returns the group number (1 through `n`) for each row.
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
Returns the percentile rank of a row within a partition (range 0 to 1). Rank of first row is always 0. Requires `OVER`.
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
Returns the rank of a row with gaps. Equal values receive the same rank, and the next distinct value skips ahead. Requires `OVER`.
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
Assigns a unique sequential integer to each row within the partition, starting from 1. Order determines the numbering sequence.
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
Concatenates two or more string expressions into a single string. `NULL` arguments are treated as empty strings in most dialects.
```go
function := uast.Concat(uast.Column[string]("t", "string"), uast.Value("old"), uast.Value("new"))
```
Output MariaDB:
```text
CONCAT(`t`.`string`, ?, ?)
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
Concatenates two or more string expressions with a specified separator between them. Skips `NULL` arguments.
```go
function := uast.ConcatWs(uast.Value("_"), uast.Column[string]("t", "string"), uast.Value("old"),uast.Value("new"))
```
Output MariaDB:
```text
CONCAT_WS(?, `t`.`string`, ?, ?)
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
Returns the leftmost `count` characters from a string expression.
```go
function := uast.LeftString(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MariaDB:
```text
LEFT(`t`.`string`, ?)
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
Converts a string expression to lowercase.
```go
function := uast.Lower(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
LOWER(`t`.`string`)
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
Left-pads a string expression with the specified separator to a total length of `count` characters.
```go
function := uast.LPad(uast.Column[string]("t", "string"), uast.Value(2), uast.Value(","))
```
Output MariaDB:
```text
LPAD(`t`.`string`, ?, ?)
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
Removes leading spaces from a string expression.
```go
function := uast.LTrim(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
LTRIM(`t`.`string`)
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
Repeats a string expression `count` times.
```go
function := uast.Repeat(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MariaDB:
```text
REPEAT(`t`.`string`, ?)
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
Replaces all occurrences of a substring in a string with a new substring.
```go
function := uast.Replace(uast.Column[string]("t", "string"), uast.Value("old"), uast.Value("new"))
```
Output MariaDB:
```text
REPLACE(`t`.`string`, ?, ?)
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
Reverses the characters in a string expression.
```go
function := uast.Reverse(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
REVERSE(`t`.`string`)
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
Returns the rightmost `count` characters from a string expression.
```go
function := uast.RightString(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MariaDB:
```text
RIGHT(`t`.`string`, ?)
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
Right-pads a string expression with the specified separator to a total length of `count` characters.
```go
function := uast.RPad(uast.Column[string]("t", "string"), uast.Value(2), uast.Value(","))
```
Output MariaDB:
```text
RPAD(`t`.`string`, ?, ?)
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
Removes trailing spaces from a string expression.
```go
function := uast.RTrim(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
RTRIM(`t`.`string`)
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
Extracts a substring from a string expression starting at `startPos` (1-based) for `lengthStr` characters.
```go
function := uast.SubString(uast.Column[string]("t", "string"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
SUBSTRING(`t`.`string`, ?, ?)
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
Removes both leading and trailing spaces from a string expression.
```go
function := uast.Trim(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
TRIM(`t`.`string`)
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
Converts a string expression to uppercase.
```go
function := uast.Upper(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
UPPER(`t`.`string`)
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
    uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    uast.Greater(uast.Column[int]("t", "number"), uast.Value(2)),
)
```
Output MariaDB:
```text
(`t`.`string` = ? AND `t`.`number` > ?)
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
Combines multiple conditions with a logical `OR`. At least one condition must be true for the combined expression to be true.
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
Wraps a `SELECT` statement as a typed expression that can be used in comparisons (`In`, `Exists`, `Equal`, etc.) or as a column in a `SELECT` clause. The generic parameter `T` specifies the scalar type of the single column returned by the subquery.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Column[int64]("t", "id")).From(uast.NewTable("test").As("t")))
```
Output MariaDB:
```text
(SELECT `t`.`id` FROM `test` AS `t`)
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
Wraps a Go value as a parameterized expression. The value is NOT inserted into the SQL string directly — instead, a placeholder (`?`, `$1`, etc.) is generated and the value is appended to the arguments slice returned by `Build()`. This is the safe way to pass user-supplied data and prevents SQL injection. 
Supported types: `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `string`, `time.Time`.
```go
var data string = "ivan"
value := uast.Value(data)
```
Output MariaDB:
```text
?
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