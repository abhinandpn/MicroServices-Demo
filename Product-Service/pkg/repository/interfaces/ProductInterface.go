package interfaces

import (
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/model"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/util/req"
)

type ProductRepository interface {

	// CURD
	AddProduct(product req.ReqProduct) error
	GetProductById(Id uint) (model.Product, error)
	DeleteProduct(Id uint) error
	UpdateProduct(product req.ReqProduct, Id uint) error
	GetProducts() ([]model.Product, error)

	// Checking
	GetProductByName(name string)(model.Product,error)
}
