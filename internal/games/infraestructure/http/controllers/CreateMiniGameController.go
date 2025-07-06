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

type CreateMiniGameController struct {
	CreateMiniGameUseCase *application.CreateMiniGameUseCase
	Validate *validator.Validate
}

func NewCreateMiniGameControlle(createMiniGameUsecase *application.CreateMiniGameUseCase) *CreateMiniGameController{
	return &CreateMiniGameController{
		CreateMiniGameUseCase: createMiniGameUsecase,
		Validate: validator.New(),
	}
}


func (ctr *CreateMiniGameController) Run(ctx *gin.Context){
	var req request.CreateMiniGameRequest

	if err := ctx.ShouldBind(&req); err != nil{
        ctx.JSON(http.StatusBadRequest, response.Response{
            Success: false,
            Message: "Error en el formato de la petición",
            Data: nil,
            Error: err.Error(),
        })
        return 
    }


	if err := ctr.Validate.Struct(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, response.Response{
            Success: false,
            Message: "Error de validación",
            Data: nil,
            Error: err.Error(),
        })
        return
    }

	miniGameEntity, err := mapper.MapCreateMiniGameRequest(req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "Error al validar los tipos",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	miniGame, err := ctr.CreateMiniGameUseCase.Run(miniGameEntity)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message: "Error al crear minijuego",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Success: true,
		Message: "Minijuego creado exitosamente",
		Data: miniGame,
		Error: nil,
	})
}