package services

import (
	"github.com/nihilKnight/pbl-platform-backend-gin/models"

	"github.com/nihilKnight/pbl-platform-backend-gin/database"
)

func CreateUser(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}

func DeleteUser(id uint) error {
	result := database.DB.Delete(&models.User{}, id)
	return result.Error
}

func UpdateUser(user *models.User) error {
	result := database.DB.Save(user)
	return result.Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	return &user, result.Error
}


func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}