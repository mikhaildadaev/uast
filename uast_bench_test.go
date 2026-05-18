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
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					stmtSelect := NewSelect(Test.Table).
						Field(
							Test.Column.ID,
						).
						Where(
							Equal(Test.Column.Number, Value(i%1000)),
						)
					_, _, err := builder.Build(stmtSelect)
					if err != nil {
						b.Errorf("Build failed: %v", err)
					}
					i++
				}
			})
		})
		b.Run("Complex", func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					querySub := NewSelect(Test1.Table).
						Field(
							Count(Test1.Column.ID, false),
						).
						Where(
							Equal(Test1.Column.Number, Test.Column.Number),
						)
					queryInSub := NewSelect(Test1.Table).
						Field(
							Test1.Column.ID.As("uid"),
						).
						Where(
							Greater(Test1.Column.Number, Value(10)),
						)
					queryExistsSub := NewSelect(Test2.Table).
						Field(
							Test2.Column.ID,
						).
						Where(
							And(
								Equal(Test2.Column.Number, Test.Column.Number),
								Equal(Test2.Column.String, Value("active")),
							),
						)
					stmtSelect := NewSelect(Test.Table).
						Field(
							Test.Column.ID.As("test_id"),
							Test.Column.Name.As("test_name"),
							Test.Column.String.As("test_string"),
							Test.Column.Number.As("test_number"),
							Subquery[int](querySub).As("sub_count"),
						).
						Join(
							Inner(Test1.Table, Equal(Test.Column.ID, Test1.Column.ID)),
							Left(Test2.Table, Equal(Test2.Column.String, Test.Column.String)),
						).
						Where(
							And(
								Equal(Test.Column.String, Value("active")),
								Greater(Test.Column.Number, Value(2)),
								In(Test.Column.ID, Subquery[int64](queryInSub)),
								Exists(Subquery[int](queryExistsSub)),
							),
						).
						GroupBy(
							Test.Column.ID,
							Test.Column.Name,
							Test.Column.String,
							Test.Column.Number,
						).
						Having(
							Greater(Count(Test1.Column.ID, false), Value[int64](0)),
						).
						OrderBy(
							Desc(Test.Column.Number),
							Asc(Test.Column.Name),
						).
						Limit(48)
					_, _, err := builder.Build(stmtSelect)
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
			for i := 0; i < b.N; i++ {
				stmtSelect := NewSelect(Test.Table).
					Field(
						Test.Column.ID,
					).
					Where(
						Equal(Test.Column.Number, Value(i%1000)),
					)
				builder.Build(stmtSelect)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				querySub := NewSelect(Test1.Table).
					Field(
						Count(Test1.Column.ID, false),
					).
					Where(
						Equal(Test1.Column.Number, Test.Column.Number),
					)
				queryInSub := NewSelect(Test1.Table).
					Field(
						Test1.Column.ID.As("uid"),
					).
					Where(
						Greater(Test1.Column.Number, Value(10)),
					)
				queryExistsSub := NewSelect(Test2.Table).
					Field(
						Test2.Column.ID,
					).
					Where(
						And(
							Equal(Test2.Column.Number, Test.Column.Number),
							Equal(Test2.Column.String, Value("active")),
						),
					)
				stmtSelect := NewSelect(Test.Table).
					Field(
						Test.Column.ID.As("test_id"),
						Test.Column.Name.As("test_name"),
						Test.Column.String.As("test_string"),
						Test.Column.Number.As("test_number"),
						Subquery[int](querySub).As("sub_count"),
					).
					Join(
						Inner(Test1.Table, Equal(Test.Column.ID, Test1.Column.ID)),
						Left(Test2.Table, Equal(Test2.Column.String, Test.Column.String)),
					).
					Where(
						And(
							Equal(Test.Column.String, Value("active")),
							Greater(Test.Column.Number, Value(2)),
							In(Test.Column.ID, Subquery[int64](queryInSub)),
							Exists(Subquery[int](queryExistsSub)),
						),
					).
					GroupBy(
						Test.Column.ID,
						Test.Column.Name,
						Test.Column.String,
						Test.Column.Number,
					).
					Having(
						Greater(Count(Test1.Column.ID, false), Value[int64](0)),
					).
					OrderBy(
						Desc(Test.Column.Number),
						Asc(Test.Column.Name),
					).
					Limit(48)
				builder.Build(stmtSelect)
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
