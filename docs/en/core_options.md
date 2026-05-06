---
outline: deep
---

# API / Core / Options

::: info **Info**
This page covers all configuration options: `Array`, `Binary`, `Comparison`, `Constant`, `Function`, `Literal`, `Logical`, `Order`, `Subquery`, `Value`. Each option is shown with a working code example and expected output.
:::

## Array
Constructs an array expression for use in SQL queries.
```go
array := uast.Array(1, 2, 3)
```
Output:
```text
ARRAY[1, 2, 3]
```

## Binary
### BitwiseAnd
Performs a bitwise AND operation between two expressions.
```go
binary := uast.BitwiseAnd(uast.Column[int]("t", "flags"), uast.Value(0b0011))
```
Output:
```text
"t"."flags" & 3
```

### BitwiseOr
Performs a bitwise OR operation between two expressions.
```go
binary := uast.BitwiseOr(uast.Column[int]("t", "flags"), uast.Value(0b1100))
```
Output:
```text
"t"."flags" | 12
```

### BitwiseXor
Performs a bitwise XOR operation between two expressions.
```go
binary := uast.BitwiseXor(uast.Column[int]("t", "flags"), uast.Value(0b1111))
```
Output:
```text
"t"."flags" ^ 15
```

### Divide
Divides the left expression by the right expression.
```go
binary := uast.Divide(uast.Column[float64]("t", "price"), uast.Value(2.0))
```
Output:
```text
"t"."price" / 2.000000
```

### Minus
Subtracts the right expression from the left expression.
```go
binary := uast.Minus(uast.Column[int]("t", "total"), uast.Value(10))
```
Output:
```text
"t"."total" - 10
```

### Modulo
Returns the remainder of dividing the left expression by the right expression.
```go
binary := uast.Modulo(uast.Column[int]("t", "id"), uast.Value(5))
```
Output:
```text
"t"."id" % 5
```

### Multiply
Multiplies the left expression by the right expression.
```go
binary := uast.Multiply(uast.Column[float64]("t", "price"), uast.Value(1.2))
```
Output:
```text
"t"."price" * 1.200000
```

### Plus
Adds the left expression to the right expression.
```go
binary := uast.Plus(uast.Column[int]("t", "amount"), uast.Value(100))
```
Output:
```text
"t"."amount" + 100
```

### ShiftLeft
Performs a bitwise left shift on the left expression by the number of bits specified in the right expression.
```go
binary := uast.ShiftLeft(uast.Column[int]("t", "mask"), uast.Value(4))
```
Output:
```text
"t"."mask" << 4
```

### ShiftRight
Performs a bitwise right shift on the left expression by the number of bits specified in the right expression.
```go
binary := uast.ShiftRight(uast.Column[int]("t", "mask"), uast.Value(2))
```
Output:
```text
"t"."mask" >> 2
```

## Column
Creates a reference to a table column, optionally qualified with a table alias. This is the primary way to reference database columns in expressions.
```go
column := uast.Column[string]("u", "email")
```
Output:
```text
"u"."email"
```

## Comparison
### Between
Checks if the left expression falls within the range defined by `valueStart` and `valueEnd` (inclusive).
```go
comparison := uast.Between(uast.Column[int]("t", "age"), uast.Value(18), uast.Value(65))
```
Output:
```text
"t"."age" BETWEEN 18 AND 65
```

### Equal
Compares two expressions for equality (`=`).
```go
comparison := uast.Equal(uast.Column[int]("t", "status"), uast.Value("active"))
```
Output:
```text
"t"."status" = 'active'
```

### Exists
Checks if the subquery returns any rows. Returns `true` if at least one row exists.
```go
nested := uast.NewSelect(uast.Column[int]("*")).From(uast.Table("orders")).Where(
    uast.Equal(
        uast.Column[int]("orders", "user_id"),
        uast.Column[int]("users", "id"),
    ),
)
comparison := uast.Exists(uast.Subquery[int](nested))
```
Output:
```text
EXISTS (SELECT * FROM "orders" WHERE "orders"."user_id" = "users"."id")
```

### Greater
Compares if the left expression is greater than the right expression (`>`).
```go
comparison := uast.Greater(uast.Column[float64]("t", "price"), uast.Value(100.0))
```
Output:
```text
"t"."price" > 100.000000
```

### GreaterEqual
Compares if the left expression is greater than or equal to the right expression (`>=`).
```go
comparison := uast.GreaterEqual(uast.Column[int]("t", "quantity"), uast.Value(1))
```
Output:
```text
"t"."quantity" >= 1
```

### ILike
Performs a case-insensitive pattern matching comparison. The right expression should contain a pattern with `%` (any sequence) and `_` (single character) wildcards.
```go
comparison := uast.ILike(uast.Column[string]("t", "name"), uast.Value("%ivan%"))
```
Output:
```text
"t"."name" ILIKE '%ivan%'
```

### In
Checks if the left expression matches any value contained within the right expression (typically a subquery or array).
```go
comparison := uast.In(
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
comparison := uast.IsNotNull(uast.Column[string]("t", "deleted_at"))
```
Output:
```text
"t"."deleted_at" IS NOT NULL
```

