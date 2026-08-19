---
outline: deep
---

# API / Ядро / Конструкторы

::: info **Информация**
На этой странице описано, как создать экземпляр sql, настроить все параметры и разобраться в каждом типе данных и конструкторе полей.
:::

## NewSQL
Создание экземпляра SQL, со всеми параметрами конфигурации
```go
builder := uast.NewSQL(
    uast.WithDialect(uast.DialectMySQL),
    uast.WithMutate(false),
)
defer builder.Close()
stmt := uast.NewSelect(uast.NewTable("users", "u")).
    Fields(
        uast.Field[int64]("u", "id"),
        uast.Field[string]("u", "name"),
    ).
    Where(
        uast.Equal(uast.Field[string]("u", "status"), uast.Value("active")),
    )
query1, args1, err1 := builder.Build(stmt)
if err1 != nil {
    log.Fatal(err1)
}
builder.SetDialect(uast.DialectPostgreSQL)
query2, args2, err2 := builder.Build(stmt)
```
Output MySQL:
```text
SELECT `u`.`id`, `u`.`name` FROM `users` AS `u` WHERE `u`.`status` = ?
```
Output PostgreSQL:
```text
SELECT "u"."id", "u"."name" FROM "users" AS "u" WHERE "u"."status" = $1
```

| Name                                                      | Description	                                                                      | Values	                                                                    | Default           |
|-----------------------------------------------------------|-------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|-------------------|
| [`WithDialect()`](/ru/sql_options#withdialect-setdialect) | Sets the SQL dialect for the builder                                                | DialectMariaDB, DialectMsSQL, DialectMySQL DialectPostgreSQL, DialectSQLite | DialectPostgreSQL |
| [`WithMutate()`](/ru/sql_options#withmutate-setmutate)    | Controls whether the builder modifies the AST in place or clones it before building | true, false                                                                 | false             |

| Name	                                   | Description	                                                                                                      | Returns                |
|------------------------------------------|----------------------------------------------------------------------------------------------------------------------|------------------------|
| [`Build()`](/ru/sql_methods#build)	   | Compiles a statement into a SQL string and a list of arguments                                                       | (string, []any, error) |
| [`Exec()`](/ru/sql_methods#exec)         | Builds and executes a statement that doesn't return rows (e.g., INSERT, UPDATE, DELETE). Uses `db.Exec()` internally | (sql.Result, error)    |
| [`Query()`](/ru/sql_methods#query)	   | Builds and executes a query that returns multiple rows (e.g., SELECT). Uses `db.Query()` internally                  | (*sql.Rows, error)     |
| [`QueryRow()`](/ru/sql_methods#queryrow) | Builds and executes a query that returns at most one row. Uses `db.QueryRow()` internally                            | (*sql.Row, error)      |