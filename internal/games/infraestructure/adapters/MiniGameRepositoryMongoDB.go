package adapters

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tLALOck64/microservicio-juegos/internal/database"
	"github.com/tLALOck64/microservicio-juegos/internal/database/models"
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/entities"
)

type MiniGameRepositoryMongoDB struct {
	DB *database.MongoDB
}

func NewMiniGameRepositoryMongoDB() (*MiniGameRepositoryMongoDB, error) {
	db, err := database.Connect()
	if err != nil {
		panic("Error connecting to MongoDB: " + err.Error())
	}

	return &MiniGameRepositoryMongoDB{DB: db}, nil
}

func (r *MiniGameRepositoryMongoDB) Create(miniGame *entities.MiniGame) (*entities.MiniGame, error) {
	collection := r.DB.Database.Collection("minigames")

	model := models.FromDomainMiniGame(miniGame)

	if model.ID.IsZero() {
		model.ID = primitive.NewObjectID()
	}

	now := time.Now()
	model.CreatedAt = now
	model.UpdatedAt = now

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		model.ID = oid
	}

	miniGame.ID = model.ID.Hex()

	return miniGame, nil
}

func (r *MiniGameRepositoryMongoDB) GetById(id string) (*entities.MiniGame, error) {
	collection := r.DB.Database.Collection("minigames")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err, 1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var model models.MiniGameModel
	filter := bson.M{"_id": objectID}

	err = collection.FindOne(ctx, filter).Decode(&model)
	if err != nil {
		return &entities.MiniGame{}, err
	}

	return models.ToDomainMiniGame(&model), nil
}

func (r *MiniGameRepositoryMongoDB) Get() ([]*entities.MiniGame, error) {
	collection := r.DB.Database.Collection("minigames")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filtro para obtener solo juegos activos
	filter := bson.M{"is_active": true}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error al obtener minijuegos: %w", err)
	}
	defer cursor.Close(ctx)

	var modelos []models.MiniGameModel
	if err = cursor.All(ctx, &modelos); err != nil {
		return nil, fmt.Errorf("error al decodificar minijuegos: %w", err)
	}

	// Convertir modelos a entidades de dominio
	miniGames := make([]*entities.MiniGame, len(modelos))
	for i, model := range modelos {
		miniGames[i] = models.ToDomainMiniGame(&model)
	}

	return miniGames, nil
}