### IsNull
Checks if the expression is `NULL`.
```go
comparison := uast.IsNull(uast.Column[string]("t", "deleted_at"))
```
Output:
```text
"t"."deleted_at" IS NULL
```

### Less
Compares if the left expression is less than the right expression (`<`).
```go
comparison := uast.Less(uast.Column[int]("t", "stock"), uast.Value(10))
```
Output:
```text
"t"."stock" < 10
```

### LessEqual
Compares if the left expression is less than or equal to the right expression (`<=`).
```go
comparison := uast.LessEqual(uast.Column[float64]("t", "discount"), uast.Value(0.5))
```
Output:
```text
"t"."discount" <= 0.500000
```

### Like
Performs a case-sensitive pattern matching comparison. The right expression should contain a pattern with `%` and `_` wildcards.
```go
comparison := uast.Like(uast.Column[string]("t", "code"), uast.Value("UA-%"))
```
Output:
```text
"t"."code" LIKE 'UA-%'
```

### NotBetween
Checks if the left expression falls outside the range defined by `valueStart` and `valueEnd`.
```go
comparison := uast.NotBetween(uast.Column[int]("t", "age"), uast.Value(18), uast.Value(65))
```
Output:
```text
"t"."age" NOT BETWEEN 18 AND 65
```

### NotEqual
Compares two expressions for inequality (`!=` or `<>`).
```go
comparison := uast.NotEqual(uast.Column[string]("t", "status"), uast.Value("banned"))
```
Output:
```text
"t"."status" != 'banned'
```

### NotExists
Checks if the subquery returns no rows. Returns `true` if the subquery result is empty.
```go
nested := uast.NewSelect(uast.Column[int]("*")).From(uast.Table("bans")).Where(
    uast.Equal(
        uast.Column[int]("bans", "user_id"),
        uast.Column[int]("users", "id"),
    ),
)
comparison := uast.NotExists(uast.Subquery[int](nested))
```
Output:
```text
NOT EXISTS (SELECT * FROM "bans" WHERE "bans"."user_id" = "users"."id")
```

### NotILike
Performs a negated case-insensitive pattern matching comparison.
```go
comparison := uast.NotILike(uast.Column[string]("t", "email"), uast.Value("%@spam.com"))
```
Output:
```text
"t"."email" NOT ILIKE '%@spam.com'
```

### NotIn
Checks if the left expression does not match any value contained within the right expression.
```go
comparison := uast.NotIn(
    uast.Column[int]("t", "status_id"),
    uast.Subquery[int](
        uast.NewSelect(uast.Column[int]("s", "id")).
            From(uast.Table("statuses")).
            Where(uast.Equal(uast.Column[string]("s", "archived"), uast.Value(true))),
    ),
)
```
Output:
```text
"t"."status_id" NOT IN (SELECT "s"."id" FROM "statuses" WHERE "s"."archived" = TRUE)
```

### NotLike
Performs a negated case-sensitive pattern matching comparison.
```go
comparison := uast.NotLike(uast.Column[string]("t", "phone"), uast.Value("+7-000-%"))
```
Output:
```text
"t"."phone" NOT LIKE '+7-000-%'
```

## Constant
### ConstBoolFalse
Returns a constant boolean `FALSE` expression. Useful as a default placeholder or for constructing logical expressions without value binding.
```go
constant := uast.ConstBoolFalse()
```
Output:
```text
FALSE
```

### ConstBoolTrue
Returns a constant boolean `TRUE` expression.
```go
constant := uast.ConstBoolTrue()
```
Output:
```text
TRUE
```

### ConstFloat32One
Returns a constant `float32` value of `1.0`. Optimized for internal comparisons and arithmetic where a unit value is required without binding a placeholder.
```go
constant := uast.ConstFloat32One()
```
Output:
```text
1.000000
```

### ConstFloat64One
Returns a constant `float64` value of `1.0`.
```go
constant := uast.ConstFloat64One()
```
Output:
```text
1.000000
```

### ConstIntOne
Returns a constant `int` value of `1`.
```go
constant := uast.ConstIntOne()
```
Output:
```text
1
```

### ConstInt8One
Returns a constant `int8` value of `1`.
```go
constant := uast.ConstInt8One()
```
Output:
```text
1
```

### ConstInt16One
Returns a constant `int16` value of `1`.
```go
constant := uast.ConstInt16One()
```
Output:
```text
1
```

### ConstInt32One
Returns a constant `int32` value of `1`.
```go
constant := uast.ConstInt32One()
```
Output:
```text
1
```

### ConstInt64One
Returns a constant `int64` value of `1`.
```go
constant := uast.ConstInt64One()
```
Output:
```text
1
```

### ConstNullDefault
Returns a typed `NULL` constant. The type is determined by the generic parameter `T`. This allows `NULL` to be used in typed contexts without explicit casting.
```go
constant := uast.ConstNullDefault[string]()
```
Output:
```text
NULL
```

### ConstStringDefault
Returns a constant empty string expression (` `). Used when an empty string needs to be embedded directly in the SQL without a placeholder.
```go
constant := uast.ConstStringDefault()
```
Output:
```text
 
```

