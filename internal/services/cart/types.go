package cart_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"

type CreateCartRequest struct {
	CustomerID int                `json:"user_id"`
	Products   []*product.Product `json:"products"`
}

type UpdateCartRequest struct {
	ID         int                `json:"id"`
	CustomerID int                `json:"user_id"`
	Products   []*product.Product `json:"products"`
}

type DeleteCartRequest struct {
	ID int `json:"id"`
}

type FindByIDRequest struct {
	ID int `json:"id"`
}
