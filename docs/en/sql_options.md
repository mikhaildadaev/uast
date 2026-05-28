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
builder := uast.NewSQL(
    uast.WithDialect(uast.DialectMariaDB)
)
defer builder.Close()
mariadbQuery, mariadbArgs, _ := builder.Build(stmt)
builder.SetDialect(uast.DialectMsSQL)
mssqlQuery, mssqlArgs, _ := builder.Build(stmt)
builder.SetDialect(uast.DialectMySQL)
mysqlQuery, mysqlArgs, _ := builder.Build(stmt)
builder.SetDialect(uast.DialectPostgreSQL)
postgresqlQuery, postgresqlArgs, _ := builder.Build(stmt)
builder.SetDialect(uast.DialectSQLite)
sqliteQuery, sqliteArgs, _ := builder.Build(stmt)
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

## WithMutable/SetMutable
`WithMutable` marks the builder as mutable at creation time. `SetMutable` switches an existing builder to mutable mode. When mutable, `Build()` mutates the original statement instead of cloning it, improving performance for single-use statements. `SetDialect` is blocked for mutable builders. Once a statement is built in mutable mode, it is modified and cannot be safely reused — subsequent builds produce undefined results.
```go
stmt := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "id"), uast.Value(1)),
    )
immutableSQL := uast.NewSQL(
    uast.WithDialect(uast.DialectPostgreSQL),
)
defer immutableSQL.Close()
query1, _, _ := immutableSQL.Build(stmt)
query2, _, _ := immutableSQL.Build(stmt)
immutableSQL.SetMutable()
query3, _, _ := immutableSQL.Build(stmt)
query4, _, _ := immutableSQL.Build(stmt)
mutableSQL := uast.NewSQL(
    uast.WithDialect(uast.DialectPostgreSQL),
    uast.WithMutable(),
)
defer mutableSQL.Close()
query5, _, _ := mutableSQL.Build(stmt)
query6, _, _ := mutableSQL.Build(stmt)
```
Output Query1:
```text
SELECT "t"."string" FROM "test" AS "t" WHERE "t"."id" = $1
```
Output Query2:
```text
SELECT "t"."string" FROM "test" AS "t" WHERE "t"."id" = $1
```
Output Query3:
```text
SELECT "t"."string" FROM "test" AS "t" WHERE "t"."id" = $1
```
Output Query4:
```text
// Undefined result — stmt was mutated
```
Output Query5:
```text
SELECT "t"."string" FROM "test" AS "t" WHERE "t"."id" = $1
```
Output Query6:
```text
// Undefined result — stmt was mutated
```