### ConstUintOne
Returns a constant `uint` value of `1`.
```go
constant := uast.ConstUintOne()
```
Output:
```text
1
```

### ConstUint8One
Returns a constant `uint8` value of `1`.
```go
constant := uast.ConstUint8One()
```
Output:
```text
1
```

### ConstUint16One
Returns a constant `uint16` value of `1`.
```go
constant := uast.ConstUint16One()
```
Output:
```text
1
```

### ConstUint32One
Returns a constant `uint32` value of `1`.
```go
constant := uast.ConstUint32One()
```
Output:
```text
1
```

### ConstUint64One
Returns a constant `uint64` value of `1`.
```go
constant := uast.ConstUint64One()
```
Output:
```text
1
```

## Function
### Aggregate
#### Avg
Returns the average (arithmetic mean) of all non-NULL values in the expression. If `distinct` is `true`, the average is calculated over distinct values only.
```go
function := uast.Avg(uast.Column[float64]("o", "total_price"), false)
```
Output:
```text
AVG("o"."total_price")
```

#### BitAnd
Returns the bitwise AND of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitAnd(uast.Column[int]("t", "permissions"), false)
```
Output:
```text
BIT_AND("t"."permissions")
```

#### BitOr
Returns the bitwise OR of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitOr(uast.Column[int]("t", "flags"), true)
```
Output:
```text
BIT_OR(DISTINCT "t"."flags")
```

#### BitXor
Returns the bitwise XOR of all bits in the expression. Only meaningful for integer types.
```go
function := uast.BitXor(uast.Column[int]("t", "checksum"), false)
```
Output:
```text
BIT_XOR("t"."checksum")
```

#### Count
Returns the number of rows matching the query, or the number of non-NULL values if an expression is provided. When `distinct` is `true`, counts only distinct values.
```go
function := uast.Count(uast.Column[int]("*"), false)
functionWithDistinct := uast.Count(uast.Column[string]("u", "email"), true)
```
Output:
```text
COUNT(*)
COUNT(DISTINCT "u"."email")
```

#### GroupConcat
Concatenates values from a group into a single string, separated by a default delimiter (typically a comma). The `distinct` flag removes duplicates before concatenation.
```go
function := uast.GroupConcat(uast.Column[string]("t", "tag"), true)
```
Output:
```text
GROUP_CONCAT(DISTINCT "t"."tag")
```

#### Max
Returns the maximum value of the expression across all rows in the group.
```go
function := uast.Max(uast.Column[float64]("o", "amount"), false)
```
Output:
```text
MAX("o"."amount")
```

#### Min
Returns the minimum value of the expression across all rows in the group.
```go
function := uast.Min(uast.Column[time.Time]("o", "created_at"), false)
```
Output:
```text
MIN("o"."created_at")
```

#### StdDev
Returns the population standard deviation of the expression. Supported mainly by MySQL dialect; check PostgreSQL compatibility.
```go
function := uast.StdDev(uast.Column[float64]("t", "score"), false)
```
Output:
```text
STDDEV("t"."score")
```

#### Sum
Returns the sum of all values in the expression. If `distinct` is `true`, sums only distinct values.
```go
function := uast.Sum(uast.Column[float64]("o", "tax"), false)
```
Output:
```text
SUM("o"."tax")
```

#### Variance
Returns the population variance of the expression. Supported mainly by MySQL dialect; check PostgreSQL compatibility.
```go
function := uast.Variance(uast.Column[float64]("t", "latency"), false)
```
Output:
```text
VARIANCE("t"."latency")
```

### Analytical
#### FirstValue
Returns the value of the expression from the first row of the window frame. Requires an `OVER` clause with window specification.
```go
function := uast.FirstValue(uast.Column[string]("t", "event")).Over(
    uast.PartitionBy(uast.Column[int]("t", "user_id")),
    uast.OrderBy(uast.Asc(uast.Column[time.Time]("t", "created_at"))),
)
```
Output:
```text
FIRST_VALUE("t"."event") OVER (PARTITION BY "t"."user_id" ORDER BY "t"."created_at" ASC)
```

#### Lag
Returns the value of the expression from a row that is `offset` rows before the current row within the partition.
```go
function := uast.Lag(uast.Column[float64]("t", "price"), 1).Over(
    uast.PartitionBy(uast.Column[int]("t", "symbol_id")),
    uast.OrderBy(uast.Asc(uast.Column[time.Time]("t", "tick_time"))),
)
```
Output:
```text
LAG("t"."price", 1) OVER (PARTITION BY "t"."symbol_id" ORDER BY "t"."tick_time" ASC)
```

#### LastValue
Returns the value of the expression from the last row of the window frame.
```go
function := uast.LastValue(uast.Column[string]("t", "status")).Over(
    uast.PartitionBy(uast.Column[int]("t", "order_id")),
    uast.OrderBy(uast.Asc(uast.Column[time.Time]("t", "updated_at"))),
    uast.Frame(uast.RowsCurrent(), uast.UnboundedFollowing()),
)
```
Output:
```text
LAST_VALUE("t"."status") OVER (PARTITION BY "t"."order_id" ORDER BY "t"."updated_at" ASC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)
```

