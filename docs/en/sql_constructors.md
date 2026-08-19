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

| Name                                                      | Description	                                                                      | Values	                                                                    | Default           |
|-----------------------------------------------------------|-------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|-------------------|
| [`WithDialect()`](/en/sql_options#withdialect-setdialect) | Sets the SQL dialect for the builder                                                | DialectMariaDB, DialectMsSQL, DialectMySQL DialectPostgreSQL, DialectSQLite | DialectPostgreSQL |
| [`WithMutate()`](/en/sql_options#withmutate-setmutate)    | Controls whether the builder modifies the AST in place or clones it before building | true, false                                                                 | false             |

| Name	                                   | Description	                                                                                                      | Returns                |
|------------------------------------------|----------------------------------------------------------------------------------------------------------------------|------------------------|
| [`Exec()`](/en/sql_methods#exec)         | Builds and executes a statement that doesn't return rows (e.g., INSERT, UPDATE, DELETE). Uses `db.Exec()` internally | (sql.Result, error)    |
| [`Query()`](/en/sql_methods#query)	   | Builds and executes a query that returns multiple rows (e.g., SELECT). Uses `db.Query()` internally                  | (*sql.Rows, error)     |
| [`QueryRow()`](/en/sql_methods#queryrow) | Builds and executes a query that returns at most one row. Uses `db.QueryRow()` internally                            | (*sql.Row, error)      |