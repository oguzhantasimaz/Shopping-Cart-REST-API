package order

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID            int
	CustomerID    int             `gorm:"foreignKey:CustomerID"`
	OrderProducts []*OrderProduct `gorm:"foreignKey:OrderID"`
	TotalPrice    float64
	Active        bool
	CreatedAt     time.Time `gorm:"<-:create"`
}

type OrderProduct struct {
	gorm.Model
	OrderID   int `gorm:"foreignKey:OrderID"`
	ProductID int `gorm:"foreignKey:ProductID"`
	Quantity  int
	CreatedAt time.Time `gorm:"<-:create"`
}

func Create(r Repository, order *Order) error {
	return r.Create(order)
}

func FindAllByCustomerID(r Repository, customerID int) (*[]Order, error) {
	return r.FindAllByCustomerID(customerID)
}

func FindByID(r Repository, id int) (*Order, error) {
	return r.FindByID(id)
}

func Update(r Repository, order *Order) error {
	return r.Update(order)
}

func Delete(r Repository, id int) error {
	return r.Delete(id)
}
