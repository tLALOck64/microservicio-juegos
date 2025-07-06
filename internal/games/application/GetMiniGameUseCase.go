package application

import (
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/entities"
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/ports"
)

type GetMiniGameUseCase struct {
	MiniGameRepository ports.MiniGameRepository
}

func NewGetMiniGameUseCase(miniGameRepository ports.MiniGameRepository) *GetMiniGameUseCase{
	return &GetMiniGameUseCase{MiniGameRepository: miniGameRepository}
}

func (mg *GetMiniGameUseCase) Run() ([]*entities.MiniGame, error) {
	miniGames, err := mg.MiniGameRepository.Get()
	if err != nil {
		return nil, err
	}
	
	return miniGames, nil
}