package repositories

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/irahardianto/service-pattern-go/infrastructures"
	"github.com/irahardianto/service-pattern-go/interfaces"
	"github.com/irahardianto/service-pattern-go/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PlayerRepositoryWithCircuitBreaker struct {
	PlayerRepository interfaces.IPlayerRepository
}

func (repository *PlayerRepositoryWithCircuitBreaker) GetPlayerByName(name string) (models.PlayerModel, error) {

	output := make(chan models.PlayerModel, 1)
	hystrix.ConfigureCommand("get_player_by_name", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_player_by_name", func() error {

		player, _ := repository.PlayerRepository.GetPlayerByName(name)

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

type PlayerRepository struct {
	Db infrastructures.SqlConnection
}

func (repository *PlayerRepository) GetPlayerByName(name string) (models.PlayerModel, error) {

	conn := repository.Db.GetDB()

	player := models.PlayerModel{}
	conn.First(&player, "Name = ?", name)

	return player, nil
}
