package repository

import (
	"github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type AdminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &AdminDatabase{DB: DB}
}
