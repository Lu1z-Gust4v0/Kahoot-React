package dtos

import "kahoot-api/internal/models"

type MessageType = uint8

const (
	SCORES MessageType = iota
	QUESTION
	REGISTER
	ANSWER
)

type (
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

	RegisterRequest struct {
		Type MessageType `json:"type"`
		Name string      `json:"name"`
	}

	AnswerRequest struct {
		Type   MessageType `json:"type"`
		Answer string      `json:"answer"`
	}
)
