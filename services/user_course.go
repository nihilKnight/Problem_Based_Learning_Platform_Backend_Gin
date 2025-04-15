package services

import (
	"github.com/nihilKnight/pbl-platform-backend-gin/models"

	"github.com/nihilKnight/pbl-platform-backend-gin/database"
)

func CreateUserCourse(userCourse *models.UserCourse) error {
	result := database.DB.Create(userCourse)
	return result.Error
}

func DeleteUserCourse(id uint) error {
	result := database.DB.Delete(&models.UserCourse{}, id)
	return result.Error
}

func UpdateUserCourse(userCourse *models.UserCourse) error {
	result := database.DB.Save(userCourse)
	return result.Error
}

// PROBLEM: multiple primary key in UserCourse, need to use composite primary key
func GetUserCourseByID(id uint) (*models.UserCourse, error) {
	var userCourse models.UserCourse
	result := database.DB.First(&userCourse, id)
	return &userCourse, result.Error
}

func GetAllUserCourses() ([]models.UserCourse, error) {
	var userCourses []models.UserCourse
	result := database.DB.Find(&userCourses)
	return userCourses, result.Error
}