UAST - Universal Abstract SQL Transformer
Пакет UAST предоставляет типобезопасный построитель SQL запросов для Go с использованием Fluent-интерфейса.

# API
# API - Builder [data]
   - [+] NewBuilder(dialect)
# API - Builder - Dialect [data]
   - [+] DialectClickHouse
   - [+] DialectDefault
   - [+] DialectMySQL
   - [+] DialectMsSQL
   - [+] DialectPostgreSQL
   - [+] DialectSQLite
# API - Query [data]
   - [+] Build(statement)
# API - Statement [list]
   - [-] DDL
   - [+] DML
# API - Statement - DDL [data]
   - [-] Alter(...)                                           # ALTER/ALTER/
   - [-] Comment(...)                                         # COMMENT/COMMENT/
   - [-] Create(...)                                          # CREATE/CREATE/
   - [-] Drop(...)                                            # DROP/DROP/
   - [-] Rename(...)                                          # RENAME/RENAME/
   - [-] Truncate(...)                                        # TRUNCATE/TRUNCATE/
# API - Statement - DML [data]
   - [+] Delete(from)                                         # DELETE/DELETE/
   - [+] Insert(columns...)                                   # INSERT/INSERT/
   - [+] Select(fields...)                                    # SELECT/SELECT/
   - [+] Update(onto)                                         # UPDATE/UPDATE/
# API - Statement - DML - Delete [data]
...
# API - Statement - DML - Insert [data]
...
# API - Statement - DML - Select [data]
   - [+] Distinct(bool)                                       # DISTINCT/DISTINCT/
   - [+] Field(fields...)                                     # FIELD/FIELD/
   - [+] From(from)                                           # FROM/FROM/
   - [+] GroupBy(groupbys...)                                 # GROUP BY/GROUP BY/
   - [+] Having(having)                                       # HAVING/HAVING/
   - [+] Join(joins...)                                       # JOIN/JOIN/
   - [+] Limit(limit)                                         # LIMIT/LIMIT/
   - [+] Offset(offset)                                       # OFFSET/OFFSET/
   - [+] OrderBy(orderbys...)                                 # ORDER BY/ORDER BY/
   - [+] Unions(unions...)                                    # UNIONS/UNIONS/
   - [+] Where(where)                                         # WHERE/WHERE
   - [+] With(withs...)                                       # WITH/WITH/
# API - Statement - DML - Update [data]
...
# API - Structure - Field [data]
   - [+] Column(tableAlias, columnName)
   - [+] Function()
   - [+] Subquery(statement)
# API - Structure - From [data]
   - [+] Cte(cteName, aliasName)
   - [+] Query(statement, aliasName)
   - [+] Table(tableName, aliasName)
# API - Structure - GroupBy [data]
   - [+] Column(tableAlias, columnName)
   - [+] Subquery(statement)
# API - Structure - Having & Where [data]
   - [+] And(expressions...)
   - [+] Or(expressions...)
# API - Structure - Join [data]
   - [+] Cross(source)                                        # CROSS JOIN/CROSS JOIN/
   - [+] Full(source, expression)                             # FULL JOIN/FULL JOIN/
   - [+] FullOuter(source, expression)                        # FULL OUTER JOIN/FULL OUTER JOIN/
   - [+] Inner(source, expression)                            # INNER JOIN/INNER JOIN/
   - [+] Left(source, expression)                             # LEFT JOIN/LEFT JOIN/
   - [+] LeftOuter(source, expression)                        # LEFT OUTER JOIN/LEFT OUTER JOIN/
   - [+] Right(source, expression)                            # RIGHT JOIN/RIGHT JOIN/
   - [+] RightOuter(source, expression)                       # RIGHT OUTER JOIN/RIGHT OUTER JOIN/
# API - Structure - Limit [data]
   - [+] Value(value)
# API - Structure - Offset [data]
   - [+] Value(value)
# API - Structure - OrderBy [data]
   - [+] Asc(expression)                                      # ASC/ASC/
   - [+] Desc(expression)                                     # DESC/DESC/
# API - Structure - Union [data]
   - [+] Union(statement)                                     # UNION/UNION/
   - [+] UnionAll(statement)                                  # UNION ALL/UNION ALL/
   - [+] UnionExcept(statement)                               # EXCEPT/EXCEPT/
   - [+] UnionIntersect(statement)                            # INTERSECT/INTERSECT/
# API - Structure - With [data]
   - [+] With(...)
