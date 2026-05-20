package uast

// Публичные переменные
var (
	// Диалект по умолчанию
	DialectDefault = DialectPostgreSQL
	// Cовместимые с MySQL
	DialectAuroraMySQL = DialectMySQL
	DialectAzureMySQL  = DialectMySQL
	DialectDoltDB      = DialectMySQL
	DialectGoogleMySQL = DialectMySQL
	DialectPlanetScale = DialectMySQL
	DialectSingleStore = DialectMySQL
	DialectTiDB        = DialectMySQL
	// Cовместимые с PostgreSQL
	DialectAlloyDB          = DialectPostgreSQL
	DialectAuroraPostgreSQL = DialectPostgreSQL
	DialectAzurePostgreSQL  = DialectPostgreSQL
	DialectCitus            = DialectPostgreSQL
	DialectCockroachDB      = DialectPostgreSQL
	DialectGooglePostgreSQL = DialectPostgreSQL
	DialectNeon             = DialectPostgreSQL
	DialectSupabase         = DialectPostgreSQL
	DialectTimescaleDB      = DialectPostgreSQL
	DialectYugabyteDB       = DialectPostgreSQL
	// Cовместимые с SQLite
	DialectCloudflareD1 = DialectSQLite
	DialectLiteFS       = DialectSQLite
	DialectTurso        = DialectSQLite
)

// Публичные структуры
type SupportDialect struct {
	config    *config
	name      string
	processor processor
	strateger strateger
}
