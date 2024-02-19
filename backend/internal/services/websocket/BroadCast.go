package websocket

import (
	"kahoot-api/internal/dtos"
	"kahoot-api/internal/models"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
)

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
	gameHub.CurrentQuestion++
	// sleep for 10 seconds before showing the scores
	time.Sleep(10 * time.Second)
	gameHub.GameEventChannel <- SHOW_SCORES

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
	// We are at the last question
	if gameHub.CurrentQuestion == uint(len(gameHub.Questions)) {
		gameHub.GameEventChannel <- FINISH_GAME
		return
	}

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

	// sleep for 10 seconds before going to the next question
	time.Sleep(10 * time.Second)
	gameHub.GameEventChannel <- NEXT_QUESTION
}

func (gameHub *GameHub) GetActivePlayers() uint8 {
	var players uint8 = 0

	for _, client := range gameHub.Clients {
		if !client.Closed {
			players++
		}
	}

	return players
}

func (gameHub *GameHub) BroadCastGameState() {
	if gameHub.GameMaster.Closed {
		return
	}

	broadcastError := gameHub.GameMaster.Connection.WriteJSON(dtos.GameStateMessage{
		Type:            dtos.GAME_STATE,
		Title:           gameHub.Game.Title,
		Code:            gameHub.Game.Code,
		CurrentQuestion: gameHub.CurrentQuestion,
		QuestionCount:   uint(len(gameHub.Questions)),
		MaxPlayers:      gameHub.Game.MaxPlayers,
		ActivePlayers:   gameHub.GetActivePlayers(),
	})

	if broadcastError != nil {
		log.Println("Failed to broadcast message to game master")

		gameHub.GameMaster.Connection.Close()
	}
}
