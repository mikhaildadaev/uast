---
outline: deep
---

# API / 核心 / 构造函数

::: info **关于**
本页文档介绍如何创建遥测实例、配置所有设置以及了解每个数据类型和字段构造函数。
:::

## NewDelete
创建一个新的 DELETE 语句实例
```go
stmtDelete := NewDelete(...)
```
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```

## NewInsert
创建一个新的 INSERT 语句实例
```go
stmtInsert := NewInsert(...)
```
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```

## NewSelect
创建一个新的 SELECT 语句实例
```go
stmtSelect := NewSelect(...)
```
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```

## NewUpdate
创建一个新的 UPDATE 语句实例
```go
stmtUpdate := NewUpdate(...)
```
Output:
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```
