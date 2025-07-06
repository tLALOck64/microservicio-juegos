package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-juegos/internal/games/infraestructure/http"
)

func RoutesGameSession(route *gin.RouterGroup){
	createController := http.SetUpCreateGameSession()
	getByIdController := http.SetUpGetByIDGameSession()
	updateController := http.SetUpUpdateGameSession()

	route.POST("/", createController.Run)
	route.GET("/:id", getByIdController.Run)
	route.PUT("/:id", updateController.Run)
}