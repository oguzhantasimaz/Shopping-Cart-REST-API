package category_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"

type CategoryService struct {
	repository category.Repository
}

func NewCategoryService(repository category.Repository) *CategoryService {
	return &CategoryService{repository: repository}
}

func (s *CategoryService) FindAll() (*[]category.Category, error) {
	return category.FindAll(s.repository)
}

func (s *CategoryService) FindByID(req *FindByIDCategoryRequest) (*category.Category, error) {
	if err := FindByIDCategoryValidate(req); err != nil {
		return nil, err
	}
	return category.FindByID(s.repository, req.ID)
}

func (s *CategoryService) Create(req *CreateCategoryRequest) error {
	if err := CreateCategoryValidate(req); err != nil {
		return err
	}
	newCategory := &category.Category{
		Name: req.Name,
	}
	return category.Create(s.repository, newCategory)
}

func (s *CategoryService) Update(req *UpdateCategoryRequest) error {
	if err := UpdateCategoryValidate(req); err != nil {
		return err
	}
	categoryToUpdate := &category.Category{
		ID:   req.ID,
		Name: req.Name,
	}
	return category.Update(s.repository, categoryToUpdate)
}

func (s *CategoryService) Delete(req *DeleteCategoryRequest) error {
	if err := DeleteCategoryValidate(req); err != nil {
		return err
	}
	return category.Delete(s.repository, req.ID)
}

func (s *CategoryService) CreateBulked(req *CreateCategoryBulkedRequest) error {
	if err := CreateCategoryBulkedValidate(req); err != nil {
		return err
	}
	for _, c := range req.Categories {
		newCategory := &category.Category{
			Name: c.Name,
		}
		if err := category.Create(s.repository, newCategory); err != nil {
			return err
		}
	}
	return nil
}
