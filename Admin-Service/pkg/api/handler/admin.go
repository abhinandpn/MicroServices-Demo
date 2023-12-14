package handler

import (
	handlerinterface "github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/api/handler/interfaces"
	services "github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/usecase/interfaces"
)

type AdminHandler struct {
	AdminUsecase services.AdminUsecase
}

func NewAdminHandler(AdminUsecase services.AdminUsecase) handlerinterface.Adminhandler {
	return &AdminHandler{AdminUsecase: AdminUsecase}
}
