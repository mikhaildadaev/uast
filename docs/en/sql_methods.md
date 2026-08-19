---
outline: deep
---

# API / SQL / Methods

::: info **Info**
This page documents shortcut methods available on the SQL builder instance: `Exec`, `Query`, and `QueryRow`. These methods combine `Build()` with the corresponding `database/sql` methods, reducing boilerplate code. Each method is shown with a working code example and expected behavior.
:::

## sqlBuilder
### Build
Compiles a statement into a SQL string and a list of arguments. This is the core method of UAST — it takes an AST statement and returns the final SQL query ready for execution.
```go
stmt := uast.NewSelect(uast.NewTable("users", "u")).
    Fields(
        uast.Field[int64]("u", "id"),
        uast.Field[string]("u", "name"),
    ).
    Where(
        uast.Equal(uast.Field[string]("u", "status"), uast.Value("active")),
    )
query, args, err := builder.Build(stmt)
```
Output:
```text
// Executes: SELECT "u"."id", "u"."name" FROM "users" AS "u" WHERE "u"."status" = $1
// Returns: [active]
```

### Exec
Builds the statement and executes it via `db.Exec()`. Returns `sql.Result` and any error. Suitable for INSERT, UPDATE, DELETE statements that do not return rows.
```go
stmt := uast.NewInsert(uast.NewTable("users").As("u")).
    Values(
        uast.Pair(uast.Field[string]("u", "string"), uast.Value("ivan")),
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
// Executes: INSERT INTO "users" AS "u" ("string") VALUES ($1)
// Returns: sql.Result with LastInsertId and RowsAffected
```

### Query
Builds the statement and executes it via `db.Query()`. Returns `*sql.Rows` and any error. Suitable for SELECT statements that return multiple rows.
```go
stmt := uast.NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[int64]("u", "id"),
        uast.Field[string]("u", "string"),
    ).
    Where(
        uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    )
builder := uast.NewSQL(uast.WithDialect(uast.DialectPostgreSQL))
defer builder.Close()
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
// Executes: SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = $1
// Returns: *sql.Rows iterator
```

### QueryRow
Builds the statement and executes it via `db.QueryRow()`. Returns `*sql.Row` and any error from `Build()`. Suitable for SELECT statements that return a single row.
```go
stmt := uast.NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[int64]("u", "id"),
        uast.Field[string]("u", "string"),
    ).
    Where(
        uast.Equal(uast.Field[int64]("u", "id"), uast.Value(1)),
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
// Executes: SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."id" = $1
// Returns: *sql.Row, scanned via row.Scan()
```

::: tip Note
These methods are shortcuts that combine `Build()` with d`atabase/sql` execution. For full control, use `Build()` directly and execute with your own `*sql.DB` or `*sql.Tx`. `QueryRow` returns `*sql.Row` — always call `Scan()` on the result. If `Build()` fails, the error is returned before any database call is made.
:::