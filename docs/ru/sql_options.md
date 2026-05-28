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
`WithMutable` помечает построитель как изменяемый при создании. `SetMutable` переключает существующий построитель в изменяемый режим. В изменяемом режиме `Build()` изменяет исходный statement вместо клонирования, что повышает производительность для одноразовых запросов. `SetDialect` заблокирован для изменяемых построителей. После сборки statement в изменяемом режиме он модифицирован и не может быть безопасно переиспользован — последующие сборки дают неопределённый результат.
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