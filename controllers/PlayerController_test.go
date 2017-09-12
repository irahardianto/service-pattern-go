package controllers

import (
	"encoding/json"
	"net/http/httptest"

	"github.com/irahardianto/service-pattern-go/interfaces/mocks"
	"github.com/irahardianto/service-pattern-go/models"
	"github.com/irahardianto/service-pattern-go/viewmodels"

	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestPlayerScore(t *testing.T) {

	player := models.PlayerModel{}
	player.Id = 101
	player.Name = "Rafael"
	player.Score = 3

	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	// setup expectations
	playerService.On("GetScores", "Rafael", "Serena").Return("Forty-Fifteen", nil)

	playerController := PlayerController{}
	playerController.PlayerService = playerService

	// call the code we are testing
	req := httptest.NewRequest("GET", "http://localhost:8080/getScore/Rafael/vs/Serena", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/getScore/{player1}/vs/{player2}", playerController.GetPlayerScore)

	r.ServeHTTP(w, req)

	expectedResult := viewmodels.ScoresVM{}
	expectedResult.Score = "Forty-Fifteen"

	actualResult := viewmodels.ScoresVM{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
}
