// UAST (Universal Abstract Syntax Tree)
// Copyright (C) 2026 Mikhail Dadaev
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package uast

// Публичные структуры
type WindowFrame struct {
	Type  string
	Start string
	End   string
}
type WindowSpec struct {
	partition []ExpressionBase
	order     []*clauseOrderBy
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
func OrderBy(orders ...*clauseOrderBy) WindowOption {
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
