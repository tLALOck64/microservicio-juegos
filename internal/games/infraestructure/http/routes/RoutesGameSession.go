package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-juegos/internal/games/infraestructure/http"
	"github.com/tLALOck64/microservicio-juegos/internal/shared/middleware"
)

func RoutesGameSession(router *gin.RouterGroup){
	createController := http.SetUpCreateGameSession()
	getByIdController := http.SetUpGetByIDGameSession()
	updateController := http.SetUpUpdateGameSession()

	router.Use(middleware.JWTAuthMiddleware())
	router.POST("/", createController.Run)
	router.GET("/:id", getByIdController.Run)
	router.PUT("/:id", updateController.Run)
}