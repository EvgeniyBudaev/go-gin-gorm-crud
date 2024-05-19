package controller

import (
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/helper"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/http/request"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/http/response"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (controller *CategoryController) Create(ctx *gin.Context) {
	createCategoryRequest := request.CreateCategoryRequest{}
	err := ctx.ShouldBindJSON(&createCategoryRequest)
	helper.ErrorPanic(err)
	controller.categoryService.Create(createCategoryRequest)
	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)
}

func (controller *CategoryController) Update(ctx *gin.Context) {
	updateCategoryRequest := request.UpdateCategoryRequest{}
	err := ctx.ShouldBindJSON(&updateCategoryRequest)
	categoryId := ctx.Param("id")
	id, err := uuid.Parse(categoryId)
	helper.ErrorPanic(err)
	updateCategoryRequest.ID = &id
	controller.categoryService.Update(updateCategoryRequest)
	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)
}

func (controller *CategoryController) Delete(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	id, err := uuid.Parse(categoryId)
	helper.ErrorPanic(err)
	controller.categoryService.Delete(&id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryController) FundById(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	id, err := uuid.Parse(categoryId)
	helper.ErrorPanic(err)
	categoryResponse := controller.categoryService.FindById(&id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryController) FundAll(ctx *gin.Context) {
	categoryListResponse := controller.categoryService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryListResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
