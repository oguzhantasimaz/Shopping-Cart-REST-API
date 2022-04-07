package category

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID        int
	Name      string
	CreatedAt time.Time `gorm:"<-:create"`
}

func FindAll(r Repository) (*[]Category, error) {
	categories, err := r.FindAll()
	return categories, err
}

func FindByID(r Repository, id int) (*Category, error) {
	category, err := r.FindByID(id)
	return category, err
}

func Create(r Repository, category *Category) error {
	err := r.Create(category)
	return err
}

func Update(r Repository, category *Category) error {
	err := r.Update(category)
	return err
}

func Delete(r Repository, id int) error {
	err := r.Delete(id)
	return err
}

func CreateBulked(r Repository, categories *[]Category) error {
	err := r.CreateBulked(categories)
	return err
}
