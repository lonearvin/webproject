package main

import (
	"github.com/gin-gonic/gin"
	"webproject/config"
	"webproject/router"
)

type Header struct {
	MessageKey string
}

func main() {

	// 日志颜色化
	gin.ForceConsoleColor()

	config.InitConfig()
	//fmt.Println(config.AppConfig.App.Port
	if config.AppConfig.App.Port == "" {
		config.AppConfig.App.Port = "8080"
	}
	r := router.SetRouter()
	err := r.Run(config.AppConfig.App.Port)
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
