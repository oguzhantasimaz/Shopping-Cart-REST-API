package cart_service

import (
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/cart"
)

type CreateCartRequest struct {
	CustomerID   int                 `json:"customer_id"`
	CartProducts []*cart.CartProduct `json:"products"`
}

type UpdateCartRequest struct {
	ID           int                 `json:"id"`
	CustomerID   int                 `json:"customer_id"`
	CartProducts []*cart.CartProduct `json:"products"`
}

type DeleteCartRequest struct {
	ID int `json:"id"`
}

type FindByIDRequest struct {
	ID int `json:"id"`
}
