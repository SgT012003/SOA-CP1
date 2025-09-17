package service

import (
	"marketplace-soa/dao"
	"marketplace-soa/model"
)

type ProductService interface {
	Create(product model.Product) (int, error)
	Update(product model.Product) error
	Delete(id int) error
	GetByID(id int) (model.Product, error)
	GetAll() ([]model.Product, error)
}

type productService struct{}

func NewProductService() ProductService {
	return &productService{}
}

func (s *productService) Create(product model.Product) (int, error) {
	return dao.InsertProduct(product)
}

func (s *productService) Update(product model.Product) error {
	return dao.UpdateProduct(product)
}

func (s *productService) Delete(id int) error {
	return dao.DeleteProduct(id)
}

func (s *productService) GetByID(id int) (model.Product, error) {
	return dao.GetProductByID(id)
}

func (s *productService) GetAll() ([]model.Product, error) {
	return dao.GetAllProducts()
}
