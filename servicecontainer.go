package main

import (
	"sync"

	"github.com/irahardianto/service-pattern-go/controllers"
	"github.com/irahardianto/service-pattern-go/repositories"
	"github.com/irahardianto/service-pattern-go/infrastructures"
	"github.com/irahardianto/service-pattern-go/services"
	"database/sql"
)

type IServiceContainer interface {
	InjectPlayerController() controllers.PlayerController
}

type kernel struct{}

func (k *kernel) InjectPlayerController() controllers.PlayerController {

	sqlConn, _ := sql.Open("sqlite3", "/var/tmp/tennis.db")
	sqliteHandler := &infrastructures.SQLiteHandler{}
	sqliteHandler.Conn = sqlConn

	playerRepository := &repositories.PlayerRepository{sqliteHandler}
	playerService := &services.PlayerService{&repositories.PlayerRepositoryWithCircuitBreaker{playerRepository}}
	playerController := controllers.PlayerController{playerService}

	return playerController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
