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

	err = utils.HandleContactPost(data)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "received", "data": data})
}
