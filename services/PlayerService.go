package services

import (
	"service-pattern-go/interfaces"
	"service-pattern-go/models"
)

type PlayerService struct {
	PlayerRepository interfaces.IPlayerRepository
}

func (repository *PlayerService) FindById(playerId int) models.PlayerModel {

	player := repository.PlayerRepository.GetPlayerById(playerId)

	return player
}
