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
```text
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
```
