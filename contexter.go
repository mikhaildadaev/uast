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

import (
	"strings"
)

// Приватные структуры
type contexter struct {
	bufferByte             []byte
	bufferExpr             []ExpressionBase
	bufferQuery            strings.Builder
	bufferValue            []any
	collectionComparison   []transformComparison
	collectionFunction     []transformFunction
	currentCountComparison int
	currentCountDepth      int
	currentCountField      int
	currentCountFunction   int
	currentCountSubquery   int
	currentCountUnions     int
	currentCountWith       int
}

// Приватные конструкторы
func newContext() *contexter {
	contexter := &contexter{
		bufferByte:           make([]byte, uastSizeInitByte),
		bufferExpr:           make([]ExpressionBase, 0, uastSizeInitExpr),
		bufferQuery:          strings.Builder{},
		bufferValue:          make([]any, 0, uastSizeInitValue),
		collectionComparison: make([]transformComparison, 0, uastSizeInitComparison),
		collectionFunction:   make([]transformFunction, 0, uastSizeInitFunction),
	}
	contexter.bufferQuery.Grow(uastSizeInitQuery)
	return contexter
}

// Приватные методы
func (contexter *contexter) countComparisonPlus() error {
	contexter.currentCountComparison++
	if contexter.currentCountComparison > uastCountMaxComparison {
		return ErrExcessMaxComparison
	}
	return nil
}
func (contexter *contexter) countDepthPlus() error {
	contexter.currentCountDepth++
	if contexter.currentCountDepth > uastCountMaxDepth {
		return ErrExcessMaxDepth
	}
	return nil
}
func (contexter *contexter) countDepthMinus() error {
	contexter.currentCountDepth--
	if contexter.currentCountDepth < 0 {
		contexter.currentCountDepth = 0
	}
	return nil
}
func (contexter *contexter) countFieldPlus() error {
	contexter.currentCountField++
	if contexter.currentCountField > uastCountMaxField {
		return ErrExcessMaxField
	}
	return nil
}
func (contexter *contexter) countFunctionPlus() error {
	contexter.currentCountFunction++
	if contexter.currentCountFunction > uastCountMaxFunction {
		return ErrExcessMaxFunction
	}
	return nil
}
func (contexter *contexter) countSubqueryPlus() error {
	contexter.currentCountSubquery++
	if contexter.currentCountSubquery > uastCountMaxSubquery {
		return ErrExcessMaxSubquery
	}
	return nil
}
func (contexter *contexter) countUnionsPlus(count int) error {
	contexter.currentCountUnions += count
	if contexter.currentCountUnions > uastCountMaxUnions {
		return ErrExcessMaxUnions
	}
	return nil
}
func (contexter *contexter) countUnionsMinus(count int) error {
	contexter.currentCountUnions -= count
	if contexter.currentCountUnions < 0 {
		contexter.currentCountUnions = 0
	}
	return nil
}
func (contexter *contexter) countWithPlus(count int) error {
	contexter.currentCountWith += count
	if contexter.currentCountWith > uastCountMaxWith {
		return ErrExcessMaxWith
	}
	return nil
}
func (contexter *contexter) countWithMinus(count int) error {
	contexter.currentCountWith -= count
	if contexter.currentCountWith < 0 {
		contexter.currentCountWith = 0
	}
	return nil
}
func (contexter *contexter) prependCollectionComparison(expr transformComparison) {
	contexter.collectionComparison = append([]transformComparison{expr}, contexter.collectionComparison...)
}
func (contexter *contexter) prependCollectionFunction(expr transformFunction) {
	contexter.collectionFunction = append([]transformFunction{expr}, contexter.collectionFunction...)
}
func (contexter *contexter) resetAll() {
	contexter.resetBufferByte()
	contexter.resetBufferExpr()
	contexter.resetBufferQuery()
	contexter.resetBufferValue()
	contexter.resetCollectionComparison()
	contexter.resetCollectionFunction()
	contexter.currentCountComparison = 0
	contexter.currentCountDepth = 0
	contexter.currentCountFunction = 0
	contexter.currentCountSubquery = 0
	contexter.currentCountUnions = 0
	contexter.currentCountWith = 0
}
func (contexter *contexter) resetBufferByte() {
	for i := range contexter.bufferByte {
		contexter.bufferByte[i] = 0
	}
}
func (contexter *contexter) resetBufferExpr() {
	contexter.bufferExpr = contexter.bufferExpr[:0]
}
func (contexter *contexter) resetBufferQuery() {
	contexter.bufferQuery.Reset()
	contexter.bufferQuery.Grow(uastSizeInitQuery)
}
func (contexter *contexter) resetBufferValue() {
	contexter.bufferValue = contexter.bufferValue[:0]
}
func (contexter *contexter) resetCollectionComparison() {
	contexter.collectionComparison = contexter.collectionComparison[:0]
}
func (contexter *contexter) resetCollectionFunction() {
	contexter.collectionFunction = contexter.collectionFunction[:0]
}
