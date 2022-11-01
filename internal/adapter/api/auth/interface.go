package auth

import "github.com/gin-gonic/gin"

type AuthorizationHandler interface {
	InitRoutes(ctx *gin.Engine)
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}
