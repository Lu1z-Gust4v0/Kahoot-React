package websocket

import (
	"kahoot-api/internal/models"
	"kahoot-api/internal/services"

	"github.com/gofiber/contrib/websocket"
)

type GameEvent = uint8

const (
	START_GAME GameEvent = iota
	NEXT_QUESTION
	SHOW_SCORES
	FINISH_GAME
)

const CORRECT_ANSWER_POINTS = 1000

type (
	GameMaster struct {
		Connection *websocket.Conn
		Closed     bool
	}

	Client struct {
		Player *models.Player
		Connection *websocket.Conn
	}

	Register struct {
		Name       string
		Connection *websocket.Conn
	}

	Answer struct {
		Answer     string
		PlayerId   string
	}

	GameHub struct {
		Game              *models.Game
		Questions         []models.Question
		CurrentQuestion   uint
		GameService       *services.GameService
		UnregisterChannel chan string
		GameMaster        *GameMaster
		Clients           map[string]*Client
		GameEventChannel  chan GameEvent
    Done             chan bool
	}
)

func NewGameHub(gameMaster *websocket.Conn, game *models.Game, questions []models.Question, service *services.GameService) *GameHub {
	return &GameHub{
		Game:              game,
		Questions:         questions,
		CurrentQuestion:   0,
		GameService:       service,
		UnregisterChannel: make(chan string),
		GameMaster:        &GameMaster{Connection: gameMaster, Closed: false},
		Clients:           make(map[string]*Client),
		GameEventChannel:  make(chan GameEvent, 2),
		Done:             make(chan bool),
	}
}

func RunGameHub(gameHub *GameHub) {
	for {
		select {
		case unregister := <-gameHub.UnregisterChannel:
			gameHub.HandleDisconnect(unregister)

		case gameEvent := <-gameHub.GameEventChannel:
			gameHub.HandleGameEvent(gameEvent)

    case <-gameHub.Done:
      return
		}
	}
}
