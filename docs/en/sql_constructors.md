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
Output:
```text
...
```
