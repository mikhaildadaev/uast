---
outline: deep
---

# API / Core / Типы

::: info **Информация**
Эта страница охватывает все 21 тип данных в категориях `Binary`, `Datetime`, `Numeric`, `String`, `Special` Каждый тип демонстрируется с помощью `Cast` и включает пример кода с диалектно-специфичным выводом SQL.
:::

## Binary
### TypeBinary
Бинарная строка фиксированной длины.
```go
binary := uast.Cast(uast.Field[int]("t", "number"), uast.TypeBinary)
```
Output MariaDB:
```text
CAST(`t`.`number` AS BINARY)
```
Output MsSQL:
```text
CAST([t].[number] AS BINARY)
```
Output MySQL:
```text
CAST(`t`.`number` AS BINARY)
```
Output PostgreSQL:
```text
CAST("t"."number" AS BYTEA)
```
Output SQLite:
```text
CAST("t"."number" AS BLOB)
```

### TypeVarBinary
Бинарная строка переменной длины.
```go
binary := uast.Cast(uast.Field[int]("t", "number"), uast.TypeVarBinary)
```
Output MariaDB:
```text
CAST(`t`.`number` AS VARBINARY)
```
Output MsSQL:
```text
CAST([t].[number] AS VARBINARY)
```
Output MySQL:
```text
CAST(`t`.`number` AS VARBINARY)
```
Output PostgreSQL:
```text
CAST("t"."number" AS BYTEA)
```
Output SQLite:
```text
CAST("t"."number" AS BLOB)
```

## Datetime
### TypeDate
Представляет значение даты (год, месяц, день).
```go
datetime := uast.Cast(uast.Field[int]("t", "number"), uast.TypeDate)
```
Output MariaDB:
```text
CAST(`t`.`number` AS DATE)
```
Output MsSQL:
```text
CAST([t].[number] AS DATE)
```
Output MySQL:
```text
CAST(`t`.`number` AS DATE)
```
Output PostgreSQL:
```text
CAST("t"."number" AS DATE)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

### TypeDateTime
Представляет комбинированное значение даты и времени.
```go
datetime := uast.Cast(uast.Field[int]("t", "number"), uast.TypeDateTime)
```
Output MariaDB:
```text
CAST(`t`.`number` AS DATETIME)
```
Output MsSQL:
```text
CAST([t].[number] AS DATETIME2)
```
Output MySQL:
```text
CAST(`t`.`number` AS DATETIME)
```
Output PostgreSQL:
```text
CAST("t"."number" AS TIMESTAMP)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

### TypeTime
Представляет значение времени (час, минута, секунда).
```go
datetime := uast.Cast(uast.Field[int]("t", "number"), uast.TypeTime)
```
Output MariaDB:
```text
CAST(`t`.`number` AS TIME)
```
Output MsSQL:
```text
CAST([t].[number] AS TIME)
```
Output MySQL:
```text
CAST(`t`.`number` AS TIME)
```
Output PostgreSQL:
```text
CAST("t"."number" AS TIME)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

### TypeTimestamp
Представляет значение временной метки.
```go
datetime := uast.Cast(uast.Field[int]("t", "number"), uast.TypeTimestamp)
```
Output MariaDB:
```text
CAST(`t`.`number` AS TIMESTAMP)
```
Output MsSQL:
```text
CAST([t].[number] AS DATETIME2)
```
Output MySQL:
```text
CAST(`t`.`number` AS TIMESTAMP)
```
Output PostgreSQL:
```text
CAST("t"."number" AS TIMESTAMPTZ)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

## Numeric
### TypeBigInt
Большой целочисленный тип.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeBigInt)
```
Output MariaDB:
```text
CAST(`t`.`number` AS SIGNED)
```
Output MsSQL:
```text
CAST([t].[number] AS BIGINT)
```
Output MySQL:
```text
CAST(`t`.`number` AS SIGNED)
```
Output PostgreSQL:
```text
CAST("t"."number" AS BIGINT)
```
Output SQLite:
```text
CAST("t"."number" AS INTEGER)
```

### TypeDecimal
Десятичное число с фиксированной запятой.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeDecimal)
```
Output MariaDB:
```text
CAST(`t`.`number` AS DECIMAL)
```
Output MsSQL:
```text
CAST([t].[number] AS DECIMAL)
```
Output MySQL:
```text
CAST(`t`.`number` AS DECIMAL)
```
Output PostgreSQL:
```text
CAST("t"."number" AS DECIMAL)
```
Output SQLite:
```text
CAST("t"."number" AS REAL)
```

### TypeDouble
Число с плавающей запятой двойной точности.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeDouble)
```
Output MariaDB:
```text
CAST(`t`.`number` AS DECIMAL)
```
Output MsSQL:
```text
CAST([t].[number] AS FLOAT)
```
Output MySQL:
```text
CAST(`t`.`number` AS DECIMAL)
```
Output PostgreSQL:
```text
CAST("t"."number" AS DOUBLE PRECISION)
```
Output SQLite:
```text
CAST("t"."number" AS REAL)
```

### TypeFloat
Число с плавающей запятой одинарной точности.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeFloat)
```
Output MariaDB:
```text
CAST(`t`.`number` AS DECIMAL)
```
Output MsSQL:
```text
CAST([t].[number] AS REAL)
```
Output MySQL:
```text
CAST(`t`.`number` AS DECIMAL)
```
Output PostgreSQL:
```text
CAST("t"."number" AS REAL)
```
Output SQLite:
```text
CAST("t"."number" AS REAL)
```

