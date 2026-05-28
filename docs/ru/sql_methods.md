---
outline: deep
---

# API / SQL / Методы

::: info **Информация**
Эта страница документирует методы, доступные для выражений: `As` для назначения псевдонимов и `Over` для добавления оконных спецификаций. Каждый метод показан с рабочим примером кода и ожидаемым выводом SQL.
:::

## exprColumn
### As
Назначает псевдоним выражению-колонке.
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
