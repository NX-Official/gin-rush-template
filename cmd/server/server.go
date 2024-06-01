package server

import (
	"gin-rush-template/config"
	"gin-rush-template/internal/global/database"
	"gin-rush-template/internal/global/middleware"
	"gin-rush-template/internal/global/otel"
	"gin-rush-template/internal/module"
	"gin-rush-template/tools"
	"github.com/gin-gonic/gin"
	"log"
)

const configPath = "config.yaml"

func Init() {
	config.Read(configPath)
	database.Init()

	if config.Get().OTel.Enable {
		otel.Init()
	}

	for _, m := range module.Modules {
		log.Println("Init Module: " + m.GetName())
		m.Init()
	}
}

func Run() {
	r := gin.New()
	gin.SetMode(string(config.Get().Mode))
	r.Use(gin.Logger(), middleware.Recovery())

	if config.Get().OTel.Enable {
		r.Use(middleware.Trace())
	}

	for _, m := range module.Modules {
		log.Println("InitRouter: " + m.GetName())
		m.InitRouter(r.Group("/" + config.Get().Prefix))
	}
	err := r.Run(config.Get().Host + ":" + config.Get().Port)
	tools.PanicOnErr(err)
}
