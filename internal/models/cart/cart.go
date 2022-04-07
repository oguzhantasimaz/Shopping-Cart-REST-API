package cart

import (
	"time"

	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Id         int
	CustomerId int
	Products   []*product.Product
	TotalPrice float64
	CreatedAt  time.Time `gorm:"<-:create"`
}

func CreateCart(r Repository, cart *Cart) error {
	return r.Create(cart)
}

func UpdateCart(r Repository, cart *Cart) error {
	return r.Update(cart)
}

func DeleteCart(r Repository, id int) error {
	return r.Delete(id)
}

func FindByIdCart(r Repository, id int) (*Cart, error) {
	cart, err := r.FindById(id)
	return cart, err
}
