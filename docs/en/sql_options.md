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
stmt := uast.NewSelect(uast.Column[string]("t", "name")).
    From(uast.Table("test")).
    Where(uast.Equal(uast.Column[int]("t", "id"), uast.Value(1)))
sql := uast.NewSQL(
    uast.WithDialect(uast.DialectMySQL)
)
defer sql.Close()
mysqlQuery, mysqlArgs, _ := sql.Build(stmt)
sql.SetDialect(uast.DialectPostgreSQL)
pgsqlQuery, pgsqlArgs, _ := sql.Build(stmt)
```
Output MySQL:
```text
SELECT `t`.`name` FROM `test` AS `t` WHERE `t`.`id` = ?
```
Output PostgreSQL:
```text
SELECT "t"."name" FROM "test" AS "t" WHERE "t"."id" = $1
```