package main

import (
	"fmt"
	cart_controller "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/controllers/cart"
	product_controller "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/controllers/product"
	cart_repository "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/repositories/cart"
	category_repository "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/repositories/category"
	order_repository "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/repositories/order"
	product_repository "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/repositories/product"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/pkg/database_handler"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/pkg/graceful"
	middleware "github.com/oguzhantasimaz/Shopping-Cart-REST-API/pkg/middlewares"
)

func main() {
	r := gin.Default()
	registerMiddlewares(r)

	db := database_handler.NewMySQLDB("root:Ot123456@tcp(127.0.0.1:3306)/shopping?charset=utf8mb4&parseTime=True&loc=Local")

	cartRepo := cart_repository.NewCartRepository(db)
	categoryRepo := category_repository.NewCategoryRepository(db)
	orderRepo := order_repository.NewOrderRepository(db)
	productRepo := product_repository.NewProductRepository(db)
	//userRepo := user_repository.NewUserRepository(db)

	if cartRepo.Migration() != nil {
		log.Fatal("Cart Migration failed")
	}
	if categoryRepo.Migration() != nil {
		log.Fatal("Category Migration failed")
	}
	if orderRepo.Migration() != nil {
		log.Fatal("Order Migration failed")
	}
	if productRepo.Migration() != nil {
		log.Fatal("Product Migration failed")
	}
	//if userRepo.Migration() != nil {
	//	log.Fatal("User Migration failed")
	//}

	cartCtrl := cart_controller.NewCartController(cartRepo)
	//categoryCtrl := category_controller.NewCategoryController(categoryRepo)
	//orderCtrl := order_controller.NewOrderController(orderRepo)
	productCtrl := product_controller.NewProductController(productRepo)
	//userCtrl := user_controller.NewUserController(userRepo)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/cart/create", cartCtrl.CreateCart)
	r.POST("/product/create", productCtrl.CreateProduct)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s \n", err)
		}
	}()

	graceful.ShutdownGin(srv, time.Second*5)
}

func registerMiddlewares(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
}
