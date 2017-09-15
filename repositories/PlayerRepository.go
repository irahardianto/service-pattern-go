package repositories

import (
	"encoding/json"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/irahardianto/service-pattern-go/helpers"
	"github.com/irahardianto/service-pattern-go/infrastructures"
	"github.com/irahardianto/service-pattern-go/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PlayerRepository struct {
	Db          infrastructures.SqlConnection
	SafeAPICall helpers.SafeAPICall
}

func (repository *PlayerRepository) GetPlayerByName(name string) (models.PlayerModel, error) {

	output := make(chan models.PlayerModel, 1)
	hystrix.ConfigureCommand("get_player_by_name", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_player_by_name", func() error {

		conn := repository.Db.GetDB()

		player := models.PlayerModel{}
		conn.First(&player, "Name = ?", name)
		output <- player
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.PlayerModel{}, err
	}
}

func (repository *PlayerRepository) GetPlayerMessageFromAPI() models.MessageModel {

	//brokenEndpoint := "http://www.mocky.io/v2/599969590f0000880206f125"
	workingEndpoint := "http://www.mocky.io/v2/599967ae0f0000880206f11e"

	callAPI := repository.SafeAPICall.Get("hello_world", workingEndpoint, 1000)
	var data models.MessageModel
	json.Unmarshal(callAPI, &data)

	return data
}
