package controllers

import (
	"webproject/config"

	"github.com/gin-gonic/gin"
)

func GetCases(ctx *gin.Context) {
	htmlFilePath := config.AppConfig.App.TemplatePath + "/cases.html"
	ctx.File(htmlFilePath)
}
