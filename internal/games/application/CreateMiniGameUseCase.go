package application

import (
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/entities"
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/ports"
)

type CreateMiniGameUseCase struct {
	MiniGameRepository ports.MiniGameRepository
}

func NewCreateMiniGameUseCase(miniGameRepository ports.MiniGameRepository) *CreateMiniGameUseCase {
	return &CreateMiniGameUseCase{MiniGameRepository: miniGameRepository}
}

func (mg CreateMiniGameUseCase) Run(miniGame entities.MiniGame) (*entities.MiniGame, error) {

	newMiniGame, err := mg.MiniGameRepository.Create(&miniGame)
	if err != nil {
		return &entities.MiniGame{}, err
	}

	return newMiniGame, nil

}
