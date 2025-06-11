package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webproject/utils"
)

func ContactPost(context *gin.Context) {
	var data utils.ContactPostData
	err := context.ShouldBind(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 处理数据
	utils.HandleContactPost(context)
}
