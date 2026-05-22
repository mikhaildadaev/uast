package uast

import (
	"strings"
)

// Приватные структуры
type contexter struct {
	bufferByte           []byte
	bufferExpr           []ExpressionBase
	bufferQuery          strings.Builder
	bufferValue          []any
	collectionComparison []transformComparison
	collectionFunction   []transformFunction
	countMaxDepth        int
	countMaxComparison   int
	countMaxFunction     int
	countMaxSubquery     int
	countMaxUnions       int
	countMaxWith         int
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
	contexter.countMaxDepth = 0
	contexter.countMaxComparison = 0
	contexter.countMaxFunction = 0
	contexter.countMaxSubquery = 0
	contexter.countMaxUnions = 0
	contexter.countMaxWith = 0
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
