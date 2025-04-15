package services

import (
	"github.com/nihilKnight/pbl-platform-backend-gin/models"

	"github.com/nihilKnight/pbl-platform-backend-gin/database"
)

func CreateCourse(course *models.Course) error {
	result := database.DB.Create(course)
	return result.Error
}

func DeleteCourse(id uint) error {
	result := database.DB.Delete(&models.Course{}, id)
	return result.Error
}

func UpdateCourse(course *models.Course) error {
	result := database.DB.Save(course)
	return result.Error
}

func GetCourseByID(id uint) (*models.Course, error) {
	var course models.Course
	result := database.DB.First(&course, id)
	return &course, result.Error
}

func GetAllCourses() ([]models.Course, error) {
	var courses []models.Course
	result := database.DB.Find(&courses)
	return courses, result.Error
}