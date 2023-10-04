package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"message": msg,
		"statusCode": code,
	})
}

func sendSuccess(ctx *gin.Context, msg string, data interface{}) {
	status := http.StatusOK
	ctx.JSON(status, gin.H{
		"message": msg,
		"statusCode": status,
		"data": data,
	})
}