package helpers

import (
	"service-pattern-go/models"
	"service-pattern-go/viewmodels"
)

type PlayerHelper struct{}

func (helper *PlayerHelper) BuildPlayerVM(model models.PlayerModel) viewmodels.PlayerVM {

	playerVM := viewmodels.PlayerVM{}
	playerVM.Name = model.Name
	playerVM.Score = model.Score

	return playerVM
}