### TypeInt
Целочисленный тип.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeInt)
```
Output MariaDB:
```text
CAST(`t`.`number` AS SIGNED)
```
Output MsSQL:
```text
CAST([t].[number] AS INT)
```
Output MySQL:
```text
CAST(`t`.`number` AS SIGNED)
```
Output PostgreSQL:
```text
CAST("t"."number" AS INTEGER)
```
Output SQLite:
```text
CAST("t"."number" AS INTEGER)
```

### TypeSmallInt
Малый целочисленный тип.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeSmallInt)
```
Output MariaDB:
```text
CAST(`t`.`number` AS SIGNED)
```
Output MsSQL:
```text
CAST([t].[number] AS SMALLINT)
```
Output MySQL:
```text
CAST(`t`.`number` AS SIGNED)
```
Output PostgreSQL:
```text
CAST("t"."number" AS SMALLINT)
```
Output SQLite:
```text
CAST("t"."number" AS INTEGER)
```

## String
### TypeChar 
Символьная строка фиксированной длины.
```go
str := uast.Cast(uast.Field[int]("t", "number"), uast.TypeChar)
```
Output MariaDB:
```text
CAST(`t`.`number` AS CHAR)
```
Output MsSQL:
```text
CAST([t].[number] AS CHAR)
```
Output MySQL:
```text
CAST(`t`.`number` AS CHAR)
```
Output PostgreSQL:
```text
CAST("t"."number" AS CHAR)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

### TypeString
Символьная строка переменной длины.
```go
str := uast.Cast(uast.Field[int]("t", "number"), uast.TypeString)
```
Output MariaDB:
```text
CAST(`t`.`number` AS VARCHAR)
```
Output MsSQL:
```text
CAST([t].[number] AS NVARCHAR)
```
Output MySQL:
```text
CAST(`t`.`number` AS VARCHAR)
```
Output PostgreSQL:
```text
CAST("t"."number" AS VARCHAR)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

### TypeText
Текстовая строка переменной длины.
```go
str := uast.Cast(uast.Field[int]("t", "number"), uast.TypeText)
```
Output MariaDB:
```text
CAST(`t`.`number` AS TEXT)
```
Output MsSQL:
```text
CAST([t].[number] AS NVARCHAR(MAX))
```
Output MySQL:
```text
CAST(`t`.`number` AS TEXT)
```
Output PostgreSQL:
```text
CAST("t"."number" AS TEXT)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

### TypeVarChar
Символьная строка переменной длины с указанным максимумом.
```go
str := uast.Cast(uast.Field[int]("t", "number"), uast.TypeVarChar)
```
Output MariaDB:
```text
CAST(`t`.`number` AS VARCHAR)
```
Output MsSQL:
```text
CAST([t].[number] AS NVARCHAR)
```
Output MySQL:
```text
CAST(`t`.`number` AS VARCHAR)
```
Output PostgreSQL:
```text
CAST("t"."number" AS VARCHAR)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

## Special
### TypeArray
Представляет тип массива.
```go
special := uast.Cast(uast.Field[int]("t", "number"), uast.TypeArray)
```
Output MariaDB:
```text
CAST(`t`.`number` AS JSON)
```
Output MsSQL:
```text
CAST([t].[number] AS NVARCHAR(MAX))
```
Output MySQL:
```text
CAST(`t`.`number` AS JSON)
```
Output PostgreSQL:
```text
CAST("t"."number" AS ARRAY)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

### TypeBoolean
Представляет булевый тип (истина/ложь).
```go
special := uast.Cast(uast.Field[int]("t", "number"), uast.TypeBoolean)
```
Output MariaDB:
```text
CAST(`t`.`number` AS TINYINT(1))
```
Output MsSQL:
```text
CAST([t].[number] AS BIT)
```
Output MySQL:
```text
CAST(`t`.`number` AS TINYINT(1))
```
Output PostgreSQL:
```text
CAST("t"."number" AS BOOLEAN)
```
Output SQLite:
```text
CAST("t"."number" AS INTEGER)
```

### TypeJSON
Представляет тип данных JSON.
```go
special := uast.Cast(uast.Field[int]("t", "number"), uast.TypeJSON)
```
Output MariaDB:
```text
CAST(`t`.`number` AS JSON)
```
Output MsSQL:
```text
CAST([t].[number] AS NVARCHAR(MAX))
```
Output MySQL:
```text
CAST(`t`.`number` AS JSON)
```
Output PostgreSQL:
```text
CAST("t"."number" AS JSONB)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

### TypeUUID
Представляет универсальный уникальный идентификатор (UUID).
```go
special := uast.Cast(uast.Field[int]("t", "number"), uast.TypeUUID)
```
Output MariaDB:
```text
CAST(`t`.`number` AS UUID)
```
Output MsSQL:
```text
CAST([t].[number] AS UNIQUEIDENTIFIER)
```
Output MySQL:
```text
CAST(`t`.`number` AS CHAR(36))
```
Output PostgreSQL:
```text
CAST("t"."number" AS UUID)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```

### TypeXML
Представляет тип данных XML.
```go
special := uast.Cast(uast.Field[int]("t", "number"), uast.TypeXML)
```
Output MariaDB:
```text
CAST(`t`.`number` AS TEXT)
```
Output MsSQL:
```text
CAST([t].[number] AS XML)
```
Output MySQL:
```text
CAST(`t`.`number` AS TEXT)
```
Output PostgreSQL:
```text
CAST("t"."number" AS XML)
```
Output SQLite:
```text
CAST("t"."number" AS TEXT)
```
