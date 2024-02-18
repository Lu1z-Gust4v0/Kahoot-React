package dtos

import "kahoot-api/internal/services"

type (
	CreateGameRequest struct {
		Title      string                  `json:"title"`
		Code       string                  `json:"code"`
		MaxPlayers uint8                   `json:"max_players"`
		Questions  []services.QuestionData `json:"questions"`
	}
)
