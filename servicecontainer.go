package main

import (
	"sync"

	"github.com/irahardianto/service-pattern-go/controllers"
	"github.com/irahardianto/service-pattern-go/infrastructures"
	"github.com/irahardianto/service-pattern-go/repositories"
	"github.com/irahardianto/service-pattern-go/services"
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
