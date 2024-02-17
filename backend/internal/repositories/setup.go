package repositories

import "gorm.io/gorm"

type Repository struct {
	DB       *gorm.DB
	Game     GameRepository
	Player   PlayerRepository
	Question QuestionRepository
}

func SetupRepositories(connection *gorm.DB) *Repository {
	return &Repository{
		DB:       connection,
		Game:     NewGameRepository(connection),
		Player:   NewPlayerRepository(connection),
		Question: NewQuestionRepository(connection),
	}
}
