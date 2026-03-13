package controllers

import (
	"net/http"

	"gintama/config"
	"gintama/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {

	var products []models.Product
	config.DB.Find(&products)

	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {

	var product models.Product

	c.ShouldBindJSON(&product)

	config.DB.Create(&product)

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {

	id := c.Param("id")

	var product models.Product

	config.DB.First(&product, id)

	c.ShouldBindJSON(&product)

	config.DB.Save(&product)

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	config.DB.Delete(&models.Product{}, id)

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}