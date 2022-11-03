package product

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	InitRoutes(router *gin.Engine)
	GetProducts(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
