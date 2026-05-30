---
outline: deep
---

# API / Ядро / Конструкторы

::: info **Информация**
Эта страница охватывает четыре конструктора операторов: `NewDelete`, `NewInsert`, `NewSelect`, `NewUpdate`. Каждый конструктор создаёт новый экземпляр оператора, который может быть настроен с помощью методов и скомпилирован в SQL с помощью `Build()`.
:::

## NewComment
Создаёт новый экземпляр оператора COMMENT. Принимает текст комментария и возвращает оператор, который можно настроить с помощью `OnColumn` или `OnTable`.
```go
stmtCommentColumn := uast.NewComment("Test comment").
    OnColumn(uast.Column[int64]("t", "id"))
stmtCommentTable := uast.NewComment("Test comment").
    OnTable(uast.NewTable("test", "t"))
```
Output MariaDB:
```text
COMMENT ON COLUMN `t`.`id` IS 'Test comment'
COMMENT ON TABLE `test` AS `t` IS 'Test comment'
```
Output MsSQL:
```text
// Not supported
// Not supported
```
Output MySQL:
```text
COMMENT ON COLUMN `t`.`id` IS 'Test comment'
COMMENT ON TABLE `test` AS `t` IS 'Test comment'
```
Output PostgreSQL:
```text
COMMENT ON COLUMN "t"."id" IS 'Test comment'
COMMENT ON TABLE "test" AS "t" IS 'Test comment'
```
Output SQLite:
```text
COMMENT ON COLUMN "t"."id" IS 'Test comment'
COMMENT ON TABLE "test" AS "t" IS 'Test comment'
```

## NewDelete
Создаёт новый экземпляр оператора DELETE. Принимает источник таблицы и возвращает оператор, который может быть настроен с помощью `Join`, `Returning`, `Where`, `With`.
```go
stmtDeleteDefault := uast.NewDelete(uast.NewTable("test", "t")).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    )
```
Output MariaDB:
```text
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output MsSQL:
```text
DELETE [t] FROM [test] AS [t] WHERE [t].[string] = @p1
```
Output MySQL:
```text
DELETE `t` FROM `test` AS `t` WHERE `t`.`string` = ?
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
stmtInsertDefault := uast.NewInsert(uast.NewTable("test", "t")).
    Values(
		uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
		uast.Pair(uast.Column[int]("t", "number"), uast.Value(2)),
	)
stmtInsertSource := uast.NewInsert(uast.NewTable("test", "t")).
	Source(NewSelect(uast.NewTable("test", "t")).
		Field(
			uast.Column[string]("t", "string"),
			uast.Column[int]("t", "number"),
		).
		Where(
			uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
		),
	)
```
Output MariaDB:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
INSERT INTO `test` AS `t` (`string`, `number`) SELECT `t`.`string`, `t`.`number` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output MsSQL:
```text
INSERT INTO [test] AS [t] ([string], [number]) VALUES (@p1, @p2)
INSERT INTO [test] AS [t] ([string], [number]) SELECT [t].[string], [t].[number] FROM [test] AS [t] WHERE [t].[string] = @p1
```
Output MySQL:
```text
INSERT INTO `test` AS `t` (`string`, `number`) VALUES (?, ?)
INSERT INTO `test` AS `t` (`string`, `number`) SELECT `t`.`string`, `t`.`number` FROM `test` AS `t` WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES ($1, $2)
INSERT INTO "test" AS "t" ("string", "number") SELECT "t"."string", "t"."number" FROM "test" AS "t" WHERE "t"."string" = $1
```
Output SQLite:
```text
INSERT INTO "test" AS "t" ("string", "number") VALUES (?, ?)
INSERT INTO "test" AS "t" ("string", "number") SELECT "t"."string", "t"."number" FROM "test" AS "t" WHERE "t"."string" = ?
```

## NewSelect
Создаёт новый экземпляр оператора SELECT. Принимает источник таблицы и возвращает оператор, который может быть настроен с помощью `Distinct`, `GroupBy`, `Having`, `Join`, `OrderBy`, `Pagination`, `Unions`, `Where`, `With`.
```go
stmtSelectDefault := uast.NewSelect(uast.NewTable("test", "t")).
    Field(
        uast.Column[int64]("t", "id"),
    ).
    Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	)
stmtSelectDistinct := uast.NewSelect(uast.NewTable("test", "t")).
    Distinct().
    Field(
        uast.Column[int64]("t", "id"),
    ).
    Where(
		uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
	)
```
Output MariaDB:
```text
SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
SELECT DISTINCT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
```
Output MsSQL:
```text
SELECT [t].[id] FROM [test] AS [t] WHERE [t].[number] = @p1
SELECT DISTINCT [t].[id] FROM [test] AS [t] WHERE [t].[number] = @p1
```
Output MySQL:
```text
SELECT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
SELECT DISTINCT `t`.`id` FROM `test` AS `t` WHERE `t`.`number` = ?
```
Output PostgreSQL:
```text
SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = $1
SELECT DISTINCT "t"."id" FROM "test" AS "t" WHERE "t"."number" = $1
```
Output SQLite:
```text
SELECT "t"."id" FROM "test" AS "t" WHERE "t"."number" = ?
SELECT DISTINCT "t"."id" FROM "test" AS "t" WHERE "t"."number" = ?
```

## NewTruncate
Создаёт новый экземпляр оператора TRUNCATE. Принимает источник таблицы и возвращает оператор, который можно настроить с помощью `Cascade()` или `RestartIdentity()`.
```go
stmtTruncateDefault := uast.NewTruncate(uast.NewTable("test", "t"))
stmtTruncateCascade := uast.NewTruncate(uast.NewTable("test", "t")).
    Cascade()
stmtTruncateRestartIdentity := uast.NewTruncate(uast.NewTable("test", "t")).
    RestartIdentity()
```
Output MariaDB:
```text
TRUNCATE TABLE `test`
TRUNCATE TABLE `test` CASCADE
TRUNCATE TABLE `test` RESTART IDENTITY
```
Output MsSQL:
```text
TRUNCATE TABLE [test]
// Not supported
// Not supported
```
Output MySQL:
```text
TRUNCATE TABLE `test`
// Not supported
// Not supported
```
Output PostgreSQL:
```text
TRUNCATE TABLE "test"
TRUNCATE TABLE "test" CASCADE
TRUNCATE TABLE "test" RESTART IDENTITY

```
Output SQLite:
```text
TRUNCATE TABLE "test"
// Not supported
// Not supported
```

## NewUpdate
Создаёт новый экземпляр оператора UPDATE. Принимает источник таблицы и возвращает оператор, который может быть настроен с помощью `Join`, `Returning`, `Set`, `Where`, `With`.
```go
stmtUpdateDefault := uast.NewUpdate(uast.NewTable("test", "t")).
    Set(
        Pair(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Where(
        uast.Equal(uast.Column[int]("t", "number"), uast.Value(2)),
    )
```
Output MariaDB:
```text
UPDATE `test` AS `t` SET `t`.`string` = ? WHERE `t`.`number` = ?
```
Output MsSQL:
```text
UPDATE [test] AS [t] SET [t].[string] = @p1 WHERE [t].[number] = @p2
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
