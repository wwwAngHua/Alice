package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, HttpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(HttpStatus, gin.H{"code": code, "data": data, "message": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
