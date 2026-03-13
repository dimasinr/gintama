package controllers

import (
	"net/http"
	"time"

	"gintama/config"
	"gintama/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET = []byte("secretkey")

func Login(c *gin.Context) {

	var input models.User
	var user models.User

	c.ShouldBindJSON(&input)

	config.DB.Where("username = ?", input.Username).First(&user)

	if user.ID == 0 || user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid login"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(SECRET)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}