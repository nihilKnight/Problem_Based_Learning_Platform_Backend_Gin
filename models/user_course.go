package models

type UserCourse struct {
	UserID   uint    `gorm:"primaryKey" json:"user_id"`
	CourseID uint    `gorm:"primaryKey" json:"course_id"`
	Score    float32 `json:"score"`
	Grade    float32 `json:"grade"`
	Progress float32 `json:"progress"`
	
	User   User   `gorm:"foreignKey:UserID"`
	Course Course `gorm:"foreignKey:CourseID"`
}