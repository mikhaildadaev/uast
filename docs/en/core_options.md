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
## Binary
### BitwiseAnd
### BitwiseOr
### BitwiseXor
### Divide
### Minus
### Modulo
### Multiply
### Plus
### ShiftLeft
### ShiftRight
## Column
## Comparison
### Between
### Equal
### Exists
### Greater
### GreaterEqual
### ILike
### In
### IsNotNull
### IsNull
### Less
### LessEqual
### Like
### NotBetween
### NotEqual
### NotExists
### NotILike
### NotIn
### NotLike
## Constant
### ConstBoolFalse
### ConstBoolTrue
### ConstFloat32One
### ConstFloat64One
### ConstIntOne
### ConstInt8One
### ConstInt16One
### ConstInt32One
### ConstInt64One
### ConstNullDefault
### ConstStringDefault
### ConstUintOne
### ConstUint8One
### ConstUint16One
### ConstUint32One
### ConstUint64One
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
## Logical
### And
### Or
## Order
### Asc
### Desc
## Subquery
## Value