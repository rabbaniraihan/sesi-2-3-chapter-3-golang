package main

import (
	"sesi-2/controller"
	"sesi-2/middleware"
	"sesi-2/model"
	"sesi-2/repository"
	"sesi-2/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=rabbani11 dbname=postgres sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.User{}, model.Product{})
}

func main() {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(*productRepository)
	productController := controller.NewProductController(*productService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository)
	userController := controller.NewUserController(*userService)

	g := gin.Default()

	g.POST("/user/register", userController.Register)
	g.POST("/user/login", userController.Login)

	productGroup := g.Group("/product", middleware.AuthMiddleware)
	productGroup.POST("/", productController.CreateProduct)
	productGroup.GET("/", productController.GetAllProduct)
	productGroup.PUT("/:id", productController.UpdateProduct)
	productGroup.DELETE("/:id", productController.DeleteProduct)
	productGroup.GET("/:id", productController.GetProductById)

	err := g.Run("localhost:8084")
	if err != nil {
		panic(err)
	}
}
