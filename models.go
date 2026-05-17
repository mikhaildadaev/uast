package uast

import (
	"time"
)

// Публичные переменные
var Categories = struct {
	Column struct {
		ID   *exprColumn[int64]
		Name *exprColumn[string]
		Type *exprColumn[string]
	}
	Table *TableSource
}{
	Table: NewTable("categories", "c"),
}
var Departments = struct {
	Column struct {
		Budget   *exprColumn[int64]
		ID       *exprColumn[int64]
		Name     *exprColumn[string]
		ParentID *exprColumn[int64]
	}
	Table *TableSource
}{
	Table: NewTable("departments", "d"),
}
var Levels = struct {
	Column struct {
		ID     *exprColumn[int64]
		Status *exprColumn[string]
		UserID *exprColumn[int64]
	}
	Table *TableSource
}{
	Table: NewTable("level", "l"),
}
var Orders = struct {
	Column struct {
		Amount   *exprColumn[int]
		CreateAt *exprColumn[time.Time]
		Date     *exprColumn[time.Time]
		ID       *exprColumn[int64]
		Name     *exprColumn[string]
		Status   *exprColumn[string]
		UpdateAt *exprColumn[time.Time]
		UserID   *exprColumn[int64]
	}
	Table *TableSource
}{
	Table: NewTable("orders", "o"),
}
var Products = struct {
	Column struct {
		Count *exprColumn[int]
		ID    *exprColumn[int64]
		Name  *exprColumn[string]
		Type  *exprColumn[string]
	}
	Table *TableSource
}{
	Table: NewTable("products", "p"),
}
var Stats = struct {
	Column struct {
		OrderCount *exprColumn[int]
		TotalSpent *exprColumn[int]
		UserID     *exprColumn[int64]
	}
	Table *TableSource
}{
	Table: NewTable("stats", "s"),
}
var Test = struct {
	Column struct {
		CreateAt *exprColumn[time.Time]
		Date     *exprColumn[time.Time]
		ID       *exprColumn[int64]
		Json     *exprColumn[string]
		Name     *exprColumn[string]
		Number   *exprColumn[int]
		String   *exprColumn[string]
		UpdateAt *exprColumn[time.Time]
		X        *exprColumn[int]
		Y        *exprColumn[int]
	}
	Table *TableSource
}{
	Table: NewTable("test", "t"),
}
var Test1 = struct {
	Column struct {
		Date   *exprColumn[time.Time]
		ID     *exprColumn[int64]
		Json   *exprColumn[string]
		Number *exprColumn[int]
		String *exprColumn[string]
		Time   *exprColumn[time.Time]
	}
	Table *TableSource
}{
	Table: NewTable("test1", "t1"),
}
var Test2 = struct {
	Column struct {
		Date   *exprColumn[time.Time]
		ID     *exprColumn[int64]
		Json   *exprColumn[string]
		Number *exprColumn[int]
		String *exprColumn[string]
		Time   *exprColumn[time.Time]
	}
	Table *TableSource
}{
	Table: NewTable("test2", "t2"),
}
var Users = struct {
	Column struct {
		Age          *exprColumn[int]
		CreateAt     *exprColumn[time.Time]
		Data         *exprColumn[string]
		DepartmentID *exprColumn[int64]
		Email        *exprColumn[string]
		HireDate     *exprColumn[time.Time]
		ID           *exprColumn[int64]
		Name         *exprColumn[string]
		Salary       *exprColumn[int64]
		Status       *exprColumn[string]
		UpdateAt     *exprColumn[time.Time]
	}
	Table *TableSource
}{
	Table: NewTable("users", "u"),
}

