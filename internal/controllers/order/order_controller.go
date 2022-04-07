package order_controller

import (
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/order"
	order_service "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/services/order"
)

type OrderController struct {
	service order_service.OrderService
}

func NewOrderController(repository order.Repository) *OrderController {
	return &OrderController{service: *order_service.NewOrderService(repository)}
}
