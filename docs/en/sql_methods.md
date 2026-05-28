---
outline: deep
---

# API / SQL / Methods

::: info **Info**
This page documents methods available on expressions: `As` for assigning aliases and `Over` for adding window specifications. Each method is shown with a working code example and expected SQL output.
:::

## exprColumn
### As
Assigns an alias to a column expression.
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
