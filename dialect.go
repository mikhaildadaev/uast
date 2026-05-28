package uast

// Публичные переменные
var (
	// Диалект по умолчанию
	DialectDefault = DialectPostgreSQL
	// Cовместимые с MariaDB
	DialectDoltDB      = DialectMariaDB
	DialectSingleStore = DialectMariaDB
	// Cовместимые с MsSQL
	DialectAmazonRDS = DialectMsSQL
	DialectAzureSQL  = DialectMsSQL
	DialectSynapse   = DialectMsSQL
	// Cовместимые с MySQL
	DialectAuroraMySQL = DialectMySQL
	DialectAzureMySQL  = DialectMySQL
	DialectGoogleMySQL = DialectMySQL
	DialectOceanBase   = DialectMySQL
	DialectPlanetScale = DialectMySQL
	DialectTDSQL       = DialectMySQL
	DialectTiDB        = DialectMySQL
	// Cовместимые с PostgreSQL
	DialectAlloyDB          = DialectPostgreSQL
	DialectArenadataDB      = DialectPostgreSQL
	DialectAuroraPostgreSQL = DialectPostgreSQL
	DialectAzurePostgreSQL  = DialectPostgreSQL
	DialectCitus            = DialectPostgreSQL
	DialectCockroachDB      = DialectPostgreSQL
	DialectGooglePostgreSQL = DialectPostgreSQL
	DialectGreenplum        = DialectPostgreSQL
	DialectKingbaseES       = DialectPostgreSQL
	DialectNeon             = DialectPostgreSQL
	DialectOpenGauss        = DialectPostgreSQL
	DialectSupabase         = DialectPostgreSQL
	DialectTimescaleDB      = DialectPostgreSQL
	DialectYandexDB         = DialectPostgreSQL
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
