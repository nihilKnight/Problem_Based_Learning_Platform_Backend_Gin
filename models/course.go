package models

type Course struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Video       string `json:"video"`
	Doc         string `json:"doc"`
	Description string `json:"description"`
	CreatedAt   int64  `gorm:"autoCreateTime" json:"created_at"`
}