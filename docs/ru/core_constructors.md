---
outline: deep
---

# API / Ядро / Конструкторы

::: info **Информация**
На этой странице описано, как создать экземпляр телеметрии, настроить все параметры и разобраться в каждом типе данных и конструкторе полей.
:::

## NewDelete
Создаёт новый экземпляр DELETE-оператора
```go
stmtDelete := NewDelete(...)
```
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```

## NewInsert
Создаёт новый экземпляр INSERT-оператора
```go
stmtInsert := NewInsert(...)
```
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```

## NewSelect
Создаёт новый экземпляр SELECT-оператора
```go
stmtSelect := NewSelect(...)
```
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```

## NewUpdate
Создаёт новый экземпляр UPDATE-оператора
```go
stmtUpdate := NewUpdate(...)
```
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```
