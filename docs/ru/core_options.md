---
outline: deep
---

# API / Ядро / Опции

::: info **Информация**
Эта страница охватывает все параметры конфигурации: `exprArray`, `exprBinary`, `exprComparison`, `exprConstant`, `exprFunction`, `exprLiteral`, `exprLogical`, `exprOrderBy`, `exprSubquery`, `exprValue`. Каждый параметр показан с рабочим примером кода и ожидаемым выводом.
:::

## exprArray
### Array
Создаёт выражение массива для использования в SQL-запросах.
```go
array := uast.Array(0, 1, 2)
```
Output MySQL:
```text
ARRAY[?, ?, ?]
```
Output PostgreSQL:
```text
ARRAY[$1, $2, $3]
```

## exprBinary
### BitwiseAnd
Выполняет побитовую операцию И между двумя выражениями.
```go
binary := uast.BitwiseAnd(uast.Column[int]("t", "number"), uast.Value(0b0010))
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
Output MySQL:
```text
`t`.`number` | ?
```
Output PostgreSQL:
```text
"t"."number" | $1
```

### BitwiseXor
Выполняет побитовую операцию исключающего ИЛИ между двумя выражениями.
```go
binary := uast.BitwiseXor(uast.Column[int]("t", "number"), uast.Value(0b0010))
```
Output MySQL:
```text
`t`.`number` ^ ?
```
Output PostgreSQL:
```text
"t"."number" ^ $1
```

### Divide
Делит левое выражение на правое.
```go
binary := uast.Divide(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` / ?
```
Output PostgreSQL:
```text
"t"."number" / $1
```

### Minus
Вычитает правое выражение из левого.
```go
binary := uast.Minus(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` - ?
```
Output PostgreSQL:
```text
"t"."number" - $1
```

### Modulo
Возвращает остаток от деления левого выражения на правое.
```go
binary := uast.Modulo(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` % ?
```
Output PostgreSQL:
```text
"t"."number" % $1
```

### Multiply
Умножает левое выражение на правое.
```go
binary := uast.Multiply(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` * ?
```
Output PostgreSQL:
```text
"t"."number" * $1
```

### Plus
Складывает левое выражение с правым.
```go
binary := uast.Plus(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` + ?
```
Output PostgreSQL:
```text
"t"."number" + $1
```

### ShiftLeft
Выполняет побитовый сдвиг влево левого выражения на количество бит, указанное в правом выражении.
```go
binary := uast.ShiftLeft(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` << ?
```
Output PostgreSQL:
```text
"t"."number" << $1
```

### ShiftRight
Выполняет побитовый сдвиг вправо левого выражения на количество бит, указанное в правом выражении.
```go
binary := uast.ShiftRight(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` >> ?
```
Output PostgreSQL:
```text
"t"."number" >> $1
```

## exprColumn
### Column
Создаёт ссылку на колонку таблицы, опционально квалифицированную псевдонимом таблицы. Это основной способ ссылаться на колонки базы данных в выражениях.
```go
column := uast.Column[string]("t", "string")
```
Output MySQL:
```text
`t`.`string`
```
Output PostgreSQL:
```text
"t"."string"
```

## exprComparison
### Between
Проверяет, попадает ли левое выражение в диапазон, заданный valueStart и valueEnd (включительно).
```go
comparison := uast.Between(uast.Column[int]("t", "number"), uast.Value(0), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` BETWEEN ? AND ?
```
Output PostgreSQL:
```text
"t"."number" BETWEEN $1 AND $2
```

### Equal
Сравнивает два выражения на равенство (`=`).
```go
comparison := uast.Equal(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` = ?
```
Output PostgreSQL:
```text
"t"."number" = $1
```

### Exists
Проверяет, возвращает ли подзапрос какие-либо строки. Возвращает `true` если существует хотя бы одна строка.
```go
comparison := uast.Exists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.Table("test"))))
```
Output MySQL:
```text
EXISTS (SELECT 1 FROM `test` AS `t`)
```
Output PostgreSQL:
```text
EXISTS (SELECT 1 FROM "test" AS "t")
```

### Greater
Сравнивает, больше ли левое выражение правого (`>`).
```go
comparison := uast.Greater(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` > ?
```
Output PostgreSQL:
```text
"t"."number" > $1
```

### GreaterEqual
Сравнивает, больше или равно ли левое выражение правому (`>=`).
```go
comparison := uast.GreaterEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` >= ?
```
Output PostgreSQL:
```text
"t"."number" >= $1
```

### ILike
Выполняет регистронезависимое сравнение с шаблоном. Правое выражение должно содержать шаблон с `%` (любая последовательность) и `_` (один символ).
```go
comparison := uast.ILike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MySQL:
```text
`t`.`string` ILIKE ?
```
Output PostgreSQL:
```text
"t"."string" ILIKE $1
```

### In
Проверяет, соответствует ли левое выражение любому значению, содержащемуся в правом выражении (обычно подзапрос или массив).
```go
comparison := uast.In(uast.Column[string]("t", "string"), uast.Array("active", "pending"))
```
Output MySQL:
```text
`t`.`string` IN (?, ?)
```
Output PostgreSQL:
```text
"t"."string" IN ($1, $2)
```

### IsNotNull
Проверяет, что выражение не `NULL`.
```go
comparison := uast.IsNotNull(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
`t`.`string` IS NOT NULL
```
Output PostgreSQL:
```text
"t"."string" IS NOT NULL
```

### IsNull
Проверяет, что выражение является `NULL`.
```go
comparison := uast.IsNull(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
`t`.`string` IS NULL
```
Output PostgreSQL:
```text
"t"."string" IS NULL
```

### Less
Сравнивает, меньше ли левое выражение правого (`<`).
```go
comparison := uast.Less(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` < ?
```
Output PostgreSQL:
```text
"t"."number" < $1
```

### LessEqual
Сравнивает, меньше или равно ли левое выражение правому (`<=`).
```go
comparison := uast.LessEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` <= ?
```
Output PostgreSQL:
```text
"t"."number" <= $1
```

### Like
Выполняет регистрозависимое сравнение с шаблоном. Правое выражение должно содержать шаблон с `%` и `_`.
```go
comparison := uast.Like(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MySQL:
```text
`t`.`string` LIKE ?
```
Output PostgreSQL:
```text
"t"."string" LIKE $1
```

### NotBetween
Проверяет, находится ли левое выражение вне диапазона, заданного `valueStart` и `valueEnd`.
```go
comparison := uast.NotBetween(uast.Column[int]("t", "number"), uast.Value(0), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` NOT BETWEEN ? AND ?
```
Output PostgreSQL:
```text
"t"."number" NOT BETWEEN $1 AND $2
```

### NotEqual
Сравнивает два выражения на неравенство (`!=` or `<>`).
```go
comparison := uast.NotEqual(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
`t`.`number` != ?
```
Output PostgreSQL:
```text
"t"."number" != $1
```

### NotExists
Проверяет, что подзапрос не возвращает строк. Возвращает `true` если результат подзапроса пуст.
```go
comparison := uast.NotExists(uast.Subquery[int](uast.NewSelect(uast.ConstIntOne()).From(uast.Table("test"))))
```
Output MySQL:
```text
NOT EXISTS (SELECT 1 FROM `test` AS `t`)
```
Output PostgreSQL:
```text
NOT EXISTS (SELECT 1 FROM "test" AS "t")
```

### NotILike
Выполняет отрицательное регистронезависимое сравнение с шаблоном.
```go
comparison := uast.NotILike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MySQL:
```text
`t`.`string` NOT ILIKE ?
```
Output PostgreSQL:
```text
"t"."string" NOT ILIKE $1
```

### NotIn
Проверяет, что левое выражение не соответствует ни одному значению, содержащемуся в правом выражении.
```go
comparison := uast.NotIn(uast.Column[string]("t", "string"), uast.Array("active", "pending"))
```
Output MySQL:
```text
`t`.`string` NOT IN (?, ?)
```
Output PostgreSQL:
```text
"t"."string" NOT IN ($1, $2)
```

### NotLike
Выполняет отрицательное регистрозависимое сравнение с шаблоном.
```go
comparison := uast.NotLike(uast.Column[string]("t", "string"), uast.Value("%ivan%"))
```
Output MySQL:
```text
`t`.`string` NOT LIKE ?
```
Output PostgreSQL:
```text
"t"."string" NOT LIKE $1
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

#### BitAnd
Возвращает побитовое И всех битов в выражении. Имеет смысл только для целочисленных типов.
```go
function := uast.BitAnd(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitAnd(uast.Column[int]("t", "number"), true)
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

#### BitOr
Возвращает побитовое ИЛИ всех битов в выражении.
```go
function := uast.BitOr(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitOr(uast.Column[int]("t", "number"), true)
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

#### BitXor
Возвращает побитовое исключающее ИЛИ всех битов в выражении.
```go
function := uast.BitXor(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.BitXor(uast.Column[int]("t", "number"), true)
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

#### Count
Возвращает количество строк, соответствующих запросу, или количество не-NULL значений, если указано выражение. Когда `distinct` равен `true`, подсчитываются только уникальные значения.
```go
function := uast.Count(uast.Column[string]("t", "string"), false)
functionWithDistinct := uast.Count(uast.Column[string]("t", "string"), true)
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

#### GroupConcat
Объединяет значения из группы в одну строку, разделённую стандартным разделителем (обычно запятая). Флаг `distinct` удаляет дубликаты перед объединением.
```go
function := uast.GroupConcat(uast.Column[string]("t", "string"), false)
functionWithDistinct := uast.GroupConcat(uast.Column[string]("t", "string"), true)
```
Output MySQL:
```text
GROUP_CONCAT(`t`.`string` SEPARATOR ',')
GROUP_CONCAT(DISTINCT `t`.`string` SEPARATOR ',')
```
Output PostgreSQL:
```text
STRING_AGG("t"."string", ', ')
STRING_AGG(DISTINCT "t"."string", ', ')
```

#### Max
Возвращает максимальное значение выражения по всем строкам в группе.
```go
function := uast.Max(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Max(uast.Column[int]("t", "number"), true)
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

#### Min
Возвращает минимальное значение выражения по всем строкам в группе.
```go
function := uast.Min(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Min(uast.Column[int]("t", "number"), true)
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

#### StdDev
Возвращает популяционное стандартное отклонение выражения.
```go
function := uast.StdDev(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.StdDev(uast.Column[int]("t", "number"), true)
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

#### Sum
Возвращает сумму всех значений в выражении. Если `distinct` равен `true`, суммируются только уникальные значения.
```go
function := uast.Sum(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Sum(uast.Column[int]("t", "number"), true)
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

#### Variance
Возвращает популяционную дисперсию выражения.
```go
function := uast.Variance(uast.Column[int]("t", "number"), false)
functionWithDistinct := uast.Variance(uast.Column[int]("t", "number"), true)
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

### Analytical
#### FirstValue
Возвращает значение выражения из первой строки оконного фрейма. Требует оператор `OVER` с оконной спецификацией.
```go
function := uast.FirstValue(uast.Column[string]("t", "string")).Over(
    uast.PartitionBy(uast.Column[int64]("t", "id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "number"))),
)
```
Output MySQL:
```text
FIRST_VALUE(`t`.`string`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
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
Output MySQL:
```text
LAG(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)
```
Output PostgreSQL:
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
Output MySQL:
```text
LAST_VALUE(`t`.`string`) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```
Output PostgreSQL:
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
Output MySQL:
```text
LEAD(`t`.`number`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`date` ASC)
```
Output PostgreSQL:
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
Output MySQL:
```text
NTH_VALUE(`t`.`string`, 2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
```
Output PostgreSQL:
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
Output MySQL:
```text
CASE WHEN `t`.`number` < ? THEN ? ELSE ? END
```
Output PostgreSQL:
```text
CASE WHEN "t"."number" < $1 THEN $2 ELSE $3 END
```

#### Coalesce
Возвращает первое не-NULL выражение из предоставленного списка. Полезно для указания запасных значений.
```go
function := uast.Coalesce(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MySQL:
```text
COALESCE(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
COALESCE("t"."createat", "t"."updateat")
```

#### Greatest
Возвращает наибольшее значение из предоставленного списка выражений.
```go
function := uast.Greatest(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MySQL:
```text
GREATEST(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
GREATEST("t"."createat", "t"."updateat")
```

#### Least
Возвращает наименьшее значение из предоставленного списка выражений.
```go
function := uast.Least(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MySQL:
```text
LEAST(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
LEAST("t"."createat", "t"."updateat")
```

#### NullIf
Возвращает `NULL` если два выражения равны; иначе возвращает первое выражение.
```go
function := uast.NullIf(uast.Column[time.Time]("t", "createat"), uast.Column[time.Time]("t", "updateat"))
```
Output MySQL:
```text
NULLIF(`t`.`createat`, `t`.`updateat`)
```
Output PostgreSQL:
```text
NULLIF("t"."createat", "t"."updateat")
```

### Convert
#### Cast
Преобразует выражение к указанному типу данных.
```go
function := uast.Cast(uast.Column[int]("t", "number"), uast.TypeString)
```
Output MySQL:
```text
CAST(`t`.`number` AS CHAR)
```
Output PostgreSQL:
```text
CAST("t"."number" AS VARCHAR)
```

#### CharLength
Возвращает количество символов в строковом выражении.
```go
function := uast.CharLength(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
CHAR_LENGTH(`t`.`string`)
```
Output PostgreSQL:
```text
CHAR_LENGTH("t"."string")
```

#### DateFormat
Форматирует выражение даты/времени в соответствии с указанной маской формата.
```go
function := uast.DateFormat(uast.Column[time.Time]("t", "createat"), uast.Value("%Y-%m-%d"))
```
Output MySQL:
```text
DATE_FORMAT(`t`.`createat`, '%Y-%m-%d')
```
Output PostgreSQL:
```text
TO_CHAR("t"."createat", '%Y-%m-%d')
```

#### Degrees
Преобразует угол из радиан в градусы.
```go
function := uast.Degrees(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
DEGREES(`t`.`number`)
```
Output PostgreSQL:
```text
DEGREES("t"."number")
```

#### Length
Возвращает длину строкового выражения в байтах.
```go
function := uast.Length(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
LENGTH(`t`.`string`)
```
Output PostgreSQL:
```text
LENGTH("t"."string")
```

#### Position
Возвращает начальную позицию первого вхождения подстроки в строку.
```go
function := uast.Position(uast.Column[string]("t", "string"), uast.Value("old"))
```
Output MySQL:
```text
POSITION(? IN `t`.`string`)
```
Output PostgreSQL:
```text
POSITION($1 IN "t"."string")
```

#### Radians
Преобразует угол из градусов в радианы.
```go
function := uast.Radians(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
RADIANS(`t`.`number`)
```
Output PostgreSQL:
```text
RADIANS("t"."number")
```

### Date and time
#### CurDate
Возвращает текущую дату (без времени).
```go
function := uast.CurDate()
```
Output MySQL:
```text
CURDATE()
```
Output PostgreSQL:
```text
CURRENT_DATE
```

#### CurTime
Возвращает текущее время (без даты).
```go
function := uast.CurTime()
```
Output MySQL:
```text
CURTIME()
```
Output PostgreSQL:
```text
CURRENT_TIME
```

#### DateAdd
Добавляет интервал даты/времени к выражению даты/времени и возвращает результирующую дату/время.
```go
function := uast.DateAdd(uast.Column[time.Time]("t", "createat"), uast.Value("2 DAY"))
```
Output MySQL:
```text
DATE_ADD(`t`.`createat`, INTERVAL 2 DAY)
```
Output PostgreSQL:
```text
("t"."createat" + INTERVAL '2 DAY')
```

#### DateDiff
Возвращает разницу в днях между двумя выражениями даты/времени (`datetimeEnd` - `datetimeStart`).
```go
function := uast.DateDiff(uast.Column[time.Time]("t", "updateat"), uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
DATEDIFF(`t`.`updateat`, `t`.`createat`)
```
Output PostgreSQL:
```text
DATE_PART('day', "t"."updateat" - "t"."createat")
```

#### DateSub
Вычитает интервал даты/времени из выражения даты/времени и возвращает результирующую дату/время.
```go
function := uast.DateSub(uast.Column[time.Time]("t", "createat"), uast.Value("2 DAY"))
```
Output MySQL:
```text
DATE_SUB(`t`.`createat`, INTERVAL 2 DAY)
```
Output PostgreSQL:
```text
("t"."createat" - INTERVAL '2 DAY')
```

#### Day
Извлекает день месяца (1–31) из выражения даты/времени.
```go
function := uast.Day(uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
DAY(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(DAY FROM "t"."createat")
```

#### DayName
Возвращает название дня недели (например, 'Понедельник', 'Вторник') для заданного выражения даты/времени.
```go
function := uast.DayName(uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
DAYNAME(`t`.`createat`)
```
Output PostgreSQL:
```text
TO_CHAR("t"."createat", 'Day')
```

#### Hour
Извлекает час (0–23) из выражения даты/времени.
```go
function := uast.Hour(uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
HOUR(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(HOUR FROM "t"."createat")
```

#### Minute
Извлекает минуту (0–59) из выражения даты/времени.
```go
function := uast.Minute(uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
MINUTE(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(MINUTE FROM "t"."createat")
```

#### Month
Извлекает месяц (1–12) из выражения даты/времени.
```go
function := uast.Month(uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
MONTH(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(MONTH FROM "t"."createat")
```

#### MonthName
Возвращает название месяца (например, 'Январь', 'Февраль') для заданного выражения даты/времени.
```go
function := uast.MonthName(uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
MONTHNAME(`t`.`createat`)
```
Output PostgreSQL:
```text
TO_CHAR("t"."createat", 'Month')
```

#### Now
Возвращает текущую дату и время.
```go
function := uast.Now()
```
Output MySQL:
```text
NOW()
```
Output PostgreSQL:
```text
CURRENT_TIMESTAMP
```

#### Quarter
Извлекает квартал (1–4) из выражения даты/времени.
```go
function := uast.Quarter(uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
QUARTER(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(QUARTER FROM "t"."createat")
```

#### Second
Извлекает секунду (0–59) из выражения даты/времени.
```go
function := uast.Second(uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
SECOND(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(SECOND FROM "t"."createat")
```

#### TimeAdd
Добавляет интервал времени к выражению времени/даты/времени и возвращает результирующее время.
```go
function := uast.TimeAdd(uast.Column[time.Time]("t", "createat"), uast.Value("2 HOUR"))
```
Output MySQL:
```text
TIME_ADD(`t`.`createat`, '2 HOUR')
```
Output PostgreSQL:
```text
("t"."createat" + INTERVAL '2 HOUR')
```

#### TimeDiff
Возвращает разницу между двумя выражениями времени/даты/времени (`timeEnd` - `timeStart`).
```go
function := uast.TimeDiff(uast.Column[time.Time]("t", "updateat"), uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
TIMEDIFF(`t`.`updateat`, `t`.`createat`)
```
Output PostgreSQL:
```text
DATE_PART('time', "t"."updateat" - "t"."createat")
```

#### TimeSub
Вычитает интервал времени из выражения времени/даты/времени и возвращает результирующее время.
```go
function := uast.TimeSub(uast.Column[time.Time]("t", "createat"), uast.Value("2 HOUR"))
```
Output MySQL:
```text
TIME_SUB(`t`.`createat`, '2 HOUR')
```
Output PostgreSQL:
```text
("t"."createat" - INTERVAL '2 HOUR')
```

#### Week
Извлекает номер недели (1–53) из выражения даты/времени.
```go
function := uast.Week(uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
WEEK(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(WEEK FROM "t"."createat")
```

#### Year
Извлекает год из выражения даты/времени.
```go
function := uast.Year(uast.Column[time.Time]("t", "createat"))
```
Output MySQL:
```text
YEAR(`t`.`createat`)
```
Output PostgreSQL:
```text
EXTRACT(YEAR FROM "t"."createat")
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
Output MySQL:
```text
JSON_ARRAY(`t`.`json`, ?, ?)
```
Output PostgreSQL:
```text
JSON_ARRAY("t"."json", $1, $2)
```

#### JsonArrayAgg
Агрегирует значения из группы в JSON-массив.
```go
function := uast.JsonArrayAgg(
    uast.Column[string]("t", "json"),
)
```
Output MySQL:
```text
JSON_ARRAYAGG(`t`.`json`)
```
Output PostgreSQL:
```text
JSON_AGG("t"."json")
```

#### JsonContains
Проверяет, содержит ли JSON-документ указанное значение.
```go
function := uast.JsonContains(
    uast.Column[string]("t", "json"),
    uast.Value(`{"key":"val"}`),
)
```
Output MySQL:
```text
JSON_CONTAINS(`t`.`json`, '{"key":"val"}')
```
Output PostgreSQL:
```text
("t"."json" @> '{"key":"val"}')
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
Output MySQL:
```text
(`t`.`json` ->> '$.parent[0].child')
```
Output PostgreSQL:
```text
("t"."json" #>> '{parent,0,child}')
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
Output MySQL:
```text
JSON_OBJECT('key', COUNT(`t`.`json`))
```
Output PostgreSQL:
```text
JSON_BUILD_OBJECT('key', COUNT("t"."json"))
```

#### JsonObjectAgg
Агрегирует пары ключ-значение из группы в один JSON-объект.
```go
function := uast.JsonObjectAgg(
    uast.Column[string]("t", "json"),
    uast.Column[int]("t", "number"),
)
```
Output MySQL:
```text
JSON_OBJECTAGG(`t`.`json`, `t`.`number`)
```
Output PostgreSQL:
```text
JSON_OBJECT_AGG("t"."json", "t"."number")
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
Output MySQL:
```text
JSON_REMOVE(`t`.`json`, '$.key1', '$.key2')
```
Output PostgreSQL:
```text
("t"."json" - '{key1}' - '{key2}')
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
Output MySQL:
```text
JSON_SET(`t`.`json`, '$.key1', ?, '$.key2', ?)
```
Output PostgreSQL:
```text
-
```

#### JsonType
Возвращает тип JSON-значения (например, 'OBJECT', 'ARRAY', 'STRING', 'INTEGER', 'NULL').
```go
function := uast.JsonType(uast.Column[string]("t", "json"))
```
Output MySQL:
```text
JSON_TYPE(`t`.`json`)
```
Output PostgreSQL:
```text
jsonb_typeof("t"."json")
```

### Math
#### Abs
Возвращает абсолютное (неотрицательное) значение числового выражения.
```go
function := uast.Abs(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
ABS(`t`.`number`)
```
Output PostgreSQL:
```text
ABS("t"."number")
```

#### ACos
Возвращает арккосинус (обратный косинус) выражения в радианах.
```go
function := uast.ACos(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
ACOS(`t`.`number`)
```
Output PostgreSQL:
```text
ACOS("t"."number")
```

#### ASin
Возвращает арксинус (обратный синус) выражения в радианах.
```go
function := uast.ASin(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
ASIN(`t`.`number`)
```
Output PostgreSQL:
```text
ASIN("t"."number")
```

#### ATan
Возвращает арктангенс (обратный тангенс) выражения в радианах.
```go
function := uast.ATan(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
ATAN(`t`.`number`)
```
Output PostgreSQL:
```text
ATAN("t"."number")
```

#### ATan2
Возвращает арктангенс частного двух аргументов (`y`/`x`), используя их знаки для определения квадранта.
```go
function := uast.ATan2(uast.Column[int]("t", "y"), uast.Column[int]("t", "x"))
```
Output MySQL:
```text
ATAN2(`t`.`y`, `t`.`x`)
```
Output PostgreSQL:
```text
ATAN2("t"."y", "t"."x")
```

#### Cbrt
Возвращает кубический корень числового выражения.
```go
function := uast.Cbrt(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
CBRT(`t`.`number`)
```
Output PostgreSQL:
```text
CBRT("t"."number")
```

#### Ceil
Возвращает наименьшее целое значение, не меньшее аргумента (округление вверх).
```go
function := uast.Ceil(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
CEILING(`t`.`number`)
```
Output PostgreSQL:
```text
CEIL("t"."number")
```

#### Cos
Возвращает косинус выражения в радианах.
```go
function := uast.Cos(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
COS(`t`.`number`)
```
Output PostgreSQL:
```text
COS("t"."number")
```

#### Exp
Возвращает число Эйлера `e` (~2.71828) возведённое в степень выражения.
```go
function := uast.Exp(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
EXP(`t`.`number`)
```
Output PostgreSQL:
```text
EXP("t"."number")
```

#### Floor
Возвращает наибольшее целое значение, не большее аргумента (округление вниз).
```go
function := uast.Floor(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
FLOOR(`t`.`number`)
```
Output PostgreSQL:
```text
FLOOR("t"."number")
```

#### Ln
Возвращает натуральный логарифм (по основанию `e`) выражения.
```go
function := uast.Ln(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
LN(`t`.`number`)
```
Output PostgreSQL:
```text
LN("t"."number")
```

#### Log
Возвращает логарифм выражения по указанному основанию.
```go
function := uast.Log(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
LOG(`t`.`number`, ?)
```
Output PostgreSQL:
```text
LOG("t"."number", $1)
```

#### Mod
Возвращает остаток (модуль) от деления первого выражения на второе.
```go
function := uast.Mod(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
MOD(`t`.`number`, ?)
```
Output PostgreSQL:
```text
MOD("t"."number", $1)
```

#### Pi
Возвращает математическую константу `p` (~3.14159).
```go
function := uast.Pi()
```
Output MySQL:
```text
PI()
```
Output PostgreSQL:
```text
PI()
```

#### Power
Возвращает выражение, возведённое в степень экспоненты.
```go
function := uast.Power(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
POWER(`t`.`number`, ?)
```
Output PostgreSQL:
```text
POWER("t"."number", $1)
```

#### Rand
Возвращает случайное значение с плавающей запятой в диапазоне [0, 1].
```go
function := uast.Rand()
```
Output MySQL:
```text
RAND()
```
Output PostgreSQL:
```text
RANDOM()
```

#### Round
Округляет выражение до указанного количества знаков после запятой.
```go
function := uast.Round(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
ROUND(`t`.`number`, ?)
```
Output PostgreSQL:
```text
ROUND("t"."number", $1)
```

#### Sin
Возвращает синус выражения в радианах.
```go
function := uast.Sin(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
SIN(`t`.`number`)
```
Output PostgreSQL:
```text
SIN("t"."number")
```

#### Sqrt
Возвращает квадратный корень выражения.
```go
function := uast.Sqrt(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
SQRT(`t`.`number`)
```
Output PostgreSQL:
```text
SQRT("t"."number")
```

#### Tan
Возвращает тангенс выражения в радианах.
```go
function := uast.Tan(uast.Column[int]("t", "number"))
```
Output MySQL:
```text
TAN(`t`.`number`)
```
Output PostgreSQL:
```text
TAN("t"."number")
```

#### Trunc
Усекает числовое выражение до указанного количества знаков после запятой (без округления).
```go
function := uast.Trunc(uast.Column[int]("t", "number"), uast.Value(2))
```
Output MySQL:
```text
TRUNCATE(`t`.`number`, ?)
```
Output PostgreSQL:
```text
TRUNC("t"."number", $1)
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
Output MySQL:
```text
CUME_DIST() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
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
Output MySQL:
```text
DENSE_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
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
Output MySQL:
```text
NTILE(2) OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
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
Output MySQL:
```text
PERCENT_RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
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
Output MySQL:
```text
RANK() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
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
Output MySQL:
```text
ROW_NUMBER() OVER (PARTITION BY `t`.`id` ORDER BY `t`.`number` DESC)
```
Output PostgreSQL:
```text
ROW_NUMBER() OVER (PARTITION BY "t"."id" ORDER BY "t"."number" DESC)
```

### String
#### Concat
Объединяет два или более строковых выражения в одну строку. Аргументы `NULL` рассматриваются как пустые строки в большинстве диалектов.
```go
function := uast.Concat(uast.Column[string]("t", "string"), uast.Value("old"), uast.Value("new"))
```
Output MySQL:
```text
CONCAT(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
CONCAT("t"."string", $1, $2)
```

#### ConcatWs
Объединяет два или более строковых выражения с указанным разделителем между ними. Пропускает аргументы `NULL`.
```go
function := uast.ConcatWs(uast.Value("_"), uast.Column[string]("t", "string"), uast.Value("old"),uast.Value("new"))
```
Output MySQL:
```text
CONCAT_WS(?, `t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
CONCAT_WS($1, "t"."string", $2, $3)
```

#### LeftString
Возвращает крайние слева `count` символов из строкового выражения.
```go
function := uast.LeftString(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MySQL:
```text
LEFT(`t`.`string`, ?)
```
Output PostgreSQL:
```text
LEFT("t"."string", $1)
```

#### Lower
Преобразует строковое выражение в нижний регистр.
```go
function := uast.Lower(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
LOWER(`t`.`string`)
```
Output PostgreSQL:
```text
LOWER("t"."string")
```

#### LPad
Дополняет строковое выражение слева указанным разделителем до общей длины `count` символов.
```go
function := uast.LPad(uast.Column[string]("t", "string"), uast.Value(2), uast.Value(","))
```
Output MySQL:
```text
LPAD(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
LPAD("t"."string", $1, $2)
```

#### LTrim
Удаляет начальные пробелы из строкового выражения.
```go
function := uast.LTrim(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
LTRIM(`t`.`string`)
```
Output PostgreSQL:
```text
LTRIM("t"."string")
```

#### Repeat
Повторяет строковое выражение `count` раз.
```go
function := uast.Repeat(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MySQL:
```text
REPEAT(`t`.`string`, ?)
```
Output PostgreSQL:
```text
REPEAT("t"."string", $1)
```

#### Replace
Заменяет все вхождения подстроки в строке на новую подстроку.
```go
function := uast.Replace(uast.Column[string]("t", "string"), uast.Value("old"), uast.Value("new"))
```
Output MySQL:
```text
REPLACE(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
REPLACE("t"."string", $1, $2)
```

#### Reverse
Переворачивает символы в строковом выражении.
```go
function := uast.Reverse(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
REVERSE(`t`.`string`)
```
Output PostgreSQL:
```text
REVERSE("t"."string")
```

#### RightString
Возвращает крайние справа `count` символов из строкового выражения.
```go
function := uast.RightString(uast.Column[string]("t", "string"), uast.Value(2))
```
Output MySQL:
```text
RIGHT(`t`.`string`, ?)
```
Output PostgreSQL:
```text
RIGHT("t"."string", $1)
```

#### RPad
Дополняет строковое выражение справа указанным разделителем до общей длины `count` символов.
```go
function := uast.RPad(uast.Column[string]("t", "string"), uast.Value(2), uast.Value(","))
```
Output MySQL:
```text
RPAD(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
RPAD("t"."string", $1, $2)
```

#### RTrim
Удаляет конечные пробелы из строкового выражения.
```go
function := uast.RTrim(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
RTRIM(`t`.`string`)
```
Output PostgreSQL:
```text
RTRIM("t"."string")
```

#### SubString
Извлекает подстроку из строкового выражения, начиная с `startPos` (начиная с 1) длиной `lengthStr` символов.
```go
function := uast.SubString(uast.Column[string]("t", "string"), uast.Value(0), uast.Value(2))
```
Output MySQL:
```text
SUBSTRING(`t`.`string`, ?, ?)
```
Output PostgreSQL:
```text
SUBSTRING("t"."string", $1, $2)
```

#### Trim
Удаляет как начальные, так и конечные пробелы из строкового выражения.
```go
function := uast.Trim(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
TRIM(`t`.`string`)
```
Output PostgreSQL:
```text
TRIM("t"."string")
```

#### Upper
Преобразует строковое выражение в верхний регистр.
```go
function := uast.Upper(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
UPPER(`t`.`string`)
```
Output PostgreSQL:
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
Output MySQL:
```text
(`t`.`string` = ? AND `t`.`number` > ?)
```
Output PostgreSQL:
```text
("t"."string" = $1 AND "t"."number" > $2)
```

### Or
Комбинирует несколько условий логическим `OR`. Хотя бы одно условие должно быть истинным для истинности комбинированного выражения.
```go
logical := uast.Or(
    uast.Equal(uast.Column[string]("t", "string"), uast.Value("active")),
    uast.Greater(uast.Column[int]("t", "number"), uast.Value(2)),
)
```
Output MySQL:
```text
(`t`.`string` = ? OR `t`.`number` > ?)
```
Output PostgreSQL:
```text
("t"."string" = $1 OR "t"."number" > $2)
```

## exprOrderBy
### Asc
Указывает порядок сортировки по возрастанию (сначала наименьшие, от А до Я). Используется для сортировки строк в запросе или в рамках оконной функции.
```go
order := uast.Asc(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
`t`.`string` ASC
```
Output PostgreSQL:
```text
"t"."string" ASC
```

### Desc
Указывает порядок сортировки по убыванию (сначала наибольшие, от Я до А). Используется для сортировки строк в запросе или в рамках оконной функции.
```go
order := uast.Desc(uast.Column[string]("t", "string"))
```
Output MySQL:
```text
`t`.`string` DESC
```
Output PostgreSQL:
```text
"t"."string" DESC
```

## exprSubquery
### Subquery
Оборачивает оператор `SELECT` как типизированное выражение, которое может использоваться в сравнениях (`In`, `Exists`, `Equal` и т.д.) или как колонка в операторе `SELECT`. Обобщённый параметр `T` указывает скалярный тип единственной колонки, возвращаемой подзапросом.
```go
subquery := uast.Subquery[int64](uast.NewSelect(uast.Column[int64]("t", "id")).From(uast.Table("test")))
```
Output MySQL:
```text
(SELECT `t`.`id` FROM `test` AS `t`)
```
Output PostgreSQL:
```text
(SELECT "t"."id" FROM "test" AS "t")
```

## exprValue
### Value
Оборачивает Go-значение как параметризованное выражение. Значение НЕ вставляется в SQL-строку напрямую — вместо этого генерируется плейсхолдер (`?`, `$1`, и т.д.), а значение добавляется в слайс аргументов, возвращаемый `Build()`. Это безопасный способ передачи пользовательских данных, предотвращающий SQL-инъекции. Поддерживаемые типы: `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `string`, `time.Time`.
```go
var data string = "ivan"
value := uast.Value(data)
```
Output MySQL:
```text
?
```
Output PostgreSQL:
```text
$1
```