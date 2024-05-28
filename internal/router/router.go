package router

import (
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(cc *controller.CategoryController, pc *controller.ProductController) *gin.Engine {
	router := gin.Default()
	baseRouter := router.Group("/api/v1")

	categoryRouter := baseRouter.Group("/category")
	categoryRouter.GET("", cc.FindAll)
	categoryRouter.GET("/:id", cc.FindById)
	categoryRouter.POST("", cc.Create)
	categoryRouter.PATCH("/:id", cc.Update)
	categoryRouter.DELETE("/:id", cc.Delete)

	productRouter := baseRouter.Group("/product")
	productRouter.GET("", pc.FindAll)
	productRouter.GET("/:id", pc.FindById)
	productRouter.POST("", pc.Create)
	productRouter.PATCH("/:id", pc.Update)
	productRouter.DELETE("/:id", pc.Delete)

	return router
}
