package valueobjects

import "fmt"

type SessionStatus string

const (
	StatusWaiting     SessionStatus = "esperando"
	StatusPlaying     SessionStatus = "jugando"
	StatusPaused      SessionStatus = "pausado"
	StatusCompleted   SessionStatus = "completado"
	StatusAbandoned SessionStatus = "abandonado"
	StatusExpired     SessionStatus = "expirado"
)

func NewSessionStatus(value string) (SessionStatus, error) {
	status := SessionStatus(value)
	if !status.IsValid() {
		return "", fmt.Errorf("estado de sesión inválido: %s. Estados válidos: waiting, playing, paused, completed, abandoned, expired", value)
	}

	return status, nil
}

func (ss SessionStatus) IsValid() bool {
	validStatus := []SessionStatus{
		StatusWaiting, StatusPlaying, StatusCompleted,
		StatusPaused, StatusAbandoned, StatusExpired,
	}

	for _, valid := range validStatus {
		if ss == valid {
			return true
		}
	}

	return false
}

func (ss SessionStatus) CanTransitionTo(newStatus SessionStatus) bool {
    validTransitions := map[SessionStatus][]SessionStatus{
        StatusWaiting: {StatusPlaying, StatusAbandoned, StatusExpired},
        StatusPlaying: {StatusPaused, StatusCompleted, StatusAbandoned, StatusExpired},
        StatusPaused:  {StatusPlaying, StatusAbandoned, StatusExpired},
        // Estados finales no pueden cambiar
        StatusCompleted: {},
        StatusAbandoned: {},
        StatusExpired:   {},
    }
    
    allowedTransitions := validTransitions[ss]
    for _, allowed := range allowedTransitions {
        if newStatus == allowed {
            return true
        }
    }
    return false
}


func (ss SessionStatus) GetDisplayName() string {
    names := map[SessionStatus]string{
        StatusWaiting:   "Esperando ",
        StatusPlaying:   "Jugando ",
        StatusPaused:    "Pausado ⏸",
        StatusCompleted: "Completado ",
        StatusAbandoned: "Abandonado ",
        StatusExpired:   "Expirado ",
    }
    return names[ss]
}

// comportamientos


func (ss SessionStatus) IsActive() bool {
    return ss == StatusWaiting || ss == StatusPlaying || ss == StatusPaused
}

func (ss SessionStatus) IsFinished() bool {
    return ss == StatusCompleted || ss == StatusAbandoned || ss == StatusExpired
}

func (ss SessionStatus) IsSuccessful() bool {
    return ss == StatusCompleted
}

func (ss SessionStatus) CanBePlayed() bool {
    return ss == StatusWaiting || ss == StatusPaused
}

func (ss SessionStatus) CanBePaused() bool {
    return ss == StatusPlaying
}

func (ss SessionStatus) CanBeResumed() bool {
    return ss == StatusPaused
}