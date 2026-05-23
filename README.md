[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mikhaildadaev/uast/blob/main/LICENSE.md)
[![Go Version](https://img.shields.io/github/go-mod/go-version/mikhaildadaev/uast)](https://github.com/mikhaildadaev/uast)
[![Go Reference](https://pkg.go.dev/badge/github.com/mikhaildadaev/uast.svg)](https://pkg.go.dev/github.com/mikhaildadaev/uast)
[![Go Report Card](https://goreportcard.com/badge/github.com/mikhaildadaev/uast)](https://goreportcard.com/report/github.com/mikhaildadaev/uast)
[![CI](https://github.com/mikhaildadaev/uast/actions/workflows/ci.yml/badge.svg)](https://github.com/mikhaildadaev/uast/actions/workflows/ci.yml)

# UAST

A high-performance, zero‑allocation type‑safe SQL builder.

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

## Key Features
- **Type-safe** — Full generics support, compile-time type checking for columns and values.
- **Multi-dialect** — MariaDB, MySQL, PostgreSQL, SQLite from a single AST.
- **Secure by design** — Three-level `Value` / `Literal` / `Constant` system prevents SQL injection.
- **High performance** — `sync.Pool` for context reuse, ~360 ns/op for simple queries.
- **Zero dependencies** — Only Go standard library.
- **Quad outputs** — Every function documented with SQL output for all 4 dialects.
- **Hot dialect switch** — `SetDialect()` changes dialect at runtime without recreating the pool.
- **Complete DML** — SELECT, INSERT, UPDATE, DELETE with all standard clauses (JOIN, CTE, UPSERT, window functions, JSON).
- **150+ functions** — Aggregate, analytical, conditional, conversion, date/time, JSON, math, ranking, string.

## Limits
- **No DDL yet**: CREATE, ALTER, DROP, TRUNCATE, COMMENT coming in v2.
- **No code generation**: Table schemas are defined manually (code-gen planned).
- **MySQL**: `RETURNING` clause not supported (MySQL limitation).
- **SQLite**: `RIGHT JOIN` / `RIGHT OUTER JOIN` not supported (SQLite limitation).
- **PostgreSQL**: `JsonSet` in development.

## Benchmarks
> **Info**
>
> The best way to compare libraries is to run benchmarks in **your own environment** with **your own workload**. Each project has unique requirements — latency, throughput, memory usage, and integration complexity — and no single test can cover them all.
>
> I recommend that you test `uast` alongside other libraries and choose the tool that best suits your needs.

## Core Performance
These benchmarks measure the cost of building SQL queries. Simple queries select one column with a WHERE clause. Complex queries include JOINs, subqueries, GROUP BY, HAVING, ORDER BY, and LIMIT.

### MultiThread
| Query   | Dialect    | Operations | Time (ns/op) | Memory (B/op) | Allocs |
|---------|------------|------------|--------------|---------------|--------|
| Complex | MariaDB    |       383K |        2,965 |         4,971 |     54 |
| Complex | MySQL      |       371K |        3,136 |         4,972 |     54 |
| Complex | PostgreSQL |       380K |        3,299 |         4,970 |     54 |
| Complex | SQLite     |       376K |        3,399 |         4,972 |     54 |
| Simple  | MariaDB    |       3.7M |        335.5 |           720 |      8 |
| Simple  | MySQL      |       3.5M |        349.0 |           720 |      8 |
| Simple  | PostgreSQL |       3.3M |        398.4 |           720 |      8 |
| Simple  | SQLite     |       3.3M |        358.3 |           720 |      8 |

### SingleThread
| Query   | Dialect    | Operations | Time (ns/op) | Memory (B/op) | Allocs |
|---------|------------|------------|--------------|---------------|--------|
| Complex | MariaDB    |       197K |        5,852 |         4,948 |     54 |
| Complex | MySQL      |       204K |        6,279 |         4,948 |     54 |
| Complex | PostgreSQL |       196K |        5,874 |         4,948 |     54 |
| Complex | SQLite     |       196K |        5,845 |         4,948 |     54 |
| Simple  | MariaDB    |       1.5M |        789.8 |           718 |      8 |
| Simple  | MySQL      |       1.5M |        778.6 |           718 |      8 |
| Simple  | PostgreSQL |       1.4M |        795.1 |           718 |      8 |
| Simple  | SQLite     |       1.4M |        787.9 |           718 |      8 |

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
- **DDL** — `CREATE` (TABLE, INDEX, VIEW), `ALTER` (ADD, DROP, ALTER, RENAME COLUMN, CONSTRAINT), `DROP` (TABLE, INDEX, VIEW), `TRUNCATE`, `COMMENT`.
- **Dialects** — `ClickHouse`, `MsSQL`, `Oracle` support.