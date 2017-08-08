package interfaces

import (
	"service-pattern-go/models"
)

type IPlayerRepository interface {
	GetAllPlayers() []models.PlayerModel
	GetPlayerById(id int) models.PlayerModel
	CreatePlayer(player models.PlayerModel) (bool, error)
	UpdatePlayer(id int, player models.PlayerModel) (bool, error)
	DeletePlayer(id int) (bool, error)
}
