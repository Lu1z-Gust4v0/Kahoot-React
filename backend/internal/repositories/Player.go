package repositories

import (
	"kahoot-api/internal/models"

	"gorm.io/gorm"
)

type (
	PlayerRepository struct {
		GetDB func() *gorm.DB
	}

	ICreatePlayer struct {
		GameId string
		Name   string
	}

	PlayerInterface interface {
		Create(ICreatePlayer) (*models.Player, error)
		GetById(id string) (*models.Player, error)
		UpdatePlayerScore(Id string, score uint16) (*models.Player, error)
	}
)

func NewPlayerRepository(database *gorm.DB) PlayerRepository {
	return PlayerRepository{
		GetDB: func() *gorm.DB {
			return database
		},
	}
}

func (r *PlayerRepository) Create(data ICreatePlayer) (*models.Player, error) {
	var player = models.Player{
		GameId: data.GameId,
		Name:   data.Name,
		Score:  uint16(0),
	}

	result := r.GetDB().Create(&player)

	if result.Error != nil {
		return nil, result.Error
	}

	return &player, nil
}

func (r *PlayerRepository) GetById(id string) (*models.Player, error) {
	var player = models.Player{}

	result := r.GetDB().First(&player, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &player, nil
}

func (r *PlayerRepository) UpdatePlayerScore(Id string, score uint16) (*models.Player, error) {
	var fieldMap = make(map[string]interface{})
	var player = models.Player{Id: Id}

	fieldMap["score"] = score

	result := r.GetDB().Model(&player).Updates(fieldMap)

	if result.Error != nil {
		return nil, result.Error
	}

	return &player, nil
}
