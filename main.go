package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/go-rest-api/app"
	"github.com/go-rest-api/controller"
	"github.com/go-rest-api/helpers"
	"github.com/go-rest-api/middleware"
	"github.com/go-rest-api/repositories"
	"github.com/go-rest-api/services"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(app.NewRouter(categoryController)),
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
