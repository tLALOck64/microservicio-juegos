package entities

import "github.com/tLALOck64/microservicio-juegos/internal/games/domain/valueobjects"

type MiniGame struct {
	ID          string
	Type        valueobjects.GameType
	Language    valueobjects.Language
	Level       valueobjects.GameLevel
	ContentJSON map[string]interface{}
	IsActive    bool
}