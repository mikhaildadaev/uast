---
outline: deep
---

# API / Ядро / Опции

::: info **Информация**
Эта страница охватывает все параметры конфигурации: `clauseGroupBy`, `clauseHaving`, `clauseJoin`, `clauseOrderBy`, `clausePagination`, `clauseReturning`, `clauseSet`, `clauseUnions`, `clauseValues`, `clauseWhere`, `clauseWith`, `exprArray`, `exprBinary`, `exprColumn`, `exprComparison`, `exprConstant`, `exprFunction`, `exprLiteral`, `exprLogical`, `exprSubquery`, `exprValue`. Каждый параметр показан с рабочим примером кода и ожидаемым выводом.
:::

## clauseGroupBy
Добавляет оператор GROUP BY для группировки строк по указанным колонкам или выражениям.
```go
groupBy := GroupBy(
	uast.Column[string]("t", "string"),
)
```
Output MariaDB:
```text
GROUP BY `t`.`string`
```
Output MsSQL:
```text
GROUP BY [t].[string]
```
Output MySQL:
```text
GROUP BY `t`.`string`
```
Output PostgreSQL:
```text
GROUP BY "t"."string"
```
Output SQLite:
```text
GROUP BY "t"."string"
```

## clauseHaving
Добавляет оператор HAVING для фильтрации групп. Используется с GROUP BY для фильтрации агрегированных результатов.
```go
having := Having(
	uast.Greater(uast.Count(uast.Column[int64]("t", "id"), false), uast.Value[int64](2)),
)
```
Output MariaDB:
```text
HAVING COUNT(`t`.`id`) > ?
```
Output MsSQL:
```text
HAVING COUNT([t].[id]) > @p1
```
Output MySQL:
```text
HAVING COUNT(`t`.`id`) > ?
```
Output PostgreSQL:
```text
HAVING COUNT("t"."id") > $1
```
Output SQLite:
```text
HAVING COUNT("t"."id") > ?
```

## clauseJoin
### Cross
Добавляет CROSS JOIN к запросу. Возвращает декартово произведение обеих таблиц.
```go
join := uast.Cross(uast.NewTable("test").As("t"))
```
Output MariaDB:
```text
CROSS JOIN `test` AS `t`
```
Output MsSQL:
```text
CROSS JOIN [test] AS [t]
```
Output MySQL:
```text
CROSS JOIN `test` AS `t`
```
Output PostgreSQL:
```text
CROSS JOIN "test" AS "t"
```
Output SQLite:
```text
CROSS JOIN "test" AS "t"
```

### Full
Добавляет FULL JOIN к запросу. Возвращает все строки из обеих таблиц, с NULL там, где нет совпадений.
```go
join := uast.Full(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
FULL JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
FULL JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
FULL JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
FULL JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
FULL JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```

### FullOuter
Добавляет FULL OUTER JOIN к запросу. Возвращает все строки из обеих таблиц, с NULL там, где нет совпадений.
```go
join := uast.FullOuter(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
FULL OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
FULL OUTER JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
FULL OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
FULL OUTER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
FULL OUTER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```

### Inner
Добавляет INNER JOIN к запросу. Возвращает строки, имеющие совпадающие значения в обеих таблицах.
```go
join := uast.Inner(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
INNER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
INNER JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
INNER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
INNER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
INNER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```

### Left
Добавляет LEFT JOIN к запросу. Возвращает все строки из левой таблицы и совпадающие строки из правой таблицы.
```go
join := uast.Left(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
LEFT JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
LEFT JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
LEFT JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
LEFT JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
LEFT JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```

### LeftOuter
Добавляет LEFT OUTER JOIN к запросу. Возвращает все строки из левой таблицы и совпадающие строки из правой таблицы.
```go
join := uast.LeftOuter(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
LEFT OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
LEFT OUTER JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
LEFT OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
LEFT OUTER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
LEFT OUTER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```

### Right
Добавляет RIGHT JOIN к запросу. Возвращает все строки из правой таблицы и совпадающие строки из левой таблицы. Не поддерживается SQLite.
```go
join := uast.Right(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
RIGHT JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
RIGHT JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
RIGHT JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
RIGHT JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
// Not supported
```

### RightOuter
Добавляет RIGHT OUTER JOIN к запросу. Возвращает все строки из правой таблицы и совпадающие строки из левой таблицы. Не поддерживается SQLite.
```go
join := uast.RightOuter(uast.NewTable("test").As("t"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("t1", "id")))
```
Output MariaDB:
```text
RIGHT OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output MsSQL:
```text
RIGHT OUTER JOIN [test] AS [t] ON [t].[id] = [t1].[id]
```
Output MySQL:
```text
RIGHT OUTER JOIN `test` AS `t` ON `t`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
RIGHT OUTER JOIN "test" AS "t" ON "t"."id" = "t1"."id"
```
Output SQLite:
```text
// Not supported
```

## clauseOrderBy
### Asc
Указывает порядок сортировки по возрастанию (сначала наименьшие, от А до Я). Используется для сортировки строк в запросе или в рамках оконной функции.
```go
orderBy := uast.Asc(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` ASC
```
Output MsSQL:
```text
[t].[string] ASC
```
Output MySQL:
```text
`t`.`string` ASC
```
Output PostgreSQL:
```text
"t"."string" ASC
```
Output SQLite:
```text
"t"."string" ASC
```

### Desc
Указывает порядок сортировки по убыванию (сначала наибольшие, от Я до А). Используется для сортировки строк в запросе или в рамках оконной функции.
```go
orderBy := uast.Desc(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` DESC
```
Output MsSQL:
```text
[t].[string] DESC
```
Output MySQL:
```text
`t`.`string` DESC
```
Output PostgreSQL:
```text
"t"."string" DESC
```
Output SQLite:
```text
"t"."string" DESC
```

## clausePagination
Определяет пагинацию для оператора SELECT с помощью `Pagination(limit, offset)`. `limit` задаёт максимальное количество возвращаемых строк. `offset` указывает количество строк, которые нужно пропустить перед возвратом результатов. Порядок отрисовки и синтаксис автоматически адаптируются к каждому диалекту.
```go
pagination := Pagination(10,0)
```
Output MariaDB:
```text
LIMIT ? OFFSET ?
```
Output MsSQL:
```text
OFFSET @p1 ROWS FETCH NEXT @p2 ROWS ONLY
```
Output MySQL:
```text
LIMIT ? OFFSET ?
```
Output PostgreSQL:
```text
LIMIT $1 OFFSET $2
```
Output SQLite:
```text
LIMIT ? OFFSET ?
```

## clauseReturning
Добавляет оператор RETURNING для возврата изменённых строк. Поддерживается MariaDB, PostgreSQL и SQLite. MySQL не поддерживает этот оператор.
```go
returning = Returning(
	uast.Column[int64]("t", "id")
    uast.Column[string]("t", "string")
)
```
Output MariaDB:
```text
RETURNING `t`.`id`, `t`.`string`
```
Output MsSQL:
```text
OUTPUT [t].[id], [t].[string]
```
Output MySQL:
```text
// Not support
```
Output PostgreSQL:
```text
RETURNING "t"."id", "t"."string"
```
Output SQLite:
```text
RETURNING "t"."id", "t"."string"
```

## clauseSet
### Assign
Указывает колонки и их новые значения с помощью `Assign` для связывания колонок со значениями. Поддерживает несколько пар для обновления нескольких колонок.
```go
set := Set(
	uast.Assign(uast.Column[string]("t", "string"), uast.Value("active")),
)
```
Output MariaDB:
```text
UPDATE `test` AS `t` SET `t`.`string` = ?
```
Output MsSQL:
```text
UPDATE [test] AS [t] SET [t].[string] = @p1
```
Output MySQL:
```text
UPDATE `test` AS `t` SET `t`.`string` = ?
```
Output PostgreSQL:
```text
UPDATE "test" AS "t" SET "t"."string" = $1
```
Output SQLite:
```text
UPDATE "test" AS "t" SET "t"."string" = ?
```

## clauseUnions
### Union
Объединяет результаты нескольких операторов SELECT. UNION возвращает уникальные строки.
```go
unions := uast.Union(uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ),
)
```
Output MariaDB:
```text
UNION SELECT `t`.`string` FROM `test` AS `t` 
```
Output MsSQL:
```text
UNION SELECT [t].[string] FROM [test] AS [t]
```
Output MySQL:
```text
UNION SELECT `t`.`string` FROM `test` AS `t`
```
Output PostgreSQL:
```text
UNION SELECT "t"."string" FROM "test" AS "t"
```
Output SQLite:
```text
UNION SELECT "t"."string" FROM "test" AS "t"
```

### UnionAll
Объединяет результаты нескольких операторов SELECT. UNION ALL возвращает все строки, включая дубликаты.
```go
unions := uast.UnionAll(uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ),
)
```
Output MariaDB:
```text
UNION ALL SELECT `t`.`string` FROM `test` AS `t`
```
Output MsSQL:
```text
UNION ALL SELECT [t].[string] FROM [test] AS [t]
```
Output MySQL:
```text
UNION ALL SELECT `t`.`string` FROM `test` AS `t`
```
Output PostgreSQL:
```text
UNION ALL SELECT "t"."string" FROM "test" AS "t"
```
Output SQLite:
```text
UNION ALL SELECT "t"."string" FROM "test" AS "t"
```

### UnionExcept
Объединяет результаты нескольких операторов SELECT. EXCEPT возвращает уникальные строки из первого запроса, которых нет во втором.
```go
unions := uast.UnionExcept(uast.NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[string]("t", "string"),
    ),
)
```
Output MariaDB:
```text
EXCEPT SELECT `t`.`string` FROM `test` AS `t`
```
Output MsSQL:
```text
EXCEPT SELECT [t].[string] FROM [test] AS [t]
```
Output MySQL:
```text
EXCEPT SELECT `t`.`string` FROM `test` AS `t`
```
Output PostgreSQL:
```text
EXCEPT SELECT "t"."string" FROM "test" AS "t"
```
Output SQLite:
```text
EXCEPT SELECT "t"."string" FROM "test" AS "t"
```

### UnionIntersect
Объединяет результаты нескольких операторов SELECT. INTERSECT возвращает уникальные строки, общие для обоих запросов.
```go
unions := uast.UnionIntersect(uast.NewSelect(uast.NewTable("test").As("t")).
	Field(
		uast.Column[string]("t", "string"),
	),
)
```
Output MariaDB:
```text
INTERSECT SELECT `t`.`string` FROM `test` AS `t`
```
Output MsSQL:
```text
INTERSECT SELECT [t].[string] FROM [test] AS [t]
```
Output MySQL:
```text
INTERSECT SELECT `t`.`string` FROM `test` AS `t`
```
Output PostgreSQL:
```text
INTERSECT SELECT "t"."string" FROM "test" AS "t"
```
Output SQLite:
```text
INTERSECT SELECT "t"."string" FROM "test" AS "t"
```

## clauseValues
### Pair
Указывает значения для вставки с помощью `Pair` для связывания колонок со значениями. Колонки автоматически определяются из пар.
```go
values := Values(
    uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
	uast.Pair(uast.Column[int]("t", "number"), uast.Value(2)),
)
```
Output MariaDB:
```text
VALUES (?, ?)
```
Output MsSQL:
```text
VALUES (@p1, @p2)
```
Output MySQL:
```text
VALUES (?, ?)
```
Output PostgreSQL:
```text
VALUES ($1, $2)
```
Output SQLite:
```text
VALUES (?, ?)
```

### Upsert
Добавляет оператор upsert к INSERT ... VALUES с помощью `Upsert`. Связывает колонки со значениями.
```go
values := Values(
    uast.Pair(uast.Column[string]("t", "string"), uast.Value("ivan")),
	uast.Pair(uast.Column[int]("t", "number"), uast.Value(2)),
).
Upsert(
    uast.Pair(uast.Column[string]("t", "string"), uast.Value("updated")),
)
```
Output MariaDB:
```text
VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
```
Output MsSQL:
```text
// Not supported
```
Output MySQL:
```text
VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
```
Output PostgreSQL:
```text
VALUES ($1, $2) ON CONFLICT DO UPDATE SET "string" = $3
```
Output SQLite:
```text
VALUES (?, ?) ON CONFLICT DO UPDATE SET "string" = ?
```

## clauseWhere
Добавляет оператор WHERE для фильтрации строк перед группировкой или агрегацией. Принимает выражения сравнения, логические операторы и подзапросы.
```go
where = Where(
	uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
)
```
Output MariaDB:
```text
WHERE `t`.`string` = ?
```
Output MsSQL:
```text
WHERE [t].[string] = @p1
```
Output MySQL:
```text
WHERE `t`.`string` = ?
```
Output PostgreSQL:
```text
WHERE "t"."string" = $1
```
Output SQLite:
```text
WHERE "t"."string" = ?
```

## clauseWith
### Norecursive
Добавляет нерекурсивное общее табличное выражение (CTE) к запросу с помощью `WithN`. Колонки получают псевдонимы через вариативные строковые аргументы.
```go
with := WithN("cte_norecursive", NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[int64]("t", "id"),
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    ),
    "id", "string",
)
```
Output MariaDB:
```text
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ?)
```
Output MsSQL:
```text
WITH [cte_norecursive] ([id], [string]) AS (SELECT [t].[id], [t].[string] FROM [test] AS [t] WHERE [t].[string] = @p1)
```
Output MySQL:
```text
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ?)
```
Output PostgreSQL:
```text
WITH "cte_norecursive" ("id", "string") AS (SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = $1)
```
Output SQLite:
```text
WITH "cte_norecursive" ("id", "string") AS (SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = ?)
```

### Recursive
Добавляет рекурсивное общее табличное выражение (CTE) к запросу с помощью `WithR`. Требует оператор `Unions` с `UnionAll` для определения рекурсивного шага.
```go
with := WithR("cte_recursive", NewSelect(uast.NewTable("test").As("t")).
    Field(
        uast.Column[int64]("t", "id"),
        uast.Column[string]("t", "string"),
    ).
    Where(
        uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    ).
    Unions(
        uast.UnionAll(uast.NewSelect(uast.NewTable("test").As("t")).
            Field(
                uast.Column[int64]("t", "id"),
                uast.Column[string]("t", "string"),
            ).
            Join(
                uast.Inner(uast.NewCTE("cte_recursive", "rec"), uast.Equal(uast.Column[int64]("t", "id"), uast.Column[int64]("rec", "id"))),
            ),
        ),
    ),
    "id", "string",
)
```
Output MariaDB:
```text
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ? UNION ALL SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` INNER JOIN `cte_recursive` AS `rec` ON `t`.`id` = `rec`.`id`)
```
Output MsSQL:
```text
WITH RECURSIVE [cte_recursive] ([id], [string]) AS (SELECT [t].[id], [t].[string] FROM [test] AS [t] WHERE [t].[string] = @p1 UNION ALL SELECT [t].[id], [t].[string] FROM [test] AS [t] INNER JOIN [cte_recursive] AS [rec] ON [t].[id] = [rec].[id])
```
Output MySQL:
```text
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` WHERE `t`.`string` = ? UNION ALL SELECT `t`.`id`, `t`.`string` FROM `test` AS `t` INNER JOIN `cte_recursive` AS `rec` ON `t`.`id` = `rec`.`id`)
```
Output PostgreSQL:
```text
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = $1 UNION ALL SELECT "t"."id", "t"."string" FROM "test" AS "t" INNER JOIN "cte_recursive" AS "rec" ON "t"."id" = "rec"."id")
```
Output SQLite:
```text
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "t"."id", "t"."string" FROM "test" AS "t" WHERE "t"."string" = ? UNION ALL SELECT "t"."id", "t"."string" FROM "test" AS "t" INNER JOIN "cte_recursive" AS "rec" ON "t"."id" = "rec"."id")
```

## exprArray
### Array
Создаёт выражение массива для использования в SQL-запросах.
```go
array := uast.Array(0, 1, 2)
```
Output MariaDB:
```text
ARRAY[?, ?, ?]
```
Output MsSQL:
```text
ARRAY[@p1, @p2, @p3]
```
Output MySQL:
```text
ARRAY[?, ?, ?]
```
Output PostgreSQL:
```text
ARRAY[$1, $2, $3]
```
Output SQLite:
```text
ARRAY[?, ?, ?]

