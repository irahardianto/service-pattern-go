package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/irahardianto/service-pattern-go/helpers"
	"github.com/irahardianto/service-pattern-go/interfaces"

	"github.com/go-chi/chi"
)

type PlayerController struct {
	PlayerService interfaces.IPlayerService
	PlayerHelper  helpers.PlayerHelper
}

func (controller *PlayerController) GetPlayer(res http.ResponseWriter, req *http.Request) {

	playerId, _ := strconv.Atoi(chi.URLParam(req, "id"))
	player := controller.PlayerService.FindById(playerId)
	playerVM := controller.PlayerHelper.BuildPlayerVM(player)

	json.NewEncoder(res).Encode(playerVM)
}

func (controller *PlayerController) GetPlayerMessage(res http.ResponseWriter, req *http.Request) {

	data := controller.PlayerService.GetPlayerMessage()
	json.NewEncoder(res).Encode(data)
}
