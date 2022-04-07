package category_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"

type CategoryService struct {
	repository category.Repository
}

func NewCategoryService(repository category.Repository) *CategoryService {
	return &CategoryService{repository: repository}
}

func (s *CategoryService) FindAll() ([]*category.Category, error) {
	return category.FindAll(s.repository)
}

func (s *CategoryService) FindByID(id int) (*category.Category, error) {
	return category.FindByID(s.repository, id)
}

func (s *CategoryService) Create(newCategory *category.Category) error {
	return category.Create(s.repository, newCategory)
}

func (s *CategoryService) Update(updatedCategory *category.Category) error {
	return category.Update(s.repository, updatedCategory)
}

func (s *CategoryService) Delete(id int) error {
	return category.Delete(s.repository, id)
}

func (s *CategoryService) CreateBulked(newCategories []*category.Category) error {
	return category.CreateBulked(s.repository, newCategories)
}
