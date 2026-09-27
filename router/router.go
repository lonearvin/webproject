package router

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"webproject/config"
	"webproject/controllers"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	// gzip 压缩（排除本身已压缩的静态资源格式）
	r.Use(gzip.Gzip(gzip.DefaultCompression,
		gzip.WithExcludedExtensions([]string{".png", ".jpg", ".jpeg", ".webp", ".gif", ".mp4", ".webm", ".pdf", ".woff2", ".ico"}),
	))

	r.SetFuncMap(template.FuncMap{})
	r.LoadHTMLFiles(
		config.AppConfig.App.TemplatePath+"/homepage.html",
		config.AppConfig.App.TemplatePath+"/cases.html",
		config.AppConfig.App.TemplatePath+"/404.html",
		config.AppConfig.App.TemplatePath+"/components/header.html",
		config.AppConfig.App.TemplatePath+"/components/footer.html",
		config.AppConfig.App.TemplatePath+"/components/contact-form.html",
		config.AppConfig.App.TemplatePath+"/ServicePages/ServicePages3C.html",
		config.AppConfig.App.TemplatePath+"/ServicePages/New_Energy_Services.html",
		config.AppConfig.App.TemplatePath+"/ServicePages/automotive_automation.html",
		config.AppConfig.App.TemplatePath+"/ServicePages/semiconductor_automation.html",
		config.AppConfig.App.TemplatePath+"/ServicePages/medical_equipment_automation.html",
		config.AppConfig.App.TemplatePath+"/ServicePages/chemical_automation.html",
	)

	// 静态资源缓存控制中间件
	r.Use(func(c *gin.Context) {
		path := c.Request.URL.Path
		if len(path) > 8 && path[:8] == "/static/" {
			c.Writer.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		} else if len(path) > 9 && path[:9] == "/picture/" {
			c.Writer.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		} else if len(path) > 7 && path[:7] == "/video/" {
			c.Writer.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		} else {
			if gin.Mode() == gin.DebugMode {
				c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				c.Writer.Header().Set("Pragma", "no-cache")
				c.Writer.Header().Set("Expires", "0")
			}
		}
		c.Next()
	})
	// 注册静态资源路径
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/picture/logo.webp")

	// 首页
	r.GET("/", controllers.Home)

	// 案例页面
	r.GET("/cases", controllers.GetCases)

	// 服务页面（根据 id 显示对应 HTML）
	r.GET("/service", controllers.GetService)

	// 可选支持 REST 风格 URL: /service/3C
	r.GET("/service/:id", func(ctx *gin.Context) {
		ctx.Request.URL.RawQuery = "id=" + ctx.Param("id")
		controllers.GetService(ctx)
	})
	r.POST("/contact", controllers.ContactPost)
	r.POST("/subscribe", controllers.Subscribe)

	// 默认 404 页面处理（以 404 状态码返回，而不是 200）
	r.NoRoute(func(ctx *gin.Context) {
		htmlFilePath := config.AppConfig.App.TemplatePath + "/404.html"
		if body, err := os.ReadFile(htmlFilePath); err == nil {
			ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", body)
		} else {
			log.Printf("找不到404页面文件: %v", err)
			ctx.String(http.StatusNotFound, "404 Not Found")
		}
	})

	return r
}
