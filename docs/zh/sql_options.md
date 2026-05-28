---
outline: deep
---

# API / SQL / 选项

::: info **关于**
本页涵盖所有配置选项：`Dialect`。每个选项都附有可运行的代码示例和预期输出。
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

## WithMutable
`WithMutable` 在创建时将构建器标记为可变。`SetMutable` 将现有构建器切换到可变模式。在可变模式下，`Build()` 会直接修改原始语句而不是克隆它，从而提高了单次使用语句的性能。可变构建器不支持 `SetDialect`。在可变模式下构建语句后，该语句已被修改，无法安全重用——后续构建将产生未定义的结果。
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