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
			stmt := NewSelect(Test.Tables.Users).
				Fields(
					Test.Users.ID.Expr(),
				).
				Where(
					Equal(Test.Users.Number.Expr(), Value(0)),
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
			stmt := NewSelect(Test.Tables.Orders).
				Fields(
					Test.Orders.ID.Expr().As("test_id"),
					Test.Orders.String.Expr().As("test_string"),
					Test.Orders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Tables.Users).
						Fields(
							Count(Test.Users.ID.Expr(), false),
						).
						Where(
							Equal(Test.Users.Number.Expr(), Test.Orders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Tables.Users, Equal(Test.Orders.ID.Expr(), Test.Users.ID.Expr())),
					Left(Test.Tables.Users, Equal(Test.Users.String.Expr(), Test.Orders.String.Expr())),
				).
				Where(
					And(
						Equal(Test.Users.String.Expr(), Value("active")),
						Greater(Test.Users.Number.Expr(), Value(2)),
						In(Test.Users.ID.Expr(), Subquery[int64](
							NewSelect(Test.Tables.Users).
								Fields(
									Test.Users.ID.Expr().As("uid"),
								).
								Where(
									Greater(Test.Users.Number.Expr(), Value(10)),
								),
						)),
						Exists(Subquery[int](NewSelect(Test.Tables.Users).
							Fields(
								Test.Users.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Users.Number.Expr(), Test.Orders.Number.Expr()),
									Equal(Test.Users.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.Orders.ID.Expr(),
					Test.Orders.String.Expr(),
					Test.Orders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Users.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.Orders.Number.Expr()),
					Asc(Test.Orders.ID.Expr()),
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
			stmt := NewSelect(Test.Tables.Users).
				Fields(
					Test.Users.ID.Expr(),
				).
				Where(
					Equal(Test.Users.Number.Expr(), Value(0)),
				)
			for i := 0; i < b.N; i++ {
				builder.Build(stmt)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			stmt := NewSelect(Test.Tables.Orders).
				Fields(
					Test.Orders.ID.Expr().As("test_id"),
					Test.Orders.String.Expr().As("test_string"),
					Test.Orders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Tables.Users).
						Fields(
							Count(Test.Users.ID.Expr(), false),
						).
						Where(
							Equal(Test.Users.Number.Expr(), Test.Orders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Tables.Users, Equal(Test.Orders.ID.Expr(), Test.Users.ID.Expr())),
					Left(Test.Tables.Users, Equal(Test.Users.String.Expr(), Test.Users.String.Expr())),
				).
				Where(
					And(
						Equal(Test.Orders.String.Expr(), Value("active")),
						Greater(Test.Orders.Number.Expr(), Value(2)),
						In(Test.Orders.ID.Expr(), Subquery[int64](NewSelect(Test.Tables.Users).
							Fields(
								Test.Users.ID.Expr().As("uid"),
							).
							Where(
								Greater(Test.Users.Number.Expr(), Value(10)),
							),
						)),
						Exists(Subquery[int](NewSelect(Test.Tables.Users).
							Fields(
								Test.Users.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Users.Number.Expr(), Test.Orders.Number.Expr()),
									Equal(Test.Users.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.Orders.ID.Expr(),
					Test.Orders.String.Expr(),
					Test.Orders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Users.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.Orders.Number.Expr()),
					Asc(Test.Orders.ID.Expr()),
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
			stmt := NewSelect(Test.Tables.Users).
				Fields(
					Test.Users.ID.Expr(),
				).
				Where(
					Equal(Test.Users.Number.Expr(), Value(0)),
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
			stmt := NewSelect(Test.Tables.Orders).
				Fields(
					Test.Orders.ID.Expr().As("test_id"),
					Test.Orders.String.Expr().As("test_string"),
					Test.Orders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Tables.Users).
						Fields(
							Count(Test.Users.ID.Expr(), false),
						).
						Where(
							Equal(Test.Users.Number.Expr(), Test.Orders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Tables.Users, Equal(Test.Orders.ID.Expr(), Test.Users.ID.Expr())),
					Left(Test.Tables.Users, Equal(Test.Users.String.Expr(), Test.Orders.String.Expr())),
				).
				Where(
					And(
						Equal(Test.Users.String.Expr(), Value("active")),
						Greater(Test.Users.Number.Expr(), Value(2)),
						In(Test.Users.ID.Expr(), Subquery[int64](
							NewSelect(Test.Tables.Users).
								Fields(
									Test.Users.ID.Expr().As("uid"),
								).
								Where(
									Greater(Test.Users.Number.Expr(), Value(10)),
								),
						)),
						Exists(Subquery[int](NewSelect(Test.Tables.Users).
							Fields(
								Test.Users.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Users.Number.Expr(), Test.Orders.Number.Expr()),
									Equal(Test.Users.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.Orders.ID.Expr(),
					Test.Orders.String.Expr(),
					Test.Orders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Users.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.Orders.Number.Expr()),
					Asc(Test.Orders.ID.Expr()),
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
			stmt := NewSelect(Test.Tables.Users).
				Fields(
					Test.Users.ID.Expr(),
				).
				Where(
					Equal(Test.Users.Number.Expr(), Value(0)),
				)
			for i := 0; i < b.N; i++ {
				builder.Build(stmt)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			stmt := NewSelect(Test.Tables.Orders).
				Fields(
					Test.Orders.ID.Expr().As("test_id"),
					Test.Orders.String.Expr().As("test_string"),
					Test.Orders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Tables.Users).
						Fields(
							Count(Test.Users.ID.Expr(), false),
						).
						Where(
							Equal(Test.Users.Number.Expr(), Test.Orders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Tables.Users, Equal(Test.Orders.ID.Expr(), Test.Users.ID.Expr())),
					Left(Test.Tables.Users, Equal(Test.Users.String.Expr(), Test.Users.String.Expr())),
				).
				Where(
					And(
						Equal(Test.Orders.String.Expr(), Value("active")),
						Greater(Test.Orders.Number.Expr(), Value(2)),
						In(Test.Orders.ID.Expr(), Subquery[int64](NewSelect(Test.Tables.Users).
							Fields(
								Test.Users.ID.Expr().As("uid"),
							).
							Where(
								Greater(Test.Users.Number.Expr(), Value(10)),
							),
						)),
						Exists(Subquery[int](NewSelect(Test.Tables.Users).
							Fields(
								Test.Users.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Users.Number.Expr(), Test.Orders.Number.Expr()),
									Equal(Test.Users.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.Orders.ID.Expr(),
					Test.Orders.String.Expr(),
					Test.Orders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Users.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.Orders.Number.Expr()),
					Asc(Test.Orders.ID.Expr()),
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
