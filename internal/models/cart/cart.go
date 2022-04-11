package cart

import (
	"sync"
	"time"

	"gorm.io/gorm"
)

var mutex = &sync.Mutex{}

type Cart struct {
	gorm.Model
	ID           int
	CustomerID   int `gorm:"foreignkey:CustomerID"`
	CartProducts []*CartProduct
	TotalPrice   float64
	CreatedAt    time.Time `gorm:"<-:create"`
}

type CartProduct struct {
	gorm.Model
	CartID    int
	ProductID int
	Quantity  int
}

func Create(r Repository, cart *Cart) error {
	mutex.Lock()
	defer mutex.Unlock()
	return r.Create(cart)
}

func Update(r Repository, cart *Cart) error {
	mutex.Lock()
	defer mutex.Unlock()
	return r.Update(cart)
}

func Delete(r Repository, id int) error {
	mutex.Lock()
	defer mutex.Unlock()
	return r.Delete(id)
}

func FindByID(r Repository, id int) (*Cart, error) {
	cart, err := r.FindByID(id)
	return cart, err
}
