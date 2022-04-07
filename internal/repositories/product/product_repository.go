package product_repository

import (
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) Migration() error {
	err := r.db.Migrator().DropTable(&product.Product{})
	if err != nil {
		return err
	}
	err = r.db.AutoMigrate(&product.Product{})
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) Create(p *product.Product) error {
	return r.db.Create(p).Error
}

func (r *productRepository) Update(p *product.Product) error {
	return r.db.Save(p).Error
}

func (r *productRepository) Delete(id int) error {
	return r.db.Where("id = ?", id).Delete(&product.Product{}).Error
}

func (r *productRepository) FindById(id int) (*product.Product, error) {
	p := new(product.Product)
	err := r.db.Where("id = ?", id).First(p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}
