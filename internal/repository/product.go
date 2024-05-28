package repository

import (
	"errors"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/helper"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/http/request"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product model.Product)
	Update(product model.Product)
	Delete(productId *uuid.UUID)
	FindById(productId *uuid.UUID) (model.Product, error)
	FindAll() []model.Product
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (cri ProductRepositoryImpl) Save(product model.Product) {
	result := cri.DB.Create(&product)
	helper.ErrorPanic(result.Error)
}

func (cri ProductRepositoryImpl) Update(product model.Product) {
	var updateProduct = request.UpdateProductRequest{
		ID:   product.ID,
		Name: product.Name,
	}
	result := cri.DB.Model(&product).Updates(updateProduct)
	helper.ErrorPanic(result.Error)
}

func (cri ProductRepositoryImpl) Delete(productId *uuid.UUID) {
	var c model.Product
	result := cri.DB.Where("id = ?", productId).Delete(&c)
	helper.ErrorPanic(result.Error)
}

func (cri ProductRepositoryImpl) FindById(productId *uuid.UUID) (model.Product, error) {
	var c model.Product
	result := cri.DB.Find(&c, productId)
	if result != nil {
		return c, nil
	}
	return c, errors.New("product is not found")
}

func (cri ProductRepositoryImpl) FindAll() []model.Product {
	var c []model.Product
	result := cri.DB.Find(&c)
	helper.ErrorPanic(result.Error)
	return c
}
