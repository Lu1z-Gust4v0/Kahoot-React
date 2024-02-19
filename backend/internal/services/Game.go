package services

import (
	"kahoot-api/internal/models"
	"kahoot-api/internal/repositories"
)

type (
	QuestionData struct {
		Title       string
		Body        string
		OptionOne   string
		OptionTwo   string
		OptionThree string
		OptionFour  string
		Correct     string
	}

	GameService struct {
		GameRepo     repositories.GameRepository
		PlayerRepo   repositories.PlayerRepository
		QuestionRepo repositories.QuestionRepository
	}

	GameServiceInterface interface {
		CreateNewGame(data repositories.ICreateGame) (*models.Game, error)
		UpdateGameStatus(gameId string, status models.GameStatus) (*models.Game, []models.Question, error)
		AddNewPlayer(gameId string, name string) (*models.Player, error)
		IncreasePlayerScore(playerId string, score uint16) (*models.Player, error)
	}
)

func (service *GameService) CreateNewGame(gameData repositories.ICreateGame, questions []QuestionData) (*models.Game, []models.Question, error) {
	newGame, createError := service.GameRepo.Create(gameData)

	newQuestions := make([]models.Question, len(questions))

	if createError != nil {
		return nil, nil, createError
	}

	for _, question := range questions {
		newQuestion, createError := service.QuestionRepo.Create(repositories.ICreateQuestion{
			GameId:      newGame.Id,
			Title:       question.Title,
			Body:        question.Body,
			OptionOne:   question.OptionOne,
			OptionTwo:   question.OptionTwo,
			OptionThree: question.OptionThree,
			OptionFour:  question.OptionFour,
			Correct:     question.Correct,
		})

		if createError != nil {
			return nil, nil, createError
		}

		newQuestions = append(newQuestions, *newQuestion)
	}

	return newGame, newQuestions, nil
}

func (service *GameService) UpdateGameStatus(gameId string, status models.GameStatus) (*models.Game, error) {
	game, updateError := service.GameRepo.Update(repositories.IUpdateGame{Id: gameId, Status: status})

	if updateError != nil {
		return nil, updateError
	}

	return game, nil
}

func (service *GameService) AddNewPlayer(gameId string, name string) (*models.Player, error) {
	player, createError := service.PlayerRepo.Create(repositories.ICreatePlayer{
		GameId: gameId,
		Name:   name,
	})

	if createError != nil {
		return nil, createError
	}

	return player, nil
}

func (service *GameService) IncreasePlayerScore(playerId string, score uint16) (*models.Player, error) {
	player, getError := service.PlayerRepo.GetById(playerId)

	if getError != nil {
		return nil, getError
	}

	player, updateError := service.PlayerRepo.UpdatePlayerScore(playerId, player.Score+score)

	if updateError != nil {
		return nil, updateError
	}

	return player, nil
}
