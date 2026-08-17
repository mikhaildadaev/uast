---
outline: deep
---

# Go
```bash
go get github.com/mikhaildadaev/uast
```

::: info **关于**
`uast` 的最新稳定版本是 **v1.26.11**.
:::

## Run Test 
```bash
go test ./...
go test -bench=. ./...
go test -cover ./...
go test -race ./...
```

## Key Features
- **类型安全** — 完整的泛型支持，在编译时检查列和值的类型。
- **多方言支持** — 从单个 AST 生成 MariaDB、MsSQL、MySQL、PostgreSQL、SQLite 的 SQL。
- **架构级安全** — 三级 `Value` / `Literal` / `Constant` 系统防止 SQL 注入。
- **高性能** — `sync.Pool` 重用上下文，简单查询约 360 ns/op。
- **零依赖** — 仅使用 Go 标准库。
- **跨数据库文档** — 每个函数都附带了在所有支持的数据库方言中的 SQL 输出示例。
- **热切换方言** — `SetDialect()` 在运行时切换方言，无需重建连接池。
- **完整的 DDL** — ALTER、COMMENT、CREATE、DROP、TRUNCATE。
- **完整的 DML** — DELETE, INSERT, SELECT, UPDATE 支持所有标准子句（JOIN、CTE、UPSERT、窗口函数、JSON）。
- **150+ 函数** — 聚合、分析、条件、转换、日期/时间、JSON、数学、排名、字符串函数。

## Limits
- **MsSQL**：不支持 `JSON CONTAINS` / `JSON TYPE` 函数（MsSQL 自身限制）。
- **MySQL**：不支持 `RETURNING` 子句（MySQL 限制）。
- **SQLite**：不支持 `RIGHT JOIN` / `RIGHT OUTER JOIN`（SQLite 限制）。

## Supported Databases
| Database       | Version | Compatible                                                                                                                                                                         |
|----------------|---------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **MariaDB**    | 10.7.0+ | DoltDB, SingleStore                                                                                                                                                                |
| **MsSQL**      | 16.0.0+ | AmazonRDS, AzureSQL, Synapse                                                                                                                                                       |
| **MySQL**      | 8.0.31+ | AuroraMySQL, AzureMySQL, GoogleMySQL, OceanBase, PlanetScale, TDSQL                                                                                                                |
| **PostgreSQL** | 9.5.0+  | AlloyDB, ArenadataDB, AuroraPostgreSQL, AzurePostgreSQL, Citus, CockroachDB, GooglePostgreSQL, Greenplum, KingbaseES, Neon, OpenGauss, Supabase, TimescaleDB, YandexDB, YugabyteDB |
| **SQLite**     | 3.35.0+ | CloudflareD1, LiteFS, Turso                                                                                                                                                        |

::: tip **注** 
该库不会在运行时检查数据库版本。在旧版本上使用功能将导致数据库返回 SQL 错误。请确保您的数据库满足最低版本要求。
:::