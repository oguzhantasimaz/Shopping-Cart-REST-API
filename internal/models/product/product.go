package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID         int
	Name       string
	SKU        string
	UnitPrice  float64
	Stock      int
	CategoryID int
	CreatedAt  time.Time `gorm:"<-:create"`
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

func FindByID(r Repository, id int) (*Product, error) {
	return r.FindByID(id)
}

func FindAll(r Repository) ([]*Product, error) {
	return r.FindAll()
}
