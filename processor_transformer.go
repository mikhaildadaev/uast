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
