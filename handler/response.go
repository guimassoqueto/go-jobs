package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guimassoqueto/go-jobs/schemas"
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

// All below are the schemas for the swagger docs

type ErrorResponse struct {
	Message string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}

type ShowOpeningResponse struct {
	Message string `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}

type ListOpeningsResponse struct {
	Message string `json:"message"`
	Data []schemas.OpeningResponse `json:"data"`
}

type UpdateOpeningsResponse struct {
	Message string `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}