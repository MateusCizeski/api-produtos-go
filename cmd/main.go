package main

import (
	"api-products/controller"
	"api-products/db"
	"api-products/repository"
	"api-products/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProcuts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:id", ProductController.GetProductById)

	server.Run(":8000")
}
