package order_service

import (
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/order"
)

type CreateOrderRequest struct {
	CustomerID    int                   `json:"customer_id"`
	OrderProducts []*order.OrderProduct `json:"order_products"`
}

type FindByIDRequest struct {
	ID int `json:"id"`
}

type UpdateOrderRequest struct {
	ID            int                   `json:"id"`
	CustomerID    int                   `json:"customer_id"`
	OrderProducts []*order.OrderProduct `json:"order_products"`
}

type DeleteOrderRequest struct {
	ID int `json:"id"`
}

type FindAllByCustomerIDRequest struct {
	CustomerID int `json:"customer_id"`
}
