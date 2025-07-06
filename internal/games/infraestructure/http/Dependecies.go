package http

import (
	"github.com/tLALOck64/microservicio-juegos/internal/games/application"
	"github.com/tLALOck64/microservicio-juegos/internal/games/application/services"
	"github.com/tLALOck64/microservicio-juegos/internal/games/domain/ports"
	"github.com/tLALOck64/microservicio-juegos/internal/games/infraestructure/adapters"
	"github.com/tLALOck64/microservicio-juegos/internal/games/infraestructure/http/controllers"
	"github.com/tLALOck64/microservicio-juegos/internal/games/infraestructure/http/controllers/helper"
)

var (
	MiniGameRepository ports.MiniGameRepository
	UUIDGen services.UUIDGenerator
	GameSessionRepository ports.GameSessionRepository
)

func init(){
	var err error
	MiniGameRepository, err = adapters.NewMiniGameRepositoryMongoDB()

	if err != nil {
		panic("Error initializing MiniGameRepository: " + err.Error())
	}

	UUIDGen, err = helper.NewUUID()

	if err != nil {
		panic("Error initializing UUIDGen: "+ err.Error())
	}

	GameSessionRepository, err = adapters.NewGameSessionRepositoryMongoDB()

	if err != nil{
		panic("Error initializing GameSessionRepository")
	}
}


func SetUpCreate() *controllers.CreateMiniGameController{
	createUseCase := application.NewCreateMiniGameUseCase(MiniGameRepository)
	return controllers.NewCreateMiniGameControlle(createUseCase)
}

func SetUpGet() *controllers.GetMiniGameController{
	getUseCase := application.NewGetMiniGameUseCase(MiniGameRepository)
	return controllers.NewGetMiniGameController(getUseCase)
}

func SetUpGetByID() *controllers.GetMiniGameByIdController{
	getByIdUseCase := application.NewGetMiniGameByIdUseCase(MiniGameRepository)
	return controllers.NewGetMiniGameByIdController(getByIdUseCase)
}

func SetUpCreateGameSession() *controllers.CreateGameSessionController{
	createGameSessionController := application.NewCreateGameSessionUseCase(GameSessionRepository)
	return controllers.NewCreateGameSessionController(createGameSessionController)
}

func SetUpGetByIDGameSession() *controllers.GetGameSessionByIdController{
	getGameSessionByIdController := application.NewGetGameSessionByIdUseCase(GameSessionRepository)
	return controllers.NewGetGameSessionByIdController(getGameSessionByIdController)
}

func SetUpUpdateGameSession() *controllers.UpdateGameSessionController{
	updateGameSessionController := application.NewUpdateGameSessionUseCase(GameSessionRepository)
	return controllers.NewUpdateGameSessionController(updateGameSessionController)
}