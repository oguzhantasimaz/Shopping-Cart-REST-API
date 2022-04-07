package order

import (
	"time"

	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID         int
	CustomerID int                `gorm:"foreignkey:CustomerID"`
	Products   []*product.Product `gorm:"many2many:order_products"`
	TotalPrice float64
	Active     bool
	CreatedAt  time.Time `gorm:"<-:create"`
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
