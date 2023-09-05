package main

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/go-rest-api/app"
	"github.com/go-rest-api/controller"
	"github.com/go-rest-api/exception"
	"github.com/go-rest-api/helpers"
	"github.com/go-rest-api/repositories"
	"github.com/go-rest-api/services"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		helpers.PanicIfError(nil)
		helpers.PanicIfError(errors.New("error"))
		helpers.PanicIfError(nil)
	})
	router.GET("/api/categories", categoryController.GetAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
