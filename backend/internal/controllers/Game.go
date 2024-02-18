package controllers

import (
	"kahoot-api/internal/dtos"
	"kahoot-api/internal/repositories"
	"kahoot-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type BaseController struct {
	Service *services.GameService
}

func NewBaseController(
	service *services.GameService,
) *BaseController {
	return &BaseController{
		Service: service,
	}
}

func (handler *BaseController) CreateGame(context *fiber.Ctx) error {
	body := new(dtos.CreateGameRequest)

	if parseError := context.BodyParser(body); parseError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   parseError.Error(),
		})
	}

	newGame, serviceError := handler.Service.CreateNewGame(repositories.ICreateGame{
		Title:      body.Title,
		Code:       body.Code,
		MaxPlayers: body.MaxPlayers,
	}, body.Questions)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to create error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "game created successfully",
		"game":    newGame,
	})
}
