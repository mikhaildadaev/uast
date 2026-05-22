---
outline: deep
---

# API / SQL / Constructors

::: info **Info**
This page describes how to create a sql instance, configure all the settings, and understand each data type and field constructor.
:::

## NewSQL
SQL instance configured for the specified dialect
```go
sql, err := uast.NewSQL()
defer sql.Close()
```
Output:
```text
...
```
