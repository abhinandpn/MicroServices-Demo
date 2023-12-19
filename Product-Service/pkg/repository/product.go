package repository

import (
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/model"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/repository/interfaces"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/util/req"
	"gorm.io/gorm"
)

type ProductDatabse struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &ProductDatabse{DB: db}
}

func (p *ProductDatabse) AddProduct(product req.ReqProduct) error {

	var ProductTable model.Product
	tx := p.DB.Begin() // Transaction begin
	qry := `insert into products (product_name,
				description,
				colour,
				price)values($1,$2,$3,$4);`
	err := p.DB.Raw(qry, product.ProductName,
		product.Description,
		product.Colour,
		product.Price).Scan(&ProductTable).Error
	if err != nil { //transaction rollback
		tx.Rollback()
	}
	// transaction commit
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
	}
	return nil
}

func (p *ProductDatabse) GetProductById(Id uint) (model.Product, error) {

	var body model.Product
	qry := `select * from products where id = $1;`
	err := p.DB.Raw(qry, Id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *ProductDatabse) DeleteProduct(Id uint) error {

	var body model.Product
	qry := `delete from products where id = $1;`
	err := p.DB.Raw(qry, Id).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductDatabse) UpdateProduct(product req.ReqProduct, Id uint) error {

	var body model.Product
	qry := `UPDATE proructs
				SET 
				  product_name = $1,
				  description = $2,
				  price = $3,
				  colour = $4
				WHERE id = $5;`
	err := p.DB.Raw(qry, product.ProductName,
		product.Description,
		product.Price,
		product.Colour,
		Id).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductDatabse) GetProducts() ([]model.Product, error) {

	var body []model.Product
	qry := `select * from products`
	err := p.DB.Raw(qry).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *ProductDatabse) GetProductByName(name string) (model.Product, error) {

	var body model.Product
	qry := `select * from products where product_name = $1;`
	err := p.DB.Raw(qry, name).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}
