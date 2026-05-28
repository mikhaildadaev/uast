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

## WithMutate/SetMutate
`WithMutate` marks the builder as mutable at creation time. `SetMutate` switches mutation mode on or off at runtime. When mutation is enabled, `Build()` mutates the original statement instead of cloning it, improving performance for single-use statements. When mutation is disabled, `Build()` clones the statement before building, preserving the original for reuse. `SetDialect` is blocked while mutation is enabled.

Once a statement is built with mutation enabled, it is modified and cannot be safely reused — subsequent builds produce undefined results. To safely reuse a statement after mutation mode, create a new statement instance.
```go
stmt1 := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "id"), uast.Value(1)),
    )
immutableBuilder := uast.NewSQL(
    uast.WithDialect(uast.DialectPostgreSQL),
)
defer immutableBuilder.Close()
query1, _, _ := immutableBuilder.Build(stmt1)
query2, _, _ := immutableBuilder.Build(stmt1)
immutableBuilder.SetMutate(true)
query3, _, _ := immutableBuilder.Build(stmt1)
query4, _, _ := immutableBuilder.Build(stmt1)
stmt2 := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "id"), uast.Value(1)),
    )
stmt3 := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "id"), uast.Value(1)),
    )
mutableBuilder := uast.NewSQL(
    uast.WithDialect(uast.DialectPostgreSQL),
    uast.WithMutate(true),
)
defer mutableBuilder.Close()
query5, _, _ := mutableBuilder.Build(stmt2)
query6, _, _ := mutableBuilder.Build(stmt2)
mutableBuilder.SetMutate(false)
query7, _, _ := mutableBuilder.Build(stmt2)
query8, _, _ := mutableBuilder.Build(stmt3)
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
Output Query7:
```text
// Undefined result — stmt was mutated
```
Output Query8:
```text
SELECT "t"."string" FROM "test" AS "t" WHERE "t"."id" = $1
```