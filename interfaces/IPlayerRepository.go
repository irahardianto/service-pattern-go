package interfaces

import (
	"github.com/irahardianto/service-pattern-go/models"
)

type IPlayerRepository interface {
	GetPlayerByName(name string) (models.PlayerModel, error)
}
