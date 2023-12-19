package interfaces

import (
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/model"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/util/req"
)

type ProductUsecase interface {
	CreateProduct(product req.ReqProduct) error
	GetProductById(Id uint) (model.Product, error)
	GetProducts() ([]model.Product, error)
	UpdateProduct(product req.ReqProduct, Id uint) error
	DeleteProduct(Id uint) error
}
