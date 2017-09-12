package services

import (
	"testing"

	"github.com/irahardianto/service-pattern-go/interfaces/mocks"
	"github.com/irahardianto/service-pattern-go/models"

	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {

	playerRepository := new(mocks.IPlayerRepository)

	player := models.PlayerModel{}
	player.Id = 101
	player.Name = "Rafael"
	player.Score = 3

	playerRepository.On("GetPlayerById", 101).Return(player)

	playerService := PlayerService{}
	playerService.PlayerRepository = playerRepository

	expectedResult := models.PlayerModel{}
	expectedResult.Id = 101
	expectedResult.Name = "Rafael"
	expectedResult.Score = 3

	actualResult := playerService.FindById(101)

	assert.Equal(t, expectedResult, actualResult)
}
