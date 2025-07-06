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
    "github.com/tLALOck64/microservicio-juegos/internal/games/domain/aggregates"
)

type GameSessionRepositoryMongoDB struct {
    DB *database.MongoDB
}

func NewGameSessionRepositoryMongoDB() (*GameSessionRepositoryMongoDB, error) {
    db, err := database.Connect()
    if err != nil {
        panic("Error connecting to MongoDB: " + err.Error())
    }

    return &GameSessionRepositoryMongoDB{DB: db}, nil
}

func (r *GameSessionRepositoryMongoDB) Create(gameSession *aggregates.GameSession) (*aggregates.GameSession, error) {
    collection := r.DB.Database.Collection("game_sessions")

    model := models.FromDomainGameSession(gameSession)

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

    gameSession.ID = model.ID.Hex()

    return gameSession, nil
}

func (r *GameSessionRepositoryMongoDB) GetById(id string) (*aggregates.GameSession, error) {
    collection := r.DB.Database.Collection("game_sessions")

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Fatal(err, 1)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var model models.GameSessionModel
    filter := bson.M{"_id": objectID}

    err = collection.FindOne(ctx, filter).Decode(&model)
    if err != nil {
        return &aggregates.GameSession{}, err
    }

    return models.ToDomainGameSession(&model), nil
}

func (r *GameSessionRepositoryMongoDB) Update(gameSession *aggregates.GameSession) error {
    collection := r.DB.Database.Collection("game_sessions")

    objectID, err := primitive.ObjectIDFromHex(gameSession.ID)
    if err != nil {
        log.Fatal(err, 1)
    }

    model := models.FromDomainGameSession(gameSession)
    model.UpdatedAt = time.Now()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{"_id": objectID}
    update := bson.M{"$set": model}

    result, err := collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    if result.MatchedCount == 0 {
        return fmt.Errorf("no game session found with id: %s", gameSession.ID)
    }

    fmt.Printf("Game session updated successfully. Modified count: %d\n", result.ModifiedCount)

    return nil
}

func (r *GameSessionRepositoryMongoDB) GetActiveByUserId(userID string) ([]*aggregates.GameSession, error) {
    collection := r.DB.Database.Collection("game_sessions")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Filtro para sesiones activas (jugando o pausadas)
    filter := bson.M{
        "user_id": userID,
        "status": bson.M{"$in": []string{"jugando", "pausado"}},
    }

    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        return nil, fmt.Errorf("error al obtener sesiones activas: %w", err)
    }
    defer cursor.Close(ctx)

    var modelos []models.GameSessionModel
    if err = cursor.All(ctx, &modelos); err != nil {
        return nil, fmt.Errorf("error al decodificar sesiones: %w", err)
    }

    // Convertir modelos a agregados de dominio
    gameSessions := make([]*aggregates.GameSession, len(modelos))
    for i, model := range modelos {
        gameSessions[i] = models.ToDomainGameSession(&model)
    }

    return gameSessions, nil
}

func (r *GameSessionRepositoryMongoDB) GetByUserId(userID string) ([]*aggregates.GameSession, error) {
    collection := r.DB.Database.Collection("game_sessions")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Filtro por usuario
    filter := bson.M{"user_id": userID}

    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        return nil, fmt.Errorf("error al obtener sesiones del usuario: %w", err)
    }
    defer cursor.Close(ctx)

    var modelos []models.GameSessionModel
    if err = cursor.All(ctx, &modelos); err != nil {
        return nil, fmt.Errorf("error al decodificar sesiones: %w", err)
    }

    // Convertir modelos a agregados de dominio
    gameSessions := make([]*aggregates.GameSession, len(modelos))
    for i, model := range modelos {
        gameSessions[i] = models.ToDomainGameSession(&model)
    }

    return gameSessions, nil
}