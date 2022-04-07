package order_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"

type CreateOrderRequest struct {
	CustomerId int                `json:"customer_id"`
	Products   []*product.Product `json:"products"`
}

type FindByIdRequest struct {
	Id int `json:"id"`
}

type UpdateOrderRequest struct {
	Id         int                `json:"id"`
	CustomerId int                `json:"customer_id"`
	Products   []*product.Product `json:"products"`
}

type DeleteOrderRequest struct {
	Id int `json:"id"`
}

type FindAllByCustomerIdRequest struct {
	CustomerId int `json:"customer_id"`
}
