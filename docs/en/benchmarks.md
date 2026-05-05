---
outline: deep
---

# Benchmarks
::: info **Info**
The best way to compare libraries is to run benchmarks in **your own environment** with **your own workload**. Each project has unique requirements — latency, throughput, memory usage, and integration complexity — and no single test can cover them all.

I recommend that you test `uast` alongside other libraries and choose the tool that best suits your needs.
:::

## Core Performance
These benchmarks measure the **cost of formatting and extracting context** by writing to `io.Discard`.

### MultiThread
| Name                 | Operations | Time (ns/op) | Memory (B/op) | Allocs |
|----------------------|------------|--------------|---------------|--------|

### SingleThread
| Name                 | Operations | Time (ns/op) | Memory (B/op) | Allocs |
|----------------------|------------|--------------|---------------|--------|

::: tip **Note** 

*Benchmarked on Intel Core i9-9880H (2.30 GHz).*
:::
