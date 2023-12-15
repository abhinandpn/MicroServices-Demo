package repository

import (
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type UserDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserReporitory {
	return UserDatabase{DB: DB}
}
