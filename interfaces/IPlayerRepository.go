package interfaces

import (
	"service-pattern-go/models"
)

type IPlayerRepository interface {
	GetAllPlayers() []models.Player
	GetPlayerById(id int) models.Player
	CreatePlayer(player models.Player) (bool, error)
	UpdatePlayer(id int, player models.Player) (bool, error)
	DeletePlayer(id int) (bool, error)
}
