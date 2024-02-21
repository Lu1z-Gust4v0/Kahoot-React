package websocket

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"kahoot-api/internal/dtos"
	"kahoot-api/internal/models"
	ws "kahoot-api/internal/services/websocket"
)

var ActiveGames = make(map[string]*ws.GameHub)

func GameClientWebsocket() func(*fiber.Ctx) error {
	return websocket.New(func(connection *websocket.Conn) {
		var player models.Player

		gameCode := connection.Params("code")

		gameHub, validCode := ActiveGames[gameCode]

		if !validCode {
			connection.WriteJSON(dtos.ErrorMessage{
				Type:  dtos.ERROR,
				Error: "Invalid game code",
			})
			connection.Close()
		}

		defer func() {
			connection.Close()
      if player.Id != "" {
        gameHub.UnregisterChannel <- player.Id
      }
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

			HandleGameClientMessages(gameHub, connection, &player, incommingMessage, data)
		}
	})
}

func HandleGameClientMessages(gameHub *ws.GameHub, connection *websocket.Conn, player *models.Player, message map[string]interface{}, data []byte) {
	switch messageType := dtos.MessageType(message["type"].(float64)); messageType {
	case dtos.REGISTER:
		var register dtos.RegisterRequest

		parseError := json.Unmarshal(data, &register)

		if parseError != nil {
			connection.WriteJSON(dtos.ErrorMessage{
				Type:  dtos.ERROR,
				Error: "Invalid register request",
			})
			return
		}

    newPlayer, connectionError := gameHub.HandleConnection(&ws.Register{Name: register.Name, Connection: connection})

    if connectionError != nil {
			connection.WriteJSON(dtos.ErrorMessage{
				Type:  dtos.ERROR,
				Error: connectionError.Error(),
			})
			return
    }
    *player = *newPlayer

	case dtos.ANSWER:
		var answer dtos.AnswerRequest

		parseError := json.Unmarshal(data, &answer)

		if parseError != nil {
			connection.WriteJSON(dtos.ErrorMessage{
				Type:  dtos.ERROR,
				Error: "Invalid answer request",
			})
			return
		}
  
    if player == nil {
			connection.WriteJSON(dtos.ErrorMessage{
				Type:  dtos.ERROR,
				Error: "Player not registered",
			})
			return
    }

		gameHub.HandleAnswer(&ws.Answer{PlayerId: player.Id, Answer: answer.Answer})

	default:
		connection.WriteJSON(dtos.ErrorMessage{
			Type:  dtos.ERROR,
			Error: "Message type not supported",
		})
	}
}
