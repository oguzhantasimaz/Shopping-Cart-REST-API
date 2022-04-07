package cart_controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/cart"
	cart_service "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/services/cart"
)

type CartController struct {
	service cart_service.CartService
}

func NewCartController(repository cart.Repository) *CartController {
	return &CartController{
		service: *cart_service.NewCartService(repository),
	}
}

func (c *CartController) CreateCart(gc *gin.Context) {
	request := new(cart_service.CreateCartRequest)
	err := json.NewDecoder(gc.Request.Body).Decode(&request)
	if err != nil {
		gc.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = c.service.Create(request)
	if err != nil {
		gc.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	gc.JSON(200, gin.H{
		"message": "Cart created successfully",
	})
}
