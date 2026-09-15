---
outline: deep
---

# API / Core / Опции

::: info **Информация**
Эта страница охватывает все параметры конфигурации: `clauseGroupBy`, `clauseHaving`, `clauseJoin`, `clauseOrderBy`, `clausePagination`, `clauseReturning`, `clauseSet`, `clauseUnions`, `clauseValues`, `clauseWhere`, `clauseWith`, `exprArray`, `exprBinary`, `exprComparison`, `exprConstant`, `exprField`, `exprFunction`, `exprLiteral`, `exprLogical`, `exprSubquery`, `exprValue`. Каждый параметр показан с рабочим примером кода и ожидаемым выводом.
:::**

## clauseGroupBy
Добавляет оператор GROUP BY для группировки строк по указанным колонкам или выражениям.
```go
groupBy := GroupBy(
	uast.Field[string]("u", "string"),
)
```
**Output MariaDB:**
```sql
GROUP BY `u`.`string`
```
**Output MsSQL:**
```sql
GROUP BY [u].[string]
```
**Output MySQL:**
```sql
GROUP BY `u`.`string`
```
**Output PostgreSQL:**
```sql
GROUP BY "u"."string"
```
**Output SQLite:**
```sql
GROUP BY "u"."string"
```

## clauseHaving
Добавляет оператор HAVING для фильтрации групп. Используется с GROUP BY для фильтрации агрегированных результатов.
```go
having := Having(
	uast.Greater(uast.Count(uast.Field[int64]("u", "id"), false), uast.Value[int64](2)),
)
```
**Output MariaDB:**
```sql
HAVING COUNT(`u`.`id`) > ?
```
**Output MsSQL:**
```sql
HAVING COUNT([u].[id]) > @p1
```
**Output MySQL:**
```sql
HAVING COUNT(`u`.`id`) > ?
```
**Output PostgreSQL:**
```sql
HAVING COUNT("u"."id") > $1
```
**Output SQLite:**
```sql
HAVING COUNT("u"."id") > ?
```

## clauseJoin
### Cross
Добавляет CROSS JOIN к запросу. Возвращает декартово произведение обеих таблиц.
```go
join := uast.Cross(uast.NewTable("users").As("u"))
```
**Output MariaDB:**
```sql
CROSS JOIN `users` AS `u`
```
**Output MsSQL:**
```sql
CROSS JOIN [users] AS [u]
```
**Output MySQL:**
```sql
CROSS JOIN `users` AS `u`
```
**Output PostgreSQL:**
```sql
CROSS JOIN "users" AS "u"
```
**Output SQLite:**
```sql
CROSS JOIN "users" AS "u"
```

