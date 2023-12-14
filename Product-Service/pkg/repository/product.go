package repository

import (
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type ProductDatabse struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return ProductDatabse{DB: db}
}
