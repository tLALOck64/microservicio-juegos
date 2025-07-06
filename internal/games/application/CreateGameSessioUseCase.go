package application

import (
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/aggregates"
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/ports"
)

type CreateGameSessionUseCase struct {
	GameSessionRepository ports.GameSessionRepository
}

func NewCreateGameSessionUseCase(gameSessionRepository ports.GameSessionRepository)*CreateGameSessionUseCase{
	return&CreateGameSessionUseCase{GameSessionRepository: gameSessionRepository}
}

func (gs *CreateGameSessionUseCase) Run(gameSession *aggregates.GameSession)(*aggregates.GameSession, error){
	newGameSession, err := gs.GameSessionRepository.Create(gameSession)

	if err != nil{
		return &aggregates.GameSession{}, err
	}

	return newGameSession, nil
}