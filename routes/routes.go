package routes

import (
	"go-restfulapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()
	productRoutes := router.Group("/api/products")
	{
		productRoutes.GET("/", controllers.GetProducts)
		productRoutes.GET("/:id", controllers.GetProductById)
		productRoutes.POST("/", controllers.CreateProduct)
		productRoutes.PUT("/:id", controllers.UpdateProduct)
		productRoutes.DELETE("/:id", controllers.DeleteProductById)
	}

	customerRoutes := router.Group("/api/customers")
	{
		customerRoutes.GET("/", controllers.GetCustomer)
		customerRoutes.GET("/:id", controllers.GetCustomerById)
		customerRoutes.POST("/", controllers.CreateCustomer)
		customerRoutes.PUT("/:id", controllers.UpdateCustomer)
		customerRoutes.DELETE("/:id", controllers.DeleteCustomerById)
	}

	return router
}