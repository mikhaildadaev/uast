---
outline: deep
---

# API / SQL / Constructors

::: info **Info**
This page describes how to create a telemetry instance, configure all the settings, and understand each data type and field constructor.
:::

## NewSQL
SQL instance configured for the specified dialect
```go
sql := uast.NewSQL(
    WithDialect(uast.DialectMySQL),
)
defer sql.Close()
sql.SetDialect(uast.DialectPostgreSQL)
```
Output:
```text
...
```
