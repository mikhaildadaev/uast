UAST - Universal Abstract SQL Transformer
Пакет UAST предоставляет типобезопасный построитель SQL запросов для Go с использованием Fluent-интерфейса.
Поддерживаемые DDL операции:
   - [.] Alter
   - [.] Comment
   - [.] Create
   - [.] Drop
   - [.] Rename
   - [.] Truncate
Поддерживаемые DML операции:
   - [+] Delete
   - [+] Insert
   - [+] Select
   - [+] Update
Поддерживаемые SQL диалекты:
   - [.] ClickHouse
   - [+] MySQL
   - [.] MsSQL
   - [+] PostgreSQL
   - [-] SQLite

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
# API - Static [list]
   - [+] DataType
# API - Static - DataType [list]
   - [+] TypeBinary
   - [+] TypeDatetime
   - [+] TypeNumeric
   - [+] TypeString
   - [+] TypeSpecial
# API - Static - DataType | TypeBinary [data]
   - [+] TypeBinary                                           # BINARY/BYTEA/
   - [+] TypeVarBinary                                        # VARBINARY/BYTEA/
# API - Static - DataType | TypeDatetime [data]
   - [+] TypeDate                                             # DATE/DATE/
   - [+] TypeDateTime                                         # DATETIME/TIMESTAMP/
   - [+] TypeTime                                             # TIME/TIME/
   - [+] TypeTimestamp                                        # DATETIME/TIMESTAMP/
# API - Static - DataType | TypeNumeric [data]
   - [+] TypeBigInt                                           # SIGNED/BIGINT/
   - [+] TypeDecimal                                          # DECIMAL/DECIMAL/
   - [+] TypeDouble                                           # DECIMAL/DOUBLE PRECISION/
   - [+] TypeFloat                                            # DECIMAL/REAL/
   - [+] TypeInt                                              # SIGNED/INTEGER/
   - [+] TypeSmallInt                                         # SIGNED/SMALLINT/
# API - Static - DataType | TypeString [data]
   - [+] TypeChar                                             # CHAR/CHAR/
   - [+] TypeString                                           # VARCHAR/VARCHAR/
   - [+] TypeText                                             # TEXT/TEXT/
   - [+] TypeVarChar                                          # VARCHAR/VARCHAR/
# API - Static - DataType | TypeSpecial [data]
   - [+] TypeArray                                            # JSON/ARRAY/
   - [+] TypeBoolean                                          # TINYINT(1)/BOOLEAN/
   - [+] TypeJson                                             # JSON/JSONB/
   - [+] TypeUUID                                             # CHAR(36)/UUID/
   - [+] TypeXML                                              # TEXT/XML/
# API - Structure [list]
   - [+] Column
   - [+] Const
   - [+] Function
   - [+] Literal
   - [+] Subquery
   - [+] Value
# API - Structure - Field [data]
   - [+] Column(tableAlias, columnName)
   - [+] Function()
   - [+] Subquery(statement)
# API - Structure - From [data]
   - [+] Cte(cteName, aliasName)
   - [+] Query(statement, aliasName)
   - [+] Table(tableName, aliasName)
# API - Structure - Function [list]
   - [+] Aggregate
   - [+] Condition
   - [+] Convertation
   - [+] Date and time
   - [+] Math
   - [+] String
# API - Structure - Function | Aggregate [data]
   - [+] Avg(number, distinct)                                # AVG/AVG/
   - [+] BitAnd(number, distinct)                             # BIT_AND/BIT_AND/
   - [+] BitOr(number, distinct)                              # BIT_OR/BIT_OR/
   - [+] BitXor(number, distinct)                             # BIT_XOR/BIT_XOR/
   - [+] Count(number, distinct)                              # COUNT/COUNT/
   - [+] GroupConcat(number, distinct)                        # GROUP_CONCAT/STRING_AGG/
   - [+] Max(number, distinct)                                # MAX/MAX/
   - [+] Min(number, distinct)                                # MIN/MIN/
   - [+] StdDev(number, distinct)                             # STDDEV/STDDEV_SAMP/
   - [+] Sum(number, distinct)                                # SUM/SUM/
   - [+] Variance(number, distinct)                           # VARIANCE/VAR_SAMP/
# API - Structure - Function | Condition [data]
   - [+] Case(whens, thens)                                   # CASE/CASE/
   - [+] Coalesce(expressions...)                             # COALESCE/COALESCE/
   - [+] Greatest(expressions...)                             # GREATEST/GREATEST/
   - [+] Least(expressions...)                                # LEAST/LEAST/
   - [+] NullIf(a, b)                                         # NULLIF/NULLIF/
