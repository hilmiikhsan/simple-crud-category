package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang/go-crud-category-api/app"
	"github.com/golang/go-crud-category-api/controller"
	"github.com/golang/go-crud-category-api/middleware"
	"github.com/golang/go-crud-category-api/repository"
	"github.com/golang/go-crud-category-api/service"
)

func main() {
	db := app.ConnectDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
