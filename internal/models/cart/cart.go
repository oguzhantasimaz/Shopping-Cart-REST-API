package cart

import (
	"time"

	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID         int
	CustomerID int                `gorm:"foreignkey:CustomerID"`
	Products   []*product.Product `gorm:"many2many:cart_products"`
	TotalPrice float64
	CreatedAt  time.Time `gorm:"<-:create"`
}

func Create(r Repository, cart *Cart) error {
	return r.Create(cart)
}

func Update(r Repository, cart *Cart) error {
	return r.Update(cart)
}

func Delete(r Repository, id int) error {
	return r.Delete(id)
}

func FindByID(r Repository, id int) (*Cart, error) {
	cart, err := r.FindByID(id)
	return cart, err
}
