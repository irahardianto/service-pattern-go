package helpers

import (
	"github.com/irahardianto/service-pattern-go/viewmodels"
)

type PlayerHelper struct{}

func (helper *PlayerHelper) BuildScoresVM(scores string) viewmodels.ScoresVM {

	scoresVM := viewmodels.ScoresVM{}
	scoresVM.Score = scores

	return scoresVM
}
