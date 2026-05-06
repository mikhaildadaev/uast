---
outline: deep
---

# API / Core / Options

::: info **Info**
This page covers all configuration options: `Extractor`, `Format`, `Level`, `Mode`, `Theme`. Each option is shown with a working code example and expected output.
:::

## Array
Constructs an array expression for use in SQL queries.
```go
arr := uast.Array(1, 2, 3)
```
Output:
```text
ARRAY[1, 2, 3]
```

## Binary
### BitwiseAnd
Performs a bitwise AND operation between two expressions.
```go
expr := uast.BitwiseAnd(uast.Column[int]("t", "flags"), uast.Value(0b0011))
```
Output:
```text
"t"."flags" & 3
```

### BitwiseOr
Performs a bitwise OR operation between two expressions.
```go
expr := uast.BitwiseOr(uast.Column[int]("t", "flags"), uast.Value(0b1100))
```
Output:
```text
"t"."flags" | 12
```

### BitwiseXor
Performs a bitwise XOR operation between two expressions.
```go
expr := uast.BitwiseXor(uast.Column[int]("t", "flags"), uast.Value(0b1111))
```
Output:
```text
"t"."flags" ^ 15
```

### Divide
Divides the left expression by the right expression.
```go
expr := uast.Divide(uast.Column[float64]("t", "price"), uast.Value(2.0))
```
Output:
```text
"t"."price" / 2.000000
```

### Minus
Subtracts the right expression from the left expression.
```go
expr := uast.Minus(uast.Column[int]("t", "total"), uast.Value(10))
```
Output:
```text
"t"."total" - 10
```

### Modulo
Returns the remainder of dividing the left expression by the right expression.
```go
expr := uast.Modulo(uast.Column[int]("t", "id"), uast.Value(5))
```
Output:
```text
"t"."id" % 5
```

### Multiply
Multiplies the left expression by the right expression.
```go
expr := uast.Multiply(uast.Column[float64]("t", "price"), uast.Value(1.2))
```
Output:
```text
"t"."price" * 1.200000
```

### Plus
Adds the left expression to the right expression.
```go
expr := uast.Plus(uast.Column[int]("t", "amount"), uast.Value(100))
```
Output:
```text
"t"."amount" + 100
```

### ShiftLeft
Performs a bitwise left shift on the left expression by the number of bits specified in the right expression.
```go
expr := uast.ShiftLeft(uast.Column[int]("t", "mask"), uast.Value(4))
```
Output:
```text
"t"."mask" << 4
```

### ShiftRight
Performs a bitwise right shift on the left expression by the number of bits specified in the right expression.
```go
expr := uast.ShiftRight(uast.Column[int]("t", "mask"), uast.Value(2))
```
Output:
```text
"t"."mask" >> 2
```

## Column
Creates a reference to a table column, optionally qualified with a table alias. This is the primary way to reference database columns in expressions.
```go
col := uast.Column[string]("u", "email")
```
Output:
```text
"u"."email"
```

## Comparison
### Between
Checks if the left expression falls within the range defined by `valueStart` and `valueEnd` (inclusive).
```go
expr := uast.Between(uast.Column[int]("t", "age"), uast.Value(18), uast.Value(65))
```
Output:
```text
"t"."age" BETWEEN 18 AND 65
```

### Equal
Compares two expressions for equality (=).
```go
expr := uast.Equal(uast.Column[int]("t", "status"), uast.Value("active"))
```
Output:
```text
"t"."status" = 'active'
```

### Exists
Checks if the subquery returns any rows. Returns `true` if at least one row exists.
```go
sub := uast.NewSelect(uast.Column[int]("*")).From(uast.Table("orders")).Where(
    uast.Equal(
        uast.Column[int]("orders", "user_id"),
        uast.Column[int]("users", "id"),
    ),
)
expr := uast.Exists(uast.Subquery[int](sub))
```
Output:
```text
EXISTS (SELECT * FROM "orders" WHERE "orders"."user_id" = "users"."id")
```

### Greater
Compares if the left expression is greater than the right expression (>).
```go
expr := uast.Greater(uast.Column[float64]("t", "price"), uast.Value(100.0))
```
Output:
```text
"t"."price" > 100.000000
```

### GreaterEqual
Compares if the left expression is greater than or equal to the right expression (>=).
```go
expr := uast.GreaterEqual(uast.Column[int]("t", "quantity"), uast.Value(1))
```
Output:
```text
"t"."quantity" >= 1
```

### ILike
Performs a case-insensitive pattern matching comparison. The right expression should contain a pattern with `%` (any sequence) and `_` (single character) wildcards.
```go
expr := uast.ILike(uast.Column[string]("t", "name"), uast.Value("%ivan%"))
```
Output:
```text
"t"."name" ILIKE '%ivan%'
```

