package service

import (
	"sesi-2/helper"
	"sesi-2/model"
	"sesi-2/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (ps *ProductService) Create(request model.ProductCreateRequest, userId string) (model.ProductCreateResponse, error) {
	id := helper.GenerateID()

	product := model.Product{
		Id:     id,
		UserId: userId,
		Price:  request.Price,
	}

	err := ps.ProductRepository.Add(product)
	return model.ProductCreateResponse{
		Id:     id,
		UserId: userId,
		Price:  request.Price,
	}, err
}

func (ps *ProductService) Get() ([]model.Product, error) {
	var products []model.Product
	return products, nil
}

func (ps *ProductService) GetById(getProduct *model.Product) error {
	return nil
}
