package controller

import (
	"fmt"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/helper"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/http/request"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/http/response"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (controller *ProductController) Create(ctx *gin.Context) {
	createProductRequest := request.CreateProductRequest{}
	err := ctx.ShouldBindJSON(&createProductRequest)
	helper.ErrorPanic(err)
	categoryID, err := uuid.Parse(createProductRequest.CategoryID.String())
	helper.ErrorPanic(err)
	createProductRequest.CategoryID = &categoryID
	controller.productService.Create(createProductRequest)
	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)
}

func (controller *ProductController) Update(ctx *gin.Context) {
	updateProductRequest := request.UpdateProductRequest{}
	err := ctx.ShouldBindJSON(&updateProductRequest)
	productId := ctx.Param("id")
	id, err := uuid.Parse(productId)
	helper.ErrorPanic(err)
	updateProductRequest.ID = &id
	controller.productService.Update(updateProductRequest)
	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)
}

func (controller *ProductController) Delete(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	id, err := uuid.Parse(categoryId)
	helper.ErrorPanic(err)
	controller.productService.Delete(&id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductController) FindById(ctx *gin.Context) {
	productId := ctx.Param("id")
	id, err := uuid.Parse(productId)
	helper.ErrorPanic(err)
	productResponse := controller.productService.FindById(&id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductController) FindAll(ctx *gin.Context) {
	productListResponse := controller.productService.FindAll()
	fmt.Println(productListResponse)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productListResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
