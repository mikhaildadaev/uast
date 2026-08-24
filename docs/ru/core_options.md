---
outline: deep
---

# API / Core / Опции

::: info **Информация**
Эта страница охватывает все параметры конфигурации: `clauseGroupBy`, `clauseHaving`, `clauseJoin`, `clauseOrderBy`, `clausePagination`, `clauseReturning`, `clauseSet`, `clauseUnions`, `clauseValues`, `clauseWhere`, `clauseWith`, `exprArray`, `exprBinary`, `exprComparison`, `exprConstant`, `exprField`, `exprFunction`, `exprLiteral`, `exprLogical`, `exprSubquery`, `exprValue`. Каждый параметр показан с рабочим примером кода и ожидаемым выводом.
:::

## clauseGroupBy
Добавляет оператор GROUP BY для группировки строк по указанным колонкам или выражениям.
```go
groupBy := GroupBy(
	uast.Field[string]("u", "string"),
)
```
Output MariaDB:
```text
GROUP BY `u`.`string`
```
Output MsSQL:
```text
GROUP BY [u].[string]
```
Output MySQL:
```text
GROUP BY `u`.`string`
```
Output PostgreSQL:
```text
GROUP BY "u"."string"
```
Output SQLite:
```text
GROUP BY "u"."string"
```

## clauseHaving
Добавляет оператор HAVING для фильтрации групп. Используется с GROUP BY для фильтрации агрегированных результатов.
```go
having := Having(
	uast.Greater(uast.Count(uast.Field[int64]("u", "id"), false), uast.Value[int64](2)),
)
```
Output MariaDB:
```text
HAVING COUNT(`u`.`id`) > ?
```
Output MsSQL:
```text
HAVING COUNT([u].[id]) > @p1
```
Output MySQL:
```text
HAVING COUNT(`u`.`id`) > ?
```
Output PostgreSQL:
```text
HAVING COUNT("u"."id") > $1
```
Output SQLite:
```text
HAVING COUNT("u"."id") > ?
```

## clauseJoin
### Cross
Добавляет CROSS JOIN к запросу. Возвращает декартово произведение обеих таблиц.
```go
join := uast.Cross(uast.NewTable("users").As("u"))
```
Output MariaDB:
```text
CROSS JOIN `users` AS `u`
```
Output MsSQL:
```text
CROSS JOIN [users] AS [u]
```
Output MySQL:
```text
CROSS JOIN `users` AS `u`
```
Output PostgreSQL:
```text
CROSS JOIN "users" AS "u"
```
Output SQLite:
```text
CROSS JOIN "users" AS "u"
```

