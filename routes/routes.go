package routes

import (
	"go-restfulapi/config"
	"go-restfulapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	protected := router.Group("/")
	protected.Use(config.AuthMiddleware())
	{
		productRoutes := protected.Group("/api/products")
		{
			productRoutes.GET("/", controllers.GetProducts)
			productRoutes.GET("/:id", controllers.GetProductById)
			productRoutes.POST("/", controllers.CreateProduct)
			productRoutes.PUT("/:id", controllers.UpdateProduct)
			productRoutes.DELETE("/:id", controllers.DeleteProductById)
		}

		customerRoutes := protected.Group("/api/customers")
		{
			customerRoutes.GET("/", controllers.GetCustomer)
			customerRoutes.GET("/:id", controllers.GetCustomerById)
			customerRoutes.POST("/", controllers.CreateCustomer)
			customerRoutes.PUT("/:id", controllers.UpdateCustomer)
			customerRoutes.DELETE("/:id", controllers.DeleteCustomerById)
		}

		oderRoutes : 

		protected.GET("/profiles", controllers.GetProfile)
	}
	return router
}
