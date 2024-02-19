package websocket

import (
	"kahoot-api/internal/dtos"
	"kahoot-api/internal/models"
	"log"

	"github.com/gofiber/contrib/websocket"
)

func (gameHub *GameHub) HandleConnection(request *Register) {
	if gameHub.Game.Status != models.Waiting {
		log.Println("Game cannot receive new players")
		request.Connection.Close()
	}

	if gameHub.Game.MaxPlayers == uint8(len(gameHub.Clients)) {
		log.Println("Game is already full")
		request.Connection.Close()
	}

	player, createError := gameHub.GameService.AddNewPlayer(gameHub.Game.Id, request.Name)

	if createError != nil {
		log.Println("Failed to create player")
		log.Println("Closing connection...")
		request.Connection.Close()
	}

	gameHub.Clients[request.Connection] = &Client{Player: player, Closed: false}
}

func (gameHub *GameHub) HandleDisconnect(connection *websocket.Conn) {
	gameHub.Clients[connection].Closed = true
	log.Printf("Player %s left the game\n", gameHub.Clients[connection].Player.Id)
}

func (gameHub *GameHub) HandleAnswer(request *Answer) {
	client := gameHub.Clients[request.Connection]

	if client.Closed {
		log.Println("This client is already closed")
		return
	}

	if gameHub.Questions[gameHub.CurrentQuestion].Correct == request.Answer {
		gameHub.Clients[request.Connection].Player.Score += CORRECT_ANSWER_POINTS
	}
}

func (gameHub *GameHub) HandleGameEvent(event GameEvent) {
	switch event {
	case START_GAME:
		gameHub.Game.Status = models.Started

		log.Println("Game started successfully")
		gameHub.BroadCastQuestion()
	case NEXT_QUESTION:
		gameHub.CurrentQuestion += 1
		gameHub.BroadCastQuestion()

	case SHOW_SCORES:
		gameHub.BroadCastScores()

	case FINISH_GAME:
		gameHub.Game.Status = models.Finished
		log.Println("Game finished successfully")
	default:
		log.Println("Unknown game event")
	}
}

func (gameHub *GameHub) BroadCastQuestion() {
	for connection, client := range gameHub.Clients {
		go func(connection *websocket.Conn, client *Client) {
			if client.Closed {
				return
			}

			currentQuestion := gameHub.Questions[gameHub.CurrentQuestion]

			broadcastError := connection.WriteJSON(dtos.QuestionMessage{
				Type:        dtos.QUESTION,
				GameId:      gameHub.Game.Id,
				Title:       currentQuestion.Title,
				Body:        currentQuestion.Body,
				OptionOne:   currentQuestion.OptionOne,
				OptionTwo:   currentQuestion.OptionTwo,
				OptionThree: currentQuestion.OptionThree,
				OptionFour:  currentQuestion.OptionFour,
				Correct:     currentQuestion.Correct,
			})

			if broadcastError != nil {
				log.Printf("Failed to broadcast message to %s\n", client.Player.Id)

				connection.Close()
				gameHub.UnregisterChannel <- connection
			}
		}(connection, client)
	}
}

func (gameHub *GameHub) GetPlayers() []models.Player {
	var players []models.Player

	for _, client := range gameHub.Clients {
		players = append(players, *client.Player)
	}

	return players
}

func (gameHub *GameHub) BroadCastScores() {
	players := gameHub.GetPlayers()

	for connection, client := range gameHub.Clients {
		go func(connection *websocket.Conn, client *Client) {
			if client.Closed {
				return
			}

			broadcastError := connection.WriteJSON(dtos.ScoresMessage{
				Type:    dtos.SCORES,
				GameId:  gameHub.Game.Id,
				Players: players,
			})

			if broadcastError != nil {
				log.Printf("Failed to broadcast message to %s\n", client.Player.Id)

				connection.Close()
				gameHub.UnregisterChannel <- connection
			}
		}(connection, client)
	}
}
