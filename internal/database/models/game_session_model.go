package models

import (
    "time"

    "github.com/tLALOck64/microservicio-juegos/internal/games/domain/aggregates"
    "github.com/tLALOck64/microservicio-juegos/internal/games/domain/valueobjects"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// GameSessionModel representa la estructura en MongoDB
type GameSessionModel struct {
    ID           primitive.ObjectID     `bson:"_id,omitempty"`
    UserID       string                 `bson:"user_id"`
    MiniGameID   primitive.ObjectID     `bson:"minigame_id"`
    Status       string                 `bson:"status"`
    Score        ScoreModel             `bson:"score"`
    StartedAt    time.Time              `bson:"started_at"`
    CompletedAt  *time.Time             `bson:"completed_at,omitempty"`
    PausedAt     *time.Time             `bson:"paused_at,omitempty"`
    Attempts     int                    `bson:"attempts"`
    CurrentData  map[string]interface{} `bson:"current_data"`
    LastActivity time.Time              `bson:"last_activity"`
    CreatedAt    time.Time              `bson:"created_at"`
    UpdatedAt    time.Time              `bson:"updated_at"`
}

// ScoreModel representa la puntuación en MongoDB
type ScoreModel struct {
    Current     int `bson:"current"`
    MaxPossible int `bson:"max_possible"`
}

// ToDomainGameSession convierte el modelo MongoDB a agregado de dominio
func ToDomainGameSession(model *GameSessionModel) *aggregates.GameSession {
    // Convertir strings a value objects
    status, _ := valueobjects.NewSessionStatus(model.Status)
    score, _ := valueobjects.NewScore(model.Score.Current, model.Score.MaxPossible)

    return &aggregates.GameSession{
        ID:           model.ID.Hex(), // ObjectID a string
        UserID:       model.UserID,
        MiniGameID:   model.MiniGameID.Hex(), // ObjectID a string
        Status:       status,
        Score:        score,
        StartedAt:    model.StartedAt,
        CompletedAt:  model.CompletedAt,
        PausedAt:     model.PausedAt,
        Attempts:     model.Attempts,
        CurrentData:  model.CurrentData,
        LastActivity: model.LastActivity,
    }
}

// FromDomainGameSession convierte agregado de dominio a modelo MongoDB
func FromDomainGameSession(entity *aggregates.GameSession) *GameSessionModel {
    var objectID primitive.ObjectID
    var miniGameObjectID primitive.ObjectID

    // Si tiene ID, convertir de string a ObjectID
    if entity.ID != "" {
        if oid, err := primitive.ObjectIDFromHex(entity.ID); err == nil {
            objectID = oid
        }
    }

    // Convertir MiniGameID de string a ObjectID
    if entity.MiniGameID != "" {
        if oid, err := primitive.ObjectIDFromHex(entity.MiniGameID); err == nil {
            miniGameObjectID = oid
        }
    }

    return &GameSessionModel{
        ID:           objectID,
        UserID:       entity.UserID,
        MiniGameID:   miniGameObjectID,
        Status:       string(entity.Status), // value object a string
        Score: ScoreModel{
            Current:     entity.Score.Current(),
            MaxPossible: entity.Score.MaxPossible(),
        },
        StartedAt:    entity.StartedAt,
        CompletedAt:  entity.CompletedAt,
        PausedAt:     entity.PausedAt,
        Attempts:     entity.Attempts,
        CurrentData:  entity.CurrentData,
        LastActivity: entity.LastActivity,
    }
}