# API - Structure - Function | Convertation [data]
   - [+] Cast(value, valueType)                               # CAST/CAST/
   - [+] CharLength(str)                                      # CHAR_LENGTH/CHAR_LENGTH/
   - [+] DateFormat(value, mask)                              # DATE_FORMAT/TO_CHAR/
   - [+] Degrees(angle)                                       # DEGREES/DEGREES/
   - [+] Length(str)                                          # LENGTH/LENGTH/
   - [+] Position(str, subStr)                                # POSITION/POSITION/
   - [+] Radians(angle)                                       # RADIANS/RADIANS/
# API - Structure - Function | Date and time [data]
   - [+] CurDate()                                            # CURDATE/CURRENT_DATE/
   - [+] CurTime()                                            # CURTIME/CURRENT_TIME/
   - [+] DateAdd(datetime, interval)                          # DATE_ADD//
   - [+] DateDiff(datetimeEnd, datetimeStart)                 # DATEDIFF/DATE_PART/
   - [+] DateSub(datetime, interval)                          # DATE_SUB//
   - [+] Day(datetime)                                        # DAY/EXTRACT/
   - [+] DayName(datetime)                                    # DAYNAME/TO_CHAR/
   - [+] Hour(datetime)                                       # HOUR/EXTRACT/
   - [+] Minute(datetime)                                     # MINUTE/EXTRACT/
   - [+] Month(datetime)                                      # MONTH/EXTRACT/
   - [+] MonthName(datetime)                                  # MONTHNAME/TO_CHAR/
   - [+] Now()                                                # NOW/CURRENT_TIMESTAMP/
   - [+] Quarter(datetime)                                    # QUARTER/EXTRACT/
   - [+] Second(datetime)                                     # SECOND/EXTRACT/
   - [+] TimeAdd(datetime, interval)                          # TIME_ADD//
   - [+] TimeDiff(datetimeEnd, datetimeStart)                 # TIMEDIFF/DATE_PART/
   - [+] TimeSub(datetime, interval)                          # TIME_SUB//
   - [+] Week(datetime)                                       # WEEK/EXTRACT/
   - [+] Year(datetime)                                       # YEAR/EXTRACT/
# API - Structure - Function | Json [data]
   - [+] JsonArray(values...)                                 # JSON_ARRAY/JSON_ARRAY/
   - [+] JsonArrayAgg(valueAggregator)                        # JSON_ARRAYAGG/JSON_AGG/
   - [+] JsonContains(haystack, needle)                       # JSON_CONTAINS//
   - [+] JsonExtract(json, path, valueType)                   # //
   - [+] JsonObject(pairs...)                                 # JSON_OBJECT/JSON_BUILD_OBJECT/
   - [+] JsonObjectAgg(keySelector, valueAggregator)          # JSON_OBJECTAGG/JSON_OBJECT_AGG/
   - [+] JsonRemove(json, paths)                              # //
   - [-] JsonSet(json, paths, values)                         # //
   - [+] JsonType(subject)                                    # JSON_TYPE/jsonb_typeof/
# API - Structure - Function | Math [data]
   - [+] Abs(numeric)                                         # ABS/ABS/
   - [+] ACos(angle)                                          # ACOS/ACOS/
   - [+] ASin(angle)                                          # ASIN/ASIN/
   - [+] ATan(angle)                                          # ATAN/ATAN/
   - [+] ATan2(y, x)                                          # ATAN2/ATAN2/
   - [+] Cbrt(numeric)                                        # CBRT/CBRT/
   - [+] Ceil(numeric)                                        # CEILING/CEIL/
   - [+] Cos(angle)                                           # COS/COS/
   - [+] Exp(numeric)                                         # EXP/EXP/
   - [+] Floor(numeric)                                       # FLOOR/FLOOR/
   - [+] Ln(numeric)                                          # LN/LN/
   - [+] Log(numeric, base)                                   # LOG/LOG/
   - [+] Mod(numeric, divisor)                                # MOD/MOD/
   - [+] Pi()                                                 # PI/PI/
   - [+] Power(numeric, exponent)                             # POWER/POWER/
   - [+] Rand()                                               # RAND/RANDOM/
   - [+] Round(numeric, precision)                            # ROUND/ROUND/
   - [+] Sin(angle)                                           # SIN/SIN/
   - [+] Sqrt(numeric)                                        # SQRT/SQRT/
   - [+] Tan(angle)                                           # TAN/TAN/
   - [+] Trunc(numeric, places)                               # TRUNCATE/TRUNC/
