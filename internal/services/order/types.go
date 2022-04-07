package order_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"

type CreateOrderRequest struct {
	CustomerID int                `json:"customer_id"`
	Products   []*product.Product `json:"products"`
}

type FindByIDRequest struct {
	ID int `json:"id"`
}

type UpdateOrderRequest struct {
	ID         int                `json:"id"`
	CustomerID int                `json:"customer_id"`
	Products   []*product.Product `json:"products"`
}

type DeleteOrderRequest struct {
	ID int `json:"id"`
}

type FindAllByCustomerIDRequest struct {
	CustomerID int `json:"customer_id"`
}
