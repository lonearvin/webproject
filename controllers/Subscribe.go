package controllers

import (
	"net/http"
	"webproject/global"
	"webproject/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
)

func Subscribe(ctx *gin.Context) {
	var data utils.SubscribeData
	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		log.Errorf(ctx, "subscribe faild: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	var existing utils.SubscribeData
	if err := global.Db.Where("email=?", data.Email).First(&existing).Error; err == nil {
		ctx.JSON(409, gin.H{
			"code":    409,
			"message": "信息重复提交",
		})
		return
	}

	// 进行数据库保存
	if err := global.Db.Create(&data).Error; err != nil {
		log.Errorf(ctx, "subscribe faild: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		log.Errorf(ctx, "subscribe faild: %v", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	})
}
