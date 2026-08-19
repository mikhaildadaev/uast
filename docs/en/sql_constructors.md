---
outline: deep
---

# API / SQL / Constructors

::: info **Info**
This page describes how to create a sql instance, configure all the settings, and understand each data type and field constructor.
:::

## NewSQL
SQL instance configured for the specified dialect
```go
builder := uast.NewSQL()
defer builder.Close()
```
Output:
```text
...
```

| Name            | Description	                                                                        | Values	                                                                  | Default           |
|-----------------|-------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|-------------------|
| `WithDialect()` |	Sets the SQL dialect for the builder                                                | DialectMariaDB, DialectMsSQL, DialectMySQL DialectPostgreSQL, DialectSQLite | DialectPostgreSQL |
| `WithMutate()`  | Controls whether the builder modifies the AST in place or clones it before building | true, false                                                                 | false             |

::: tip **Note**
`SetDialect` is blocked when mutation is enabled `WithMutate()`.
:::

| Name	     | Description	                                                                                                      | Returns                |
|------------|--------------------------------------------------------------------------------------------------------------------|------------------------|
| Build()	 | Compiles a statement into a SQL string and a list of arguments                                                     | (string, []any, error) |
| Exec()     | Builds and executes a statement that doesn't return rows (e.g., INSERT, UPDATE, DELETE). Uses db.Exec() internally | (sql.Result, error)    |
| Query()	 | Builds and executes a query that returns multiple rows (e.g., SELECT). Uses db.Query() internally                  | (*sql.Rows, error)     |
| QueryRow() | Builds and executes a query that returns at most one row. Uses db.QueryRow() internally                            | (*sql.Row, error)      |