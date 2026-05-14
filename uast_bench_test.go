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
					stmtSelect := NewSelect(Users.Table).Field(Users.ID).Where(Equal(Users.Age, Value(i%1000)))
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
						Field(Count(Orders.ID, false)).
						Where(
							Equal(Orders.UserID, Users.ID),
						)
					queryInSub := NewSelect(Orders.Table).
						Field(Orders.UserID.As("uid")).
						Where(
							Greater(Orders.Amount, Value(10)),
						)
					queryExistsSub := NewSelect(Levels.Table).
						Field(Levels.ID).
						Where(
							And(
								Equal(Levels.UserID, Users.ID),
								Equal(Levels.Status, Value("active")),
							),
						)
					stmtSelect := NewSelect(Users.Table).
						Field(
							Users.ID.As("user_id"),
							Users.Name.As("user_name"),
							Users.Email.As("user_email"),
							Users.Age.As("user_age"),
							Users.Status.As("user_status"),
							Subquery[int](querySub).As("order_count"),
							Subquery[int](NewSelect(Orders.Table).
								Field(Sum(Orders.Amount, false)).
								Where(
									Equal(Orders.UserID, Users.ID),
								),
							).As("total_spent"),
						).
						Join(
							Inner(Levels.Table, Equal(Users.ID, Levels.UserID)),
							Left(Categories.Table, Equal(Categories.Type, Users.Status)),
						).
						Where(And(
							Equal(Users.Status, Value("active")),
							Greater(Users.Age, Value(18)),
							In(Users.ID, Subquery[int64](queryInSub)),
							Exists(Subquery[int](queryExistsSub)),
							NotExists(Subquery[int](NewSelect(Users.Table).
								Field(ConstIntOne()).
								Where(
									And(
										Equal(Users.ID, Users.ID), IsNull(Users.Email),
									),
								),
							)),
						)).
						GroupBy(
							Users.ID,
							Users.Name,
							Users.Email,
							Users.Age,
							Users.Status,
						).
						Having(
							And(
								Greater(Subquery[int64](NewSelect(Orders.Table).
									Field(Count(Orders.ID, false)).
									Where(Equal(Orders.UserID, Users.ID)),
								), Value(int64(0))),
								Greater(Subquery[int64](NewSelect(Orders.Table).
									Field(Sum(Orders.Amount, false)).
									Where(Equal(Orders.UserID, Users.ID)),
								), Value(int64(100))),
							),
						).
						OrderBy(
							Desc(Subquery[int](NewSelect(Orders.Table).
								Field(Sum(Orders.Amount, false)).
								Where(Equal(Orders.UserID, Users.ID)),
							)),
							Asc(Users.Name),
							Asc(Users.Age),
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
					Field(Users.ID).
					Where(Equal(Users.Age, Value(i%1000)))
				builder.Build(stmtSelect)
			}
		})
		b.Run("Complex", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				querySub := NewSelect(Orders.Table).
					Field(Count(Orders.ID, false)).
					Where(
						Equal(Orders.UserID, Users.ID),
					)
				queryInSub := NewSelect(Orders.Table).
					Field(Orders.UserID.As("uid")).
					Where(
						Greater(Orders.Amount, Value(10)),
					)
				queryExistsSub := NewSelect(Levels.Table).
					Field(Levels.ID).
					Where(
						And(
							Equal(Levels.UserID, Users.ID),
							Equal(Levels.Status, Value("active")),
						),
					)
				stmtSelect := NewSelect(Users.Table).
					Field(
						Users.ID.As("user_id"),
						Users.Name.As("user_name"),
						Users.Email.As("user_email"),
						Users.Age.As("user_age"),
						Users.Status.As("user_status"),
						Subquery[int](querySub).As("order_count"),
						Subquery[int](NewSelect(Orders.Table).
							Field(Sum(Orders.Amount, false)).
							Where(
								Equal(Orders.UserID, Users.ID),
							),
						).As("total_spent"),
					).
					Join(
						Inner(Levels.Table, Equal(Users.ID, Levels.UserID)),
						Left(Categories.Table, Equal(Categories.Type, Users.Status)),
					).
					Where(And(
						Equal(Users.Status, Value("active")),
						Greater(Users.Age, Value(18)),
						In(Users.ID, Subquery[int64](queryInSub)),
						Exists(Subquery[int](queryExistsSub)),
						NotExists(Subquery[int](NewSelect(Users.Table).
							Field(ConstIntOne()).
							Where(
								And(
									Equal(Users.ID, Users.ID), IsNull(Users.Email),
								),
							),
						)),
					)).
					GroupBy(
						Users.ID,
						Users.Name,
						Users.Email,
						Users.Age,
						Users.Status,
					).
					Having(
						And(
							Greater(Subquery[int64](NewSelect(Orders.Table).
								Field(Count(Orders.ID, false)).
								Where(Equal(Orders.UserID, Users.ID)),
							), Value(int64(0))),
							Greater(Subquery[int64](NewSelect(Orders.Table).
								Field(Sum(Orders.Amount, false)).
								Where(Equal(Orders.UserID, Users.ID)),
							), Value(int64(100))),
						),
					).
					OrderBy(
						Desc(Subquery[int](NewSelect(Orders.Table).
							Field(Sum(Orders.Amount, false)).
							Where(Equal(Orders.UserID, Users.ID)),
						)),
						Asc(Users.Name),
						Asc(Users.Age),
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
