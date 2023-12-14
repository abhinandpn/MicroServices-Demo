package usecase

import (
	service "github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/usecase/interfaces"

	"github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/repository/interfaces"
)

type AdminUsecase struct {
	AdminRepo interfaces.AdminRepository
}

func NewAdminUseCase(AdminRepo interfaces.AdminRepository) service.AdminUsecase {
	return &AdminUsecase{AdminRepo: AdminRepo}
}
