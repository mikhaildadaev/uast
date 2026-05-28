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
`WithMutable` 将构建器标记为可变。在可变模式下，`Build()` 会直接修改原始语句而不是克隆它，从而提高了单次使用语句的性能。可变构建器不支持 `SetDialect`。
```go
mariadbStmt := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "id"), uast.Value(1)),
    )
mssqlStmt := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "id"), uast.Value(1)),
    )
mysqlStmt := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "id"), uast.Value(1)),
    )
postgresqlStmt := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "id"), uast.Value(1)),
    )
sqliteStmt := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "id"), uast.Value(1)),
    )
mariadbBuilder := uast.NewSQL(
    uast.WithDialect(uast.DialectMariaDB),
    uast.WithMutable(),
)
defer mariadbBuilder.Close()
mssqlBuilder := uast.NewSQL(
    uast.WithDialect(uast.DialectMsSQL),
    uast.WithMutable(),
)
defer mssqlBuilder.Close()
mysqlBuilder := uast.NewSQL(
    uast.WithDialect(uast.DialectMySQL),
    uast.WithMutable(),
)
defer mysqlBuilder.Close()
postgresqlBuilder := uast.NewSQL(
    uast.WithDialect(uast.DialectPostgreSQL),
    uast.WithMutable(),
)
defer postgresqlBuilder.Close()
sqliteBuilder := uast.NewSQL(
    uast.WithDialect(uast.DialectSQLite),
    uast.WithMutable(),
)
defer sqliteBuilder.Close()
mariadbQuery, mariadbArgs, _ := mariadbBuilder.Build(mariadbStmt)
mssqlQuery, mssqlArgs, _ := mssqlBuilder.Build(mssqlStmt)
mysqlQuery, mysqlArgs, _ := mysqlBuilder.Build(mysqlStmt)
postgresqlQuery, postgresqlArgs, _ := postgresqlBuilder.Build(postgresqlStmt)
sqliteQuery, sqliteArgs, _ := sqliteBuilder.Build(sqliteStmt)
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