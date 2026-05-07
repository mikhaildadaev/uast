package uast

import (
	"time"
)

// Публичные переменные
var Categories = struct {
	Table *TableSource
	ID    *exprColumn[int64]
	Name  *exprColumn[string]
	Type  *exprColumn[string]
}{
	Table: NewTable("categories", "c"),
}
var Departments = struct {
	Table    *TableSource
	Budget   *exprColumn[int64]
	ID       *exprColumn[int64]
	Name     *exprColumn[string]
	ParentID *exprColumn[int64]
}{
	Table: NewTable("departments", "d"),
}
var Levels = struct {
	Table  *TableSource
	ID     *exprColumn[int64]
	Status *exprColumn[string]
	UserID *exprColumn[int64]
}{
	Table: NewTable("level", "l"),
}
var Orders = struct {
	Table    *TableSource
	Amount   *exprColumn[int]
	CreateAt *exprColumn[time.Time]
	Date     *exprColumn[time.Time]
	ID       *exprColumn[int64]
	Name     *exprColumn[string]
	Status   *exprColumn[string]
	UpdateAt *exprColumn[time.Time]
	UserID   *exprColumn[int64]
}{
	Table: NewTable("orders", "o"),
}
var Products = struct {
	Table *TableSource
	Count *exprColumn[int]
	ID    *exprColumn[int64]
	Name  *exprColumn[string]
	Type  *exprColumn[string]
}{
	Table: NewTable("products", "p"),
}
var Stats = struct {
	Table      *TableSource
	OrderCount *exprColumn[int]
	TotalSpent *exprColumn[int]
	UserID     *exprColumn[int64]
}{
	Table: NewTable("stats", "s"),
}
var Test = struct {
	Table    *TableSource
	CreateAt *exprColumn[time.Time]
	Number   *exprColumn[int]
	String   *exprColumn[string]
	UpdateAt *exprColumn[time.Time]
	X        *exprColumn[int]
	Y        *exprColumn[int]
}{
	Table: NewTable("test", "t"),
}
var Users = struct {
	Table        *TableSource
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
}{
	Table: NewTable("users", "u"),
}

// Приватные функции
func init() {
	// Categories
	Categories.ID = Column[int64](Categories.Table.aliasName, "id")
	Categories.Name = Column[string](Categories.Table.aliasName, "name")
	Categories.Type = Column[string](Categories.Table.aliasName, "type")
	// Departments
	Departments.Budget = Column[int64](Departments.Table.aliasName, "budget")
	Departments.ID = Column[int64](Departments.Table.aliasName, "id")
	Departments.Name = Column[string](Departments.Table.aliasName, "name")
	Departments.ParentID = Column[int64](Departments.Table.aliasName, "parent_id")
	// Levels
	Levels.ID = Column[int64](Levels.Table.aliasName, "id")
	Levels.Status = Column[string](Levels.Table.aliasName, "status")
	Levels.UserID = Column[int64](Levels.Table.aliasName, "user_id")
	// Orders
	Orders.Amount = Column[int](Orders.Table.aliasName, "amount")
	Orders.CreateAt = Column[time.Time](Orders.Table.aliasName, "createat")
	Orders.Date = Column[time.Time](Orders.Table.aliasName, "date")
	Orders.ID = Column[int64](Orders.Table.aliasName, "id")
	Orders.Name = Column[string](Orders.Table.aliasName, "name")
	Orders.Status = Column[string](Orders.Table.aliasName, "status")
	Orders.UpdateAt = Column[time.Time](Orders.Table.aliasName, "updateat")
	Orders.UserID = Column[int64](Orders.Table.aliasName, "user_id")
	// Products
	Products.Count = Column[int](Products.Table.aliasName, "count")
	Products.ID = Column[int64](Products.Table.aliasName, "id")
	Products.Name = Column[string](Products.Table.aliasName, "name")
	Products.Type = Column[string](Products.Table.aliasName, "type")
	// Stats
	Stats.OrderCount = Column[int](Stats.Table.aliasName, "order_count")
	Stats.TotalSpent = Column[int](Stats.Table.aliasName, "total_spent")
	Stats.UserID = Column[int64](Stats.Table.aliasName, "user_id")
	// Test
	Test.CreateAt = Column[time.Time](Test.Table.aliasName, "createat")
	Test.Number = Column[int](Test.Table.aliasName, "number")
	Test.String = Column[string](Test.Table.aliasName, "string")
	Test.UpdateAt = Column[time.Time](Test.Table.aliasName, "updateat")
	Test.X = Column[int](Test.Table.aliasName, "x")
	Test.Y = Column[int](Test.Table.aliasName, "y")
	// Users
	Users.Age = Column[int](Users.Table.aliasName, "age")
	Users.CreateAt = Column[time.Time](Users.Table.aliasName, "createat")
	Users.Data = Column[string](Users.Table.aliasName, "data")
	Users.DepartmentID = Column[int64](Users.Table.aliasName, "department_id")
	Users.Email = Column[string](Users.Table.aliasName, "email")
	Users.HireDate = Column[time.Time](Users.Table.aliasName, "hire_date")
	Users.ID = Column[int64](Users.Table.aliasName, "id")
	Users.Name = Column[string](Users.Table.aliasName, "name")
	Users.Salary = Column[int64](Users.Table.aliasName, "salary")
	Users.Status = Column[string](Users.Table.aliasName, "status")
	Users.UpdateAt = Column[time.Time](Users.Table.aliasName, "updateat")
}