### Full
Добавляет FULL JOIN к запросу. Возвращает все строки из обеих таблиц, с NULL там, где нет совпадений.
```go
join := uast.Full(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
FULL JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
FULL JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
FULL JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
FULL JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
FULL JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### FullOuter
Добавляет FULL OUTER JOIN к запросу. Возвращает все строки из обеих таблиц, с NULL там, где нет совпадений.
```go
join := uast.FullOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
FULL OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
FULL OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
FULL OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
FULL OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
FULL OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Inner
Добавляет INNER JOIN к запросу. Возвращает строки, имеющие совпадающие значения в обеих таблицах.
```go
join := uast.Inner(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
INNER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
INNER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
INNER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
INNER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
INNER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Left
Добавляет LEFT JOIN к запросу. Возвращает все строки из левой таблицы и совпадающие строки из правой таблицы.
```go
join := uast.Left(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
LEFT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
LEFT JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
LEFT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
LEFT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
LEFT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### LeftOuter
Добавляет LEFT OUTER JOIN к запросу. Возвращает все строки из левой таблицы и совпадающие строки из правой таблицы.
```go
join := uast.LeftOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
LEFT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
LEFT OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
LEFT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
LEFT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
LEFT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Right
Добавляет RIGHT JOIN к запросу. Возвращает все строки из правой таблицы и совпадающие строки из левой таблицы. Не поддерживается SQLite.
```go
join := uast.Right(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
RIGHT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
RIGHT JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
RIGHT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
RIGHT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
// Not supported
```

### RightOuter
Добавляет RIGHT OUTER JOIN к запросу. Возвращает все строки из правой таблицы и совпадающие строки из левой таблицы. Не поддерживается SQLite.
```go
join := uast.RightOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
Output MariaDB:
```text
RIGHT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output MsSQL:
```text
RIGHT OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
Output MySQL:
```text
RIGHT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
Output PostgreSQL:
```text
RIGHT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
Output SQLite:
```text
// Not supported
```

## clauseOrderBy
### Asc
Указывает порядок сортировки по возрастанию (сначала наименьшие, от А до Я). Используется для сортировки строк в запросе или в рамках оконной функции.
```go
orderBy := uast.Asc(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
`u`.`string` ASC
```
Output MsSQL:
```text
[u].[string] ASC
```
Output MySQL:
```text
`u`.`string` ASC
```
Output PostgreSQL:
```text
"u"."string" ASC
```
Output SQLite:
```text
"u"."string" ASC
```

### Desc
Указывает порядок сортировки по убыванию (сначала наибольшие, от Я до А). Используется для сортировки строк в запросе или в рамках оконной функции.
```go
orderBy := uast.Desc(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
`u`.`string` DESC
```
Output MsSQL:
```text
[u].[string] DESC
```
Output MySQL:
```text
`u`.`string` DESC
```
Output PostgreSQL:
```text
"u"."string" DESC
```
Output SQLite:
```text
"u"."string" DESC
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
	uast.Field[int64]("u", "id"),
    uast.Field[string]("u", "string"),
)
```
Output MariaDB:
```text
RETURNING `u`.`id`, `u`.`string`
```
Output MsSQL:
```text
OUTPUT [u].[id], [u].[string]
```
Output MySQL:
```text
// Not support
```
Output PostgreSQL:
```text
RETURNING "u"."id", "u"."string"
```
Output SQLite:
```text
RETURNING "u"."id", "u"."string"
```

## clauseSet
### Assign
Указывает колонки и их новые значения с помощью `Assign` для связывания колонок со значениями. Поддерживает несколько пар для обновления нескольких колонок.
```go
set := Set(
	uast.Assign(uast.Field[string]("u", "string"), uast.Value("active")),
)
```
Output MariaDB:
```text
UPDATE `users` AS `u` SET `u`.`string` = ?
```
Output MsSQL:
```text
UPDATE [users] AS [u] SET [u].[string] = @p1
```
Output MySQL:
```text
UPDATE `users` AS `u` SET `u`.`string` = ?
```
Output PostgreSQL:
```text
UPDATE "users" AS "u" SET "u"."string" = $1
```
Output SQLite:
```text
UPDATE "users" AS "u" SET "u"."string" = ?
```

## clauseUnions
### Union
Объединяет результаты нескольких операторов SELECT. UNION возвращает уникальные строки.
```go
unions := uast.Union(uast.NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[string]("u", "string"),
    ),
)
```
Output MariaDB:
```text
UNION SELECT `u`.`string` FROM `users` AS `u` 
```
Output MsSQL:
```text
UNION SELECT [u].[string] FROM [users] AS [u]
```
Output MySQL:
```text
UNION SELECT `u`.`string` FROM `users` AS `u`
```
Output PostgreSQL:
```text
UNION SELECT "u"."string" FROM "users" AS "u"
```
Output SQLite:
```text
UNION SELECT "u"."string" FROM "users" AS "u"
```

### UnionAll
Объединяет результаты нескольких операторов SELECT. UNION ALL возвращает все строки, включая дубликаты.
```go
unions := uast.UnionAll(uast.NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[string]("u", "string"),
    ),
)
```
Output MariaDB:
```text
UNION ALL SELECT `u`.`string` FROM `users` AS `u`
```
Output MsSQL:
```text
UNION ALL SELECT [u].[string] FROM [users] AS [u]
```
Output MySQL:
```text
UNION ALL SELECT `u`.`string` FROM `users` AS `u`
```
Output PostgreSQL:
```text
UNION ALL SELECT "u"."string" FROM "users" AS "u"
```
Output SQLite:
```text
UNION ALL SELECT "u"."string" FROM "users" AS "u"
```

### UnionExcept
Объединяет результаты нескольких операторов SELECT. EXCEPT возвращает уникальные строки из первого запроса, которых нет во втором.
```go
unions := uast.UnionExcept(uast.NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[string]("u", "string"),
    ),
)
```
Output MariaDB:
```text
EXCEPT SELECT `u`.`string` FROM `users` AS `u`
```
Output MsSQL:
```text
EXCEPT SELECT [u].[string] FROM [users] AS [u]
```
Output MySQL:
```text
EXCEPT SELECT `u`.`string` FROM `users` AS `u`
```
Output PostgreSQL:
```text
EXCEPT SELECT "u"."string" FROM "users" AS "u"
```
Output SQLite:
```text
EXCEPT SELECT "u"."string" FROM "users" AS "u"
```

### UnionIntersect
Объединяет результаты нескольких операторов SELECT. INTERSECT возвращает уникальные строки, общие для обоих запросов.
```go
unions := uast.UnionIntersect(uast.NewSelect(uast.NewTable("users").As("u")).
	Field(
		uast.Field[string]("u", "string"),
	),
)
```
Output MariaDB:
```text
INTERSECT SELECT `u`.`string` FROM `users` AS `u`
```
Output MsSQL:
```text
INTERSECT SELECT [u].[string] FROM [users] AS [u]
```
Output MySQL:
```text
INTERSECT SELECT `u`.`string` FROM `users` AS `u`
```
Output PostgreSQL:
```text
INTERSECT SELECT "u"."string" FROM "users" AS "u"
```
Output SQLite:
```text
INTERSECT SELECT "u"."string" FROM "users" AS "u"
```

## clauseValues
### Pair
Указывает значения для вставки с помощью `Pair` для связывания колонок со значениями. Колонки автоматически определяются из пар.
```go
values := Values(
    uast.Pair(uast.Field[string]("u", "string"), uast.Value("ivan")),
	uast.Pair(uast.Field[int]("u", "number"), uast.Value(2)),
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
    uast.Pair(uast.Field[string]("u", "string"), uast.Value("ivan")),
	uast.Pair(uast.Field[int]("u", "number"), uast.Value(2)),
).
Upsert(
    uast.Pair(uast.Field[string]("u", "string"), uast.Value("updated")),
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
	uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
)
```
Output MariaDB:
```text
WHERE `u`.`string` = ?
```
Output MsSQL:
```text
WHERE [u].[string] = @p1
```
Output MySQL:
```text
WHERE `u`.`string` = ?
```
Output PostgreSQL:
```text
WHERE "u"."string" = $1
```
Output SQLite:
```text
WHERE "u"."string" = ?
```

## clauseWith
### Norecursive
Добавляет нерекурсивное общее табличное выражение (CTE) к запросу с помощью `WithN`. Колонки получают псевдонимы через вариативные строковые аргументы.
```go
with := WithN("cte_norecursive", NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[int64]("u", "id"),
        uast.Field[string]("u", "string"),
    ).
    Where(
        uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    ),
    "id", "string",
)
```
Output MariaDB:
```text
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ?)
```
Output MsSQL:
```text
WITH [cte_norecursive] ([id], [string]) AS (SELECT [u].[id], [u].[string] FROM [users] AS [u] WHERE [u].[string] = @p1)
```
Output MySQL:
```text
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ?)
```
Output PostgreSQL:
```text
WITH "cte_norecursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = $1)
```
Output SQLite:
```text
WITH "cte_norecursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = ?)
```

### Recursive
Добавляет рекурсивное общее табличное выражение (CTE) к запросу с помощью `WithR`. Требует оператор `Unions` с `UnionAll` для определения рекурсивного шага.
```go
with := WithR("cte_recursive", NewSelect(uast.NewTable("users").As("u")).
    Field(
        uast.Field[int64]("u", "id"),
        uast.Field[string]("u", "string"),
    ).
    Where(
        uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    ).
    Unions(
        uast.UnionAll(uast.NewSelect(uast.NewTable("users").As("u")).
            Field(
                uast.Field[int64]("u", "id"),
                uast.Field[string]("u", "string"),
            ).
            Join(
                uast.Inner(uast.NewCTE("cte_recursive", "rec"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("rec", "id"))),
            ),
        ),
    ),
    "id", "string",
)
```
Output MariaDB:
```text
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ? UNION ALL SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` INNER JOIN `cte_recursive` AS `rec` ON `u`.`id` = `rec`.`id`)
```
Output MsSQL:
```text
WITH RECURSIVE [cte_recursive] ([id], [string]) AS (SELECT [u].[id], [u].[string] FROM [users] AS [u] WHERE [u].[string] = @p1 UNION ALL SELECT [u].[id], [u].[string] FROM [users] AS [u] INNER JOIN [cte_recursive] AS [rec] ON [u].[id] = [rec].[id])
```
Output MySQL:
```text
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ? UNION ALL SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` INNER JOIN `cte_recursive` AS `rec` ON `u`.`id` = `rec`.`id`)
```
Output PostgreSQL:
```text
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = $1 UNION ALL SELECT "u"."id", "u"."string" FROM "users" AS "u" INNER JOIN "cte_recursive" AS "rec" ON "u"."id" = "rec"."id")
```
Output SQLite:
```text
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = ? UNION ALL SELECT "u"."id", "u"."string" FROM "users" AS "u" INNER JOIN "cte_recursive" AS "rec" ON "u"."id" = "rec"."id")
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
binary := uast.BitwiseAnd(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
Output MsSQL:
```text
[u].[number] & @p1
```
Output MySQL:
```text
`u`.`number` & ?
```
Output PostgreSQL:
```text
"u"."number" & $1
```

### BitwiseOr
Выполняет побитовую операцию ИЛИ между двумя выражениями.
```go
binary := uast.BitwiseOr(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`u`.`number` | ?
```
Output MsSQL:
```text
[u].[number] | @p1
```
Output MySQL:
```text
`u`.`number` | ?
```
Output PostgreSQL:
```text
"u"."number" | $1
```
Output SQLite:
```text
"u"."number" | ?
```

### BitwiseXor
Выполняет побитовую операцию исключающего ИЛИ между двумя выражениями.
```go
binary := uast.BitwiseXor(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
Output MariaDB:
```text
`u`.`number` ^ ?
```
Output MsSQL:
```text
[u].[number] ^ @p1
```
Output MySQL:
```text
`u`.`number` ^ ?
```
Output PostgreSQL:
```text
"u"."number" ^ $1
```
Output SQLite:
```text
"u"."number" ^ ?
```

### Divide
Делит левое выражение на правое.
```go
binary := uast.Divide(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` / ?
```
Output MsSQL:
```text
[u].[number] / @p1
```
Output MySQL:
```text
`u`.`number` / ?
```
Output PostgreSQL:
```text
"u"."number" / $1
```
Output SQLite:
```text
"u"."number" / ?
```

### Minus
Вычитает правое выражение из левого.
```go
binary := uast.Minus(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` - ?
```
Output MsSQL:
```text
[u].[number] - @p1
```
Output MySQL:
```text
`u`.`number` - ?
```
Output PostgreSQL:
```text
"u"."number" - $1
```
Output SQLite:
```text
"u"."number" - ?
```

### Modulo
Возвращает остаток от деления левого выражения на правое.
```go
binary := uast.Modulo(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` % ?
```
Output MsSQL:
```text
[u].[number] % @p1
```
Output MySQL:
```text
`u`.`number` % ?
```
Output PostgreSQL:
```text
"u"."number" % $1
```
Output SQLite:
```text
"u"."number" % ?
```

### Multiply
Умножает левое выражение на правое.
```go
binary := uast.Multiply(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` * ?
```
Output MsSQL:
```text
[u].[number] * @p1
```
Output MySQL:
```text
`u`.`number` * ?
```
Output PostgreSQL:
```text
"u"."number" * $1
```
Output SQLite:
```text
"u"."number" * ?
```

### Plus
Складывает левое выражение с правым.
```go
binary := uast.Plus(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` + ?
```
Output MsSQL:
```text
[u].[number] + @p1
```
Output MySQL:
```text
`u`.`number` + ?
```
Output PostgreSQL:
```text
"u"."number" + $1
```
Output SQLite:
```text
"u"."number" + ?
```

### ShiftLeft
Выполняет побитовый сдвиг влево левого выражения на количество бит, указанное в правом выражении.
```go
binary := uast.ShiftLeft(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` << ?
```
Output MsSQL:
```text
[u].[number] << @p1
```
Output MySQL:
```text
`u`.`number` << ?
```
Output PostgreSQL:
```text
"u"."number" << $1
```
Output SQLite:
```text
"u"."number" << ?
```

### ShiftRight
Выполняет побитовый сдвиг вправо левого выражения на количество бит, указанное в правом выражении.
```go
binary := uast.ShiftRight(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` >> ?
```
Output MsSQL:
```text
[u].[number] >> @p1
```
Output MySQL:
```text
`u`.`number` >> ?
```
Output PostgreSQL:
```text
"u"."number" >> $1
```
Output SQLite:
```text
"u"."number" >> ?
```

## exprComparison
### Between
Проверяет, попадает ли левое выражение в диапазон, заданный valueStart и valueEnd (включительно).
```go
comparison := uast.Between(uast.Field[int]("u", "number"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` BETWEEN ? AND ?
```
Output MsSQL:
```text
[u].[number] BETWEEN @p1 AND @p2
```
Output MySQL:
```text
`u`.`number` BETWEEN ? AND ?
```
Output PostgreSQL:
```text
"u"."number" BETWEEN $1 AND $2
```
Output SQLite:
```text
"u"."number" BETWEEN ? AND ?
```

### Equal
Сравнивает два выражения на равенство (`=`).
```go
comparison := uast.Equal(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` = ?
```
Output MsSQL:
```text
[u].[number] = @p1
```
Output MySQL:
```text
`u`.`number` = ?
```
Output PostgreSQL:
```text
"u"."number" = $1
```
Output SQLite:
```text
"u"."number" = ?
```

### Exists
Проверяет, возвращает ли подзапрос какие-либо строки. Возвращает `true` если существует хотя бы одна строка.
```go
comparison := uast.Exists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("users").As("u"))))
```
Output MariaDB:
```text
EXISTS (SELECT 1 FROM `users` AS `u`)
```
Output MsSQL:
```text
EXISTS (SELECT 1 FROM [users] AS [u])
```
Output MySQL:
```text
EXISTS (SELECT 1 FROM `users` AS `u`)
```
Output PostgreSQL:
```text
EXISTS (SELECT 1 FROM "users" AS "u")
```
Output SQLite:
```text
EXISTS (SELECT 1 FROM "users" AS "u")
```

### Greater
Сравнивает, больше ли левое выражение правого (`>`).
```go
comparison := uast.Greater(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` > ?
```
Output MsSQL:
```text
[u].[number] > @p1
```
Output MySQL:
```text
`u`.`number` > ?
```
Output PostgreSQL:
```text
"u"."number" > $1
```
Output SQLite:
```text
"u"."number" > ?
```

### GreaterEqual
Сравнивает, больше или равно ли левое выражение правому (`>=`).
```go
comparison := uast.GreaterEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` >= ?
```
Output MsSQL:
```text
[u].[number] >= @p1
```
Output MySQL:
```text
`u`.`number` >= ?
```
Output PostgreSQL:
```text
"u"."number" >= $1
```
Output SQLite:
```text
"u"."number" >= ?
```

### ILike
Выполняет регистронезависимое сравнение с шаблоном. Правое выражение должно содержать шаблон с `%` (любая последовательность) и `_` (один символ).
```go
comparison := uast.ILike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
LOWER(`u`.`string`) LIKE LOWER(?)
```
Output MsSQL:
```text
LOWER([u].[string]) LIKE LOWER(@p1)
```
Output MySQL:
```text
LOWER(`u`.`string`) LIKE LOWER(?)
```
Output PostgreSQL:
```text
"u"."string" ILIKE $1
```
Output SQLite:
```text
LOWER("u"."string") LIKE LOWER(?)
```

### In
Проверяет, соответствует ли левое выражение любому значению, содержащемуся в правом выражении (обычно подзапрос или массив).
```go
comparison := uast.In(uast.Field[string]("u", "string"), uast.Array("active", "pending"))
```
Output MariaDB:
```text
`u`.`string` IN (?, ?)
```
Output MsSQL:
```text
[u].[string] IN (@p1, @p2)
```
Output MySQL:
```text
`u`.`string` IN (?, ?)
```
Output PostgreSQL:
```text
"u"."string" IN ($1, $2)
```
Output SQLite:
```text
"u"."string" IN (?, ?)
```

### IsNotNull
Проверяет, что выражение не `NULL`.
```go
comparison := uast.IsNotNull(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
`u`.`string` IS NOT NULL
```
Output MsSQL:
```text
[u].[string] IS NOT NULL
```
Output MySQL:
```text
`u`.`string` IS NOT NULL
```
Output PostgreSQL:
```text
"u"."string" IS NOT NULL
```
Output SQLite:
```text
"u"."string" IS NOT NULL
```

### IsNull
Проверяет, что выражение является `NULL`.
```go
comparison := uast.IsNull(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
`u`.`string` IS NULL
```
Output MsSQL:
```text
[u].[string] IS NULL
```
Output MySQL:
```text
`u`.`string` IS NULL
```
Output PostgreSQL:
```text
"u"."string" IS NULL
```
Output SQLite:
```text
"u"."string" IS NULL
```

### Less
Сравнивает, меньше ли левое выражение правого (`<`).
```go
comparison := uast.Less(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` < ?
```
Output MsSQL:
```text
[u].[number] < @p1
```
Output MySQL:
```text
`u`.`number` < ?
```
Output PostgreSQL:
```text
"u"."number" < $1
```
Output SQLite:
```text
"u"."number" < ?
```

### LessEqual
Сравнивает, меньше или равно ли левое выражение правому (`<=`).
```go
comparison := uast.LessEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` <= ?
```
Output MsSQL:
```text
[u].[number] <= @p1
```
Output MySQL:
```text
`u`.`number` <= ?
```
Output PostgreSQL:
```text
"u"."number" <= $1
```
Output SQLite:
```text
"u"."number" <= ?
```

### Like
Выполняет регистрозависимое сравнение с шаблоном. Правое выражение должно содержать шаблон с `%` и `_`.
```go
comparison := uast.Like(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
`u`.`string` LIKE ?
```
Output MsSQL:
```text
[u].[number] LIKE @p1
```
Output MySQL:
```text
`u`.`string` LIKE ?
```
Output PostgreSQL:
```text
"u"."string" LIKE $1
```
Output SQLite:
```text
"u"."string" LIKE ?
```

### NotBetween
Проверяет, находится ли левое выражение вне диапазона, заданного `valueStart` и `valueEnd`.
```go
comparison := uast.NotBetween(uast.Field[int]("u", "number"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` NOT BETWEEN ? AND ?
```
Output MsSQL:
```text
[u].[number] NOT BETWEEN @p1 AND @p2
```
Output MySQL:
```text
`u`.`number` NOT BETWEEN ? AND ?
```
Output PostgreSQL:
```text
"u"."number" NOT BETWEEN $1 AND $2
```
Output SQLite:
```text
"u"."number" NOT BETWEEN ? AND ?
```

### NotEqual
Сравнивает два выражения на неравенство (`!=` or `<>`).
```go
comparison := uast.NotEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
Output MariaDB:
```text
`u`.`number` != ?
```
Output MsSQL:
```text
[u].[number] != @p1
```
Output MySQL:
```text
`u`.`number` != ?
```
Output PostgreSQL:
```text
"u"."number" != $1
```
Output SQLite:
```text
"u"."number" != ?
```

### NotExists
Проверяет, что подзапрос не возвращает строк. Возвращает `true` если результат подзапроса пуст.
```go
comparison := uast.NotExists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("users").As("u"))))
```
Output MariaDB:
```text
NOT EXISTS (SELECT 1 FROM `users` AS `u`)
```
Output MsSQL:
```text
NOT EXISTS (SELECT 1 FROM [users] AS [u])
```
Output MySQL:
```text
NOT EXISTS (SELECT 1 FROM `users` AS `u`)
```
Output PostgreSQL:
```text
NOT EXISTS (SELECT 1 FROM "users" AS "u")
```
Output SQLite:
```text
NOT EXISTS (SELECT 1 FROM "users" AS "u")
```

### NotILike
Выполняет отрицательное регистронезависимое сравнение с шаблоном.
```go
comparison := uast.NotILike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
LOWER(`u`.`string`) NOT LIKE LOWER(?)
```
Output MsSQL:
```text
LOWER([u].[string]) NOT LIKE LOWER(@p1)
```
Output MySQL:
```text
LOWER(`u`.`string`) NOT LIKE LOWER(?)
```
Output PostgreSQL:
```text
"u"."string" NOT ILIKE $1
```
Output SQLite:
```text
LOWER("u"."string") NOT LIKE LOWER(?)
```

### NotIn
Проверяет, что левое выражение не соответствует ни одному значению, содержащемуся в правом выражении.
```go
comparison := uast.NotIn(uast.Field[string]("u", "string"), uast.Array("active", "pending"))
```
Output MariaDB:
```text
`u`.`string` NOT IN (?, ?)
```
Output MsSQL:
```text
[u].[string] NOT IN (@p1, @p2)
```
Output MySQL:
```text
`u`.`string` NOT IN (?, ?)
```
Output PostgreSQL:
```text
"u"."string" NOT IN ($1, $2)
```
Output SQLite:
```text
"u"."string" NOT IN (?, ?)
```

### NotLike
Выполняет отрицательное регистрозависимое сравнение с шаблоном.
```go
comparison := uast.NotLike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
Output MariaDB:
```text
`u`.`string` NOT LIKE ?
```
Output MsSQL:
```text
[u].[string] NOT LIKE @p1
```
Output MySQL:
```text
`u`.`string` NOT LIKE ?
```
Output PostgreSQL:
```text
"u"."string" NOT LIKE $1
```
Output SQLite:
```text
"u"."string" NOT LIKE ?
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

## exprField
### Field
Создаёт ссылку на колонку таблицы, опционально квалифицированную псевдонимом таблицы. Это основной способ ссылаться на колонки базы данных в выражениях.
```go
field := uast.Field[string]("u", "string")
```
Output MariaDB:
```text
`u`.`string`
```
Output MsSQL:
```text
[u].[string]
```
Output MySQL:
```text
`u`.`string`
```
Output PostgreSQL:
```text
"u"."string"
```
Output SQLite:
```text
"u"."string"
```

## exprFunction
### Aggregate
#### Avg
Возвращает среднее арифметическое всех не-NULL значений в выражении. Если `distinct` равен `true`, среднее вычисляется только по уникальным значениям.
```go
function := uast.Avg(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Avg(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
AVG(`u`.`number`)
AVG(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
AVG([u].[number])
AVG(DISTINCT [u].[number])
```
Output MySQL:
```text
AVG(`u`.`number`)
AVG(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
AVG("u"."number")
AVG(DISTINCT "u"."number")
```
Output SQLite:
```text
AVG("u"."number")
AVG(DISTINCT "u"."number")
```

#### BitAnd
Возвращает побитовое И всех битов в выражении. Имеет смысл только для целочисленных типов.
```go
function := uast.BitAnd(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitAnd(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
BIT_AND(`u`.`number`)
BIT_AND(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
BIT_AND([u].[number])
BIT_AND(DISTINCT [u].[number])
```
Output MySQL:
```text
BIT_AND(`u`.`number`)
BIT_AND(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
BIT_AND("u"."number")
BIT_AND(DISTINCT "u"."number")
```
Output SQLite:
```text
BIT_AND("u"."number")
BIT_AND(DISTINCT "u"."number")
```

#### BitOr
Возвращает побитовое ИЛИ всех битов в выражении.
```go
function := uast.BitOr(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitOr(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
BIT_OR(`u`.`number`)
BIT_OR(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
BIT_OR([u].[number])
BIT_OR(DISTINCT [u].[number])
```
Output MySQL:
```text
BIT_OR(`u`.`number`)
BIT_OR(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
BIT_OR("u"."number")
BIT_OR(DISTINCT "u"."number")
```
Output SQLite:
```text
BIT_OR("u"."number")
BIT_OR(DISTINCT "u"."number")
```

#### BitXor
Возвращает побитовое исключающее ИЛИ всех битов в выражении.
```go
function := uast.BitXor(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitXor(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
BIT_XOR(`u`.`number`)
BIT_XOR(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
BIT_XOR([u].[number])
BIT_XOR(DISTINCT [u].[number])
```
Output MySQL:
```text
BIT_XOR(`u`.`number`)
BIT_XOR(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
BIT_XOR("u"."number")
BIT_XOR(DISTINCT "u"."number")
```
Output SQLite:
```text
BIT_XOR("u"."number")
BIT_XOR(DISTINCT "u"."number")
```

#### Count
Возвращает количество строк, соответствующих запросу, или количество не-NULL значений, если указано выражение. Когда `distinct` равен `true`, подсчитываются только уникальные значения.
```go
function := uast.Count(uast.Field[string]("u", "string"), false)
functionWithDistinct := uast.Count(uast.Field[string]("u", "string"), true)
```
Output MariaDB:
```text
COUNT(`u`.`string`)
COUNT(DISTINCT `u`.`string`)
```
Output MsSQL:
```text
COUNT([u].[string])
COUNT(DISTINCT [u].[string])
```
Output MySQL:
```text
COUNT(`u`.`string`)
COUNT(DISTINCT `u`.`string`)
```
Output PostgreSQL:
```text
COUNT("u"."string")
COUNT(DISTINCT "u"."string")
```
Output SQLite:
```text
COUNT("u"."string")
COUNT(DISTINCT "u"."string")
```

#### GroupConcat
Объединяет значения из группы в одну строку, разделённую стандартным разделителем (обычно запятая). Флаг `distinct` удаляет дубликаты перед объединением.
```go
function := uast.GroupConcat(uast.Field[string]("u", "string"), false)
functionWithDistinct := uast.GroupConcat(uast.Field[string]("u", "string"), true)
```
Output MariaDB:
```text
GROUP_CONCAT(`u`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `u`.`string` SEPARATOR ',')
```
Output MsSQL:
```text
GROUP_CONCAT([u].[string], ',')
GROUP_CONCAT(DISTINCT [u].[string], ',')
```
Output MySQL:
```text
GROUP_CONCAT(`u`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `u`.`string` SEPARATOR ',')
```
Output PostgreSQL:
```text
STRING_AGG("u"."string", ',')
STRING_AGG(DISTINCT "u"."string", ',')
```
Output SQLite:
```text
GROUP_CONCAT("u"."string" SEPARATOR ',')
GROUP_CONCAT(DISTINCT "u"."string" SEPARATOR ',')
```

#### Max
Возвращает максимальное значение выражения по всем строкам в группе.
```go
function := uast.Max(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Max(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
MAX(`u`.`number`)
MAX(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
MAX([u].[number])
MAX(DISTINCT [u].[number])
```
Output MySQL:
```text
MAX(`u`.`number`)
MAX(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
MAX("u"."number")
MAX(DISTINCT "u"."number")
```
Output SQLite:
```text
MAX("u"."number")
MAX(DISTINCT "u"."number")
```

#### Min
Возвращает минимальное значение выражения по всем строкам в группе.
```go
function := uast.Min(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Min(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
MIN(`u`.`number`)
MIN(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
MIN([u].[number])
MIN(DISTINCT [u].[number])
```
Output MySQL:
```text
MIN(`u`.`number`)
MIN(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
MIN("u"."number")
MIN(DISTINCT "u"."number")
```
Output SQLite:
```text
MIN("u"."number")
MIN(DISTINCT "u"."number")
```

#### StdDev
Возвращает популяционное стандартное отклонение выражения.
```go
function := uast.StdDev(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.StdDev(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
STDDEV(`u`.`number`)
STDDEV(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
STDEV([u].[number])
STDEV(DISTINCT [u].[number])
```
Output MySQL:
```text
STDDEV(`u`.`number`)
STDDEV(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
STDDEV_SAMP("u"."number")
STDDEV_SAMP(DISTINCT "u"."number")
```
Output SQLite:
```text
STDEV("u"."number")
STDEV(DISTINCT "u"."number")
```

#### Sum
Возвращает сумму всех значений в выражении. Если `distinct` равен `true`, суммируются только уникальные значения.
```go
function := uast.Sum(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Sum(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
SUM(`u`.`number`)
SUM(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
SUM([u].[number])
SUM(DISTINCT [u].[number])
```
Output MySQL:
```text
SUM(`u`.`number`)
SUM(DISTINCT `u`.`number`)
```
Output PostgreSQL:
```text
SUM("u"."number")
SUM(DISTINCT "u"."number")
```
Output SQLite:
```text
SUM("u"."number")
SUM(DISTINCT "u"."number")
```

#### Variance
Возвращает популяционную дисперсию выражения.
```go
function := uast.Variance(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Variance(uast.Field[int]("u", "number"), true)
```
Output MariaDB:
```text
VARIANCE(`u`.`number`)
VARIANCE(DISTINCT `u`.`number`)
```
Output MsSQL:
```text
VAR([u].[number])
VAR(DISTINCT [u].[number])
```
Output MySQL:
```text
VARIANCE(`u`.`number`)
VARIANCE(DISTINCT "u"."number")
```
Output PostgreSQL:
```text
VAR_SAMP("u"."number")
VAR_SAMP(DISTINCT "u"."number")
```
Output SQLite:
```text
VARIANCE("u"."number")
VARIANCE(DISTINCT "u"."number")
```

### Analytical
#### FirstValue
Возвращает значение выражения из первой строки оконного фрейма. Требует оператор `OVER` с оконной спецификацией.
```go
function := uast.FirstValue(uast.Field[string]("u", "string")).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
FIRST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
FIRST_VALUE([u].[string]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
FIRST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
FIRST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
FIRST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### Lag
Возвращает значение выражения из строки, смещённой на `offset` строк назад от текущей строки в рамках раздела.
```go
function := uast.Lag(uast.Field[int]("u", "number"), 2).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Asc(uast.Field[time.Time]("u", "date"))),
)
```
Output MariaDB:
```text
LAG(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
Output MsSQL:
```text
LAG([u].[number], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[date] ASC)
```
Output MySQL:
```text
LAG(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
Output PostgreSQL:
```text
LAG("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```
Output SQLite:
```text
LAG("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```

#### LastValue
Возвращает значение выражения из последней строки оконного фрейма.
```go
function := uast.LastValue(uast.Field[string]("u", "string")).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Asc(uast.Field[int]("u", "number"))),
    uast.RowsBetween("CURRENT ROW", "UNBOUNDED FOLLOWING"),
)
```
Output MariaDB:
```text
LAST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output MsSQL:
```text
LAST_VALUE([u].[string]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output MySQL:
```text
LAST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output PostgreSQL:
```text
LAST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output SQLite:
```text
LAST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```

#### Lead
Возвращает значение выражения из строки, смещённой на `offset` строк вперёд от текущей строки в рамках раздела.
```go
function := uast.Lead(uast.Field[int]("u", "number"), 2).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Asc(uast.Field[time.Time]("u", "date"))),
)
```
Output MariaDB:
```text
LEAD(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
Output MsSQL:
```text
LEAD([u].[number], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[date] ASC)
```
Output MySQL:
```text
LEAD(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
Output PostgreSQL:
```text
LEAD("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```
Output SQLite:
```text
LEAD("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```

#### NthValue
Возвращает значение выражения из `n-й` строки оконного фрейма.
```go
function := uast.NthValue(uast.Field[string]("u", "string"), 2).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
    uast.RowsBetween("UNBOUNDED PRECEDING", "CURRENT ROW"),
)
```
Output MariaDB:
```text
NTH_VALUE(`u`.`string`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output MsSQL:
```text
NTH_VALUE([u].[string], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output MySQL:
```text
NTH_VALUE(`u`.`string`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output PostgreSQL:
```text
NTH_VALUE("u"."string", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output SQLite:
```text
NTH_VALUE("u"."string", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```

### Condition
#### Case
Вычисляет список пар `WHEN`-`THEN` и возвращает выражение `THEN` для первого истинного WHEN. Если ни одно условие не истинно, возвращает выражение  `ELSE` если оно указано, иначе `NULL`.
```go
pairs := uast.CaseIf(
    uast.CasePair(
        uast.Less(uast.Field[int]("u", "number"), uast.Value(2)),
        uast.Value("old"),
    ),
)
elseExpr := uast.CaseElse(uast.Value("new"))
function := uast.Case(pairs, elseExpr)
```
Output MariaDB:
```text
CASE WHEN `u`.`number` < ? THEN ? ELSE ? END
```
Output MsSQL:
```text
CASE WHEN [u].[number] < @p1 THEN @p2 ELSE @p3 END
```
Output MySQL:
```text
CASE WHEN `u`.`number` < ? THEN ? ELSE ? END
```
Output PostgreSQL:
```text
CASE WHEN "u"."number" < $1 THEN $2 ELSE $3 END
```
Output SQLite:
```text
CASE WHEN "u"."number" < ? THEN ? ELSE ? END
```

#### Coalesce
Возвращает первое не-NULL выражение из предоставленного списка. Полезно для указания запасных значений.
```go
function := uast.Coalesce(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
Output MariaDB:
```text
COALESCE(`u`.`createat`, `u`.`updateat`)
```
Output MsSQL:
```text
COALESCE([u].[createat], [u].[updateat])
```
Output MySQL:
```text
COALESCE(`u`.`createat`, `u`.`updateat`)
```
Output PostgreSQL:
```text
COALESCE("u"."createat", "u"."updateat")
```
Output SQLite:
```text
COALESCE("u"."createat", "u"."updateat")
```

#### Greatest
Возвращает наибольшее значение из предоставленного списка выражений.
```go
function := uast.Greatest(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
Output MariaDB:
```text
GREATEST(`u`.`createat`, `u`.`updateat`)
```
Output MsSQL:
```text
GREATEST([u].[createat], [u].[updateat])
```
Output MySQL:
```text
GREATEST(`u`.`createat`, `u`.`updateat`)
```
Output PostgreSQL:
```text
GREATEST("u"."createat", "u"."updateat")
```
Output SQLite:
```text
GREATEST("u"."createat", "u"."updateat")
```

#### Least
Возвращает наименьшее значение из предоставленного списка выражений.
```go
function := uast.Least(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
Output MariaDB:
```text
LEAST(`u`.`createat`, `u`.`updateat`)
```
Output MsSQL:
```text
LEAST([u].[createat], [u].[updateat])
```
Output MySQL:
```text
LEAST(`u`.`createat`, `u`.`updateat`)
```
Output PostgreSQL:
```text
LEAST("u"."createat", "u"."updateat")
```
Output SQLite:
```text
LEAST("u"."createat", "u"."updateat")
```

#### NullIf
Возвращает `NULL` если два выражения равны; иначе возвращает первое выражение.
```go
function := uast.NullIf(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
Output MariaDB:
```text
NULLIF(`u`.`createat`, `u`.`updateat`)
```
Output MsSQL:
```text
NULLIF([u].[createat], [u].[updateat])
```
Output MySQL:
```text
NULLIF(`u`.`createat`, `u`.`updateat`)
```
Output PostgreSQL:
```text
NULLIF("u"."createat", "u"."updateat")
```
Output SQLite:
```text
NULLIF("u"."createat", "u"."updateat")
```

### Convert
#### Cast
Преобразует выражение к указанному типу данных.
```go
function := uast.Cast(uast.Field[int]("u", "number"), uast.TypeString)
```
Output MariaDB:
```text
CAST(`u`.`number` AS CHAR)
```
Output MsSQL:
```text
CAST([u].[number] AS NVARCHAR)
```
Output MySQL:
```text
CAST(`u`.`number` AS CHAR)
```
Output PostgreSQL:
```text
CAST("u"."number" AS VARCHAR)
```
Output SQLite:
```text
CAST("u"."number" AS TEXT)
```

#### CharLength
Возвращает количество символов в строковом выражении.
```go
function := uast.CharLength(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
CHAR_LENGTH(`u`.`string`)
```
Output MsSQL:
```text
CHAR_LENGTH([u].[string])
```
Output MySQL:
```text
CHAR_LENGTH(`u`.`string`)
```
Output PostgreSQL:
```text
CHAR_LENGTH("u"."string")
```
Output SQLite:
```text
CHAR_LENGTH("u"."string")
```

#### DateFormat
Форматирует выражение даты/времени в соответствии с указанной маской формата.
```go
function := uast.DateFormat(uast.Field[time.Time]("u", "createat"), uast.Value("%Y-%m-%d"))
```
Output MariaDB:
```text
DATE_FORMAT(`u`.`createat`, '%Y-%m-%d')
```
Output MsSQL:
```text
FORMAT([u].[createat], '%Y-%m-%d')
```
Output MySQL:
```text
DATE_FORMAT(`u`.`createat`, '%Y-%m-%d')
```
Output PostgreSQL:
```text
TO_CHAR("u"."createat", '%Y-%m-%d')
```
Output SQLite:
```text
STRFTIME("u"."createat", '%Y-%m-%d')
```

#### Degrees
Преобразует угол из радиан в градусы.
```go
function := uast.Degrees(uast.Field[int]("u", "number"))
```
Output MariaDB:
```text
DEGREES(`u`.`number`)
```
Output MsSQL:
```text
DEGREES([u].[number])
```
Output MySQL:
```text
DEGREES(`u`.`number`)
```
Output PostgreSQL:
```text
DEGREES("u"."number")
```
Output SQLite:
```text
DEGREES("u"."number")
```

#### Length
Возвращает длину строкового выражения в байтах.
```go
function := uast.Length(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
LENGTH(`u`.`string`)
```
Output MsSQL:
```text
LEN([u].[string])
```
Output MySQL:
```text
LENGTH(`u`.`string`)
```
Output PostgreSQL:
```text
LENGTH("u"."string")
```
Output SQLite:
```text
LENGTH("u"."string")
```

#### Position
Возвращает начальную позицию первого вхождения подстроки в строку.
```go
function := uast.Position(uast.Field[string]("u", "string"), uast.Value("old"))
```
Output MariaDB:
```text
POSITION(? IN `u`.`string`)
```
Output MsSQL:
```text
CHARINDEX(@p1, [u].[string])
```
Output MySQL:
```text
POSITION(? IN `u`.`string`)
```
Output PostgreSQL:
```text
POSITION($1 IN "u"."string")
```
Output SQLite:
```text
POSITION(? IN "u"."string")
```

#### Radians
Преобразует угол из градусов в радианы.
```go
function := uast.Radians(uast.Field[int]("u", "number"))
```
Output MariaDB:
```text
RADIANS(`u`.`number`)
```
Output MsSQL:
```text
RADIANS([u].[number])
```
Output MySQL:
```text
RADIANS(`u`.`number`)
```
Output PostgreSQL:
```text
RADIANS("u"."number")
```
Output SQLite:
```text
RADIANS("u"."number")
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
DATE('now')
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
TIME('now')
```

#### DateAdd
Добавляет интервал даты/времени к выражению даты/времени и возвращает результирующую дату/время.
```go
function := uast.DateAdd(uast.Field[time.Time]("u", "createat"), uast.Value("2 DAY"))
```
Output MariaDB:
```text
DATE_ADD(`u`.`createat`, INTERVAL 2 DAY)
```
Output MsSQL:
```text
DATEADD(DAY, 2, [u].[createat])
```
Output MySQL:
```text
DATE_ADD(`u`.`createat`, INTERVAL 2 DAY)
```
Output PostgreSQL:
```text
("u"."createat" + INTERVAL '2 DAY')
```
Output SQLite:
```text
DATETIME("u"."createat", '+2 DAY')
```

#### DateDiff
Возвращает разницу в днях между двумя выражениями даты/времени (`datetimeEnd` - `datetimeStart`).
```go
function := uast.DateDiff(uast.Field[time.Time]("u", "updateat"), uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
DATEDIFF(`u`.`updateat`, `u`.`createat`)
```
Output MsSQL:
```text
DATEDIFF([u].[updateat], [u].[createat])
```
Output MySQL:
```text
DATEDIFF(`u`.`updateat`, `u`.`createat`)
```
Output PostgreSQL:
```text
DATE_PART('day', "u"."updateat" - "u"."createat")
```
Output SQLite:
```text
DATEDIFF("u"."updateat", "u"."createat")
```

#### DateSub
Вычитает интервал даты/времени из выражения даты/времени и возвращает результирующую дату/время.
```go
function := uast.DateSub(uast.Field[time.Time]("u", "createat"), uast.Value("2 DAY"))
```
Output MariaDB:
```text
DATE_SUB(`u`.`createat`, INTERVAL 2 DAY)
```
Output MsSQL:
```text
DATEADD(DAY, -2, [u].[createat])
```
Output MySQL:
```text
DATE_SUB(`u`.`createat`, INTERVAL 2 DAY)
```
Output PostgreSQL:
```text
("u"."createat" - INTERVAL '2 DAY')
```
Output SQLite:
```text
DATETIME("u"."createat", '-2 DAY')
```

#### Day
Извлекает день месяца (1–31) из выражения даты/времени.
```go
function := uast.Day(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
DAY(`u`.`createat`)
```
Output MsSQL:
```text
DAY([u].[createat])
```
Output MySQL:
```text
DAY(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(DAY FROM "u"."createat")
```
Output SQLite:
```text
DAY("u"."createat")
```

#### DayName
Возвращает название дня недели (например, 'Понедельник', 'Вторник') для заданного выражения даты/времени.
```go
function := uast.DayName(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
DAYNAME(`u`.`createat`)
```
Output MsSQL:
```text
DATENAME(WEEKDAY, [u].[createat])
```
Output MySQL:
```text
DAYNAME(`u`.`createat`)
```
Output PostgreSQL:
```text
TO_CHAR("u"."createat", 'Day')
```
Output SQLite:
```text
STRFTIME('%w', "u"."createat")
```

#### Hour
Извлекает час (0–23) из выражения даты/времени.
```go
function := uast.Hour(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
HOUR(`u`.`createat`)
```
Output MsSQL:
```text
DATEPART(HOUR, [u].[createat])
```
Output MySQL:
```text
HOUR(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(HOUR FROM "u"."createat")
```
Output SQLite:
```text
HOUR("u"."createat")
```

#### Minute
Извлекает минуту (0–59) из выражения даты/времени.
```go
function := uast.Minute(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
MINUTE(`u`.`createat`)
```
Output MsSQL:
```text
DATEPART(MINUTE, [u].[createat])
```
Output MySQL:
```text
MINUTE(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(MINUTE FROM "u"."createat")
```
Output SQLite:
```text
MINUTE("u"."createat")
```

#### Month
Извлекает месяц (1–12) из выражения даты/времени.
```go
function := uast.Month(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
MONTH(`u`.`createat`)
```
Output MsSQL:
```text
MONTH([u].[createat])
```
Output MySQL:
```text
MONTH(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(MONTH FROM "u"."createat")
```
Output SQLite:
```text
MONTH("u"."createat")
```

#### MonthName
Возвращает название месяца (например, 'Январь', 'Февраль') для заданного выражения даты/времени.
```go
function := uast.MonthName(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
MONTHNAME(`u`.`createat`)
```
Output MsSQL:
```text
DATENAME(MONTH, [u].[createat])
```
Output MySQL:
```text
MONTHNAME(`u`.`createat`)
```
Output PostgreSQL:
```text
TO_CHAR("u"."createat", 'Month')
```
Output SQLite:
```text
STRFTIME('%m', "u"."createat")
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
DATETIME('now')
```

#### Quarter
Извлекает квартал (1–4) из выражения даты/времени.
```go
function := uast.Quarter(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
QUARTER(`u`.`createat`)
```
Output MsSQL:
```text
DATEPART(QUARTER, [u].[createat])
```
Output MySQL:
```text
QUARTER(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(QUARTER FROM "u"."createat")
```
Output SQLite:
```text
QUARTER("u"."createat")
```

#### Second
Извлекает секунду (0–59) из выражения даты/времени.
```go
function := uast.Second(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
SECOND(`u`.`createat`)
```
Output MsSQL:
```text
DATEPART(SECOND, [u].[createat])
```
Output MySQL:
```text
SECOND(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(SECOND FROM "u"."createat")
```
Output SQLite:
```text
SECOND("u"."createat")
```

#### TimeAdd
Добавляет интервал времени к выражению времени/даты/времени и возвращает результирующее время.
```go
function := uast.TimeAdd(uast.Field[time.Time]("u", "createat"), uast.Value("2 HOUR"))
```
Output MariaDB:
```text
TIME_ADD(`u`.`createat`, '2 HOUR')
```
Output MsSQL:
```text
DATEADD(HOUR, 2, [u].[createat])
```
Output MySQL:
```text
TIME_ADD(`u`.`createat`, '2 HOUR')
```
Output PostgreSQL:
```text
("u"."createat" + INTERVAL '2 HOUR')
```
Output SQLite:
```text
TIME("u"."createat", '+2 HOUR')
```

#### TimeDiff
Возвращает разницу между двумя выражениями времени/даты/времени (`timeEnd` - `timeStart`).
```go
function := uast.TimeDiff(uast.Field[time.Time]("u", "updateat"), uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
TIMEDIFF(`u`.`updateat`, `u`.`createat`)
```
Output MsSQL:
```text
TIMEDIFF([u].[updateat], [u].[createat])
```
Output MySQL:
```text
TIMEDIFF(`u`.`updateat`, `u`.`createat`)
```
Output PostgreSQL:
```text
DATE_PART('time', "u"."updateat" - "u"."createat")
```
Output SQLite:
```text
TIMEDIFF("u"."updateat", "u"."createat")
```

#### TimeSub
Вычитает интервал времени из выражения времени/даты/времени и возвращает результирующее время.
```go
function := uast.TimeSub(uast.Field[time.Time]("u", "createat"), uast.Value("2 HOUR"))
```
Output MariaDB:
```text
TIME_SUB(`u`.`createat`, '2 HOUR')
```
Output MsSQL:
```text
DATEADD(HOUR, -2, [u].[createat])
```
Output MySQL:
```text
TIME_SUB(`u`.`createat`, '2 HOUR')
```
Output PostgreSQL:
```text
("u"."createat" - INTERVAL '2 HOUR')
```
Output SQLite:
```text
TIME("u"."createat", '-2 HOUR')
```

#### Week
Извлекает номер недели (1–53) из выражения даты/времени.
```go
function := uast.Week(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
WEEK(`u`.`createat`)
```
Output MsSQL:
```text
DATEPART(WEEK, [u].[createat])
```
Output MySQL:
```text
WEEK(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(WEEK FROM "u"."createat")
```
Output SQLite:
```text
WEEK("u"."createat")
```

#### Year
Извлекает год из выражения даты/времени.
```go
function := uast.Year(uast.Field[time.Time]("u", "createat"))
```
Output MariaDB:
```text
YEAR(`u`.`createat`)
```
Output MsSQL:
```text
YEAR([u].[createat])
```
Output MySQL:
```text
YEAR(`u`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(YEAR FROM "u"."createat")
```
Output SQLite:
```text
YEAR("u"."createat")
```

### Json
#### JsonArray
Создаёт JSON-массив из заданного выражения и опциональных дополнительных значений.
```go
function := uast.JsonArray(
    uast.Field[string]("u", "json"), 
    uast.Value("val1"), 
    uast.Value("val2"),
)
```
Output MariaDB:
```text
JSON_ARRAY(`u`.`json`, ?, ?)
```
Output MsSQL:
```text
JSON_ARRAY([u].[json], @p1, @p2)
```
Output MySQL:
```text
JSON_ARRAY(`u`.`json`, ?, ?)
```
Output PostgreSQL:
```text
JSON_ARRAY("u"."json", $1, $2)
```
Output SQLite:
```text
JSON_ARRAY("u"."json", ?, ?)
```

#### JsonArrayAgg
Агрегирует значения из группы в JSON-массив.
```go
function := uast.JsonArrayAgg(
    uast.Field[string]("u", "json"),
)
```
Output MariaDB:
```text
JSON_ARRAYAGG(`u`.`json`)
```
Output MsSQL:
```text
JSON_ARRAYAGG([u].[json])
```
Output MySQL:
```text
JSON_ARRAYAGG(`u`.`json`)
```
Output PostgreSQL:
```text
JSON_AGG("u"."json")
```
Output SQLite:
```text
JSON_GROUP_ARRAY("u"."json")
```

#### JsonContains
Проверяет, содержит ли JSON-документ указанное значение.
```go
function := uast.JsonContains(
    uast.Field[string]("u", "json"),
    uast.Value(`{"key":"val"}`),
)
```
Output MariaDB:
```text
JSON_CONTAINS(`u`.`json`, '{"key":"val"}')
```
Output MsSQL:
```text
// Not supported
```
Output MySQL:
```text
JSON_CONTAINS(`u`.`json`, '{"key":"val"}')
```
Output PostgreSQL:
```text
("u"."json" @> '{"key":"val"}')
```
Output SQLite:
```text
JSON_CONTAINS("u"."json", '{"key":"val"}')
```

#### JsonExtract
Извлекает значение из JSON-документа по указанному пути. Параметр `json` строится с помощью `JsonPath` и опциональных `JsonKey`/`JsonIndex`.
```go
function := JsonExtract(
    uast.Field[string]("u", "json"), 
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
(`u`.`json` ->> '$.parent[0].child')
```
Output MsSQL:
```text
JSON_VALUE([u].[json], '$.parent[0].child')
```
Output MySQL:
```text
(`u`.`json` ->> '$.parent[0].child')
```
Output PostgreSQL:
```text
("u"."json" #>> '{parent,0,child}')
```
Output SQLite:
```text
("u"."json" ->> '$.parent[0].child')
```

#### JsonObject
Создаёт JSON-объект из пар ключ-значение.
```go
function := uast.JsonObject(
    uast.JsonPair(
        uast.JsonKey("key"), 
        uast.Count(uast.Field[string]("u", "json"), false),
    ),
)
```
Output MariaDB:
```text
JSON_OBJECT('key', COUNT(`u`.`json`))
```
Output MsSQL:
```text
JSON_OBJECT('key', COUNT([u].[json]))
```
Output MySQL:
```text
JSON_OBJECT('key', COUNT(`u`.`json`))
```
Output PostgreSQL:
```text
JSON_BUILD_OBJECT('key', COUNT("u"."json"))
```
Output SQLite:
```text
JSON_OBJECT('key', COUNT("u"."json"))
```

#### JsonObjectAgg
Агрегирует пары ключ-значение из группы в один JSON-объект.
```go
function := uast.JsonObjectAgg(
    uast.Field[string]("u", "json"),
    uast.Field[int]("u", "number"),
)
```
Output MariaDB:
```text
JSON_OBJECTAGG(`u`.`json`, `u`.`number`)
```
Output MsSQL:
```text
JSON_OBJECTAGG([u].[json], [u].[number])
```
Output MySQL:
```text
JSON_OBJECTAGG(`u`.`json`, `u`.`number`)
```
Output PostgreSQL:
```text
JSON_OBJECT_AGG("u"."json", "u"."number")
```
Output SQLite:
```text
JSON_GROUP_OBJECT("u"."json", "u"."number")
```

#### JsonRemove
Удаляет значение из JSON-документа по указанному пути(ям).
```go
function := uast.JsonRemove(
    uast.Field[string]("u", "json"),
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
JSON_REMOVE(`u`.`json`, '$.key1', '$.key2')
```
Output MsSQL:
```text
JSON_MODIFY(JSON_MODIFY([u].[json], '$.key1', NULL), '$.key2', NULL)
```
Output MySQL:
```text
JSON_REMOVE(`u`.`json`, '$.key1', '$.key2')
```
Output PostgreSQL:
```text
("u"."json" - '{key1}' - '{key2}')
```
Output SQLite:
```text
JSON_REMOVE("u"."json", '$.key1', '$.key2')
```

#### JsonSet
Устанавливает значение в JSON-документе по указанному пути(ям). Создаёт путь, если он не существует.
```go
function := uast.JsonSet(
    uast.Field[string]("u", "json"),
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
JSON_SET(`u`.`json`, '$.key1', ?, '$.key2', ?)
```
Output MsSQL:
```text
JSON_MODIFY(JSON_MODIFY([u].[json], '$.key1', @p1), '$.key2', @p2)
```
Output MySQL:
```text
JSON_SET(`u`.`json`, '$.key1', ?, '$.key2', ?)
```
Output PostgreSQL:
```text
jsonb_set(jsonb_set("u"."json", '{key1}', $1), '{key2}', $2)
```
Output SQLite:
```text
JSON_SET("u"."json", '$.key1', ?, '$.key2', ?)
```

#### JsonType
Возвращает тип JSON-значения (например, 'OBJECT', 'ARRAY', 'STRING', 'INTEGER', 'NULL').
```go
function := uast.JsonType(uast.Field[string]("u", "json"))
```
Output MariaDB:
```text
JSON_TYPE(`u`.`json`)
```
Output MsSQL:
```text
// Not supported
```
Output MySQL:
```text
JSON_TYPE(`u`.`json`)
```
Output PostgreSQL:
```text
jsonb_typeof("u"."json")
```
Output SQLite:
```text
JSON_TYPE("u"."json")
```

### Math
#### Abs
Возвращает абсолютное (неотрицательное) значение числового выражения.
```go
function := uast.Abs(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
ABS(`u`.`x`)
```
Output MsSQL:
```text
ABS([u].[x])
```
Output MySQL:
```text
ABS(`u`.`x`)
```
Output PostgreSQL:
```text
ABS("u"."x")
```
Output SQLite:
```text
ABS("u"."x")
```

#### ACos
Возвращает арккосинус (обратный косинус) выражения в радианах.
```go
function := uast.ACos(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
ACOS(`u`.`x`)
```
Output MsSQL:
```text
ACOS([u].[x])
```
Output MySQL:
```text
ACOS(`u`.`x`)
```
Output PostgreSQL:
```text
ACOS("u"."x")
```
Output SQLite:
```text
ACOS("u"."x")
```

#### ASin
Возвращает арксинус (обратный синус) выражения в радианах.
```go
function := uast.ASin(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
ASIN(`u`.`x`)
```
Output MsSQL:
```text
ASIN([u].[x])
```
Output MySQL:
```text
ASIN(`u`.`x`)
```
Output PostgreSQL:
```text
ASIN("u"."x")
```
Output SQLite:
```text
ASIN("u"."x")
```

#### ATan
Возвращает арктангенс (обратный тангенс) выражения в радианах.
```go
function := uast.ATan(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
ATAN(`u`.`x`)
```
Output MsSQL:
```text
ATAN([u].[x])
```
Output MySQL:
```text
ATAN(`u`.`x`)
```
Output PostgreSQL:
```text
ATAN("u"."x")
```
Output SQLite:
```text
ATAN("u"."x")
```

#### ATan2
Возвращает арктангенс частного двух аргументов (`y`/`x`), используя их знаки для определения квадранта.
```go
function := uast.ATan2(uast.Field[int]("u", "y"), uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
ATAN2(`u`.`y`, `u`.`x`)
```
Output MsSQL:
```text
ATAN2([u].[y], [u].[x])
```
Output MySQL:
```text
ATAN2(`u`.`y`, `u`.`x`)
```
Output PostgreSQL:
```text
ATAN2("u"."y", "u"."x")
```
Output SQLite:
```text
ATAN2("u"."y", "u"."x")
```

#### Cbrt
Возвращает кубический корень числового выражения.
```go
function := uast.Cbrt(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
CBRT(`u`.`x`)
```
Output MsSQL:
```text
CBRT([u].[x])
```
Output MySQL:
```text
CBRT(`u`.`x`)
```
Output PostgreSQL:
```text
CBRT("u"."x")
```
Output SQLite:
```text
CBRT("u"."x")
```

#### Ceil
Возвращает наименьшее целое значение, не меньшее аргумента (округление вверх).
```go
function := uast.Ceil(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
CEILING(`u`.`x`)
```
Output MsSQL:
```text
CEILING([u].[x])
```
Output MySQL:
```text
CEILING(`u`.`x`)
```
Output PostgreSQL:
```text
CEIL("u"."x")
```
Output SQLite:
```text
CEIL("u"."x")
```

#### Cos
Возвращает косинус выражения в радианах.
```go
function := uast.Cos(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
COS(`u`.`x`)
```
Output MsSQL:
```text
COS([u].[x])
```
Output MySQL:
```text
COS(`u`.`x`)
```
Output PostgreSQL:
```text
COS("u"."x")
```
Output SQLite:
```text
COS("u"."x")
```

#### Exp
Возвращает число Эйлера `e` (~2.71828) возведённое в степень выражения.
```go
function := uast.Exp(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
EXP(`u`.`x`)
```
Output MsSQL:
```text
EXP([u].[x])
```
Output MySQL:
```text
EXP(`u`.`x`)
```
Output PostgreSQL:
```text
EXP("u"."x")
```
Output SQLite:
```text
EXP("u"."x")
```

#### Floor
Возвращает наибольшее целое значение, не большее аргумента (округление вниз).
```go
function := uast.Floor(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
FLOOR(`u`.`x`)
```
Output MsSQL:
```text
FLOOR([u].[x])
```
Output MySQL:
```text
FLOOR(`u`.`x`)
```
Output PostgreSQL:
```text
FLOOR("u"."x")
```
Output SQLite:
```text
FLOOR("u"."x")
```

#### Ln
Возвращает натуральный логарифм (по основанию `e`) выражения.
```go
function := uast.Ln(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
LN(`u`.`x`)
```
Output MsSQL:
```text
LN([u].[x])
```
Output MySQL:
```text
LN(`u`.`x`)
```
Output PostgreSQL:
```text
LN("u"."x")
```
Output SQLite:
```text
LN("u"."x")
```

#### Log
Возвращает логарифм выражения по указанному основанию.
```go
function := uast.Log(uast.Field[int]("u", "x"), uast.Value(2))
```
Output MariaDB:
```text
LOG(`u`.`x`, ?)
```
Output MsSQL:
```text
LOG([u].[x], @p1)
```
Output MySQL:
```text
LOG(`u`.`x`, ?)
```
Output PostgreSQL:
```text
LOG("u"."x", $1)
```
Output SQLite:
```text
LOG("u"."x", ?)
```

#### Mod
Возвращает остаток (модуль) от деления первого выражения на второе.
```go
function := uast.Mod(uast.Field[int]("u", "x"), uast.Value(2))
```
Output MariaDB:
```text
MOD(`u`.`x`, ?)
```
Output MsSQL:
```text
MOD([u].[x], @p1)
```
Output MySQL:
```text
MOD(`u`.`x`, ?)
```
Output PostgreSQL:
```text
MOD("u"."x", $1)
```
Output SQLite:
```text
MOD("u"."x", ?)
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
function := uast.Power(uast.Field[int]("u", "x"), uast.Value(2))
```
Output MariaDB:
```text
POWER(`u`.`x`, ?)
```
Output MsSQL:
```text
POWER([u].[x], @p1)
```
Output MySQL:
```text
POWER(`u`.`x`, ?)
```
Output PostgreSQL:
```text
POWER("u"."x", $1)
```
Output SQLite:
```text
POWER("u"."x", ?)
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
function := uast.Round(uast.Field[int]("u", "x"), uast.Value(2))
```
Output MariaDB:
```text
ROUND(`u`.`x`, ?)
```
Output MsSQL:
```text
ROUND([u].[x], @p1)
```
Output MySQL:
```text
ROUND(`u`.`x`, ?)
```
Output PostgreSQL:
```text
ROUND("u"."x", $1)
```
Output SQLite:
```text
ROUND("u"."x", ?)
```

#### Sin
Возвращает синус выражения в радианах.
```go
function := uast.Sin(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
SIN(`u`.`x`)
```
Output MsSQL:
```text
SIN([u].[x])
```
Output MySQL:
```text
SIN(`u`.`x`)
```
Output PostgreSQL:
```text
SIN("u"."x")
```
Output SQLite:
```text
SIN("u"."x")
```

#### Sqrt
Возвращает квадратный корень выражения.
```go
function := uast.Sqrt(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
SQRT(`u`.`x`)
```
Output MsSQL:
```text
SQRT([u].[x])
```
Output MySQL:
```text
SQRT(`u`.`x`)
```
Output PostgreSQL:
```text
SQRT("u"."x")
```
Output SQLite:
```text
SQRT("u"."x")
```

#### Tan
Возвращает тангенс выражения в радианах.
```go
function := uast.Tan(uast.Field[int]("u", "x"))
```
Output MariaDB:
```text
TAN(`u`.`x`)
```
Output MsSQL:
```text
TAN([u].[x])
```
Output MySQL:
```text
TAN(`u`.`x`)
```
Output PostgreSQL:
```text
TAN("u"."x")
```
Output SQLite:
```text
TAN("u"."x")
```

#### Trunc
Усекает числовое выражение до указанного количества знаков после запятой (без округления).
```go
function := uast.Trunc(uast.Field[int]("u", "x"), uast.Value(2))
```
Output MariaDB:
```text
TRUNCATE(`u`.`x`, ?)
```
Output MsSQL:
```text
ROUND([u].[x], @p1, 1)
```
Output MySQL:
```text
TRUNCATE(`u`.`x`, ?)
```
Output PostgreSQL:
```text
TRUNC("u"."x", $1)
```
Output SQLite:
```text
TRUNC("u"."x", ?)
```

### Ranking
#### CumeDist
Возвращает кумулятивное распределение значения в рамках раздела (отношение строк, которые идут до или равны текущей строке). Должна использоваться с оператором `OVER`.
```go
function := uast.CumeDist().Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
CUME_DIST() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
CUME_DIST() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
CUME_DIST() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
CUME_DIST() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
CUME_DIST() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### DenseRank
Возвращает ранг строки без пропусков. Строки с равными значениями получают одинаковый ранг, а следующий ранг является непосредственно следующим целым числом. Требует `OVER`.
```go
function := uast.DenseRank().Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
DENSE_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
DENSE_RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
DENSE_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
DENSE_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
DENSE_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### NTile
Делит строки в рамках раздела на `n` приблизительно равных групп и возвращает номер группы (от 1 до `n`) для каждой строки.
```go
function := uast.NTile(2).Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
NTILE(2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
NTILE(2) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
NTILE(2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
NTILE(2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
NTILE(2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### PercentRank
Возвращает процентильный ранг строки в рамках раздела (диапазон от 0 до 1). Ранг первой строки всегда равен 0. Требует `OVER`.
```go
function := uast.PercentRank().Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
PERCENT_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
PERCENT_RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
PERCENT_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
PERCENT_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
PERCENT_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### Rank
Возвращает ранг строки с пропусками. Равные значения получают одинаковый ранг, а следующее отличное значение пропускает ранги. Требует `OVER`.
```go
function := uast.Rank().Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

#### RowNumber
Присваивает уникальный последовательный номер каждой строке в рамках раздела, начиная с 1. Порядок определяет последовательность нумерации.
```go
function := uast.RowNumber().Over(
    uast.PartitionBy(uast.Field[int64]("u", "id")),
    uast.OrderBy(uast.Desc(uast.Field[int]("u", "number"))),
)
```
Output MariaDB:
```text
ROW_NUMBER() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output MsSQL:
```text
ROW_NUMBER() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
Output MySQL:
```text
ROW_NUMBER() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
Output PostgreSQL:
```text
ROW_NUMBER() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
Output SQLite:
```text
ROW_NUMBER() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

### String
#### Concat
Объединяет два или более строковых выражения в одну строку. Аргументы `NULL` рассматриваются как пустые строки в большинстве диалектов.
```go
function := uast.Concat(uast.Field[string]("u", "string"), uast.Value("old"), uast.Value("new"))
```
Output MariaDB:
```text
CONCAT(`u`.`string`, ?, ?)
```
Output MsSQL:
```text
CONCAT([u].[string], @p1, @p2)
```
Output MySQL:
```text
CONCAT(`u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
CONCAT("u"."string", $1, $2)
```
Output SQLite:
```text
CONCAT("u"."string", ?, ?)
```

#### ConcatWs
Объединяет два или более строковых выражения с указанным разделителем между ними. Пропускает аргументы `NULL`.
```go
function := uast.ConcatWs(uast.Value("_"), uast.Field[string]("u", "string"), uast.Value("old"),uast.Value("new"))
```
Output MariaDB:
```text
CONCAT_WS(?, `u`.`string`, ?, ?)
```
Output MsSQL:
```text
CONCAT_WS(@p1, [u].[string], @p2, @p3)
```
Output MySQL:
```text
CONCAT_WS(?, `u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
CONCAT_WS($1, "u"."string", $2, $3)
```
Output SQLite:
```text
CONCAT_WS(?, "u"."string", ?, ?)
```

#### LeftString
Возвращает крайние слева `count` символов из строкового выражения.
```go
function := uast.LeftString(uast.Field[string]("u", "string"), uast.Value(2))
```
Output MariaDB:
```text
LEFT(`u`.`string`, ?)
```
Output MsSQL:
```text
LEFT([u].[string], @p1)
```
Output MySQL:
```text
LEFT(`u`.`string`, ?)
```
Output PostgreSQL:
```text
LEFT("u"."string", $1)
```
Output SQLite:
```text
LEFT("u"."string", ?)
```

#### Lower
Преобразует строковое выражение в нижний регистр.
```go
function := uast.Lower(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
LOWER(`u`.`string`)
```
Output MsSQL:
```text
LOWER([u].[string])
```
Output MySQL:
```text
LOWER(`u`.`string`)
```
Output PostgreSQL:
```text
LOWER("u"."string")
```
Output SQLite:
```text
LOWER("u"."string")
```

#### LPad
Дополняет строковое выражение слева указанным разделителем до общей длины `count` символов.
```go
function := uast.LPad(uast.Field[string]("u", "string"), uast.Value(2), uast.Value(","))
```
Output MariaDB:
```text
LPAD(`u`.`string`, ?, ?)
```
Output MsSQL:
```text
LPAD([u].[string], @p1, @p2)
```
Output MySQL:
```text
LPAD(`u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
LPAD("u"."string", $1, $2)
```
Output SQLite:
```text
LPAD("u"."string", ?, ?)
```

#### LTrim
Удаляет начальные пробелы из строкового выражения.
```go
function := uast.LTrim(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
LTRIM(`u`.`string`)
```
Output MsSQL:
```text
LTRIM([u].[string])
```
Output MySQL:
```text
LTRIM(`u`.`string`)
```
Output PostgreSQL:
```text
LTRIM("u"."string")
```
Output SQLite:
```text
LTRIM("u"."string")
```

#### Repeat
Повторяет строковое выражение `count` раз.
```go
function := uast.Repeat(uast.Field[string]("u", "string"), uast.Value(2))
```
Output MariaDB:
```text
REPEAT(`u`.`string`, ?)
```
Output MsSQL:
```text
REPEAT([u].[string], @p1)
```
Output MySQL:
```text
REPEAT(`u`.`string`, ?)
```
Output PostgreSQL:
```text
REPEAT("u"."string", $1)
```
Output SQLite:
```text
REPEAT("u"."string", ?)
```

#### Replace
Заменяет все вхождения подстроки в строке на новую подстроку.
```go
function := uast.Replace(uast.Field[string]("u", "string"), uast.Value("old"), uast.Value("new"))
```
Output MariaDB:
```text
REPLACE(`u`.`string`, ?, ?)
```
Output MsSQL:
```text
REPLACE([u].[string], @p1, @p2)
```
Output MySQL:
```text
REPLACE(`u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
REPLACE("u"."string", $1, $2)
```
Output SQLite:
```text
REPLACE("u"."string", ?, ?)
```

#### Reverse
Переворачивает символы в строковом выражении.
```go
function := uast.Reverse(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
REVERSE(`u`.`string`)
```
Output MsSQL:
```text
REVERSE([u].[string])
```
Output MySQL:
```text
REVERSE(`u`.`string`)
```
Output PostgreSQL:
```text
REVERSE("u"."string")
```
Output SQLite:
```text
REVERSE("u"."string")
```

#### RightString
Возвращает крайние справа `count` символов из строкового выражения.
```go
function := uast.RightString(uast.Field[string]("u", "string"), uast.Value(2))
```
Output MariaDB:
```text
RIGHT(`u`.`string`, ?)
```
Output MsSQL:
```text
RIGHT([u].[string], @p1)
```
Output MySQL:
```text
RIGHT(`u`.`string`, ?)
```
Output PostgreSQL:
```text
RIGHT("u"."string", $1)
```
Output SQLite:
```text
RIGHT("u"."string", ?)
```

#### RPad
Дополняет строковое выражение справа указанным разделителем до общей длины `count` символов.
```go
function := uast.RPad(uast.Field[string]("u", "string"), uast.Value(2), uast.Value(","))
```
Output MariaDB:
```text
RPAD(`u`.`string`, ?, ?)
```
Output MsSQL:
```text
RPAD([u].[string], @p1, @p2)
```
Output MySQL:
```text
RPAD(`u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
RPAD("u"."string", $1, $2)
```
Output SQLite:
```text
RPAD("u"."string", ?, ?)
```

#### RTrim
Удаляет конечные пробелы из строкового выражения.
```go
function := uast.RTrim(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
RTRIM(`u`.`string`)
```
Output MsSQL:
```text
RTRIM([u].[string])
```
Output MySQL:
```text
RTRIM(`u`.`string`)
```
Output PostgreSQL:
```text
RTRIM("u"."string")
```
Output SQLite:
```text
RTRIM("u"."string")
```

#### SubString
Извлекает подстроку из строкового выражения, начиная с `startPos` (начиная с 1) длиной `lengthStr` символов.
```go
function := uast.SubString(uast.Field[string]("u", "string"), uast.Value(0), uast.Value(2))
```
Output MariaDB:
```text
SUBSTRING(`u`.`string`, ?, ?)
```
Output MsSQL:
```text
SUBSTRING([u].[string], @p1, @p2)
```
Output MySQL:
```text
SUBSTRING(`u`.`string`, ?, ?)
```
Output PostgreSQL:
```text
SUBSTRING("u"."string", $1, $2)
```
Output SQLite:
```text
SUBSTRING("u"."string", ?, ?)
```

#### Trim
Удаляет как начальные, так и конечные пробелы из строкового выражения.
```go
function := uast.Trim(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
TRIM(`u`.`string`)
```
Output MsSQL:
```text
TRIM([u].[string])
```
Output MySQL:
```text
TRIM(`u`.`string`)
```
Output PostgreSQL:
```text
TRIM("u"."string")
```
Output SQLite:
```text
TRIM("u"."string")
```

#### Upper
Преобразует строковое выражение в верхний регистр.
```go
function := uast.Upper(uast.Field[string]("u", "string"))
```
Output MariaDB:
```text
UPPER(`u`.`string`)
```
Output MsSQL:
```text
UPPER([u].[string])
```
Output MySQL:
```text
UPPER(`u`.`string`)
```
Output PostgreSQL:
```text
UPPER("u"."string")
```
Output SQLite:
```text
UPPER("u"."string")
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
    uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    uast.Greater(uast.Field[int]("u", "number"), uast.Value(2)),
)
```
Output MariaDB:
```text
(`u`.`string` = ? AND `u`.`number` > ?)
```
Output MsSQL:
```text
([u].[string] = @p1 AND [u].[number] > @p2)
```
Output MySQL:
```text
(`u`.`string` = ? AND `u`.`number` > ?)
```
Output PostgreSQL:
```text
("u"."string" = $1 AND "u"."number" > $2)
```
Output SQLite:
```text
("u"."string" = ? AND "u"."number" > ?)
```

### Or
Комбинирует несколько условий логическим `OR`. Хотя бы одно условие должно быть истинным для истинности комбинированного выражения.
```go
logical := uast.Or(
    uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
    uast.Greater(uast.Field[int]("u", "number"), uast.Value(2)),
)
```
Output MariaDB:
```text
(`u`.`string` = ? OR `u`.`number` > ?)
```
Output MsSQL:
```text
([u].[string] = @p1 OR [u].[number] > @p2)
```
Output MySQL:
```text
(`u`.`string` = ? OR `u`.`number` > ?)
```
Output PostgreSQL:
```text
("u"."string" = $1 OR "u"."number" > $2)
```
Output SQLite:
```text
("u"."string" = ? OR "u"."number" > ?)
```

## exprSubquery
### Subquery
Оборачивает оператор `SELECT` как типизированное выражение, которое может использоваться в сравнениях (`In`, `Exists`, `Equal` и т.д.) или как колонка в операторе `SELECT`. Обобщённый параметр `u` указывает скалярный тип единственной колонки, возвращаемой подзапросом.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Field[int64]("u", "id")).From(uast.NewTable("users").As("u")))
```
Output MariaDB:
```text
(SELECT `u`.`id` FROM `users` AS `u`)
```
Output MsSQL:
```text
(SELECT [u].[id] FROM [users] AS [u])
```
Output MySQL:
```text
(SELECT `u`.`id` FROM `users` AS `u`)
```
Output PostgreSQL:
```text
(SELECT "u"."id" FROM "users" AS "u")
```
Output SQLite:
```text
(SELECT "u"."id" FROM "users" AS "u")
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