## exprBinary
### BitwiseAnd
Выполняет побитовую операцию И между двумя выражениями.
```go
binary := uast.BitwiseAnd(uast.Column[int]("t", "number"), uast.Value(0b0010))
```
Output MsSQL:
```text
[t].[number] & @p1
```
Output MySQL:
```text
`t`.`number` & ?
```
Output PostgreSQL:
```text
"t"."number" & $1
```

### BitwiseOr
Выполняет побитовую операцию ИЛИ между двумя выражениями.
```go
binary := uast.BitwiseOr(uast.Column[int]("t", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`t`.`number` | ?
```
Output MsSQL:
```text
[t].[number] | @p1
```
Output MySQL:
```text
`t`.`number` | ?
```
Output PostgreSQL:
```text
"t"."number" | $1
```
Output SQLite:
```text
"t"."number" | ?
```

### BitwiseXor
Выполняет побитовую операцию исключающего ИЛИ между двумя выражениями.
```go
binary := uast.BitwiseXor(uast.Column[int]("t", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`t`.`number` ^ ?
```
Output MsSQL:
```text
[t].[number] ^ @p1
```
Output MySQL:
```text
`t`.`number` ^ ?
```
Output PostgreSQL:
```text
"t"."number" ^ $1
```
Output SQLite:
```text
"t"."number" ^ ?
```

