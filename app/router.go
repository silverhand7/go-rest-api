package app

import (
	"errors"
	"net/http"

	"github.com/go-rest-api/controller"
	"github.com/go-rest-api/exception"
	"github.com/go-rest-api/helpers"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
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
	return router
}