#### Lead
Returns the value of the expression from a row that is `offset` rows after the current row within the partition.
```go
function := uast.Lead(uast.Column[float64]("t", "temperature"), 1).Over(
    uast.PartitionBy(uast.Column[int]("t", "sensor_id")),
    uast.OrderBy(uast.Asc(uast.Column[time.Time]("t", "measured_at"))),
)
```
Output:
```text
LEAD("t"."temperature", 1) OVER (PARTITION BY "t"."sensor_id" ORDER BY "t"."measured_at" ASC)
```

#### NthValue
Returns the value of the expression from the `n-th` row of the window frame (1-based).
```go
function := uast.NthValue(uast.Column[string]("t", "log_entry"), 3).Over(
    uast.PartitionBy(uast.Column[int]("t", "batch_id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "severity"))),
)
```
Output:
```text
NTH_VALUE("t"."log_entry", 3) OVER (PARTITION BY "t"."batch_id" ORDER BY "t"."severity" DESC)
```

### Condition
#### Case
Evaluates a list of `WHEN`-`THEN` pairs and returns the `THEN` expression for the first true WHEN. If no condition is true, returns the `ELSE` expression if provided, or `NULL`.
```go
pairs := uast.CaseIf(
    uast.CasePair(
        uast.Greater(uast.Column[int]("t", "score"), uast.Value(90)),
        uast.Value("A"),
    ),
    uast.CasePair(
        uast.Greater(uast.Column[int]("t", "score"), uast.Value(75)),
        uast.Value("B"),
    ),
)
elseExpr := uast.CaseElse(uast.Value("C"))
function := uast.Case(pairs, elseExpr)
```
Output:
```text
CASE WHEN "t"."score" > 90 THEN 'A' WHEN "t"."score" > 75 THEN 'B' ELSE 'C' END
```

#### Coalesce
Returns the first non-NULL expression from the provided list. Useful for providing fallback values.
```go
function := uast.Coalesce(
    uast.Column[string]("u", "nickname"),
    uast.Column[string]("u", "username"),
    uast.Value("Anonymous"),
)
```
Output:
```text
COALESCE("u"."nickname", "u"."username", 'Anonymous')
```

#### Greatest
Returns the largest value from the provided list of expressions.
```go
function := uast.Greatest(
    uast.Column[int]("t", "score_a"),
    uast.Column[int]("t", "score_b"),
    uast.Column[int]("t", "score_c"),
)
```
Output:
```text
GREATEST("t"."score_a", "t"."score_b", "t"."score_c")
```

#### Least
Returns the smallest value from the provided list of expressions.
```go
function := uast.Least(
    uast.Column[time.Time]("t", "start_a"),
    uast.Column[time.Time]("t", "start_b"),
)
```
Output:
```text
LEAST("t"."start_a", "t"."start_b")
```

#### NullIf
Returns `NULL` if the two expressions are equal; otherwise returns the first expression.
```go
function := uast.NullIf(uast.Column[int]("t", "value"), uast.Value(0))
```
Output:
```text
NULLIF("t"."value", 0)
```

### Convert
#### Cast
Converts an expression to a specified data type.
```go
function := uast.Cast(uast.Column[string]("t", "json_str"), uast.TypeJSON)
```
Output:
```text
CAST("t"."json_str" AS JSON)
```

#### CharLength
Returns the number of characters in a string expression.
```go
function := uast.CharLength(uast.Column[string]("t", "body"))
```
Output:
```text
CHAR_LENGTH("t"."body")
```

#### DateFormat
Formats a datetime expression according to a specified format mask.
```go
function := uast.DateFormat(uast.Column[time.Time]("t", "event_date"), uast.Value("%Y-%m-%d"))
```
Output:
```text
DATE_FORMAT("t"."event_date", '%Y-%m-%d')
```

#### Degrees
Converts an angle from radians to degrees.
```go
function := uast.Degrees(uast.Column[float64]("t", "angle_rad"))
```
Output:
```text
DEGREES("t"."angle_rad")
```

#### Length
Returns the byte length of a string expression.
```go
function := uast.Length(uast.Column[string]("t", "payload"))
```
Output:
```text
LENGTH("t"."payload")
```

#### Position
Returns the starting position of the first occurrence of a substring within a string.
```go
function := uast.Position(uast.Column[string]("t", "url"), uast.Value("https://"))
```
Output:
```text
POSITION('https://' IN "t"."url")
```

#### Radians
Converts an angle from degrees to radians.
```go
function := uast.Radians(uast.Column[float64]("t", "angle_deg"))
```
Output:
```text
RADIANS("t"."angle_deg")
```

### Date and time
#### CurDate
Returns the current date (without time).
```go
function := uast.CurDate()
```
Output:
```text
CURDATE()
```

#### CurTime
Returns the current time (without date).
```go
function := uast.CurTime()
```
Output:
```text
CURTIME()
```

