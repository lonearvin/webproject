package router

import (
	"github.com/gin-gonic/gin"
	"webproject/controllers"
	"webproject/middlewares"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	// 增加静态文件
	r.Static("./picture", "./picture/")
	r.Static("/template", "./template")
	r.Static("/video", "./video")
	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.GET("/", controllers.Home)
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	api := r.Group("/api/")
	api.GET("/exchangeRate", controllers.GetExchangeRates)
	// 需要登陆和注册才可以使用的功能
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchangeRate", controllers.CreateExchangeRate)
	}
	return r
}
