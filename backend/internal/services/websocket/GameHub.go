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
		Closed bool
	}

	Register struct {
		Name       string
		Connection *websocket.Conn
	}

	Answer struct {
		Answer     string
		Connection *websocket.Conn
	}

	GameHub struct {
		Game              *models.Game
		Questions         []models.Question
		CurrentQuestion   uint
		GameService       *services.GameService
		RegisterChannel   chan *Register
		UnregisterChannel chan *websocket.Conn
		AnswerChannel     chan *Answer
		GameMaster        *GameMaster
		Clients           map[*websocket.Conn]*Client
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
		RegisterChannel:   make(chan *Register),
		UnregisterChannel: make(chan *websocket.Conn),
		AnswerChannel:     make(chan *Answer),
		GameMaster:        &GameMaster{Connection: gameMaster, Closed: false},
		Clients:           make(map[*websocket.Conn]*Client),
		GameEventChannel:  make(chan GameEvent, 2),
		Done:             make(chan bool),
	}
}

func RunGameHub(gameHub *GameHub) {
	for {
		select {
		case request := <-gameHub.RegisterChannel:
			gameHub.HandleConnection(request)

		case unregister := <-gameHub.UnregisterChannel:
			gameHub.HandleDisconnect(unregister)

		case answer := <-gameHub.AnswerChannel:
			gameHub.HandleAnswer(answer)

		case gameEvent := <-gameHub.GameEventChannel:
			gameHub.HandleGameEvent(gameEvent)

    case <-gameHub.Done:
      return
		}
	}
}
