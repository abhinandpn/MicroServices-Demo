package usecase

import (
	"fmt"

	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/model"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/repository/interfaces"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/util/req"

	service "github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/usecase/interfaces"
)

type ProductUsecase struct {
	productRepo interfaces.ProductRepository
}

func NewProductRepository(ProductRepop interfaces.ProductRepository) service.ProductUsecase {
	return &ProductUsecase{productRepo: ProductRepop}
}

func (p *ProductUsecase) CreateProduct(product req.ReqProduct) error {

	body, err := p.productRepo.GetProductByName(product.ProductName)
	if err != nil {
		return err
	}
	if body.Id != 0 {
		res := fmt.Errorf("product alredy exist")
		return res
	} else {
		err = p.productRepo.AddProduct(product)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *ProductUsecase) GetProductById(Id uint) (model.Product, error) {

	body, err := p.productRepo.GetProductById(Id)
	if err != nil {
		return body, err
	}
	if body.Id == 0 {
		res := fmt.Errorf("product does not exist")
		return body, res
	}
	return body, nil
}

func (p *ProductUsecase) GetProducts() ([]model.Product, error) {

	body, err := p.productRepo.GetProducts()
	if err != nil {
		return body, err
	}
	if body == nil {
		res := fmt.Errorf("no product exist")
		return body, res
	}
	return body, nil
}

func (p *ProductUsecase) UpdateProduct(product req.ReqProduct, Id uint) error {

	body, err := p.productRepo.GetProductById(Id)
	if err != nil {
		return err
	}
	if body.Id != 0 {
		err := p.productRepo.UpdateProduct(product, Id)
		if err != nil {
			return err
		}
	} else {
		res := fmt.Errorf("product does not exist")
		return res
	}
	return nil
}

func (p *ProductUsecase) DeleteProduct(Id uint) error {

	body, err := p.productRepo.GetProductById(Id)
	if err != nil {
		return err
	}
	if body.Id != 0 {
		err := p.productRepo.DeleteProduct(Id)
		if err != nil {
			return err
		}
	}
	return nil
}
