---
outline: deep
---

# 基准

::: info **关于**
比较库的最佳方式是在**您自己的环境**中使用**您自己的工作负载**运行基准测试。每个项目都有独特的需求——延迟、吞吐量、内存使用和集成复杂性——没有任何单一的测试能够覆盖所有情况。

我建议您将 `uast` 与其他库一起测试，并选择最适合您需求的工具。
:::

## Core Performance
这些基准测试测量构建 SQL 查询的成本。简单查询选择一列并带有 WHERE 条件。复杂查询包括 JOIN、子查询、GROUP BY、HAVING、ORDER BY 和 LIMIT。

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

::: tip **注** 
Multi 模式下的 `sync.Pool` 重用 `contexter` 缓冲区，减少内存分配并降低 GC 压力。

*Benchmarked on Intel Core i9-9880H (2.30 GHz).*
:::
