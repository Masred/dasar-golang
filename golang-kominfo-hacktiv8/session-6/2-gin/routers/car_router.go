package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/masred/dasar-golang/golang-kominfo-hactiv8/session-6/2-gin/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
