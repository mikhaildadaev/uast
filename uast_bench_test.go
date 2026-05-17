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
					stmtSelect := NewSelect(Users.Table).Field(Users.Column.ID).Where(Equal(Users.Column.Age, Value(i%1000)))
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
					querySub := NewSelect(Orders.Table).
						Field(Count(Orders.Column.ID, false)).
						Where(
							Equal(Orders.Column.UserID, Users.Column.ID),
						)
					queryInSub := NewSelect(Orders.Table).
						Field(Orders.Column.UserID.As("uid")).
						Where(
							Greater(Orders.Column.Amount, Value(10)),
						)
					queryExistsSub := NewSelect(Levels.Table).
						Field(Levels.Column.ID).
						Where(
							And(
								Equal(Levels.Column.UserID, Users.Column.ID),
								Equal(Levels.Column.Status, Value("active")),
							),
						)
					stmtSelect := NewSelect(Users.Table).
						Field(
							Users.Column.ID.As("user_id"),
							Users.Column.Name.As("user_name"),
							Users.Column.Email.As("user_email"),
							Users.Column.Age.As("user_age"),
							Users.Column.Status.As("user_status"),
							Subquery[int](querySub).As("order_count"),
							Subquery[int](NewSelect(Orders.Table).
								Field(Sum(Orders.Column.Amount, false)).
								Where(
									Equal(Orders.Column.UserID, Users.Column.ID),
								),
							).As("total_spent"),
						).
						Join(
							Inner(Levels.Table, Equal(Users.Column.ID, Levels.Column.UserID)),
							Left(Categories.Table, Equal(Categories.Column.Type, Users.Column.Status)),
						).
						Where(And(
							Equal(Users.Column.Status, Value("active")),
							Greater(Users.Column.Age, Value(18)),
							In(Users.Column.ID, Subquery[int64](queryInSub)),
							Exists(Subquery[int](queryExistsSub)),
							NotExists(Subquery[int](NewSelect(Users.Table).
								Field(ConstIntOne()).
								Where(
									And(
										Equal(Users.Column.ID, Users.Column.ID), IsNull(Users.Column.Email),
									),
								),
							)),
						)).
						GroupBy(
							Users.Column.ID,
							Users.Column.Name,
							Users.Column.Email,
							Users.Column.Age,
							Users.Column.Status,
						).
						Having(
							And(
								Greater(Subquery[int64](NewSelect(Orders.Table).
									Field(Count(Orders.Column.ID, false)).
									Where(Equal(Orders.Column.UserID, Users.Column.ID)),
								), Value(int64(0))),
								Greater(Subquery[int64](NewSelect(Orders.Table).
									Field(Sum(Orders.Column.Amount, false)).
									Where(Equal(Orders.Column.UserID, Users.Column.ID)),
								), Value(int64(100))),
							),
						).
						OrderBy(
							Desc(Subquery[int](NewSelect(Orders.Table).
								Field(Sum(Orders.Column.Amount, false)).
								Where(Equal(Orders.Column.UserID, Users.Column.ID)),
							)),
							Asc(Users.Column.Name),
							Asc(Users.Column.Age),
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
				stmtSelect := NewSelect(Users.Table).
					Field(Users.Column.ID).
					Where(Equal(Users.Column.Age, Value(i%1000)))
				builder.Build(stmtSelect)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				querySub := NewSelect(Orders.Table).
					Field(Count(Orders.Column.ID, false)).
					Where(
						Equal(Orders.Column.UserID, Users.Column.ID),
					)
				queryInSub := NewSelect(Orders.Table).
					Field(Orders.Column.UserID.As("uid")).
					Where(
						Greater(Orders.Column.Amount, Value(10)),
					)
				queryExistsSub := NewSelect(Levels.Table).
					Field(Levels.Column.ID).
					Where(
						And(
							Equal(Levels.Column.UserID, Users.Column.ID),
							Equal(Levels.Column.Status, Value("active")),
						),
					)
				stmtSelect := NewSelect(Users.Table).
					Field(
						Users.Column.ID.As("user_id"),
						Users.Column.Name.As("user_name"),
						Users.Column.Email.As("user_email"),
						Users.Column.Age.As("user_age"),
						Users.Column.Status.As("user_status"),
						Subquery[int](querySub).As("order_count"),
						Subquery[int](NewSelect(Orders.Table).
							Field(Sum(Orders.Column.Amount, false)).
							Where(
								Equal(Orders.Column.UserID, Users.Column.ID),
							),
						).As("total_spent"),
					).
					Join(
						Inner(Levels.Table, Equal(Users.Column.ID, Levels.Column.UserID)),
						Left(Categories.Table, Equal(Categories.Column.Type, Users.Column.Status)),
					).
					Where(And(
						Equal(Users.Column.Status, Value("active")),
						Greater(Users.Column.Age, Value(18)),
						In(Users.Column.ID, Subquery[int64](queryInSub)),
						Exists(Subquery[int](queryExistsSub)),
						NotExists(Subquery[int](NewSelect(Users.Table).
							Field(ConstIntOne()).
							Where(
								And(
									Equal(Users.Column.ID, Users.Column.ID), IsNull(Users.Column.Email),
								),
							),
						)),
					)).
					GroupBy(
						Users.Column.ID,
						Users.Column.Name,
						Users.Column.Email,
						Users.Column.Age,
						Users.Column.Status,
					).
					Having(
						And(
							Greater(Subquery[int64](NewSelect(Orders.Table).
								Field(Count(Orders.Column.ID, false)).
								Where(Equal(Orders.Column.UserID, Users.Column.ID)),
							), Value(int64(0))),
							Greater(Subquery[int64](NewSelect(Orders.Table).
								Field(Sum(Orders.Column.Amount, false)).
								Where(Equal(Orders.Column.UserID, Users.Column.ID)),
							), Value(int64(100))),
						),
					).
					OrderBy(
						Desc(Subquery[int](NewSelect(Orders.Table).
							Field(Sum(Orders.Column.Amount, false)).
							Where(Equal(Orders.Column.UserID, Users.Column.ID)),
						)),
						Asc(Users.Column.Name),
						Asc(Users.Column.Age),
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
