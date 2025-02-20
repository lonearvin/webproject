package router

import (
	"github.com/gin-gonic/gin"
	"webproject/controllers"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	return r
}
