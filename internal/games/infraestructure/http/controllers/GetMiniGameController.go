package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-juegos/internal/games/application"
	"github.com/tLALOck64/microservicio-juegos/internal/shared/response"
)

type GetMiniGameController struct {
	GetMiniGameUseCase *application.GetMiniGameUseCase
}

func NewGetMiniGameController(getMiniGameUseCase *application.GetMiniGameUseCase) *GetMiniGameController{
	return &GetMiniGameController{GetMiniGameUseCase: getMiniGameUseCase}
}

func(ctr *GetMiniGameController) Run(ctx *gin.Context){
	miniGames, err := ctr.GetMiniGameUseCase.Run()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message: "Error retrieved minigames",
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Message: "Successfully retrieved minigames",
		Data: miniGames,
		Error: nil,
	})
}