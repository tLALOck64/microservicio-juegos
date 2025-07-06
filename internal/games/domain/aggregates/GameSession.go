package aggregates

import (
	"fmt"
	"time"

	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/valueobjects"
)

type GameSession struct {
	ID           string                     `bson:"_id,omitempty"`
	UserID       string                     `bson:"user_id"`
	MiniGameID   string                     `bson:"minigame_id"`
	Status       valueobjects.SessionStatus `bson:"status"`
	Score        valueobjects.Score         `bson:"score"`
	StartedAt    time.Time                  `bson:"started_at"`
	CompletedAt  *time.Time                 `bson:"completed_at,omitempty"`
	PausedAt     *time.Time                 `bson:"paused_at,omitempty"`
	Attempts     int                        `bson:"attempts"`
	CurrentData  map[string]interface{}     `bson:"current_data"`
	LastActivity time.Time                  `bson:"last_activity"`
}

func (gs *GameSession) StartGame() error {
	if !gs.Status.CanTransitionTo(valueobjects.StatusPlaying) {
		return fmt.Errorf("no se puede iniciar juego en estado: %s", gs.Status.GetDisplayName())
	}

	now := time.Now()
	gs.Status = valueobjects.StatusPlaying
	gs.StartedAt = now
	gs.LastActivity = now
	return nil
}

func (gs *GameSession) CompleteGame() error {
	if !gs.Status.CanTransitionTo(valueobjects.StatusCompleted) {
		return fmt.Errorf("no se puede completar juego en estado: %s", gs.Status.GetDisplayName())
	}

	gs.Status = valueobjects.StatusCompleted
	now := time.Now()
	gs.CompletedAt = &now
	gs.LastActivity = now
	return nil
}

func (gs *GameSession) AbandonGame() error {
	if !gs.Status.CanTransitionTo(valueobjects.StatusAbandoned) {
		return fmt.Errorf("no se puede abandonar juego en estado: %s", gs.Status.GetDisplayName())
	}

	gs.Status = valueobjects.StatusAbandoned
	gs.LastActivity = time.Now()
	return nil
}

func (gs *GameSession) PauseGame() error {
	if !gs.Status.CanBePaused() {
		return fmt.Errorf("no se puede pausar juego en estado: %s", gs.Status.GetDisplayName())
	}

	gs.Status = valueobjects.StatusPaused
	now := time.Now()
	gs.PausedAt = &now
	gs.LastActivity = now
	return nil
}

func (gs *GameSession) ResumeGame() error {
	if !gs.Status.CanBeResumed() {
		return fmt.Errorf("no se puede reanudar juego en estado: %s", gs.Status.GetDisplayName())
	}

	gs.Status = valueobjects.StatusPlaying
	gs.PausedAt = nil
	gs.LastActivity = time.Now()
	return nil
}

func (gs *GameSession) UpdateScore(newScore valueobjects.Score) {
	gs.Score = newScore
	gs.LastActivity = time.Now()
}

func (gs *GameSession) UpdateCurrentData(data map[string]interface{}) {
	gs.CurrentData = data
	gs.LastActivity = time.Now()
}