package ports

import "github.com/tLALOck64/microservicio-juegos/internal/games/domain/entities"

type MiniGameRepository interface {
	Create(MiniGame *entities.MiniGame) (*entities.MiniGame, error)
	Get() ([]*entities.MiniGame, error)
	GetById(id string) (*entities.MiniGame, error)
}