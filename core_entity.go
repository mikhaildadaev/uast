package uast

// Публичный интерфейс
type EntityTarget interface {
	isEntityTarget()
}

// Публичные конструкторы
func NewIndex(name string) *targetIndex {
	return &targetIndex{
		name: name,
	}
}
func NewSchema(name string) *targetSchema {
	return &targetSchema{
		name: name,
	}
}
func NewView(name, alias string) *targetView {
	return &targetView{
		TableSource: NewTable(name, alias),
	}
}

// Публичные методы
func (target *targetIndex) isEntityTarget() {}
func (target *targetIndex) Name() string {
	return target.name
}
func (target *targetSchema) isEntityTarget() {}
func (target *targetSchema) Name() string {
	return target.name
}
func (target *targetTable) isEntityTarget() {}
func (target *targetView) isEntityTarget()  {}

// Приватные структуры
type targetIndex struct {
	name string
}
type targetSchema struct {
	name string
}
type targetTable = TableSource
type targetView struct {
	*TableSource
}
