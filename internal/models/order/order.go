package order

import (
	"sync"
	"time"

	"gorm.io/gorm"
)

var mutex = &sync.Mutex{}

type Order struct {
	gorm.Model
	ID            int
	CustomerID    int
	OrderProducts []*OrderProduct
	TotalPrice    float64
	Active        bool
	CreatedAt     time.Time `gorm:"<-:create"`
}

type OrderProduct struct {
	gorm.Model
	OrderID   int
	ProductID int
	Quantity  int
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
	mutex.Lock()
	defer mutex.Unlock()
	return r.Update(order)
}

func Delete(r Repository, id int) error {
	mutex.Lock()
	defer mutex.Unlock()
	return r.Delete(id)
}
