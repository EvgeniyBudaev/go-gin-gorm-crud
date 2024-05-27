package main

import (
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/controller"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/database"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/model"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/repository"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/router"
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/services"
	"github.com/go-playground/validator/v10"

	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	db, err := database.NewConnectionToDB()
	if err != nil {
		log.Printf("Connection to DB with ERROR %d", err)
		return
	}
	err = db.Table("categories").AutoMigrate(&model.Category{})
	if err != nil {
		log.Printf("AutoMigrate categories with ERROR %d", err)
		return
	}
	err = db.Table("products").AutoMigrate(&model.Product{})
	if err != nil {
		log.Printf("AutoMigrate products with ERROR %d", err)
		return
	}

	validate := validator.New()
	// Repository
	cr := repository.NewCategoryRepositoryImpl(db)

	// Service
	cs := services.NewCategoryServiceImpl(cr, validate)

	// Controller
	cc := controller.NewCategoryController(cs)

	// Router
	r := router.NewRouter(cc)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Printf("PORT .env not parsing with ERROR %d", err)
		return
	}

	err = http.ListenAndServe(":"+strconv.Itoa(port), r)
	if err != nil {
		log.Printf("Server Not Running on Port %d with ERROR %d : \n", port, err)
		return
	}
}
