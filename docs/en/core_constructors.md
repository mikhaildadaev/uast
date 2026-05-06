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
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```

## NewInsert
Сreates a new INSERT statement instance
```go
stmtInsert := NewInsert(...)
```
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```

## NewSelect
Сreates a new SELECT statement instance
```go
stmtSelect := NewSelect(...)
```
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```

## NewUpdate
Сreates a new UPDATE statement instance
```go
stmtUpdate := NewUpdate(...)
```
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```
