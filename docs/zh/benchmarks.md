---
outline: deep
---

# Benchmarks

::: info **关于**
比较库的最佳方式是在**您自己的环境**中使用**您自己的工作负载**运行基准测试。每个项目都有独特的需求——延迟、吞吐量、内存使用和集成复杂性——没有任何单一的测试能够覆盖所有情况。

我建议您将 `uast` 与其他库一起测试，并选择最适合您需求的工具。
:::

## Core Performance
这些基准测试衡量了使用不可变 AST 构建 SQL 查询的成本，这保证了零突变，并允许跨多种方言安全地重用。

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

::: tip **注** 
Multi 模式下的 `sync.Pool` 重用 `contexter` 缓冲区，减少内存分配并降低 GC 压力。

*Benchmarked on Intel Core i9-9880H (2.30 GHz).*
:::
