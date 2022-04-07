package product_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"

type ProductService struct {
	repository product.Repository
}

func NewProductService(repository product.Repository) *ProductService {
	return &ProductService{repository: repository}
}

func (s *ProductService) Create(req *CreateProductRequest) error {
	err := CreateProductValidate(req)
	if err != nil {
		return err
	}
	newProduct := &product.Product{
		Name:      req.Name,
		SKU:       req.SKU,
		UnitPrice: req.UnitPrice,
		Quantity:  req.Quantity,
		Category:  req.Category,
	}
	return product.Create(s.repository, newProduct)
}

func (s *ProductService) Update(req *UpdateProductRequest) error {
	err := UpdateProductValidate(req)
	if err != nil {
		return err
	}
	productToUpdate := &product.Product{
		ID:        req.ID,
		Name:      req.Name,
		SKU:       req.SKU,
		UnitPrice: req.UnitPrice,
		Quantity:  req.Quantity,
		Category:  req.Category,
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
