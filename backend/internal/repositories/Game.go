package repositories

import (
	"kahoot-api/internal/models"

	"gorm.io/gorm"
)

type (
	GameRepository struct {
		GetDB func() *gorm.DB
	}

	ICreateGame struct {
		Title      string
		Code       string
		MaxPlayers uint8
	}

	IUpdateGame struct {
		Id     string
		Status models.GameStatus
	}

	GameInterface interface {
		Create(ICreateGame) (*models.Game, error)
		GetById(id string) (*models.Game, error)
		Update(IUpdateGame) (*models.Game, error)
	}
)

func NewGameRepository(database *gorm.DB) GameRepository {
	return GameRepository{
		GetDB: func() *gorm.DB {
			return database
		},
	}
}

func (r *GameRepository) Create(data ICreateGame) (*models.Game, error) {
	var game = models.Game{
		Title:      data.Title,
		Code:       data.Code,
		MaxPlayers: data.MaxPlayers,
		Status:     models.Waiting,
	}

	result := r.GetDB().Create(&game)

	if result.Error != nil {
		return nil, result.Error
	}

	return &game, nil
}

func (r *GameRepository) GetById(id string) (*models.Game, error) {
	var game = models.Game{}

	result := r.GetDB().First(&game, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &game, nil
}

func (r *GameRepository) Update(data IUpdateGame) (*models.Game, error) {
	var fieldMap = make(map[string]interface{})
	var game = models.Game{Id: data.Id}

	fieldMap["status"] = data.Status

	result := r.GetDB().Model(&game).Updates(fieldMap)

	if result.Error != nil {
		return nil, result.Error
	}

	return &game, nil
}
