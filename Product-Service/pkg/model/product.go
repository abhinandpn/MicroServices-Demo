package model

type Product struct {
	Id          uint    `json:"id" gorm:"primaryKey;not null"`
	ProductName string  `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Description string  `json:"description" gorm:"not null"`
	Colour      string  `json:"colour" gorm:"not null" `
	Price       float64 `json:"price" gorm:"not null"`
}
