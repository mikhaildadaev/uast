package uast

// Публичные переменные
var (
	DialectAlloyDB          = DialectPostgreSQL
	DialectAuroraMySQL      = DialectMySQL
	DialectAuroraPostgreSQL = DialectPostgreSQL
	DialectCockroachDB      = DialectPostgreSQL
	DialectDefault          = DialectPostgreSQL
	DialectTiDB             = DialectMySQL
	DialectYugabyteDB       = DialectPostgreSQL
)

// Публичные структуры
type SupportDialect struct {
	config    *config
	name      string
	processor processor
	strateger strateger
}
