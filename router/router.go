package router

import (
	"github.com/gin-gonic/gin"
	"webproject/controllers"
	"webproject/middlewares"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	// 设置静态文件
	r.Static("/picture", "./picture")

	r.GET("/", controllers.Home)
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	api := r.Group("/api/")
	api.GET("/exchangeRate", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchangeRate", controllers.CreateExchangeRate)
	}
	return r
}
