package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-juegos/internal/games/application"
	"github.com/tLALOck64/microservicio-juegos/internal/shared/response"
)

type GetMiniGameByIdController struct {
	GetMiniGameByIdUseCase *application.GetMiniGameByIdUseCase
}


func NewGetMiniGameByIdController(getMiniGameByIdUseCase *application.GetMiniGameByIdUseCase) *GetMiniGameByIdController{
	return &GetMiniGameByIdController{GetMiniGameByIdUseCase: getMiniGameByIdUseCase}
}

func (ctr *GetMiniGameByIdController) Run(ctx *gin.Context){
	id := ctx.Param("id")

	miniGame, err := ctr.GetMiniGameByIdUseCase.Run(id)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message:"Error retrieving miniGame",
			Data:  nil,
			Error: err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Message: "MiniGame retrieved succesfully",
		Data: miniGame,
		Error: nil,
	})
}