package product_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"

type CreateProductRequest struct {
	Name      string            `json:"name"`
	SKU       string            `json:"description"`
	UnitPrice float64           `json:"unit_price"`
	Quantity  int               `json:"quantity"`
	Category  category.Category `json:"category"`
}

type UpdateProductRequest struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	SKU       string            `json:"description"`
	UnitPrice float64           `json:"unit_price"`
	Quantity  int               `json:"quantity"`
	Category  category.Category `json:"category"`
}

type DeleteProductRequest struct {
	Id int `json:"id"`
}

type FindProductByIdRequest struct {
	Id int `json:"id"`
}
