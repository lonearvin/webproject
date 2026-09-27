package controllers

import (
	"log"
	"net/http"
	"webproject/global"
	"webproject/utils"

	"github.com/gin-gonic/gin"
)

func Subscribe(ctx *gin.Context) {
	var data utils.SubscribeData
	if err := ctx.ShouldBind(&data); err != nil {
		log.Printf("subscribe bind failed: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	var existing utils.SubscribeData
	if err := global.Db.Where("email=?", data.Email).First(&existing).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"code":    http.StatusConflict,
			"message": "信息重复提交",
		})
		return
	}

	// 进行数据库保存
	if err := global.Db.Create(&data).Error; err != nil {
		log.Printf("subscribe save failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	})
}
