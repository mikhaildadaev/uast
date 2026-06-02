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
			stmt := NewSelect(Test.Table).
				Fields(
					Test.Column.ID.Expr(),
				).
				Where(
					Equal(Test.Column.Number.Expr(), Value(0)),
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
			stmt := NewSelect(Data.Table).
				Fields(
					Data.Column.ID.Expr().As("test_id"),
					Data.Column.String.Expr().As("test_string"),
					Data.Column.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table).
						Fields(
							Count(Test.Column.ID.Expr(), false),
						).
						Where(
							Equal(Test.Column.Number.Expr(), Data.Column.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table, Equal(Data.Column.ID.Expr(), Test.Column.ID.Expr())),
					Left(Test.Table, Equal(Test.Column.String.Expr(), Data.Column.String.Expr())),
				).
				Where(
					And(
						Equal(Test.Column.String.Expr(), Value("active")),
						Greater(Test.Column.Number.Expr(), Value(2)),
						In(Test.Column.ID.Expr(), Subquery[int64](
							NewSelect(Test.Table).
								Fields(
									Test.Column.ID.Expr().As("uid"),
								).
								Where(
									Greater(Test.Column.Number.Expr(), Value(10)),
								),
						)),
						Exists(Subquery[int](NewSelect(Test.Table).
							Fields(
								Test.Column.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Column.Number.Expr(), Data.Column.Number.Expr()),
									Equal(Test.Column.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Data.Column.ID.Expr(),
					Data.Column.String.Expr(),
					Data.Column.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Column.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Data.Column.Number.Expr()),
					Asc(Data.Column.ID.Expr()),
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
			stmt := NewSelect(Test.Table).
				Fields(
					Test.Column.ID.Expr(),
				).
				Where(
					Equal(Test.Column.Number.Expr(), Value(0)),
				)
			for i := 0; i < b.N; i++ {
				builder.Build(stmt)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			stmt := NewSelect(Data.Table).
				Fields(
					Data.Column.ID.Expr().As("test_id"),
					Data.Column.String.Expr().As("test_string"),
					Data.Column.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table).
						Fields(
							Count(Test.Column.ID.Expr(), false),
						).
						Where(
							Equal(Test.Column.Number.Expr(), Data.Column.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table, Equal(Data.Column.ID.Expr(), Test.Column.ID.Expr())),
					Left(Test.Table, Equal(Test.Column.String.Expr(), Test.Column.String.Expr())),
				).
				Where(
					And(
						Equal(Data.Column.String.Expr(), Value("active")),
						Greater(Data.Column.Number.Expr(), Value(2)),
						In(Data.Column.ID.Expr(), Subquery[int64](NewSelect(Test.Table).
							Fields(
								Test.Column.ID.Expr().As("uid"),
							).
							Where(
								Greater(Test.Column.Number.Expr(), Value(10)),
							),
						)),
						Exists(Subquery[int](NewSelect(Test.Table).
							Fields(
								Test.Column.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Column.Number.Expr(), Data.Column.Number.Expr()),
									Equal(Test.Column.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Data.Column.ID.Expr(),
					Data.Column.String.Expr(),
					Data.Column.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Column.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Data.Column.Number.Expr()),
					Asc(Data.Column.ID.Expr()),
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
			stmt := NewSelect(Test.Table).
				Fields(
					Test.Column.ID.Expr(),
				).
				Where(
					Equal(Test.Column.Number.Expr(), Value(0)),
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
			stmt := NewSelect(Data.Table).
				Fields(
					Data.Column.ID.Expr().As("test_id"),
					Data.Column.String.Expr().As("test_string"),
					Data.Column.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table).
						Fields(
							Count(Test.Column.ID.Expr(), false),
						).
						Where(
							Equal(Test.Column.Number.Expr(), Data.Column.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table, Equal(Data.Column.ID.Expr(), Test.Column.ID.Expr())),
					Left(Test.Table, Equal(Test.Column.String.Expr(), Data.Column.String.Expr())),
				).
				Where(
					And(
						Equal(Test.Column.String.Expr(), Value("active")),
						Greater(Test.Column.Number.Expr(), Value(2)),
						In(Test.Column.ID.Expr(), Subquery[int64](
							NewSelect(Test.Table).
								Fields(
									Test.Column.ID.Expr().As("uid"),
								).
								Where(
									Greater(Test.Column.Number.Expr(), Value(10)),
								),
						)),
						Exists(Subquery[int](NewSelect(Test.Table).
							Fields(
								Test.Column.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Column.Number.Expr(), Data.Column.Number.Expr()),
									Equal(Test.Column.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Data.Column.ID.Expr(),
					Data.Column.String.Expr(),
					Data.Column.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Column.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Data.Column.Number.Expr()),
					Asc(Data.Column.ID.Expr()),
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
			stmt := NewSelect(Test.Table).
				Fields(
					Test.Column.ID.Expr(),
				).
				Where(
					Equal(Test.Column.Number.Expr(), Value(0)),
				)
			for i := 0; i < b.N; i++ {
				builder.Build(stmt)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			stmt := NewSelect(Data.Table).
				Fields(
					Data.Column.ID.Expr().As("test_id"),
					Data.Column.String.Expr().As("test_string"),
					Data.Column.Number.Expr().As("test_number"),
					Subquery[int](NewSelect(Test.Table).
						Fields(
							Count(Test.Column.ID.Expr(), false),
						).
						Where(
							Equal(Test.Column.Number.Expr(), Data.Column.Number.Expr()),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table, Equal(Data.Column.ID.Expr(), Test.Column.ID.Expr())),
					Left(Test.Table, Equal(Test.Column.String.Expr(), Test.Column.String.Expr())),
				).
				Where(
					And(
						Equal(Data.Column.String.Expr(), Value("active")),
						Greater(Data.Column.Number.Expr(), Value(2)),
						In(Data.Column.ID.Expr(), Subquery[int64](NewSelect(Test.Table).
							Fields(
								Test.Column.ID.Expr().As("uid"),
							).
							Where(
								Greater(Test.Column.Number.Expr(), Value(10)),
							),
						)),
						Exists(Subquery[int](NewSelect(Test.Table).
							Fields(
								Test.Column.ID.Expr(),
							).
							Where(
								And(
									Equal(Test.Column.Number.Expr(), Data.Column.Number.Expr()),
									Equal(Test.Column.String.Expr(), Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Data.Column.ID.Expr(),
					Data.Column.String.Expr(),
					Data.Column.Number.Expr(),
				).
				Having(
					Greater(Count(Test.Column.ID.Expr(), false), Value[int64](0)),
				).
				OrderBy(
					Desc(Data.Column.Number.Expr()),
					Asc(Data.Column.ID.Expr()),
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
