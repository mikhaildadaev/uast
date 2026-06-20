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
			stmt := NewSelect(Test.Table.Users).
				Fields(
					Test.TableUsers.ID.Expr(),
				).
				Where(
					Equal(Test.TableUsers.Number.Expr(), Value(0)),
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
			stmt := NewSelect(Test.Table.Orders).
				Fields(
					Test.TableOrders.ID.Expr().As("test_id"),
					Test.TableOrders.String.Expr().As("test_string"),
					Test.TableOrders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table.Users).
						Fields(
							Count(Test.TableUsers.ID.Expr(), false),
						).
						Where(
							Equal(Test.TableUsers.Number.Expr(), Test.TableOrders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table.Users, Equal(Test.TableOrders.ID.Expr(), Test.TableUsers.ID.Expr())),
					Left(Test.Table.Users, Equal(Test.TableUsers.String.Expr(), Test.TableOrders.String.Expr())),
				).
				Where(
					And(
						Equal(Test.TableUsers.String.Expr(), Value("active")),
						Greater(Test.TableUsers.Number.Expr(), Value(2)),
						In(Test.TableUsers.ID.Expr(), Subquery[int64](
							NewSelect(Test.Table.Users).
								Fields(
									Test.TableUsers.ID.Expr().As("uid"),
								).
								Where(
									Greater(Test.TableUsers.Number.Expr(), Value(10)),
								),
						)),
						Exists(Subquery[int](NewSelect(Test.Table.Users).
							Fields(
								Test.TableUsers.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.TableUsers.Number.Expr(), Test.TableOrders.Number.Expr()),
									Equal(Test.TableUsers.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.TableOrders.ID.Expr(),
					Test.TableOrders.String.Expr(),
					Test.TableOrders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.TableUsers.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.TableOrders.Number.Expr()),
					Asc(Test.TableOrders.ID.Expr()),
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
			stmt := NewSelect(Test.Table.Users).
				Fields(
					Test.TableUsers.ID.Expr(),
				).
				Where(
					Equal(Test.TableUsers.Number.Expr(), Value(0)),
				)
			for i := 0; i < b.N; i++ {
				builder.Build(stmt)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			stmt := NewSelect(Test.Table.Orders).
				Fields(
					Test.TableOrders.ID.Expr().As("test_id"),
					Test.TableOrders.String.Expr().As("test_string"),
					Test.TableOrders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table.Users).
						Fields(
							Count(Test.TableUsers.ID.Expr(), false),
						).
						Where(
							Equal(Test.TableUsers.Number.Expr(), Test.TableOrders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table.Users, Equal(Test.TableOrders.ID.Expr(), Test.TableUsers.ID.Expr())),
					Left(Test.Table.Users, Equal(Test.TableUsers.String.Expr(), Test.TableUsers.String.Expr())),
				).
				Where(
					And(
						Equal(Test.TableOrders.String.Expr(), Value("active")),
						Greater(Test.TableOrders.Number.Expr(), Value(2)),
						In(Test.TableOrders.ID.Expr(), Subquery[int64](NewSelect(Test.Table.Users).
							Fields(
								Test.TableUsers.ID.Expr().As("uid"),
							).
							Where(
								Greater(Test.TableUsers.Number.Expr(), Value(10)),
							),
						)),
						Exists(Subquery[int](NewSelect(Test.Table.Users).
							Fields(
								Test.TableUsers.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.TableUsers.Number.Expr(), Test.TableOrders.Number.Expr()),
									Equal(Test.TableUsers.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.TableOrders.ID.Expr(),
					Test.TableOrders.String.Expr(),
					Test.TableOrders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.TableUsers.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.TableOrders.Number.Expr()),
					Asc(Test.TableOrders.ID.Expr()),
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
			stmt := NewSelect(Test.Table.Users).
				Fields(
					Test.TableUsers.ID.Expr(),
				).
				Where(
					Equal(Test.TableUsers.Number.Expr(), Value(0)),
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
			stmt := NewSelect(Test.Table.Orders).
				Fields(
					Test.TableOrders.ID.Expr().As("test_id"),
					Test.TableOrders.String.Expr().As("test_string"),
					Test.TableOrders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table.Users).
						Fields(
							Count(Test.TableUsers.ID.Expr(), false),
						).
						Where(
							Equal(Test.TableUsers.Number.Expr(), Test.TableOrders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table.Users, Equal(Test.TableOrders.ID.Expr(), Test.TableUsers.ID.Expr())),
					Left(Test.Table.Users, Equal(Test.TableUsers.String.Expr(), Test.TableOrders.String.Expr())),
				).
				Where(
					And(
						Equal(Test.TableUsers.String.Expr(), Value("active")),
						Greater(Test.TableUsers.Number.Expr(), Value(2)),
						In(Test.TableUsers.ID.Expr(), Subquery[int64](
							NewSelect(Test.Table.Users).
								Fields(
									Test.TableUsers.ID.Expr().As("uid"),
								).
								Where(
									Greater(Test.TableUsers.Number.Expr(), Value(10)),
								),
						)),
						Exists(Subquery[int](NewSelect(Test.Table.Users).
							Fields(
								Test.TableUsers.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.TableUsers.Number.Expr(), Test.TableOrders.Number.Expr()),
									Equal(Test.TableUsers.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.TableOrders.ID.Expr(),
					Test.TableOrders.String.Expr(),
					Test.TableOrders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.TableUsers.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.TableOrders.Number.Expr()),
					Asc(Test.TableOrders.ID.Expr()),
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
			stmt := NewSelect(Test.Table.Users).
				Fields(
					Test.TableUsers.ID.Expr(),
				).
				Where(
					Equal(Test.TableUsers.Number.Expr(), Value(0)),
				)
			for i := 0; i < b.N; i++ {
				builder.Build(stmt)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			stmt := NewSelect(Test.Table.Orders).
				Fields(
					Test.TableOrders.ID.Expr().As("test_id"),
					Test.TableOrders.String.Expr().As("test_string"),
					Test.TableOrders.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table.Users).
						Fields(
							Count(Test.TableUsers.ID.Expr(), false),
						).
						Where(
							Equal(Test.TableUsers.Number.Expr(), Test.TableOrders.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table.Users, Equal(Test.TableOrders.ID.Expr(), Test.TableUsers.ID.Expr())),
					Left(Test.Table.Users, Equal(Test.TableUsers.String.Expr(), Test.TableUsers.String.Expr())),
				).
				Where(
					And(
						Equal(Test.TableOrders.String.Expr(), Value("active")),
						Greater(Test.TableOrders.Number.Expr(), Value(2)),
						In(Test.TableOrders.ID.Expr(), Subquery[int64](NewSelect(Test.Table.Users).
							Fields(
								Test.TableUsers.ID.Expr().As("uid"),
							).
							Where(
								Greater(Test.TableUsers.Number.Expr(), Value(10)),
							),
						)),
						Exists(Subquery[int](NewSelect(Test.Table.Users).
							Fields(
								Test.TableUsers.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.TableUsers.Number.Expr(), Test.TableOrders.Number.Expr()),
									Equal(Test.TableUsers.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Test.TableOrders.ID.Expr(),
					Test.TableOrders.String.Expr(),
					Test.TableOrders.Number.Expr(),
				).
				Having(
					Greater(Count(Test.TableUsers.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Test.TableOrders.Number.Expr()),
					Asc(Test.TableOrders.ID.Expr()),
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
