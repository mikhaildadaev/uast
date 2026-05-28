---
outline: deep
---

# API / SQL / 方法

::: info **关于**
本页面记录了可用于表达式的方法： `As` 用于分配别名，`Over` 用于添加窗口规范。每个方法都配有可运行的代码示例和预期 SQL 输出。
:::

## exprColumn
### As
为列表达式分配别名。
```go
column := uast.Column[string]("t", "string").As("alias")
```
Output MariaDB:
```text
`t`.`string` AS `alias`
```
Output MsSQL:
```text
[t].[string] AS [alias]
```
Output MySQL:
```text
`t`.`string` AS `alias`
```
Output PostgreSQL:
```text
"t"."string" AS "alias"
```
Output SQLite:
```text
"t"."string" AS "alias"
```

