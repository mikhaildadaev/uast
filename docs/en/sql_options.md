---
outline: deep
---

# API / SQL / Options

::: info **Info**
This page covers all configuration options: `Dialect`. Each option is shown with a working code example and expected output.
:::

## WithDialect/SetDialect
`WithDialect` sets the dialect at creation time. `SetDialect` switches the dialect of an existing instance at runtime without recreating the connection pool.
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
postgresqlQuery, postgresqlArgs, _ := sql.Build(stmt)
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