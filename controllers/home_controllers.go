package controllers

import (
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.HTML(200, "homepage.html", gin.H{})
}
