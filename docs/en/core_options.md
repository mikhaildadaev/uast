---
outline: deep
---

# API / Core / Options

::: info **Info**
This page covers all configuration options: `Extractor`, `Format`, `Level`, `Mode`, `Theme`. Each option is shown with a working code example and expected output.
:::

## WithExtractor/SetExtractor
Automatic context extraction. Fields from `context.Context` are added to every log, metric, and trace automatically
```go
...
```
Output:
```text
...
```

func CaseIf[OutT typeScalar](pairs ...casePair[OutT]) []casePair[OutT] {
	return pairs
}
func CaseElse[OutT typeScalar](else_ ExpressionSafe[OutT]) ExpressionSafe[OutT] {
	return else_
}
func CasePair[InT, OutT typeScalar](when ExpressionSafe[InT], then ExpressionSafe[OutT]) casePair[OutT] {
	return casePair[OutT]{
		then: then,
		when: when,
	}
}
func JsonGroup(expressions []ExpressionBase, values ...ExpressionBase) *exprJson {
	return &exprJson{
		expressions: expressions,
		values:      values,
	}
}
func JsonIndex(index int) *exprLiteral[int] {
	return &exprLiteral[int]{
		value: index,
	}
}
func JsonKey(key string) *exprLiteral[string] {
	return &exprLiteral[string]{
		value: key,
	}
}
func JsonPair(key ExpressionSafe[string], value ExpressionBase) *exprJson {
	return &exprJson{
		expressions: []ExpressionBase{key},
		operator:    uastCompositeCommaSpace,
		values:      []ExpressionBase{value},
	}
}
func JsonPath(path ...ExpressionBase) []ExpressionBase {
	return path
}
func ValueRow[T typeScalar](row ...T) []T {
	return row
}

## Array
...
```go
...
```
Output:
```text
...
```

## Binary
### BitwiseAnd
...
```go
...
```
Output:
```text
...
```

### BitwiseOr
...
```go
...
```
Output:
```text
...
```

### BitwiseXor
...
```go
...
```
Output:
```text
...
```

### Divide
...
```go
...
```
Output:
```text
...
```

### Minus
...
```go
...
```
Output:
```text
...
```

### Modulo
...
```go
...
```
Output:
```text
...
```

### Multiply
...
```go
...
```
Output:
```text
...
```

### Plus
...
```go
...
```
Output:
```text
...
```

### ShiftLeft
...
```go
...
```
Output:
```text
...
```

### ShiftRight
...
```go
...
```
Output:
```text
...
```

## Column
...
```go
...
```
Output:
```text
...
```

## Comparison
### Between
...
```go
...
```
Output:
```text
...
```

### Equal
...
```go
...
```
Output:
```text
...
```

### Exists
...
```go
...
```
Output:
```text
...
```

### Greater
...
```go
...
```
Output:
```text
...
```

### GreaterEqual
...
```go
...
```
Output:
```text
...
```

### ILike
...
```go
...
```
Output:
```text
...
```

### In
...
```go
...
```
Output:
```text
...
```

### IsNotNull
...
```go
...
```
Output:
```text
...
```

### IsNull
...
```go
...
```
Output:
```text
...
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
#### Avg
#### BitAnd
#### BitOr
#### BitXor
#### Count
#### GroupConcat
#### Max
#### Min
#### StdDev
#### Sum
#### Variance
### Analytical
#### FirstValue
#### Lag
#### LastValue
#### Lead
#### NthValue
### Condition
#### Case
#### Coalesce
#### Greatest
#### Least
#### NullIf
### Convert
#### Cast
#### CharLength
#### DateFormat
#### Degrees
#### Length
#### Position
#### Radians
### Date and time
#### CurDate
#### CurTime
#### DateAdd
#### DateDiff
#### DateSub
#### Day
#### DayName
#### Hour
#### Minute
#### Month
#### MonthName
#### Now
#### Quarter
#### Second
#### TimeAdd
#### TimeDiff
#### TimeSub
#### Week
#### Year
### Json
#### JsonArray
#### JsonArrayAgg
#### JsonContains
#### JsonExtract
#### JsonObject
#### JsonObjectAgg
#### JsonRemove
#### JsonSet
#### JsonType
### Math
#### Abs
#### ACos
#### ASin
#### ATan
#### ATan2
#### Cbrt
#### Ceil
#### Cos
#### Exp
#### Floor
#### Ln
#### Log
#### Mod
#### Pi
#### Power
#### Rand
#### Round
#### Sin
#### Sqrt
#### Tan
#### Trunc
### String
#### Concat
#### ConcatWs
#### LeftString
#### Lower
#### LPad
#### LTrim
#### Repeat
#### Replace
#### Reverse
#### RightString
#### RPad
#### RTrim
#### SubString
#### Trim
#### Upper
### Ranking
#### CumeDist
#### DenseRank
#### NTile
#### PercentRank
#### Rank
#### RowNumber
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