---
outline: deep
---

# API / Core / Types

::: info **Info**
This page covers all 21 data types in `Binary`, `Datetime`, `Numeric`, `String`, `Special` categories. Each type is demonstrated using `Cast` and includes a code example with dialect-specific SQL **Output.
:::**

## Binary
### TypeBinary
Fixed-length binary string.
```go
binary := uast.Cast(uast.Field[int]("t", "number"), uast.TypeBinary)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS BINARY)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS BINARY)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS BINARY)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS BYTEA)
```
**Output SQLite:**
```sql
CAST("t"."number" AS BLOB)
```

### TypeVarBinary
Variable-length binary string.
```go
binary := uast.Cast(uast.Field[int]("t", "number"), uast.TypeVarBinary)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS VARBINARY)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS VARBINARY)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS VARBINARY)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS BYTEA)
```
**Output SQLite:**
```sql
CAST("t"."number" AS BLOB)
```

## Datetime
### TypeDate
Represents a date value (year, month, day).
```go
datetime := uast.Cast(uast.Field[int]("t", "number"), uast.TypeDate)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS DATE)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS DATE)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS DATE)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS DATE)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```

### TypeDateTime
Represents a combined date and time value.
```go
datetime := uast.Cast(uast.Field[int]("t", "number"), uast.TypeDateTime)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS DATETIME)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS DATETIME2)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS DATETIME)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS TIMESTAMP)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```

### TypeTime
Represents a time value (hour, minute, second).
```go
datetime := uast.Cast(uast.Field[int]("t", "number"), uast.TypeTime)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS TIME)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS TIME)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS TIME)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS TIME)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```

### TypeTimestamp
Represents a timestamp value.
```go
datetime := uast.Cast(uast.Field[int]("t", "number"), uast.TypeTimestamp)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS TIMESTAMP)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS DATETIME2)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS TIMESTAMP)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS TIMESTAMPTZ)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```

## Numeric
### TypeBigInt
Large integer type.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeBigInt)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS SIGNED)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS BIGINT)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS SIGNED)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS BIGINT)
```
**Output SQLite:**
```sql
CAST("t"."number" AS INTEGER)
```

### TypeDecimal
Fixed-point decimal number.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeDecimal)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS DECIMAL)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS DECIMAL)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS DECIMAL)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS DECIMAL)
```
**Output SQLite:**
```sql
CAST("t"."number" AS REAL)
```

### TypeDouble
Double-precision floating-point number.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeDouble)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS DECIMAL)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS FLOAT)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS DECIMAL)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS DOUBLE PRECISION)
```
**Output SQLite:**
```sql
CAST("t"."number" AS REAL)
```

### TypeFloat
Single-precision floating-point number.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeFloat)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS DECIMAL)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS REAL)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS DECIMAL)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS REAL)
```
**Output SQLite:**
```sql
CAST("t"."number" AS REAL)
```

### TypeInt
Integer type.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeInt)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS SIGNED)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS INT)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS SIGNED)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS INTEGER)
```
**Output SQLite:**
```sql
CAST("t"."number" AS INTEGER)
```

### TypeSmallInt
Small integer type.
```go
math := uast.Cast(uast.Field[int]("t", "number"), uast.TypeSmallInt)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS SIGNED)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS SMALLINT)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS SIGNED)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS SMALLINT)
```
**Output SQLite:**
```sql
CAST("t"."number" AS INTEGER)
```

## String
### TypeChar 
Fixed-length character string.
```go
str := uast.Cast(uast.Field[int]("t", "number"), uast.TypeChar)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS CHAR)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS CHAR)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS CHAR)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS CHAR)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```

### TypeString
Variable-length character string.
```go
str := uast.Cast(uast.Field[int]("t", "number"), uast.TypeString)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS VARCHAR)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS NVARCHAR)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS VARCHAR)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS VARCHAR)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```

### TypeText
Variable-length text string.
```go
str := uast.Cast(uast.Field[int]("t", "number"), uast.TypeText)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS TEXT)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS NVARCHAR(MAX))
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS TEXT)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS TEXT)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```

### TypeVarChar
Variable-length character string with specified maximum.
```go
str := uast.Cast(uast.Field[int]("t", "number"), uast.TypeVarChar)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS VARCHAR)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS NVARCHAR)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS VARCHAR)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS VARCHAR)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```

## Special
### TypeArray
Represents an array type.
```go
special := uast.Cast(uast.Field[int]("t", "number"), uast.TypeArray)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS JSON)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS NVARCHAR(MAX))
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS JSON)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS ARRAY)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```

### TypeBoolean
Represents a boolean (true/false) type.
```go
special := uast.Cast(uast.Field[int]("t", "number"), uast.TypeBoolean)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS TINYINT(1))
```
**Output MsSQL:**
```sql
CAST([t].[number] AS BIT)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS TINYINT(1))
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS BOOLEAN)
```
**Output SQLite:**
```sql
CAST("t"."number" AS INTEGER)
```

### TypeJSON
Represents a JSON data type.
```go
special := uast.Cast(uast.Field[int]("t", "number"), uast.TypeJSON)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS JSON)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS NVARCHAR(MAX))
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS JSON)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS JSONB)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```

### TypeUUID
Represents a universally unique identifier (UUID).
```go
special := uast.Cast(uast.Field[int]("t", "number"), uast.TypeUUID)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS UUID)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS UNIQUEIDENTIFIER)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS CHAR(36))
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS UUID)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```

### TypeXML
Represents an XML data type.
```go
special := uast.Cast(uast.Field[int]("t", "number"), uast.TypeXML)
```
**Output MariaDB:**
```sql
CAST(`t`.`number` AS TEXT)
```
**Output MsSQL:**
```sql
CAST([t].[number] AS XML)
```
**Output MySQL:**
```sql
CAST(`t`.`number` AS TEXT)
```
**Output PostgreSQL:**
```sql
CAST("t"."number" AS XML)
```
**Output SQLite:**
```sql
CAST("t"."number" AS TEXT)
```
