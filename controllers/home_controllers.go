package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"webproject/config"
)

// Home 该函数将主页面返回给请求
func Home(ctx *gin.Context) {
	htmlFilePath := config.AppConfig.App.TemplatePath + "/homepage.html"
	// 检查是否存在这个地址
	_, err := os.Open(htmlFilePath)
	if err != nil {
		log.Printf("Failed to find Html file: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find HTML file"})
		return
	}
	ctx.File(htmlFilePath)
}
