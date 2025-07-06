package application

import (
    "fmt"
    
    "github.com/tLALOck64/microservicio-juegos/internal/games/domain/aggregates"
    "github.com/tLALOck64/microservicio-juegos/internal/games/domain/ports"
    "github.com/tLALOck64/microservicio-juegos/internal/games/domain/valueobjects"
)

type UpdateGameSessionDTO struct {
    ID     string
    Status string
    Score  *valueobjects.Score
}

type UpdateGameSessionUseCase struct {
    GameSessionRepository ports.GameSessionRepository
}

func NewUpdateGameSessionUseCase(
    gameSessionRepository ports.GameSessionRepository,
) *UpdateGameSessionUseCase {
    return &UpdateGameSessionUseCase{
        GameSessionRepository: gameSessionRepository,
    }
}

func (uc *UpdateGameSessionUseCase) Run(id string, status string, score *valueobjects.Score) (*aggregates.GameSession, error) {
    // Obtener sesión actual
    gameSession, err := uc.GameSessionRepository.GetById(id)
    if err != nil {
        return nil, fmt.Errorf("error al obtener sesión: %w", err)
    }
    
    // Aplicar cambios según el status solicitado
    switch status {
    case "jugando":
        if err := gameSession.ResumeGame(); err != nil {
            return nil, err
        }
    case "pausado":
        if err := gameSession.PauseGame(); err != nil {
            return nil, err
        }
    case "completado":
        if err := gameSession.CompleteGame(); err != nil {
            return nil, err
        }
    case "abandonado":
        if err := gameSession.AbandonGame(); err != nil {
            return nil, err
        }
    default:
        return nil, fmt.Errorf("estado no reconocido: %s", status)
    }
    
    // Actualizar puntuación si se proporcionó
    if score != nil {
        gameSession.UpdateScore(*score)
    }
    
    // Guardar cambios
    if err := uc.GameSessionRepository.Update(gameSession); err != nil {
        return nil, fmt.Errorf("error al guardar cambios: %w", err)
    }
    
    return gameSession, nil
}