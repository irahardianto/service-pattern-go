package repositories

import (
	"service-pattern-go/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type GormPlayerRepository struct {
	db *gorm.DB
}

func InitGormDB(conn *gorm.DB) *GormPlayerRepository {

	dbHandler := new(GormPlayerRepository)
	dbHandler.db = conn

	return dbHandler
}

func (conn *GormPlayerRepository) GetAllPlayers() []models.Player {
	//not implemented yet
	return nil
}

func (conn *GormPlayerRepository) GetPlayerById(id int) models.Player {

	player := models.Player{}
	conn.db.First(&player, id)

	return player
}

func (conn *GormPlayerRepository) CreatePlayer(player models.Player) (bool, error) {
	//not implemented yet
	return false, nil
}

func (conn *GormPlayerRepository) UpdatePlayer(id int, player models.Player) (bool, error) {
	//not implemented yet
	return false, nil
}

func (conn *GormPlayerRepository) DeletePlayer(id int) (bool, error) {
	//not implemented yet
	return false, nil
}
