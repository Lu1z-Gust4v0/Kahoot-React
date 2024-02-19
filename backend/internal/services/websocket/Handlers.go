package websocket

import (
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
    log.Println("Game started successfully")
		gameHub.Game.Status = models.Started
		gameHub.GameEventChannel <- NEXT_QUESTION

	case NEXT_QUESTION:
		go gameHub.BroadCastQuestion()

	case SHOW_SCORES:
		go gameHub.BroadCastScores()

	case FINISH_GAME:
    log.Println("Game finished successfully")
		gameHub.Game.Status = models.Finished

	default:
		log.Println("Unknown game event")
	}
}
