package uast

import (
	"testing"
)

// Бенчмарки компонентов
func Benchmark_Select_Multi(b *testing.B) {
	benchmarkAllDialects(b, func(b *testing.B, supportDialect *SupportDialect) {
		builder := NewSQL(
			WithDialect(supportDialect),
		)
		defer builder.Close()
		b.Run("Simple", func(b *testing.B) {
			stmt := NewSelect(Test.Table).
				Field(
					Test.Column.ID,
				).
				Where(
					Equal(Test.Column.Number, Value(0)),
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
				Field(
					Data.Column.ID.As("test_id"),
					Data.Column.String.As("test_string"),
					Data.Column.Number.As("test_number"),
					Subquery[int](NewSelect(Test.Table).
						Field(
							Count(Test.Column.ID, false),
						).
						Where(
							Equal(Test.Column.Number, Data.Column.Number),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table, Equal(Data.Column.ID, Test.Column.ID)),
					Left(Test.Table, Equal(Test.Column.String, Data.Column.String)),
				).
				Where(
					And(
						Equal(Test.Column.String, Value("active")),
						Greater(Test.Column.Number, Value(2)),
						In(Test.Column.ID, Subquery[int64](
							NewSelect(Test.Table).
								Field(
									Test.Column.ID.As("uid"),
								).
								Where(
									Greater(Test.Column.Number, Value(10)),
								),
						)),
						Exists(Subquery[int](NewSelect(Test.Table).
							Field(
								Test.Column.ID,
							).
							Where(
								And(
									Equal(Test.Column.Number, Data.Column.Number),
									Equal(Test.Column.String, Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Data.Column.ID,
					Data.Column.String,
					Data.Column.Number,
				).
				Having(
					Greater(Count(Test.Column.ID, false), Value[int64](0)),
				).
				OrderBy(
					Desc(Data.Column.Number),
					Asc(Data.Column.ID),
				).
				Limit(48)
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
func Benchmark_Select_Single(b *testing.B) {
	benchmarkAllDialects(b, func(b *testing.B, supportDialect *SupportDialect) {
		builder := NewSQL(
			WithDialect(supportDialect),
		)
		defer builder.Close()
		b.Run("Simple", func(b *testing.B) {
			stmt := NewSelect(Test.Table).
				Field(
					Test.Column.ID,
				).
				Where(
					Equal(Test.Column.Number, Value(0)),
				)
			for i := 0; i < b.N; i++ {
				builder.Build(stmt)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			stmt := NewSelect(Data.Table).
				Field(
					Data.Column.ID.As("test_id"),
					Data.Column.String.As("test_string"),
					Data.Column.Number.As("test_number"),
					Subquery[int](NewSelect(Test.Table).
						Field(
							Count(Test.Column.ID, false),
						).
						Where(
							Equal(Test.Column.Number, Data.Column.Number),
						),
					).As("sub_count"),
				).
				Join(
					Inner(Test.Table, Equal(Data.Column.ID, Test.Column.ID)),
					Left(Test.Table, Equal(Test.Column.String, Test.Column.String)),
				).
				Where(
					And(
						Equal(Data.Column.String, Value("active")),
						Greater(Data.Column.Number, Value(2)),
						In(Data.Column.ID, Subquery[int64](NewSelect(Test.Table).
							Field(
								Test.Column.ID.As("uid"),
							).
							Where(
								Greater(Test.Column.Number, Value(10)),
							),
						)),
						Exists(Subquery[int](NewSelect(Test.Table).
							Field(
								Test.Column.ID,
							).
							Where(
								And(
									Equal(Test.Column.Number, Data.Column.Number),
									Equal(Test.Column.String, Value("active")),
								),
							),
						)),
					),
				).
				GroupBy(
					Data.Column.ID,
					Data.Column.String,
					Data.Column.Number,
				).
				Having(
					Greater(Count(Test.Column.ID, false), Value[int64](0)),
				).
				OrderBy(
					Desc(Data.Column.Number),
					Asc(Data.Column.ID),
				).
				Limit(48)
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
