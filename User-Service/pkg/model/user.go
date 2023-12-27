package model

type User struct {
	Id       uint   `json:"id" gorm:"not null" binding:"required"`
	Name     string `json:"name" gorm:"not null" binding:"required"`
	Email    string `json:"email" gorm:"not null" binding:"required"`
	Phone    string `json:"phone" gorm:"not null" binding:"required"`
	Password string `json:"password" gorm:"not null" binding:"required"`
}
