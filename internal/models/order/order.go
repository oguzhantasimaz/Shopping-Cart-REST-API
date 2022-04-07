package order

import (
	"time"

	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id         int
	CustomerId int
	Products   []*product.Product
	TotalPrice float64
	Active     bool
	CreatedAt  time.Time `gorm:"<-:create"`
}

func Create(r Repository, order *Order) error {
	return r.Create(order)
}

func FindAllByCustomerId(r Repository, customerId int) ([]*Order, error) {
	return r.FindAllByCustomerId(customerId)
}

func FindById(r Repository, id int) (*Order, error) {
	return r.FindById(id)
}

func Update(r Repository, order *Order) error {
	return r.Update(order)
}

func Delete(r Repository, id int) error {
	return r.Delete(id)
}