#### DateAdd
Adds a time/date interval to a datetime expression and returns the resulting datetime.
```go
function := uast.DateAdd(uast.Column[time.Time]("t", "start_date"), uast.Value("1 YEAR"))
```
Output:
```text
DATE_ADD("t"."start_date", '1 YEAR')
```

#### DateDiff
Returns the difference in days between two datetime expressions (`datetimeEnd` - `datetimeStart`).
```go
function := uast.DateDiff(
    uast.Column[time.Time]("t", "closed_at"),
    uast.Column[time.Time]("t", "opened_at"),
)
```
Output:
```text
DATEDIFF("t"."closed_at", "t"."opened_at")
```

#### DateSub
Subtracts a time/date interval from a datetime expression and returns the resulting datetime.
```go
function := uast.DateSub(uast.Column[time.Time]("t", "due_date"), uast.Value("3 MONTH"))
```
Output:
```text
DATE_SUB("t"."due_date", '3 MONTH')
```

#### Day
Extracts the day of the month (1–31) from a datetime expression.
```go
function := uast.Day(uast.Column[time.Time]("t", "created_at"))
```
Output:
```text
DAY("t"."created_at")
```

#### DayName
Returns the name of the weekday (e.g., 'Monday', 'Tuesday') for a given datetime expression.
```go
function := uast.DayName(uast.Column[time.Time]("t", "event_date"))
```
Output:
```text
DAYNAME("t"."event_date")
```

#### Hour
Extracts the hour (0–23) from a datetime expression.
```go
function := uast.Hour(uast.Column[time.Time]("t", "log_time"))
```
Output:
```text
HOUR("t"."log_time")
```

#### Minute
Extracts the minute (0–59) from a datetime expression.
```go
function := uast.Minute(uast.Column[time.Time]("t", "log_time"))
```
Output:
```text
MINUTE("t"."log_time")
```

#### Month
Extracts the month (1–12) from a datetime expression.
```go
function := uast.Month(uast.Column[time.Time]("t", "birth_date"))
```
Output:
```text
MONTH("t"."birth_date")
```

#### MonthName
Returns the name of the month (e.g., 'January', 'February') for a given datetime expression.
```go
function := uast.MonthName(uast.Column[time.Time]("t", "holiday"))
```
Output:
```text
MONTHNAME("t"."holiday")
```

#### Now
Returns the current date and time.
```go
function := uast.Now()
```
Output:
```text
NOW()
```

#### Quarter
Extracts the quarter (1–4) from a datetime expression.
```go
function := uast.Quarter(uast.Column[time.Time]("t", "fiscal_date"))
```
Output:
```text
QUARTER("t"."fiscal_date")
```

#### Second
Extracts the second (0–59) from a datetime expression.
```go
function := uast.Second(uast.Column[time.Time]("t", "response_time"))
```
Output:
```text
SECOND("t"."response_time")
```

#### TimeAdd
Adds a time interval to a time/datetime expression and returns the resulting time.
```go
function := uast.TimeAdd(uast.Column[time.Time]("t", "shift_start"), uast.Value("8 HOUR"))
```
Output:
```text
TIME_ADD("t"."shift_start", '8 HOUR')
```

#### TimeDiff
Returns the difference between two time/datetime expressions (`timeEnd` - `timeStart`).
```go
function := uast.TimeDiff(
    uast.Column[time.Time]("t", "check_out"),
    uast.Column[time.Time]("t", "check_in"),
)
```
Output:
```text
TIMEDIFF("t"."check_out", "t"."check_in")
```

#### TimeSub
Subtracts a time interval from a time/datetime expression and returns the resulting time.
```go
function := uast.TimeSub(uast.Column[time.Time]("t", "lunch_end"), uast.Value("30 MINUTE"))
```
Output:
```text
TIME_SUB("t"."lunch_end", '30 MINUTE')
```

#### Week
Extracts the week number (1–53) from a datetime expression.
```go
function := uast.Week(uast.Column[time.Time]("t", "ship_date"))
```
Output:
```text
WEEK("t"."ship_date")
```

#### Year
Extracts the year from a datetime expression.
```go
function := uast.Year(uast.Column[time.Time]("t", "founded"))
```
Output:
```text
YEAR("t"."founded")
```

### Json
#### JsonArray
Creates a JSON array from the given expression and optional additional values.
```go
function := uast.JsonArray(uast.Column[int]("t", "tag_id"), uast.Value("urgent"))
```
Output:
```text
JSON_ARRAY("t"."tag_id", 'urgent')
```

#### JsonArrayAgg
Aggregates values from a group into a JSON array.
```go
function := uast.JsonArrayAgg(uast.Column[string]("t", "label"))
```
Output:
```text
JSON_ARRAYAGG("t"."label")
```

#### JsonContains
Checks whether a JSON document contains a specified value.
```go
function := uast.JsonContains(uast.Column[string]("t", "metadata"), uast.Value(`"key1"`))
```
Output:
```text
JSON_CONTAINS("t"."metadata", '"key1"')
```

