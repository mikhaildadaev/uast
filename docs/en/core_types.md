---
outline: deep
---

# API / Core / Types

::: info **Info**
This page covers all 21 data types in `Binary`, `Datetime`, `Numeric`, `String`, `Special` categories. Each type is demonstrated using `Cast` and includes a code example with dialect-specific SQL output.
:::

## Binary
Binary types store raw byte data. Fixed-length and variable-length binary strings have different representations depending on the SQL dialect.

### TypeBinary
Fixed-length binary string.
```go
binary := uast.Cast(uast.Column[int]("t", "number"), uast.TypeBinary)
```
Output MySQL:
```text
CAST("t"."number" AS BINARY)
```
Output PostgreSQL:
```text
CAST("t"."number" AS BYTEA)
```

### TypeVarBinary
Variable-length binary string.
```go
binary := uast.Cast(uast.Column[int]("t", "number"), uast.TypeVarBinary)
```
Output MySQL:
```text
CAST("t"."number" AS VARBINARY)
```
Output PostgreSQL:
```text
CAST("t"."number" AS BYTEA)
```

## Datetime
Date and time types store temporal values. Some dialects use a single type for multiple temporal representations, while others distinguish between them with specific types.

### TypeDate
Represents a date value (year, month, day).
```go
datetime := uast.Cast(uast.Column[int]("t", "number"), uast.TypeDate)
```
Output:
```text
CAST("t"."number" AS DATE)
```

### TypeDateTime
Represents a combined date and time value.
```go
datetime := uast.Cast(uast.Column[int]("t", "number"), uast.TypeDateTime)
```
Output MySQL:
```text
CAST("t"."number" AS DATETIME)
```
Output PostgreSQL:
```text
CAST("t"."number" AS TIMESTAMP)
```

### TypeTime
Represents a time value (hour, minute, second).
```go
datetime := uast.Cast(uast.Column[int]("t", "number"), uast.TypeTime)
```
Output:
```text
CAST("t"."number" AS TIME)
```

### TypeTimestamp
Represents a timestamp value.
```go
datetime := uast.Cast(uast.Column[int]("t", "number"), uast.TypeTimestamp)
```
Output MySQL:
```text
CAST("t"."number" AS TIMESTAMP)
```
Output PostgreSQL:
```text
CAST("t"."number" AS TIMESTAMPTZ)
```

## Numeric
Numeric types store integer and floating-point values. Some dialects use general-purpose numeric types, while others provide a richer set of specific types.

### TypeBigInt
Large integer type.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeBigInt)
```
Output MySQL:
```text
CAST("t"."number" AS SIGNED)
```
Output PostgreSQL:
```text
CAST("t"."number" AS BIGINT)
```

### TypeDecimal
Fixed-point decimal number.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeDecimal)
```
Output:
```text
CAST("t"."number" AS DECIMAL)
```

### TypeDouble
Double-precision floating-point number.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeDouble)
```
Output MySQL:
```text
CAST("t"."number" AS DECIMAL)
```
Output PostgreSQL:
```text
CAST("t"."number" AS DOUBLE PRECISION)
```

### TypeFloat
Single-precision floating-point number.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeFloat)
```
Output MySQL:
```text
CAST("t"."number" AS DECIMAL)
```
Output PostgreSQL:
```text
CAST("t"."number" AS REAL)
```

### TypeInt
Integer type.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeInt)
```
Output MySQL:
```text
CAST("t"."number" AS SIGNED)
```
Output PostgreSQL:
```text
CAST("t"."number" AS INTEGER)
```

### TypeSmallInt
Small integer type.
```go
math := uast.Cast(uast.Column[int]("t", "number"), uast.TypeSmallInt)
```
Output MySQL:
```text
CAST("t"."number" AS SIGNED)
```
Output PostgreSQL:
```text
CAST("t"."number" AS SMALLINT)
```

## String
String types store character and text data. Variable-length, fixed-length, and large text types are represented differently across dialects.

### TypeChar 
Fixed-length character string.
```go
str := uast.Cast(uast.Column[int]("t", "number"), uast.TypeChar)
```
Output:
```text
CAST("t"."number" AS CHAR)
```

### TypeString
Variable-length character string.
```go
str := uast.Cast(uast.Column[int]("t", "number"), uast.TypeString)
```
Output:
```text
CAST("t"."number" AS VARCHAR)
```

### TypeText
Variable-length text string.
```go
str := uast.Cast(uast.Column[int]("t", "number"), uast.TypeText)
```
Output:
```text
CAST("t"."number" AS TEXT)
```

### TypeVarChar
Variable-length character string with specified maximum.
```go
str := uast.Cast(uast.Column[int]("t", "number"), uast.TypeVarChar)
```
Output:
```text
CAST("t"."number" AS VARCHAR)
```

## Special
Special types cover `Array`, `Boolean`, `JSON`, `UUID`, `XML` data. Some dialects use native types, while others rely on compatible alternatives.

### TypeArray
Represents an array type.
```go
special := uast.Cast(uast.Column[int]("t", "number"), uast.TypeArray)
```
Output MySQL:
```text
CAST("t"."number" AS JSON)
```
Output PostgreSQL:
```text
CAST("t"."number" AS ARRAY)
```

### TypeBoolean
Represents a boolean (true/false) type.
```go
special := uast.Cast(uast.Column[int]("t", "number"), uast.TypeBoolean)
```
Output MySQL:
```text
CAST("t"."number" AS TINYINT(1))
```
Output PostgreSQL:
```text
CAST("t"."number" AS BOOLEAN)
```

### TypeJSON
Represents a JSON data type.
```go
special := uast.Cast(uast.Column[int]("t", "number"), uast.TypeJSON)
```
Output MySQL:
```text
CAST("t"."number" AS JSON)
```
Output PostgreSQL:
```text
CAST("t"."number" AS JSONB)
```

### TypeUUID
Represents a universally unique identifier (UUID).
```go
special := uast.Cast(uast.Column[int]("t", "number"), uast.TypeUUID)
```
Output MySQL:
```text
CAST("t"."number" AS CHAR(36))
```
Output PostgreSQL:
```text
CAST("t"."number" AS UUID)
```

### TypeXML
Represents an XML data type.
```go
special := uast.Cast(uast.Column[int]("t", "number"), uast.TypeXML)
```
Output MySQL:
```text
CAST("t"."number" AS TEXT)
```
Output PostgreSQL:
```text
CAST("t"."number" AS XML)
```
