package uast

// Приватные интерфейсы
type processor interface {
	renderer
	transformer
	validator
}