### Full
Добавляет FULL JOIN к запросу. Возвращает все строки из обеих таблиц, с NULL там, где нет совпадений.
```go
join := uast.Full(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
FULL JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
FULL JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
FULL JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
FULL JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
FULL JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### FullOuter
Добавляет FULL OUTER JOIN к запросу. Возвращает все строки из обеих таблиц, с NULL там, где нет совпадений.
```go
join := uast.FullOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
FULL OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
FULL OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
FULL OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
FULL OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
FULL OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Inner
Добавляет INNER JOIN к запросу. Возвращает строки, имеющие совпадающие значения в обеих таблицах.
```go
join := uast.Inner(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
INNER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
INNER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
INNER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
INNER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
INNER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Left
Добавляет LEFT JOIN к запросу. Возвращает все строки из левой таблицы и совпадающие строки из правой таблицы.
```go
join := uast.Left(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
LEFT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
LEFT JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
LEFT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
LEFT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
LEFT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### LeftOuter
Добавляет LEFT OUTER JOIN к запросу. Возвращает все строки из левой таблицы и совпадающие строки из правой таблицы.
```go
join := uast.LeftOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
LEFT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
LEFT OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
LEFT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
LEFT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
LEFT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```

### Right
Добавляет RIGHT JOIN к запросу. Возвращает все строки из правой таблицы и совпадающие строки из левой таблицы. Не поддерживается SQLite.
```go
join := uast.Right(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
RIGHT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
RIGHT JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
RIGHT JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
RIGHT JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
// Not supported
```

### RightOuter
Добавляет RIGHT OUTER JOIN к запросу. Возвращает все строки из правой таблицы и совпадающие строки из левой таблицы. Не поддерживается SQLite.
```go
join := uast.RightOuter(uast.NewTable("users").As("u"), uast.Equal(uast.Field[int64]("u", "id"), uast.Field[int64]("t1", "id")))
```
**Output MariaDB:**
```sql
RIGHT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output MsSQL:**
```sql
RIGHT OUTER JOIN [users] AS [u] ON [u].[id] = [t1].[id]
```
**Output MySQL:**
```sql
RIGHT OUTER JOIN `users` AS `u` ON `u`.`id` = `t1`.`id`
```
**Output PostgreSQL:**
```sql
RIGHT OUTER JOIN "users" AS "u" ON "u"."id" = "t1"."id"
```
**Output SQLite:**
```sql
// Not supported
```

## clauseOrderBy
### Asc
Указывает порядок сортировки по возрастанию (сначала наименьшие, от А до Я). Используется для сортировки строк в запросе или в рамках оконной функции.
```go
orderBy := uast.Asc(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
`u`.`string` ASC
```
**Output MsSQL:**
```sql
[u].[string] ASC
```
**Output MySQL:**
```sql
`u`.`string` ASC
```
**Output PostgreSQL:**
```sql
"u"."string" ASC
```
**Output SQLite:**
```sql
"u"."string" ASC
```

### Desc
Указывает порядок сортировки по убыванию (сначала наибольшие, от Я до А). Используется для сортировки строк в запросе или в рамках оконной функции.
```go
orderBy := uast.Desc(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
`u`.`string` DESC
```
**Output MsSQL:**
```sql
[u].[string] DESC
```
**Output MySQL:**
```sql
`u`.`string` DESC
```
**Output PostgreSQL:**
```sql
"u"."string" DESC
```
**Output SQLite:**
```sql
"u"."string" DESC
```

## clausePagination
Определяет пагинацию для оператора SELECT с помощью `Pagination(limit, offset)`. `limit` задаёт максимальное количество возвращаемых строк. `offset` указывает количество строк, которые нужно пропустить перед возвратом результатов. Порядок отрисовки и синтаксис автоматически адаптируются к каждому диалекту.
```go
pagination := Pagination(10,0)
```
**Output MariaDB:**
```sql
LIMIT ? OFFSET ?
```
**Output MsSQL:**
```sql
OFFSET @p1 ROWS FETCH NEXT @p2 ROWS ONLY
```
**Output MySQL:**
```sql
LIMIT ? OFFSET ?
```
**Output PostgreSQL:**
```sql
LIMIT $1 OFFSET $2
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
RETURNING `u`.`id`, `u`.`string`
```
**Output MsSQL:**
```sql
**Output [u].[id], [u].[string]
```
**Output MySQL:**
```sql
// Not support
```
**Output PostgreSQL:**
```sql
RETURNING "u"."id", "u"."string"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
UPDATE `users` AS `u` SET `u`.`string` = ?
```
**Output MsSQL:**
```sql
UPDATE [users] AS [u] SET [u].[string] = @p1
```
**Output MySQL:**
```sql
UPDATE `users` AS `u` SET `u`.`string` = ?
```
**Output PostgreSQL:**
```sql
UPDATE "users" AS "u" SET "u"."string" = $1
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
UNION SELECT `u`.`string` FROM `users` AS `u` 
```
**Output MsSQL:**
```sql
UNION SELECT [u].[string] FROM [users] AS [u]
```
**Output MySQL:**
```sql
UNION SELECT `u`.`string` FROM `users` AS `u`
```
**Output PostgreSQL:**
```sql
UNION SELECT "u"."string" FROM "users" AS "u"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
UNION ALL SELECT `u`.`string` FROM `users` AS `u`
```
**Output MsSQL:**
```sql
UNION ALL SELECT [u].[string] FROM [users] AS [u]
```
**Output MySQL:**
```sql
UNION ALL SELECT `u`.`string` FROM `users` AS `u`
```
**Output PostgreSQL:**
```sql
UNION ALL SELECT "u"."string" FROM "users" AS "u"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
EXCEPT SELECT `u`.`string` FROM `users` AS `u`
```
**Output MsSQL:**
```sql
EXCEPT SELECT [u].[string] FROM [users] AS [u]
```
**Output MySQL:**
```sql
EXCEPT SELECT `u`.`string` FROM `users` AS `u`
```
**Output PostgreSQL:**
```sql
EXCEPT SELECT "u"."string" FROM "users" AS "u"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
INTERSECT SELECT `u`.`string` FROM `users` AS `u`
```
**Output MsSQL:**
```sql
INTERSECT SELECT [u].[string] FROM [users] AS [u]
```
**Output MySQL:**
```sql
INTERSECT SELECT `u`.`string` FROM `users` AS `u`
```
**Output PostgreSQL:**
```sql
INTERSECT SELECT "u"."string" FROM "users" AS "u"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
VALUES (?, ?)
```
**Output MsSQL:**
```sql
VALUES (@p1, @p2)
```
**Output MySQL:**
```sql
VALUES (?, ?)
```
**Output PostgreSQL:**
```sql
VALUES ($1, $2)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
```
**Output MsSQL:**
```sql
// Not supported
```
**Output MySQL:**
```sql
VALUES (?, ?) ON DUPLICATE KEY UPDATE `string` = ?
```
**Output PostgreSQL:**
```sql
VALUES ($1, $2) ON CONFLICT DO UPDATE SET "string" = $3
```
**Output SQLite:**
```sql
VALUES (?, ?) ON CONFLICT DO UPDATE SET "string" = ?
```

## clauseWhere
Добавляет оператор WHERE для фильтрации строк перед группировкой или агрегацией. Принимает выражения сравнения, логические операторы и подзапросы.
```go
where = Where(
	uast.Equal(uast.Field[string]("u", "string"), uast.Value("active")),
)
```
**Output MariaDB:**
```sql
WHERE `u`.`string` = ?
```
**Output MsSQL:**
```sql
WHERE [u].[string] = @p1
```
**Output MySQL:**
```sql
WHERE `u`.`string` = ?
```
**Output PostgreSQL:**
```sql
WHERE "u"."string" = $1
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ?)
```
**Output MsSQL:**
```sql
WITH [cte_norecursive] ([id], [string]) AS (SELECT [u].[id], [u].[string] FROM [users] AS [u] WHERE [u].[string] = @p1)
```
**Output MySQL:**
```sql
WITH `cte_norecursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ?)
```
**Output PostgreSQL:**
```sql
WITH "cte_norecursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = $1)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ? UNION ALL SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` INNER JOIN `cte_recursive` AS `rec` ON `u`.`id` = `rec`.`id`)
```
**Output MsSQL:**
```sql
WITH RECURSIVE [cte_recursive] ([id], [string]) AS (SELECT [u].[id], [u].[string] FROM [users] AS [u] WHERE [u].[string] = @p1 UNION ALL SELECT [u].[id], [u].[string] FROM [users] AS [u] INNER JOIN [cte_recursive] AS [rec] ON [u].[id] = [rec].[id])
```
**Output MySQL:**
```sql
WITH RECURSIVE `cte_recursive` (`id`, `string`) AS (SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` WHERE `u`.`string` = ? UNION ALL SELECT `u`.`id`, `u`.`string` FROM `users` AS `u` INNER JOIN `cte_recursive` AS `rec` ON `u`.`id` = `rec`.`id`)
```
**Output PostgreSQL:**
```sql
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = $1 UNION ALL SELECT "u"."id", "u"."string" FROM "users" AS "u" INNER JOIN "cte_recursive" AS "rec" ON "u"."id" = "rec"."id")
```
**Output SQLite:**
```sql
WITH RECURSIVE "cte_recursive" ("id", "string") AS (SELECT "u"."id", "u"."string" FROM "users" AS "u" WHERE "u"."string" = ? UNION ALL SELECT "u"."id", "u"."string" FROM "users" AS "u" INNER JOIN "cte_recursive" AS "rec" ON "u"."id" = "rec"."id")
```

## exprArray
### Array
Создаёт выражение массива для использования в SQL-запросах.
```go
array := uast.Array(0, 1, 2)
```
**Output MariaDB:**
```sql
ARRAY[?, ?, ?]
```
**Output MsSQL:**
```sql
ARRAY[@p1, @p2, @p3]
```
**Output MySQL:**
```sql
ARRAY[?, ?, ?]
```
**Output PostgreSQL:**
```sql
ARRAY[$1, $2, $3]
```
**Output SQLite:**
```sql
ARRAY[?, ?, ?]

## exprBinary
### BitwiseAnd
Выполняет побитовую операцию И между двумя выражениями.
```go
binary := uast.BitwiseAnd(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
**Output MsSQL:**
```sql
[u].[number] & @p1
```
**Output MySQL:**
```sql
`u`.`number` & ?
```
**Output PostgreSQL:**
```sql
"u"."number" & $1
```

### BitwiseOr
Выполняет побитовую операцию ИЛИ между двумя выражениями.
```go
binary := uast.BitwiseOr(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
**Output MariaDB:**
```sql
`u`.`number` | ?
```
**Output MsSQL:**
```sql
[u].[number] | @p1
```
**Output MySQL:**
```sql
`u`.`number` | ?
```
**Output PostgreSQL:**
```sql
"u"."number" | $1
```
**Output SQLite:**
```sql
"u"."number" | ?
```

### BitwiseXor
Выполняет побитовую операцию исключающего ИЛИ между двумя выражениями.
```go
binary := uast.BitwiseXor(uast.Field[int]("u", "number"), uast.Value(0b0010))
```
**Output MariaDB:**
```sql
`u`.`number` ^ ?
```
**Output MsSQL:**
```sql
[u].[number] ^ @p1
```
**Output MySQL:**
```sql
`u`.`number` ^ ?
```
**Output PostgreSQL:**
```sql
"u"."number" ^ $1
```
**Output SQLite:**
```sql
"u"."number" ^ ?
```

### Divide
Делит левое выражение на правое.
```go
binary := uast.Divide(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` / ?
```
**Output MsSQL:**
```sql
[u].[number] / @p1
```
**Output MySQL:**
```sql
`u`.`number` / ?
```
**Output PostgreSQL:**
```sql
"u"."number" / $1
```
**Output SQLite:**
```sql
"u"."number" / ?
```

### Minus
Вычитает правое выражение из левого.
```go
binary := uast.Minus(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` - ?
```
**Output MsSQL:**
```sql
[u].[number] - @p1
```
**Output MySQL:**
```sql
`u`.`number` - ?
```
**Output PostgreSQL:**
```sql
"u"."number" - $1
```
**Output SQLite:**
```sql
"u"."number" - ?
```

### Modulo
Возвращает остаток от деления левого выражения на правое.
```go
binary := uast.Modulo(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` % ?
```
**Output MsSQL:**
```sql
[u].[number] % @p1
```
**Output MySQL:**
```sql
`u`.`number` % ?
```
**Output PostgreSQL:**
```sql
"u"."number" % $1
```
**Output SQLite:**
```sql
"u"."number" % ?
```

### Multiply
Умножает левое выражение на правое.
```go
binary := uast.Multiply(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` * ?
```
**Output MsSQL:**
```sql
[u].[number] * @p1
```
**Output MySQL:**
```sql
`u`.`number` * ?
```
**Output PostgreSQL:**
```sql
"u"."number" * $1
```
**Output SQLite:**
```sql
"u"."number" * ?
```

### Plus
Складывает левое выражение с правым.
```go
binary := uast.Plus(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` + ?
```
**Output MsSQL:**
```sql
[u].[number] + @p1
```
**Output MySQL:**
```sql
`u`.`number` + ?
```
**Output PostgreSQL:**
```sql
"u"."number" + $1
```
**Output SQLite:**
```sql
"u"."number" + ?
```

### ShiftLeft
Выполняет побитовый сдвиг влево левого выражения на количество бит, указанное в правом выражении.
```go
binary := uast.ShiftLeft(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` << ?
```
**Output MsSQL:**
```sql
[u].[number] << @p1
```
**Output MySQL:**
```sql
`u`.`number` << ?
```
**Output PostgreSQL:**
```sql
"u"."number" << $1
```
**Output SQLite:**
```sql
"u"."number" << ?
```

### ShiftRight
Выполняет побитовый сдвиг вправо левого выражения на количество бит, указанное в правом выражении.
```go
binary := uast.ShiftRight(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` >> ?
```
**Output MsSQL:**
```sql
[u].[number] >> @p1
```
**Output MySQL:**
```sql
`u`.`number` >> ?
```
**Output PostgreSQL:**
```sql
"u"."number" >> $1
```
**Output SQLite:**
```sql
"u"."number" >> ?
```

## exprComparison
### Between
Проверяет, попадает ли левое выражение в диапазон, заданный valueStart и valueEnd (включительно).
```go
comparison := uast.Between(uast.Field[int]("u", "number"), uast.Value(0), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` BETWEEN ? AND ?
```
**Output MsSQL:**
```sql
[u].[number] BETWEEN @p1 AND @p2
```
**Output MySQL:**
```sql
`u`.`number` BETWEEN ? AND ?
```
**Output PostgreSQL:**
```sql
"u"."number" BETWEEN $1 AND $2
```
**Output SQLite:**
```sql
"u"."number" BETWEEN ? AND ?
```

### Equal
Сравнивает два выражения на равенство (`=`).
```go
comparison := uast.Equal(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` = ?
```
**Output MsSQL:**
```sql
[u].[number] = @p1
```
**Output MySQL:**
```sql
`u`.`number` = ?
```
**Output PostgreSQL:**
```sql
"u"."number" = $1
```
**Output SQLite:**
```sql
"u"."number" = ?
```

### Exists
Проверяет, возвращает ли подзапрос какие-либо строки. Возвращает `true` если существует хотя бы одна строка.
```go
comparison := uast.Exists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("users").As("u"))))
```
**Output MariaDB:**
```sql
EXISTS (SELECT 1 FROM `users` AS `u`)
```
**Output MsSQL:**
```sql
EXISTS (SELECT 1 FROM [users] AS [u])
```
**Output MySQL:**
```sql
EXISTS (SELECT 1 FROM `users` AS `u`)
```
**Output PostgreSQL:**
```sql
EXISTS (SELECT 1 FROM "users" AS "u")
```
**Output SQLite:**
```sql
EXISTS (SELECT 1 FROM "users" AS "u")
```

### Greater
Сравнивает, больше ли левое выражение правого (`>`).
```go
comparison := uast.Greater(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` > ?
```
**Output MsSQL:**
```sql
[u].[number] > @p1
```
**Output MySQL:**
```sql
`u`.`number` > ?
```
**Output PostgreSQL:**
```sql
"u"."number" > $1
```
**Output SQLite:**
```sql
"u"."number" > ?
```

### GreaterEqual
Сравнивает, больше или равно ли левое выражение правому (`>=`).
```go
comparison := uast.GreaterEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` >= ?
```
**Output MsSQL:**
```sql
[u].[number] >= @p1
```
**Output MySQL:**
```sql
`u`.`number` >= ?
```
**Output PostgreSQL:**
```sql
"u"."number" >= $1
```
**Output SQLite:**
```sql
"u"."number" >= ?
```

### ILike
Выполняет регистронезависимое сравнение с шаблоном. Правое выражение должно содержать шаблон с `%` (любая последовательность) и `_` (один символ).
```go
comparison := uast.ILike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
**Output MariaDB:**
```sql
LOWER(`u`.`string`) LIKE LOWER(?)
```
**Output MsSQL:**
```sql
LOWER([u].[string]) LIKE LOWER(@p1)
```
**Output MySQL:**
```sql
LOWER(`u`.`string`) LIKE LOWER(?)
```
**Output PostgreSQL:**
```sql
"u"."string" ILIKE $1
```
**Output SQLite:**
```sql
LOWER("u"."string") LIKE LOWER(?)
```

### In
Проверяет, соответствует ли левое выражение любому значению, содержащемуся в правом выражении (обычно подзапрос или массив).
```go
comparison := uast.In(uast.Field[string]("u", "string"), uast.Array("active", "pending"))
```
**Output MariaDB:**
```sql
`u`.`string` IN (?, ?)
```
**Output MsSQL:**
```sql
[u].[string] IN (@p1, @p2)
```
**Output MySQL:**
```sql
`u`.`string` IN (?, ?)
```
**Output PostgreSQL:**
```sql
"u"."string" IN ($1, $2)
```
**Output SQLite:**
```sql
"u"."string" IN (?, ?)
```

### IsNotNull
Проверяет, что выражение не `NULL`.
```go
comparison := uast.IsNotNull(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
`u`.`string` IS NOT NULL
```
**Output MsSQL:**
```sql
[u].[string] IS NOT NULL
```
**Output MySQL:**
```sql
`u`.`string` IS NOT NULL
```
**Output PostgreSQL:**
```sql
"u"."string" IS NOT NULL
```
**Output SQLite:**
```sql
"u"."string" IS NOT NULL
```

### IsNull
Проверяет, что выражение является `NULL`.
```go
comparison := uast.IsNull(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
`u`.`string` IS NULL
```
**Output MsSQL:**
```sql
[u].[string] IS NULL
```
**Output MySQL:**
```sql
`u`.`string` IS NULL
```
**Output PostgreSQL:**
```sql
"u"."string" IS NULL
```
**Output SQLite:**
```sql
"u"."string" IS NULL
```

### Less
Сравнивает, меньше ли левое выражение правого (`<`).
```go
comparison := uast.Less(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` < ?
```
**Output MsSQL:**
```sql
[u].[number] < @p1
```
**Output MySQL:**
```sql
`u`.`number` < ?
```
**Output PostgreSQL:**
```sql
"u"."number" < $1
```
**Output SQLite:**
```sql
"u"."number" < ?
```

### LessEqual
Сравнивает, меньше или равно ли левое выражение правому (`<=`).
```go
comparison := uast.LessEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` <= ?
```
**Output MsSQL:**
```sql
[u].[number] <= @p1
```
**Output MySQL:**
```sql
`u`.`number` <= ?
```
**Output PostgreSQL:**
```sql
"u"."number" <= $1
```
**Output SQLite:**
```sql
"u"."number" <= ?
```

### Like
Выполняет регистрозависимое сравнение с шаблоном. Правое выражение должно содержать шаблон с `%` и `_`.
```go
comparison := uast.Like(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
**Output MariaDB:**
```sql
`u`.`string` LIKE ?
```
**Output MsSQL:**
```sql
[u].[number] LIKE @p1
```
**Output MySQL:**
```sql
`u`.`string` LIKE ?
```
**Output PostgreSQL:**
```sql
"u"."string" LIKE $1
```
**Output SQLite:**
```sql
"u"."string" LIKE ?
```

### NotBetween
Проверяет, находится ли левое выражение вне диапазона, заданного `valueStart` и `valueEnd`.
```go
comparison := uast.NotBetween(uast.Field[int]("u", "number"), uast.Value(0), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` NOT BETWEEN ? AND ?
```
**Output MsSQL:**
```sql
[u].[number] NOT BETWEEN @p1 AND @p2
```
**Output MySQL:**
```sql
`u`.`number` NOT BETWEEN ? AND ?
```
**Output PostgreSQL:**
```sql
"u"."number" NOT BETWEEN $1 AND $2
```
**Output SQLite:**
```sql
"u"."number" NOT BETWEEN ? AND ?
```

### NotEqual
Сравнивает два выражения на неравенство (`!=` or `<>`).
```go
comparison := uast.NotEqual(uast.Field[int]("u", "number"), uast.Value(2))
```
**Output MariaDB:**
```sql
`u`.`number` != ?
```
**Output MsSQL:**
```sql
[u].[number] != @p1
```
**Output MySQL:**
```sql
`u`.`number` != ?
```
**Output PostgreSQL:**
```sql
"u"."number" != $1
```
**Output SQLite:**
```sql
"u"."number" != ?
```

### NotExists
Проверяет, что подзапрос не возвращает строк. Возвращает `true` если результат подзапроса пуст.
```go
comparison := uast.NotExists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.NewTable("users").As("u"))))
```
**Output MariaDB:**
```sql
NOT EXISTS (SELECT 1 FROM `users` AS `u`)
```
**Output MsSQL:**
```sql
NOT EXISTS (SELECT 1 FROM [users] AS [u])
```
**Output MySQL:**
```sql
NOT EXISTS (SELECT 1 FROM `users` AS `u`)
```
**Output PostgreSQL:**
```sql
NOT EXISTS (SELECT 1 FROM "users" AS "u")
```
**Output SQLite:**
```sql
NOT EXISTS (SELECT 1 FROM "users" AS "u")
```

### NotILike
Выполняет отрицательное регистронезависимое сравнение с шаблоном.
```go
comparison := uast.NotILike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
**Output MariaDB:**
```sql
LOWER(`u`.`string`) NOT LIKE LOWER(?)
```
**Output MsSQL:**
```sql
LOWER([u].[string]) NOT LIKE LOWER(@p1)
```
**Output MySQL:**
```sql
LOWER(`u`.`string`) NOT LIKE LOWER(?)
```
**Output PostgreSQL:**
```sql
"u"."string" NOT ILIKE $1
```
**Output SQLite:**
```sql
LOWER("u"."string") NOT LIKE LOWER(?)
```

### NotIn
Проверяет, что левое выражение не соответствует ни одному значению, содержащемуся в правом выражении.
```go
comparison := uast.NotIn(uast.Field[string]("u", "string"), uast.Array("active", "pending"))
```
**Output MariaDB:**
```sql
`u`.`string` NOT IN (?, ?)
```
**Output MsSQL:**
```sql
[u].[string] NOT IN (@p1, @p2)
```
**Output MySQL:**
```sql
`u`.`string` NOT IN (?, ?)
```
**Output PostgreSQL:**
```sql
"u"."string" NOT IN ($1, $2)
```
**Output SQLite:**
```sql
"u"."string" NOT IN (?, ?)
```

### NotLike
Выполняет отрицательное регистрозависимое сравнение с шаблоном.
```go
comparison := uast.NotLike(uast.Field[string]("u", "string"), uast.Value("%ivan%"))
```
**Output MariaDB:**
```sql
`u`.`string` NOT LIKE ?
```
**Output MsSQL:**
```sql
[u].[string] NOT LIKE @p1
```
**Output MySQL:**
```sql
`u`.`string` NOT LIKE ?
```
**Output PostgreSQL:**
```sql
"u"."string" NOT LIKE $1
```
**Output SQLite:**
```sql
"u"."string" NOT LIKE ?
```

## exprConstant
### ConstBoolFalse
Возвращает константное булево выражение `FALSE`.
```go
constant := uast.ConstBoolFalse()
```
**Output SQL:**
```sql
FALSE
```

### ConstBoolTrue
Возвращает константное булево выражение `TRUE`.
```go
constant := uast.ConstBoolTrue()
```
**Output SQL:**
```sql
TRUE
```

### ConstFloat32One
Возвращает константное значение `float32` равное `1.0`. 
```go
constant := uast.ConstFloat32One()
```
**Output SQL:**
```sql
1.0
```

### ConstFloat64One
Возвращает константное значение `float64` равное `1.000000`.
```go
constant := uast.ConstFloat64One()
```
**Output SQL:**
```sql
1.000000
```

### ConstIntOne
Возвращает константное значение `int` равное `1`.
```go
constant := uast.ConstIntOne()
```
**Output SQL:**
```sql
1
```

### ConstInt8One
Возвращает константное значение `int8` равное `1`.
```go
constant := uast.ConstInt8One()
```
**Output SQL:**
```sql
1
```

### ConstInt16One
Возвращает константное значение `int16` равное `1`.
```go
constant := uast.ConstInt16One()
```
**Output SQL:**
```sql
1
```

### ConstInt32One
Возвращает константное значение `int32` равное `1`.
```go
constant := uast.ConstInt32One()
```
**Output SQL:**
```sql
1
```

### ConstInt64One
Возвращает константное значение `int64` равное `1`.
```go
constant := uast.ConstInt64One()
```
**Output SQL:**
```sql
1
```

### ConstStringDefault
Возвращает константное значение `string` равное `DEFAULT`.
```go
constant := uast.ConstStringDefault()
```
**Output SQL:**
```sql
DEFAULT
```

### ConstStringNull
Возвращает константное значение `string` равное `NULL`.
```go
constant := uast.ConstStringNull()
```
**Output SQL:**
```sql
NULL
```

### ConstUintOne
Возвращает константное значение `uint` равное `1`.
```go
constant := uast.ConstUintOne()
```
**Output SQL:**
```sql
1
```

### ConstUint8One
Возвращает константное значение `uint8` равное `1`.
```go
constant := uast.ConstUint8One()
```
**Output SQL:**
```sql
1
```

### ConstUint16One
Возвращает константное значение `uint16` равное `1`.
```go
constant := uast.ConstUint16One()
```
**Output SQL:**
```sql
1
```

### ConstUint32One
Возвращает константное значение `uint32` равное `1`.
```go
constant := uast.ConstUint32One()
```
**Output SQL:**
```sql
1
```

### ConstUint64One
Возвращает константное значение `uint64` равное `1`.
```go
constant := uast.ConstUint64One()
```
**Output SQL:**
```sql
1
```

## exprField
### Field
Создаёт ссылку на колонку таблицы, опционально квалифицированную псевдонимом таблицы. Это основной способ ссылаться на колонки базы данных в выражениях.
```go
field := uast.Field[string]("u", "string")
```
**Output MariaDB:**
```sql
`u`.`string`
```
**Output MsSQL:**
```sql
[u].[string]
```
**Output MySQL:**
```sql
`u`.`string`
```
**Output PostgreSQL:**
```sql
"u"."string"
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
AVG(`u`.`number`)
AVG(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
AVG([u].[number])
AVG(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
AVG(`u`.`number`)
AVG(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
AVG("u"."number")
AVG(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
AVG("u"."number")
AVG(DISTINCT "u"."number")
```

#### BitAnd
Возвращает побитовое И всех битов в выражении. Имеет смысл только для целочисленных типов.
```go
function := uast.BitAnd(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitAnd(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
BIT_AND(`u`.`number`)
BIT_AND(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
BIT_AND([u].[number])
BIT_AND(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
BIT_AND(`u`.`number`)
BIT_AND(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
BIT_AND("u"."number")
BIT_AND(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
BIT_AND("u"."number")
BIT_AND(DISTINCT "u"."number")
```

#### BitOr
Возвращает побитовое ИЛИ всех битов в выражении.
```go
function := uast.BitOr(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitOr(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
BIT_OR(`u`.`number`)
BIT_OR(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
BIT_OR([u].[number])
BIT_OR(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
BIT_OR(`u`.`number`)
BIT_OR(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
BIT_OR("u"."number")
BIT_OR(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
BIT_OR("u"."number")
BIT_OR(DISTINCT "u"."number")
```

#### BitXor
Возвращает побитовое исключающее ИЛИ всех битов в выражении.
```go
function := uast.BitXor(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.BitXor(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
BIT_XOR(`u`.`number`)
BIT_XOR(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
BIT_XOR([u].[number])
BIT_XOR(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
BIT_XOR(`u`.`number`)
BIT_XOR(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
BIT_XOR("u"."number")
BIT_XOR(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
BIT_XOR("u"."number")
BIT_XOR(DISTINCT "u"."number")
```

#### Count
Возвращает количество строк, соответствующих запросу, или количество не-NULL значений, если указано выражение. Когда `distinct` равен `true`, подсчитываются только уникальные значения.
```go
function := uast.Count(uast.Field[string]("u", "string"), false)
functionWithDistinct := uast.Count(uast.Field[string]("u", "string"), true)
```
**Output MariaDB:**
```sql
COUNT(`u`.`string`)
COUNT(DISTINCT `u`.`string`)
```
**Output MsSQL:**
```sql
COUNT([u].[string])
COUNT(DISTINCT [u].[string])
```
**Output MySQL:**
```sql
COUNT(`u`.`string`)
COUNT(DISTINCT `u`.`string`)
```
**Output PostgreSQL:**
```sql
COUNT("u"."string")
COUNT(DISTINCT "u"."string")
```
**Output SQLite:**
```sql
COUNT("u"."string")
COUNT(DISTINCT "u"."string")
```

#### GroupConcat
Объединяет значения из группы в одну строку, разделённую стандартным разделителем (обычно запятая). Флаг `distinct` удаляет дубликаты перед объединением.
```go
function := uast.GroupConcat(uast.Field[string]("u", "string"), false)
functionWithDistinct := uast.GroupConcat(uast.Field[string]("u", "string"), true)
```
**Output MariaDB:**
```sql
GROUP_CONCAT(`u`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `u`.`string` SEPARATOR ',')
```
**Output MsSQL:**
```sql
GROUP_CONCAT([u].[string], ',')
GROUP_CONCAT(DISTINCT [u].[string], ',')
```
**Output MySQL:**
```sql
GROUP_CONCAT(`u`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `u`.`string` SEPARATOR ',')
```
**Output PostgreSQL:**
```sql
STRING_AGG("u"."string", ',')
STRING_AGG(DISTINCT "u"."string", ',')
```
**Output SQLite:**
```sql
GROUP_CONCAT("u"."string" SEPARATOR ',')
GROUP_CONCAT(DISTINCT "u"."string" SEPARATOR ',')
```

#### Max
Возвращает максимальное значение выражения по всем строкам в группе.
```go
function := uast.Max(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Max(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
MAX(`u`.`number`)
MAX(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
MAX([u].[number])
MAX(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
MAX(`u`.`number`)
MAX(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
MAX("u"."number")
MAX(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
MAX("u"."number")
MAX(DISTINCT "u"."number")
```

#### Min
Возвращает минимальное значение выражения по всем строкам в группе.
```go
function := uast.Min(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Min(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
MIN(`u`.`number`)
MIN(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
MIN([u].[number])
MIN(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
MIN(`u`.`number`)
MIN(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
MIN("u"."number")
MIN(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
MIN("u"."number")
MIN(DISTINCT "u"."number")
```

#### StdDev
Возвращает популяционное стандартное отклонение выражения.
```go
function := uast.StdDev(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.StdDev(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
STDDEV(`u`.`number`)
STDDEV(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
STDEV([u].[number])
STDEV(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
STDDEV(`u`.`number`)
STDDEV(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
STDDEV_SAMP("u"."number")
STDDEV_SAMP(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
STDEV("u"."number")
STDEV(DISTINCT "u"."number")
```

#### Sum
Возвращает сумму всех значений в выражении. Если `distinct` равен `true`, суммируются только уникальные значения.
```go
function := uast.Sum(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Sum(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
SUM(`u`.`number`)
SUM(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
SUM([u].[number])
SUM(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
SUM(`u`.`number`)
SUM(DISTINCT `u`.`number`)
```
**Output PostgreSQL:**
```sql
SUM("u"."number")
SUM(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
SUM("u"."number")
SUM(DISTINCT "u"."number")
```

#### Variance
Возвращает популяционную дисперсию выражения.
```go
function := uast.Variance(uast.Field[int]("u", "number"), false)
functionWithDistinct := uast.Variance(uast.Field[int]("u", "number"), true)
```
**Output MariaDB:**
```sql
VARIANCE(`u`.`number`)
VARIANCE(DISTINCT `u`.`number`)
```
**Output MsSQL:**
```sql
VAR([u].[number])
VAR(DISTINCT [u].[number])
```
**Output MySQL:**
```sql
VARIANCE(`u`.`number`)
VARIANCE(DISTINCT "u"."number")
```
**Output PostgreSQL:**
```sql
VAR_SAMP("u"."number")
VAR_SAMP(DISTINCT "u"."number")
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
FIRST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
FIRST_VALUE([u].[string]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
FIRST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
FIRST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
LAG(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
**Output MsSQL:**
```sql
LAG([u].[number], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[date] ASC)
```
**Output MySQL:**
```sql
LAG(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
**Output PostgreSQL:**
```sql
LAG("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
LAST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
**Output MsSQL:**
```sql
LAST_VALUE([u].[string]) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
**Output MySQL:**
```sql
LAST_VALUE(`u`.`string`) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
**Output PostgreSQL:**
```sql
LAST_VALUE("u"."string") OVER (PARTITION BY "u"."id" ORDER BY "u"."number" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
LEAD(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
**Output MsSQL:**
```sql
LEAD([u].[number], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[date] ASC)
```
**Output MySQL:**
```sql
LEAD(`u`.`number`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`date` ASC)
```
**Output PostgreSQL:**
```sql
LEAD("u"."number", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."date" ASC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
NTH_VALUE(`u`.`string`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
**Output MsSQL:**
```sql
NTH_VALUE([u].[string], 2) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
**Output MySQL:**
```sql
NTH_VALUE(`u`.`string`, 2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
**Output PostgreSQL:**
```sql
NTH_VALUE("u"."string", 2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
CASE WHEN `u`.`number` < ? THEN ? ELSE ? END
```
**Output MsSQL:**
```sql
CASE WHEN [u].[number] < @p1 THEN @p2 ELSE @p3 END
```
**Output MySQL:**
```sql
CASE WHEN `u`.`number` < ? THEN ? ELSE ? END
```
**Output PostgreSQL:**
```sql
CASE WHEN "u"."number" < $1 THEN $2 ELSE $3 END
```
**Output SQLite:**
```sql
CASE WHEN "u"."number" < ? THEN ? ELSE ? END
```

#### Coalesce
Возвращает первое не-NULL выражение из предоставленного списка. Полезно для указания запасных значений.
```go
function := uast.Coalesce(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
**Output MariaDB:**
```sql
COALESCE(`u`.`createat`, `u`.`updateat`)
```
**Output MsSQL:**
```sql
COALESCE([u].[createat], [u].[updateat])
```
**Output MySQL:**
```sql
COALESCE(`u`.`createat`, `u`.`updateat`)
```
**Output PostgreSQL:**
```sql
COALESCE("u"."createat", "u"."updateat")
```
**Output SQLite:**
```sql
COALESCE("u"."createat", "u"."updateat")
```

#### Greatest
Возвращает наибольшее значение из предоставленного списка выражений.
```go
function := uast.Greatest(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
**Output MariaDB:**
```sql
GREATEST(`u`.`createat`, `u`.`updateat`)
```
**Output MsSQL:**
```sql
GREATEST([u].[createat], [u].[updateat])
```
**Output MySQL:**
```sql
GREATEST(`u`.`createat`, `u`.`updateat`)
```
**Output PostgreSQL:**
```sql
GREATEST("u"."createat", "u"."updateat")
```
**Output SQLite:**
```sql
GREATEST("u"."createat", "u"."updateat")
```

#### Least
Возвращает наименьшее значение из предоставленного списка выражений.
```go
function := uast.Least(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
**Output MariaDB:**
```sql
LEAST(`u`.`createat`, `u`.`updateat`)
```
**Output MsSQL:**
```sql
LEAST([u].[createat], [u].[updateat])
```
**Output MySQL:**
```sql
LEAST(`u`.`createat`, `u`.`updateat`)
```
**Output PostgreSQL:**
```sql
LEAST("u"."createat", "u"."updateat")
```
**Output SQLite:**
```sql
LEAST("u"."createat", "u"."updateat")
```

#### NullIf
Возвращает `NULL` если два выражения равны; иначе возвращает первое выражение.
```go
function := uast.NullIf(uast.Field[time.Time]("u", "createat"), uast.Field[time.Time]("u", "updateat"))
```
**Output MariaDB:**
```sql
NULLIF(`u`.`createat`, `u`.`updateat`)
```
**Output MsSQL:**
```sql
NULLIF([u].[createat], [u].[updateat])
```
**Output MySQL:**
```sql
NULLIF(`u`.`createat`, `u`.`updateat`)
```
**Output PostgreSQL:**
```sql
NULLIF("u"."createat", "u"."updateat")
```
**Output SQLite:**
```sql
NULLIF("u"."createat", "u"."updateat")
```

### Convert
#### Cast
Преобразует выражение к указанному типу данных.
```go
function := uast.Cast(uast.Field[int]("u", "number"), uast.TypeString)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS CHAR)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS NVARCHAR)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS CHAR)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS VARCHAR)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

#### CharLength
Возвращает количество символов в строковом выражении.
```go
function := uast.CharLength(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
CHAR_LENGTH(`u`.`string`)
```
**Output MsSQL:**
```sql
CHAR_LENGTH([u].[string])
```
**Output MySQL:**
```sql
CHAR_LENGTH(`u`.`string`)
```
**Output PostgreSQL:**
```sql
CHAR_LENGTH("u"."string")
```
**Output SQLite:**
```sql
CHAR_LENGTH("u"."string")
```

#### DateFormat
Форматирует выражение даты/времени в соответствии с указанной маской формата.
```go
function := uast.DateFormat(uast.Field[time.Time]("u", "createat"), uast.Value("%Y-%m-%d"))
```
**Output MariaDB:**
```sql
DATE_FORMAT(`u`.`createat`, '%Y-%m-%d')
```
**Output MsSQL:**
```sql
FORMAT([u].[createat], '%Y-%m-%d')
```
**Output MySQL:**
```sql
DATE_FORMAT(`u`.`createat`, '%Y-%m-%d')
```
**Output PostgreSQL:**
```sql
TO_CHAR("u"."createat", '%Y-%m-%d')
```
**Output SQLite:**
```sql
STRFTIME("u"."createat", '%Y-%m-%d')
```

#### Degrees
Преобразует угол из радиан в градусы.
```go
function := uast.Degrees(uast.Field[int]("u", "number"))
```
**Output MariaDB:**
```sql
DEGREES(`u`.`number`)
```
**Output MsSQL:**
```sql
DEGREES([u].[number])
```
**Output MySQL:**
```sql
DEGREES(`u`.`number`)
```
**Output PostgreSQL:**
```sql
DEGREES("u"."number")
```
**Output SQLite:**
```sql
DEGREES("u"."number")
```

#### Length
Возвращает длину строкового выражения в байтах.
```go
function := uast.Length(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
LENGTH(`u`.`string`)
```
**Output MsSQL:**
```sql
LEN([u].[string])
```
**Output MySQL:**
```sql
LENGTH(`u`.`string`)
```
**Output PostgreSQL:**
```sql
LENGTH("u"."string")
```
**Output SQLite:**
```sql
LENGTH("u"."string")
```

#### Position
Возвращает начальную позицию первого вхождения подстроки в строку.
```go
function := uast.Position(uast.Field[string]("u", "string"), uast.Value("old"))
```
**Output MariaDB:**
```sql
POSITION(? IN `u`.`string`)
```
**Output MsSQL:**
```sql
CHARINDEX(@p1, [u].[string])
```
**Output MySQL:**
```sql
POSITION(? IN `u`.`string`)
```
**Output PostgreSQL:**
```sql
POSITION($1 IN "u"."string")
```
**Output SQLite:**
```sql
POSITION(? IN "u"."string")
```

#### Radians
Преобразует угол из градусов в радианы.
```go
function := uast.Radians(uast.Field[int]("u", "number"))
```
**Output MariaDB:**
```sql
RADIANS(`u`.`number`)
```
**Output MsSQL:**
```sql
RADIANS([u].[number])
```
**Output MySQL:**
```sql
RADIANS(`u`.`number`)
```
**Output PostgreSQL:**
```sql
RADIANS("u"."number")
```
**Output SQLite:**
```sql
RADIANS("u"."number")
```

### Date and time
#### CurDate
Возвращает текущую дату (без времени).
```go
function := uast.CurDate()
```
**Output MariaDB:**
```sql
CURDATE()
```
**Output MsSQL:**
```sql
CAST(GETDATE() AS DATE)
```
**Output MySQL:**
```sql
CURDATE()
```
**Output PostgreSQL:**
```sql
CURRENT_DATE
```
**Output SQLite:**
```sql
DATE('now')
```

#### CurTime
Возвращает текущее время (без даты).
```go
function := uast.CurTime()
```
**Output MariaDB:**
```sql
CURTIME()
```
**Output MsSQL:**
```sql
CAST(GETDATE() AS TIME)
```
**Output MySQL:**
```sql
CURTIME()
```
**Output PostgreSQL:**
```sql
CURRENT_TIME
```
**Output SQLite:**
```sql
TIME('now')
```

#### DateAdd
Добавляет интервал даты/времени к выражению даты/времени и возвращает результирующую дату/время.
```go
function := uast.DateAdd(uast.Field[time.Time]("u", "createat"), uast.Value("2 DAY"))
```
**Output MariaDB:**
```sql
DATE_ADD(`u`.`createat`, INTERVAL 2 DAY)
```
**Output MsSQL:**
```sql
DATEADD(DAY, 2, [u].[createat])
```
**Output MySQL:**
```sql
DATE_ADD(`u`.`createat`, INTERVAL 2 DAY)
```
**Output PostgreSQL:**
```sql
("u"."createat" + INTERVAL '2 DAY')
```
**Output SQLite:**
```sql
DATETIME("u"."createat", '+2 DAY')
```

#### DateDiff
Возвращает разницу в днях между двумя выражениями даты/времени (`datetimeEnd` - `datetimeStart`).
```go
function := uast.DateDiff(uast.Field[time.Time]("u", "updateat"), uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
DATEDIFF(`u`.`updateat`, `u`.`createat`)
```
**Output MsSQL:**
```sql
DATEDIFF([u].[updateat], [u].[createat])
```
**Output MySQL:**
```sql
DATEDIFF(`u`.`updateat`, `u`.`createat`)
```
**Output PostgreSQL:**
```sql
DATE_PART('day', "u"."updateat" - "u"."createat")
```
**Output SQLite:**
```sql
DATEDIFF("u"."updateat", "u"."createat")
```

#### DateSub
Вычитает интервал даты/времени из выражения даты/времени и возвращает результирующую дату/время.
```go
function := uast.DateSub(uast.Field[time.Time]("u", "createat"), uast.Value("2 DAY"))
```
**Output MariaDB:**
```sql
DATE_SUB(`u`.`createat`, INTERVAL 2 DAY)
```
**Output MsSQL:**
```sql
DATEADD(DAY, -2, [u].[createat])
```
**Output MySQL:**
```sql
DATE_SUB(`u`.`createat`, INTERVAL 2 DAY)
```
**Output PostgreSQL:**
```sql
("u"."createat" - INTERVAL '2 DAY')
```
**Output SQLite:**
```sql
DATETIME("u"."createat", '-2 DAY')
```

#### Day
Извлекает день месяца (1–31) из выражения даты/времени.
```go
function := uast.Day(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
DAY(`u`.`createat`)
```
**Output MsSQL:**
```sql
DAY([u].[createat])
```
**Output MySQL:**
```sql
DAY(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(DAY FROM "u"."createat")
```
**Output SQLite:**
```sql
DAY("u"."createat")
```

#### DayName
Возвращает название дня недели (например, 'Понедельник', 'Вторник') для заданного выражения даты/времени.
```go
function := uast.DayName(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
DAYNAME(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATENAME(WEEKDAY, [u].[createat])
```
**Output MySQL:**
```sql
DAYNAME(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
TO_CHAR("u"."createat", 'Day')
```
**Output SQLite:**
```sql
STRFTIME('%w', "u"."createat")
```

#### Hour
Извлекает час (0–23) из выражения даты/времени.
```go
function := uast.Hour(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
HOUR(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATEPART(HOUR, [u].[createat])
```
**Output MySQL:**
```sql
HOUR(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(HOUR FROM "u"."createat")
```
**Output SQLite:**
```sql
HOUR("u"."createat")
```

#### Minute
Извлекает минуту (0–59) из выражения даты/времени.
```go
function := uast.Minute(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
MINUTE(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATEPART(MINUTE, [u].[createat])
```
**Output MySQL:**
```sql
MINUTE(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(MINUTE FROM "u"."createat")
```
**Output SQLite:**
```sql
MINUTE("u"."createat")
```

#### Month
Извлекает месяц (1–12) из выражения даты/времени.
```go
function := uast.Month(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
MONTH(`u`.`createat`)
```
**Output MsSQL:**
```sql
MONTH([u].[createat])
```
**Output MySQL:**
```sql
MONTH(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(MONTH FROM "u"."createat")
```
**Output SQLite:**
```sql
MONTH("u"."createat")
```

#### MonthName
Возвращает название месяца (например, 'Январь', 'Февраль') для заданного выражения даты/времени.
```go
function := uast.MonthName(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
MONTHNAME(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATENAME(MONTH, [u].[createat])
```
**Output MySQL:**
```sql
MONTHNAME(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
TO_CHAR("u"."createat", 'Month')
```
**Output SQLite:**
```sql
STRFTIME('%m', "u"."createat")
```

#### Now
Возвращает текущую дату и время.
```go
function := uast.Now()
```
**Output MariaDB:**
```sql
NOW()
```
**Output MsSQL:**
```sql
GETDATE()
```
**Output MySQL:**
```sql
NOW()
```
**Output PostgreSQL:**
```sql
CURRENT_TIMESTAMP
```
**Output SQLite:**
```sql
DATETIME('now')
```

#### Quarter
Извлекает квартал (1–4) из выражения даты/времени.
```go
function := uast.Quarter(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
QUARTER(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATEPART(QUARTER, [u].[createat])
```
**Output MySQL:**
```sql
QUARTER(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(QUARTER FROM "u"."createat")
```
**Output SQLite:**
```sql
QUARTER("u"."createat")
```

#### Second
Извлекает секунду (0–59) из выражения даты/времени.
```go
function := uast.Second(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
SECOND(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATEPART(SECOND, [u].[createat])
```
**Output MySQL:**
```sql
SECOND(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(SECOND FROM "u"."createat")
```
**Output SQLite:**
```sql
SECOND("u"."createat")
```

#### TimeAdd
Добавляет интервал времени к выражению времени/даты/времени и возвращает результирующее время.
```go
function := uast.TimeAdd(uast.Field[time.Time]("u", "createat"), uast.Value("2 HOUR"))
```
**Output MariaDB:**
```sql
TIME_ADD(`u`.`createat`, '2 HOUR')
```
**Output MsSQL:**
```sql
DATEADD(HOUR, 2, [u].[createat])
```
**Output MySQL:**
```sql
TIME_ADD(`u`.`createat`, '2 HOUR')
```
**Output PostgreSQL:**
```sql
("u"."createat" + INTERVAL '2 HOUR')
```
**Output SQLite:**
```sql
TIME("u"."createat", '+2 HOUR')
```

#### TimeDiff
Возвращает разницу между двумя выражениями времени/даты/времени (`timeEnd` - `timeStart`).
```go
function := uast.TimeDiff(uast.Field[time.Time]("u", "updateat"), uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
TIMEDIFF(`u`.`updateat`, `u`.`createat`)
```
**Output MsSQL:**
```sql
TIMEDIFF([u].[updateat], [u].[createat])
```
**Output MySQL:**
```sql
TIMEDIFF(`u`.`updateat`, `u`.`createat`)
```
**Output PostgreSQL:**
```sql
DATE_PART('time', "u"."updateat" - "u"."createat")
```
**Output SQLite:**
```sql
TIMEDIFF("u"."updateat", "u"."createat")
```

#### TimeSub
Вычитает интервал времени из выражения времени/даты/времени и возвращает результирующее время.
```go
function := uast.TimeSub(uast.Field[time.Time]("u", "createat"), uast.Value("2 HOUR"))
```
**Output MariaDB:**
```sql
TIME_SUB(`u`.`createat`, '2 HOUR')
```
**Output MsSQL:**
```sql
DATEADD(HOUR, -2, [u].[createat])
```
**Output MySQL:**
```sql
TIME_SUB(`u`.`createat`, '2 HOUR')
```
**Output PostgreSQL:**
```sql
("u"."createat" - INTERVAL '2 HOUR')
```
**Output SQLite:**
```sql
TIME("u"."createat", '-2 HOUR')
```

#### Week
Извлекает номер недели (1–53) из выражения даты/времени.
```go
function := uast.Week(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
WEEK(`u`.`createat`)
```
**Output MsSQL:**
```sql
DATEPART(WEEK, [u].[createat])
```
**Output MySQL:**
```sql
WEEK(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(WEEK FROM "u"."createat")
```
**Output SQLite:**
```sql
WEEK("u"."createat")
```

#### Year
Извлекает год из выражения даты/времени.
```go
function := uast.Year(uast.Field[time.Time]("u", "createat"))
```
**Output MariaDB:**
```sql
YEAR(`u`.`createat`)
```
**Output MsSQL:**
```sql
YEAR([u].[createat])
```
**Output MySQL:**
```sql
YEAR(`u`.`createat`)
```
**Output PostgreSQL:**
```sql
EXTRACT(YEAR FROM "u"."createat")
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_ARRAY(`u`.`json`, ?, ?)
```
**Output MsSQL:**
```sql
JSON_ARRAY([u].[json], @p1, @p2)
```
**Output MySQL:**
```sql
JSON_ARRAY(`u`.`json`, ?, ?)
```
**Output PostgreSQL:**
```sql
JSON_ARRAY("u"."json", $1, $2)
```
**Output SQLite:**
```sql
JSON_ARRAY("u"."json", ?, ?)
```

#### JsonArrayAgg
Агрегирует значения из группы в JSON-массив.
```go
function := uast.JsonArrayAgg(
    uast.Field[string]("u", "json"),
)
```
**Output MariaDB:**
```sql
JSON_ARRAYAGG(`u`.`json`)
```
**Output MsSQL:**
```sql
JSON_ARRAYAGG([u].[json])
```
**Output MySQL:**
```sql
JSON_ARRAYAGG(`u`.`json`)
```
**Output PostgreSQL:**
```sql
JSON_AGG("u"."json")
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_CONTAINS(`u`.`json`, '{"key":"val"}')
```
**Output MsSQL:**
```sql
// Not supported
```
**Output MySQL:**
```sql
JSON_CONTAINS(`u`.`json`, '{"key":"val"}')
```
**Output PostgreSQL:**
```sql
("u"."json" @> '{"key":"val"}')
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
(`u`.`json` ->> '$.parent[0].child')
```
**Output MsSQL:**
```sql
JSON_VALUE([u].[json], '$.parent[0].child')
```
**Output MySQL:**
```sql
(`u`.`json` ->> '$.parent[0].child')
```
**Output PostgreSQL:**
```sql
("u"."json" #>> '{parent,0,child}')
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_OBJECT('key', COUNT(`u`.`json`))
```
**Output MsSQL:**
```sql
JSON_OBJECT('key', COUNT([u].[json]))
```
**Output MySQL:**
```sql
JSON_OBJECT('key', COUNT(`u`.`json`))
```
**Output PostgreSQL:**
```sql
JSON_BUILD_OBJECT('key', COUNT("u"."json"))
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_OBJECTAGG(`u`.`json`, `u`.`number`)
```
**Output MsSQL:**
```sql
JSON_OBJECTAGG([u].[json], [u].[number])
```
**Output MySQL:**
```sql
JSON_OBJECTAGG(`u`.`json`, `u`.`number`)
```
**Output PostgreSQL:**
```sql
JSON_OBJECT_AGG("u"."json", "u"."number")
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_REMOVE(`u`.`json`, '$.key1', '$.key2')
```
**Output MsSQL:**
```sql
JSON_MODIFY(JSON_MODIFY([u].[json], '$.key1', NULL), '$.key2', NULL)
```
**Output MySQL:**
```sql
JSON_REMOVE(`u`.`json`, '$.key1', '$.key2')
```
**Output PostgreSQL:**
```sql
("u"."json" - '{key1}' - '{key2}')
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
JSON_SET(`u`.`json`, '$.key1', ?, '$.key2', ?)
```
**Output MsSQL:**
```sql
JSON_MODIFY(JSON_MODIFY([u].[json], '$.key1', @p1), '$.key2', @p2)
```
**Output MySQL:**
```sql
JSON_SET(`u`.`json`, '$.key1', ?, '$.key2', ?)
```
**Output PostgreSQL:**
```sql
jsonb_set(jsonb_set("u"."json", '{key1}', $1), '{key2}', $2)
```
**Output SQLite:**
```sql
JSON_SET("u"."json", '$.key1', ?, '$.key2', ?)
```

#### JsonType
Возвращает тип JSON-значения (например, 'OBJECT', 'ARRAY', 'STRING', 'INTEGER', 'NULL').
```go
function := uast.JsonType(uast.Field[string]("u", "json"))
```
**Output MariaDB:**
```sql
JSON_TYPE(`u`.`json`)
```
**Output MsSQL:**
```sql
// Not supported
```
**Output MySQL:**
```sql
JSON_TYPE(`u`.`json`)
```
**Output PostgreSQL:**
```sql
jsonb_typeof("u"."json")
```
**Output SQLite:**
```sql
JSON_TYPE("u"."json")
```

### Math
#### Abs
Возвращает абсолютное (неотрицательное) значение числового выражения.
```go
function := uast.Abs(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
ABS(`u`.`x`)
```
**Output MsSQL:**
```sql
ABS([u].[x])
```
**Output MySQL:**
```sql
ABS(`u`.`x`)
```
**Output PostgreSQL:**
```sql
ABS("u"."x")
```
**Output SQLite:**
```sql
ABS("u"."x")
```

#### ACos
Возвращает арккосинус (обратный косинус) выражения в радианах.
```go
function := uast.ACos(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
ACOS(`u`.`x`)
```
**Output MsSQL:**
```sql
ACOS([u].[x])
```
**Output MySQL:**
```sql
ACOS(`u`.`x`)
```
**Output PostgreSQL:**
```sql
ACOS("u"."x")
```
**Output SQLite:**
```sql
ACOS("u"."x")
```

#### ASin
Возвращает арксинус (обратный синус) выражения в радианах.
```go
function := uast.ASin(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
ASIN(`u`.`x`)
```
**Output MsSQL:**
```sql
ASIN([u].[x])
```
**Output MySQL:**
```sql
ASIN(`u`.`x`)
```
**Output PostgreSQL:**
```sql
ASIN("u"."x")
```
**Output SQLite:**
```sql
ASIN("u"."x")
```

#### ATan
Возвращает арктангенс (обратный тангенс) выражения в радианах.
```go
function := uast.ATan(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
ATAN(`u`.`x`)
```
**Output MsSQL:**
```sql
ATAN([u].[x])
```
**Output MySQL:**
```sql
ATAN(`u`.`x`)
```
**Output PostgreSQL:**
```sql
ATAN("u"."x")
```
**Output SQLite:**
```sql
ATAN("u"."x")
```

#### ATan2
Возвращает арктангенс частного двух аргументов (`y`/`x`), используя их знаки для определения квадранта.
```go
function := uast.ATan2(uast.Field[int]("u", "y"), uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
ATAN2(`u`.`y`, `u`.`x`)
```
**Output MsSQL:**
```sql
ATAN2([u].[y], [u].[x])
```
**Output MySQL:**
```sql
ATAN2(`u`.`y`, `u`.`x`)
```
**Output PostgreSQL:**
```sql
ATAN2("u"."y", "u"."x")
```
**Output SQLite:**
```sql
ATAN2("u"."y", "u"."x")
```

#### Cbrt
Возвращает кубический корень числового выражения.
```go
function := uast.Cbrt(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
CBRT(`u`.`x`)
```
**Output MsSQL:**
```sql
CBRT([u].[x])
```
**Output MySQL:**
```sql
CBRT(`u`.`x`)
```
**Output PostgreSQL:**
```sql
CBRT("u"."x")
```
**Output SQLite:**
```sql
CBRT("u"."x")
```

#### Ceil
Возвращает наименьшее целое значение, не меньшее аргумента (округление вверх).
```go
function := uast.Ceil(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
CEILING(`u`.`x`)
```
**Output MsSQL:**
```sql
CEILING([u].[x])
```
**Output MySQL:**
```sql
CEILING(`u`.`x`)
```
**Output PostgreSQL:**
```sql
CEIL("u"."x")
```
**Output SQLite:**
```sql
CEIL("u"."x")
```

#### Cos
Возвращает косинус выражения в радианах.
```go
function := uast.Cos(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
COS(`u`.`x`)
```
**Output MsSQL:**
```sql
COS([u].[x])
```
**Output MySQL:**
```sql
COS(`u`.`x`)
```
**Output PostgreSQL:**
```sql
COS("u"."x")
```
**Output SQLite:**
```sql
COS("u"."x")
```

#### Exp
Возвращает число Эйлера `e` (~2.71828) возведённое в степень выражения.
```go
function := uast.Exp(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
EXP(`u`.`x`)
```
**Output MsSQL:**
```sql
EXP([u].[x])
```
**Output MySQL:**
```sql
EXP(`u`.`x`)
```
**Output PostgreSQL:**
```sql
EXP("u"."x")
```
**Output SQLite:**
```sql
EXP("u"."x")
```

#### Floor
Возвращает наибольшее целое значение, не большее аргумента (округление вниз).
```go
function := uast.Floor(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
FLOOR(`u`.`x`)
```
**Output MsSQL:**
```sql
FLOOR([u].[x])
```
**Output MySQL:**
```sql
FLOOR(`u`.`x`)
```
**Output PostgreSQL:**
```sql
FLOOR("u"."x")
```
**Output SQLite:**
```sql
FLOOR("u"."x")
```

#### Ln
Возвращает натуральный логарифм (по основанию `e`) выражения.
```go
function := uast.Ln(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
LN(`u`.`x`)
```
**Output MsSQL:**
```sql
LN([u].[x])
```
**Output MySQL:**
```sql
LN(`u`.`x`)
```
**Output PostgreSQL:**
```sql
LN("u"."x")
```
**Output SQLite:**
```sql
LN("u"."x")
```

#### Log
Возвращает логарифм выражения по указанному основанию.
```go
function := uast.Log(uast.Field[int]("u", "x"), uast.Value(2))
```
**Output MariaDB:**
```sql
LOG(`u`.`x`, ?)
```
**Output MsSQL:**
```sql
LOG([u].[x], @p1)
```
**Output MySQL:**
```sql
LOG(`u`.`x`, ?)
```
**Output PostgreSQL:**
```sql
LOG("u"."x", $1)
```
**Output SQLite:**
```sql
LOG("u"."x", ?)
```

#### Mod
Возвращает остаток (модуль) от деления первого выражения на второе.
```go
function := uast.Mod(uast.Field[int]("u", "x"), uast.Value(2))
```
**Output MariaDB:**
```sql
MOD(`u`.`x`, ?)
```
**Output MsSQL:**
```sql
MOD([u].[x], @p1)
```
**Output MySQL:**
```sql
MOD(`u`.`x`, ?)
```
**Output PostgreSQL:**
```sql
MOD("u"."x", $1)
```
**Output SQLite:**
```sql
MOD("u"."x", ?)
```

#### Pi
Возвращает математическую константу `π` (~3.14159).
```go
function := uast.Pi()
```
**Output MariaDB:**
```sql
PI()
```
**Output MsSQL:**
```sql
PI()
```
**Output MySQL:**
```sql
PI()
```
**Output PostgreSQL:**
```sql
PI()
```
**Output SQLite:**
```sql
PI()
```

#### Power
Возвращает выражение, возведённое в степень экспоненты.
```go
function := uast.Power(uast.Field[int]("u", "x"), uast.Value(2))
```
**Output MariaDB:**
```sql
POWER(`u`.`x`, ?)
```
**Output MsSQL:**
```sql
POWER([u].[x], @p1)
```
**Output MySQL:**
```sql
POWER(`u`.`x`, ?)
```
**Output PostgreSQL:**
```sql
POWER("u"."x", $1)
```
**Output SQLite:**
```sql
POWER("u"."x", ?)
```

#### Rand
Возвращает случайное значение с плавающей запятой в диапазоне [0, 1].
```go
function := uast.Rand()
```
**Output MariaDB:**
```sql
RAND()
```
**Output MsSQL:**
```sql
RAND()
```
**Output MySQL:**
```sql
RAND()
```
**Output PostgreSQL:**
```sql
RANDOM()
```
**Output SQLite:**
```sql
RANDOM()
```

#### Round
Округляет выражение до указанного количества знаков после запятой.
```go
function := uast.Round(uast.Field[int]("u", "x"), uast.Value(2))
```
**Output MariaDB:**
```sql
ROUND(`u`.`x`, ?)
```
**Output MsSQL:**
```sql
ROUND([u].[x], @p1)
```
**Output MySQL:**
```sql
ROUND(`u`.`x`, ?)
```
**Output PostgreSQL:**
```sql
ROUND("u"."x", $1)
```
**Output SQLite:**
```sql
ROUND("u"."x", ?)
```

#### Sin
Возвращает синус выражения в радианах.
```go
function := uast.Sin(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
SIN(`u`.`x`)
```
**Output MsSQL:**
```sql
SIN([u].[x])
```
**Output MySQL:**
```sql
SIN(`u`.`x`)
```
**Output PostgreSQL:**
```sql
SIN("u"."x")
```
**Output SQLite:**
```sql
SIN("u"."x")
```

#### Sqrt
Возвращает квадратный корень выражения.
```go
function := uast.Sqrt(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
SQRT(`u`.`x`)
```
**Output MsSQL:**
```sql
SQRT([u].[x])
```
**Output MySQL:**
```sql
SQRT(`u`.`x`)
```
**Output PostgreSQL:**
```sql
SQRT("u"."x")
```
**Output SQLite:**
```sql
SQRT("u"."x")
```

#### Tan
Возвращает тангенс выражения в радианах.
```go
function := uast.Tan(uast.Field[int]("u", "x"))
```
**Output MariaDB:**
```sql
TAN(`u`.`x`)
```
**Output MsSQL:**
```sql
TAN([u].[x])
```
**Output MySQL:**
```sql
TAN(`u`.`x`)
```
**Output PostgreSQL:**
```sql
TAN("u"."x")
```
**Output SQLite:**
```sql
TAN("u"."x")
```

#### Trunc
Усекает числовое выражение до указанного количества знаков после запятой (без округления).
```go
function := uast.Trunc(uast.Field[int]("u", "x"), uast.Value(2))
```
**Output MariaDB:**
```sql
TRUNCATE(`u`.`x`, ?)
```
**Output MsSQL:**
```sql
ROUND([u].[x], @p1, 1)
```
**Output MySQL:**
```sql
TRUNCATE(`u`.`x`, ?)
```
**Output PostgreSQL:**
```sql
TRUNC("u"."x", $1)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
CUME_DIST() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
CUME_DIST() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
CUME_DIST() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
CUME_DIST() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
DENSE_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
DENSE_RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
DENSE_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
DENSE_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
NTILE(2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
NTILE(2) OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
NTILE(2) OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
NTILE(2) OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
PERCENT_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
PERCENT_RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
PERCENT_RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
PERCENT_RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
RANK() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
RANK() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
RANK() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
ROW_NUMBER() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output MsSQL:**
```sql
ROW_NUMBER() OVER (PARTITION BY [u].[id] ORDER BY [u].[number] DESC)
```
**Output MySQL:**
```sql
ROW_NUMBER() OVER (PARTITION BY `u`.`id` ORDER BY `u`.`number` DESC)
```
**Output PostgreSQL:**
```sql
ROW_NUMBER() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```
**Output SQLite:**
```sql
ROW_NUMBER() OVER (PARTITION BY "u"."id" ORDER BY "u"."number" DESC)
```

### String
#### Concat
Объединяет два или более строковых выражения в одну строку. Аргументы `NULL` рассматриваются как пустые строки в большинстве диалектов.
```go
function := uast.Concat(uast.Field[string]("u", "string"), uast.Value("old"), uast.Value("new"))
```
**Output MariaDB:**
```sql
CONCAT(`u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
CONCAT([u].[string], @p1, @p2)
```
**Output MySQL:**
```sql
CONCAT(`u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
CONCAT("u"."string", $1, $2)
```
**Output SQLite:**
```sql
CONCAT("u"."string", ?, ?)
```

#### ConcatWs
Объединяет два или более строковых выражения с указанным разделителем между ними. Пропускает аргументы `NULL`.
```go
function := uast.ConcatWs(uast.Value("_"), uast.Field[string]("u", "string"), uast.Value("old"),uast.Value("new"))
```
**Output MariaDB:**
```sql
CONCAT_WS(?, `u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
CONCAT_WS(@p1, [u].[string], @p2, @p3)
```
**Output MySQL:**
```sql
CONCAT_WS(?, `u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
CONCAT_WS($1, "u"."string", $2, $3)
```
**Output SQLite:**
```sql
CONCAT_WS(?, "u"."string", ?, ?)
```

#### LeftString
Возвращает крайние слева `count` символов из строкового выражения.
```go
function := uast.LeftString(uast.Field[string]("u", "string"), uast.Value(2))
```
**Output MariaDB:**
```sql
LEFT(`u`.`string`, ?)
```
**Output MsSQL:**
```sql
LEFT([u].[string], @p1)
```
**Output MySQL:**
```sql
LEFT(`u`.`string`, ?)
```
**Output PostgreSQL:**
```sql
LEFT("u"."string", $1)
```
**Output SQLite:**
```sql
LEFT("u"."string", ?)
```

#### Lower
Преобразует строковое выражение в нижний регистр.
```go
function := uast.Lower(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
LOWER(`u`.`string`)
```
**Output MsSQL:**
```sql
LOWER([u].[string])
```
**Output MySQL:**
```sql
LOWER(`u`.`string`)
```
**Output PostgreSQL:**
```sql
LOWER("u"."string")
```
**Output SQLite:**
```sql
LOWER("u"."string")
```

#### LPad
Дополняет строковое выражение слева указанным разделителем до общей длины `count` символов.
```go
function := uast.LPad(uast.Field[string]("u", "string"), uast.Value(2), uast.Value(","))
```
**Output MariaDB:**
```sql
LPAD(`u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
LPAD([u].[string], @p1, @p2)
```
**Output MySQL:**
```sql
LPAD(`u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
LPAD("u"."string", $1, $2)
```
**Output SQLite:**
```sql
LPAD("u"."string", ?, ?)
```

#### LTrim
Удаляет начальные пробелы из строкового выражения.
```go
function := uast.LTrim(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
LTRIM(`u`.`string`)
```
**Output MsSQL:**
```sql
LTRIM([u].[string])
```
**Output MySQL:**
```sql
LTRIM(`u`.`string`)
```
**Output PostgreSQL:**
```sql
LTRIM("u"."string")
```
**Output SQLite:**
```sql
LTRIM("u"."string")
```

#### Repeat
Повторяет строковое выражение `count` раз.
```go
function := uast.Repeat(uast.Field[string]("u", "string"), uast.Value(2))
```
**Output MariaDB:**
```sql
REPEAT(`u`.`string`, ?)
```
**Output MsSQL:**
```sql
REPEAT([u].[string], @p1)
```
**Output MySQL:**
```sql
REPEAT(`u`.`string`, ?)
```
**Output PostgreSQL:**
```sql
REPEAT("u"."string", $1)
```
**Output SQLite:**
```sql
REPEAT("u"."string", ?)
```

#### Replace
Заменяет все вхождения подстроки в строке на новую подстроку.
```go
function := uast.Replace(uast.Field[string]("u", "string"), uast.Value("old"), uast.Value("new"))
```
**Output MariaDB:**
```sql
REPLACE(`u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
REPLACE([u].[string], @p1, @p2)
```
**Output MySQL:**
```sql
REPLACE(`u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
REPLACE("u"."string", $1, $2)
```
**Output SQLite:**
```sql
REPLACE("u"."string", ?, ?)
```

#### Reverse
Переворачивает символы в строковом выражении.
```go
function := uast.Reverse(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
REVERSE(`u`.`string`)
```
**Output MsSQL:**
```sql
REVERSE([u].[string])
```
**Output MySQL:**
```sql
REVERSE(`u`.`string`)
```
**Output PostgreSQL:**
```sql
REVERSE("u"."string")
```
**Output SQLite:**
```sql
REVERSE("u"."string")
```

#### RightString
Возвращает крайние справа `count` символов из строкового выражения.
```go
function := uast.RightString(uast.Field[string]("u", "string"), uast.Value(2))
```
**Output MariaDB:**
```sql
RIGHT(`u`.`string`, ?)
```
**Output MsSQL:**
```sql
RIGHT([u].[string], @p1)
```
**Output MySQL:**
```sql
RIGHT(`u`.`string`, ?)
```
**Output PostgreSQL:**
```sql
RIGHT("u"."string", $1)
```
**Output SQLite:**
```sql
RIGHT("u"."string", ?)
```

#### RPad
Дополняет строковое выражение справа указанным разделителем до общей длины `count` символов.
```go
function := uast.RPad(uast.Field[string]("u", "string"), uast.Value(2), uast.Value(","))
```
**Output MariaDB:**
```sql
RPAD(`u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
RPAD([u].[string], @p1, @p2)
```
**Output MySQL:**
```sql
RPAD(`u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
RPAD("u"."string", $1, $2)
```
**Output SQLite:**
```sql
RPAD("u"."string", ?, ?)
```

#### RTrim
Удаляет конечные пробелы из строкового выражения.
```go
function := uast.RTrim(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
RTRIM(`u`.`string`)
```
**Output MsSQL:**
```sql
RTRIM([u].[string])
```
**Output MySQL:**
```sql
RTRIM(`u`.`string`)
```
**Output PostgreSQL:**
```sql
RTRIM("u"."string")
```
**Output SQLite:**
```sql
RTRIM("u"."string")
```

#### SubString
Извлекает подстроку из строкового выражения, начиная с `startPos` (начиная с 1) длиной `lengthStr` символов.
```go
function := uast.SubString(uast.Field[string]("u", "string"), uast.Value(0), uast.Value(2))
```
**Output MariaDB:**
```sql
SUBSTRING(`u`.`string`, ?, ?)
```
**Output MsSQL:**
```sql
SUBSTRING([u].[string], @p1, @p2)
```
**Output MySQL:**
```sql
SUBSTRING(`u`.`string`, ?, ?)
```
**Output PostgreSQL:**
```sql
SUBSTRING("u"."string", $1, $2)
```
**Output SQLite:**
```sql
SUBSTRING("u"."string", ?, ?)
```

#### Trim
Удаляет как начальные, так и конечные пробелы из строкового выражения.
```go
function := uast.Trim(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
TRIM(`u`.`string`)
```
**Output MsSQL:**
```sql
TRIM([u].[string])
```
**Output MySQL:**
```sql
TRIM(`u`.`string`)
```
**Output PostgreSQL:**
```sql
TRIM("u"."string")
```
**Output SQLite:**
```sql
TRIM("u"."string")
```

#### Upper
Преобразует строковое выражение в верхний регистр.
```go
function := uast.Upper(uast.Field[string]("u", "string"))
```
**Output MariaDB:**
```sql
UPPER(`u`.`string`)
```
**Output MsSQL:**
```sql
UPPER([u].[string])
```
**Output MySQL:**
```sql
UPPER(`u`.`string`)
```
**Output PostgreSQL:**
```sql
UPPER("u"."string")
```
**Output SQLite:**
```sql
UPPER("u"."string")
```

## exprLiteral
### Literal
Встраивает необработанное литеральное значение непосредственно в сгенерированную SQL-строку (не параметризуется). Используйте с осторожностью — значения записываются как есть. Предпочитайте `Value` для пользовательских данных.
```go
literal := uast.Literal("%Y-%m-%d")
```
**Output SQL:**
```sql
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
**Output MariaDB:**
```sql
(`u`.`string` = ? AND `u`.`number` > ?)
```
**Output MsSQL:**
```sql
([u].[string] = @p1 AND [u].[number] > @p2)
```
**Output MySQL:**
```sql
(`u`.`string` = ? AND `u`.`number` > ?)
```
**Output PostgreSQL:**
```sql
("u"."string" = $1 AND "u"."number" > $2)
```
**Output SQLite:**
```sql
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
**Output MariaDB:**
```sql
(`u`.`string` = ? OR `u`.`number` > ?)
```
**Output MsSQL:**
```sql
([u].[string] = @p1 OR [u].[number] > @p2)
```
**Output MySQL:**
```sql
(`u`.`string` = ? OR `u`.`number` > ?)
```
**Output PostgreSQL:**
```sql
("u"."string" = $1 OR "u"."number" > $2)
```
**Output SQLite:**
```sql
("u"."string" = ? OR "u"."number" > ?)
```

## exprSubquery
### Subquery
Оборачивает оператор `SELECT` как типизированное выражение, которое может использоваться в сравнениях (`In`, `Exists`, `Equal` и т.д.) или как колонка в операторе `SELECT`. Обобщённый параметр `u` указывает скалярный тип единственной колонки, возвращаемой подзапросом.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Field[int64]("u", "id")).From(uast.NewTable("users").As("u")))
```
**Output MariaDB:**
```sql
(SELECT `u`.`id` FROM `users` AS `u`)
```
**Output MsSQL:**
```sql
(SELECT [u].[id] FROM [users] AS [u])
```
**Output MySQL:**
```sql
(SELECT `u`.`id` FROM `users` AS `u`)
```
**Output PostgreSQL:**
```sql
(SELECT "u"."id" FROM "users" AS "u")
```
**Output SQLite:**
```sql
(SELECT "u"."id" FROM "users" AS "u")
```

## exprValue
### Value
Оборачивает Go-значение как параметризованное выражение. Значение НЕ вставляется в SQL-строку напрямую — вместо этого генерируется плейсхолдер (`?`, `$1`, и т.д.), а значение добавляется в слайс аргументов, возвращаемый `Build()`. Это безопасный способ передачи пользовательских данных, предотвращающий SQL-инъекции. Поддерживаемые типы: `bool`, `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `string`, `time.Time`.
```go
var data string = "ivan"
value := uast.Value(data)
```
**Output MariaDB:**
```sql
?
```
**Output MsSQL:**
```sql
@p1
```
**Output MySQL:**
```sql
?
```
**Output PostgreSQL:**
```sql
$1
```
**Output SQLite:**
```sql
?
```