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
`WithMutable` помечает построитель как изменяемый. В изменяемом режиме `Build()` изменяет исходный statement вместо клонирования, что повышает производительность для одноразовых запросов. `SetDialect` заблокирован для изменяемых построителей.
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