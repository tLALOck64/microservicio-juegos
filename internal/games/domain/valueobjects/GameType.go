package valueobjects

import "fmt"

type GameType string

const (
	Memorama    GameType = "memorama"
	Listening   GameType = "escucha"
	Translation GameType = "traduccion"
)

func NewGameType(value string) (GameType, error) {
	gameType := GameType(value)
	if !gameType.IsValid() {
		return "", fmt.Errorf("tipo de juego inválido: %s", value)
	}
	return gameType, nil
}

func (gt GameType) IsValid() bool {
	validTypes := []GameType{
		Memorama, Listening, Translation,
	}

	for _, valid := range validTypes {
		if gt == valid {
			return true
		}
	}
	return false
}