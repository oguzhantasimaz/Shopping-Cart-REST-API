package category_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"
	category_service "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/services/category"
	"strconv"
)

type CategoryController struct {
	service category_service.CategoryService
}

func NewCategoryController(repository category.Repository) *CategoryController {
	return &CategoryController{
		service: *category_service.NewCategoryService(repository),
	}
}

func (c *CategoryController) CreateCategory(gc *gin.Context) {
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
	gc.JSON(201, gin.H{"message": "Category created successfully"})
}

func (c *CategoryController) CreateBulkCategory(gc *gin.Context) {
	request := new(category_service.CreateCategoryBulkedRequest)
	if err := gc.Bind(request); err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := c.service.CreateBulked(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(201, gin.H{"message": "Categories created successfully"})
}

func (c *CategoryController) FindAllCategories(gc *gin.Context) {
	categories, err := c.service.FindAll()
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(200, categories)
}

func (c *CategoryController) FindCategory(gc *gin.Context) {
	request := new(category_service.FindByIDCategoryRequest)
	id := gc.Param("id")
	var err error
	request.ID, err = strconv.Atoi(id)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	category, err := c.service.FindByID(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(200, category)
}

func (c *CategoryController) UpdateCategory(gc *gin.Context) {
	request := new(category_service.UpdateCategoryRequest)
	if err := gc.Bind(request); err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := c.service.Update(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(200, gin.H{"message": "Category updated successfully"})
}

func (c *CategoryController) DeleteCategory(gc *gin.Context) {
	request := new(category_service.DeleteCategoryRequest)
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
	gc.JSON(200, gin.H{"message": "Category deleted successfully"})
}
