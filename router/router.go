package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"webproject/config"
	"webproject/controllers"
)

func SetRouter() *gin.Engine {
	// 初始化 gin 引擎
	r := gin.Default()
	if gin.Mode() == gin.DebugMode {
		r.Use(func(c *gin.Context) {
			c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Writer.Header().Set("Pragma", "no-cache")
			c.Writer.Header().Set("Expires", "0")
			c.Next()
		})
	}
	// 注册静态资源路径
	staticPaths := map[string]string{
		"/static":  "./static",
		"/picture": "./picture",
		"/video":   "./video",
	}
	for route, path := range staticPaths {
		r.Static(route, path)
	}
	r.StaticFile("/favicon.ico", "./picture/favicon.ico")

	// 首页
	r.GET("/", controllers.Home)

	// 视频页
	r.GET("/api/video", controllers.GetProductVideo)

	// 服务页面（根据 id 显示对应 HTML）
	r.GET("/service", controllers.GetService)

	// 可选支持 REST 风格 URL: /service/3C
	r.GET("/service/:id", func(ctx *gin.Context) {
		ctx.Request.URL.RawQuery = "id=" + ctx.Param("id")
		controllers.GetService(ctx)
	})
	r.POST("/contact", controllers.ContactPost)
	r.POST("/subscribe", controllers.Subscribe)

	// 默认 404 页面处理
	r.NoRoute(func(ctx *gin.Context) {
		htmlFilePath := config.AppConfig.App.TemplatePath + "/404.html"
		if _, err := os.Stat(htmlFilePath); err == nil {
			ctx.File(htmlFilePath)
		} else {
			log.Printf("找不到404页面文件: %v", err)
			ctx.String(http.StatusNotFound, "404 Not Found")
		}
	})

	return r
}
