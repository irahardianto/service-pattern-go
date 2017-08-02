package main

import (
	"net/http"
	"service-pattern-go/controllers"
	"service-pattern-go/repositories"
	"service-pattern-go/services"

	"github.com/jinzhu/gorm"
)

func main() {

	dbHandlers, err := gorm.Open("sqlite3", "/var/tmp/gorm.db")
	if err != nil {
		//handle error
	}

	playerRepository := repositories.InitGormDB(dbHandlers)

	playerService := new(services.PlayerService)
	playerService.PlayerRepository = playerRepository

	playerController := controllers.PlayerController{}
	playerController.PlayerService = playerService

	defer dbHandlers.Close()

	http.HandleFunc("/getPlayer", func(res http.ResponseWriter, req *http.Request) {
		playerController.GetPlayer(res, req)
	})

	http.ListenAndServe(":8080", nil)
}
