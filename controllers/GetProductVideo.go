package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func GetProductVideo(ctx *gin.Context) {
	// htmlFilePath := config.AppConfig.App.TemplatePath
	Path := "templates/producter.html"
	// 检查是否存在这个地址
	_, err := os.Open(Path)
	if err != nil {
		log.Printf("Failed to find Html file: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find HTML file"})
		return
	}
	ctx.File(Path)
}
