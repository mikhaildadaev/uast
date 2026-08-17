---
outline: deep
---

# Go
```bash
go get github.com/mikhaildadaev/uast
```

::: info **Информация**
Последняя стабильная версия `uast` — **v1.26.11**.
:::

## Run Test 
```bash
go test ./...
go test -bench=. ./...
go test -cover ./...
go test -race ./...
```

## Key Features
- **Типобезопасность** — Полная поддержка дженериков, проверка типов колонок и значений на этапе компиляции.
- **Мультидиалектность** — MariaDB, MsSQL, MySQL, PostgreSQL, SQLite из одного AST.
- **Безопасность на уровне архитектуры** — Трёхуровневая система `Value` / `Literal` / `Constant` предотвращает SQL-инъекции.
- **Высокая производительность** — `sync.Pool` для переиспользования контекста, ~360 нс/оп для простых запросов.
- **Ноль зависимостей** — Только стандартная библиотека Go.
- **Кросс-диалектная документация** — В каждой функции показан SQL для всех поддерживаемых диалектов.
- **Горячая смена диалекта** — `SetDialect()` меняет диалект во время выполнения без пересоздания пула.
- **Полный DDL** — ALTER, COMMENT, CREATE, DROP, TRUNCATE.
- **Полный DML** — DELETE, INSERT, SELECT, UPDATE со всеми стандартными клаузами (JOIN, CTE, UPSERT, оконные функции, JSON).
- **150+ функций** — Агрегатные, аналитические, условные, конвертации, дата/время, JSON, математические, ранжирующие, строковые.

## Limits
- **MsSQL**: Функция `JSON CONTAINS` / `JSON TYPE` не поддерживается (ограничение MsSQL).
- **MySQL**: Клауза `RETURNING` не поддерживается (ограничение MySQL).
- **SQLite**: `RIGHT JOIN` / `RIGHT OUTER JOIN` не поддерживаются (ограничение SQLite).

## Supported Databases
| Database       | Version | Compatible                                                                                                                                                                         |
|----------------|---------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **MariaDB**    | 10.7.0+ | DoltDB, SingleStore                                                                                                                                                                |
| **MsSQL**      | 16.0.0+ | AmazonRDS, AzureSQL, Synapse                                                                                                                                                       |
| **MySQL**      | 8.0.31+ | AuroraMySQL, AzureMySQL, GoogleMySQL, OceanBase, PlanetScale, TDSQL                                                                                                                |
| **PostgreSQL** | 9.5.0+  | AlloyDB, ArenadataDB, AuroraPostgreSQL, AzurePostgreSQL, Citus, CockroachDB, GooglePostgreSQL, Greenplum, KingbaseES, Neon, OpenGauss, Supabase, TimescaleDB, YandexDB, YugabyteDB |
| **SQLite**     | 3.35.0+ | CloudflareD1, LiteFS, Turso                                                                                                                                                        |

::: tip **Примечание** 
Библиотека не проверяет версию базы данных во время выполнения. Использование функций на более старых версиях приведёт к ошибкам SQL со стороны базы данных. Убедитесь, что ваша база данных соответствует минимальным требованиям к версии.
:::