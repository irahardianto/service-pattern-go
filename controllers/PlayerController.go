package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/irahardianto/service-pattern-go/interfaces"

	"github.com/go-chi/chi"
	"github.com/irahardianto/service-pattern-go/viewmodels"
)

type PlayerController struct {
	interfaces.IPlayerService
}

func (controller *PlayerController) GetPlayerScore(res http.ResponseWriter, req *http.Request) {

	player1Name := chi.URLParam(req, "player1")
	player2Name := chi.URLParam(req, "player2")

	scores, err := controller.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	json.NewEncoder(res).Encode(viewmodels.ScoresVM{scores})
}
