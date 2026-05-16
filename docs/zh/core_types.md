---
outline: deep
---

# API / 核心 / 类型

::: info **关于**
本页面涵盖了 `Binary`、`Datetime`、`Numeric`、`String`、`Special` 类别中的所有 21 种数据类型。每种类型都使用 `Cast` 进行演示，并包含带有方言特定 SQL 输出的代码示例。
:::

## Binary
### TypeBinary
定长二进制字符串。
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
变长二进制字符串。
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
### TypeDate
表示日期值（年、月、日）。
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
表示组合的日期和时间值。
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
表示时间值（时、分、秒）。
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
表示时间戳值。
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
### TypeBigInt
大整数类型。
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
定点十进制数。
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
双精度浮点数。
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
单精度浮点数。
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
整数类型。
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
小整数类型。
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
### TypeChar 
定长字符串。
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
变长字符串。
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
变长文本字符串。
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
指定最大长度的变长字符串。
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
### TypeArray
表示数组类型。
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
表示布尔类型（真/假）。
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
表示 JSON 数据类型。
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
表示通用唯一标识符（UUID）。
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
表示 XML 数据类型。
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
