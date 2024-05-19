package router

import (
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(categoryController *controller.CategoryController) *gin.Engine {
	router := gin.Default()
	baseRouter := router.Group("/api/v1")
	categoryRouter := baseRouter.Group("/category")
	categoryRouter.GET("", categoryController.FundAll)
	categoryRouter.GET("/:id", categoryController.FundById)
	categoryRouter.POST("", categoryController.Create)
	categoryRouter.PATCH("/:id", categoryController.Update)
	categoryRouter.DELETE("/:id", categoryController.Delete)
	return router
}
