---
outline: deep
---

# API / 核心 / 构造函数

::: info **关于**
本页文档介绍如何创建遥测实例、配置所有设置以及了解每个数据类型和字段构造函数。
:::

## NewSQL
为指定方言配置的 SQL 实例
```go
builder := uast.NewSQL()
defer builder.Close()
```
Output:
```text
...
```

| Name                                                      | Description	                                                                      | Values	                                                                    | Default           |
|-----------------------------------------------------------|-------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|-------------------|
| [`WithDialect()`](/zh/sql_options#withdialect-setdialect) | Sets the SQL dialect for the builder                                                | DialectMariaDB, DialectMsSQL, DialectMySQL DialectPostgreSQL, DialectSQLite | DialectPostgreSQL |
| [`WithMutate()`](/zh/sql_options#withmutate-setmutate)    | Controls whether the builder modifies the AST in place or clones it before building | true, false                                                                 | false             |

| Name	                                   | Description	                                                                                                      | Returns                |
|------------------------------------------|----------------------------------------------------------------------------------------------------------------------|------------------------|
| [`Build()`](/zh/sql_methods#build)	   | Compiles a statement into a SQL string and a list of arguments                                                       | (string, []any, error) |
| [`Exec()`](/zh/sql_methods#exec)         | Builds and executes a statement that doesn't return rows (e.g., INSERT, UPDATE, DELETE). Uses `db.Exec()` internally | (sql.Result, error)    |
| [`Query()`](/zh/sql_methods#query)	   | Builds and executes a query that returns multiple rows (e.g., SELECT). Uses `db.Query()` internally                  | (*sql.Rows, error)     |
| [`QueryRow()`](/zh/sql_methods#queryrow) | Builds and executes a query that returns at most one row. Uses `db.QueryRow()` internally                            | (*sql.Row, error)      |