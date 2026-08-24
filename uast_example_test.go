// UAST (Universal Abstract Syntax Tree)
// Copyright (C) 2026 Mikhail Dadaev
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package uast_test

import (
	"fmt"
	"log"

	"github.com/mikhaildadaev/uast"
)

// Примеры использования публичных конструкторов
func ExampleNewSQL() {
	builder := uast.NewSQL(
		uast.WithDialect(uast.DialectMySQL),
		uast.WithMutate(false),
	)
	defer builder.Close()
	stmt := uast.NewSelect(uast.NewTable("users", "u")).
		Fields(
			uast.Field[int64]("u", "id"),
			uast.Field[string]("u", "name"),
		).
		Where(
			uast.Equal(uast.Field[string]("u", "status"), uast.Value("active")),
		)
	queryMySQL, argsMySQL, errMySQL := builder.Build(stmt)
	if errMySQL != nil {
		log.Fatal(errMySQL)
	}
	fmt.Printf("MySQL: %s | %v\n\n", queryMySQL, argsMySQL)
	builder.SetDialect(uast.DialectPostgreSQL)
	queryPostgreSQL, argsPostgreSQL, errPostgreSQL := builder.Build(stmt)
	if errPostgreSQL != nil {
		log.Fatal(errPostgreSQL)
	}
	fmt.Printf("PostgreSQL: %s | %v\n", queryPostgreSQL, argsPostgreSQL)
	// Output:
	// MySQL: SELECT `u`.`id`, `u`.`name` FROM `users` AS `u` WHERE `u`.`status` = ? | [active]
	//
	// PostgreSQL: SELECT "u"."id", "u"."name" FROM "users" AS "u" WHERE "u"."status" = $1 | [active]
}
