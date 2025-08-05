package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetProductVideo(ctx *gin.Context) {
	Path := "templates/producter.html"
	_, err := os.Open(Path)
	if err != nil {
		log.Printf("Failed to find Html file: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find HTML file"})
		return
	}
	ctx.File(Path)
}