### In
Checks if the left expression matches any value contained within the right expression (typically a subquery or array).
```go
expr := uast.In(
    uast.Column[int]("t", "category_id"),
    uast.Subquery[int](
        uast.NewSelect(uast.Column[int]("c", "id")).
            From(uast.Table("categories")).
            Where(uast.Equal(uast.Column[string]("c", "type"), uast.Value("premium"))),
    ),
)
```
Output:
```text
"t"."category_id" IN (SELECT "c"."id" FROM "categories" WHERE "c"."type" = 'premium')
```

### IsNotNull
Checks if the expression is not `NULL`.
```go
expr := uast.IsNotNull(uast.Column[string]("t", "deleted_at"))
```
Output:
```text
"t"."deleted_at" IS NOT NULL
```

### IsNull
Checks if the expression is `NULL`.
```go
expr := uast.IsNull(uast.Column[string]("t", "deleted_at"))
```
Output:
```text
"t"."deleted_at" IS NULL
```

### Less
...
```go
...
```
Output:
```text
...
```

### LessEqual
...
```go
...
```
Output:
```text
...
```

### Like
...
```go
...
```
Output:
```text
...
```

### NotBetween
...
```go
...
```
Output:
```text
...
```

### NotEqual
...
```go
...
```
Output:
```text
...
```

### NotExists
...
```go
...
```
Output:
```text
...
```

### NotILike
...
```go
...
```
Output:
```text
...
```

### NotIn
...
```go
...
```
Output:
```text
...
```

### NotLike
...
```go
...
```
Output:
```text
...
```

## Constant
### ConstBoolFalse
...
```go
...
```
Output:
```text
...
```

### ConstBoolTrue
...
```go
...
```
Output:
```text
...
```

### ConstFloat32One
...
```go
...
```
Output:
```text
...
```

### ConstFloat64One
...
```go
...
```
Output:
```text
...
```

### ConstIntOne
...
```go
...
```
Output:
```text
...
```

### ConstInt8One
...
```go
...
```
Output:
```text
...
```

### ConstInt16One
...
```go
...
```
Output:
```text
...
```

### ConstInt32One
...
```go
...
```
Output:
```text
...
```

### ConstInt64One
...
```go
...
```
Output:
```text
...
```

### ConstNullDefault
...
```go
...
```
Output:
```text
...
```

### ConstStringDefault
...
```go
...
```
Output:
```text
...
```

### ConstUintOne
...
```go
...
```
Output:
```text
...
```

### ConstUint8One
...
```go
...
```
Output:
```text
...
```

### ConstUint16One
...
```go
...
```
Output:
```text
...
```

### ConstUint32One
...
```go
...
```
Output:
```text
...
```

### ConstUint64One
...
```go
...
```
Output:
```text
...
```

## Function
### Aggregate
#### Avg...
```go
...
```
Output:
```text
...
```

