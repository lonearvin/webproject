package router

import (
	"github.com/gin-gonic/gin"
	"webproject/controllers"
	"webproject/middlewares"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	// 注册静态资源路径
	staticPaths := map[string]string{
		"/picture":  "./picture",
		"/template": "./template",
		"/video":    "./video",
	}
	for route, path := range staticPaths {
		r.Static(route, path)
	}
	r.StaticFile("/favicon.ico", "./favicon.ico")

	// 首页
	r.GET("/", controllers.Home)

	// 用户认证相关路由
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	// API 路由（包含登录认证中间件）
	api := r.Group("/api")
	api.GET("/exchangeRate", controllers.GetExchangeRates)

	// 需要认证的路由
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchangeRate", controllers.CreateExchangeRate)
	}

	return r
}
