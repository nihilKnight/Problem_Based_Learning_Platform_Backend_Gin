package services

import (
	"github.com/nihilKnight/pbl-platform-backend-gin/models"

	"github.com/nihilKnight/pbl-platform-backend-gin/database"
)

func CreateUser(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	return &user, result.Error
}