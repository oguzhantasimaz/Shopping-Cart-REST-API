package category_controller

import (
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
