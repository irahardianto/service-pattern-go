package main

import (
	"service-pattern-go/controllers"
	"service-pattern-go/infrastructures"
	"service-pattern-go/repositories"
	"service-pattern-go/services"
	"sync"
)

type IServiceContainer interface {
	InjectPlayerController() controllers.PlayerController
}

type kernel struct{}

func (k *kernel) InjectPlayerController() controllers.PlayerController {

	sqlconn := new(infrastructures.SqlConnection)
	sqlconn.InitDB()

	playerRepository := new(repositories.PlayerRepository)
	playerRepository.Db.Db = sqlconn.GetDB()

	playerService := new(services.PlayerService)
	playerService.PlayerRepository = playerRepository

	playerController := controllers.PlayerController{}
	playerController.PlayerService = playerService

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
