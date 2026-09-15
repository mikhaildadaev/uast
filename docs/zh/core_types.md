---
outline: deep
---

# API / Core / 类型

::: info **关于**
本页面涵盖了 `Binary`、`Datetime`、`Numeric`、`String`、`Special` 类别中的所有 21 种数据类型。每种类型都使用 `Cast` 进行演示，并包含带有方言特定 SQL 输出的代码示例。
:::**

## Binary
### TypeBinary
定长二进制字符串。
```go
binary := uast.Cast(uast.Field[int]("u", "number"), uast.TypeBinary)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS BINARY)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS BINARY)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS BINARY)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS BYTEA)
```
**Output SQLite:**
```sql
CAST("u"."number" AS BLOB)
```

### TypeVarBinary
变长二进制字符串。
```go
binary := uast.Cast(uast.Field[int]("u", "number"), uast.TypeVarBinary)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS VARBINARY)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS VARBINARY)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS VARBINARY)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS BYTEA)
```
**Output SQLite:**
```sql
CAST("u"."number" AS BLOB)
```

## Datetime
### TypeDate
表示日期值（年、月、日）。
```go
datetime := uast.Cast(uast.Field[int]("u", "number"), uast.TypeDate)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS DATE)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS DATE)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS DATE)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS DATE)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

### TypeDateTime
表示组合的日期和时间值。
```go
datetime := uast.Cast(uast.Field[int]("u", "number"), uast.TypeDateTime)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS DATETIME)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS DATETIME2)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS DATETIME)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS TIMESTAMP)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

### TypeTime
表示时间值（时、分、秒）。
```go
datetime := uast.Cast(uast.Field[int]("u", "number"), uast.TypeTime)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS TIME)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS TIME)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS TIME)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS TIME)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

### TypeTimestamp
表示时间戳值。
```go
datetime := uast.Cast(uast.Field[int]("u", "number"), uast.TypeTimestamp)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS TIMESTAMP)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS DATETIME2)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS TIMESTAMP)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS TIMESTAMPTZ)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

## Numeric
### TypeBigInt
大整数类型。
```go
math := uast.Cast(uast.Field[int]("u", "number"), uast.TypeBigInt)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS SIGNED)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS BIGINT)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS SIGNED)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS BIGINT)
```
**Output SQLite:**
```sql
CAST("u"."number" AS INTEGER)
```

### TypeDecimal
定点十进制数。
```go
math := uast.Cast(uast.Field[int]("u", "number"), uast.TypeDecimal)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS DECIMAL)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS DECIMAL)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS DECIMAL)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS DECIMAL)
```
**Output SQLite:**
```sql
CAST("u"."number" AS REAL)
```

### TypeDouble
双精度浮点数。
```go
math := uast.Cast(uast.Field[int]("u", "number"), uast.TypeDouble)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS DECIMAL)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS FLOAT)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS DECIMAL)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS DOUBLE PRECISION)
```
**Output SQLite:**
```sql
CAST("u"."number" AS REAL)
```

### TypeFloat
单精度浮点数。
```go
math := uast.Cast(uast.Field[int]("u", "number"), uast.TypeFloat)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS DECIMAL)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS REAL)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS DECIMAL)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS REAL)
```
**Output SQLite:**
```sql
CAST("u"."number" AS REAL)
```

### TypeInt
整数类型。
```go
math := uast.Cast(uast.Field[int]("u", "number"), uast.TypeInt)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS SIGNED)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS INT)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS SIGNED)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS INTEGER)
```
**Output SQLite:**
```sql
CAST("u"."number" AS INTEGER)
```

### TypeSmallInt
小整数类型。
```go
math := uast.Cast(uast.Field[int]("u", "number"), uast.TypeSmallInt)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS SIGNED)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS SMALLINT)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS SIGNED)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS SMALLINT)
```
**Output SQLite:**
```sql
CAST("u"."number" AS INTEGER)
```

## String
### TypeChar 
定长字符串。
```go
str := uast.Cast(uast.Field[int]("u", "number"), uast.TypeChar)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS CHAR)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS CHAR)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS CHAR)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS CHAR)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

### TypeString
变长字符串。
```go
str := uast.Cast(uast.Field[int]("u", "number"), uast.TypeString)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS VARCHAR)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS NVARCHAR)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS VARCHAR)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS VARCHAR)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

### TypeText
变长文本字符串。
```go
str := uast.Cast(uast.Field[int]("u", "number"), uast.TypeText)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS TEXT)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS NVARCHAR(MAX))
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS TEXT)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS TEXT)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

### TypeVarChar
指定最大长度的变长字符串。
```go
str := uast.Cast(uast.Field[int]("u", "number"), uast.TypeVarChar)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS VARCHAR)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS NVARCHAR)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS VARCHAR)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS VARCHAR)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

## Special
### TypeArray
表示数组类型。
```go
special := uast.Cast(uast.Field[int]("u", "number"), uast.TypeArray)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS JSON)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS NVARCHAR(MAX))
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS JSON)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS ARRAY)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

### TypeBoolean
表示布尔类型（真/假）。
```go
special := uast.Cast(uast.Field[int]("u", "number"), uast.TypeBoolean)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS TINYINT(1))
```
**Output MsSQL:**
```sql
CAST([u].[number] AS BIT)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS TINYINT(1))
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS BOOLEAN)
```
**Output SQLite:**
```sql
CAST("u"."number" AS INTEGER)
```

### TypeJSON
表示 JSON 数据类型。
```go
special := uast.Cast(uast.Field[int]("u", "number"), uast.TypeJSON)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS JSON)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS NVARCHAR(MAX))
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS JSON)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS JSONB)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

### TypeUUID
表示通用唯一标识符（UUID）。
```go
special := uast.Cast(uast.Field[int]("u", "number"), uast.TypeUUID)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS UUID)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS UNIQUEIDENTIFIER)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS CHAR(36))
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS UUID)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```

### TypeXML
表示 XML 数据类型。
```go
special := uast.Cast(uast.Field[int]("u", "number"), uast.TypeXML)
```
**Output MariaDB:**
```sql
CAST(`u`.`number` AS TEXT)
```
**Output MsSQL:**
```sql
CAST([u].[number] AS XML)
```
**Output MySQL:**
```sql
CAST(`u`.`number` AS TEXT)
```
**Output PostgreSQL:**
```sql
CAST("u"."number" AS XML)
```
**Output SQLite:**
```sql
CAST("u"."number" AS TEXT)
```
