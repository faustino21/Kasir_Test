package main

import (
	"Kasir_Test/Delivery/api"
	"Kasir_Test/Delivery/middleware"
	"Kasir_Test/config"
	"Kasir_Test/util"
	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type appServer struct {
	r *gin.Engine
	c config.Config
}

func (a *appServer) initHandler() {
	a.r.Use(middleware.AuthTokenMiddleware())
	a.v1()
}

func (a *appServer) v1() {
	cashierGroup := a.r.Group("/cashiers")
	productGroup := a.r.Group("products")
	api.CashierApiRoute(cashierGroup, a.c.UseCaseManager.CashierUseCase())
	api.LoginApiRoute(cashierGroup, a.c.UseCaseManager.LoginUseCase())
	api.ProductApiRoute(productGroup, a.c.UseCaseManager.ProductUseCase())
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
