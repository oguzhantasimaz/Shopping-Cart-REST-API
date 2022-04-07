package cart_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"

type CreateCartRequest struct {
	CustomerId int                `json:"user_id"`
	Products   []*product.Product `json:"products"`
}

type UpdateCartRequest struct {
	Id         int                `json:"id"`
	CustomerId int                `json:"user_id"`
	Products   []*product.Product `json:"products"`
}

type DeleteCartRequest struct {
	Id int `json:"id"`
}

type FindByIdRequest struct {
	Id int `json:"id"`
}
