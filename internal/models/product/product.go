package product

import (
	"time"

	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id        int
	Name      string
	SKU       string
	UnitPrice float64
	Quantity  int
	Category  category.Category
	CreatedAt time.Time `gorm:"<-:create"`
}

func Create(r Repository, p *Product) error {
	return r.Create(p)
}

func Update(r Repository, p *Product) error {
	return r.Update(p)
}

func Delete(r Repository, id int) error {
	return r.Delete(id)
}

func FindById(r Repository, id int) (*Product, error) {
	return r.FindById(id)
}
