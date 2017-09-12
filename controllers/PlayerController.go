package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/irahardianto/service-pattern-go/helpers"
	"github.com/irahardianto/service-pattern-go/interfaces"

	"github.com/go-chi/chi"
)

type PlayerController struct {
	PlayerService interfaces.IPlayerService
	PlayerHelper  helpers.PlayerHelper
}

func (controller *PlayerController) GetPlayerScore(res http.ResponseWriter, req *http.Request) {

	player1Name := chi.URLParam(req, "player1")
	player2Name := chi.URLParam(req, "player2")

	scores, err := controller.PlayerService.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	response := controller.PlayerHelper.BuildScoresVM(scores)

	json.NewEncoder(res).Encode(response)
}

func (controller *PlayerController) GetPlayerMessage(res http.ResponseWriter, req *http.Request) {

	data := controller.PlayerService.GetPlayerMessage()
	json.NewEncoder(res).Encode(data)
}
