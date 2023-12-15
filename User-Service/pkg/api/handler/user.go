package handler

import(
	services "github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/usecase/interfaces"
	handlerinterface "github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/api/handler/interfaces"

)
type UserHandler struct{
	UserUsecase services.UserUsecase
}

func NewUserUseHandler (usecase services.UserUsecase)handlerinterface.UserHandler{
	return UserHandler{UserUsecase: usecase}
}