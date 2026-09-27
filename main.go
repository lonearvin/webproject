package main

import (
	"log"
	"webproject/config"
	"webproject/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	config.InitConfig()
	model := config.AppConfig.GinMode.Model
	if model == "debug" {
		gin.SetMode(gin.DebugMode)
	} else if model == "test" {
		gin.SetMode(gin.TestMode)
	} else if model == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	if config.AppConfig.App.Port == "" {
		config.AppConfig.App.Port = "8080"
	}
	r := router.SetRouter()
	if err := r.Run(config.AppConfig.App.Port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
