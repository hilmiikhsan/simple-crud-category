package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang/go-crud-category-api/app"
	"github.com/golang/go-crud-category-api/controller"
	"github.com/golang/go-crud-category-api/exception"
	"github.com/golang/go-crud-category-api/middleware"
	"github.com/golang/go-crud-category-api/repository"
	"github.com/golang/go-crud-category-api/service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.ConnectDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
