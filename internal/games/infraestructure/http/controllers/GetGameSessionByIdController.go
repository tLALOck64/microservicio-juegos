package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-juegos/internal/games/application"
	"github.com/tLALOck64/microservicio-juegos/internal/shared/response"
)

type GetGameSessionByIdController struct {
	GetMiniSessionByIdUseCase *application.GetGameSessionByIdUseCase
}

func NewGetGameSessionByIdController(getGameSessionByIdUseCase *application.GetGameSessionByIdUseCase)*GetGameSessionByIdController{
	return &GetGameSessionByIdController{GetMiniSessionByIdUseCase: getGameSessionByIdUseCase}
}


func (ctr *GetGameSessionByIdController) Run(ctx *gin.Context){
	id := ctx.Param("id")

	gameSession, err := ctr.GetMiniSessionByIdUseCase.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message: "Invalid retrieved gamesession",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Message: "GameSession retrieved Succesfully",
		Data: gameSession,
		Error: nil,
	})
}