#### BitAnd```go
...
```
Output:
```text
...
```

#### BitOr```go
...
```
Output:
```text
...
```

#### BitXor
```go
...
```
Output:
```text
...
```

#### Count
```go
...
```
Output:
```text
...
```

#### GroupConcat
```go
...
```
Output:
```text
...
```

#### Max
```go
...
```
Output:
```text
...
```

#### Min
```go
...
```
Output:
```text
...
```

#### StdDev
```go
...
```
Output:
```text
...
```

#### Sum
```go
...
```
Output:
```text
...
```

#### Variance
```go
...
```
Output:
```text
...
```

### Analytical
#### FirstValue
```go
...
```
Output:
```text
...
```

#### Lag
```go
...
```
Output:
```text
...
```

#### LastValue
```go
...
```
Output:
```text
...
```

#### Lead
```go
...
```
Output:
```text
...
```

#### NthValue
```go
...
```
Output:
```text
...
```

### Condition
#### Case
```go
...
```
Output:
```text
...
```

#### Coalesce
```go
...
```
Output:
```text
...
```

#### Greatest
```go
...
```
Output:
```text
...
```

#### Least
```go
...
```
Output:
```text
...
```

#### NullIf
```go
...
```
Output:
```text
...
```

### Convert
#### Cast
```go
...
```
Output:
```text
...
```

#### CharLength
```go
...
```
Output:
```text
...
```

#### DateFormat
```go
...
```
Output:
```text
...
```

#### Degrees
```go
...
```
Output:
```text
...
```

#### Length
```go
...
```
Output:
```text
...
```

#### Position
```go
...
```
Output:
```text
...
```

#### Radians
```go
...
```
Output:
```text
...
```

### Date and time
#### CurDate
```go
...
```
Output:
```text
...
```

#### CurTime
```go
...
```
Output:
```text
...
```

#### DateAdd
```go
...
```
Output:
```text
...
```

#### DateDiff
```go
...
```
Output:
```text
...
```

#### DateSub
```go
...
```
Output:
```text
...
```

#### Day
```go
...
```
Output:
```text
...
```

#### DayName
```go
...
```
Output:
```text
...
```

#### Hour
```go
...
```
Output:
```text
...
```

#### Minute
```go
...
```
Output:
```text
...
```

#### Month
```go
...
```
Output:
```text
...
```

#### MonthName
```go
...
```
Output:
```text
...
```

#### Now
```go
...
```
Output:
```text
...
```

#### Quarter
```go
...
```
Output:
```text
...
```

#### Second
```go
...
```
Output:
```text
...
```

#### TimeAdd
```go
...
```
Output:
```text
...
```

#### TimeDiff
```go
...
```
Output:
```text
...
```

#### TimeSub
```go
...
```
Output:
```text
...
```

#### Week
```go
...
```
Output:
```text
...
```

#### Year
```go
...
```
Output:
```text
...
```

### Json
#### JsonArray
```go
...
```
Output:
```text
...
```

#### JsonArrayAgg
```go
...
```
Output:
```text
...
```

#### JsonContains
```go
...
```
Output:
```text
...
```

#### JsonExtract
```go
...
```
Output:
```text
...
```

#### JsonObject
```go
...
```
Output:
```text
...
```

#### JsonObjectAgg
```go
...
```
Output:
```text
...
```

#### JsonRemove
```go
...
```
Output:
```text
...
```

#### JsonSet
```go
...
```
Output:
```text
...
```

#### JsonType
```go
...
```
Output:
```text
...
```

### Math
#### Abs
```go
...
```
Output:
```text
...
```

#### ACos
```go
...
```
Output:
```text
...
```

#### ASin
```go
...
```
Output:
```text
...
```

#### ATan
```go
...
```
Output:
```text
...
```

#### ATan2
```go
...
```
Output:
```text
...
```

#### Cbrt
```go
...
```
Output:
```text
...
```

#### Ceil
```go
...
```
Output:
```text
...
```

#### Cos
```go
...
```
Output:
```text
...
```

#### Exp
```go
...
```
Output:
```text
...
```

#### Floor
```go
...
```
Output:
```text
...
```

#### Ln
```go
...
```
Output:
```text
...
```

#### Log
```go
...
```
Output:
```text
...
```

#### Mod
```go
...
```
Output:
```text
...
```

#### Pi
```go
...
```
Output:
```text
...
```

#### Power
```go
...
```
Output:
```text
...
```

#### Rand
```go
...
```
Output:
```text
...
```

#### Round
```go
...
```
Output:
```text
...
```

#### Sin
```go
...
```
Output:
```text
...
```

#### Sqrt
```go
...
```
Output:
```text
...
```

#### Tan
```go
...
```
Output:
```text
...
```

#### Trunc
```go
...
```
Output:
```text
...
```

### String
#### Concat
```go
...
```
Output:
```text
...
```

#### ConcatWs
```go
...
```
Output:
```text
...
```

#### LeftString
```go
...
```
Output:
```text
...
```

#### Lower
```go
...
```
Output:
```text
...
```

#### LPad
```go
...
```
Output:
```text
...
```

#### LTrim
```go
...
```
Output:
```text
...
```

#### Repeat
```go
...
```
Output:
```text
...
```

#### Replace
```go
...
```
Output:
```text
...
```

#### Reverse
```go
...
```
Output:
```text
...
```

#### RightString
```go
...
```
Output:
```text
...
```

#### RPad
```go
...
```
Output:
```text
...
```

#### RTrim
```go
...
```
Output:
```text
...
```

#### SubString
```go
...
```
Output:
```text
...
```

#### Trim
```go
...
```
Output:
```text
...
```

#### Upper
```go
...
```
Output:
```text
...
```

### Ranking
#### CumeDist
```go
...
```
Output:
```text
...
```

#### DenseRank
```go
...
```
Output:
```text
...
```

#### NTile
```go
...
```
Output:
```text
...
```

#### PercentRank
```go
...
```
Output:
```text
...
```

#### Rank
```go
...
```
Output:
```text
...
```

#### RowNumber
```go
...
```
Output:
```text
...
```

## Literal
...
```go
...
```
Output:
```text
...
```

## Logical
### And
...
```go
...
```
Output:
```text
...
```

### Or
...
```go
...
```
Output:
```text
...
```

## Order
### Asc
...
```go
...
```
Output:
```text
...
```

### Desc
...
```go
...
```
Output:
```text
...
```

## Subquery
...
```go
...
```
Output:
```text
...
```

## Value
...
```go
...
```
Output:
```text
...
```