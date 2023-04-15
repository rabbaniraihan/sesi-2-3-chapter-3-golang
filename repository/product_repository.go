package repository

import (
	"sesi-2/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) Add(newProduct model.Product) error {
	tx := pr.db.Create(&newProduct)
	return tx.Error
}

func (pr *ProductRepository) Get(products model.Product) error {
	tx := pr.db.Find(&products)
	return tx.Error
}

func (pr *ProductRepository) GetById(getProduct model.Product) error {
	product := &model.Product{}
	tx := pr.db.First(product, getProduct)
	return tx.Error
}
