package uast

// Публичные переменные
var (
	DialectDefault = DialectPostgreSQL
)

// Публичные структуры
type SupportDialect struct {
	config    *config
	name      string
	processor processor
	strateger strateger
}
