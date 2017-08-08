package main

import (
	"sync"

	"github.com/gorilla/mux"
)

type IGorillaRouter interface {
	InitRouter() *mux.Router
}

type muxRouter struct{}

func (router *muxRouter) InitRouter() *mux.Router {

	playerController := ServiceContainer().InjectPlayerController()

	r := mux.NewRouter()
	r.HandleFunc("/getPlayer/{id}", playerController.GetPlayer)

	return r
}

var (
	m          *muxRouter
	routerOnce sync.Once
)

func GorillaMuxRouter() IGorillaRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &muxRouter{}
		})
	}
	return m
}
