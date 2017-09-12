package services

import (
	"github.com/irahardianto/service-pattern-go/interfaces"
	"github.com/irahardianto/service-pattern-go/models"
)

type PlayerService struct {
	PlayerRepository interfaces.IPlayerRepository
}

func (service *PlayerService) FindById(playerId int) models.PlayerModel {

	player := service.PlayerRepository.GetPlayerById(playerId)

	return player
}

func (service *PlayerService) GetPlayerMessage() models.MessageModel {

	data := service.PlayerRepository.GetPlayerMessageFromAPI()

	return data
}
