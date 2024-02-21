package websocket

import (
	"errors"
	"kahoot-api/internal/models"
	"log"
)

func (gameHub *GameHub) HandleConnection(request *Register) (*models.Player, error) {
	if gameHub.Game.Status != models.Waiting {
		log.Println("Game cannot receive new players")
    return nil, errors.New("Game cannot receive new players")
	}

	if gameHub.Game.MaxPlayers == uint8(len(gameHub.Clients)) {
		log.Println("Game is already full")
    return nil, errors.New("Game is already full")
	}

	player, createError := gameHub.GameService.AddNewPlayer(gameHub.Game.Id, request.Name)

	if createError != nil {
		log.Println("Failed to create player")
    return nil, errors.New("Failed to create player")
	}

	gameHub.Clients[player.Id] = &Client{Player: player, Connection: request.Connection}
  go gameHub.BroadCastGameState()

  return player, nil
}

func (gameHub *GameHub) HandleDisconnect(playerId string) {
  _, valid := gameHub.Clients[playerId]

  // If the connection was invalid, it means the connection was not from a registed player
  if !valid {
    return
  }

	log.Printf("Player %s left the game\n", playerId)
  delete(gameHub.Clients, playerId)
  go gameHub.BroadCastGameState()
}

func (gameHub *GameHub) HandleAnswer(request *Answer) {
  _, closed := gameHub.Clients[request.PlayerId]

	if !closed {
		log.Println("Invalid connection")
		return
	}
  
  var correct = gameHub.Questions[gameHub.CurrentQuestion].Correct

	if request.Answer == correct {
		gameHub.Clients[request.PlayerId].Player.Score += CORRECT_ANSWER_POINTS
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
