---
outline: deep
---

# API / Ядро / Конструкторы

::: info **Информация**
Эта страница охватывает четыре конструктора операторов: `NewDelete`, `NewInsert`, `NewSelect`, `NewUpdate`. Каждый конструктор создаёт новый экземпляр оператора, который может быть настроен с помощью методов и скомпилирован в SQL с помощью `Build()`.
:::

## NewDelete
Создаёт новый экземпляр оператора DELETE. Принимает источник таблицы и возвращает оператор, который может быть настроен с помощью `Join`, `Returning`, `Where`, `With`.
```go
statement := uast.NewDelete(uast.Table("test")).
    Where(uast.Equal(uast.Column[string]("test", "status"), uast.Value("active")))
```
Output MySQL:
```text
DELETE FROM `test` WHERE `test`.`status` = ?
```
Output PostgreSQL:
```text
DELETE FROM "test" WHERE "test"."status" = $1
```

## NewInsert
Создаёт новый экземпляр оператора INSERT. Принимает колонки и возвращает оператор, который может быть настроен с помощью `Into`, `Returning`, `Source`, `Values`, `With`.
```go
statement := uast.NewInsert(uast.Column[string]("test", "name"), uast.Column[int]("test", "age")).
    Into(uast.Table("test")).
    Values(uast.Value("ivan"), uast.Value(2))
```
Output MySQL:
```text
INSERT INTO `test` (`test`.`name`, `test`.`age`) VALUES (?, ?)
```
Output PostgreSQL:
```text
INSERT INTO "test" ("test"."name", "test"."age") VALUES ($1, $2)
```

## NewSelect
Создаёт новый экземпляр оператора SELECT. Принимает поля и возвращает оператор, который может быть настроен с помощью `Distinct`, `From`, `GroupBy`, `Having`, `Join`, `Limit`, `Offset`, `OrderBy`, `Unions`, `Where`, `With`.
```go
statement := uast.NewSelect(uast.Column[string]("test", "email")).
    From(uast.Table("test")).
    Where(uast.Equal(uast.Column[string]("test", "status"), uast.Value("active")))
```
Output MySQL:
```text
SELECT `test`.`email` FROM `test` WHERE `test`.`status` = ?
```
Output PostgreSQL:
```text
SELECT "test"."email" FROM "test" WHERE "test"."status" = $1
```

## NewUpdate
Создаёт новый экземпляр оператора UPDATE. Принимает источник таблицы и возвращает оператор, который может быть настроен с помощью `Join`, `Returning`, `Set`, `Where`, `With`.
```go
statement := uast.NewUpdate(uast.Table("test")).
    Set(uast.Set(uast.Column[string]("test", "status"), uast.Value("active"))).
    Where(uast.Equal(uast.Column[int]("test", "id"), uast.Value(2)))
```
Output MySQL:
```text
UPDATE `test` SET `test`.`status` = ? WHERE `test`.`id` = ?
```
Output PostgreSQL:
```text
UPDATE "test" SET "test"."status" = $1 WHERE "test"."id" = $2
```
