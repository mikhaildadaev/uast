---
outline: deep
---

# Benchmarks
::: info **Info**
The best way to compare libraries is to run benchmarks in **your own environment** with **your own workload**. Each project has unique requirements — latency, throughput, memory usage, and integration complexity — and no single test can cover them all.

I recommend that you test `uast` alongside other libraries and choose the tool that best suits your needs.
:::

## Core Performance
These benchmarks measure the cost of building SQL queries. Simple queries select one column with a WHERE clause. Complex queries include JOINs, subqueries, GROUP BY, HAVING, ORDER BY, and LIMIT.

### MultiThread
| Mode    | Dialect    | Operations | Time (ns/op) | Memory (B/op) | Allocs |
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
| Mode    | Dialect    | Operations | Time (ns/op) | Memory (B/op) | Allocs |
|---------|------------|------------|--------------|---------------|--------|
| Complex | MariaDB    |       383K |        5,852 |         4,948 |     54 |
| Complex | MySQL      |       371K |        6,279 |         4,948 |     54 |
| Complex | PostgreSQL |       380K |        5,874 |         4,948 |     54 |
| Complex | SQLite     |       376K |        5,845 |         4,948 |     54 |
| Simple  | MariaDB    |       1.5M |        789.8 |           718 |      8 |
| Simple  | MySQL      |       1.5M |        778.6 |           718 |      8 |
| Simple  | PostgreSQL |       1.4M |        795.1 |           718 |      8 |
| Simple  | SQLite     |       1.4M |        787.9 |           718 |      8 |

::: tip **Note** 
Simple queries select one column with a basic WHERE clause. Complex queries include 2 JOINs, 3 subqueries, GROUP BY, HAVING, ORDER BY, and LIMIT. `sync.Pool` in Multi mode reuses `contexter` buffers, reducing allocations and GC pressure.

*Benchmarked on Intel Core i9-9880H (2.30 GHz).*
:::
