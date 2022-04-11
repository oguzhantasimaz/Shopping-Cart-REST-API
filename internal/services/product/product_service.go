package product_service

import (
	"fmt"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
)

type ProductService struct {
	repository   product.Repository
	categoryRepo category.Repository
}

func NewProductService(repository product.Repository, catRepo category.Repository) *ProductService {
	return &ProductService{repository: repository, categoryRepo: catRepo}
}

func (s *ProductService) Create(req *CreateProductRequest) error {
	err := CreateProductValidate(req)
	if err != nil {
		return err
	}

	cat, catError := category.FindByID(s.categoryRepo, req.CategoryID)
	if catError != nil {
		return catError
	}
	if cat == nil {
		return fmt.Errorf("category not found")
	}
	newProduct := &product.Product{
		Name:       req.Name,
		SKU:        req.SKU,
		UnitPrice:  req.UnitPrice,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
	}
	return product.Create(s.repository, newProduct)
}

func (s *ProductService) Update(req *UpdateProductRequest) error {
	err := UpdateProductValidate(req)
	if err != nil {
		return err
	}
	productToUpdate := &product.Product{
		ID:         req.ID,
		Name:       req.Name,
		SKU:        req.SKU,
		UnitPrice:  req.UnitPrice,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
	}
	return product.Update(s.repository, productToUpdate)
}

func (s *ProductService) Delete(req *DeleteProductRequest) error {
	err := DeleteProductValidate(req)
	if err != nil {
		return err
	}
	return product.Delete(s.repository, req.ID)
}

func (s *ProductService) FindByID(req *FindProductByIDRequest) (*product.Product, error) {
	err := FindProductByIDValidate(req)
	if err != nil {
		return nil, err
	}
	return product.FindByID(s.repository, req.ID)
}

func (s *ProductService) FindAll() ([]*product.Product, error) {
	return product.FindAll(s.repository)
}
