package main

import (
	"Kasir_Test/config"
	"Kasir_Test/manager"
	"Kasir_Test/util"
	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type appServer struct {
	r              *gin.Engine
	c              config.Config
	InfraManager   manager.InfraManager
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

func (a *appServer) initHandler() {
	a.v1()
}

func (a *appServer) v1() {

}

func (a *appServer) Run() {
	a.initHandler()
	err := a.r.Run(a.c.ApiConfig.Url)
	if err != nil {
		util.Log.Fatal().Msg("Server Failed to run")
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.NewConfig(".", "config")

	return &appServer{
		r: r,
		c: c,
	}
}
