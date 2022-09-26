package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Masred/dasar-golang/golang-restful-api/app"
	"github.com/Masred/dasar-golang/golang-restful-api/controller"
	"github.com/Masred/dasar-golang/golang-restful-api/helper"
	"github.com/Masred/dasar-golang/golang-restful-api/middleware"
	"github.com/Masred/dasar-golang/golang-restful-api/repository"
	"github.com/Masred/dasar-golang/golang-restful-api/service"
	"github.com/go-playground/validator"
)

func main() {

	db := app.NewDB()
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
	helper.PanicIfError(err)
}
