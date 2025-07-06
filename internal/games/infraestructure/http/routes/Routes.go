package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-juegos/internal/games/infraestructure/http"
	"github.com/tLALOck64/microservicio-juegos/internal/shared/middleware"
)

func Routes(router *gin.RouterGroup) {
	createController := http.SetUpCreate()
	getController := http.SetUpGet()
	getByIdController := http.SetUpGetByID()
	
	router.Use(middleware.JWTAuthMiddleware())
	router.POST("/", createController.Run)
	router.GET("/", getController.Run)
	router.GET("/:id", getByIdController.Run)
}