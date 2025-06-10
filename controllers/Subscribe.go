package controllers

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
	"net/http"
	"webproject/global"
)
import "webproject/utils"

func Subscribe(ctx *gin.Context) {
	var data utils.SubscribeData
	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		log.Errorf(ctx, "subscribe faild: %v", err.Error())
		return
	}

	if err := global.GlobalDB.AutoMigrate(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		log.Errorf(ctx, "subscribe faild: %v", err.Error())
		return
	}

	// 进行数据库保存
	if err := global.GlobalDB.Create(&data).Error; err != nil {
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
