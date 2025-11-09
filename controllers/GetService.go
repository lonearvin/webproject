package controllers

import (
	"fmt"
	"os"
	"webproject/config"

	"github.com/gin-gonic/gin"
)

func GetService(ctx *gin.Context) {
	// 这里是解析案例
	caseID := ctx.Query("id")
	//fmt.Println(caseID)
	var htmlFilePath string
	if caseID == "3C" {
		htmlFilePath = config.AppConfig.App.TemplatePath + "/ServicePages/ServicePages3C.html"
	} else if caseID == "New_Energy_Services" {
		htmlFilePath = config.AppConfig.App.TemplatePath + "/ServicePages/New_Energy_Services.html"
	} else if caseID == "automotive_automation" {
		htmlFilePath = config.AppConfig.App.TemplatePath + "/ServicePages/automotive_automation.html"
	} else if caseID == "semiconductor_automation" {
		htmlFilePath = config.AppConfig.App.TemplatePath + "/ServicePages/semiconductor_automation.html"
	} else if caseID == "medical_equipment_automation" {
		htmlFilePath = config.AppConfig.App.TemplatePath + "/ServicePages/medical_equipment_automation.html"
	} else if caseID == "chemical_automation" {
		htmlFilePath = config.AppConfig.App.TemplatePath + "/ServicePages/chemical_automation.html"
	}
	// 检查是否存在这个地址
	_, err := os.Open(htmlFilePath)
	if err != nil {
		ctx.File(config.AppConfig.App.TemplatePath + "/404.html")
		fmt.Println(err)
		return
	}
	ctx.File(htmlFilePath)
}
