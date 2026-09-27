package controllers

import (
	"os"
	"webproject/config"

	"github.com/gin-gonic/gin"
)

// servicePages 维护合法的服务页面 ID 到模板文件的映射
var servicePages = map[string]string{
	"3C":                            "ServicePages3C.html",
	"New_Energy_Services":           "New_Energy_Services.html",
	"automotive_automation":         "automotive_automation.html",
	"semiconductor_automation":      "semiconductor_automation.html",
	"medical_equipment_automation":  "medical_equipment_automation.html",
	"chemical_automation":           "chemical_automation.html",
}

func GetService(ctx *gin.Context) {
	caseID := ctx.Query("id")
	filename, ok := servicePages[caseID]
	if !ok {
		ctx.HTML(404, "404.html", gin.H{})
		return
	}

	htmlFilePath := config.AppConfig.App.TemplatePath + "/ServicePages/" + filename
	// 用 Stat 检查文件存在，避免 Open 泄漏文件描述符
	if _, err := os.Stat(htmlFilePath); err != nil {
		ctx.HTML(404, "404.html", gin.H{})
		return
	}
	ctx.File(htmlFilePath)
}
