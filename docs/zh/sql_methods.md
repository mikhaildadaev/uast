---
outline: deep
---

# API / SQL / 方法

::: info **关于**
本页面记录了 SQL 构建器实例上可用的快捷方法：`Exec`、`Query` 和 `QueryRow`。这些方法将 `Build()` 与相应的 `database/sql` 方法结合，减少样板代码。每个方法都配有可运行的代码示例和预期行为。
:::

## sqlBuilder
### Exec
构建语句并通过 `db.Exec()` 执行。返回 `sql.Result` 和错误。适用于不返回行的 INSERT、UPDATE、DELETE 语句。
```go
stmt := uast.NewInsert(uast.NewTable("test").As("t")).
    Values(
        uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
    )
builder := uast.NewSQL(uast.WithDialect(uast.DialectPostgreSQL))
defer builder.Close()
db, _ := sql.Open("postgres", "postgres://user:pass@localhost/db")
defer db.Close()
result, err := builder.Exec(stmt, db)
if err != nil {
    log.Fatal(err)
}
rowsAffected, _ := result.RowsAffected()
```
Output:
```text
// Executes: INSERT INTO "test" AS "t" ("string") VALUES ($1)
// Returns: sql.Result with LastInsertId and RowsAffected
```

### Query
构建语句并通过 `db.Query()` 执行。返回 `*sql.Rows` 和错误。适用于返回多行的 SELECT 语句。
```go
stmt := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[int64]("t", "id"),
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    )
builder := uast.NewSQL(uast.WithDialect(uast.DialectPostgreSQL))
defer sbuilderql.Close()
db, _ := sql.Open("postgres", "postgres://user:pass@localhost/db")
defer db.Close()
rows, err := builder.Query(stmt, db)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
    var id int64
    var str string
    rows.Scan(&id, &str)
    fmt.Printf("id: %d, string: %s\n", id, str)
}
```
Output:
```text
// Executes: SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = $1
// Returns: *sql.Rows iterator
```

### QueryRow
构建语句并通过 `db.QueryRow()` 执行。返回 `*sql.Row` 和来自 `Build()` 的错误。适用于返回单行的 SELECT 语句。
```go
stmt := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[int64]("t", "id"),
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[int64]("t", "id"), uast.Value(1)),
    )
builder := uast.NewSQL(uast.WithDialect(uast.DialectPostgreSQL))
defer builder.Close()
db, _ := sql.Open("postgres", "postgres://user:pass@localhost/db")
defer db.Close()
row, err := builder.QueryRow(stmt, db)
if err != nil {
    log.Fatal(err)
}
var id int64
var str string
err = row.Scan(&id, &str)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("id: %d, string: %s\n", id, str)
```
Output:
```text
// Executes: SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."id" = $1
// Returns: *sql.Row, scanned via row.Scan()
```

::: tip 注
这些方法是结合 `Build()` 与 `database/sql` 执行的快捷方式。如需完全控制，请直接使用 `Build()` 并通过自己的 `*sql.DB` 或 `*sql.Tx` 执行。`QueryRow` 返回 `*sql.Row` — 请始终调用 `Scan()` 获取结果。如果 `Build()` 失败，错误会在任何数据库调用之前返回。
:::