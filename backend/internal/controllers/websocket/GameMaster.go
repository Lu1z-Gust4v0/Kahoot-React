package websocket

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"kahoot-api/internal/dtos"
	"kahoot-api/internal/models"
	"kahoot-api/internal/services"
	ws "kahoot-api/internal/services/websocket"
)

func GameMasterWebsocket(service *services.GameService) func(*fiber.Ctx) error {
	return websocket.New(func(connection *websocket.Conn) {
		gameId := connection.Params("gameId")

		game, error := service.GetGameById(gameId)

		if error != nil {
			connection.WriteJSON(dtos.ErrorMessage{
				Type:  dtos.ERROR,
				Error: error.Error(),
			})
      connection.Close()
      return
		}

		gameHub, error := GetGameHub(game, connection, service)

		if error != nil {
			connection.WriteJSON(dtos.ErrorMessage{
				Type:  dtos.ERROR,
				Error: error.Error(),
			})
      connection.Close()
      return
		}

		go gameHub.BroadCastGameState()

		defer func() {
			gameHub.GameMaster.Closed = true
			gameHub.GameMaster.Connection.Close()
		}()

		for {
			var incommingMessage map[string]interface{}

			_, data, readError := connection.ReadMessage()

			if readError != nil {
				log.Println("Error on reading incomming message: ", readError)
				if websocket.IsUnexpectedCloseError(readError, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					return
				}
				return
			}

			parseError := json.Unmarshal(data, &incommingMessage)

			if parseError != nil {
				connection.WriteJSON(dtos.ErrorMessage{
					Type:  dtos.ERROR,
					Error: "Failed to parsed incomming message",
				})
				return
			}

			HandleGameMasterMessages(gameHub, connection, incommingMessage, data)
		}
	})
}

func GetGameHub(game *models.Game, connection *websocket.Conn, service *services.GameService) (*ws.GameHub, error) {
	_, validCode := ActiveGames[game.Code]

	if validCode || game.Status != models.Waiting {
		return nil, errors.New("Game hub already exists")
	}

	questions, error := service.GetGameQuestions(game.Id)

	if error != nil {
		return nil, error
	}

	gameHub := ws.NewGameHub(connection, game, questions, service)
	ActiveGames[game.Code] = gameHub

	return gameHub, nil
}

func HandleGameMasterMessages(gameHub *ws.GameHub, connection *websocket.Conn, message map[string]interface{}, data []byte) {
	switch messageType := message["type"].(dtos.MessageType); messageType {
	case dtos.START_GAME:
		var startGame dtos.StartGameRequest

		parseError := json.Unmarshal(data, &startGame)

		if parseError != nil {
			connection.WriteJSON(dtos.ErrorMessage{
				Type:  dtos.ERROR,
				Error: "Invalid start game request",
			})
			return
		}

		gameHub.GameEventChannel <- ws.START_GAME

	default:
		connection.WriteJSON(dtos.ErrorMessage{
			Type:  dtos.ERROR,
			Error: "Message type not supported",
		})
	}
}
