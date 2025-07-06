package ports

import "github.com/tLALOck64/microservicio-juegos/internal/games/domain/aggregates"

type GameSessionRepository interface {
	// Crear nueva sesión de juego
	Create(gameSession *aggregates.GameSession) (*aggregates.GameSession, error)
	
	// Obtener sesión por ID
	GetById(id string) (*aggregates.GameSession, error)
	
	// Actualizar sesión existente
	Update(gameSession *aggregates.GameSession) error
}
