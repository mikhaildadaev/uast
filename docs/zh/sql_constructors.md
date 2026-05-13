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
stmtDelete := uast.NewDelete(...)
stmtInsert := uast.NewInsert(...)
stmtSelect := uast.NewSelect(...)
stmtUpdate := uast.NewUpdate(...)
sql, err := uast.NewSQL(uast.DialectMySQL)
defer sql.Close()
mySQLDeleteQuery, mySQLDeleteArguments, err := sql.Build(stmtDelete)
if err != nil {
    return fmt.Errorf("build delete: %w", err)
}
mySQLInsertQuery, mySQLInsertArguments, err := sql.Build(stmtInsert)
if err != nil {
    return fmt.Errorf("build insert: %w", err)
}
sql.SetDialect(uast.DialectPostgreSQL)
pgSQLSelectQuery, pgSQLSelectArguments, err := sql.Build(stmtSelect)
if err != nil {
    return fmt.Errorf("build select: %w", err)
}
pgSQLUpdateQuery, pgSQLUpdateArguments, err := sql.Build(stmtUpdate)
if err != nil {
    return fmt.Errorf("build update: %w", err)
}
```
Output MySQL:
```text
...
```
Output PostgreSQL:
```text
...
```
