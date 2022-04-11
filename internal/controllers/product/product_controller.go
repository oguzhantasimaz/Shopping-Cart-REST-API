package product_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
	category_service "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/services/category"
	product_service "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/services/product"
	"strconv"
)

type ProductController struct {
	service      product_service.ProductService
	categoryServ category_service.CategoryService
}

func NewProductController(repository product.Repository, catRepository category.Repository) *ProductController {
	return &ProductController{service: *product_service.NewProductService(repository, catRepository)}
}

func (c *ProductController) CreateProduct(gc *gin.Context) {
	request := new(product_service.CreateProductRequest)
	if err := gc.Bind(request); err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := c.service.Create(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(200, gin.H{"message": "Product created successfully"})
}

func (c *ProductController) UpdateProduct(gc *gin.Context) {
	request := new(product_service.UpdateProductRequest)
	if err := gc.Bind(request); err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := c.service.Update(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(200, gin.H{"message": "Product updated successfully"})
}

func (c *ProductController) DeleteProduct(gc *gin.Context) {
	request := new(product_service.DeleteProductRequest)
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
	gc.JSON(200, gin.H{"message": "Product deleted successfully"})
}

func (c *ProductController) GetProduct(gc *gin.Context) {
	request := new(product_service.FindProductByIDRequest)
	id := gc.Param("id")
	var err error
	request.ID, err = strconv.Atoi(id)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	product, err := c.service.FindByID(request)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(200, gin.H{"product": product})
}

func (c *ProductController) GetProducts(gc *gin.Context) {
	queryParams := gc.Request.URL.Query()
	fmt.Println(queryParams)
	products, err := c.service.FindAll()
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(200, gin.H{"products": products})
}
