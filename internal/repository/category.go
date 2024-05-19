package repository

import (
	"errors"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/helper"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/http/request"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(category model.Category)
	Update(category model.Category)
	Delete(categoryId *uuid.UUID)
	FindById(categoryId *uuid.UUID) (model.Category, error)
	FindAll() []model.Category
}

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (cri CategoryRepositoryImpl) Save(category model.Category) {
	result := cri.DB.Create(&category)
	helper.ErrorPanic(result.Error)
}

func (cri CategoryRepositoryImpl) Update(category model.Category) {
	var updateCategory = request.UpdateCategoryRequest{
		ID:   category.ID,
		Name: category.Name,
	}
	result := cri.DB.Model(&category).Updates(updateCategory)
	helper.ErrorPanic(result.Error)
}

func (cri CategoryRepositoryImpl) Delete(categoryId *uuid.UUID) {
	var c model.Category
	result := cri.DB.Where("id = ?", categoryId).Delete(&c)
	helper.ErrorPanic(result.Error)
}

func (cri CategoryRepositoryImpl) FindById(categoryId *uuid.UUID) (model.Category, error) {
	var c model.Category
	result := cri.DB.Find(&c, categoryId)
	if result != nil {
		return c, nil
	}
	return c, errors.New("category is not found")
}

func (cri CategoryRepositoryImpl) FindAll() []model.Category {
	var c []model.Category
	result := cri.DB.Find(&c)
	helper.ErrorPanic(result.Error)
	return c
}
