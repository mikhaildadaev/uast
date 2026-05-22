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
Output MariaDB:
```text
DELETE FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output MySQL:
```text
DELETE FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = $1
```
Output SQLite:
```text
DELETE FROM "test" AS "t" WHERE "t"."string" = ?
```

## NewInsert
Создаёт новый экземпляр оператора INSERT. Принимает источник таблицы и возвращает оператор, который может быть настроен с помощью `Returning`, `Source/Values`, `With`.
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
Output MariaDB:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
```
Output MySQL:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
```
Output PostgreSQL:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES ($1, $2)
```
Output SQLite:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES (?, ?)
```

## NewSelect
Создаёт новый экземпляр оператора SELECT. Принимает источник таблицы и возвращает оператор, который может быть настроен с помощью `Distinct`, `GroupBy`, `Having`, `Join`, `Limit`, `Offset`, `OrderBy`, `Unions`, `Where`, `With`.
```go
statement := uast.NewSelect(uast.Table("test")).
    Field(
        uast.Column[string]("test", "string"),
    ).
    Where(
        uast.Equal(uast.Column[string]("test", "string"), uast.Value("active")),
    )
```
Output MariaDB:
```text
SELECT `t`.`string` FROM `test` AS `t`
```
Output MySQL:
```text
SELECT `t`.`string` FROM `test` AS `t`
```
Output PostgreSQL:
```text
SELECT "t"."string" FROM "test" AS "t"
```
Output SQLite:
```text
SELECT "t"."string" FROM "test" AS "t"
```

## NewUpdate
Создаёт новый экземпляр оператора UPDATE. Принимает источник таблицы и возвращает оператор, который может быть настроен с помощью `Join`, `Returning`, `Set`, `Where`, `With`.
```go
statement := uast.NewUpdate(uast.Table("test")).
    Set(
        Pair(uast.Column[string]("test", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Column[int]("test", "number"), uast.Value(2)),
    )
```
Output MariaDB:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
```
Output MySQL:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" SET "t"."string" = $1 WHERE "t"."number" = $2
```
Output SQLite:
```text
UPDATE "test" AS "t" SET "t"."string" = ? WHERE "t"."number" = ?
```
