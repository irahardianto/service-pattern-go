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
func TestGetScore(t *testing.T) {

	player := models.PlayerModel{}
	player.Id = 101
	player.Name = "Rafael"
	player.Score = 3

	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	// setup expectations
	playerService.On("FindById", 101).Return(player)

	playerController := PlayerController{}
	playerController.PlayerService = playerService

	// call the code we are testing
	req := httptest.NewRequest("GET", "http://localhost:8080/getPlayer/101", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/getPlayer/{id}", playerController.GetPlayer)

	r.ServeHTTP(w, req)

	expectedResult := viewmodels.PlayerVM{}
	expectedResult.Name = "Rafael"
	expectedResult.Score = 3

	actualResult := viewmodels.PlayerVM{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
}
