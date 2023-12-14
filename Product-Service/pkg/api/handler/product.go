package handler

import (
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/api/handler/interfaces"
	service "github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/usecase/interfaces"
)

type ProductHandler struct {
	productusecase service.ProductUsecase
}

func NewProductHandler(ProductUserCase service.ProductUsecase) interfaces.ProductHandler {
	return ProductHandler{productusecase: ProductUserCase}
}
