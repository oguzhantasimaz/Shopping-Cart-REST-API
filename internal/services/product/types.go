package product_service

type CreateProductRequest struct {
	Name       string  `json:"name"`
	SKU        string  `json:"SKU"`
	UnitPrice  float64 `json:"unit_price"`
	Stock      int     `json:"stock"`
	CategoryID int     `json:"category_id"`
}

type UpdateProductRequest struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	SKU        string  `json:"SKU"`
	UnitPrice  float64 `json:"unit_price"`
	Stock      int     `json:"stock"`
	CategoryID int     `json:"category_id"`
}

type DeleteProductRequest struct {
	ID int `json:"id"`
}

type FindProductByIDRequest struct {
	ID int `json:"id"`
}
