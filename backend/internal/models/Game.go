package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameStatus int8

const (
	Waiting GameStatus = iota
	Started
	Finished
)

type Question struct {
	Id          string `json:"id" gorm:"primaryKey"`
	GameId      string `json:"game_id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	OptionOne   string `json:"option_one"`
	OptionTwo   string `json:"option_two"`
	OptionThree string `json:"option_three"`
	OptionFour  string `json:"option_four"`
	Correct     string `json:"correct"`
}

type Player struct {
	Id     string `json:"id" gorm:"primaryKey"`
	GameId string `json:"game_id"`
	Name   string `json:"name"`
	Score uint16 `json:"score"`
}

type Game struct {
	Id         string     `json:"id" gorm:"primaryKey"`
	Title      string     `json:"title"`
	Code       string     `json:"code"`
	MaxPlayers uint8      `json:"max_players"`
	Status     uint8      `json:"status"`
}

func (player *Player) BeforeCreate(tx *gorm.DB) error {
	player.Id = uuid.NewString()

	return nil
}

func (quiz *Game) BeforeCreate(tx *gorm.DB) error {
	quiz.Id = uuid.NewString()

	return nil
}

func (question *Question) BeforeCreate(tx *gorm.DB) error {
	question.Id = uuid.NewString()

	return nil
}
