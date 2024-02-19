package websocket

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"kahoot-api/internal/dtos"
	ws "kahoot-api/internal/services/websocket"
)

var ActiveGames = make(map[string]*ws.GameHub)

func GameClientWebsocket() func(*fiber.Ctx) error {
	return websocket.New(func(connection *websocket.Conn) {
		go func() {
			connection.Close()
		}()

		gameCode := connection.Params("code")

		gameHub, validCode := ActiveGames[gameCode]

		if !validCode {
			connection.WriteJSON(dtos.ErrorMessage{
				Type:  dtos.ERROR,
				Error: "Invalid game code",
			})
			connection.Close()
		}

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

			HandleGameClientMessages(gameHub, connection, incommingMessage, data)
		}
	})
}

func HandleGameClientMessages(gameHub *ws.GameHub, connection *websocket.Conn, message map[string]interface{}, data []byte) {
	switch messageType := message["type"].(dtos.MessageType); messageType {
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

		gameHub.RegisterChannel <- &ws.Register{Name: register.Name, Connection: connection}
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

		gameHub.AnswerChannel <- &ws.Answer{Answer: answer.Answer}

	default:
		connection.WriteJSON(dtos.ErrorMessage{
			Type:  dtos.ERROR,
			Error: "Message type not supported",
		})
	}
}
