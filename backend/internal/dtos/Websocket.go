package dtos

import "kahoot-api/internal/models"

type MessageType = uint8

const (
	SCORES MessageType = iota
	QUESTION
	REGISTER
	ANSWER
	START_GAME
	GAME_STATE
	ERROR
)

type (
	GameStateMessage struct {
		Type            MessageType       `json:"type"`
		Title           string            `json:"title"`
		Code            string            `json:"code"`
		CurrentQuestion uint              `json:"current_question"`
		QuestionCount   uint              `json:"question_count"`
		MaxPlayers      uint8             `json:"max_players"`
		ActivePlayers   uint8             `json:"active_players"`
		Players         []models.Player   `json:"players"`
		Status          models.GameStatus `json:"status"`
	}

	QuestionMessage struct {
		Type        MessageType `json:"type"`
		GameId      string      `json:"game_id"`
		Title       string      `json:"title"`
		Body        string      `json:"body"`
		OptionOne   string      `json:"option_one"`
		OptionTwo   string      `json:"option_two"`
		OptionThree string      `json:"option_three"`
		OptionFour  string      `json:"option_four"`
		Correct     string      `json:"correct"`
	}

	ScoresMessage struct {
		Type    MessageType     `json:"type"`
		GameId  string          `json:"game_id"`
		Players []models.Player `json:"players"`
	}

	ErrorMessage struct {
		Type  MessageType `json:"type"`
		Error string      `json:"error"`
	}

	RegisterRequest struct {
		Type MessageType `json:"type"`
		Name string      `json:"name"`
	}

	AnswerRequest struct {
		Type   MessageType `json:"type"`
		Answer string      `json:"answer"`
	}

	StartGameRequest struct {
		Type MessageType `json:"type"`
	}
)
