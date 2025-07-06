package application

import (
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/aggregates"
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/ports"
)

type GetGameSessionByIdUseCase struct {
	GameSessionRepository ports.GameSessionRepository
}

func NewGetGameSessionByIdUseCase(gameSessionRepository ports.GameSessionRepository)*GetGameSessionByIdUseCase{
	return &GetGameSessionByIdUseCase{GameSessionRepository: gameSessionRepository}
}

func (gs *GetGameSessionByIdUseCase) Run(id string) (*aggregates.GameSession, error){
	gameSession, err := gs.GameSessionRepository.GetById(id)

	if err != nil{
		return &aggregates.GameSession{}, err
	}

	return gameSession, nil
}