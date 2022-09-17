//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/Masred/dasar-golang/golang-dependency-injection/app"
	"github.com/Masred/dasar-golang/golang-dependency-injection/controller"
	"github.com/Masred/dasar-golang/golang-dependency-injection/middleware"
	"github.com/Masred/dasar-golang/golang-dependency-injection/repository"
	"github.com/Masred/dasar-golang/golang-dependency-injection/service"
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryServiceImpl,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryControllerImpl,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		app.NewServer,
	)
	return nil
}
