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