package main

import (
	"fmt"
	"webproject/config"
)

func main() {
	config.InitConfig()
	fmt.Println(config.AppConfig.App.Port)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(config.AppConfig.App.Port) // listen and serve on 0.0.0.0:8080
}