// Приватные функции
func init() {
	// Categories
	Categories.Column.ID = Column[int64](Categories.Table.aliasName, "id")
	Categories.Column.Name = Column[string](Categories.Table.aliasName, "name")
	Categories.Column.Type = Column[string](Categories.Table.aliasName, "type")
	// Departments
	Departments.Column.Budget = Column[int64](Departments.Table.aliasName, "budget")
	Departments.Column.ID = Column[int64](Departments.Table.aliasName, "id")
	Departments.Column.Name = Column[string](Departments.Table.aliasName, "name")
	Departments.Column.ParentID = Column[int64](Departments.Table.aliasName, "parent_id")
	// Levels
	Levels.Column.ID = Column[int64](Levels.Table.aliasName, "id")
	Levels.Column.Status = Column[string](Levels.Table.aliasName, "status")
	Levels.Column.UserID = Column[int64](Levels.Table.aliasName, "user_id")
	// Orders
	Orders.Column.Amount = Column[int](Orders.Table.aliasName, "amount")
	Orders.Column.CreateAt = Column[time.Time](Orders.Table.aliasName, "createat")
	Orders.Column.Date = Column[time.Time](Orders.Table.aliasName, "date")
	Orders.Column.ID = Column[int64](Orders.Table.aliasName, "id")
	Orders.Column.Name = Column[string](Orders.Table.aliasName, "name")
	Orders.Column.Status = Column[string](Orders.Table.aliasName, "status")
	Orders.Column.UpdateAt = Column[time.Time](Orders.Table.aliasName, "updateat")
	Orders.Column.UserID = Column[int64](Orders.Table.aliasName, "user_id")
	// Products
	Products.Column.Count = Column[int](Products.Table.aliasName, "count")
	Products.Column.ID = Column[int64](Products.Table.aliasName, "id")
	Products.Column.Name = Column[string](Products.Table.aliasName, "name")
	Products.Column.Type = Column[string](Products.Table.aliasName, "type")
	// Stats
	Stats.Column.OrderCount = Column[int](Stats.Table.aliasName, "order_count")
	Stats.Column.TotalSpent = Column[int](Stats.Table.aliasName, "total_spent")
	Stats.Column.UserID = Column[int64](Stats.Table.aliasName, "user_id")
	// Test
	Test.Column.CreateAt = Column[time.Time](Test.Table.aliasName, "createat")
	Test.Column.Date = Column[time.Time](Test.Table.aliasName, "date")
	Test.Column.ID = Column[int64](Test.Table.aliasName, "id")
	Test.Column.Json = Column[string](Test.Table.aliasName, "json")
	Test.Column.Name = Column[string](Test.Table.aliasName, "name")
	Test.Column.Number = Column[int](Test.Table.aliasName, "number")
	Test.Column.String = Column[string](Test.Table.aliasName, "string")
	Test.Column.UpdateAt = Column[time.Time](Test.Table.aliasName, "updateat")
	Test.Column.X = Column[int](Test.Table.aliasName, "x")
	Test.Column.Y = Column[int](Test.Table.aliasName, "y")
	// Test1
	Test1.Column.Date = Column[time.Time](Test1.Table.aliasName, "date")
	Test1.Column.ID = Column[int64](Test1.Table.aliasName, "id")
	Test1.Column.Json = Column[string](Test1.Table.aliasName, "json")
	Test1.Column.Number = Column[int](Test1.Table.aliasName, "number")
	Test1.Column.String = Column[string](Test1.Table.aliasName, "string")
	Test1.Column.Time = Column[time.Time](Test1.Table.aliasName, "time")
	// Test2
	Test2.Column.Date = Column[time.Time](Test2.Table.aliasName, "date")
	Test2.Column.ID = Column[int64](Test2.Table.aliasName, "id")
	Test2.Column.Json = Column[string](Test2.Table.aliasName, "json")
	Test2.Column.Number = Column[int](Test2.Table.aliasName, "number")
	Test2.Column.String = Column[string](Test2.Table.aliasName, "string")
	Test2.Column.Time = Column[time.Time](Test2.Table.aliasName, "time")
	// Users
	Users.Column.Age = Column[int](Users.Table.aliasName, "age")
	Users.Column.CreateAt = Column[time.Time](Users.Table.aliasName, "createat")
	Users.Column.Data = Column[string](Users.Table.aliasName, "data")
	Users.Column.DepartmentID = Column[int64](Users.Table.aliasName, "department_id")
	Users.Column.Email = Column[string](Users.Table.aliasName, "email")
	Users.Column.HireDate = Column[time.Time](Users.Table.aliasName, "hire_date")
	Users.Column.ID = Column[int64](Users.Table.aliasName, "id")
	Users.Column.Name = Column[string](Users.Table.aliasName, "name")
	Users.Column.Salary = Column[int64](Users.Table.aliasName, "salary")
	Users.Column.Status = Column[string](Users.Table.aliasName, "status")
	Users.Column.UpdateAt = Column[time.Time](Users.Table.aliasName, "updateat")
}
