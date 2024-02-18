package main

import (
	"github.com/gofiber/fiber/v2"
	"kahoot-api/configs"
	"kahoot-api/internal/controllers"
	repo "kahoot-api/internal/repositories"
	"kahoot-api/internal/services"
	"log"
)

func main() {
	app := fiber.New()

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

	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Hello World")
	})

	app.Listen(":3000")
}
