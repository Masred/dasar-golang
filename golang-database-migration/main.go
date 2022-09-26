package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Masred/dasar-golang/golang-database-migration/app"
	"github.com/Masred/dasar-golang/golang-database-migration/controller"
	"github.com/Masred/dasar-golang/golang-database-migration/helper"
	"github.com/Masred/dasar-golang/golang-database-migration/middleware"
	"github.com/Masred/dasar-golang/golang-database-migration/repository"
	"github.com/Masred/dasar-golang/golang-database-migration/service"
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
