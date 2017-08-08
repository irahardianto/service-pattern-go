package main

import (
	"sync"

	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type muxRouter struct{}

func (router *muxRouter) InitRouter() *chi.Mux {

	playerController := ServiceContainer().InjectPlayerController()

	r := chi.NewRouter()
	r.HandleFunc("/getPlayer/{id}", playerController.GetPlayer)

	return r
}

var (
	m          *muxRouter
	routerOnce sync.Once
)

func ChiMuxRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &muxRouter{}
		})
	}
	return m
}
