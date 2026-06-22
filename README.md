[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mikhaildadaev/uast/blob/main/LICENSE.md)
[![Go Version](https://img.shields.io/github/go-mod/go-version/mikhaildadaev/uast)](https://github.com/mikhaildadaev/uast)
[![Go Reference](https://pkg.go.dev/badge/github.com/mikhaildadaev/uast.svg)](https://pkg.go.dev/github.com/mikhaildadaev/uast)
[![Go Report Card](https://goreportcard.com/badge/github.com/mikhaildadaev/uast)](https://goreportcard.com/report/github.com/mikhaildadaev/uast)
[![CI](https://github.com/mikhaildadaev/uast/actions/workflows/ci.yml/badge.svg)](https://github.com/mikhaildadaev/uast/actions/workflows/ci.yml)

# UAST

A high-performance, zero-dependency type‑safe query builder.

## Go
```bash
go get github.com/mikhaildadaev/uast
```

> **Info**
>
> The latest stable version of uast is v1.26.11.

### Run Test 
```bash
go test ./...
go test -bench=. ./...
go test -cover ./...
go test -race ./...
```

### Key Features
- **Type-safe** — Full generics support, compile-time type checking for columns and values.
- **Multi-dialect** — MariaDB, MsSQL, MySQL, PostgreSQL, SQLite from a single AST.
- **Secure by design** — Three-level `Constant` / `Literal` / `Value` system prevents SQL injection.
- **High performance** — `sync.Pool` for context reuse, ~360 ns/op for simple queries.
- **Zero dependencies** — Only Go standard library.
- **Cross-dialect docs** — Every function documented with SQL output for each supported dialect.
- **Hot dialect switch** — `SetDialect()` changes dialect at runtime without recreating the pool.
- **Complete DDL** — ALTER, COMMENT, CREATE, DROP, TRUNCATE.
- **Complete DML** — SELECT, INSERT, UPDATE, DELETE with all standard clauses (JOIN, CTE, UPSERT, window functions, JSON).
- **150+ functions** — Aggregate, analytical, conditional, conversion, date/time, JSON, math, ranking, string.

### Limits
- **MsSQL**: `JSON CONTAINS` / `JSON TYPE` function not supported (MsSQL limitation).
- **MySQL**: `RETURNING` clause not supported (MySQL limitation).
- **SQLite**: `RIGHT JOIN` / `RIGHT OUTER JOIN` not supported (SQLite limitation).

## Supported Databases
| Database       | Version | Compatible                                                                                                                                                                         |
|----------------|---------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **MariaDB**    | 10.5.0+ | DoltDB, SingleStore                                                                                                                                                                |
| **MsSQL**      | 16.0.0+ | AmazonRDS, AzureSQL, Synapse                                                                                                                                                       |
| **MySQL**      | 8.0.31+ | AuroraMySQL, AzureMySQL, GoogleMySQL, OceanBase, PlanetScale, TDSQL                                                                                                                |
| **PostgreSQL** | 9.5.0+  | AlloyDB, ArenadataDB, AuroraPostgreSQL, AzurePostgreSQL, Citus, CockroachDB, GooglePostgreSQL, Greenplum, KingbaseES, Neon, OpenGauss, Supabase, TimescaleDB, YandexDB, YugabyteDB |
| **SQLite**     | 3.35.0+ | CloudflareD1, LiteFS, Turso                                                                                                                                                        |

> **Note** 
> 
> The library does not verify the database version at runtime. Using features on older versions will result in SQL errors from the database. Ensure your database meets the minimum version requirements.

## Benchmarks
> **Info**
>
> The best way to compare libraries is to run benchmarks in **your own environment** with **your own workload**. Each project has unique requirements — latency, throughput, memory usage, and integration complexity — and no single test can cover them all.
>
> I recommend that you test `uast` alongside other libraries and choose the tool that best suits your needs.

### Core Performance
These benchmarks measure the cost of building SQL queries with an immutable AST, which guarantees zero mutations and allows safe reuse across multiple dialects. 

#### MultiThread
| Query   | Dialect    | Operations | Time (ns/op) | Memory (B/op) | Allocs |
|---------|------------|------------|--------------|---------------|--------|
| Complex | MariaDB    |       383K |        2,965 |         4,502 |     66 |
| Complex | MsSQL      |       371K |        3,136 |         4,502 |     66 |
| Complex | MySQL      |       371K |        3,136 |         4,502 |     66 |
| Complex | PostgreSQL |       380K |        3,299 |         4,502 |     66 |
| Complex | SQLite     |       376K |        3,399 |         4,518 |     67 |
| Simple  | MariaDB    |       3.2M |        431.9 |           824 |     12 |
| Simple  | MsSQL      |       3.0M |        416.9 |           824 |     12 |
| Simple  | MySQL      |       2.8M |        427.1 |           825 |     12 |
| Simple  | PostgreSQL |       2.8M |        434.0 |           825 |     12 |
| Simple  | SQLite     |       2.6M |        433.2 |           824 |     12 |

#### SingleThread
| Query   | Dialect    | Operations | Time (ns/op) | Memory (B/op) | Allocs |
|---------|------------|------------|--------------|---------------|--------|
| Complex | MariaDB    |       169K |        6,722 |         4,499 |     66 |
| Complex | MsSQL      |       171K |        6,819 |         4,499 |     66 |
| Complex | MySQL      |       175K |        6,707 |         4,499 |     66 |
| Complex | PostgreSQL |       177K |        6,680 |         4,499 |     66 |
| Complex | SQLite     |       168K |        6,727 |         4,499 |     67 |
| Simple  | MariaDB    |       1.3M |        944.3 |           824 |     12 |
| Simple  | MsSQL      |       1.3M |        941.9 |           824 |     12 |
| Simple  | MySQL      |       1.3M |        936.4 |           824 |     12 |
| Simple  | PostgreSQL |       1.3M |        931.3 |           824 |     12 |
| Simple  | SQLite     |       1.3M |        937.2 |           824 |     12 |

> **Note**
>
> Simple queries select one column with a basic WHERE clause. Complex queries include 2 JOINs, 3 subqueries, GROUP BY, HAVING, ORDER BY, and LIMIT. `sync.Pool` in Multi mode reuses `contexter` buffers, reducing allocations and GC pressure.
>
>*Benchmarked on Intel Core i9-9880H (2.30 GHz).*

## Usage
```go
import (
    "fmt"
    "log"
    "github.com/mikhaildadaev/uast"
)
func main() {
   ...
}
```

## Roadmap
- **Code generation** — Planned for future releases. Currently, schemas are defined manually.
- **Dialects** — `ClickHouse`, `Oracle` support.