#### JsonExtract
Extracts a value from a JSON document at the specified path. The `json` parameter is built with JsonPath and optional `JsonKey`/`JsonIndex`.
```go
path := uast.JsonGroup(nil, uast.JsonKey("address"), uast.JsonKey("city"))
function := uast.JsonExtract(uast.Column[string]("t", "profile"), uast.JsonPair(uast.Literal("$.address.city"), nil), uast.TypeString)
```
Output:
```text
JSON_EXTRACT("t"."profile", '$.address.city') AS STRING
```

#### JsonObject
Builds a JSON object from key-value pairs.
```go
function := uast.JsonObject(
    uast.JsonPair(uast.Value("name"), uast.Column[string]("u", "full_name")),
    uast.JsonPair(uast.Value("age"), uast.Column[int]("u", "age")),
)
```
Output:
```text
JSON_OBJECT('name', "u"."full_name", 'age', "u"."age")
```

#### JsonObjectAgg
Aggregates key-value pairs from a group into a single JSON object.
```go
function := uast.JsonObjectAgg(
    uast.Column[string]("t", "config_key"),
    uast.Column[string]("t", "config_value"),
)
```
Output:
```text
JSON_OBJECTAGG("t"."config_key", "t"."config_value")
```

#### JsonRemove
Removes a value from a JSON document at the specified path(s).
```go
paths := uast.JsonGroup(nil, uast.JsonKey("temp_field"))
function := uast.JsonRemove(uast.Column[string]("t", "data"), paths)
```
Output:
```text
JSON_REMOVE("t"."data", '$.temp_field')
```

#### JsonSet
Sets a value in a JSON document at the specified path(s). Creates the path if it does not exist.
```go
function := uast.JsonSet(
    uast.Column[string]("t", "settings"),
    uast.JsonPair(uast.Value("$.theme"), uast.Value("dark")),
)
```
Output:
```text
JSON_SET("t"."settings", '$.theme', 'dark')
```

#### JsonType
Returns the JSON type of a JSON value (e.g., 'OBJECT', 'ARRAY', 'STRING', 'INTEGER', 'NULL').
```go
function := uast.JsonType(uast.Column[string]("t", "attributes"))
```
Output:
```text
JSON_TYPE("t"."attributes")
```

### Math
#### Abs
Returns the absolute (non-negative) value of a numeric expression.
```go
function := uast.Abs(uast.Column[float64]("t", "deviation"))
```
Output:
```text
ABS("t"."deviation")
```

#### ACos
Returns the arc cosine (inverse cosine) of the expression, in radians.
```go
function := uast.ACos(uast.Column[float64]("t", "cosine_val"))
```
Output:
```text
ACOS("t"."cosine_val")
```

#### ASin
Returns the arc sine (inverse sine) of the expression, in radians.
```go
function := uast.ASin(uast.Column[float64]("t", "sine_val"))
```
Output:
```text
ASIN("t"."sine_val")
```

#### ATan
Returns the arc tangent (inverse tangent) of the expression, in radians.
```go
function := uast.ATan(uast.Column[float64]("t", "slope"))
```
Output:
```text
ATAN("t"."slope")
```

#### ATan2
Returns the arc tangent of the quotient of its two arguments (`y`/`x`), using their signs to determine the quadrant.
```go
function := uast.ATan2(uast.Column[float64]("t", "y"), uast.Column[float64]("t", "x"))
```
Output:
```text
ATAN2("t"."y", "t"."x")
```

#### Cbrt
Returns the cube root of a numeric expression.
```go
function := uast.Cbrt(uast.Column[float64]("t", "volume"))
```
Output:
```text
CBRT("t"."volume")
```

#### Ceil
Returns the smallest integer value not less than the argument (rounds up).
```go
function := uast.Ceil(uast.Column[float64]("t", "rating"))
```
Output:
```text
CEIL("t"."rating")
```

#### Cos
Returns the cosine of the expression, where the expression is in radians.
```go
function := uast.Cos(uast.Column[float64]("t", "angle"))
```
Output:
```text
COS("t"."angle")
```

