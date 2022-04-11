package product

import (
	"sync"
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

var mutex = &sync.Mutex{}

func Create(r Repository, p *Product) error {
	return r.Create(p)
}

func Update(r Repository, p *Product) error {
	mutex.Lock()
	defer mutex.Unlock()
	return r.Update(p)
}

func Delete(r Repository, id int) error {
	mutex.Lock()
	defer mutex.Unlock()
	return r.Delete(id)
}

func FindByID(r Repository, id int) (*Product, error) {
	return r.FindByID(id)
}

func FindAll(r Repository) ([]*Product, error) {
	return r.FindAll()
}
