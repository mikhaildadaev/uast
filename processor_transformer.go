package uast

// Приватные интерфейсы
type transformer interface {
	elementTransformer
	componentTransformer
	statementTransformer
}
type elementTransformer interface {
	transformComparison() error
	transformFunction() error
}
type componentTransformer interface {
}
type statementTransformer interface {
	transformDelete(baseTransformer *baseTransformer, stmtDelete *stmtDelete) error
	transformInsert(baseTransformer *baseTransformer, stmtInsert *stmtInsert) error
	transformSelect(baseTransformer *baseTransformer, stmtSelect *stmtSelect) error
	transformUpdate(baseTransformer *baseTransformer, stmtUpdate *stmtUpdate) error
}

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
func (baseTransformer *baseTransformer) transformComparison() error {
	for _, expr := range baseTransformer.contexter.collectionComparison {
		if dialectComparison, exists := baseTransformer.config.listComparisons[expr.transformGetOperator()]; exists {
			if err := dialectComparison(baseTransformer, expr); err != nil {
				return err
			}
		}
	}
	return nil
}
func (baseTransformer *baseTransformer) transformFunction() error {
	for _, expr := range baseTransformer.contexter.collectionFunction {
		if dialectFunction, exists := baseTransformer.config.listFunctions[expr.transformGetService()]; exists {
			if err := dialectFunction(baseTransformer, expr); err != nil {
				return err
			}
		}
	}
	return nil
}
