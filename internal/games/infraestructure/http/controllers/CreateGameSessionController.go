package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tLALOck64/microservicio-juegos/internal/games/application"
	"github.com/tLALOck64/microservicio-juegos/internal/games/infraestructure/http/mapper"
	"github.com/tLALOck64/microservicio-juegos/internal/games/infraestructure/http/request"
	"github.com/tLALOck64/microservicio-juegos/internal/shared/response"
)

type CreateGameSessionController struct {
	CreateGameSessionUseCase *application.CreateGameSessionUseCase
	Validate *validator.Validate
}


func NewCreateGameSessionController(createGameSessionUseCase *application.CreateGameSessionUseCase)*CreateGameSessionController{
	return&CreateGameSessionController{
		CreateGameSessionUseCase: createGameSessionUseCase,
		Validate: validator.New(),
	}
}


func (ctr *CreateGameSessionController) Run(ctx *gin.Context){
	var req request.CreateGameSessionRequest

	if err := ctx.ShouldBind(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "Error al validar formato",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	if err := ctr.Validate.Struct(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "Error de validación",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	gameSessionEntity, err := mapper.MapCreateGameSessionRequest(req);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "Error de tipos",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	gameSession, err := ctr.CreateGameSessionUseCase.Run(&gameSessionEntity)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message: "Error ",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Success: true,
		Message: "GameSession Created successfully",
		Data: gameSession,
		Error: nil,
	})

}