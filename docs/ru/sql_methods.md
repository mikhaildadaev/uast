---
outline: deep
---

# API / SQL / Методы

::: info **Информация**
Эта страница документирует сокращённые методы, доступные для экземпляра SQL-построителя: `Exec`, `Query` и `QueryRow`. Эти методы объединяют `Build()` с соответствующими методами `database/sql`, сокращая шаблонный код. Каждый метод показан с рабочим примером кода и ожидаемым поведением.
:::

## sqlBuilder
### Build
Компилирует оператор в SQL-строку и список аргументов. Это основной метод UAST — он принимает AST-оператор и возвращает финальный SQL-запрос, готовый к выполнению.
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
Собирает оператор и выполняет его через `db.Exec()`. Возвращает `sql.Result` и ошибку. Подходит для операторов INSERT, UPDATE, DELETE, которые не возвращают строки.
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
Собирает оператор и выполняет его через `db.Query()`. Возвращает `*sql.Rows` и ошибку. Подходит для операторов SELECT, которые возвращают несколько строк.
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
Собирает оператор и выполняет его через `db.QueryRow()`. Возвращает `*sql.Row` и ошибку от `Build()`. Подходит для операторов SELECT, которые возвращают одну строку.
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

::: tip Примечание
Эти методы — сокращённый путь, объединяющий `Build()` с выполнением через `database/sql`. Для полного контроля используйте `Build()` напрямую и выполняйте запрос с собственным `*sql.DB` или `*sql.Tx`. `QueryRow` возвращает `*sql.Row` — всегда вызывайте `Scan()` для получения результата. Если `Build()` завершается с ошибкой, она возвращается до выполнения запроса к базе данных.
:::