package repositories

import (
	"kahoot-api/internal/models"

	"gorm.io/gorm"
)

type (
	QuestionRepository struct {
		GetDB func() *gorm.DB
	}

	ICreateQuestion struct {
		GameId      string
		Title       string
		Body        string
		OptionOne   string
		OptionTwo   string
		OptionThree string
		OptionFour  string
		Correct     string
	}

  QuestionInterface interface {
    Create(ICreateQuestion) (*models.Question, error)
    GetGameQuestions(gameId string) ([]models.Question, error)
  }
)

func NewQuestionRepository(database *gorm.DB) QuestionRepository {
	return QuestionRepository{
		GetDB: func() *gorm.DB {
			return database
		},
	}
}

func (r *QuestionRepository) Create(data ICreateQuestion) (*models.Question, error) {
  var question = models.Question {
    GameId: data.GameId,
    Title: data.Title,
    Body: data.Body,
    OptionOne: data.OptionOne,
    OptionTwo: data.OptionTwo,
    OptionThree: data.OptionThree,
    OptionFour: data.OptionFour,
    Correct: data.Correct,
  }

  result := r.GetDB().Create(&question)

	if result.Error != nil {
		return nil, result.Error
	}

  return &question, nil
}

func (r *QuestionRepository) GetGameQuestions(gameId string) ([]models.Question, error) {
  var condition = make(map[string]interface{})
  var questions []models.Question

  condition["game_id"] = gameId

  result := r.GetDB().Where(condition).Find(&questions)

	if result.Error != nil {
		return nil, result.Error
	}

  return questions, nil
}
