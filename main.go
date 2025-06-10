package main

import (
	"github.com/gin-gonic/gin"
	"webproject/config"
	"webproject/router"
)

func main() {

	// 日志颜色化
	gin.ForceConsoleColor()

	config.InitConfig()
	//fmt.Println(config.AppConfig.App.Port)
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
	err := r.Run(config.AppConfig.App.Port)
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
