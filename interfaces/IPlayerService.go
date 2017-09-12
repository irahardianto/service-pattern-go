package interfaces

import (
	"github.com/irahardianto/service-pattern-go/models"
)

type IPlayerService interface {
	GetScores(player1Name string, player2Name string) (string, error)
	GetPlayerMessage() models.MessageModel
}
