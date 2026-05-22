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
sql := uast.NewSQL(
    uast.WithDialect(uast.DialectMariaDB)
)
defer sql.Close()
mariadbQuery, mariadbArgs, _ := sql.Build(stmt)
sql.SetDialect(uast.DialectMySQL)
mysqlQuery, mysqlArgs, _ := sql.Build(stmt)
sql.SetDialect(uast.DialectPostgreSQL)
pgsqlQuery, pgsqlArgs, _ := sql.Build(stmt)
sql.SetDialect(uast.DialectSQLite)
sqliteQuery, sqliteArgs, _ := sql.Build(stmt)
```
Output MariaDB:
```text
SELECT `t`.`string` FROM `test` AS `t` WHERE `t`.`id` = ?
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