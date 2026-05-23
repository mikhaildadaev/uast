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
- **多方言支持** — 从单个 AST 生成 MariaDB、MySQL、PostgreSQL、SQLite 的 SQL。
- **架构级安全** — 三级 `Value` / `Literal` / `Constant` 系统防止 SQL 注入。
- **高性能** — `sync.Pool` 重用上下文，简单查询约 360 ns/op。
- **零依赖** — 仅使用 Go 标准库。
- **四方言输出** — 每个函数都附有全部 4 种方言的 SQL 输出文档。
- **热切换方言** — `SetDialect()` 在运行时切换方言，无需重建连接池。
- **完整的 DML** — SELECT、INSERT、UPDATE、DELETE 支持所有标准子句（JOIN、CTE、UPSERT、窗口函数、JSON）。
- **150+ 函数** — 聚合、分析、条件、转换、日期/时间、JSON、数学、排名、字符串函数。

## Limits
- **暂不支持 DDL**：CREATE、ALTER、DROP、TRUNCATE、COMMENT 计划在 v2 中推出。
- **无代码生成**：表结构需手动定义（代码生成已列入计划）。
- **MySQL**：不支持 `RETURNING` 子句（MySQL 限制）。
- **SQLite**：不支持 `RIGHT JOIN` / `RIGHT OUTER JOIN`（SQLite 限制）。
- **PostgreSQL**：`JsonSet` 正在开发中。
