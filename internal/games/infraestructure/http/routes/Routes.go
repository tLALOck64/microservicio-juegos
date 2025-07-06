package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-juegos/internal/games/infraestructure/http"
)

func Routes(route *gin.RouterGroup) {
	createController := http.SetUpCreate()
	getController := http.SetUpGet()
	getByIdController := http.SetUpGetByID()

	route.POST("/", createController.Run)
	route.GET("/", getController.Run)
	route.GET("/:id", getByIdController.Run)
}