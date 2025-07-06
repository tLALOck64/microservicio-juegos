package models

import (
	"time"

	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/entities"
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/valueobjects"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MiniGameModel representa la estructura en MongoDB
type MiniGameModel struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty"`
	Type        string                 `bson:"type"`
	Language    string                 `bson:"language"`
	Level       string                 `bson:"level"`
	ContentJSON map[string]interface{} `bson:"content_json"`
	IsActive    bool                   `bson:"is_active"`
	CreatedAt   time.Time              `bson:"created_at"`
	UpdatedAt   time.Time              `bson:"updated_at"`
}

// ToDomainMiniGame convierte el modelo MongoDB a entidad de dominio
func ToDomainMiniGame(model *MiniGameModel) *entities.MiniGame {
	// Convertir strings a value objects
	gameType, _ := valueobjects.NewGameType(model.Type)
	language, _ := valueobjects.NewLanguage(model.Language)
	level, _ := valueobjects.NewGameLevel(model.Level)

	return &entities.MiniGame{
		ID:          model.ID.Hex(), // ObjectID a string
		Type:        gameType,
		Language:    language,
		Level:       level,
		ContentJSON: model.ContentJSON,
		IsActive:    model.IsActive,
	}
}

// FromDomainMiniGame convierte entidad de dominio a modelo MongoDB
func FromDomainMiniGame(entity *entities.MiniGame) *MiniGameModel {
	var objectID primitive.ObjectID

	// Si tiene ID, convertir de string a ObjectID
	if entity.ID != "" {
		if oid, err := primitive.ObjectIDFromHex(entity.ID); err == nil {
			objectID = oid
		}
	}

	return &MiniGameModel{
		ID:          objectID,
		Type:        string(entity.Type),     // value object a string
		Language:    string(entity.Language), // value object a string
		Level:       string(entity.Level),    // value object a string
		ContentJSON: entity.ContentJSON,
		IsActive:    entity.IsActive,
	}
}
