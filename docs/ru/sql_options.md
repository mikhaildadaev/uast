---
outline: deep
---

# API / SQL / Опции

::: info **Информация**
На этой странице описаны все параметры конфигурации: `Dialect`. Каждая опция показана с рабочим примером кода и ожидаемым выводом.
:::

## WithDialect/SetDialect
`WithDialect` устанавливает диалект при создании экземпляра. `SetDialect` переключает диалект существующего экземпляра во время выполнения без пересоздания пула соединений.
```go
stmt := uast.NewSelect(uast.Column[string]("t", "string")).
    From(
        uast.NewTable("test").As("t"),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "id"), uast.Value(1)),
    )
sql := uast.NewSQL(
    uast.WithDialect(uast.DialectMariaDB)
)
defer sql.Close()
mariadbQuery, mariadbArgs, _ := sql.Build(stmt)
sql.SetDialect(uast.DialectMsSQL)
mssqlQuery, mssqlArgs, _ := sql.Build(stmt)
sql.SetDialect(uast.DialectMySQL)
mysqlQuery, mysqlArgs, _ := sql.Build(stmt)
sql.SetDialect(uast.DialectPostgreSQL)
pgsqlQuery, pgsqlArgs, _ := sql.Build(stmt)
sql.SetDialect(uast.DialectSQLite)
sqliteQuery, sqliteArgs, _ := sql.Build(stmt)
```
Output MariaDB:
```text
SELECT `t`.`string` FROM `test` AS `t` WHERE `t`.`id` = ?
```
Output MsSQL:
```text
SELECT [t].[string] FROM [test] AS [t] WHERE [t].[id] = @p1
```
Output MySQL:
```text
SELECT `t`.`string` FROM `test` AS `t` WHERE `t`.`id` = ?
```
Output PostgreSQL:
```text
SELECT "t"."string" FROM "test" AS "t" WHERE "t"."id" = $1
```
Output SQLite:
```text
SELECT "t"."string" FROM "test" AS "t" WHERE "t"."id" = ?
```