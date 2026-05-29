---
outline: deep
---

# Go
```bash
go get github.com/mikhaildadaev/uast
```

::: info **Info**
The latest stable version of `uast` is **v1.26.11**.
:::

## Run Test 
```bash
go test ./...
go test -bench=. ./...
go test -cover ./...
go test -race ./...
```

## Key Features
- **Type-safe** — Full generics support, compile-time type checking for columns and values.
- **Multi-dialect** — MariaDB, MySQL, PostgreSQL, SQLite from a single AST.
- **Secure by design** — Three-level `Value` / `Literal` / `Constant` system prevents SQL injection.
- **High performance** — `sync.Pool` for context reuse, ~360 ns/op for simple queries.
- **Zero dependencies** — Only Go standard library.
- **Quad outputs** — Every function documented with SQL output for all 4 dialects.
- **Hot dialect switch** — `SetDialect()` changes dialect at runtime without recreating the pool.
- **Complete DML** — DELETE, INSERT, SELECT, UPDATE with all standard clauses (JOIN, CTE, UPSERT, window functions, JSON).
- **150+ functions** — Aggregate, analytical, conditional, conversion, date/time, JSON, math, ranking, string.

## Limits
- **No DDL yet**: ALTER, CREATE, COMMENT, DROP, TRUNCATE coming in v2.
- **No code generation**: Table schemas are defined manually (code-gen planned).
- **MySQL**: `RETURNING` clause not supported (MySQL limitation).
- **SQLite**: `RIGHT JOIN` / `RIGHT OUTER JOIN` not supported (SQLite limitation).
- **PostgreSQL**: `JsonSet` in development.

## Supported Databases
| Database       | Version | Compatible                                                                                                                                                                         |
|----------------|---------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **MariaDB**    | 10.5.0+ | DoltDB, SingleStore                                                                                                                                                                |
| **MsSQL**      | 16.0.0+ | AmazonRDS, AzureSQL, Synapse                                                                                                                                                       |
| **MySQL**      | 8.0.31+ | AuroraMySQL, AzureMySQL, GoogleMySQL, OceanBase, PlanetScale, TDSQL                                                                                                                |
| **PostgreSQL** | 9.5.0+  | AlloyDB, ArenadataDB, AuroraPostgreSQL, AzurePostgreSQL, Citus, CockroachDB, GooglePostgreSQL, Greenplum, KingbaseES, Neon, OpenGauss, Supabase, TimescaleDB, YandexDB, YugabyteDB |
| **SQLite**     | 3.35.0+ | CloudflareD1, LiteFS, Turso                                                                                                                                                        |

::: tip **Note** 
The library does not verify the database version at runtime. Using features on older versions will result in SQL errors from the database. Ensure your database meets the minimum version requirements.
:::