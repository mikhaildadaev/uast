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

// Приватные структуры
type baseTransformer struct {
	config    *config
	contexter *contexter
	strateger strateger
}

// Приватные типофункции
type comparisonTransform func(*baseTransformer, transformComparison) error
type functionTransform func(*baseTransformer, transformFunction) error

// Приватные конструкторы
func newTransformer(config *config, contexter *contexter, strateger strateger) *baseTransformer {
	return &baseTransformer{
		config:    config,
		contexter: contexter,
		strateger: strateger,
	}
}

// Приватные методы
func (transformer *baseTransformer) transformComparison() error {
	for _, expr := range transformer.contexter.collectionComparison {
		if dialectComparison, exists := transformer.config.listComparisons[expr.getOperator()]; exists {
			if err := dialectComparison(transformer, expr); err != nil {
				return err
			}
		}
	}
	return nil
}
func (transformer *baseTransformer) transformFunction() error {
	for _, expr := range transformer.contexter.collectionFunction {
		if dialectFunction, exists := transformer.config.listFunctions[expr.getService()]; exists {
			if err := dialectFunction(transformer, expr); err != nil {
				return err
			}
		}
	}
	return nil
}
