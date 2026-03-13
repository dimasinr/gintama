package routes

import (
	"gintama/controllers"
	"gintama/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/login", controllers.Login)

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/products", controllers.GetProducts)
	auth.POST("/products", controllers.CreateProduct)
	auth.PUT("/products/:id", controllers.UpdateProduct)
	auth.DELETE("/products/:id", controllers.DeleteProduct)
}