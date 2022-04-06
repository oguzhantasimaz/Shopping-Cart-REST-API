package category_repository

import (
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Migration() error {
	err := r.db.Migrator().DropTable(&category.Category{})
	if err != nil {
		return err
	}
	err = r.db.AutoMigrate(&category.Category{})
	if err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) Create(category *category.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) Update(category *category.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id int) error {
	return r.db.Where("id = ?", id).Delete(&category.Category{}).Error
}

func (r *categoryRepository) FindAll() ([]category.Category, error) {
	var categories []category.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) FindByID(id int) (*category.Category, error) {
	c := new(category.Category)
	err := r.db.Where("id = ?", id).First(c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *categoryRepository) CreateBulked(categories []category.Category) error {
	return r.db.Create(&categories).Error
}
