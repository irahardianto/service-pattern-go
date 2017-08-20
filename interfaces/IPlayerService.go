package interfaces

import (
	"service-pattern-go/models"
)

type IPlayerService interface {
	FindById(playerId int) models.PlayerModel
	GetPlayerMessage() models.MessageModel
}
