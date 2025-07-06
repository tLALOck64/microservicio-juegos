package mapper

import (
	"fmt"

	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/aggregates"
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/entities"
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/valueobjects"
	"github.com/tLALOck64/microservicio-juegos/internal/games/infraestructure/http/request"
)


func MapCreateMiniGameRequest(req request.CreateMiniGameRequest) (entities.MiniGame, error) {
	gameType, err := valueobjects.NewGameType(req.Type)
	if err != nil {
		return entities.MiniGame{}, fmt.Errorf("tipo inválido: %w", err)
	}

	lang, err := valueobjects.NewLanguage(string(req.Language))
	if err != nil {
		return entities.MiniGame{}, fmt.Errorf("idioma inválido: %w", err)
	}

	level, err := valueobjects.NewGameLevel(req.Level)
	if err != nil {
		return entities.MiniGame{}, fmt.Errorf("nivel inválido: %w", err)
	}

	return entities.MiniGame{
		Type:        gameType,
		Language:    lang,
		Level:       level,
		ContentJSON: req.ContentJSON,
		IsActive:    req.IsActive,
	}, nil
}


func MapCreateGameSessionRequest(req request.CreateGameSessionRequest) (aggregates.GameSession, error) {

    status, _ := valueobjects.NewSessionStatus("jugando") // Siempre empieza jugando
    score, _ := valueobjects.NewScore(0, 100)            // Siempre empieza en 0

    return aggregates.GameSession{
        UserID:      req.UserID,
        MiniGameID:  req.MiniGameID,
        Status:      status,
        Score:       score,
        CurrentData: req.CurrentData, 
    }, nil
}