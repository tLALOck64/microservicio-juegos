package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "github.com/tLALOck64/microservicio-juegos/internal/games/application"
    "github.com/tLALOck64/microservicio-juegos/internal/games/domain/valueobjects"
    "github.com/tLALOck64/microservicio-juegos/internal/games/infraestructure/http/request"
    "github.com/tLALOck64/microservicio-juegos/internal/shared/response"
)

type UpdateGameSessionController struct {
    UpdateGameSessionUseCase *application.UpdateGameSessionUseCase // ✅ Solo este use case
    Validate                 *validator.Validate
}

func NewUpdateGameSessionController(
    updateGameSessionUseCase *application.UpdateGameSessionUseCase,
) *UpdateGameSessionController {
    return &UpdateGameSessionController{
        UpdateGameSessionUseCase: updateGameSessionUseCase,
        Validate:                 validator.New(),
    }
}

func (ctr *UpdateGameSessionController) Run(ctx *gin.Context) {
    // Obtener ID de la sesión desde URL
    id := ctx.Param("id")
    if id == "" {
        ctx.JSON(http.StatusBadRequest, response.Response{
            Success: false,
            Message: "ID de sesión requerido",
            Data:    nil,
            Error:   "session_id is required",
        })
        return
    }

    var req request.UpdateGameSessionRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, response.Response{
            Success: false,
            Message: "Error en el formato de la petición",
            Data:    nil,
            Error:   err.Error(),
        })
        return
    }

    if err := ctr.Validate.Struct(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, response.Response{
            Success: false,
            Message: "Error de validación",
            Data:    nil,
            Error:   err.Error(),
        })
        return
    }

    // Crear score si se proporcionó
    var score *valueobjects.Score
    if req.Score > 0 {
        scoreObj, err := valueobjects.NewScore(req.Score, 100) // Asumimos máximo de 100
        if err != nil {
            ctx.JSON(http.StatusBadRequest, response.Response{
                Success: false,
                Message: "Score inválido",
                Data:    nil,
                Error:   err.Error(),
            })
            return
        }
        score = &scoreObj
    }

    // Ejecutar use case con el status deseado
    gameSession, err := ctr.UpdateGameSessionUseCase.Run(id, req.Status, score)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, response.Response{
            Success: false,
            Message: "Error al actualizar sesión de juego",
            Data:    nil,
            Error:   err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, response.Response{
        Success: true,
        Message: "Sesión de juego actualizada exitosamente",
        Data:    gameSession,
        Error:   nil,
    })
}