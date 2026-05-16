---
outline: deep
---

# API / SQL / Methods

::: info **Info**
This page documents methods available on statement instances: `Delete`, `Insert`, `Select`, `Update`. Each method configures a specific clause and returns the statement for chaining. Every method is shown with a working code example and expected SQL output.
:::

## stmtDelete
### Join
Adds JOIN clauses to the DELETE statement. Supports INNER, LEFT, RIGHT, FULL, CROSS and their OUTER variants.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Returning
Adds a RETURNING clause to return deleted rows. Supported by PostgreSQL. MySQL does not support this clause natively.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Where
Adds a WHERE clause to filter rows for deletion. Accepts comparison expressions, logical operators, and subqueries.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### With
Adds a Common Table Expression (CTE) to the DELETE statement using `WithN` (non-recursive) or `WithR` (recursive).
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

## stmtInsert
### Returning
Adds a RETURNING clause to return inserted rows. Supported by PostgreSQL. MySQL does not support this clause natively.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Source
Specifies a subquery as the data source for INSERT. Used for `INSERT ... SELECT` statements. When using `Source`, columns are inferred from the subquery fields.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Values
Specifies values for insertion using `Pair` to associate columns with values. Columns are automatically inferred from the pairs.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### With
Adds a Common Table Expression (CTE) to the INSERT statement.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

## stmtSelect
### Distinct
Adds the DISTINCT modifier to remove duplicate rows from the result set.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Field
Specifies the fields to select. Accepts columns, functions, subqueries, and aliases.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### GroupBy
Adds a GROUP BY clause to group rows by specified columns or expressions.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Having
Adds a HAVING clause to filter groups. Used with GROUP BY to filter aggregated results.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Join
Adds JOIN clauses to combine rows from multiple tables. Supports all 8 join types.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Limit
Limits the number of rows returned by the query.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Offset
Skips a specified number of rows before returning results. Used for pagination with Limit.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### OrderBy
Adds an ORDER BY clause to sort results by specified columns or expressions in ascending or descending order.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Unions
Combines results from multiple SELECT statements using UNION, UNION ALL, EXCEPT, or INTERSECT.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Where
Adds a WHERE clause to filter rows before grouping or aggregation.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### With
Adds a Common Table Expression (CTE) to the SELECT statement.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

## stmtUpdate
### Join
Adds JOIN clauses to the UPDATE statement for updating rows based on related table data.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Returning
Adds a RETURNING clause to return updated rows. Supported by PostgreSQL.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Set
Specifies columns and their new values using `Pair` to associate columns with values. Supports multiple pairs for updating multiple columns.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### Where
Adds a WHERE clause to filter rows for updating.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```

### With
Adds a Common Table Expression (CTE) to the UPDATE statement.
```go
...
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```
