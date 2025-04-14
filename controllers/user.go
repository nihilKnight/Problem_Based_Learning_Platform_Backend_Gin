package controllers

import (
	"net/http"

	"github.com/nihilKnight/pbl-platform-backend-gin/utils"

	"github.com/nihilKnight/pbl-platform-backend-gin/services"

	"github.com/gin-gonic/gin"
	"github.com/nihilKnight/pbl-platform-backend-gin/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password hashing failed"})
		return
	}
	user.Password = hashedPassword

	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}