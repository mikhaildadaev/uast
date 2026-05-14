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
    Where(
        uast.Equal(uast.Column[string]("test", "string"), uast.Value("active")),
    )
```
Output MySQL:
```text
DELETE FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = $1
```

## NewInsert
Создаёт новый экземпляр оператора INSERT. Принимает колонки и возвращает оператор, который может быть настроен с помощью `Into`, `Returning`, `Source`, `Values`, `With`.
```go
statement := uast.NewInsert(uast.Table("test")).
    Column(
        uast.Column[string]("test", "string"), 
        uast.Column[int]("test", "number"),
    ).
    Values(
        uast.Row(
            uast.Value("ivan"), 
            uast.Value(2),
        ),
    )
```
Output MySQL:
```text
INSERT INTO `test` (`test`.`string`, `test`.`number`) VALUES (?, ?)
```
Output PostgreSQL:
```text
INSERT INTO "test" ("test"."string", "test"."number") VALUES ($1, $2)
```

## NewSelect
Создаёт новый экземпляр оператора SELECT. Принимает поля и возвращает оператор, который может быть настроен с помощью `Distinct`, `From`, `GroupBy`, `Having`, `Join`, `Limit`, `Offset`, `OrderBy`, `Unions`, `Where`, `With`.
```go
statement := uast.NewSelect(uast.Table("test")).
    Field(
        uast.Column[string]("test", "string"),
    ).
    Where(
        uast.Equal(uast.Column[string]("test", "string"), uast.Value("active")),
    )
```
Output MySQL:
```text
SELECT `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
SELECT "t"."string" FROM "test" AS "t" WHERE "t"."string" = $1
```

## NewUpdate
Создаёт новый экземпляр оператора UPDATE. Принимает источник таблицы и возвращает оператор, который может быть настроен с помощью `Join`, `Returning`, `Set`, `Where`, `With`.
```go
statement := uast.NewUpdate(uast.Table("test")).
    Set(
        Assign(
            uast.Column[string]("test", "string"), 
            uast.Value("active"),
        ),
    ).
    Where(
        uast.Equal(uast.Column[int]("test", "number"), uast.Value(2)),
    )
```
Output MySQL:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2
```

