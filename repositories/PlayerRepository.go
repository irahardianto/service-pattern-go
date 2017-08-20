package repositories

import (
	"encoding/json"
	"service-pattern-go/helpers"
	"service-pattern-go/infrastructures"
	"service-pattern-go/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PlayerRepository struct {
	Db          infrastructures.SqlConnection
	SafeAPICall helpers.SafeAPICall
}

// func InitGormDB(conn *gorm.DB) *GormPlayerRepository {
//
// 	dbHandler := new(GormPlayerRepository)
// 	dbHandler.db = conn
//
// 	return dbHandler
// }

func (repository *PlayerRepository) GetAllPlayers() []models.PlayerModel {
	//not implemented yet
	return nil
}

func (repository *PlayerRepository) GetPlayerById(id int) models.PlayerModel {
	conn := repository.Db.GetDB()
	player := models.PlayerModel{}
	//conn.db.First(&player, id)
	conn.First(&player, id)

	return player
}

func (repository *PlayerRepository) CreatePlayer(player models.PlayerModel) (bool, error) {
	//not implemented yet
	return false, nil
}

func (repository *PlayerRepository) UpdatePlayer(id int, player models.PlayerModel) (bool, error) {
	//not implemented yet
	return false, nil
}

func (repository *PlayerRepository) DeletePlayer(id int) (bool, error) {
	//not implemented yet
	return false, nil
}

func (repository *PlayerRepository) GetPlayerMessageFromAPI() models.MessageModel {

	//brokenEndpoint := "http://www.mocky.io/v2/599969590f0000880206f125"
	workingEndpoint := "http://www.mocky.io/v2/599967ae0f0000880206f11e"

	callAPI := repository.SafeAPICall.Get("hello_world", workingEndpoint, 1000)
	var data models.MessageModel
	json.Unmarshal(callAPI, &data)

	return data
}
