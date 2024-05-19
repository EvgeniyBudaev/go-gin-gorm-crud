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

type CategoryService interface {
	Create(category request.CreateCategoryRequest)
	Update(category request.UpdateCategoryRequest)
	Delete(categoryId *uuid.UUID)
	FindById(categoryId *uuid.UUID) response.CategoryResponse
	FindAll() []response.CategoryResponse
}

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	validate           *validator.Validate
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		validate:           validate,
	}
}

func (csi CategoryServiceImpl) Create(c request.CreateCategoryRequest) {
	err := csi.validate.Struct(c)
	helper.ErrorPanic(err)
	categoryModel := model.Category{
		Name: c.Name,
	}
	csi.CategoryRepository.Save(categoryModel)
}

func (csi CategoryServiceImpl) Update(c request.UpdateCategoryRequest) {
	categoryData, err := csi.CategoryRepository.FindById(c.ID)
	helper.ErrorPanic(err)
	categoryData.Name = c.Name
	csi.CategoryRepository.Update(categoryData)
}

func (csi CategoryServiceImpl) Delete(categoryId *uuid.UUID) {
	csi.CategoryRepository.Delete(categoryId)
}

func (csi CategoryServiceImpl) FindById(categoryId *uuid.UUID) response.CategoryResponse {
	categoryData, err := csi.CategoryRepository.FindById(categoryId)
	helper.ErrorPanic(err)
	categoryResponse := response.CategoryResponse{
		ID:   categoryData.ID,
		Name: categoryData.Name,
	}
	return categoryResponse
}

func (csi CategoryServiceImpl) FindAll() []response.CategoryResponse {
	result := csi.CategoryRepository.FindAll()
	var categoryList []response.CategoryResponse
	for _, value := range result {
		c := response.CategoryResponse{
			ID:   value.ID,
			Name: value.Name,
		}
		categoryList = append(categoryList, c)
	}
	return categoryList
}
