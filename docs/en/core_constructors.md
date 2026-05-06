---
outline: deep
---

# API / Core / Constructors

::: info **Info**
This page describes how to create a telemetry instance, configure all the settings, and understand each data type and field constructor.
:::

## NewDelete
Сreates a new DELETE statement instance
```go
stmtDelete := NewDelete(...)
```

## NewInsert
Сreates a new INSERT statement instance
```go
stmtInsert := NewInsert(...)
```

## NewSelect
Сreates a new SELECT statement instance
```go
stmtSelect := NewSelect(...)
```

## NewUpdate
Сreates a new UPDATE statement instance
```go
stmtUpdate := NewUpdate(...)
```

## NewSQL
SQL instance configured for the specified dialect
```go
mySQL, err := uast.NewSQL(uast.DialectMySQL)
if err != nil {
    return fmt.Errorf("mysql builder: %w", err)
}
defer builderMySQL.Close()
postgreSQL, err := uast.NewSQL(uast.DialectPostgreSQL)
if err != nil {
    return fmt.Errorf("postgresql builder: %w", err)
}
defer builderPostgreSQL.Close()
stmtDelete := NewDelete(...)
stmtInsert := NewInsert(...)
stmtSelect := NewSelect(...)
stmtUpdate := NewUpdate(...)
mySQLDeleteQuery, mySQLDeleteArguments, err := builderMySQL.Build(stmtDelete)
if err != nil {
    return fmt.Errorf("build delete: %w", err)
}
mySQLInsertQuery, mySQLInsertArguments, err := builderMySQL.Build(stmtInsert)
if err != nil {
    return fmt.Errorf("build insert: %w", err)
}
postgreSQLSelectQuery, postgreSQLSelectArguments, err := builderPostgreSQL.Build(stmtSelect)
if err != nil {
    return fmt.Errorf("build select: %w", err)
}
postgreSQLUpdateQuery, postgreSQLUpdateArguments, err := builderPostgreSQL.Build(stmtUpdate)
if err != nil {
    return fmt.Errorf("build update: %w", err)
}
```
Output:
```json
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
{"level":"info","type":"metric","name":"payments","value":99.99,"node_id":"123-abc","trace_id":"abc-123"}
{"level":"info","type":"trace","name":"payment_processing","duration":150,"span_id":"span-456","node_id":"123-abc","trace_id":"abc-123"}
```

| Name                                                            | Description                                                                     | Values                                                             | Default      |
|-----------------------------------------------------------------|---------------------------------------------------------------------------------|--------------------------------------------------------------------|--------------|
| [`WithExtractor()`](/en/core_options#withextractor-setextractor)| Auto-extract fields from `context.Context` by key names                         | `keys ...string`                                                   |              |
| [`WithFormat()`](/en/core_options#withformat-setformat)         | Output format: structured JSON or human-readable TEXT with optional ANSI colors | `FormatJson`, `FormatText`                                         | `FormatJson` |
| [`WithLevel()`](/en/core_options#withlevel-setlevel)            | Minimum log severity. Only messages at or above this level are written          | `LevelDebug`, `LevelError`, `LevelFatal`, `LevelInfo`, `LevelWarn` | `LevelInfo`  |
| [`WithMode()`](/en/core_options#withmode-setmode)               | Write mode: non-blocking `ModeAsync` with buffer or blocking `ModeSync`         | `ModeAsync`, `ModeSync`                                            | `ModeSync`   |
| [`WithTheme()`](/en/core_options#withtheme-settheme)            | ANSI color theme for TEXT output: optimized for dark or light terminals         | `ThemeDark`, `ThemeLight`                                          | `ThemeDark`  |

| Name                                | Description                                    | Values                                                                                                                                                     |
|-------------------------------------|------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [`TypeData`](/en/core_types#data)   | Log messages, Prometheus metrics, Tempo traces | `DataLog`, `DataMetric`, `DataTrace`                                                                                                                       |
| [`TypeField`](/en/core_types#field) | 16 type-safe field constructors                | `Bool`, `Bools`, `Duration`, `Durations`, `Error`, `Errors`, `Float64`, `Floats64`, `Int`, `Ints`, `Int64`, `Ints64`, `String`, `Strings`, `Time`, `Times` |
