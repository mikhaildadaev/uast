package uast

// Публичные структуры
type WindowFrame struct {
	Type  string
	Start string
	End   string
}
type WindowSpec struct {
	partition []ExpressionBase
	order     []*exprOrderBy
	frame     *WindowFrame
}
type WindowOption func(*WindowSpec)

// Публичные методы
func GroupsBetween(start, end string) WindowOption {
	return func(windowSpec *WindowSpec) {
		windowSpec.frame = &WindowFrame{
			Type:  "GROUPS",
			Start: start,
			End:   end,
		}
	}
}
func OrderBy(orders ...*exprOrderBy) WindowOption {
	return func(windowSpec *WindowSpec) {
		windowSpec.order = orders
	}
}
func PartitionBy(exprs ...ExpressionBase) WindowOption {
	return func(windowSpec *WindowSpec) {
		windowSpec.partition = exprs
	}
}
func RangeBetween(start, end string) WindowOption {
	return func(windowSpec *WindowSpec) {
		windowSpec.frame = &WindowFrame{
			Type:  "RANGE",
			Start: start,
			End:   end,
		}
	}
}
func RowsBetween(start, end string) WindowOption {
	return func(windowSpec *WindowSpec) {
		windowSpec.frame = &WindowFrame{
			Type:  "ROWS",
			Start: start,
			End:   end,
		}
	}
}