# API - Structure - Function | String [data]
   - [+] Concat(strs...)                                      # CONCAT/CONCAT/
   - [+] ConcatWs(separator, strs...)                         # CONCAT_WS/CONCAT_WS/
   - [+] LeftString(str, count)                               # LEFT/LEFT/
   - [+] Lower(str)                                           # LOWER/LOWER/
   - [+] LPad(str, count, separator)                          # LPAD/LPAD/
   - [+] LTrim(str)                                           # LTRIM/LTRIM/
   - [+] Repeat(str, count)                                   # REPEAT/REPEAT/
   - [+] Replace(str, strOld, strNew)                         # REPLACE/REPLACE/
   - [+] Reverse(str)                                         # REVERSE/REVERSE/
   - [+] RightString(str, count)                              # RIGHT/RIGHT/
   - [+] RPad(str, count, separator)                          # RPAD/RPAD/
   - [+] RTrim(str)                                           # RTRIM/RTRIM/
   - [+] SubString(str, startPos, lengthStr)                  # SUBSTRING/SUBSTRING/
   - [+] Trim(str)                                            # TRIM/TRIM/
   - [+] Upper(str)                                           # UPPER/UPPER/
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
# API - Structure - Predicate [list]
   - [+] Binary
   - [+] Comparison
   - [+] Logical
# API - Structure - Predicate | Binary [data]
   - [+] BitwiseAnd(left, right)                              # &
   - [+] BitwiseOr(left, right)                               # |
   - [+] BitwiseXor(left, right)                              # ^
   - [+] Divide(left, right)                                  # /
   - [+] Minus(left, right)                                   # -
   - [+] Modulo(left, right)                                  # %
   - [+] Multiply(left, right)                                # *
   - [+] Plus(left, right)                                    # +
   - [+] ShiftL(left, right)                                  # <<
   - [+] ShiftR(left, right)                                  # >>
# API - Structure - Predicate | Comparison [data]
   - [+] Between(left, start, end)                            # BETWEEN/BETWEEN/
   - [+] Equal(left, right)                                   # =
   - [+] Exists(left)                                         # EXISTS/EXISTS/
   - [+] GE(left, right)                                      # >=
   - [+] GT(left, right)                                      # >
   - [+] ILike(left, right)                                   # LIKE/ILIKE/
   - [+] In(left, right)                                      # IN/IN/
   - [+] IsNotNull(left)                                      # IS NOT NULL/IS NOT NULL/
   - [+] IsNull(left)                                         # IS NULL/IS NULL/
   - [+] LE(left, right)                                      # <=
   - [+] Like(left, right)                                    # LIKE/LIKE/
   - [+] LT(left, right)                                      # <
   - [+] NotBetween(left, start, end)                         # NOT BETWEEN/NOT BETWEEN/
   - [+] NotEqual(left, right)                                # <>
   - [+] NotExists(left)                                      # NOT EXISTS/NOT EXISTS/
   - [+] NotILike(left, right)                                # NOT LIKE/NOT ILIKE/
   - [+] NotIn(left, right)                                   # NOT IN/NOT IN/
   - [+] NotLike(left, right)                                 # NOT LIKE/NOT LIKE/
# API - Structure - Predicate | Logical [data]
   - [+] And(expressions...)                                  # AND/AND/
   - [+] Or(expressions...)                                   # OR/OR/
# API - Structure - Union [data]
   - [+] Union(statement)                                     # UNION/UNION/
   - [+] UnionAll(statement)                                  # UNION ALL/UNION ALL/
   - [+] UnionExcept(statement)                               # EXCEPT/EXCEPT/
   - [+] UnionIntersect(statement)                            # INTERSECT/INTERSECT/
# API - Structure - With [data]
   - [+] With(...)

# Examples
## QueryDelete
...
## QueryInsert
...
## QuerySelect
...
## QueryUpdate
...

# Description [list]
   - [+] Function - функции
# Description | Function [list]
   - [+] Aggregate - функции агрегатные
   - [+] Condition - функции условий
   - [+] Convertation - функции конвертации
   - [+] Date and time  - функции даты и времени
   - [+] Json - функции обмена данными
   - [+] Math - функции математические
   - [+] String - функции строковые
# Description | Function - Aggregate [data]
   - [+] Avg - вычисляет среднее арифметическое значений
   - [+] BitAnd - выполняет побитовое И для всех значений
   - [+] BitOr - выполняет побитовое ИЛИ для всех значений
   - [+] BitXor - выполняет побитовое исключающее ИЛИ для всех значений
   - [+] Count - подсчитывает количество строк или не-NULL значений
   - [+] GroupConcat - объединяет значения в строку с разделителем
   - [+] Max - находит максимальное значение
   - [+] Min - находит минимальное значение
   - [+] StdDev - вычисляет стандартное отклонение
   - [+] Sum - вычисляет сумму значений
   - [+] Variance - вычисляет дисперсию
