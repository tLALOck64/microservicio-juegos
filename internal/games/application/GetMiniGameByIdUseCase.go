package application

import (
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/entities"
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/ports"
)

type GetMiniGameByIdUseCase struct {
	MiniGameRepository ports.MiniGameRepository
}

func NewGetMiniGameByIdUseCase(miniGameRepository ports.MiniGameRepository) *GetMiniGameByIdUseCase{
	return&GetMiniGameByIdUseCase{MiniGameRepository: miniGameRepository}
}

func (mg *GetMiniGameByIdUseCase) Run(id string)(*entities.MiniGame, error){
	miniGame, err := mg.MiniGameRepository.GetById(id)

	if err != nil{
		return &entities.MiniGame{}, err
	}

	return miniGame, nil
}