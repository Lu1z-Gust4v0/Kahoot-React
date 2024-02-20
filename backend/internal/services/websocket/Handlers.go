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
  go gameHub.BroadCastGameState()
}

func (gameHub *GameHub) HandleDisconnect(connection *websocket.Conn) {
	gameHub.Clients[connection].Closed = true
	log.Printf("Player %s left the game\n", gameHub.Clients[connection].Player.Id)
  go gameHub.BroadCastGameState()
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
    gameHub.BroadCastGameState()
		gameHub.GameEventChannel <- NEXT_QUESTION

	case NEXT_QUESTION:
    log.Println("Next question")
		gameHub.BroadCastQuestion()
    gameHub.BroadCastGameState()
    gameHub.CurrentQuestion++;
    gameHub.GameEventChannel <- SHOW_SCORES

	case SHOW_SCORES:
    log.Println("Show players score")
		gameHub.BroadCastScores()

    if gameHub.CurrentQuestion == uint(len(gameHub.Questions)) {
      gameHub.GameEventChannel <- FINISH_GAME
    }

    gameHub.BroadCastGameState()
    gameHub.GameEventChannel <- NEXT_QUESTION

	case FINISH_GAME:
    log.Println("Game finished successfully")
		gameHub.Game.Status = models.Finished
    gameHub.BroadCastGameState()
    gameHub.Done <- true

	default:
		log.Println("Unknown game event")
	}
}
