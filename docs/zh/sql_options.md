---
outline: deep
---

# API / SQL / 选项

::: info **关于**
本页面涵盖 SQL 构建器的配置选项：用于创建时设置的 `WithDialect` 和 `WithMutate`，以及用于运行时修改的 `SetDialect` 和 `SetMutate`。每个选项都配有可运行的代码示例和预期输出。
:::

## WithDialect/SetDialect
`WithDialect` 在创建实例时设置方言。 `SetDialect` 在运行时切换现有实例的方言，无需重新创建连接池。
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
`WithMutate` 在创建时将构建器标记为可变。`SetMutate` 在运行时打开或关闭可变模式。启用可变模式时，`Build()` 直接修改原始语句而不是克隆它，从而提高一次性查询的性能。禁用可变模式时，`Build()` 在构建之前克隆语句，保留原始语句以供重用。可变模式启用时，`SetDialect` 被阻止。

一旦在可变模式下构建了语句，该语句已被修改，无法安全重用 — 后续构建会产生未定义的结果。要在可变模式后安全重用语句，请创建一个新的语句实例。
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