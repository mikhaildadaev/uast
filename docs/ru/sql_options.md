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