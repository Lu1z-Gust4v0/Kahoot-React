package main

import (
	"kahoot-api/configs"
	"kahoot-api/internal/controllers"
	ws "kahoot-api/internal/controllers/websocket"
	repo "kahoot-api/internal/repositories"
	"kahoot-api/internal/services"

	"github.com/gofiber/fiber/v2"
  "github.com/gofiber/contrib/websocket"

	"log"
)

func main() {
	app := fiber.New()
 
  // Protect websocket routes
  app.Use("/api/ws", func(context *fiber.Ctx) error {
    if websocket.IsWebSocketUpgrade(context) {
      return context.Next()
    }
    
    return context.Status(fiber.StatusUpgradeRequired).JSON(fiber.Map{
      "message": "This is a websocker route",
    })
  })

	database, setupError := configs.SetUpDatabase()
	defer configs.CloseConnection(database)

	if setupError != nil {
		log.Fatal(setupError.Error())
	}

	repositories := repo.SetupRepositories(database)
	service := &services.GameService{
		GameRepo:     repositories.Game,
		QuestionRepo: repositories.Question,
		PlayerRepo:   repositories.Player,
	}

	controllers := controllers.NewBaseController(service)

	app.Post("/api/game", controllers.CreateGame)
  app.Get("/api/ws/client/:code", ws.GameClientWebsocket())
  app.Get("/api/ws/master/:gameId", ws.GameMasterWebsocket(service))

	app.Listen(":3000")
}
