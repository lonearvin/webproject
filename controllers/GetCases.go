package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetCases(ctx *gin.Context) {
	ctx.HTML(200, "cases.html", gin.H{})
}
