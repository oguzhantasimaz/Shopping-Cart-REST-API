package product_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"

type CreateProductRequest struct {
	Name      string            `json:"name"`
	SKU       string            `json:"SKU"`
	UnitPrice float64           `json:"unit_price"`
	Quantity  int               `json:"quantity"`
	Category  category.Category `json:"category"`
}

type UpdateProductRequest struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	SKU       string            `json:"description"`
	UnitPrice float64           `json:"unit_price"`
	Quantity  int               `json:"quantity"`
	Category  category.Category `json:"category"`
}

type DeleteProductRequest struct {
	ID int `json:"id"`
}

type FindProductByIDRequest struct {
	ID int `json:"id"`
}
