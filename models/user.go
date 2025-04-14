package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique" json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
}