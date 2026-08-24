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