# Description | Function - Condition [data]
   - [+] Case - реализует условную логику
   - [+] Coalesce - возвращает первое не-NULL значение из списка
   - [+] Greatest - возвращает наибольшее значение из списка
   - [+] Least - возвращает наименьшее значение из списка
   - [+] NullIf - возвращает NULL если два значения равны
# Description | Function - Convertation [data]
   - [+] Cast - преобразует значение в указанный тип данных
   - [+] CharLength - возвращает длину строки в символах
   - [+] DateFormat - форматирует дату/время по указанному шаблону
   - [+] Degrees - преобразует радианы в градусы
   - [+] Length - возвращает длину строки в байтах
   - [+] Position - возвращает позицию подстроки в строке
   - [+] Radians - преобразует градусы в радианы
# Description | Function - Date and time [data]
   - [+] CurDate - возвращает текущую дату
   - [+] CurTime - возвращает текущее время
   - [+] DateAdd - добавляет интервал к дате
   - [+] DateDiff - возвращает разницу между двумя датами в днях
   - [+] DateSub - вычитает интервал из даты
   - [+] Day - извлекает день из даты
   - [+] DayName - возвращает название дня недели
   - [+] Hour - извлекает час из времени
   - [+] Minute - извлекает минуты из времени
   - [+] Month - извлекает месяц из даты
   - [+] MonthName - возвращает название месяца
   - [+] Now - возвращает текущую дату и время
   - [+] Quarter - извлекает квартал из даты
   - [+] Second - извлекает секунды из времени
   - [+] TimeAdd - добавляет интервал к времени
   - [+] TimeDiff - возвращает разницу между двумя временными промежутками
   - [+] TimeSub - вычитает интервал из времени
   - [+] Week - извлекает номер недели из даты
   - [+] Year - извлекает год из даты
# Description | Function - Json [data]
   - [+] JsonArray - создаёт массив в JSON;
   - [+] JsonArrayAgg - агрегирует значения в массив JSON;
   - [+] JsonContains - поиск значение по всему JSON;
   - [+] JsonExtract - извлекает значение по указанному пути в JSON;
   - [+] JsonObject - создаёт объект из пар ключ-значение в JSON;
   - [+] JsonObjectAgg - агрегирует значения в объект JSON;
   - [+] JsonRemove - удаляет элемент по указанному пути в JSON;
   - [.] JsonSet - устанавливает значение по указанному пути в JSON;
   - [+] JsonType - возвращает тип значения в JSON;
# Description | Function - Math [data]
   - [+] Abs - возвращает абсолютное значение числа
   - [+] ACos - возвращает арккосинус числа
   - [+] ASin - возвращает арксинус числа
   - [+] ATan - возвращает арктангенс числа
   - [+] ATan2 - возвращает арктангенс отношения y/x
   - [+] Cbrt - возвращает кубический корень числа
   - [+] Ceil - округляет число вверх до ближайшего целого
   - [+] Cos - возвращает косинус угла
   - [+] Exp - возвращает e в степени числа
   - [+] Floor - округляет число вниз до ближайшего целого
   - [+] Ln - возвращает натуральный логарифм числа
   - [+] Log - возвращает логарифм числа по указанному основанию
   - [+] Mod - возвращает остаток от деления
   - [+] Pi - возвращает число пи
   - [+] Power - возводит число в указанную степень
   - [+] Rand - генерирует случайное число
   - [+] Round - округляет число до указанного количества знаков
   - [+] Sin - возвращает синус угла
   - [+] Sqrt - возвращает квадратный корень числа
   - [+] Tan - возвращает тангенс угла
   - [+] Trunc - усекает число до указанного количества знаков
# Description | Function - String [data]
   - [+] Concat - объединяет строки
   - [+] ConcatWs - объединяет строки с указанным разделителем
   - [+] LeftString - возвращает указанное количество первых символов строки
   - [+] Lower - преобразует строку в нижний регистр
   - [+] LPad - дополняет строку слева до указанной длины
   - [+] LTrim - удаляет пробелы в начале строки
   - [+] Repeat - повторяет строку указанное количество раз
   - [+] Replace - заменяет все вхождения подстроки
   - [+] Reverse - переворачивает строку
   - [+] RightString - возвращает указанное количество последних символов строки
   - [+] RPad - дополняет строку справа до указанной длины
   - [+] RTrim - удаляет пробелы в конце строки
   - [+] SubString - извлекает подстроку с указанной позиции и длины
   - [+] Trim - удаляет пробелы в начале и конце строки
   - [+] Upper - преобразует строку в верхний регистр

# Future [list]
   - [.] Structure
# Future | Structure [data]
   - [+] Window