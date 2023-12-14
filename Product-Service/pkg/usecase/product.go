package usecase

import (
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/repository/interfaces"

	service "github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/usecase/interfaces"
)

type ProductUsecase struct {
	productRepo interfaces.ProductRepository
}

func NewProductRepository(ProductRepop interfaces.ProductRepository) service.ProductUsecase {
	return ProductUsecase{productRepo: ProductRepop}
}
