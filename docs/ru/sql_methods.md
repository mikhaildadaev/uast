---
outline: deep
---

# API / SQL / Методы

::: info **Информация**
Эта страница документирует сокращённые методы, доступные для экземпляра SQL-построителя: `Exec`, `Query` и `QueryRow`. Эти методы объединяют `Build()` с соответствующими методами `database/sql`, сокращая шаблонный код. Каждый метод показан с рабочим примером кода и ожидаемым поведением.
:::

## sqlBuilder
### Exec
Собирает оператор и выполняет его через `db.Exec()`. Возвращает `sql.Result` и ошибку. Подходит для операторов INSERT, UPDATE, DELETE, которые не возвращают строки.
```go
sql := uast.NewSQL(uast.WithDialect(uast.DialectPostgreSQL))
defer sql.Close()
db, _ := sql.Open("postgres", "postgres://user:pass@localhost/db")
defer db.Close()
stmt := uast.NewInsert(uast.NewTable("test").As("t")).
    Values(
        uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
    )
result, err := sql.Exec(stmt, db)
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
Builds the statement and executes it via `db.Exec()`. Returns `sql.Result` and any error. Suitable for INSERT, UPDATE, DELETE statements that do not return rows.
```go
sql := uast.NewSQL(uast.WithDialect(uast.DialectPostgreSQL))
defer sql.Close()
db, _ := sql.Open("postgres", "postgres://user:pass@localhost/db")
defer db.Close()
stmt := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[int64]("t", "id"),
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    )
rows, err := sql.Query(stmt, db)
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
Собирает оператор и выполняет его через `db.QueryRow()`. Возвращает `*sql.Row` и ошибку от `Build()`. Подходит для операторов SELECT, которые возвращают одну строку.
```go
sql := uast.NewSQL(uast.WithDialect(uast.DialectPostgreSQL))
defer sql.Close()
db, _ := sql.Open("postgres", "postgres://user:pass@localhost/db")
defer db.Close()
stmt := uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[int64]("t", "id"),
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[int64]("t", "id"), uast.Value(1)),
    )
row, err := sql.QueryRow(stmt, db)
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

::: tip Примечание
Эти методы — сокращённый путь, объединяющий `Build()` с выполнением через `database/sql`. Для полного контроля используйте `Build()` напрямую и выполняйте запрос с собственным `*sql.DB` или `*sql.Tx`. `QueryRow` возвращает `*sql.Row` — всегда вызывайте `Scan()` для получения результата. Если `Build()` завершается с ошибкой, она возвращается до выполнения запроса к базе данных.
:::