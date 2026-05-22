---
outline: deep
---

# API / 核心 / 构造函数

::: info **关于**
本页文档介绍如何创建遥测实例、配置所有设置以及了解每个数据类型和字段构造函数。
:::

## NewSQL
为指定方言配置的 SQL 实例
```go
sql, err := uast.NewSQL()
defer sql.Close()
```
Output:
```text
...
```
