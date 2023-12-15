package usecase

import (
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/repository/interfaces"
	services "github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/usecase/interfaces"
)

type UserUsecase struct {
	userRepo interfaces.UserReporitory
}

func NewUserUsecase(repo interfaces.UserReporitory) services.UserUsecase {
	return UserUsecase{userRepo: repo}
}
