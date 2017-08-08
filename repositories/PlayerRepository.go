package repositories

import (
	"service-pattern-go/infrastructures"
	"service-pattern-go/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PlayerRepository struct {
	Db infrastructures.SqlConnection
}

// func InitGormDB(conn *gorm.DB) *GormPlayerRepository {
//
// 	dbHandler := new(GormPlayerRepository)
// 	dbHandler.db = conn
//
// 	return dbHandler
// }

func (repository *PlayerRepository) GetAllPlayers() []models.Player {
	//not implemented yet
	return nil
}

func (repository *PlayerRepository) GetPlayerById(id int) models.Player {
	conn := repository.Db.GetDB()
	player := models.Player{}
	//conn.db.First(&player, id)
	conn.First(&player, id)

	return player
}

func (repository *PlayerRepository) CreatePlayer(player models.Player) (bool, error) {
	//not implemented yet
	return false, nil
}

func (repository *PlayerRepository) UpdatePlayer(id int, player models.Player) (bool, error) {
	//not implemented yet
	return false, nil
}

func (repository *PlayerRepository) DeletePlayer(id int) (bool, error) {
	//not implemented yet
	return false, nil
}
