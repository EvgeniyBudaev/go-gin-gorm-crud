package services

import (
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/helper"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/http/request"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/http/response"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/model"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductService interface {
	Create(product request.CreateProductRequest)
	Update(product request.UpdateProductRequest)
	Delete(productId *uuid.UUID)
	FindById(productId *uuid.UUID) response.ProductResponse
	FindAll() []response.ProductResponse
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	validate          *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		validate:          validate,
	}
}

func (csi ProductServiceImpl) Create(c request.CreateProductRequest) {
	err := csi.validate.Struct(c)
	helper.ErrorPanic(err)
	productModel := model.Product{
		Name:       c.Name,
		CategoryID: c.CategoryID,
	}
	csi.ProductRepository.Save(productModel)
}

func (csi ProductServiceImpl) Update(c request.UpdateProductRequest) {
	productData, err := csi.ProductRepository.FindById(c.ID)
	helper.ErrorPanic(err)
	productData.Name = c.Name
	csi.ProductRepository.Update(productData)
}

func (csi ProductServiceImpl) Delete(productId *uuid.UUID) {
	csi.ProductRepository.Delete(productId)
}

func (csi ProductServiceImpl) FindById(productId *uuid.UUID) response.ProductResponse {
	productData, err := csi.ProductRepository.FindById(productId)
	helper.ErrorPanic(err)
	productResponse := response.ProductResponse{
		ID:         productData.ID,
		Name:       productData.Name,
		CategoryID: productData.CategoryID,
	}
	return productResponse
}

func (csi ProductServiceImpl) FindAll() []response.ProductResponse {
	result := csi.ProductRepository.FindAll()
	var productList []response.ProductResponse
	for _, value := range result {
		c := response.ProductResponse{
			ID:         value.ID,
			Name:       value.Name,
			CategoryID: value.CategoryID,
		}
		productList = append(productList, c)
	}
	return productList
}
