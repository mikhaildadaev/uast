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

package uast

import (
	"testing"
)

// Бенчмарки компонентов
func Benchmark_Immutable_Multi(b *testing.B) {
	builder := NewSQL()
	defer builder.Close()
	benchmarkAllDialects(b, func(b *testing.B, supportDialect *SupportDialect) {
		builder.SetDialect(supportDialect)
		b.Run("Simple", func(b *testing.B) {
			stmt := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.ID.Expr(),
				).
				Where(
					Equal(Test.Table.Users.Number.Expr(), Value(0)),
				)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _, err := builder.Build(stmt)
					if err != nil {
						b.Errorf("Build failed: %v", err)
					}
				}
			})
		})
		b.Run("Complex", func(b *testing.B) {
			stmt := NewSelect(Test.Table.Order).
				Fields(
					Test.Table.Orders.ID.Expr().As("test_id"),
					Test.Table.Orders.String.Expr().As("test_string"),
					Test.Table.Orders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table.User).
						Fields(
							Count(Test.Table.Users.ID.Expr(), false),
						).
						Where(
							Equal(Test.Table.Users.Number.Expr(), Test.Table.Orders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table.User, Equal(Test.Table.Orders.ID.Expr(), Test.Table.Users.ID.Expr())),
					Left(Test.Table.User, Equal(Test.Table.Users.String.Expr(), Test.Table.Orders.String.Expr())),
				).
				Where(
					And(
						Equal(Test.Table.Users.String.Expr(), Value("active")),
						Greater(Test.Table.Users.Number.Expr(), Value(2)),
						In(Test.Table.Users.ID.Expr(), Subquery[int64](
							NewSelect(Test.Table.User).
								Fields(
									Test.Table.Users.ID.Expr().As("uid"),
								).
								Where(
									Greater(Test.Table.Users.Number.Expr(), Value(10)),
								),
						)),
						Exists(Subquery[int](NewSelect(Test.Table.User).
							Fields(
								Test.Table.Users.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Table.Users.Number.Expr(), Test.Table.Orders.Number.Expr()),
									Equal(Test.Table.Users.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.Table.Orders.ID.Expr(),
					Test.Table.Orders.String.Expr(),
					Test.Table.Orders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Table.Users.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.Table.Orders.Number.Expr()),
					Asc(Test.Table.Orders.ID.Expr()),
				).
				Pagination(48, 0)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					_, _, err := builder.Build(stmt)
					if err != nil {
						b.Errorf("Build failed: %v", err)
					}
					i++
				}
			})
		})
	})
}
func Benchmark_Immutable_Single(b *testing.B) {
	builder := NewSQL()
	defer builder.Close()
	benchmarkAllDialects(b, func(b *testing.B, supportDialect *SupportDialect) {
		builder.SetDialect(supportDialect)
		b.Run("Simple", func(b *testing.B) {
			stmt := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.ID.Expr(),
				).
				Where(
					Equal(Test.Table.Users.Number.Expr(), Value(0)),
				)
			for i := 0; i < b.N; i++ {
				builder.Build(stmt)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			stmt := NewSelect(Test.Table.Order).
				Fields(
					Test.Table.Orders.ID.Expr().As("test_id"),
					Test.Table.Orders.String.Expr().As("test_string"),
					Test.Table.Orders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table.User).
						Fields(
							Count(Test.Table.Users.ID.Expr(), false),
						).
						Where(
							Equal(Test.Table.Users.Number.Expr(), Test.Table.Orders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table.User, Equal(Test.Table.Orders.ID.Expr(), Test.Table.Users.ID.Expr())),
					Left(Test.Table.User, Equal(Test.Table.Users.String.Expr(), Test.Table.Users.String.Expr())),
				).
				Where(
					And(
						Equal(Test.Table.Orders.String.Expr(), Value("active")),
						Greater(Test.Table.Orders.Number.Expr(), Value(2)),
						In(Test.Table.Orders.ID.Expr(), Subquery[int64](NewSelect(Test.Table.User).
							Fields(
								Test.Table.Users.ID.Expr().As("uid"),
							).
							Where(
								Greater(Test.Table.Users.Number.Expr(), Value(10)),
							),
						)),
						Exists(Subquery[int](NewSelect(Test.Table.User).
							Fields(
								Test.Table.Users.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Table.Users.Number.Expr(), Test.Table.Orders.Number.Expr()),
									Equal(Test.Table.Users.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.Table.Orders.ID.Expr(),
					Test.Table.Orders.String.Expr(),
					Test.Table.Orders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Table.Users.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.Table.Orders.Number.Expr()),
					Asc(Test.Table.Orders.ID.Expr()),
				).
				Pagination(48, 0)
			for i := 0; i < b.N; i++ {
				builder.Build(stmt)
			}
		})
	})
}
func Benchmark_Mutable_Multi(b *testing.B) {
	benchmarkAllDialects(b, func(b *testing.B, supportDialect *SupportDialect) {
		builder := NewSQL(
			WithDialect(supportDialect),
			WithMutate(true),
		)
		defer builder.Close()
		b.Run("Simple", func(b *testing.B) {
			stmt := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.ID.Expr(),
				).
				Where(
					Equal(Test.Table.Users.Number.Expr(), Value(0)),
				)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _, err := builder.Build(stmt)
					if err != nil {
						b.Errorf("Build failed: %v", err)
					}
				}
			})
		})
		b.Run("Complex", func(b *testing.B) {
			stmt := NewSelect(Test.Table.Order).
				Fields(
					Test.Table.Orders.ID.Expr().As("test_id"),
					Test.Table.Orders.String.Expr().As("test_string"),
					Test.Table.Orders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table.User).
						Fields(
							Count(Test.Table.Users.ID.Expr(), false),
						).
						Where(
							Equal(Test.Table.Users.Number.Expr(), Test.Table.Orders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table.User, Equal(Test.Table.Orders.ID.Expr(), Test.Table.Users.ID.Expr())),
					Left(Test.Table.User, Equal(Test.Table.Users.String.Expr(), Test.Table.Orders.String.Expr())),
				).
				Where(
					And(
						Equal(Test.Table.Users.String.Expr(), Value("active")),
						Greater(Test.Table.Users.Number.Expr(), Value(2)),
						In(Test.Table.Users.ID.Expr(), Subquery[int64](
							NewSelect(Test.Table.User).
								Fields(
									Test.Table.Users.ID.Expr().As("uid"),
								).
								Where(
									Greater(Test.Table.Users.Number.Expr(), Value(10)),
								),
						)),
						Exists(Subquery[int](NewSelect(Test.Table.User).
							Fields(
								Test.Table.Users.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Table.Users.Number.Expr(), Test.Table.Orders.Number.Expr()),
									Equal(Test.Table.Users.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.Table.Orders.ID.Expr(),
					Test.Table.Orders.String.Expr(),
					Test.Table.Orders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Table.Users.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.Table.Orders.Number.Expr()),
					Asc(Test.Table.Orders.ID.Expr()),
				).
				Pagination(48, 0)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					_, _, err := builder.Build(stmt)
					if err != nil {
						b.Errorf("Build failed: %v", err)
					}
					i++
				}
			})
		})
	})
}
func Benchmark_Mutable_Single(b *testing.B) {
	benchmarkAllDialects(b, func(b *testing.B, supportDialect *SupportDialect) {
		builder := NewSQL(
			WithDialect(supportDialect),
			WithMutate(true),
		)
		defer builder.Close()
		b.Run("Simple", func(b *testing.B) {
			stmt := NewSelect(Test.Table.User).
				Fields(
					Test.Table.Users.ID.Expr(),
				).
				Where(
					Equal(Test.Table.Users.Number.Expr(), Value(0)),
				)
			for i := 0; i < b.N; i++ {
				builder.Build(stmt)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			stmt := NewSelect(Test.Table.Order).
				Fields(
					Test.Table.Orders.ID.Expr().As("test_id"),
					Test.Table.Orders.String.Expr().As("test_string"),
					Test.Table.Orders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table.User).
						Fields(
							Count(Test.Table.Users.ID.Expr(), false),
						).
						Where(
							Equal(Test.Table.Users.Number.Expr(), Test.Table.Orders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table.User, Equal(Test.Table.Orders.ID.Expr(), Test.Table.Users.ID.Expr())),
					Left(Test.Table.User, Equal(Test.Table.Users.String.Expr(), Test.Table.Users.String.Expr())),
				).
				Where(
					And(
						Equal(Test.Table.Orders.String.Expr(), Value("active")),
						Greater(Test.Table.Orders.Number.Expr(), Value(2)),
						In(Test.Table.Orders.ID.Expr(), Subquery[int64](NewSelect(Test.Table.User).
							Fields(
								Test.Table.Users.ID.Expr().As("uid"),
							).
							Where(
								Greater(Test.Table.Users.Number.Expr(), Value(10)),
							),
						)),
						Exists(Subquery[int](NewSelect(Test.Table.User).
							Fields(
								Test.Table.Users.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Table.Users.Number.Expr(), Test.Table.Orders.Number.Expr()),
									Equal(Test.Table.Users.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.Table.Orders.ID.Expr(),
					Test.Table.Orders.String.Expr(),
					Test.Table.Orders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Table.Users.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.Table.Orders.Number.Expr()),
					Asc(Test.Table.Orders.ID.Expr()),
				).
				Pagination(48, 0)
			for i := 0; i < b.N; i++ {
				builder.Build(stmt)
			}
		})
	})
}

// Приватные функции
func benchmarkAllDialects(b *testing.B, testFunc func(b *testing.B, supportDialect *SupportDialect)) {
	for _, supportDialect := range listSupportDialects {
		currentDialect := supportDialect
		b.Run(currentDialect.name, func(b *testing.B) {
			testFunc(b, currentDialect)
		})
	}
}
