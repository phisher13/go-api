package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewErrorHandler(err error, ctx *gin.Context) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return
}
