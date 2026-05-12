---
outline: deep
---

# API / Ядро / Типы

::: info **Информация**
Эта страница охватывает все 21 тип данных в категориях `Binary`, `Datetime`, `Numeric`, `String`, `Special` Каждый тип демонстрируется с помощью `Cast` и включает пример кода с диалектно-специфичным выводом SQL.
:::

## Binary
Бинарные типы хранят необработанные байтовые данные. Строки фиксированной и переменной длины имеют разные представления в зависимости от диалекта SQL.

### TypeBinary
Бинарная строка фиксированной длины.
```go
binary := uast.Cast(uast.Column[int]("t", "number"), uast.TypeBinary)
```
Output MySQL:
```text
CAST(`t`.`number` AS BINARY)
```
Output PostgreSQL:
```text
CAST("t"."number" AS BYTEA)
```

### TypeVarBinary
Бинарная строка переменной длины.
```go
binary := uast.Cast(uast.Column[int]("t", "number"), uast.TypeVarBinary)
```
Output MySQL:
```text
CAST(`t`.`number` AS VARBINARY)
```
Output PostgreSQL:
```text
CAST("t"."number" AS BYTEA)
```

## Datetime
Типы даты и времени хранят временные значения. Некоторые диалекты используют один тип для нескольких временных представлений, в то время как другие различают их с помощью специфичных типов.

### TypeDate
Представляет значение даты (год, месяц, день).
```go
datetime := uast.Cast(uast.Column[int]("t", "number"), uast.TypeDate)
```
Output MySQL:
```text
CAST(`t`.`number` AS DATE)
```
Output PostgreSQL:
```text
CAST("t"."number" AS DATE)
```

### TypeDateTime
Представляет комбинированное значение даты и времени.
```go
datetime := uast.Cast(uast.Column[int]("t", "number"), uast.TypeDateTime)
```
Output MySQL:
```text
CAST(`t`.`number` AS DATETIME)
```
Output PostgreSQL:
```text
CAST("t"."number" AS TIMESTAMP)
```

### TypeTime
Представляет значение времени (час, минута, секунда).
```go
datetime := uast.Cast(uast.Column[int]("t", "number"), uast.TypeTime)
```
Output MySQL:
```text
CAST(`t`.`number` AS TIME)
```
Output PostgreSQL:
```text
CAST("t"."number" AS TIME)
```

### TypeTimestamp
Представляет значение временной метки.
```go
datetime := uast.Cast(uast.Column[int]("t", "number"), uast.TypeTimestamp)
```
Output MySQL:
```text
CAST(`t`.`number` AS TIMESTAMP)
```
Output PostgreSQL:
```text
CAST("t"."number" AS TIMESTAMPTZ)
```

## Numeric
Числовые типы хранят целочисленные значения и значения с плавающей запятой. Некоторые диалекты используют универсальные числовые типы, в то время как другие предоставляют более богатый набор специфичных типов.

### TypeBigInt
Большой целочисленный тип.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeBigInt)
```
Output MySQL:
```text
CAST(`t`.`number` AS SIGNED)
```
Output PostgreSQL:
```text
CAST("t"."number" AS BIGINT)
```

### TypeDecimal
Десятичное число с фиксированной запятой.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeDecimal)
```
Output MySQL:
```text
CAST(`t`.`number` AS DECIMAL)
```
Output PostgreSQL:
```text
CAST("t"."number" AS DECIMAL)
```

### TypeDouble
Число с плавающей запятой двойной точности.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeDouble)
```
Output MySQL:
```text
CAST(`t`.`number` AS DECIMAL)
```
Output PostgreSQL:
```text
CAST("t"."number" AS DOUBLE PRECISION)
```

### TypeFloat
Число с плавающей запятой одинарной точности.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeFloat)
```
Output MySQL:
```text
CAST(`t`.`number` AS DECIMAL)
```
Output PostgreSQL:
```text
CAST("t"."number" AS REAL)
```

### TypeInt
Целочисленный тип.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeInt)
```
Output MySQL:
```text
CAST(`t`.`number` AS SIGNED)
```
Output PostgreSQL:
```text
CAST("t"."number" AS INTEGER)
```

### TypeSmallInt
Малый целочисленный тип.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeSmallInt)
```
Output MySQL:
```text
CAST(`t`.`number` AS SIGNED)
```
Output PostgreSQL:
```text
CAST("t"."number" AS SMALLINT)
```

## String
Строковые типы хранят символьные и текстовые данные. Типы переменной длины, фиксированной длины и большие текстовые типы представлены по-разному в разных диалектах.

### TypeChar 
Символьная строка фиксированной длины.
```go
str := uast.Cast(uast.Column[int]("t", "number"), uast.TypeChar)
```
Output MySQL:
```text
CAST(`t`.`number` AS CHAR)
```
Output PostgreSQL:
```text
CAST("t"."number" AS CHAR)
```

### TypeString
Символьная строка переменной длины.
```go
str := uast.Cast(uast.Column[int]("t", "number"), uast.TypeString)
```
Output MySQL:
```text
CAST(`t`.`number` AS VARCHAR)
```
Output PostgreSQL:
```text
CAST("t"."number" AS VARCHAR)
```

### TypeText
Текстовая строка переменной длины.
```go
str := uast.Cast(uast.Column[int]("t", "number"), uast.TypeText)
```
Output MySQL:
```text
CAST(`t`.`number` AS TEXT)
```
Output PostgreSQL:
```text
CAST("t"."number" AS TEXT)
```

### TypeVarChar
Символьная строка переменной длины с указанным максимумом.
```go
str := uast.Cast(uast.Column[int]("t", "number"), uast.TypeVarChar)
```
Output MySQL:
```text
CAST(`t`.`number` AS VARCHAR)
```
Output PostgreSQL:
```text
CAST("t"."number" AS VARCHAR)
```

## Special
Специальные типы охватывают `Array`, `Boolean`, `JSON`, `UUID`, `XML` данные. Некоторые диалекты используют нативные типы, в то время как другие полагаются на совместимые альтернативы.

### TypeArray
Представляет тип массива.
```go
special := uast.Cast(uast.Column[int]("t", "number"), uast.TypeArray)
```
Output MySQL:
```text
CAST(`t`.`number` AS JSON)
```
Output PostgreSQL:
```text
CAST("t"."number" AS ARRAY)
```

### TypeBoolean
Представляет булевый тип (истина/ложь).
```go
special := uast.Cast(uast.Column[int]("t", "number"), uast.TypeBoolean)
```
Output MySQL:
```text
CAST(`t`.`number` AS TINYINT(1))
```
Output PostgreSQL:
```text
CAST("t"."number" AS BOOLEAN)
```

### TypeJSON
Представляет тип данных JSON.
```go
special := uast.Cast(uast.Column[int]("t", "number"), uast.TypeJSON)
```
Output MySQL:
```text
CAST(`t`.`number` AS JSON)
```
Output PostgreSQL:
```text
CAST("t"."number" AS JSONB)
```

### TypeUUID
Представляет универсальный уникальный идентификатор (UUID).
```go
special := uast.Cast(uast.Column[int]("t", "number"), uast.TypeUUID)
```
Output MySQL:
```text
CAST(`t`.`number` AS CHAR(36))
```
Output PostgreSQL:
```text
CAST("t"."number" AS UUID)
```

### TypeXML
Представляет тип данных XML.
```go
special := uast.Cast(uast.Column[int]("t", "number"), uast.TypeXML)
```
Output MySQL:
```text
CAST(`t`.`number` AS TEXT)
```
Output PostgreSQL:
```text
CAST("t"."number" AS XML)
```
