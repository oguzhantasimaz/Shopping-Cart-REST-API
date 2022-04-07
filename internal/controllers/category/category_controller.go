package category_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"
	category_service "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/services/category"
)

type CategoryController struct {
	service category_service.CategoryService
}

func NewCategoryController(repository category.Repository) *CategoryController {
	return &CategoryController{
		service: *category_service.NewCategoryService(repository),
	}
}

func (c *CategoryController) Create(gc *gin.Context) {
	request := new(category_service.CreateCategoryRequest)
	if err := gc.Bind(request); err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := c.service.Create(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
