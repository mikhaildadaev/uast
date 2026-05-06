---
outline: deep
---

# API / Ядро / Конструкторы

::: info **Информация**
На этой странице описано, как создать экземпляр телеметрии, настроить все параметры и разобраться в каждом типе данных и конструкторе полей.
:::

## NewSQL
SQL-инстанс, сконфигурированный под указанный диалект
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