#### Exp
Returns `e` (Euler's number, ~2.71828) raised to the power of the expression.
```go
function := uast.Exp(uast.Column[float64]("t", "log_odds"))
```
Output:
```text
EXP("t"."log_odds")
```

#### Floor
Returns the largest integer value not greater than the argument (rounds down).
```go
function := uast.Floor(uast.Column[float64]("t", "amount"))
```
Output:
```text
FLOOR("t"."amount")
```

#### Ln
Returns the natural logarithm (base `e`) of the expression.
```go
function := uast.Ln(uast.Column[float64]("t", "ratio"))
```
Output:
```text
LN("t"."ratio")
```

#### Log
Returns the logarithm of the expression to the specified base.
```go
function := uast.Log(uast.Column[float64]("t", "value"), uast.Value(10.0))
```
Output:
```text
LOG(10.000000, "t"."value")
```

#### Mod
Returns the remainder (modulo) of the division of the first expression by the second.
```go
function := uast.Mod(uast.Column[int]("t", "serial"), uast.Value(16))
```
Output:
```text
MOD("t"."serial", 16)
```

#### Pi
Returns the mathematical constant `p` (~3.14159).
```go
function := uast.Pi()
```
Output:
```text
PI()
```

#### Power
Returns the expression raised to the power of the exponent.
```go
function := uast.Power(uast.Column[float64]("t", "base"), uast.Column[float64]("t", "exponent"))
```
Output:
```text
POWER("t"."base", "t"."exponent")
```

#### Rand
Returns a random floating-point value in the range [0, 1].
```go
function := uast.Rand()
```
Output:
```text
RAND()
```

#### Round
Rounds the expression to the specified number of decimal places.
```go
function := uast.Round(uast.Column[float64]("t", "price"), uast.Value(2))
```
Output:
```text
ROUND("t"."price", 2)
```

#### Sin
Returns the sine of the expression, where the expression is in radians.
```go
function := uast.Sin(uast.Column[float64]("t", "phase"))
```
Output:
```text
SIN("t"."phase")
```

#### Sqrt
Returns the square root of the expression.
```go
function := uast.Sqrt(uast.Column[float64]("t", "area"))
```
Output:
```text
SQRT("t"."area")
```

#### Tan
Returns the tangent of the expression, where the expression is in radians.
```go
function := uast.Tan(uast.Column[float64]("t", "incidence"))
```
Output:
```text
TAN("t"."incidence")
```

#### Trunc
Truncates the numeric expression to the specified number of decimal places (without rounding).
```go
function := uast.Trunc(uast.Column[float64]("t", "measurement"), uast.Value(3))
```
Output:
```text
TRUNC("t"."measurement", 3)
```

### String
#### Concat
Concatenates two or more string expressions into a single string. `NULL` arguments are treated as empty strings in most dialects.
```go
function := uast.Concat(
    uast.Column[string]("u", "first_name"),
    uast.Value(" "),
    uast.Column[string]("u", "last_name"),
)
```
Output:
```text
CONCAT("u"."first_name", ' ', "u"."last_name")
```

#### ConcatWs
Concatenates two or more string expressions with a specified separator between them. Skips `NULL` arguments.
```go
function := uast.ConcatWs(
    uast.Value(", "),
    uast.Column[string]("a", "city"),
    uast.Column[string]("a", "country"),
)
```
Output:
```text
CONCAT_WS(', ', "a"."city", "a"."country")
```

#### LeftString
Returns the leftmost `count` characters from a string expression.
```go
function := uast.LeftString(uast.Column[string]("t", "sku"), uast.Value(3))
```
Output:
```text
LEFT("t"."sku", 3)
```

#### Lower
Converts a string expression to lowercase.
```go
function := uast.Lower(uast.Column[string]("t", "email"))
```
Output:
```text
LOWER("t"."email")
```

#### LPad
Left-pads a string expression with the specified separator to a total length of `count` characters.
```go
function := uast.LPad(uast.Column[string]("t", "code"), uast.Value(10), uast.Value("0"))
```
Output:
```text
LPAD("t"."code", 10, '0')
```

#### LTrim
Removes leading spaces from a string expression.
```go
function := uast.LTrim(uast.Column[string]("t", "raw_input"))
```
Output:
```text
LTRIM("t"."raw_input")
```

#### Repeat
Repeats a string expression `count` times.
```go
function := uast.Repeat(uast.Value("-"), uast.Value(5))
```
Output:
```text
REPEAT('-', 5)
```

#### Replace
Replaces all occurrences of a substring in a string with a new substring.
```go
function := uast.Replace(uast.Column[string]("t", "url"), uast.Value("http://"), uast.Value("https://"))
```
Output:
```text
REPLACE("t"."url", 'http://', 'https://')
```

#### Reverse
Reverses the characters in a string expression.
```go
function := uast.Reverse(uast.Column[string]("t", "dna_sequence"))
```
Output:
```text
REVERSE("t"."dna_sequence")
```

#### RightString
Returns the rightmost `count` characters from a string expression.
```go
function := uast.RightString(uast.Column[string]("t", "filename"), uast.Value(4))
```
Output:
```text
RIGHT("t"."filename", 4)
```

#### RPad
Right-pads a string expression with the specified separator to a total length of `count` characters.
```go
function := uast.RPad(uast.Column[string]("t", "title"), uast.Value(30), uast.Value("."))
```
Output:
```text
RPAD("t"."title", 30, '.')
```

#### RTrim
Removes trailing spaces from a string expression.
```go
function := uast.RTrim(uast.Column[string]("t", "comment"))
```
Output:
```text
RTRIM("t"."comment")
```

#### SubString
Extracts a substring from a string expression starting at `startPos` (1-based) for `lengthStr` characters.
```go
function := uast.SubString(uast.Column[string]("t", "isbn"), uast.Value(1), uast.Value(3))
```
Output:
```text
SUBSTRING("t"."isbn", 1, 3)
```

#### Trim
Removes both leading and trailing spaces from a string expression.
```go
function := uast.Trim(uast.Column[string]("t", "username"))
```
Output:
```text
TRIM("t"."username")
```

#### Upper
Converts a string expression to uppercase.
```go
function := uast.Upper(uast.Column[string]("t", "country_code"))
```
Output:
```text
UPPER("t"."country_code")
```

### Ranking
#### CumeDist
Returns the cumulative distribution of a value within a partition (the ratio of rows that come before or are peers with the current row). Must be used with an `OVER` clause.
```go
function := uast.CumeDist().Over(
    uast.OrderBy(uast.Desc(uast.Column[float64]("t", "score"))),
)
```
Output:
```text
CUME_DIST() OVER (ORDER BY "t"."score" DESC)
```

#### DenseRank
Returns the rank of a row without gaps. Rows with equal values receive the same rank, and the next rank is the immediate next integer. Requires `OVER`.
```go
function := uast.DenseRank().Over(
    uast.PartitionBy(uast.Column[int]("t", "category_id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "points"))),
)
```
Output:
```text
DENSE_RANK() OVER (PARTITION BY "t"."category_id" ORDER BY "t"."points" DESC)
```

#### NTile
Divides the rows within a partition into `n` approximately equal groups and returns the group number (1 through `n`) for each row.
```go
function := uast.NTile(4).Over(
    uast.OrderBy(uast.Asc(uast.Column[float64]("t", "gpa"))),
)
```
Output:
```text
NTILE(4) OVER (ORDER BY "t"."gpa" ASC)
```

#### PercentRank
Returns the percentile rank of a row within a partition (range 0 to 1). Rank of first row is always 0. Requires `OVER`.
```go
function := uast.PercentRank().Over(
    uast.OrderBy(uast.Asc(uast.Column[float64]("t", "latency"))),
)
```
Output:
```text
PERCENT_RANK() OVER (ORDER BY "t"."latency" ASC)
```

#### Rank
Returns the rank of a row with gaps. Equal values receive the same rank, and the next distinct value skips ahead. Requires `OVER`.
```go
function := uast.Rank().Over(
    uast.PartitionBy(uast.Column[int]("t", "league_id")),
    uast.OrderBy(uast.Desc(uast.Column[int]("t", "wins"))),
)
```
Output:
```text
RANK() OVER (PARTITION BY "t"."league_id" ORDER BY "t"."wins" DESC)
```

#### RowNumber
Assigns a unique sequential integer to each row within the partition, starting from 1. Order determines the numbering sequence.
```go
function := uast.RowNumber().Over(
    uast.PartitionBy(uast.Column[int]("t", "session_id")),
    uast.OrderBy(uast.Asc(uast.Column[time.Time]("t", "event_time"))),
)
```
Output:
```text
ROW_NUMBER() OVER (PARTITION BY "t"."session_id" ORDER BY "t"."event_time" ASC)
```

## Literal
Embeds a raw literal value directly into the generated SQL string (not parameterized). Use with caution — values are written as-is. Prefer `Value` for user-supplied data.
```go
literal := uast.Literal("CURRENT_TIMESTAMP")
```
Output:
```text
CURRENT_TIMESTAMP
```

## Logical
### And
Combines multiple conditions with a logical `AND`. All conditions must be true for the combined expression to be true.
```go
logical := uast.And(
    uast.Equal(uast.Column[string]("t", "status"), uast.Value("active")),
    uast.Greater(uast.Column[int]("t", "login_count"), uast.Value(0)),
)
```
Output:
```text
("t"."status" = 'active' AND "t"."login_count" > 0)
```

### Or
Combines multiple conditions with a logical `OR`. At least one condition must be true for the combined expression to be true.
```go
logical := uast.Or(
    uast.IsNull(uast.Column[string]("t", "closed_at")),
    uast.Greater(uast.Column[time.Time]("t", "closed_at"), uast.Value("2026-01-01")),
)
```
Output:
```text
("t"."closed_at" IS NULL OR "t"."closed_at" > '2026-01-01')
```

## Order
### Asc
Specifies ascending order for an `ORDER BY` clause or window function ordering.
```go
order := uast.Asc(uast.Column[string]("u", "last_name"))
```
Output:
```text
"u"."last_name" ASC
```

### Desc
Specifies descending order for an `ORDER BY` clause or window function ordering.
```go
order := uast.Desc(uast.Column[int]("o", "total"))
```
Output:
```text
"o"."total" DESC
```

## Subquery
Wraps a `SELECT` statement as a typed expression that can be used in comparisons (`In`, `Exists`, `Equal`, etc.) or as a column in a `SELECT` clause. The generic parameter `T` specifies the scalar type of the single column returned by the subquery.
```go
subquery := uast.Subquery[int](
    uast.NewSelect(uast.Column[int]("c", "id")).
        From(uast.Table("categories")).
        Where(uast.Equal(uast.Column[string]("c", "slug"), uast.Value("books"))),
)
```
Output:
```text
(SELECT "c"."id" FROM "categories" WHERE "c"."slug" = 'books')
```

## Value
Wraps a Go value as a parameterized expression. The value is NOT inserted into the SQL string directly — instead, a placeholder (`$1`, `?`, etc.) is generated and the value is appended to the arguments slice returned by `Build()`. This is the safe way to pass user-supplied data.
```go
value := uast.Value("hello@world.com")
```
Output:
```text
$1
```