### Divide
Делит левое выражение на правое.
```go
binary := uast.Divide(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` / ?
```
Output MsSQL:
```text
[t].[number] / @p1
```
Output MySQL:
```text
`t`.`number` / ?
```
Output PostgreSQL:
```text
"t"."number" / $1
```
Output SQLite:
```text
"t"."number" / ?
```

### Minus
Вычитает правое выражение из левого.
```go
binary := uast.Minus(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` - ?
```
Output MsSQL:
```text
[t].[number] - @p1
```
Output MySQL:
```text
`t`.`number` - ?
```
Output PostgreSQL:
```text
"t"."number" - $1
```
Output SQLite:
```text
"t"."number" - ?
```

### Modulo
Возвращает остаток от деления левого выражения на правое.
```go
binary := uast.Modulo(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` % ?
```
Output MsSQL:
```text
[t].[number] % @p1
```
Output MySQL:
```text
`t`.`number` % ?
```
Output PostgreSQL:
```text
"t"."number" % $1
```
Output SQLite:
```text
"t"."number" % ?
```

### Multiply
Умножает левое выражение на правое.
```go
binary := uast.Multiply(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` * ?
```
Output MsSQL:
```text
[t].[number] * @p1
```
Output MySQL:
```text
`t`.`number` * ?
```
Output PostgreSQL:
```text
"t"."number" * $1
```
Output SQLite:
```text
"t"."number" * ?
```

### Plus
Складывает левое выражение с правым.
```go
binary := uast.Plus(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` + ?
```
Output MsSQL:
```text
[t].[number] + @p1
```
Output MySQL:
```text
`t`.`number` + ?
```
Output PostgreSQL:
```text
"t"."number" + $1
```
Output SQLite:
```text
"t"."number" + ?
```

### ShiftLeft
Выполняет побитовый сдвиг влево левого выражения на количество бит, указанное в правом выражении.
```go
binary := uast.ShiftLeft(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` << ?
```
Output MsSQL:
```text
[t].[number] << @p1
```
Output MySQL:
```text
`t`.`number` << ?
```
Output PostgreSQL:
```text
"t"."number" << $1
```
Output SQLite:
```text
"t"."number" << ?
```

### ShiftRight
Выполняет побитовый сдвиг вправо левого выражения на количество бит, указанное в правом выражении.
```go
binary := uast.ShiftRight(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` >> ?
```
Output MsSQL:
```text
[t].[number] >> @p1
```
Output MySQL:
```text
`t`.`number` >> ?
```
Output PostgreSQL:
```text
"t"."number" >> $1
```
Output SQLite:
```text
"t"."number" >> ?
```

## exprColumn
### Column
Создаёт ссылку на колонку таблицы, опционально квалифицированную псевдонимом таблицы. Это основной способ ссылаться на колонки базы данных в выражениях.
```go
column := uast.Column[string]("t", "string")
```
Output MariaDB:
```text
`t`.`string`
```
Output MsSQL:
```text
[t].[string]
```
Output MySQL:
```text
`t`.`string`
```
Output PostgreSQL:
```text
"t"."string"
```
Output SQLite:
```text
"t"."string"
```

## exprComparison
### Between
Проверяет, попадает ли левое выражение в диапазон, заданный valueStart и valueEnd (включительно).
```go
comparison := uast.Between(uast.Column[int]("t", "number"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` BETWEEN ? AND ?
```
Output MsSQL:
```text
[t].[number] BETWEEN @p1 AND @p2
```
Output MySQL:
```text
`t`.`number` BETWEEN ? AND ?
```
Output PostgreSQL:
```text
"t"."number" BETWEEN $1 AND $2
```
Output SQLite:
```text
"t"."number" BETWEEN ? AND ?
```

### Equal
Сравнивает два выражения на равенство (`=`).
```go
comparison := uast.Equal(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` = ?
```
Output MsSQL:
```text
[t].[number] = @p1
```
Output MySQL:
```text
`t`.`number` = ?
```
Output PostgreSQL:
```text
"t"."number" = $1
```
Output SQLite:
```text
"t"."number" = ?
```

### Exists
Проверяет, возвращает ли подзапрос какие-либо строки. Возвращает `true` если существует хотя бы одна строка.
```go
comparison := uast.Exists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("test").As("t"))))
```
Output MariaDB:
```text
EXISTS (SELECT 1 FROM `test` AS `t`)
```
Output MsSQL:
```text
EXISTS (SELECT 1 FROM [test] AS [t])
```
Output MySQL:
```text
EXISTS (SELECT 1 FROM `test` AS `t`)
```
Output PostgreSQL:
```text
EXISTS (SELECT 1 FROM "test" AS "t")
```
Output SQLite:
```text
EXISTS (SELECT 1 FROM "test" AS "t")
```

### Greater
Сравнивает, больше ли левое выражение правого (`>`).
```go
comparison := uast.Greater(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` > ?
```
Output MsSQL:
```text
[t].[number] > @p1
```
Output MySQL:
```text
`t`.`number` > ?
```
Output PostgreSQL:
```text
"t"."number" > $1
```
Output SQLite:
```text
"t"."number" > ?
```

### GreaterEqual
Сравнивает, больше или равно ли левое выражение правому (`>=`).
```go
comparison := uast.GreaterEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` >= ?
```
Output MsSQL:
```text
[t].[number] >= @p1
```
Output MySQL:
```text
`t`.`number` >= ?
```
Output PostgreSQL:
```text
"t"."number" >= $1
```
Output SQLite:
```text
"t"."number" >= ?
```

### ILike
Выполняет регистронезависимое сравнение с шаблоном. Правое выражение должно содержать шаблон с `%` (любая последовательность) и `_` (один символ).
```go
comparison := uast.ILike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
LOWER(`t`.`string`) LIKE LOWER(?)
```
Output MsSQL:
```text
LOWER([t].[string]) LIKE LOWER(@p1)
```
Output MySQL:
```text
LOWER(`t`.`string`) LIKE LOWER(?)
```
Output PostgreSQL:
```text
"t"."string" ILIKE $1
```
Output SQLite:
```text
LOWER("t"."string") LIKE LOWER(?)
```

### In
Проверяет, соответствует ли левое выражение любому значению, содержащемуся в правом выражении (обычно подзапрос или массив).
```go
comparison := uast.In(uast.Column[string]("t", "string"), uast.Array("active", "pending"))
```
Output MariaDB:
```text
`t`.`string` IN (?, ?)
```
Output MsSQL:
```text
[t].[string] IN (@p1, @p2)
```
Output MySQL:
```text
`t`.`string` IN (?, ?)
```
Output PostgreSQL:
```text
"t"."string" IN ($1, $2)
```
Output SQLite:
```text
"t"."string" IN (?, ?)
```

### IsNotNull
Проверяет, что выражение не `NULL`.
```go
comparison := uast.IsNotNull(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` IS NOT NULL
```
Output MsSQL:
```text
[t].[string] IS NOT NULL
```
Output MySQL:
```text
`t`.`string` IS NOT NULL
```
Output PostgreSQL:
```text
"t"."string" IS NOT NULL
```
Output SQLite:
```text
"t"."string" IS NOT NULL
```

### IsNull
Проверяет, что выражение является `NULL`.
```go
comparison := uast.IsNull(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
`t`.`string` IS NULL
```
Output MsSQL:
```text
[t].[string] IS NULL
```
Output MySQL:
```text
`t`.`string` IS NULL
```
Output PostgreSQL:
```text
"t"."string" IS NULL
```
Output SQLite:
```text
"t"."string" IS NULL
```

### Less
Сравнивает, меньше ли левое выражение правого (`<`).
```go
comparison := uast.Less(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` < ?
```
Output MsSQL:
```text
[t].[number] < @p1
```
Output MySQL:
```text
`t`.`number` < ?
```
Output PostgreSQL:
```text
"t"."number" < $1
```
Output SQLite:
```text
"t"."number" < ?
```

### LessEqual
Сравнивает, меньше или равно ли левое выражение правому (`<=`).
```go
comparison := uast.LessEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` <= ?
```
Output MsSQL:
```text
[t].[number] <= @p1
```
Output MySQL:
```text
`t`.`number` <= ?
```
Output PostgreSQL:
```text
"t"."number" <= $1
```
Output SQLite:
```text
"t"."number" <= ?
```

### Like
Выполняет регистрозависимое сравнение с шаблоном. Правое выражение должно содержать шаблон с `%` и `_`.
```go
comparison := uast.Like(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
`t`.`string` LIKE ?
```
Output MsSQL:
```text
[t].[number] LIKE @p1
```
Output MySQL:
```text
`t`.`string` LIKE ?
```
Output PostgreSQL:
```text
"t"."string" LIKE $1
```
Output SQLite:
```text
"t"."string" LIKE ?
```

### NotBetween
Проверяет, находится ли левое выражение вне диапазона, заданного `valueStart` и `valueEnd`.
```go
comparison := uast.NotBetween(uast.Column[int]("t", "number"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` NOT BETWEEN ? AND ?
```
Output MsSQL:
```text
[t].[number] NOT BETWEEN @p1 AND @p2
```
Output MySQL:
```text
`t`.`number` NOT BETWEEN ? AND ?
```
Output PostgreSQL:
```text
"t"."number" NOT BETWEEN $1 AND $2
```
Output SQLite:
```text
"t"."number" NOT BETWEEN ? AND ?
```

### NotEqual
Сравнивает два выражения на неравенство (`!=` or `<>`).
```go
comparison := uast.NotEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
`t`.`number` != ?
```
Output MsSQL:
```text
[t].[number] != @p1
```
Output MySQL:
```text
`t`.`number` != ?
```
Output PostgreSQL:
```text
"t"."number" != $1
```
Output SQLite:
```text
"t"."number" != ?
```

### NotExists
Проверяет, что подзапрос не возвращает строк. Возвращает `true` если результат подзапроса пуст.
```go
comparison := uast.NotExists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("test").As("t"))))
```
Output MariaDB:
```text
NOT EXISTS (SELECT 1 FROM `test` AS `t`)
```
Output MsSQL:
```text
NOT EXISTS (SELECT 1 FROM [test] AS [t])
```
Output MySQL:
```text
NOT EXISTS (SELECT 1 FROM `test` AS `t`)
```
Output PostgreSQL:
```text
NOT EXISTS (SELECT 1 FROM "test" AS "t")
```
Output SQLite:
```text
NOT EXISTS (SELECT 1 FROM "test" AS "t")
```

### NotILike
Выполняет отрицательное регистронезависимое сравнение с шаблоном.
```go
comparison := uast.NotILike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
LOWER(`t`.`string`) NOT LIKE LOWER(?)
```
Output MsSQL:
```text
LOWER([t].[string]) NOT LIKE LOWER(@p1)
```
Output MySQL:
```text
LOWER(`t`.`string`) NOT LIKE LOWER(?)
```
Output PostgreSQL:
```text
"t"."string" NOT ILIKE $1
```
Output SQLite:
```text
LOWER("t"."string") NOT LIKE LOWER(?)
```

### NotIn
Проверяет, что левое выражение не соответствует ни одному значению, содержащемуся в правом выражении.
```go
comparison := uast.NotIn(uast.Column[string]("t", "string"), uast.Array("active", "pending"))
```
Output MariaDB:
```text
`t`.`string` NOT IN (?, ?)
```
Output MsSQL:
```text
[t].[string] NOT IN (@p1, @p2)
```
Output MySQL:
```text
`t`.`string` NOT IN (?, ?)
```
Output PostgreSQL:
```text
"t"."string" NOT IN ($1, $2)
```
Output SQLite:
```text
"t"."string" NOT IN (?, ?)
```

### NotLike
Выполняет отрицательное регистрозависимое сравнение с шаблоном.
```go
comparison := uast.NotLike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
`t`.`string` NOT LIKE ?
```
Output MsSQL:
```text
[t].[string] NOT LIKE @p1
```
Output MySQL:
```text
`t`.`string` NOT LIKE ?
```
Output PostgreSQL:
```text
"t"."string" NOT LIKE $1
```
Output SQLite:
```text
"t"."string" NOT LIKE ?
```

## exprConstant
### ConstBoolFalse
Возвращает константное булево выражение `FALSE`.
```go
constant := uast.ConstBoolFalse()
```
Output:
```text
FALSE
```

### ConstBoolTrue
Возвращает константное булево выражение `TRUE`.
```go
constant := uast.ConstBoolTrue()
```
Output:
```text
TRUE
```

### ConstFloat32One
Возвращает константное значение `float32` равное `1.0`. 
```go
constant := uast.ConstFloat32One()
```
Output:
```text
1.0
```

### ConstFloat64One
Возвращает константное значение `float64` равное `1.000000`.
```go
constant := uast.ConstFloat64One()
```
Output:
```text
1.000000
```

### ConstIntOne
Возвращает константное значение `int` равное `1`.
```go
constant := uast.ConstIntOne()
```
Output:
```text
1
```

### ConstInt8One
Возвращает константное значение `int8` равное `1`.
```go
constant := uast.ConstInt8One()
```
Output:
```text
1
```

### ConstInt16One
Возвращает константное значение `int16` равное `1`.
```go
constant := uast.ConstInt16One()
```
Output:
```text
1
```

### ConstInt32One
Возвращает константное значение `int32` равное `1`.
```go
constant := uast.ConstInt32One()
```
Output:
```text
1
```

### ConstInt64One
Возвращает константное значение `int64` равное `1`.
```go
constant := uast.ConstInt64One()
```
Output:
```text
1
```

### ConstStringDefault
Возвращает константное значение `string` равное `DEFAULT`.
```go
constant := uast.ConstStringDefault()
```
Output:
```text
DEFAULT
```

### ConstStringNull
Возвращает константное значение `string` равное `NULL`.
```go
constant := uast.ConstStringNull()
```
Output:
```text
NULL
```

### ConstUintOne
Возвращает константное значение `uint` равное `1`.
```go
constant := uast.ConstUintOne()
```
Output:
```text
1
```

### ConstUint8One
Возвращает константное значение `uint8` равное `1`.
```go
constant := uast.ConstUint8One()
```
Output:
```text
1
```

### ConstUint16One
Возвращает константное значение `uint16` равное `1`.
```go
constant := uast.ConstUint16One()
```
Output:
```text
1
```

### ConstUint32One
Возвращает константное значение `uint32` равное `1`.
```go
constant := uast.ConstUint32One()
```
Output:
```text
1
```

### ConstUint64One
Возвращает константное значение `uint64` равное `1`.
```go
constant := uast.ConstUint64One()
```
Output:
```text
1
```

## exprFunction
### Aggregate
#### Avg
Возвращает среднее арифметическое всех не-NULL значений в выражении. Если `distinct` равен `true`, среднее вычисляется только по уникальным значениям.
```go
function := uast.Avg(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Avg(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
AVG(`t`.`number`)
AVG(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
AVG([t].[number])
AVG(DISTINCT [t].[number])
```
Output MySQL:
```text
AVG(`t`.`number`)
AVG(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
AVG("t"."number")
AVG(DISTINCT "t"."number")
```
Output SQLite:
```text
AVG("t"."number")
AVG(DISTINCT "t"."number")
```

#### BitAnd
Возвращает побитовое И всех битов в выражении. Имеет смысл только для целочисленных типов.
```go
function := uast.BitAnd(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitAnd(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
BIT_AND(`t`.`number`)
BIT_AND(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
BIT_AND([t].[number])
BIT_AND(DISTINCT [t].[number])
```
Output MySQL:
```text
BIT_AND(`t`.`number`)
BIT_AND(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
BIT_AND("t"."number")
BIT_AND(DISTINCT "t"."number")
```
Output SQLite:
```text
BIT_AND("t"."number")
BIT_AND(DISTINCT "t"."number")
```

#### BitOr
Возвращает побитовое ИЛИ всех битов в выражении.
```go
function := uast.BitOr(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitOr(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
BIT_OR(`t`.`number`)
BIT_OR(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
BIT_OR([t].[number])
BIT_OR(DISTINCT [t].[number])
```
Output MySQL:
```text
BIT_OR(`t`.`number`)
BIT_OR(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
BIT_OR("t"."number")
BIT_OR(DISTINCT "t"."number")
```
Output SQLite:
```text
BIT_OR("t"."number")
BIT_OR(DISTINCT "t"."number")
```

#### BitXor
Возвращает побитовое исключающее ИЛИ всех битов в выражении.
```go
function := uast.BitXor(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitXor(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
BIT_XOR(`t`.`number`)
BIT_XOR(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
BIT_XOR([t].[number])
BIT_XOR(DISTINCT [t].[number])
```
Output MySQL:
```text
BIT_XOR(`t`.`number`)
BIT_XOR(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
BIT_XOR("t"."number")
BIT_XOR(DISTINCT "t"."number")
```
Output SQLite:
```text
BIT_XOR("t"."number")
BIT_XOR(DISTINCT "t"."number")
```

#### Count
Возвращает количество строк, соответствующих запросу, или количество не-NULL значений, если указано выражение. Когда `distinct` равен `true`, подсчитываются только уникальные значения.
```go
function := uast.Count(uast.Column[string]("t", "string"), false)
functionWithDistinct := uast.Count(uast.Column[string]("t", "string"), true)
```
Output MariaDB:
```text
COUNT(`t`.`string`)
COUNT(DISTINCT `t`.`string`)
```
Output MsSQL:
```text
COUNT([t].[string])
COUNT(DISTINCT [t].[string])
```
Output MySQL:
```text
COUNT(`t`.`string`)
COUNT(DISTINCT `t`.`string`)
```
Output PostgreSQL:
```text
COUNT("t"."string")
COUNT(DISTINCT "t"."string")
```
Output SQLite:
```text
COUNT("t"."string")
COUNT(DISTINCT "t"."string")
```

#### GroupConcat
Объединяет значения из группы в одну строку, разделённую стандартным разделителем (обычно запятая). Флаг `distinct` удаляет дубликаты перед объединением.
```go
function := uast.GroupConcat(uast.Column[string]("t", "string"), false)
functionWithDistinct := uast.GroupConcat(uast.Column[string]("t", "string"), true)
```
Output MariaDB:
```text
GROUP_CONCAT(`t`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `t`.`string` SEPARATOR ',')
```
Output MsSQL:
```text
GROUP_CONCAT([t].[string], ',')
GROUP_CONCAT(DISTINCT [t].[string], ',')
```
Output MySQL:
```text
GROUP_CONCAT(`t`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `t`.`string` SEPARATOR ',')
```
Output PostgreSQL:
```text
STRING_AGG("t"."string", ',')
STRING_AGG(DISTINCT "t"."string", ',')
```
Output SQLite:
```text
GROUP_CONCAT("t"."string" SEPARATOR ',')
GROUP_CONCAT(DISTINCT "t"."string" SEPARATOR ',')
```

#### Max
Возвращает максимальное значение выражения по всем строкам в группе.
```go
function := uast.Max(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Max(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
MAX(`t`.`number`)
MAX(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
MAX([t].[number])
MAX(DISTINCT [t].[number])
```
Output MySQL:
```text
MAX(`t`.`number`)
MAX(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
MAX("t"."number")
MAX(DISTINCT "t"."number")
```
Output SQLite:
```text
MAX("t"."number")
MAX(DISTINCT "t"."number")
```

#### Min
Возвращает минимальное значение выражения по всем строкам в группе.
```go
function := uast.Min(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Min(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
MIN(`t`.`number`)
MIN(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
MIN([t].[number])
MIN(DISTINCT [t].[number])
```
Output MySQL:
```text
MIN(`t`.`number`)
MIN(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
MIN("t"."number")
MIN(DISTINCT "t"."number")
```
Output SQLite:
```text
MIN("t"."number")
MIN(DISTINCT "t"."number")
```

#### StdDev
Возвращает популяционное стандартное отклонение выражения.
```go
function := uast.StdDev(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.StdDev(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
STDDEV(`t`.`number`)
STDDEV(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
STDEV([t].[number])
STDEV(DISTINCT [t].[number])
```
Output MySQL:
```text
STDDEV(`t`.`number`)
STDDEV(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
STDDEV_SAMP("t"."number")
STDDEV_SAMP(DISTINCT "t"."number")
```
Output SQLite:
```text
STDEV("t"."number")
STDEV(DISTINCT "t"."number")
```

#### Sum
Возвращает сумму всех значений в выражении. Если `distinct` равен `true`, суммируются только уникальные значения.
```go
function := uast.Sum(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Sum(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
SUM(`t`.`number`)
SUM(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
SUM([t].[number])
SUM(DISTINCT [t].[number])
```
Output MySQL:
```text
SUM(`t`.`number`)
SUM(DISTINCT `t`.`number`)
```
Output PostgreSQL:
```text
SUM("t"."number")
SUM(DISTINCT "t"."number")
```
Output SQLite:
```text
SUM("t"."number")
SUM(DISTINCT "t"."number")
```

#### Variance
Возвращает популяционную дисперсию выражения.
```go
function := uast.Variance(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Variance(uast.Column[int]("t", "number"), true)
```
Output MariaDB:
```text
VARIANCE(`t`.`number`)
VARIANCE(DISTINCT `t`.`number`)
```
Output MsSQL:
```text
VAR([t].[number])
VAR(DISTINCT [t].[number])
```
Output MySQL:
```text
VARIANCE(`t`.`number`)
VARIANCE(DISTINCT "t"."number")
```
Output PostgreSQL:
```text
VAR_SAMP("t"."number")
VAR_SAMP(DISTINCT "t"."number")
```
Output SQLite:
```text
VARIANCE("t"."number")
VARIANCE(DISTINCT "t"."number")
```

### Analytical
#### FirstValue
Возвращает значение выражения из первой строки оконного фрейма. Требует оператор `OVER` с оконной спецификацией.
```go
function := uast.FirstValue(uast.Column[string]("t", "string")).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
FIRST_VALUE(`t`.`string`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
FIRST_VALUE([t].[string]) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
FIRST_VALUE(`t`.`string`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
FIRST_VALUE("t"."string") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
FIRST_VALUE("t"."string") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### Lag
Возвращает значение выражения из строки, смещённой на `offset` строк назад от текущей строки в рамках раздела.
```go
function := uast.Lag(uast.Column[int]("t", "number"), 2).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Asc(uast.Column[time.Time]("t", "date"))),
)
```
Output MariaDB:
```text
LAG(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)
```
Output MsSQL:
```text
LAG([t].[number], 2) OVER (PARTITION BY [t].[id] ORDER BY [t].[date] ASC)
```
Output MySQL:
```text
LAG(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)
```
Output PostgreSQL:
```text
LAG("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)
```
Output SQLite:
```text
LAG("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)
```

#### LastValue
Возвращает значение выражения из последней строки оконного фрейма.
```go
function := uast.LastValue(uast.Column[string]("t", "string")).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Asc(uast.Column[int]("t", "number"))),
    uast.RowsBetween("CURRENT ROW", "UNBOUNDED FOLLOWING"),
)
```
Output MariaDB:
```text
LAST_VALUE(`t`.`string`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output MsSQL:
```text
LAST_VALUE([t].[string]) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output MySQL:
```text
LAST_VALUE(`t`.`string`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output PostgreSQL:
```text
LAST_VALUE("t"."string") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output SQLite:
```text
LAST_VALUE("t"."string") OVER (PARTITION BY "t"."id" ORDER BY "t"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```

#### Lead
Возвращает значение выражения из строки, смещённой на `offset` строк вперёд от текущей строки в рамках раздела.
```go
function := uast.Lead(uast.Column[int]("t", "number"), 2).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Asc(uast.Column[time.Time]("t", "date"))),
)
```
Output MariaDB:
```text
LEAD(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)
```
Output MsSQL:
```text
LEAD([t].[number], 2) OVER (PARTITION BY [t].[id] ORDER BY [t].[date] ASC)
```
Output MySQL:
```text
LEAD(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)
```
Output PostgreSQL:
```text
LEAD("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)
```
Output SQLite:
```text
LEAD("t"."number", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."date" ASC)
```

#### NthValue
Возвращает значение выражения из `n-й` строки оконного фрейма.
```go
function := uast.NthValue(uast.Column[string]("t", "string"), 2).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
    uast.RowsBetween("UNBOUNDED PRECEDING", "CURRENT ROW"),
)
```
Output MariaDB:
```text
NTH_VALUE(`t`.`string`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output MsSQL:
```text
NTH_VALUE([t].[string], 2) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output MySQL:
```text
NTH_VALUE(`t`.`string`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output PostgreSQL:
```text
NTH_VALUE("t"."string", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output SQLite:
```text
NTH_VALUE("t"."string", 2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```

### Condition
#### Case
Вычисляет список пар `WHEN`-`THEN` и возвращает выражение `THEN` для первого истинного WHEN. Если ни одно условие не истинно, возвращает выражение  `ELSE` если оно указано, иначе `NULL`.
```go
pairs := uast.CaseIf(
    uast.CasePair(
        uast.Less(uast.Column[int]("t", "number"), uast.Value(2)),
        uast.Value("old"),
    ),
)
elseExpr := uast.CaseElse(uast.Value("new"))
function := uast.Case(pairs, elseExpr)
```
Output MariaDB:
```text
CASE WHEN `t`.`number` < ? THEN ? ELSE ? END
```
Output MsSQL:
```text
CASE WHEN [t].[number] < @p1 THEN @p2 ELSE @p3 END
```
Output MySQL:
```text
CASE WHEN `t`.`number` < ? THEN ? ELSE ? END
```
Output PostgreSQL:
```text
CASE WHEN "t"."number" < $1 THEN $2 ELSE $3 END
```
Output SQLite:
```text
CASE WHEN "t"."number" < ? THEN ? ELSE ? END
```

#### Coalesce
Возвращает первое не-NULL выражение из предоставленного списка. Полезно для указания запасных значений.
```go
function := uast.Coalesce(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
COALESCE(`t`.`createat`, `t`.`updateat`)
```
Output MsSQL:
```text
COALESCE([t].[createat], [t].[updateat])
```
Output MySQL:
```text
COALESCE(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
COALESCE("t"."createat", "t"."updateat")
```
Output SQLite:
```text
COALESCE("t"."createat", "t"."updateat")
```

#### Greatest
Возвращает наибольшее значение из предоставленного списка выражений.
```go
function := uast.Greatest(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
GREATEST(`t`.`createat`, `t`.`updateat`)
```
Output MsSQL:
```text
GREATEST([t].[createat], [t].[updateat])
```
Output MySQL:
```text
GREATEST(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
GREATEST("t"."createat", "t"."updateat")
```
Output SQLite:
```text
GREATEST("t"."createat", "t"."updateat")
```

#### Least
Возвращает наименьшее значение из предоставленного списка выражений.
```go
function := uast.Least(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
LEAST(`t`.`createat`, `t`.`updateat`)
```
Output MsSQL:
```text
LEAST([t].[createat], [t].[updateat])
```
Output MySQL:
```text
LEAST(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
LEAST("t"."createat", "t"."updateat")
```
Output SQLite:
```text
LEAST("t"."createat", "t"."updateat")
```

#### NullIf
Возвращает `NULL` если два выражения равны; иначе возвращает первое выражение.
```go
function := uast.NullIf(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MariaDB:
```text
NULLIF(`t`.`createat`, `t`.`updateat`)
```
Output MsSQL:
```text
NULLIF([t].[createat], [t].[updateat])
```
Output MySQL:
```text
NULLIF(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
NULLIF("t"."createat", "t"."updateat")
```
Output SQLite:
```text
NULLIF("t"."createat", "t"."updateat")
```

### Convert
#### Cast
Преобразует выражение к указанному типу данных.
```go
function := uast.Cast(uast.Column[int]("t", "number"), uast.TypeString)
```
Output MariaDB:
```text
CAST(`t`.`number` AS CHAR)
```
Output MsSQL:
```text
CAST([t].[number] AS NVARCHAR)
```
Output MySQL:
```text
CAST(`t`.`number` AS CHAR)
```
Output PostgreSQL:
```text
CAST("t"."number" AS VARCHAR)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

#### CharLength
Возвращает количество символов в строковом выражении.
```go
function := uast.CharLength(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
CHAR_LENGTH(`t`.`string`)
```
Output MsSQL:
```text
CHAR_LENGTH([t].[string])
```
Output MySQL:
```text
CHAR_LENGTH(`t`.`string`)
```
Output PostgreSQL:
```text
CHAR_LENGTH("t"."string")
```
Output SQLite:
```text
CHAR_LENGTH("t"."string")
```

#### DateFormat
Форматирует выражение даты/времени в соответствии с указанной маской формата.
```go
function := uast.DateFormat(uast.Column[time.Time]("t", "createat"), uast.Value("%Y-%m-%d"))
```
Output MariaDB:
```text
DATE_FORMAT(`t`.`createat`, '%Y-%m-%d')
```
Output MsSQL:
```text
FORMAT([t].[createat], '%Y-%m-%d')
```
Output MySQL:
```text
DATE_FORMAT(`t`.`createat`, '%Y-%m-%d')
```
Output PostgreSQL:
```text
TO_CHAR("t"."createat", '%Y-%m-%d')
```
Output SQLite:
```text
strftime("t"."createat", '%Y-%m-%d')
```

#### Degrees
Преобразует угол из радиан в градусы.
```go
function := uast.Degrees(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
DEGREES(`t`.`number`)
```
Output MsSQL:
```text
DEGREES([t].[number])
```
Output MySQL:
```text
DEGREES(`t`.`number`)
```
Output PostgreSQL:
```text
DEGREES("t"."number")
```
Output SQLite:
```text
DEGREES("t"."number")
```

#### Length
Возвращает длину строкового выражения в байтах.
```go
function := uast.Length(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
LENGTH(`t`.`string`)
```
Output MsSQL:
```text
LEN([t].[string])
```
Output MySQL:
```text
LENGTH(`t`.`string`)
```
Output PostgreSQL:
```text
LENGTH("t"."string")
```
Output SQLite:
```text
LENGTH("t"."string")
```

#### Position
Возвращает начальную позицию первого вхождения подстроки в строку.
```go
function := uast.Position(uast.Column[string]("t", "string"), uast.Value("old"))
```
Output MariaDB:
```text
POSITION(? IN `t`.`string`)
```
Output MsSQL:
```text
CHARINDEX(@p1, [t].[string])
```
Output MySQL:
```text
POSITION(? IN `t`.`string`)
```
Output PostgreSQL:
```text
POSITION($1 IN "t"."string")
```
Output SQLite:
```text
POSITION(? IN "t"."string")
```

#### Radians
Преобразует угол из градусов в радианы.
```go
function := uast.Radians(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
RADIANS(`t`.`number`)
```
Output MsSQL:
```text
RADIANS([t].[number])
```
Output MySQL:
```text
RADIANS(`t`.`number`)
```
Output PostgreSQL:
```text
RADIANS("t"."number")
```
Output SQLite:
```text
RADIANS("t"."number")
```

### Date and time
#### CurDate
Возвращает текущую дату (без времени).
```go
function := uast.CurDate()
```
Output MariaDB:
```text
CURDATE()
```
Output MsSQL:
```text
CAST(GETDATE() AS DATE)
```
Output MySQL:
```text
CURDATE()
```
Output PostgreSQL:
```text
CURRENT_DATE
```
Output SQLite:
```text
date('now')
```

#### CurTime
Возвращает текущее время (без даты).
```go
function := uast.CurTime()
```
Output MariaDB:
```text
CURTIME()
```
Output MsSQL:
```text
CAST(GETDATE() AS TIME)
```
Output MySQL:
```text
CURTIME()
```
Output PostgreSQL:
```text
CURRENT_TIME
```
Output SQLite:
```text
time('now')
```

#### DateAdd
Добавляет интервал даты/времени к выражению даты/времени и возвращает результирующую дату/время.
```go
function := uast.DateAdd(uast.Column[time.Time]("t", "createat"), uast.Value("2 DAY"))
```
Output MariaDB:
```text
DATE_ADD(`t`.`createat`, INTERVAL 2 DAY)
```
Output MsSQL:
```text
// In development
```
Output MySQL:
```text
DATE_ADD(`t`.`createat`, INTERVAL 2 DAY)
```
Output PostgreSQL:
```text
("t"."createat" + INTERVAL '2 DAY')
```
Output SQLite:
```text
datetime("t"."createat",  '+2 DAY')
```

#### DateDiff
Возвращает разницу в днях между двумя выражениями даты/времени (`datetimeEnd` - `datetimeStart`).
```go
function := uast.DateDiff(uast.Column[time.Time]("t", "updateat"), uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
DATEDIFF(`t`.`updateat`, `t`.`createat`)
```
Output MsSQL:
```text
DATEDIFF([t].[updateat], [t].[createat])
```
Output MySQL:
```text
DATEDIFF(`t`.`updateat`, `t`.`createat`)
```
Output PostgreSQL:
```text
DATE_PART('day', "t"."updateat" - "t"."createat")
```
Output SQLite:
```text
DATEDIFF("t"."updateat", "t"."createat")
```

#### DateSub
Вычитает интервал даты/времени из выражения даты/времени и возвращает результирующую дату/время.
```go
function := uast.DateSub(uast.Column[time.Time]("t", "createat"), uast.Value("2 DAY"))
```
Output MariaDB:
```text
DATE_SUB(`t`.`createat`, INTERVAL 2 DAY)
```
Output MsSQL:
```text
// In development
```
Output MySQL:
```text
DATE_SUB(`t`.`createat`, INTERVAL 2 DAY)
```
Output PostgreSQL:
```text
("t"."createat" - INTERVAL '2 DAY')
```
Output SQLite:
```text
datetime("t"."createat", '-2 DAY')
```

#### Day
Извлекает день месяца (1–31) из выражения даты/времени.
```go
function := uast.Day(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
DAY(`t`.`createat`)
```
Output MsSQL:
```text
DAY([t].[createat])
```
Output MySQL:
```text
DAY(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(DAY FROM "t"."createat")
```
Output SQLite:
```text
DAY("t"."createat")
```

#### DayName
Возвращает название дня недели (например, 'Понедельник', 'Вторник') для заданного выражения даты/времени.
```go
function := uast.DayName(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
DAYNAME(`t`.`createat`)
```
Output MsSQL:
```text
DATENAME(WEEKDAY, [t].[createat])
```
Output MySQL:
```text
DAYNAME(`t`.`createat`)
```
Output PostgreSQL:
```text
TO_CHAR("t"."createat", 'Day')
```
Output SQLite:
```text
strftime('%w', "t"."createat")
```

#### Hour
Извлекает час (0–23) из выражения даты/времени.
```go
function := uast.Hour(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
HOUR(`t`.`createat`)
```
Output MsSQL:
```text
DATEPART(HOUR, [t].[createat])
```
Output MySQL:
```text
HOUR(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(HOUR FROM "t"."createat")
```
Output SQLite:
```text
HOUR("t"."createat")
```

#### Minute
Извлекает минуту (0–59) из выражения даты/времени.
```go
function := uast.Minute(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
MINUTE(`t`.`createat`)
```
Output MsSQL:
```text
DATEPART(MINUTE, [t].[createat])
```
Output MySQL:
```text
MINUTE(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(MINUTE FROM "t"."createat")
```
Output SQLite:
```text
MINUTE("t"."createat")
```

#### Month
Извлекает месяц (1–12) из выражения даты/времени.
```go
function := uast.Month(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
MONTH(`t`.`createat`)
```
Output MsSQL:
```text
MONTH([t].[createat])
```
Output MySQL:
```text
MONTH(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(MONTH FROM "t"."createat")
```
Output SQLite:
```text
MONTH("t"."createat")
```

#### MonthName
Возвращает название месяца (например, 'Январь', 'Февраль') для заданного выражения даты/времени.
```go
function := uast.MonthName(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
MONTHNAME(`t`.`createat`)
```
Output MsSQL:
```text
DATENAME(MONTH, [t].[createat])
```
Output MySQL:
```text
MONTHNAME(`t`.`createat`)
```
Output PostgreSQL:
```text
TO_CHAR("t"."createat", 'Month')
```
Output SQLite:
```text
strftime('%m', "t"."createat")
```

#### Now
Возвращает текущую дату и время.
```go
function := uast.Now()
```
Output MariaDB:
```text
NOW()
```
Output MsSQL:
```text
GETDATE()
```
Output MySQL:
```text
NOW()
```
Output PostgreSQL:
```text
CURRENT_TIMESTAMP
```
Output SQLite:
```text
datetime('now')
```

#### Quarter
Извлекает квартал (1–4) из выражения даты/времени.
```go
function := uast.Quarter(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
QUARTER(`t`.`createat`)
```
Output MsSQL:
```text
DATEPART(QUARTER, [t].[createat])
```
Output MySQL:
```text
QUARTER(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(QUARTER FROM "t"."createat")
```
Output SQLite:
```text
QUARTER("t"."createat")
```

#### Second
Извлекает секунду (0–59) из выражения даты/времени.
```go
function := uast.Second(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
SECOND(`t`.`createat`)
```
Output MsSQL:
```text
DATEPART(SECOND, [t].[createat])
```
Output MySQL:
```text
SECOND(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(SECOND FROM "t"."createat")
```
Output SQLite:
```text
SECOND("t"."createat")
```

#### TimeAdd
Добавляет интервал времени к выражению времени/даты/времени и возвращает результирующее время.
```go
function := uast.TimeAdd(uast.Column[time.Time]("t", "createat"), uast.Value("2 HOUR"))
```
Output MariaDB:
```text
TIME_ADD(`t`.`createat`, '2 HOUR')
```
Output MsSQL:
```text
// In development
```
Output MySQL:
```text
TIME_ADD(`t`.`createat`, '2 HOUR')
```
Output PostgreSQL:
```text
("t"."createat" + INTERVAL '2 HOUR')
```
Output SQLite:
```text
time("t"."createat", '+2 HOUR')
```

#### TimeDiff
Возвращает разницу между двумя выражениями времени/даты/времени (`timeEnd` - `timeStart`).
```go
function := uast.TimeDiff(uast.Column[time.Time]("t", "updateat"), uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
TIMEDIFF(`t`.`updateat`, `t`.`createat`)
```
Output MsSQL:
```text
TIMEDIFF([t].[updateat], [t].[createat])
```
Output MySQL:
```text
TIMEDIFF(`t`.`updateat`, `t`.`createat`)
```
Output PostgreSQL:
```text
DATE_PART('time', "t"."updateat" - "t"."createat")
```
Output SQLite:
```text
TIMEDIFF("t"."updateat", "t"."createat")
```

#### TimeSub
Вычитает интервал времени из выражения времени/даты/времени и возвращает результирующее время.
```go
function := uast.TimeSub(uast.Column[time.Time]("t", "createat"), uast.Value("2 HOUR"))
```
Output MariaDB:
```text
TIME_SUB(`t`.`createat`, '2 HOUR')
```
Output MsSQL:
```text
// In development
```
Output MySQL:
```text
TIME_SUB(`t`.`createat`, '2 HOUR')
```
Output PostgreSQL:
```text
("t"."createat" - INTERVAL '2 HOUR')
```
Output SQLite:
```text
time("t"."createat", '-2 HOUR')
```

#### Week
Извлекает номер недели (1–53) из выражения даты/времени.
```go
function := uast.Week(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
WEEK(`t`.`createat`)
```
Output MsSQL:
```text
DATEPART(WEEK, [t].[createat])
```
Output MySQL:
```text
WEEK(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(WEEK FROM "t"."createat")
```
Output SQLite:
```text
WEEK("t"."createat")
```

#### Year
Извлекает год из выражения даты/времени.
```go
function := uast.Year(uast.Column[time.Time]("t", "createat"))
```
Output MariaDB:
```text
YEAR(`t`.`createat`)
```
Output MsSQL:
```text
YEAR([t].[createat])
```
Output MySQL:
```text
YEAR(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(YEAR FROM "t"."createat")
```
Output SQLite:
```text
YEAR("t"."createat")
```

### Json
#### JsonArray
Создаёт JSON-массив из заданного выражения и опциональных дополнительных значений.
```go
function := uast.JsonArray(
    uast.Column[string]("t", "json"), 
    uast.Value("val1"), 
    uast.Value("val2"),
)
```
Output MariaDB:
```text
JSON_ARRAY(`t`.`json`, ?, ?)
```
Output MsSQL:
```text
JSON_ARRAY([t].[json], @p1, @p2)
```
Output MySQL:
```text
JSON_ARRAY(`t`.`json`, ?, ?)
```
Output PostgreSQL:
```text
JSON_ARRAY("t"."json", $1, $2)
```
Output SQLite:
```text
JSON_ARRAY("t"."json", ?, ?)
```

#### JsonArrayAgg
Агрегирует значения из группы в JSON-массив.
```go
function := uast.JsonArrayAgg(
    uast.Column[string]("t", "json"),
)
```
Output MariaDB:
```text
JSON_ARRAYAGG(`t`.`json`)
```
Output MsSQL:
```text
JSON_ARRAYAGG([t].[json])
```
Output MySQL:
```text
JSON_ARRAYAGG(`t`.`json`)
```
Output PostgreSQL:
```text
JSON_AGG("t"."json")
```
Output SQLite:
```text
JSON_GROUP_ARRAY("t"."json")
```

#### JsonContains
Проверяет, содержит ли JSON-документ указанное значение.
```go
function := uast.JsonContains(
    uast.Column[string]("t", "json"),
    uast.Value(`{"key":"val"}`),
)
```
Output MariaDB:
```text
JSON_CONTAINS(`t`.`json`, '{"key":"val"}')
```
Output MsSQL:
```text
// In development
```
Output MySQL:
```text
JSON_CONTAINS(`t`.`json`, '{"key":"val"}')
```
Output PostgreSQL:
```text
("t"."json" @> '{"key":"val"}')
```
Output SQLite:
```text
JSON_CONTAINS("t"."json", '{"key":"val"}')
```

#### JsonExtract
Извлекает значение из JSON-документа по указанному пути. Параметр `json` строится с помощью `JsonPath` и опциональных `JsonKey`/`JsonIndex`.
```go
function := JsonExtract(
    uast.Column[string]("t", "json"), 
    uast.JsonGroup(
        uast.JsonPath(
            uast.JsonKey("parent"), 
            uast.JsonIndex(0), 
            uast.JsonKey("child"),
        ),
    ),
    uast.TypeString,
)
```
Output MariaDB:
```text
(`t`.`json` ->> '$.parent[0].child')
```
Output MsSQL:
```text
// In development
```
Output MySQL:
```text
(`t`.`json` ->> '$.parent[0].child')
```
Output PostgreSQL:
```text
("t"."json" #>> '{parent,0,child}')
```
Output SQLite:
```text
("t"."json" ->> '$.parent[0].child')
```

#### JsonObject
Создаёт JSON-объект из пар ключ-значение.
```go
function := uast.JsonObject(
    uast.JsonPair(
        uast.JsonKey("key"), 
        uast.Count(uast.Column[string]("t", "json"), false),
    ),
)
```
Output MariaDB:
```text
JSON_OBJECT('key', COUNT(`t`.`json`))
```
Output MsSQL:
```text
JSON_OBJECT('key', COUNT([t].[json]))
```
Output MySQL:
```text
JSON_OBJECT('key', COUNT(`t`.`json`))
```
Output PostgreSQL:
```text
JSON_BUILD_OBJECT('key', COUNT("t"."json"))
```
Output SQLite:
```text
JSON_OBJECT('key', COUNT("t"."json"))
```

#### JsonObjectAgg
Агрегирует пары ключ-значение из группы в один JSON-объект.
```go
function := uast.JsonObjectAgg(
    uast.Column[string]("t", "json"),
    uast.Column[int]("t", "number"),
)
```
Output MariaDB:
```text
JSON_OBJECTAGG(`t`.`json`, `t`.`number`)
```
Output MsSQL:
```text
JSON_OBJECTAGG([t].[json], [t].[number])
```
Output MySQL:
```text
JSON_OBJECTAGG(`t`.`json`, `t`.`number`)
```
Output PostgreSQL:
```text
JSON_OBJECT_AGG("t"."json", "t"."number")
```
Output SQLite:
```text
JSON_GROUP_OBJECT("t"."json", "t"."number")
```

#### JsonRemove
Удаляет значение из JSON-документа по указанному пути(ям).
```go
function := uast.JsonRemove(
    uast.Column[string]("t", "json"),
    uast.JsonGroup(
        uast.JsonPath(
            uast.JsonKey("key1"),
        ),
    ), 
    uast.JsonGroup(
        uast.JsonPath(
            uast.JsonKey("key2"),
        ),
    ),
)
```
Output MariaDB:
```text
JSON_REMOVE(`t`.`json`, '$.key1', '$.key2')
```
Output MsSQL:
```text
// In development
```
Output MySQL:
```text
JSON_REMOVE(`t`.`json`, '$.key1', '$.key2')
```
Output PostgreSQL:
```text
("t"."json" - '{key1}' - '{key2}')
```
Output SQLite:
```text
JSON_REMOVE("t"."json", '$.key1', '$.key2')
```

#### JsonSet
Устанавливает значение в JSON-документе по указанному пути(ям). Создаёт путь, если он не существует.
```go
function := uast.JsonSet(
    uast.Column[string]("t", "json"),
    uast.JsonGroup(
        uast.JsonPath(
            uast.JsonKey("key1"),
        ), 
        uast.Value("val1"),
    ),
    uast.JsonGroup(
        uast.JsonPath(
            uast.JsonKey("key2"),
        ),
        uast.Value("val2"),
    ),
)
```
Output MariaDB:
```text
JSON_SET(`t`.`json`, '$.key1', ?, '$.key2', ?)
```
Output MsSQL:
```text
// In development
```
Output MySQL:
```text
JSON_SET(`t`.`json`, '$.key1', ?, '$.key2', ?)
```
Output PostgreSQL:
```text
// In development
```
Output SQLite:
```text
JSON_SET("t"."json", '$.key1', ?, '$.key2', ?)
```

#### JsonType
Возвращает тип JSON-значения (например, 'OBJECT', 'ARRAY', 'STRING', 'INTEGER', 'NULL').
```go
function := uast.JsonType(uast.Column[string]("t", "json"))
```
Output MariaDB:
```text
JSON_TYPE(`t`.`json`)
```
Output MsSQL:
```text
// In development
```
Output MySQL:
```text
JSON_TYPE(`t`.`json`)
```
Output PostgreSQL:
```text
jsonb_typeof("t"."json")
```
Output SQLite:
```text
JSON_TYPE("t"."json")
```

### Math
#### Abs
Возвращает абсолютное (неотрицательное) значение числового выражения.
```go
function := uast.Abs(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ABS(`t`.`number`)
```
Output MsSQL:
```text
ABS([t].[number])
```
Output MySQL:
```text
ABS(`t`.`number`)
```
Output PostgreSQL:
```text
ABS("t"."number")
```
Output SQLite:
```text
ABS("t"."number")
```

#### ACos
Возвращает арккосинус (обратный косинус) выражения в радианах.
```go
function := uast.ACos(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ACOS(`t`.`number`)
```
Output MsSQL:
```text
ACOS([t].[number])
```
Output MySQL:
```text
ACOS(`t`.`number`)
```
Output PostgreSQL:
```text
ACOS("t"."number")
```
Output SQLite:
```text
ACOS("t"."number")
```

#### ASin
Возвращает арксинус (обратный синус) выражения в радианах.
```go
function := uast.ASin(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ASIN(`t`.`number`)
```
Output MsSQL:
```text
ASIN([t].[number])
```
Output MySQL:
```text
ASIN(`t`.`number`)
```
Output PostgreSQL:
```text
ASIN("t"."number")
```
Output SQLite:
```text
ASIN("t"."number")
```

#### ATan
Возвращает арктангенс (обратный тангенс) выражения в радианах.
```go
function := uast.ATan(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
ATAN(`t`.`number`)
```
Output MsSQL:
```text
ATAN([t].[number])
```
Output MySQL:
```text
ATAN(`t`.`number`)
```
Output PostgreSQL:
```text
ATAN("t"."number")
```
Output SQLite:
```text
ATAN("t"."number")
```

#### ATan2
Возвращает арктангенс частного двух аргументов (`y`/`x`), используя их знаки для определения квадранта.
```go
function := uast.ATan2(uast.Column[int]("t", "y"), uast.Column[int]("t", "x"))
```
Output MariaDB:
```text
ATAN2(`t`.`y`, `t`.`x`)
```
Output MsSQL:
```text
ATAN2([t].[y], [t].[x])
```
Output MySQL:
```text
ATAN2(`t`.`y`, `t`.`x`)
```
Output PostgreSQL:
```text
ATAN2("t"."y", "t"."x")
```
Output SQLite:
```text
ATAN2("t"."y", "t"."x")
```

#### Cbrt
Возвращает кубический корень числового выражения.
```go
function := uast.Cbrt(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
CBRT(`t`.`number`)
```
Output MsSQL:
```text
CBRT([t].[number])
```
Output MySQL:
```text
CBRT(`t`.`number`)
```
Output PostgreSQL:
```text
CBRT("t"."number")
```
Output SQLite:
```text
CBRT("t"."number")
```

#### Ceil
Возвращает наименьшее целое значение, не меньшее аргумента (округление вверх).
```go
function := uast.Ceil(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
CEILING(`t`.`number`)
```
Output MsSQL:
```text
CEILING([t].[number])
```
Output MySQL:
```text
CEILING(`t`.`number`)
```
Output PostgreSQL:
```text
CEIL("t"."number")
```
Output SQLite:
```text
CEIL("t"."number")
```

#### Cos
Возвращает косинус выражения в радианах.
```go
function := uast.Cos(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
COS(`t`.`number`)
```
Output MsSQL:
```text
COS([t].[number])
```
Output MySQL:
```text
COS(`t`.`number`)
```
Output PostgreSQL:
```text
COS("t"."number")
```
Output SQLite:
```text
COS("t"."number")
```

#### Exp
Возвращает число Эйлера `e` (~2.71828) возведённое в степень выражения.
```go
function := uast.Exp(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
EXP(`t`.`number`)
```
Output MsSQL:
```text
EXP([t].[number])
```
Output MySQL:
```text
EXP(`t`.`number`)
```
Output PostgreSQL:
```text
EXP("t"."number")
```
Output SQLite:
```text
EXP("t"."number")
```

#### Floor
Возвращает наибольшее целое значение, не большее аргумента (округление вниз).
```go
function := uast.Floor(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
FLOOR(`t`.`number`)
```
Output MsSQL:
```text
FLOOR([t].[number])
```
Output MySQL:
```text
FLOOR(`t`.`number`)
```
Output PostgreSQL:
```text
FLOOR("t"."number")
```
Output SQLite:
```text
FLOOR("t"."number")
```

#### Ln
Возвращает натуральный логарифм (по основанию `e`) выражения.
```go
function := uast.Ln(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
LN(`t`.`number`)
```
Output MsSQL:
```text
LN([t].[number])
```
Output MySQL:
```text
LN(`t`.`number`)
```
Output PostgreSQL:
```text
LN("t"."number")
```
Output SQLite:
```text
LN("t"."number")
```

#### Log
Возвращает логарифм выражения по указанному основанию.
```go
function := uast.Log(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
LOG(`t`.`number`, ?)
```
Output MsSQL:
```text
LOG([t].[number], @p1)
```
Output MySQL:
```text
LOG(`t`.`number`, ?)
```
Output PostgreSQL:
```text
LOG("t"."number", $1)
```
Output SQLite:
```text
LOG("t"."number", ?)
```

#### Mod
Возвращает остаток (модуль) от деления первого выражения на второе.
```go
function := uast.Mod(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
MOD(`t`.`number`, ?)
```
Output MsSQL:
```text
MOD([t].[number], @p1)
```
Output MySQL:
```text
MOD(`t`.`number`, ?)
```
Output PostgreSQL:
```text
MOD("t"."number", $1)
```
Output SQLite:
```text
MOD("t"."number", ?)
```

#### Pi
Возвращает математическую константу `π` (~3.14159).
```go
function := uast.Pi()
```
Output MariaDB:
```text
PI()
```
Output MsSQL:
```text
PI()
```
Output MySQL:
```text
PI()
```
Output PostgreSQL:
```text
PI()
```
Output SQLite:
```text
PI()
```

#### Power
Возвращает выражение, возведённое в степень экспоненты.
```go
function := uast.Power(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
POWER(`t`.`number`, ?)
```
Output MsSQL:
```text
POWER([t].[number], @p1)
```
Output MySQL:
```text
POWER(`t`.`number`, ?)
```
Output PostgreSQL:
```text
POWER("t"."number", $1)
```
Output SQLite:
```text
POWER("t"."number", ?)
```

#### Rand
Возвращает случайное значение с плавающей запятой в диапазоне [0, 1].
```go
function := uast.Rand()
```
Output MariaDB:
```text
RAND()
```
Output MsSQL:
```text
RAND()
```
Output MySQL:
```text
RAND()
```
Output PostgreSQL:
```text
RANDOM()
```
Output SQLite:
```text
RANDOM()
```

#### Round
Округляет выражение до указанного количества знаков после запятой.
```go
function := uast.Round(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
ROUND(`t`.`number`, ?)
```
Output MsSQL:
```text
ROUND([t].[number], @p1)
```
Output MySQL:
```text
ROUND(`t`.`number`, ?)
```
Output PostgreSQL:
```text
ROUND("t"."number", $1)
```
Output SQLite:
```text
ROUND("t"."number", ?)
```

#### Sin
Возвращает синус выражения в радианах.
```go
function := uast.Sin(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
SIN(`t`.`number`)
```
Output MsSQL:
```text
SIN([t].[number])
```
Output MySQL:
```text
SIN(`t`.`number`)
```
Output PostgreSQL:
```text
SIN("t"."number")
```
Output SQLite:
```text
SIN("t"."number")
```

#### Sqrt
Возвращает квадратный корень выражения.
```go
function := uast.Sqrt(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
SQRT(`t`.`number`)
```
Output MsSQL:
```text
SQRT([t].[number])
```
Output MySQL:
```text
SQRT(`t`.`number`)
```
Output PostgreSQL:
```text
SQRT("t"."number")
```
Output SQLite:
```text
SQRT("t"."number")
```

#### Tan
Возвращает тангенс выражения в радианах.
```go
function := uast.Tan(uast.Column[int]("t", "number"))
```
Output MariaDB:
```text
TAN(`t`.`number`)
```
Output MsSQL:
```text
TAN([t].[number])
```
Output MySQL:
```text
TAN(`t`.`number`)
```
Output PostgreSQL:
```text
TAN("t"."number")
```
Output SQLite:
```text
TAN("t"."number")
```

#### Trunc
Усекает числовое выражение до указанного количества знаков после запятой (без округления).
```go
function := uast.Trunc(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MariaDB:
```text
TRUNCATE(`t`.`number`, ?)
```
Output MsSQL:
```text
ROUND([t].[number], @p1, 1)
```
Output MySQL:
```text
TRUNCATE(`t`.`number`, ?)
```
Output PostgreSQL:
```text
TRUNC("t"."number", $1)
```
Output SQLite:
```text
TRUNC("t"."number", ?)
```

### Ranking
#### CumeDist
Возвращает кумулятивное распределение значения в рамках раздела (отношение строк, которые идут до или равны текущей строке). Должна использоваться с оператором `OVER`.
```go
function := uast.CumeDist().Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
CUME_DIST() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
CUME_DIST() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
CUME_DIST() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
CUME_DIST() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
CUME_DIST() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### DenseRank
Возвращает ранг строки без пропусков. Строки с равными значениями получают одинаковый ранг, а следующий ранг является непосредственно следующим целым числом. Требует `OVER`.
```go
function := uast.DenseRank().Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
DENSE_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
DENSE_RANK() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
DENSE_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
DENSE_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
DENSE_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### NTile
Делит строки в рамках раздела на `n` приблизительно равных групп и возвращает номер группы (от 1 до `n`) для каждой строки.
```go
function := uast.NTile(2).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
NTILE(2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
NTILE(2) OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
NTILE(2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
NTILE(2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
NTILE(2) OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### PercentRank
Возвращает процентильный ранг строки в рамках раздела (диапазон от 0 до 1). Ранг первой строки всегда равен 0. Требует `OVER`.
```go
function := uast.PercentRank().Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
PERCENT_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
PERCENT_RANK() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
PERCENT_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
PERCENT_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
PERCENT_RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### Rank
Возвращает ранг строки с пропусками. Равные значения получают одинаковый ранг, а следующее отличное значение пропускает ранги. Требует `OVER`.
```go
function := uast.Rank().Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
RANK() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
RANK() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

#### RowNumber
Присваивает уникальный последовательный номер каждой строке в рамках раздела, начиная с 1. Порядок определяет последовательность нумерации.
```go
function := uast.RowNumber().Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MariaDB:
```text
ROW_NUMBER() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output MsSQL:
```text
ROW_NUMBER() OVER (PARTITION BY [t].[id] ORDER BY [t].[number] DESC)
```
Output MySQL:
```text
ROW_NUMBER() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
ROW_NUMBER() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```
Output SQLite:
```text
ROW_NUMBER() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

### String
#### Concat
Объединяет два или более строковых выражения в одну строку. Аргументы `NULL` рассматриваются как пустые строки в большинстве диалектов.
```go
function := uast.Concat(uast.Column[string]("t", "string"), uast.Value("old"), uast.Value("new"))
```
Output MariaDB:
```text
CONCAT(`t`.`string`, ?, ?)
```
Output MsSQL:
```text
CONCAT([t].[string], @p1, @p2)
```
Output MySQL:
```text
CONCAT(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
CONCAT("t"."string", $1, $2)
```
Output SQLite:
```text
CONCAT("t"."string", ?, ?)
```

#### ConcatWs
Объединяет два или более строковых выражения с указанным разделителем между ними. Пропускает аргументы `NULL`.
```go
function := uast.ConcatWs(uast.Value("_"), uast.Column[string]("t", "string"), uast.Value("old"),uast.Value("new"))
```
Output MariaDB:
```text
CONCAT_WS(?, `t`.`string`, ?, ?)
```
Output MsSQL:
```text
CONCAT_WS(@p1, [t].[string], @p2, @p3)
```
Output MySQL:
```text
CONCAT_WS(?, `t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
CONCAT_WS($1, "t"."string", $2, $3)
```
Output SQLite:
```text
CONCAT_WS(?, "t"."string", ?, ?)
```

#### LeftString
Возвращает крайние слева `count` символов из строкового выражения.
```go
function := uast.LeftString(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MariaDB:
```text
LEFT(`t`.`string`, ?)
```
Output MsSQL:
```text
LEFT([t].[string], @p1)
```
Output MySQL:
```text
LEFT(`t`.`string`, ?)
```
Output PostgreSQL:
```text
LEFT("t"."string", $1)
```
Output SQLite:
```text
LEFT("t"."string", ?)
```

#### Lower
Преобразует строковое выражение в нижний регистр.
```go
function := uast.Lower(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
LOWER(`t`.`string`)
```
Output MsSQL:
```text
LOWER([t].[string])
```
Output MySQL:
```text
LOWER(`t`.`string`)
```
Output PostgreSQL:
```text
LOWER("t"."string")
```
Output SQLite:
```text
LOWER("t"."string")
```

#### LPad
Дополняет строковое выражение слева указанным разделителем до общей длины `count` символов.
```go
function := uast.LPad(uast.Column[string]("t", "string"), uast.Value(2), uast.Value(","))
```
Output MariaDB:
```text
LPAD(`t`.`string`, ?, ?)
```
Output MsSQL:
```text
LPAD([t].[string], @p1, @p2)
```
Output MySQL:
```text
LPAD(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
LPAD("t"."string", $1, $2)
```
Output SQLite:
```text
LPAD("t"."string", ?, ?)
```

#### LTrim
Удаляет начальные пробелы из строкового выражения.
```go
function := uast.LTrim(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
LTRIM(`t`.`string`)
```
Output MsSQL:
```text
LTRIM([t].[string])
```
Output MySQL:
```text
LTRIM(`t`.`string`)
```
Output PostgreSQL:
```text
LTRIM("t"."string")
```
Output SQLite:
```text
LTRIM("t"."string")
```

#### Repeat
Повторяет строковое выражение `count` раз.
```go
function := uast.Repeat(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MariaDB:
```text
REPEAT(`t`.`string`, ?)
```
Output MsSQL:
```text
REPEAT([t].[string], @p1)
```
Output MySQL:
```text
REPEAT(`t`.`string`, ?)
```
Output PostgreSQL:
```text
REPEAT("t"."string", $1)
```
Output SQLite:
```text
REPEAT("t"."string", ?)
```

#### Replace
Заменяет все вхождения подстроки в строке на новую подстроку.
```go
function := uast.Replace(uast.Column[string]("t", "string"), uast.Value("old"), uast.Value("new"))
```
Output MariaDB:
```text
REPLACE(`t`.`string`, ?, ?)
```
Output MsSQL:
```text
REPLACE([t].[string], @p1, @p2)
```
Output MySQL:
```text
REPLACE(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
REPLACE("t"."string", $1, $2)
```
Output SQLite:
```text
REPLACE("t"."string", ?, ?)
```

#### Reverse
Переворачивает символы в строковом выражении.
```go
function := uast.Reverse(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
REVERSE(`t`.`string`)
```
Output MsSQL:
```text
REVERSE([t].[string])
```
Output MySQL:
```text
REVERSE(`t`.`string`)
```
Output PostgreSQL:
```text
REVERSE("t"."string")
```
Output SQLite:
```text
REVERSE("t"."string")
```

#### RightString
Возвращает крайние справа `count` символов из строкового выражения.
```go
function := uast.RightString(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MariaDB:
```text
RIGHT(`t`.`string`, ?)
```
Output MsSQL:
```text
RIGHT([t].[string], @p1)
```
Output MySQL:
```text
RIGHT(`t`.`string`, ?)
```
Output PostgreSQL:
```text
RIGHT("t"."string", $1)
```
Output SQLite:
```text
RIGHT("t"."string", ?)
```

#### RPad
Дополняет строковое выражение справа указанным разделителем до общей длины `count` символов.
```go
function := uast.RPad(uast.Column[string]("t", "string"), uast.Value(2), uast.Value(","))
```
Output MariaDB:
```text
RPAD(`t`.`string`, ?, ?)
```
Output MsSQL:
```text
RPAD([t].[string], @p1, @p2)
```
Output MySQL:
```text
RPAD(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
RPAD("t"."string", $1, $2)
```
Output SQLite:
```text
RPAD("t"."string", ?, ?)
```

#### RTrim
Удаляет конечные пробелы из строкового выражения.
```go
function := uast.RTrim(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
RTRIM(`t`.`string`)
```
Output MsSQL:
```text
RTRIM([t].[string])
```
Output MySQL:
```text
RTRIM(`t`.`string`)
```
Output PostgreSQL:
```text
RTRIM("t"."string")
```
Output SQLite:
```text
RTRIM("t"."string")
```

#### SubString
Извлекает подстроку из строкового выражения, начиная с `startPos` (начиная с 1) длиной `lengthStr` символов.
```go
function := uast.SubString(uast.Column[string]("t", "string"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
SUBSTRING(`t`.`string`, ?, ?)
```
Output MsSQL:
```text
SUBSTRING([t].[string], @p1, @p2)
```
Output MySQL:
```text
SUBSTRING(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
SUBSTRING("t"."string", $1, $2)
```
Output SQLite:
```text
SUBSTRING("t"."string", ?, ?)
```

#### Trim
Удаляет как начальные, так и конечные пробелы из строкового выражения.
```go
function := uast.Trim(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
TRIM(`t`.`string`)
```
Output MsSQL:
```text
TRIM([t].[string])
```
Output MySQL:
```text
TRIM(`t`.`string`)
```
Output PostgreSQL:
```text
TRIM("t"."string")
```
Output SQLite:
```text
TRIM("t"."string")
```

#### Upper
Преобразует строковое выражение в верхний регистр.
```go
function := uast.Upper(uast.Column[string]("t", "string"))
```
Output MariaDB:
```text
UPPER(`t`.`string`)
```
Output MsSQL:
```text
UPPER([t].[string])
```
Output MySQL:
```text
UPPER(`t`.`string`)
```
Output PostgreSQL:
```text
UPPER("t"."string")
```
Output SQLite:
```text
UPPER("t"."string")
```

## exprLiteral
### Literal
Встраивает необработанное литеральное значение непосредственно в сгенерированную SQL-строку (не параметризуется). Используйте с осторожностью — значения записываются как есть. Предпочитайте `Value` для пользовательских данных.
```go
literal := uast.Literal("%Y-%m-%d")
```
Output:
```text
'%Y-%m-%d'
```

## exprLogical
### And
Комбинирует несколько условий логическим `AND`. Все условия должны быть истинными для истинности комбинированного выражения.
```go
logical := uast.And(
    uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    uast.Greater(uast.Column[int]("t", "number"), uast.Value(2)),
)
```
Output MariaDB:
```text
(`t`.`string` = ? AND `t`.`number` > ?)
```
Output MsSQL:
```text
([t].[string] = @p1 AND [t].[number] > @p2)
```
Output MySQL:
```text
(`t`.`string` = ? AND `t`.`number` > ?)
```
Output PostgreSQL:
```text
("t"."string" = $1 AND "t"."number" > $2)
```
Output SQLite:
```text
("t"."string" = ? AND "t"."number" > ?)
```

### Or
Комбинирует несколько условий логическим `OR`. Хотя бы одно условие должно быть истинным для истинности комбинированного выражения.
```go
logical := uast.Or(
    uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    uast.Greater(uast.Column[int]("t", "number"), uast.Value(2)),
)
```
Output MariaDB:
```text
(`t`.`string` = ? OR `t`.`number` > ?)
```
Output MsSQL:
```text
([t].[string] = @p1 OR [t].[number] > @p2)
```
Output MySQL:
```text
(`t`.`string` = ? OR `t`.`number` > ?)
```
Output PostgreSQL:
```text
("t"."string" = $1 OR "t"."number" > $2)
```
Output SQLite:
```text
("t"."string" = ? OR "t"."number" > ?)
```

## exprSubquery
### Subquery
Оборачивает оператор `SELECT` как типизированное выражение, которое может использоваться в сравнениях (`In`, `Exists`, `Equal` и т.д.) или как колонка в операторе `SELECT`. Обобщённый параметр `T` указывает скалярный тип единственной колонки, возвращаемой подзапросом.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Column[int64]("t", "id")).From(uast.NewTable("test").As("t")))
```
Output MariaDB:
```text
(SELECT `t`.`id` FROM `test` AS `t`)
```
Output MsSQL:
```text
(SELECT [t].[id] FROM [test] AS [t])
```
Output MySQL:
```text
(SELECT `t`.`id` FROM `test` AS `t`)
```
Output PostgreSQL:
```text
(SELECT "t"."id" FROM "test" AS "t")
```
Output SQLite:
```text
(SELECT "t"."id" FROM "test" AS "t")
```

## exprValue
### Value
Оборачивает Go-значение как параметризованное выражение. Значение НЕ вставляется в SQL-строку напрямую — вместо этого генерируется плейсхолдер (`?`, `$1`, и т.д.), а значение добавляется в слайс аргументов, возвращаемый `Build()`. Это безопасный способ передачи пользовательских данных, предотвращающий SQL-инъекции. Поддерживаемые типы: `bool`, `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `string`, `time.Time`.
```go
var data string = "ivan"
value := uast.Value(data)
```
Output MariaDB:
```text
?
```
Output MsSQL:
```text
@p1
```
Output MySQL:
```text
?
```
Output PostgreSQL:
```text
$1
```
Output SQLite:
```text
?
```