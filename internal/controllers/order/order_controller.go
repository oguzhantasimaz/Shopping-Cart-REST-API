package order_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/order"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
	order_service "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/services/order"
	"strconv"
)

type OrderController struct {
	service order_service.OrderService
}

func NewOrderController(repository order.Repository, proRepo product.Repository) *OrderController {
	return &OrderController{service: *order_service.NewOrderService(repository, proRepo)}
}

func (c *OrderController) CreateOrder(gc *gin.Context) {
	request := new(order_service.CreateOrderRequest)
	if err := gc.Bind(request); err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := c.service.Create(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(201, "Order created successfully")
}

func (c *OrderController) UpdateOrder(gc *gin.Context) {
	request := new(order_service.UpdateOrderRequest)
	if err := gc.Bind(request); err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := c.service.Update(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(200, "Order updated successfully")
}

func (c *OrderController) DeleteOrder(gc *gin.Context) {
	request := new(order_service.DeleteOrderRequest)
	id := gc.Param("id")
	var err error
	request.ID, err = strconv.Atoi(id)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = c.service.Delete(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(200, "Order deleted successfully")
}

func (c *OrderController) GetOrder(gc *gin.Context) {
	request := new(order_service.FindByIDRequest)
	id := gc.Param("id")
	var err error
	request.ID, err = strconv.Atoi(id)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	order, err := c.service.FindByID(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(200, order)
}

func (c *OrderController) GetOrdersByCustomerID(gc *gin.Context) {
	request := new(order_service.FindAllByCustomerIDRequest)
	id := gc.Param("customerID")
	var err error
	request.CustomerID, err = strconv.Atoi(id)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	orders, err := c.service.FindAllByCustomerID(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(200, orders)
}
