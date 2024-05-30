package server

import (
	"gin-rush-template/config"
	"gin-rush-template/internal/global/database"
	"gin-rush-template/internal/global/middleware"
	"gin-rush-template/internal/module"
	"github.com/gin-gonic/gin"
	"log"
)

const configPath = "config.yaml"

func Init() {
	config.Read(configPath)
	database.Init()

	for _, m := range module.Modules {
		log.Println("Init Module: " + m.GetName())
		m.Init()
	}
}

func Run() {
	r := gin.New()
	r.Use(gin.Logger(), middleware.Recovery())
	for _, m := range module.Modules {
		log.Println("InitRouter: " + m.GetName())
		m.InitRouter(r.Group("/" + config.Get().Prefix))
	}
	err := r.Run(config.Get().Host + ":" + config.Get().Port)
	if err != nil {
		panic(err)
	}
}
