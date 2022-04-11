package main

import (
	"fmt"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/config"
	auth_controller "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/controllers/auth"
	cart_controller "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/controllers/cart"
	category_controller "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/controllers/category"
	order_controller "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/controllers/order"
	product_controller "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/controllers/product"
	cart_repository "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/repositories/cart"
	category_repository "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/repositories/category"
	order_repository "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/repositories/order"
	product_repository "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/repositories/product"
	user_repository "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/repositories/user"
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
	userRepo := user_repository.NewUserRepository(db)

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
	if userRepo.Migration() != nil {
		log.Fatal("User Migration failed")
	}

	userRepo.InsertSampleData()
	appConfig, err := config.GetAllConfigValues("/Users/oguzhantasimaz/Desktop/Shopping-Cart-REST-API/config/location.qa.yaml")
	if err != nil {
		log.Fatal(err)
	}

	cartCtrl := cart_controller.NewCartController(cartRepo)
	categoryCtrl := category_controller.NewCategoryController(categoryRepo)
	orderCtrl := order_controller.NewOrderController(orderRepo, productRepo)
	productCtrl := product_controller.NewProductController(productRepo, categoryRepo)
	authCtrl := auth_controller.NewAuthController(appConfig, userRepo)
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

	//auth routes
	r.POST("/login", authCtrl.Login)
	r.POST("/signup", authCtrl.Signup)
	r.POST("/logout", authCtrl.Logout)

	//cart routes
	r.GET("/cart/get/:id", middleware.AuthMiddleware(appConfig.SecretKey), cartCtrl.GetCart)
	r.POST("/cart/create", middleware.AuthMiddleware(appConfig.SecretKey), cartCtrl.CreateCart)
	r.PUT("/cart/update/:id", middleware.AuthMiddleware(appConfig.SecretKey), cartCtrl.UpdateCart)
	r.DELETE("/cart/delete/:id", middleware.AuthMiddleware(appConfig.SecretKey), cartCtrl.DeleteCart)

	//category routes
	r.GET("/category/get/:id", categoryCtrl.GetCategory)
	r.GET("/category/get", categoryCtrl.GetAllCategories)
	r.POST("/category/create", middleware.AuthMiddleware(appConfig.SecretKey), categoryCtrl.CreateCategory)
	r.POST("/category/create/bulk", middleware.AuthMiddleware(appConfig.SecretKey), categoryCtrl.CreateCategories)
	r.PUT("/category/update/:id", middleware.AuthMiddleware(appConfig.SecretKey), categoryCtrl.UpdateCategory)
	r.DELETE("/category/delete/:id", middleware.AuthMiddleware(appConfig.SecretKey), categoryCtrl.DeleteCategory)

	//product routes
	r.GET("/product/get/:id", productCtrl.GetProduct)
	r.GET("/product/get", productCtrl.GetProducts)
	r.POST("/product/create", middleware.AuthMiddleware(appConfig.SecretKey), productCtrl.CreateProduct)
	r.PUT("/product/update/:id", middleware.AuthMiddleware(appConfig.SecretKey), productCtrl.UpdateProduct)
	r.DELETE("/product/delete/:id", middleware.AuthMiddleware(appConfig.SecretKey), productCtrl.DeleteProduct)

	//order routes
	r.GET("/order/get/:id", middleware.AuthMiddleware(appConfig.SecretKey), orderCtrl.GetOrder)
	r.GET("/order/get/customer/:customerID", middleware.AuthMiddleware(appConfig.SecretKey), orderCtrl.GetOrdersByCustomerID)
	r.POST("/order/create", middleware.AuthMiddleware(appConfig.SecretKey), orderCtrl.CreateOrder)
	r.PUT("/order/update/:id", middleware.AuthMiddleware(appConfig.SecretKey), orderCtrl.UpdateOrder)
	r.DELETE("/order/delete/:id", middleware.AuthMiddleware(appConfig.SecretKey), orderCtrl.DeleteOrder)

	//user routes

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
