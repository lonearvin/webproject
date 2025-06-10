package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"webproject/config"
)

func GetService(ctx *gin.Context) {
	// 这里是解析案例
	caseID := ctx.Query("id")
	//fmt.Println(caseID)
	var htmlFilePath string
	if caseID == "3C" {
		htmlFilePath = config.AppConfig.App.TemplatePath + "/ServicePages/ServicePages3C.html"
	}
	//fmt.Println(htmlFilePath)
	// 检查是否存在这个地址
	_, err := os.Open(htmlFilePath)
	if err != nil {
		//ctx.HTML(http.StatusNotFound, config.AppConfig.App.TemplatePath+"404.html", gin.H{})
		ctx.File(config.AppConfig.App.TemplatePath + "/404.html")
		fmt.Println(err)
		return
	}
	ctx.File(htmlFilePath)
}
