package cart_controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/cart"
	cart_service "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/services/cart"
	"strconv"
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
	gc.JSON(201, gin.H{
		"message": "Cart created successfully",
	})
}

func (c *CartController) UpdateCart(gc *gin.Context) {
	request := new(cart_service.UpdateCartRequest)
	err := json.NewDecoder(gc.Request.Body).Decode(&request)
	if err != nil {
		gc.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = c.service.Update(request)
	if err != nil {
		gc.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	gc.JSON(200, gin.H{
		"message": "Cart updated successfully",
	})
}

func (c *CartController) DeleteCart(gc *gin.Context) {
	request := new(cart_service.DeleteCartRequest)
	id := gc.Param("id")
	var err error
	if err != nil {
		gc.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	request.ID, err = strconv.Atoi(id)
	if err != nil {
		gc.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = c.service.Delete(request)
}

func (c *CartController) FindByIdCart(gc *gin.Context) {
	request := new(cart_service.FindByIDRequest)
	id := gc.Param("id")
	var err error
	if err != nil {
		gc.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	request.ID, err = strconv.Atoi(id)
	if err != nil {
		gc.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	cart, err := c.service.FindByID(request)
	if err != nil {
		gc.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	gc.JSON(200, gin.H{
		"cart": cart,
	})
}
