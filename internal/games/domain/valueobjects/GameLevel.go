package valueobjects

import "fmt"

type GameLevel string

const (
	Beginner     GameLevel = "principiante"
	Intermediate GameLevel = "intermedio"
	Advanced     GameLevel = "avanzado"
	Expert       GameLevel = "experto"
)

func NewGameLevel(value string) (GameLevel, error) {
	level := GameLevel(value)
	if !level.IsValid() {
		return "", fmt.Errorf("nivel inválido: %s. Niveles válidos: principiante, intermedio, avanzado, experto", level)
	}

	return level, nil
}

func (gl GameLevel) IsValid() bool {
	return gl == Beginner || gl == Intermediate || gl == Advanced || gl == Expert
}