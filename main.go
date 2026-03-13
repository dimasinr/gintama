package main

import (
	"gintama/config"
	"gintama/models"
	"gintama/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{}, &models.Product{})

	routes.SetupRoutes(r)

	r.Run(":8080")
}